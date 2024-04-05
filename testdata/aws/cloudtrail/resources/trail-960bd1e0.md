Provides a CloudTrail resource.

> **Tip:** For a multi-region trail, this resource must be in the home region of the trail.

> **Tip:** For an organization trail, this resource must be in the master account of the organization.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

Enable CloudTrail to capture all compatible management events in region.
For capturing events from services like IAM, `include_global_service_events` must be enabled.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {forceDestroy: true});
const currentCallerIdentity = aws.getCallerIdentity({});
const currentPartition = aws.getPartition({});
const currentRegion = aws.getRegion({});
const examplePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            sid: "AWSCloudTrailAclCheck",
            effect: "Allow",
            principals: [{
                type: "Service",
                identifiers: ["cloudtrail.amazonaws.com"],
            }],
            actions: ["s3:GetBucketAcl"],
            resources: [exampleBucketV2.arn],
            conditions: [{
                test: "StringEquals",
                variable: "aws:SourceArn",
                values: [Promise.all([currentPartition, currentRegion, currentCallerIdentity]).then(([currentPartition, currentRegion, currentCallerIdentity]) => `arn:${currentPartition.partition}:cloudtrail:${currentRegion.name}:${currentCallerIdentity.accountId}:trail/example`)],
            }],
        },
        {
            sid: "AWSCloudTrailWrite",
            effect: "Allow",
            principals: [{
                type: "Service",
                identifiers: ["cloudtrail.amazonaws.com"],
            }],
            actions: ["s3:PutObject"],
            resources: [pulumi.all([exampleBucketV2.arn, currentCallerIdentity]).apply(([arn, currentCallerIdentity]) => `${arn}/prefix/AWSLogs/${currentCallerIdentity.accountId}/*`)],
            conditions: [
                {
                    test: "StringEquals",
                    variable: "s3:x-amz-acl",
                    values: ["bucket-owner-full-control"],
                },
                {
                    test: "StringEquals",
                    variable: "aws:SourceArn",
                    values: [Promise.all([currentPartition, currentRegion, currentCallerIdentity]).then(([currentPartition, currentRegion, currentCallerIdentity]) => `arn:${currentPartition.partition}:cloudtrail:${currentRegion.name}:${currentCallerIdentity.accountId}:trail/example`)],
                },
            ],
        },
    ],
});
const exampleBucketPolicy = new aws.s3.BucketPolicy("exampleBucketPolicy", {
    bucket: exampleBucketV2.id,
    policy: examplePolicyDocument.apply(examplePolicyDocument => examplePolicyDocument.json),
});
const exampleTrail = new aws.cloudtrail.Trail("exampleTrail", {
    s3BucketName: exampleBucketV2.id,
    s3KeyPrefix: "prefix",
    includeGlobalServiceEvents: false,
}, {
    dependsOn: [exampleBucketPolicy],
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2", force_destroy=True)
current_caller_identity = aws.get_caller_identity()
current_partition = aws.get_partition()
current_region = aws.get_region()
example_policy_document = aws.iam.get_policy_document_output(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="AWSCloudTrailAclCheck",
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=["cloudtrail.amazonaws.com"],
        )],
        actions=["s3:GetBucketAcl"],
        resources=[example_bucket_v2.arn],
        conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
            test="StringEquals",
            variable="aws:SourceArn",
            values=[f"arn:{current_partition.partition}:cloudtrail:{current_region.name}:{current_caller_identity.account_id}:trail/example"],
        )],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="AWSCloudTrailWrite",
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=["cloudtrail.amazonaws.com"],
        )],
        actions=["s3:PutObject"],
        resources=[example_bucket_v2.arn.apply(lambda arn: f"{arn}/prefix/AWSLogs/{current_caller_identity.account_id}/*")],
        conditions=[
            aws.iam.GetPolicyDocumentStatementConditionArgs(
                test="StringEquals",
                variable="s3:x-amz-acl",
                values=["bucket-owner-full-control"],
            ),
            aws.iam.GetPolicyDocumentStatementConditionArgs(
                test="StringEquals",
                variable="aws:SourceArn",
                values=[f"arn:{current_partition.partition}:cloudtrail:{current_region.name}:{current_caller_identity.account_id}:trail/example"],
            ),
        ],
    ),
])
example_bucket_policy = aws.s3.BucketPolicy("exampleBucketPolicy",
    bucket=example_bucket_v2.id,
    policy=example_policy_document.json)
