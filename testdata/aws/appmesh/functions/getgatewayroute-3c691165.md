The App Mesh Gateway Route data source allows details of an App Mesh Gateway Route to be retrieved by its name, mesh_name, virtual_gateway_name, and optionally the mesh_owner.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.appmesh.getGatewayRoute({
    meshName: "test-mesh",
    name: "test-route",
    virtualGatewayName: "test-gateway",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.appmesh.get_gateway_route(mesh_name="test-mesh",
    name="test-route",
    virtual_gateway_name="test-gateway")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.AppMesh.GetGatewayRoute.Invoke(new()
    {
        MeshName = "test-mesh",
        Name = "test-route",
        VirtualGatewayName = "test-gateway",
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
		_, err := appmesh.LookupGatewayRoute(ctx, &appmesh.LookupGatewayRouteArgs{
			MeshName:           "test-mesh",
			Name:               "test-route",
			VirtualGatewayName: "test-gateway",
		}, nil)
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
import com.pulumi.aws.appmesh.AppmeshFunctions;
import com.pulumi.aws.appmesh.inputs.GetGatewayRouteArgs;
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
        final var test = AppmeshFunctions.getGatewayRoute(GetGatewayRouteArgs.builder()
            .meshName("test-mesh")
            .name("test-route")
            .virtualGatewayName("test-gateway")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:appmesh:getGatewayRoute
      Arguments:
        meshName: test-mesh
        name: test-route
        virtualGatewayName: test-gateway
```
{{% /example %}}
{{% /examples %}}