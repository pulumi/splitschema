Provides a DB event subscription resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultInstance = new aws.rds.Instance("defaultInstance", {
    allocatedStorage: 10,
    engine: "mysql",
    engineVersion: "5.6.17",
    instanceClass: "db.t2.micro",
    dbName: "mydb",
    username: "foo",
    password: "bar",
    dbSubnetGroupName: "my_database_subnet_group",
    parameterGroupName: "default.mysql5.6",
});
const defaultTopic = new aws.sns.Topic("defaultTopic", {});
const defaultEventSubscription = new aws.rds.EventSubscription("defaultEventSubscription", {
    snsTopic: defaultTopic.arn,
    sourceType: "db-instance",
    sourceIds: [defaultInstance.identifier],
    eventCategories: [
        "availability",
        "deletion",
        "failover",
        "failure",
        "low storage",
        "maintenance",
        "notification",
        "read replica",
        "recovery",
        "restoration",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

default_instance = aws.rds.Instance("defaultInstance",
    allocated_storage=10,
    engine="mysql",
    engine_version="5.6.17",
    instance_class="db.t2.micro",
    db_name="mydb",
    username="foo",
    password="bar",
    db_subnet_group_name="my_database_subnet_group",
    parameter_group_name="default.mysql5.6")
default_topic = aws.sns.Topic("defaultTopic")
default_event_subscription = aws.rds.EventSubscription("defaultEventSubscription",
    sns_topic=default_topic.arn,
    source_type="db-instance",
    source_ids=[default_instance.identifier],
    event_categories=[
        "availability",
        "deletion",
        "failover",
        "failure",
        "low storage",
        "maintenance",
        "notification",
        "read replica",
        "recovery",
        "restoration",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultInstance = new Aws.Rds.Instance("defaultInstance", new()
    {
        AllocatedStorage = 10,
        Engine = "mysql",
        EngineVersion = "5.6.17",
        InstanceClass = "db.t2.micro",
        DbName = "mydb",
        Username = "foo",
        Password = "bar",
        DbSubnetGroupName = "my_database_subnet_group",
        ParameterGroupName = "default.mysql5.6",
    });

    var defaultTopic = new Aws.Sns.Topic("defaultTopic");

    var defaultEventSubscription = new Aws.Rds.EventSubscription("defaultEventSubscription", new()
    {
        SnsTopic = defaultTopic.Arn,
        SourceType = "db-instance",
        SourceIds = new[]
        {
            defaultInstance.Identifier,
        },
        EventCategories = new[]
        {
            "availability",
            "deletion",
            "failover",
            "failure",
            "low storage",
            "maintenance",
            "notification",
            "read replica",
            "recovery",
            "restoration",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultInstance, err := rds.NewInstance(ctx, "defaultInstance", &rds.InstanceArgs{
			AllocatedStorage:   pulumi.Int(10),
			Engine:             pulumi.String("mysql"),
			EngineVersion:      pulumi.String("5.6.17"),
			InstanceClass:      pulumi.String("db.t2.micro"),
			DbName:             pulumi.String("mydb"),
			Username:           pulumi.String("foo"),
			Password:           pulumi.String("bar"),
			DbSubnetGroupName:  pulumi.String("my_database_subnet_group"),
			ParameterGroupName: pulumi.String("default.mysql5.6"),
		})
		if err != nil {
			return err
		}
		defaultTopic, err := sns.NewTopic(ctx, "defaultTopic", nil)
		if err != nil {
			return err
		}
		_, err = rds.NewEventSubscription(ctx, "defaultEventSubscription", &rds.EventSubscriptionArgs{
			SnsTopic:   defaultTopic.Arn,
			SourceType: pulumi.String("db-instance"),
			SourceIds: pulumi.StringArray{
				defaultInstance.Identifier,
			},
			EventCategories: pulumi.StringArray{
				pulumi.String("availability"),
				pulumi.String("deletion"),
				pulumi.String("failover"),
				pulumi.String("failure"),
				pulumi.String("low storage"),
				pulumi.String("maintenance"),
				pulumi.String("notification"),
				pulumi.String("read replica"),
				pulumi.String("recovery"),
				pulumi.String("restoration"),
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
import com.pulumi.aws.rds.Instance;
import com.pulumi.aws.rds.InstanceArgs;
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.rds.EventSubscription;
import com.pulumi.aws.rds.EventSubscriptionArgs;
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
        var defaultInstance = new Instance("defaultInstance", InstanceArgs.builder()        
            .allocatedStorage(10)
            .engine("mysql")
            .engineVersion("5.6.17")
            .instanceClass("db.t2.micro")
            .dbName("mydb")
            .username("foo")
            .password("bar")
            .dbSubnetGroupName("my_database_subnet_group")
            .parameterGroupName("default.mysql5.6")
            .build());

        var defaultTopic = new Topic("defaultTopic");

        var defaultEventSubscription = new EventSubscription("defaultEventSubscription", EventSubscriptionArgs.builder()        
            .snsTopic(defaultTopic.arn())
            .sourceType("db-instance")
            .sourceIds(defaultInstance.identifier())
            .eventCategories(            
                "availability",
                "deletion",
                "failover",
                "failure",
                "low storage",
                "maintenance",
                "notification",
                "read replica",
                "recovery",
                "restoration")
            .build());

    }
}
```
```yaml
resources:
  defaultInstance:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 10
      engine: mysql
      engineVersion: 5.6.17
      instanceClass: db.t2.micro
      dbName: mydb
      username: foo
      password: bar
      dbSubnetGroupName: my_database_subnet_group
      parameterGroupName: default.mysql5.6
  defaultTopic:
    type: aws:sns:Topic
  defaultEventSubscription:
    type: aws:rds:EventSubscription
    properties:
      snsTopic: ${defaultTopic.arn}
      sourceType: db-instance
      sourceIds:
        - ${defaultInstance.identifier}
      eventCategories:
        - availability
        - deletion
        - failover
        - failure
        - low storage
        - maintenance
        - notification
        - read replica
        - recovery
        - restoration
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DB Event Subscriptions using the `name`. For example:

```sh
 $ pulumi import aws:rds/eventSubscription:EventSubscription default rds-event-sub
```
 