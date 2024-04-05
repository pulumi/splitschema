The App Mesh Mesh data source allows details of an App Mesh Mesh to be retrieved by its name and optionally the mesh_owner.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const simple = aws.appmesh.getMesh({
    name: "simpleapp",
});
```
```python
import pulumi
import pulumi_aws as aws

simple = aws.appmesh.get_mesh(name="simpleapp")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var simple = Aws.AppMesh.GetMesh.Invoke(new()
    {
        Name = "simpleapp",
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
		_, err := appmesh.LookupMesh(ctx, &appmesh.LookupMeshArgs{
			Name: "simpleapp",
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
import com.pulumi.aws.appmesh.inputs.GetMeshArgs;
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
        final var simple = AppmeshFunctions.getMesh(GetMeshArgs.builder()
            .name("simpleapp")
            .build());

    }
}
```
```yaml
variables:
  simple:
    fn::invoke:
      Function: aws:appmesh:getMesh
      Arguments:
        name: simpleapp
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const simple = current.then(current => aws.appmesh.getMesh({
    name: "simpleapp",
    meshOwner: current.accountId,
}));
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
simple = aws.appmesh.get_mesh(name="simpleapp",
    mesh_owner=current.account_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var simple = Aws.AppMesh.GetMesh.Invoke(new()
    {
        Name = "simpleapp",
        MeshOwner = current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appmesh"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = appmesh.LookupMesh(ctx, &appmesh.LookupMeshArgs{
			Name:      "simpleapp",
			MeshOwner: pulumi.StringRef(current.AccountId),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.appmesh.AppmeshFunctions;
import com.pulumi.aws.appmesh.inputs.GetMeshArgs;
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

        final var simple = AppmeshFunctions.getMesh(GetMeshArgs.builder()
            .name("simpleapp")
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
  simple:
    fn::invoke:
      Function: aws:appmesh:getMesh
      Arguments:
        name: simpleapp
        meshOwner: ${current.accountId}
```
{{% /example %}}
{{% /examples %}}