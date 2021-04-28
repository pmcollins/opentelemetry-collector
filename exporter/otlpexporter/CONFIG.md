## otlpexporter-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| timeout |time.Duration| 5s | Timeout is the timeout for every attempt to send data to the backend.  |
| sending_queue |[exporterhelper-QueueSettings](#exporterhelper-QueueSettings)| <no value> |  |
| retry_on_failure |[exporterhelper-RetrySettings](#exporterhelper-RetrySettings)| <no value> |  |
| endpoint |string| <no value> | The target to which the exporter is going to send traces or metrics, using the gRPC protocol. The valid syntax is described at https://github.com/grpc/grpc/blob/master/doc/naming.md.  |
| compression |string| <no value> | The compression key for supported compression types within collector. Currently the only supported mode is `gzip`.  |
| ca_file |string| <no value> | Path to the CA cert. For a client this verifies the server certificate. For a server this verifies client certificates. If empty uses system root CA. (optional)  |
| cert_file |string| <no value> | Path to the TLS cert to use for TLS required connections. (optional)  |
| key_file |string| <no value> | Path to the TLS key to use for TLS required connections. (optional)  |
| insecure |bool| <no value> | In gRPC when set to true, this is used to disable the client transport security. See https://godoc.org/google.golang.org/grpc#WithInsecure. In HTTP, this disables verifying the server's certificate chain and host name (InsecureSkipVerify in the tls Config). Please refer to https://godoc.org/crypto/tls#Config for more information. (optional, default false)  |
| insecure_skip_verify |bool| <no value> | InsecureSkipVerify will enable TLS but not verify the certificate.  |
| server_name_override |string| <no value> | ServerName requested by client for virtual hosting. This sets the ServerName in the TLSConfig. Please refer to https://godoc.org/crypto/tls#Config for more information. (optional)  |
| keepalive |[configgrpc-KeepaliveClientConfig](#configgrpc-KeepaliveClientConfig)| <no value> | The keepalive parameters for gRPC client. See grpc.WithKeepaliveParams (https://godoc.org/google.golang.org/grpc#WithKeepaliveParams).  |
| read_buffer_size |int| <no value> | ReadBufferSize for gRPC client. See grpc.WithReadBufferSize (https://godoc.org/google.golang.org/grpc#WithReadBufferSize).  |
| write_buffer_size |int| 524288 | WriteBufferSize for gRPC gRPC. See grpc.WithWriteBufferSize (https://godoc.org/google.golang.org/grpc#WithWriteBufferSize).  |
| wait_for_ready |bool| <no value> | WaitForReady parameter configures client to wait for ready state before sending data. (https://github.com/grpc/grpc/blob/master/doc/wait-for-ready.md)  |
| headers |map[string]string| <no value> | The headers associated with gRPC requests.  |
| per_rpc_auth |[configgrpc-PerRPCAuthConfig](#configgrpc-PerRPCAuthConfig)| <no value> | PerRPCAuth parameter configures the client to send authentication data on a per-RPC basis.  |
| balancer_name |string| <no value> | Sets the balancer in grpclb_policy to discover the servers. Default is pick_first https://github.com/grpc/grpc-go/blob/master/examples/features/load_balancing/README.md  |

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

## configgrpc-KeepaliveClientConfig

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| time |time.Duration| <no value> |  |
| timeout |time.Duration| <no value> |  |
| permit_without_stream |bool| <no value> |  |

## configgrpc-PerRPCAuthConfig

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| type |string| <no value> | AuthType represents the authentication type to use. Currently, only 'bearer' is supported.  |
| bearer_token |string| <no value> | BearerToken specifies the bearer token to use for every RPC.  |

