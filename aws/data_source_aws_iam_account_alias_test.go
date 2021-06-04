package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/terraform-providers/terraform-provider-aws/atest"
)

func testAccAWSIAMAccountAliasDataSource_basic(t *testing.T) {
	dataSourceName := "data.aws_iam_account_alias.test"
	resourceName := "aws_iam_account_alias.test"

	rName := acctest.RandomWithPrefix("tf-acc-test")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { atest.PreCheck(t) },
		ErrorCheck:   atest.ErrorCheck(t, iam.EndpointsID),
		Providers:    atest.Providers,
		CheckDestroy: testAccCheckAWSIAMAccountAliasDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSIAMAccountAliasDataSourceConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "account_alias", resourceName, "account_alias"),
				),
			},
		},
	})
}

func testAccAWSIAMAccountAliasDataSourceConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_iam_account_alias" "test" {
  account_alias = %[1]q
}

data "aws_iam_account_alias" "test" {
  depends_on = [aws_iam_account_alias.test]
}
`, rName)
}
