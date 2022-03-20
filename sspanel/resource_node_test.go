package sspanel

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceNode(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceNode,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"sspanel_node.foo", "id", regexp.MustCompile("^[0-9]*")),
				),
			},
		},
	})
}

const testAccResourceNode = `
resource "sspanel_node" "foo" {
  name = "Test node"
  server = "test.trojan.one;port=443"
  sort = 14
}
`
