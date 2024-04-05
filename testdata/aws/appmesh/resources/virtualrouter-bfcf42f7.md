Provides an AWS App Mesh virtual router resource.

## Breaking Changes

Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92) and [here](https://github.com/awslabs/aws-app-mesh-examples/issues/94)), `aws.appmesh.VirtualRouter` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:

* Remove service `service_names` from the `spec` argument. AWS has created a `aws.appmesh.VirtualService` resource for each service name. Import these resource using `pulumi import`.

* Add a `listener` configuration block to the `spec` argument.

The state associated with existing resources will automatically be migrated.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb = new aws.appmesh.VirtualRouter("serviceb", {
    meshName: aws_appmesh_mesh.simple.id,
    spec: {
        listeners: [{
            portMapping: {
                port: 8080,
                protocol: "http",
            },
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

serviceb = aws.appmesh.VirtualRouter("serviceb",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec=aws.appmesh.VirtualRouterSpecArgs(
        listeners=[aws.appmesh.VirtualRouterSpecListenerArgs(
            port_mapping=aws.appmesh.VirtualRouterSpecListenerPortMappingArgs(
                port=8080,
                protocol="http",
            ),
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var serviceb = new Aws.AppMesh.VirtualRouter("serviceb", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        Spec = new Aws.AppMesh.Inputs.VirtualRouterSpecArgs
        {
            Listeners = new[]
            {
                new Aws.AppMesh.Inputs.VirtualRouterSpecListenerArgs
                {
                    PortMapping = new Aws.AppMesh.Inputs.VirtualRouterSpecListenerPortMappingArgs
                    {
                        Port = 8080,
                        Protocol = "http",
                    },
                },
            },
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
		_, err := appmesh.NewVirtualRouter(ctx, "serviceb", &appmesh.VirtualRouterArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.VirtualRouterSpecArgs{
				Listeners: appmesh.VirtualRouterSpecListenerArray{
					&appmesh.VirtualRouterSpecListenerArgs{
						PortMapping: &appmesh.VirtualRouterSpecListenerPortMappingArgs{
							Port:     pulumi.Int(8080),
							Protocol: pulumi.String("http"),
						},
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
import com.pulumi.aws.appmesh.VirtualRouter;
import com.pulumi.aws.appmesh.VirtualRouterArgs;
import com.pulumi.aws.appmesh.inputs.VirtualRouterSpecArgs;
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
        var serviceb = new VirtualRouter("serviceb", VirtualRouterArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .spec(VirtualRouterSpecArgs.builder()
                .listeners(VirtualRouterSpecListenerArgs.builder()
                    .portMapping(VirtualRouterSpecListenerPortMappingArgs.builder()
                        .port(8080)
                        .protocol("http")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  serviceb:
    type: aws:appmesh:VirtualRouter
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      spec:
        listeners:
          - portMapping:
              port: 8080
              protocol: http
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Mesh virtual routers using `mesh_name` together with the virtual router's `name`. For example:

```sh
 $ pulumi import aws:appmesh/virtualRouter:VirtualRouter serviceb simpleapp/serviceB
```
 