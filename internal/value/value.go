package value

import (
	"github.com/hashicorp/go-cty/cty"
)

func String(rawValue cty.Value) string {
	if rawValue.IsNull() {
		return ""
	}
	return rawValue.AsString()
}

func Bool(rawValue cty.Value) *bool {
	if rawValue.IsNull() {
		return nil
	}

	value := rawValue.True()
	return &value
}

func Int(rawValue cty.Value) *int {
	if rawValue.IsNull() {
		return nil
	}

	int64Value, _ := rawValue.AsBigFloat().Int64()
	value := int(int64Value)
	return &value
}

func StringArray(rawValue cty.Value) []string {
	results := make([]string, rawValue.LengthInt())
	if rawValue.IsNull() {
		return nil
	}

	i := 0
	rawValue.ForEachElement(func(key, val cty.Value) (stop bool) {
		results[i] = String(val)
		i++
		return stop
	})

	return results
}
