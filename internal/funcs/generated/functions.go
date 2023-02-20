// Code generated by "gen"; DO NOT EDIT.
package funcs

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl-lang/schema"
)

var (
	v1_4_0 = version.Must(version.NewVersion("1.4.0"))
)

func Functions(v *version.Version) map[string]schema.FunctionSignature {
	if v.GreaterThanOrEqual(v1_4_0) {
		return v1_4_0_Functions()
	}

	return v1_4_0_Functions()
}
