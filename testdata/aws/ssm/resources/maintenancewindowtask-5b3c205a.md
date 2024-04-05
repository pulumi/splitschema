Provides an SSM Maintenance Window Task resource

{{% examples %}}
## Example Usage
{{% example %}}
### Automation Tasks

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssm.MaintenanceWindowTask("example", {
    maxConcurrency: "2",
    maxErrors: "1",
    priority: 1,
    taskArn: "AWS-RestartEC2Instance",
    taskType: "AUTOMATION",
    windowId: aws_ssm_maintenance_window.example.id,
    targets: [{
        key: "InstanceIds",
        values: [aws_instance.example.id],
    }],
    taskInvocationParameters: {
        automationParameters: {
            documentVersion: "$LATEST",
            parameters: [{
                name: "InstanceId",
                values: [aws_instance.example.id],
            }],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssm.MaintenanceWindowTask("example",
    max_concurrency="2",
    max_errors="1",
    priority=1,
    task_arn="AWS-RestartEC2Instance",
    task_type="AUTOMATION",
    window_id=aws_ssm_maintenance_window["example"]["id"],
    targets=[aws.ssm.MaintenanceWindowTaskTargetArgs(
        key="InstanceIds",
        values=[aws_instance["example"]["id"]],
    )],
    task_invocation_parameters=aws.ssm.MaintenanceWindowTaskTaskInvocationParametersArgs(
        automation_parameters=aws.ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersArgs(
            document_version="$LATEST",
            parameters=[aws.ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterArgs(
                name="InstanceId",
                values=[aws_instance["example"]["id"]],
            )],
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
    var example = new Aws.Ssm.MaintenanceWindowTask("example", new()
    {
        MaxConcurrency = "2",
        MaxErrors = "1",
        Priority = 1,
        TaskArn = "AWS-RestartEC2Instance",
        TaskType = "AUTOMATION",
        WindowId = aws_ssm_maintenance_window.Example.Id,
        Targets = new[]
        {
            new Aws.Ssm.Inputs.MaintenanceWindowTaskTargetArgs
            {
                Key = "InstanceIds",
                Values = new[]
                {
                    aws_instance.Example.Id,
                },
            },
        },
        TaskInvocationParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersArgs
        {
            AutomationParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersArgs
            {
                DocumentVersion = "$LATEST",
                Parameters = new[]
                {
                    new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterArgs
                    {
                        Name = "InstanceId",
                        Values = new[]
                        {
                            aws_instance.Example.Id,
                        },
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewMaintenanceWindowTask(ctx, "example", &ssm.MaintenanceWindowTaskArgs{
			MaxConcurrency: pulumi.String("2"),
			MaxErrors:      pulumi.String("1"),
			Priority:       pulumi.Int(1),
			TaskArn:        pulumi.String("AWS-RestartEC2Instance"),
			TaskType:       pulumi.String("AUTOMATION"),
			WindowId:       pulumi.Any(aws_ssm_maintenance_window.Example.Id),
			Targets: ssm.MaintenanceWindowTaskTargetArray{
				&ssm.MaintenanceWindowTaskTargetArgs{
					Key: pulumi.String("InstanceIds"),
					Values: pulumi.StringArray{
						aws_instance.Example.Id,
					},
				},
			},
			TaskInvocationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersArgs{
				AutomationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersArgs{
					DocumentVersion: pulumi.String("$LATEST"),
					Parameters: ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterArray{
						&ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterArgs{
							Name: pulumi.String("InstanceId"),
							Values: pulumi.StringArray{
								aws_instance.Example.Id,
							},
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
import com.pulumi.aws.ssm.MaintenanceWindowTask;
import com.pulumi.aws.ssm.MaintenanceWindowTaskArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTaskTargetArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTaskTaskInvocationParametersArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersArgs;
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
        var example = new MaintenanceWindowTask("example", MaintenanceWindowTaskArgs.builder()        
            .maxConcurrency(2)
            .maxErrors(1)
            .priority(1)
            .taskArn("AWS-RestartEC2Instance")
            .taskType("AUTOMATION")
            .windowId(aws_ssm_maintenance_window.example().id())
            .targets(MaintenanceWindowTaskTargetArgs.builder()
                .key("InstanceIds")
                .values(aws_instance.example().id())
                .build())
            .taskInvocationParameters(MaintenanceWindowTaskTaskInvocationParametersArgs.builder()
                .automationParameters(MaintenanceWindowTaskTaskInvocationParametersAutomationParametersArgs.builder()
                    .documentVersion("$LATEST")
                    .parameters(MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterArgs.builder()
                        .name("InstanceId")
                        .values(aws_instance.example().id())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssm:MaintenanceWindowTask
    properties:
      maxConcurrency: 2
      maxErrors: 1
      priority: 1
      taskArn: AWS-RestartEC2Instance
      taskType: AUTOMATION
      windowId: ${aws_ssm_maintenance_window.example.id}
      targets:
        - key: InstanceIds
          values:
            - ${aws_instance.example.id}
      taskInvocationParameters:
        automationParameters:
          documentVersion: $LATEST
          parameters:
            - name: InstanceId
              values:
                - ${aws_instance.example.id}
```
{{% /example %}}
{{% example %}}
### Run Command Tasks

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssm.MaintenanceWindowTask("example", {
    maxConcurrency: "2",
    maxErrors: "1",
    priority: 1,
    taskArn: "AWS-RunShellScript",
    taskType: "RUN_COMMAND",
    windowId: aws_ssm_maintenance_window.example.id,
    targets: [{
        key: "InstanceIds",
        values: [aws_instance.example.id],
    }],
    taskInvocationParameters: {
        runCommandParameters: {
            outputS3Bucket: aws_s3_bucket.example.id,
            outputS3KeyPrefix: "output",
            serviceRoleArn: aws_iam_role.example.arn,
            timeoutSeconds: 600,
            notificationConfig: {
                notificationArn: aws_sns_topic.example.arn,
                notificationEvents: ["All"],
                notificationType: "Command",
            },
            parameters: [{
                name: "commands",
                values: ["date"],
            }],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssm.MaintenanceWindowTask("example",
    max_concurrency="2",
    max_errors="1",
    priority=1,
    task_arn="AWS-RunShellScript",
    task_type="RUN_COMMAND",
    window_id=aws_ssm_maintenance_window["example"]["id"],
    targets=[aws.ssm.MaintenanceWindowTaskTargetArgs(
        key="InstanceIds",
        values=[aws_instance["example"]["id"]],
    )],
    task_invocation_parameters=aws.ssm.MaintenanceWindowTaskTaskInvocationParametersArgs(
        run_command_parameters=aws.ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersArgs(
            output_s3_bucket=aws_s3_bucket["example"]["id"],
            output_s3_key_prefix="output",
            service_role_arn=aws_iam_role["example"]["arn"],
            timeout_seconds=600,
            notification_config=aws.ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigArgs(
                notification_arn=aws_sns_topic["example"]["arn"],
                notification_events=["All"],
                notification_type="Command",
            ),
            parameters=[aws.ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterArgs(
                name="commands",
                values=["date"],
            )],
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
    var example = new Aws.Ssm.MaintenanceWindowTask("example", new()
    {
        MaxConcurrency = "2",
        MaxErrors = "1",
        Priority = 1,
        TaskArn = "AWS-RunShellScript",
        TaskType = "RUN_COMMAND",
        WindowId = aws_ssm_maintenance_window.Example.Id,
        Targets = new[]
        {
            new Aws.Ssm.Inputs.MaintenanceWindowTaskTargetArgs
            {
                Key = "InstanceIds",
                Values = new[]
                {
                    aws_instance.Example.Id,
                },
            },
        },
        TaskInvocationParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersArgs
        {
            RunCommandParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersArgs
            {
                OutputS3Bucket = aws_s3_bucket.Example.Id,
                OutputS3KeyPrefix = "output",
                ServiceRoleArn = aws_iam_role.Example.Arn,
                TimeoutSeconds = 600,
                NotificationConfig = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigArgs
                {
                    NotificationArn = aws_sns_topic.Example.Arn,
                    NotificationEvents = new[]
                    {
                        "All",
                    },
                    NotificationType = "Command",
                },
                Parameters = new[]
                {
                    new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterArgs
                    {
                        Name = "commands",
                        Values = new[]
                        {
                            "date",
                        },
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewMaintenanceWindowTask(ctx, "example", &ssm.MaintenanceWindowTaskArgs{
			MaxConcurrency: pulumi.String("2"),
			MaxErrors:      pulumi.String("1"),
			Priority:       pulumi.Int(1),
			TaskArn:        pulumi.String("AWS-RunShellScript"),
			TaskType:       pulumi.String("RUN_COMMAND"),
			WindowId:       pulumi.Any(aws_ssm_maintenance_window.Example.Id),
			Targets: ssm.MaintenanceWindowTaskTargetArray{
				&ssm.MaintenanceWindowTaskTargetArgs{
					Key: pulumi.String("InstanceIds"),
					Values: pulumi.StringArray{
						aws_instance.Example.Id,
					},
				},
			},
			TaskInvocationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersArgs{
				RunCommandParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersArgs{
					OutputS3Bucket:    pulumi.Any(aws_s3_bucket.Example.Id),
					OutputS3KeyPrefix: pulumi.String("output"),
					ServiceRoleArn:    pulumi.Any(aws_iam_role.Example.Arn),
					TimeoutSeconds:    pulumi.Int(600),
					NotificationConfig: &ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigArgs{
						NotificationArn: pulumi.Any(aws_sns_topic.Example.Arn),
						NotificationEvents: pulumi.StringArray{
							pulumi.String("All"),
						},
						NotificationType: pulumi.String("Command"),
					},
					Parameters: ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterArray{
						&ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterArgs{
							Name: pulumi.String("commands"),
							Values: pulumi.StringArray{
								pulumi.String("date"),
							},
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
import com.pulumi.aws.ssm.MaintenanceWindowTask;
import com.pulumi.aws.ssm.MaintenanceWindowTaskArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTaskTargetArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTaskTaskInvocationParametersArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigArgs;
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
        var example = new MaintenanceWindowTask("example", MaintenanceWindowTaskArgs.builder()        
            .maxConcurrency(2)
            .maxErrors(1)
            .priority(1)
            .taskArn("AWS-RunShellScript")
            .taskType("RUN_COMMAND")
            .windowId(aws_ssm_maintenance_window.example().id())
            .targets(MaintenanceWindowTaskTargetArgs.builder()
                .key("InstanceIds")
                .values(aws_instance.example().id())
                .build())
            .taskInvocationParameters(MaintenanceWindowTaskTaskInvocationParametersArgs.builder()
                .runCommandParameters(MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersArgs.builder()
                    .outputS3Bucket(aws_s3_bucket.example().id())
                    .outputS3KeyPrefix("output")
                    .serviceRoleArn(aws_iam_role.example().arn())
                    .timeoutSeconds(600)
                    .notificationConfig(MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigArgs.builder()
                        .notificationArn(aws_sns_topic.example().arn())
                        .notificationEvents("All")
                        .notificationType("Command")
                        .build())
                    .parameters(MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterArgs.builder()
                        .name("commands")
                        .values("date")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssm:MaintenanceWindowTask
    properties:
      maxConcurrency: 2
      maxErrors: 1
      priority: 1
      taskArn: AWS-RunShellScript
      taskType: RUN_COMMAND
      windowId: ${aws_ssm_maintenance_window.example.id}
      targets:
        - key: InstanceIds
          values:
            - ${aws_instance.example.id}
      taskInvocationParameters:
        runCommandParameters:
          outputS3Bucket: ${aws_s3_bucket.example.id}
          outputS3KeyPrefix: output
          serviceRoleArn: ${aws_iam_role.example.arn}
          timeoutSeconds: 600
          notificationConfig:
            notificationArn: ${aws_sns_topic.example.arn}
            notificationEvents:
              - All
            notificationType: Command
          parameters:
            - name: commands
              values:
                - date
```
{{% /example %}}
{{% example %}}
### Step Function Tasks

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssm.MaintenanceWindowTask("example", {
    maxConcurrency: "2",
    maxErrors: "1",
    priority: 1,
    taskArn: aws_sfn_activity.example.id,
    taskType: "STEP_FUNCTIONS",
    windowId: aws_ssm_maintenance_window.example.id,
    targets: [{
        key: "InstanceIds",
        values: [aws_instance.example.id],
    }],
    taskInvocationParameters: {
        stepFunctionsParameters: {
            input: "{\"key1\":\"value1\"}",
            name: "example",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssm.MaintenanceWindowTask("example",
    max_concurrency="2",
    max_errors="1",
    priority=1,
    task_arn=aws_sfn_activity["example"]["id"],
    task_type="STEP_FUNCTIONS",
    window_id=aws_ssm_maintenance_window["example"]["id"],
    targets=[aws.ssm.MaintenanceWindowTaskTargetArgs(
        key="InstanceIds",
        values=[aws_instance["example"]["id"]],
    )],
    task_invocation_parameters=aws.ssm.MaintenanceWindowTaskTaskInvocationParametersArgs(
        step_functions_parameters=aws.ssm.MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersArgs(
            input="{\"key1\":\"value1\"}",
            name="example",
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
    var example = new Aws.Ssm.MaintenanceWindowTask("example", new()
    {
        MaxConcurrency = "2",
        MaxErrors = "1",
        Priority = 1,
        TaskArn = aws_sfn_activity.Example.Id,
        TaskType = "STEP_FUNCTIONS",
        WindowId = aws_ssm_maintenance_window.Example.Id,
        Targets = new[]
        {
            new Aws.Ssm.Inputs.MaintenanceWindowTaskTargetArgs
            {
                Key = "InstanceIds",
                Values = new[]
                {
                    aws_instance.Example.Id,
                },
            },
        },
        TaskInvocationParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersArgs
        {
            StepFunctionsParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersArgs
            {
                Input = "{\"key1\":\"value1\"}",
                Name = "example",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewMaintenanceWindowTask(ctx, "example", &ssm.MaintenanceWindowTaskArgs{
			MaxConcurrency: pulumi.String("2"),
			MaxErrors:      pulumi.String("1"),
			Priority:       pulumi.Int(1),
			TaskArn:        pulumi.Any(aws_sfn_activity.Example.Id),
			TaskType:       pulumi.String("STEP_FUNCTIONS"),
			WindowId:       pulumi.Any(aws_ssm_maintenance_window.Example.Id),
			Targets: ssm.MaintenanceWindowTaskTargetArray{
				&ssm.MaintenanceWindowTaskTargetArgs{
					Key: pulumi.String("InstanceIds"),
					Values: pulumi.StringArray{
						aws_instance.Example.Id,
					},
				},
			},
			TaskInvocationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersArgs{
				StepFunctionsParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersArgs{
					Input: pulumi.String("{\"key1\":\"value1\"}"),
					Name:  pulumi.String("example"),
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
import com.pulumi.aws.ssm.MaintenanceWindowTask;
import com.pulumi.aws.ssm.MaintenanceWindowTaskArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTaskTargetArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTaskTaskInvocationParametersArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersArgs;
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
        var example = new MaintenanceWindowTask("example", MaintenanceWindowTaskArgs.builder()        
            .maxConcurrency(2)
            .maxErrors(1)
            .priority(1)
            .taskArn(aws_sfn_activity.example().id())
            .taskType("STEP_FUNCTIONS")
            .windowId(aws_ssm_maintenance_window.example().id())
            .targets(MaintenanceWindowTaskTargetArgs.builder()
                .key("InstanceIds")
                .values(aws_instance.example().id())
                .build())
            .taskInvocationParameters(MaintenanceWindowTaskTaskInvocationParametersArgs.builder()
                .stepFunctionsParameters(MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersArgs.builder()
                    .input("{\"key1\":\"value1\"}")
                    .name("example")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssm:MaintenanceWindowTask
    properties:
      maxConcurrency: 2
      maxErrors: 1
      priority: 1
      taskArn: ${aws_sfn_activity.example.id}
      taskType: STEP_FUNCTIONS
      windowId: ${aws_ssm_maintenance_window.example.id}
      targets:
        - key: InstanceIds
          values:
            - ${aws_instance.example.id}
      taskInvocationParameters:
        stepFunctionsParameters:
          input: '{"key1":"value1"}'
          name: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS Maintenance Window Task using the `window_id` and `window_task_id` separated by `/`. For example:

```sh
 $ pulumi import aws:ssm/maintenanceWindowTask:MaintenanceWindowTask task <window_id>/<window_task_id>
```
 