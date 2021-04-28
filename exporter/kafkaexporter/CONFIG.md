## kafkaexporter-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| timeout |time.Duration| 5s | Timeout is the timeout for every attempt to send data to the backend.  |
| sending_queue |[exporterhelper-QueueSettings](#exporterhelper-QueueSettings)| <no value> |  |
| retry_on_failure |[exporterhelper-RetrySettings](#exporterhelper-RetrySettings)| <no value> |  |
| brokers |[]string| <no value> | The list of kafka brokers (default localhost:9092)  |
| protocol_version |string| <no value> | Kafka protocol version  |
| topic |string| <no value> | The name of the kafka topic to export to (default otlp_spans for traces, otlp_metrics for metrics)  |
| encoding |string| otlp_proto | Encoding of messages (default "otlp_proto")  |
| metadata |[kafkaexporter-Metadata](#kafkaexporter-Metadata)| <no value> | Metadata is the namespace for metadata management properties used by the Client, and shared by the Producer/Consumer.  |
| auth |[kafkaexporter-Authentication](#kafkaexporter-Authentication)| <no value> | Authentication defines used authentication mechanism.  |

## exporterhelper-QueueSettings

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| enabled |bool| true | Enabled indicates whether to not enqueue batches before sending to the consumerSender.  |
| num_consumers |int| 10 | NumConsumers is the number of consumers from the queue.  |
| queue_size |int| 5000 | QueueSize is the maximum number of batches allowed in queue at a given time.  |

## exporterhelper-RetrySettings

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| enabled |bool| true | Enabled indicates whether to not retry sending batches in case of export failure.  |
| initial_interval |time.Duration| 5s | InitialInterval the time to wait after the first failure before retrying.  |
| max_interval |time.Duration| 30s | MaxInterval is the upper bound on backoff interval. Once this value is reached the delay between consecutive retries will always be `MaxInterval`.  |
| max_elapsed_time |time.Duration| 5m0s | MaxElapsedTime is the maximum amount of time (including retries) spent trying to send a request/batch. Once this value is reached, the data is discarded.  |

## kafkaexporter-Metadata

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| full |bool| true | Whether to maintain a full set of metadata for all topics, or just the minimal set that has been necessary so far. The full set is simpler and usually more convenient, but can take up a substantial amount of memory if you have many topics and partitions. Defaults to true.  |
| retry |[kafkaexporter-MetadataRetry](#kafkaexporter-MetadataRetry)| <no value> | Retry configuration for metadata. This configuration is useful to avoid race conditions when broker is starting at the same time as collector.  |

## kafkaexporter-MetadataRetry

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| max |int| 3 | The total number of times to retry a metadata request when the cluster is in the middle of a leader election or at startup (default 3).  |
| backoff |time.Duration| 250ms | How long to wait for leader election to occur before retrying (default 250ms). Similar to the JVM's `retry.backoff.ms`.  |

## kafkaexporter-Authentication

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| plain_text |[kafkaexporter-PlainTextConfig](#kafkaexporter-PlainTextConfig)| <no value> |  |
| sasl |[kafkaexporter-SASLConfig](#kafkaexporter-SASLConfig)| <no value> |  |
| tls |[configtls-TLSClientSetting](#configtls-TLSClientSetting)| <no value> |  |
| kerberos |[kafkaexporter-KerberosConfig](#kafkaexporter-KerberosConfig)| <no value> |  |

## kafkaexporter-PlainTextConfig

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| username |string| <no value> |  |
| password |string| <no value> |  |

## kafkaexporter-SASLConfig

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| username |string| <no value> | Username to be used on authentication  |
| password |string| <no value> | Password to be used on authentication  |
| mechanism |string| <no value> | SASL Mechanism to be used, possible values are: (PLAIN, SCRAM-SHA-256 or SCRAM-SHA-512).  |

## configtls-TLSClientSetting

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| ca_file |string| <no value> | Path to the CA cert. For a client this verifies the server certificate. For a server this verifies client certificates. If empty uses system root CA. (optional)  |
| cert_file |string| <no value> | Path to the TLS cert to use for TLS required connections. (optional)  |
| key_file |string| <no value> | Path to the TLS key to use for TLS required connections. (optional)  |
| insecure |bool| <no value> | In gRPC when set to true, this is used to disable the client transport security. See https://godoc.org/google.golang.org/grpc#WithInsecure. In HTTP, this disables verifying the server's certificate chain and host name (InsecureSkipVerify in the tls Config). Please refer to https://godoc.org/crypto/tls#Config for more information. (optional, default false)  |
| insecure_skip_verify |bool| <no value> | InsecureSkipVerify will enable TLS but not verify the certificate.  |
| server_name_override |string| <no value> | ServerName requested by client for virtual hosting. This sets the ServerName in the TLSConfig. Please refer to https://godoc.org/crypto/tls#Config for more information. (optional)  |

## kafkaexporter-KerberosConfig

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| service_name |string| <no value> |  |
| realm |string| <no value> |  |
| use_keytab |bool| <no value> |  |
| username |string| <no value> |  |
| password |string| <no value> |  |
| config_file |string| <no value> |  |
| keytab_file |string| <no value> |  |

