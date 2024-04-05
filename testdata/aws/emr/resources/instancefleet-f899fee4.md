Provides an Elastic MapReduce Cluster Instance Fleet configuration.
See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/emr/) for more information.

> **NOTE:** At this time, Instance Fleets cannot be destroyed through the API nor
web interface. Instance Fleets are destroyed when the EMR Cluster is destroyed.
the provider will resize any Instance Fleet to zero when destroying the resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const task = new aws.emr.InstanceFleet("task", {
    clusterId: aws_emr_cluster.cluster.id,
    instanceTypeConfigs: [
        {
            bidPriceAsPercentageOfOnDemandPrice: 100,
            ebsConfigs: [{
                size: 100,
                type: "gp2",
                volumesPerInstance: 1,
            }],
            instanceType: "m4.xlarge",
            weightedCapacity: 1,
        },
        {
            bidPriceAsPercentageOfOnDemandPrice: 100,
            ebsConfigs: [{
                size: 100,
                type: "gp2",
                volumesPerInstance: 1,
            }],
            instanceType: "m4.2xlarge",
            weightedCapacity: 2,
        },
    ],
    launchSpecifications: {
        spotSpecifications: [{
            allocationStrategy: "capacity-optimized",
            blockDurationMinutes: 0,
            timeoutAction: "TERMINATE_CLUSTER",
            timeoutDurationMinutes: 10,
        }],
    },
    targetOnDemandCapacity: 1,
    targetSpotCapacity: 1,
});
```
```python
import pulumi
import pulumi_aws as aws

