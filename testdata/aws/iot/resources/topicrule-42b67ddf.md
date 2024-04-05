{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mytopic = new aws.sns.Topic("mytopic", {});
const myerrortopic = new aws.sns.Topic("myerrortopic", {});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["iot.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const role = new aws.iam.Role("role", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const rule = new aws.iot.TopicRule("rule", {
    description: "Example rule",
    enabled: true,
    sql: "SELECT * FROM 'topic/test'",
    sqlVersion: "2016-03-23",
    sns: [{
        messageFormat: "RAW",
        roleArn: role.arn,
        targetArn: mytopic.arn,
    }],
    errorAction: {
        sns: {
            messageFormat: "RAW",
            roleArn: role.arn,
            targetArn: myerrortopic.arn,
        },
    },
});
const iamPolicyForLambdaPolicyDocument = mytopic.arn.apply(arn => aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        actions: ["sns:Publish"],
        resources: [arn],
    }],
}));
const iamPolicyForLambdaRolePolicy = new aws.iam.RolePolicy("iamPolicyForLambdaRolePolicy", {
    role: role.id,
    policy: iamPolicyForLambdaPolicyDocument.apply(iamPolicyForLambdaPolicyDocument => iamPolicyForLambdaPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

mytopic = aws.sns.Topic("mytopic")
myerrortopic = aws.sns.Topic("myerrortopic")
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["iot.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
role = aws.iam.Role("role", assume_role_policy=assume_role.json)
rule = aws.iot.TopicRule("rule",
    description="Example rule",
    enabled=True,
    sql="SELECT * FROM 'topic/test'",
    sql_version="2016-03-23",
    sns=[aws.iot.TopicRuleSnsArgs(
        message_format="RAW",
        role_arn=role.arn,
        target_arn=mytopic.arn,
    )],
    error_action=aws.iot.TopicRuleErrorActionArgs(
        sns=aws.iot.TopicRuleErrorActionSnsArgs(
            message_format="RAW",
            role_arn=role.arn,
            target_arn=myerrortopic.arn,
        ),
    ))
iam_policy_for_lambda_policy_document = mytopic.arn.apply(lambda arn: aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["sns:Publish"],
    resources=[arn],
)]))
iam_policy_for_lambda_role_policy = aws.iam.RolePolicy("iamPolicyForLambdaRolePolicy",
    role=role.id,
    policy=iam_policy_for_lambda_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mytopic = new Aws.Sns.Topic("mytopic");

    var myerrortopic = new Aws.Sns.Topic("myerrortopic");

    var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            "iot.amazonaws.com",
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

    var role = new Aws.Iam.Role("role", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var rule = new Aws.Iot.TopicRule("rule", new()
    {
        Description = "Example rule",
        Enabled = true,
        Sql = "SELECT * FROM 'topic/test'",
        SqlVersion = "2016-03-23",
        Sns = new[]
        {
            new Aws.Iot.Inputs.TopicRuleSnsArgs
            {
                MessageFormat = "RAW",
                RoleArn = role.Arn,
                TargetArn = mytopic.Arn,
            },
        },
        ErrorAction = new Aws.Iot.Inputs.TopicRuleErrorActionArgs
        {
            Sns = new Aws.Iot.Inputs.TopicRuleErrorActionSnsArgs
            {
                MessageFormat = "RAW",
                RoleArn = role.Arn,
                TargetArn = myerrortopic.Arn,
            },
        },
    });

    var iamPolicyForLambdaPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "sns:Publish",
                },
                Resources = new[]
                {
                    mytopic.Arn,
                },
            },
        },
    });

    var iamPolicyForLambdaRolePolicy = new Aws.Iam.RolePolicy("iamPolicyForLambdaRolePolicy", new()
    {
        Role = role.Id,
        Policy = iamPolicyForLambdaPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
mytopic, err := sns.NewTopic(ctx, "mytopic", nil)
if err != nil {
return err
}
myerrortopic, err := sns.NewTopic(ctx, "myerrortopic", nil)
if err != nil {
return err
}
assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Effect: pulumi.StringRef("Allow"),
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "Service",
Identifiers: []string{
"iot.amazonaws.com",
},
},
},
Actions: []string{
"sts:AssumeRole",
},
},
},
}, nil);
if err != nil {
return err
}
role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
AssumeRolePolicy: *pulumi.String(assumeRole.Json),
})
if err != nil {
return err
}
_, err = iot.NewTopicRule(ctx, "rule", &iot.TopicRuleArgs{
Description: pulumi.String("Example rule"),
Enabled: pulumi.Bool(true),
Sql: pulumi.String("SELECT * FROM 'topic/test'"),
SqlVersion: pulumi.String("2016-03-23"),
Sns: iot.TopicRuleSnsArray{
&iot.TopicRuleSnsArgs{
MessageFormat: pulumi.String("RAW"),
RoleArn: role.Arn,
TargetArn: mytopic.Arn,
},
},
ErrorAction: &iot.TopicRuleErrorActionArgs{
Sns: &iot.TopicRuleErrorActionSnsArgs{
MessageFormat: pulumi.String("RAW"),
RoleArn: role.Arn,
TargetArn: myerrortopic.Arn,
},
},
})
if err != nil {
return err
}
iamPolicyForLambdaPolicyDocument := mytopic.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Effect: "Allow",
Actions: []string{
"sns:Publish",
},
Resources: interface{}{
arn,
},
},
},
}, nil), nil
}).(iam.GetPolicyDocumentResultOutput)
_, err = iam.NewRolePolicy(ctx, "iamPolicyForLambdaRolePolicy", &iam.RolePolicyArgs{
Role: role.ID(),
Policy: iamPolicyForLambdaPolicyDocument.ApplyT(func(iamPolicyForLambdaPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
return &iamPolicyForLambdaPolicyDocument.Json, nil
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iot.TopicRule;
import com.pulumi.aws.iot.TopicRuleArgs;
import com.pulumi.aws.iot.inputs.TopicRuleSnsArgs;
import com.pulumi.aws.iot.inputs.TopicRuleErrorActionArgs;
import com.pulumi.aws.iot.inputs.TopicRuleErrorActionSnsArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
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
        var mytopic = new Topic("mytopic");

        var myerrortopic = new Topic("myerrortopic");

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("iot.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var role = new Role("role", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var rule = new TopicRule("rule", TopicRuleArgs.builder()        
            .description("Example rule")
            .enabled(true)
            .sql("SELECT * FROM 'topic/test'")
            .sqlVersion("2016-03-23")
            .sns(TopicRuleSnsArgs.builder()
                .messageFormat("RAW")
                .roleArn(role.arn())
                .targetArn(mytopic.arn())
                .build())
            .errorAction(TopicRuleErrorActionArgs.builder()
                .sns(TopicRuleErrorActionSnsArgs.builder()
                    .messageFormat("RAW")
                    .roleArn(role.arn())
                    .targetArn(myerrortopic.arn())
                    .build())
                .build())
            .build());

        final var iamPolicyForLambdaPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("sns:Publish")
                .resources(mytopic.arn())
                .build())
            .build());

        var iamPolicyForLambdaRolePolicy = new RolePolicy("iamPolicyForLambdaRolePolicy", RolePolicyArgs.builder()        
            .role(role.id())
            .policy(iamPolicyForLambdaPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(iamPolicyForLambdaPolicyDocument -> iamPolicyForLambdaPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  rule:
    type: aws:iot:TopicRule
    properties:
      description: Example rule
      enabled: true
      sql: SELECT * FROM 'topic/test'
      sqlVersion: 2016-03-23
      sns:
        - messageFormat: RAW
          roleArn: ${role.arn}
          targetArn: ${mytopic.arn}
      errorAction:
        sns:
          messageFormat: RAW
          roleArn: ${role.arn}
          targetArn: ${myerrortopic.arn}
  mytopic:
    type: aws:sns:Topic
  myerrortopic:
    type: aws:sns:Topic
  role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  iamPolicyForLambdaRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${role.id}
      policy: ${iamPolicyForLambdaPolicyDocument.json}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - iot.amazonaws.com
            actions:
              - sts:AssumeRole
  iamPolicyForLambdaPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - sns:Publish
            resources:
              - ${mytopic.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IoT Topic Rules using the `name`. For example:

```sh
 $ pulumi import aws:iot/topicRule:TopicRule rule <name>
```
 