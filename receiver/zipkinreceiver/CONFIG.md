## zipkinreceiver-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| endpoint |string| 0.0.0.0:9411 | Endpoint configures the listening address for the server.  |
| tls_settings |[configtls-TLSServerSetting](#configtls-TLSServerSetting)| <no value> | TLSSetting struct exposes TLS client configuration.  |
| cors_allowed_origins |[]string| <no value> | CorsOrigins are the allowed CORS origins for HTTP/JSON requests to grpc-gateway adapter for the OTLP receiver. See github.com/rs/cors An empty list means that CORS is not enabled at all. A wildcard (*) can be used to match any origin or one or more characters of an origin.  |
| cors_allowed_headers |[]string| <no value> | CorsHeaders are the allowed CORS headers for HTTP/JSON requests to grpc-gateway adapter for the OTLP receiver. See github.com/rs/cors CORS needs to be enabled first by providing a non-empty list in CorsOrigins A wildcard (*) can be used to match any header.  |
| parse_string_tags |bool| <no value> | If enabled the zipkin receiver will attempt to parse string tags/binary annotations into int/bool/float. Disabled by default  |

## configtls-TLSServerSetting

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| ca_file |string| <no value> | Path to the CA cert. For a client this verifies the server certificate. For a server this verifies client certificates. If empty uses system root CA. (optional)  |
| cert_file |string| <no value> | Path to the TLS cert to use for TLS required connections. (optional)  |
| key_file |string| <no value> | Path to the TLS key to use for TLS required connections. (optional)  |
| client_ca_file |string| <no value> | Path to the TLS cert to use by the server to verify a client certificate. (optional) This sets the ClientCAs and ClientAuth to RequireAndVerifyClientCert in the TLSConfig. Please refer to https://godoc.org/crypto/tls#Config for more information. (optional)  |

