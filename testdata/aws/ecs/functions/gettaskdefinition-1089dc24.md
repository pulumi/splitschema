The ECS task definition data source allows access to details of
a specific AWS ECS task definition.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mongoTaskDefinition = aws.ecs.getTaskDefinition({
    taskDefinition: mongoEcs / taskDefinitionTaskDefinition.family,
});
const foo = new aws.ecs.Cluster("foo", {});
const mongoEcs_taskDefinitionTaskDefinition = new aws.ecs.TaskDefinition("mongoEcs/taskDefinitionTaskDefinition", {
    family: "mongodb",
    containerDefinitions: `[
  {
    "cpu": 128,
    "environment": [{
      "name": "SECRET",
      "value": "KEY"
    }],
    "essential": true,
    "image": "mongo:latest",
    "memory": 128,
    "memoryReservation": 64,
    "name": "mongodb"
  }
]
`,
});
const mongoService = new aws.ecs.Service("mongoService", {
    cluster: foo.id,
    desiredCount: 2,
    taskDefinition: mongoTaskDefinition.then(mongoTaskDefinition => mongoTaskDefinition.arn),
});
```
```python
import pulumi
import pulumi_aws as aws

mongo_task_definition = aws.ecs.get_task_definition(task_definition=mongo_ecs / task_definition_task_definition["family"])
foo = aws.ecs.Cluster("foo")
mongo_ecs_task_definition_task_definition = aws.ecs.TaskDefinition("mongoEcs/taskDefinitionTaskDefinition",
    family="mongodb",
    container_definitions="""[
  {
    "cpu": 128,
    "environment": [{
      "name": "SECRET",
      "value": "KEY"
    }],
    "essential": true,
    "image": "mongo:latest",
    "memory": 128,
    "memoryReservation": 64,
    "name": "mongodb"
  }
]
""")
mongo_service = aws.ecs.Service("mongoService",
    cluster=foo.id,
    desired_count=2,
    task_definition=mongo_task_definition.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mongoTaskDefinition = Aws.Ecs.GetTaskDefinition.Invoke(new()
    {
        TaskDefinition = mongoEcs / taskDefinitionTaskDefinition.Family,
    });

    var foo = new Aws.Ecs.Cluster("foo");

    var mongoEcs_taskDefinitionTaskDefinition = new Aws.Ecs.TaskDefinition("mongoEcs/taskDefinitionTaskDefinition", new()
    {
        Family = "mongodb",
        ContainerDefinitions = @"[
  {
    ""cpu"": 128,
    ""environment"": [{
      ""name"": ""SECRET"",
      ""value"": ""KEY""
    }],
    ""essential"": true,
    ""image"": ""mongo:latest"",
    ""memory"": 128,
    ""memoryReservation"": 64,
    ""name"": ""mongodb""
  }
]
",
    });

    var mongoService = new Aws.Ecs.Service("mongoService", new()
    {
        Cluster = foo.Id,
        DesiredCount = 2,
        TaskDefinition = mongoTaskDefinition.Apply(getTaskDefinitionResult => getTaskDefinitionResult.Arn),
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
		mongoTaskDefinition, err := ecs.LookupTaskDefinition(ctx, &ecs.LookupTaskDefinitionArgs{
			TaskDefinition: mongoEcs / taskDefinitionTaskDefinition.Family,
		}, nil)
		if err != nil {
			return err
		}
		foo, err := ecs.NewCluster(ctx, "foo", nil)
		if err != nil {
			return err
		}
		_, err = ecs.NewTaskDefinition(ctx, "mongoEcs/taskDefinitionTaskDefinition", &ecs.TaskDefinitionArgs{
			Family: pulumi.String("mongodb"),
			ContainerDefinitions: pulumi.String(`[
  {
    "cpu": 128,
    "environment": [{
      "name": "SECRET",
      "value": "KEY"
    }],
    "essential": true,
    "image": "mongo:latest",
    "memory": 128,
    "memoryReservation": 64,
    "name": "mongodb"
  }
]
`),
		})
		if err != nil {
			return err
		}
		_, err = ecs.NewService(ctx, "mongoService", &ecs.ServiceArgs{
			Cluster:        foo.ID(),
			DesiredCount:   pulumi.Int(2),
			TaskDefinition: *pulumi.String(mongoTaskDefinition.Arn),
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
import com.pulumi.aws.ecs.EcsFunctions;
import com.pulumi.aws.ecs.inputs.GetTaskDefinitionArgs;
import com.pulumi.aws.ecs.Cluster;
import com.pulumi.aws.ecs.TaskDefinition;
import com.pulumi.aws.ecs.TaskDefinitionArgs;
import com.pulumi.aws.ecs.Service;
import com.pulumi.aws.ecs.ServiceArgs;
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
        final var mongoTaskDefinition = EcsFunctions.getTaskDefinition(GetTaskDefinitionArgs.builder()
            .taskDefinition(mongoEcs / taskDefinitionTaskDefinition.family())
            .build());

        var foo = new Cluster("foo");

        var mongoEcs_taskDefinitionTaskDefinition = new TaskDefinition("mongoEcs/taskDefinitionTaskDefinition", TaskDefinitionArgs.builder()        
            .family("mongodb")
            .containerDefinitions("""
[
  {
    "cpu": 128,
    "environment": [{
      "name": "SECRET",
      "value": "KEY"
    }],
    "essential": true,
    "image": "mongo:latest",
    "memory": 128,
    "memoryReservation": 64,
    "name": "mongodb"
  }
]
            """)
            .build());

        var mongoService = new Service("mongoService", ServiceArgs.builder()        
            .cluster(foo.id())
            .desiredCount(2)
            .taskDefinition(mongoTaskDefinition.applyValue(getTaskDefinitionResult -> getTaskDefinitionResult.arn()))
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}