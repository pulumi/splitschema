Provides an API Gateway REST API Policy.

> **Note:** Amazon API Gateway Version 1 resources are used for creating and deploying REST APIs. To create and deploy WebSocket and HTTP APIs, use Amazon API Gateway Version 2 resources.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testRestApi = new aws.apigateway.RestApi("testRestApi", {});
const testPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: ["*"],
        }],
        actions: ["execute-api:Invoke"],
        resources: [testRestApi.executionArn],
        conditions: [{
            test: "IpAddress",
            variable: "aws:SourceIp",
            values: ["123.123.123.123/32"],
        }],
    }],
});
const testRestApiPolicy = new aws.apigateway.RestApiPolicy("testRestApiPolicy", {
    restApiId: testRestApi.id,
    policy: testPolicyDocument.apply(testPolicyDocument => testPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

test_rest_api = aws.apigateway.RestApi("testRestApi")
test_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["*"],
    )],
    actions=["execute-api:Invoke"],
    resources=[test_rest_api.execution_arn],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="IpAddress",
        variable="aws:SourceIp",
        values=["123.123.123.123/32"],
    )],
)])
test_rest_api_policy = aws.apigateway.RestApiPolicy("testRestApiPolicy",
    rest_api_id=test_rest_api.id,
    policy=test_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testRestApi = new Aws.ApiGateway.RestApi("testRestApi");

    var testPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            "*",
                        },
                    },
                },
                Actions = new[]
                {
                    "execute-api:Invoke",
                },
                Resources = new[]
                {
                    testRestApi.ExecutionArn,
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "IpAddress",
                        Variable = "aws:SourceIp",
                        Values = new[]
                        {
                            "123.123.123.123/32",
                        },
                    },
                },
            },
        },
    });

    var testRestApiPolicy = new Aws.ApiGateway.RestApiPolicy("testRestApiPolicy", new()
    {
        RestApiId = testRestApi.Id,
        Policy = testPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
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
		testRestApi, err := apigateway.NewRestApi(ctx, "testRestApi", nil)
		if err != nil {
			return err
		}
		testPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("AWS"),
							Identifiers: pulumi.StringArray{
								pulumi.String("*"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("execute-api:Invoke"),
					},
					Resources: pulumi.StringArray{
						testRestApi.ExecutionArn,
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("IpAddress"),
							Variable: pulumi.String("aws:SourceIp"),
							Values: pulumi.StringArray{
								pulumi.String("123.123.123.123/32"),
							},
						},
					},
				},
			},
		}, nil)
		_, err = apigateway.NewRestApiPolicy(ctx, "testRestApiPolicy", &apigateway.RestApiPolicyArgs{
			RestApiId: testRestApi.ID(),
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
import com.pulumi.aws.apigateway.RestApi;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.apigateway.RestApiPolicy;
import com.pulumi.aws.apigateway.RestApiPolicyArgs;
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
        var testRestApi = new RestApi("testRestApi");

        final var testPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("*")
                    .build())
                .actions("execute-api:Invoke")
                .resources(testRestApi.executionArn())
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("IpAddress")
                    .variable("aws:SourceIp")
                    .values("123.123.123.123/32")
                    .build())
                .build())
            .build());

        var testRestApiPolicy = new RestApiPolicy("testRestApiPolicy", RestApiPolicyArgs.builder()        
            .restApiId(testRestApi.id())
            .policy(testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(testPolicyDocument -> testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  testRestApi:
    type: aws:apigateway:RestApi
  testRestApiPolicy:
    type: aws:apigateway:RestApiPolicy
    properties:
      restApiId: ${testRestApi.id}
      policy: ${testPolicyDocument.json}
variables:
  testPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - '*'
            actions:
              - execute-api:Invoke
            resources:
              - ${testRestApi.executionArn}
            conditions:
              - test: IpAddress
                variable: aws:SourceIp
                values:
                  - 123.123.123.123/32
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_api_gateway_rest_api_policy` using the REST API ID. For example:

```sh
 $ pulumi import aws:apigateway/restApiPolicy:RestApiPolicy example 12345abcde
```
 