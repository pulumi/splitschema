Provides a CloudWatch Metric Stream resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Filters

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const streamsAssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["streams.metrics.cloudwatch.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const metricStreamToFirehoseRole = new aws.iam.Role("metricStreamToFirehoseRole", {assumeRolePolicy: streamsAssumeRole.then(streamsAssumeRole => streamsAssumeRole.json)});
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
const firehoseToS3Role = new aws.iam.Role("firehoseToS3Role", {assumeRolePolicy: firehoseAssumeRole.then(firehoseAssumeRole => firehoseAssumeRole.json)});
const s3Stream = new aws.kinesis.FirehoseDeliveryStream("s3Stream", {
    destination: "extended_s3",
    extendedS3Configuration: {
        roleArn: firehoseToS3Role.arn,
        bucketArn: bucket.arn,
    },
});
const main = new aws.cloudwatch.MetricStream("main", {
    roleArn: metricStreamToFirehoseRole.arn,
    firehoseArn: s3Stream.arn,
    outputFormat: "json",
    includeFilters: [
        {
            namespace: "AWS/EC2",
            metricNames: [
                "CPUUtilization",
                "NetworkOut",
            ],
        },
        {
            namespace: "AWS/EBS",
            metricNames: [],
        },
    ],
});
const metricStreamToFirehosePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        actions: [
            "firehose:PutRecord",
            "firehose:PutRecordBatch",
        ],
        resources: [s3Stream.arn],
    }],
});
const metricStreamToFirehoseRolePolicy = new aws.iam.RolePolicy("metricStreamToFirehoseRolePolicy", {
    role: metricStreamToFirehoseRole.id,
    policy: metricStreamToFirehosePolicyDocument.apply(metricStreamToFirehosePolicyDocument => metricStreamToFirehosePolicyDocument.json),
});
const bucketAcl = new aws.s3.BucketAclV2("bucketAcl", {
    bucket: bucket.id,
    acl: "private",
});
const firehoseToS3PolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        actions: [
            "s3:AbortMultipartUpload",
            "s3:GetBucketLocation",
            "s3:GetObject",
            "s3:ListBucket",
            "s3:ListBucketMultipartUploads",
            "s3:PutObject",
        ],
        resources: [
            bucket.arn,
            pulumi.interpolate`${bucket.arn}/*`,
        ],
    }],
});
const firehoseToS3RolePolicy = new aws.iam.RolePolicy("firehoseToS3RolePolicy", {
    role: firehoseToS3Role.id,
    policy: firehoseToS3PolicyDocument.apply(firehoseToS3PolicyDocument => firehoseToS3PolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

streams_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["streams.metrics.cloudwatch.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
metric_stream_to_firehose_role = aws.iam.Role("metricStreamToFirehoseRole", assume_role_policy=streams_assume_role.json)
bucket = aws.s3.BucketV2("bucket")
firehose_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["firehose.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
firehose_to_s3_role = aws.iam.Role("firehoseToS3Role", assume_role_policy=firehose_assume_role.json)
s3_stream = aws.kinesis.FirehoseDeliveryStream("s3Stream",
    destination="extended_s3",
    extended_s3_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs(
        role_arn=firehose_to_s3_role.arn,
        bucket_arn=bucket.arn,
    ))
main = aws.cloudwatch.MetricStream("main",
    role_arn=metric_stream_to_firehose_role.arn,
    firehose_arn=s3_stream.arn,
    output_format="json",
    include_filters=[
        aws.cloudwatch.MetricStreamIncludeFilterArgs(
            namespace="AWS/EC2",
            metric_names=[
                "CPUUtilization",
                "NetworkOut",
            ],
        ),
        aws.cloudwatch.MetricStreamIncludeFilterArgs(
            namespace="AWS/EBS",
            metric_names=[],
        ),
    ])
metric_stream_to_firehose_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=[
        "firehose:PutRecord",
        "firehose:PutRecordBatch",
    ],
    resources=[s3_stream.arn],
)])
metric_stream_to_firehose_role_policy = aws.iam.RolePolicy("metricStreamToFirehoseRolePolicy",
    role=metric_stream_to_firehose_role.id,
    policy=metric_stream_to_firehose_policy_document.json)
bucket_acl = aws.s3.BucketAclV2("bucketAcl",
    bucket=bucket.id,
    acl="private")
firehose_to_s3_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=[
        "s3:AbortMultipartUpload",
        "s3:GetBucketLocation",
        "s3:GetObject",
        "s3:ListBucket",
        "s3:ListBucketMultipartUploads",
        "s3:PutObject",
    ],
    resources=[
        bucket.arn,
        bucket.arn.apply(lambda arn: f"{arn}/*"),
    ],
)])
firehose_to_s3_role_policy = aws.iam.RolePolicy("firehoseToS3RolePolicy",
    role=firehose_to_s3_role.id,
    policy=firehose_to_s3_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var streamsAssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            "streams.metrics.cloudwatch.amazonaws.com",
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

    var metricStreamToFirehoseRole = new Aws.Iam.Role("metricStreamToFirehoseRole", new()
    {
        AssumeRolePolicy = streamsAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

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

    var firehoseToS3Role = new Aws.Iam.Role("firehoseToS3Role", new()
    {
        AssumeRolePolicy = firehoseAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var s3Stream = new Aws.Kinesis.FirehoseDeliveryStream("s3Stream", new()
    {
        Destination = "extended_s3",
        ExtendedS3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs
        {
            RoleArn = firehoseToS3Role.Arn,
            BucketArn = bucket.Arn,
        },
    });

    var main = new Aws.CloudWatch.MetricStream("main", new()
    {
        RoleArn = metricStreamToFirehoseRole.Arn,
        FirehoseArn = s3Stream.Arn,
        OutputFormat = "json",
        IncludeFilters = new[]
        {
            new Aws.CloudWatch.Inputs.MetricStreamIncludeFilterArgs
            {
                Namespace = "AWS/EC2",
                MetricNames = new[]
                {
                    "CPUUtilization",
                    "NetworkOut",
                },
            },
            new Aws.CloudWatch.Inputs.MetricStreamIncludeFilterArgs
            {
                Namespace = "AWS/EBS",
                MetricNames = new() { },
            },
        },
    });

    var metricStreamToFirehosePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "firehose:PutRecord",
                    "firehose:PutRecordBatch",
                },
                Resources = new[]
                {
                    s3Stream.Arn,
                },
            },
        },
    });

    var metricStreamToFirehoseRolePolicy = new Aws.Iam.RolePolicy("metricStreamToFirehoseRolePolicy", new()
    {
        Role = metricStreamToFirehoseRole.Id,
        Policy = metricStreamToFirehosePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var bucketAcl = new Aws.S3.BucketAclV2("bucketAcl", new()
    {
        Bucket = bucket.Id,
        Acl = "private",
    });

    var firehoseToS3PolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "s3:AbortMultipartUpload",
                    "s3:GetBucketLocation",
                    "s3:GetObject",
                    "s3:ListBucket",
                    "s3:ListBucketMultipartUploads",
                    "s3:PutObject",
                },
                Resources = new[]
                {
                    bucket.Arn,
                    $"{bucket.Arn}/*",
                },
            },
        },
    });

    var firehoseToS3RolePolicy = new Aws.Iam.RolePolicy("firehoseToS3RolePolicy", new()
    {
        Role = firehoseToS3Role.Id,
        Policy = firehoseToS3PolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		streamsAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"streams.metrics.cloudwatch.amazonaws.com",
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
		metricStreamToFirehoseRole, err := iam.NewRole(ctx, "metricStreamToFirehoseRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(streamsAssumeRole.Json),
		})
		if err != nil {
			return err
		}
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
		firehoseToS3Role, err := iam.NewRole(ctx, "firehoseToS3Role", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(firehoseAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		s3Stream, err := kinesis.NewFirehoseDeliveryStream(ctx, "s3Stream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("extended_s3"),
			ExtendedS3Configuration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs{
				RoleArn:   firehoseToS3Role.Arn,
				BucketArn: bucket.Arn,
			},
		})
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewMetricStream(ctx, "main", &cloudwatch.MetricStreamArgs{
			RoleArn:      metricStreamToFirehoseRole.Arn,
			FirehoseArn:  s3Stream.Arn,
			OutputFormat: pulumi.String("json"),
			IncludeFilters: cloudwatch.MetricStreamIncludeFilterArray{
				&cloudwatch.MetricStreamIncludeFilterArgs{
					Namespace: pulumi.String("AWS/EC2"),
					MetricNames: pulumi.StringArray{
						pulumi.String("CPUUtilization"),
						pulumi.String("NetworkOut"),
					},
				},
				&cloudwatch.MetricStreamIncludeFilterArgs{
					Namespace:   pulumi.String("AWS/EBS"),
					MetricNames: pulumi.StringArray{},
				},
			},
		})
		if err != nil {
			return err
		}
		metricStreamToFirehosePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("firehose:PutRecord"),
						pulumi.String("firehose:PutRecordBatch"),
					},
					Resources: pulumi.StringArray{
						s3Stream.Arn,
					},
				},
			},
		}, nil)
		_, err = iam.NewRolePolicy(ctx, "metricStreamToFirehoseRolePolicy", &iam.RolePolicyArgs{
			Role: metricStreamToFirehoseRole.ID(),
			Policy: metricStreamToFirehosePolicyDocument.ApplyT(func(metricStreamToFirehosePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &metricStreamToFirehosePolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
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
		firehoseToS3PolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:AbortMultipartUpload"),
						pulumi.String("s3:GetBucketLocation"),
						pulumi.String("s3:GetObject"),
						pulumi.String("s3:ListBucket"),
						pulumi.String("s3:ListBucketMultipartUploads"),
						pulumi.String("s3:PutObject"),
					},
					Resources: pulumi.StringArray{
						bucket.Arn,
						bucket.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
				},
			},
		}, nil)
		_, err = iam.NewRolePolicy(ctx, "firehoseToS3RolePolicy", &iam.RolePolicyArgs{
			Role: firehoseToS3Role.ID(),
			Policy: firehoseToS3PolicyDocument.ApplyT(func(firehoseToS3PolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &firehoseToS3PolicyDocument.Json, nil
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs;
import com.pulumi.aws.cloudwatch.MetricStream;
import com.pulumi.aws.cloudwatch.MetricStreamArgs;
import com.pulumi.aws.cloudwatch.inputs.MetricStreamIncludeFilterArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
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
        final var streamsAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("streams.metrics.cloudwatch.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var metricStreamToFirehoseRole = new Role("metricStreamToFirehoseRole", RoleArgs.builder()        
            .assumeRolePolicy(streamsAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

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

        var firehoseToS3Role = new Role("firehoseToS3Role", RoleArgs.builder()        
            .assumeRolePolicy(firehoseAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var s3Stream = new FirehoseDeliveryStream("s3Stream", FirehoseDeliveryStreamArgs.builder()        
            .destination("extended_s3")
            .extendedS3Configuration(FirehoseDeliveryStreamExtendedS3ConfigurationArgs.builder()
                .roleArn(firehoseToS3Role.arn())
                .bucketArn(bucket.arn())
                .build())
            .build());

        var main = new MetricStream("main", MetricStreamArgs.builder()        
            .roleArn(metricStreamToFirehoseRole.arn())
            .firehoseArn(s3Stream.arn())
            .outputFormat("json")
            .includeFilters(            
                MetricStreamIncludeFilterArgs.builder()
                    .namespace("AWS/EC2")
                    .metricNames(                    
                        "CPUUtilization",
                        "NetworkOut")
                    .build(),
                MetricStreamIncludeFilterArgs.builder()
                    .namespace("AWS/EBS")
                    .metricNames()
                    .build())
            .build());

        final var metricStreamToFirehosePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions(                
                    "firehose:PutRecord",
                    "firehose:PutRecordBatch")
                .resources(s3Stream.arn())
                .build())
            .build());

        var metricStreamToFirehoseRolePolicy = new RolePolicy("metricStreamToFirehoseRolePolicy", RolePolicyArgs.builder()        
            .role(metricStreamToFirehoseRole.id())
            .policy(metricStreamToFirehosePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(metricStreamToFirehosePolicyDocument -> metricStreamToFirehosePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var bucketAcl = new BucketAclV2("bucketAcl", BucketAclV2Args.builder()        
            .bucket(bucket.id())
            .acl("private")
            .build());

        final var firehoseToS3PolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions(                
                    "s3:AbortMultipartUpload",
                    "s3:GetBucketLocation",
                    "s3:GetObject",
                    "s3:ListBucket",
                    "s3:ListBucketMultipartUploads",
                    "s3:PutObject")
                .resources(                
                    bucket.arn(),
                    bucket.arn().applyValue(arn -> String.format("%s/*", arn)))
                .build())
            .build());

        var firehoseToS3RolePolicy = new RolePolicy("firehoseToS3RolePolicy", RolePolicyArgs.builder()        
            .role(firehoseToS3Role.id())
            .policy(firehoseToS3PolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(firehoseToS3PolicyDocument -> firehoseToS3PolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:cloudwatch:MetricStream
    properties:
      roleArn: ${metricStreamToFirehoseRole.arn}
      firehoseArn: ${s3Stream.arn}
      outputFormat: json
      includeFilters:
        - namespace: AWS/EC2
          metricNames:
            - CPUUtilization
            - NetworkOut
        - namespace: AWS/EBS
          metricNames: []
  metricStreamToFirehoseRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${streamsAssumeRole.json}
  metricStreamToFirehoseRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${metricStreamToFirehoseRole.id}
      policy: ${metricStreamToFirehosePolicyDocument.json}
  bucket:
    type: aws:s3:BucketV2
  bucketAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${bucket.id}
      acl: private
  firehoseToS3Role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${firehoseAssumeRole.json}
  firehoseToS3RolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${firehoseToS3Role.id}
      policy: ${firehoseToS3PolicyDocument.json}
  s3Stream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: extended_s3
      extendedS3Configuration:
        roleArn: ${firehoseToS3Role.arn}
        bucketArn: ${bucket.arn}
variables:
  streamsAssumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - streams.metrics.cloudwatch.amazonaws.com
            actions:
              - sts:AssumeRole
  metricStreamToFirehosePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - firehose:PutRecord
              - firehose:PutRecordBatch
            resources:
              - ${s3Stream.arn}
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
  firehoseToS3PolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - s3:AbortMultipartUpload
              - s3:GetBucketLocation
              - s3:GetObject
              - s3:ListBucket
              - s3:ListBucketMultipartUploads
              - s3:PutObject
            resources:
              - ${bucket.arn}
              - ${bucket.arn}/*
```
{{% /example %}}
{{% example %}}
### Additional Statistics

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.cloudwatch.MetricStream("main", {
    roleArn: aws_iam_role.metric_stream_to_firehose.arn,
    firehoseArn: aws_kinesis_firehose_delivery_stream.s3_stream.arn,
    outputFormat: "json",
    statisticsConfigurations: [
        {
            additionalStatistics: [
                "p1",
                "tm99",
            ],
            includeMetrics: [{
                metricName: "CPUUtilization",
                namespace: "AWS/EC2",
            }],
        },
        {
            additionalStatistics: ["TS(50.5:)"],
            includeMetrics: [{
                metricName: "CPUUtilization",
                namespace: "AWS/EC2",
            }],
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.cloudwatch.MetricStream("main",
    role_arn=aws_iam_role["metric_stream_to_firehose"]["arn"],
    firehose_arn=aws_kinesis_firehose_delivery_stream["s3_stream"]["arn"],
    output_format="json",
    statistics_configurations=[
        aws.cloudwatch.MetricStreamStatisticsConfigurationArgs(
            additional_statistics=[
                "p1",
                "tm99",
            ],
            include_metrics=[aws.cloudwatch.MetricStreamStatisticsConfigurationIncludeMetricArgs(
                metric_name="CPUUtilization",
                namespace="AWS/EC2",
            )],
        ),
        aws.cloudwatch.MetricStreamStatisticsConfigurationArgs(
            additional_statistics=["TS(50.5:)"],
            include_metrics=[aws.cloudwatch.MetricStreamStatisticsConfigurationIncludeMetricArgs(
                metric_name="CPUUtilization",
                namespace="AWS/EC2",
            )],
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.CloudWatch.MetricStream("main", new()
    {
        RoleArn = aws_iam_role.Metric_stream_to_firehose.Arn,
        FirehoseArn = aws_kinesis_firehose_delivery_stream.S3_stream.Arn,
        OutputFormat = "json",
        StatisticsConfigurations = new[]
        {
            new Aws.CloudWatch.Inputs.MetricStreamStatisticsConfigurationArgs
            {
                AdditionalStatistics = new[]
                {
                    "p1",
                    "tm99",
                },
                IncludeMetrics = new[]
                {
                    new Aws.CloudWatch.Inputs.MetricStreamStatisticsConfigurationIncludeMetricArgs
                    {
                        MetricName = "CPUUtilization",
                        Namespace = "AWS/EC2",
                    },
                },
            },
            new Aws.CloudWatch.Inputs.MetricStreamStatisticsConfigurationArgs
            {
                AdditionalStatistics = new[]
                {
                    "TS(50.5:)",
                },
                IncludeMetrics = new[]
                {
                    new Aws.CloudWatch.Inputs.MetricStreamStatisticsConfigurationIncludeMetricArgs
                    {
                        MetricName = "CPUUtilization",
                        Namespace = "AWS/EC2",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewMetricStream(ctx, "main", &cloudwatch.MetricStreamArgs{
			RoleArn:      pulumi.Any(aws_iam_role.Metric_stream_to_firehose.Arn),
			FirehoseArn:  pulumi.Any(aws_kinesis_firehose_delivery_stream.S3_stream.Arn),
			OutputFormat: pulumi.String("json"),
			StatisticsConfigurations: cloudwatch.MetricStreamStatisticsConfigurationArray{
				&cloudwatch.MetricStreamStatisticsConfigurationArgs{
					AdditionalStatistics: pulumi.StringArray{
						pulumi.String("p1"),
						pulumi.String("tm99"),
					},
					IncludeMetrics: cloudwatch.MetricStreamStatisticsConfigurationIncludeMetricArray{
						&cloudwatch.MetricStreamStatisticsConfigurationIncludeMetricArgs{
							MetricName: pulumi.String("CPUUtilization"),
							Namespace:  pulumi.String("AWS/EC2"),
						},
					},
				},
				&cloudwatch.MetricStreamStatisticsConfigurationArgs{
					AdditionalStatistics: pulumi.StringArray{
						pulumi.String("TS(50.5:)"),
					},
					IncludeMetrics: cloudwatch.MetricStreamStatisticsConfigurationIncludeMetricArray{
						&cloudwatch.MetricStreamStatisticsConfigurationIncludeMetricArgs{
							MetricName: pulumi.String("CPUUtilization"),
							Namespace:  pulumi.String("AWS/EC2"),
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
import com.pulumi.aws.cloudwatch.MetricStream;
import com.pulumi.aws.cloudwatch.MetricStreamArgs;
import com.pulumi.aws.cloudwatch.inputs.MetricStreamStatisticsConfigurationArgs;
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
        var main = new MetricStream("main", MetricStreamArgs.builder()        
            .roleArn(aws_iam_role.metric_stream_to_firehose().arn())
            .firehoseArn(aws_kinesis_firehose_delivery_stream.s3_stream().arn())
            .outputFormat("json")
            .statisticsConfigurations(            
                MetricStreamStatisticsConfigurationArgs.builder()
                    .additionalStatistics(                    
                        "p1",
                        "tm99")
                    .includeMetrics(MetricStreamStatisticsConfigurationIncludeMetricArgs.builder()
                        .metricName("CPUUtilization")
                        .namespace("AWS/EC2")
                        .build())
                    .build(),
                MetricStreamStatisticsConfigurationArgs.builder()
                    .additionalStatistics("TS(50.5:)")
                    .includeMetrics(MetricStreamStatisticsConfigurationIncludeMetricArgs.builder()
                        .metricName("CPUUtilization")
                        .namespace("AWS/EC2")
                        .build())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:cloudwatch:MetricStream
    properties:
      roleArn: ${aws_iam_role.metric_stream_to_firehose.arn}
      firehoseArn: ${aws_kinesis_firehose_delivery_stream.s3_stream.arn}
      outputFormat: json
      statisticsConfigurations:
        - additionalStatistics:
            - p1
            - tm99
          includeMetrics:
            - metricName: CPUUtilization
              namespace: AWS/EC2
        - additionalStatistics:
            - TS(50.5:)
          includeMetrics:
            - metricName: CPUUtilization
              namespace: AWS/EC2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch metric streams using the `name`. For example:

```sh
 $ pulumi import aws:cloudwatch/metricStream:MetricStream sample sample-stream-name
```
 