example_trail = aws.cloudtrail.Trail("exampleTrail",
    s3_bucket_name=example_bucket_v2.id,
    s3_key_prefix="prefix",
    include_global_service_events=False,
    opts=pulumi.ResourceOptions(depends_on=[example_bucket_policy]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2", new()
    {
        ForceDestroy = true,
    });

    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var currentPartition = Aws.GetPartition.Invoke();

    var currentRegion = Aws.GetRegion.Invoke();

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "AWSCloudTrailAclCheck",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "cloudtrail.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:GetBucketAcl",
                },
                Resources = new[]
                {
                    exampleBucketV2.Arn,
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "StringEquals",
                        Variable = "aws:SourceArn",
                        Values = new[]
                        {
                            $"arn:{currentPartition.Apply(getPartitionResult => getPartitionResult.Partition)}:cloudtrail:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:trail/example",
                        },
                    },
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "AWSCloudTrailWrite",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "cloudtrail.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:PutObject",
                },
                Resources = new[]
                {
                    $"{exampleBucketV2.Arn}/prefix/AWSLogs/{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}/*",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "StringEquals",
                        Variable = "s3:x-amz-acl",
                        Values = new[]
                        {
                            "bucket-owner-full-control",
                        },
                    },
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "StringEquals",
                        Variable = "aws:SourceArn",
                        Values = new[]
                        {
                            $"arn:{currentPartition.Apply(getPartitionResult => getPartitionResult.Partition)}:cloudtrail:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:trail/example",
                        },
                    },
                },
            },
        },
    });

    var exampleBucketPolicy = new Aws.S3.BucketPolicy("exampleBucketPolicy", new()
    {
        Bucket = exampleBucketV2.Id,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleTrail = new Aws.CloudTrail.Trail("exampleTrail", new()
    {
        S3BucketName = exampleBucketV2.Id,
        S3KeyPrefix = "prefix",
        IncludeGlobalServiceEvents = false,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleBucketPolicy,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudtrail"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", &s3.BucketV2Args{
			ForceDestroy: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
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
		examplePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Sid:    pulumi.String("AWSCloudTrailAclCheck"),
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("Service"),
							Identifiers: pulumi.StringArray{
								pulumi.String("cloudtrail.amazonaws.com"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("s3:GetBucketAcl"),
					},
					Resources: pulumi.StringArray{
						exampleBucketV2.Arn,
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("StringEquals"),
							Variable: pulumi.String("aws:SourceArn"),
							Values: pulumi.StringArray{
								pulumi.String(fmt.Sprintf("arn:%v:cloudtrail:%v:%v:trail/example", currentPartition.Partition, currentRegion.Name, currentCallerIdentity.AccountId)),
							},
						},
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Sid:    pulumi.String("AWSCloudTrailWrite"),
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("Service"),
							Identifiers: pulumi.StringArray{
								pulumi.String("cloudtrail.amazonaws.com"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("s3:PutObject"),
					},
					Resources: pulumi.StringArray{
						exampleBucketV2.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/prefix/AWSLogs/%v/*", arn, currentCallerIdentity.AccountId), nil
						}).(pulumi.StringOutput),
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("StringEquals"),
							Variable: pulumi.String("s3:x-amz-acl"),
							Values: pulumi.StringArray{
								pulumi.String("bucket-owner-full-control"),
							},
						},
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("StringEquals"),
							Variable: pulumi.String("aws:SourceArn"),
							Values: pulumi.StringArray{
								pulumi.String(fmt.Sprintf("arn:%v:cloudtrail:%v:%v:trail/example", currentPartition.Partition, currentRegion.Name, currentCallerIdentity.AccountId)),
							},
						},
					},
				},
			},
		}, nil)
		exampleBucketPolicy, err := s3.NewBucketPolicy(ctx, "exampleBucketPolicy", &s3.BucketPolicyArgs{
			Bucket: exampleBucketV2.ID(),
			Policy: examplePolicyDocument.ApplyT(func(examplePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &examplePolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
		})
		if err != nil {
			return err
		}
		_, err = cloudtrail.NewTrail(ctx, "exampleTrail", &cloudtrail.TrailArgs{
			S3BucketName:               exampleBucketV2.ID(),
			S3KeyPrefix:                pulumi.String("prefix"),
			IncludeGlobalServiceEvents: pulumi.Bool(false),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleBucketPolicy,
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
import com.pulumi.aws.cloudtrail.Trail;
import com.pulumi.aws.cloudtrail.TrailArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2", BucketV2Args.builder()        
            .forceDestroy(true)
            .build());

        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        final var currentPartition = AwsFunctions.getPartition();

        final var currentRegion = AwsFunctions.getRegion();

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .sid("AWSCloudTrailAclCheck")
                    .effect("Allow")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("Service")
                        .identifiers("cloudtrail.amazonaws.com")
                        .build())
                    .actions("s3:GetBucketAcl")
                    .resources(exampleBucketV2.arn())
                    .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                        .test("StringEquals")
                        .variable("aws:SourceArn")
                        .values(String.format("arn:%s:cloudtrail:%s:%s:trail/example", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                        .build())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .sid("AWSCloudTrailWrite")
                    .effect("Allow")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("Service")
                        .identifiers("cloudtrail.amazonaws.com")
                        .build())
                    .actions("s3:PutObject")
                    .resources(exampleBucketV2.arn().applyValue(arn -> String.format("%s/prefix/AWSLogs/%s/*", arn,currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))))
                    .conditions(                    
                        GetPolicyDocumentStatementConditionArgs.builder()
                            .test("StringEquals")
                            .variable("s3:x-amz-acl")
                            .values("bucket-owner-full-control")
                            .build(),
                        GetPolicyDocumentStatementConditionArgs.builder()
                            .test("StringEquals")
                            .variable("aws:SourceArn")
                            .values(String.format("arn:%s:cloudtrail:%s:%s:trail/example", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                            .build())
                    .build())
            .build());

        var exampleBucketPolicy = new BucketPolicy("exampleBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(exampleBucketV2.id())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(examplePolicyDocument -> examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var exampleTrail = new Trail("exampleTrail", TrailArgs.builder()        
            .s3BucketName(exampleBucketV2.id())
            .s3KeyPrefix("prefix")
            .includeGlobalServiceEvents(false)
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleBucketPolicy)
                .build());

    }
}
```
```yaml
resources:
  exampleTrail:
    type: aws:cloudtrail:Trail
    properties:
      s3BucketName: ${exampleBucketV2.id}
      s3KeyPrefix: prefix
      includeGlobalServiceEvents: false
    options:
      dependson:
        - ${exampleBucketPolicy}
  exampleBucketV2:
    type: aws:s3:BucketV2
    properties:
      forceDestroy: true
  exampleBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${exampleBucketV2.id}
      policy: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: AWSCloudTrailAclCheck
            effect: Allow
            principals:
              - type: Service
                identifiers:
                  - cloudtrail.amazonaws.com
            actions:
              - s3:GetBucketAcl
            resources:
              - ${exampleBucketV2.arn}
            conditions:
              - test: StringEquals
                variable: aws:SourceArn
                values:
                  - arn:${currentPartition.partition}:cloudtrail:${currentRegion.name}:${currentCallerIdentity.accountId}:trail/example
          - sid: AWSCloudTrailWrite
            effect: Allow
            principals:
              - type: Service
                identifiers:
                  - cloudtrail.amazonaws.com
            actions:
              - s3:PutObject
            resources:
              - ${exampleBucketV2.arn}/prefix/AWSLogs/${currentCallerIdentity.accountId}/*
            conditions:
              - test: StringEquals
                variable: s3:x-amz-acl
                values:
                  - bucket-owner-full-control
              - test: StringEquals
                variable: aws:SourceArn
                values:
                  - arn:${currentPartition.partition}:cloudtrail:${currentRegion.name}:${currentCallerIdentity.accountId}:trail/example
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
```
{{% /example %}}
### Data Event Logging

CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 objects and Lambda function invocations. Additional information about data event configuration can be found in the following links:

* [CloudTrail API DataResource documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_DataResource.html) (for basic event selector).
* [CloudTrail API AdvancedFieldSelector documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedFieldSelector.html) (for advanced event selector).
{{% example %}}
### Logging All Lambda Function Invocations By Using Basic Event Selectors

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudtrail.Trail("example", {eventSelectors: [{
    dataResources: [{
        type: "AWS::Lambda::Function",
        values: ["arn:aws:lambda"],
    }],
    includeManagementEvents: true,
    readWriteType: "All",
}]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudtrail.Trail("example", event_selectors=[aws.cloudtrail.TrailEventSelectorArgs(
    data_resources=[aws.cloudtrail.TrailEventSelectorDataResourceArgs(
        type="AWS::Lambda::Function",
        values=["arn:aws:lambda"],
    )],
    include_management_events=True,
    read_write_type="All",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudTrail.Trail("example", new()
    {
        EventSelectors = new[]
        {
            new Aws.CloudTrail.Inputs.TrailEventSelectorArgs
            {
                DataResources = new[]
                {
                    new Aws.CloudTrail.Inputs.TrailEventSelectorDataResourceArgs
                    {
                        Type = "AWS::Lambda::Function",
                        Values = new[]
                        {
                            "arn:aws:lambda",
                        },
                    },
                },
                IncludeManagementEvents = true,
                ReadWriteType = "All",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudtrail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
			EventSelectors: cloudtrail.TrailEventSelectorArray{
				&cloudtrail.TrailEventSelectorArgs{
					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
						&cloudtrail.TrailEventSelectorDataResourceArgs{
							Type: pulumi.String("AWS::Lambda::Function"),
							Values: pulumi.StringArray{
								pulumi.String("arn:aws:lambda"),
							},
						},
					},
					IncludeManagementEvents: pulumi.Bool(true),
					ReadWriteType:           pulumi.String("All"),
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
import com.pulumi.aws.cloudtrail.Trail;
import com.pulumi.aws.cloudtrail.TrailArgs;
import com.pulumi.aws.cloudtrail.inputs.TrailEventSelectorArgs;
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
        var example = new Trail("example", TrailArgs.builder()        
            .eventSelectors(TrailEventSelectorArgs.builder()
                .dataResources(TrailEventSelectorDataResourceArgs.builder()
                    .type("AWS::Lambda::Function")
                    .values("arn:aws:lambda")
                    .build())
                .includeManagementEvents(true)
                .readWriteType("All")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudtrail:Trail
    properties:
      eventSelectors:
        - dataResources:
            - type: AWS::Lambda::Function
              values:
                - arn:aws:lambda
          includeManagementEvents: true
          readWriteType: All
```
{{% /example %}}
{{% example %}}
### Logging All S3 Object Events By Using Basic Event Selectors

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudtrail.Trail("example", {eventSelectors: [{
    dataResources: [{
        type: "AWS::S3::Object",
        values: ["arn:aws:s3"],
    }],
    includeManagementEvents: true,
    readWriteType: "All",
}]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudtrail.Trail("example", event_selectors=[aws.cloudtrail.TrailEventSelectorArgs(
    data_resources=[aws.cloudtrail.TrailEventSelectorDataResourceArgs(
        type="AWS::S3::Object",
        values=["arn:aws:s3"],
    )],
    include_management_events=True,
    read_write_type="All",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudTrail.Trail("example", new()
    {
        EventSelectors = new[]
        {
            new Aws.CloudTrail.Inputs.TrailEventSelectorArgs
            {
                DataResources = new[]
                {
                    new Aws.CloudTrail.Inputs.TrailEventSelectorDataResourceArgs
                    {
                        Type = "AWS::S3::Object",
                        Values = new[]
                        {
                            "arn:aws:s3",
                        },
                    },
                },
                IncludeManagementEvents = true,
                ReadWriteType = "All",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudtrail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
			EventSelectors: cloudtrail.TrailEventSelectorArray{
				&cloudtrail.TrailEventSelectorArgs{
					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
						&cloudtrail.TrailEventSelectorDataResourceArgs{
							Type: pulumi.String("AWS::S3::Object"),
							Values: pulumi.StringArray{
								pulumi.String("arn:aws:s3"),
							},
						},
					},
					IncludeManagementEvents: pulumi.Bool(true),
					ReadWriteType:           pulumi.String("All"),
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
import com.pulumi.aws.cloudtrail.Trail;
import com.pulumi.aws.cloudtrail.TrailArgs;
import com.pulumi.aws.cloudtrail.inputs.TrailEventSelectorArgs;
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
        var example = new Trail("example", TrailArgs.builder()        
            .eventSelectors(TrailEventSelectorArgs.builder()
                .dataResources(TrailEventSelectorDataResourceArgs.builder()
                    .type("AWS::S3::Object")
                    .values("arn:aws:s3")
                    .build())
                .includeManagementEvents(true)
                .readWriteType("All")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudtrail:Trail
    properties:
      eventSelectors:
        - dataResources:
            - type: AWS::S3::Object
              values:
                - arn:aws:s3
          includeManagementEvents: true
          readWriteType: All
```
{{% /example %}}
{{% example %}}
### Logging Individual S3 Bucket Events By Using Basic Event Selectors

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const important-bucket = aws.s3.getBucket({
    bucket: "important-bucket",
});
const example = new aws.cloudtrail.Trail("example", {eventSelectors: [{
    dataResources: [{
        type: "AWS::S3::Object",
        values: [important_bucket.then(important_bucket => `${important_bucket.arn}/`)],
    }],
    includeManagementEvents: true,
    readWriteType: "All",
}]});
```
```python
import pulumi
import pulumi_aws as aws

important_bucket = aws.s3.get_bucket(bucket="important-bucket")
example = aws.cloudtrail.Trail("example", event_selectors=[aws.cloudtrail.TrailEventSelectorArgs(
    data_resources=[aws.cloudtrail.TrailEventSelectorDataResourceArgs(
        type="AWS::S3::Object",
        values=[f"{important_bucket.arn}/"],
    )],
    include_management_events=True,
    read_write_type="All",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var important_bucket = Aws.S3.GetBucket.Invoke(new()
    {
        Bucket = "important-bucket",
    });

    var example = new Aws.CloudTrail.Trail("example", new()
    {
        EventSelectors = new[]
        {
            new Aws.CloudTrail.Inputs.TrailEventSelectorArgs
            {
                DataResources = new[]
                {
                    new Aws.CloudTrail.Inputs.TrailEventSelectorDataResourceArgs
                    {
                        Type = "AWS::S3::Object",
                        Values = new[]
                        {
                            important_bucket.Apply(important_bucket => $"{important_bucket.Apply(getBucketResult => getBucketResult.Arn)}/"),
                        },
                    },
                },
                IncludeManagementEvents = true,
                ReadWriteType = "All",
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudtrail"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		important_bucket, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
			Bucket: "important-bucket",
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
			EventSelectors: cloudtrail.TrailEventSelectorArray{
				&cloudtrail.TrailEventSelectorArgs{
					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
						&cloudtrail.TrailEventSelectorDataResourceArgs{
							Type: pulumi.String("AWS::S3::Object"),
							Values: pulumi.StringArray{
								pulumi.String(fmt.Sprintf("%v/", important_bucket.Arn)),
							},
						},
					},
					IncludeManagementEvents: pulumi.Bool(true),
					ReadWriteType:           pulumi.String("All"),
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
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.inputs.GetBucketArgs;
import com.pulumi.aws.cloudtrail.Trail;
import com.pulumi.aws.cloudtrail.TrailArgs;
import com.pulumi.aws.cloudtrail.inputs.TrailEventSelectorArgs;
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
        final var important-bucket = S3Functions.getBucket(GetBucketArgs.builder()
            .bucket("important-bucket")
            .build());

        var example = new Trail("example", TrailArgs.builder()        
            .eventSelectors(TrailEventSelectorArgs.builder()
                .dataResources(TrailEventSelectorDataResourceArgs.builder()
                    .type("AWS::S3::Object")
                    .values(String.format("%s/", important_bucket.arn()))
                    .build())
                .includeManagementEvents(true)
                .readWriteType("All")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudtrail:Trail
    properties:
      eventSelectors:
        - dataResources:
            - type: AWS::S3::Object
              values:
                - ${["important-bucket"].arn}/
          includeManagementEvents: true
          readWriteType: All
variables:
  important-bucket:
    fn::invoke:
      Function: aws:s3:getBucket
      Arguments:
        bucket: important-bucket
```
{{% /example %}}
{{% example %}}
### Logging All S3 Object Events Except For Two S3 Buckets By Using Advanced Event Selectors

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.inputs.GetBucketArgs;
import com.pulumi.aws.cloudtrail.Trail;
import com.pulumi.aws.cloudtrail.TrailArgs;
import com.pulumi.aws.cloudtrail.inputs.TrailAdvancedEventSelectorArgs;
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
        final var not-important-bucket-1 = S3Functions.getBucket(GetBucketArgs.builder()
            .bucket("not-important-bucket-1")
            .build());

        final var not-important-bucket-2 = S3Functions.getBucket(GetBucketArgs.builder()
            .bucket("not-important-bucket-2")
            .build());

        var example = new Trail("example", TrailArgs.builder()        
            .advancedEventSelectors(            
                TrailAdvancedEventSelectorArgs.builder()
                    .fieldSelectors(                    
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .equals("Data")
                            .field("eventCategory")
                            .build(),
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .field("resources.ARN")
                            .notStartsWith(                            
                                String.format("%s/", not_important_bucket_1.arn()),
                                String.format("%s/", not_important_bucket_2.arn()))
                            .build(),
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .equals("AWS::S3::Object")
                            .field("resources.type")
                            .build())
                    .name("Log all S3 objects events except for two S3 buckets")
                    .build(),
                TrailAdvancedEventSelectorArgs.builder()
                    .fieldSelectors(TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                        .equals("Management")
                        .field("eventCategory")
                        .build())
                    .name("Log readOnly and writeOnly management events")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudtrail:Trail
    properties:
      advancedEventSelectors:
        - fieldSelectors:
            - equals:
                - Data
              field: eventCategory
            - field: resources.ARN
              notStartsWith:
                - ${["not-important-bucket-1"].arn}/
                - ${["not-important-bucket-2"].arn}/
            - equals:
                - AWS::S3::Object
              field: resources.type
          name: Log all S3 objects events except for two S3 buckets
        - fieldSelectors:
            - equals:
                - Management
              field: eventCategory
          name: Log readOnly and writeOnly management events
variables:
  not-important-bucket-1:
    fn::invoke:
      Function: aws:s3:getBucket
      Arguments:
        bucket: not-important-bucket-1
  not-important-bucket-2:
    fn::invoke:
      Function: aws:s3:getBucket
      Arguments:
        bucket: not-important-bucket-2
```
{{% /example %}}
{{% example %}}
### Logging Individual S3 Buckets And Specific Event Names By Using Advanced Event Selectors

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.inputs.GetBucketArgs;
import com.pulumi.aws.cloudtrail.Trail;
import com.pulumi.aws.cloudtrail.TrailArgs;
import com.pulumi.aws.cloudtrail.inputs.TrailAdvancedEventSelectorArgs;
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
        final var important-bucket-1 = S3Functions.getBucket(GetBucketArgs.builder()
            .bucket("important-bucket-1")
            .build());

        final var important-bucket-2 = S3Functions.getBucket(GetBucketArgs.builder()
            .bucket("important-bucket-2")
            .build());

        final var important-bucket-3 = S3Functions.getBucket(GetBucketArgs.builder()
            .bucket("important-bucket-3")
            .build());

        var example = new Trail("example", TrailArgs.builder()        
            .advancedEventSelectors(            
                TrailAdvancedEventSelectorArgs.builder()
                    .fieldSelectors(                    
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .equals("Data")
                            .field("eventCategory")
                            .build(),
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .equals(                            
                                "PutObject",
                                "DeleteObject")
                            .field("eventName")
                            .build(),
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .field("resources.ARN")
                            .startsWith(                            
                                String.format("%s/", important_bucket_1.arn()),
                                String.format("%s/", important_bucket_2.arn()))
                            .build(),
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .equals("false")
                            .field("readOnly")
                            .build(),
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .equals("AWS::S3::Object")
                            .field("resources.type")
                            .build())
                    .name("Log PutObject and DeleteObject events for two S3 buckets")
                    .build(),
                TrailAdvancedEventSelectorArgs.builder()
                    .fieldSelectors(                    
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .equals("Data")
                            .field("eventCategory")
                            .build(),
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .field("eventName")
                            .startsWith("Delete")
                            .build(),
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .equals(String.format("%s/important-prefix", important_bucket_3.arn()))
                            .field("resources.ARN")
                            .build(),
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .equals("false")
                            .field("readOnly")
                            .build(),
                        TrailAdvancedEventSelectorFieldSelectorArgs.builder()
                            .equals("AWS::S3::Object")
                            .field("resources.type")
                            .build())
                    .name("Log Delete* events for one S3 bucket")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudtrail:Trail
    properties:
      advancedEventSelectors:
        - fieldSelectors:
            - equals:
                - Data
              field: eventCategory
            - equals:
                - PutObject
                - DeleteObject
              field: eventName
            - field: resources.ARN
              startsWith:
                - ${["important-bucket-1"].arn}/
                - ${["important-bucket-2"].arn}/
            - equals:
                - 'false'
              field: readOnly
            - equals:
                - AWS::S3::Object
              field: resources.type
          name: Log PutObject and DeleteObject events for two S3 buckets
        - fieldSelectors:
            - equals:
                - Data
              field: eventCategory
            - field: eventName
              startsWith:
                - Delete
            - equals:
                - ${["important-bucket-3"].arn}/important-prefix
              field: resources.ARN
            - equals:
                - 'false'
              field: readOnly
            - equals:
                - AWS::S3::Object
              field: resources.type
          name: Log Delete* events for one S3 bucket
variables:
  important-bucket-1:
    fn::invoke:
      Function: aws:s3:getBucket
      Arguments:
        bucket: important-bucket-1
  important-bucket-2:
    fn::invoke:
      Function: aws:s3:getBucket
      Arguments:
        bucket: important-bucket-2
  important-bucket-3:
    fn::invoke:
      Function: aws:s3:getBucket
      Arguments:
        bucket: important-bucket-3
```
{{% /example %}}
{{% example %}}
### Sending Events to CloudWatch Logs

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
const exampleTrail = new aws.cloudtrail.Trail("exampleTrail", {cloudWatchLogsGroupArn: pulumi.interpolate`${exampleLogGroup.arn}:*`});
// CloudTrail requires the Log Stream wildcard
```
```python
import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
example_trail = aws.cloudtrail.Trail("exampleTrail", cloud_watch_logs_group_arn=example_log_group.arn.apply(lambda arn: f"{arn}:*"))
# CloudTrail requires the Log Stream wildcard
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");

    var exampleTrail = new Aws.CloudTrail.Trail("exampleTrail", new()
    {
        CloudWatchLogsGroupArn = exampleLogGroup.Arn.Apply(arn => $"{arn}:*"),
    });

    // CloudTrail requires the Log Stream wildcard
});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudtrail"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
		if err != nil {
			return err
		}
		_, err = cloudtrail.NewTrail(ctx, "exampleTrail", &cloudtrail.TrailArgs{
			CloudWatchLogsGroupArn: exampleLogGroup.Arn.ApplyT(func(arn string) (string, error) {
				return fmt.Sprintf("%v:*", arn), nil
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.cloudtrail.Trail;
import com.pulumi.aws.cloudtrail.TrailArgs;
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
        var exampleLogGroup = new LogGroup("exampleLogGroup");

        var exampleTrail = new Trail("exampleTrail", TrailArgs.builder()        
            .cloudWatchLogsGroupArn(exampleLogGroup.arn().applyValue(arn -> String.format("%s:*", arn)))
            .build());

    }
}
```
```yaml
resources:
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
  exampleTrail:
    type: aws:cloudtrail:Trail
    properties:
      cloudWatchLogsGroupArn: ${exampleLogGroup.arn}:*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cloudtrails using the `arn`. For example:

```sh
 $ pulumi import aws:cloudtrail/trail:Trail sample arn:aws:cloudtrail:us-east-1:123456789012:trail/my-sample-trail
```
 