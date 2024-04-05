Provides a SageMaker Flow Definition resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.FlowDefinition("example", {
    flowDefinitionName: "example",
    roleArn: aws_iam_role.example.arn,
    humanLoopConfig: {
        humanTaskUiArn: aws_sagemaker_human_task_ui.example.arn,
        taskAvailabilityLifetimeInSeconds: 1,
        taskCount: 1,
        taskDescription: "example",
        taskTitle: "example",
        workteamArn: aws_sagemaker_workteam.example.arn,
    },
    outputConfig: {
        s3OutputPath: `s3://${aws_s3_bucket.example.bucket}/`,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.FlowDefinition("example",
    flow_definition_name="example",
    role_arn=aws_iam_role["example"]["arn"],
    human_loop_config=aws.sagemaker.FlowDefinitionHumanLoopConfigArgs(
        human_task_ui_arn=aws_sagemaker_human_task_ui["example"]["arn"],
        task_availability_lifetime_in_seconds=1,
        task_count=1,
        task_description="example",
        task_title="example",
        workteam_arn=aws_sagemaker_workteam["example"]["arn"],
    ),
    output_config=aws.sagemaker.FlowDefinitionOutputConfigArgs(
        s3_output_path=f"s3://{aws_s3_bucket['example']['bucket']}/",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.FlowDefinition("example", new()
    {
        FlowDefinitionName = "example",
        RoleArn = aws_iam_role.Example.Arn,
        HumanLoopConfig = new Aws.Sagemaker.Inputs.FlowDefinitionHumanLoopConfigArgs
        {
            HumanTaskUiArn = aws_sagemaker_human_task_ui.Example.Arn,
            TaskAvailabilityLifetimeInSeconds = 1,
            TaskCount = 1,
            TaskDescription = "example",
            TaskTitle = "example",
            WorkteamArn = aws_sagemaker_workteam.Example.Arn,
        },
        OutputConfig = new Aws.Sagemaker.Inputs.FlowDefinitionOutputConfigArgs
        {
            S3OutputPath = $"s3://{aws_s3_bucket.Example.Bucket}/",
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewFlowDefinition(ctx, "example", &sagemaker.FlowDefinitionArgs{
			FlowDefinitionName: pulumi.String("example"),
			RoleArn:            pulumi.Any(aws_iam_role.Example.Arn),
			HumanLoopConfig: &sagemaker.FlowDefinitionHumanLoopConfigArgs{
				HumanTaskUiArn:                    pulumi.Any(aws_sagemaker_human_task_ui.Example.Arn),
				TaskAvailabilityLifetimeInSeconds: pulumi.Int(1),
				TaskCount:                         pulumi.Int(1),
				TaskDescription:                   pulumi.String("example"),
				TaskTitle:                         pulumi.String("example"),
				WorkteamArn:                       pulumi.Any(aws_sagemaker_workteam.Example.Arn),
			},
			OutputConfig: &sagemaker.FlowDefinitionOutputConfigArgs{
				S3OutputPath: pulumi.String(fmt.Sprintf("s3://%v/", aws_s3_bucket.Example.Bucket)),
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
import com.pulumi.aws.sagemaker.FlowDefinition;
import com.pulumi.aws.sagemaker.FlowDefinitionArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionHumanLoopConfigArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionOutputConfigArgs;
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
        var example = new FlowDefinition("example", FlowDefinitionArgs.builder()        
            .flowDefinitionName("example")
            .roleArn(aws_iam_role.example().arn())
            .humanLoopConfig(FlowDefinitionHumanLoopConfigArgs.builder()
                .humanTaskUiArn(aws_sagemaker_human_task_ui.example().arn())
                .taskAvailabilityLifetimeInSeconds(1)
                .taskCount(1)
                .taskDescription("example")
                .taskTitle("example")
                .workteamArn(aws_sagemaker_workteam.example().arn())
                .build())
            .outputConfig(FlowDefinitionOutputConfigArgs.builder()
                .s3OutputPath(String.format("s3://%s/", aws_s3_bucket.example().bucket()))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:FlowDefinition
    properties:
      flowDefinitionName: example
      roleArn: ${aws_iam_role.example.arn}
      humanLoopConfig:
        humanTaskUiArn: ${aws_sagemaker_human_task_ui.example.arn}
        taskAvailabilityLifetimeInSeconds: 1
        taskCount: 1
        taskDescription: example
        taskTitle: example
        workteamArn: ${aws_sagemaker_workteam.example.arn}
      outputConfig:
        s3OutputPath: s3://${aws_s3_bucket.example.bucket}/
```
{{% /example %}}
{{% example %}}
### Public Workteam Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.FlowDefinition("example", {
    flowDefinitionName: "example",
    roleArn: aws_iam_role.example.arn,
    humanLoopConfig: {
        humanTaskUiArn: aws_sagemaker_human_task_ui.example.arn,
        taskAvailabilityLifetimeInSeconds: 1,
        taskCount: 1,
        taskDescription: "example",
        taskTitle: "example",
        workteamArn: `arn:aws:sagemaker:${data.aws_region.current.name}:394669845002:workteam/public-crowd/default`,
        publicWorkforceTaskPrice: {
            amountInUsd: {
                cents: 1,
                tenthFractionsOfACent: 2,
            },
        },
    },
    outputConfig: {
        s3OutputPath: `s3://${aws_s3_bucket.example.bucket}/`,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.FlowDefinition("example",
    flow_definition_name="example",
    role_arn=aws_iam_role["example"]["arn"],
    human_loop_config=aws.sagemaker.FlowDefinitionHumanLoopConfigArgs(
        human_task_ui_arn=aws_sagemaker_human_task_ui["example"]["arn"],
        task_availability_lifetime_in_seconds=1,
        task_count=1,
        task_description="example",
        task_title="example",
        workteam_arn=f"arn:aws:sagemaker:{data['aws_region']['current']['name']}:394669845002:workteam/public-crowd/default",
        public_workforce_task_price=aws.sagemaker.FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceArgs(
            amount_in_usd=aws.sagemaker.FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsdArgs(
                cents=1,
                tenth_fractions_of_a_cent=2,
            ),
        ),
    ),
    output_config=aws.sagemaker.FlowDefinitionOutputConfigArgs(
        s3_output_path=f"s3://{aws_s3_bucket['example']['bucket']}/",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.FlowDefinition("example", new()
    {
        FlowDefinitionName = "example",
        RoleArn = aws_iam_role.Example.Arn,
        HumanLoopConfig = new Aws.Sagemaker.Inputs.FlowDefinitionHumanLoopConfigArgs
        {
            HumanTaskUiArn = aws_sagemaker_human_task_ui.Example.Arn,
            TaskAvailabilityLifetimeInSeconds = 1,
            TaskCount = 1,
            TaskDescription = "example",
            TaskTitle = "example",
            WorkteamArn = $"arn:aws:sagemaker:{data.Aws_region.Current.Name}:394669845002:workteam/public-crowd/default",
            PublicWorkforceTaskPrice = new Aws.Sagemaker.Inputs.FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceArgs
            {
                AmountInUsd = new Aws.Sagemaker.Inputs.FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsdArgs
                {
                    Cents = 1,
                    TenthFractionsOfACent = 2,
                },
            },
        },
        OutputConfig = new Aws.Sagemaker.Inputs.FlowDefinitionOutputConfigArgs
        {
            S3OutputPath = $"s3://{aws_s3_bucket.Example.Bucket}/",
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewFlowDefinition(ctx, "example", &sagemaker.FlowDefinitionArgs{
			FlowDefinitionName: pulumi.String("example"),
			RoleArn:            pulumi.Any(aws_iam_role.Example.Arn),
			HumanLoopConfig: &sagemaker.FlowDefinitionHumanLoopConfigArgs{
				HumanTaskUiArn:                    pulumi.Any(aws_sagemaker_human_task_ui.Example.Arn),
				TaskAvailabilityLifetimeInSeconds: pulumi.Int(1),
				TaskCount:                         pulumi.Int(1),
				TaskDescription:                   pulumi.String("example"),
				TaskTitle:                         pulumi.String("example"),
				WorkteamArn:                       pulumi.String(fmt.Sprintf("arn:aws:sagemaker:%v:394669845002:workteam/public-crowd/default", data.Aws_region.Current.Name)),
				PublicWorkforceTaskPrice: &sagemaker.FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceArgs{
					AmountInUsd: &sagemaker.FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsdArgs{
						Cents:                 pulumi.Int(1),
						TenthFractionsOfACent: pulumi.Int(2),
					},
				},
			},
			OutputConfig: &sagemaker.FlowDefinitionOutputConfigArgs{
				S3OutputPath: pulumi.String(fmt.Sprintf("s3://%v/", aws_s3_bucket.Example.Bucket)),
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
import com.pulumi.aws.sagemaker.FlowDefinition;
import com.pulumi.aws.sagemaker.FlowDefinitionArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionHumanLoopConfigArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsdArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionOutputConfigArgs;
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
        var example = new FlowDefinition("example", FlowDefinitionArgs.builder()        
            .flowDefinitionName("example")
            .roleArn(aws_iam_role.example().arn())
            .humanLoopConfig(FlowDefinitionHumanLoopConfigArgs.builder()
                .humanTaskUiArn(aws_sagemaker_human_task_ui.example().arn())
                .taskAvailabilityLifetimeInSeconds(1)
                .taskCount(1)
                .taskDescription("example")
                .taskTitle("example")
                .workteamArn(String.format("arn:aws:sagemaker:%s:394669845002:workteam/public-crowd/default", data.aws_region().current().name()))
                .publicWorkforceTaskPrice(FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceArgs.builder()
                    .amountInUsd(FlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsdArgs.builder()
                        .cents(1)
                        .tenthFractionsOfACent(2)
                        .build())
                    .build())
                .build())
            .outputConfig(FlowDefinitionOutputConfigArgs.builder()
                .s3OutputPath(String.format("s3://%s/", aws_s3_bucket.example().bucket()))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:FlowDefinition
    properties:
      flowDefinitionName: example
      roleArn: ${aws_iam_role.example.arn}
      humanLoopConfig:
        humanTaskUiArn: ${aws_sagemaker_human_task_ui.example.arn}
        taskAvailabilityLifetimeInSeconds: 1
        taskCount: 1
        taskDescription: example
        taskTitle: example
        workteamArn: arn:aws:sagemaker:${data.aws_region.current.name}:394669845002:workteam/public-crowd/default
        publicWorkforceTaskPrice:
          amountInUsd:
            cents: 1
            tenthFractionsOfACent: 2
      outputConfig:
        s3OutputPath: s3://${aws_s3_bucket.example.bucket}/
```
{{% /example %}}
{{% example %}}
### Human Loop Activation Config Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.FlowDefinition("example", {
    flowDefinitionName: "example",
    roleArn: aws_iam_role.example.arn,
    humanLoopConfig: {
        humanTaskUiArn: aws_sagemaker_human_task_ui.example.arn,
        taskAvailabilityLifetimeInSeconds: 1,
        taskCount: 1,
        taskDescription: "example",
        taskTitle: "example",
        workteamArn: aws_sagemaker_workteam.example.arn,
    },
    humanLoopRequestSource: {
        awsManagedHumanLoopRequestSource: "AWS/Textract/AnalyzeDocument/Forms/V1",
    },
    humanLoopActivationConfig: {
        humanLoopActivationConditionsConfig: {
            humanLoopActivationConditions: `        {
			"Conditions": [
			  {
				"ConditionType": "Sampling",
				"ConditionParameters": {
				  "RandomSamplingPercentage": 5
				}
			  }
			]
		}
`,
        },
    },
    outputConfig: {
        s3OutputPath: `s3://${aws_s3_bucket.example.bucket}/`,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.FlowDefinition("example",
    flow_definition_name="example",
    role_arn=aws_iam_role["example"]["arn"],
    human_loop_config=aws.sagemaker.FlowDefinitionHumanLoopConfigArgs(
        human_task_ui_arn=aws_sagemaker_human_task_ui["example"]["arn"],
        task_availability_lifetime_in_seconds=1,
        task_count=1,
        task_description="example",
        task_title="example",
        workteam_arn=aws_sagemaker_workteam["example"]["arn"],
    ),
    human_loop_request_source=aws.sagemaker.FlowDefinitionHumanLoopRequestSourceArgs(
        aws_managed_human_loop_request_source="AWS/Textract/AnalyzeDocument/Forms/V1",
    ),
    human_loop_activation_config=aws.sagemaker.FlowDefinitionHumanLoopActivationConfigArgs(
        human_loop_activation_conditions_config=aws.sagemaker.FlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfigArgs(
            human_loop_activation_conditions="""        {
			"Conditions": [
			  {
				"ConditionType": "Sampling",
				"ConditionParameters": {
				  "RandomSamplingPercentage": 5
				}
			  }
			]
		}
""",
        ),
    ),
    output_config=aws.sagemaker.FlowDefinitionOutputConfigArgs(
        s3_output_path=f"s3://{aws_s3_bucket['example']['bucket']}/",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.FlowDefinition("example", new()
    {
        FlowDefinitionName = "example",
        RoleArn = aws_iam_role.Example.Arn,
        HumanLoopConfig = new Aws.Sagemaker.Inputs.FlowDefinitionHumanLoopConfigArgs
        {
            HumanTaskUiArn = aws_sagemaker_human_task_ui.Example.Arn,
            TaskAvailabilityLifetimeInSeconds = 1,
            TaskCount = 1,
            TaskDescription = "example",
            TaskTitle = "example",
            WorkteamArn = aws_sagemaker_workteam.Example.Arn,
        },
        HumanLoopRequestSource = new Aws.Sagemaker.Inputs.FlowDefinitionHumanLoopRequestSourceArgs
        {
            AwsManagedHumanLoopRequestSource = "AWS/Textract/AnalyzeDocument/Forms/V1",
        },
        HumanLoopActivationConfig = new Aws.Sagemaker.Inputs.FlowDefinitionHumanLoopActivationConfigArgs
        {
            HumanLoopActivationConditionsConfig = new Aws.Sagemaker.Inputs.FlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfigArgs
            {
                HumanLoopActivationConditions = @"        {
			""Conditions"": [
			  {
				""ConditionType"": ""Sampling"",
				""ConditionParameters"": {
				  ""RandomSamplingPercentage"": 5
				}
			  }
			]
		}
",
            },
        },
        OutputConfig = new Aws.Sagemaker.Inputs.FlowDefinitionOutputConfigArgs
        {
            S3OutputPath = $"s3://{aws_s3_bucket.Example.Bucket}/",
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewFlowDefinition(ctx, "example", &sagemaker.FlowDefinitionArgs{
			FlowDefinitionName: pulumi.String("example"),
			RoleArn:            pulumi.Any(aws_iam_role.Example.Arn),
			HumanLoopConfig: &sagemaker.FlowDefinitionHumanLoopConfigArgs{
				HumanTaskUiArn:                    pulumi.Any(aws_sagemaker_human_task_ui.Example.Arn),
				TaskAvailabilityLifetimeInSeconds: pulumi.Int(1),
				TaskCount:                         pulumi.Int(1),
				TaskDescription:                   pulumi.String("example"),
				TaskTitle:                         pulumi.String("example"),
				WorkteamArn:                       pulumi.Any(aws_sagemaker_workteam.Example.Arn),
			},
			HumanLoopRequestSource: &sagemaker.FlowDefinitionHumanLoopRequestSourceArgs{
				AwsManagedHumanLoopRequestSource: pulumi.String("AWS/Textract/AnalyzeDocument/Forms/V1"),
			},
			HumanLoopActivationConfig: &sagemaker.FlowDefinitionHumanLoopActivationConfigArgs{
				HumanLoopActivationConditionsConfig: &sagemaker.FlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfigArgs{
					HumanLoopActivationConditions: pulumi.String(`        {
			"Conditions": [
			  {
				"ConditionType": "Sampling",
				"ConditionParameters": {
				  "RandomSamplingPercentage": 5
				}
			  }
			]
		}
`),
				},
			},
			OutputConfig: &sagemaker.FlowDefinitionOutputConfigArgs{
				S3OutputPath: pulumi.String(fmt.Sprintf("s3://%v/", aws_s3_bucket.Example.Bucket)),
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
import com.pulumi.aws.sagemaker.FlowDefinition;
import com.pulumi.aws.sagemaker.FlowDefinitionArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionHumanLoopConfigArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionHumanLoopRequestSourceArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionHumanLoopActivationConfigArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfigArgs;
import com.pulumi.aws.sagemaker.inputs.FlowDefinitionOutputConfigArgs;
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
        var example = new FlowDefinition("example", FlowDefinitionArgs.builder()        
            .flowDefinitionName("example")
            .roleArn(aws_iam_role.example().arn())
            .humanLoopConfig(FlowDefinitionHumanLoopConfigArgs.builder()
                .humanTaskUiArn(aws_sagemaker_human_task_ui.example().arn())
                .taskAvailabilityLifetimeInSeconds(1)
                .taskCount(1)
                .taskDescription("example")
                .taskTitle("example")
                .workteamArn(aws_sagemaker_workteam.example().arn())
                .build())
            .humanLoopRequestSource(FlowDefinitionHumanLoopRequestSourceArgs.builder()
                .awsManagedHumanLoopRequestSource("AWS/Textract/AnalyzeDocument/Forms/V1")
                .build())
            .humanLoopActivationConfig(FlowDefinitionHumanLoopActivationConfigArgs.builder()
                .humanLoopActivationConditionsConfig(FlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfigArgs.builder()
                    .humanLoopActivationConditions("""
        {
			"Conditions": [
			  {
				"ConditionType": "Sampling",
				"ConditionParameters": {
				  "RandomSamplingPercentage": 5
				}
			  }
			]
		}
                    """)
                    .build())
                .build())
            .outputConfig(FlowDefinitionOutputConfigArgs.builder()
                .s3OutputPath(String.format("s3://%s/", aws_s3_bucket.example().bucket()))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:FlowDefinition
    properties:
      flowDefinitionName: example
      roleArn: ${aws_iam_role.example.arn}
      humanLoopConfig:
        humanTaskUiArn: ${aws_sagemaker_human_task_ui.example.arn}
        taskAvailabilityLifetimeInSeconds: 1
        taskCount: 1
        taskDescription: example
        taskTitle: example
        workteamArn: ${aws_sagemaker_workteam.example.arn}
      humanLoopRequestSource:
        awsManagedHumanLoopRequestSource: AWS/Textract/AnalyzeDocument/Forms/V1
      humanLoopActivationConfig:
        humanLoopActivationConditionsConfig:
          humanLoopActivationConditions: |2
                    {
            			"Conditions": [
            			  {
            				"ConditionType": "Sampling",
            				"ConditionParameters": {
            				  "RandomSamplingPercentage": 5
            				}
            			  }
            			]
            		}
      outputConfig:
        s3OutputPath: s3://${aws_s3_bucket.example.bucket}/
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Flow Definitions using the `flow_definition_name`. For example:

```sh
 $ pulumi import aws:sagemaker/flowDefinition:FlowDefinition example example
```
 