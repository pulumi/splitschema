Provides an Application AutoScaling ScheduledAction resource.

{{% examples %}}
## Example Usage
{{% example %}}
### DynamoDB Table Autoscaling

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dynamodbTarget = new aws.appautoscaling.Target("dynamodbTarget", {
    maxCapacity: 100,
    minCapacity: 5,
    resourceId: "table/tableName",
    scalableDimension: "dynamodb:table:ReadCapacityUnits",
    serviceNamespace: "dynamodb",
});
const dynamodbScheduledAction = new aws.appautoscaling.ScheduledAction("dynamodbScheduledAction", {
    serviceNamespace: dynamodbTarget.serviceNamespace,
    resourceId: dynamodbTarget.resourceId,
    scalableDimension: dynamodbTarget.scalableDimension,
    schedule: "at(2006-01-02T15:04:05)",
    scalableTargetAction: {
        minCapacity: 1,
        maxCapacity: 200,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

dynamodb_target = aws.appautoscaling.Target("dynamodbTarget",
    max_capacity=100,
    min_capacity=5,
    resource_id="table/tableName",
    scalable_dimension="dynamodb:table:ReadCapacityUnits",
    service_namespace="dynamodb")
dynamodb_scheduled_action = aws.appautoscaling.ScheduledAction("dynamodbScheduledAction",
    service_namespace=dynamodb_target.service_namespace,
    resource_id=dynamodb_target.resource_id,
    scalable_dimension=dynamodb_target.scalable_dimension,
    schedule="at(2006-01-02T15:04:05)",
    scalable_target_action=aws.appautoscaling.ScheduledActionScalableTargetActionArgs(
        min_capacity=1,
        max_capacity=200,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var dynamodbTarget = new Aws.AppAutoScaling.Target("dynamodbTarget", new()
    {
        MaxCapacity = 100,
        MinCapacity = 5,
        ResourceId = "table/tableName",
        ScalableDimension = "dynamodb:table:ReadCapacityUnits",
        ServiceNamespace = "dynamodb",
    });

    var dynamodbScheduledAction = new Aws.AppAutoScaling.ScheduledAction("dynamodbScheduledAction", new()
    {
        ServiceNamespace = dynamodbTarget.ServiceNamespace,
        ResourceId = dynamodbTarget.ResourceId,
        ScalableDimension = dynamodbTarget.ScalableDimension,
        Schedule = "at(2006-01-02T15:04:05)",
        ScalableTargetAction = new Aws.AppAutoScaling.Inputs.ScheduledActionScalableTargetActionArgs
        {
            MinCapacity = 1,
            MaxCapacity = 200,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dynamodbTarget, err := appautoscaling.NewTarget(ctx, "dynamodbTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(100),
			MinCapacity:       pulumi.Int(5),
			ResourceId:        pulumi.String("table/tableName"),
			ScalableDimension: pulumi.String("dynamodb:table:ReadCapacityUnits"),
			ServiceNamespace:  pulumi.String("dynamodb"),
		})
		if err != nil {
			return err
		}
		_, err = appautoscaling.NewScheduledAction(ctx, "dynamodbScheduledAction", &appautoscaling.ScheduledActionArgs{
			ServiceNamespace:  dynamodbTarget.ServiceNamespace,
			ResourceId:        dynamodbTarget.ResourceId,
			ScalableDimension: dynamodbTarget.ScalableDimension,
			Schedule:          pulumi.String("at(2006-01-02T15:04:05)"),
			ScalableTargetAction: &appautoscaling.ScheduledActionScalableTargetActionArgs{
				MinCapacity: pulumi.Int(1),
				MaxCapacity: pulumi.Int(200),
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
import com.pulumi.aws.appautoscaling.Target;
import com.pulumi.aws.appautoscaling.TargetArgs;
import com.pulumi.aws.appautoscaling.ScheduledAction;
import com.pulumi.aws.appautoscaling.ScheduledActionArgs;
import com.pulumi.aws.appautoscaling.inputs.ScheduledActionScalableTargetActionArgs;
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
        var dynamodbTarget = new Target("dynamodbTarget", TargetArgs.builder()        
            .maxCapacity(100)
            .minCapacity(5)
            .resourceId("table/tableName")
            .scalableDimension("dynamodb:table:ReadCapacityUnits")
            .serviceNamespace("dynamodb")
            .build());

        var dynamodbScheduledAction = new ScheduledAction("dynamodbScheduledAction", ScheduledActionArgs.builder()        
            .serviceNamespace(dynamodbTarget.serviceNamespace())
            .resourceId(dynamodbTarget.resourceId())
            .scalableDimension(dynamodbTarget.scalableDimension())
            .schedule("at(2006-01-02T15:04:05)")
            .scalableTargetAction(ScheduledActionScalableTargetActionArgs.builder()
                .minCapacity(1)
                .maxCapacity(200)
                .build())
            .build());

    }
}
```
```yaml
resources:
  dynamodbTarget:
    type: aws:appautoscaling:Target
    properties:
      maxCapacity: 100
      minCapacity: 5
      resourceId: table/tableName
      scalableDimension: dynamodb:table:ReadCapacityUnits
      serviceNamespace: dynamodb
  dynamodbScheduledAction:
    type: aws:appautoscaling:ScheduledAction
    properties:
      serviceNamespace: ${dynamodbTarget.serviceNamespace}
      resourceId: ${dynamodbTarget.resourceId}
      scalableDimension: ${dynamodbTarget.scalableDimension}
      schedule: at(2006-01-02T15:04:05)
      scalableTargetAction:
        minCapacity: 1
        maxCapacity: 200
```
{{% /example %}}
{{% example %}}
### ECS Service Autoscaling

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecsTarget = new aws.appautoscaling.Target("ecsTarget", {
    maxCapacity: 4,
    minCapacity: 1,
    resourceId: "service/clusterName/serviceName",
    scalableDimension: "ecs:service:DesiredCount",
    serviceNamespace: "ecs",
});
const ecsScheduledAction = new aws.appautoscaling.ScheduledAction("ecsScheduledAction", {
    serviceNamespace: ecsTarget.serviceNamespace,
    resourceId: ecsTarget.resourceId,
    scalableDimension: ecsTarget.scalableDimension,
    schedule: "at(2006-01-02T15:04:05)",
    scalableTargetAction: {
        minCapacity: 1,
        maxCapacity: 10,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

ecs_target = aws.appautoscaling.Target("ecsTarget",
    max_capacity=4,
    min_capacity=1,
    resource_id="service/clusterName/serviceName",
    scalable_dimension="ecs:service:DesiredCount",
    service_namespace="ecs")
ecs_scheduled_action = aws.appautoscaling.ScheduledAction("ecsScheduledAction",
    service_namespace=ecs_target.service_namespace,
    resource_id=ecs_target.resource_id,
    scalable_dimension=ecs_target.scalable_dimension,
    schedule="at(2006-01-02T15:04:05)",
    scalable_target_action=aws.appautoscaling.ScheduledActionScalableTargetActionArgs(
        min_capacity=1,
        max_capacity=10,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ecsTarget = new Aws.AppAutoScaling.Target("ecsTarget", new()
    {
        MaxCapacity = 4,
        MinCapacity = 1,
        ResourceId = "service/clusterName/serviceName",
        ScalableDimension = "ecs:service:DesiredCount",
        ServiceNamespace = "ecs",
    });

    var ecsScheduledAction = new Aws.AppAutoScaling.ScheduledAction("ecsScheduledAction", new()
    {
        ServiceNamespace = ecsTarget.ServiceNamespace,
        ResourceId = ecsTarget.ResourceId,
        ScalableDimension = ecsTarget.ScalableDimension,
        Schedule = "at(2006-01-02T15:04:05)",
        ScalableTargetAction = new Aws.AppAutoScaling.Inputs.ScheduledActionScalableTargetActionArgs
        {
            MinCapacity = 1,
            MaxCapacity = 10,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ecsTarget, err := appautoscaling.NewTarget(ctx, "ecsTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(4),
			MinCapacity:       pulumi.Int(1),
			ResourceId:        pulumi.String("service/clusterName/serviceName"),
			ScalableDimension: pulumi.String("ecs:service:DesiredCount"),
			ServiceNamespace:  pulumi.String("ecs"),
		})
		if err != nil {
			return err
		}
		_, err = appautoscaling.NewScheduledAction(ctx, "ecsScheduledAction", &appautoscaling.ScheduledActionArgs{
			ServiceNamespace:  ecsTarget.ServiceNamespace,
			ResourceId:        ecsTarget.ResourceId,
			ScalableDimension: ecsTarget.ScalableDimension,
			Schedule:          pulumi.String("at(2006-01-02T15:04:05)"),
			ScalableTargetAction: &appautoscaling.ScheduledActionScalableTargetActionArgs{
				MinCapacity: pulumi.Int(1),
				MaxCapacity: pulumi.Int(10),
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
import com.pulumi.aws.appautoscaling.Target;
import com.pulumi.aws.appautoscaling.TargetArgs;
import com.pulumi.aws.appautoscaling.ScheduledAction;
import com.pulumi.aws.appautoscaling.ScheduledActionArgs;
import com.pulumi.aws.appautoscaling.inputs.ScheduledActionScalableTargetActionArgs;
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
        var ecsTarget = new Target("ecsTarget", TargetArgs.builder()        
            .maxCapacity(4)
            .minCapacity(1)
            .resourceId("service/clusterName/serviceName")
            .scalableDimension("ecs:service:DesiredCount")
            .serviceNamespace("ecs")
            .build());

        var ecsScheduledAction = new ScheduledAction("ecsScheduledAction", ScheduledActionArgs.builder()        
            .serviceNamespace(ecsTarget.serviceNamespace())
            .resourceId(ecsTarget.resourceId())
            .scalableDimension(ecsTarget.scalableDimension())
            .schedule("at(2006-01-02T15:04:05)")
            .scalableTargetAction(ScheduledActionScalableTargetActionArgs.builder()
                .minCapacity(1)
                .maxCapacity(10)
                .build())
            .build());

    }
}
```
```yaml
resources:
  ecsTarget:
    type: aws:appautoscaling:Target
    properties:
      maxCapacity: 4
      minCapacity: 1
      resourceId: service/clusterName/serviceName
      scalableDimension: ecs:service:DesiredCount
      serviceNamespace: ecs
  ecsScheduledAction:
    type: aws:appautoscaling:ScheduledAction
    properties:
      serviceNamespace: ${ecsTarget.serviceNamespace}
      resourceId: ${ecsTarget.resourceId}
      scalableDimension: ${ecsTarget.scalableDimension}
      schedule: at(2006-01-02T15:04:05)
      scalableTargetAction:
        minCapacity: 1
        maxCapacity: 10
```
{{% /example %}}
{{% /examples %}}