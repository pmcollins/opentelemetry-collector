## spanprocessor-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| include |[filterconfig-MatchProperties](#filterconfig-MatchProperties)| <no value> | Include specifies the set of span/log properties that must be present in order for this processor to apply to it. Note: If `exclude` is specified, the span/log is compared against those properties after the `include` properties. This is an optional field. If neither `include` and `exclude` are set, all span/logs are processed. If `include` is set and `exclude` isn't set, then all span/logs matching the properties in this structure are processed.  |
| exclude |[filterconfig-MatchProperties](#filterconfig-MatchProperties)| <no value> | Exclude specifies when this processor will not be applied to the span/logs which match the specified properties. Note: The `exclude` properties are checked after the `include` properties, if they exist, are checked. If `include` isn't specified, the `exclude` properties are checked against all span/logs. This is an optional field. If neither `include` and `exclude` are set, all span/logs are processed. If `exclude` is set and `include` isn't set, then all span/logs  that do no match the properties in this structure are processed.  |
| name |[spanprocessor-Name](#spanprocessor-Name)| <no value> | Rename specifies the components required to re-name a span. The `from_attributes` field needs to be set for this processor to be properly configured. Note: The field name is `Rename` to avoid collision with the Name() method from configmodels.ProcessorSettings.NamedEntity  |

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

## spanprocessor-Name

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| from_attributes |[]string| <no value> | FromAttributes represents the attribute keys to pull the values from to generate the new span name. All attribute keys are required in the span to re-name a span. If any attribute is missing from the span, no re-name will occur. Note: The new span name is constructed in order of the `from_attributes` specified in the configuration. This field is required and cannot be empty.  |
| separator |string| <no value> | Separator is the string used to separate attributes values in the new span name. If no value is set, no separator is used between attribute values. Used with FromAttributes only.  |
| to_attributes |[spanprocessor-ToAttributes](#spanprocessor-ToAttributes)| <no value> | ToAttributes specifies a configuration to extract attributes from span name.  |

## spanprocessor-ToAttributes

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| rules |[]string| <no value> | Rules is a list of rules to extract attribute values from span name. The values in the span name are replaced by extracted attribute names. Each rule in the list is a regex pattern string. Span name is checked against the regex. If it matches then all named subexpressions of the regex are extracted as attributes and are added to the span. Each subexpression name becomes an attribute name and subexpression matched portion becomes the attribute value. The matched portion in the span name is replaced by extracted attribute name. If the attributes already exist in the span then they will be overwritten. The process is repeated for all rules in the order they are specified. Each subsequent rule works on the span name that is the output after processing the previous rule.  |
| break_after_match |bool| <no value> | BreakAfterMatch specifies if processing of rules should stop after the first match. If it is false rule processing will continue to be performed over the modified span name.  |

