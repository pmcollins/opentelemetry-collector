## loggingexporter-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| loglevel |string| info | LogLevel defines log level of the logging exporter; options are debug, info, warn, error.  |
| sampling_initial |int| 2 | SamplingInitial defines how many samples are initially logged during each second.  |
| sampling_thereafter |int| 500 | SamplingThereafter defines the sampling rate after the initial samples are logged.  |

