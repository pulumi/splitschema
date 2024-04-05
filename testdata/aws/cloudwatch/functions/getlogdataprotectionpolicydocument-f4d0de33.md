Generates a CloudWatch Log Group Data Protection Policy document in JSON format for use with the `aws.cloudwatch.LogDataProtectionPolicy` resource.

> For more information about data protection policies, see the [Help protect sensitive log data with masking](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogDataProtectionPolicyDocument = aws.cloudwatch.getLogDataProtectionPolicyDocument({
    name: "Example",
    statements: [
        {
            sid: "Audit",
            dataIdentifiers: [
                "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US",
            ],
            operation: {
                audit: {
                    findingsDestination: {
                        cloudwatchLogs: {
                            logGroup: aws_cloudwatch_log_group.audit.name,
                        },
                        firehose: {
                            deliveryStream: aws_kinesis_firehose_delivery_stream.audit.name,
                        },
                        s3: {
                            bucket: aws_s3_bucket.audit.bucket,
                        },
                    },
                },
            },
        },
        {
            sid: "Deidentify",
            dataIdentifiers: [
                "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US",
            ],
            operation: {
                deidentify: {
                    maskConfig: {},
                },
            },
        },
    ],
});
const exampleLogDataProtectionPolicy = new aws.cloudwatch.LogDataProtectionPolicy("exampleLogDataProtectionPolicy", {
    logGroupName: aws_cloudwatch_log_group.example.name,
    policyDocument: exampleLogDataProtectionPolicyDocument.then(exampleLogDataProtectionPolicyDocument => exampleLogDataProtectionPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_log_data_protection_policy_document = aws.cloudwatch.get_log_data_protection_policy_document(name="Example",
    statements=[
        aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementArgs(
            sid="Audit",
            data_identifiers=[
                "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US",
            ],
            operation=aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementOperationArgs(
                audit=aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementOperationAuditArgs(
                    findings_destination=aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationArgs(
                        cloudwatch_logs=aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs(
                            log_group=aws_cloudwatch_log_group["audit"]["name"],
                        ),
                        firehose=aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehoseArgs(
                            delivery_stream=aws_kinesis_firehose_delivery_stream["audit"]["name"],
                        ),
                        s3=aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3Args(
                            bucket=aws_s3_bucket["audit"]["bucket"],
                        ),
                    ),
                ),
            ),
        ),
        aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementArgs(
            sid="Deidentify",
            data_identifiers=[
                "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US",
            ],
            operation=aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementOperationArgs(
                deidentify=aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementOperationDeidentifyArgs(
                    mask_config=aws.cloudwatch.GetLogDataProtectionPolicyDocumentStatementOperationDeidentifyMaskConfigArgs(),
                ),
            ),
        ),
    ])
example_log_data_protection_policy = aws.cloudwatch.LogDataProtectionPolicy("exampleLogDataProtectionPolicy",
    log_group_name=aws_cloudwatch_log_group["example"]["name"],
    policy_document=example_log_data_protection_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLogDataProtectionPolicyDocument = Aws.CloudWatch.GetLogDataProtectionPolicyDocument.Invoke(new()
    {
        Name = "Example",
        Statements = new[]
        {
            new Aws.CloudWatch.Inputs.GetLogDataProtectionPolicyDocumentStatementInputArgs
            {
                Sid = "Audit",
                DataIdentifiers = new[]
                {
                    "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                    "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US",
                },
                Operation = new Aws.CloudWatch.Inputs.GetLogDataProtectionPolicyDocumentStatementOperationInputArgs
                {
                    Audit = new Aws.CloudWatch.Inputs.GetLogDataProtectionPolicyDocumentStatementOperationAuditInputArgs
                    {
                        FindingsDestination = new Aws.CloudWatch.Inputs.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationInputArgs
                        {
                            CloudwatchLogs = new Aws.CloudWatch.Inputs.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsInputArgs
                            {
                                LogGroup = aws_cloudwatch_log_group.Audit.Name,
                            },
                            Firehose = new Aws.CloudWatch.Inputs.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehoseInputArgs
                            {
                                DeliveryStream = aws_kinesis_firehose_delivery_stream.Audit.Name,
                            },
                            S3 = new Aws.CloudWatch.Inputs.GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3InputArgs
                            {
                                Bucket = aws_s3_bucket.Audit.Bucket,
                            },
                        },
                    },
                },
            },
            new Aws.CloudWatch.Inputs.GetLogDataProtectionPolicyDocumentStatementInputArgs
            {
                Sid = "Deidentify",
                DataIdentifiers = new[]
                {
                    "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                    "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US",
                },
                Operation = new Aws.CloudWatch.Inputs.GetLogDataProtectionPolicyDocumentStatementOperationInputArgs
                {
                    Deidentify = new Aws.CloudWatch.Inputs.GetLogDataProtectionPolicyDocumentStatementOperationDeidentifyInputArgs
                    {
                        MaskConfig = null,
                    },
                },
            },
        },
    });

    var exampleLogDataProtectionPolicy = new Aws.CloudWatch.LogDataProtectionPolicy("exampleLogDataProtectionPolicy", new()
    {
        LogGroupName = aws_cloudwatch_log_group.Example.Name,
        PolicyDocument = exampleLogDataProtectionPolicyDocument.Apply(getLogDataProtectionPolicyDocumentResult => getLogDataProtectionPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogDataProtectionPolicyDocument, err := cloudwatch.GetLogDataProtectionPolicyDocument(ctx, &cloudwatch.GetLogDataProtectionPolicyDocumentArgs{
			Name: "Example",
			Statements: []cloudwatch.GetLogDataProtectionPolicyDocumentStatement{
				{
					Sid: pulumi.StringRef("Audit"),
					DataIdentifiers: []string{
						"arn:aws:dataprotection::aws:data-identifier/EmailAddress",
						"arn:aws:dataprotection::aws:data-identifier/DriversLicense-US",
					},
					Operation: {
						Audit: {
							FindingsDestination: {
								CloudwatchLogs: {
									LogGroup: aws_cloudwatch_log_group.Audit.Name,
								},
								Firehose: {
									DeliveryStream: aws_kinesis_firehose_delivery_stream.Audit.Name,
								},
								S3: {
									Bucket: aws_s3_bucket.Audit.Bucket,
								},
							},
						},
					},
				},
				{
					Sid: pulumi.StringRef("Deidentify"),
					DataIdentifiers: []string{
						"arn:aws:dataprotection::aws:data-identifier/EmailAddress",
						"arn:aws:dataprotection::aws:data-identifier/DriversLicense-US",
					},
					Operation: {
						Deidentify: {
							MaskConfig: nil,
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogDataProtectionPolicy(ctx, "exampleLogDataProtectionPolicy", &cloudwatch.LogDataProtectionPolicyArgs{
			LogGroupName:   pulumi.Any(aws_cloudwatch_log_group.Example.Name),
			PolicyDocument: *pulumi.String(exampleLogDataProtectionPolicyDocument.Json),
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
import com.pulumi.aws.cloudwatch.CloudwatchFunctions;
import com.pulumi.aws.cloudwatch.inputs.GetLogDataProtectionPolicyDocumentArgs;
import com.pulumi.aws.cloudwatch.LogDataProtectionPolicy;
import com.pulumi.aws.cloudwatch.LogDataProtectionPolicyArgs;
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
        final var exampleLogDataProtectionPolicyDocument = CloudwatchFunctions.getLogDataProtectionPolicyDocument(GetLogDataProtectionPolicyDocumentArgs.builder()
            .name("Example")
            .statements(            
                GetLogDataProtectionPolicyDocumentStatementArgs.builder()
                    .sid("Audit")
                    .dataIdentifiers(                    
                        "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                        "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US")
                    .operation(GetLogDataProtectionPolicyDocumentStatementOperationArgs.builder()
                        .audit(GetLogDataProtectionPolicyDocumentStatementOperationAuditArgs.builder()
                            .findingsDestination(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationArgs.builder()
                                .cloudwatchLogs(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogsArgs.builder()
                                    .logGroup(aws_cloudwatch_log_group.audit().name())
                                    .build())
                                .firehose(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehoseArgs.builder()
                                    .deliveryStream(aws_kinesis_firehose_delivery_stream.audit().name())
                                    .build())
                                .s3(GetLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3Args.builder()
                                    .bucket(aws_s3_bucket.audit().bucket())
                                    .build())
                                .build())
                            .build())
                        .build())
                    .build(),
                GetLogDataProtectionPolicyDocumentStatementArgs.builder()
                    .sid("Deidentify")
                    .dataIdentifiers(                    
                        "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                        "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US")
                    .operation(GetLogDataProtectionPolicyDocumentStatementOperationArgs.builder()
                        .deidentify(GetLogDataProtectionPolicyDocumentStatementOperationDeidentifyArgs.builder()
                            .maskConfig()
                            .build())
                        .build())
                    .build())
            .build());

        var exampleLogDataProtectionPolicy = new LogDataProtectionPolicy("exampleLogDataProtectionPolicy", LogDataProtectionPolicyArgs.builder()        
            .logGroupName(aws_cloudwatch_log_group.example().name())
            .policyDocument(exampleLogDataProtectionPolicyDocument.applyValue(getLogDataProtectionPolicyDocumentResult -> getLogDataProtectionPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  exampleLogDataProtectionPolicy:
    type: aws:cloudwatch:LogDataProtectionPolicy
    properties:
      logGroupName: ${aws_cloudwatch_log_group.example.name}
      policyDocument: ${exampleLogDataProtectionPolicyDocument.json}
variables:
  exampleLogDataProtectionPolicyDocument:
    fn::invoke:
      Function: aws:cloudwatch:getLogDataProtectionPolicyDocument
      Arguments:
        name: Example
        statements:
          - sid: Audit
            dataIdentifiers:
              - arn:aws:dataprotection::aws:data-identifier/EmailAddress
              - arn:aws:dataprotection::aws:data-identifier/DriversLicense-US
            operation:
              audit:
                findingsDestination:
                  cloudwatchLogs:
                    logGroup: ${aws_cloudwatch_log_group.audit.name}
                  firehose:
                    deliveryStream: ${aws_kinesis_firehose_delivery_stream.audit.name}
                  s3:
                    bucket: ${aws_s3_bucket.audit.bucket}
          - sid: Deidentify
            dataIdentifiers:
              - arn:aws:dataprotection::aws:data-identifier/EmailAddress
              - arn:aws:dataprotection::aws:data-identifier/DriversLicense-US
            operation:
              deidentify:
                maskConfig: {}
```
{{% /example %}}
{{% /examples %}}