Manages a Config Organization Custom Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Managed Rules (those invoking an AWS managed rule), see the `aws_config_organization_managed__rule` resource.

> **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excluded_accounts` argument.

> **NOTE:** The proper Lambda permission to allow the AWS Config service invoke the Lambda Function must be in place before the rule will successfully create or update. See also the `aws.lambda.Permission` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePermission = new aws.lambda.Permission("examplePermission", {
    action: "lambda:InvokeFunction",
    "function": aws_lambda_function.example.arn,
    principal: "config.amazonaws.com",
});
const exampleOrganization = new aws.organizations.Organization("exampleOrganization", {
    awsServiceAccessPrincipals: ["config-multiaccountsetup.amazonaws.com"],
    featureSet: "ALL",
});
const exampleOrganizationCustomRule = new aws.cfg.OrganizationCustomRule("exampleOrganizationCustomRule", {
    lambdaFunctionArn: aws_lambda_function.example.arn,
    triggerTypes: ["ConfigurationItemChangeNotification"],
}, {
    dependsOn: [
        examplePermission,
        exampleOrganization,
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_permission = aws.lambda_.Permission("examplePermission",
    action="lambda:InvokeFunction",
    function=aws_lambda_function["example"]["arn"],
    principal="config.amazonaws.com")
example_organization = aws.organizations.Organization("exampleOrganization",
    aws_service_access_principals=["config-multiaccountsetup.amazonaws.com"],
    feature_set="ALL")
example_organization_custom_rule = aws.cfg.OrganizationCustomRule("exampleOrganizationCustomRule",
    lambda_function_arn=aws_lambda_function["example"]["arn"],
    trigger_types=["ConfigurationItemChangeNotification"],
    opts=pulumi.ResourceOptions(depends_on=[
            example_permission,
            example_organization,
        ]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplePermission = new Aws.Lambda.Permission("examplePermission", new()
    {
        Action = "lambda:InvokeFunction",
        Function = aws_lambda_function.Example.Arn,
        Principal = "config.amazonaws.com",
    });

    var exampleOrganization = new Aws.Organizations.Organization("exampleOrganization", new()
    {
        AwsServiceAccessPrincipals = new[]
        {
            "config-multiaccountsetup.amazonaws.com",
        },
        FeatureSet = "ALL",
    });

    var exampleOrganizationCustomRule = new Aws.Cfg.OrganizationCustomRule("exampleOrganizationCustomRule", new()
    {
        LambdaFunctionArn = aws_lambda_function.Example.Arn,
        TriggerTypes = new[]
        {
            "ConfigurationItemChangeNotification",
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            examplePermission,
            exampleOrganization,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplePermission, err := lambda.NewPermission(ctx, "examplePermission", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  pulumi.Any(aws_lambda_function.Example.Arn),
			Principal: pulumi.String("config.amazonaws.com"),
		})
		if err != nil {
			return err
		}
		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("config-multiaccountsetup.amazonaws.com"),
			},
			FeatureSet: pulumi.String("ALL"),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewOrganizationCustomRule(ctx, "exampleOrganizationCustomRule", &cfg.OrganizationCustomRuleArgs{
			LambdaFunctionArn: pulumi.Any(aws_lambda_function.Example.Arn),
			TriggerTypes: pulumi.StringArray{
				pulumi.String("ConfigurationItemChangeNotification"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			examplePermission,
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
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
import com.pulumi.aws.organizations.Organization;
import com.pulumi.aws.organizations.OrganizationArgs;
import com.pulumi.aws.cfg.OrganizationCustomRule;
import com.pulumi.aws.cfg.OrganizationCustomRuleArgs;
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
        var examplePermission = new Permission("examplePermission", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function(aws_lambda_function.example().arn())
            .principal("config.amazonaws.com")
            .build());

        var exampleOrganization = new Organization("exampleOrganization", OrganizationArgs.builder()        
            .awsServiceAccessPrincipals("config-multiaccountsetup.amazonaws.com")
            .featureSet("ALL")
            .build());

        var exampleOrganizationCustomRule = new OrganizationCustomRule("exampleOrganizationCustomRule", OrganizationCustomRuleArgs.builder()        
            .lambdaFunctionArn(aws_lambda_function.example().arn())
            .triggerTypes("ConfigurationItemChangeNotification")
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    examplePermission,
                    exampleOrganization)
                .build());

    }
}
```
```yaml
resources:
  examplePermission:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: ${aws_lambda_function.example.arn}
      principal: config.amazonaws.com
  exampleOrganization:
    type: aws:organizations:Organization
    properties:
      awsServiceAccessPrincipals:
        - config-multiaccountsetup.amazonaws.com
      featureSet: ALL
  exampleOrganizationCustomRule:
    type: aws:cfg:OrganizationCustomRule
    properties:
      lambdaFunctionArn: ${aws_lambda_function.example.arn}
      triggerTypes:
        - ConfigurationItemChangeNotification
    options:
      dependson:
        - ${examplePermission}
        - ${exampleOrganization}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Config Organization Custom Rules using the name. For example:

```sh
 $ pulumi import aws:cfg/organizationCustomRule:OrganizationCustomRule example example
```
 