Provides an EventBridge Custom Schema Registry resource.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.schemas.Registry("test", {description: "A custom schema registry"});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.schemas.Registry("test", description="A custom schema registry")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Schemas.Registry("test", new()
    {
        Description = "A custom schema registry",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/schemas"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := schemas.NewRegistry(ctx, "test", &schemas.RegistryArgs{
			Description: pulumi.String("A custom schema registry"),
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
import com.pulumi.aws.schemas.Registry;
import com.pulumi.aws.schemas.RegistryArgs;
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
        var test = new Registry("test", RegistryArgs.builder()        
            .description("A custom schema registry")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:schemas:Registry
    properties:
      description: A custom schema registry
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge schema registries using the `name`. For example:

```sh
 $ pulumi import aws:schemas/registry:Registry test my_own_registry
```
 