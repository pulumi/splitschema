Provides a CodeBuild Resource Policy Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleReportGroup = new aws.codebuild.ReportGroup("exampleReportGroup", {
    type: "TEST",
    exportConfig: {
        type: "NO_EXPORT",
    },
});
const currentPartition = aws.getPartition({});
const currentCallerIdentity = aws.getCallerIdentity({});
const exampleResourcePolicy = new aws.codebuild.ResourcePolicy("exampleResourcePolicy", {
    resourceArn: exampleReportGroup.arn,
    policy: pulumi.all([currentPartition, currentCallerIdentity, exampleReportGroup.arn]).apply(([currentPartition, currentCallerIdentity, arn]) => JSON.stringify({
        Version: "2012-10-17",
        Id: "default",
        Statement: [{
            Sid: "default",
            Effect: "Allow",
            Principal: {
                AWS: `arn:${currentPartition.partition}:iam::${currentCallerIdentity.accountId}:root`,
            },
            Action: [
                "codebuild:BatchGetReportGroups",
                "codebuild:BatchGetReports",
                "codebuild:ListReportsForReportGroup",
                "codebuild:DescribeTestCases",
            ],
            Resource: arn,
        }],
    })),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_report_group = aws.codebuild.ReportGroup("exampleReportGroup",
    type="TEST",
    export_config=aws.codebuild.ReportGroupExportConfigArgs(
        type="NO_EXPORT",
    ))
current_partition = aws.get_partition()
current_caller_identity = aws.get_caller_identity()
example_resource_policy = aws.codebuild.ResourcePolicy("exampleResourcePolicy",
    resource_arn=example_report_group.arn,
    policy=example_report_group.arn.apply(lambda arn: json.dumps({
        "Version": "2012-10-17",
        "Id": "default",
        "Statement": [{
            "Sid": "default",
            "Effect": "Allow",
            "Principal": {
                "AWS": f"arn:{current_partition.partition}:iam::{current_caller_identity.account_id}:root",
            },
            "Action": [
                "codebuild:BatchGetReportGroups",
                "codebuild:BatchGetReports",
                "codebuild:ListReportsForReportGroup",
                "codebuild:DescribeTestCases",
            ],
            "Resource": arn,
        }],
    })))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleReportGroup = new Aws.CodeBuild.ReportGroup("exampleReportGroup", new()
    {
        Type = "TEST",
        ExportConfig = new Aws.CodeBuild.Inputs.ReportGroupExportConfigArgs
        {
            Type = "NO_EXPORT",
        },
    });

    var currentPartition = Aws.GetPartition.Invoke();

    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var exampleResourcePolicy = new Aws.CodeBuild.ResourcePolicy("exampleResourcePolicy", new()
    {
        ResourceArn = exampleReportGroup.Arn,
        Policy = Output.Tuple(currentPartition, currentCallerIdentity, exampleReportGroup.Arn).Apply(values =>
        {
            var currentPartition = values.Item1;
            var currentCallerIdentity = values.Item2;
            var arn = values.Item3;
            return JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["Version"] = "2012-10-17",
                ["Id"] = "default",
                ["Statement"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["Sid"] = "default",
                        ["Effect"] = "Allow",
                        ["Principal"] = new Dictionary<string, object?>
                        {
                            ["AWS"] = $"arn:{currentPartition.Apply(getPartitionResult => getPartitionResult.Partition)}:iam::{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:root",
                        },
                        ["Action"] = new[]
                        {
                            "codebuild:BatchGetReportGroups",
                            "codebuild:BatchGetReports",
                            "codebuild:ListReportsForReportGroup",
                            "codebuild:DescribeTestCases",
                        },
                        ["Resource"] = arn,
                    },
                },
            });
        }),
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codebuild"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleReportGroup, err := codebuild.NewReportGroup(ctx, "exampleReportGroup", &codebuild.ReportGroupArgs{
			Type: pulumi.String("TEST"),
			ExportConfig: &codebuild.ReportGroupExportConfigArgs{
				Type: pulumi.String("NO_EXPORT"),
			},
		})
		if err != nil {
			return err
		}
		currentPartition, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = codebuild.NewResourcePolicy(ctx, "exampleResourcePolicy", &codebuild.ResourcePolicyArgs{
			ResourceArn: exampleReportGroup.Arn,
			Policy: exampleReportGroup.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"Version": "2012-10-17",
					"Id":      "default",
					"Statement": []map[string]interface{}{
						map[string]interface{}{
							"Sid":    "default",
							"Effect": "Allow",
							"Principal": map[string]interface{}{
								"AWS": fmt.Sprintf("arn:%v:iam::%v:root", currentPartition.Partition, currentCallerIdentity.AccountId),
							},
							"Action": []string{
								"codebuild:BatchGetReportGroups",
								"codebuild:BatchGetReports",
								"codebuild:ListReportsForReportGroup",
								"codebuild:DescribeTestCases",
							},
							"Resource": arn,
						},
					},
				})
				if err != nil {
					return _zero, err
				}
				json0 := string(tmpJSON0)
				return pulumi.String(json0), nil
			}).(pulumi.StringOutput),
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
import com.pulumi.aws.codebuild.ReportGroup;
import com.pulumi.aws.codebuild.ReportGroupArgs;
import com.pulumi.aws.codebuild.inputs.ReportGroupExportConfigArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.codebuild.ResourcePolicy;
import com.pulumi.aws.codebuild.ResourcePolicyArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var exampleReportGroup = new ReportGroup("exampleReportGroup", ReportGroupArgs.builder()        
            .type("TEST")
            .exportConfig(ReportGroupExportConfigArgs.builder()
                .type("NO_EXPORT")
                .build())
            .build());

        final var currentPartition = AwsFunctions.getPartition();

        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        var exampleResourcePolicy = new ResourcePolicy("exampleResourcePolicy", ResourcePolicyArgs.builder()        
            .resourceArn(exampleReportGroup.arn())
            .policy(exampleReportGroup.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Id", "default"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Sid", "default"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", String.format("arn:%s:iam::%s:root", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                        )),
                        jsonProperty("Action", jsonArray(
                            "codebuild:BatchGetReportGroups", 
                            "codebuild:BatchGetReports", 
                            "codebuild:ListReportsForReportGroup", 
                            "codebuild:DescribeTestCases"
                        )),
                        jsonProperty("Resource", arn)
                    )))
                ))))
            .build());

    }
}
```
```yaml
resources:
  exampleReportGroup:
    type: aws:codebuild:ReportGroup
    properties:
      type: TEST
      exportConfig:
        type: NO_EXPORT
  exampleResourcePolicy:
    type: aws:codebuild:ResourcePolicy
    properties:
      resourceArn: ${exampleReportGroup.arn}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Id: default
          Statement:
            - Sid: default
              Effect: Allow
              Principal:
                AWS: arn:${currentPartition.partition}:iam::${currentCallerIdentity.accountId}:root
              Action:
                - codebuild:BatchGetReportGroups
                - codebuild:BatchGetReports
                - codebuild:ListReportsForReportGroup
                - codebuild:DescribeTestCases
              Resource: ${exampleReportGroup.arn}
variables:
  currentPartition:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeBuild Resource Policy using the CodeBuild Resource Policy arn. For example:

```sh
 $ pulumi import aws:codebuild/resourcePolicy:ResourcePolicy example arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
```
 