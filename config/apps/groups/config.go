package groups

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("googleworkspace_group_settings", func(r *config.Resource) {
		r.ShortGroup = "groups"
		r.References["email"] = config.Reference{
			Type: "github.com/alanraison/provider-googleworkspace/apis/directory/v1alpha1.Group",
		}
	})
}
