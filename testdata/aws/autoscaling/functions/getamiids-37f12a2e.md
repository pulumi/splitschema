The Autoscaling Groups data source allows access to the list of AWS
ASGs within a specific region. This will allow you to pass a list of AutoScaling Groups to other resources.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const groups = aws.autoscaling.getAmiIds({
    filters: [
        {
            name: "tag:Team",
            values: ["Pets"],
        },
        {
            name: "tag-key",
            values: ["Environment"],
        },
    ],
});
const slackNotifications = new aws.autoscaling.Notification("slackNotifications", {
    groupNames: groups.then(groups => groups.names),
    notifications: [
        "autoscaling:EC2_INSTANCE_LAUNCH",
        "autoscaling:EC2_INSTANCE_TERMINATE",
        "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
        "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
    ],
    topicArn: "TOPIC ARN",
});
```
```python
import pulumi
import pulumi_aws as aws

groups = aws.autoscaling.get_ami_ids(filters=[
    aws.autoscaling.GetAmiIdsFilterArgs(
        name="tag:Team",
        values=["Pets"],
    ),
    aws.autoscaling.GetAmiIdsFilterArgs(
        name="tag-key",
        values=["Environment"],
    ),
])
slack_notifications = aws.autoscaling.Notification("slackNotifications",
    group_names=groups.names,
    notifications=[
        "autoscaling:EC2_INSTANCE_LAUNCH",
        "autoscaling:EC2_INSTANCE_TERMINATE",
        "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
        "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
    ],
    topic_arn="TOPIC ARN")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var groups = Aws.AutoScaling.GetAmiIds.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.AutoScaling.Inputs.GetAmiIdsFilterInputArgs
            {
                Name = "tag:Team",
                Values = new[]
                {
                    "Pets",
                },
            },
            new Aws.AutoScaling.Inputs.GetAmiIdsFilterInputArgs
            {
                Name = "tag-key",
                Values = new[]
                {
                    "Environment",
                },
            },
        },
    });

    var slackNotifications = new Aws.AutoScaling.Notification("slackNotifications", new()
    {
        GroupNames = groups.Apply(getAmiIdsResult => getAmiIdsResult.Names),
        Notifications = new[]
        {
            "autoscaling:EC2_INSTANCE_LAUNCH",
            "autoscaling:EC2_INSTANCE_TERMINATE",
            "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
            "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
        },
        TopicArn = "TOPIC ARN",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		groups, err := autoscaling.GetAmiIds(ctx, &autoscaling.GetAmiIdsArgs{
			Filters: []autoscaling.GetAmiIdsFilter{
				{
					Name: "tag:Team",
					Values: []string{
						"Pets",
					},
				},
				{
					Name: "tag-key",
					Values: []string{
						"Environment",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = autoscaling.NewNotification(ctx, "slackNotifications", &autoscaling.NotificationArgs{
			GroupNames: interface{}(groups.Names),
			Notifications: pulumi.StringArray{
				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH"),
				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE"),
				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH_ERROR"),
				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE_ERROR"),
			},
			TopicArn: pulumi.String("TOPIC ARN"),
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
import com.pulumi.aws.autoscaling.AutoscalingFunctions;
import com.pulumi.aws.autoscaling.inputs.GetAmiIdsArgs;
import com.pulumi.aws.autoscaling.Notification;
import com.pulumi.aws.autoscaling.NotificationArgs;
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
        final var groups = AutoscalingFunctions.getAmiIds(GetAmiIdsArgs.builder()
            .filters(            
                GetAmiIdsFilterArgs.builder()
                    .name("tag:Team")
                    .values("Pets")
                    .build(),
                GetAmiIdsFilterArgs.builder()
                    .name("tag-key")
                    .values("Environment")
                    .build())
            .build());

        var slackNotifications = new Notification("slackNotifications", NotificationArgs.builder()        
            .groupNames(groups.applyValue(getAmiIdsResult -> getAmiIdsResult.names()))
            .notifications(            
                "autoscaling:EC2_INSTANCE_LAUNCH",
                "autoscaling:EC2_INSTANCE_TERMINATE",
                "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
                "autoscaling:EC2_INSTANCE_TERMINATE_ERROR")
            .topicArn("TOPIC ARN")
            .build());

    }
}
```
```yaml
resources:
  slackNotifications:
    type: aws:autoscaling:Notification
    properties:
      groupNames: ${groups.names}
      notifications:
        - autoscaling:EC2_INSTANCE_LAUNCH
        - autoscaling:EC2_INSTANCE_TERMINATE
        - autoscaling:EC2_INSTANCE_LAUNCH_ERROR
        - autoscaling:EC2_INSTANCE_TERMINATE_ERROR
      topicArn: TOPIC ARN
variables:
  groups:
    fn::invoke:
      Function: aws:autoscaling:getAmiIds
      Arguments:
        filters:
          - name: tag:Team
            values:
              - Pets
          - name: tag-key
            values:
              - Environment
```
{{% /example %}}
{{% /examples %}}