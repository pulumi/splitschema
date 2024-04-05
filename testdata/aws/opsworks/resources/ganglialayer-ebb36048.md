Provides an OpsWorks Ganglia layer resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const monitor = new aws.opsworks.GangliaLayer("monitor", {
    stackId: aws_opsworks_stack.main.id,
    password: "foobarbaz",
});
```
```python
import pulumi
import pulumi_aws as aws

monitor = aws.opsworks.GangliaLayer("monitor",
    stack_id=aws_opsworks_stack["main"]["id"],
    password="foobarbaz")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var monitor = new Aws.OpsWorks.GangliaLayer("monitor", new()
    {
        StackId = aws_opsworks_stack.Main.Id,
        Password = "foobarbaz",
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
		_, err := opsworks.NewGangliaLayer(ctx, "monitor", &opsworks.GangliaLayerArgs{
			StackId:  pulumi.Any(aws_opsworks_stack.Main.Id),
			Password: pulumi.String("foobarbaz"),
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
import com.pulumi.aws.opsworks.GangliaLayer;
import com.pulumi.aws.opsworks.GangliaLayerArgs;
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
        var monitor = new GangliaLayer("monitor", GangliaLayerArgs.builder()        
            .stackId(aws_opsworks_stack.main().id())
            .password("foobarbaz")
            .build());

    }
}
```
```yaml
resources:
  monitor:
    type: aws:opsworks:GangliaLayer
    properties:
      stackId: ${aws_opsworks_stack.main.id}
      password: foobarbaz
```
{{% /example %}}
{{% /examples %}}