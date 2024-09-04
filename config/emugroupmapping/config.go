package emugroupmapping

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure github_emu_group_mapping resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_emu_group_mapping", func(r *config.Resource) {

		r.ShortGroup = "team"

		r.References["team_slug"] = config.Reference{
			TerraformName: "github_team",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("slug",true)`,
		}
	})
}
