Provides a settings of an API Gateway Account. Settings is applied region-wide per `provider` block.

> **Note:** As there is no API method for deleting account settings or resetting it to defaults, destroying this resource will keep your account settings intact

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["apigateway.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const cloudwatchRole = new aws.iam.Role("cloudwatchRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const demo = new aws.apigateway.Account("demo", {cloudwatchRoleArn: cloudwatchRole.arn});
const cloudwatchPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: [
            "logs:CreateLogGroup",
            "logs:CreateLogStream",
            "logs:DescribeLogGroups",
            "logs:DescribeLogStreams",
            "logs:PutLogEvents",
            "logs:GetLogEvents",
            "logs:FilterLogEvents",
        ],
        resources: ["*"],
    }],
});
const cloudwatchRolePolicy = new aws.iam.RolePolicy("cloudwatchRolePolicy", {
    role: cloudwatchRole.id,
    policy: cloudwatchPolicyDocument.then(cloudwatchPolicyDocument => cloudwatchPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["apigateway.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
cloudwatch_role = aws.iam.Role("cloudwatchRole", assume_role_policy=assume_role.json)
demo = aws.apigateway.Account("demo", cloudwatch_role_arn=cloudwatch_role.arn)
cloudwatch_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=[
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:DescribeLogGroups",
        "logs:DescribeLogStreams",
        "logs:PutLogEvents",
        "logs:GetLogEvents",
        "logs:FilterLogEvents",
    ],
    resources=["*"],
)])
cloudwatch_role_policy = aws.iam.RolePolicy("cloudwatchRolePolicy",
    role=cloudwatch_role.id,
    policy=cloudwatch_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
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
                            "apigateway.amazonaws.com",
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

    var cloudwatchRole = new Aws.Iam.Role("cloudwatchRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var demo = new Aws.ApiGateway.Account("demo", new()
    {
        CloudwatchRoleArn = cloudwatchRole.Arn,
    });

    var cloudwatchPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:DescribeLogGroups",
                    "logs:DescribeLogStreams",
                    "logs:PutLogEvents",
                    "logs:GetLogEvents",
                    "logs:FilterLogEvents",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var cloudwatchRolePolicy = new Aws.Iam.RolePolicy("cloudwatchRolePolicy", new()
    {
        Role = cloudwatchRole.Id,
        Policy = cloudwatchPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"apigateway.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		cloudwatchRole, err := iam.NewRole(ctx, "cloudwatchRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewAccount(ctx, "demo", &apigateway.AccountArgs{
			CloudwatchRoleArn: cloudwatchRole.Arn,
		})
		if err != nil {
			return err
		}
		cloudwatchPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"logs:CreateLogGroup",
						"logs:CreateLogStream",
						"logs:DescribeLogGroups",
						"logs:DescribeLogStreams",
						"logs:PutLogEvents",
						"logs:GetLogEvents",
						"logs:FilterLogEvents",
					},
					Resources: []string{
						"*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "cloudwatchRolePolicy", &iam.RolePolicyArgs{
			Role:   cloudwatchRole.ID(),
			Policy: *pulumi.String(cloudwatchPolicyDocument.Json),
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.apigateway.Account;
import com.pulumi.aws.apigateway.AccountArgs;
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
        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("apigateway.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var cloudwatchRole = new Role("cloudwatchRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var demo = new Account("demo", AccountArgs.builder()        
            .cloudwatchRoleArn(cloudwatchRole.arn())
            .build());

        final var cloudwatchPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions(                
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:DescribeLogGroups",
                    "logs:DescribeLogStreams",
                    "logs:PutLogEvents",
                    "logs:GetLogEvents",
                    "logs:FilterLogEvents")
                .resources("*")
                .build())
            .build());

        var cloudwatchRolePolicy = new RolePolicy("cloudwatchRolePolicy", RolePolicyArgs.builder()        
            .role(cloudwatchRole.id())
            .policy(cloudwatchPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  demo:
    type: aws:apigateway:Account
    properties:
      cloudwatchRoleArn: ${cloudwatchRole.arn}
  cloudwatchRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  cloudwatchRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${cloudwatchRole.id}
      policy: ${cloudwatchPolicyDocument.json}
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
                  - apigateway.amazonaws.com
            actions:
              - sts:AssumeRole
  cloudwatchPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - logs:CreateLogGroup
              - logs:CreateLogStream
              - logs:DescribeLogGroups
              - logs:DescribeLogStreams
              - logs:PutLogEvents
              - logs:GetLogEvents
              - logs:FilterLogEvents
            resources:
              - '*'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import API Gateway Accounts using the word `api-gateway-account`. For example:

```sh
 $ pulumi import aws:apigateway/account:Account demo api-gateway-account
```
 