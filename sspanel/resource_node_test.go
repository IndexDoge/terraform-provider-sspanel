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
			{
				Config: testAccResourceNodeUpdateNodeIp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sspanel_node.foo", "node_ip", "1.1.1.1")),
			},
		},
	})
}

const testAccResourceNode = `
resource "sspanel_node" "foo" {
  name = "Test node"
  server = "1.1.1.1.nip.io;port=443"
  node_ip = "127.0.0.1"
  sort = 14
}
`

const testAccResourceNodeUpdateNodeIp = `
resource "sspanel_node" "foo" {
  name = "Test node"
  server = "1.1.1.1.nip.io;port=443"
  node_ip = "1.1.1.1"
  sort = 14
}
`
