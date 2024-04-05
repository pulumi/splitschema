Attaches a policy to an S3 bucket resource.

> Policies can be attached to both S3 general purpose buckets and S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketV2("example", {});
const allowAccessFromAnotherAccountPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        principals: [{
            type: "AWS",
            identifiers: ["123456789012"],
        }],
        actions: [
            "s3:GetObject",
            "s3:ListBucket",
        ],
        resources: [
            example.arn,
            pulumi.interpolate`${example.arn}/*`,
        ],
    }],
});
const allowAccessFromAnotherAccountBucketPolicy = new aws.s3.BucketPolicy("allowAccessFromAnotherAccountBucketPolicy", {
    bucket: example.id,
    policy: allowAccessFromAnotherAccountPolicyDocument.apply(allowAccessFromAnotherAccountPolicyDocument => allowAccessFromAnotherAccountPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketV2("example")
allow_access_from_another_account_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["123456789012"],
    )],
    actions=[
        "s3:GetObject",
        "s3:ListBucket",
    ],
    resources=[
        example.arn,
        example.arn.apply(lambda arn: f"{arn}/*"),
    ],
)])
allow_access_from_another_account_bucket_policy = aws.s3.BucketPolicy("allowAccessFromAnotherAccountBucketPolicy",
    bucket=example.id,
    policy=allow_access_from_another_account_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketV2("example");

    var allowAccessFromAnotherAccountPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
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
                    "s3:GetObject",
                    "s3:ListBucket",
                },
                Resources = new[]
                {
                    example.Arn,
                    $"{example.Arn}/*",
                },
            },
        },
    });

    var allowAccessFromAnotherAccountBucketPolicy = new Aws.S3.BucketPolicy("allowAccessFromAnotherAccountBucketPolicy", new()
    {
        Bucket = example.Id,
        Policy = allowAccessFromAnotherAccountPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := s3.NewBucketV2(ctx, "example", nil)
		if err != nil {
			return err
		}
		allowAccessFromAnotherAccountPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("AWS"),
							Identifiers: pulumi.StringArray{
								pulumi.String("123456789012"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("s3:GetObject"),
						pulumi.String("s3:ListBucket"),
					},
					Resources: pulumi.StringArray{
						example.Arn,
						example.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
				},
			},
		}, nil)
		_, err = s3.NewBucketPolicy(ctx, "allowAccessFromAnotherAccountBucketPolicy", &s3.BucketPolicyArgs{
			Bucket: example.ID(),
			Policy: allowAccessFromAnotherAccountPolicyDocument.ApplyT(func(allowAccessFromAnotherAccountPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &allowAccessFromAnotherAccountPolicyDocument.Json, nil
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
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
        var example = new BucketV2("example");

        final var allowAccessFromAnotherAccountPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("123456789012")
                    .build())
                .actions(                
                    "s3:GetObject",
                    "s3:ListBucket")
                .resources(                
                    example.arn(),
                    example.arn().applyValue(arn -> String.format("%s/*", arn)))
                .build())
            .build());

        var allowAccessFromAnotherAccountBucketPolicy = new BucketPolicy("allowAccessFromAnotherAccountBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(example.id())
            .policy(allowAccessFromAnotherAccountPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(allowAccessFromAnotherAccountPolicyDocument -> allowAccessFromAnotherAccountPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketV2
  allowAccessFromAnotherAccountBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${example.id}
      policy: ${allowAccessFromAnotherAccountPolicyDocument.json}
variables:
  allowAccessFromAnotherAccountPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - principals:
              - type: AWS
                identifiers:
                  - '123456789012'
            actions:
              - s3:GetObject
              - s3:ListBucket
            resources:
              - ${example.arn}
              - ${example.arn}/*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 bucket policies using the bucket name. For example:

```sh
 $ pulumi import aws:s3/bucketPolicy:BucketPolicy allow_access_from_another_account my-tf-test-bucket
```
 