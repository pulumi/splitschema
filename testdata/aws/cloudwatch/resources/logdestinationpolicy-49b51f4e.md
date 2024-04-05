Provides a CloudWatch Logs destination policy resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testDestination = new aws.cloudwatch.LogDestination("testDestination", {
    roleArn: aws_iam_role.iam_for_cloudwatch.arn,
    targetArn: aws_kinesis_stream.kinesis_for_cloudwatch.arn,
});
const testDestinationPolicyPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: ["123456789012"],
        }],
        actions: ["logs:PutSubscriptionFilter"],
        resources: [testDestination.arn],
    }],
});
const testDestinationPolicyLogDestinationPolicy = new aws.cloudwatch.LogDestinationPolicy("testDestinationPolicyLogDestinationPolicy", {
    destinationName: testDestination.name,
    accessPolicy: testDestinationPolicyPolicyDocument.apply(testDestinationPolicyPolicyDocument => testDestinationPolicyPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

test_destination = aws.cloudwatch.LogDestination("testDestination",
    role_arn=aws_iam_role["iam_for_cloudwatch"]["arn"],
    target_arn=aws_kinesis_stream["kinesis_for_cloudwatch"]["arn"])
test_destination_policy_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["123456789012"],
    )],
    actions=["logs:PutSubscriptionFilter"],
    resources=[test_destination.arn],
)])
test_destination_policy_log_destination_policy = aws.cloudwatch.LogDestinationPolicy("testDestinationPolicyLogDestinationPolicy",
    destination_name=test_destination.name,
    access_policy=test_destination_policy_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testDestination = new Aws.CloudWatch.LogDestination("testDestination", new()
    {
        RoleArn = aws_iam_role.Iam_for_cloudwatch.Arn,
        TargetArn = aws_kinesis_stream.Kinesis_for_cloudwatch.Arn,
    });

    var testDestinationPolicyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            "123456789012",
                        },
                    },
                },
                Actions = new[]
                {
                    "logs:PutSubscriptionFilter",
                },
                Resources = new[]
                {
                    testDestination.Arn,
                },
            },
        },
    });

    var testDestinationPolicyLogDestinationPolicy = new Aws.CloudWatch.LogDestinationPolicy("testDestinationPolicyLogDestinationPolicy", new()
    {
        DestinationName = testDestination.Name,
        AccessPolicy = testDestinationPolicyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testDestination, err := cloudwatch.NewLogDestination(ctx, "testDestination", &cloudwatch.LogDestinationArgs{
			RoleArn:   pulumi.Any(aws_iam_role.Iam_for_cloudwatch.Arn),
			TargetArn: pulumi.Any(aws_kinesis_stream.Kinesis_for_cloudwatch.Arn),
		})
		if err != nil {
			return err
		}
		testDestinationPolicyPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("AWS"),
							Identifiers: pulumi.StringArray{
								pulumi.String("123456789012"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("logs:PutSubscriptionFilter"),
					},
					Resources: pulumi.StringArray{
						testDestination.Arn,
					},
				},
			},
		}, nil)
		_, err = cloudwatch.NewLogDestinationPolicy(ctx, "testDestinationPolicyLogDestinationPolicy", &cloudwatch.LogDestinationPolicyArgs{
			DestinationName: testDestination.Name,
			AccessPolicy: testDestinationPolicyPolicyDocument.ApplyT(func(testDestinationPolicyPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &testDestinationPolicyPolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
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
import com.pulumi.aws.cloudwatch.LogDestination;
import com.pulumi.aws.cloudwatch.LogDestinationArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.cloudwatch.LogDestinationPolicy;
import com.pulumi.aws.cloudwatch.LogDestinationPolicyArgs;
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
        var testDestination = new LogDestination("testDestination", LogDestinationArgs.builder()        
            .roleArn(aws_iam_role.iam_for_cloudwatch().arn())
            .targetArn(aws_kinesis_stream.kinesis_for_cloudwatch().arn())
            .build());

        final var testDestinationPolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("123456789012")
                    .build())
                .actions("logs:PutSubscriptionFilter")
                .resources(testDestination.arn())
                .build())
            .build());

        var testDestinationPolicyLogDestinationPolicy = new LogDestinationPolicy("testDestinationPolicyLogDestinationPolicy", LogDestinationPolicyArgs.builder()        
            .destinationName(testDestination.name())
            .accessPolicy(testDestinationPolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(testDestinationPolicyPolicyDocument -> testDestinationPolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  testDestination:
    type: aws:cloudwatch:LogDestination
    properties:
      roleArn: ${aws_iam_role.iam_for_cloudwatch.arn}
      targetArn: ${aws_kinesis_stream.kinesis_for_cloudwatch.arn}
  testDestinationPolicyLogDestinationPolicy:
    type: aws:cloudwatch:LogDestinationPolicy
    properties:
      destinationName: ${testDestination.name}
      accessPolicy: ${testDestinationPolicyPolicyDocument.json}
variables:
  testDestinationPolicyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - '123456789012'
            actions:
              - logs:PutSubscriptionFilter
            resources:
              - ${testDestination.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch Logs destination policies using the `destination_name`. For example:

```sh
 $ pulumi import aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy test_destination_policy test_destination
```
 