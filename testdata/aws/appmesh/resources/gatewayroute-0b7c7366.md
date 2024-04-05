Provides an AWS App Mesh gateway route resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appmesh.GatewayRoute("example", {
    meshName: "example-service-mesh",
    virtualGatewayName: aws_appmesh_virtual_gateway.example.name,
    spec: {
        httpRoute: {
            action: {
                target: {
                    virtualService: {
                        virtualServiceName: aws_appmesh_virtual_service.example.name,
                    },
                },
            },
            match: {
                prefix: "/",
            },
        },
    },
    tags: {
        Environment: "test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appmesh.GatewayRoute("example",
    mesh_name="example-service-mesh",
    virtual_gateway_name=aws_appmesh_virtual_gateway["example"]["name"],
    spec=aws.appmesh.GatewayRouteSpecArgs(
        http_route=aws.appmesh.GatewayRouteSpecHttpRouteArgs(
            action=aws.appmesh.GatewayRouteSpecHttpRouteActionArgs(
                target=aws.appmesh.GatewayRouteSpecHttpRouteActionTargetArgs(
                    virtual_service=aws.appmesh.GatewayRouteSpecHttpRouteActionTargetVirtualServiceArgs(
                        virtual_service_name=aws_appmesh_virtual_service["example"]["name"],
                    ),
                ),
            ),
            match=aws.appmesh.GatewayRouteSpecHttpRouteMatchArgs(
                prefix="/",
            ),
        ),
    ),
    tags={
        "Environment": "test",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppMesh.GatewayRoute("example", new()
    {
        MeshName = "example-service-mesh",
        VirtualGatewayName = aws_appmesh_virtual_gateway.Example.Name,
        Spec = new Aws.AppMesh.Inputs.GatewayRouteSpecArgs
        {
            HttpRoute = new Aws.AppMesh.Inputs.GatewayRouteSpecHttpRouteArgs
            {
                Action = new Aws.AppMesh.Inputs.GatewayRouteSpecHttpRouteActionArgs
                {
                    Target = new Aws.AppMesh.Inputs.GatewayRouteSpecHttpRouteActionTargetArgs
                    {
                        VirtualService = new Aws.AppMesh.Inputs.GatewayRouteSpecHttpRouteActionTargetVirtualServiceArgs
                        {
                            VirtualServiceName = aws_appmesh_virtual_service.Example.Name,
                        },
                    },
                },
                Match = new Aws.AppMesh.Inputs.GatewayRouteSpecHttpRouteMatchArgs
                {
                    Prefix = "/",
                },
            },
        },
        Tags = 
        {
            { "Environment", "test" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appmesh"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appmesh.NewGatewayRoute(ctx, "example", &appmesh.GatewayRouteArgs{
			MeshName:           pulumi.String("example-service-mesh"),
			VirtualGatewayName: pulumi.Any(aws_appmesh_virtual_gateway.Example.Name),
			Spec: &appmesh.GatewayRouteSpecArgs{
				HttpRoute: &appmesh.GatewayRouteSpecHttpRouteArgs{
					Action: &appmesh.GatewayRouteSpecHttpRouteActionArgs{
						Target: &appmesh.GatewayRouteSpecHttpRouteActionTargetArgs{
							VirtualService: &appmesh.GatewayRouteSpecHttpRouteActionTargetVirtualServiceArgs{
								VirtualServiceName: pulumi.Any(aws_appmesh_virtual_service.Example.Name),
							},
						},
					},
					Match: &appmesh.GatewayRouteSpecHttpRouteMatchArgs{
						Prefix: pulumi.String("/"),
					},
				},
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("test"),
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
import com.pulumi.aws.appmesh.GatewayRoute;
import com.pulumi.aws.appmesh.GatewayRouteArgs;
import com.pulumi.aws.appmesh.inputs.GatewayRouteSpecArgs;
import com.pulumi.aws.appmesh.inputs.GatewayRouteSpecHttpRouteArgs;
import com.pulumi.aws.appmesh.inputs.GatewayRouteSpecHttpRouteActionArgs;
import com.pulumi.aws.appmesh.inputs.GatewayRouteSpecHttpRouteActionTargetArgs;
import com.pulumi.aws.appmesh.inputs.GatewayRouteSpecHttpRouteActionTargetVirtualServiceArgs;
import com.pulumi.aws.appmesh.inputs.GatewayRouteSpecHttpRouteMatchArgs;
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
        var example = new GatewayRoute("example", GatewayRouteArgs.builder()        
            .meshName("example-service-mesh")
            .virtualGatewayName(aws_appmesh_virtual_gateway.example().name())
            .spec(GatewayRouteSpecArgs.builder()
                .httpRoute(GatewayRouteSpecHttpRouteArgs.builder()
                    .action(GatewayRouteSpecHttpRouteActionArgs.builder()
                        .target(GatewayRouteSpecHttpRouteActionTargetArgs.builder()
                            .virtualService(GatewayRouteSpecHttpRouteActionTargetVirtualServiceArgs.builder()
                                .virtualServiceName(aws_appmesh_virtual_service.example().name())
                                .build())
                            .build())
                        .build())
                    .match(GatewayRouteSpecHttpRouteMatchArgs.builder()
                        .prefix("/")
                        .build())
                    .build())
                .build())
            .tags(Map.of("Environment", "test"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appmesh:GatewayRoute
    properties:
      meshName: example-service-mesh
      virtualGatewayName: ${aws_appmesh_virtual_gateway.example.name}
      spec:
        httpRoute:
          action:
            target:
              virtualService:
                virtualServiceName: ${aws_appmesh_virtual_service.example.name}
          match:
            prefix: /
      tags:
        Environment: test
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Mesh gateway routes using `mesh_name` and `virtual_gateway_name` together with the gateway route's `name`. For example:

```sh
 $ pulumi import aws:appmesh/gatewayRoute:GatewayRoute example mesh/gw1/example-gateway-route
```
 