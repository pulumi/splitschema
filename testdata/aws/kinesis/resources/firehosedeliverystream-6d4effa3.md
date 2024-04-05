Provides a Kinesis Firehose Delivery Stream resource. Amazon Kinesis Firehose is a fully managed, elastic service to easily deliver real-time data streams to destinations such as Amazon S3 and Amazon Redshift.

For more details, see the [Amazon Kinesis Firehose Documentation](https://aws.amazon.com/documentation/firehose/).

{{% examples %}}
## Example Usage
{{% example %}}
### Extended S3 Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.BucketV2("bucket", {});
const firehoseAssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["firehose.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const firehoseRole = new aws.iam.Role("firehoseRole", {assumeRolePolicy: firehoseAssumeRole.then(firehoseAssumeRole => firehoseAssumeRole.json)});
const lambdaAssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["lambda.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const lambdaIam = new aws.iam.Role("lambdaIam", {assumeRolePolicy: lambdaAssumeRole.then(lambdaAssumeRole => lambdaAssumeRole.json)});
const lambdaProcessor = new aws.lambda.Function("lambdaProcessor", {
    code: new pulumi.asset.FileArchive("lambda.zip"),
    role: lambdaIam.arn,
    handler: "exports.handler",
    runtime: "nodejs16.x",
});
const extendedS3Stream = new aws.kinesis.FirehoseDeliveryStream("extendedS3Stream", {
    destination: "extended_s3",
    extendedS3Configuration: {
        roleArn: firehoseRole.arn,
        bucketArn: bucket.arn,
        processingConfiguration: {
            enabled: true,
            processors: [{
                type: "Lambda",
                parameters: [{
                    parameterName: "LambdaArn",
                    parameterValue: pulumi.interpolate`${lambdaProcessor.arn}:$LATEST`,
                }],
            }],
        },
    },
});
const bucketAcl = new aws.s3.BucketAclV2("bucketAcl", {
    bucket: bucket.id,
    acl: "private",
});
```
```python
import pulumi
import pulumi_aws as aws

bucket = aws.s3.BucketV2("bucket")
firehose_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["firehose.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
firehose_role = aws.iam.Role("firehoseRole", assume_role_policy=firehose_assume_role.json)
lambda_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["lambda.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
lambda_iam = aws.iam.Role("lambdaIam", assume_role_policy=lambda_assume_role.json)
lambda_processor = aws.lambda_.Function("lambdaProcessor",
    code=pulumi.FileArchive("lambda.zip"),
    role=lambda_iam.arn,
    handler="exports.handler",
    runtime="nodejs16.x")
extended_s3_stream = aws.kinesis.FirehoseDeliveryStream("extendedS3Stream",
    destination="extended_s3",
    extended_s3_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs(
        role_arn=firehose_role.arn,
        bucket_arn=bucket.arn,
        processing_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs(
            enabled=True,
            processors=[aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs(
                type="Lambda",
                parameters=[aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs(
                    parameter_name="LambdaArn",
                    parameter_value=lambda_processor.arn.apply(lambda arn: f"{arn}:$LATEST"),
                )],
            )],
        ),
    ))
bucket_acl = aws.s3.BucketAclV2("bucketAcl",
    bucket=bucket.id,
    acl="private")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bucket = new Aws.S3.BucketV2("bucket");

    var firehoseAssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "firehose.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var firehoseRole = new Aws.Iam.Role("firehoseRole", new()
    {
        AssumeRolePolicy = firehoseAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var lambdaAssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "lambda.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var lambdaIam = new Aws.Iam.Role("lambdaIam", new()
    {
        AssumeRolePolicy = lambdaAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var lambdaProcessor = new Aws.Lambda.Function("lambdaProcessor", new()
    {
        Code = new FileArchive("lambda.zip"),
        Role = lambdaIam.Arn,
        Handler = "exports.handler",
        Runtime = "nodejs16.x",
    });

    var extendedS3Stream = new Aws.Kinesis.FirehoseDeliveryStream("extendedS3Stream", new()
    {
        Destination = "extended_s3",
        ExtendedS3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs
        {
            RoleArn = firehoseRole.Arn,
            BucketArn = bucket.Arn,
            ProcessingConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs
            {
                Enabled = true,
                Processors = new[]
                {
                    new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs
                    {
                        Type = "Lambda",
                        Parameters = new[]
                        {
                            new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs
                            {
                                ParameterName = "LambdaArn",
                                ParameterValue = lambdaProcessor.Arn.Apply(arn => $"{arn}:$LATEST"),
                            },
                        },
                    },
                },
            },
        },
    });

    var bucketAcl = new Aws.S3.BucketAclV2("bucketAcl", new()
    {
        Bucket = bucket.Id,
        Acl = "private",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		firehoseAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"firehose.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		firehoseRole, err := iam.NewRole(ctx, "firehoseRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(firehoseAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		lambdaAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"lambda.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		lambdaIam, err := iam.NewRole(ctx, "lambdaIam", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(lambdaAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		lambdaProcessor, err := lambda.NewFunction(ctx, "lambdaProcessor", &lambda.FunctionArgs{
			Code:    pulumi.NewFileArchive("lambda.zip"),
			Role:    lambdaIam.Arn,
			Handler: pulumi.String("exports.handler"),
			Runtime: pulumi.String("nodejs16.x"),
		})
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "extendedS3Stream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("extended_s3"),
			ExtendedS3Configuration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs{
				RoleArn:   firehoseRole.Arn,
				BucketArn: bucket.Arn,
				ProcessingConfiguration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs{
					Enabled: pulumi.Bool(true),
					Processors: kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArray{
						&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs{
							Type: pulumi.String("Lambda"),
							Parameters: kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArray{
								&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs{
									ParameterName: pulumi.String("LambdaArn"),
									ParameterValue: lambdaProcessor.Arn.ApplyT(func(arn string) (string, error) {
										return fmt.Sprintf("%v:$LATEST", arn), nil
									}).(pulumi.StringOutput),
								},
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "bucketAcl", &s3.BucketAclV2Args{
			Bucket: bucket.ID(),
			Acl:    pulumi.String("private"),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.asset.FileArchive;
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
        var bucket = new BucketV2("bucket");

        final var firehoseAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("firehose.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var firehoseRole = new Role("firehoseRole", RoleArgs.builder()        
            .assumeRolePolicy(firehoseAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var lambdaAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("lambda.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var lambdaIam = new Role("lambdaIam", RoleArgs.builder()        
            .assumeRolePolicy(lambdaAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var lambdaProcessor = new Function("lambdaProcessor", FunctionArgs.builder()        
            .code(new FileArchive("lambda.zip"))
            .role(lambdaIam.arn())
            .handler("exports.handler")
            .runtime("nodejs16.x")
            .build());

        var extendedS3Stream = new FirehoseDeliveryStream("extendedS3Stream", FirehoseDeliveryStreamArgs.builder()        
            .destination("extended_s3")
            .extendedS3Configuration(FirehoseDeliveryStreamExtendedS3ConfigurationArgs.builder()
                .roleArn(firehoseRole.arn())
                .bucketArn(bucket.arn())
                .processingConfiguration(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs.builder()
                    .enabled("true")
                    .processors(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs.builder()
                        .type("Lambda")
                        .parameters(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs.builder()
                            .parameterName("LambdaArn")
                            .parameterValue(lambdaProcessor.arn().applyValue(arn -> String.format("%s:$LATEST", arn)))
                            .build())
                        .build())
                    .build())
                .build())
            .build());

        var bucketAcl = new BucketAclV2("bucketAcl", BucketAclV2Args.builder()        
            .bucket(bucket.id())
            .acl("private")
            .build());

    }
}
```
```yaml
resources:
  extendedS3Stream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: extended_s3
      extendedS3Configuration:
        roleArn: ${firehoseRole.arn}
        bucketArn: ${bucket.arn}
        processingConfiguration:
          enabled: 'true'
          processors:
            - type: Lambda
              parameters:
                - parameterName: LambdaArn
                  parameterValue: ${lambdaProcessor.arn}:$LATEST
  bucket:
    type: aws:s3:BucketV2
  bucketAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${bucket.id}
      acl: private
  firehoseRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${firehoseAssumeRole.json}
  lambdaIam:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${lambdaAssumeRole.json}
  lambdaProcessor:
    type: aws:lambda:Function
    properties:
      code:
        fn::FileArchive: lambda.zip
      role: ${lambdaIam.arn}
      handler: exports.handler
      runtime: nodejs16.x
variables:
  firehoseAssumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - firehose.amazonaws.com
            actions:
              - sts:AssumeRole
  lambdaAssumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - lambda.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% example %}}
### Extended S3 Destination with dynamic partitioning

These examples use built-in Firehose functionality, rather than requiring a lambda.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const extendedS3Stream = new aws.kinesis.FirehoseDeliveryStream("extendedS3Stream", {
    destination: "extended_s3",
    extendedS3Configuration: {
        roleArn: aws_iam_role.firehose_role.arn,
        bucketArn: aws_s3_bucket.bucket.arn,
        bufferingSize: 64,
        dynamicPartitioningConfiguration: {
            enabled: true,
        },
        prefix: "data/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/",
        errorOutputPrefix: "errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/",
        processingConfiguration: {
            enabled: true,
            processors: [
                {
                    type: "RecordDeAggregation",
                    parameters: [{
                        parameterName: "SubRecordType",
                        parameterValue: "JSON",
                    }],
                },
                {
                    type: "AppendDelimiterToRecord",
                },
                {
                    type: "MetadataExtraction",
                    parameters: [
                        {
                            parameterName: "JsonParsingEngine",
                            parameterValue: "JQ-1.6",
                        },
                        {
                            parameterName: "MetadataExtractionQuery",
                            parameterValue: "{customer_id:.customer_id}",
                        },
                    ],
                },
            ],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

extended_s3_stream = aws.kinesis.FirehoseDeliveryStream("extendedS3Stream",
    destination="extended_s3",
    extended_s3_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs(
        role_arn=aws_iam_role["firehose_role"]["arn"],
        bucket_arn=aws_s3_bucket["bucket"]["arn"],
        buffering_size=64,
        dynamic_partitioning_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfigurationArgs(
            enabled=True,
        ),
        prefix="data/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/",
        error_output_prefix="errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/",
        processing_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs(
            enabled=True,
            processors=[
                aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs(
                    type="RecordDeAggregation",
                    parameters=[aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs(
                        parameter_name="SubRecordType",
                        parameter_value="JSON",
                    )],
                ),
                aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs(
                    type="AppendDelimiterToRecord",
                ),
                aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs(
                    type="MetadataExtraction",
                    parameters=[
                        aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs(
                            parameter_name="JsonParsingEngine",
                            parameter_value="JQ-1.6",
                        ),
                        aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs(
                            parameter_name="MetadataExtractionQuery",
                            parameter_value="{customer_id:.customer_id}",
                        ),
                    ],
                ),
            ],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var extendedS3Stream = new Aws.Kinesis.FirehoseDeliveryStream("extendedS3Stream", new()
    {
        Destination = "extended_s3",
        ExtendedS3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs
        {
            RoleArn = aws_iam_role.Firehose_role.Arn,
            BucketArn = aws_s3_bucket.Bucket.Arn,
            BufferingSize = 64,
            DynamicPartitioningConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfigurationArgs
            {
                Enabled = true,
            },
            Prefix = "data/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/",
            ErrorOutputPrefix = "errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/",
            ProcessingConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs
            {
                Enabled = true,
                Processors = new[]
                {
                    new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs
                    {
                        Type = "RecordDeAggregation",
                        Parameters = new[]
                        {
                            new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs
                            {
                                ParameterName = "SubRecordType",
                                ParameterValue = "JSON",
                            },
                        },
                    },
                    new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs
                    {
                        Type = "AppendDelimiterToRecord",
                    },
                    new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs
                    {
                        Type = "MetadataExtraction",
                        Parameters = new[]
                        {
                            new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs
                            {
                                ParameterName = "JsonParsingEngine",
                                ParameterValue = "JQ-1.6",
                            },
                            new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs
                            {
                                ParameterName = "MetadataExtractionQuery",
                                ParameterValue = "{customer_id:.customer_id}",
                            },
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kinesis.NewFirehoseDeliveryStream(ctx, "extendedS3Stream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("extended_s3"),
			ExtendedS3Configuration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs{
				RoleArn:       pulumi.Any(aws_iam_role.Firehose_role.Arn),
				BucketArn:     pulumi.Any(aws_s3_bucket.Bucket.Arn),
				BufferingSize: pulumi.Int(64),
				DynamicPartitioningConfiguration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfigurationArgs{
					Enabled: pulumi.Bool(true),
				},
				Prefix:            pulumi.String("data/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/"),
				ErrorOutputPrefix: pulumi.String("errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/"),
				ProcessingConfiguration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs{
					Enabled: pulumi.Bool(true),
					Processors: kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArray{
						&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs{
							Type: pulumi.String("RecordDeAggregation"),
							Parameters: kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArray{
								&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs{
									ParameterName:  pulumi.String("SubRecordType"),
									ParameterValue: pulumi.String("JSON"),
								},
							},
						},
						&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs{
							Type: pulumi.String("AppendDelimiterToRecord"),
						},
						&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs{
							Type: pulumi.String("MetadataExtraction"),
							Parameters: kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArray{
								&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs{
									ParameterName:  pulumi.String("JsonParsingEngine"),
									ParameterValue: pulumi.String("JQ-1.6"),
								},
								&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs{
									ParameterName:  pulumi.String("MetadataExtractionQuery"),
									ParameterValue: pulumi.String("{customer_id:.customer_id}"),
								},
							},
						},
					},
				},
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
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs;
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
        var extendedS3Stream = new FirehoseDeliveryStream("extendedS3Stream", FirehoseDeliveryStreamArgs.builder()        
            .destination("extended_s3")
            .extendedS3Configuration(FirehoseDeliveryStreamExtendedS3ConfigurationArgs.builder()
                .roleArn(aws_iam_role.firehose_role().arn())
                .bucketArn(aws_s3_bucket.bucket().arn())
                .bufferingSize(64)
                .dynamicPartitioningConfiguration(FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfigurationArgs.builder()
                    .enabled("true")
                    .build())
                .prefix("data/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/")
                .errorOutputPrefix("errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/")
                .processingConfiguration(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs.builder()
                    .enabled("true")
                    .processors(                    
                        FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs.builder()
                            .type("RecordDeAggregation")
                            .parameters(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs.builder()
                                .parameterName("SubRecordType")
                                .parameterValue("JSON")
                                .build())
                            .build(),
                        FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs.builder()
                            .type("AppendDelimiterToRecord")
                            .build(),
                        FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs.builder()
                            .type("MetadataExtraction")
                            .parameters(                            
                                FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs.builder()
                                    .parameterName("JsonParsingEngine")
                                    .parameterValue("JQ-1.6")
                                    .build(),
                                FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs.builder()
                                    .parameterName("MetadataExtractionQuery")
                                    .parameterValue("{customer_id:.customer_id}")
                                    .build())
                            .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  extendedS3Stream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: extended_s3
      extendedS3Configuration:
        roleArn: ${aws_iam_role.firehose_role.arn}
        bucketArn: ${aws_s3_bucket.bucket.arn}
        bufferingSize: 64
        dynamicPartitioningConfiguration:
          enabled: 'true'
        prefix: data/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/
        errorOutputPrefix: errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/
        processingConfiguration:
          enabled: 'true'
          processors:
            - type: RecordDeAggregation
              parameters:
                - parameterName: SubRecordType
                  parameterValue: JSON
            - type: AppendDelimiterToRecord
            - type: MetadataExtraction
              parameters:
                - parameterName: JsonParsingEngine
                  parameterValue: JQ-1.6
                - parameterName: MetadataExtractionQuery
                  parameterValue: '{customer_id:.customer_id}'
```

Multiple Dynamic Partitioning Keys (maximum of 50) can be added by comma separating the `parameter_value`.

The following example adds the Dynamic Partitioning Keys: `store_id` and `customer_id` to the S3 prefix.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const extendedS3Stream = new aws.kinesis.FirehoseDeliveryStream("extendedS3Stream", {
    destination: "extended_s3",
    extendedS3Configuration: {
        roleArn: aws_iam_role.firehose_role.arn,
        bucketArn: aws_s3_bucket.bucket.arn,
        bufferingSize: 64,
        dynamicPartitioningConfiguration: {
            enabled: true,
        },
        prefix: "data/store_id=!{partitionKeyFromQuery:store_id}/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/",
        errorOutputPrefix: "errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/",
        processingConfiguration: {
            enabled: true,
            processors: [{
                type: "MetadataExtraction",
                parameters: [
                    {
                        parameterName: "JsonParsingEngine",
                        parameterValue: "JQ-1.6",
                    },
                    {
                        parameterName: "MetadataExtractionQuery",
                        parameterValue: "{store_id:.store_id,customer_id:.customer_id}",
                    },
                ],
            }],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

extended_s3_stream = aws.kinesis.FirehoseDeliveryStream("extendedS3Stream",
    destination="extended_s3",
    extended_s3_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs(
        role_arn=aws_iam_role["firehose_role"]["arn"],
        bucket_arn=aws_s3_bucket["bucket"]["arn"],
        buffering_size=64,
        dynamic_partitioning_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfigurationArgs(
            enabled=True,
        ),
        prefix="data/store_id=!{partitionKeyFromQuery:store_id}/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/",
        error_output_prefix="errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/",
        processing_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs(
            enabled=True,
            processors=[aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs(
                type="MetadataExtraction",
                parameters=[
                    aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs(
                        parameter_name="JsonParsingEngine",
                        parameter_value="JQ-1.6",
                    ),
                    aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs(
                        parameter_name="MetadataExtractionQuery",
                        parameter_value="{store_id:.store_id,customer_id:.customer_id}",
                    ),
                ],
            )],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var extendedS3Stream = new Aws.Kinesis.FirehoseDeliveryStream("extendedS3Stream", new()
    {
        Destination = "extended_s3",
        ExtendedS3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs
        {
            RoleArn = aws_iam_role.Firehose_role.Arn,
            BucketArn = aws_s3_bucket.Bucket.Arn,
            BufferingSize = 64,
            DynamicPartitioningConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfigurationArgs
            {
                Enabled = true,
            },
            Prefix = "data/store_id=!{partitionKeyFromQuery:store_id}/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/",
            ErrorOutputPrefix = "errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/",
            ProcessingConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs
            {
                Enabled = true,
                Processors = new[]
                {
                    new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs
                    {
                        Type = "MetadataExtraction",
                        Parameters = new[]
                        {
                            new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs
                            {
                                ParameterName = "JsonParsingEngine",
                                ParameterValue = "JQ-1.6",
                            },
                            new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs
                            {
                                ParameterName = "MetadataExtractionQuery",
                                ParameterValue = "{store_id:.store_id,customer_id:.customer_id}",
                            },
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kinesis.NewFirehoseDeliveryStream(ctx, "extendedS3Stream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("extended_s3"),
			ExtendedS3Configuration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs{
				RoleArn:       pulumi.Any(aws_iam_role.Firehose_role.Arn),
				BucketArn:     pulumi.Any(aws_s3_bucket.Bucket.Arn),
				BufferingSize: pulumi.Int(64),
				DynamicPartitioningConfiguration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfigurationArgs{
					Enabled: pulumi.Bool(true),
				},
				Prefix:            pulumi.String("data/store_id=!{partitionKeyFromQuery:store_id}/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/"),
				ErrorOutputPrefix: pulumi.String("errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/"),
				ProcessingConfiguration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs{
					Enabled: pulumi.Bool(true),
					Processors: kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArray{
						&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs{
							Type: pulumi.String("MetadataExtraction"),
							Parameters: kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArray{
								&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs{
									ParameterName:  pulumi.String("JsonParsingEngine"),
									ParameterValue: pulumi.String("JQ-1.6"),
								},
								&kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs{
									ParameterName:  pulumi.String("MetadataExtractionQuery"),
									ParameterValue: pulumi.String("{store_id:.store_id,customer_id:.customer_id}"),
								},
							},
						},
					},
				},
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
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs;
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
        var extendedS3Stream = new FirehoseDeliveryStream("extendedS3Stream", FirehoseDeliveryStreamArgs.builder()        
            .destination("extended_s3")
            .extendedS3Configuration(FirehoseDeliveryStreamExtendedS3ConfigurationArgs.builder()
                .roleArn(aws_iam_role.firehose_role().arn())
                .bucketArn(aws_s3_bucket.bucket().arn())
                .bufferingSize(64)
                .dynamicPartitioningConfiguration(FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfigurationArgs.builder()
                    .enabled("true")
                    .build())
                .prefix("data/store_id=!{partitionKeyFromQuery:store_id}/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/")
                .errorOutputPrefix("errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/")
                .processingConfiguration(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationArgs.builder()
                    .enabled("true")
                    .processors(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs.builder()
                        .type("MetadataExtraction")
                        .parameters(                        
                            FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs.builder()
                                .parameterName("JsonParsingEngine")
                                .parameterValue("JQ-1.6")
                                .build(),
                            FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArgs.builder()
                                .parameterName("MetadataExtractionQuery")
                                .parameterValue("{store_id:.store_id,customer_id:.customer_id}")
                                .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  extendedS3Stream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: extended_s3
      extendedS3Configuration:
        roleArn: ${aws_iam_role.firehose_role.arn}
        bucketArn: ${aws_s3_bucket.bucket.arn}
        bufferingSize: 64
        dynamicPartitioningConfiguration:
          enabled: 'true'
        prefix: data/store_id=!{partitionKeyFromQuery:store_id}/customer_id=!{partitionKeyFromQuery:customer_id}/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/
        errorOutputPrefix: errors/year=!{timestamp:yyyy}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/!{firehose:error-output-type}/
        processingConfiguration:
          enabled: 'true'
          processors:
            - type: MetadataExtraction
              parameters:
                - parameterName: JsonParsingEngine
                  parameterValue: JQ-1.6
                - parameterName: MetadataExtractionQuery
                  parameterValue: '{store_id:.store_id,customer_id:.customer_id}'
```
{{% /example %}}
{{% example %}}
### Redshift Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCluster = new aws.redshift.Cluster("testCluster", {
    clusterIdentifier: "tf-redshift-cluster",
    databaseName: "test",
    masterUsername: "testuser",
    masterPassword: "T3stPass",
    nodeType: "dc1.large",
    clusterType: "single-node",
});
const testStream = new aws.kinesis.FirehoseDeliveryStream("testStream", {
    destination: "redshift",
    redshiftConfiguration: {
        roleArn: aws_iam_role.firehose_role.arn,
        clusterJdbcurl: pulumi.interpolate`jdbc:redshift://${testCluster.endpoint}/${testCluster.databaseName}`,
        username: "testuser",
        password: "T3stPass",
        dataTableName: "test-table",
        copyOptions: "delimiter '|'",
        dataTableColumns: "test-col",
        s3BackupMode: "Enabled",
        s3Configuration: {
            roleArn: aws_iam_role.firehose_role.arn,
            bucketArn: aws_s3_bucket.bucket.arn,
            bufferingSize: 10,
            bufferingInterval: 400,
            compressionFormat: "GZIP",
        },
        s3BackupConfiguration: {
            roleArn: aws_iam_role.firehose_role.arn,
            bucketArn: aws_s3_bucket.bucket.arn,
            bufferingSize: 15,
            bufferingInterval: 300,
            compressionFormat: "GZIP",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_cluster = aws.redshift.Cluster("testCluster",
    cluster_identifier="tf-redshift-cluster",
    database_name="test",
    master_username="testuser",
    master_password="T3stPass",
    node_type="dc1.large",
    cluster_type="single-node")
test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="redshift",
    redshift_configuration=aws.kinesis.FirehoseDeliveryStreamRedshiftConfigurationArgs(
        role_arn=aws_iam_role["firehose_role"]["arn"],
        cluster_jdbcurl=pulumi.Output.all(test_cluster.endpoint, test_cluster.database_name).apply(lambda endpoint, database_name: f"jdbc:redshift://{endpoint}/{database_name}"),
        username="testuser",
        password="T3stPass",
        data_table_name="test-table",
        copy_options="delimiter '|'",
        data_table_columns="test-col",
        s3_backup_mode="Enabled",
        s3_configuration=aws.kinesis.FirehoseDeliveryStreamRedshiftConfigurationS3ConfigurationArgs(
            role_arn=aws_iam_role["firehose_role"]["arn"],
            bucket_arn=aws_s3_bucket["bucket"]["arn"],
            buffering_size=10,
            buffering_interval=400,
            compression_format="GZIP",
        ),
        s3_backup_configuration=aws.kinesis.FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationArgs(
            role_arn=aws_iam_role["firehose_role"]["arn"],
            bucket_arn=aws_s3_bucket["bucket"]["arn"],
            buffering_size=15,
            buffering_interval=300,
            compression_format="GZIP",
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCluster = new Aws.RedShift.Cluster("testCluster", new()
    {
        ClusterIdentifier = "tf-redshift-cluster",
        DatabaseName = "test",
        MasterUsername = "testuser",
        MasterPassword = "T3stPass",
        NodeType = "dc1.large",
        ClusterType = "single-node",
    });

    var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new()
    {
        Destination = "redshift",
        RedshiftConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamRedshiftConfigurationArgs
        {
            RoleArn = aws_iam_role.Firehose_role.Arn,
            ClusterJdbcurl = Output.Tuple(testCluster.Endpoint, testCluster.DatabaseName).Apply(values =>
            {
                var endpoint = values.Item1;
                var databaseName = values.Item2;
                return $"jdbc:redshift://{endpoint}/{databaseName}";
            }),
            Username = "testuser",
            Password = "T3stPass",
            DataTableName = "test-table",
            CopyOptions = "delimiter '|'",
            DataTableColumns = "test-col",
            S3BackupMode = "Enabled",
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamRedshiftConfigurationS3ConfigurationArgs
            {
                RoleArn = aws_iam_role.Firehose_role.Arn,
                BucketArn = aws_s3_bucket.Bucket.Arn,
                BufferingSize = 10,
                BufferingInterval = 400,
                CompressionFormat = "GZIP",
            },
            S3BackupConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationArgs
            {
                RoleArn = aws_iam_role.Firehose_role.Arn,
                BucketArn = aws_s3_bucket.Bucket.Arn,
                BufferingSize = 15,
                BufferingInterval = 300,
                CompressionFormat = "GZIP",
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testCluster, err := redshift.NewCluster(ctx, "testCluster", &redshift.ClusterArgs{
			ClusterIdentifier: pulumi.String("tf-redshift-cluster"),
			DatabaseName:      pulumi.String("test"),
			MasterUsername:    pulumi.String("testuser"),
			MasterPassword:    pulumi.String("T3stPass"),
			NodeType:          pulumi.String("dc1.large"),
			ClusterType:       pulumi.String("single-node"),
		})
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("redshift"),
			RedshiftConfiguration: &kinesis.FirehoseDeliveryStreamRedshiftConfigurationArgs{
				RoleArn: pulumi.Any(aws_iam_role.Firehose_role.Arn),
				ClusterJdbcurl: pulumi.All(testCluster.Endpoint, testCluster.DatabaseName).ApplyT(func(_args []interface{}) (string, error) {
					endpoint := _args[0].(string)
					databaseName := _args[1].(string)
					return fmt.Sprintf("jdbc:redshift://%v/%v", endpoint, databaseName), nil
				}).(pulumi.StringOutput),
				Username:         pulumi.String("testuser"),
				Password:         pulumi.String("T3stPass"),
				DataTableName:    pulumi.String("test-table"),
				CopyOptions:      pulumi.String("delimiter '|'"),
				DataTableColumns: pulumi.String("test-col"),
				S3BackupMode:     pulumi.String("Enabled"),
				S3Configuration: &kinesis.FirehoseDeliveryStreamRedshiftConfigurationS3ConfigurationArgs{
					RoleArn:           pulumi.Any(aws_iam_role.Firehose_role.Arn),
					BucketArn:         pulumi.Any(aws_s3_bucket.Bucket.Arn),
					BufferingSize:     pulumi.Int(10),
					BufferingInterval: pulumi.Int(400),
					CompressionFormat: pulumi.String("GZIP"),
				},
				S3BackupConfiguration: &kinesis.FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationArgs{
					RoleArn:           pulumi.Any(aws_iam_role.Firehose_role.Arn),
					BucketArn:         pulumi.Any(aws_s3_bucket.Bucket.Arn),
					BufferingSize:     pulumi.Int(15),
					BufferingInterval: pulumi.Int(300),
					CompressionFormat: pulumi.String("GZIP"),
				},
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
import com.pulumi.aws.redshift.Cluster;
import com.pulumi.aws.redshift.ClusterArgs;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamRedshiftConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamRedshiftConfigurationS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationArgs;
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
        var testCluster = new Cluster("testCluster", ClusterArgs.builder()        
            .clusterIdentifier("tf-redshift-cluster")
            .databaseName("test")
            .masterUsername("testuser")
            .masterPassword("T3stPass")
            .nodeType("dc1.large")
            .clusterType("single-node")
            .build());

        var testStream = new FirehoseDeliveryStream("testStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("redshift")
            .redshiftConfiguration(FirehoseDeliveryStreamRedshiftConfigurationArgs.builder()
                .roleArn(aws_iam_role.firehose_role().arn())
                .clusterJdbcurl(Output.tuple(testCluster.endpoint(), testCluster.databaseName()).applyValue(values -> {
                    var endpoint = values.t1;
                    var databaseName = values.t2;
                    return String.format("jdbc:redshift://%s/%s", endpoint,databaseName);
                }))
                .username("testuser")
                .password("T3stPass")
                .dataTableName("test-table")
                .copyOptions("delimiter '|'")
                .dataTableColumns("test-col")
                .s3BackupMode("Enabled")
                .s3Configuration(FirehoseDeliveryStreamRedshiftConfigurationS3ConfigurationArgs.builder()
                    .roleArn(aws_iam_role.firehose_role().arn())
                    .bucketArn(aws_s3_bucket.bucket().arn())
                    .bufferingSize(10)
                    .bufferingInterval(400)
                    .compressionFormat("GZIP")
                    .build())
                .s3BackupConfiguration(FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationArgs.builder()
                    .roleArn(aws_iam_role.firehose_role().arn())
                    .bucketArn(aws_s3_bucket.bucket().arn())
                    .bufferingSize(15)
                    .bufferingInterval(300)
                    .compressionFormat("GZIP")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testCluster:
    type: aws:redshift:Cluster
    properties:
      clusterIdentifier: tf-redshift-cluster
      databaseName: test
      masterUsername: testuser
      masterPassword: T3stPass
      nodeType: dc1.large
      clusterType: single-node
  testStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: redshift
      redshiftConfiguration:
        roleArn: ${aws_iam_role.firehose_role.arn}
        clusterJdbcurl: jdbc:redshift://${testCluster.endpoint}/${testCluster.databaseName}
        username: testuser
        password: T3stPass
        dataTableName: test-table
        copyOptions: delimiter '|'
        dataTableColumns: test-col
        s3BackupMode: Enabled
        s3Configuration:
          roleArn: ${aws_iam_role.firehose_role.arn}
          bucketArn: ${aws_s3_bucket.bucket.arn}
          bufferingSize: 10
          bufferingInterval: 400
          compressionFormat: GZIP
        s3BackupConfiguration:
          roleArn: ${aws_iam_role.firehose_role.arn}
          bucketArn: ${aws_s3_bucket.bucket.arn}
          bufferingSize: 15
          bufferingInterval: 300
          compressionFormat: GZIP
```
{{% /example %}}
{{% example %}}
### Elasticsearch Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCluster = new aws.elasticsearch.Domain("testCluster", {});
const testStream = new aws.kinesis.FirehoseDeliveryStream("testStream", {
    destination: "elasticsearch",
    elasticsearchConfiguration: {
        domainArn: testCluster.arn,
        roleArn: aws_iam_role.firehose_role.arn,
        indexName: "test",
        typeName: "test",
        s3Configuration: {
            roleArn: aws_iam_role.firehose_role.arn,
            bucketArn: aws_s3_bucket.bucket.arn,
            bufferingSize: 10,
            bufferingInterval: 400,
            compressionFormat: "GZIP",
        },
        processingConfiguration: {
            enabled: true,
            processors: [{
                type: "Lambda",
                parameters: [{
                    parameterName: "LambdaArn",
                    parameterValue: `${aws_lambda_function.lambda_processor.arn}:$LATEST`,
                }],
            }],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_cluster = aws.elasticsearch.Domain("testCluster")
test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="elasticsearch",
    elasticsearch_configuration=aws.kinesis.FirehoseDeliveryStreamElasticsearchConfigurationArgs(
        domain_arn=test_cluster.arn,
        role_arn=aws_iam_role["firehose_role"]["arn"],
        index_name="test",
        type_name="test",
        s3_configuration=aws.kinesis.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs(
            role_arn=aws_iam_role["firehose_role"]["arn"],
            bucket_arn=aws_s3_bucket["bucket"]["arn"],
            buffering_size=10,
            buffering_interval=400,
            compression_format="GZIP",
        ),
        processing_configuration=aws.kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationArgs(
            enabled=True,
            processors=[aws.kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorArgs(
                type="Lambda",
                parameters=[aws.kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorParameterArgs(
                    parameter_name="LambdaArn",
                    parameter_value=f"{aws_lambda_function['lambda_processor']['arn']}:$LATEST",
                )],
            )],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCluster = new Aws.ElasticSearch.Domain("testCluster");

    var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new()
    {
        Destination = "elasticsearch",
        ElasticsearchConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationArgs
        {
            DomainArn = testCluster.Arn,
            RoleArn = aws_iam_role.Firehose_role.Arn,
            IndexName = "test",
            TypeName = "test",
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs
            {
                RoleArn = aws_iam_role.Firehose_role.Arn,
                BucketArn = aws_s3_bucket.Bucket.Arn,
                BufferingSize = 10,
                BufferingInterval = 400,
                CompressionFormat = "GZIP",
            },
            ProcessingConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationArgs
            {
                Enabled = true,
                Processors = new[]
                {
                    new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorArgs
                    {
                        Type = "Lambda",
                        Parameters = new[]
                        {
                            new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorParameterArgs
                            {
                                ParameterName = "LambdaArn",
                                ParameterValue = $"{aws_lambda_function.Lambda_processor.Arn}:$LATEST",
                            },
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticsearch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testCluster, err := elasticsearch.NewDomain(ctx, "testCluster", nil)
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("elasticsearch"),
			ElasticsearchConfiguration: &kinesis.FirehoseDeliveryStreamElasticsearchConfigurationArgs{
				DomainArn: testCluster.Arn,
				RoleArn:   pulumi.Any(aws_iam_role.Firehose_role.Arn),
				IndexName: pulumi.String("test"),
				TypeName:  pulumi.String("test"),
				S3Configuration: &kinesis.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs{
					RoleArn:           pulumi.Any(aws_iam_role.Firehose_role.Arn),
					BucketArn:         pulumi.Any(aws_s3_bucket.Bucket.Arn),
					BufferingSize:     pulumi.Int(10),
					BufferingInterval: pulumi.Int(400),
					CompressionFormat: pulumi.String("GZIP"),
				},
				ProcessingConfiguration: &kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationArgs{
					Enabled: pulumi.Bool(true),
					Processors: kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorArray{
						&kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorArgs{
							Type: pulumi.String("Lambda"),
							Parameters: kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorParameterArray{
								&kinesis.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorParameterArgs{
									ParameterName:  pulumi.String("LambdaArn"),
									ParameterValue: pulumi.String(fmt.Sprintf("%v:$LATEST", aws_lambda_function.Lambda_processor.Arn)),
								},
							},
						},
					},
				},
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
import com.pulumi.aws.elasticsearch.Domain;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamElasticsearchConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationArgs;
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
        var testCluster = new Domain("testCluster");

        var testStream = new FirehoseDeliveryStream("testStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("elasticsearch")
            .elasticsearchConfiguration(FirehoseDeliveryStreamElasticsearchConfigurationArgs.builder()
                .domainArn(testCluster.arn())
                .roleArn(aws_iam_role.firehose_role().arn())
                .indexName("test")
                .typeName("test")
                .s3Configuration(FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs.builder()
                    .roleArn(aws_iam_role.firehose_role().arn())
                    .bucketArn(aws_s3_bucket.bucket().arn())
                    .bufferingSize(10)
                    .bufferingInterval(400)
                    .compressionFormat("GZIP")
                    .build())
                .processingConfiguration(FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationArgs.builder()
                    .enabled("true")
                    .processors(FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorArgs.builder()
                        .type("Lambda")
                        .parameters(FirehoseDeliveryStreamElasticsearchConfigurationProcessingConfigurationProcessorParameterArgs.builder()
                            .parameterName("LambdaArn")
                            .parameterValue(String.format("%s:$LATEST", aws_lambda_function.lambda_processor().arn()))
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testCluster:
    type: aws:elasticsearch:Domain
  testStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: elasticsearch
      elasticsearchConfiguration:
        domainArn: ${testCluster.arn}
        roleArn: ${aws_iam_role.firehose_role.arn}
        indexName: test
        typeName: test
        s3Configuration:
          roleArn: ${aws_iam_role.firehose_role.arn}
          bucketArn: ${aws_s3_bucket.bucket.arn}
          bufferingSize: 10
          bufferingInterval: 400
          compressionFormat: GZIP
        processingConfiguration:
          enabled: 'true'
          processors:
            - type: Lambda
              parameters:
                - parameterName: LambdaArn
                  parameterValue: ${aws_lambda_function.lambda_processor.arn}:$LATEST
```
{{% /example %}}
{{% example %}}
### Elasticsearch Destination With VPC

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCluster = new aws.elasticsearch.Domain("testCluster", {
    clusterConfig: {
        instanceCount: 2,
        zoneAwarenessEnabled: true,
        instanceType: "t2.small.elasticsearch",
    },
    ebsOptions: {
        ebsEnabled: true,
        volumeSize: 10,
    },
    vpcOptions: {
        securityGroupIds: [aws_security_group.first.id],
        subnetIds: [
            aws_subnet.first.id,
            aws_subnet.second.id,
        ],
    },
});
const firehose-elasticsearchPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            effect: "Allow",
            actions: ["es:*"],
            resources: [
                testCluster.arn,
                pulumi.interpolate`${testCluster.arn}/*`,
            ],
        },
        {
            effect: "Allow",
            actions: [
                "ec2:DescribeVpcs",
                "ec2:DescribeVpcAttribute",
                "ec2:DescribeSubnets",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeNetworkInterfaces",
                "ec2:CreateNetworkInterface",
                "ec2:CreateNetworkInterfacePermission",
                "ec2:DeleteNetworkInterface",
            ],
            resources: ["*"],
        },
    ],
});
const firehose_elasticsearchRolePolicy = new aws.iam.RolePolicy("firehose-elasticsearchRolePolicy", {
    role: aws_iam_role.firehose.id,
    policy: firehose_elasticsearchPolicyDocument.apply(firehose_elasticsearchPolicyDocument => firehose_elasticsearchPolicyDocument.json),
});
const test = new aws.kinesis.FirehoseDeliveryStream("test", {
    destination: "elasticsearch",
    elasticsearchConfiguration: {
        domainArn: testCluster.arn,
        roleArn: aws_iam_role.firehose.arn,
        indexName: "test",
        typeName: "test",
        s3Configuration: {
            roleArn: aws_iam_role.firehose.arn,
            bucketArn: aws_s3_bucket.bucket.arn,
        },
        vpcConfig: {
            subnetIds: [
                aws_subnet.first.id,
                aws_subnet.second.id,
            ],
            securityGroupIds: [aws_security_group.first.id],
            roleArn: aws_iam_role.firehose.arn,
        },
    },
}, {
    dependsOn: [firehose_elasticsearchRolePolicy],
});
```
```python
import pulumi
import pulumi_aws as aws

test_cluster = aws.elasticsearch.Domain("testCluster",
    cluster_config=aws.elasticsearch.DomainClusterConfigArgs(
        instance_count=2,
        zone_awareness_enabled=True,
        instance_type="t2.small.elasticsearch",
    ),
    ebs_options=aws.elasticsearch.DomainEbsOptionsArgs(
        ebs_enabled=True,
        volume_size=10,
    ),
    vpc_options=aws.elasticsearch.DomainVpcOptionsArgs(
        security_group_ids=[aws_security_group["first"]["id"]],
        subnet_ids=[
            aws_subnet["first"]["id"],
            aws_subnet["second"]["id"],
        ],
    ))
firehose_elasticsearch_policy_document = aws.iam.get_policy_document_output(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=["es:*"],
        resources=[
            test_cluster.arn,
            test_cluster.arn.apply(lambda arn: f"{arn}/*"),
        ],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=[
            "ec2:DescribeVpcs",
            "ec2:DescribeVpcAttribute",
            "ec2:DescribeSubnets",
            "ec2:DescribeSecurityGroups",
            "ec2:DescribeNetworkInterfaces",
            "ec2:CreateNetworkInterface",
            "ec2:CreateNetworkInterfacePermission",
            "ec2:DeleteNetworkInterface",
        ],
        resources=["*"],
    ),
])
firehose_elasticsearch_role_policy = aws.iam.RolePolicy("firehose-elasticsearchRolePolicy",
    role=aws_iam_role["firehose"]["id"],
    policy=firehose_elasticsearch_policy_document.json)
test = aws.kinesis.FirehoseDeliveryStream("test",
    destination="elasticsearch",
    elasticsearch_configuration=aws.kinesis.FirehoseDeliveryStreamElasticsearchConfigurationArgs(
        domain_arn=test_cluster.arn,
        role_arn=aws_iam_role["firehose"]["arn"],
        index_name="test",
        type_name="test",
        s3_configuration=aws.kinesis.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs(
            role_arn=aws_iam_role["firehose"]["arn"],
            bucket_arn=aws_s3_bucket["bucket"]["arn"],
        ),
        vpc_config=aws.kinesis.FirehoseDeliveryStreamElasticsearchConfigurationVpcConfigArgs(
            subnet_ids=[
                aws_subnet["first"]["id"],
                aws_subnet["second"]["id"],
            ],
            security_group_ids=[aws_security_group["first"]["id"]],
            role_arn=aws_iam_role["firehose"]["arn"],
        ),
    ),
    opts=pulumi.ResourceOptions(depends_on=[firehose_elasticsearch_role_policy]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCluster = new Aws.ElasticSearch.Domain("testCluster", new()
    {
        ClusterConfig = new Aws.ElasticSearch.Inputs.DomainClusterConfigArgs
        {
            InstanceCount = 2,
            ZoneAwarenessEnabled = true,
            InstanceType = "t2.small.elasticsearch",
        },
        EbsOptions = new Aws.ElasticSearch.Inputs.DomainEbsOptionsArgs
        {
            EbsEnabled = true,
            VolumeSize = 10,
        },
        VpcOptions = new Aws.ElasticSearch.Inputs.DomainVpcOptionsArgs
        {
            SecurityGroupIds = new[]
            {
                aws_security_group.First.Id,
            },
            SubnetIds = new[]
            {
                aws_subnet.First.Id,
                aws_subnet.Second.Id,
            },
        },
    });

    var firehose_elasticsearchPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "es:*",
                },
                Resources = new[]
                {
                    testCluster.Arn,
                    $"{testCluster.Arn}/*",
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "ec2:DescribeVpcs",
                    "ec2:DescribeVpcAttribute",
                    "ec2:DescribeSubnets",
                    "ec2:DescribeSecurityGroups",
                    "ec2:DescribeNetworkInterfaces",
                    "ec2:CreateNetworkInterface",
                    "ec2:CreateNetworkInterfacePermission",
                    "ec2:DeleteNetworkInterface",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var firehose_elasticsearchRolePolicy = new Aws.Iam.RolePolicy("firehose-elasticsearchRolePolicy", new()
    {
        Role = aws_iam_role.Firehose.Id,
        Policy = firehose_elasticsearchPolicyDocument.Apply(firehose_elasticsearchPolicyDocument => firehose_elasticsearchPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json)),
    });

    var test = new Aws.Kinesis.FirehoseDeliveryStream("test", new()
    {
        Destination = "elasticsearch",
        ElasticsearchConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationArgs
        {
            DomainArn = testCluster.Arn,
            RoleArn = aws_iam_role.Firehose.Arn,
            IndexName = "test",
            TypeName = "test",
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs
            {
                RoleArn = aws_iam_role.Firehose.Arn,
                BucketArn = aws_s3_bucket.Bucket.Arn,
            },
            VpcConfig = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamElasticsearchConfigurationVpcConfigArgs
            {
                SubnetIds = new[]
                {
                    aws_subnet.First.Id,
                    aws_subnet.Second.Id,
                },
                SecurityGroupIds = new[]
                {
                    aws_security_group.First.Id,
                },
                RoleArn = aws_iam_role.Firehose.Arn,
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            firehose_elasticsearchRolePolicy,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticsearch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testCluster, err := elasticsearch.NewDomain(ctx, "testCluster", &elasticsearch.DomainArgs{
			ClusterConfig: &elasticsearch.DomainClusterConfigArgs{
				InstanceCount:        pulumi.Int(2),
				ZoneAwarenessEnabled: pulumi.Bool(true),
				InstanceType:         pulumi.String("t2.small.elasticsearch"),
			},
			EbsOptions: &elasticsearch.DomainEbsOptionsArgs{
				EbsEnabled: pulumi.Bool(true),
				VolumeSize: pulumi.Int(10),
			},
			VpcOptions: &elasticsearch.DomainVpcOptionsArgs{
				SecurityGroupIds: pulumi.StringArray{
					aws_security_group.First.Id,
				},
				SubnetIds: pulumi.StringArray{
					aws_subnet.First.Id,
					aws_subnet.Second.Id,
				},
			},
		})
		if err != nil {
			return err
		}
		firehose_elasticsearchPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("es:*"),
					},
					Resources: pulumi.StringArray{
						testCluster.Arn,
						testCluster.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("ec2:DescribeVpcs"),
						pulumi.String("ec2:DescribeVpcAttribute"),
						pulumi.String("ec2:DescribeSubnets"),
						pulumi.String("ec2:DescribeSecurityGroups"),
						pulumi.String("ec2:DescribeNetworkInterfaces"),
						pulumi.String("ec2:CreateNetworkInterface"),
						pulumi.String("ec2:CreateNetworkInterfacePermission"),
						pulumi.String("ec2:DeleteNetworkInterface"),
					},
					Resources: pulumi.StringArray{
						pulumi.String("*"),
					},
				},
			},
		}, nil)
		_, err = iam.NewRolePolicy(ctx, "firehose-elasticsearchRolePolicy", &iam.RolePolicyArgs{
			Role: pulumi.Any(aws_iam_role.Firehose.Id),
			Policy: firehose_elasticsearchPolicyDocument.ApplyT(func(firehose_elasticsearchPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &firehose_elasticsearchPolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
		})
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "test", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("elasticsearch"),
			ElasticsearchConfiguration: &kinesis.FirehoseDeliveryStreamElasticsearchConfigurationArgs{
				DomainArn: testCluster.Arn,
				RoleArn:   pulumi.Any(aws_iam_role.Firehose.Arn),
				IndexName: pulumi.String("test"),
				TypeName:  pulumi.String("test"),
				S3Configuration: &kinesis.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs{
					RoleArn:   pulumi.Any(aws_iam_role.Firehose.Arn),
					BucketArn: pulumi.Any(aws_s3_bucket.Bucket.Arn),
				},
				VpcConfig: &kinesis.FirehoseDeliveryStreamElasticsearchConfigurationVpcConfigArgs{
					SubnetIds: pulumi.StringArray{
						aws_subnet.First.Id,
						aws_subnet.Second.Id,
					},
					SecurityGroupIds: pulumi.StringArray{
						aws_security_group.First.Id,
					},
					RoleArn: pulumi.Any(aws_iam_role.Firehose.Arn),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			firehose_elasticsearchRolePolicy,
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
import com.pulumi.aws.elasticsearch.Domain;
import com.pulumi.aws.elasticsearch.DomainArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainClusterConfigArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainEbsOptionsArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainVpcOptionsArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamElasticsearchConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamElasticsearchConfigurationVpcConfigArgs;
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
        var testCluster = new Domain("testCluster", DomainArgs.builder()        
            .clusterConfig(DomainClusterConfigArgs.builder()
                .instanceCount(2)
                .zoneAwarenessEnabled(true)
                .instanceType("t2.small.elasticsearch")
                .build())
            .ebsOptions(DomainEbsOptionsArgs.builder()
                .ebsEnabled(true)
                .volumeSize(10)
                .build())
            .vpcOptions(DomainVpcOptionsArgs.builder()
                .securityGroupIds(aws_security_group.first().id())
                .subnetIds(                
                    aws_subnet.first().id(),
                    aws_subnet.second().id())
                .build())
            .build());

        final var firehose-elasticsearchPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions("es:*")
                    .resources(                    
                        testCluster.arn(),
                        testCluster.arn().applyValue(arn -> String.format("%s/*", arn)))
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions(                    
                        "ec2:DescribeVpcs",
                        "ec2:DescribeVpcAttribute",
                        "ec2:DescribeSubnets",
                        "ec2:DescribeSecurityGroups",
                        "ec2:DescribeNetworkInterfaces",
                        "ec2:CreateNetworkInterface",
                        "ec2:CreateNetworkInterfacePermission",
                        "ec2:DeleteNetworkInterface")
                    .resources("*")
                    .build())
            .build());

        var firehose_elasticsearchRolePolicy = new RolePolicy("firehose-elasticsearchRolePolicy", RolePolicyArgs.builder()        
            .role(aws_iam_role.firehose().id())
            .policy(firehose_elasticsearchPolicyDocument.applyValue(firehose_elasticsearchPolicyDocument -> firehose_elasticsearchPolicyDocument.json()))
            .build());

        var test = new FirehoseDeliveryStream("test", FirehoseDeliveryStreamArgs.builder()        
            .destination("elasticsearch")
            .elasticsearchConfiguration(FirehoseDeliveryStreamElasticsearchConfigurationArgs.builder()
                .domainArn(testCluster.arn())
                .roleArn(aws_iam_role.firehose().arn())
                .indexName("test")
                .typeName("test")
                .s3Configuration(FirehoseDeliveryStreamElasticsearchConfigurationS3ConfigurationArgs.builder()
                    .roleArn(aws_iam_role.firehose().arn())
                    .bucketArn(aws_s3_bucket.bucket().arn())
                    .build())
                .vpcConfig(FirehoseDeliveryStreamElasticsearchConfigurationVpcConfigArgs.builder()
                    .subnetIds(                    
                        aws_subnet.first().id(),
                        aws_subnet.second().id())
                    .securityGroupIds(aws_security_group.first().id())
                    .roleArn(aws_iam_role.firehose().arn())
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(firehose_elasticsearchRolePolicy)
                .build());

    }
}
```
```yaml
resources:
  testCluster:
    type: aws:elasticsearch:Domain
    properties:
      clusterConfig:
        instanceCount: 2
        zoneAwarenessEnabled: true
        instanceType: t2.small.elasticsearch
      ebsOptions:
        ebsEnabled: true
        volumeSize: 10
      vpcOptions:
        securityGroupIds:
          - ${aws_security_group.first.id}
        subnetIds:
          - ${aws_subnet.first.id}
          - ${aws_subnet.second.id}
  firehose-elasticsearchRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${aws_iam_role.firehose.id}
      policy: ${["firehose-elasticsearchPolicyDocument"].json}
  test:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: elasticsearch
      elasticsearchConfiguration:
        domainArn: ${testCluster.arn}
        roleArn: ${aws_iam_role.firehose.arn}
        indexName: test
        typeName: test
        s3Configuration:
          roleArn: ${aws_iam_role.firehose.arn}
          bucketArn: ${aws_s3_bucket.bucket.arn}
        vpcConfig:
          subnetIds:
            - ${aws_subnet.first.id}
            - ${aws_subnet.second.id}
          securityGroupIds:
            - ${aws_security_group.first.id}
          roleArn: ${aws_iam_role.firehose.arn}
    options:
      dependson:
        - ${["firehose-elasticsearchRolePolicy"]}
variables:
  firehose-elasticsearchPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - es:*
            resources:
              - ${testCluster.arn}
              - ${testCluster.arn}/*
          - effect: Allow
            actions:
              - ec2:DescribeVpcs
              - ec2:DescribeVpcAttribute
              - ec2:DescribeSubnets
              - ec2:DescribeSecurityGroups
              - ec2:DescribeNetworkInterfaces
              - ec2:CreateNetworkInterface
              - ec2:CreateNetworkInterfacePermission
              - ec2:DeleteNetworkInterface
            resources:
              - '*'
```
{{% /example %}}
{{% example %}}
### OpenSearch Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCluster = new aws.opensearch.Domain("testCluster", {});
const testStream = new aws.kinesis.FirehoseDeliveryStream("testStream", {
    destination: "opensearch",
    opensearchConfiguration: {
        domainArn: testCluster.arn,
        roleArn: aws_iam_role.firehose_role.arn,
        indexName: "test",
        s3Configuration: {
            roleArn: aws_iam_role.firehose_role.arn,
            bucketArn: aws_s3_bucket.bucket.arn,
            bufferingSize: 10,
            bufferingInterval: 400,
            compressionFormat: "GZIP",
        },
        processingConfiguration: {
            enabled: true,
            processors: [{
                type: "Lambda",
                parameters: [{
                    parameterName: "LambdaArn",
                    parameterValue: `${aws_lambda_function.lambda_processor.arn}:$LATEST`,
                }],
            }],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_cluster = aws.opensearch.Domain("testCluster")
test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="opensearch",
    opensearch_configuration=aws.kinesis.FirehoseDeliveryStreamOpensearchConfigurationArgs(
        domain_arn=test_cluster.arn,
        role_arn=aws_iam_role["firehose_role"]["arn"],
        index_name="test",
        s3_configuration=aws.kinesis.FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs(
            role_arn=aws_iam_role["firehose_role"]["arn"],
            bucket_arn=aws_s3_bucket["bucket"]["arn"],
            buffering_size=10,
            buffering_interval=400,
            compression_format="GZIP",
        ),
        processing_configuration=aws.kinesis.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationArgs(
            enabled=True,
            processors=[aws.kinesis.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessorArgs(
                type="Lambda",
                parameters=[aws.kinesis.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessorParameterArgs(
                    parameter_name="LambdaArn",
                    parameter_value=f"{aws_lambda_function['lambda_processor']['arn']}:$LATEST",
                )],
            )],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCluster = new Aws.OpenSearch.Domain("testCluster");

    var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new()
    {
        Destination = "opensearch",
        OpensearchConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchConfigurationArgs
        {
            DomainArn = testCluster.Arn,
            RoleArn = aws_iam_role.Firehose_role.Arn,
            IndexName = "test",
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs
            {
                RoleArn = aws_iam_role.Firehose_role.Arn,
                BucketArn = aws_s3_bucket.Bucket.Arn,
                BufferingSize = 10,
                BufferingInterval = 400,
                CompressionFormat = "GZIP",
            },
            ProcessingConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationArgs
            {
                Enabled = true,
                Processors = new[]
                {
                    new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessorArgs
                    {
                        Type = "Lambda",
                        Parameters = new[]
                        {
                            new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessorParameterArgs
                            {
                                ParameterName = "LambdaArn",
                                ParameterValue = $"{aws_lambda_function.Lambda_processor.Arn}:$LATEST",
                            },
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testCluster, err := opensearch.NewDomain(ctx, "testCluster", nil)
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("opensearch"),
			OpensearchConfiguration: &kinesis.FirehoseDeliveryStreamOpensearchConfigurationArgs{
				DomainArn: testCluster.Arn,
				RoleArn:   pulumi.Any(aws_iam_role.Firehose_role.Arn),
				IndexName: pulumi.String("test"),
				S3Configuration: &kinesis.FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs{
					RoleArn:           pulumi.Any(aws_iam_role.Firehose_role.Arn),
					BucketArn:         pulumi.Any(aws_s3_bucket.Bucket.Arn),
					BufferingSize:     pulumi.Int(10),
					BufferingInterval: pulumi.Int(400),
					CompressionFormat: pulumi.String("GZIP"),
				},
				ProcessingConfiguration: &kinesis.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationArgs{
					Enabled: pulumi.Bool(true),
					Processors: kinesis.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessorArray{
						&kinesis.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessorArgs{
							Type: pulumi.String("Lambda"),
							Parameters: kinesis.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessorParameterArray{
								&kinesis.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessorParameterArgs{
									ParameterName:  pulumi.String("LambdaArn"),
									ParameterValue: pulumi.String(fmt.Sprintf("%v:$LATEST", aws_lambda_function.Lambda_processor.Arn)),
								},
							},
						},
					},
				},
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
import com.pulumi.aws.opensearch.Domain;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamOpensearchConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationArgs;
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
        var testCluster = new Domain("testCluster");

        var testStream = new FirehoseDeliveryStream("testStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("opensearch")
            .opensearchConfiguration(FirehoseDeliveryStreamOpensearchConfigurationArgs.builder()
                .domainArn(testCluster.arn())
                .roleArn(aws_iam_role.firehose_role().arn())
                .indexName("test")
                .s3Configuration(FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs.builder()
                    .roleArn(aws_iam_role.firehose_role().arn())
                    .bucketArn(aws_s3_bucket.bucket().arn())
                    .bufferingSize(10)
                    .bufferingInterval(400)
                    .compressionFormat("GZIP")
                    .build())
                .processingConfiguration(FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationArgs.builder()
                    .enabled("true")
                    .processors(FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessorArgs.builder()
                        .type("Lambda")
                        .parameters(FirehoseDeliveryStreamOpensearchConfigurationProcessingConfigurationProcessorParameterArgs.builder()
                            .parameterName("LambdaArn")
                            .parameterValue(String.format("%s:$LATEST", aws_lambda_function.lambda_processor().arn()))
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testCluster:
    type: aws:opensearch:Domain
  testStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: opensearch
      opensearchConfiguration:
        domainArn: ${testCluster.arn}
        roleArn: ${aws_iam_role.firehose_role.arn}
        indexName: test
        s3Configuration:
          roleArn: ${aws_iam_role.firehose_role.arn}
          bucketArn: ${aws_s3_bucket.bucket.arn}
          bufferingSize: 10
          bufferingInterval: 400
          compressionFormat: GZIP
        processingConfiguration:
          enabled: 'true'
          processors:
            - type: Lambda
              parameters:
                - parameterName: LambdaArn
                  parameterValue: ${aws_lambda_function.lambda_processor.arn}:$LATEST
```
{{% /example %}}
{{% example %}}
### OpenSearch Destination With VPC

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCluster = new aws.opensearch.Domain("testCluster", {
    clusterConfig: {
        instanceCount: 2,
        zoneAwarenessEnabled: true,
        instanceType: "m4.large.search",
    },
    ebsOptions: {
        ebsEnabled: true,
        volumeSize: 10,
    },
    vpcOptions: {
        securityGroupIds: [aws_security_group.first.id],
        subnetIds: [
            aws_subnet.first.id,
            aws_subnet.second.id,
        ],
    },
});
const firehose_opensearch = new aws.iam.RolePolicy("firehose-opensearch", {
    role: aws_iam_role.firehose.id,
    policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "es:*"
      ],
      "Resource": [
        "${testCluster.arn}",
        "${testCluster.arn}/*"
      ]
        },
        {
          "Effect": "Allow",
          "Action": [
            "ec2:DescribeVpcs",
            "ec2:DescribeVpcAttribute",
            "ec2:DescribeSubnets",
            "ec2:DescribeSecurityGroups",
            "ec2:DescribeNetworkInterfaces",
            "ec2:CreateNetworkInterface",
            "ec2:CreateNetworkInterfacePermission",
            "ec2:DeleteNetworkInterface"
          ],
          "Resource": [
            "*"
          ]
        }
  ]
}
`,
});
const test = new aws.kinesis.FirehoseDeliveryStream("test", {
    destination: "opensearch",
    opensearchConfiguration: {
        domainArn: testCluster.arn,
        roleArn: aws_iam_role.firehose.arn,
        indexName: "test",
        s3Configuration: {
            roleArn: aws_iam_role.firehose.arn,
            bucketArn: aws_s3_bucket.bucket.arn,
        },
        vpcConfig: {
            subnetIds: [
                aws_subnet.first.id,
                aws_subnet.second.id,
            ],
            securityGroupIds: [aws_security_group.first.id],
            roleArn: aws_iam_role.firehose.arn,
        },
    },
}, {
    dependsOn: [firehose_opensearch],
});
```
```python
import pulumi
import pulumi_aws as aws

test_cluster = aws.opensearch.Domain("testCluster",
    cluster_config=aws.opensearch.DomainClusterConfigArgs(
        instance_count=2,
        zone_awareness_enabled=True,
        instance_type="m4.large.search",
    ),
    ebs_options=aws.opensearch.DomainEbsOptionsArgs(
        ebs_enabled=True,
        volume_size=10,
    ),
    vpc_options=aws.opensearch.DomainVpcOptionsArgs(
        security_group_ids=[aws_security_group["first"]["id"]],
        subnet_ids=[
            aws_subnet["first"]["id"],
            aws_subnet["second"]["id"],
        ],
    ))
firehose_opensearch = aws.iam.RolePolicy("firehose-opensearch",
    role=aws_iam_role["firehose"]["id"],
    policy=pulumi.Output.all(test_cluster.arn, test_cluster.arn).apply(lambda testClusterArn, testClusterArn1: f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Effect": "Allow",
      "Action": [
        "es:*"
      ],
      "Resource": [
        "{test_cluster_arn}",
        "{test_cluster_arn1}/*"
      ]
        }},
        {{
          "Effect": "Allow",
          "Action": [
            "ec2:DescribeVpcs",
            "ec2:DescribeVpcAttribute",
            "ec2:DescribeSubnets",
            "ec2:DescribeSecurityGroups",
            "ec2:DescribeNetworkInterfaces",
            "ec2:CreateNetworkInterface",
            "ec2:CreateNetworkInterfacePermission",
            "ec2:DeleteNetworkInterface"
          ],
          "Resource": [
            "*"
          ]
        }}
  ]
}}
"""))
test = aws.kinesis.FirehoseDeliveryStream("test",
    destination="opensearch",
    opensearch_configuration=aws.kinesis.FirehoseDeliveryStreamOpensearchConfigurationArgs(
        domain_arn=test_cluster.arn,
        role_arn=aws_iam_role["firehose"]["arn"],
        index_name="test",
        s3_configuration=aws.kinesis.FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs(
            role_arn=aws_iam_role["firehose"]["arn"],
            bucket_arn=aws_s3_bucket["bucket"]["arn"],
        ),
        vpc_config=aws.kinesis.FirehoseDeliveryStreamOpensearchConfigurationVpcConfigArgs(
            subnet_ids=[
                aws_subnet["first"]["id"],
                aws_subnet["second"]["id"],
            ],
            security_group_ids=[aws_security_group["first"]["id"]],
            role_arn=aws_iam_role["firehose"]["arn"],
        ),
    ),
    opts=pulumi.ResourceOptions(depends_on=[firehose_opensearch]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCluster = new Aws.OpenSearch.Domain("testCluster", new()
    {
        ClusterConfig = new Aws.OpenSearch.Inputs.DomainClusterConfigArgs
        {
            InstanceCount = 2,
            ZoneAwarenessEnabled = true,
            InstanceType = "m4.large.search",
        },
        EbsOptions = new Aws.OpenSearch.Inputs.DomainEbsOptionsArgs
        {
            EbsEnabled = true,
            VolumeSize = 10,
        },
        VpcOptions = new Aws.OpenSearch.Inputs.DomainVpcOptionsArgs
        {
            SecurityGroupIds = new[]
            {
                aws_security_group.First.Id,
            },
            SubnetIds = new[]
            {
                aws_subnet.First.Id,
                aws_subnet.Second.Id,
            },
        },
    });

    var firehose_opensearch = new Aws.Iam.RolePolicy("firehose-opensearch", new()
    {
        Role = aws_iam_role.Firehose.Id,
        Policy = Output.Tuple(testCluster.Arn, testCluster.Arn).Apply(values =>
        {
            var testClusterArn = values.Item1;
            var testClusterArn1 = values.Item2;
            return @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Effect"": ""Allow"",
      ""Action"": [
        ""es:*""
      ],
      ""Resource"": [
        ""{testClusterArn}"",
        ""{testClusterArn1}/*""
      ]
        }},
        {{
          ""Effect"": ""Allow"",
          ""Action"": [
            ""ec2:DescribeVpcs"",
            ""ec2:DescribeVpcAttribute"",
            ""ec2:DescribeSubnets"",
            ""ec2:DescribeSecurityGroups"",
            ""ec2:DescribeNetworkInterfaces"",
            ""ec2:CreateNetworkInterface"",
            ""ec2:CreateNetworkInterfacePermission"",
            ""ec2:DeleteNetworkInterface""
          ],
          ""Resource"": [
            ""*""
          ]
        }}
  ]
}}
";
        }),
    });

    var test = new Aws.Kinesis.FirehoseDeliveryStream("test", new()
    {
        Destination = "opensearch",
        OpensearchConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchConfigurationArgs
        {
            DomainArn = testCluster.Arn,
            RoleArn = aws_iam_role.Firehose.Arn,
            IndexName = "test",
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs
            {
                RoleArn = aws_iam_role.Firehose.Arn,
                BucketArn = aws_s3_bucket.Bucket.Arn,
            },
            VpcConfig = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchConfigurationVpcConfigArgs
            {
                SubnetIds = new[]
                {
                    aws_subnet.First.Id,
                    aws_subnet.Second.Id,
                },
                SecurityGroupIds = new[]
                {
                    aws_security_group.First.Id,
                },
                RoleArn = aws_iam_role.Firehose.Arn,
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            firehose_opensearch,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testCluster, err := opensearch.NewDomain(ctx, "testCluster", &opensearch.DomainArgs{
			ClusterConfig: &opensearch.DomainClusterConfigArgs{
				InstanceCount:        pulumi.Int(2),
				ZoneAwarenessEnabled: pulumi.Bool(true),
				InstanceType:         pulumi.String("m4.large.search"),
			},
			EbsOptions: &opensearch.DomainEbsOptionsArgs{
				EbsEnabled: pulumi.Bool(true),
				VolumeSize: pulumi.Int(10),
			},
			VpcOptions: &opensearch.DomainVpcOptionsArgs{
				SecurityGroupIds: pulumi.StringArray{
					aws_security_group.First.Id,
				},
				SubnetIds: pulumi.StringArray{
					aws_subnet.First.Id,
					aws_subnet.Second.Id,
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "firehose-opensearch", &iam.RolePolicyArgs{
			Role: pulumi.Any(aws_iam_role.Firehose.Id),
			Policy: pulumi.All(testCluster.Arn, testCluster.Arn).ApplyT(func(_args []interface{}) (string, error) {
				testClusterArn := _args[0].(string)
				testClusterArn1 := _args[1].(string)
				return fmt.Sprintf(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "es:*"
      ],
      "Resource": [
        "%v",
        "%v/*"
      ]
        },
        {
          "Effect": "Allow",
          "Action": [
            "ec2:DescribeVpcs",
            "ec2:DescribeVpcAttribute",
            "ec2:DescribeSubnets",
            "ec2:DescribeSecurityGroups",
            "ec2:DescribeNetworkInterfaces",
            "ec2:CreateNetworkInterface",
            "ec2:CreateNetworkInterfacePermission",
            "ec2:DeleteNetworkInterface"
          ],
          "Resource": [
            "*"
          ]
        }
  ]
}
`, testClusterArn, testClusterArn1), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "test", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("opensearch"),
			OpensearchConfiguration: &kinesis.FirehoseDeliveryStreamOpensearchConfigurationArgs{
				DomainArn: testCluster.Arn,
				RoleArn:   pulumi.Any(aws_iam_role.Firehose.Arn),
				IndexName: pulumi.String("test"),
				S3Configuration: &kinesis.FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs{
					RoleArn:   pulumi.Any(aws_iam_role.Firehose.Arn),
					BucketArn: pulumi.Any(aws_s3_bucket.Bucket.Arn),
				},
				VpcConfig: &kinesis.FirehoseDeliveryStreamOpensearchConfigurationVpcConfigArgs{
					SubnetIds: pulumi.StringArray{
						aws_subnet.First.Id,
						aws_subnet.Second.Id,
					},
					SecurityGroupIds: pulumi.StringArray{
						aws_security_group.First.Id,
					},
					RoleArn: pulumi.Any(aws_iam_role.Firehose.Arn),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			firehose_opensearch,
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
import com.pulumi.aws.opensearch.Domain;
import com.pulumi.aws.opensearch.DomainArgs;
import com.pulumi.aws.opensearch.inputs.DomainClusterConfigArgs;
import com.pulumi.aws.opensearch.inputs.DomainEbsOptionsArgs;
import com.pulumi.aws.opensearch.inputs.DomainVpcOptionsArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamOpensearchConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamOpensearchConfigurationVpcConfigArgs;
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
        var testCluster = new Domain("testCluster", DomainArgs.builder()        
            .clusterConfig(DomainClusterConfigArgs.builder()
                .instanceCount(2)
                .zoneAwarenessEnabled(true)
                .instanceType("m4.large.search")
                .build())
            .ebsOptions(DomainEbsOptionsArgs.builder()
                .ebsEnabled(true)
                .volumeSize(10)
                .build())
            .vpcOptions(DomainVpcOptionsArgs.builder()
                .securityGroupIds(aws_security_group.first().id())
                .subnetIds(                
                    aws_subnet.first().id(),
                    aws_subnet.second().id())
                .build())
            .build());

        var firehose_opensearch = new RolePolicy("firehose-opensearch", RolePolicyArgs.builder()        
            .role(aws_iam_role.firehose().id())
            .policy(Output.tuple(testCluster.arn(), testCluster.arn()).applyValue(values -> {
                var testClusterArn = values.t1;
                var testClusterArn1 = values.t2;
                return """
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "es:*"
      ],
      "Resource": [
        "%s",
        "%s/*"
      ]
        },
        {
          "Effect": "Allow",
          "Action": [
            "ec2:DescribeVpcs",
            "ec2:DescribeVpcAttribute",
            "ec2:DescribeSubnets",
            "ec2:DescribeSecurityGroups",
            "ec2:DescribeNetworkInterfaces",
            "ec2:CreateNetworkInterface",
            "ec2:CreateNetworkInterfacePermission",
            "ec2:DeleteNetworkInterface"
          ],
          "Resource": [
            "*"
          ]
        }
  ]
}
", testClusterArn,testClusterArn1);
            }))
            .build());

        var test = new FirehoseDeliveryStream("test", FirehoseDeliveryStreamArgs.builder()        
            .destination("opensearch")
            .opensearchConfiguration(FirehoseDeliveryStreamOpensearchConfigurationArgs.builder()
                .domainArn(testCluster.arn())
                .roleArn(aws_iam_role.firehose().arn())
                .indexName("test")
                .s3Configuration(FirehoseDeliveryStreamOpensearchConfigurationS3ConfigurationArgs.builder()
                    .roleArn(aws_iam_role.firehose().arn())
                    .bucketArn(aws_s3_bucket.bucket().arn())
                    .build())
                .vpcConfig(FirehoseDeliveryStreamOpensearchConfigurationVpcConfigArgs.builder()
                    .subnetIds(                    
                        aws_subnet.first().id(),
                        aws_subnet.second().id())
                    .securityGroupIds(aws_security_group.first().id())
                    .roleArn(aws_iam_role.firehose().arn())
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(firehose_opensearch)
                .build());

    }
}
```
```yaml
resources:
  testCluster:
    type: aws:opensearch:Domain
    properties:
      clusterConfig:
        instanceCount: 2
        zoneAwarenessEnabled: true
        instanceType: m4.large.search
      ebsOptions:
        ebsEnabled: true
        volumeSize: 10
      vpcOptions:
        securityGroupIds:
          - ${aws_security_group.first.id}
        subnetIds:
          - ${aws_subnet.first.id}
          - ${aws_subnet.second.id}
  firehose-opensearch:
    type: aws:iam:RolePolicy
    properties:
      role: ${aws_iam_role.firehose.id}
      policy: |
        {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Action": [
                "es:*"
              ],
              "Resource": [
                "${testCluster.arn}",
                "${testCluster.arn}/*"
              ]
                },
                {
                  "Effect": "Allow",
                  "Action": [
                    "ec2:DescribeVpcs",
                    "ec2:DescribeVpcAttribute",
                    "ec2:DescribeSubnets",
                    "ec2:DescribeSecurityGroups",
                    "ec2:DescribeNetworkInterfaces",
                    "ec2:CreateNetworkInterface",
                    "ec2:CreateNetworkInterfacePermission",
                    "ec2:DeleteNetworkInterface"
                  ],
                  "Resource": [
                    "*"
                  ]
                }
          ]
        }
  test:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: opensearch
      opensearchConfiguration:
        domainArn: ${testCluster.arn}
        roleArn: ${aws_iam_role.firehose.arn}
        indexName: test
        s3Configuration:
          roleArn: ${aws_iam_role.firehose.arn}
          bucketArn: ${aws_s3_bucket.bucket.arn}
        vpcConfig:
          subnetIds:
            - ${aws_subnet.first.id}
            - ${aws_subnet.second.id}
          securityGroupIds:
            - ${aws_security_group.first.id}
          roleArn: ${aws_iam_role.firehose.arn}
    options:
      dependson:
        - ${["firehose-opensearch"]}
```
{{% /example %}}
{{% example %}}
### OpenSearch Serverless Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCollection = new aws.opensearch.ServerlessCollection("testCollection", {});
const testStream = new aws.kinesis.FirehoseDeliveryStream("testStream", {
    destination: "opensearchserverless",
    opensearchserverlessConfiguration: {
        collectionEndpoint: testCollection.collectionEndpoint,
        roleArn: aws_iam_role.firehose_role.arn,
        indexName: "test",
        s3Configuration: {
            roleArn: aws_iam_role.firehose_role.arn,
            bucketArn: aws_s3_bucket.bucket.arn,
            bufferingSize: 10,
            bufferingInterval: 400,
            compressionFormat: "GZIP",
        },
        processingConfiguration: {
            enabled: true,
            processors: [{
                type: "Lambda",
                parameters: [{
                    parameterName: "LambdaArn",
                    parameterValue: `${aws_lambda_function.lambda_processor.arn}:$LATEST`,
                }],
            }],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_collection = aws.opensearch.ServerlessCollection("testCollection")
test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="opensearchserverless",
    opensearchserverless_configuration=aws.kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationArgs(
        collection_endpoint=test_collection.collection_endpoint,
        role_arn=aws_iam_role["firehose_role"]["arn"],
        index_name="test",
        s3_configuration=aws.kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationArgs(
            role_arn=aws_iam_role["firehose_role"]["arn"],
            bucket_arn=aws_s3_bucket["bucket"]["arn"],
            buffering_size=10,
            buffering_interval=400,
            compression_format="GZIP",
        ),
        processing_configuration=aws.kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationArgs(
            enabled=True,
            processors=[aws.kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationProcessorArgs(
                type="Lambda",
                parameters=[aws.kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationProcessorParameterArgs(
                    parameter_name="LambdaArn",
                    parameter_value=f"{aws_lambda_function['lambda_processor']['arn']}:$LATEST",
                )],
            )],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCollection = new Aws.OpenSearch.ServerlessCollection("testCollection");

    var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new()
    {
        Destination = "opensearchserverless",
        OpensearchserverlessConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchserverlessConfigurationArgs
        {
            CollectionEndpoint = testCollection.CollectionEndpoint,
            RoleArn = aws_iam_role.Firehose_role.Arn,
            IndexName = "test",
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationArgs
            {
                RoleArn = aws_iam_role.Firehose_role.Arn,
                BucketArn = aws_s3_bucket.Bucket.Arn,
                BufferingSize = 10,
                BufferingInterval = 400,
                CompressionFormat = "GZIP",
            },
            ProcessingConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationArgs
            {
                Enabled = true,
                Processors = new[]
                {
                    new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationProcessorArgs
                    {
                        Type = "Lambda",
                        Parameters = new[]
                        {
                            new Aws.Kinesis.Inputs.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationProcessorParameterArgs
                            {
                                ParameterName = "LambdaArn",
                                ParameterValue = $"{aws_lambda_function.Lambda_processor.Arn}:$LATEST",
                            },
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testCollection, err := opensearch.NewServerlessCollection(ctx, "testCollection", nil)
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("opensearchserverless"),
			OpensearchserverlessConfiguration: &kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationArgs{
				CollectionEndpoint: testCollection.CollectionEndpoint,
				RoleArn:            pulumi.Any(aws_iam_role.Firehose_role.Arn),
				IndexName:          pulumi.String("test"),
				S3Configuration: &kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationArgs{
					RoleArn:           pulumi.Any(aws_iam_role.Firehose_role.Arn),
					BucketArn:         pulumi.Any(aws_s3_bucket.Bucket.Arn),
					BufferingSize:     pulumi.Int(10),
					BufferingInterval: pulumi.Int(400),
					CompressionFormat: pulumi.String("GZIP"),
				},
				ProcessingConfiguration: &kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationArgs{
					Enabled: pulumi.Bool(true),
					Processors: kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationProcessorArray{
						&kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationProcessorArgs{
							Type: pulumi.String("Lambda"),
							Parameters: kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationProcessorParameterArray{
								&kinesis.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationProcessorParameterArgs{
									ParameterName:  pulumi.String("LambdaArn"),
									ParameterValue: pulumi.String(fmt.Sprintf("%v:$LATEST", aws_lambda_function.Lambda_processor.Arn)),
								},
							},
						},
					},
				},
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
import com.pulumi.aws.opensearch.ServerlessCollection;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamOpensearchserverlessConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationArgs;
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
        var testCollection = new ServerlessCollection("testCollection");

        var testStream = new FirehoseDeliveryStream("testStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("opensearchserverless")
            .opensearchserverlessConfiguration(FirehoseDeliveryStreamOpensearchserverlessConfigurationArgs.builder()
                .collectionEndpoint(testCollection.collectionEndpoint())
                .roleArn(aws_iam_role.firehose_role().arn())
                .indexName("test")
                .s3Configuration(FirehoseDeliveryStreamOpensearchserverlessConfigurationS3ConfigurationArgs.builder()
                    .roleArn(aws_iam_role.firehose_role().arn())
                    .bucketArn(aws_s3_bucket.bucket().arn())
                    .bufferingSize(10)
                    .bufferingInterval(400)
                    .compressionFormat("GZIP")
                    .build())
                .processingConfiguration(FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationArgs.builder()
                    .enabled("true")
                    .processors(FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationProcessorArgs.builder()
                        .type("Lambda")
                        .parameters(FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfigurationProcessorParameterArgs.builder()
                            .parameterName("LambdaArn")
                            .parameterValue(String.format("%s:$LATEST", aws_lambda_function.lambda_processor().arn()))
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testCollection:
    type: aws:opensearch:ServerlessCollection
  testStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: opensearchserverless
      opensearchserverlessConfiguration:
        collectionEndpoint: ${testCollection.collectionEndpoint}
        roleArn: ${aws_iam_role.firehose_role.arn}
        indexName: test
        s3Configuration:
          roleArn: ${aws_iam_role.firehose_role.arn}
          bucketArn: ${aws_s3_bucket.bucket.arn}
          bufferingSize: 10
          bufferingInterval: 400
          compressionFormat: GZIP
        processingConfiguration:
          enabled: 'true'
          processors:
            - type: Lambda
              parameters:
                - parameterName: LambdaArn
                  parameterValue: ${aws_lambda_function.lambda_processor.arn}:$LATEST
```
{{% /example %}}
{{% example %}}
### Splunk Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testStream = new aws.kinesis.FirehoseDeliveryStream("testStream", {
    destination: "splunk",
    splunkConfiguration: {
        hecEndpoint: "https://http-inputs-mydomain.splunkcloud.com:443",
        hecToken: "51D4DA16-C61B-4F5F-8EC7-ED4301342A4A",
        hecAcknowledgmentTimeout: 600,
        hecEndpointType: "Event",
        s3BackupMode: "FailedEventsOnly",
        s3Configuration: {
            roleArn: aws_iam_role.firehose.arn,
            bucketArn: aws_s3_bucket.bucket.arn,
            bufferingSize: 10,
            bufferingInterval: 400,
            compressionFormat: "GZIP",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="splunk",
    splunk_configuration=aws.kinesis.FirehoseDeliveryStreamSplunkConfigurationArgs(
        hec_endpoint="https://http-inputs-mydomain.splunkcloud.com:443",
        hec_token="51D4DA16-C61B-4F5F-8EC7-ED4301342A4A",
        hec_acknowledgment_timeout=600,
        hec_endpoint_type="Event",
        s3_backup_mode="FailedEventsOnly",
        s3_configuration=aws.kinesis.FirehoseDeliveryStreamSplunkConfigurationS3ConfigurationArgs(
            role_arn=aws_iam_role["firehose"]["arn"],
            bucket_arn=aws_s3_bucket["bucket"]["arn"],
            buffering_size=10,
            buffering_interval=400,
            compression_format="GZIP",
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new()
    {
        Destination = "splunk",
        SplunkConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamSplunkConfigurationArgs
        {
            HecEndpoint = "https://http-inputs-mydomain.splunkcloud.com:443",
            HecToken = "51D4DA16-C61B-4F5F-8EC7-ED4301342A4A",
            HecAcknowledgmentTimeout = 600,
            HecEndpointType = "Event",
            S3BackupMode = "FailedEventsOnly",
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamSplunkConfigurationS3ConfigurationArgs
            {
                RoleArn = aws_iam_role.Firehose.Arn,
                BucketArn = aws_s3_bucket.Bucket.Arn,
                BufferingSize = 10,
                BufferingInterval = 400,
                CompressionFormat = "GZIP",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("splunk"),
			SplunkConfiguration: &kinesis.FirehoseDeliveryStreamSplunkConfigurationArgs{
				HecEndpoint:              pulumi.String("https://http-inputs-mydomain.splunkcloud.com:443"),
				HecToken:                 pulumi.String("51D4DA16-C61B-4F5F-8EC7-ED4301342A4A"),
				HecAcknowledgmentTimeout: pulumi.Int(600),
				HecEndpointType:          pulumi.String("Event"),
				S3BackupMode:             pulumi.String("FailedEventsOnly"),
				S3Configuration: &kinesis.FirehoseDeliveryStreamSplunkConfigurationS3ConfigurationArgs{
					RoleArn:           pulumi.Any(aws_iam_role.Firehose.Arn),
					BucketArn:         pulumi.Any(aws_s3_bucket.Bucket.Arn),
					BufferingSize:     pulumi.Int(10),
					BufferingInterval: pulumi.Int(400),
					CompressionFormat: pulumi.String("GZIP"),
				},
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
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamSplunkConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamSplunkConfigurationS3ConfigurationArgs;
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
        var testStream = new FirehoseDeliveryStream("testStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("splunk")
            .splunkConfiguration(FirehoseDeliveryStreamSplunkConfigurationArgs.builder()
                .hecEndpoint("https://http-inputs-mydomain.splunkcloud.com:443")
                .hecToken("51D4DA16-C61B-4F5F-8EC7-ED4301342A4A")
                .hecAcknowledgmentTimeout(600)
                .hecEndpointType("Event")
                .s3BackupMode("FailedEventsOnly")
                .s3Configuration(FirehoseDeliveryStreamSplunkConfigurationS3ConfigurationArgs.builder()
                    .roleArn(aws_iam_role.firehose().arn())
                    .bucketArn(aws_s3_bucket.bucket().arn())
                    .bufferingSize(10)
                    .bufferingInterval(400)
                    .compressionFormat("GZIP")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: splunk
      splunkConfiguration:
        hecEndpoint: https://http-inputs-mydomain.splunkcloud.com:443
        hecToken: 51D4DA16-C61B-4F5F-8EC7-ED4301342A4A
        hecAcknowledgmentTimeout: 600
        hecEndpointType: Event
        s3BackupMode: FailedEventsOnly
        s3Configuration:
          roleArn: ${aws_iam_role.firehose.arn}
          bucketArn: ${aws_s3_bucket.bucket.arn}
          bufferingSize: 10
          bufferingInterval: 400
          compressionFormat: GZIP
```
{{% /example %}}
{{% example %}}
### HTTP Endpoint (e.g., New Relic) Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testStream = new aws.kinesis.FirehoseDeliveryStream("testStream", {
    destination: "http_endpoint",
    httpEndpointConfiguration: {
        url: "https://aws-api.newrelic.com/firehose/v1",
        name: "New Relic",
        accessKey: "my-key",
        bufferingSize: 15,
        bufferingInterval: 600,
        roleArn: aws_iam_role.firehose.arn,
        s3BackupMode: "FailedDataOnly",
        s3Configuration: {
            roleArn: aws_iam_role.firehose.arn,
            bucketArn: aws_s3_bucket.bucket.arn,
            bufferingSize: 10,
            bufferingInterval: 400,
            compressionFormat: "GZIP",
        },
        requestConfiguration: {
            contentEncoding: "GZIP",
            commonAttributes: [
                {
                    name: "testname",
                    value: "testvalue",
                },
                {
                    name: "testname2",
                    value: "testvalue2",
                },
            ],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
    destination="http_endpoint",
    http_endpoint_configuration=aws.kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationArgs(
        url="https://aws-api.newrelic.com/firehose/v1",
        name="New Relic",
        access_key="my-key",
        buffering_size=15,
        buffering_interval=600,
        role_arn=aws_iam_role["firehose"]["arn"],
        s3_backup_mode="FailedDataOnly",
        s3_configuration=aws.kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationS3ConfigurationArgs(
            role_arn=aws_iam_role["firehose"]["arn"],
            bucket_arn=aws_s3_bucket["bucket"]["arn"],
            buffering_size=10,
            buffering_interval=400,
            compression_format="GZIP",
        ),
        request_configuration=aws.kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationArgs(
            content_encoding="GZIP",
            common_attributes=[
                aws.kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttributeArgs(
                    name="testname",
                    value="testvalue",
                ),
                aws.kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttributeArgs(
                    name="testname2",
                    value="testvalue2",
                ),
            ],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new()
    {
        Destination = "http_endpoint",
        HttpEndpointConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamHttpEndpointConfigurationArgs
        {
            Url = "https://aws-api.newrelic.com/firehose/v1",
            Name = "New Relic",
            AccessKey = "my-key",
            BufferingSize = 15,
            BufferingInterval = 600,
            RoleArn = aws_iam_role.Firehose.Arn,
            S3BackupMode = "FailedDataOnly",
            S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamHttpEndpointConfigurationS3ConfigurationArgs
            {
                RoleArn = aws_iam_role.Firehose.Arn,
                BucketArn = aws_s3_bucket.Bucket.Arn,
                BufferingSize = 10,
                BufferingInterval = 400,
                CompressionFormat = "GZIP",
            },
            RequestConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationArgs
            {
                ContentEncoding = "GZIP",
                CommonAttributes = new[]
                {
                    new Aws.Kinesis.Inputs.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttributeArgs
                    {
                        Name = "testname",
                        Value = "testvalue",
                    },
                    new Aws.Kinesis.Inputs.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttributeArgs
                    {
                        Name = "testname2",
                        Value = "testvalue2",
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("http_endpoint"),
			HttpEndpointConfiguration: &kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationArgs{
				Url:               pulumi.String("https://aws-api.newrelic.com/firehose/v1"),
				Name:              pulumi.String("New Relic"),
				AccessKey:         pulumi.String("my-key"),
				BufferingSize:     pulumi.Int(15),
				BufferingInterval: pulumi.Int(600),
				RoleArn:           pulumi.Any(aws_iam_role.Firehose.Arn),
				S3BackupMode:      pulumi.String("FailedDataOnly"),
				S3Configuration: &kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationS3ConfigurationArgs{
					RoleArn:           pulumi.Any(aws_iam_role.Firehose.Arn),
					BucketArn:         pulumi.Any(aws_s3_bucket.Bucket.Arn),
					BufferingSize:     pulumi.Int(10),
					BufferingInterval: pulumi.Int(400),
					CompressionFormat: pulumi.String("GZIP"),
				},
				RequestConfiguration: &kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationArgs{
					ContentEncoding: pulumi.String("GZIP"),
					CommonAttributes: kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttributeArray{
						&kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttributeArgs{
							Name:  pulumi.String("testname"),
							Value: pulumi.String("testvalue"),
						},
						&kinesis.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttributeArgs{
							Name:  pulumi.String("testname2"),
							Value: pulumi.String("testvalue2"),
						},
					},
				},
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
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamHttpEndpointConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamHttpEndpointConfigurationS3ConfigurationArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationArgs;
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
        var testStream = new FirehoseDeliveryStream("testStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("http_endpoint")
            .httpEndpointConfiguration(FirehoseDeliveryStreamHttpEndpointConfigurationArgs.builder()
                .url("https://aws-api.newrelic.com/firehose/v1")
                .name("New Relic")
                .accessKey("my-key")
                .bufferingSize(15)
                .bufferingInterval(600)
                .roleArn(aws_iam_role.firehose().arn())
                .s3BackupMode("FailedDataOnly")
                .s3Configuration(FirehoseDeliveryStreamHttpEndpointConfigurationS3ConfigurationArgs.builder()
                    .roleArn(aws_iam_role.firehose().arn())
                    .bucketArn(aws_s3_bucket.bucket().arn())
                    .bufferingSize(10)
                    .bufferingInterval(400)
                    .compressionFormat("GZIP")
                    .build())
                .requestConfiguration(FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationArgs.builder()
                    .contentEncoding("GZIP")
                    .commonAttributes(                    
                        FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttributeArgs.builder()
                            .name("testname")
                            .value("testvalue")
                            .build(),
                        FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfigurationCommonAttributeArgs.builder()
                            .name("testname2")
                            .value("testvalue2")
                            .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: http_endpoint
      httpEndpointConfiguration:
        url: https://aws-api.newrelic.com/firehose/v1
        name: New Relic
        accessKey: my-key
        bufferingSize: 15
        bufferingInterval: 600
        roleArn: ${aws_iam_role.firehose.arn}
        s3BackupMode: FailedDataOnly
        s3Configuration:
          roleArn: ${aws_iam_role.firehose.arn}
          bucketArn: ${aws_s3_bucket.bucket.arn}
          bufferingSize: 10
          bufferingInterval: 400
          compressionFormat: GZIP
        requestConfiguration:
          contentEncoding: GZIP
          commonAttributes:
            - name: testname
              value: testvalue
            - name: testname2
              value: testvalue2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Kinesis Firehose Delivery streams using the stream ARN. For example:

```sh
 $ pulumi import aws:kinesis/firehoseDeliveryStream:FirehoseDeliveryStream foo arn:aws:firehose:us-east-1:XXX:deliverystream/example
```
 NoteImport does not work for stream destination `s3`. Consider using `extended_s3` since `s3` destination is deprecated.

