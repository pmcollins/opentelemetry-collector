## prometheusexporter-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| endpoint |string| <no value> | The address on which the Prometheus scrape handler will be run on.  |
| namespace |string| <no value> | Namespace if set, exports metrics under the provided value.  |
| const_labels |prometheus.Labels| <no value> | ConstLabels are values that are applied for every exported metric.  |
| send_timestamps |bool| <no value> | SendTimestamps will send the underlying scrape timestamp with the export  |
| metric_expiration |time.Duration| 5m0s | MetricExpiration defines how long metrics are kept without updates  |
| resource_to_telemetry_conversion |[exporterhelper-ResourceToTelemetrySettings](#exporterhelper-ResourceToTelemetrySettings)| <no value> | ResourceToTelemetrySettings defines configuration for converting resource attributes to metric labels.  |

## exporterhelper-ResourceToTelemetrySettings

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| enabled |bool| <no value> | Enabled indicates whether to not convert resource attributes to metric labels  |

