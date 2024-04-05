Provides a Glue resource policy. Only one can exist per region.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentCallerIdentity = aws.getCallerIdentity({});
const currentPartition = aws.getPartition({});
const currentRegion = aws.getRegion({});
const glue-example-policy = Promise.all([currentPartition, currentRegion, currentCallerIdentity]).then(([currentPartition, currentRegion, currentCallerIdentity]) => aws.iam.getPolicyDocument({
    statements: [{
        actions: ["glue:CreateTable"],
        resources: [`arn:${currentPartition.partition}:glue:${currentRegion.name}:${currentCallerIdentity.accountId}:*`],
        principals: [{
            identifiers: ["*"],
            type: "AWS",
        }],
    }],
}));
const example = new aws.glue.ResourcePolicy("example", {policy: glue_example_policy.then(glue_example_policy => glue_example_policy.json)});
```
```python
import pulumi
import pulumi_aws as aws

current_caller_identity = aws.get_caller_identity()
current_partition = aws.get_partition()
current_region = aws.get_region()
glue_example_policy = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["glue:CreateTable"],
    resources=[f"arn:{current_partition.partition}:glue:{current_region.name}:{current_caller_identity.account_id}:*"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        identifiers=["*"],
        type="AWS",
    )],
)])
example = aws.glue.ResourcePolicy("example", policy=glue_example_policy.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var currentPartition = Aws.GetPartition.Invoke();

    var currentRegion = Aws.GetRegion.Invoke();

    var glue_example_policy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "glue:CreateTable",
                },
                Resources = new[]
                {
                    $"arn:{currentPartition.Apply(getPartitionResult => getPartitionResult.Partition)}:glue:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:*",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Identifiers = new[]
                        {
                            "*",
                        },
                        Type = "AWS",
                    },
                },
            },
        },
    });

    var example = new Aws.Glue.ResourcePolicy("example", new()
    {
        Policy = glue_example_policy.Apply(glue_example_policy => glue_example_policy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json)),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentPartition, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		glue_example_policy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"glue:CreateTable",
					},
					Resources: []string{
						fmt.Sprintf("arn:%v:glue:%v:%v:*", currentPartition.Partition, currentRegion.Name, currentCallerIdentity.AccountId),
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Identifiers: []string{
								"*",
							},
							Type: "AWS",
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = glue.NewResourcePolicy(ctx, "example", &glue.ResourcePolicyArgs{
			Policy: *pulumi.String(glue_example_policy.Json),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.glue.ResourcePolicy;
import com.pulumi.aws.glue.ResourcePolicyArgs;
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
        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        final var currentPartition = AwsFunctions.getPartition();

        final var currentRegion = AwsFunctions.getRegion();

        final var glue-example-policy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("glue:CreateTable")
                .resources(String.format("arn:%s:glue:%s:%s:*", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .identifiers("*")
                    .type("AWS")
                    .build())
                .build())
            .build());

        var example = new ResourcePolicy("example", ResourcePolicyArgs.builder()        
            .policy(glue_example_policy.json())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:ResourcePolicy
    properties:
      policy: ${["glue-example-policy"].json}
variables:
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  currentPartition:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  glue-example-policy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - glue:CreateTable
            resources:
              - arn:${currentPartition.partition}:glue:${currentRegion.name}:${currentCallerIdentity.accountId}:*
            principals:
              - identifiers:
                  - '*'
                type: AWS
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Resource Policy using the account ID. For example:

```sh
 $ pulumi import aws:glue/resourcePolicy:ResourcePolicy Test 12356789012
```
 