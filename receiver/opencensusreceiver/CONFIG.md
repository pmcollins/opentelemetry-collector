## opencensusreceiver-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| endpoint |string| 0.0.0.0:55678 | Endpoint configures the address for this network connection. For TCP and UDP networks, the address has the form "host:port". The host must be a literal IP address, or a host name that can be resolved to IP addresses. The port must be a literal port number or a service name. If the host is a literal IPv6 address it must be enclosed in square brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80". The zone specifies the scope of the literal IPv6 address as defined in RFC 4007.  |
| transport |string| tcp | Transport to use. Known protocols are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket".  |
| tls_settings |[configtls-TLSServerSetting](#configtls-TLSServerSetting)| <no value> | Configures the protocol to use TLS. The default value is nil, which will cause the protocol to not use TLS.  |
| max_recv_msg_size_mib |uint64| <no value> | MaxRecvMsgSizeMiB sets the maximum size (in MiB) of messages accepted by the server.  |
| max_concurrent_streams |uint32| <no value> | MaxConcurrentStreams sets the limit on the number of concurrent streams to each ServerTransport. It has effect only for streaming RPCs.  |
| read_buffer_size |int| 524288 | ReadBufferSize for gRPC server. See grpc.ReadBufferSize (https://godoc.org/google.golang.org/grpc#ReadBufferSize).  |
| write_buffer_size |int| <no value> | WriteBufferSize for gRPC server. See grpc.WriteBufferSize (https://godoc.org/google.golang.org/grpc#WriteBufferSize).  |
| keepalive |[configgrpc-KeepaliveServerConfig](#configgrpc-KeepaliveServerConfig)| <no value> | Keepalive anchor for all the settings related to keepalive.  |
| auth |[configauth-Authentication](#configauth-Authentication)| <no value> | Auth for this receiver  |
| cors_allowed_origins |[]string| <no value> | CorsOrigins are the allowed CORS origins for HTTP/JSON requests to grpc-gateway adapter for the OpenCensus receiver. See github.com/rs/cors An empty list means that CORS is not enabled at all. A wildcard (*) can be used to match any origin or one or more characters of an origin.  |

## configtls-TLSServerSetting

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| ca_file |string| <no value> | Path to the CA cert. For a client this verifies the server certificate. For a server this verifies client certificates. If empty uses system root CA. (optional)  |
| cert_file |string| <no value> | Path to the TLS cert to use for TLS required connections. (optional)  |
| key_file |string| <no value> | Path to the TLS key to use for TLS required connections. (optional)  |
| client_ca_file |string| <no value> | Path to the TLS cert to use by the server to verify a client certificate. (optional) This sets the ClientCAs and ClientAuth to RequireAndVerifyClientCert in the TLSConfig. Please refer to https://godoc.org/crypto/tls#Config for more information. (optional)  |

## configgrpc-KeepaliveServerConfig

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| server_parameters |[configgrpc-KeepaliveServerParameters](#configgrpc-KeepaliveServerParameters)| <no value> |  |
| enforcement_policy |[configgrpc-KeepaliveEnforcementPolicy](#configgrpc-KeepaliveEnforcementPolicy)| <no value> |  |

## configgrpc-KeepaliveServerParameters

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| max_connection_idle |time.Duration| <no value> |  |
| max_connection_age |time.Duration| <no value> |  |
| max_connection_age_grace |time.Duration| <no value> |  |
| time |time.Duration| <no value> |  |
| timeout |time.Duration| <no value> |  |

## configgrpc-KeepaliveEnforcementPolicy

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| min_time |time.Duration| <no value> |  |
| permit_without_stream |bool| <no value> |  |

## configauth-Authentication

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| attribute |string| <no value> | The attribute (header name) to look for auth data. Optional, default value: "authentication".  |
| oidc |[configauth-OIDC](#configauth-OIDC)| <no value> | OIDC configures this receiver to use the given OIDC provider as the backend for the authentication mechanism. Required.  |

## configauth-OIDC

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| issuer_url |string| <no value> | IssuerURL is the base URL for the OIDC provider. Required.  |
| audience |string| <no value> | Audience of the token, used during the verification. For example: "https://accounts.google.com" or "https://login.salesforce.com". Required.  |
| issuer_ca_path |string| <no value> | The local path for the issuer CA's TLS server cert. Optional.  |
| username_claim |string| <no value> | The claim to use as the username, in case the token's 'sub' isn't the suitable source. Optional.  |
| groups_claim |string| <no value> | The claim that holds the subject's group membership information. Optional.  |

