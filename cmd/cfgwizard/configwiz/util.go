package configwiz

import (
	"reflect"
	"strings"
)

func splitTag(s reflect.StructTag) (string, string) {
	trimmed := strings.Trim(strings.TrimPrefix(string(s), "mapstructure:"), `"`)
	split := strings.Split(trimmed, ",")
	meta := ""
	if len(split) == 2 {
		meta = strings.TrimSpace(split[1])
	}
	return split[0], meta
}
