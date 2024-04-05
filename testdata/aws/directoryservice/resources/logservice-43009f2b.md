Provides a Log subscription for AWS Directory Service that pushes logs to cloudwatch.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {retentionInDays: 14});
const ad-log-policyPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        actions: [
            "logs:CreateLogStream",
            "logs:PutLogEvents",
        ],
        principals: [{
            identifiers: ["ds.amazonaws.com"],
            type: "Service",
        }],
        resources: [pulumi.interpolate`${exampleLogGroup.arn}:*`],
        effect: "Allow",
    }],
});
const ad_log_policyLogResourcePolicy = new aws.cloudwatch.LogResourcePolicy("ad-log-policyLogResourcePolicy", {
    policyDocument: ad_log_policyPolicyDocument.apply(ad_log_policyPolicyDocument => ad_log_policyPolicyDocument.json),
    policyName: "ad-log-policy",
});
const exampleLogService = new aws.directoryservice.LogService("exampleLogService", {
    directoryId: aws_directory_service_directory.example.id,
    logGroupName: exampleLogGroup.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup", retention_in_days=14)
ad_log_policy_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=[
        "logs:CreateLogStream",
        "logs:PutLogEvents",
    ],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        identifiers=["ds.amazonaws.com"],
        type="Service",
    )],
    resources=[example_log_group.arn.apply(lambda arn: f"{arn}:*")],
    effect="Allow",
)])
ad_log_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("ad-log-policyLogResourcePolicy",
    policy_document=ad_log_policy_policy_document.json,
    policy_name="ad-log-policy")
example_log_service = aws.directoryservice.LogService("exampleLogService",
    directory_id=aws_directory_service_directory["example"]["id"],
    log_group_name=example_log_group.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new()
    {
        RetentionInDays = 14,
    });

    var ad_log_policyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "logs:CreateLogStream",
                    "logs:PutLogEvents",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Identifiers = new[]
                        {
                            "ds.amazonaws.com",
                        },
                        Type = "Service",
                    },
                },
                Resources = new[]
                {
                    $"{exampleLogGroup.Arn}:*",
                },
                Effect = "Allow",
            },
        },
    });

    var ad_log_policyLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("ad-log-policyLogResourcePolicy", new()
    {
        PolicyDocument = ad_log_policyPolicyDocument.Apply(ad_log_policyPolicyDocument => ad_log_policyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json)),
        PolicyName = "ad-log-policy",
    });

    var exampleLogService = new Aws.DirectoryService.LogService("exampleLogService", new()
    {
        DirectoryId = aws_directory_service_directory.Example.Id,
        LogGroupName = exampleLogGroup.Name,
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", &cloudwatch.LogGroupArgs{
			RetentionInDays: pulumi.Int(14),
		})
		if err != nil {
			return err
		}
		ad_log_policyPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Actions: pulumi.StringArray{
						pulumi.String("logs:CreateLogStream"),
						pulumi.String("logs:PutLogEvents"),
					},
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Identifiers: pulumi.StringArray{
								pulumi.String("ds.amazonaws.com"),
							},
							Type: pulumi.String("Service"),
						},
					},
					Resources: pulumi.StringArray{
						exampleLogGroup.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v:*", arn), nil
						}).(pulumi.StringOutput),
					},
					Effect: pulumi.String("Allow"),
				},
			},
		}, nil)
		_, err = cloudwatch.NewLogResourcePolicy(ctx, "ad-log-policyLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
			PolicyDocument: ad_log_policyPolicyDocument.ApplyT(func(ad_log_policyPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &ad_log_policyPolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
			PolicyName: pulumi.String("ad-log-policy"),
		})
		if err != nil {
			return err
		}
		_, err = directoryservice.NewLogService(ctx, "exampleLogService", &directoryservice.LogServiceArgs{
			DirectoryId:  pulumi.Any(aws_directory_service_directory.Example.Id),
			LogGroupName: exampleLogGroup.Name,
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.cloudwatch.LogGroupArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.cloudwatch.LogResourcePolicy;
import com.pulumi.aws.cloudwatch.LogResourcePolicyArgs;
import com.pulumi.aws.directoryservice.LogService;
import com.pulumi.aws.directoryservice.LogServiceArgs;
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
        var exampleLogGroup = new LogGroup("exampleLogGroup", LogGroupArgs.builder()        
            .retentionInDays(14)
            .build());

        final var ad-log-policyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions(                
                    "logs:CreateLogStream",
                    "logs:PutLogEvents")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .identifiers("ds.amazonaws.com")
                    .type("Service")
                    .build())
                .resources(exampleLogGroup.arn().applyValue(arn -> String.format("%s:*", arn)))
                .effect("Allow")
                .build())
            .build());

        var ad_log_policyLogResourcePolicy = new LogResourcePolicy("ad-log-policyLogResourcePolicy", LogResourcePolicyArgs.builder()        
            .policyDocument(ad_log_policyPolicyDocument.applyValue(ad_log_policyPolicyDocument -> ad_log_policyPolicyDocument.json()))
            .policyName("ad-log-policy")
            .build());

        var exampleLogService = new LogService("exampleLogService", LogServiceArgs.builder()        
            .directoryId(aws_directory_service_directory.example().id())
            .logGroupName(exampleLogGroup.name())
            .build());

    }
}
```
```yaml
resources:
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
    properties:
      retentionInDays: 14
  ad-log-policyLogResourcePolicy:
    type: aws:cloudwatch:LogResourcePolicy
    properties:
      policyDocument: ${["ad-log-policyPolicyDocument"].json}
      policyName: ad-log-policy
  exampleLogService:
    type: aws:directoryservice:LogService
    properties:
      directoryId: ${aws_directory_service_directory.example.id}
      logGroupName: ${exampleLogGroup.name}
variables:
  ad-log-policyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - logs:CreateLogStream
              - logs:PutLogEvents
            principals:
              - identifiers:
                  - ds.amazonaws.com
                type: Service
            resources:
              - ${exampleLogGroup.arn}:*
            effect: Allow
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Directory Service Log Subscriptions using the directory id. For example:

```sh
 $ pulumi import aws:directoryservice/logService:LogService msad d-1234567890
```
 