package sql_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/sql"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/sql/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type SqlVirtualNetworkRuleResource struct{}

/*
	---Testing for Success---
	Test a basic SQL virtual network rule configuration setup and update scenario, and
	validate that new property is set correctly.
*/
func TestAccSqlVirtualNetworkRule_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sql_virtual_network_rule", "test")
	r := SqlVirtualNetworkRuleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("ignore_missing_vnet_service_endpoint").HasValue("false"),
			),
		},
		data.ImportStep(),
		{
			Config: r.withUpdates(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("ignore_missing_vnet_service_endpoint").HasValue("true"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSqlVirtualNetworkRule_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sql_virtual_network_rule", "test")
	r := SqlVirtualNetworkRuleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("ignore_missing_vnet_service_endpoint").HasValue("false"),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

/*
	---Testing for Success---
	Test an update to the SQL Virtual Network Rule to connect to a different subnet, and
	validate that new subnet is set correctly.
*/
func TestAccSqlVirtualNetworkRule_switchSubnets(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sql_virtual_network_rule", "test")
	r := SqlVirtualNetworkRuleResource{}

	// Create regex strings that will ensure that one subnet name exists, but not the other
	preConfigRegex := regexp.MustCompile(fmt.Sprintf("(subnet1%d)$|(subnet[^2]%d)$", data.RandomInteger, data.RandomInteger))  // subnet 1 but not 2
	postConfigRegex := regexp.MustCompile(fmt.Sprintf("(subnet2%d)$|(subnet[^1]%d)$", data.RandomInteger, data.RandomInteger)) // subnet 2 but not 1

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.subnetSwitchPre(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("subnet_id").MatchesRegex(preConfigRegex),
			),
		},
		{
			Config: r.subnetSwitchPost(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("subnet_id").MatchesRegex(postConfigRegex),
			),
		},
	})
}

/*
	---Testing for Success---
*/
func TestAccSqlVirtualNetworkRule_disappears(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sql_virtual_network_rule", "test")
	r := SqlVirtualNetworkRuleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		data.DisappearsStep(acceptance.DisappearsStepData{
			Config:       r.basic,
			TestResource: r,
		}),
	})
}

/*
	--Testing for Success--
	Test if we are able to create a vnet without the SQL endpoint, but SQL rule
	is still applied since the endpoint validation will be set to false.
*/
func TestAccSqlVirtualNetworkRule_ignoreEndpointValid(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sql_virtual_network_rule", "test")
	r := SqlVirtualNetworkRuleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.ignoreEndpointValid(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

/*
	--Testing for Failure--
	Test if we are able to create a vnet with out the SQL endpoint, but SQL rule
	is still applied since the endpoint validation will be set to false.
*/
func TestAccSqlVirtualNetworkRule_IgnoreEndpointInvalid(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sql_virtual_network_rule", "test")
	r := SqlVirtualNetworkRuleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config:      r.ignoreEndpointInvalid(data),
			ExpectError: regexp.MustCompile("Code=\"VirtualNetworkRuleBadRequest\""),
		},
	})
}

/*
	--Testing for Success--
	Test if we are able to create multiple subnets and connect multiple subnets to the
	SQL server.
*/
func TestAccSqlVirtualNetworkRule_multipleSubnets(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sql_virtual_network_rule", "test")
	resourceName2 := "azurerm_sql_virtual_network_rule.rule2"
	resourceName3 := "azurerm_sql_virtual_network_rule.rule3"
	r := SqlVirtualNetworkRuleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.multipleSubnets(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(resourceName2).ExistsInAzure(r),
				check.That(resourceName3).ExistsInAzure(r),
			),
		},
	})
}

/*
	--Testing for Failure--
	Validation Function Tests - Invalid Name Validations
*/
func TestResourceAzureRMSqlVirtualNetworkRule_invalidNameValidation(t *testing.T) {
	cases := []struct {
		Value    string
		ErrCount int
	}{
		// Must only contain alphanumeric characters, periods, underscores or hyphens (4 cases)
		{
			Value:    "test!Rule",
			ErrCount: 1,
		},
		{
			Value:    "test&Rule",
			ErrCount: 1,
		},
		{
			Value:    "test:Rule",
			ErrCount: 1,
		},
		{
			Value:    "test'Rule",
			ErrCount: 1,
		},
		// Cannot be more than 64 characters (1 case - ensure starts with a letter)
		{
			Value:    fmt.Sprintf("v%s", acctest.RandString(64)),
			ErrCount: 1,
		},
		// Cannot be empty (1 case)
		{
			Value:    "",
			ErrCount: 1,
		},
		// Cannot be single character (1 case)
		{
			Value:    "a",
			ErrCount: 1,
		},
		// Cannot end in a hyphen (1 case)
		{
			Value:    "testRule-",
			ErrCount: 1,
		},
		// Cannot end in a period (1 case)
		{
			Value:    "testRule.",
			ErrCount: 1,
		},
		// Cannot start with a period, underscore or hyphen (3 cases)
		{
			Value:    ".testRule",
			ErrCount: 1,
		},
		{
			Value:    "_testRule",
			ErrCount: 1,
		},
		{
			Value:    "-testRule",
			ErrCount: 1,
		},
	}

	for _, tc := range cases {
		_, errors := sql.ValidateSqlVirtualNetworkRuleName(tc.Value, "azurerm_sql_virtual_network_rule")

		if len(errors) != tc.ErrCount {
			t.Fatalf("Bad: Expected the Azure RM SQL Virtual Network Rule Name to trigger a validation error.")
		}
	}
}

