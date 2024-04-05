Provides a Glue Workflow resource.
The workflow graph (DAG) can be build using the `aws.glue.Trigger` resource.
See the example below for creating a graph with four nodes (two triggers and two jobs).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Workflow("example", {});
const example_start = new aws.glue.Trigger("example-start", {
    type: "ON_DEMAND",
    workflowName: example.name,
    actions: [{
        jobName: "example-job",
    }],
});
const example_inner = new aws.glue.Trigger("example-inner", {
    type: "CONDITIONAL",
    workflowName: example.name,
    predicate: {
        conditions: [{
            jobName: "example-job",
            state: "SUCCEEDED",
        }],
    },
    actions: [{
        jobName: "another-example-job",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Workflow("example")
example_start = aws.glue.Trigger("example-start",
    type="ON_DEMAND",
    workflow_name=example.name,
    actions=[aws.glue.TriggerActionArgs(
        job_name="example-job",
    )])
example_inner = aws.glue.Trigger("example-inner",
    type="CONDITIONAL",
    workflow_name=example.name,
    predicate=aws.glue.TriggerPredicateArgs(
        conditions=[aws.glue.TriggerPredicateConditionArgs(
            job_name="example-job",
            state="SUCCEEDED",
        )],
    ),
    actions=[aws.glue.TriggerActionArgs(
        job_name="another-example-job",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Workflow("example");

    var example_start = new Aws.Glue.Trigger("example-start", new()
    {
        Type = "ON_DEMAND",
        WorkflowName = example.Name,
        Actions = new[]
        {
            new Aws.Glue.Inputs.TriggerActionArgs
            {
                JobName = "example-job",
            },
        },
    });

    var example_inner = new Aws.Glue.Trigger("example-inner", new()
    {
        Type = "CONDITIONAL",
        WorkflowName = example.Name,
        Predicate = new Aws.Glue.Inputs.TriggerPredicateArgs
        {
            Conditions = new[]
            {
                new Aws.Glue.Inputs.TriggerPredicateConditionArgs
                {
                    JobName = "example-job",
                    State = "SUCCEEDED",
                },
            },
        },
        Actions = new[]
        {
            new Aws.Glue.Inputs.TriggerActionArgs
            {
                JobName = "another-example-job",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := glue.NewWorkflow(ctx, "example", nil)
		if err != nil {
			return err
		}
		_, err = glue.NewTrigger(ctx, "example-start", &glue.TriggerArgs{
			Type:         pulumi.String("ON_DEMAND"),
			WorkflowName: example.Name,
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					JobName: pulumi.String("example-job"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = glue.NewTrigger(ctx, "example-inner", &glue.TriggerArgs{
			Type:         pulumi.String("CONDITIONAL"),
			WorkflowName: example.Name,
			Predicate: &glue.TriggerPredicateArgs{
				Conditions: glue.TriggerPredicateConditionArray{
					&glue.TriggerPredicateConditionArgs{
						JobName: pulumi.String("example-job"),
						State:   pulumi.String("SUCCEEDED"),
					},
				},
			},
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					JobName: pulumi.String("another-example-job"),
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
import com.pulumi.aws.glue.Workflow;
import com.pulumi.aws.glue.Trigger;
import com.pulumi.aws.glue.TriggerArgs;
import com.pulumi.aws.glue.inputs.TriggerActionArgs;
import com.pulumi.aws.glue.inputs.TriggerPredicateArgs;
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
        var example = new Workflow("example");

        var example_start = new Trigger("example-start", TriggerArgs.builder()        
            .type("ON_DEMAND")
            .workflowName(example.name())
            .actions(TriggerActionArgs.builder()
                .jobName("example-job")
                .build())
            .build());

        var example_inner = new Trigger("example-inner", TriggerArgs.builder()        
            .type("CONDITIONAL")
            .workflowName(example.name())
            .predicate(TriggerPredicateArgs.builder()
                .conditions(TriggerPredicateConditionArgs.builder()
                    .jobName("example-job")
                    .state("SUCCEEDED")
                    .build())
                .build())
            .actions(TriggerActionArgs.builder()
                .jobName("another-example-job")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Workflow
  example-start:
    type: aws:glue:Trigger
    properties:
      type: ON_DEMAND
      workflowName: ${example.name}
      actions:
        - jobName: example-job
  example-inner:
    type: aws:glue:Trigger
    properties:
      type: CONDITIONAL
      workflowName: ${example.name}
      predicate:
        conditions:
          - jobName: example-job
            state: SUCCEEDED
      actions:
        - jobName: another-example-job
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Workflows using `name`. For example:

```sh
 $ pulumi import aws:glue/workflow:Workflow MyWorkflow MyWorkflow
```
 