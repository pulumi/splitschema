Provides a CE Cost Category.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.costexplorer.CostCategory("test", {
    rules: [
        {
            rule: {
                dimension: {
                    key: "LINKED_ACCOUNT_NAME",
                    matchOptions: ["ENDS_WITH"],
                    values: ["-prod"],
                },
            },
            value: "production",
        },
        {
            rule: {
                dimension: {
                    key: "LINKED_ACCOUNT_NAME",
                    matchOptions: ["ENDS_WITH"],
                    values: ["-stg"],
                },
            },
            value: "staging",
        },
        {
            rule: {
                dimension: {
                    key: "LINKED_ACCOUNT_NAME",
                    matchOptions: ["ENDS_WITH"],
                    values: ["-dev"],
                },
            },
            value: "testing",
        },
    ],
    ruleVersion: "CostCategoryExpression.v1",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.costexplorer.CostCategory("test",
    rules=[
        aws.costexplorer.CostCategoryRuleArgs(
            rule=aws.costexplorer.CostCategoryRuleRuleArgs(
                dimension=aws.costexplorer.CostCategoryRuleRuleDimensionArgs(
                    key="LINKED_ACCOUNT_NAME",
                    match_options=["ENDS_WITH"],
                    values=["-prod"],
                ),
            ),
            value="production",
        ),
        aws.costexplorer.CostCategoryRuleArgs(
            rule=aws.costexplorer.CostCategoryRuleRuleArgs(
                dimension=aws.costexplorer.CostCategoryRuleRuleDimensionArgs(
                    key="LINKED_ACCOUNT_NAME",
                    match_options=["ENDS_WITH"],
                    values=["-stg"],
                ),
            ),
            value="staging",
        ),
        aws.costexplorer.CostCategoryRuleArgs(
            rule=aws.costexplorer.CostCategoryRuleRuleArgs(
                dimension=aws.costexplorer.CostCategoryRuleRuleDimensionArgs(
                    key="LINKED_ACCOUNT_NAME",
                    match_options=["ENDS_WITH"],
                    values=["-dev"],
                ),
            ),
            value="testing",
        ),
    ],
    rule_version="CostCategoryExpression.v1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.CostExplorer.CostCategory("test", new()
    {
        Rules = new[]
        {
            new Aws.CostExplorer.Inputs.CostCategoryRuleArgs
            {
                Rule = new Aws.CostExplorer.Inputs.CostCategoryRuleRuleArgs
                {
                    Dimension = new Aws.CostExplorer.Inputs.CostCategoryRuleRuleDimensionArgs
                    {
                        Key = "LINKED_ACCOUNT_NAME",
                        MatchOptions = new[]
                        {
                            "ENDS_WITH",
                        },
                        Values = new[]
                        {
                            "-prod",
                        },
                    },
                },
                Value = "production",
            },
            new Aws.CostExplorer.Inputs.CostCategoryRuleArgs
            {
                Rule = new Aws.CostExplorer.Inputs.CostCategoryRuleRuleArgs
                {
                    Dimension = new Aws.CostExplorer.Inputs.CostCategoryRuleRuleDimensionArgs
                    {
                        Key = "LINKED_ACCOUNT_NAME",
                        MatchOptions = new[]
                        {
                            "ENDS_WITH",
                        },
                        Values = new[]
                        {
                            "-stg",
                        },
                    },
                },
                Value = "staging",
            },
            new Aws.CostExplorer.Inputs.CostCategoryRuleArgs
            {
                Rule = new Aws.CostExplorer.Inputs.CostCategoryRuleRuleArgs
                {
                    Dimension = new Aws.CostExplorer.Inputs.CostCategoryRuleRuleDimensionArgs
                    {
                        Key = "LINKED_ACCOUNT_NAME",
                        MatchOptions = new[]
                        {
                            "ENDS_WITH",
                        },
                        Values = new[]
                        {
                            "-dev",
                        },
                    },
                },
                Value = "testing",
            },
        },
        RuleVersion = "CostCategoryExpression.v1",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/costexplorer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costexplorer.NewCostCategory(ctx, "test", &costexplorer.CostCategoryArgs{
			Rules: costexplorer.CostCategoryRuleArray{
				&costexplorer.CostCategoryRuleArgs{
					Rule: &costexplorer.CostCategoryRuleRuleArgs{
						Dimension: &costexplorer.CostCategoryRuleRuleDimensionArgs{
							Key: pulumi.String("LINKED_ACCOUNT_NAME"),
							MatchOptions: pulumi.StringArray{
								pulumi.String("ENDS_WITH"),
							},
							Values: pulumi.StringArray{
								pulumi.String("-prod"),
							},
						},
					},
					Value: pulumi.String("production"),
				},
				&costexplorer.CostCategoryRuleArgs{
					Rule: &costexplorer.CostCategoryRuleRuleArgs{
						Dimension: &costexplorer.CostCategoryRuleRuleDimensionArgs{
							Key: pulumi.String("LINKED_ACCOUNT_NAME"),
							MatchOptions: pulumi.StringArray{
								pulumi.String("ENDS_WITH"),
							},
							Values: pulumi.StringArray{
								pulumi.String("-stg"),
							},
						},
					},
					Value: pulumi.String("staging"),
				},
				&costexplorer.CostCategoryRuleArgs{
					Rule: &costexplorer.CostCategoryRuleRuleArgs{
						Dimension: &costexplorer.CostCategoryRuleRuleDimensionArgs{
							Key: pulumi.String("LINKED_ACCOUNT_NAME"),
							MatchOptions: pulumi.StringArray{
								pulumi.String("ENDS_WITH"),
							},
							Values: pulumi.StringArray{
								pulumi.String("-dev"),
							},
						},
					},
					Value: pulumi.String("testing"),
				},
			},
			RuleVersion: pulumi.String("CostCategoryExpression.v1"),
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
import com.pulumi.aws.costexplorer.CostCategory;
import com.pulumi.aws.costexplorer.CostCategoryArgs;
import com.pulumi.aws.costexplorer.inputs.CostCategoryRuleArgs;
import com.pulumi.aws.costexplorer.inputs.CostCategoryRuleRuleArgs;
import com.pulumi.aws.costexplorer.inputs.CostCategoryRuleRuleDimensionArgs;
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
        var test = new CostCategory("test", CostCategoryArgs.builder()        
            .rules(            
                CostCategoryRuleArgs.builder()
                    .rule(CostCategoryRuleRuleArgs.builder()
                        .dimension(CostCategoryRuleRuleDimensionArgs.builder()
                            .key("LINKED_ACCOUNT_NAME")
                            .matchOptions("ENDS_WITH")
                            .values("-prod")
                            .build())
                        .build())
                    .value("production")
                    .build(),
                CostCategoryRuleArgs.builder()
                    .rule(CostCategoryRuleRuleArgs.builder()
                        .dimension(CostCategoryRuleRuleDimensionArgs.builder()
                            .key("LINKED_ACCOUNT_NAME")
                            .matchOptions("ENDS_WITH")
                            .values("-stg")
                            .build())
                        .build())
                    .value("staging")
                    .build(),
                CostCategoryRuleArgs.builder()
                    .rule(CostCategoryRuleRuleArgs.builder()
                        .dimension(CostCategoryRuleRuleDimensionArgs.builder()
                            .key("LINKED_ACCOUNT_NAME")
                            .matchOptions("ENDS_WITH")
                            .values("-dev")
                            .build())
                        .build())
                    .value("testing")
                    .build())
            .ruleVersion("CostCategoryExpression.v1")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:costexplorer:CostCategory
    properties:
      rules:
        - rule:
            dimension:
              key: LINKED_ACCOUNT_NAME
              matchOptions:
                - ENDS_WITH
              values:
                - -prod
          value: production
        - rule:
            dimension:
              key: LINKED_ACCOUNT_NAME
              matchOptions:
                - ENDS_WITH
              values:
                - -stg
          value: staging
        - rule:
            dimension:
              key: LINKED_ACCOUNT_NAME
              matchOptions:
                - ENDS_WITH
              values:
                - -dev
          value: testing
      ruleVersion: CostCategoryExpression.v1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ce_cost_category` using the id. For example:

```sh
 $ pulumi import aws:costexplorer/costCategory:CostCategory example costCategoryARN
```
 