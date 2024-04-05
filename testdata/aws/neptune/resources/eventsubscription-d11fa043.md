{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCluster = new aws.neptune.Cluster("defaultCluster", {
    clusterIdentifier: "neptune-cluster-demo",
    engine: "neptune",
    backupRetentionPeriod: 5,
    preferredBackupWindow: "07:00-09:00",
    skipFinalSnapshot: true,
    iamDatabaseAuthenticationEnabled: true,
    applyImmediately: true,
});
const example = new aws.neptune.ClusterInstance("example", {
    clusterIdentifier: defaultCluster.id,
    engine: "neptune",
    instanceClass: "db.r4.large",
    applyImmediately: true,
});
const defaultTopic = new aws.sns.Topic("defaultTopic", {});
const defaultEventSubscription = new aws.neptune.EventSubscription("defaultEventSubscription", {
    snsTopicArn: defaultTopic.arn,
    sourceType: "db-instance",
    sourceIds: [example.id],
    eventCategories: [
        "maintenance",
        "availability",
        "creation",
        "backup",
        "restoration",
        "recovery",
        "deletion",
        "failover",
        "failure",
        "notification",
        "configuration change",
        "read replica",
    ],
    tags: {
        env: "test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

default_cluster = aws.neptune.Cluster("defaultCluster",
    cluster_identifier="neptune-cluster-demo",
    engine="neptune",
    backup_retention_period=5,
    preferred_backup_window="07:00-09:00",
    skip_final_snapshot=True,
    iam_database_authentication_enabled=True,
    apply_immediately=True)
example = aws.neptune.ClusterInstance("example",
    cluster_identifier=default_cluster.id,
    engine="neptune",
    instance_class="db.r4.large",
    apply_immediately=True)
default_topic = aws.sns.Topic("defaultTopic")
default_event_subscription = aws.neptune.EventSubscription("defaultEventSubscription",
    sns_topic_arn=default_topic.arn,
    source_type="db-instance",
    source_ids=[example.id],
    event_categories=[
        "maintenance",
        "availability",
        "creation",
        "backup",
        "restoration",
        "recovery",
        "deletion",
        "failover",
        "failure",
        "notification",
        "configuration change",
        "read replica",
    ],
    tags={
        "env": "test",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultCluster = new Aws.Neptune.Cluster("defaultCluster", new()
    {
        ClusterIdentifier = "neptune-cluster-demo",
        Engine = "neptune",
        BackupRetentionPeriod = 5,
        PreferredBackupWindow = "07:00-09:00",
        SkipFinalSnapshot = true,
        IamDatabaseAuthenticationEnabled = true,
        ApplyImmediately = true,
    });

    var example = new Aws.Neptune.ClusterInstance("example", new()
    {
        ClusterIdentifier = defaultCluster.Id,
        Engine = "neptune",
        InstanceClass = "db.r4.large",
        ApplyImmediately = true,
    });

    var defaultTopic = new Aws.Sns.Topic("defaultTopic");

    var defaultEventSubscription = new Aws.Neptune.EventSubscription("defaultEventSubscription", new()
    {
        SnsTopicArn = defaultTopic.Arn,
        SourceType = "db-instance",
        SourceIds = new[]
        {
            example.Id,
        },
        EventCategories = new[]
        {
            "maintenance",
            "availability",
            "creation",
            "backup",
            "restoration",
            "recovery",
            "deletion",
            "failover",
            "failure",
            "notification",
            "configuration change",
            "read replica",
        },
        Tags = 
        {
            { "env", "test" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/neptune"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultCluster, err := neptune.NewCluster(ctx, "defaultCluster", &neptune.ClusterArgs{
			ClusterIdentifier:                pulumi.String("neptune-cluster-demo"),
			Engine:                           pulumi.String("neptune"),
			BackupRetentionPeriod:            pulumi.Int(5),
			PreferredBackupWindow:            pulumi.String("07:00-09:00"),
			SkipFinalSnapshot:                pulumi.Bool(true),
			IamDatabaseAuthenticationEnabled: pulumi.Bool(true),
			ApplyImmediately:                 pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		example, err := neptune.NewClusterInstance(ctx, "example", &neptune.ClusterInstanceArgs{
			ClusterIdentifier: defaultCluster.ID(),
			Engine:            pulumi.String("neptune"),
			InstanceClass:     pulumi.String("db.r4.large"),
			ApplyImmediately:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		defaultTopic, err := sns.NewTopic(ctx, "defaultTopic", nil)
		if err != nil {
			return err
		}
		_, err = neptune.NewEventSubscription(ctx, "defaultEventSubscription", &neptune.EventSubscriptionArgs{
			SnsTopicArn: defaultTopic.Arn,
			SourceType:  pulumi.String("db-instance"),
			SourceIds: pulumi.StringArray{
				example.ID(),
			},
			EventCategories: pulumi.StringArray{
				pulumi.String("maintenance"),
				pulumi.String("availability"),
				pulumi.String("creation"),
				pulumi.String("backup"),
				pulumi.String("restoration"),
				pulumi.String("recovery"),
				pulumi.String("deletion"),
				pulumi.String("failover"),
				pulumi.String("failure"),
				pulumi.String("notification"),
				pulumi.String("configuration change"),
				pulumi.String("read replica"),
			},
			Tags: pulumi.StringMap{
				"env": pulumi.String("test"),
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
import com.pulumi.aws.neptune.Cluster;
import com.pulumi.aws.neptune.ClusterArgs;
import com.pulumi.aws.neptune.ClusterInstance;
import com.pulumi.aws.neptune.ClusterInstanceArgs;
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.neptune.EventSubscription;
import com.pulumi.aws.neptune.EventSubscriptionArgs;
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
            .clusterIdentifier("neptune-cluster-demo")
            .engine("neptune")
            .backupRetentionPeriod(5)
            .preferredBackupWindow("07:00-09:00")
            .skipFinalSnapshot(true)
            .iamDatabaseAuthenticationEnabled("true")
            .applyImmediately("true")
            .build());

        var example = new ClusterInstance("example", ClusterInstanceArgs.builder()        
            .clusterIdentifier(defaultCluster.id())
            .engine("neptune")
            .instanceClass("db.r4.large")
            .applyImmediately("true")
            .build());

        var defaultTopic = new Topic("defaultTopic");

        var defaultEventSubscription = new EventSubscription("defaultEventSubscription", EventSubscriptionArgs.builder()        
            .snsTopicArn(defaultTopic.arn())
            .sourceType("db-instance")
            .sourceIds(example.id())
            .eventCategories(            
                "maintenance",
                "availability",
                "creation",
                "backup",
                "restoration",
                "recovery",
                "deletion",
                "failover",
                "failure",
                "notification",
                "configuration change",
                "read replica")
            .tags(Map.of("env", "test"))
            .build());

    }
}
```
```yaml
resources:
  defaultCluster:
    type: aws:neptune:Cluster
    properties:
      clusterIdentifier: neptune-cluster-demo
      engine: neptune
      backupRetentionPeriod: 5
      preferredBackupWindow: 07:00-09:00
      skipFinalSnapshot: true
      iamDatabaseAuthenticationEnabled: 'true'
      applyImmediately: 'true'
  example:
    type: aws:neptune:ClusterInstance
    properties:
      clusterIdentifier: ${defaultCluster.id}
      engine: neptune
      instanceClass: db.r4.large
      applyImmediately: 'true'
  defaultTopic:
    type: aws:sns:Topic
  defaultEventSubscription:
    type: aws:neptune:EventSubscription
    properties:
      snsTopicArn: ${defaultTopic.arn}
      sourceType: db-instance
      sourceIds:
        - ${example.id}
      eventCategories:
        - maintenance
        - availability
        - creation
        - backup
        - restoration
        - recovery
        - deletion
        - failover
        - failure
        - notification
        - configuration change
        - read replica
      tags:
        env: test
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_neptune_event_subscription` using the event subscription name. For example:

```sh
 $ pulumi import aws:neptune/eventSubscription:EventSubscription example my-event-subscription
```
 