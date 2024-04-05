Provides a Batch Scheduling Policy resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.batch.SchedulingPolicy("example", {
    fairSharePolicy: {
        computeReservation: 1,
        shareDecaySeconds: 3600,
        shareDistributions: [
            {
                shareIdentifier: "A1*",
                weightFactor: 0.1,
            },
            {
                shareIdentifier: "A2",
                weightFactor: 0.2,
            },
        ],
    },
    tags: {
        Name: "Example Batch Scheduling Policy",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.batch.SchedulingPolicy("example",
    fair_share_policy=aws.batch.SchedulingPolicyFairSharePolicyArgs(
        compute_reservation=1,
        share_decay_seconds=3600,
        share_distributions=[
            aws.batch.SchedulingPolicyFairSharePolicyShareDistributionArgs(
                share_identifier="A1*",
                weight_factor=0.1,
            ),
            aws.batch.SchedulingPolicyFairSharePolicyShareDistributionArgs(
                share_identifier="A2",
                weight_factor=0.2,
            ),
        ],
    ),
    tags={
        "Name": "Example Batch Scheduling Policy",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Batch.SchedulingPolicy("example", new()
    {
        FairSharePolicy = new Aws.Batch.Inputs.SchedulingPolicyFairSharePolicyArgs
        {
            ComputeReservation = 1,
            ShareDecaySeconds = 3600,
            ShareDistributions = new[]
            {
                new Aws.Batch.Inputs.SchedulingPolicyFairSharePolicyShareDistributionArgs
                {
                    ShareIdentifier = "A1*",
                    WeightFactor = 0.1,
                },
                new Aws.Batch.Inputs.SchedulingPolicyFairSharePolicyShareDistributionArgs
                {
                    ShareIdentifier = "A2",
                    WeightFactor = 0.2,
                },
            },
        },
        Tags = 
        {
            { "Name", "Example Batch Scheduling Policy" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/batch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.NewSchedulingPolicy(ctx, "example", &batch.SchedulingPolicyArgs{
			FairSharePolicy: &batch.SchedulingPolicyFairSharePolicyArgs{
				ComputeReservation: pulumi.Int(1),
				ShareDecaySeconds:  pulumi.Int(3600),
				ShareDistributions: batch.SchedulingPolicyFairSharePolicyShareDistributionArray{
					&batch.SchedulingPolicyFairSharePolicyShareDistributionArgs{
						ShareIdentifier: pulumi.String("A1*"),
						WeightFactor:    pulumi.Float64(0.1),
					},
					&batch.SchedulingPolicyFairSharePolicyShareDistributionArgs{
						ShareIdentifier: pulumi.String("A2"),
						WeightFactor:    pulumi.Float64(0.2),
					},
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Batch Scheduling Policy"),
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
import com.pulumi.aws.batch.SchedulingPolicy;
import com.pulumi.aws.batch.SchedulingPolicyArgs;
import com.pulumi.aws.batch.inputs.SchedulingPolicyFairSharePolicyArgs;
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
        var example = new SchedulingPolicy("example", SchedulingPolicyArgs.builder()        
            .fairSharePolicy(SchedulingPolicyFairSharePolicyArgs.builder()
                .computeReservation(1)
                .shareDecaySeconds(3600)
                .shareDistributions(                
                    SchedulingPolicyFairSharePolicyShareDistributionArgs.builder()
                        .shareIdentifier("A1*")
                        .weightFactor(0.1)
                        .build(),
                    SchedulingPolicyFairSharePolicyShareDistributionArgs.builder()
                        .shareIdentifier("A2")
                        .weightFactor(0.2)
                        .build())
                .build())
            .tags(Map.of("Name", "Example Batch Scheduling Policy"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:batch:SchedulingPolicy
    properties:
      fairSharePolicy:
        computeReservation: 1
        shareDecaySeconds: 3600
        shareDistributions:
          - shareIdentifier: A1*
            weightFactor: 0.1
          - shareIdentifier: A2
            weightFactor: 0.2
      tags:
        Name: Example Batch Scheduling Policy
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Batch Scheduling Policy using the `arn`. For example:

```sh
 $ pulumi import aws:batch/schedulingPolicy:SchedulingPolicy test_policy arn:aws:batch:us-east-1:123456789012:scheduling-policy/sample
```
 