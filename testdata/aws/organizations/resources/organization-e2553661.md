Provides a resource to create an organization.

!> **WARNING:** When migrating from a `feature_set` of `CONSOLIDATED_BILLING` to `ALL`, the Organization account owner will received an email stating the following: "You started the process to enable all features for your AWS organization. As part of that process, all member accounts that joined your organization by invitation must approve the change. You don’t need approval from member accounts that you directly created from within your AWS organization." After all member accounts have accepted the invitation, the Organization account owner must then finalize the changes via the [AWS Console](https://console.aws.amazon.com/organizations/home#/organization/settings/migration-progress). Until these steps are performed, the provider will perpetually show a difference, and the `DescribeOrganization` API will continue to show the `FeatureSet` as `CONSOLIDATED_BILLING`. See the [AWS Organizations documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html) for more information.

!> **WARNING:** [Warning from the AWS Docs](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnableAWSServiceAccess.html): "We recommend that you enable integration between AWS Organizations and the specified AWS service by using the console or commands that are provided by the specified service. Doing so ensures that the service is aware that it can create the resources that are required for the integration. How the service creates those resources in the organization's accounts depends on that service. For more information, see the documentation for the other AWS service."

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const org = new aws.organizations.Organization("org", {
    awsServiceAccessPrincipals: [
        "cloudtrail.amazonaws.com",
        "config.amazonaws.com",
    ],
    featureSet: "ALL",
});
```
```python
import pulumi
import pulumi_aws as aws

org = aws.organizations.Organization("org",
    aws_service_access_principals=[
        "cloudtrail.amazonaws.com",
        "config.amazonaws.com",
    ],
    feature_set="ALL")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var org = new Aws.Organizations.Organization("org", new()
    {
        AwsServiceAccessPrincipals = new[]
        {
            "cloudtrail.amazonaws.com",
            "config.amazonaws.com",
        },
        FeatureSet = "ALL",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewOrganization(ctx, "org", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("cloudtrail.amazonaws.com"),
				pulumi.String("config.amazonaws.com"),
			},
			FeatureSet: pulumi.String("ALL"),
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
        var org = new Organization("org", OrganizationArgs.builder()        
            .awsServiceAccessPrincipals(            
                "cloudtrail.amazonaws.com",
                "config.amazonaws.com")
            .featureSet("ALL")
            .build());

    }
}
```
```yaml
resources:
  org:
    type: aws:organizations:Organization
    properties:
      awsServiceAccessPrincipals:
        - cloudtrail.amazonaws.com
        - config.amazonaws.com
      featureSet: ALL
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the AWS organization using the `id`. For example:

```sh
 $ pulumi import aws:organizations/organization:Organization my_org o-1234567
```
 