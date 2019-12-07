/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var (
	// log is for logging in this package.
	eventloggerlog = logf.Log.WithName("eventlogger-resource")
	cl             client.Client
)

// SetupWebhookWithManager setup webhook
func (r *EventLogger) SetupWebhookWithManager(mgr ctrl.Manager) error {
	cl = mgr.GetClient()
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-eventlogger-bakito-ch-v1-eventlogger,mutating=true,failurePolicy=fail,groups=webapp.my.domain,resources=guestbooks,verbs=create;update,versions=v1,name=mguestbook.kb.io
var _ webhook.Defaulter = &EventLogger{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *EventLogger) Default() {
	eventloggerlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-eventlogger-bakito-ch-v1-eventlogger,mutating=false,failurePolicy=fail,groups=webapp.my.domain,resources=eventLoggers,versions=v1,name=veventLogger.kb.io
var _ webhook.Validator = &EventLogger{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *EventLogger) ValidateCreate() error {
	eventloggerlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *EventLogger) ValidateUpdate(old runtime.Object) error {
	eventloggerlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *EventLogger) ValidateDelete() error {
	eventloggerlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
