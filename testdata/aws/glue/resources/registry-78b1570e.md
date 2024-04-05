Provides a Glue Registry resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Registry("example", {registryName: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Registry("example", registry_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Registry("example", new()
    {
        RegistryName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewRegistry(ctx, "example", &glue.RegistryArgs{
			RegistryName: pulumi.String("example"),
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
import com.pulumi.aws.glue.Registry;
import com.pulumi.aws.glue.RegistryArgs;
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
        var example = new Registry("example", RegistryArgs.builder()        
            .registryName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Registry
    properties:
      registryName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Registries using `arn`. For example:

```sh
 $ pulumi import aws:glue/registry:Registry example arn:aws:glue:us-west-2:123456789012:registry/example
```
 