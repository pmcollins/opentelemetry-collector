## memorylimiter-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| check_interval |time.Duration| <no value> | CheckInterval is the time between measurements of memory usage for the purposes of avoiding going over the limits. Defaults to zero, so no checks will be performed.  |
| limit_mib |uint32| <no value> | MemoryLimitMiB is the maximum amount of memory, in MiB, targeted to be allocated by the process.  |
| spike_limit_mib |uint32| <no value> | MemorySpikeLimitMiB is the maximum, in MiB, spike expected between the measurements of memory usage.  |
| ballast_size_mib |uint32| <no value> | BallastSizeMiB is the size, in MiB, of the ballast size being used by the process.  |
| limit_percentage |uint32| <no value> | MemoryLimitPercentage is the maximum amount of memory, in %, targeted to be allocated by the process. The fixed memory settings MemoryLimitMiB has a higher precedence.  |
| spike_limit_percentage |uint32| <no value> | MemorySpikePercentage is the maximum, in percents against the total memory, spike expected between the measurements of memory usage.  |