/*
	--Testing for Success--
	Validation Function Tests - (Barely) Valid Name Validations
*/
func TestResourceAzureRMSqlVirtualNetworkRule_validNameValidation(t *testing.T) {
	cases := []struct {
		Value    string
		ErrCount int
	}{
		// Test all lowercase
		{
			Value:    "thisisarule",
			ErrCount: 0,
		},
		// Test all uppercase
		{
			Value:    "THISISARULE",
			ErrCount: 0,
		},
		// Test alternating cases
		{
			Value:    "tHiSiSaRuLe",
			ErrCount: 0,
		},
		// Test hyphens
		{
			Value:    "this-is-a-rule",
			ErrCount: 0,
		},
		// Test multiple hyphens in a row
		{
			Value:    "this----1s----a----ru1e",
			ErrCount: 0,
		},
		// Test underscores
		{
			Value:    "this_is_a_rule",
			ErrCount: 0,
		},
		// Test ending with underscore
		{
			Value:    "this_is_a_rule_",
			ErrCount: 0,
		},
		// Test multiple underscoress in a row
		{
			Value:    "this____1s____a____ru1e",
			ErrCount: 0,
		},
		// Test periods
		{
			Value:    "this.is.a.rule",
			ErrCount: 0,
		},
		// Test multiple periods in a row
		{
			Value:    "this....1s....a....ru1e",
			ErrCount: 0,
		},
		// Test numbers
		{
			Value:    "1108501298509850810258091285091820-5",
			ErrCount: 0,
		},
		// Test a lot of hyphens and numbers
		{
			Value:    "x-5-4-1-2-5-2-6-1-5-2-5-1-2-5-6-2-2",
			ErrCount: 0,
		},
		// Test a lot of underscores and numbers
		{
			Value:    "x_5_4_1_2_5_2_6_1_5_2_5_1_2_5_6_2_2",
			ErrCount: 0,
		},
		// Test a lot of periods and numbers
		{
			Value:    "x.5.4.1.2.5.2.6.1.5.2.5.1.2.5.6.2.2",
			ErrCount: 0,
		},
		// Test exactly 64 characters
		{
			Value:    fmt.Sprintf("v%s", acctest.RandString(63)),
			ErrCount: 0,
		},
	}

	for _, tc := range cases {
		_, errors := sql.ValidateSqlVirtualNetworkRuleName(tc.Value, "azurerm_sql_virtual_network_rule")

		if len(errors) != tc.ErrCount {
			t.Fatalf("Bad: Expected the Azure RM SQL Virtual Network Rule Name pass name validation successfully but triggered a validation error.")
		}
	}
}

