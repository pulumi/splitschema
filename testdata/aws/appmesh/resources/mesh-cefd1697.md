Provides an AWS App Mesh service mesh resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const simple = new aws.appmesh.Mesh("simple", {});
```
```python
import pulumi
import pulumi_aws as aws

simple = aws.appmesh.Mesh("simple")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var simple = new Aws.AppMesh.Mesh("simple");

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
		_, err := appmesh.NewMesh(ctx, "simple", nil)
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
import com.pulumi.aws.appmesh.Mesh;
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
        var simple = new Mesh("simple");

    }
}
```
```yaml
resources:
  simple:
    type: aws:appmesh:Mesh
```
{{% /example %}}
{{% example %}}
### Egress Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const simple = new aws.appmesh.Mesh("simple", {spec: {
    egressFilter: {
        type: "ALLOW_ALL",
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

simple = aws.appmesh.Mesh("simple", spec=aws.appmesh.MeshSpecArgs(
    egress_filter=aws.appmesh.MeshSpecEgressFilterArgs(
        type="ALLOW_ALL",
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
    var simple = new Aws.AppMesh.Mesh("simple", new()
    {
        Spec = new Aws.AppMesh.Inputs.MeshSpecArgs
        {
            EgressFilter = new Aws.AppMesh.Inputs.MeshSpecEgressFilterArgs
            {
                Type = "ALLOW_ALL",
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
		_, err := appmesh.NewMesh(ctx, "simple", &appmesh.MeshArgs{
			Spec: &appmesh.MeshSpecArgs{
				EgressFilter: &appmesh.MeshSpecEgressFilterArgs{
					Type: pulumi.String("ALLOW_ALL"),
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
import com.pulumi.aws.appmesh.Mesh;
import com.pulumi.aws.appmesh.MeshArgs;
import com.pulumi.aws.appmesh.inputs.MeshSpecArgs;
import com.pulumi.aws.appmesh.inputs.MeshSpecEgressFilterArgs;
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
        var simple = new Mesh("simple", MeshArgs.builder()        
            .spec(MeshSpecArgs.builder()
                .egressFilter(MeshSpecEgressFilterArgs.builder()
                    .type("ALLOW_ALL")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  simple:
    type: aws:appmesh:Mesh
    properties:
      spec:
        egressFilter:
          type: ALLOW_ALL
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Mesh service meshes using the `name`. For example:

```sh
 $ pulumi import aws:appmesh/mesh:Mesh simple simpleapp
```
 