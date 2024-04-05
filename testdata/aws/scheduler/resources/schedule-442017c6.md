Provides an EventBridge Scheduler Schedule resource.

You can find out more about EventBridge Scheduler in the [User Guide](https://docs.aws.amazon.com/scheduler/latest/UserGuide/what-is-scheduler.html).

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.scheduler.Schedule("example", {
    groupName: "default",
    flexibleTimeWindow: {
        mode: "OFF",
    },
    scheduleExpression: "rate(1 hours)",
    target: {
        arn: aws_sqs_queue.example.arn,
        roleArn: aws_iam_role.example.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.scheduler.Schedule("example",
    group_name="default",
    flexible_time_window=aws.scheduler.ScheduleFlexibleTimeWindowArgs(
        mode="OFF",
    ),
    schedule_expression="rate(1 hours)",
    target=aws.scheduler.ScheduleTargetArgs(
        arn=aws_sqs_queue["example"]["arn"],
        role_arn=aws_iam_role["example"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Scheduler.Schedule("example", new()
    {
        GroupName = "default",
        FlexibleTimeWindow = new Aws.Scheduler.Inputs.ScheduleFlexibleTimeWindowArgs
        {
            Mode = "OFF",
        },
        ScheduleExpression = "rate(1 hours)",
        Target = new Aws.Scheduler.Inputs.ScheduleTargetArgs
        {
            Arn = aws_sqs_queue.Example.Arn,
            RoleArn = aws_iam_role.Example.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/scheduler"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := scheduler.NewSchedule(ctx, "example", &scheduler.ScheduleArgs{
			GroupName: pulumi.String("default"),
			FlexibleTimeWindow: &scheduler.ScheduleFlexibleTimeWindowArgs{
				Mode: pulumi.String("OFF"),
			},
			ScheduleExpression: pulumi.String("rate(1 hours)"),
			Target: &scheduler.ScheduleTargetArgs{
				Arn:     pulumi.Any(aws_sqs_queue.Example.Arn),
				RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
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
import com.pulumi.aws.scheduler.Schedule;
import com.pulumi.aws.scheduler.ScheduleArgs;
import com.pulumi.aws.scheduler.inputs.ScheduleFlexibleTimeWindowArgs;
import com.pulumi.aws.scheduler.inputs.ScheduleTargetArgs;
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
        var example = new Schedule("example", ScheduleArgs.builder()        
            .groupName("default")
            .flexibleTimeWindow(ScheduleFlexibleTimeWindowArgs.builder()
                .mode("OFF")
                .build())
            .scheduleExpression("rate(1 hours)")
            .target(ScheduleTargetArgs.builder()
                .arn(aws_sqs_queue.example().arn())
                .roleArn(aws_iam_role.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:scheduler:Schedule
    properties:
      groupName: default
      flexibleTimeWindow:
        mode: OFF
      scheduleExpression: rate(1 hours)
      target:
        arn: ${aws_sqs_queue.example.arn}
        roleArn: ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% example %}}
### Universal Target

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleQueue = new aws.sqs.Queue("exampleQueue", {});
const exampleSchedule = new aws.scheduler.Schedule("exampleSchedule", {
    flexibleTimeWindow: {
        mode: "OFF",
    },
    scheduleExpression: "rate(1 hours)",
    target: {
        arn: "arn:aws:scheduler:::aws-sdk:sqs:sendMessage",
        roleArn: aws_iam_role.example.arn,
        input: exampleQueue.url.apply(url => JSON.stringify({
            MessageBody: "Greetings, programs!",
            QueueUrl: url,
        })),
    },
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_queue = aws.sqs.Queue("exampleQueue")
example_schedule = aws.scheduler.Schedule("exampleSchedule",
    flexible_time_window=aws.scheduler.ScheduleFlexibleTimeWindowArgs(
        mode="OFF",
    ),
    schedule_expression="rate(1 hours)",
    target=aws.scheduler.ScheduleTargetArgs(
        arn="arn:aws:scheduler:::aws-sdk:sqs:sendMessage",
        role_arn=aws_iam_role["example"]["arn"],
        input=example_queue.url.apply(lambda url: json.dumps({
            "MessageBody": "Greetings, programs!",
            "QueueUrl": url,
        })),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleQueue = new Aws.Sqs.Queue("exampleQueue");

    var exampleSchedule = new Aws.Scheduler.Schedule("exampleSchedule", new()
    {
        FlexibleTimeWindow = new Aws.Scheduler.Inputs.ScheduleFlexibleTimeWindowArgs
        {
            Mode = "OFF",
        },
        ScheduleExpression = "rate(1 hours)",
        Target = new Aws.Scheduler.Inputs.ScheduleTargetArgs
        {
            Arn = "arn:aws:scheduler:::aws-sdk:sqs:sendMessage",
            RoleArn = aws_iam_role.Example.Arn,
            Input = exampleQueue.Url.Apply(url => JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["MessageBody"] = "Greetings, programs!",
                ["QueueUrl"] = url,
            })),
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/scheduler"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleQueue, err := sqs.NewQueue(ctx, "exampleQueue", nil)
		if err != nil {
			return err
		}
		_, err = scheduler.NewSchedule(ctx, "exampleSchedule", &scheduler.ScheduleArgs{
			FlexibleTimeWindow: &scheduler.ScheduleFlexibleTimeWindowArgs{
				Mode: pulumi.String("OFF"),
			},
			ScheduleExpression: pulumi.String("rate(1 hours)"),
			Target: &scheduler.ScheduleTargetArgs{
				Arn:     pulumi.String("arn:aws:scheduler:::aws-sdk:sqs:sendMessage"),
				RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
				Input: exampleQueue.Url.ApplyT(func(url string) (pulumi.String, error) {
					var _zero pulumi.String
					tmpJSON0, err := json.Marshal(map[string]interface{}{
						"MessageBody": "Greetings, programs!",
						"QueueUrl":    url,
					})
					if err != nil {
						return _zero, err
					}
					json0 := string(tmpJSON0)
					return pulumi.String(json0), nil
				}).(pulumi.StringOutput),
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
import com.pulumi.aws.sqs.Queue;
import com.pulumi.aws.scheduler.Schedule;
import com.pulumi.aws.scheduler.ScheduleArgs;
import com.pulumi.aws.scheduler.inputs.ScheduleFlexibleTimeWindowArgs;
import com.pulumi.aws.scheduler.inputs.ScheduleTargetArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var exampleQueue = new Queue("exampleQueue");

        var exampleSchedule = new Schedule("exampleSchedule", ScheduleArgs.builder()        
            .flexibleTimeWindow(ScheduleFlexibleTimeWindowArgs.builder()
                .mode("OFF")
                .build())
            .scheduleExpression("rate(1 hours)")
            .target(ScheduleTargetArgs.builder()
                .arn("arn:aws:scheduler:::aws-sdk:sqs:sendMessage")
                .roleArn(aws_iam_role.example().arn())
                .input(exampleQueue.url().applyValue(url -> serializeJson(
                    jsonObject(
                        jsonProperty("MessageBody", "Greetings, programs!"),
                        jsonProperty("QueueUrl", url)
                    ))))
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleQueue:
    type: aws:sqs:Queue
  exampleSchedule:
    type: aws:scheduler:Schedule
    properties:
      flexibleTimeWindow:
        mode: OFF
      scheduleExpression: rate(1 hours)
      target:
        arn: arn:aws:scheduler:::aws-sdk:sqs:sendMessage
        roleArn: ${aws_iam_role.example.arn}
        input:
          fn::toJSON:
            MessageBody: Greetings, programs!
            QueueUrl: ${exampleQueue.url}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import schedules using the combination `group_name/name`. For example:

```sh
 $ pulumi import aws:scheduler/schedule:Schedule example my-schedule-group/my-schedule
```
 