Provides an OpsWorks static web server layer resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const web = new aws.opsworks.StaticWebLayer("web", {stackId: aws_opsworks_stack.main.id});
```
```python
import pulumi
import pulumi_aws as aws

web = aws.opsworks.StaticWebLayer("web", stack_id=aws_opsworks_stack["main"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var web = new Aws.OpsWorks.StaticWebLayer("web", new()
    {
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
		_, err := opsworks.NewStaticWebLayer(ctx, "web", &opsworks.StaticWebLayerArgs{
			StackId: pulumi.Any(aws_opsworks_stack.Main.Id),
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
import com.pulumi.aws.opsworks.StaticWebLayer;
import com.pulumi.aws.opsworks.StaticWebLayerArgs;
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
        var web = new StaticWebLayer("web", StaticWebLayerArgs.builder()        
            .stackId(aws_opsworks_stack.main().id())
            .build());

    }
}
```
```yaml
resources:
  web:
    type: aws:opsworks:StaticWebLayer
    properties:
      stackId: ${aws_opsworks_stack.main.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import OpsWorks static web server Layers using the `id`. For example:

```sh
 $ pulumi import aws:opsworks/staticWebLayer:StaticWebLayer bar 00000000-0000-0000-0000-000000000000
```
 