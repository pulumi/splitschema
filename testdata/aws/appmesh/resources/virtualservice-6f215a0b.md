Provides an AWS App Mesh virtual service resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Virtual Node Provider

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const servicea = new aws.appmesh.VirtualService("servicea", {
    meshName: aws_appmesh_mesh.simple.id,
    spec: {
        provider: {
            virtualNode: {
                virtualNodeName: aws_appmesh_virtual_node.serviceb1.name,
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

servicea = aws.appmesh.VirtualService("servicea",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec=aws.appmesh.VirtualServiceSpecArgs(
        provider=aws.appmesh.VirtualServiceSpecProviderArgs(
            virtual_node=aws.appmesh.VirtualServiceSpecProviderVirtualNodeArgs(
                virtual_node_name=aws_appmesh_virtual_node["serviceb1"]["name"],
            ),
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var servicea = new Aws.AppMesh.VirtualService("servicea", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        Spec = new Aws.AppMesh.Inputs.VirtualServiceSpecArgs
        {
            Provider = new Aws.AppMesh.Inputs.VirtualServiceSpecProviderArgs
            {
                VirtualNode = new Aws.AppMesh.Inputs.VirtualServiceSpecProviderVirtualNodeArgs
                {
                    VirtualNodeName = aws_appmesh_virtual_node.Serviceb1.Name,
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
		_, err := appmesh.NewVirtualService(ctx, "servicea", &appmesh.VirtualServiceArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.VirtualServiceSpecArgs{
				Provider: &appmesh.VirtualServiceSpecProviderArgs{
					VirtualNode: &appmesh.VirtualServiceSpecProviderVirtualNodeArgs{
						VirtualNodeName: pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
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
import com.pulumi.aws.appmesh.VirtualService;
import com.pulumi.aws.appmesh.VirtualServiceArgs;
import com.pulumi.aws.appmesh.inputs.VirtualServiceSpecArgs;
import com.pulumi.aws.appmesh.inputs.VirtualServiceSpecProviderArgs;
import com.pulumi.aws.appmesh.inputs.VirtualServiceSpecProviderVirtualNodeArgs;
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
        var servicea = new VirtualService("servicea", VirtualServiceArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .spec(VirtualServiceSpecArgs.builder()
                .provider(VirtualServiceSpecProviderArgs.builder()
                    .virtualNode(VirtualServiceSpecProviderVirtualNodeArgs.builder()
                        .virtualNodeName(aws_appmesh_virtual_node.serviceb1().name())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  servicea:
    type: aws:appmesh:VirtualService
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      spec:
        provider:
          virtualNode:
            virtualNodeName: ${aws_appmesh_virtual_node.serviceb1.name}
```
{{% /example %}}
{{% example %}}
### Virtual Router Provider

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const servicea = new aws.appmesh.VirtualService("servicea", {
    meshName: aws_appmesh_mesh.simple.id,
    spec: {
        provider: {
            virtualRouter: {
                virtualRouterName: aws_appmesh_virtual_router.serviceb.name,
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

servicea = aws.appmesh.VirtualService("servicea",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec=aws.appmesh.VirtualServiceSpecArgs(
        provider=aws.appmesh.VirtualServiceSpecProviderArgs(
            virtual_router=aws.appmesh.VirtualServiceSpecProviderVirtualRouterArgs(
                virtual_router_name=aws_appmesh_virtual_router["serviceb"]["name"],
            ),
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var servicea = new Aws.AppMesh.VirtualService("servicea", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        Spec = new Aws.AppMesh.Inputs.VirtualServiceSpecArgs
        {
            Provider = new Aws.AppMesh.Inputs.VirtualServiceSpecProviderArgs
            {
                VirtualRouter = new Aws.AppMesh.Inputs.VirtualServiceSpecProviderVirtualRouterArgs
                {
                    VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
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
		_, err := appmesh.NewVirtualService(ctx, "servicea", &appmesh.VirtualServiceArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.VirtualServiceSpecArgs{
				Provider: &appmesh.VirtualServiceSpecProviderArgs{
					VirtualRouter: &appmesh.VirtualServiceSpecProviderVirtualRouterArgs{
						VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
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
import com.pulumi.aws.appmesh.VirtualService;
import com.pulumi.aws.appmesh.VirtualServiceArgs;
import com.pulumi.aws.appmesh.inputs.VirtualServiceSpecArgs;
import com.pulumi.aws.appmesh.inputs.VirtualServiceSpecProviderArgs;
import com.pulumi.aws.appmesh.inputs.VirtualServiceSpecProviderVirtualRouterArgs;
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
        var servicea = new VirtualService("servicea", VirtualServiceArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .spec(VirtualServiceSpecArgs.builder()
                .provider(VirtualServiceSpecProviderArgs.builder()
                    .virtualRouter(VirtualServiceSpecProviderVirtualRouterArgs.builder()
                        .virtualRouterName(aws_appmesh_virtual_router.serviceb().name())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  servicea:
    type: aws:appmesh:VirtualService
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      spec:
        provider:
          virtualRouter:
            virtualRouterName: ${aws_appmesh_virtual_router.serviceb.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Mesh virtual services using `mesh_name` together with the virtual service's `name`. For example:

```sh
 $ pulumi import aws:appmesh/virtualService:VirtualService servicea simpleapp/servicea.simpleapp.local
```
 