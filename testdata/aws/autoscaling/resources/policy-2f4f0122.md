Provides an AutoScaling Scaling Policy resource.

> **NOTE:** You may want to omit `desired_capacity` attribute from attached `aws.autoscaling.Group`
when using autoscaling policies. It's good practice to pick either
[manual](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-manual-scaling.html)
or [dynamic](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-scale-based-on-demand.html)
(policy-based) scaling.



{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.autoscaling.Group("bar", {
    availabilityZones: ["us-east-1a"],
    maxSize: 5,
    minSize: 2,
    healthCheckGracePeriod: 300,
    healthCheckType: "ELB",
    forceDelete: true,
    launchConfiguration: aws_launch_configuration.foo.name,
});
const bat = new aws.autoscaling.Policy("bat", {
    scalingAdjustment: 4,
    adjustmentType: "ChangeInCapacity",
    cooldown: 300,
    autoscalingGroupName: bar.name,
});
```
```python
import pulumi
import pulumi_aws as aws

bar = aws.autoscaling.Group("bar",
    availability_zones=["us-east-1a"],
    max_size=5,
    min_size=2,
    health_check_grace_period=300,
    health_check_type="ELB",
    force_delete=True,
    launch_configuration=aws_launch_configuration["foo"]["name"])
bat = aws.autoscaling.Policy("bat",
    scaling_adjustment=4,
    adjustment_type="ChangeInCapacity",
    cooldown=300,
    autoscaling_group_name=bar.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bar = new Aws.AutoScaling.Group("bar", new()
    {
        AvailabilityZones = new[]
        {
            "us-east-1a",
        },
        MaxSize = 5,
        MinSize = 2,
        HealthCheckGracePeriod = 300,
        HealthCheckType = "ELB",
        ForceDelete = true,
        LaunchConfiguration = aws_launch_configuration.Foo.Name,
    });

    var bat = new Aws.AutoScaling.Policy("bat", new()
    {
        ScalingAdjustment = 4,
        AdjustmentType = "ChangeInCapacity",
        Cooldown = 300,
        AutoscalingGroupName = bar.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bar, err := autoscaling.NewGroup(ctx, "bar", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			MaxSize:                pulumi.Int(5),
			MinSize:                pulumi.Int(2),
			HealthCheckGracePeriod: pulumi.Int(300),
			HealthCheckType:        pulumi.String("ELB"),
			ForceDelete:            pulumi.Bool(true),
			LaunchConfiguration:    pulumi.Any(aws_launch_configuration.Foo.Name),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewPolicy(ctx, "bat", &autoscaling.PolicyArgs{
			ScalingAdjustment:    pulumi.Int(4),
			AdjustmentType:       pulumi.String("ChangeInCapacity"),
			Cooldown:             pulumi.Int(300),
			AutoscalingGroupName: bar.Name,
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
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.Policy;
import com.pulumi.aws.autoscaling.PolicyArgs;
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
        var bar = new Group("bar", GroupArgs.builder()        
            .availabilityZones("us-east-1a")
            .maxSize(5)
            .minSize(2)
            .healthCheckGracePeriod(300)
            .healthCheckType("ELB")
            .forceDelete(true)
            .launchConfiguration(aws_launch_configuration.foo().name())
            .build());

        var bat = new Policy("bat", PolicyArgs.builder()        
            .scalingAdjustment(4)
            .adjustmentType("ChangeInCapacity")
            .cooldown(300)
            .autoscalingGroupName(bar.name())
            .build());

    }
}
```
```yaml
resources:
  bat:
    type: aws:autoscaling:Policy
    properties:
      scalingAdjustment: 4
      adjustmentType: ChangeInCapacity
      cooldown: 300
      autoscalingGroupName: ${bar.name}
  bar:
    type: aws:autoscaling:Group
    properties:
      availabilityZones:
        - us-east-1a
      maxSize: 5
      minSize: 2
      healthCheckGracePeriod: 300
      healthCheckType: ELB
      forceDelete: true
      launchConfiguration: ${aws_launch_configuration.foo.name}
```
{{% /example %}}
{{% example %}}
### Create target tracking scaling policy using metric math

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.autoscaling.Policy("example", {
    autoscalingGroupName: "my-test-asg",
    policyType: "TargetTrackingScaling",
    targetTrackingConfiguration: {
        customizedMetricSpecification: {
            metrics: [
                {
                    id: "m1",
                    label: "Get the queue size (the number of messages waiting to be processed)",
                    metricStat: {
                        metric: {
                            dimensions: [{
                                name: "QueueName",
                                value: "my-queue",
                            }],
                            metricName: "ApproximateNumberOfMessagesVisible",
                            namespace: "AWS/SQS",
                        },
                        stat: "Sum",
                    },
                    returnData: false,
                },
                {
                    id: "m2",
                    label: "Get the group size (the number of InService instances)",
                    metricStat: {
                        metric: {
                            dimensions: [{
                                name: "AutoScalingGroupName",
                                value: "my-asg",
                            }],
                            metricName: "GroupInServiceInstances",
                            namespace: "AWS/AutoScaling",
                        },
                        stat: "Average",
                    },
                    returnData: false,
                },
                {
                    expression: "m1 / m2",
                    id: "e1",
                    label: "Calculate the backlog per instance",
                    returnData: true,
                },
            ],
        },
        targetValue: 100,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.autoscaling.Policy("example",
    autoscaling_group_name="my-test-asg",
    policy_type="TargetTrackingScaling",
    target_tracking_configuration=aws.autoscaling.PolicyTargetTrackingConfigurationArgs(
        customized_metric_specification=aws.autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationArgs(
            metrics=[
                aws.autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs(
                    id="m1",
                    label="Get the queue size (the number of messages waiting to be processed)",
                    metric_stat=aws.autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatArgs(
                        metric=aws.autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs(
                            dimensions=[aws.autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs(
                                name="QueueName",
                                value="my-queue",
                            )],
                            metric_name="ApproximateNumberOfMessagesVisible",
                            namespace="AWS/SQS",
                        ),
                        stat="Sum",
                    ),
                    return_data=False,
                ),
                aws.autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs(
                    id="m2",
                    label="Get the group size (the number of InService instances)",
                    metric_stat=aws.autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatArgs(
                        metric=aws.autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs(
                            dimensions=[aws.autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs(
                                name="AutoScalingGroupName",
                                value="my-asg",
                            )],
                            metric_name="GroupInServiceInstances",
                            namespace="AWS/AutoScaling",
                        ),
                        stat="Average",
                    ),
                    return_data=False,
                ),
                aws.autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs(
                    expression="m1 / m2",
                    id="e1",
                    label="Calculate the backlog per instance",
                    return_data=True,
                ),
            ],
        ),
        target_value=100,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AutoScaling.Policy("example", new()
    {
        AutoscalingGroupName = "my-test-asg",
        PolicyType = "TargetTrackingScaling",
        TargetTrackingConfiguration = new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationArgs
        {
            CustomizedMetricSpecification = new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationArgs
            {
                Metrics = new[]
                {
                    new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs
                    {
                        Id = "m1",
                        Label = "Get the queue size (the number of messages waiting to be processed)",
                        MetricStat = new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatArgs
                        {
                            Metric = new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs
                            {
                                Dimensions = new[]
                                {
                                    new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs
                                    {
                                        Name = "QueueName",
                                        Value = "my-queue",
                                    },
                                },
                                MetricName = "ApproximateNumberOfMessagesVisible",
                                Namespace = "AWS/SQS",
                            },
                            Stat = "Sum",
                        },
                        ReturnData = false,
                    },
                    new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs
                    {
                        Id = "m2",
                        Label = "Get the group size (the number of InService instances)",
                        MetricStat = new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatArgs
                        {
                            Metric = new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs
                            {
                                Dimensions = new[]
                                {
                                    new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs
                                    {
                                        Name = "AutoScalingGroupName",
                                        Value = "my-asg",
                                    },
                                },
                                MetricName = "GroupInServiceInstances",
                                Namespace = "AWS/AutoScaling",
                            },
                            Stat = "Average",
                        },
                        ReturnData = false,
                    },
                    new Aws.AutoScaling.Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs
                    {
                        Expression = "m1 / m2",
                        Id = "e1",
                        Label = "Calculate the backlog per instance",
                        ReturnData = true,
                    },
                },
            },
            TargetValue = 100,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := autoscaling.NewPolicy(ctx, "example", &autoscaling.PolicyArgs{
			AutoscalingGroupName: pulumi.String("my-test-asg"),
			PolicyType:           pulumi.String("TargetTrackingScaling"),
			TargetTrackingConfiguration: &autoscaling.PolicyTargetTrackingConfigurationArgs{
				CustomizedMetricSpecification: &autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationArgs{
					Metrics: autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArray{
						&autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs{
							Id:    pulumi.String("m1"),
							Label: pulumi.String("Get the queue size (the number of messages waiting to be processed)"),
							MetricStat: &autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatArgs{
								Metric: &autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs{
									Dimensions: autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArray{
										&autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs{
											Name:  pulumi.String("QueueName"),
											Value: pulumi.String("my-queue"),
										},
									},
									MetricName: pulumi.String("ApproximateNumberOfMessagesVisible"),
									Namespace:  pulumi.String("AWS/SQS"),
								},
								Stat: pulumi.String("Sum"),
							},
							ReturnData: pulumi.Bool(false),
						},
						&autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs{
							Id:    pulumi.String("m2"),
							Label: pulumi.String("Get the group size (the number of InService instances)"),
							MetricStat: &autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatArgs{
								Metric: &autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs{
									Dimensions: autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArray{
										&autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs{
											Name:  pulumi.String("AutoScalingGroupName"),
											Value: pulumi.String("my-asg"),
										},
									},
									MetricName: pulumi.String("GroupInServiceInstances"),
									Namespace:  pulumi.String("AWS/AutoScaling"),
								},
								Stat: pulumi.String("Average"),
							},
							ReturnData: pulumi.Bool(false),
						},
						&autoscaling.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs{
							Expression: pulumi.String("m1 / m2"),
							Id:         pulumi.String("e1"),
							Label:      pulumi.String("Calculate the backlog per instance"),
							ReturnData: pulumi.Bool(true),
						},
					},
				},
				TargetValue: pulumi.Float64(100),
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
import com.pulumi.aws.autoscaling.Policy;
import com.pulumi.aws.autoscaling.PolicyArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyTargetTrackingConfigurationArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationArgs;
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
        var example = new Policy("example", PolicyArgs.builder()        
            .autoscalingGroupName("my-test-asg")
            .policyType("TargetTrackingScaling")
            .targetTrackingConfiguration(PolicyTargetTrackingConfigurationArgs.builder()
                .customizedMetricSpecification(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationArgs.builder()
                    .metrics(                    
                        PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs.builder()
                            .id("m1")
                            .label("Get the queue size (the number of messages waiting to be processed)")
                            .metricStat(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatArgs.builder()
                                .metric(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs.builder()
                                    .dimensions(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs.builder()
                                        .name("QueueName")
                                        .value("my-queue")
                                        .build())
                                    .metricName("ApproximateNumberOfMessagesVisible")
                                    .namespace("AWS/SQS")
                                    .build())
                                .stat("Sum")
                                .build())
                            .returnData(false)
                            .build(),
                        PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs.builder()
                            .id("m2")
                            .label("Get the group size (the number of InService instances)")
                            .metricStat(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatArgs.builder()
                                .metric(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricArgs.builder()
                                    .dimensions(PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimensionArgs.builder()
                                        .name("AutoScalingGroupName")
                                        .value("my-asg")
                                        .build())
                                    .metricName("GroupInServiceInstances")
                                    .namespace("AWS/AutoScaling")
                                    .build())
                                .stat("Average")
                                .build())
                            .returnData(false)
                            .build(),
                        PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricArgs.builder()
                            .expression("m1 / m2")
                            .id("e1")
                            .label("Calculate the backlog per instance")
                            .returnData(true)
                            .build())
                    .build())
                .targetValue(100)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:autoscaling:Policy
    properties:
      autoscalingGroupName: my-test-asg
      policyType: TargetTrackingScaling
      targetTrackingConfiguration:
        customizedMetricSpecification:
          metrics:
            - id: m1
              label: Get the queue size (the number of messages waiting to be processed)
              metricStat:
                metric:
                  dimensions:
                    - name: QueueName
                      value: my-queue
                  metricName: ApproximateNumberOfMessagesVisible
                  namespace: AWS/SQS
                stat: Sum
              returnData: false
            - id: m2
              label: Get the group size (the number of InService instances)
              metricStat:
                metric:
                  dimensions:
                    - name: AutoScalingGroupName
                      value: my-asg
                  metricName: GroupInServiceInstances
                  namespace: AWS/AutoScaling
                stat: Average
              returnData: false
            - expression: m1 / m2
              id: e1
              label: Calculate the backlog per instance
              returnData: true
        targetValue: 100
```
{{% /example %}}
{{% example %}}
### Create predictive scaling policy using customized metrics

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.autoscaling.Policy("example", {
    autoscalingGroupName: "my-test-asg",
    policyType: "PredictiveScaling",
    predictiveScalingConfiguration: {
        metricSpecification: {
            customizedCapacityMetricSpecification: {
                metricDataQueries: [{
                    expression: "SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))",
                    id: "capacity_sum",
                }],
            },
            customizedLoadMetricSpecification: {
                metricDataQueries: [{
                    expression: "SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 3600))",
                    id: "load_sum",
                }],
            },
            customizedScalingMetricSpecification: {
                metricDataQueries: [
                    {
                        expression: "SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))",
                        id: "capacity_sum",
                        returnData: false,
                    },
                    {
                        expression: "SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 300))",
                        id: "load_sum",
                        returnData: false,
                    },
                    {
                        expression: "load_sum / (capacity_sum * PERIOD(capacity_sum) / 60)",
                        id: "weighted_average",
                    },
                ],
            },
            targetValue: 10,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.autoscaling.Policy("example",
    autoscaling_group_name="my-test-asg",
    policy_type="PredictiveScaling",
    predictive_scaling_configuration=aws.autoscaling.PolicyPredictiveScalingConfigurationArgs(
        metric_specification=aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationArgs(
            customized_capacity_metric_specification=aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationArgs(
                metric_data_queries=[aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryArgs(
                    expression="SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))",
                    id="capacity_sum",
                )],
            ),
            customized_load_metric_specification=aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationArgs(
                metric_data_queries=[aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQueryArgs(
                    expression="SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 3600))",
                    id="load_sum",
                )],
            ),
            customized_scaling_metric_specification=aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs(
                metric_data_queries=[
                    aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs(
                        expression="SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))",
                        id="capacity_sum",
                        return_data=False,
                    ),
                    aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs(
                        expression="SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 300))",
                        id="load_sum",
                        return_data=False,
                    ),
                    aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs(
                        expression="load_sum / (capacity_sum * PERIOD(capacity_sum) / 60)",
                        id="weighted_average",
                    ),
                ],
            ),
            target_value=10,
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
    var example = new Aws.AutoScaling.Policy("example", new()
    {
        AutoscalingGroupName = "my-test-asg",
        PolicyType = "PredictiveScaling",
        PredictiveScalingConfiguration = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationArgs
        {
            MetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationArgs
            {
                CustomizedCapacityMetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationArgs
                {
                    MetricDataQueries = new[]
                    {
                        new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryArgs
                        {
                            Expression = "SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))",
                            Id = "capacity_sum",
                        },
                    },
                },
                CustomizedLoadMetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationArgs
                {
                    MetricDataQueries = new[]
                    {
                        new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQueryArgs
                        {
                            Expression = "SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 3600))",
                            Id = "load_sum",
                        },
                    },
                },
                CustomizedScalingMetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs
                {
                    MetricDataQueries = new[]
                    {
                        new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs
                        {
                            Expression = "SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))",
                            Id = "capacity_sum",
                            ReturnData = false,
                        },
                        new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs
                        {
                            Expression = "SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 300))",
                            Id = "load_sum",
                            ReturnData = false,
                        },
                        new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs
                        {
                            Expression = "load_sum / (capacity_sum * PERIOD(capacity_sum) / 60)",
                            Id = "weighted_average",
                        },
                    },
                },
                TargetValue = 10,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := autoscaling.NewPolicy(ctx, "example", &autoscaling.PolicyArgs{
			AutoscalingGroupName: pulumi.String("my-test-asg"),
			PolicyType:           pulumi.String("PredictiveScaling"),
			PredictiveScalingConfiguration: &autoscaling.PolicyPredictiveScalingConfigurationArgs{
				MetricSpecification: &autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationArgs{
					CustomizedCapacityMetricSpecification: &autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationArgs{
						MetricDataQueries: autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryArray{
							&autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryArgs{
								Expression: pulumi.String("SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))"),
								Id:         pulumi.String("capacity_sum"),
							},
						},
					},
					CustomizedLoadMetricSpecification: &autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationArgs{
						MetricDataQueries: autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQueryArray{
							&autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQueryArgs{
								Expression: pulumi.String("SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 3600))"),
								Id:         pulumi.String("load_sum"),
							},
						},
					},
					CustomizedScalingMetricSpecification: &autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs{
						MetricDataQueries: autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArray{
							&autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs{
								Expression: pulumi.String("SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))"),
								Id:         pulumi.String("capacity_sum"),
								ReturnData: pulumi.Bool(false),
							},
							&autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs{
								Expression: pulumi.String("SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 300))"),
								Id:         pulumi.String("load_sum"),
								ReturnData: pulumi.Bool(false),
							},
							&autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs{
								Expression: pulumi.String("load_sum / (capacity_sum * PERIOD(capacity_sum) / 60)"),
								Id:         pulumi.String("weighted_average"),
							},
						},
					},
					TargetValue: pulumi.Float64(10),
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
import com.pulumi.aws.autoscaling.Policy;
import com.pulumi.aws.autoscaling.PolicyArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs;
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
        var example = new Policy("example", PolicyArgs.builder()        
            .autoscalingGroupName("my-test-asg")
            .policyType("PredictiveScaling")
            .predictiveScalingConfiguration(PolicyPredictiveScalingConfigurationArgs.builder()
                .metricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationArgs.builder()
                    .customizedCapacityMetricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationArgs.builder()
                        .metricDataQueries(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryArgs.builder()
                            .expression("SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))")
                            .id("capacity_sum")
                            .build())
                        .build())
                    .customizedLoadMetricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationArgs.builder()
                        .metricDataQueries(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQueryArgs.builder()
                            .expression("SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 3600))")
                            .id("load_sum")
                            .build())
                        .build())
                    .customizedScalingMetricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs.builder()
                        .metricDataQueries(                        
                            PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs.builder()
                                .expression("SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName=\"GroupInServiceIntances\" my-test-asg', 'Average', 300))")
                                .id("capacity_sum")
                                .returnData(false)
                                .build(),
                            PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs.builder()
                                .expression("SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName=\"CPUUtilization\" my-test-asg', 'Sum', 300))")
                                .id("load_sum")
                                .returnData(false)
                                .build(),
                            PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs.builder()
                                .expression("load_sum / (capacity_sum * PERIOD(capacity_sum) / 60)")
                                .id("weighted_average")
                                .build())
                        .build())
                    .targetValue(10)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:autoscaling:Policy
    properties:
      autoscalingGroupName: my-test-asg
      policyType: PredictiveScaling
      predictiveScalingConfiguration:
        metricSpecification:
          customizedCapacityMetricSpecification:
            metricDataQueries:
              - expression: SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName="GroupInServiceIntances" my-test-asg', 'Average', 300))
                id: capacity_sum
          customizedLoadMetricSpecification:
            metricDataQueries:
              - expression: SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName="CPUUtilization" my-test-asg', 'Sum', 3600))
                id: load_sum
          customizedScalingMetricSpecification:
            metricDataQueries:
              - expression: SUM(SEARCH('{AWS/AutoScaling,AutoScalingGroupName} MetricName="GroupInServiceIntances" my-test-asg', 'Average', 300))
                id: capacity_sum
                returnData: false
              - expression: SUM(SEARCH('{AWS/EC2,AutoScalingGroupName} MetricName="CPUUtilization" my-test-asg', 'Sum', 300))
                id: load_sum
                returnData: false
              - expression: load_sum / (capacity_sum * PERIOD(capacity_sum) / 60)
                id: weighted_average
          targetValue: 10
```
{{% /example %}}
{{% example %}}
### Create predictive scaling policy using customized scaling and predefined load metric

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.autoscaling.Policy("example", {
    autoscalingGroupName: "my-test-asg",
    policyType: "PredictiveScaling",
    predictiveScalingConfiguration: {
        metricSpecification: {
            customizedScalingMetricSpecification: {
                metricDataQueries: [{
                    id: "scaling",
                    metricStat: {
                        metric: {
                            dimensions: [{
                                name: "AutoScalingGroupName",
                                value: "my-test-asg",
                            }],
                            metricName: "CPUUtilization",
                            namespace: "AWS/EC2",
                        },
                        stat: "Average",
                    },
                }],
            },
            predefinedLoadMetricSpecification: {
                predefinedMetricType: "ASGTotalCPUUtilization",
                resourceLabel: "app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff",
            },
            targetValue: 10,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.autoscaling.Policy("example",
    autoscaling_group_name="my-test-asg",
    policy_type="PredictiveScaling",
    predictive_scaling_configuration=aws.autoscaling.PolicyPredictiveScalingConfigurationArgs(
        metric_specification=aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationArgs(
            customized_scaling_metric_specification=aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs(
                metric_data_queries=[aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs(
                    id="scaling",
                    metric_stat=aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatArgs(
                        metric=aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricArgs(
                            dimensions=[aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricDimensionArgs(
                                name="AutoScalingGroupName",
                                value="my-test-asg",
                            )],
                            metric_name="CPUUtilization",
                            namespace="AWS/EC2",
                        ),
                        stat="Average",
                    ),
                )],
            ),
            predefined_load_metric_specification=aws.autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecificationArgs(
                predefined_metric_type="ASGTotalCPUUtilization",
                resource_label="app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff",
            ),
            target_value=10,
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
    var example = new Aws.AutoScaling.Policy("example", new()
    {
        AutoscalingGroupName = "my-test-asg",
        PolicyType = "PredictiveScaling",
        PredictiveScalingConfiguration = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationArgs
        {
            MetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationArgs
            {
                CustomizedScalingMetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs
                {
                    MetricDataQueries = new[]
                    {
                        new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs
                        {
                            Id = "scaling",
                            MetricStat = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatArgs
                            {
                                Metric = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricArgs
                                {
                                    Dimensions = new[]
                                    {
                                        new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricDimensionArgs
                                        {
                                            Name = "AutoScalingGroupName",
                                            Value = "my-test-asg",
                                        },
                                    },
                                    MetricName = "CPUUtilization",
                                    Namespace = "AWS/EC2",
                                },
                                Stat = "Average",
                            },
                        },
                    },
                },
                PredefinedLoadMetricSpecification = new Aws.AutoScaling.Inputs.PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecificationArgs
                {
                    PredefinedMetricType = "ASGTotalCPUUtilization",
                    ResourceLabel = "app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff",
                },
                TargetValue = 10,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := autoscaling.NewPolicy(ctx, "example", &autoscaling.PolicyArgs{
			AutoscalingGroupName: pulumi.String("my-test-asg"),
			PolicyType:           pulumi.String("PredictiveScaling"),
			PredictiveScalingConfiguration: &autoscaling.PolicyPredictiveScalingConfigurationArgs{
				MetricSpecification: &autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationArgs{
					CustomizedScalingMetricSpecification: &autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs{
						MetricDataQueries: autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArray{
							&autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs{
								Id: pulumi.String("scaling"),
								MetricStat: &autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatArgs{
									Metric: &autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricArgs{
										Dimensions: autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricDimensionArray{
											&autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricDimensionArgs{
												Name:  pulumi.String("AutoScalingGroupName"),
												Value: pulumi.String("my-test-asg"),
											},
										},
										MetricName: pulumi.String("CPUUtilization"),
										Namespace:  pulumi.String("AWS/EC2"),
									},
									Stat: pulumi.String("Average"),
								},
							},
						},
					},
					PredefinedLoadMetricSpecification: &autoscaling.PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecificationArgs{
						PredefinedMetricType: pulumi.String("ASGTotalCPUUtilization"),
						ResourceLabel:        pulumi.String("app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff"),
					},
					TargetValue: pulumi.Float64(10),
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
import com.pulumi.aws.autoscaling.Policy;
import com.pulumi.aws.autoscaling.PolicyArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs;
import com.pulumi.aws.autoscaling.inputs.PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecificationArgs;
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
        var example = new Policy("example", PolicyArgs.builder()        
            .autoscalingGroupName("my-test-asg")
            .policyType("PredictiveScaling")
            .predictiveScalingConfiguration(PolicyPredictiveScalingConfigurationArgs.builder()
                .metricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationArgs.builder()
                    .customizedScalingMetricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationArgs.builder()
                        .metricDataQueries(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryArgs.builder()
                            .id("scaling")
                            .metricStat(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatArgs.builder()
                                .metric(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricArgs.builder()
                                    .dimensions(PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricDimensionArgs.builder()
                                        .name("AutoScalingGroupName")
                                        .value("my-test-asg")
                                        .build())
                                    .metricName("CPUUtilization")
                                    .namespace("AWS/EC2")
                                    .build())
                                .stat("Average")
                                .build())
                            .build())
                        .build())
                    .predefinedLoadMetricSpecification(PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecificationArgs.builder()
                        .predefinedMetricType("ASGTotalCPUUtilization")
                        .resourceLabel("app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff")
                        .build())
                    .targetValue(10)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:autoscaling:Policy
    properties:
      autoscalingGroupName: my-test-asg
      policyType: PredictiveScaling
      predictiveScalingConfiguration:
        metricSpecification:
          customizedScalingMetricSpecification:
            metricDataQueries:
              - id: scaling
                metricStat:
                  metric:
                    dimensions:
                      - name: AutoScalingGroupName
                        value: my-test-asg
                    metricName: CPUUtilization
                    namespace: AWS/EC2
                  stat: Average
          predefinedLoadMetricSpecification:
            predefinedMetricType: ASGTotalCPUUtilization
            resourceLabel: app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff
          targetValue: 10
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AutoScaling scaling policy using the role autoscaling_group_name and name separated by `/`. For example:

```sh
 $ pulumi import aws:autoscaling/policy:Policy test-policy asg-name/policy-name
```
 