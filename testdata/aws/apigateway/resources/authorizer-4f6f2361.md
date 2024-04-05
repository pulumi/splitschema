Provides an API Gateway Authorizer.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const demoRestApi = new aws.apigateway.RestApi("demoRestApi", {});
const invocationAssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["apigateway.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const invocationRole = new aws.iam.Role("invocationRole", {
    path: "/",
    assumeRolePolicy: invocationAssumeRole.then(invocationAssumeRole => invocationAssumeRole.json),
});
const lambdaAssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["lambda.amazonaws.com"],
        }],
    }],
});
const lambda = new aws.iam.Role("lambda", {assumeRolePolicy: lambdaAssumeRole.then(lambdaAssumeRole => lambdaAssumeRole.json)});
const authorizer = new aws.lambda.Function("authorizer", {
    code: new pulumi.asset.FileArchive("lambda-function.zip"),
    role: lambda.arn,
    handler: "exports.example",
});
const demoAuthorizer = new aws.apigateway.Authorizer("demoAuthorizer", {
    restApi: demoRestApi.id,
    authorizerUri: authorizer.invokeArn,
    authorizerCredentials: invocationRole.arn,
});
const invocationPolicyPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        actions: ["lambda:InvokeFunction"],
        resources: [authorizer.arn],
    }],
});
const invocationPolicyRolePolicy = new aws.iam.RolePolicy("invocationPolicyRolePolicy", {
    role: invocationRole.id,
    policy: invocationPolicyPolicyDocument.apply(invocationPolicyPolicyDocument => invocationPolicyPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

demo_rest_api = aws.apigateway.RestApi("demoRestApi")
invocation_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["apigateway.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
invocation_role = aws.iam.Role("invocationRole",
    path="/",
    assume_role_policy=invocation_assume_role.json)
lambda_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["lambda.amazonaws.com"],
    )],
)])
lambda_ = aws.iam.Role("lambda", assume_role_policy=lambda_assume_role.json)
authorizer = aws.lambda_.Function("authorizer",
    code=pulumi.FileArchive("lambda-function.zip"),
    role=lambda_.arn,
    handler="exports.example")
demo_authorizer = aws.apigateway.Authorizer("demoAuthorizer",
    rest_api=demo_rest_api.id,
    authorizer_uri=authorizer.invoke_arn,
    authorizer_credentials=invocation_role.arn)
