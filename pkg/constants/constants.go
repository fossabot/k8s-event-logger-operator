package constants

const (
	// EnvConfigFilePath env variable name for config path
	EnvConfigFilePath = "CONFIG_PATH"

	// EnvEventLoggerImage env variable name for the image if the event logger
	EnvEventLoggerImage = "EVENT_LOGGER_IMAGE"

	// EnvLoggerPodReqCPU set the logger pod request cpu
	EnvLoggerPodReqCPU = "LOGGER_POD_REQUEST_CPU"

	// EnvLoggerPodReqMem set the logger pod request memory
	EnvLoggerPodReqMem = "LOGGER_POD_REQUEST_MEM"

	// EnvLoggerPodMaxCPU set the logger pod max cpu
	EnvLoggerPodMaxCPU = "LOGGER_POD_LIMIT_CPU"

	// EnvLoggerPodMaxMem set the logger pod mac memory
	EnvLoggerPodMaxMem = "LOGGER_POD_LIMIT_MEM"

	// MetricsHost host for the metrics
	MetricsHost = "0.0.0.0"

	// WebhookHost host for the webhook
	WebhookHost = MetricsHost

	// MetricsPort port for the metrics
	MetricsPort int32 = 8383

	// OperatorMetricsPort port for the operator metrics
	OperatorMetricsPort int32 = 8686

	// WebhookPort port for the webhook
	WebhookPort int = 8443
)
