Provides an Application AutoScaling Policy resource.

{{% examples %}}
## Example Usage
{{% example %}}
### DynamoDB Table Autoscaling

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dynamodbTableReadTarget = new aws.appautoscaling.Target("dynamodbTableReadTarget", {
    maxCapacity: 100,
    minCapacity: 5,
    resourceId: "table/tableName",
    scalableDimension: "dynamodb:table:ReadCapacityUnits",
    serviceNamespace: "dynamodb",
});
const dynamodbTableReadPolicy = new aws.appautoscaling.Policy("dynamodbTableReadPolicy", {
    policyType: "TargetTrackingScaling",
    resourceId: dynamodbTableReadTarget.resourceId,
    scalableDimension: dynamodbTableReadTarget.scalableDimension,
    serviceNamespace: dynamodbTableReadTarget.serviceNamespace,
    targetTrackingScalingPolicyConfiguration: {
        predefinedMetricSpecification: {
            predefinedMetricType: "DynamoDBReadCapacityUtilization",
        },
        targetValue: 70,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

dynamodb_table_read_target = aws.appautoscaling.Target("dynamodbTableReadTarget",
    max_capacity=100,
    min_capacity=5,
    resource_id="table/tableName",
    scalable_dimension="dynamodb:table:ReadCapacityUnits",
    service_namespace="dynamodb")
dynamodb_table_read_policy = aws.appautoscaling.Policy("dynamodbTableReadPolicy",
    policy_type="TargetTrackingScaling",
    resource_id=dynamodb_table_read_target.resource_id,
    scalable_dimension=dynamodb_table_read_target.scalable_dimension,
    service_namespace=dynamodb_table_read_target.service_namespace,
    target_tracking_scaling_policy_configuration=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs(
        predefined_metric_specification=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs(
            predefined_metric_type="DynamoDBReadCapacityUtilization",
        ),
        target_value=70,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var dynamodbTableReadTarget = new Aws.AppAutoScaling.Target("dynamodbTableReadTarget", new()
    {
        MaxCapacity = 100,
        MinCapacity = 5,
        ResourceId = "table/tableName",
        ScalableDimension = "dynamodb:table:ReadCapacityUnits",
        ServiceNamespace = "dynamodb",
    });

    var dynamodbTableReadPolicy = new Aws.AppAutoScaling.Policy("dynamodbTableReadPolicy", new()
    {
        PolicyType = "TargetTrackingScaling",
        ResourceId = dynamodbTableReadTarget.ResourceId,
        ScalableDimension = dynamodbTableReadTarget.ScalableDimension,
        ServiceNamespace = dynamodbTableReadTarget.ServiceNamespace,
        TargetTrackingScalingPolicyConfiguration = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationArgs
        {
            PredefinedMetricSpecification = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs
            {
                PredefinedMetricType = "DynamoDBReadCapacityUtilization",
            },
            TargetValue = 70,
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
		dynamodbTableReadTarget, err := appautoscaling.NewTarget(ctx, "dynamodbTableReadTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(100),
			MinCapacity:       pulumi.Int(5),
			ResourceId:        pulumi.String("table/tableName"),
			ScalableDimension: pulumi.String("dynamodb:table:ReadCapacityUnits"),
			ServiceNamespace:  pulumi.String("dynamodb"),
		})
		if err != nil {
			return err
		}
		_, err = appautoscaling.NewPolicy(ctx, "dynamodbTableReadPolicy", &appautoscaling.PolicyArgs{
			PolicyType:        pulumi.String("TargetTrackingScaling"),
			ResourceId:        dynamodbTableReadTarget.ResourceId,
			ScalableDimension: dynamodbTableReadTarget.ScalableDimension,
			ServiceNamespace:  dynamodbTableReadTarget.ServiceNamespace,
			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
				PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
					PredefinedMetricType: pulumi.String("DynamoDBReadCapacityUtilization"),
				},
				TargetValue: pulumi.Float64(70),
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
import com.pulumi.aws.appautoscaling.Policy;
import com.pulumi.aws.appautoscaling.PolicyArgs;
import com.pulumi.aws.appautoscaling.inputs.PolicyTargetTrackingScalingPolicyConfigurationArgs;
import com.pulumi.aws.appautoscaling.inputs.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs;
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
        var dynamodbTableReadTarget = new Target("dynamodbTableReadTarget", TargetArgs.builder()        
            .maxCapacity(100)
            .minCapacity(5)
            .resourceId("table/tableName")
            .scalableDimension("dynamodb:table:ReadCapacityUnits")
            .serviceNamespace("dynamodb")
            .build());

        var dynamodbTableReadPolicy = new Policy("dynamodbTableReadPolicy", PolicyArgs.builder()        
            .policyType("TargetTrackingScaling")
            .resourceId(dynamodbTableReadTarget.resourceId())
            .scalableDimension(dynamodbTableReadTarget.scalableDimension())
            .serviceNamespace(dynamodbTableReadTarget.serviceNamespace())
            .targetTrackingScalingPolicyConfiguration(PolicyTargetTrackingScalingPolicyConfigurationArgs.builder()
                .predefinedMetricSpecification(PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs.builder()
                    .predefinedMetricType("DynamoDBReadCapacityUtilization")
                    .build())
                .targetValue(70)
                .build())
            .build());

    }
}
```
```yaml
resources:
  dynamodbTableReadTarget:
    type: aws:appautoscaling:Target
    properties:
      maxCapacity: 100
      minCapacity: 5
      resourceId: table/tableName
      scalableDimension: dynamodb:table:ReadCapacityUnits
      serviceNamespace: dynamodb
  dynamodbTableReadPolicy:
    type: aws:appautoscaling:Policy
    properties:
      policyType: TargetTrackingScaling
      resourceId: ${dynamodbTableReadTarget.resourceId}
      scalableDimension: ${dynamodbTableReadTarget.scalableDimension}
      serviceNamespace: ${dynamodbTableReadTarget.serviceNamespace}
      targetTrackingScalingPolicyConfiguration:
        predefinedMetricSpecification:
          predefinedMetricType: DynamoDBReadCapacityUtilization
        targetValue: 70
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
const ecsPolicy = new aws.appautoscaling.Policy("ecsPolicy", {
    policyType: "StepScaling",
    resourceId: ecsTarget.resourceId,
    scalableDimension: ecsTarget.scalableDimension,
    serviceNamespace: ecsTarget.serviceNamespace,
    stepScalingPolicyConfiguration: {
        adjustmentType: "ChangeInCapacity",
        cooldown: 60,
        metricAggregationType: "Maximum",
        stepAdjustments: [{
            metricIntervalUpperBound: "0",
            scalingAdjustment: -1,
        }],
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
ecs_policy = aws.appautoscaling.Policy("ecsPolicy",
    policy_type="StepScaling",
    resource_id=ecs_target.resource_id,
    scalable_dimension=ecs_target.scalable_dimension,
    service_namespace=ecs_target.service_namespace,
    step_scaling_policy_configuration=aws.appautoscaling.PolicyStepScalingPolicyConfigurationArgs(
        adjustment_type="ChangeInCapacity",
        cooldown=60,
        metric_aggregation_type="Maximum",
        step_adjustments=[aws.appautoscaling.PolicyStepScalingPolicyConfigurationStepAdjustmentArgs(
            metric_interval_upper_bound="0",
            scaling_adjustment=-1,
        )],
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

    var ecsPolicy = new Aws.AppAutoScaling.Policy("ecsPolicy", new()
    {
        PolicyType = "StepScaling",
        ResourceId = ecsTarget.ResourceId,
        ScalableDimension = ecsTarget.ScalableDimension,
        ServiceNamespace = ecsTarget.ServiceNamespace,
        StepScalingPolicyConfiguration = new Aws.AppAutoScaling.Inputs.PolicyStepScalingPolicyConfigurationArgs
        {
            AdjustmentType = "ChangeInCapacity",
            Cooldown = 60,
            MetricAggregationType = "Maximum",
            StepAdjustments = new[]
            {
                new Aws.AppAutoScaling.Inputs.PolicyStepScalingPolicyConfigurationStepAdjustmentArgs
                {
                    MetricIntervalUpperBound = "0",
                    ScalingAdjustment = -1,
                },
            },
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
		_, err = appautoscaling.NewPolicy(ctx, "ecsPolicy", &appautoscaling.PolicyArgs{
			PolicyType:        pulumi.String("StepScaling"),
			ResourceId:        ecsTarget.ResourceId,
			ScalableDimension: ecsTarget.ScalableDimension,
			ServiceNamespace:  ecsTarget.ServiceNamespace,
			StepScalingPolicyConfiguration: &appautoscaling.PolicyStepScalingPolicyConfigurationArgs{
				AdjustmentType:        pulumi.String("ChangeInCapacity"),
				Cooldown:              pulumi.Int(60),
				MetricAggregationType: pulumi.String("Maximum"),
				StepAdjustments: appautoscaling.PolicyStepScalingPolicyConfigurationStepAdjustmentArray{
					&appautoscaling.PolicyStepScalingPolicyConfigurationStepAdjustmentArgs{
						MetricIntervalUpperBound: pulumi.String("0"),
						ScalingAdjustment:        -1,
					},
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
{{% /example %}}
{{% example %}}
### Preserve desired count when updating an autoscaled ECS Service

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecsService = new aws.ecs.Service("ecsService", {
    cluster: "clusterName",
    taskDefinition: "taskDefinitionFamily:1",
    desiredCount: 2,
});
```
```python
import pulumi
import pulumi_aws as aws

ecs_service = aws.ecs.Service("ecsService",
    cluster="clusterName",
    task_definition="taskDefinitionFamily:1",
    desired_count=2)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ecsService = new Aws.Ecs.Service("ecsService", new()
    {
        Cluster = "clusterName",
        TaskDefinition = "taskDefinitionFamily:1",
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
		_, err := ecs.NewService(ctx, "ecsService", &ecs.ServiceArgs{
			Cluster:        pulumi.String("clusterName"),
			TaskDefinition: pulumi.String("taskDefinitionFamily:1"),
			DesiredCount:   pulumi.Int(2),
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
        var ecsService = new Service("ecsService", ServiceArgs.builder()        
            .cluster("clusterName")
            .taskDefinition("taskDefinitionFamily:1")
            .desiredCount(2)
            .build());

    }
}
```
```yaml
resources:
  ecsService:
    type: aws:ecs:Service
    properties:
      cluster: clusterName
      taskDefinition: taskDefinitionFamily:1
      desiredCount: 2
```
{{% /example %}}
{{% example %}}
### Aurora Read Replica Autoscaling

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const replicasTarget = new aws.appautoscaling.Target("replicasTarget", {
    serviceNamespace: "rds",
    scalableDimension: "rds:cluster:ReadReplicaCount",
    resourceId: `cluster:${aws_rds_cluster.example.id}`,
    minCapacity: 1,
    maxCapacity: 15,
});
const replicasPolicy = new aws.appautoscaling.Policy("replicasPolicy", {
    serviceNamespace: replicasTarget.serviceNamespace,
    scalableDimension: replicasTarget.scalableDimension,
    resourceId: replicasTarget.resourceId,
    policyType: "TargetTrackingScaling",
    targetTrackingScalingPolicyConfiguration: {
        predefinedMetricSpecification: {
            predefinedMetricType: "RDSReaderAverageCPUUtilization",
        },
        targetValue: 75,
        scaleInCooldown: 300,
        scaleOutCooldown: 300,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

replicas_target = aws.appautoscaling.Target("replicasTarget",
    service_namespace="rds",
    scalable_dimension="rds:cluster:ReadReplicaCount",
    resource_id=f"cluster:{aws_rds_cluster['example']['id']}",
    min_capacity=1,
    max_capacity=15)
replicas_policy = aws.appautoscaling.Policy("replicasPolicy",
    service_namespace=replicas_target.service_namespace,
    scalable_dimension=replicas_target.scalable_dimension,
    resource_id=replicas_target.resource_id,
    policy_type="TargetTrackingScaling",
    target_tracking_scaling_policy_configuration=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs(
        predefined_metric_specification=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs(
            predefined_metric_type="RDSReaderAverageCPUUtilization",
        ),
        target_value=75,
        scale_in_cooldown=300,
        scale_out_cooldown=300,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var replicasTarget = new Aws.AppAutoScaling.Target("replicasTarget", new()
    {
        ServiceNamespace = "rds",
        ScalableDimension = "rds:cluster:ReadReplicaCount",
        ResourceId = $"cluster:{aws_rds_cluster.Example.Id}",
        MinCapacity = 1,
        MaxCapacity = 15,
    });

    var replicasPolicy = new Aws.AppAutoScaling.Policy("replicasPolicy", new()
    {
        ServiceNamespace = replicasTarget.ServiceNamespace,
        ScalableDimension = replicasTarget.ScalableDimension,
        ResourceId = replicasTarget.ResourceId,
        PolicyType = "TargetTrackingScaling",
        TargetTrackingScalingPolicyConfiguration = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationArgs
        {
            PredefinedMetricSpecification = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs
            {
                PredefinedMetricType = "RDSReaderAverageCPUUtilization",
            },
            TargetValue = 75,
            ScaleInCooldown = 300,
            ScaleOutCooldown = 300,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		replicasTarget, err := appautoscaling.NewTarget(ctx, "replicasTarget", &appautoscaling.TargetArgs{
			ServiceNamespace:  pulumi.String("rds"),
			ScalableDimension: pulumi.String("rds:cluster:ReadReplicaCount"),
			ResourceId:        pulumi.String(fmt.Sprintf("cluster:%v", aws_rds_cluster.Example.Id)),
			MinCapacity:       pulumi.Int(1),
			MaxCapacity:       pulumi.Int(15),
		})
		if err != nil {
			return err
		}
		_, err = appautoscaling.NewPolicy(ctx, "replicasPolicy", &appautoscaling.PolicyArgs{
			ServiceNamespace:  replicasTarget.ServiceNamespace,
			ScalableDimension: replicasTarget.ScalableDimension,
			ResourceId:        replicasTarget.ResourceId,
			PolicyType:        pulumi.String("TargetTrackingScaling"),
			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
				PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
					PredefinedMetricType: pulumi.String("RDSReaderAverageCPUUtilization"),
				},
				TargetValue:      pulumi.Float64(75),
				ScaleInCooldown:  pulumi.Int(300),
				ScaleOutCooldown: pulumi.Int(300),
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
import com.pulumi.aws.appautoscaling.Policy;
import com.pulumi.aws.appautoscaling.PolicyArgs;
import com.pulumi.aws.appautoscaling.inputs.PolicyTargetTrackingScalingPolicyConfigurationArgs;
import com.pulumi.aws.appautoscaling.inputs.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs;
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
        var replicasTarget = new Target("replicasTarget", TargetArgs.builder()        
            .serviceNamespace("rds")
            .scalableDimension("rds:cluster:ReadReplicaCount")
            .resourceId(String.format("cluster:%s", aws_rds_cluster.example().id()))
            .minCapacity(1)
            .maxCapacity(15)
            .build());

        var replicasPolicy = new Policy("replicasPolicy", PolicyArgs.builder()        
            .serviceNamespace(replicasTarget.serviceNamespace())
            .scalableDimension(replicasTarget.scalableDimension())
            .resourceId(replicasTarget.resourceId())
            .policyType("TargetTrackingScaling")
            .targetTrackingScalingPolicyConfiguration(PolicyTargetTrackingScalingPolicyConfigurationArgs.builder()
                .predefinedMetricSpecification(PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs.builder()
                    .predefinedMetricType("RDSReaderAverageCPUUtilization")
                    .build())
                .targetValue(75)
                .scaleInCooldown(300)
                .scaleOutCooldown(300)
                .build())
            .build());

    }
}
```
```yaml
resources:
  replicasTarget:
    type: aws:appautoscaling:Target
    properties:
      serviceNamespace: rds
      scalableDimension: rds:cluster:ReadReplicaCount
      resourceId: cluster:${aws_rds_cluster.example.id}
      minCapacity: 1
      maxCapacity: 15
  replicasPolicy:
    type: aws:appautoscaling:Policy
    properties:
      serviceNamespace: ${replicasTarget.serviceNamespace}
      scalableDimension: ${replicasTarget.scalableDimension}
      resourceId: ${replicasTarget.resourceId}
      policyType: TargetTrackingScaling
      targetTrackingScalingPolicyConfiguration:
        predefinedMetricSpecification:
          predefinedMetricType: RDSReaderAverageCPUUtilization
        targetValue: 75
        scaleInCooldown: 300
        scaleOutCooldown: 300
```
{{% /example %}}
{{% example %}}
### Create target tracking scaling policy using metric math

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
const example = new aws.appautoscaling.Policy("example", {
    policyType: "TargetTrackingScaling",
    resourceId: ecsTarget.resourceId,
    scalableDimension: ecsTarget.scalableDimension,
    serviceNamespace: ecsTarget.serviceNamespace,
    targetTrackingScalingPolicyConfiguration: {
        targetValue: 100,
        customizedMetricSpecification: {
            metrics: [
                {
                    label: "Get the queue size (the number of messages waiting to be processed)",
                    id: "m1",
                    metricStat: {
                        metric: {
                            metricName: "ApproximateNumberOfMessagesVisible",
                            namespace: "AWS/SQS",
                            dimensions: [{
                                name: "QueueName",
                                value: "my-queue",
                            }],
                        },
                        stat: "Sum",
                    },
                    returnData: false,
                },
                {
                    label: "Get the ECS running task count (the number of currently running tasks)",
                    id: "m2",
                    metricStat: {
                        metric: {
                            metricName: "RunningTaskCount",
                            namespace: "ECS/ContainerInsights",
                            dimensions: [
                                {
                                    name: "ClusterName",
                                    value: "default",
                                },
                                {
                                    name: "ServiceName",
                                    value: "web-app",
                                },
                            ],
                        },
                        stat: "Average",
                    },
                    returnData: false,
                },
                {
                    label: "Calculate the backlog per instance",
                    id: "e1",
                    expression: "m1 / m2",
                    returnData: true,
                },
            ],
        },
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
example = aws.appautoscaling.Policy("example",
    policy_type="TargetTrackingScaling",
    resource_id=ecs_target.resource_id,
    scalable_dimension=ecs_target.scalable_dimension,
    service_namespace=ecs_target.service_namespace,
    target_tracking_scaling_policy_configuration=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs(
        target_value=100,
        customized_metric_specification=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs(
            metrics=[
                aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs(
                    label="Get the queue size (the number of messages waiting to be processed)",
                    id="m1",
                    metric_stat=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatArgs(
                        metric=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs(
                            metric_name="ApproximateNumberOfMessagesVisible",
                            namespace="AWS/SQS",
                            dimensions=[aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs(
                                name="QueueName",
                                value="my-queue",
                            )],
                        ),
                        stat="Sum",
                    ),
                    return_data=False,
                ),
                aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs(
                    label="Get the ECS running task count (the number of currently running tasks)",
                    id="m2",
                    metric_stat=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatArgs(
                        metric=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs(
                            metric_name="RunningTaskCount",
                            namespace="ECS/ContainerInsights",
                            dimensions=[
                                aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs(
                                    name="ClusterName",
                                    value="default",
                                ),
                                aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs(
                                    name="ServiceName",
                                    value="web-app",
                                ),
                            ],
                        ),
                        stat="Average",
                    ),
                    return_data=False,
                ),
                aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs(
                    label="Calculate the backlog per instance",
                    id="e1",
                    expression="m1 / m2",
                    return_data=True,
                ),
            ],
        ),
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

    var example = new Aws.AppAutoScaling.Policy("example", new()
    {
        PolicyType = "TargetTrackingScaling",
        ResourceId = ecsTarget.ResourceId,
        ScalableDimension = ecsTarget.ScalableDimension,
        ServiceNamespace = ecsTarget.ServiceNamespace,
        TargetTrackingScalingPolicyConfiguration = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationArgs
        {
            TargetValue = 100,
            CustomizedMetricSpecification = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs
            {
                Metrics = new[]
                {
                    new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs
                    {
                        Label = "Get the queue size (the number of messages waiting to be processed)",
                        Id = "m1",
                        MetricStat = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatArgs
                        {
                            Metric = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs
                            {
                                MetricName = "ApproximateNumberOfMessagesVisible",
                                Namespace = "AWS/SQS",
                                Dimensions = new[]
                                {
                                    new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs
                                    {
                                        Name = "QueueName",
                                        Value = "my-queue",
                                    },
                                },
                            },
                            Stat = "Sum",
                        },
                        ReturnData = false,
                    },
                    new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs
                    {
                        Label = "Get the ECS running task count (the number of currently running tasks)",
                        Id = "m2",
                        MetricStat = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatArgs
                        {
                            Metric = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs
                            {
                                MetricName = "RunningTaskCount",
                                Namespace = "ECS/ContainerInsights",
                                Dimensions = new[]
                                {
                                    new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs
                                    {
                                        Name = "ClusterName",
                                        Value = "default",
                                    },
                                    new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs
                                    {
                                        Name = "ServiceName",
                                        Value = "web-app",
                                    },
                                },
                            },
                            Stat = "Average",
                        },
                        ReturnData = false,
                    },
                    new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs
                    {
                        Label = "Calculate the backlog per instance",
                        Id = "e1",
                        Expression = "m1 / m2",
                        ReturnData = true,
                    },
                },
            },
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
		_, err = appautoscaling.NewPolicy(ctx, "example", &appautoscaling.PolicyArgs{
			PolicyType:        pulumi.String("TargetTrackingScaling"),
			ResourceId:        ecsTarget.ResourceId,
			ScalableDimension: ecsTarget.ScalableDimension,
			ServiceNamespace:  ecsTarget.ServiceNamespace,
			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
				TargetValue: pulumi.Float64(100),
				CustomizedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs{
					Metrics: appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArray{
						&appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs{
							Label: pulumi.String("Get the queue size (the number of messages waiting to be processed)"),
							Id:    pulumi.String("m1"),
							MetricStat: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatArgs{
								Metric: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs{
									MetricName: pulumi.String("ApproximateNumberOfMessagesVisible"),
									Namespace:  pulumi.String("AWS/SQS"),
									Dimensions: appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArray{
										&appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs{
											Name:  pulumi.String("QueueName"),
											Value: pulumi.String("my-queue"),
										},
									},
								},
								Stat: pulumi.String("Sum"),
							},
							ReturnData: pulumi.Bool(false),
						},
						&appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs{
							Label: pulumi.String("Get the ECS running task count (the number of currently running tasks)"),
							Id:    pulumi.String("m2"),
							MetricStat: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatArgs{
								Metric: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs{
									MetricName: pulumi.String("RunningTaskCount"),
									Namespace:  pulumi.String("ECS/ContainerInsights"),
									Dimensions: appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArray{
										&appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs{
											Name:  pulumi.String("ClusterName"),
											Value: pulumi.String("default"),
										},
										&appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs{
											Name:  pulumi.String("ServiceName"),
											Value: pulumi.String("web-app"),
										},
									},
								},
								Stat: pulumi.String("Average"),
							},
							ReturnData: pulumi.Bool(false),
						},
						&appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs{
							Label:      pulumi.String("Calculate the backlog per instance"),
							Id:         pulumi.String("e1"),
							Expression: pulumi.String("m1 / m2"),
							ReturnData: pulumi.Bool(true),
						},
					},
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
import com.pulumi.aws.appautoscaling.Target;
import com.pulumi.aws.appautoscaling.TargetArgs;
import com.pulumi.aws.appautoscaling.Policy;
import com.pulumi.aws.appautoscaling.PolicyArgs;
import com.pulumi.aws.appautoscaling.inputs.PolicyTargetTrackingScalingPolicyConfigurationArgs;
import com.pulumi.aws.appautoscaling.inputs.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs;
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

        var example = new Policy("example", PolicyArgs.builder()        
            .policyType("TargetTrackingScaling")
            .resourceId(ecsTarget.resourceId())
            .scalableDimension(ecsTarget.scalableDimension())
            .serviceNamespace(ecsTarget.serviceNamespace())
            .targetTrackingScalingPolicyConfiguration(PolicyTargetTrackingScalingPolicyConfigurationArgs.builder()
                .targetValue(100)
                .customizedMetricSpecification(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs.builder()
                    .metrics(                    
                        PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs.builder()
                            .label("Get the queue size (the number of messages waiting to be processed)")
                            .id("m1")
                            .metricStat(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatArgs.builder()
                                .metric(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs.builder()
                                    .metricName("ApproximateNumberOfMessagesVisible")
                                    .namespace("AWS/SQS")
                                    .dimensions(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs.builder()
                                        .name("QueueName")
                                        .value("my-queue")
                                        .build())
                                    .build())
                                .stat("Sum")
                                .build())
                            .returnData(false)
                            .build(),
                        PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs.builder()
                            .label("Get the ECS running task count (the number of currently running tasks)")
                            .id("m2")
                            .metricStat(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatArgs.builder()
                                .metric(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs.builder()
                                    .metricName("RunningTaskCount")
                                    .namespace("ECS/ContainerInsights")
                                    .dimensions(                                    
                                        PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs.builder()
                                            .name("ClusterName")
                                            .value("default")
                                            .build(),
                                        PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs.builder()
                                            .name("ServiceName")
                                            .value("web-app")
                                            .build())
                                    .build())
                                .stat("Average")
                                .build())
                            .returnData(false)
                            .build(),
                        PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricArgs.builder()
                            .label("Calculate the backlog per instance")
                            .id("e1")
                            .expression("m1 / m2")
                            .returnData(true)
                            .build())
                    .build())
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
  example:
    type: aws:appautoscaling:Policy
    properties:
      policyType: TargetTrackingScaling
      resourceId: ${ecsTarget.resourceId}
      scalableDimension: ${ecsTarget.scalableDimension}
      serviceNamespace: ${ecsTarget.serviceNamespace}
      targetTrackingScalingPolicyConfiguration:
        targetValue: 100
        customizedMetricSpecification:
          metrics:
            - label: Get the queue size (the number of messages waiting to be processed)
              id: m1
              metricStat:
                metric:
                  metricName: ApproximateNumberOfMessagesVisible
                  namespace: AWS/SQS
                  dimensions:
                    - name: QueueName
                      value: my-queue
                stat: Sum
              returnData: false
            - label: Get the ECS running task count (the number of currently running tasks)
              id: m2
              metricStat:
                metric:
                  metricName: RunningTaskCount
                  namespace: ECS/ContainerInsights
                  dimensions:
                    - name: ClusterName
                      value: default
                    - name: ServiceName
                      value: web-app
                stat: Average
              returnData: false
            - label: Calculate the backlog per instance
              id: e1
              expression: m1 / m2
              returnData: true
```
{{% /example %}}
{{% example %}}
### MSK / Kafka Autoscaling

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mskTarget = new aws.appautoscaling.Target("mskTarget", {
    serviceNamespace: "kafka",
    scalableDimension: "kafka:broker-storage:VolumeSize",
    resourceId: aws_msk_cluster.example.arn,
    minCapacity: 1,
    maxCapacity: 8,
});
const targets = new aws.appautoscaling.Policy("targets", {
    serviceNamespace: mskTarget.serviceNamespace,
    scalableDimension: mskTarget.scalableDimension,
    resourceId: mskTarget.resourceId,
    policyType: "TargetTrackingScaling",
    targetTrackingScalingPolicyConfiguration: {
        predefinedMetricSpecification: {
            predefinedMetricType: "KafkaBrokerStorageUtilization",
        },
        targetValue: 55,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

msk_target = aws.appautoscaling.Target("mskTarget",
    service_namespace="kafka",
    scalable_dimension="kafka:broker-storage:VolumeSize",
    resource_id=aws_msk_cluster["example"]["arn"],
    min_capacity=1,
    max_capacity=8)
targets = aws.appautoscaling.Policy("targets",
    service_namespace=msk_target.service_namespace,
    scalable_dimension=msk_target.scalable_dimension,
    resource_id=msk_target.resource_id,
    policy_type="TargetTrackingScaling",
    target_tracking_scaling_policy_configuration=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs(
        predefined_metric_specification=aws.appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs(
            predefined_metric_type="KafkaBrokerStorageUtilization",
        ),
        target_value=55,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mskTarget = new Aws.AppAutoScaling.Target("mskTarget", new()
    {
        ServiceNamespace = "kafka",
        ScalableDimension = "kafka:broker-storage:VolumeSize",
        ResourceId = aws_msk_cluster.Example.Arn,
        MinCapacity = 1,
        MaxCapacity = 8,
    });

    var targets = new Aws.AppAutoScaling.Policy("targets", new()
    {
        ServiceNamespace = mskTarget.ServiceNamespace,
        ScalableDimension = mskTarget.ScalableDimension,
        ResourceId = mskTarget.ResourceId,
        PolicyType = "TargetTrackingScaling",
        TargetTrackingScalingPolicyConfiguration = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationArgs
        {
            PredefinedMetricSpecification = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs
            {
                PredefinedMetricType = "KafkaBrokerStorageUtilization",
            },
            TargetValue = 55,
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
		mskTarget, err := appautoscaling.NewTarget(ctx, "mskTarget", &appautoscaling.TargetArgs{
			ServiceNamespace:  pulumi.String("kafka"),
			ScalableDimension: pulumi.String("kafka:broker-storage:VolumeSize"),
			ResourceId:        pulumi.Any(aws_msk_cluster.Example.Arn),
			MinCapacity:       pulumi.Int(1),
			MaxCapacity:       pulumi.Int(8),
		})
		if err != nil {
			return err
		}
		_, err = appautoscaling.NewPolicy(ctx, "targets", &appautoscaling.PolicyArgs{
			ServiceNamespace:  mskTarget.ServiceNamespace,
			ScalableDimension: mskTarget.ScalableDimension,
			ResourceId:        mskTarget.ResourceId,
			PolicyType:        pulumi.String("TargetTrackingScaling"),
			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
				PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
					PredefinedMetricType: pulumi.String("KafkaBrokerStorageUtilization"),
				},
				TargetValue: pulumi.Float64(55),
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
import com.pulumi.aws.appautoscaling.Policy;
import com.pulumi.aws.appautoscaling.PolicyArgs;
import com.pulumi.aws.appautoscaling.inputs.PolicyTargetTrackingScalingPolicyConfigurationArgs;
import com.pulumi.aws.appautoscaling.inputs.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs;
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
        var mskTarget = new Target("mskTarget", TargetArgs.builder()        
            .serviceNamespace("kafka")
            .scalableDimension("kafka:broker-storage:VolumeSize")
            .resourceId(aws_msk_cluster.example().arn())
            .minCapacity(1)
            .maxCapacity(8)
            .build());

        var targets = new Policy("targets", PolicyArgs.builder()        
            .serviceNamespace(mskTarget.serviceNamespace())
            .scalableDimension(mskTarget.scalableDimension())
            .resourceId(mskTarget.resourceId())
            .policyType("TargetTrackingScaling")
            .targetTrackingScalingPolicyConfiguration(PolicyTargetTrackingScalingPolicyConfigurationArgs.builder()
                .predefinedMetricSpecification(PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs.builder()
                    .predefinedMetricType("KafkaBrokerStorageUtilization")
                    .build())
                .targetValue(55)
                .build())
            .build());

    }
}
```
```yaml
resources:
  mskTarget:
    type: aws:appautoscaling:Target
    properties:
      serviceNamespace: kafka
      scalableDimension: kafka:broker-storage:VolumeSize
      resourceId: ${aws_msk_cluster.example.arn}
      minCapacity: 1
      maxCapacity: 8
  targets:
    type: aws:appautoscaling:Policy
    properties:
      serviceNamespace: ${mskTarget.serviceNamespace}
      scalableDimension: ${mskTarget.scalableDimension}
      resourceId: ${mskTarget.resourceId}
      policyType: TargetTrackingScaling
      targetTrackingScalingPolicyConfiguration:
        predefinedMetricSpecification:
          predefinedMetricType: KafkaBrokerStorageUtilization
        targetValue: 55
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Application AutoScaling Policy using the `service-namespace` , `resource-id`, `scalable-dimension` and `policy-name` separated by `/`. For example:

```sh
 $ pulumi import aws:appautoscaling/policy:Policy test-policy service-namespace/resource-id/scalable-dimension/policy-name
```
 