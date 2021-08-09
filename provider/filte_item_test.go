package provider

import (
	"fmt"
	"testing"

	"github.com/sergief/terraform-provider-synology/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSynologyFileItem(t *testing.T) {
	filename := "/home/test-acc-synology.txt"
	content := "Content for acceptance tests"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSynologyFileItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSynologyFileItemConfig(filename, content),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSynologyFileItemExists("synology_file.acc_test_new", filename, content),
				),
			},
		},
	})
}

func testAccCheckSynologyFileItemDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(client.SynologyClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "synology_file" {
			continue
		}

		filename := rs.Primary.Attributes["filename"]

		err := client.Delete(filename, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckSynologyFileItemConfig(filename, content string) string {
	return fmt.Sprintf(`
	resource "synology_file" "acc_test_new" {
		filename = "%s"
		content = "%s"
	}
	`, filename, content)
}

func testAccCheckSynologyFileItemExists(n string, filename string, content string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID set")
		}

		if rs.Primary.Attributes["filename"] != filename {
			return fmt.Errorf("Filename doesn't match")
		}

		if rs.Primary.Attributes["content"] != content {
			return fmt.Errorf("Content doesn't match")
		}

		return nil
	}
}
