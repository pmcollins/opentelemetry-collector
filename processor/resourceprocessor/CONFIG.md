## resourceprocessor-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| attributes |[]processorhelper.ActionKeyValue| <no value> | AttributesActions specifies the list of actions to be applied on resource attributes. The set of actions are {INSERT, UPDATE, UPSERT, DELETE, HASH, EXTRACT}.  |

## []processorhelper-ActionKeyValue

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| key |string| <no value> | Key specifies the attribute to act upon. This is a required field.  |
| value |interface {}| <no value> | Value specifies the value to populate for the key. The type of the value is inferred from the configuration.  |
| pattern |string| <no value> | A regex pattern  must be specified for the action EXTRACT. It uses the attribute specified by `key' to extract values from The target keys are inferred based on the names of the matcher groups provided and the names will be inferred based on the values of the matcher group. Note: All subexpressions must have a name. Note: The value type of the source key must be a string. If it isn't, no extraction will occur.  |
| from_attribute |string| <no value> | FromAttribute specifies the attribute to use to populate the value. If the attribute doesn't exist, no action is performed.  |
| action |processorhelper.Action| <no value> | Action specifies the type of action to perform. The set of values are {INSERT, UPDATE, UPSERT, DELETE, HASH}. Both lower case and upper case are supported. INSERT -  Inserts the key/value to attributes when the key does not exist.           No action is applied to attributes where the key already exists.           Either Value or FromAttribute must be set. UPDATE -  Updates an existing key with a value. No action is applied           to attributes where the key does not exist.           Either Value or FromAttribute must be set. UPSERT -  Performs insert or update action depending on the attributes           containing the key. The key/value is insert to attributes           that did not originally have the key. The key/value is updated           for attributes where the key already existed.           Either Value or FromAttribute must be set. DELETE  - Deletes the attribute. If the key doesn't exist,           no action is performed. HASH    - Calculates the SHA-1 hash of an existing value and overwrites the           value with it's SHA-1 hash result. EXTRACT - Extracts values using a regular expression rule from the input           'key' to target keys specified in the 'rule'. If a target key           already exists, it will be overridden. This is a required field.  |

