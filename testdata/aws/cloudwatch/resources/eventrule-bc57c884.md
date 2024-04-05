Provides an EventBridge Rule resource.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const console = new aws.cloudwatch.EventRule("console", {
    description: "Capture each AWS Console Sign In",
    eventPattern: JSON.stringify({
        "detail-type": ["AWS Console Sign In via CloudTrail"],
    }),
});
const awsLogins = new aws.sns.Topic("awsLogins", {});
const sns = new aws.cloudwatch.EventTarget("sns", {
    rule: console.name,
    arn: awsLogins.arn,
});
const snsTopicPolicy = awsLogins.arn.apply(arn => aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        actions: ["SNS:Publish"],
        principals: [{
            type: "Service",
            identifiers: ["events.amazonaws.com"],
        }],
        resources: [arn],
    }],
}));
const _default = new aws.sns.TopicPolicy("default", {
    arn: awsLogins.arn,
    policy: snsTopicPolicy.apply(snsTopicPolicy => snsTopicPolicy.json),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

console = aws.cloudwatch.EventRule("console",
    description="Capture each AWS Console Sign In",
    event_pattern=json.dumps({
        "detail-type": ["AWS Console Sign In via CloudTrail"],
    }))
aws_logins = aws.sns.Topic("awsLogins")
sns = aws.cloudwatch.EventTarget("sns",
    rule=console.name,
    arn=aws_logins.arn)
sns_topic_policy = aws_logins.arn.apply(lambda arn: aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["SNS:Publish"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["events.amazonaws.com"],
    )],
    resources=[arn],
)]))
default = aws.sns.TopicPolicy("default",
    arn=aws_logins.arn,
    policy=sns_topic_policy.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var console = new Aws.CloudWatch.EventRule("console", new()
    {
        Description = "Capture each AWS Console Sign In",
        EventPattern = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["detail-type"] = new[]
            {
                "AWS Console Sign In via CloudTrail",
            },
        }),
    });

    var awsLogins = new Aws.Sns.Topic("awsLogins");

    var sns = new Aws.CloudWatch.EventTarget("sns", new()
    {
        Rule = console.Name,
        Arn = awsLogins.Arn,
    });

    var snsTopicPolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "SNS:Publish",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "events.amazonaws.com",
                        },
                    },
                },
                Resources = new[]
                {
                    awsLogins.Arn,
                },
            },
        },
    });

    var @default = new Aws.Sns.TopicPolicy("default", new()
    {
        Arn = awsLogins.Arn,
        Policy = snsTopicPolicy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
tmpJSON0, err := json.Marshal(map[string]interface{}{
"detail-type": []string{
"AWS Console Sign In via CloudTrail",
},
})
if err != nil {
return err
}
json0 := string(tmpJSON0)
console, err := cloudwatch.NewEventRule(ctx, "console", &cloudwatch.EventRuleArgs{
Description: pulumi.String("Capture each AWS Console Sign In"),
EventPattern: pulumi.String(json0),
})
if err != nil {
return err
}
awsLogins, err := sns.NewTopic(ctx, "awsLogins", nil)
if err != nil {
return err
}
_, err = cloudwatch.NewEventTarget(ctx, "sns", &cloudwatch.EventTargetArgs{
Rule: console.Name,
Arn: awsLogins.Arn,
})
if err != nil {
return err
}
snsTopicPolicy := awsLogins.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Effect: "Allow",
Actions: []string{
"SNS:Publish",
},
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "Service",
Identifiers: []string{
"events.amazonaws.com",
},
},
},
Resources: interface{}{
arn,
},
},
},
}, nil), nil
}).(iam.GetPolicyDocumentResultOutput)
_, err = sns.NewTopicPolicy(ctx, "default", &sns.TopicPolicyArgs{
Arn: awsLogins.Arn,
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
import com.pulumi.aws.cloudwatch.EventRule;
import com.pulumi.aws.cloudwatch.EventRuleArgs;
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.cloudwatch.EventTarget;
import com.pulumi.aws.cloudwatch.EventTargetArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.sns.TopicPolicy;
import com.pulumi.aws.sns.TopicPolicyArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var console = new EventRule("console", EventRuleArgs.builder()        
            .description("Capture each AWS Console Sign In")
            .eventPattern(serializeJson(
                jsonObject(
                    jsonProperty("detail-type", jsonArray("AWS Console Sign In via CloudTrail"))
                )))
            .build());

        var awsLogins = new Topic("awsLogins");

        var sns = new EventTarget("sns", EventTargetArgs.builder()        
            .rule(console.name())
            .arn(awsLogins.arn())
            .build());

        final var snsTopicPolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("SNS:Publish")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("events.amazonaws.com")
                    .build())
                .resources(awsLogins.arn())
                .build())
            .build());

        var default_ = new TopicPolicy("default", TopicPolicyArgs.builder()        
            .arn(awsLogins.arn())
            .policy(snsTopicPolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(snsTopicPolicy -> snsTopicPolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  console:
    type: aws:cloudwatch:EventRule
    properties:
      description: Capture each AWS Console Sign In
      eventPattern:
        fn::toJSON:
          detail-type:
            - AWS Console Sign In via CloudTrail
  sns:
    type: aws:cloudwatch:EventTarget
    properties:
      rule: ${console.name}
      arn: ${awsLogins.arn}
  awsLogins:
    type: aws:sns:Topic
  default:
    type: aws:sns:TopicPolicy
    properties:
      arn: ${awsLogins.arn}
      policy: ${snsTopicPolicy.json}
variables:
  snsTopicPolicy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - SNS:Publish
            principals:
              - type: Service
                identifiers:
                  - events.amazonaws.com
            resources:
              - ${awsLogins.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge Rules using the `event_bus_name/rule_name` (if you omit `event_bus_name`, the `default` event bus will be used). For example:

```sh
 $ pulumi import aws:cloudwatch/eventRule:EventRule console example-event-bus/capture-console-sign-in
```
 