Provides an AutoScaling Group with Notification support, via SNS Topics. Each of
the `notifications` map to a [Notification Configuration](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeNotificationConfigurations.html) inside Amazon Web
Services, and are applied to each AutoScaling Group you supply.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sns.Topic("example", {});
// arn is an exported attribute
const bar = new aws.autoscaling.Group("bar", {});
// ...
const foo = new aws.autoscaling.Group("foo", {});
// ...
const exampleNotifications = new aws.autoscaling.Notification("exampleNotifications", {
    groupNames: [
        bar.name,
        foo.name,
    ],
    notifications: [
        "autoscaling:EC2_INSTANCE_LAUNCH",
        "autoscaling:EC2_INSTANCE_TERMINATE",
        "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
        "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
    ],
    topicArn: example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sns.Topic("example")
# arn is an exported attribute
bar = aws.autoscaling.Group("bar")
# ...
foo = aws.autoscaling.Group("foo")
# ...
example_notifications = aws.autoscaling.Notification("exampleNotifications",
    group_names=[
        bar.name,
        foo.name,
    ],
    notifications=[
        "autoscaling:EC2_INSTANCE_LAUNCH",
        "autoscaling:EC2_INSTANCE_TERMINATE",
        "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
        "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
    ],
    topic_arn=example.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sns.Topic("example");

    // arn is an exported attribute
    var bar = new Aws.AutoScaling.Group("bar");

    // ...
    var foo = new Aws.AutoScaling.Group("foo");

    // ...
    var exampleNotifications = new Aws.AutoScaling.Notification("exampleNotifications", new()
    {
        GroupNames = new[]
        {
            bar.Name,
            foo.Name,
        },
        Notifications = new[]
        {
            "autoscaling:EC2_INSTANCE_LAUNCH",
            "autoscaling:EC2_INSTANCE_TERMINATE",
            "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
            "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
        },
        TopicArn = example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := sns.NewTopic(ctx, "example", nil)
		if err != nil {
			return err
		}
		bar, err := autoscaling.NewGroup(ctx, "bar", nil)
		if err != nil {
			return err
		}
		foo, err := autoscaling.NewGroup(ctx, "foo", nil)
		if err != nil {
			return err
		}
		_, err = autoscaling.NewNotification(ctx, "exampleNotifications", &autoscaling.NotificationArgs{
			GroupNames: pulumi.StringArray{
				bar.Name,
				foo.Name,
			},
			Notifications: pulumi.StringArray{
				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH"),
				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE"),
				pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH_ERROR"),
				pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE_ERROR"),
			},
			TopicArn: example.Arn,
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
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.autoscaling.Group;
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
        var example = new Topic("example");

        var bar = new Group("bar");

        var foo = new Group("foo");

        var exampleNotifications = new Notification("exampleNotifications", NotificationArgs.builder()        
            .groupNames(            
                bar.name(),
                foo.name())
            .notifications(            
                "autoscaling:EC2_INSTANCE_LAUNCH",
                "autoscaling:EC2_INSTANCE_TERMINATE",
                "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
                "autoscaling:EC2_INSTANCE_TERMINATE_ERROR")
            .topicArn(example.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleNotifications:
    type: aws:autoscaling:Notification
    properties:
      groupNames:
        - ${bar.name}
        - ${foo.name}
      notifications:
        - autoscaling:EC2_INSTANCE_LAUNCH
        - autoscaling:EC2_INSTANCE_TERMINATE
        - autoscaling:EC2_INSTANCE_LAUNCH_ERROR
        - autoscaling:EC2_INSTANCE_TERMINATE_ERROR
      topicArn: ${example.arn}
  example:
    type: aws:sns:Topic
  bar:
    type: aws:autoscaling:Group
  foo:
    type: aws:autoscaling:Group
```
{{% /example %}}
{{% /examples %}}