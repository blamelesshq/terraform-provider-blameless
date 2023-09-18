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
