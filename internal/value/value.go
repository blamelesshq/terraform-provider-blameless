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
