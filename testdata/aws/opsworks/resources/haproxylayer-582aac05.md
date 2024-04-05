Provides an OpsWorks haproxy layer resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lb = new aws.opsworks.HaproxyLayer("lb", {
    stackId: aws_opsworks_stack.main.id,
    statsPassword: "foobarbaz",
});
```
```python
import pulumi
import pulumi_aws as aws

lb = aws.opsworks.HaproxyLayer("lb",
    stack_id=aws_opsworks_stack["main"]["id"],
    stats_password="foobarbaz")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var lb = new Aws.OpsWorks.HaproxyLayer("lb", new()
    {
        StackId = aws_opsworks_stack.Main.Id,
        StatsPassword = "foobarbaz",
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
		_, err := opsworks.NewHaproxyLayer(ctx, "lb", &opsworks.HaproxyLayerArgs{
			StackId:       pulumi.Any(aws_opsworks_stack.Main.Id),
			StatsPassword: pulumi.String("foobarbaz"),
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
import com.pulumi.aws.opsworks.HaproxyLayer;
import com.pulumi.aws.opsworks.HaproxyLayerArgs;
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
        var lb = new HaproxyLayer("lb", HaproxyLayerArgs.builder()        
            .stackId(aws_opsworks_stack.main().id())
            .statsPassword("foobarbaz")
            .build());

    }
}
```
```yaml
resources:
  lb:
    type: aws:opsworks:HaproxyLayer
    properties:
      stackId: ${aws_opsworks_stack.main.id}
      statsPassword: foobarbaz
```
{{% /example %}}
{{% /examples %}}