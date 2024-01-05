/*
Copyright 2021 Upbound Inc.
*/

package null

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("null_resource", func(r *ujconfig.Resource) {
		r.Kind = "Resource"
		// And other overrides.

		// an example renaming from triggers -> triggers_renamed
		r.Version = "v1beta1"

		r.Conversions = append(r.Conversions, conversion.NewFieldRenameConversion("v1beta1", "spec.forProvider.triggers", "v1alpha1", "spec.forProvider.triggersRenamed"))
		r.Conversions = append(r.Conversions, conversion.NewFieldRenameConversion("v1alpha1", "spec.forProvider.triggersRenamed", "v1beta1", "spec.forProvider.triggers"))
	})
}
