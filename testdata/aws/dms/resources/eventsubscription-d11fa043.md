Provides a DMS (Data Migration Service) event subscription resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dms.EventSubscription("example", {
    enabled: true,
    eventCategories: [
        "creation",
        "failure",
    ],
    snsTopicArn: aws_sns_topic.example.arn,
    sourceIds: [aws_dms_replication_task.example.replication_task_id],
    sourceType: "replication-task",
    tags: {
        Name: "example",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.dms.EventSubscription("example",
    enabled=True,
    event_categories=[
        "creation",
        "failure",
    ],
    sns_topic_arn=aws_sns_topic["example"]["arn"],
    source_ids=[aws_dms_replication_task["example"]["replication_task_id"]],
    source_type="replication-task",
    tags={
        "Name": "example",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Dms.EventSubscription("example", new()
    {
        Enabled = true,
        EventCategories = new[]
        {
            "creation",
            "failure",
        },
        SnsTopicArn = aws_sns_topic.Example.Arn,
        SourceIds = new[]
        {
            aws_dms_replication_task.Example.Replication_task_id,
        },
        SourceType = "replication-task",
        Tags = 
        {
            { "Name", "example" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewEventSubscription(ctx, "example", &dms.EventSubscriptionArgs{
			Enabled: pulumi.Bool(true),
			EventCategories: pulumi.StringArray{
				pulumi.String("creation"),
				pulumi.String("failure"),
			},
			SnsTopicArn: pulumi.Any(aws_sns_topic.Example.Arn),
			SourceIds: pulumi.StringArray{
				aws_dms_replication_task.Example.Replication_task_id,
			},
			SourceType: pulumi.String("replication-task"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example"),
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
import com.pulumi.aws.dms.EventSubscription;
import com.pulumi.aws.dms.EventSubscriptionArgs;
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
        var example = new EventSubscription("example", EventSubscriptionArgs.builder()        
            .enabled(true)
            .eventCategories(            
                "creation",
                "failure")
            .snsTopicArn(aws_sns_topic.example().arn())
            .sourceIds(aws_dms_replication_task.example().replication_task_id())
            .sourceType("replication-task")
            .tags(Map.of("Name", "example"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:dms:EventSubscription
    properties:
      enabled: true
      eventCategories:
        - creation
        - failure
      snsTopicArn: ${aws_sns_topic.example.arn}
      sourceIds:
        - ${aws_dms_replication_task.example.replication_task_id}
      sourceType: replication-task
      tags:
        Name: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import event subscriptions using the `name`. For example:

```sh
 $ pulumi import aws:dms/eventSubscription:EventSubscription test my-awesome-event-subscription
```
 