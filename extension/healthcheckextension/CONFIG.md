## healthcheckextension-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| port |uint16| <no value> | Port is the port used to publish the health check status. The default value is 13133. Deprecated: use Endpoint instead.  |
| endpoint |string| 0.0.0.0:13133 | Endpoint configures the address for this network connection. The address has the form "host:port". The host must be a literal IP address, or a host name that can be resolved to IP addresses. The port must be a literal port number or a service name. If the host is a literal IPv6 address it must be enclosed in square brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80". The zone specifies the scope of the literal IPv6 address as defined in RFC 4007.  |