task = aws.emr.InstanceFleet("task",
    cluster_id=aws_emr_cluster["cluster"]["id"],
    instance_type_configs=[
        aws.emr.InstanceFleetInstanceTypeConfigArgs(
            bid_price_as_percentage_of_on_demand_price=100,
            ebs_configs=[aws.emr.InstanceFleetInstanceTypeConfigEbsConfigArgs(
                size=100,
                type="gp2",
                volumes_per_instance=1,
            )],
            instance_type="m4.xlarge",
            weighted_capacity=1,
        ),
        aws.emr.InstanceFleetInstanceTypeConfigArgs(
            bid_price_as_percentage_of_on_demand_price=100,
            ebs_configs=[aws.emr.InstanceFleetInstanceTypeConfigEbsConfigArgs(
                size=100,
                type="gp2",
                volumes_per_instance=1,
            )],
            instance_type="m4.2xlarge",
            weighted_capacity=2,
        ),
    ],
    launch_specifications=aws.emr.InstanceFleetLaunchSpecificationsArgs(
        spot_specifications=[aws.emr.InstanceFleetLaunchSpecificationsSpotSpecificationArgs(
            allocation_strategy="capacity-optimized",
            block_duration_minutes=0,
            timeout_action="TERMINATE_CLUSTER",
            timeout_duration_minutes=10,
        )],
    ),
    target_on_demand_capacity=1,
    target_spot_capacity=1)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var task = new Aws.Emr.InstanceFleet("task", new()
    {
        ClusterId = aws_emr_cluster.Cluster.Id,
        InstanceTypeConfigs = new[]
        {
            new Aws.Emr.Inputs.InstanceFleetInstanceTypeConfigArgs
            {
                BidPriceAsPercentageOfOnDemandPrice = 100,
                EbsConfigs = new[]
                {
                    new Aws.Emr.Inputs.InstanceFleetInstanceTypeConfigEbsConfigArgs
                    {
                        Size = 100,
                        Type = "gp2",
                        VolumesPerInstance = 1,
                    },
                },
                InstanceType = "m4.xlarge",
                WeightedCapacity = 1,
            },
            new Aws.Emr.Inputs.InstanceFleetInstanceTypeConfigArgs
            {
                BidPriceAsPercentageOfOnDemandPrice = 100,
                EbsConfigs = new[]
                {
                    new Aws.Emr.Inputs.InstanceFleetInstanceTypeConfigEbsConfigArgs
                    {
                        Size = 100,
                        Type = "gp2",
                        VolumesPerInstance = 1,
                    },
                },
                InstanceType = "m4.2xlarge",
                WeightedCapacity = 2,
            },
        },
        LaunchSpecifications = new Aws.Emr.Inputs.InstanceFleetLaunchSpecificationsArgs
        {
            SpotSpecifications = new[]
            {
                new Aws.Emr.Inputs.InstanceFleetLaunchSpecificationsSpotSpecificationArgs
                {
                    AllocationStrategy = "capacity-optimized",
                    BlockDurationMinutes = 0,
                    TimeoutAction = "TERMINATE_CLUSTER",
                    TimeoutDurationMinutes = 10,
                },
            },
        },
        TargetOnDemandCapacity = 1,
        TargetSpotCapacity = 1,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.NewInstanceFleet(ctx, "task", &emr.InstanceFleetArgs{
			ClusterId: pulumi.Any(aws_emr_cluster.Cluster.Id),
			InstanceTypeConfigs: emr.InstanceFleetInstanceTypeConfigArray{
				&emr.InstanceFleetInstanceTypeConfigArgs{
					BidPriceAsPercentageOfOnDemandPrice: pulumi.Float64(100),
					EbsConfigs: emr.InstanceFleetInstanceTypeConfigEbsConfigArray{
						&emr.InstanceFleetInstanceTypeConfigEbsConfigArgs{
							Size:               pulumi.Int(100),
							Type:               pulumi.String("gp2"),
							VolumesPerInstance: pulumi.Int(1),
						},
					},
					InstanceType:     pulumi.String("m4.xlarge"),
					WeightedCapacity: pulumi.Int(1),
				},
				&emr.InstanceFleetInstanceTypeConfigArgs{
					BidPriceAsPercentageOfOnDemandPrice: pulumi.Float64(100),
					EbsConfigs: emr.InstanceFleetInstanceTypeConfigEbsConfigArray{
						&emr.InstanceFleetInstanceTypeConfigEbsConfigArgs{
							Size:               pulumi.Int(100),
							Type:               pulumi.String("gp2"),
							VolumesPerInstance: pulumi.Int(1),
						},
					},
					InstanceType:     pulumi.String("m4.2xlarge"),
					WeightedCapacity: pulumi.Int(2),
				},
			},
			LaunchSpecifications: &emr.InstanceFleetLaunchSpecificationsArgs{
				SpotSpecifications: emr.InstanceFleetLaunchSpecificationsSpotSpecificationArray{
					&emr.InstanceFleetLaunchSpecificationsSpotSpecificationArgs{
						AllocationStrategy:     pulumi.String("capacity-optimized"),
						BlockDurationMinutes:   pulumi.Int(0),
						TimeoutAction:          pulumi.String("TERMINATE_CLUSTER"),
						TimeoutDurationMinutes: pulumi.Int(10),
					},
				},
			},
			TargetOnDemandCapacity: pulumi.Int(1),
			TargetSpotCapacity:     pulumi.Int(1),
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
import com.pulumi.aws.emr.InstanceFleet;
import com.pulumi.aws.emr.InstanceFleetArgs;
import com.pulumi.aws.emr.inputs.InstanceFleetInstanceTypeConfigArgs;
import com.pulumi.aws.emr.inputs.InstanceFleetLaunchSpecificationsArgs;
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
        var task = new InstanceFleet("task", InstanceFleetArgs.builder()        
            .clusterId(aws_emr_cluster.cluster().id())
            .instanceTypeConfigs(            
                InstanceFleetInstanceTypeConfigArgs.builder()
                    .bidPriceAsPercentageOfOnDemandPrice(100)
                    .ebsConfigs(InstanceFleetInstanceTypeConfigEbsConfigArgs.builder()
                        .size(100)
                        .type("gp2")
                        .volumesPerInstance(1)
                        .build())
                    .instanceType("m4.xlarge")
                    .weightedCapacity(1)
                    .build(),
                InstanceFleetInstanceTypeConfigArgs.builder()
                    .bidPriceAsPercentageOfOnDemandPrice(100)
                    .ebsConfigs(InstanceFleetInstanceTypeConfigEbsConfigArgs.builder()
                        .size(100)
                        .type("gp2")
                        .volumesPerInstance(1)
                        .build())
                    .instanceType("m4.2xlarge")
                    .weightedCapacity(2)
                    .build())
            .launchSpecifications(InstanceFleetLaunchSpecificationsArgs.builder()
                .spotSpecifications(InstanceFleetLaunchSpecificationsSpotSpecificationArgs.builder()
                    .allocationStrategy("capacity-optimized")
                    .blockDurationMinutes(0)
                    .timeoutAction("TERMINATE_CLUSTER")
                    .timeoutDurationMinutes(10)
                    .build())
                .build())
            .targetOnDemandCapacity(1)
            .targetSpotCapacity(1)
            .build());

    }
}
```
```yaml
resources:
  task:
    type: aws:emr:InstanceFleet
    properties:
      clusterId: ${aws_emr_cluster.cluster.id}
      instanceTypeConfigs:
        - bidPriceAsPercentageOfOnDemandPrice: 100
          ebsConfigs:
            - size: 100
              type: gp2
              volumesPerInstance: 1
          instanceType: m4.xlarge
          weightedCapacity: 1
        - bidPriceAsPercentageOfOnDemandPrice: 100
          ebsConfigs:
            - size: 100
              type: gp2
              volumesPerInstance: 1
          instanceType: m4.2xlarge
          weightedCapacity: 2
      launchSpecifications:
        spotSpecifications:
          - allocationStrategy: capacity-optimized
            blockDurationMinutes: 0
            timeoutAction: TERMINATE_CLUSTER
            timeoutDurationMinutes: 10
      targetOnDemandCapacity: 1
      targetSpotCapacity: 1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EMR Instance Fleet using the EMR Cluster identifier and Instance Fleet identifier separated by a forward slash (`/`). For example:

```sh
 $ pulumi import aws:emr/instanceFleet:InstanceFleet example j-123456ABCDEF/if-15EK4O09RZLNR
```
 