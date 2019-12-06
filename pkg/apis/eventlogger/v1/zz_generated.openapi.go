// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.EventLogger":       schema_pkg_apis_eventlogger_v1_EventLogger(ref),
		"github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.EventLoggerConf":   schema_pkg_apis_eventlogger_v1_EventLoggerConf(ref),
		"github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.EventLoggerSpec":   schema_pkg_apis_eventlogger_v1_EventLoggerSpec(ref),
		"github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.EventLoggerStatus": schema_pkg_apis_eventlogger_v1_EventLoggerStatus(ref),
		"github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.Kind":              schema_pkg_apis_eventlogger_v1_Kind(ref),
	}
}

func schema_pkg_apis_eventlogger_v1_EventLogger(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventLogger is the Schema for the eventloggers API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.EventLoggerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.EventLoggerStatus"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.EventLoggerSpec", "github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.EventLoggerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_eventlogger_v1_EventLoggerConf(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventLoggerConf defines the configuration of EventLogger",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kinds": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Kinds the kinds to logg the events for",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.Kind"),
									},
								},
							},
						},
					},
					"eventTypes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "EventTypes the event types to log. If empty all events are logged.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.Kind"},
	}
}

func schema_pkg_apis_eventlogger_v1_EventLoggerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventLoggerSpec defines the desired state of EventLogger",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kinds": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Kinds the kinds to logg the events for",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.Kind"),
									},
								},
							},
						},
					},
					"eventTypes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "EventTypes the event types to log. If empty all events are logged.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Labels additional labels for the logger pod",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Description: "Labels additional annotations for the logger pod",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"scrapeMetrics": {
						SchemaProps: spec.SchemaProps{
							Description: "ScrapeMetrics if true, prometheus scrape annotations are added to the pod",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace the namespace to watch on, may be an empty string",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serviceAccount": {
						SchemaProps: spec.SchemaProps{
							Description: "ServiceAccount the service account to use for the logger pod",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/bakito/k8s-event-logger-operator/pkg/apis/eventlogger/v1.Kind"},
	}
}

func schema_pkg_apis_eventlogger_v1_EventLoggerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventLoggerStatus defines the observed state of EventLogger",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"operatorVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "OperatorVersion the version of the operator that processed the cr",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastProcessed": {
						SchemaProps: spec.SchemaProps{
							Description: "LastProcessed the timestamp the cr was last processed",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"error": {
						SchemaProps: spec.SchemaProps{
							Description: "Error",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"operatorVersion", "lastProcessed"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_eventlogger_v1_Kind(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Kind defines a kind to loge events for",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"eventTypes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "EventTypes the event types to log. If empty events are logged as defined in spec.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"matchingPatterns": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "MatchingPatterns optional regex pattern that must be contained in the message to be logged",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"skipOnMatch": {
						SchemaProps: spec.SchemaProps{
							Description: "SkipOnMatch skip the entry if matched",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}
