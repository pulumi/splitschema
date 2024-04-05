Manages a Config Organization Managed Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Custom Rules (those invoking a custom Lambda Function), see the `aws.cfg.OrganizationCustomRule` resource.

> **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excluded_accounts` argument.

> **NOTE:** Every Organization account except those configured in the `excluded_accounts` argument must have a Configuration Recorder with proper IAM permissions before the rule will successfully create or update. See also the `aws.cfg.Recorder` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOrganization = new aws.organizations.Organization("exampleOrganization", {
    awsServiceAccessPrincipals: ["config-multiaccountsetup.amazonaws.com"],
    featureSet: "ALL",
});
const exampleOrganizationManagedRule = new aws.cfg.OrganizationManagedRule("exampleOrganizationManagedRule", {ruleIdentifier: "IAM_PASSWORD_POLICY"}, {
    dependsOn: [exampleOrganization],
});
```
```python
import pulumi
import pulumi_aws as aws

example_organization = aws.organizations.Organization("exampleOrganization",
    aws_service_access_principals=["config-multiaccountsetup.amazonaws.com"],
    feature_set="ALL")
example_organization_managed_rule = aws.cfg.OrganizationManagedRule("exampleOrganizationManagedRule", rule_identifier="IAM_PASSWORD_POLICY",
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
            "config-multiaccountsetup.amazonaws.com",
        },
        FeatureSet = "ALL",
    });

    var exampleOrganizationManagedRule = new Aws.Cfg.OrganizationManagedRule("exampleOrganizationManagedRule", new()
    {
        RuleIdentifier = "IAM_PASSWORD_POLICY",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("config-multiaccountsetup.amazonaws.com"),
			},
			FeatureSet: pulumi.String("ALL"),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewOrganizationManagedRule(ctx, "exampleOrganizationManagedRule", &cfg.OrganizationManagedRuleArgs{
			RuleIdentifier: pulumi.String("IAM_PASSWORD_POLICY"),
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
import com.pulumi.aws.cfg.OrganizationManagedRule;
import com.pulumi.aws.cfg.OrganizationManagedRuleArgs;
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
            .awsServiceAccessPrincipals("config-multiaccountsetup.amazonaws.com")
            .featureSet("ALL")
            .build());

        var exampleOrganizationManagedRule = new OrganizationManagedRule("exampleOrganizationManagedRule", OrganizationManagedRuleArgs.builder()        
            .ruleIdentifier("IAM_PASSWORD_POLICY")
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
        - config-multiaccountsetup.amazonaws.com
      featureSet: ALL
  exampleOrganizationManagedRule:
    type: aws:cfg:OrganizationManagedRule
    properties:
      ruleIdentifier: IAM_PASSWORD_POLICY
    options:
      dependson:
        - ${exampleOrganization}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Config Organization Managed Rules using the name. For example:

```sh
 $ pulumi import aws:cfg/organizationManagedRule:OrganizationManagedRule example example
```
 