Manages an Amazon API Gateway Version 2 authorizer.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic WebSocket API

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Authorizer("example", {
    apiId: aws_apigatewayv2_api.example.id,
    authorizerType: "REQUEST",
    authorizerUri: aws_lambda_function.example.invoke_arn,
    identitySources: ["route.request.header.Auth"],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Authorizer("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    authorizer_type="REQUEST",
    authorizer_uri=aws_lambda_function["example"]["invoke_arn"],
    identity_sources=["route.request.header.Auth"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.Authorizer("example", new()
    {
        ApiId = aws_apigatewayv2_api.Example.Id,
        AuthorizerType = "REQUEST",
        AuthorizerUri = aws_lambda_function.Example.Invoke_arn,
        IdentitySources = new[]
        {
            "route.request.header.Auth",
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
		_, err := apigatewayv2.NewAuthorizer(ctx, "example", &apigatewayv2.AuthorizerArgs{
			ApiId:          pulumi.Any(aws_apigatewayv2_api.Example.Id),
			AuthorizerType: pulumi.String("REQUEST"),
			AuthorizerUri:  pulumi.Any(aws_lambda_function.Example.Invoke_arn),
			IdentitySources: pulumi.StringArray{
				pulumi.String("route.request.header.Auth"),
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
import com.pulumi.aws.apigatewayv2.Authorizer;
import com.pulumi.aws.apigatewayv2.AuthorizerArgs;
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
        var example = new Authorizer("example", AuthorizerArgs.builder()        
            .apiId(aws_apigatewayv2_api.example().id())
            .authorizerType("REQUEST")
            .authorizerUri(aws_lambda_function.example().invoke_arn())
            .identitySources("route.request.header.Auth")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:Authorizer
    properties:
      apiId: ${aws_apigatewayv2_api.example.id}
      authorizerType: REQUEST
      authorizerUri: ${aws_lambda_function.example.invoke_arn}
      identitySources:
        - route.request.header.Auth
```
{{% /example %}}
{{% example %}}
### Basic HTTP API

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Authorizer("example", {
    apiId: aws_apigatewayv2_api.example.id,
    authorizerType: "REQUEST",
    authorizerUri: aws_lambda_function.example.invoke_arn,
    identitySources: ["$request.header.Authorization"],
    authorizerPayloadFormatVersion: "2.0",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Authorizer("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    authorizer_type="REQUEST",
    authorizer_uri=aws_lambda_function["example"]["invoke_arn"],
    identity_sources=["$request.header.Authorization"],
    authorizer_payload_format_version="2.0")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.Authorizer("example", new()
    {
        ApiId = aws_apigatewayv2_api.Example.Id,
        AuthorizerType = "REQUEST",
        AuthorizerUri = aws_lambda_function.Example.Invoke_arn,
        IdentitySources = new[]
        {
            "$request.header.Authorization",
        },
        AuthorizerPayloadFormatVersion = "2.0",
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
		_, err := apigatewayv2.NewAuthorizer(ctx, "example", &apigatewayv2.AuthorizerArgs{
			ApiId:          pulumi.Any(aws_apigatewayv2_api.Example.Id),
			AuthorizerType: pulumi.String("REQUEST"),
			AuthorizerUri:  pulumi.Any(aws_lambda_function.Example.Invoke_arn),
			IdentitySources: pulumi.StringArray{
				pulumi.String("$request.header.Authorization"),
			},
			AuthorizerPayloadFormatVersion: pulumi.String("2.0"),
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
import com.pulumi.aws.apigatewayv2.Authorizer;
import com.pulumi.aws.apigatewayv2.AuthorizerArgs;
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
        var example = new Authorizer("example", AuthorizerArgs.builder()        
            .apiId(aws_apigatewayv2_api.example().id())
            .authorizerType("REQUEST")
            .authorizerUri(aws_lambda_function.example().invoke_arn())
            .identitySources("$request.header.Authorization")
            .authorizerPayloadFormatVersion("2.0")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:Authorizer
    properties:
      apiId: ${aws_apigatewayv2_api.example.id}
      authorizerType: REQUEST
      authorizerUri: ${aws_lambda_function.example.invoke_arn}
      identitySources:
        - $request.header.Authorization
      authorizerPayloadFormatVersion: '2.0'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_apigatewayv2_authorizer` using the API identifier and authorizer identifier. For example:

```sh
 $ pulumi import aws:apigatewayv2/authorizer:Authorizer example aabbccddee/1122334
```
 