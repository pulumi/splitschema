Manages an Amazon API Gateway Version 2 route.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html) for [WebSocket](https://docs.aws.amazon.com/apigateway/latest/developerguide/websocket-api-develop-routes.html) and [HTTP](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html) APIs.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleApi = new aws.apigatewayv2.Api("exampleApi", {
    protocolType: "WEBSOCKET",
    routeSelectionExpression: "$request.body.action",
});
const exampleRoute = new aws.apigatewayv2.Route("exampleRoute", {
    apiId: exampleApi.id,
    routeKey: "$default",
});
```
```python
import pulumi
import pulumi_aws as aws

example_api = aws.apigatewayv2.Api("exampleApi",
    protocol_type="WEBSOCKET",
    route_selection_expression="$request.body.action")
example_route = aws.apigatewayv2.Route("exampleRoute",
    api_id=example_api.id,
    route_key="$default")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleApi = new Aws.ApiGatewayV2.Api("exampleApi", new()
    {
        ProtocolType = "WEBSOCKET",
        RouteSelectionExpression = "$request.body.action",
    });

    var exampleRoute = new Aws.ApiGatewayV2.Route("exampleRoute", new()
    {
        ApiId = exampleApi.Id,
        RouteKey = "$default",
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
		exampleApi, err := apigatewayv2.NewApi(ctx, "exampleApi", &apigatewayv2.ApiArgs{
			ProtocolType:             pulumi.String("WEBSOCKET"),
			RouteSelectionExpression: pulumi.String("$request.body.action"),
		})
		if err != nil {
			return err
		}
		_, err = apigatewayv2.NewRoute(ctx, "exampleRoute", &apigatewayv2.RouteArgs{
			ApiId:    exampleApi.ID(),
			RouteKey: pulumi.String("$default"),
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
import com.pulumi.aws.apigatewayv2.Route;
import com.pulumi.aws.apigatewayv2.RouteArgs;
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
        var exampleApi = new Api("exampleApi", ApiArgs.builder()        
            .protocolType("WEBSOCKET")
            .routeSelectionExpression("$request.body.action")
            .build());

        var exampleRoute = new Route("exampleRoute", RouteArgs.builder()        
            .apiId(exampleApi.id())
            .routeKey("$default")
            .build());

    }
}
```
```yaml
resources:
  exampleApi:
    type: aws:apigatewayv2:Api
    properties:
      protocolType: WEBSOCKET
      routeSelectionExpression: $request.body.action
  exampleRoute:
    type: aws:apigatewayv2:Route
    properties:
      apiId: ${exampleApi.id}
      routeKey: $default
```
{{% /example %}}
{{% example %}}
### HTTP Proxy Integration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleApi = new aws.apigatewayv2.Api("exampleApi", {protocolType: "HTTP"});
const exampleIntegration = new aws.apigatewayv2.Integration("exampleIntegration", {
    apiId: exampleApi.id,
    integrationType: "HTTP_PROXY",
    integrationMethod: "ANY",
    integrationUri: "https://example.com/{proxy}",
});
const exampleRoute = new aws.apigatewayv2.Route("exampleRoute", {
    apiId: exampleApi.id,
    routeKey: "ANY /example/{proxy+}",
    target: pulumi.interpolate`integrations/${exampleIntegration.id}`,
});
```
```python
import pulumi
import pulumi_aws as aws

example_api = aws.apigatewayv2.Api("exampleApi", protocol_type="HTTP")
example_integration = aws.apigatewayv2.Integration("exampleIntegration",
    api_id=example_api.id,
    integration_type="HTTP_PROXY",
    integration_method="ANY",
    integration_uri="https://example.com/{proxy}")
example_route = aws.apigatewayv2.Route("exampleRoute",
    api_id=example_api.id,
    route_key="ANY /example/{proxy+}",
    target=example_integration.id.apply(lambda id: f"integrations/{id}"))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleApi = new Aws.ApiGatewayV2.Api("exampleApi", new()
    {
        ProtocolType = "HTTP",
    });

    var exampleIntegration = new Aws.ApiGatewayV2.Integration("exampleIntegration", new()
    {
        ApiId = exampleApi.Id,
        IntegrationType = "HTTP_PROXY",
        IntegrationMethod = "ANY",
        IntegrationUri = "https://example.com/{proxy}",
    });

    var exampleRoute = new Aws.ApiGatewayV2.Route("exampleRoute", new()
    {
        ApiId = exampleApi.Id,
        RouteKey = "ANY /example/{proxy+}",
        Target = exampleIntegration.Id.Apply(id => $"integrations/{id}"),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleApi, err := apigatewayv2.NewApi(ctx, "exampleApi", &apigatewayv2.ApiArgs{
			ProtocolType: pulumi.String("HTTP"),
		})
		if err != nil {
			return err
		}
		exampleIntegration, err := apigatewayv2.NewIntegration(ctx, "exampleIntegration", &apigatewayv2.IntegrationArgs{
			ApiId:             exampleApi.ID(),
			IntegrationType:   pulumi.String("HTTP_PROXY"),
			IntegrationMethod: pulumi.String("ANY"),
			IntegrationUri:    pulumi.String("https://example.com/{proxy}"),
		})
		if err != nil {
			return err
		}
		_, err = apigatewayv2.NewRoute(ctx, "exampleRoute", &apigatewayv2.RouteArgs{
			ApiId:    exampleApi.ID(),
			RouteKey: pulumi.String("ANY /example/{proxy+}"),
			Target: exampleIntegration.ID().ApplyT(func(id string) (string, error) {
				return fmt.Sprintf("integrations/%v", id), nil
			}).(pulumi.StringOutput),
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
import com.pulumi.aws.apigatewayv2.Integration;
import com.pulumi.aws.apigatewayv2.IntegrationArgs;
import com.pulumi.aws.apigatewayv2.Route;
import com.pulumi.aws.apigatewayv2.RouteArgs;
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
        var exampleApi = new Api("exampleApi", ApiArgs.builder()        
            .protocolType("HTTP")
            .build());

        var exampleIntegration = new Integration("exampleIntegration", IntegrationArgs.builder()        
            .apiId(exampleApi.id())
            .integrationType("HTTP_PROXY")
            .integrationMethod("ANY")
            .integrationUri("https://example.com/{proxy}")
            .build());

        var exampleRoute = new Route("exampleRoute", RouteArgs.builder()        
            .apiId(exampleApi.id())
            .routeKey("ANY /example/{proxy+}")
            .target(exampleIntegration.id().applyValue(id -> String.format("integrations/%s", id)))
            .build());

    }
}
```
```yaml
resources:
  exampleApi:
    type: aws:apigatewayv2:Api
    properties:
      protocolType: HTTP
  exampleIntegration:
    type: aws:apigatewayv2:Integration
    properties:
      apiId: ${exampleApi.id}
      integrationType: HTTP_PROXY
      integrationMethod: ANY
      integrationUri: https://example.com/{proxy}
  exampleRoute:
    type: aws:apigatewayv2:Route
    properties:
      apiId: ${exampleApi.id}
      routeKey: ANY /example/{proxy+}
      target: integrations/${exampleIntegration.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_apigatewayv2_route` using the API identifier and route identifier. For example:

```sh
 $ pulumi import aws:apigatewayv2/route:Route example aabbccddee/1122334
```
 -> __Note:__ The API Gateway managed route created as part of [_quick_create_](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-basic-concept.html#apigateway-definition-quick-create) cannot be imported.

