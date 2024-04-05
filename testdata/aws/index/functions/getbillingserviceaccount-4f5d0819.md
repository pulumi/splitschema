Use this data source to get the Account ID of the [AWS Billing and Cost Management Service Account](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-getting-started.html#step-2) for the purpose of permitting in S3 bucket policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = aws.getBillingServiceAccount({});
const billingLogs = new aws.s3.BucketV2("billingLogs", {});
const billingLogsAcl = new aws.s3.BucketAclV2("billingLogsAcl", {
    bucket: billingLogs.id,
    acl: "private",
});
const allowBillingLoggingPolicyDocument = pulumi.all([main, billingLogs.arn, main, billingLogs.arn]).apply(([main, billingLogsArn, main1, billingLogsArn1]) => aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            effect: "Allow",
            principals: [{
                type: "AWS",
                identifiers: [main.arn],
            }],
            actions: [
                "s3:GetBucketAcl",
                "s3:GetBucketPolicy",
            ],
            resources: [billingLogsArn],
        },
        {
            effect: "Allow",
            principals: [{
                type: "AWS",
                identifiers: [main1.arn],
            }],
            actions: ["s3:PutObject"],
            resources: [`${billingLogsArn1}/*`],
        },
    ],
}));
const allowBillingLoggingBucketPolicy = new aws.s3.BucketPolicy("allowBillingLoggingBucketPolicy", {
    bucket: billingLogs.id,
    policy: allowBillingLoggingPolicyDocument.apply(allowBillingLoggingPolicyDocument => allowBillingLoggingPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.get_billing_service_account()
billing_logs = aws.s3.BucketV2("billingLogs")
billing_logs_acl = aws.s3.BucketAclV2("billingLogsAcl",
    bucket=billing_logs.id,
    acl="private")
allow_billing_logging_policy_document = pulumi.Output.all(billing_logs.arn, billing_logs.arn).apply(lambda billingLogsArn, billingLogsArn1: aws.iam.get_policy_document_output(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=[main.arn],
        )],
        actions=[
            "s3:GetBucketAcl",
            "s3:GetBucketPolicy",
        ],
        resources=[billing_logs_arn],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=[main.arn],
        )],
        actions=["s3:PutObject"],
        resources=[f"{billing_logs_arn1}/*"],
    ),
]))
allow_billing_logging_bucket_policy = aws.s3.BucketPolicy("allowBillingLoggingBucketPolicy",
    bucket=billing_logs.id,
    policy=allow_billing_logging_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = Aws.GetBillingServiceAccount.Invoke();

    var billingLogs = new Aws.S3.BucketV2("billingLogs");

    var billingLogsAcl = new Aws.S3.BucketAclV2("billingLogsAcl", new()
    {
        Bucket = billingLogs.Id,
        Acl = "private",
    });

    var allowBillingLoggingPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            main.Apply(getBillingServiceAccountResult => getBillingServiceAccountResult.Arn),
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:GetBucketAcl",
                    "s3:GetBucketPolicy",
                },
                Resources = new[]
                {
                    billingLogs.Arn,
                },
            },
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
                            main.Apply(getBillingServiceAccountResult => getBillingServiceAccountResult.Arn),
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:PutObject",
                },
                Resources = new[]
                {
                    $"{billingLogs.Arn}/*",
                },
            },
        },
    });

    var allowBillingLoggingBucketPolicy = new Aws.S3.BucketPolicy("allowBillingLoggingBucketPolicy", new()
    {
        Bucket = billingLogs.Id,
        Policy = allowBillingLoggingPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
main, err := aws.GetBillingServiceAccount(ctx, nil, nil);
if err != nil {
return err
}
billingLogs, err := s3.NewBucketV2(ctx, "billingLogs", nil)
if err != nil {
return err
}
_, err = s3.NewBucketAclV2(ctx, "billingLogsAcl", &s3.BucketAclV2Args{
Bucket: billingLogs.ID(),
Acl: pulumi.String("private"),
})
if err != nil {
return err
}
allowBillingLoggingPolicyDocument := pulumi.All(billingLogs.Arn,billingLogs.Arn).ApplyT(func(_args []interface{}) (iam.GetPolicyDocumentResult, error) {
billingLogsArn := _args[0].(string)
billingLogsArn1 := _args[1].(string)
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
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
"s3:GetBucketPolicy",
},
Resources: []string{
billingLogsArn,
},
},
{
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
fmt.Sprintf("%v/*", billingLogsArn1),
},
},
},
}, nil), nil
}).(iam.GetPolicyDocumentResultOutput)
_, err = s3.NewBucketPolicy(ctx, "allowBillingLoggingBucketPolicy", &s3.BucketPolicyArgs{
Bucket: billingLogs.ID(),
Policy: allowBillingLoggingPolicyDocument.ApplyT(func(allowBillingLoggingPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
return &allowBillingLoggingPolicyDocument.Json, nil
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetBillingServiceAccountArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
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
        final var main = AwsFunctions.getBillingServiceAccount();

        var billingLogs = new BucketV2("billingLogs");

        var billingLogsAcl = new BucketAclV2("billingLogsAcl", BucketAclV2Args.builder()        
            .bucket(billingLogs.id())
            .acl("private")
            .build());

        final var allowBillingLoggingPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("AWS")
                        .identifiers(main.applyValue(getBillingServiceAccountResult -> getBillingServiceAccountResult.arn()))
                        .build())
                    .actions(                    
                        "s3:GetBucketAcl",
                        "s3:GetBucketPolicy")
                    .resources(billingLogs.arn())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("AWS")
                        .identifiers(main.applyValue(getBillingServiceAccountResult -> getBillingServiceAccountResult.arn()))
                        .build())
                    .actions("s3:PutObject")
                    .resources(billingLogs.arn().applyValue(arn -> String.format("%s/*", arn)))
                    .build())
            .build());

        var allowBillingLoggingBucketPolicy = new BucketPolicy("allowBillingLoggingBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(billingLogs.id())
            .policy(allowBillingLoggingPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(allowBillingLoggingPolicyDocument -> allowBillingLoggingPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  billingLogs:
    type: aws:s3:BucketV2
  billingLogsAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${billingLogs.id}
      acl: private
  allowBillingLoggingBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${billingLogs.id}
      policy: ${allowBillingLoggingPolicyDocument.json}
variables:
  main:
    fn::invoke:
      Function: aws:getBillingServiceAccount
      Arguments: {}
  allowBillingLoggingPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - ${main.arn}
            actions:
              - s3:GetBucketAcl
              - s3:GetBucketPolicy
            resources:
              - ${billingLogs.arn}
          - effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - ${main.arn}
            actions:
              - s3:PutObject
            resources:
              - ${billingLogs.arn}/*
```
{{% /example %}}
{{% /examples %}}