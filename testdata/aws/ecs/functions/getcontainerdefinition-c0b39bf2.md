The ECS container definition data source allows access to details of
a specific container within an AWS ECS service.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecs-mongo = aws.ecs.getContainerDefinition({
    taskDefinition: aws_ecs_task_definition.mongo.id,
    containerName: "mongodb",
});
```
```python
import pulumi
import pulumi_aws as aws

ecs_mongo = aws.ecs.get_container_definition(task_definition=aws_ecs_task_definition["mongo"]["id"],
    container_name="mongodb")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ecs_mongo = Aws.Ecs.GetContainerDefinition.Invoke(new()
    {
        TaskDefinition = aws_ecs_task_definition.Mongo.Id,
        ContainerName = "mongodb",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.GetContainerDefinition(ctx, &ecs.GetContainerDefinitionArgs{
			TaskDefinition: aws_ecs_task_definition.Mongo.Id,
			ContainerName:  "mongodb",
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
import com.pulumi.aws.ecs.EcsFunctions;
import com.pulumi.aws.ecs.inputs.GetContainerDefinitionArgs;
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
        final var ecs-mongo = EcsFunctions.getContainerDefinition(GetContainerDefinitionArgs.builder()
            .taskDefinition(aws_ecs_task_definition.mongo().id())
            .containerName("mongodb")
            .build());

    }
}
```
```yaml
variables:
  ecs-mongo:
    fn::invoke:
      Function: aws:ecs:getContainerDefinition
      Arguments:
        taskDefinition: ${aws_ecs_task_definition.mongo.id}
        containerName: mongodb
```
{{% /example %}}
{{% /examples %}}