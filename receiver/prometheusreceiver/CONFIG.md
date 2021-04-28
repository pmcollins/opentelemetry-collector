## prometheusreceiver-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| buffer_period |time.Duration| <no value> |  |
| buffer_count |int| <no value> |  |
| use_start_time_metric |bool| <no value> |  |
| start_time_metric_regex |string| <no value> |  |
| config |interface {}| <no value> | ConfigPlaceholder is just an entry to make the configuration pass a check that requires that all keys present in the config actually exist on the structure, ie.: it will error if an unknown key is present.  |

