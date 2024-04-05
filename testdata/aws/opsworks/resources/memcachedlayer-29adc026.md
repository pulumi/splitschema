Provides an OpsWorks memcached layer resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cache = new aws.opsworks.MemcachedLayer("cache", {stackId: aws_opsworks_stack.main.id});
```
```python
import pulumi
import pulumi_aws as aws

cache = aws.opsworks.MemcachedLayer("cache", stack_id=aws_opsworks_stack["main"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var cache = new Aws.OpsWorks.MemcachedLayer("cache", new()
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
		_, err := opsworks.NewMemcachedLayer(ctx, "cache", &opsworks.MemcachedLayerArgs{
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
import com.pulumi.aws.opsworks.MemcachedLayer;
import com.pulumi.aws.opsworks.MemcachedLayerArgs;
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
        var cache = new MemcachedLayer("cache", MemcachedLayerArgs.builder()        
            .stackId(aws_opsworks_stack.main().id())
            .build());

    }
}
```
```yaml
resources:
  cache:
    type: aws:opsworks:MemcachedLayer
    properties:
      stackId: ${aws_opsworks_stack.main.id}
```
{{% /example %}}
{{% /examples %}}