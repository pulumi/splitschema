> **Note:** To prevent a race condition during service deletion, make sure to set `depends_on` to the related `aws.iam.RolePolicy`; otherwise, the policy may be destroyed too soon and the ECS service will then get stuck in the `DRAINING` state.

Provides an ECS service - effectively a task that is expected to run until an error occurs or a user terminates it (typically a webserver or a database).

See [ECS Services section in AWS developer guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mongo = new aws.ecs.Service("mongo", {
    cluster: aws_ecs_cluster.foo.id,
    taskDefinition: aws_ecs_task_definition.mongo.arn,
    desiredCount: 3,
    iamRole: aws_iam_role.foo.arn,
    orderedPlacementStrategies: [{
        type: "binpack",
        field: "cpu",
    }],
    loadBalancers: [{
        targetGroupArn: aws_lb_target_group.foo.arn,
        containerName: "mongo",
        containerPort: 8080,
    }],
    placementConstraints: [{
        type: "memberOf",
        expression: "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]",
    }],
}, {
    dependsOn: [aws_iam_role_policy.foo],
});
```
```python
import pulumi
import pulumi_aws as aws

mongo = aws.ecs.Service("mongo",
    cluster=aws_ecs_cluster["foo"]["id"],
    task_definition=aws_ecs_task_definition["mongo"]["arn"],
    desired_count=3,
    iam_role=aws_iam_role["foo"]["arn"],
    ordered_placement_strategies=[aws.ecs.ServiceOrderedPlacementStrategyArgs(
        type="binpack",
        field="cpu",
    )],
    load_balancers=[aws.ecs.ServiceLoadBalancerArgs(
        target_group_arn=aws_lb_target_group["foo"]["arn"],
        container_name="mongo",
        container_port=8080,
    )],
    placement_constraints=[aws.ecs.ServicePlacementConstraintArgs(
        type="memberOf",
        expression="attribute:ecs.availability-zone in [us-west-2a, us-west-2b]",
    )],
    opts=pulumi.ResourceOptions(depends_on=[aws_iam_role_policy["foo"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mongo = new Aws.Ecs.Service("mongo", new()
    {
        Cluster = aws_ecs_cluster.Foo.Id,
        TaskDefinition = aws_ecs_task_definition.Mongo.Arn,
        DesiredCount = 3,
        IamRole = aws_iam_role.Foo.Arn,
        OrderedPlacementStrategies = new[]
        {
            new Aws.Ecs.Inputs.ServiceOrderedPlacementStrategyArgs
            {
                Type = "binpack",
                Field = "cpu",
            },
        },
        LoadBalancers = new[]
        {
            new Aws.Ecs.Inputs.ServiceLoadBalancerArgs
            {
                TargetGroupArn = aws_lb_target_group.Foo.Arn,
                ContainerName = "mongo",
                ContainerPort = 8080,
            },
        },
        PlacementConstraints = new[]
        {
            new Aws.Ecs.Inputs.ServicePlacementConstraintArgs
            {
                Type = "memberOf",
                Expression = "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]",
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_role_policy.Foo,
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
		_, err := ecs.NewService(ctx, "mongo", &ecs.ServiceArgs{
			Cluster:        pulumi.Any(aws_ecs_cluster.Foo.Id),
			TaskDefinition: pulumi.Any(aws_ecs_task_definition.Mongo.Arn),
			DesiredCount:   pulumi.Int(3),
			IamRole:        pulumi.Any(aws_iam_role.Foo.Arn),
			OrderedPlacementStrategies: ecs.ServiceOrderedPlacementStrategyArray{
				&ecs.ServiceOrderedPlacementStrategyArgs{
					Type:  pulumi.String("binpack"),
					Field: pulumi.String("cpu"),
				},
			},
			LoadBalancers: ecs.ServiceLoadBalancerArray{
				&ecs.ServiceLoadBalancerArgs{
					TargetGroupArn: pulumi.Any(aws_lb_target_group.Foo.Arn),
					ContainerName:  pulumi.String("mongo"),
					ContainerPort:  pulumi.Int(8080),
				},
			},
			PlacementConstraints: ecs.ServicePlacementConstraintArray{
				&ecs.ServicePlacementConstraintArgs{
					Type:       pulumi.String("memberOf"),
					Expression: pulumi.String("attribute:ecs.availability-zone in [us-west-2a, us-west-2b]"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_iam_role_policy.Foo,
		}))
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
import com.pulumi.aws.ecs.Service;
import com.pulumi.aws.ecs.ServiceArgs;
import com.pulumi.aws.ecs.inputs.ServiceOrderedPlacementStrategyArgs;
import com.pulumi.aws.ecs.inputs.ServiceLoadBalancerArgs;
import com.pulumi.aws.ecs.inputs.ServicePlacementConstraintArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var mongo = new Service("mongo", ServiceArgs.builder()        
            .cluster(aws_ecs_cluster.foo().id())
            .taskDefinition(aws_ecs_task_definition.mongo().arn())
            .desiredCount(3)
            .iamRole(aws_iam_role.foo().arn())
            .orderedPlacementStrategies(ServiceOrderedPlacementStrategyArgs.builder()
                .type("binpack")
                .field("cpu")
                .build())
            .loadBalancers(ServiceLoadBalancerArgs.builder()
                .targetGroupArn(aws_lb_target_group.foo().arn())
                .containerName("mongo")
                .containerPort(8080)
                .build())
            .placementConstraints(ServicePlacementConstraintArgs.builder()
                .type("memberOf")
                .expression("attribute:ecs.availability-zone in [us-west-2a, us-west-2b]")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_iam_role_policy.foo())
                .build());

    }
}
```
```yaml
resources:
  mongo:
    type: aws:ecs:Service
    properties:
      cluster: ${aws_ecs_cluster.foo.id}
      taskDefinition: ${aws_ecs_task_definition.mongo.arn}
      desiredCount: 3
      iamRole: ${aws_iam_role.foo.arn}
      orderedPlacementStrategies:
        - type: binpack
          field: cpu
      loadBalancers:
        - targetGroupArn: ${aws_lb_target_group.foo.arn}
          containerName: mongo
          containerPort: 8080
      placementConstraints:
        - type: memberOf
          expression: attribute:ecs.availability-zone in [us-west-2a, us-west-2b]
    options:
      dependson:
        - ${aws_iam_role_policy.foo}
```
{{% /example %}}
{{% example %}}
### Ignoring Changes to Desired Count

You can use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to create an ECS service with an initial count of running instances, then ignore any changes to that count caused externally (e.g. Application Autoscaling).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ... other configurations ...
const example = new aws.ecs.Service("example", {desiredCount: 2});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configurations ...
example = aws.ecs.Service("example", desired_count=2)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ... other configurations ...
    var example = new Aws.Ecs.Service("example", new()
    {
        DesiredCount = 2,
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
		_, err := ecs.NewService(ctx, "example", &ecs.ServiceArgs{
			DesiredCount: pulumi.Int(2),
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
        var example = new Service("example", ServiceArgs.builder()        
            .desiredCount(2)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ecs:Service
    properties:
      # Example: Create service with 2 instances to start
      desiredCount: 2
```
{{% /example %}}
{{% example %}}
### Daemon Scheduling Strategy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.ecs.Service("bar", {
    cluster: aws_ecs_cluster.foo.id,
    taskDefinition: aws_ecs_task_definition.bar.arn,
    schedulingStrategy: "DAEMON",
});
```
```python
import pulumi
import pulumi_aws as aws

bar = aws.ecs.Service("bar",
    cluster=aws_ecs_cluster["foo"]["id"],
    task_definition=aws_ecs_task_definition["bar"]["arn"],
    scheduling_strategy="DAEMON")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bar = new Aws.Ecs.Service("bar", new()
    {
        Cluster = aws_ecs_cluster.Foo.Id,
        TaskDefinition = aws_ecs_task_definition.Bar.Arn,
        SchedulingStrategy = "DAEMON",
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
		_, err := ecs.NewService(ctx, "bar", &ecs.ServiceArgs{
			Cluster:            pulumi.Any(aws_ecs_cluster.Foo.Id),
			TaskDefinition:     pulumi.Any(aws_ecs_task_definition.Bar.Arn),
			SchedulingStrategy: pulumi.String("DAEMON"),
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
        var bar = new Service("bar", ServiceArgs.builder()        
            .cluster(aws_ecs_cluster.foo().id())
            .taskDefinition(aws_ecs_task_definition.bar().arn())
            .schedulingStrategy("DAEMON")
            .build());

    }
}
```
```yaml
resources:
  bar:
    type: aws:ecs:Service
    properties:
      cluster: ${aws_ecs_cluster.foo.id}
      taskDefinition: ${aws_ecs_task_definition.bar.arn}
      schedulingStrategy: DAEMON
```
{{% /example %}}
{{% example %}}
### CloudWatch Deployment Alarms

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ecs.Service("example", {
    cluster: aws_ecs_cluster.example.id,
    alarms: {
        enable: true,
        rollback: true,
        alarmNames: [aws_cloudwatch_metric_alarm.example.alarm_name],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ecs.Service("example",
    cluster=aws_ecs_cluster["example"]["id"],
    alarms=aws.ecs.ServiceAlarmsArgs(
        enable=True,
        rollback=True,
        alarm_names=[aws_cloudwatch_metric_alarm["example"]["alarm_name"]],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ecs.Service("example", new()
    {
        Cluster = aws_ecs_cluster.Example.Id,
        Alarms = new Aws.Ecs.Inputs.ServiceAlarmsArgs
        {
            Enable = true,
            Rollback = true,
            AlarmNames = new[]
            {
                aws_cloudwatch_metric_alarm.Example.Alarm_name,
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
		_, err := ecs.NewService(ctx, "example", &ecs.ServiceArgs{
			Cluster: pulumi.Any(aws_ecs_cluster.Example.Id),
			Alarms: &ecs.ServiceAlarmsArgs{
				Enable:   pulumi.Bool(true),
				Rollback: pulumi.Bool(true),
				AlarmNames: pulumi.StringArray{
					aws_cloudwatch_metric_alarm.Example.Alarm_name,
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
import com.pulumi.aws.ecs.Service;
import com.pulumi.aws.ecs.ServiceArgs;
import com.pulumi.aws.ecs.inputs.ServiceAlarmsArgs;
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
        var example = new Service("example", ServiceArgs.builder()        
            .cluster(aws_ecs_cluster.example().id())
            .alarms(ServiceAlarmsArgs.builder()
                .enable(true)
                .rollback(true)
                .alarmNames(aws_cloudwatch_metric_alarm.example().alarm_name())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ecs:Service
    properties:
      cluster: ${aws_ecs_cluster.example.id}
      alarms:
        enable: true
        rollback: true
        alarmNames:
          - ${aws_cloudwatch_metric_alarm.example.alarm_name}
```
{{% /example %}}
{{% example %}}
### External Deployment Controller

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ecs.Service("example", {
    cluster: aws_ecs_cluster.example.id,
    deploymentController: {
        type: "EXTERNAL",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ecs.Service("example",
    cluster=aws_ecs_cluster["example"]["id"],
    deployment_controller=aws.ecs.ServiceDeploymentControllerArgs(
        type="EXTERNAL",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ecs.Service("example", new()
    {
        Cluster = aws_ecs_cluster.Example.Id,
        DeploymentController = new Aws.Ecs.Inputs.ServiceDeploymentControllerArgs
        {
            Type = "EXTERNAL",
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
		_, err := ecs.NewService(ctx, "example", &ecs.ServiceArgs{
			Cluster: pulumi.Any(aws_ecs_cluster.Example.Id),
			DeploymentController: &ecs.ServiceDeploymentControllerArgs{
				Type: pulumi.String("EXTERNAL"),
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
import com.pulumi.aws.ecs.Service;
import com.pulumi.aws.ecs.ServiceArgs;
import com.pulumi.aws.ecs.inputs.ServiceDeploymentControllerArgs;
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
        var example = new Service("example", ServiceArgs.builder()        
            .cluster(aws_ecs_cluster.example().id())
            .deploymentController(ServiceDeploymentControllerArgs.builder()
                .type("EXTERNAL")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ecs:Service
    properties:
      cluster: ${aws_ecs_cluster.example.id}
      deploymentController:
        type: EXTERNAL
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECS services using the `name` together with ecs cluster `name`. For example:

```sh
 $ pulumi import aws:ecs/service:Service imported cluster-name/service-name
```
 