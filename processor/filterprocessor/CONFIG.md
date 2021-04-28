## filterprocessor-Config

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| metrics |[filterprocessor-MetricFilters](#filterprocessor-MetricFilters)| <no value> |  |

## filterprocessor-MetricFilters

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| include |[filtermetric-MatchProperties](#filtermetric-MatchProperties)| <no value> | Include match properties describe metrics that should be included in the Collector Service pipeline, all other metrics should be dropped from further processing. If both Include and Exclude are specified, Include filtering occurs first.  |
| exclude |[filtermetric-MatchProperties](#filtermetric-MatchProperties)| <no value> | Exclude match properties describe metrics that should be excluded from the Collector Service pipeline, all other metrics should be included. If both Include and Exclude are specified, Include filtering occurs first.  |

## filtermetric-MatchProperties

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| match_type |filtermetric.MatchType| <no value> | MatchType specifies the type of matching desired  |
| regexp |[regexp-Config](#regexp-Config)| <no value> | RegexpConfig specifies options for the Regexp match type  |
| metric_names |[]string| <no value> | MetricNames specifies the list of string patterns to match metric names against. A match occurs if the metric name matches at least one string pattern in this list.  |
| expressions |[]string| <no value> | Expressions specifies the list of expr expressions to match metrics against. A match occurs if any datapoint in a metric matches at least one expression in this list.  |
| resource_attributes |[]filterconfig.Attribute| <no value> | ResourceAttributes defines a list of possible resource attributes to match metrics against. A match occurs if any resource attribute matches at least one expression in this given list.  |

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

## filtermetric-MatchProperties

| Name | Type | Default | Docs |
| ---- | ---- | ------- | ---- |
| match_type |filtermetric.MatchType| <no value> | MatchType specifies the type of matching desired  |
| regexp |[regexp-Config](#regexp-Config)| <no value> | RegexpConfig specifies options for the Regexp match type  |
| metric_names |[]string| <no value> | MetricNames specifies the list of string patterns to match metric names against. A match occurs if the metric name matches at least one string pattern in this list.  |
| expressions |[]string| <no value> | Expressions specifies the list of expr expressions to match metrics against. A match occurs if any datapoint in a metric matches at least one expression in this list.  |
| resource_attributes |[]filterconfig.Attribute| <no value> | ResourceAttributes defines a list of possible resource attributes to match metrics against. A match occurs if any resource attribute matches at least one expression in this given list.  |

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

