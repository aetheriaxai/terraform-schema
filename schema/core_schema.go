package schema

import (
	"fmt"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl-lang/schema"
	mod_v0_12 "github.com/hashicorp/terraform-schema/internal/schema/0.12"
	mod_v0_13 "github.com/hashicorp/terraform-schema/internal/schema/0.13"
	mod_v0_14 "github.com/hashicorp/terraform-schema/internal/schema/0.14"
	mod_v0_15 "github.com/hashicorp/terraform-schema/internal/schema/0.15"
	mod_v1_1 "github.com/hashicorp/terraform-schema/internal/schema/1.1"
	mod_v1_2 "github.com/hashicorp/terraform-schema/internal/schema/1.2"
)

var (
	v0_12 = version.Must(version.NewVersion("0.12"))
	v0_13 = version.Must(version.NewVersion("0.13"))
	v0_14 = version.Must(version.NewVersion("0.14"))
	v0_15 = version.Must(version.NewVersion("0.15"))
	v1_1  = version.Must(version.NewVersion("1.1"))
	v1_2  = version.Must(version.NewVersion("1.2"))
)

// CoreModuleSchemaForVersion finds a module schema which is relevant
// for the given Terraform version.
// It will return error if such schema cannot be found.
func CoreModuleSchemaForVersion(v *version.Version) (*schema.BodySchema, error) {
	ver, err := semVer(v)
	if err != nil {
		return nil, fmt.Errorf("invalid version: %w", err)
	}

	if ver.GreaterThanOrEqual(v1_2) {
		return mod_v1_2.ModuleSchema(ver), nil
	}
	if ver.GreaterThanOrEqual(v1_1) {
		return mod_v1_1.ModuleSchema(ver), nil
	}
	if ver.GreaterThanOrEqual(v0_15) {
		return mod_v0_15.ModuleSchema(ver), nil
	}
	if ver.GreaterThanOrEqual(v0_14) {
		return mod_v0_14.ModuleSchema(ver), nil
	}
	if ver.GreaterThanOrEqual(v0_13) {
		return mod_v0_13.ModuleSchema(ver), nil
	}
	if ver.GreaterThanOrEqual(v0_12) {
		return mod_v0_12.ModuleSchema(ver), nil
	}

	return nil, NoCompatibleSchemaErr{Version: ver}
}

//go:generate go run ../internal/versiongen -w ./versions_gen.go
func CoreModuleSchemaForConstraint(vc version.Constraints) (*schema.BodySchema, error) {
	for _, v := range terraformVersions {
		if vc.Check(v) {
			return CoreModuleSchemaForVersion(v)
		}
	}

	return nil, NoCompatibleSchemaErr{Constraints: vc}
}

func semVer(ver *version.Version) (*version.Version, error) {
	// Assume that alpha/beta/rc prereleases have the same compatibility
	segments := ver.Segments64()
	segmentsOnly := fmt.Sprintf("%d.%d.%d", segments[0], segments[1], segments[2])
	return version.NewVersion(segmentsOnly)
}