func (r SqlVirtualNetworkRuleResource) Exists(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.VirtualNetworkRuleID(state.ID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Sql.VirtualNetworkRulesClient.Get(ctx, id.ResourceGroup, id.ServerName, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving Sql Virtual Network Rule %q (Server %q / Resource Group %q): %+v", id.Name, id.ServerName, id.ResourceGroup, err)
	}
	return utils.Bool(true), nil
}

func (r SqlVirtualNetworkRuleResource) Destroy(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.VirtualNetworkRuleID(state.ID)
	if err != nil {
		return nil, err
	}
	rulesClient := client.Sql.VirtualNetworkRulesClient
	future, err := rulesClient.Delete(ctx, id.ResourceGroup, id.ServerName, id.Name)
	if err != nil {
		return nil, fmt.Errorf("deleting Sql Virtual Network Rule %q (Server %q / Resource Group %q): %+v", id.Name, id.ServerName, id.ResourceGroup, err)
	}
	if err := future.WaitForCompletionRef(ctx, rulesClient.Client); err != nil {
		return nil, fmt.Errorf("waiting for deletion for Sql Virtual Network Rule %q (Server %q / Resource Group %q): %+v", id.Name, id.ServerName, id.ResourceGroup, err)
	}
	return utils.Bool(true), nil
}

/*
	(This test configuration is intended to succeed.)
	Basic Provisioning Configuration
*/
func (r SqlVirtualNetworkRuleResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvnet%d"
  address_space       = ["10.7.29.0/29"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.7.29.0/29"
  service_endpoints    = ["Microsoft.Sql"]
}

resource "azurerm_sql_server" "test" {
  name                         = "acctestsqlserver%d"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  version                      = "12.0"
  administrator_login          = "missadmin"
  administrator_login_password = "${md5(%d)}!"
}

resource "azurerm_sql_virtual_network_rule" "test" {
  name                                 = "acctestsqlvnetrule%d"
  resource_group_name                  = azurerm_resource_group.test.name
  server_name                          = azurerm_sql_server.test.name
  subnet_id                            = azurerm_subnet.test.id
  ignore_missing_vnet_service_endpoint = false
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func (r SqlVirtualNetworkRuleResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_sql_virtual_network_rule" "import" {
  name                                 = azurerm_sql_virtual_network_rule.test.name
  resource_group_name                  = azurerm_sql_virtual_network_rule.test.resource_group_name
  server_name                          = azurerm_sql_virtual_network_rule.test.server_name
  subnet_id                            = azurerm_sql_virtual_network_rule.test.subnet_id
  ignore_missing_vnet_service_endpoint = azurerm_sql_virtual_network_rule.test.ignore_missing_vnet_service_endpoint
}
`, r.basic(data))
}

/*
	(This test configuration is intended to succeed.)
	Basic Provisioning Update Configuration (all other properties would recreate the rule)
	ignore_missing_vnet_service_endpoint (false ==> true)
*/
func (r SqlVirtualNetworkRuleResource) withUpdates(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvnet%d"
  address_space       = ["10.7.29.0/29"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.7.29.0/29"
  service_endpoints    = ["Microsoft.Sql"]
}

resource "azurerm_sql_server" "test" {
  name                         = "acctestsqlserver%d"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  version                      = "12.0"
  administrator_login          = "missadmin"
  administrator_login_password = "${md5(%d)}!"
}

resource "azurerm_sql_virtual_network_rule" "test" {
  name                                 = "acctestsqlvnetrule%d"
  resource_group_name                  = azurerm_resource_group.test.name
  server_name                          = azurerm_sql_server.test.name
  subnet_id                            = azurerm_subnet.test.id
  ignore_missing_vnet_service_endpoint = true
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

/*
	(This test configuration is intended to succeed.)
	This test is designed to set up a scenario where a user would want to update the subnet
	on a given SQL virtual network rule. This configuration sets up the resources initially.
*/
func (r SqlVirtualNetworkRuleResource) subnetSwitchPre(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvnet%d"
  address_space       = ["10.7.29.0/24"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test1" {
  name                 = "subnet1%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.7.29.0/25"
  service_endpoints    = ["Microsoft.Sql"]
}

resource "azurerm_subnet" "test2" {
  name                 = "subnet2%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.7.29.128/25"
  service_endpoints    = ["Microsoft.Sql"]
}

resource "azurerm_sql_server" "test" {
  name                         = "acctestsqlserver%d"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  version                      = "12.0"
  administrator_login          = "missadmin"
  administrator_login_password = "${md5(%d)}!"
}

resource "azurerm_sql_virtual_network_rule" "test" {
  name                = "acctestsqlvnetrule%d"
  resource_group_name = azurerm_resource_group.test.name
  server_name         = azurerm_sql_server.test.name
  subnet_id           = azurerm_subnet.test1.id
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

/*
	(This test configuration is intended to succeed.)
	This test is designed to set up a scenario where a user would want to update the subnet
	on a given SQL virtual network rule. This configuration contains the update from
	azurerm_subnet.test1 to azurerm_subnet.test2.
*/
func (r SqlVirtualNetworkRuleResource) subnetSwitchPost(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvnet%d"
  address_space       = ["10.7.29.0/24"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test1" {
  name                 = "subnet1%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.7.29.0/25"
  service_endpoints    = ["Microsoft.Sql"]
}

resource "azurerm_subnet" "test2" {
  name                 = "subnet2%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.7.29.128/25"
  service_endpoints    = ["Microsoft.Sql"]
}

resource "azurerm_sql_server" "test" {
  name                         = "acctestsqlserver%d"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  version                      = "12.0"
  administrator_login          = "missadmin"
  administrator_login_password = "${md5(%d)}!"
}

resource "azurerm_sql_virtual_network_rule" "test" {
  name                = "acctestsqlvnetrule%d"
  resource_group_name = azurerm_resource_group.test.name
  server_name         = azurerm_sql_server.test.name
  subnet_id           = azurerm_subnet.test2.id
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

/*
	(This test configuration is intended to succeed.)
	Succeeds because subnet's service_endpoints does not include 'Microsoft.Sql' and the SQL
    virtual network rule is set to *not* validate that the service_endpoint includes that value.
    The endpoint is purposefully set to Microsoft.Storage.
*/
func (r SqlVirtualNetworkRuleResource) ignoreEndpointValid(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvnet%d"
  address_space       = ["10.7.29.0/29"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.7.29.0/29"
  service_endpoints    = ["Microsoft.Storage"]
}

resource "azurerm_sql_server" "test" {
  name                         = "acctestsqlserver%d"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  version                      = "12.0"
  administrator_login          = "missadmin"
  administrator_login_password = "${md5(%d)}!"
}

resource "azurerm_sql_virtual_network_rule" "test" {
  name                                 = "acctestsqlvnetrule%d"
  resource_group_name                  = azurerm_resource_group.test.name
  server_name                          = azurerm_sql_server.test.name
  subnet_id                            = azurerm_subnet.test.id
  ignore_missing_vnet_service_endpoint = true
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

/*
	(This test configuration is intended to fail.)
	Fails because subnet's service_endpoints does not include 'Microsoft.Sql' and the SQL
    virtual network rule is set to validate that the service_endpoint includes that value.
    The endpoint is purposefully set to Microsoft.Storage.
*/
func (r SqlVirtualNetworkRuleResource) ignoreEndpointInvalid(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestvnet%d"
  address_space       = ["10.7.29.0/29"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.7.29.0/29"
  service_endpoints    = ["Microsoft.Storage"]
}

resource "azurerm_sql_server" "test" {
  name                         = "acctestsqlserver%d"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  version                      = "12.0"
  administrator_login          = "missadmin"
  administrator_login_password = "${md5(%d)}!"
}

resource "azurerm_sql_virtual_network_rule" "test" {
  name                                 = "acctestsqlvnetrule%d"
  resource_group_name                  = azurerm_resource_group.test.name
  server_name                          = azurerm_sql_server.test.name
  subnet_id                            = azurerm_subnet.test.id
  ignore_missing_vnet_service_endpoint = false
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

/*
	(This test configuration is intended to succeed.)
	This configuration sets up 3 subnets in 2 different virtual networks, and adds
	SQL virtual network rules for all 3 subnets to the SQL server.
*/
func (r SqlVirtualNetworkRuleResource) multipleSubnets(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_virtual_network" "vnet1" {
  name                = "acctestvnet1%d"
  address_space       = ["10.7.29.0/24"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_virtual_network" "vnet2" {
  name                = "acctestvnet2%d"
  address_space       = ["10.1.29.0/29"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "vnet1_subnet1" {
  name                 = "acctestsubnet1%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.vnet1.name
  address_prefix       = "10.7.29.0/29"
  service_endpoints    = ["Microsoft.Sql"]
}

resource "azurerm_subnet" "vnet1_subnet2" {
  name                 = "acctestsubnet2%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.vnet1.name
  address_prefix       = "10.7.29.128/29"
  service_endpoints    = ["Microsoft.Sql"]
}

resource "azurerm_subnet" "vnet2_subnet1" {
  name                 = "acctestsubnet3%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.vnet2.name
  address_prefix       = "10.1.29.0/29"
  service_endpoints    = ["Microsoft.Sql"]
}

resource "azurerm_sql_server" "test" {
  name                         = "acctestsqlserver1%d"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  version                      = "12.0"
  administrator_login          = "missadmin"
  administrator_login_password = "${md5(%d)}!"
}

resource "azurerm_sql_virtual_network_rule" "test" {
  name                                 = "acctestsqlvnetrule1%d"
  resource_group_name                  = azurerm_resource_group.test.name
  server_name                          = azurerm_sql_server.test.name
  subnet_id                            = azurerm_subnet.vnet1_subnet1.id
  ignore_missing_vnet_service_endpoint = false
}

resource "azurerm_sql_virtual_network_rule" "rule2" {
  name                                 = "acctestsqlvnetrule2%d"
  resource_group_name                  = azurerm_resource_group.test.name
  server_name                          = azurerm_sql_server.test.name
  subnet_id                            = azurerm_subnet.vnet1_subnet2.id
  ignore_missing_vnet_service_endpoint = false
}

resource "azurerm_sql_virtual_network_rule" "rule3" {
  name                                 = "acctestsqlvnetrule3%d"
  resource_group_name                  = azurerm_resource_group.test.name
  server_name                          = azurerm_sql_server.test.name
  subnet_id                            = azurerm_subnet.vnet2_subnet1.id
  ignore_missing_vnet_service_endpoint = false
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}
