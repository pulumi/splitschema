The App Mesh Virtual Router data source allows details of an App Mesh Virtual Service to be retrieved by its name and mesh_name.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.appmesh.getVirtualRouter({
    meshName: "example-mesh-name",
    name: "example-router-name",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.appmesh.get_virtual_router(mesh_name="example-mesh-name",
    name="example-router-name")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.AppMesh.GetVirtualRouter.Invoke(new()
    {
        MeshName = "example-mesh-name",
        Name = "example-router-name",
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
		_, err := appmesh.LookupVirtualRouter(ctx, &appmesh.LookupVirtualRouterArgs{
			MeshName: "example-mesh-name",
			Name:     "example-router-name",
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
import com.pulumi.aws.appmesh.inputs.GetVirtualRouterArgs;
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
        final var test = AppmeshFunctions.getVirtualRouter(GetVirtualRouterArgs.builder()
            .meshName("example-mesh-name")
            .name("example-router-name")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:appmesh:getVirtualRouter
      Arguments:
        meshName: example-mesh-name
        name: example-router-name
```
{{% /example %}}
{{% /examples %}}