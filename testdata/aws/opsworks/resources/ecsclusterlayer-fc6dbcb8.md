Provides an OpsWorks ECS Cluster layer resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.opsworks.EcsClusterLayer("example", {
    stackId: aws_opsworks_stack.example.id,
    ecsClusterArn: aws_ecs_cluster.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.opsworks.EcsClusterLayer("example",
    stack_id=aws_opsworks_stack["example"]["id"],
    ecs_cluster_arn=aws_ecs_cluster["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.OpsWorks.EcsClusterLayer("example", new()
    {
        StackId = aws_opsworks_stack.Example.Id,
        EcsClusterArn = aws_ecs_cluster.Example.Arn,
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
		_, err := opsworks.NewEcsClusterLayer(ctx, "example", &opsworks.EcsClusterLayerArgs{
			StackId:       pulumi.Any(aws_opsworks_stack.Example.Id),
			EcsClusterArn: pulumi.Any(aws_ecs_cluster.Example.Arn),
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
import com.pulumi.aws.opsworks.EcsClusterLayer;
import com.pulumi.aws.opsworks.EcsClusterLayerArgs;
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
        var example = new EcsClusterLayer("example", EcsClusterLayerArgs.builder()        
            .stackId(aws_opsworks_stack.example().id())
            .ecsClusterArn(aws_ecs_cluster.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:opsworks:EcsClusterLayer
    properties:
      stackId: ${aws_opsworks_stack.example.id}
      ecsClusterArn: ${aws_ecs_cluster.example.arn}
```
{{% /example %}}
{{% /examples %}}