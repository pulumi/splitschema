Manages a Detective Organization Admin Account. The AWS account utilizing this resource must be an Organizations primary account. More information about Organizations support in Detective can be found in the [Detective User Guide](https://docs.aws.amazon.com/detective/latest/adminguide/accounts-orgs-transition.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOrganization = new aws.organizations.Organization("exampleOrganization", {
    awsServiceAccessPrincipals: ["detective.amazonaws.com"],
    featureSet: "ALL",
});
const exampleOrganizationAdminAccount = new aws.detective.OrganizationAdminAccount("exampleOrganizationAdminAccount", {accountId: "123456789012"}, {
    dependsOn: [exampleOrganization],
});
```
```python
import pulumi
import pulumi_aws as aws

example_organization = aws.organizations.Organization("exampleOrganization",
    aws_service_access_principals=["detective.amazonaws.com"],
    feature_set="ALL")
example_organization_admin_account = aws.detective.OrganizationAdminAccount("exampleOrganizationAdminAccount", account_id="123456789012",
opts=pulumi.ResourceOptions(depends_on=[example_organization]))
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
            "detective.amazonaws.com",
        },
        FeatureSet = "ALL",
    });

    var exampleOrganizationAdminAccount = new Aws.Detective.OrganizationAdminAccount("exampleOrganizationAdminAccount", new()
    {
        AccountId = "123456789012",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleOrganization,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/detective"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("detective.amazonaws.com"),
			},
			FeatureSet: pulumi.String("ALL"),
		})
		if err != nil {
			return err
		}
		_, err = detective.NewOrganizationAdminAccount(ctx, "exampleOrganizationAdminAccount", &detective.OrganizationAdminAccountArgs{
			AccountId: pulumi.String("123456789012"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleOrganization,
		}))
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
import com.pulumi.aws.detective.OrganizationAdminAccount;
import com.pulumi.aws.detective.OrganizationAdminAccountArgs;
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
            .awsServiceAccessPrincipals("detective.amazonaws.com")
            .featureSet("ALL")
            .build());

        var exampleOrganizationAdminAccount = new OrganizationAdminAccount("exampleOrganizationAdminAccount", OrganizationAdminAccountArgs.builder()        
            .accountId("123456789012")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleOrganization)
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
        - detective.amazonaws.com
      featureSet: ALL
  exampleOrganizationAdminAccount:
    type: aws:detective:OrganizationAdminAccount
    properties:
      accountId: '123456789012'
    options:
      dependson:
        - ${exampleOrganization}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_detective_organization_admin_account` using `account_id`. For example:

```sh
 $ pulumi import aws:detective/organizationAdminAccount:OrganizationAdminAccount example 123456789012
```
 