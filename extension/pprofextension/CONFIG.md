## pprofextension-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| endpoint |string| localhost:1777 | Endpoint configures the address for this network connection. The address has the form "host:port". The host must be a literal IP address, or a host name that can be resolved to IP addresses. The port must be a literal port number or a service name. If the host is a literal IPv6 address it must be enclosed in square brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80". The zone specifies the scope of the literal IPv6 address as defined in RFC 4007.  |
| block_profile_fraction |int| <no value> | Fraction of blocking events that are profiled. A value <= 0 disables profiling. See https://golang.org/pkg/runtime/#SetBlockProfileRate for details.  |
| mutex_profile_fraction |int| <no value> | Fraction of mutex contention events that are profiled. A value <= 0 disables profiling. See https://golang.org/pkg/runtime/#SetMutexProfileFraction for details.  |
| save_to_file |string| <no value> | Optional file name to save the CPU profile to. The profiling starts when the Collector starts and is saved to the file when the Collector is terminated.  |

