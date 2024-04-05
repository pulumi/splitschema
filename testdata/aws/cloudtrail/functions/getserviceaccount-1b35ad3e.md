Use this data source to get the Account ID of the [AWS CloudTrail Service Account](http://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-supported-regions.html)
in a given region for the purpose of allowing CloudTrail to store trail data in S3.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = aws.cloudtrail.getServiceAccount({});
const bucket = new aws.s3.BucketV2("bucket", {forceDestroy: true});
const allowCloudtrailLoggingPolicyDocument = pulumi.all([main, bucket.arn, main, bucket.arn]).apply(([main, bucketArn, main1, bucketArn1]) => aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            sid: "Put bucket policy needed for trails",
            effect: "Allow",
            principals: [{
                type: "AWS",
                identifiers: [main.arn],
            }],
            actions: ["s3:PutObject"],
            resources: [`${bucketArn}/*`],
        },
        {
            sid: "Get bucket policy needed for trails",
            effect: "Allow",
            principals: [{
                type: "AWS",
                identifiers: [main1.arn],
            }],
            actions: ["s3:GetBucketAcl"],
            resources: [bucketArn1],
        },
    ],
}));
const allowCloudtrailLoggingBucketPolicy = new aws.s3.BucketPolicy("allowCloudtrailLoggingBucketPolicy", {
    bucket: bucket.id,
    policy: allowCloudtrailLoggingPolicyDocument.apply(allowCloudtrailLoggingPolicyDocument => allowCloudtrailLoggingPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.cloudtrail.get_service_account()
bucket = aws.s3.BucketV2("bucket", force_destroy=True)
allow_cloudtrail_logging_policy_document = pulumi.Output.all(bucket.arn, bucket.arn).apply(lambda bucketArn, bucketArn1: aws.iam.get_policy_document_output(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="Put bucket policy needed for trails",
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=[main.arn],
        )],
        actions=["s3:PutObject"],
        resources=[f"{bucket_arn}/*"],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="Get bucket policy needed for trails",
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=[main.arn],
        )],
        actions=["s3:GetBucketAcl"],
        resources=[bucket_arn1],
    ),
]))
allow_cloudtrail_logging_bucket_policy = aws.s3.BucketPolicy("allowCloudtrailLoggingBucketPolicy",
    bucket=bucket.id,
    policy=allow_cloudtrail_logging_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = Aws.CloudTrail.GetServiceAccount.Invoke();

    var bucket = new Aws.S3.BucketV2("bucket", new()
    {
        ForceDestroy = true,
    });

    var allowCloudtrailLoggingPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "Put bucket policy needed for trails",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            main.Apply(getServiceAccountResult => getServiceAccountResult.Arn),
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:PutObject",
                },
                Resources = new[]
                {
                    $"{bucket.Arn}/*",
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "Get bucket policy needed for trails",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            main.Apply(getServiceAccountResult => getServiceAccountResult.Arn),
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:GetBucketAcl",
                },
                Resources = new[]
                {
                    bucket.Arn,
                },
            },
        },
    });

    var allowCloudtrailLoggingBucketPolicy = new Aws.S3.BucketPolicy("allowCloudtrailLoggingBucketPolicy", new()
    {
        Bucket = bucket.Id,
        Policy = allowCloudtrailLoggingPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudtrail"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
main, err := cloudtrail.GetServiceAccount(ctx, nil, nil);
if err != nil {
return err
}
bucket, err := s3.NewBucketV2(ctx, "bucket", &s3.BucketV2Args{
ForceDestroy: pulumi.Bool(true),
})
if err != nil {
return err
}
allowCloudtrailLoggingPolicyDocument := pulumi.All(bucket.Arn,bucket.Arn).ApplyT(func(_args []interface{}) (iam.GetPolicyDocumentResult, error) {
bucketArn := _args[0].(string)
bucketArn1 := _args[1].(string)
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Sid: "Put bucket policy needed for trails",
Effect: "Allow",
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: interface{}{
main.Arn,
},
},
},
Actions: []string{
"s3:PutObject",
},
Resources: []string{
fmt.Sprintf("%v/*", bucketArn),
},
},
{
Sid: "Get bucket policy needed for trails",
Effect: "Allow",
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: interface{}{
main.Arn,
},
},
},
Actions: []string{
"s3:GetBucketAcl",
},
Resources: []string{
bucketArn1,
},
},
},
}, nil), nil
}).(iam.GetPolicyDocumentResultOutput)
_, err = s3.NewBucketPolicy(ctx, "allowCloudtrailLoggingBucketPolicy", &s3.BucketPolicyArgs{
Bucket: bucket.ID(),
Policy: allowCloudtrailLoggingPolicyDocument.ApplyT(func(allowCloudtrailLoggingPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
return &allowCloudtrailLoggingPolicyDocument.Json, nil
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
import com.pulumi.aws.cloudtrail.CloudtrailFunctions;
import com.pulumi.aws.cloudtrail.inputs.GetServiceAccountArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
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
        final var main = CloudtrailFunctions.getServiceAccount();

        var bucket = new BucketV2("bucket", BucketV2Args.builder()        
            .forceDestroy(true)
            .build());

        final var allowCloudtrailLoggingPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .sid("Put bucket policy needed for trails")
                    .effect("Allow")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("AWS")
                        .identifiers(main.applyValue(getServiceAccountResult -> getServiceAccountResult.arn()))
                        .build())
                    .actions("s3:PutObject")
                    .resources(bucket.arn().applyValue(arn -> String.format("%s/*", arn)))
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .sid("Get bucket policy needed for trails")
                    .effect("Allow")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("AWS")
                        .identifiers(main.applyValue(getServiceAccountResult -> getServiceAccountResult.arn()))
                        .build())
                    .actions("s3:GetBucketAcl")
                    .resources(bucket.arn())
                    .build())
            .build());

        var allowCloudtrailLoggingBucketPolicy = new BucketPolicy("allowCloudtrailLoggingBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(bucket.id())
            .policy(allowCloudtrailLoggingPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(allowCloudtrailLoggingPolicyDocument -> allowCloudtrailLoggingPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  bucket:
    type: aws:s3:BucketV2
    properties:
      forceDestroy: true
  allowCloudtrailLoggingBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${bucket.id}
      policy: ${allowCloudtrailLoggingPolicyDocument.json}
variables:
  main:
    fn::invoke:
      Function: aws:cloudtrail:getServiceAccount
      Arguments: {}
  allowCloudtrailLoggingPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: Put bucket policy needed for trails
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - ${main.arn}
            actions:
              - s3:PutObject
            resources:
              - ${bucket.arn}/*
          - sid: Get bucket policy needed for trails
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - ${main.arn}
            actions:
              - s3:GetBucketAcl
            resources:
              - ${bucket.arn}
```
{{% /example %}}
{{% /examples %}}