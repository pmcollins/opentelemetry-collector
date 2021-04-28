## kafkareceiver-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| brokers |[]string| <no value> | The list of kafka brokers (default localhost:9092)  |
| protocol_version |string| <no value> | Kafka protocol version  |
| topic |string| otlp_spans | The name of the kafka topic to consume from (default "otlp_spans")  |
| encoding |string| otlp_proto | Encoding of the messages (default "otlp_proto")  |
| group_id |string| otel-collector | The consumer group that receiver will be consuming messages from (default "otel-collector")  |
| client_id |string| otel-collector | The consumer client ID that receiver will use (default "otel-collector")  |
| metadata |[kafkaexporter-Metadata](#kafkaexporter-Metadata)| <no value> | Metadata is the namespace for metadata management properties used by the Client, and shared by the Producer/Consumer.  |
| auth |[kafkaexporter-Authentication](#kafkaexporter-Authentication)| <no value> |  |

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

