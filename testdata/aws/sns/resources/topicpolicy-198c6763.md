Provides an SNS topic policy resource

> **NOTE:** If a Principal is specified as just an AWS account ID rather than an ARN, AWS silently converts it to the ARN for the root user, causing future deployments to differ. To avoid this problem, just specify the full ARN, e.g. `arn:aws:iam::123456789012:root`

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.sns.Topic("test", {});
const snsTopicPolicy = test.arn.apply(arn => aws.iam.getPolicyDocumentOutput({
    policyId: "__default_policy_ID",
    statements: [{
        actions: [
            "SNS:Subscribe",
            "SNS:SetTopicAttributes",
            "SNS:RemovePermission",
            "SNS:Receive",
            "SNS:Publish",
            "SNS:ListSubscriptionsByTopic",
            "SNS:GetTopicAttributes",
            "SNS:DeleteTopic",
            "SNS:AddPermission",
        ],
        conditions: [{
            test: "StringEquals",
            variable: "AWS:SourceOwner",
            values: [_var["account-id"]],
        }],
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: ["*"],
        }],
        resources: [arn],
        sid: "__default_statement_ID",
    }],
}));
const _default = new aws.sns.TopicPolicy("default", {
    arn: test.arn,
    policy: snsTopicPolicy.apply(snsTopicPolicy => snsTopicPolicy.json),
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.sns.Topic("test")
sns_topic_policy = test.arn.apply(lambda arn: aws.iam.get_policy_document_output(policy_id="__default_policy_ID",
    statements=[aws.iam.GetPolicyDocumentStatementArgs(
        actions=[
            "SNS:Subscribe",
            "SNS:SetTopicAttributes",
            "SNS:RemovePermission",
            "SNS:Receive",
            "SNS:Publish",
            "SNS:ListSubscriptionsByTopic",
            "SNS:GetTopicAttributes",
            "SNS:DeleteTopic",
            "SNS:AddPermission",
        ],
        conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
            test="StringEquals",
            variable="AWS:SourceOwner",
            values=[var["account-id"]],
        )],
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=["*"],
        )],
        resources=[arn],
        sid="__default_statement_ID",
    )]))
default = aws.sns.TopicPolicy("default",
    arn=test.arn,
    policy=sns_topic_policy.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Sns.Topic("test");

    var snsTopicPolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        PolicyId = "__default_policy_ID",
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "SNS:Subscribe",
                    "SNS:SetTopicAttributes",
                    "SNS:RemovePermission",
                    "SNS:Receive",
                    "SNS:Publish",
                    "SNS:ListSubscriptionsByTopic",
                    "SNS:GetTopicAttributes",
                    "SNS:DeleteTopic",
                    "SNS:AddPermission",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "StringEquals",
                        Variable = "AWS:SourceOwner",
                        Values = new[]
                        {
                            @var.Account_id,
                        },
                    },
                },
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            "*",
                        },
                    },
                },
                Resources = new[]
                {
                    test.Arn,
                },
                Sid = "__default_statement_ID",
            },
        },
    });

    var @default = new Aws.Sns.TopicPolicy("default", new()
    {
        Arn = test.Arn,
        Policy = snsTopicPolicy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
test, err := sns.NewTopic(ctx, "test", nil)
if err != nil {
return err
}
snsTopicPolicy := test.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
PolicyId: "__default_policy_ID",
Statements: []iam.GetPolicyDocumentStatement{
{
Actions: []string{
"SNS:Subscribe",
"SNS:SetTopicAttributes",
"SNS:RemovePermission",
"SNS:Receive",
"SNS:Publish",
"SNS:ListSubscriptionsByTopic",
"SNS:GetTopicAttributes",
"SNS:DeleteTopic",
"SNS:AddPermission",
},
Conditions: []iam.GetPolicyDocumentStatementCondition{
{
Test: "StringEquals",
Variable: "AWS:SourceOwner",
Values: interface{}{
_var.AccountId,
},
},
},
Effect: "Allow",
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: []string{
"*",
},
},
},
Resources: interface{}{
arn,
},
Sid: "__default_statement_ID",
},
},
}, nil), nil
}).(iam.GetPolicyDocumentResultOutput)
_, err = sns.NewTopicPolicy(ctx, "default", &sns.TopicPolicyArgs{
Arn: test.Arn,
Policy: snsTopicPolicy.ApplyT(func(snsTopicPolicy iam.GetPolicyDocumentResult) (*string, error) {
return &snsTopicPolicy.Json, nil
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
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.sns.TopicPolicy;
import com.pulumi.aws.sns.TopicPolicyArgs;
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
        var test = new Topic("test");

        final var snsTopicPolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .policyId("__default_policy_ID")
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions(                
                    "SNS:Subscribe",
                    "SNS:SetTopicAttributes",
                    "SNS:RemovePermission",
                    "SNS:Receive",
                    "SNS:Publish",
                    "SNS:ListSubscriptionsByTopic",
                    "SNS:GetTopicAttributes",
                    "SNS:DeleteTopic",
                    "SNS:AddPermission")
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("StringEquals")
                    .variable("AWS:SourceOwner")
                    .values(var_.account-id())
                    .build())
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("*")
                    .build())
                .resources(test.arn())
                .sid("__default_statement_ID")
                .build())
            .build());

        var default_ = new TopicPolicy("default", TopicPolicyArgs.builder()        
            .arn(test.arn())
            .policy(snsTopicPolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(snsTopicPolicy -> snsTopicPolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:sns:Topic
  default:
    type: aws:sns:TopicPolicy
    properties:
      arn: ${test.arn}
      policy: ${snsTopicPolicy.json}
variables:
  snsTopicPolicy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        policyId: __default_policy_ID
        statements:
          - actions:
              - SNS:Subscribe
              - SNS:SetTopicAttributes
              - SNS:RemovePermission
              - SNS:Receive
              - SNS:Publish
              - SNS:ListSubscriptionsByTopic
              - SNS:GetTopicAttributes
              - SNS:DeleteTopic
              - SNS:AddPermission
            conditions:
              - test: StringEquals
                variable: AWS:SourceOwner
                values:
                  - ${var"account-id"[%!s(MISSING)]}
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - '*'
            resources:
              - ${test.arn}
            sid: __default_statement_ID
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SNS Topic Policy using the topic ARN. For example:

```sh
 $ pulumi import aws:sns/topicPolicy:TopicPolicy user_updates arn:aws:sns:us-west-2:0123456789012:my-topic
```
 