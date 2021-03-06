package pod

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"reflect"

	eventloggerv1 "github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1"
	c "github.com/bakito/k8s-event-logger-operator/pkg/constants"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var (
	log              = logf.Log.WithName("controller_pod")
	gracePeriod      int64
	eventLoggerImage = "quay.io/bakito/k8s-event-logger"
	podReqCPU        = resource.MustParse("100m")
	podReqMem        = resource.MustParse("64Mi")
	podMaxCPU        = resource.MustParse("200m")
	podMaxMem        = resource.MustParse("128Mi")
)

// Add creates a new EventLogger Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr.GetClient(), mgr.GetScheme()))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(client client.Client, scheme *runtime.Scheme) reconcile.Reconciler {

	if podImage, ok := os.LookupEnv(c.EnvEventLoggerImage); ok {
		eventLoggerImage = podImage
	}
	if cpu, ok := os.LookupEnv(c.EnvLoggerPodReqCPU); ok {
		podReqCPU = resource.MustParse(cpu)
	}
	if mem, ok := os.LookupEnv(c.EnvLoggerPodReqMem); ok {
		podReqMem = resource.MustParse(mem)
	}
	if cpu, ok := os.LookupEnv(c.EnvLoggerPodMaxCPU); ok {
		podMaxCPU = resource.MustParse(cpu)
	}
	if mem, ok := os.LookupEnv(c.EnvLoggerPodMaxMem); ok {
		podMaxMem = resource.MustParse(mem)
	}

	return &ReconcileEventLogger{client: client, scheme: scheme}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("pod-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource EventLogger
	err = c.Watch(&source.Kind{Type: &eventloggerv1.EventLogger{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pod and requeue the owner EventLogger
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &enqueueDeletedRequestForOwner{
		EnqueueRequestForOwner: handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &eventloggerv1.EventLogger{},
		},
	})
	if err != nil {
		return err
	}
	// Watch for changes to secondary resource Secret and requeue the owner EventLogger
	err = c.Watch(&source.Kind{Type: &corev1.Secret{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &eventloggerv1.EventLogger{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileEventLogger implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileEventLogger{}

// ReconcileEventLogger reconciles a EventLogger object
type ReconcileEventLogger struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a EventLogger object and makes changes based on the state read
// and what is in the EventLogger.Spec
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileEventLogger) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)

	// Fetch the EventLogger cr
	cr := &eventloggerv1.EventLogger{}
	err := r.client.Get(context.TODO(), request.NamespacedName, cr)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return r.updateCR(cr, reqLogger, err)
	}

	saccChanged, roleChanged, rbChanged, err := r.setupRbac(cr, reqLogger)
	if err != nil {
		return r.updateCR(cr, reqLogger, err)
	}

	// Define a new Pod object
	pod := podForCR(cr)
	// Check if this Pod already exists
	podChanged, err := r.createOrReplacePod(cr, pod, reqLogger)
	if err != nil {
		return r.updateCR(cr, reqLogger, err)
	}

	if saccChanged || roleChanged || rbChanged || podChanged {
		return r.updateCR(cr, reqLogger, nil)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileEventLogger) createOrReplacePod(cr *eventloggerv1.EventLogger, pod *corev1.Pod,
	reqLogger logr.Logger) (bool, error) {

	podList := &corev1.PodList{}
	opts := []client.ListOption{
		client.InNamespace(cr.Namespace),
		client.MatchingLabels(map[string]string{
			"app":        loggerName(cr),
			"created-by": "eventlogger"}),
	}
	err := r.client.List(context.TODO(), podList, opts...)

	if err != nil {
		return false, err
	}

	replacePod := false
	if len(podList.Items) == 1 {
		op := podList.Items[0]
		replacePod = podChanged(&op, pod)
	}

	if replacePod || len(podList.Items) > 1 {

		for _, p := range podList.Items {
			reqLogger.Info(fmt.Sprintf("Deleting %s", pod.Kind), "Namespace", pod.GetNamespace(), "Name", pod.GetName())
			err = r.client.Delete(context.TODO(), &p, &client.DeleteOptions{GracePeriodSeconds: &gracePeriod})
			if err != nil {
				return false, err
			}
		}
		podList = &corev1.PodList{}
	}

	if len(podList.Items) == 0 {
		// Set EventLogger cr as the owner and controller
		if err := controllerutil.SetControllerReference(cr, pod, r.scheme); err != nil {
			return false, err
		}
		reqLogger.Info(fmt.Sprintf("Creating a new %s", pod.Kind), "Namespace", pod.GetNamespace(), "Name", pod.GetName())
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

func (r *ReconcileEventLogger) setupRbac(cr *eventloggerv1.EventLogger, reqLogger logr.Logger) (bool, bool, bool, error) {
	var saccChanged, roleChanged, rbChanged bool
	var err error
	sacc, role, rb := rbacForCR(cr)

	if cr.Spec.ServiceAccount == "" {
		saccChanged, err = r.createOrReplace(cr, sacc, reqLogger, nil)
		if err != nil {
			return saccChanged, roleChanged, rbChanged, err
		}
		roleChanged, err = r.createOrReplace(cr, role, reqLogger, func(curr runtime.Object, next runtime.Object) updateReplace {
			o1 := curr.(*rbacv1.Role)
			o2 := next.(*rbacv1.Role)
			if reflect.DeepEqual(o1.Rules, o2.Rules) {
				return no
			}
			return update
		})
		if err != nil {
			return saccChanged, roleChanged, rbChanged, err
		}
		rbChanged, err = r.createOrReplace(cr, rb, reqLogger, nil)
		if err != nil {
			return saccChanged, roleChanged, rbChanged, err
		}
	} else {
		// Only delete sa if the name is different than the configured
		if cr.Spec.ServiceAccount != sacc.GetName() {
			err = r.saveDelete(sacc)
			if err != nil {
				return saccChanged, roleChanged, rbChanged, err
			}
		}
		err = r.saveDelete(role)
		if err != nil {
			return saccChanged, roleChanged, rbChanged, err
		}
		err = r.saveDelete(rb)
		if err != nil {
			return saccChanged, roleChanged, rbChanged, err
		}
	}
	return saccChanged, roleChanged, rbChanged, nil
}

func (r *ReconcileEventLogger) createOrReplace(cr *eventloggerv1.EventLogger,
	res runtime.Object,
	reqLogger logr.Logger,
	updateCheck func(curr runtime.Object, next runtime.Object) updateReplace) (bool, error) {
	query := res.DeepCopyObject()
	mo := res.(metav1.Object)
	// Check if this Resource already exists
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: mo.GetName(), Namespace: mo.GetNamespace()}, query)
	if err != nil && errors.IsNotFound(err) {
		// Set EventLogger cr as the owner and controller
		if err := controllerutil.SetControllerReference(cr, mo, r.scheme); err != nil {
			return false, err
		}

		reqLogger.Info(fmt.Sprintf("Creating a new %s", query.GetObjectKind().GroupVersionKind().Kind), "Namespace", mo.GetNamespace(), "Name", mo.GetName())
		err = r.client.Create(context.TODO(), res.(runtime.Object))
		if err != nil {
			return false, err
		}
		return true, nil

	} else if err != nil {
		return false, err
	}

	if updateCheck != nil {
		check := updateCheck(query, res)
		if check == update {
			reqLogger.Info(fmt.Sprintf("Updating %s", query.GetObjectKind().GroupVersionKind().Kind), "Namespace", mo.GetNamespace(), "Name", mo.GetName())
			err = r.client.Update(context.TODO(), res.(runtime.Object))

			if err != nil {
				return false, err
			}
			return true, nil
		} else if check == replace {
			reqLogger.Info(fmt.Sprintf("Replacing %s", query.GetObjectKind().GroupVersionKind().Kind), "Namespace", mo.GetNamespace(), "Name", mo.GetName())

			err = r.client.Delete(context.TODO(), query.(runtime.Object))

			if err != nil {
				return false, err
			}
			err = r.client.Create(context.TODO(), query.(runtime.Object))

			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	// Resource already exists
	return false, nil
}

func (r *ReconcileEventLogger) updateCR(cr *eventloggerv1.EventLogger, logger logr.Logger, err error) (reconcile.Result, error) {
	updErr := cr.UpdateStatus(logger, err, r.client)
	return reconcile.Result{}, updErr
}

func (r *ReconcileEventLogger) saveDelete(obj runtime.Object) error {
	err := r.client.Delete(context.TODO(), obj)
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
	}
	return nil
}

// podForCR returns a pod with the same name/namespace as the cr
func podForCR(cr *eventloggerv1.EventLogger) *corev1.Pod {
	labels := make(map[string]string)
	for k, v := range cr.Spec.Labels {
		labels[k] = v
	}
	labels["app"] = loggerName(cr)
	labels["created-by"] = "eventlogger"

	annotations := make(map[string]string)
	for k, v := range cr.Spec.Annotations {
		annotations[k] = v
	}
	if cr.Spec.ScrapeMetrics != nil && *cr.Spec.ScrapeMetrics {
		annotations["prometheus.io/port"] = string(c.MetricsPort)
		annotations["prometheus.io/scrape"] = "true"
	}

	watchNamespace := cr.GetNamespace()
	if cr.Spec.Namespace != nil {
		watchNamespace = *cr.Spec.Namespace
	}

	saccName := loggerName(cr)
	if cr.Spec.ServiceAccount != "" {
		saccName = cr.Spec.ServiceAccount
	}

	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind: "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        loggerName(cr) + "-" + randString(),
			Namespace:   cr.Namespace,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:            "event-logger",
					Image:           eventLoggerImage,
					ImagePullPolicy: corev1.PullAlways,
					Command:         []string{"/opt/go/k8s-event-logger"},
					Args:            os.Args[1:], // pass on the operator args
					Env: []corev1.EnvVar{
						{
							Name:  "WATCH_NAMESPACE",
							Value: watchNamespace,
						},
						{
							Name: "POD_NAME",
							ValueFrom: &corev1.EnvVarSource{
								FieldRef: &corev1.ObjectFieldSelector{
									FieldPath: "metadata.name",
								},
							},
						},
						{
							Name:  c.EnvConfigName,
							Value: cr.Name,
						},
						{
							Name:  "DEBUG_CONFIG",
							Value: "false",
						},
					},
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    podReqCPU,
							corev1.ResourceMemory: podReqMem,
						},
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    podMaxCPU,
							corev1.ResourceMemory: podMaxMem,
						},
					},
				},
			},
			ServiceAccountName: saccName,
		},
	}
	return pod
}

