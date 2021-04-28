## batchprocessor-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| timeout |time.Duration| 200ms | Timeout sets the time after which a batch will be sent regardless of size.  |
| send_batch_size |uint32| 8192 | SendBatchSize is the size of a batch which after hit, will trigger it to be sent.  |
| send_batch_max_size |uint32| <no value> | SendBatchMaxSize is the maximum size of a batch. Larger batches are split into smaller units. Default value is 0, that means no maximum size.  |

