{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCluster = new aws.redshift.Cluster("defaultCluster", {
    clusterIdentifier: "tf-redshift-cluster",
    databaseName: "mydb",
    masterUsername: "foo",
    masterPassword: "Mustbe8characters",
    nodeType: "dc1.large",
    clusterType: "single-node",
});
const defaultSnapshotSchedule = new aws.redshift.SnapshotSchedule("defaultSnapshotSchedule", {
    identifier: "tf-redshift-snapshot-schedule",
    definitions: ["rate(12 hours)"],
});
const defaultSnapshotScheduleAssociation = new aws.redshift.SnapshotScheduleAssociation("defaultSnapshotScheduleAssociation", {
    clusterIdentifier: defaultCluster.id,
    scheduleIdentifier: defaultSnapshotSchedule.id,
});
```
```python
import pulumi
import pulumi_aws as aws

default_cluster = aws.redshift.Cluster("defaultCluster",
    cluster_identifier="tf-redshift-cluster",
    database_name="mydb",
    master_username="foo",
    master_password="Mustbe8characters",
    node_type="dc1.large",
    cluster_type="single-node")
default_snapshot_schedule = aws.redshift.SnapshotSchedule("defaultSnapshotSchedule",
    identifier="tf-redshift-snapshot-schedule",
    definitions=["rate(12 hours)"])
default_snapshot_schedule_association = aws.redshift.SnapshotScheduleAssociation("defaultSnapshotScheduleAssociation",
    cluster_identifier=default_cluster.id,
    schedule_identifier=default_snapshot_schedule.id)
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
        ClusterIdentifier = "tf-redshift-cluster",
        DatabaseName = "mydb",
        MasterUsername = "foo",
        MasterPassword = "Mustbe8characters",
        NodeType = "dc1.large",
        ClusterType = "single-node",
    });

    var defaultSnapshotSchedule = new Aws.RedShift.SnapshotSchedule("defaultSnapshotSchedule", new()
    {
        Identifier = "tf-redshift-snapshot-schedule",
        Definitions = new[]
        {
            "rate(12 hours)",
        },
    });

    var defaultSnapshotScheduleAssociation = new Aws.RedShift.SnapshotScheduleAssociation("defaultSnapshotScheduleAssociation", new()
    {
        ClusterIdentifier = defaultCluster.Id,
        ScheduleIdentifier = defaultSnapshotSchedule.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultCluster, err := redshift.NewCluster(ctx, "defaultCluster", &redshift.ClusterArgs{
			ClusterIdentifier: pulumi.String("tf-redshift-cluster"),
			DatabaseName:      pulumi.String("mydb"),
			MasterUsername:    pulumi.String("foo"),
			MasterPassword:    pulumi.String("Mustbe8characters"),
			NodeType:          pulumi.String("dc1.large"),
			ClusterType:       pulumi.String("single-node"),
		})
		if err != nil {
			return err
		}
		defaultSnapshotSchedule, err := redshift.NewSnapshotSchedule(ctx, "defaultSnapshotSchedule", &redshift.SnapshotScheduleArgs{
			Identifier: pulumi.String("tf-redshift-snapshot-schedule"),
			Definitions: pulumi.StringArray{
				pulumi.String("rate(12 hours)"),
			},
		})
		if err != nil {
			return err
		}
		_, err = redshift.NewSnapshotScheduleAssociation(ctx, "defaultSnapshotScheduleAssociation", &redshift.SnapshotScheduleAssociationArgs{
			ClusterIdentifier:  defaultCluster.ID(),
			ScheduleIdentifier: defaultSnapshotSchedule.ID(),
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
import com.pulumi.aws.redshift.SnapshotSchedule;
import com.pulumi.aws.redshift.SnapshotScheduleArgs;
import com.pulumi.aws.redshift.SnapshotScheduleAssociation;
import com.pulumi.aws.redshift.SnapshotScheduleAssociationArgs;
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
            .clusterIdentifier("tf-redshift-cluster")
            .databaseName("mydb")
            .masterUsername("foo")
            .masterPassword("Mustbe8characters")
            .nodeType("dc1.large")
            .clusterType("single-node")
            .build());

        var defaultSnapshotSchedule = new SnapshotSchedule("defaultSnapshotSchedule", SnapshotScheduleArgs.builder()        
            .identifier("tf-redshift-snapshot-schedule")
            .definitions("rate(12 hours)")
            .build());

        var defaultSnapshotScheduleAssociation = new SnapshotScheduleAssociation("defaultSnapshotScheduleAssociation", SnapshotScheduleAssociationArgs.builder()        
            .clusterIdentifier(defaultCluster.id())
            .scheduleIdentifier(defaultSnapshotSchedule.id())
            .build());

    }
}
```
```yaml
resources:
  defaultCluster:
    type: aws:redshift:Cluster
    properties:
      clusterIdentifier: tf-redshift-cluster
      databaseName: mydb
      masterUsername: foo
      masterPassword: Mustbe8characters
      nodeType: dc1.large
      clusterType: single-node
  defaultSnapshotSchedule:
    type: aws:redshift:SnapshotSchedule
    properties:
      identifier: tf-redshift-snapshot-schedule
      definitions:
        - rate(12 hours)
  defaultSnapshotScheduleAssociation:
    type: aws:redshift:SnapshotScheduleAssociation
    properties:
      clusterIdentifier: ${defaultCluster.id}
      scheduleIdentifier: ${defaultSnapshotSchedule.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Snapshot Schedule Association using the `<cluster-identifier>/<schedule-identifier>`. For example:

```sh
 $ pulumi import aws:redshift/snapshotScheduleAssociation:SnapshotScheduleAssociation default tf-redshift-cluster/tf-redshift-snapshot-schedule
```
 