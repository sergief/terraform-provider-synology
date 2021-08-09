package provider

import (
	"fmt"
	"testing"

	"github.com/sergief/terraform-provider-synology/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSynologyFolderItem(t *testing.T) {
	path := "/home/test-acc-synology-folder"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSynologyFolderItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSynologyFolderItemConfig(path),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSynologyFolderItemExists("synology_folder.acc_test_new", path),
				),
			},
		},
	})
}

func testAccCheckSynologyFolderItemDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(client.SynologyClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "synology_folder" {
			continue
		}

		path := rs.Primary.Attributes["path"]

		err := client.Delete(path, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckSynologyFolderItemConfig(path string) string {
	return fmt.Sprintf(`
	resource "synology_folder" "acc_test_new" {
		path = "%s"
	}
	`, path)
}

func testAccCheckSynologyFolderItemExists(n string, path string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID set")
		}

		if rs.Primary.Attributes["path"] != path {
			return fmt.Errorf("Path doesn't match")
		}

		return nil
	}
}
