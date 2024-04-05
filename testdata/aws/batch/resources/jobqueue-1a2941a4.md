Provides a Batch Job Queue resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Job Queue

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testQueue = new aws.batch.JobQueue("testQueue", {
    state: "ENABLED",
    priority: 1,
    computeEnvironments: [
        aws_batch_compute_environment.test_environment_1.arn,
        aws_batch_compute_environment.test_environment_2.arn,
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

test_queue = aws.batch.JobQueue("testQueue",
    state="ENABLED",
    priority=1,
    compute_environments=[
        aws_batch_compute_environment["test_environment_1"]["arn"],
        aws_batch_compute_environment["test_environment_2"]["arn"],
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testQueue = new Aws.Batch.JobQueue("testQueue", new()
    {
        State = "ENABLED",
        Priority = 1,
        ComputeEnvironments = new[]
        {
            aws_batch_compute_environment.Test_environment_1.Arn,
            aws_batch_compute_environment.Test_environment_2.Arn,
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
		_, err := batch.NewJobQueue(ctx, "testQueue", &batch.JobQueueArgs{
			State:    pulumi.String("ENABLED"),
			Priority: pulumi.Int(1),
			ComputeEnvironments: pulumi.StringArray{
				aws_batch_compute_environment.Test_environment_1.Arn,
				aws_batch_compute_environment.Test_environment_2.Arn,
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
import com.pulumi.aws.batch.JobQueue;
import com.pulumi.aws.batch.JobQueueArgs;
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
        var testQueue = new JobQueue("testQueue", JobQueueArgs.builder()        
            .state("ENABLED")
            .priority(1)
            .computeEnvironments(            
                aws_batch_compute_environment.test_environment_1().arn(),
                aws_batch_compute_environment.test_environment_2().arn())
            .build());

    }
}
```
```yaml
resources:
  testQueue:
    type: aws:batch:JobQueue
    properties:
      state: ENABLED
      priority: 1
      computeEnvironments:
        - ${aws_batch_compute_environment.test_environment_1.arn}
        - ${aws_batch_compute_environment.test_environment_2.arn}
```
{{% /example %}}
{{% example %}}
### Job Queue with a fair share scheduling policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleSchedulingPolicy = new aws.batch.SchedulingPolicy("exampleSchedulingPolicy", {fairSharePolicy: {
    computeReservation: 1,
    shareDecaySeconds: 3600,
    shareDistributions: [{
        shareIdentifier: "A1*",
        weightFactor: 0.1,
    }],
}});
const exampleJobQueue = new aws.batch.JobQueue("exampleJobQueue", {
    schedulingPolicyArn: exampleSchedulingPolicy.arn,
    state: "ENABLED",
    priority: 1,
    computeEnvironments: [
        aws_batch_compute_environment.test_environment_1.arn,
        aws_batch_compute_environment.test_environment_2.arn,
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_scheduling_policy = aws.batch.SchedulingPolicy("exampleSchedulingPolicy", fair_share_policy=aws.batch.SchedulingPolicyFairSharePolicyArgs(
    compute_reservation=1,
    share_decay_seconds=3600,
    share_distributions=[aws.batch.SchedulingPolicyFairSharePolicyShareDistributionArgs(
        share_identifier="A1*",
        weight_factor=0.1,
    )],
))
example_job_queue = aws.batch.JobQueue("exampleJobQueue",
    scheduling_policy_arn=example_scheduling_policy.arn,
    state="ENABLED",
    priority=1,
    compute_environments=[
        aws_batch_compute_environment["test_environment_1"]["arn"],
        aws_batch_compute_environment["test_environment_2"]["arn"],
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleSchedulingPolicy = new Aws.Batch.SchedulingPolicy("exampleSchedulingPolicy", new()
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
            },
        },
    });

    var exampleJobQueue = new Aws.Batch.JobQueue("exampleJobQueue", new()
    {
        SchedulingPolicyArn = exampleSchedulingPolicy.Arn,
        State = "ENABLED",
        Priority = 1,
        ComputeEnvironments = new[]
        {
            aws_batch_compute_environment.Test_environment_1.Arn,
            aws_batch_compute_environment.Test_environment_2.Arn,
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
		exampleSchedulingPolicy, err := batch.NewSchedulingPolicy(ctx, "exampleSchedulingPolicy", &batch.SchedulingPolicyArgs{
			FairSharePolicy: &batch.SchedulingPolicyFairSharePolicyArgs{
				ComputeReservation: pulumi.Int(1),
				ShareDecaySeconds:  pulumi.Int(3600),
				ShareDistributions: batch.SchedulingPolicyFairSharePolicyShareDistributionArray{
					&batch.SchedulingPolicyFairSharePolicyShareDistributionArgs{
						ShareIdentifier: pulumi.String("A1*"),
						WeightFactor:    pulumi.Float64(0.1),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = batch.NewJobQueue(ctx, "exampleJobQueue", &batch.JobQueueArgs{
			SchedulingPolicyArn: exampleSchedulingPolicy.Arn,
			State:               pulumi.String("ENABLED"),
			Priority:            pulumi.Int(1),
			ComputeEnvironments: pulumi.StringArray{
				aws_batch_compute_environment.Test_environment_1.Arn,
				aws_batch_compute_environment.Test_environment_2.Arn,
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
import com.pulumi.aws.batch.JobQueue;
import com.pulumi.aws.batch.JobQueueArgs;
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
        var exampleSchedulingPolicy = new SchedulingPolicy("exampleSchedulingPolicy", SchedulingPolicyArgs.builder()        
            .fairSharePolicy(SchedulingPolicyFairSharePolicyArgs.builder()
                .computeReservation(1)
                .shareDecaySeconds(3600)
                .shareDistributions(SchedulingPolicyFairSharePolicyShareDistributionArgs.builder()
                    .shareIdentifier("A1*")
                    .weightFactor(0.1)
                    .build())
                .build())
            .build());

        var exampleJobQueue = new JobQueue("exampleJobQueue", JobQueueArgs.builder()        
            .schedulingPolicyArn(exampleSchedulingPolicy.arn())
            .state("ENABLED")
            .priority(1)
            .computeEnvironments(            
                aws_batch_compute_environment.test_environment_1().arn(),
                aws_batch_compute_environment.test_environment_2().arn())
            .build());

    }
}
```
```yaml
resources:
  exampleSchedulingPolicy:
    type: aws:batch:SchedulingPolicy
    properties:
      fairSharePolicy:
        computeReservation: 1
        shareDecaySeconds: 3600
        shareDistributions:
          - shareIdentifier: A1*
            weightFactor: 0.1
  exampleJobQueue:
    type: aws:batch:JobQueue
    properties:
      schedulingPolicyArn: ${exampleSchedulingPolicy.arn}
      state: ENABLED
      priority: 1
      computeEnvironments:
        - ${aws_batch_compute_environment.test_environment_1.arn}
        - ${aws_batch_compute_environment.test_environment_2.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Batch Job Queue using the `arn`. For example:

```sh
 $ pulumi import aws:batch/jobQueue:JobQueue test_queue arn:aws:batch:us-east-1:123456789012:job-queue/sample
```
 