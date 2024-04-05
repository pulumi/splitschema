Manages an Amazon API Gateway Version 2 integration.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Integration("example", {
    apiId: aws_apigatewayv2_api.example.id,
    integrationType: "MOCK",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Integration("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    integration_type="MOCK")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.Integration("example", new()
    {
        ApiId = aws_apigatewayv2_api.Example.Id,
        IntegrationType = "MOCK",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewIntegration(ctx, "example", &apigatewayv2.IntegrationArgs{
			ApiId:           pulumi.Any(aws_apigatewayv2_api.Example.Id),
			IntegrationType: pulumi.String("MOCK"),
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
import com.pulumi.aws.apigatewayv2.Integration;
import com.pulumi.aws.apigatewayv2.IntegrationArgs;
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
        var example = new Integration("example", IntegrationArgs.builder()        
            .apiId(aws_apigatewayv2_api.example().id())
            .integrationType("MOCK")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:Integration
    properties:
      apiId: ${aws_apigatewayv2_api.example.id}
      integrationType: MOCK
```
{{% /example %}}
{{% example %}}
### Lambda Integration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleFunction = new aws.lambda.Function("exampleFunction", {
    code: new pulumi.asset.FileArchive("example.zip"),
    role: aws_iam_role.example.arn,
    handler: "index.handler",
    runtime: "nodejs16.x",
});
const exampleIntegration = new aws.apigatewayv2.Integration("exampleIntegration", {
    apiId: aws_apigatewayv2_api.example.id,
    integrationType: "AWS_PROXY",
    connectionType: "INTERNET",
    contentHandlingStrategy: "CONVERT_TO_TEXT",
    description: "Lambda example",
    integrationMethod: "POST",
    integrationUri: exampleFunction.invokeArn,
    passthroughBehavior: "WHEN_NO_MATCH",
});
```
```python
import pulumi
import pulumi_aws as aws

example_function = aws.lambda_.Function("exampleFunction",
    code=pulumi.FileArchive("example.zip"),
    role=aws_iam_role["example"]["arn"],
    handler="index.handler",
    runtime="nodejs16.x")
example_integration = aws.apigatewayv2.Integration("exampleIntegration",
    api_id=aws_apigatewayv2_api["example"]["id"],
    integration_type="AWS_PROXY",
    connection_type="INTERNET",
    content_handling_strategy="CONVERT_TO_TEXT",
    description="Lambda example",
    integration_method="POST",
    integration_uri=example_function.invoke_arn,
    passthrough_behavior="WHEN_NO_MATCH")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleFunction = new Aws.Lambda.Function("exampleFunction", new()
    {
        Code = new FileArchive("example.zip"),
        Role = aws_iam_role.Example.Arn,
        Handler = "index.handler",
        Runtime = "nodejs16.x",
    });

    var exampleIntegration = new Aws.ApiGatewayV2.Integration("exampleIntegration", new()
    {
        ApiId = aws_apigatewayv2_api.Example.Id,
        IntegrationType = "AWS_PROXY",
        ConnectionType = "INTERNET",
        ContentHandlingStrategy = "CONVERT_TO_TEXT",
        Description = "Lambda example",
        IntegrationMethod = "POST",
        IntegrationUri = exampleFunction.InvokeArn,
        PassthroughBehavior = "WHEN_NO_MATCH",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleFunction, err := lambda.NewFunction(ctx, "exampleFunction", &lambda.FunctionArgs{
			Code:    pulumi.NewFileArchive("example.zip"),
			Role:    pulumi.Any(aws_iam_role.Example.Arn),
			Handler: pulumi.String("index.handler"),
			Runtime: pulumi.String("nodejs16.x"),
		})
		if err != nil {
			return err
		}
		_, err = apigatewayv2.NewIntegration(ctx, "exampleIntegration", &apigatewayv2.IntegrationArgs{
			ApiId:                   pulumi.Any(aws_apigatewayv2_api.Example.Id),
			IntegrationType:         pulumi.String("AWS_PROXY"),
			ConnectionType:          pulumi.String("INTERNET"),
			ContentHandlingStrategy: pulumi.String("CONVERT_TO_TEXT"),
			Description:             pulumi.String("Lambda example"),
			IntegrationMethod:       pulumi.String("POST"),
			IntegrationUri:          exampleFunction.InvokeArn,
			PassthroughBehavior:     pulumi.String("WHEN_NO_MATCH"),
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
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
import com.pulumi.aws.apigatewayv2.Integration;
import com.pulumi.aws.apigatewayv2.IntegrationArgs;
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
        var exampleFunction = new Function("exampleFunction", FunctionArgs.builder()        
            .code(new FileArchive("example.zip"))
            .role(aws_iam_role.example().arn())
            .handler("index.handler")
            .runtime("nodejs16.x")
            .build());

        var exampleIntegration = new Integration("exampleIntegration", IntegrationArgs.builder()        
            .apiId(aws_apigatewayv2_api.example().id())
            .integrationType("AWS_PROXY")
            .connectionType("INTERNET")
            .contentHandlingStrategy("CONVERT_TO_TEXT")
            .description("Lambda example")
            .integrationMethod("POST")
            .integrationUri(exampleFunction.invokeArn())
            .passthroughBehavior("WHEN_NO_MATCH")
            .build());

    }
}
```
```yaml
resources:
  exampleFunction:
    type: aws:lambda:Function
    properties:
      code:
        fn::FileArchive: example.zip
      role: ${aws_iam_role.example.arn}
      handler: index.handler
      runtime: nodejs16.x
  exampleIntegration:
    type: aws:apigatewayv2:Integration
    properties:
      apiId: ${aws_apigatewayv2_api.example.id}
      integrationType: AWS_PROXY
      connectionType: INTERNET
      contentHandlingStrategy: CONVERT_TO_TEXT
      description: Lambda example
      integrationMethod: POST
      integrationUri: ${exampleFunction.invokeArn}
      passthroughBehavior: WHEN_NO_MATCH
```
{{% /example %}}
{{% example %}}
### AWS Service Integration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Integration("example", {
    apiId: aws_apigatewayv2_api.example.id,
    credentialsArn: aws_iam_role.example.arn,
    description: "SQS example",
    integrationType: "AWS_PROXY",
    integrationSubtype: "SQS-SendMessage",
    requestParameters: {
        QueueUrl: "$request.header.queueUrl",
        MessageBody: "$request.body.message",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Integration("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    credentials_arn=aws_iam_role["example"]["arn"],
    description="SQS example",
    integration_type="AWS_PROXY",
    integration_subtype="SQS-SendMessage",
    request_parameters={
        "QueueUrl": "$request.header.queueUrl",
        "MessageBody": "$request.body.message",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.Integration("example", new()
    {
        ApiId = aws_apigatewayv2_api.Example.Id,
        CredentialsArn = aws_iam_role.Example.Arn,
        Description = "SQS example",
        IntegrationType = "AWS_PROXY",
        IntegrationSubtype = "SQS-SendMessage",
        RequestParameters = 
        {
            { "QueueUrl", "$request.header.queueUrl" },
            { "MessageBody", "$request.body.message" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewIntegration(ctx, "example", &apigatewayv2.IntegrationArgs{
			ApiId:              pulumi.Any(aws_apigatewayv2_api.Example.Id),
			CredentialsArn:     pulumi.Any(aws_iam_role.Example.Arn),
			Description:        pulumi.String("SQS example"),
			IntegrationType:    pulumi.String("AWS_PROXY"),
			IntegrationSubtype: pulumi.String("SQS-SendMessage"),
			RequestParameters: pulumi.StringMap{
				"QueueUrl":    pulumi.String("$request.header.queueUrl"),
				"MessageBody": pulumi.String("$request.body.message"),
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
import com.pulumi.aws.apigatewayv2.Integration;
import com.pulumi.aws.apigatewayv2.IntegrationArgs;
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
        var example = new Integration("example", IntegrationArgs.builder()        
            .apiId(aws_apigatewayv2_api.example().id())
            .credentialsArn(aws_iam_role.example().arn())
            .description("SQS example")
            .integrationType("AWS_PROXY")
            .integrationSubtype("SQS-SendMessage")
            .requestParameters(Map.ofEntries(
                Map.entry("QueueUrl", "$request.header.queueUrl"),
                Map.entry("MessageBody", "$request.body.message")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:Integration
    properties:
      apiId: ${aws_apigatewayv2_api.example.id}
      credentialsArn: ${aws_iam_role.example.arn}
      description: SQS example
      integrationType: AWS_PROXY
      integrationSubtype: SQS-SendMessage
      requestParameters:
        QueueUrl: $request.header.queueUrl
        MessageBody: $request.body.message
```
{{% /example %}}
{{% example %}}
### Private Integration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Integration("example", {
    apiId: aws_apigatewayv2_api.example.id,
    credentialsArn: aws_iam_role.example.arn,
    description: "Example with a load balancer",
    integrationType: "HTTP_PROXY",
    integrationUri: aws_lb_listener.example.arn,
    integrationMethod: "ANY",
    connectionType: "VPC_LINK",
    connectionId: aws_apigatewayv2_vpc_link.example.id,
    tlsConfig: {
        serverNameToVerify: "example.com",
    },
    requestParameters: {
        "append:header.authforintegration": "$context.authorizer.authorizerResponse",
        "overwrite:path": "staticValueForIntegration",
    },
    responseParameters: [
        {
            statusCode: "403",
            mappings: {
                "append:header.auth": "$context.authorizer.authorizerResponse",
            },
        },
        {
            statusCode: "200",
            mappings: {
                "overwrite:statuscode": "204",
            },
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Integration("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    credentials_arn=aws_iam_role["example"]["arn"],
    description="Example with a load balancer",
    integration_type="HTTP_PROXY",
    integration_uri=aws_lb_listener["example"]["arn"],
    integration_method="ANY",
    connection_type="VPC_LINK",
    connection_id=aws_apigatewayv2_vpc_link["example"]["id"],
    tls_config=aws.apigatewayv2.IntegrationTlsConfigArgs(
        server_name_to_verify="example.com",
    ),
    request_parameters={
        "append:header.authforintegration": "$context.authorizer.authorizerResponse",
        "overwrite:path": "staticValueForIntegration",
    },
    response_parameters=[
        aws.apigatewayv2.IntegrationResponseParameterArgs(
            status_code="403",
            mappings={
                "append:header.auth": "$context.authorizer.authorizerResponse",
            },
        ),
        aws.apigatewayv2.IntegrationResponseParameterArgs(
            status_code="200",
            mappings={
                "overwrite:statuscode": "204",
            },
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.Integration("example", new()
    {
        ApiId = aws_apigatewayv2_api.Example.Id,
        CredentialsArn = aws_iam_role.Example.Arn,
        Description = "Example with a load balancer",
        IntegrationType = "HTTP_PROXY",
        IntegrationUri = aws_lb_listener.Example.Arn,
        IntegrationMethod = "ANY",
        ConnectionType = "VPC_LINK",
        ConnectionId = aws_apigatewayv2_vpc_link.Example.Id,
        TlsConfig = new Aws.ApiGatewayV2.Inputs.IntegrationTlsConfigArgs
        {
            ServerNameToVerify = "example.com",
        },
        RequestParameters = 
        {
            { "append:header.authforintegration", "$context.authorizer.authorizerResponse" },
            { "overwrite:path", "staticValueForIntegration" },
        },
        ResponseParameters = new[]
        {
            new Aws.ApiGatewayV2.Inputs.IntegrationResponseParameterArgs
            {
                StatusCode = "403",
                Mappings = 
                {
                    { "append:header.auth", "$context.authorizer.authorizerResponse" },
                },
            },
            new Aws.ApiGatewayV2.Inputs.IntegrationResponseParameterArgs
            {
                StatusCode = "200",
                Mappings = 
                {
                    { "overwrite:statuscode", "204" },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewIntegration(ctx, "example", &apigatewayv2.IntegrationArgs{
			ApiId:             pulumi.Any(aws_apigatewayv2_api.Example.Id),
			CredentialsArn:    pulumi.Any(aws_iam_role.Example.Arn),
			Description:       pulumi.String("Example with a load balancer"),
			IntegrationType:   pulumi.String("HTTP_PROXY"),
			IntegrationUri:    pulumi.Any(aws_lb_listener.Example.Arn),
			IntegrationMethod: pulumi.String("ANY"),
			ConnectionType:    pulumi.String("VPC_LINK"),
			ConnectionId:      pulumi.Any(aws_apigatewayv2_vpc_link.Example.Id),
			TlsConfig: &apigatewayv2.IntegrationTlsConfigArgs{
				ServerNameToVerify: pulumi.String("example.com"),
			},
			RequestParameters: pulumi.StringMap{
				"append:header.authforintegration": pulumi.String("$context.authorizer.authorizerResponse"),
				"overwrite:path":                   pulumi.String("staticValueForIntegration"),
			},
			ResponseParameters: apigatewayv2.IntegrationResponseParameterArray{
				&apigatewayv2.IntegrationResponseParameterArgs{
					StatusCode: pulumi.String("403"),
					Mappings: pulumi.StringMap{
						"append:header.auth": pulumi.String("$context.authorizer.authorizerResponse"),
					},
				},
				&apigatewayv2.IntegrationResponseParameterArgs{
					StatusCode: pulumi.String("200"),
					Mappings: pulumi.StringMap{
						"overwrite:statuscode": pulumi.String("204"),
					},
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
import com.pulumi.aws.apigatewayv2.Integration;
import com.pulumi.aws.apigatewayv2.IntegrationArgs;
import com.pulumi.aws.apigatewayv2.inputs.IntegrationTlsConfigArgs;
import com.pulumi.aws.apigatewayv2.inputs.IntegrationResponseParameterArgs;
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
        var example = new Integration("example", IntegrationArgs.builder()        
            .apiId(aws_apigatewayv2_api.example().id())
            .credentialsArn(aws_iam_role.example().arn())
            .description("Example with a load balancer")
            .integrationType("HTTP_PROXY")
            .integrationUri(aws_lb_listener.example().arn())
            .integrationMethod("ANY")
            .connectionType("VPC_LINK")
            .connectionId(aws_apigatewayv2_vpc_link.example().id())
            .tlsConfig(IntegrationTlsConfigArgs.builder()
                .serverNameToVerify("example.com")
                .build())
            .requestParameters(Map.ofEntries(
                Map.entry("append:header.authforintegration", "$context.authorizer.authorizerResponse"),
                Map.entry("overwrite:path", "staticValueForIntegration")
            ))
            .responseParameters(            
                IntegrationResponseParameterArgs.builder()
                    .statusCode(403)
                    .mappings(Map.of("append:header.auth", "$context.authorizer.authorizerResponse"))
                    .build(),
                IntegrationResponseParameterArgs.builder()
                    .statusCode(200)
                    .mappings(Map.of("overwrite:statuscode", "204"))
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:Integration
    properties:
      apiId: ${aws_apigatewayv2_api.example.id}
      credentialsArn: ${aws_iam_role.example.arn}
      description: Example with a load balancer
      integrationType: HTTP_PROXY
      integrationUri: ${aws_lb_listener.example.arn}
      integrationMethod: ANY
      connectionType: VPC_LINK
      connectionId: ${aws_apigatewayv2_vpc_link.example.id}
      tlsConfig:
        serverNameToVerify: example.com
      requestParameters:
        append:header.authforintegration: $context.authorizer.authorizerResponse
        overwrite:path: staticValueForIntegration
      responseParameters:
        - statusCode: 403
          mappings:
            append:header.auth: $context.authorizer.authorizerResponse
        - statusCode: 200
          mappings:
            overwrite:statuscode: '204'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_apigatewayv2_integration` using the API identifier and integration identifier. For example:

```sh
 $ pulumi import aws:apigatewayv2/integration:Integration example aabbccddee/1122334
```
 -> __Note:__ The API Gateway managed integration created as part of [_quick_create_](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-basic-concept.html#apigateway-definition-quick-create) cannot be imported.

