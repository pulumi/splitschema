Provides a SSM resource data sync.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const hogeBucketV2 = new aws.s3.BucketV2("hogeBucketV2", {});
const hogePolicyDocument = aws.iam.getPolicyDocument({
    statements: [
        {
            sid: "SSMBucketPermissionsCheck",
            effect: "Allow",
            principals: [{
                type: "Service",
                identifiers: ["ssm.amazonaws.com"],
            }],
            actions: ["s3:GetBucketAcl"],
            resources: ["arn:aws:s3:::tf-test-bucket-1234"],
        },
        {
            sid: "SSMBucketDelivery",
            effect: "Allow",
            principals: [{
                type: "Service",
                identifiers: ["ssm.amazonaws.com"],
            }],
            actions: ["s3:PutObject"],
            resources: ["arn:aws:s3:::tf-test-bucket-1234/*"],
            conditions: [{
                test: "StringEquals",
                variable: "s3:x-amz-acl",
                values: ["bucket-owner-full-control"],
            }],
        },
    ],
});
const hogeBucketPolicy = new aws.s3.BucketPolicy("hogeBucketPolicy", {
    bucket: hogeBucketV2.id,
    policy: hogePolicyDocument.then(hogePolicyDocument => hogePolicyDocument.json),
});
const foo = new aws.ssm.ResourceDataSync("foo", {s3Destination: {
    bucketName: hogeBucketV2.bucket,
    region: hogeBucketV2.region,
}});
```
```python
import pulumi
import pulumi_aws as aws

hoge_bucket_v2 = aws.s3.BucketV2("hogeBucketV2")
hoge_policy_document = aws.iam.get_policy_document(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="SSMBucketPermissionsCheck",
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=["ssm.amazonaws.com"],
        )],
        actions=["s3:GetBucketAcl"],
        resources=["arn:aws:s3:::tf-test-bucket-1234"],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="SSMBucketDelivery",
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=["ssm.amazonaws.com"],
        )],
        actions=["s3:PutObject"],
        resources=["arn:aws:s3:::tf-test-bucket-1234/*"],
        conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
            test="StringEquals",
            variable="s3:x-amz-acl",
            values=["bucket-owner-full-control"],
        )],
    ),
])
hoge_bucket_policy = aws.s3.BucketPolicy("hogeBucketPolicy",
    bucket=hoge_bucket_v2.id,
    policy=hoge_policy_document.json)
foo = aws.ssm.ResourceDataSync("foo", s3_destination=aws.ssm.ResourceDataSyncS3DestinationArgs(
    bucket_name=hoge_bucket_v2.bucket,
    region=hoge_bucket_v2.region,
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var hogeBucketV2 = new Aws.S3.BucketV2("hogeBucketV2");

    var hogePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "SSMBucketPermissionsCheck",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "ssm.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:GetBucketAcl",
                },
                Resources = new[]
                {
                    "arn:aws:s3:::tf-test-bucket-1234",
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "SSMBucketDelivery",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "ssm.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:PutObject",
                },
                Resources = new[]
                {
                    "arn:aws:s3:::tf-test-bucket-1234/*",
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
                },
            },
        },
    });

    var hogeBucketPolicy = new Aws.S3.BucketPolicy("hogeBucketPolicy", new()
    {
        Bucket = hogeBucketV2.Id,
        Policy = hogePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var foo = new Aws.Ssm.ResourceDataSync("foo", new()
    {
        S3Destination = new Aws.Ssm.Inputs.ResourceDataSyncS3DestinationArgs
        {
            BucketName = hogeBucketV2.Bucket,
            Region = hogeBucketV2.Region,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		hogeBucketV2, err := s3.NewBucketV2(ctx, "hogeBucketV2", nil)
		if err != nil {
			return err
		}
		hogePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("SSMBucketPermissionsCheck"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"ssm.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"s3:GetBucketAcl",
					},
					Resources: []string{
						"arn:aws:s3:::tf-test-bucket-1234",
					},
				},
				{
					Sid:    pulumi.StringRef("SSMBucketDelivery"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"ssm.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"s3:PutObject",
					},
					Resources: []string{
						"arn:aws:s3:::tf-test-bucket-1234/*",
					},
					Conditions: []iam.GetPolicyDocumentStatementCondition{
						{
							Test:     "StringEquals",
							Variable: "s3:x-amz-acl",
							Values: []string{
								"bucket-owner-full-control",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketPolicy(ctx, "hogeBucketPolicy", &s3.BucketPolicyArgs{
			Bucket: hogeBucketV2.ID(),
			Policy: *pulumi.String(hogePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = ssm.NewResourceDataSync(ctx, "foo", &ssm.ResourceDataSyncArgs{
			S3Destination: &ssm.ResourceDataSyncS3DestinationArgs{
				BucketName: hogeBucketV2.Bucket,
				Region:     hogeBucketV2.Region,
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
import com.pulumi.aws.ssm.ResourceDataSync;
import com.pulumi.aws.ssm.ResourceDataSyncArgs;
import com.pulumi.aws.ssm.inputs.ResourceDataSyncS3DestinationArgs;
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
        var hogeBucketV2 = new BucketV2("hogeBucketV2");

        final var hogePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .sid("SSMBucketPermissionsCheck")
                    .effect("Allow")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("Service")
                        .identifiers("ssm.amazonaws.com")
                        .build())
                    .actions("s3:GetBucketAcl")
                    .resources("arn:aws:s3:::tf-test-bucket-1234")
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .sid("SSMBucketDelivery")
                    .effect("Allow")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("Service")
                        .identifiers("ssm.amazonaws.com")
                        .build())
                    .actions("s3:PutObject")
                    .resources("arn:aws:s3:::tf-test-bucket-1234/*")
                    .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                        .test("StringEquals")
                        .variable("s3:x-amz-acl")
                        .values("bucket-owner-full-control")
                        .build())
                    .build())
            .build());

        var hogeBucketPolicy = new BucketPolicy("hogeBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(hogeBucketV2.id())
            .policy(hogePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var foo = new ResourceDataSync("foo", ResourceDataSyncArgs.builder()        
            .s3Destination(ResourceDataSyncS3DestinationArgs.builder()
                .bucketName(hogeBucketV2.bucket())
                .region(hogeBucketV2.region())
                .build())
            .build());

    }
}
```
```yaml
resources:
  hogeBucketV2:
    type: aws:s3:BucketV2
  hogeBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${hogeBucketV2.id}
      policy: ${hogePolicyDocument.json}
  foo:
    type: aws:ssm:ResourceDataSync
    properties:
      s3Destination:
        bucketName: ${hogeBucketV2.bucket}
        region: ${hogeBucketV2.region}
variables:
  hogePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: SSMBucketPermissionsCheck
            effect: Allow
            principals:
              - type: Service
                identifiers:
                  - ssm.amazonaws.com
            actions:
              - s3:GetBucketAcl
            resources:
              - arn:aws:s3:::tf-test-bucket-1234
          - sid: SSMBucketDelivery
            effect: Allow
            principals:
              - type: Service
                identifiers:
                  - ssm.amazonaws.com
            actions:
              - s3:PutObject
            resources:
              - arn:aws:s3:::tf-test-bucket-1234/*
            conditions:
              - test: StringEquals
                variable: s3:x-amz-acl
                values:
                  - bucket-owner-full-control
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SSM resource data sync using the `name`. For example:

```sh
 $ pulumi import aws:ssm/resourceDataSync:ResourceDataSync example example-name
```
 