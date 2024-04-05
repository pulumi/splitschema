Manages an Amazon API Gateway Version 2 API.

> **Note:** Amazon API Gateway Version 2 resources are used for creating and deploying WebSocket and HTTP APIs. To create and deploy REST APIs, use Amazon API Gateway Version 1 resources.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic WebSocket API

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Api("example", {
    protocolType: "WEBSOCKET",
    routeSelectionExpression: "$request.body.action",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Api("example",
    protocol_type="WEBSOCKET",
    route_selection_expression="$request.body.action")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.Api("example", new()
    {
        ProtocolType = "WEBSOCKET",
        RouteSelectionExpression = "$request.body.action",
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
		_, err := apigatewayv2.NewApi(ctx, "example", &apigatewayv2.ApiArgs{
			ProtocolType:             pulumi.String("WEBSOCKET"),
			RouteSelectionExpression: pulumi.String("$request.body.action"),
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
import com.pulumi.aws.apigatewayv2.Api;
import com.pulumi.aws.apigatewayv2.ApiArgs;
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
        var example = new Api("example", ApiArgs.builder()        
            .protocolType("WEBSOCKET")
            .routeSelectionExpression("$request.body.action")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:Api
    properties:
      protocolType: WEBSOCKET
      routeSelectionExpression: $request.body.action
```
{{% /example %}}
{{% example %}}
### Basic HTTP API

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Api("example", {protocolType: "HTTP"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Api("example", protocol_type="HTTP")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.Api("example", new()
    {
        ProtocolType = "HTTP",
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
		_, err := apigatewayv2.NewApi(ctx, "example", &apigatewayv2.ApiArgs{
			ProtocolType: pulumi.String("HTTP"),
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
import com.pulumi.aws.apigatewayv2.Api;
import com.pulumi.aws.apigatewayv2.ApiArgs;
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
        var example = new Api("example", ApiArgs.builder()        
            .protocolType("HTTP")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:Api
    properties:
      protocolType: HTTP
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_apigatewayv2_api` using the API identifier. For example:

```sh
 $ pulumi import aws:apigatewayv2/api:Api example aabbccddee
```
 