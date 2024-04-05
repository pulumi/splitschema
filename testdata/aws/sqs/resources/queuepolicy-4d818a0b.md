Allows you to set a policy of an SQS Queue
while referencing ARN of the queue within the policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const queue = new aws.sqs.Queue("queue", {});
const testPolicyDocument = queue.arn.apply(arn => aws.iam.getPolicyDocumentOutput({
    statements: [{
        sid: "First",
        effect: "Allow",
        principals: [{
            type: "*",
            identifiers: ["*"],
        }],
        actions: ["sqs:SendMessage"],
        resources: [arn],
        conditions: [{
            test: "ArnEquals",
            variable: "aws:SourceArn",
            values: [aws_sns_topic.example.arn],
        }],
    }],
}));
const testQueuePolicy = new aws.sqs.QueuePolicy("testQueuePolicy", {
    queueUrl: queue.id,
    policy: testPolicyDocument.apply(testPolicyDocument => testPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

queue = aws.sqs.Queue("queue")
test_policy_document = queue.arn.apply(lambda arn: aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="First",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="*",
        identifiers=["*"],
    )],
    actions=["sqs:SendMessage"],
    resources=[arn],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="ArnEquals",
        variable="aws:SourceArn",
        values=[aws_sns_topic["example"]["arn"]],
    )],
)]))
test_queue_policy = aws.sqs.QueuePolicy("testQueuePolicy",
    queue_url=queue.id,
    policy=test_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var queue = new Aws.Sqs.Queue("queue");

    var testPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "First",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "*",
                        Identifiers = new[]
                        {
                            "*",
                        },
                    },
                },
                Actions = new[]
                {
                    "sqs:SendMessage",
                },
                Resources = new[]
                {
                    queue.Arn,
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "ArnEquals",
                        Variable = "aws:SourceArn",
                        Values = new[]
                        {
                            aws_sns_topic.Example.Arn,
                        },
                    },
                },
            },
        },
    });

    var testQueuePolicy = new Aws.Sqs.QueuePolicy("testQueuePolicy", new()
    {
        QueueUrl = queue.Id,
        Policy = testPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
queue, err := sqs.NewQueue(ctx, "queue", nil)
if err != nil {
return err
}
testPolicyDocument := queue.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Sid: "First",
Effect: "Allow",
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "*",
Identifiers: []string{
"*",
},
},
},
Actions: []string{
"sqs:SendMessage",
},
Resources: []string{
arn,
},
Conditions: []iam.GetPolicyDocumentStatementCondition{
{
Test: "ArnEquals",
Variable: "aws:SourceArn",
Values: interface{}{
aws_sns_topic.Example.Arn,
},
},
},
},
},
}, nil), nil
}).(iam.GetPolicyDocumentResultOutput)
_, err = sqs.NewQueuePolicy(ctx, "testQueuePolicy", &sqs.QueuePolicyArgs{
QueueUrl: queue.ID(),
Policy: testPolicyDocument.ApplyT(func(testPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
return &testPolicyDocument.Json, nil
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
import com.pulumi.aws.sqs.Queue;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.sqs.QueuePolicy;
import com.pulumi.aws.sqs.QueuePolicyArgs;
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
        var queue = new Queue("queue");

        final var testPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("First")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("*")
                    .identifiers("*")
                    .build())
                .actions("sqs:SendMessage")
                .resources(queue.arn())
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("ArnEquals")
                    .variable("aws:SourceArn")
                    .values(aws_sns_topic.example().arn())
                    .build())
                .build())
            .build());

        var testQueuePolicy = new QueuePolicy("testQueuePolicy", QueuePolicyArgs.builder()        
            .queueUrl(queue.id())
            .policy(testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(testPolicyDocument -> testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  queue:
    type: aws:sqs:Queue
  testQueuePolicy:
    type: aws:sqs:QueuePolicy
    properties:
      queueUrl: ${queue.id}
      policy: ${testPolicyDocument.json}
variables:
  testPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: First
            effect: Allow
            principals:
              - type: '*'
                identifiers:
                  - '*'
            actions:
              - sqs:SendMessage
            resources:
              - ${queue.arn}
            conditions:
              - test: ArnEquals
                variable: aws:SourceArn
                values:
                  - ${aws_sns_topic.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SQS Queue Policies using the queue URL. For example:

```sh
 $ pulumi import aws:sqs/queuePolicy:QueuePolicy test https://queue.amazonaws.com/0123456789012/myqueue
```
 