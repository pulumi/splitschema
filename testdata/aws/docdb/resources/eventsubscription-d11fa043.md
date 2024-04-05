Provides a DocumentDB event subscription resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCluster = new aws.docdb.Cluster("exampleCluster", {
    clusterIdentifier: "example",
    availabilityZones: [
        data.aws_availability_zones.available.names[0],
        data.aws_availability_zones.available.names[1],
        data.aws_availability_zones.available.names[2],
    ],
    masterUsername: "foo",
    masterPassword: "mustbeeightcharaters",
    skipFinalSnapshot: true,
});
const exampleTopic = new aws.sns.Topic("exampleTopic", {});
const exampleEventSubscription = new aws.docdb.EventSubscription("exampleEventSubscription", {
    enabled: true,
    eventCategories: [
        "creation",
        "failure",
    ],
    sourceType: "db-cluster",
    sourceIds: [exampleCluster.id],
    snsTopicArn: exampleTopic.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example_cluster = aws.docdb.Cluster("exampleCluster",
    cluster_identifier="example",
    availability_zones=[
        data["aws_availability_zones"]["available"]["names"],
        data["aws_availability_zones"]["available"]["names"],
        data["aws_availability_zones"]["available"]["names"],
    ],
    master_username="foo",
    master_password="mustbeeightcharaters",
    skip_final_snapshot=True)
example_topic = aws.sns.Topic("exampleTopic")
example_event_subscription = aws.docdb.EventSubscription("exampleEventSubscription",
    enabled=True,
    event_categories=[
        "creation",
        "failure",
    ],
    source_type="db-cluster",
    source_ids=[example_cluster.id],
    sns_topic_arn=example_topic.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleCluster = new Aws.DocDB.Cluster("exampleCluster", new()
    {
        ClusterIdentifier = "example",
        AvailabilityZones = new[]
        {
            data.Aws_availability_zones.Available.Names[0],
            data.Aws_availability_zones.Available.Names[1],
            data.Aws_availability_zones.Available.Names[2],
        },
        MasterUsername = "foo",
        MasterPassword = "mustbeeightcharaters",
        SkipFinalSnapshot = true,
    });

    var exampleTopic = new Aws.Sns.Topic("exampleTopic");

    var exampleEventSubscription = new Aws.DocDB.EventSubscription("exampleEventSubscription", new()
    {
        Enabled = true,
        EventCategories = new[]
        {
            "creation",
            "failure",
        },
        SourceType = "db-cluster",
        SourceIds = new[]
        {
            exampleCluster.Id,
        },
        SnsTopicArn = exampleTopic.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/docdb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleCluster, err := docdb.NewCluster(ctx, "exampleCluster", &docdb.ClusterArgs{
			ClusterIdentifier: pulumi.String("example"),
			AvailabilityZones: pulumi.StringArray{
				data.Aws_availability_zones.Available.Names[0],
				data.Aws_availability_zones.Available.Names[1],
				data.Aws_availability_zones.Available.Names[2],
			},
			MasterUsername:    pulumi.String("foo"),
			MasterPassword:    pulumi.String("mustbeeightcharaters"),
			SkipFinalSnapshot: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		exampleTopic, err := sns.NewTopic(ctx, "exampleTopic", nil)
		if err != nil {
			return err
		}
		_, err = docdb.NewEventSubscription(ctx, "exampleEventSubscription", &docdb.EventSubscriptionArgs{
			Enabled: pulumi.Bool(true),
			EventCategories: pulumi.StringArray{
				pulumi.String("creation"),
				pulumi.String("failure"),
			},
			SourceType: pulumi.String("db-cluster"),
			SourceIds: pulumi.StringArray{
				exampleCluster.ID(),
			},
			SnsTopicArn: exampleTopic.Arn,
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
import com.pulumi.aws.docdb.Cluster;
import com.pulumi.aws.docdb.ClusterArgs;
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.docdb.EventSubscription;
import com.pulumi.aws.docdb.EventSubscriptionArgs;
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
        var exampleCluster = new Cluster("exampleCluster", ClusterArgs.builder()        
            .clusterIdentifier("example")
            .availabilityZones(            
                data.aws_availability_zones().available().names()[0],
                data.aws_availability_zones().available().names()[1],
                data.aws_availability_zones().available().names()[2])
            .masterUsername("foo")
            .masterPassword("mustbeeightcharaters")
            .skipFinalSnapshot(true)
            .build());

        var exampleTopic = new Topic("exampleTopic");

        var exampleEventSubscription = new EventSubscription("exampleEventSubscription", EventSubscriptionArgs.builder()        
            .enabled(true)
            .eventCategories(            
                "creation",
                "failure")
            .sourceType("db-cluster")
            .sourceIds(exampleCluster.id())
            .snsTopicArn(exampleTopic.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleCluster:
    type: aws:docdb:Cluster
    properties:
      clusterIdentifier: example
      availabilityZones:
        - ${data.aws_availability_zones.available.names[0]}
        - ${data.aws_availability_zones.available.names[1]}
        - ${data.aws_availability_zones.available.names[2]}
      masterUsername: foo
      masterPassword: mustbeeightcharaters
      skipFinalSnapshot: true
  exampleTopic:
    type: aws:sns:Topic
  exampleEventSubscription:
    type: aws:docdb:EventSubscription
    properties:
      enabled: true
      eventCategories:
        - creation
        - failure
      sourceType: db-cluster
      sourceIds:
        - ${exampleCluster.id}
      snsTopicArn: ${exampleTopic.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DocumentDB Event Subscriptions using the `name`. For example:

```sh
 $ pulumi import aws:docdb/eventSubscription:EventSubscription example event-sub
```
 