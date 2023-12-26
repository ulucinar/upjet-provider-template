/*
Copyright 2021 Upbound Inc.
*/

package null

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("null_resource", func(r *ujconfig.Resource) {
		r.Kind = "Resource"
		// And other overrides.

		// an example renaming from triggers -> triggers_renamed
		s := r.TerraformResource.Schema["triggers"]
		delete(r.TerraformResource.Schema, "triggers")
		r.TerraformResource.Schema["triggers_renamed"] = s
		r.Version = "v1beta1"
	})
}
