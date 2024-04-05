Manages a GuardDuty Organization Admin Account. The AWS account utilizing this resource must be an Organizations primary account. More information about Organizations support in GuardDuty can be found in the [GuardDuty User Guide](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOrganization = new aws.organizations.Organization("exampleOrganization", {
    awsServiceAccessPrincipals: ["guardduty.amazonaws.com"],
    featureSet: "ALL",
});
const exampleDetector = new aws.guardduty.Detector("exampleDetector", {});
const exampleOrganizationAdminAccount = new aws.guardduty.OrganizationAdminAccount("exampleOrganizationAdminAccount", {adminAccountId: "123456789012"}, {
    dependsOn: [exampleOrganization],
});
```
```python
import pulumi
import pulumi_aws as aws

example_organization = aws.organizations.Organization("exampleOrganization",
    aws_service_access_principals=["guardduty.amazonaws.com"],
    feature_set="ALL")
example_detector = aws.guardduty.Detector("exampleDetector")
example_organization_admin_account = aws.guardduty.OrganizationAdminAccount("exampleOrganizationAdminAccount", admin_account_id="123456789012",
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
            "guardduty.amazonaws.com",
        },
        FeatureSet = "ALL",
    });

    var exampleDetector = new Aws.GuardDuty.Detector("exampleDetector");

    var exampleOrganizationAdminAccount = new Aws.GuardDuty.OrganizationAdminAccount("exampleOrganizationAdminAccount", new()
    {
        AdminAccountId = "123456789012",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/guardduty"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("guardduty.amazonaws.com"),
			},
			FeatureSet: pulumi.String("ALL"),
		})
		if err != nil {
			return err
		}
		_, err = guardduty.NewDetector(ctx, "exampleDetector", nil)
		if err != nil {
			return err
		}
		_, err = guardduty.NewOrganizationAdminAccount(ctx, "exampleOrganizationAdminAccount", &guardduty.OrganizationAdminAccountArgs{
			AdminAccountId: pulumi.String("123456789012"),
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
import com.pulumi.aws.guardduty.Detector;
import com.pulumi.aws.guardduty.OrganizationAdminAccount;
import com.pulumi.aws.guardduty.OrganizationAdminAccountArgs;
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
            .awsServiceAccessPrincipals("guardduty.amazonaws.com")
            .featureSet("ALL")
            .build());

        var exampleDetector = new Detector("exampleDetector");

        var exampleOrganizationAdminAccount = new OrganizationAdminAccount("exampleOrganizationAdminAccount", OrganizationAdminAccountArgs.builder()        
            .adminAccountId("123456789012")
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
        - guardduty.amazonaws.com
      featureSet: ALL
  exampleDetector:
    type: aws:guardduty:Detector
  exampleOrganizationAdminAccount:
    type: aws:guardduty:OrganizationAdminAccount
    properties:
      adminAccountId: '123456789012'
    options:
      dependson:
        - ${exampleOrganization}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GuardDuty Organization Admin Account using the AWS account ID. For example:

```sh
 $ pulumi import aws:guardduty/organizationAdminAccount:OrganizationAdminAccount example 123456789012
```
 