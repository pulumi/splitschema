Data source for managing an AWS App Mesh Virtual Gateway.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.appmesh.getVirtualGateway({
    meshName: "mesh-gateway",
    name: "example-mesh",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appmesh.get_virtual_gateway(mesh_name="mesh-gateway",
    name="example-mesh")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.AppMesh.GetVirtualGateway.Invoke(new()
    {
        MeshName = "mesh-gateway",
        Name = "example-mesh",
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
		_, err := appmesh.LookupVirtualGateway(ctx, &appmesh.LookupVirtualGatewayArgs{
			MeshName: "mesh-gateway",
			Name:     "example-mesh",
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
import com.pulumi.aws.appmesh.inputs.GetVirtualGatewayArgs;
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
        final var example = AppmeshFunctions.getVirtualGateway(GetVirtualGatewayArgs.builder()
            .meshName("mesh-gateway")
            .name("example-mesh")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:appmesh:getVirtualGateway
      Arguments:
        meshName: mesh-gateway
        name: example-mesh
```

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.appmesh.AppmeshFunctions;
import com.pulumi.aws.appmesh.inputs.GetVirtualGatewayArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        final var test = AppmeshFunctions.getVirtualGateway(GetVirtualGatewayArgs.builder()
            .name("example.mesh.local")
            .meshName("example-mesh")
            .meshOwner(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
            .build());

    }
}
```
```yaml
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  test:
    fn::invoke:
      Function: aws:appmesh:getVirtualGateway
      Arguments:
        name: example.mesh.local
        meshName: example-mesh
        meshOwner: ${current.accountId}
```
{{% /example %}}
{{% /examples %}}