func rbacForCR(cr *eventloggerv1.EventLogger) (*corev1.ServiceAccount, *rbacv1.Role, *rbacv1.RoleBinding) {
	sacc := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind: "ServiceAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      loggerName(cr),
			Namespace: cr.Namespace,
			Labels: map[string]string{
				"app": loggerName(cr),
			},
		},
	}

	role := &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{
			Kind: "Role",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      loggerName(cr),
			Namespace: cr.Namespace,
			Labels: map[string]string{
				"app": loggerName(cr),
			},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"events", "pods"},
				Verbs:     []string{"watch", "get", "list"},
			},
			{
				APIGroups: []string{"eventlogger.bakito.ch"},
				Resources: []string{"eventloggers"},
				Verbs:     []string{"get", "list", "patch", "update", "watch"},
			},
		},
	}
	rb := &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind: "RoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      loggerName(cr),
			Namespace: cr.Namespace,
			Labels: map[string]string{
				"app": loggerName(cr),
			},
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      loggerName(cr),
				Namespace: cr.Namespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "Role",
			APIGroup: "rbac.authorization.k8s.io",
			Name:     loggerName(cr),
		},
	}

	return sacc, role, rb
}

type updateReplace string

const (
	update  updateReplace = "update"
	replace updateReplace = "replace"
	no      updateReplace = "no"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func randString() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type enqueueDeletedRequestForOwner struct {
	handler.EnqueueRequestForOwner
}

// Create implements Predicate
func (h enqueueDeletedRequestForOwner) Create(event.CreateEvent, workqueue.RateLimitingInterface) {
}

// Delete implements Predicate
func (h enqueueDeletedRequestForOwner) Delete(e event.DeleteEvent, rli workqueue.RateLimitingInterface) {
	h.EnqueueRequestForOwner.Delete(e, rli)
}

// Update implements Predicate
func (h enqueueDeletedRequestForOwner) Update(event.UpdateEvent, workqueue.RateLimitingInterface) {
}

// Generic implements Predicate
func (h enqueueDeletedRequestForOwner) Generic(event.GenericEvent, workqueue.RateLimitingInterface) {
}

func loggerName(cr *eventloggerv1.EventLogger) string {
	return "event-logger-" + cr.Name
}

func podChanged(old, new *corev1.Pod) bool {
	if old.Spec.ServiceAccountName != new.Spec.ServiceAccountName {
		return true
	}
	if len(old.Spec.Containers) > 0 && len(new.Spec.Containers) > 0 && old.Spec.Containers[0].Image != new.Spec.Containers[0].Image {
		return true
	}

	return podEnv(old, "WATCH_NAMESPACE") != podEnv(new, "WATCH_NAMESPACE")
}

func podEnv(pod *corev1.Pod, name string) string {

	for _, env := range pod.Spec.Containers[0].Env {
		if env.Name == name {
			return env.Value
		}
	}
	return "N/A"
}