invocation_policy_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["lambda:InvokeFunction"],
    resources=[authorizer.arn],
)])
invocation_policy_role_policy = aws.iam.RolePolicy("invocationPolicyRolePolicy",
    role=invocation_role.id,
    policy=invocation_policy_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var demoRestApi = new Aws.ApiGateway.RestApi("demoRestApi");

    var invocationAssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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

    var invocationRole = new Aws.Iam.Role("invocationRole", new()
    {
        Path = "/",
        AssumeRolePolicy = invocationAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var lambdaAssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "lambda.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var lambda = new Aws.Iam.Role("lambda", new()
    {
        AssumeRolePolicy = lambdaAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var authorizer = new Aws.Lambda.Function("authorizer", new()
    {
        Code = new FileArchive("lambda-function.zip"),
        Role = lambda.Arn,
        Handler = "exports.example",
    });

    var demoAuthorizer = new Aws.ApiGateway.Authorizer("demoAuthorizer", new()
    {
        RestApi = demoRestApi.Id,
        AuthorizerUri = authorizer.InvokeArn,
        AuthorizerCredentials = invocationRole.Arn,
    });

    var invocationPolicyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "lambda:InvokeFunction",
                },
                Resources = new[]
                {
                    authorizer.Arn,
                },
            },
        },
    });

    var invocationPolicyRolePolicy = new Aws.Iam.RolePolicy("invocationPolicyRolePolicy", new()
    {
        Role = invocationRole.Id,
        Policy = invocationPolicyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		demoRestApi, err := apigateway.NewRestApi(ctx, "demoRestApi", nil)
		if err != nil {
			return err
		}
		invocationAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
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
		invocationRole, err := iam.NewRole(ctx, "invocationRole", &iam.RoleArgs{
			Path:             pulumi.String("/"),
			AssumeRolePolicy: *pulumi.String(invocationAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		lambdaAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"lambda.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		lambda, err := iam.NewRole(ctx, "lambda", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(lambdaAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		authorizer, err := lambda.NewFunction(ctx, "authorizer", &lambda.FunctionArgs{
			Code:    pulumi.NewFileArchive("lambda-function.zip"),
			Role:    lambda.Arn,
			Handler: pulumi.String("exports.example"),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewAuthorizer(ctx, "demoAuthorizer", &apigateway.AuthorizerArgs{
			RestApi:               demoRestApi.ID(),
			AuthorizerUri:         authorizer.InvokeArn,
			AuthorizerCredentials: invocationRole.Arn,
		})
		if err != nil {
			return err
		}
		invocationPolicyPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("lambda:InvokeFunction"),
					},
					Resources: pulumi.StringArray{
						authorizer.Arn,
					},
				},
			},
		}, nil)
		_, err = iam.NewRolePolicy(ctx, "invocationPolicyRolePolicy", &iam.RolePolicyArgs{
			Role: invocationRole.ID(),
			Policy: invocationPolicyPolicyDocument.ApplyT(func(invocationPolicyPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &invocationPolicyPolicyDocument.Json, nil
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
import com.pulumi.aws.apigateway.RestApi;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
import com.pulumi.aws.apigateway.Authorizer;
import com.pulumi.aws.apigateway.AuthorizerArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import com.pulumi.asset.FileArchive;
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
        var demoRestApi = new RestApi("demoRestApi");

        final var invocationAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("apigateway.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var invocationRole = new Role("invocationRole", RoleArgs.builder()        
            .path("/")
            .assumeRolePolicy(invocationAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var lambdaAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("lambda.amazonaws.com")
                    .build())
                .build())
            .build());

        var lambda = new Role("lambda", RoleArgs.builder()        
            .assumeRolePolicy(lambdaAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var authorizer = new Function("authorizer", FunctionArgs.builder()        
            .code(new FileArchive("lambda-function.zip"))
            .role(lambda.arn())
            .handler("exports.example")
            .build());

        var demoAuthorizer = new Authorizer("demoAuthorizer", AuthorizerArgs.builder()        
            .restApi(demoRestApi.id())
            .authorizerUri(authorizer.invokeArn())
            .authorizerCredentials(invocationRole.arn())
            .build());

        final var invocationPolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("lambda:InvokeFunction")
                .resources(authorizer.arn())
                .build())
            .build());

        var invocationPolicyRolePolicy = new RolePolicy("invocationPolicyRolePolicy", RolePolicyArgs.builder()        
            .role(invocationRole.id())
            .policy(invocationPolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(invocationPolicyPolicyDocument -> invocationPolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  demoAuthorizer:
    type: aws:apigateway:Authorizer
    properties:
      restApi: ${demoRestApi.id}
      authorizerUri: ${authorizer.invokeArn}
      authorizerCredentials: ${invocationRole.arn}
  demoRestApi:
    type: aws:apigateway:RestApi
  invocationRole:
    type: aws:iam:Role
    properties:
      path: /
      assumeRolePolicy: ${invocationAssumeRole.json}
  invocationPolicyRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${invocationRole.id}
      policy: ${invocationPolicyPolicyDocument.json}
  lambda:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${lambdaAssumeRole.json}
  authorizer:
    type: aws:lambda:Function
    properties:
      code:
        fn::FileArchive: lambda-function.zip
      role: ${lambda.arn}
      handler: exports.example
variables:
  invocationAssumeRole:
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
  invocationPolicyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - lambda:InvokeFunction
            resources:
              - ${authorizer.arn}
  lambdaAssumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - lambda.amazonaws.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS API Gateway Authorizer using the `REST-API-ID/AUTHORIZER-ID`. For example:

```sh
 $ pulumi import aws:apigateway/authorizer:Authorizer authorizer 12345abcde/example
```
 