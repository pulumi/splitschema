Provides an FIS Experiment Template, which can be used to run an experiment.
An experiment template contains one or more actions to run on specified targets during an experiment.
It also contains the stop conditions that prevent the experiment from going out of bounds.
See [Amazon Fault Injection Simulator](https://docs.aws.amazon.com/fis/index.html)
for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.fis.ExperimentTemplate("example", {
    description: "example",
    roleArn: aws_iam_role.example.arn,
    stopConditions: [{
        source: "none",
    }],
    actions: [{
        name: "example-action",
        actionId: "aws:ec2:terminate-instances",
        target: {
            key: "Instances",
            value: "example-target",
        },
    }],
    targets: [{
        name: "example-target",
        resourceType: "aws:ec2:instance",
        selectionMode: "COUNT(1)",
        resourceTags: [{
            key: "env",
            value: "example",
        }],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.fis.ExperimentTemplate("example",
    description="example",
    role_arn=aws_iam_role["example"]["arn"],
    stop_conditions=[aws.fis.ExperimentTemplateStopConditionArgs(
        source="none",
    )],
    actions=[aws.fis.ExperimentTemplateActionArgs(
        name="example-action",
        action_id="aws:ec2:terminate-instances",
        target=aws.fis.ExperimentTemplateActionTargetArgs(
            key="Instances",
            value="example-target",
        ),
    )],
    targets=[aws.fis.ExperimentTemplateTargetArgs(
        name="example-target",
        resource_type="aws:ec2:instance",
        selection_mode="COUNT(1)",
        resource_tags=[aws.fis.ExperimentTemplateTargetResourceTagArgs(
            key="env",
            value="example",
        )],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Fis.ExperimentTemplate("example", new()
    {
        Description = "example",
        RoleArn = aws_iam_role.Example.Arn,
        StopConditions = new[]
        {
            new Aws.Fis.Inputs.ExperimentTemplateStopConditionArgs
            {
                Source = "none",
            },
        },
        Actions = new[]
        {
            new Aws.Fis.Inputs.ExperimentTemplateActionArgs
            {
                Name = "example-action",
                ActionId = "aws:ec2:terminate-instances",
                Target = new Aws.Fis.Inputs.ExperimentTemplateActionTargetArgs
                {
                    Key = "Instances",
                    Value = "example-target",
                },
            },
        },
        Targets = new[]
        {
            new Aws.Fis.Inputs.ExperimentTemplateTargetArgs
            {
                Name = "example-target",
                ResourceType = "aws:ec2:instance",
                SelectionMode = "COUNT(1)",
                ResourceTags = new[]
                {
                    new Aws.Fis.Inputs.ExperimentTemplateTargetResourceTagArgs
                    {
                        Key = "env",
                        Value = "example",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fis.NewExperimentTemplate(ctx, "example", &fis.ExperimentTemplateArgs{
			Description: pulumi.String("example"),
			RoleArn:     pulumi.Any(aws_iam_role.Example.Arn),
			StopConditions: fis.ExperimentTemplateStopConditionArray{
				&fis.ExperimentTemplateStopConditionArgs{
					Source: pulumi.String("none"),
				},
			},
			Actions: fis.ExperimentTemplateActionArray{
				&fis.ExperimentTemplateActionArgs{
					Name:     pulumi.String("example-action"),
					ActionId: pulumi.String("aws:ec2:terminate-instances"),
					Target: &fis.ExperimentTemplateActionTargetArgs{
						Key:   pulumi.String("Instances"),
						Value: pulumi.String("example-target"),
					},
				},
			},
			Targets: fis.ExperimentTemplateTargetArray{
				&fis.ExperimentTemplateTargetArgs{
					Name:          pulumi.String("example-target"),
					ResourceType:  pulumi.String("aws:ec2:instance"),
					SelectionMode: pulumi.String("COUNT(1)"),
					ResourceTags: fis.ExperimentTemplateTargetResourceTagArray{
						&fis.ExperimentTemplateTargetResourceTagArgs{
							Key:   pulumi.String("env"),
							Value: pulumi.String("example"),
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
import com.pulumi.aws.fis.ExperimentTemplate;
import com.pulumi.aws.fis.ExperimentTemplateArgs;
import com.pulumi.aws.fis.inputs.ExperimentTemplateStopConditionArgs;
import com.pulumi.aws.fis.inputs.ExperimentTemplateActionArgs;
import com.pulumi.aws.fis.inputs.ExperimentTemplateActionTargetArgs;
import com.pulumi.aws.fis.inputs.ExperimentTemplateTargetArgs;
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
        var example = new ExperimentTemplate("example", ExperimentTemplateArgs.builder()        
            .description("example")
            .roleArn(aws_iam_role.example().arn())
            .stopConditions(ExperimentTemplateStopConditionArgs.builder()
                .source("none")
                .build())
            .actions(ExperimentTemplateActionArgs.builder()
                .name("example-action")
                .actionId("aws:ec2:terminate-instances")
                .target(ExperimentTemplateActionTargetArgs.builder()
                    .key("Instances")
                    .value("example-target")
                    .build())
                .build())
            .targets(ExperimentTemplateTargetArgs.builder()
                .name("example-target")
                .resourceType("aws:ec2:instance")
                .selectionMode("COUNT(1)")
                .resourceTags(ExperimentTemplateTargetResourceTagArgs.builder()
                    .key("env")
                    .value("example")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:fis:ExperimentTemplate
    properties:
      description: example
      roleArn: ${aws_iam_role.example.arn}
      stopConditions:
        - source: none
      actions:
        - name: example-action
          actionId: aws:ec2:terminate-instances
          target:
            key: Instances
            value: example-target
      targets:
        - name: example-target
          resourceType: aws:ec2:instance
          selectionMode: COUNT(1)
          resourceTags:
            - key: env
              value: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import FIS Experiment Templates using the `id`. For example:

```sh
 $ pulumi import aws:fis/experimentTemplate:ExperimentTemplate template EXT123AbCdEfGhIjK
```
 