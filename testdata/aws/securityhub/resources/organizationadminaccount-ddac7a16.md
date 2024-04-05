Manages a Security Hub administrator account for an organization. The AWS account utilizing this resource must be an Organizations primary account. More information about Organizations support in Security Hub can be found in the [Security Hub User Guide](https://docs.aws.amazon.com/securityhub/latest/userguide/designate-orgs-admin-account.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOrganization = new aws.organizations.Organization("exampleOrganization", {
    awsServiceAccessPrincipals: ["securityhub.amazonaws.com"],
    featureSet: "ALL",
});
const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleOrganizationAdminAccount = new aws.securityhub.OrganizationAdminAccount("exampleOrganizationAdminAccount", {adminAccountId: "123456789012"}, {
    dependsOn: [exampleOrganization],
});
// Auto enable security hub in organization member accounts
const exampleOrganizationConfiguration = new aws.securityhub.OrganizationConfiguration("exampleOrganizationConfiguration", {autoEnable: true});
```
```python
import pulumi
import pulumi_aws as aws

example_organization = aws.organizations.Organization("exampleOrganization",
    aws_service_access_principals=["securityhub.amazonaws.com"],
    feature_set="ALL")
example_account = aws.securityhub.Account("exampleAccount")
example_organization_admin_account = aws.securityhub.OrganizationAdminAccount("exampleOrganizationAdminAccount", admin_account_id="123456789012",
opts=pulumi.ResourceOptions(depends_on=[example_organization]))
# Auto enable security hub in organization member accounts
example_organization_configuration = aws.securityhub.OrganizationConfiguration("exampleOrganizationConfiguration", auto_enable=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleOrganization = new Aws.Organizations.Organization("exampleOrganization", new()
    {
        AwsServiceAccessPrincipals = new[]
        {
            "securityhub.amazonaws.com",
        },
        FeatureSet = "ALL",
    });

    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleOrganizationAdminAccount = new Aws.SecurityHub.OrganizationAdminAccount("exampleOrganizationAdminAccount", new()
    {
        AdminAccountId = "123456789012",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleOrganization,
        },
    });

    // Auto enable security hub in organization member accounts
    var exampleOrganizationConfiguration = new Aws.SecurityHub.OrganizationConfiguration("exampleOrganizationConfiguration", new()
    {
        AutoEnable = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("securityhub.amazonaws.com"),
			},
			FeatureSet: pulumi.String("ALL"),
		})
		if err != nil {
			return err
		}
		_, err = securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewOrganizationAdminAccount(ctx, "exampleOrganizationAdminAccount", &securityhub.OrganizationAdminAccountArgs{
			AdminAccountId: pulumi.String("123456789012"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleOrganization,
		}))
		if err != nil {
			return err
		}
		_, err = securityhub.NewOrganizationConfiguration(ctx, "exampleOrganizationConfiguration", &securityhub.OrganizationConfigurationArgs{
			AutoEnable: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.organizations.Organization;
import com.pulumi.aws.organizations.OrganizationArgs;
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.OrganizationAdminAccount;
import com.pulumi.aws.securityhub.OrganizationAdminAccountArgs;
import com.pulumi.aws.securityhub.OrganizationConfiguration;
import com.pulumi.aws.securityhub.OrganizationConfigurationArgs;
import com.pulumi.resources.CustomResourceOptions;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var exampleOrganization = new Organization("exampleOrganization", OrganizationArgs.builder()        
            .awsServiceAccessPrincipals("securityhub.amazonaws.com")
            .featureSet("ALL")
            .build());

        var exampleAccount = new Account("exampleAccount");

        var exampleOrganizationAdminAccount = new OrganizationAdminAccount("exampleOrganizationAdminAccount", OrganizationAdminAccountArgs.builder()        
            .adminAccountId("123456789012")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleOrganization)
                .build());

        var exampleOrganizationConfiguration = new OrganizationConfiguration("exampleOrganizationConfiguration", OrganizationConfigurationArgs.builder()        
            .autoEnable(true)
            .build());

    }
}
```
```yaml
resources:
  exampleOrganization:
    type: aws:organizations:Organization
    properties:
      awsServiceAccessPrincipals:
        - securityhub.amazonaws.com
      featureSet: ALL
  exampleAccount:
    type: aws:securityhub:Account
  exampleOrganizationAdminAccount:
    type: aws:securityhub:OrganizationAdminAccount
    properties:
      adminAccountId: '123456789012'
    options:
      dependson:
        - ${exampleOrganization}
  # Auto enable security hub in organization member accounts
  exampleOrganizationConfiguration:
    type: aws:securityhub:OrganizationConfiguration
    properties:
      autoEnable: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Security Hub Organization Admin Accounts using the AWS account ID. For example:

```sh
 $ pulumi import aws:securityhub/organizationAdminAccount:OrganizationAdminAccount example 123456789012
```
 