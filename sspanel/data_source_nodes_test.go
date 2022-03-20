package sspanel

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceScaffolding(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.sspanel_nodes.foo", "nodes.0.id", regexp.MustCompile("^1$")),
				),
			},
		},
	})
}

const testAccDataSourceScaffolding = `
data "sspanel_nodes" "foo" { }
`
