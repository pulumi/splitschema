Provides an OpsWorks MySQL layer resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const db = new aws.opsworks.MysqlLayer("db", {stackId: aws_opsworks_stack.main.id});
```
```python
import pulumi
import pulumi_aws as aws

db = aws.opsworks.MysqlLayer("db", stack_id=aws_opsworks_stack["main"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var db = new Aws.OpsWorks.MysqlLayer("db", new()
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
		_, err := opsworks.NewMysqlLayer(ctx, "db", &opsworks.MysqlLayerArgs{
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
import com.pulumi.aws.opsworks.MysqlLayer;
import com.pulumi.aws.opsworks.MysqlLayerArgs;
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
        var db = new MysqlLayer("db", MysqlLayerArgs.builder()        
            .stackId(aws_opsworks_stack.main().id())
            .build());

    }
}
```
```yaml
resources:
  db:
    type: aws:opsworks:MysqlLayer
    properties:
      stackId: ${aws_opsworks_stack.main.id}
```
{{% /example %}}
{{% /examples %}}