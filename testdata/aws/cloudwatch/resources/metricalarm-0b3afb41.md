Provides a CloudWatch Metric Alarm resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobar = new aws.cloudwatch.MetricAlarm("foobar", {
    alarmDescription: "This metric monitors ec2 cpu utilization",
    comparisonOperator: "GreaterThanOrEqualToThreshold",
    evaluationPeriods: 2,
    insufficientDataActions: [],
    metricName: "CPUUtilization",
    namespace: "AWS/EC2",
    period: 120,
    statistic: "Average",
    threshold: 80,
});
```
```python
import pulumi
import pulumi_aws as aws

foobar = aws.cloudwatch.MetricAlarm("foobar",
    alarm_description="This metric monitors ec2 cpu utilization",
    comparison_operator="GreaterThanOrEqualToThreshold",
    evaluation_periods=2,
    insufficient_data_actions=[],
    metric_name="CPUUtilization",
    namespace="AWS/EC2",
    period=120,
    statistic="Average",
    threshold=80)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foobar = new Aws.CloudWatch.MetricAlarm("foobar", new()
    {
        AlarmDescription = "This metric monitors ec2 cpu utilization",
        ComparisonOperator = "GreaterThanOrEqualToThreshold",
        EvaluationPeriods = 2,
        InsufficientDataActions = new[] {},
        MetricName = "CPUUtilization",
        Namespace = "AWS/EC2",
        Period = 120,
        Statistic = "Average",
        Threshold = 80,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewMetricAlarm(ctx, "foobar", &cloudwatch.MetricAlarmArgs{
			AlarmDescription:        pulumi.String("This metric monitors ec2 cpu utilization"),
			ComparisonOperator:      pulumi.String("GreaterThanOrEqualToThreshold"),
			EvaluationPeriods:       pulumi.Int(2),
			InsufficientDataActions: pulumi.Array{},
			MetricName:              pulumi.String("CPUUtilization"),
			Namespace:               pulumi.String("AWS/EC2"),
			Period:                  pulumi.Int(120),
			Statistic:               pulumi.String("Average"),
			Threshold:               pulumi.Float64(80),
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
import com.pulumi.aws.cloudwatch.MetricAlarm;
import com.pulumi.aws.cloudwatch.MetricAlarmArgs;
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
        var foobar = new MetricAlarm("foobar", MetricAlarmArgs.builder()        
            .alarmDescription("This metric monitors ec2 cpu utilization")
            .comparisonOperator("GreaterThanOrEqualToThreshold")
            .evaluationPeriods(2)
            .insufficientDataActions()
            .metricName("CPUUtilization")
            .namespace("AWS/EC2")
            .period(120)
            .statistic("Average")
            .threshold(80)
            .build());

    }
}
```
```yaml
resources:
  foobar:
    type: aws:cloudwatch:MetricAlarm
    properties:
      alarmDescription: This metric monitors ec2 cpu utilization
      comparisonOperator: GreaterThanOrEqualToThreshold
      evaluationPeriods: 2
      insufficientDataActions: []
      metricName: CPUUtilization
      namespace: AWS/EC2
      period: 120
      statistic: Average
      threshold: 80
```
{{% /example %}}
{{% /examples %}}
## Example in Conjunction with Scaling Policies

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const batPolicy = new aws.autoscaling.Policy("batPolicy", {
    scalingAdjustment: 4,
    adjustmentType: "ChangeInCapacity",
    cooldown: 300,
    autoscalingGroupName: aws_autoscaling_group.bar.name,
});
const batMetricAlarm = new aws.cloudwatch.MetricAlarm("batMetricAlarm", {
    comparisonOperator: "GreaterThanOrEqualToThreshold",
    evaluationPeriods: 2,
    metricName: "CPUUtilization",
    namespace: "AWS/EC2",
    period: 120,
    statistic: "Average",
    threshold: 80,
    dimensions: {
        AutoScalingGroupName: aws_autoscaling_group.bar.name,
    },
    alarmDescription: "This metric monitors ec2 cpu utilization",
    alarmActions: [batPolicy.arn],
});
```
```python
import pulumi
import pulumi_aws as aws

bat_policy = aws.autoscaling.Policy("batPolicy",
    scaling_adjustment=4,
    adjustment_type="ChangeInCapacity",
    cooldown=300,
    autoscaling_group_name=aws_autoscaling_group["bar"]["name"])
bat_metric_alarm = aws.cloudwatch.MetricAlarm("batMetricAlarm",
    comparison_operator="GreaterThanOrEqualToThreshold",
    evaluation_periods=2,
    metric_name="CPUUtilization",
    namespace="AWS/EC2",
    period=120,
    statistic="Average",
    threshold=80,
    dimensions={
        "AutoScalingGroupName": aws_autoscaling_group["bar"]["name"],
    },
    alarm_description="This metric monitors ec2 cpu utilization",
    alarm_actions=[bat_policy.arn])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var batPolicy = new Aws.AutoScaling.Policy("batPolicy", new()
    {
        ScalingAdjustment = 4,
        AdjustmentType = "ChangeInCapacity",
        Cooldown = 300,
        AutoscalingGroupName = aws_autoscaling_group.Bar.Name,
    });

    var batMetricAlarm = new Aws.CloudWatch.MetricAlarm("batMetricAlarm", new()
    {
        ComparisonOperator = "GreaterThanOrEqualToThreshold",
        EvaluationPeriods = 2,
        MetricName = "CPUUtilization",
        Namespace = "AWS/EC2",
        Period = 120,
        Statistic = "Average",
        Threshold = 80,
        Dimensions = 
        {
            { "AutoScalingGroupName", aws_autoscaling_group.Bar.Name },
        },
        AlarmDescription = "This metric monitors ec2 cpu utilization",
        AlarmActions = new[]
        {
            batPolicy.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		batPolicy, err := autoscaling.NewPolicy(ctx, "batPolicy", &autoscaling.PolicyArgs{
			ScalingAdjustment:    pulumi.Int(4),
			AdjustmentType:       pulumi.String("ChangeInCapacity"),
			Cooldown:             pulumi.Int(300),
			AutoscalingGroupName: pulumi.Any(aws_autoscaling_group.Bar.Name),
		})
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewMetricAlarm(ctx, "batMetricAlarm", &cloudwatch.MetricAlarmArgs{
			ComparisonOperator: pulumi.String("GreaterThanOrEqualToThreshold"),
			EvaluationPeriods:  pulumi.Int(2),
			MetricName:         pulumi.String("CPUUtilization"),
			Namespace:          pulumi.String("AWS/EC2"),
			Period:             pulumi.Int(120),
			Statistic:          pulumi.String("Average"),
			Threshold:          pulumi.Float64(80),
			Dimensions: pulumi.StringMap{
				"AutoScalingGroupName": pulumi.Any(aws_autoscaling_group.Bar.Name),
			},
			AlarmDescription: pulumi.String("This metric monitors ec2 cpu utilization"),
			AlarmActions: pulumi.Array{
				batPolicy.Arn,
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
import com.pulumi.aws.cloudwatch.MetricAlarm;
import com.pulumi.aws.cloudwatch.MetricAlarmArgs;
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
        var batPolicy = new Policy("batPolicy", PolicyArgs.builder()        
            .scalingAdjustment(4)
            .adjustmentType("ChangeInCapacity")
            .cooldown(300)
            .autoscalingGroupName(aws_autoscaling_group.bar().name())
            .build());

        var batMetricAlarm = new MetricAlarm("batMetricAlarm", MetricAlarmArgs.builder()        
            .comparisonOperator("GreaterThanOrEqualToThreshold")
            .evaluationPeriods(2)
            .metricName("CPUUtilization")
            .namespace("AWS/EC2")
            .period(120)
            .statistic("Average")
            .threshold(80)
            .dimensions(Map.of("AutoScalingGroupName", aws_autoscaling_group.bar().name()))
            .alarmDescription("This metric monitors ec2 cpu utilization")
            .alarmActions(batPolicy.arn())
            .build());

    }
}
```
```yaml
resources:
  batPolicy:
    type: aws:autoscaling:Policy
    properties:
      scalingAdjustment: 4
      adjustmentType: ChangeInCapacity
      cooldown: 300
      autoscalingGroupName: ${aws_autoscaling_group.bar.name}
  batMetricAlarm:
    type: aws:cloudwatch:MetricAlarm
    properties:
      comparisonOperator: GreaterThanOrEqualToThreshold
      evaluationPeriods: 2
      metricName: CPUUtilization
      namespace: AWS/EC2
      period: 120
      statistic: Average
      threshold: 80
      dimensions:
        AutoScalingGroupName: ${aws_autoscaling_group.bar.name}
      alarmDescription: This metric monitors ec2 cpu utilization
      alarmActions:
        - ${batPolicy.arn}
```

## Example with an Expression

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobar = new aws.cloudwatch.MetricAlarm("foobar", {
    alarmDescription: "Request error rate has exceeded 10%",
    comparisonOperator: "GreaterThanOrEqualToThreshold",
    evaluationPeriods: 2,
    insufficientDataActions: [],
    metricQueries: [
        {
            expression: "m2/m1*100",
            id: "e1",
            label: "Error Rate",
            returnData: true,
        },
        {
            id: "m1",
            metric: {
                dimensions: {
                    LoadBalancer: "app/web",
                },
                metricName: "RequestCount",
                namespace: "AWS/ApplicationELB",
                period: 120,
                stat: "Sum",
                unit: "Count",
            },
        },
        {
            id: "m2",
            metric: {
                dimensions: {
                    LoadBalancer: "app/web",
                },
                metricName: "HTTPCode_ELB_5XX_Count",
                namespace: "AWS/ApplicationELB",
                period: 120,
                stat: "Sum",
                unit: "Count",
            },
        },
    ],
    threshold: 10,
});
```
```python
import pulumi
import pulumi_aws as aws

foobar = aws.cloudwatch.MetricAlarm("foobar",
    alarm_description="Request error rate has exceeded 10%",
    comparison_operator="GreaterThanOrEqualToThreshold",
    evaluation_periods=2,
    insufficient_data_actions=[],
    metric_queries=[
        aws.cloudwatch.MetricAlarmMetricQueryArgs(
            expression="m2/m1*100",
            id="e1",
            label="Error Rate",
            return_data=True,
        ),
        aws.cloudwatch.MetricAlarmMetricQueryArgs(
            id="m1",
            metric=aws.cloudwatch.MetricAlarmMetricQueryMetricArgs(
                dimensions={
                    "LoadBalancer": "app/web",
                },
                metric_name="RequestCount",
                namespace="AWS/ApplicationELB",
                period=120,
                stat="Sum",
                unit="Count",
            ),
        ),
        aws.cloudwatch.MetricAlarmMetricQueryArgs(
            id="m2",
            metric=aws.cloudwatch.MetricAlarmMetricQueryMetricArgs(
                dimensions={
                    "LoadBalancer": "app/web",
                },
                metric_name="HTTPCode_ELB_5XX_Count",
                namespace="AWS/ApplicationELB",
                period=120,
                stat="Sum",
                unit="Count",
            ),
        ),
    ],
    threshold=10)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foobar = new Aws.CloudWatch.MetricAlarm("foobar", new()
    {
        AlarmDescription = "Request error rate has exceeded 10%",
        ComparisonOperator = "GreaterThanOrEqualToThreshold",
        EvaluationPeriods = 2,
        InsufficientDataActions = new[] {},
        MetricQueries = new[]
        {
            new Aws.CloudWatch.Inputs.MetricAlarmMetricQueryArgs
            {
                Expression = "m2/m1*100",
                Id = "e1",
                Label = "Error Rate",
                ReturnData = true,
            },
            new Aws.CloudWatch.Inputs.MetricAlarmMetricQueryArgs
            {
                Id = "m1",
                Metric = new Aws.CloudWatch.Inputs.MetricAlarmMetricQueryMetricArgs
                {
                    Dimensions = 
                    {
                        { "LoadBalancer", "app/web" },
                    },
                    MetricName = "RequestCount",
                    Namespace = "AWS/ApplicationELB",
                    Period = 120,
                    Stat = "Sum",
                    Unit = "Count",
                },
            },
            new Aws.CloudWatch.Inputs.MetricAlarmMetricQueryArgs
            {
                Id = "m2",
                Metric = new Aws.CloudWatch.Inputs.MetricAlarmMetricQueryMetricArgs
                {
                    Dimensions = 
                    {
                        { "LoadBalancer", "app/web" },
                    },
                    MetricName = "HTTPCode_ELB_5XX_Count",
                    Namespace = "AWS/ApplicationELB",
                    Period = 120,
                    Stat = "Sum",
                    Unit = "Count",
                },
            },
        },
        Threshold = 10,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewMetricAlarm(ctx, "foobar", &cloudwatch.MetricAlarmArgs{
			AlarmDescription:        pulumi.String("Request error rate has exceeded 10%"),
			ComparisonOperator:      pulumi.String("GreaterThanOrEqualToThreshold"),
			EvaluationPeriods:       pulumi.Int(2),
			InsufficientDataActions: pulumi.Array{},
			MetricQueries: cloudwatch.MetricAlarmMetricQueryArray{
				&cloudwatch.MetricAlarmMetricQueryArgs{
					Expression: pulumi.String("m2/m1*100"),
					Id:         pulumi.String("e1"),
					Label:      pulumi.String("Error Rate"),
					ReturnData: pulumi.Bool(true),
				},
				&cloudwatch.MetricAlarmMetricQueryArgs{
					Id: pulumi.String("m1"),
					Metric: &cloudwatch.MetricAlarmMetricQueryMetricArgs{
						Dimensions: pulumi.StringMap{
							"LoadBalancer": pulumi.String("app/web"),
						},
						MetricName: pulumi.String("RequestCount"),
						Namespace:  pulumi.String("AWS/ApplicationELB"),
						Period:     pulumi.Int(120),
						Stat:       pulumi.String("Sum"),
						Unit:       pulumi.String("Count"),
					},
				},
				&cloudwatch.MetricAlarmMetricQueryArgs{
					Id: pulumi.String("m2"),
					Metric: &cloudwatch.MetricAlarmMetricQueryMetricArgs{
						Dimensions: pulumi.StringMap{
							"LoadBalancer": pulumi.String("app/web"),
						},
						MetricName: pulumi.String("HTTPCode_ELB_5XX_Count"),
						Namespace:  pulumi.String("AWS/ApplicationELB"),
						Period:     pulumi.Int(120),
						Stat:       pulumi.String("Sum"),
						Unit:       pulumi.String("Count"),
					},
				},
			},
			Threshold: pulumi.Float64(10),
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
import com.pulumi.aws.cloudwatch.MetricAlarm;
import com.pulumi.aws.cloudwatch.MetricAlarmArgs;
import com.pulumi.aws.cloudwatch.inputs.MetricAlarmMetricQueryArgs;
import com.pulumi.aws.cloudwatch.inputs.MetricAlarmMetricQueryMetricArgs;
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
        var foobar = new MetricAlarm("foobar", MetricAlarmArgs.builder()        
            .alarmDescription("Request error rate has exceeded 10%")
            .comparisonOperator("GreaterThanOrEqualToThreshold")
            .evaluationPeriods(2)
            .insufficientDataActions()
            .metricQueries(            
                MetricAlarmMetricQueryArgs.builder()
                    .expression("m2/m1*100")
                    .id("e1")
                    .label("Error Rate")
                    .returnData("true")
                    .build(),
                MetricAlarmMetricQueryArgs.builder()
                    .id("m1")
                    .metric(MetricAlarmMetricQueryMetricArgs.builder()
                        .dimensions(Map.of("LoadBalancer", "app/web"))
                        .metricName("RequestCount")
                        .namespace("AWS/ApplicationELB")
                        .period(120)
                        .stat("Sum")
                        .unit("Count")
                        .build())
                    .build(),
                MetricAlarmMetricQueryArgs.builder()
                    .id("m2")
                    .metric(MetricAlarmMetricQueryMetricArgs.builder()
                        .dimensions(Map.of("LoadBalancer", "app/web"))
                        .metricName("HTTPCode_ELB_5XX_Count")
                        .namespace("AWS/ApplicationELB")
                        .period(120)
                        .stat("Sum")
                        .unit("Count")
                        .build())
                    .build())
            .threshold(10)
            .build());

    }
}
```
```yaml
resources:
  foobar:
    type: aws:cloudwatch:MetricAlarm
    properties:
      alarmDescription: Request error rate has exceeded 10%
      comparisonOperator: GreaterThanOrEqualToThreshold
      evaluationPeriods: 2
      insufficientDataActions: []
      metricQueries:
        - expression: m2/m1*100
          id: e1
          label: Error Rate
          returnData: 'true'
        - id: m1
          metric:
            dimensions:
              LoadBalancer: app/web
            metricName: RequestCount
            namespace: AWS/ApplicationELB
            period: 120
            stat: Sum
            unit: Count
        - id: m2
          metric:
            dimensions:
              LoadBalancer: app/web
            metricName: HTTPCode_ELB_5XX_Count
            namespace: AWS/ApplicationELB
            period: 120
            stat: Sum
            unit: Count
      threshold: 10
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const xxAnomalyDetection = new aws.cloudwatch.MetricAlarm("xxAnomalyDetection", {
    alarmDescription: "This metric monitors ec2 cpu utilization",
    comparisonOperator: "GreaterThanUpperThreshold",
    evaluationPeriods: 2,
    insufficientDataActions: [],
    metricQueries: [
        {
            expression: "ANOMALY_DETECTION_BAND(m1)",
            id: "e1",
            label: "CPUUtilization (Expected)",
            returnData: true,
        },
        {
            id: "m1",
            metric: {
                dimensions: {
                    InstanceId: "i-abc123",
                },
                metricName: "CPUUtilization",
                namespace: "AWS/EC2",
                period: 120,
                stat: "Average",
                unit: "Count",
            },
            returnData: true,
        },
    ],
    thresholdMetricId: "e1",
});
```
```python
import pulumi
import pulumi_aws as aws

xx_anomaly_detection = aws.cloudwatch.MetricAlarm("xxAnomalyDetection",
    alarm_description="This metric monitors ec2 cpu utilization",
    comparison_operator="GreaterThanUpperThreshold",
    evaluation_periods=2,
    insufficient_data_actions=[],
    metric_queries=[
        aws.cloudwatch.MetricAlarmMetricQueryArgs(
            expression="ANOMALY_DETECTION_BAND(m1)",
            id="e1",
            label="CPUUtilization (Expected)",
            return_data=True,
        ),
        aws.cloudwatch.MetricAlarmMetricQueryArgs(
            id="m1",
            metric=aws.cloudwatch.MetricAlarmMetricQueryMetricArgs(
                dimensions={
                    "InstanceId": "i-abc123",
                },
                metric_name="CPUUtilization",
                namespace="AWS/EC2",
                period=120,
                stat="Average",
                unit="Count",
            ),
            return_data=True,
        ),
    ],
    threshold_metric_id="e1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var xxAnomalyDetection = new Aws.CloudWatch.MetricAlarm("xxAnomalyDetection", new()
    {
        AlarmDescription = "This metric monitors ec2 cpu utilization",
        ComparisonOperator = "GreaterThanUpperThreshold",
        EvaluationPeriods = 2,
        InsufficientDataActions = new[] {},
        MetricQueries = new[]
        {
            new Aws.CloudWatch.Inputs.MetricAlarmMetricQueryArgs
            {
                Expression = "ANOMALY_DETECTION_BAND(m1)",
                Id = "e1",
                Label = "CPUUtilization (Expected)",
                ReturnData = true,
            },
            new Aws.CloudWatch.Inputs.MetricAlarmMetricQueryArgs
            {
                Id = "m1",
                Metric = new Aws.CloudWatch.Inputs.MetricAlarmMetricQueryMetricArgs
                {
                    Dimensions = 
                    {
                        { "InstanceId", "i-abc123" },
                    },
                    MetricName = "CPUUtilization",
                    Namespace = "AWS/EC2",
                    Period = 120,
                    Stat = "Average",
                    Unit = "Count",
                },
                ReturnData = true,
            },
        },
        ThresholdMetricId = "e1",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewMetricAlarm(ctx, "xxAnomalyDetection", &cloudwatch.MetricAlarmArgs{
			AlarmDescription:        pulumi.String("This metric monitors ec2 cpu utilization"),
			ComparisonOperator:      pulumi.String("GreaterThanUpperThreshold"),
			EvaluationPeriods:       pulumi.Int(2),
			InsufficientDataActions: pulumi.Array{},
			MetricQueries: cloudwatch.MetricAlarmMetricQueryArray{
				&cloudwatch.MetricAlarmMetricQueryArgs{
					Expression: pulumi.String("ANOMALY_DETECTION_BAND(m1)"),
					Id:         pulumi.String("e1"),
					Label:      pulumi.String("CPUUtilization (Expected)"),
					ReturnData: pulumi.Bool(true),
				},
				&cloudwatch.MetricAlarmMetricQueryArgs{
					Id: pulumi.String("m1"),
					Metric: &cloudwatch.MetricAlarmMetricQueryMetricArgs{
						Dimensions: pulumi.StringMap{
							"InstanceId": pulumi.String("i-abc123"),
						},
						MetricName: pulumi.String("CPUUtilization"),
						Namespace:  pulumi.String("AWS/EC2"),
						Period:     pulumi.Int(120),
						Stat:       pulumi.String("Average"),
						Unit:       pulumi.String("Count"),
					},
					ReturnData: pulumi.Bool(true),
				},
			},
			ThresholdMetricId: pulumi.String("e1"),
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
import com.pulumi.aws.cloudwatch.MetricAlarm;
import com.pulumi.aws.cloudwatch.MetricAlarmArgs;
import com.pulumi.aws.cloudwatch.inputs.MetricAlarmMetricQueryArgs;
import com.pulumi.aws.cloudwatch.inputs.MetricAlarmMetricQueryMetricArgs;
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
        var xxAnomalyDetection = new MetricAlarm("xxAnomalyDetection", MetricAlarmArgs.builder()        
            .alarmDescription("This metric monitors ec2 cpu utilization")
            .comparisonOperator("GreaterThanUpperThreshold")
            .evaluationPeriods(2)
            .insufficientDataActions()
            .metricQueries(            
                MetricAlarmMetricQueryArgs.builder()
                    .expression("ANOMALY_DETECTION_BAND(m1)")
                    .id("e1")
                    .label("CPUUtilization (Expected)")
                    .returnData("true")
                    .build(),
                MetricAlarmMetricQueryArgs.builder()
                    .id("m1")
                    .metric(MetricAlarmMetricQueryMetricArgs.builder()
                        .dimensions(Map.of("InstanceId", "i-abc123"))
                        .metricName("CPUUtilization")
                        .namespace("AWS/EC2")
                        .period(120)
                        .stat("Average")
                        .unit("Count")
                        .build())
                    .returnData("true")
                    .build())
            .thresholdMetricId("e1")
            .build());

    }
}
```
```yaml
resources:
  xxAnomalyDetection:
    type: aws:cloudwatch:MetricAlarm
    properties:
      alarmDescription: This metric monitors ec2 cpu utilization
      comparisonOperator: GreaterThanUpperThreshold
      evaluationPeriods: 2
      insufficientDataActions: []
      metricQueries:
        - expression: ANOMALY_DETECTION_BAND(m1)
          id: e1
          label: CPUUtilization (Expected)
          returnData: 'true'
        - id: m1
          metric:
            dimensions:
              InstanceId: i-abc123
            metricName: CPUUtilization
            namespace: AWS/EC2
            period: 120
            stat: Average
            unit: Count
          returnData: 'true'
      thresholdMetricId: e1
```

## Example of monitoring Healthy Hosts on NLB using Target Group and NLB

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const nlbHealthyhosts = new aws.cloudwatch.MetricAlarm("nlbHealthyhosts", {
    comparisonOperator: "LessThanThreshold",
    evaluationPeriods: 1,
    metricName: "HealthyHostCount",
    namespace: "AWS/NetworkELB",
    period: 60,
    statistic: "Average",
    threshold: _var.logstash_servers_count,
    alarmDescription: "Number of healthy nodes in Target Group",
    actionsEnabled: true,
    alarmActions: [aws_sns_topic.sns.arn],
    okActions: [aws_sns_topic.sns.arn],
    dimensions: {
        TargetGroup: aws_lb_target_group["lb-tg"].arn_suffix,
        LoadBalancer: aws_lb.lb.arn_suffix,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

nlb_healthyhosts = aws.cloudwatch.MetricAlarm("nlbHealthyhosts",
    comparison_operator="LessThanThreshold",
    evaluation_periods=1,
    metric_name="HealthyHostCount",
    namespace="AWS/NetworkELB",
    period=60,
    statistic="Average",
    threshold=var["logstash_servers_count"],
    alarm_description="Number of healthy nodes in Target Group",
    actions_enabled=True,
    alarm_actions=[aws_sns_topic["sns"]["arn"]],
    ok_actions=[aws_sns_topic["sns"]["arn"]],
    dimensions={
        "TargetGroup": aws_lb_target_group["lb-tg"]["arn_suffix"],
        "LoadBalancer": aws_lb["lb"]["arn_suffix"],
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var nlbHealthyhosts = new Aws.CloudWatch.MetricAlarm("nlbHealthyhosts", new()
    {
        ComparisonOperator = "LessThanThreshold",
        EvaluationPeriods = 1,
        MetricName = "HealthyHostCount",
        Namespace = "AWS/NetworkELB",
        Period = 60,
        Statistic = "Average",
        Threshold = @var.Logstash_servers_count,
        AlarmDescription = "Number of healthy nodes in Target Group",
        ActionsEnabled = true,
        AlarmActions = new[]
        {
            aws_sns_topic.Sns.Arn,
        },
        OkActions = new[]
        {
            aws_sns_topic.Sns.Arn,
        },
        Dimensions = 
        {
            { "TargetGroup", aws_lb_target_group.Lb_tg.Arn_suffix },
            { "LoadBalancer", aws_lb.Lb.Arn_suffix },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewMetricAlarm(ctx, "nlbHealthyhosts", &cloudwatch.MetricAlarmArgs{
			ComparisonOperator: pulumi.String("LessThanThreshold"),
			EvaluationPeriods:  pulumi.Int(1),
			MetricName:         pulumi.String("HealthyHostCount"),
			Namespace:          pulumi.String("AWS/NetworkELB"),
			Period:             pulumi.Int(60),
			Statistic:          pulumi.String("Average"),
			Threshold:          pulumi.Any(_var.Logstash_servers_count),
			AlarmDescription:   pulumi.String("Number of healthy nodes in Target Group"),
			ActionsEnabled:     pulumi.Bool(true),
			AlarmActions: pulumi.Array{
				aws_sns_topic.Sns.Arn,
			},
			OkActions: pulumi.Array{
				aws_sns_topic.Sns.Arn,
			},
			Dimensions: pulumi.StringMap{
				"TargetGroup":  pulumi.Any(aws_lb_target_group.LbTg.Arn_suffix),
				"LoadBalancer": pulumi.Any(aws_lb.Lb.Arn_suffix),
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
import com.pulumi.aws.cloudwatch.MetricAlarm;
import com.pulumi.aws.cloudwatch.MetricAlarmArgs;
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
        var nlbHealthyhosts = new MetricAlarm("nlbHealthyhosts", MetricAlarmArgs.builder()        
            .comparisonOperator("LessThanThreshold")
            .evaluationPeriods(1)
            .metricName("HealthyHostCount")
            .namespace("AWS/NetworkELB")
            .period(60)
            .statistic("Average")
            .threshold(var_.logstash_servers_count())
            .alarmDescription("Number of healthy nodes in Target Group")
            .actionsEnabled("true")
            .alarmActions(aws_sns_topic.sns().arn())
            .okActions(aws_sns_topic.sns().arn())
            .dimensions(Map.ofEntries(
                Map.entry("TargetGroup", aws_lb_target_group.lb-tg().arn_suffix()),
                Map.entry("LoadBalancer", aws_lb.lb().arn_suffix())
            ))
            .build());

    }
}
```
```yaml
resources:
  nlbHealthyhosts:
    type: aws:cloudwatch:MetricAlarm
    properties:
      comparisonOperator: LessThanThreshold
      evaluationPeriods: 1
      metricName: HealthyHostCount
      namespace: AWS/NetworkELB
      period: 60
      statistic: Average
      threshold: ${var.logstash_servers_count}
      alarmDescription: Number of healthy nodes in Target Group
      actionsEnabled: 'true'
      alarmActions:
        - ${aws_sns_topic.sns.arn}
      okActions:
        - ${aws_sns_topic.sns.arn}
      dimensions:
        TargetGroup: ${aws_lb_target_group"lb-tg"[%!s(MISSING)].arn_suffix}
        LoadBalancer: ${aws_lb.lb.arn_suffix}
```

> **NOTE:**  You cannot create a metric alarm consisting of both `statistic` and `extended_statistic` parameters.
You must choose one or the other


## Import

Using `pulumi import`, import CloudWatch Metric Alarm using the `alarm_name`. For example:

```sh
 $ pulumi import aws:cloudwatch/metricAlarm:MetricAlarm test alarm-12345
```
 