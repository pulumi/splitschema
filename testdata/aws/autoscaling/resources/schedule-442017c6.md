Provides an AutoScaling Schedule resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobarGroup = new aws.autoscaling.Group("foobarGroup", {
    availabilityZones: ["us-west-2a"],
    maxSize: 1,
    minSize: 1,
    healthCheckGracePeriod: 300,
    healthCheckType: "ELB",
    forceDelete: true,
    terminationPolicies: ["OldestInstance"],
});
const foobarSchedule = new aws.autoscaling.Schedule("foobarSchedule", {
    scheduledActionName: "foobar",
    minSize: 0,
    maxSize: 1,
    desiredCapacity: 0,
    startTime: "2016-12-11T18:00:00Z",
    endTime: "2016-12-12T06:00:00Z",
    autoscalingGroupName: foobarGroup.name,
});
```
```python
import pulumi
import pulumi_aws as aws

foobar_group = aws.autoscaling.Group("foobarGroup",
    availability_zones=["us-west-2a"],
    max_size=1,
    min_size=1,
    health_check_grace_period=300,
    health_check_type="ELB",
    force_delete=True,
    termination_policies=["OldestInstance"])
foobar_schedule = aws.autoscaling.Schedule("foobarSchedule",
    scheduled_action_name="foobar",
    min_size=0,
    max_size=1,
    desired_capacity=0,
    start_time="2016-12-11T18:00:00Z",
    end_time="2016-12-12T06:00:00Z",
    autoscaling_group_name=foobar_group.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foobarGroup = new Aws.AutoScaling.Group("foobarGroup", new()
    {
        AvailabilityZones = new[]
        {
            "us-west-2a",
        },
        MaxSize = 1,
        MinSize = 1,
        HealthCheckGracePeriod = 300,
        HealthCheckType = "ELB",
        ForceDelete = true,
        TerminationPolicies = new[]
        {
            "OldestInstance",
        },
    });

    var foobarSchedule = new Aws.AutoScaling.Schedule("foobarSchedule", new()
    {
        ScheduledActionName = "foobar",
        MinSize = 0,
        MaxSize = 1,
        DesiredCapacity = 0,
        StartTime = "2016-12-11T18:00:00Z",
        EndTime = "2016-12-12T06:00:00Z",
        AutoscalingGroupName = foobarGroup.Name,
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
		foobarGroup, err := autoscaling.NewGroup(ctx, "foobarGroup", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
			},
			MaxSize:                pulumi.Int(1),
			MinSize:                pulumi.Int(1),
			HealthCheckGracePeriod: pulumi.Int(300),
			HealthCheckType:        pulumi.String("ELB"),
			ForceDelete:            pulumi.Bool(true),
			TerminationPolicies: pulumi.StringArray{
				pulumi.String("OldestInstance"),
			},
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewSchedule(ctx, "foobarSchedule", &autoscaling.ScheduleArgs{
			ScheduledActionName:  pulumi.String("foobar"),
			MinSize:              pulumi.Int(0),
			MaxSize:              pulumi.Int(1),
			DesiredCapacity:      pulumi.Int(0),
			StartTime:            pulumi.String("2016-12-11T18:00:00Z"),
			EndTime:              pulumi.String("2016-12-12T06:00:00Z"),
			AutoscalingGroupName: foobarGroup.Name,
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
import com.pulumi.aws.autoscaling.Schedule;
import com.pulumi.aws.autoscaling.ScheduleArgs;
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
        var foobarGroup = new Group("foobarGroup", GroupArgs.builder()        
            .availabilityZones("us-west-2a")
            .maxSize(1)
            .minSize(1)
            .healthCheckGracePeriod(300)
            .healthCheckType("ELB")
            .forceDelete(true)
            .terminationPolicies("OldestInstance")
            .build());

        var foobarSchedule = new Schedule("foobarSchedule", ScheduleArgs.builder()        
            .scheduledActionName("foobar")
            .minSize(0)
            .maxSize(1)
            .desiredCapacity(0)
            .startTime("2016-12-11T18:00:00Z")
            .endTime("2016-12-12T06:00:00Z")
            .autoscalingGroupName(foobarGroup.name())
            .build());

    }
}
```
```yaml
resources:
  foobarGroup:
    type: aws:autoscaling:Group
    properties:
      availabilityZones:
        - us-west-2a
      maxSize: 1
      minSize: 1
      healthCheckGracePeriod: 300
      healthCheckType: ELB
      forceDelete: true
      terminationPolicies:
        - OldestInstance
  foobarSchedule:
    type: aws:autoscaling:Schedule
    properties:
      scheduledActionName: foobar
      minSize: 0
      maxSize: 1
      desiredCapacity: 0
      startTime: 2016-12-11T18:00:00Z
      endTime: 2016-12-12T06:00:00Z
      autoscalingGroupName: ${foobarGroup.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AutoScaling ScheduledAction using the `auto-scaling-group-name` and `scheduled-action-name`. For example:

```sh
 $ pulumi import aws:autoscaling/schedule:Schedule resource-name auto-scaling-group-name/scheduled-action-name
```
 