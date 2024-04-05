Get information about the organization that the user's account belongs to

{{% examples %}}
## Example Usage
{{% example %}}
### List all account IDs for the organization

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.organizations.getOrganization({});
export const accountIds = example.then(example => example.accounts.map(__item => __item.id));
```
```python
import pulumi
import pulumi_aws as aws

example = aws.organizations.get_organization()
pulumi.export("accountIds", [__item.id for __item in example.accounts])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Organizations.GetOrganization.Invoke();

    return new Dictionary<string, object?>
    {
        ["accountIds"] = example.Apply(getOrganizationResult => getOrganizationResult.Accounts).Select(__item => __item.Id).ToList(),
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := organizations.LookupOrganization(ctx, nil, nil)
		if err != nil {
			return err
		}
		var splat0 []*string
		for _, val0 := range example.Accounts {
			splat0 = append(splat0, val0.Id)
		}
		ctx.Export("accountIds", splat0)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.organizations.OrganizationsFunctions;
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
        final var example = OrganizationsFunctions.getOrganization();

        ctx.export("accountIds", example.applyValue(getOrganizationResult -> getOrganizationResult.accounts()).stream().map(element -> element.id()).collect(toList()));
    }
}
```
{{% /example %}}
{{% example %}}
### SNS topic that can be interacted by the organization only

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.organizations.getOrganization({});
const snsTopic = new aws.sns.Topic("snsTopic", {});
const snsTopicPolicyPolicyDocument = pulumi.all([example, snsTopic.arn]).apply(([example, arn]) => aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        actions: [
            "SNS:Subscribe",
            "SNS:Publish",
        ],
        conditions: [{
            test: "StringEquals",
            variable: "aws:PrincipalOrgID",
            values: [example.id],
        }],
        principals: [{
            type: "AWS",
            identifiers: ["*"],
        }],
        resources: [arn],
    }],
}));
const snsTopicPolicyTopicPolicy = new aws.sns.TopicPolicy("snsTopicPolicyTopicPolicy", {
    arn: snsTopic.arn,
    policy: snsTopicPolicyPolicyDocument.apply(snsTopicPolicyPolicyDocument => snsTopicPolicyPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.organizations.get_organization()
sns_topic = aws.sns.Topic("snsTopic")
sns_topic_policy_policy_document = sns_topic.arn.apply(lambda arn: aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=[
        "SNS:Subscribe",
        "SNS:Publish",
    ],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="StringEquals",
        variable="aws:PrincipalOrgID",
        values=[example.id],
    )],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["*"],
    )],
    resources=[arn],
)]))
sns_topic_policy_topic_policy = aws.sns.TopicPolicy("snsTopicPolicyTopicPolicy",
    arn=sns_topic.arn,
    policy=sns_topic_policy_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Organizations.GetOrganization.Invoke();

    var snsTopic = new Aws.Sns.Topic("snsTopic");

    var snsTopicPolicyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "SNS:Subscribe",
                    "SNS:Publish",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "StringEquals",
                        Variable = "aws:PrincipalOrgID",
                        Values = new[]
                        {
                            example.Apply(getOrganizationResult => getOrganizationResult.Id),
                        },
                    },
                },
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
                    snsTopic.Arn,
                },
            },
        },
    });

    var snsTopicPolicyTopicPolicy = new Aws.Sns.TopicPolicy("snsTopicPolicyTopicPolicy", new()
    {
        Arn = snsTopic.Arn,
        Policy = snsTopicPolicyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
example, err := organizations.LookupOrganization(ctx, nil, nil);
if err != nil {
return err
}
snsTopic, err := sns.NewTopic(ctx, "snsTopic", nil)
if err != nil {
return err
}
snsTopicPolicyPolicyDocument := snsTopic.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Effect: "Allow",
Actions: []string{
"SNS:Subscribe",
"SNS:Publish",
},
Conditions: []iam.GetPolicyDocumentStatementCondition{
{
Test: "StringEquals",
Variable: "aws:PrincipalOrgID",
Values: interface{}{
example.Id,
},
},
},
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
},
},
}, nil), nil
}).(iam.GetPolicyDocumentResultOutput)
_, err = sns.NewTopicPolicy(ctx, "snsTopicPolicyTopicPolicy", &sns.TopicPolicyArgs{
Arn: snsTopic.Arn,
Policy: snsTopicPolicyPolicyDocument.ApplyT(func(snsTopicPolicyPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
return &snsTopicPolicyPolicyDocument.Json, nil
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
import com.pulumi.aws.organizations.OrganizationsFunctions;
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
        final var example = OrganizationsFunctions.getOrganization();

        var snsTopic = new Topic("snsTopic");

        final var snsTopicPolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions(                
                    "SNS:Subscribe",
                    "SNS:Publish")
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("StringEquals")
                    .variable("aws:PrincipalOrgID")
                    .values(example.applyValue(getOrganizationResult -> getOrganizationResult.id()))
                    .build())
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("*")
                    .build())
                .resources(snsTopic.arn())
                .build())
            .build());

        var snsTopicPolicyTopicPolicy = new TopicPolicy("snsTopicPolicyTopicPolicy", TopicPolicyArgs.builder()        
            .arn(snsTopic.arn())
            .policy(snsTopicPolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(snsTopicPolicyPolicyDocument -> snsTopicPolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  snsTopic:
    type: aws:sns:Topic
  snsTopicPolicyTopicPolicy:
    type: aws:sns:TopicPolicy
    properties:
      arn: ${snsTopic.arn}
      policy: ${snsTopicPolicyPolicyDocument.json}
variables:
  example:
    fn::invoke:
      Function: aws:organizations:getOrganization
      Arguments: {}
  snsTopicPolicyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - SNS:Subscribe
              - SNS:Publish
            conditions:
              - test: StringEquals
                variable: aws:PrincipalOrgID
                values:
                  - ${example.id}
            principals:
              - type: AWS
                identifiers:
                  - '*'
            resources:
              - ${snsTopic.arn}
```
{{% /example %}}
{{% /examples %}}