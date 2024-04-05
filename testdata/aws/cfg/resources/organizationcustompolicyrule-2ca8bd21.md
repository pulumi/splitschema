Manages a Config Organization Custom Policy Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Managed Rules (those invoking an AWS managed rule), see the `aws_config_organization_managed__rule` resource.

> **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excluded_accounts` argument.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cfg.OrganizationCustomPolicyRule("example", {
    policyRuntime: "guard-2.x.x",
    policyText: `  let status = ['ACTIVE']

  rule tableisactive when
      resourceType == "AWS::DynamoDB::Table" {
      configuration.tableStatus == %status
  }

  rule checkcompliance when
      resourceType == "AWS::DynamoDB::Table"
      tableisactive {
          let pitr = supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus
          %pitr == "ENABLED"
      }

`,
    resourceTypesScopes: ["AWS::DynamoDB::Table"],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cfg.OrganizationCustomPolicyRule("example",
    policy_runtime="guard-2.x.x",
    policy_text="""  let status = ['ACTIVE']

  rule tableisactive when
      resourceType == "AWS::DynamoDB::Table" {
      configuration.tableStatus == %status
  }

  rule checkcompliance when
      resourceType == "AWS::DynamoDB::Table"
      tableisactive {
          let pitr = supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus
          %pitr == "ENABLED"
      }

""",
    resource_types_scopes=["AWS::DynamoDB::Table"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Cfg.OrganizationCustomPolicyRule("example", new()
    {
        PolicyRuntime = "guard-2.x.x",
        PolicyText = @"  let status = ['ACTIVE']

  rule tableisactive when
      resourceType == ""AWS::DynamoDB::Table"" {
      configuration.tableStatus == %status
  }

  rule checkcompliance when
      resourceType == ""AWS::DynamoDB::Table""
      tableisactive {
          let pitr = supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus
          %pitr == ""ENABLED""
      }

",
        ResourceTypesScopes = new[]
        {
            "AWS::DynamoDB::Table",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cfg.NewOrganizationCustomPolicyRule(ctx, "example", &cfg.OrganizationCustomPolicyRuleArgs{
			PolicyRuntime: pulumi.String("guard-2.x.x"),
			PolicyText: pulumi.String(`  let status = ['ACTIVE']

  rule tableisactive when
      resourceType == "AWS::DynamoDB::Table" {
      configuration.tableStatus == %status
  }

  rule checkcompliance when
      resourceType == "AWS::DynamoDB::Table"
      tableisactive {
          let pitr = supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus
          %pitr == "ENABLED"
      }

`),
			ResourceTypesScopes: pulumi.StringArray{
				pulumi.String("AWS::DynamoDB::Table"),
			},
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
import com.pulumi.aws.cfg.OrganizationCustomPolicyRule;
import com.pulumi.aws.cfg.OrganizationCustomPolicyRuleArgs;
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
        var example = new OrganizationCustomPolicyRule("example", OrganizationCustomPolicyRuleArgs.builder()        
            .policyRuntime("guard-2.x.x")
            .policyText("""
  let status = ['ACTIVE']

  rule tableisactive when
      resourceType == "AWS::DynamoDB::Table" {
      configuration.tableStatus == %status
  }

  rule checkcompliance when
      resourceType == "AWS::DynamoDB::Table"
      tableisactive {
          let pitr = supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus
          %pitr == "ENABLED"
      }

            """)
            .resourceTypesScopes("AWS::DynamoDB::Table")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cfg:OrganizationCustomPolicyRule
    properties:
      policyRuntime: guard-2.x.x
      policyText: |2+
          let status = ['ACTIVE']

          rule tableisactive when
              resourceType == "AWS::DynamoDB::Table" {
              configuration.tableStatus == %status
          }

          rule checkcompliance when
              resourceType == "AWS::DynamoDB::Table"
              tableisactive {
                  let pitr = supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus
                  %pitr == "ENABLED"
              }

      resourceTypesScopes:
        - AWS::DynamoDB::Table
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a Config Organization Custom Policy Rule using the `name` argument. For example:

```sh
 $ pulumi import aws:cfg/organizationCustomPolicyRule:OrganizationCustomPolicyRule example example_rule_name
```
 