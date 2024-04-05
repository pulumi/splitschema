Provides a Managed Scaling policy for EMR Cluster. With Amazon EMR versions 5.30.0 and later (except for Amazon EMR 6.0.0), you can enable EMR managed scaling to automatically increase or decrease the number of instances or units in your cluster based on workload. See [Using EMR Managed Scaling in Amazon EMR](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-managed-scaling.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sample = new aws.emr.Cluster("sample", {
    releaseLabel: "emr-5.30.0",
    masterInstanceGroup: {
        instanceType: "m4.large",
    },
    coreInstanceGroup: {
        instanceType: "c4.large",
    },
});
// skip ...
const samplepolicy = new aws.emr.ManagedScalingPolicy("samplepolicy", {
    clusterId: sample.id,
    computeLimits: [{
        unitType: "Instances",
        minimumCapacityUnits: 2,
        maximumCapacityUnits: 10,
        maximumOndemandCapacityUnits: 2,
        maximumCoreCapacityUnits: 10,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

sample = aws.emr.Cluster("sample",
    release_label="emr-5.30.0",
    master_instance_group=aws.emr.ClusterMasterInstanceGroupArgs(
        instance_type="m4.large",
    ),
    core_instance_group=aws.emr.ClusterCoreInstanceGroupArgs(
        instance_type="c4.large",
    ))
# skip ...
samplepolicy = aws.emr.ManagedScalingPolicy("samplepolicy",
    cluster_id=sample.id,
    compute_limits=[aws.emr.ManagedScalingPolicyComputeLimitArgs(
        unit_type="Instances",
        minimum_capacity_units=2,
        maximum_capacity_units=10,
        maximum_ondemand_capacity_units=2,
        maximum_core_capacity_units=10,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var sample = new Aws.Emr.Cluster("sample", new()
    {
        ReleaseLabel = "emr-5.30.0",
        MasterInstanceGroup = new Aws.Emr.Inputs.ClusterMasterInstanceGroupArgs
        {
            InstanceType = "m4.large",
        },
        CoreInstanceGroup = new Aws.Emr.Inputs.ClusterCoreInstanceGroupArgs
        {
            InstanceType = "c4.large",
        },
    });

    // skip ...
    var samplepolicy = new Aws.Emr.ManagedScalingPolicy("samplepolicy", new()
    {
        ClusterId = sample.Id,
        ComputeLimits = new[]
        {
            new Aws.Emr.Inputs.ManagedScalingPolicyComputeLimitArgs
            {
                UnitType = "Instances",
                MinimumCapacityUnits = 2,
                MaximumCapacityUnits = 10,
                MaximumOndemandCapacityUnits = 2,
                MaximumCoreCapacityUnits = 10,
            },
        },
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
		sample, err := emr.NewCluster(ctx, "sample", &emr.ClusterArgs{
			ReleaseLabel: pulumi.String("emr-5.30.0"),
			MasterInstanceGroup: &emr.ClusterMasterInstanceGroupArgs{
				InstanceType: pulumi.String("m4.large"),
			},
			CoreInstanceGroup: &emr.ClusterCoreInstanceGroupArgs{
				InstanceType: pulumi.String("c4.large"),
			},
		})
		if err != nil {
			return err
		}
		_, err = emr.NewManagedScalingPolicy(ctx, "samplepolicy", &emr.ManagedScalingPolicyArgs{
			ClusterId: sample.ID(),
			ComputeLimits: emr.ManagedScalingPolicyComputeLimitArray{
				&emr.ManagedScalingPolicyComputeLimitArgs{
					UnitType:                     pulumi.String("Instances"),
					MinimumCapacityUnits:         pulumi.Int(2),
					MaximumCapacityUnits:         pulumi.Int(10),
					MaximumOndemandCapacityUnits: pulumi.Int(2),
					MaximumCoreCapacityUnits:     pulumi.Int(10),
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
import com.pulumi.aws.emr.Cluster;
import com.pulumi.aws.emr.ClusterArgs;
import com.pulumi.aws.emr.inputs.ClusterMasterInstanceGroupArgs;
import com.pulumi.aws.emr.inputs.ClusterCoreInstanceGroupArgs;
import com.pulumi.aws.emr.ManagedScalingPolicy;
import com.pulumi.aws.emr.ManagedScalingPolicyArgs;
import com.pulumi.aws.emr.inputs.ManagedScalingPolicyComputeLimitArgs;
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
        var sample = new Cluster("sample", ClusterArgs.builder()        
            .releaseLabel("emr-5.30.0")
            .masterInstanceGroup(ClusterMasterInstanceGroupArgs.builder()
                .instanceType("m4.large")
                .build())
            .coreInstanceGroup(ClusterCoreInstanceGroupArgs.builder()
                .instanceType("c4.large")
                .build())
            .build());

        var samplepolicy = new ManagedScalingPolicy("samplepolicy", ManagedScalingPolicyArgs.builder()        
            .clusterId(sample.id())
            .computeLimits(ManagedScalingPolicyComputeLimitArgs.builder()
                .unitType("Instances")
                .minimumCapacityUnits(2)
                .maximumCapacityUnits(10)
                .maximumOndemandCapacityUnits(2)
                .maximumCoreCapacityUnits(10)
                .build())
            .build());

    }
}
```
```yaml
resources:
  sample:
    type: aws:emr:Cluster
    properties:
      releaseLabel: emr-5.30.0
      masterInstanceGroup:
        instanceType: m4.large
      coreInstanceGroup:
        instanceType: c4.large
  samplepolicy:
    type: aws:emr:ManagedScalingPolicy
    properties:
      clusterId: ${sample.id}
      computeLimits:
        - unitType: Instances
          minimumCapacityUnits: 2
          maximumCapacityUnits: 10
          maximumOndemandCapacityUnits: 2
          maximumCoreCapacityUnits: 10
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EMR Managed Scaling Policies using the EMR Cluster identifier. For example:

```sh
 $ pulumi import aws:emr/managedScalingPolicy:ManagedScalingPolicy example j-123456ABCDEF
```
 