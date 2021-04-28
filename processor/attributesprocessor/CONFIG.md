## attributesprocessor-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| include |[filterconfig-MatchProperties](#filterconfig-MatchProperties)| <no value> | Include specifies the set of span/log properties that must be present in order for this processor to apply to it. Note: If `exclude` is specified, the span/log is compared against those properties after the `include` properties. This is an optional field. If neither `include` and `exclude` are set, all span/logs are processed. If `include` is set and `exclude` isn't set, then all span/logs matching the properties in this structure are processed.  |
| exclude |[filterconfig-MatchProperties](#filterconfig-MatchProperties)| <no value> | Exclude specifies when this processor will not be applied to the span/logs which match the specified properties. Note: The `exclude` properties are checked after the `include` properties, if they exist, are checked. If `include` isn't specified, the `exclude` properties are checked against all span/logs. This is an optional field. If neither `include` and `exclude` are set, all span/logs are processed. If `exclude` is set and `include` isn't set, then all span/logs  that do no match the properties in this structure are processed.  |
| actions |[]processorhelper.ActionKeyValue| <no value> | Actions specifies the list of attributes to act on. The set of actions are {INSERT, UPDATE, UPSERT, DELETE, HASH, EXTRACT}. This is a required field.  |

## filterconfig-MatchProperties

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| match_type |filterset.MatchType| <no value> |  |
| regexp |[regexp-Config](#regexp-Config)| <no value> |  |
| services |[]string| <no value> | Services specify the list of of items to match service name against. A match occurs if the span's service name matches at least one item in this list. This is an optional field.  |
| span_names |[]string| <no value> | SpanNames specify the list of items to match span name against. A match occurs if the span name matches at least one item in this list. This is an optional field.  |
| log_names |[]string| <no value> | LogNames is a list of strings that the LogRecord's name field must match against.  |
| attributes |[]filterconfig.Attribute| <no value> | Attributes specifies the list of attributes to match against. All of these attributes must match exactly for a match to occur. Only match_type=strict is allowed if "attributes" are specified. This is an optional field.  |
| resources |[]filterconfig.Attribute| <no value> | Resources specify the list of items to match the resources against. A match occurs if the span's resources matches at least one item in this list. This is an optional field.  |
| libraries |[]filterconfig.InstrumentationLibrary| <no value> | Libraries specify the list of items to match the implementation library against. A match occurs if the span's implementation library matches at least one item in this list. This is an optional field.  |

## regexp-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| cacheenabled |bool| <no value> | CacheEnabled determines whether match results are LRU cached to make subsequent matches faster. Cache size is unlimited unless CacheMaxNumEntries is also specified.  |
| cachemaxnumentries |int| <no value> | CacheMaxNumEntries is the max number of entries of the LRU cache that stores match results. CacheMaxNumEntries is ignored if CacheEnabled is false.  |

## []filterconfig-Attribute

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| key |string| <no value> | Key specifies the attribute key.  |
| value |interface {}| <no value> | Values specifies the value to match against. If it is not set, any value will match.  |

## []filterconfig-Attribute

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| key |string| <no value> | Key specifies the attribute key.  |
| value |interface {}| <no value> | Values specifies the value to match against. If it is not set, any value will match.  |

## []filterconfig-InstrumentationLibrary

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| name |string| <no value> |  |
| version |[string](#string)| <no value> | version match  expected actual  match  nil      <blank> yes  nil      1       yes  <blank>  <blank> yes  <blank>  1       no  1        <blank> no  1        1       yes  |

## filterconfig-MatchProperties

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| match_type |filterset.MatchType| <no value> |  |
| regexp |[regexp-Config](#regexp-Config)| <no value> |  |
| services |[]string| <no value> | Services specify the list of of items to match service name against. A match occurs if the span's service name matches at least one item in this list. This is an optional field.  |
| span_names |[]string| <no value> | SpanNames specify the list of items to match span name against. A match occurs if the span name matches at least one item in this list. This is an optional field.  |
| log_names |[]string| <no value> | LogNames is a list of strings that the LogRecord's name field must match against.  |
| attributes |[]filterconfig.Attribute| <no value> | Attributes specifies the list of attributes to match against. All of these attributes must match exactly for a match to occur. Only match_type=strict is allowed if "attributes" are specified. This is an optional field.  |
| resources |[]filterconfig.Attribute| <no value> | Resources specify the list of items to match the resources against. A match occurs if the span's resources matches at least one item in this list. This is an optional field.  |
| libraries |[]filterconfig.InstrumentationLibrary| <no value> | Libraries specify the list of items to match the implementation library against. A match occurs if the span's implementation library matches at least one item in this list. This is an optional field.  |

## regexp-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| cacheenabled |bool| <no value> | CacheEnabled determines whether match results are LRU cached to make subsequent matches faster. Cache size is unlimited unless CacheMaxNumEntries is also specified.  |
| cachemaxnumentries |int| <no value> | CacheMaxNumEntries is the max number of entries of the LRU cache that stores match results. CacheMaxNumEntries is ignored if CacheEnabled is false.  |

## []filterconfig-Attribute

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| key |string| <no value> | Key specifies the attribute key.  |
| value |interface {}| <no value> | Values specifies the value to match against. If it is not set, any value will match.  |

## []filterconfig-Attribute

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| key |string| <no value> | Key specifies the attribute key.  |
| value |interface {}| <no value> | Values specifies the value to match against. If it is not set, any value will match.  |

## []filterconfig-InstrumentationLibrary

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| name |string| <no value> |  |
| version |[string](#string)| <no value> | version match  expected actual  match  nil      <blank> yes  nil      1       yes  <blank>  <blank> yes  <blank>  1       no  1        <blank> no  1        1       yes  |

## []processorhelper-ActionKeyValue

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| key |string| <no value> | Key specifies the attribute to act upon. This is a required field.  |
| value |interface {}| <no value> | Value specifies the value to populate for the key. The type of the value is inferred from the configuration.  |
| pattern |string| <no value> | A regex pattern  must be specified for the action EXTRACT. It uses the attribute specified by `key' to extract values from The target keys are inferred based on the names of the matcher groups provided and the names will be inferred based on the values of the matcher group. Note: All subexpressions must have a name. Note: The value type of the source key must be a string. If it isn't, no extraction will occur.  |
| from_attribute |string| <no value> | FromAttribute specifies the attribute to use to populate the value. If the attribute doesn't exist, no action is performed.  |
| action |processorhelper.Action| <no value> | Action specifies the type of action to perform. The set of values are {INSERT, UPDATE, UPSERT, DELETE, HASH}. Both lower case and upper case are supported. INSERT -  Inserts the key/value to attributes when the key does not exist.           No action is applied to attributes where the key already exists.           Either Value or FromAttribute must be set. UPDATE -  Updates an existing key with a value. No action is applied           to attributes where the key does not exist.           Either Value or FromAttribute must be set. UPSERT -  Performs insert or update action depending on the attributes           containing the key. The key/value is insert to attributes           that did not originally have the key. The key/value is updated           for attributes where the key already existed.           Either Value or FromAttribute must be set. DELETE  - Deletes the attribute. If the key doesn't exist,           no action is performed. HASH    - Calculates the SHA-1 hash of an existing value and overwrites the           value with it's SHA-1 hash result. EXTRACT - Extracts values using a regular expression rule from the input           'key' to target keys specified in the 'rule'. If a target key           already exists, it will be overridden. This is a required field.  |

