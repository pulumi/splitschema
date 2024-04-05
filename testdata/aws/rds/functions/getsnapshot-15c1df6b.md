Use this data source to get information about a DB Snapshot for use when provisioning DB instances

> **NOTE:** This data source does not apply to snapshots created on Aurora DB clusters.
See the `aws.rds.ClusterSnapshot` data source for DB Cluster snapshots.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const prod = new aws.rds.Instance("prod", {
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
const latestProdSnapshot = aws.rds.getSnapshotOutput({
    dbInstanceIdentifier: prod.identifier,
    mostRecent: true,
});
// Use the latest production snapshot to create a dev instance.
const dev = new aws.rds.Instance("dev", {
    instanceClass: "db.t2.micro",
    dbName: "mydbdev",
    snapshotIdentifier: latestProdSnapshot.apply(latestProdSnapshot => latestProdSnapshot.id),
});
```
```python
import pulumi
import pulumi_aws as aws

prod = aws.rds.Instance("prod",
    allocated_storage=10,
    engine="mysql",
    engine_version="5.6.17",
    instance_class="db.t2.micro",
    db_name="mydb",
    username="foo",
    password="bar",
    db_subnet_group_name="my_database_subnet_group",
    parameter_group_name="default.mysql5.6")
latest_prod_snapshot = aws.rds.get_snapshot_output(db_instance_identifier=prod.identifier,
    most_recent=True)
# Use the latest production snapshot to create a dev instance.
dev = aws.rds.Instance("dev",
    instance_class="db.t2.micro",
    db_name="mydbdev",
    snapshot_identifier=latest_prod_snapshot.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var prod = new Aws.Rds.Instance("prod", new()
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

    var latestProdSnapshot = Aws.Rds.GetSnapshot.Invoke(new()
    {
        DbInstanceIdentifier = prod.Identifier,
        MostRecent = true,
    });

    // Use the latest production snapshot to create a dev instance.
    var dev = new Aws.Rds.Instance("dev", new()
    {
        InstanceClass = "db.t2.micro",
        DbName = "mydbdev",
        SnapshotIdentifier = latestProdSnapshot.Apply(getSnapshotResult => getSnapshotResult.Id),
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
		prod, err := rds.NewInstance(ctx, "prod", &rds.InstanceArgs{
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
		latestProdSnapshot := rds.LookupSnapshotOutput(ctx, rds.GetSnapshotOutputArgs{
			DbInstanceIdentifier: prod.Identifier,
			MostRecent:           pulumi.Bool(true),
		}, nil)
		_, err = rds.NewInstance(ctx, "dev", &rds.InstanceArgs{
			InstanceClass: pulumi.String("db.t2.micro"),
			DbName:        pulumi.String("mydbdev"),
			SnapshotIdentifier: latestProdSnapshot.ApplyT(func(latestProdSnapshot rds.GetSnapshotResult) (*string, error) {
				return &latestProdSnapshot.Id, nil
			}).(pulumi.StringPtrOutput),
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
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetSnapshotArgs;
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
        var prod = new Instance("prod", InstanceArgs.builder()        
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

        final var latestProdSnapshot = RdsFunctions.getSnapshot(GetSnapshotArgs.builder()
            .dbInstanceIdentifier(prod.identifier())
            .mostRecent(true)
            .build());

        var dev = new Instance("dev", InstanceArgs.builder()        
            .instanceClass("db.t2.micro")
            .dbName("mydbdev")
            .snapshotIdentifier(latestProdSnapshot.applyValue(getSnapshotResult -> getSnapshotResult).applyValue(latestProdSnapshot -> latestProdSnapshot.applyValue(getSnapshotResult -> getSnapshotResult.id())))
            .build());

    }
}
```
```yaml
resources:
  prod:
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
  # Use the latest production snapshot to create a dev instance.
  dev:
    type: aws:rds:Instance
    properties:
      instanceClass: db.t2.micro
      dbName: mydbdev
      snapshotIdentifier: ${latestProdSnapshot.id}
variables:
  latestProdSnapshot:
    fn::invoke:
      Function: aws:rds:getSnapshot
      Arguments:
        dbInstanceIdentifier: ${prod.identifier}
        mostRecent: true
```
{{% /example %}}
{{% /examples %}}