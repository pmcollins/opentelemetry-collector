## prometheusremotewriteexporter-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| timeout |time.Duration| 5s | Timeout is the timeout for every attempt to send data to the backend.  |
| sending_queue |[exporterhelper-QueueSettings](#exporterhelper-QueueSettings)| <no value> |  |
| retry_on_failure |[exporterhelper-RetrySettings](#exporterhelper-RetrySettings)| <no value> |  |
| namespace |string| <no value> | prefix attached to each exported metric name See: https://prometheus.io/docs/practices/naming/#metric-names  |
| external_labels |map[string]string| <no value> | ExternalLabels defines a map of label keys and values that are allowed to start with reserved prefix "__"  |
| endpoint |string| http://some.url:9411/api/prom/push | The target URL to send data to (e.g.: http://some.url:9411/v1/traces).  |
| ca_file |string| <no value> | Path to the CA cert. For a client this verifies the server certificate. For a server this verifies client certificates. If empty uses system root CA. (optional)  |
| cert_file |string| <no value> | Path to the TLS cert to use for TLS required connections. (optional)  |
| key_file |string| <no value> | Path to the TLS key to use for TLS required connections. (optional)  |
| insecure |bool| <no value> | In gRPC when set to true, this is used to disable the client transport security. See https://godoc.org/google.golang.org/grpc#WithInsecure. In HTTP, this disables verifying the server's certificate chain and host name (InsecureSkipVerify in the tls Config). Please refer to https://godoc.org/crypto/tls#Config for more information. (optional, default false)  |
| insecure_skip_verify |bool| <no value> | InsecureSkipVerify will enable TLS but not verify the certificate.  |
| server_name_override |string| <no value> | ServerName requested by client for virtual hosting. This sets the ServerName in the TLSConfig. Please refer to https://godoc.org/crypto/tls#Config for more information. (optional)  |
| read_buffer_size |int| <no value> | ReadBufferSize for HTTP client. See http.Transport.ReadBufferSize.  |
| write_buffer_size |int| 524288 | WriteBufferSize for HTTP client. See http.Transport.WriteBufferSize.  |
| timeout |time.Duration| 5s | Timeout parameter configures `http.Client.Timeout`.  |
| headers |map[string]string| <no value> | Additional headers attached to each HTTP request sent by the client. Existing header values are overwritten if collision happens.  |
| customroundtripper |func(http.RoundTripper) (http.RoundTripper, error)| <no value> | Custom Round Tripper to allow for individual components to intercept HTTP requests  |

## exporterhelper-QueueSettings

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| enabled |bool| <no value> | Enabled indicates whether to not enqueue batches before sending to the consumerSender.  |
| num_consumers |int| <no value> | NumConsumers is the number of consumers from the queue.  |
| queue_size |int| <no value> | QueueSize is the maximum number of batches allowed in queue at a given time.  |

## exporterhelper-RetrySettings

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| enabled |bool| true | Enabled indicates whether to not retry sending batches in case of export failure.  |
| initial_interval |time.Duration| 5s | InitialInterval the time to wait after the first failure before retrying.  |
| max_interval |time.Duration| 30s | MaxInterval is the upper bound on backoff interval. Once this value is reached the delay between consecutive retries will always be `MaxInterval`.  |
| max_elapsed_time |time.Duration| 5m0s | MaxElapsedTime is the maximum amount of time (including retries) spent trying to send a request/batch. Once this value is reached, the data is discarded.  |

