package notificationhub_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

type NotificationHubNamespaceDataSource struct{}

func TestAccNotificationHubNamespaceDataSource_free(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_notification_hub_namespace", "test")
	d := NotificationHubNamespaceDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: d.free(data),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(data.ResourceName, "namespace_type", "NotificationHub"),
				resource.TestCheckResourceAttr(data.ResourceName, "sku.0.name", "Free"),
				resource.TestCheckResourceAttr(data.ResourceName, "tags.%", "1"),
			),
		},
	})
}

func (d NotificationHubNamespaceDataSource) free(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

data "azurerm_notification_hub_namespace" "test" {
  name                = azurerm_notification_hub_namespace.test.name
  resource_group_name = azurerm_notification_hub_namespace.test.resource_group_name
}
`, NotificationHubNamespaceResource{}.free(data))
}
