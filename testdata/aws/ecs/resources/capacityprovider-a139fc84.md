Provides an ECS cluster capacity provider. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).

> **NOTE:** Associating an ECS Capacity Provider to an Auto Scaling Group will automatically add the `AmazonECSManaged` tag to the Auto Scaling Group. This tag should be included in the `aws.autoscaling.Group` resource configuration to prevent the provider from removing it in subsequent executions as well as ensuring the `AmazonECSManaged` tag is propagated to all EC2 Instances in the Auto Scaling Group if `min_size` is above 0 on creation. Any EC2 Instances in the Auto Scaling Group without this tag must be manually be updated, otherwise they may cause unexpected scaling behavior and metrics.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ... other configuration, including potentially other tags ...
const testGroup = new aws.autoscaling.Group("testGroup", {tags: [{
    key: "AmazonECSManaged",
    value: "true",
    propagateAtLaunch: true,
}]});
const testCapacityProvider = new aws.ecs.CapacityProvider("testCapacityProvider", {autoScalingGroupProvider: {
    autoScalingGroupArn: testGroup.arn,
    managedTerminationProtection: "ENABLED",
    managedScaling: {
        maximumScalingStepSize: 1000,
        minimumScalingStepSize: 1,
        status: "ENABLED",
        targetCapacity: 10,
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configuration, including potentially other tags ...
test_group = aws.autoscaling.Group("testGroup", tags=[aws.autoscaling.GroupTagArgs(
    key="AmazonECSManaged",
    value="true",
    propagate_at_launch=True,
)])
test_capacity_provider = aws.ecs.CapacityProvider("testCapacityProvider", auto_scaling_group_provider=aws.ecs.CapacityProviderAutoScalingGroupProviderArgs(
    auto_scaling_group_arn=test_group.arn,
    managed_termination_protection="ENABLED",
    managed_scaling=aws.ecs.CapacityProviderAutoScalingGroupProviderManagedScalingArgs(
        maximum_scaling_step_size=1000,
        minimum_scaling_step_size=1,
        status="ENABLED",
        target_capacity=10,
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
    // ... other configuration, including potentially other tags ...
    var testGroup = new Aws.AutoScaling.Group("testGroup", new()
    {
        Tags = new[]
        {
            new Aws.AutoScaling.Inputs.GroupTagArgs
            {
                Key = "AmazonECSManaged",
                Value = "true",
                PropagateAtLaunch = true,
            },
        },
    });

    var testCapacityProvider = new Aws.Ecs.CapacityProvider("testCapacityProvider", new()
    {
        AutoScalingGroupProvider = new Aws.Ecs.Inputs.CapacityProviderAutoScalingGroupProviderArgs
        {
            AutoScalingGroupArn = testGroup.Arn,
            ManagedTerminationProtection = "ENABLED",
            ManagedScaling = new Aws.Ecs.Inputs.CapacityProviderAutoScalingGroupProviderManagedScalingArgs
            {
                MaximumScalingStepSize = 1000,
                MinimumScalingStepSize = 1,
                Status = "ENABLED",
                TargetCapacity = 10,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testGroup, err := autoscaling.NewGroup(ctx, "testGroup", &autoscaling.GroupArgs{
			Tags: autoscaling.GroupTagArray{
				&autoscaling.GroupTagArgs{
					Key:               pulumi.String("AmazonECSManaged"),
					Value:             pulumi.String("true"),
					PropagateAtLaunch: pulumi.Bool(true),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = ecs.NewCapacityProvider(ctx, "testCapacityProvider", &ecs.CapacityProviderArgs{
			AutoScalingGroupProvider: &ecs.CapacityProviderAutoScalingGroupProviderArgs{
				AutoScalingGroupArn:          testGroup.Arn,
				ManagedTerminationProtection: pulumi.String("ENABLED"),
				ManagedScaling: &ecs.CapacityProviderAutoScalingGroupProviderManagedScalingArgs{
					MaximumScalingStepSize: pulumi.Int(1000),
					MinimumScalingStepSize: pulumi.Int(1),
					Status:                 pulumi.String("ENABLED"),
					TargetCapacity:         pulumi.Int(10),
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
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupTagArgs;
import com.pulumi.aws.ecs.CapacityProvider;
import com.pulumi.aws.ecs.CapacityProviderArgs;
import com.pulumi.aws.ecs.inputs.CapacityProviderAutoScalingGroupProviderArgs;
import com.pulumi.aws.ecs.inputs.CapacityProviderAutoScalingGroupProviderManagedScalingArgs;
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
        var testGroup = new Group("testGroup", GroupArgs.builder()        
            .tags(GroupTagArgs.builder()
                .key("AmazonECSManaged")
                .value(true)
                .propagateAtLaunch(true)
                .build())
            .build());

        var testCapacityProvider = new CapacityProvider("testCapacityProvider", CapacityProviderArgs.builder()        
            .autoScalingGroupProvider(CapacityProviderAutoScalingGroupProviderArgs.builder()
                .autoScalingGroupArn(testGroup.arn())
                .managedTerminationProtection("ENABLED")
                .managedScaling(CapacityProviderAutoScalingGroupProviderManagedScalingArgs.builder()
                    .maximumScalingStepSize(1000)
                    .minimumScalingStepSize(1)
                    .status("ENABLED")
                    .targetCapacity(10)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testGroup:
    type: aws:autoscaling:Group
    properties:
      tags:
        - key: AmazonECSManaged
          value: true
          propagateAtLaunch: true
  testCapacityProvider:
    type: aws:ecs:CapacityProvider
    properties:
      autoScalingGroupProvider:
        autoScalingGroupArn: ${testGroup.arn}
        managedTerminationProtection: ENABLED
        managedScaling:
          maximumScalingStepSize: 1000
          minimumScalingStepSize: 1
          status: ENABLED
          targetCapacity: 10
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECS Capacity Providers using the `name`. For example:

```sh
 $ pulumi import aws:ecs/capacityProvider:CapacityProvider example example
```
 