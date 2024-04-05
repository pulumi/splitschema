Data source for managing an AWS App Mesh Virtual Node.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.appmesh.getVirtualNode({
    meshName: "example-mesh",
    name: "serviceBv1",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.appmesh.get_virtual_node(mesh_name="example-mesh",
    name="serviceBv1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.AppMesh.GetVirtualNode.Invoke(new()
    {
        MeshName = "example-mesh",
        Name = "serviceBv1",
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
		_, err := appmesh.LookupVirtualNode(ctx, &appmesh.LookupVirtualNodeArgs{
			MeshName: "example-mesh",
			Name:     "serviceBv1",
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
import com.pulumi.aws.appmesh.inputs.GetVirtualNodeArgs;
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
        final var test = AppmeshFunctions.getVirtualNode(GetVirtualNodeArgs.builder()
            .meshName("example-mesh")
            .name("serviceBv1")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:appmesh:getVirtualNode
      Arguments:
        meshName: example-mesh
        name: serviceBv1
```
{{% /example %}}
{{% /examples %}}