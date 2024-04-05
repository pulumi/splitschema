Use this data source to get information about a DB Cluster Snapshot for use when provisioning DB clusters.

> **NOTE:** This data source does not apply to snapshots created on DB Instances.
See the `aws.rds.Snapshot` data source for DB Instance snapshots.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const developmentFinalSnapshot = aws.rds.getClusterSnapshot({
    dbClusterIdentifier: "development_cluster",
    mostRecent: true,
});
// Use the last snapshot of the dev database before it was destroyed to create
// a new dev database.
const auroraCluster = new aws.rds.Cluster("auroraCluster", {
    clusterIdentifier: "development_cluster",
    snapshotIdentifier: developmentFinalSnapshot.then(developmentFinalSnapshot => developmentFinalSnapshot.id),
    dbSubnetGroupName: "my_db_subnet_group",
});
const auroraClusterInstance = new aws.rds.ClusterInstance("auroraClusterInstance", {
    clusterIdentifier: auroraCluster.id,
    instanceClass: "db.t2.small",
    dbSubnetGroupName: "my_db_subnet_group",
});
```
```python
import pulumi
import pulumi_aws as aws

development_final_snapshot = aws.rds.get_cluster_snapshot(db_cluster_identifier="development_cluster",
    most_recent=True)
# Use the last snapshot of the dev database before it was destroyed to create
# a new dev database.
aurora_cluster = aws.rds.Cluster("auroraCluster",
    cluster_identifier="development_cluster",
    snapshot_identifier=development_final_snapshot.id,
    db_subnet_group_name="my_db_subnet_group")
aurora_cluster_instance = aws.rds.ClusterInstance("auroraClusterInstance",
    cluster_identifier=aurora_cluster.id,
    instance_class="db.t2.small",
    db_subnet_group_name="my_db_subnet_group")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var developmentFinalSnapshot = Aws.Rds.GetClusterSnapshot.Invoke(new()
    {
        DbClusterIdentifier = "development_cluster",
        MostRecent = true,
    });

    // Use the last snapshot of the dev database before it was destroyed to create
    // a new dev database.
    var auroraCluster = new Aws.Rds.Cluster("auroraCluster", new()
    {
        ClusterIdentifier = "development_cluster",
        SnapshotIdentifier = developmentFinalSnapshot.Apply(getClusterSnapshotResult => getClusterSnapshotResult.Id),
        DbSubnetGroupName = "my_db_subnet_group",
    });

    var auroraClusterInstance = new Aws.Rds.ClusterInstance("auroraClusterInstance", new()
    {
        ClusterIdentifier = auroraCluster.Id,
        InstanceClass = "db.t2.small",
        DbSubnetGroupName = "my_db_subnet_group",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		developmentFinalSnapshot, err := rds.LookupClusterSnapshot(ctx, &rds.LookupClusterSnapshotArgs{
			DbClusterIdentifier: pulumi.StringRef("development_cluster"),
			MostRecent:          pulumi.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		auroraCluster, err := rds.NewCluster(ctx, "auroraCluster", &rds.ClusterArgs{
			ClusterIdentifier:  pulumi.String("development_cluster"),
			SnapshotIdentifier: *pulumi.String(developmentFinalSnapshot.Id),
			DbSubnetGroupName:  pulumi.String("my_db_subnet_group"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewClusterInstance(ctx, "auroraClusterInstance", &rds.ClusterInstanceArgs{
			ClusterIdentifier: auroraCluster.ID(),
			InstanceClass:     pulumi.String("db.t2.small"),
			DbSubnetGroupName: pulumi.String("my_db_subnet_group"),
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
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetClusterSnapshotArgs;
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
import com.pulumi.aws.rds.ClusterInstance;
import com.pulumi.aws.rds.ClusterInstanceArgs;
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
        final var developmentFinalSnapshot = RdsFunctions.getClusterSnapshot(GetClusterSnapshotArgs.builder()
            .dbClusterIdentifier("development_cluster")
            .mostRecent(true)
            .build());

        var auroraCluster = new Cluster("auroraCluster", ClusterArgs.builder()        
            .clusterIdentifier("development_cluster")
            .snapshotIdentifier(developmentFinalSnapshot.applyValue(getClusterSnapshotResult -> getClusterSnapshotResult.id()))
            .dbSubnetGroupName("my_db_subnet_group")
            .build());

        var auroraClusterInstance = new ClusterInstance("auroraClusterInstance", ClusterInstanceArgs.builder()        
            .clusterIdentifier(auroraCluster.id())
            .instanceClass("db.t2.small")
            .dbSubnetGroupName("my_db_subnet_group")
            .build());

    }
}
```
```yaml
resources:
  # Use the last snapshot of the dev database before it was destroyed to create
  # a new dev database.
  auroraCluster:
    type: aws:rds:Cluster
    properties:
      clusterIdentifier: development_cluster
      snapshotIdentifier: ${developmentFinalSnapshot.id}
      dbSubnetGroupName: my_db_subnet_group
  auroraClusterInstance:
    type: aws:rds:ClusterInstance
    properties:
      clusterIdentifier: ${auroraCluster.id}
      instanceClass: db.t2.small
      dbSubnetGroupName: my_db_subnet_group
variables:
  developmentFinalSnapshot:
    fn::invoke:
      Function: aws:rds:getClusterSnapshot
      Arguments:
        dbClusterIdentifier: development_cluster
        mostRecent: true
```
{{% /example %}}
{{% /examples %}}