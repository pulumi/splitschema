Provides a resource to create an EventBridge resource policy to support cross-account events.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

> **Note:** The EventBridge bus policy resource  (`aws.cloudwatch.EventBusPolicy`) is incompatible with the EventBridge permission resource (`aws.cloudwatch.EventPermission`) and will overwrite permissions.

{{% examples %}}
## Example Usage
{{% example %}}
### Account Access

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        sid: "DevAccountAccess",
        effect: "Allow",
        actions: ["events:PutEvents"],
        resources: ["arn:aws:events:eu-west-1:123456789012:event-bus/default"],
        principals: [{
            type: "AWS",
            identifiers: ["123456789012"],
        }],
    }],
});
const testEventBusPolicy = new aws.cloudwatch.EventBusPolicy("testEventBusPolicy", {
    policy: testPolicyDocument.then(testPolicyDocument => testPolicyDocument.json),
    eventBusName: aws_cloudwatch_event_bus.test.name,
});
```
```python
import pulumi
import pulumi_aws as aws

test_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="DevAccountAccess",
    effect="Allow",
    actions=["events:PutEvents"],
    resources=["arn:aws:events:eu-west-1:123456789012:event-bus/default"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["123456789012"],
    )],
)])
test_event_bus_policy = aws.cloudwatch.EventBusPolicy("testEventBusPolicy",
    policy=test_policy_document.json,
    event_bus_name=aws_cloudwatch_event_bus["test"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "DevAccountAccess",
                Effect = "Allow",
                Actions = new[]
                {
                    "events:PutEvents",
                },
                Resources = new[]
                {
                    "arn:aws:events:eu-west-1:123456789012:event-bus/default",
                },
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
            },
        },
    });

    var testEventBusPolicy = new Aws.CloudWatch.EventBusPolicy("testEventBusPolicy", new()
    {
        Policy = testPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
        EventBusName = aws_cloudwatch_event_bus.Test.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("DevAccountAccess"),
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"events:PutEvents",
					},
					Resources: []string{
						"arn:aws:events:eu-west-1:123456789012:event-bus/default",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "AWS",
							Identifiers: []string{
								"123456789012",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventBusPolicy(ctx, "testEventBusPolicy", &cloudwatch.EventBusPolicyArgs{
			Policy:       *pulumi.String(testPolicyDocument.Json),
			EventBusName: pulumi.Any(aws_cloudwatch_event_bus.Test.Name),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.cloudwatch.EventBusPolicy;
import com.pulumi.aws.cloudwatch.EventBusPolicyArgs;
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
        final var testPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("DevAccountAccess")
                .effect("Allow")
                .actions("events:PutEvents")
                .resources("arn:aws:events:eu-west-1:123456789012:event-bus/default")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("123456789012")
                    .build())
                .build())
            .build());

        var testEventBusPolicy = new EventBusPolicy("testEventBusPolicy", EventBusPolicyArgs.builder()        
            .policy(testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .eventBusName(aws_cloudwatch_event_bus.test().name())
            .build());

    }
}
```
```yaml
resources:
  testEventBusPolicy:
    type: aws:cloudwatch:EventBusPolicy
    properties:
      policy: ${testPolicyDocument.json}
      eventBusName: ${aws_cloudwatch_event_bus.test.name}
variables:
  testPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: DevAccountAccess
            effect: Allow
            actions:
              - events:PutEvents
            resources:
              - arn:aws:events:eu-west-1:123456789012:event-bus/default
            principals:
              - type: AWS
                identifiers:
                  - '123456789012'
```
{{% /example %}}
{{% example %}}
### Organization Access

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        sid: "OrganizationAccess",
        effect: "Allow",
        actions: [
            "events:DescribeRule",
            "events:ListRules",
            "events:ListTargetsByRule",
            "events:ListTagsForResource",
        ],
        resources: [
            "arn:aws:events:eu-west-1:123456789012:rule/*",
            "arn:aws:events:eu-west-1:123456789012:event-bus/default",
        ],
        principals: [{
            type: "AWS",
            identifiers: ["*"],
        }],
        conditions: [{
            test: "StringEquals",
            variable: "aws:PrincipalOrgID",
            values: [aws_organizations_organization.example.id],
        }],
    }],
});
const testEventBusPolicy = new aws.cloudwatch.EventBusPolicy("testEventBusPolicy", {
    policy: testPolicyDocument.then(testPolicyDocument => testPolicyDocument.json),
    eventBusName: aws_cloudwatch_event_bus.test.name,
});
```
```python
import pulumi
import pulumi_aws as aws

test_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="OrganizationAccess",
    effect="Allow",
    actions=[
        "events:DescribeRule",
        "events:ListRules",
        "events:ListTargetsByRule",
        "events:ListTagsForResource",
    ],
    resources=[
        "arn:aws:events:eu-west-1:123456789012:rule/*",
        "arn:aws:events:eu-west-1:123456789012:event-bus/default",
    ],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["*"],
    )],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="StringEquals",
        variable="aws:PrincipalOrgID",
        values=[aws_organizations_organization["example"]["id"]],
    )],
)])
test_event_bus_policy = aws.cloudwatch.EventBusPolicy("testEventBusPolicy",
    policy=test_policy_document.json,
    event_bus_name=aws_cloudwatch_event_bus["test"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "OrganizationAccess",
                Effect = "Allow",
                Actions = new[]
                {
                    "events:DescribeRule",
                    "events:ListRules",
                    "events:ListTargetsByRule",
                    "events:ListTagsForResource",
                },
                Resources = new[]
                {
                    "arn:aws:events:eu-west-1:123456789012:rule/*",
                    "arn:aws:events:eu-west-1:123456789012:event-bus/default",
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
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "StringEquals",
                        Variable = "aws:PrincipalOrgID",
                        Values = new[]
                        {
                            aws_organizations_organization.Example.Id,
                        },
                    },
                },
            },
        },
    });

    var testEventBusPolicy = new Aws.CloudWatch.EventBusPolicy("testEventBusPolicy", new()
    {
        Policy = testPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
        EventBusName = aws_cloudwatch_event_bus.Test.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
testPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Sid: pulumi.StringRef("OrganizationAccess"),
Effect: pulumi.StringRef("Allow"),
Actions: []string{
"events:DescribeRule",
"events:ListRules",
"events:ListTargetsByRule",
"events:ListTagsForResource",
},
Resources: []string{
"arn:aws:events:eu-west-1:123456789012:rule/*",
"arn:aws:events:eu-west-1:123456789012:event-bus/default",
},
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: []string{
"*",
},
},
},
Conditions: []iam.GetPolicyDocumentStatementCondition{
{
Test: "StringEquals",
Variable: "aws:PrincipalOrgID",
Values: interface{}{
aws_organizations_organization.Example.Id,
},
},
},
},
},
}, nil);
if err != nil {
return err
}
_, err = cloudwatch.NewEventBusPolicy(ctx, "testEventBusPolicy", &cloudwatch.EventBusPolicyArgs{
Policy: *pulumi.String(testPolicyDocument.Json),
EventBusName: pulumi.Any(aws_cloudwatch_event_bus.Test.Name),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.cloudwatch.EventBusPolicy;
import com.pulumi.aws.cloudwatch.EventBusPolicyArgs;
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
        final var testPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("OrganizationAccess")
                .effect("Allow")
                .actions(                
                    "events:DescribeRule",
                    "events:ListRules",
                    "events:ListTargetsByRule",
                    "events:ListTagsForResource")
                .resources(                
                    "arn:aws:events:eu-west-1:123456789012:rule/*",
                    "arn:aws:events:eu-west-1:123456789012:event-bus/default")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("*")
                    .build())
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("StringEquals")
                    .variable("aws:PrincipalOrgID")
                    .values(aws_organizations_organization.example().id())
                    .build())
                .build())
            .build());

        var testEventBusPolicy = new EventBusPolicy("testEventBusPolicy", EventBusPolicyArgs.builder()        
            .policy(testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .eventBusName(aws_cloudwatch_event_bus.test().name())
            .build());

    }
}
```
```yaml
resources:
  testEventBusPolicy:
    type: aws:cloudwatch:EventBusPolicy
    properties:
      policy: ${testPolicyDocument.json}
      eventBusName: ${aws_cloudwatch_event_bus.test.name}
variables:
  testPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: OrganizationAccess
            effect: Allow
            actions:
              - events:DescribeRule
              - events:ListRules
              - events:ListTargetsByRule
              - events:ListTagsForResource
            resources:
              - arn:aws:events:eu-west-1:123456789012:rule/*
              - arn:aws:events:eu-west-1:123456789012:event-bus/default
            principals:
              - type: AWS
                identifiers:
                  - '*'
            conditions:
              - test: StringEquals
                variable: aws:PrincipalOrgID
                values:
                  - ${aws_organizations_organization.example.id}
```
{{% /example %}}
{{% example %}}
### Multiple Statements

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testPolicyDocument = aws.iam.getPolicyDocument({
    statements: [
        {
            sid: "DevAccountAccess",
            effect: "Allow",
            actions: ["events:PutEvents"],
            resources: ["arn:aws:events:eu-west-1:123456789012:event-bus/default"],
            principals: [{
                type: "AWS",
                identifiers: ["123456789012"],
            }],
        },
        {
            sid: "OrganizationAccess",
            effect: "Allow",
            actions: [
                "events:DescribeRule",
                "events:ListRules",
                "events:ListTargetsByRule",
                "events:ListTagsForResource",
            ],
            resources: [
                "arn:aws:events:eu-west-1:123456789012:rule/*",
                "arn:aws:events:eu-west-1:123456789012:event-bus/default",
            ],
            principals: [{
                type: "AWS",
                identifiers: ["*"],
            }],
            conditions: [{
                test: "StringEquals",
                variable: "aws:PrincipalOrgID",
                values: [aws_organizations_organization.example.id],
            }],
        },
    ],
});
const testEventBusPolicy = new aws.cloudwatch.EventBusPolicy("testEventBusPolicy", {
    policy: testPolicyDocument.then(testPolicyDocument => testPolicyDocument.json),
    eventBusName: aws_cloudwatch_event_bus.test.name,
});
```
```python
import pulumi
import pulumi_aws as aws

test_policy_document = aws.iam.get_policy_document(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="DevAccountAccess",
        effect="Allow",
        actions=["events:PutEvents"],
        resources=["arn:aws:events:eu-west-1:123456789012:event-bus/default"],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=["123456789012"],
        )],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="OrganizationAccess",
        effect="Allow",
        actions=[
            "events:DescribeRule",
            "events:ListRules",
            "events:ListTargetsByRule",
            "events:ListTagsForResource",
        ],
        resources=[
            "arn:aws:events:eu-west-1:123456789012:rule/*",
            "arn:aws:events:eu-west-1:123456789012:event-bus/default",
        ],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=["*"],
        )],
        conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
            test="StringEquals",
            variable="aws:PrincipalOrgID",
            values=[aws_organizations_organization["example"]["id"]],
        )],
    ),
])
test_event_bus_policy = aws.cloudwatch.EventBusPolicy("testEventBusPolicy",
    policy=test_policy_document.json,
    event_bus_name=aws_cloudwatch_event_bus["test"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "DevAccountAccess",
                Effect = "Allow",
                Actions = new[]
                {
                    "events:PutEvents",
                },
                Resources = new[]
                {
                    "arn:aws:events:eu-west-1:123456789012:event-bus/default",
                },
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
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "OrganizationAccess",
                Effect = "Allow",
                Actions = new[]
                {
                    "events:DescribeRule",
                    "events:ListRules",
                    "events:ListTargetsByRule",
                    "events:ListTagsForResource",
                },
                Resources = new[]
                {
                    "arn:aws:events:eu-west-1:123456789012:rule/*",
                    "arn:aws:events:eu-west-1:123456789012:event-bus/default",
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
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "StringEquals",
                        Variable = "aws:PrincipalOrgID",
                        Values = new[]
                        {
                            aws_organizations_organization.Example.Id,
                        },
                    },
                },
            },
        },
    });

    var testEventBusPolicy = new Aws.CloudWatch.EventBusPolicy("testEventBusPolicy", new()
    {
        Policy = testPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
        EventBusName = aws_cloudwatch_event_bus.Test.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
testPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Sid: pulumi.StringRef("DevAccountAccess"),
Effect: pulumi.StringRef("Allow"),
Actions: []string{
"events:PutEvents",
},
Resources: []string{
"arn:aws:events:eu-west-1:123456789012:event-bus/default",
},
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: []string{
"123456789012",
},
},
},
},
{
Sid: pulumi.StringRef("OrganizationAccess"),
Effect: pulumi.StringRef("Allow"),
Actions: []string{
"events:DescribeRule",
"events:ListRules",
"events:ListTargetsByRule",
"events:ListTagsForResource",
},
Resources: []string{
"arn:aws:events:eu-west-1:123456789012:rule/*",
"arn:aws:events:eu-west-1:123456789012:event-bus/default",
},
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: []string{
"*",
},
},
},
Conditions: []iam.GetPolicyDocumentStatementCondition{
{
Test: "StringEquals",
Variable: "aws:PrincipalOrgID",
Values: interface{}{
aws_organizations_organization.Example.Id,
},
},
},
},
},
}, nil);
if err != nil {
return err
}
_, err = cloudwatch.NewEventBusPolicy(ctx, "testEventBusPolicy", &cloudwatch.EventBusPolicyArgs{
Policy: *pulumi.String(testPolicyDocument.Json),
EventBusName: pulumi.Any(aws_cloudwatch_event_bus.Test.Name),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.cloudwatch.EventBusPolicy;
import com.pulumi.aws.cloudwatch.EventBusPolicyArgs;
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
        final var testPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .sid("DevAccountAccess")
                    .effect("Allow")
                    .actions("events:PutEvents")
                    .resources("arn:aws:events:eu-west-1:123456789012:event-bus/default")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("AWS")
                        .identifiers("123456789012")
                        .build())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .sid("OrganizationAccess")
                    .effect("Allow")
                    .actions(                    
                        "events:DescribeRule",
                        "events:ListRules",
                        "events:ListTargetsByRule",
                        "events:ListTagsForResource")
                    .resources(                    
                        "arn:aws:events:eu-west-1:123456789012:rule/*",
                        "arn:aws:events:eu-west-1:123456789012:event-bus/default")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("AWS")
                        .identifiers("*")
                        .build())
                    .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                        .test("StringEquals")
                        .variable("aws:PrincipalOrgID")
                        .values(aws_organizations_organization.example().id())
                        .build())
                    .build())
            .build());

        var testEventBusPolicy = new EventBusPolicy("testEventBusPolicy", EventBusPolicyArgs.builder()        
            .policy(testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .eventBusName(aws_cloudwatch_event_bus.test().name())
            .build());

    }
}
```
```yaml
resources:
  testEventBusPolicy:
    type: aws:cloudwatch:EventBusPolicy
    properties:
      policy: ${testPolicyDocument.json}
      eventBusName: ${aws_cloudwatch_event_bus.test.name}
variables:
  testPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: DevAccountAccess
            effect: Allow
            actions:
              - events:PutEvents
            resources:
              - arn:aws:events:eu-west-1:123456789012:event-bus/default
            principals:
              - type: AWS
                identifiers:
                  - '123456789012'
          - sid: OrganizationAccess
            effect: Allow
            actions:
              - events:DescribeRule
              - events:ListRules
              - events:ListTargetsByRule
              - events:ListTagsForResource
            resources:
              - arn:aws:events:eu-west-1:123456789012:rule/*
              - arn:aws:events:eu-west-1:123456789012:event-bus/default
            principals:
              - type: AWS
                identifiers:
                  - '*'
            conditions:
              - test: StringEquals
                variable: aws:PrincipalOrgID
                values:
                  - ${aws_organizations_organization.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an EventBridge policy using the `event_bus_name`. For example:

```sh
 $ pulumi import aws:cloudwatch/eventBusPolicy:EventBusPolicy DevAccountAccess example-event-bus
```
 