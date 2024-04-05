Provides a CodeStar Notifications Rule.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const code = new aws.codecommit.Repository("code", {repositoryName: "example-code-repo"});
const notif = new aws.sns.Topic("notif", {});
const notifAccess = notif.arn.apply(arn => aws.iam.getPolicyDocumentOutput({
    statements: [{
        actions: ["sns:Publish"],
        principals: [{
            type: "Service",
            identifiers: ["codestar-notifications.amazonaws.com"],
        }],
        resources: [arn],
    }],
}));
const _default = new aws.sns.TopicPolicy("default", {
    arn: notif.arn,
    policy: notifAccess.apply(notifAccess => notifAccess.json),
});
const commits = new aws.codestarnotifications.NotificationRule("commits", {
    detailType: "BASIC",
    eventTypeIds: ["codecommit-repository-comments-on-commits"],
    resource: code.arn,
    targets: [{
        address: notif.arn,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

code = aws.codecommit.Repository("code", repository_name="example-code-repo")
notif = aws.sns.Topic("notif")
notif_access = notif.arn.apply(lambda arn: aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sns:Publish"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["codestar-notifications.amazonaws.com"],
    )],
    resources=[arn],
)]))
default = aws.sns.TopicPolicy("default",
    arn=notif.arn,
    policy=notif_access.json)
commits = aws.codestarnotifications.NotificationRule("commits",
    detail_type="BASIC",
    event_type_ids=["codecommit-repository-comments-on-commits"],
    resource=code.arn,
    targets=[aws.codestarnotifications.NotificationRuleTargetArgs(
        address=notif.arn,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var code = new Aws.CodeCommit.Repository("code", new()
    {
        RepositoryName = "example-code-repo",
    });

    var notif = new Aws.Sns.Topic("notif");

    var notifAccess = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sns:Publish",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "codestar-notifications.amazonaws.com",
                        },
                    },
                },
                Resources = new[]
                {
                    notif.Arn,
                },
            },
        },
    });

    var @default = new Aws.Sns.TopicPolicy("default", new()
    {
        Arn = notif.Arn,
        Policy = notifAccess.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var commits = new Aws.CodeStarNotifications.NotificationRule("commits", new()
    {
        DetailType = "BASIC",
        EventTypeIds = new[]
        {
            "codecommit-repository-comments-on-commits",
        },
        Resource = code.Arn,
        Targets = new[]
        {
            new Aws.CodeStarNotifications.Inputs.NotificationRuleTargetArgs
            {
                Address = notif.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codecommit"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codestarnotifications"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
code, err := codecommit.NewRepository(ctx, "code", &codecommit.RepositoryArgs{
RepositoryName: pulumi.String("example-code-repo"),
})
if err != nil {
return err
}
notif, err := sns.NewTopic(ctx, "notif", nil)
if err != nil {
return err
}
notifAccess := notif.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Actions: []string{
"sns:Publish",
},
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "Service",
Identifiers: []string{
"codestar-notifications.amazonaws.com",
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
Arn: notif.Arn,
Policy: notifAccess.ApplyT(func(notifAccess iam.GetPolicyDocumentResult) (*string, error) {
return &notifAccess.Json, nil
}).(pulumi.StringPtrOutput),
})
if err != nil {
return err
}
_, err = codestarnotifications.NewNotificationRule(ctx, "commits", &codestarnotifications.NotificationRuleArgs{
DetailType: pulumi.String("BASIC"),
EventTypeIds: pulumi.StringArray{
pulumi.String("codecommit-repository-comments-on-commits"),
},
Resource: code.Arn,
Targets: codestarnotifications.NotificationRuleTargetArray{
&codestarnotifications.NotificationRuleTargetArgs{
Address: notif.Arn,
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
import com.pulumi.aws.codecommit.Repository;
import com.pulumi.aws.codecommit.RepositoryArgs;
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.sns.TopicPolicy;
import com.pulumi.aws.sns.TopicPolicyArgs;
import com.pulumi.aws.codestarnotifications.NotificationRule;
import com.pulumi.aws.codestarnotifications.NotificationRuleArgs;
import com.pulumi.aws.codestarnotifications.inputs.NotificationRuleTargetArgs;
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
        var code = new Repository("code", RepositoryArgs.builder()        
            .repositoryName("example-code-repo")
            .build());

        var notif = new Topic("notif");

        final var notifAccess = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sns:Publish")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("codestar-notifications.amazonaws.com")
                    .build())
                .resources(notif.arn())
                .build())
            .build());

        var default_ = new TopicPolicy("default", TopicPolicyArgs.builder()        
            .arn(notif.arn())
            .policy(notifAccess.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(notifAccess -> notifAccess.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var commits = new NotificationRule("commits", NotificationRuleArgs.builder()        
            .detailType("BASIC")
            .eventTypeIds("codecommit-repository-comments-on-commits")
            .resource(code.arn())
            .targets(NotificationRuleTargetArgs.builder()
                .address(notif.arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  code:
    type: aws:codecommit:Repository
    properties:
      repositoryName: example-code-repo
  notif:
    type: aws:sns:Topic
  default:
    type: aws:sns:TopicPolicy
    properties:
      arn: ${notif.arn}
      policy: ${notifAccess.json}
  commits:
    type: aws:codestarnotifications:NotificationRule
    properties:
      detailType: BASIC
      eventTypeIds:
        - codecommit-repository-comments-on-commits
      resource: ${code.arn}
      targets:
        - address: ${notif.arn}
variables:
  notifAccess:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sns:Publish
            principals:
              - type: Service
                identifiers:
                  - codestar-notifications.amazonaws.com
            resources:
              - ${notif.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeStar notification rule using the ARN. For example:

```sh
 $ pulumi import aws:codestarnotifications/notificationRule:NotificationRule foo arn:aws:codestar-notifications:us-west-1:0123456789:notificationrule/2cdc68a3-8f7c-4893-b6a5-45b362bd4f2b
```
 