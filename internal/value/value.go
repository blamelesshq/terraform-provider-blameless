package value

import (
	"github.com/hashicorp/go-cty/cty"
)

func String(rawValue cty.Value) *string {
	if rawValue.IsNull() {
		return nil
	}
	val := rawValue.AsString()
	return &val
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
	if rawValue.IsNull() {
		return nil
	}
	results := make([]string, rawValue.LengthInt())

	i := 0
	rawValue.ForEachElement(func(key, val cty.Value) (stop bool) {
		v := String(val)
		result := ""
		if v != nil {
			result = *v
		}
		results[i] = result
		i++
		return stop
	})

	return results
}
