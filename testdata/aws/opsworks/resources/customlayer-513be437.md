Provides an OpsWorks custom layer resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const custlayer = new aws.opsworks.CustomLayer("custlayer", {
    shortName: "awesome",
    stackId: aws_opsworks_stack.main.id,
});
```
```python
import pulumi
import pulumi_aws as aws

custlayer = aws.opsworks.CustomLayer("custlayer",
    short_name="awesome",
    stack_id=aws_opsworks_stack["main"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var custlayer = new Aws.OpsWorks.CustomLayer("custlayer", new()
    {
        ShortName = "awesome",
        StackId = aws_opsworks_stack.Main.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewCustomLayer(ctx, "custlayer", &opsworks.CustomLayerArgs{
			ShortName: pulumi.String("awesome"),
			StackId:   pulumi.Any(aws_opsworks_stack.Main.Id),
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
import com.pulumi.aws.opsworks.CustomLayer;
import com.pulumi.aws.opsworks.CustomLayerArgs;
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
        var custlayer = new CustomLayer("custlayer", CustomLayerArgs.builder()        
            .shortName("awesome")
            .stackId(aws_opsworks_stack.main().id())
            .build());

    }
}
```
```yaml
resources:
  custlayer:
    type: aws:opsworks:CustomLayer
    properties:
      shortName: awesome
      stackId: ${aws_opsworks_stack.main.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import OpsWorks Custom Layers using the `id`. For example:

```sh
 $ pulumi import aws:opsworks/customLayer:CustomLayer bar 00000000-0000-0000-0000-000000000000
```
 