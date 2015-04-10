package inject

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func indent(text string, depth int) string {
	indent := strings.Repeat("  ", depth)
	return indent + strings.Replace(text, "\n", "\n"+indent, -1)
}

func ptrString(ptr interface{}) string {
	return fmt.Sprintf("%s=%p", reflect.TypeOf(ptr), ptr)
}

func mapString(m map[string]string) string {
	if len(m) == 0 {
		return "map[]"
	}
	b := make([]string, 0, len(m))
	for k, v := range m {
		b = append(b, fmt.Sprintf("%s: %s", k, v))
	}
	sort.Strings(b)
	return fmt.Sprintf("map[\n%s\n]", indent(strings.Join(b, ",\n"), 1))
}

func arrayString(a []string) string {
	if len(a) == 0 {
		return "[]"
	}
	for i, v := range a {
		a[i] = fmt.Sprint(v)
	}
	return fmt.Sprintf("[\n%s\n]", indent(strings.Join(a, ",\n"), 1))
}
