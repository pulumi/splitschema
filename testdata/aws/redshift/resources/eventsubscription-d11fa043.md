Provides a Redshift event subscription resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCluster = new aws.redshift.Cluster("defaultCluster", {
    clusterIdentifier: "default",
    databaseName: "default",
});
// ...
const defaultTopic = new aws.sns.Topic("defaultTopic", {});
const defaultEventSubscription = new aws.redshift.EventSubscription("defaultEventSubscription", {
    snsTopicArn: defaultTopic.arn,
    sourceType: "cluster",
    sourceIds: [defaultCluster.id],
    severity: "INFO",
    eventCategories: [
        "configuration",
        "management",
        "monitoring",
        "security",
    ],
    tags: {
        Name: "default",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

default_cluster = aws.redshift.Cluster("defaultCluster",
    cluster_identifier="default",
    database_name="default")
# ...
default_topic = aws.sns.Topic("defaultTopic")
default_event_subscription = aws.redshift.EventSubscription("defaultEventSubscription",
    sns_topic_arn=default_topic.arn,
    source_type="cluster",
    source_ids=[default_cluster.id],
    severity="INFO",
    event_categories=[
        "configuration",
        "management",
        "monitoring",
        "security",
    ],
    tags={
        "Name": "default",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultCluster = new Aws.RedShift.Cluster("defaultCluster", new()
    {
        ClusterIdentifier = "default",
        DatabaseName = "default",
    });

    // ...
    var defaultTopic = new Aws.Sns.Topic("defaultTopic");

    var defaultEventSubscription = new Aws.RedShift.EventSubscription("defaultEventSubscription", new()
    {
        SnsTopicArn = defaultTopic.Arn,
        SourceType = "cluster",
        SourceIds = new[]
        {
            defaultCluster.Id,
        },
        Severity = "INFO",
        EventCategories = new[]
        {
            "configuration",
            "management",
            "monitoring",
            "security",
        },
        Tags = 
        {
            { "Name", "default" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultCluster, err := redshift.NewCluster(ctx, "defaultCluster", &redshift.ClusterArgs{
			ClusterIdentifier: pulumi.String("default"),
			DatabaseName:      pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		defaultTopic, err := sns.NewTopic(ctx, "defaultTopic", nil)
		if err != nil {
			return err
		}
		_, err = redshift.NewEventSubscription(ctx, "defaultEventSubscription", &redshift.EventSubscriptionArgs{
			SnsTopicArn: defaultTopic.Arn,
			SourceType:  pulumi.String("cluster"),
			SourceIds: pulumi.StringArray{
				defaultCluster.ID(),
			},
			Severity: pulumi.String("INFO"),
			EventCategories: pulumi.StringArray{
				pulumi.String("configuration"),
				pulumi.String("management"),
				pulumi.String("monitoring"),
				pulumi.String("security"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("default"),
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
import com.pulumi.aws.redshift.Cluster;
import com.pulumi.aws.redshift.ClusterArgs;
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.redshift.EventSubscription;
import com.pulumi.aws.redshift.EventSubscriptionArgs;
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
        var defaultCluster = new Cluster("defaultCluster", ClusterArgs.builder()        
            .clusterIdentifier("default")
            .databaseName("default")
            .build());

        var defaultTopic = new Topic("defaultTopic");

        var defaultEventSubscription = new EventSubscription("defaultEventSubscription", EventSubscriptionArgs.builder()        
            .snsTopicArn(defaultTopic.arn())
            .sourceType("cluster")
            .sourceIds(defaultCluster.id())
            .severity("INFO")
            .eventCategories(            
                "configuration",
                "management",
                "monitoring",
                "security")
            .tags(Map.of("Name", "default"))
            .build());

    }
}
```
```yaml
resources:
  defaultCluster:
    type: aws:redshift:Cluster
    properties:
      clusterIdentifier: default
      databaseName: default
  defaultTopic:
    type: aws:sns:Topic
  defaultEventSubscription:
    type: aws:redshift:EventSubscription
    properties:
      snsTopicArn: ${defaultTopic.arn}
      sourceType: cluster
      sourceIds:
        - ${defaultCluster.id}
      severity: INFO
      eventCategories:
        - configuration
        - management
        - monitoring
        - security
      tags:
        Name: default
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Event Subscriptions using the `name`. For example:

```sh
 $ pulumi import aws:redshift/eventSubscription:EventSubscription default redshift-event-sub
```
 