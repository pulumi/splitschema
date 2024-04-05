Provides an ECS task set - effectively a task that is expected to run until an error occurs or a user terminates it (typically a webserver or a database).

See [ECS Task Set section in AWS developer guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-external.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ecs.TaskSet("example", {
    service: aws_ecs_service.example.id,
    cluster: aws_ecs_cluster.example.id,
    taskDefinition: aws_ecs_task_definition.example.arn,
    loadBalancers: [{
        targetGroupArn: aws_lb_target_group.example.arn,
        containerName: "mongo",
        containerPort: 8080,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ecs.TaskSet("example",
    service=aws_ecs_service["example"]["id"],
    cluster=aws_ecs_cluster["example"]["id"],
    task_definition=aws_ecs_task_definition["example"]["arn"],
    load_balancers=[aws.ecs.TaskSetLoadBalancerArgs(
        target_group_arn=aws_lb_target_group["example"]["arn"],
        container_name="mongo",
        container_port=8080,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ecs.TaskSet("example", new()
    {
        Service = aws_ecs_service.Example.Id,
        Cluster = aws_ecs_cluster.Example.Id,
        TaskDefinition = aws_ecs_task_definition.Example.Arn,
        LoadBalancers = new[]
        {
            new Aws.Ecs.Inputs.TaskSetLoadBalancerArgs
            {
                TargetGroupArn = aws_lb_target_group.Example.Arn,
                ContainerName = "mongo",
                ContainerPort = 8080,
            },
        },
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
		_, err := ecs.NewTaskSet(ctx, "example", &ecs.TaskSetArgs{
			Service:        pulumi.Any(aws_ecs_service.Example.Id),
			Cluster:        pulumi.Any(aws_ecs_cluster.Example.Id),
			TaskDefinition: pulumi.Any(aws_ecs_task_definition.Example.Arn),
			LoadBalancers: ecs.TaskSetLoadBalancerArray{
				&ecs.TaskSetLoadBalancerArgs{
					TargetGroupArn: pulumi.Any(aws_lb_target_group.Example.Arn),
					ContainerName:  pulumi.String("mongo"),
					ContainerPort:  pulumi.Int(8080),
				},
			},
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
import com.pulumi.aws.ecs.TaskSet;
import com.pulumi.aws.ecs.TaskSetArgs;
import com.pulumi.aws.ecs.inputs.TaskSetLoadBalancerArgs;
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
        var example = new TaskSet("example", TaskSetArgs.builder()        
            .service(aws_ecs_service.example().id())
            .cluster(aws_ecs_cluster.example().id())
            .taskDefinition(aws_ecs_task_definition.example().arn())
            .loadBalancers(TaskSetLoadBalancerArgs.builder()
                .targetGroupArn(aws_lb_target_group.example().arn())
                .containerName("mongo")
                .containerPort(8080)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ecs:TaskSet
    properties:
      service: ${aws_ecs_service.example.id}
      cluster: ${aws_ecs_cluster.example.id}
      taskDefinition: ${aws_ecs_task_definition.example.arn}
      loadBalancers:
        - targetGroupArn: ${aws_lb_target_group.example.arn}
          containerName: mongo
          containerPort: 8080
```
{{% /example %}}
{{% example %}}
### Ignoring Changes to Scale

You can utilize the generic resource lifecycle configuration block with `ignore_changes` to create an ECS service with an initial count of running instances, then ignore any changes to that count caused externally (e.g. Application Autoscaling).

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ecs.TaskSet;
import com.pulumi.aws.ecs.TaskSetArgs;
import com.pulumi.aws.ecs.inputs.TaskSetScaleArgs;
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
        var example = new TaskSet("example", TaskSetArgs.builder()        
            .lifecycle(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .scale(TaskSetScaleArgs.builder()
                .value(50)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ecs:TaskSet
    properties:
      lifecycle:
        ignoreChanges:
          - scale
      # Example: Run 50% of the servcie's desired count
      scale:
        value: 50
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECS Task Sets using the `task_set_id`, `service`, and `cluster` separated by commas (`,`). For example:

```sh
 $ pulumi import aws:ecs/taskSet:TaskSet example ecs-svc/7177320696926227436,arn:aws:ecs:us-west-2:123456789101:service/example/example-1234567890,arn:aws:ecs:us-west-2:123456789101:cluster/example
```
 