// Copyright (C) 2023 Nicola Murino
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sftpgo

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAllowListResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: `
					resource "sftpgo_allowlist_entry" "test" {
  					  ipornet = "172.16.3.0/24"
					  protocols = 0
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("sftpgo_allowlist_entry.test", "ipornet", "172.16.3.0/24"),
					resource.TestCheckResourceAttr("sftpgo_allowlist_entry.test", "id", "172.16.3.0/24"),
					resource.TestCheckNoResourceAttr("sftpgo_allowlist_entry.test", "description"),
					resource.TestCheckResourceAttr("sftpgo_allowlist_entry.test", "protocols", "0"),
					resource.TestCheckResourceAttrSet("sftpgo_allowlist_entry.test", "created_at"),
					resource.TestCheckResourceAttrSet("sftpgo_allowlist_entry.test", "updated_at"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "sftpgo_allowlist_entry.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: `
					resource "sftpgo_allowlist_entry" "test" {
					  ipornet = "172.16.3.0/24"
					  protocols = 7
					  description = "desc"
				    }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("sftpgo_allowlist_entry.test", "ipornet", "172.16.3.0/24"),
					resource.TestCheckResourceAttr("sftpgo_allowlist_entry.test", "id", "172.16.3.0/24"),
					resource.TestCheckResourceAttr("sftpgo_allowlist_entry.test", "description", "desc"),
					resource.TestCheckResourceAttr("sftpgo_allowlist_entry.test", "protocols", "7"),
					resource.TestCheckResourceAttrSet("sftpgo_allowlist_entry.test", "created_at"),
					resource.TestCheckResourceAttrSet("sftpgo_allowlist_entry.test", "updated_at"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
