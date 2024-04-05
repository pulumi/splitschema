Manages an RDS database instance snapshot. For managing RDS database cluster snapshots, see the `aws.rds.ClusterSnapshot` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.rds.Instance("bar", {
    allocatedStorage: 10,
    engine: "mysql",
    engineVersion: "5.6.21",
    instanceClass: "db.t2.micro",
    dbName: "baz",
    password: "barbarbarbar",
    username: "foo",
    maintenanceWindow: "Fri:09:00-Fri:09:30",
    backupRetentionPeriod: 0,
    parameterGroupName: "default.mysql5.6",
});
const test = new aws.rds.Snapshot("test", {
    dbInstanceIdentifier: bar.identifier,
    dbSnapshotIdentifier: "testsnapshot1234",
});
```
```python
import pulumi
import pulumi_aws as aws

bar = aws.rds.Instance("bar",
    allocated_storage=10,
    engine="mysql",
    engine_version="5.6.21",
    instance_class="db.t2.micro",
    db_name="baz",
    password="barbarbarbar",
    username="foo",
    maintenance_window="Fri:09:00-Fri:09:30",
    backup_retention_period=0,
    parameter_group_name="default.mysql5.6")
test = aws.rds.Snapshot("test",
    db_instance_identifier=bar.identifier,
    db_snapshot_identifier="testsnapshot1234")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bar = new Aws.Rds.Instance("bar", new()
    {
        AllocatedStorage = 10,
        Engine = "mysql",
        EngineVersion = "5.6.21",
        InstanceClass = "db.t2.micro",
        DbName = "baz",
        Password = "barbarbarbar",
        Username = "foo",
        MaintenanceWindow = "Fri:09:00-Fri:09:30",
        BackupRetentionPeriod = 0,
        ParameterGroupName = "default.mysql5.6",
    });

    var test = new Aws.Rds.Snapshot("test", new()
    {
        DbInstanceIdentifier = bar.Identifier,
        DbSnapshotIdentifier = "testsnapshot1234",
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
		bar, err := rds.NewInstance(ctx, "bar", &rds.InstanceArgs{
			AllocatedStorage:      pulumi.Int(10),
			Engine:                pulumi.String("mysql"),
			EngineVersion:         pulumi.String("5.6.21"),
			InstanceClass:         pulumi.String("db.t2.micro"),
			DbName:                pulumi.String("baz"),
			Password:              pulumi.String("barbarbarbar"),
			Username:              pulumi.String("foo"),
			MaintenanceWindow:     pulumi.String("Fri:09:00-Fri:09:30"),
			BackupRetentionPeriod: pulumi.Int(0),
			ParameterGroupName:    pulumi.String("default.mysql5.6"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewSnapshot(ctx, "test", &rds.SnapshotArgs{
			DbInstanceIdentifier: bar.Identifier,
			DbSnapshotIdentifier: pulumi.String("testsnapshot1234"),
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
import com.pulumi.aws.rds.Snapshot;
import com.pulumi.aws.rds.SnapshotArgs;
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
        var bar = new Instance("bar", InstanceArgs.builder()        
            .allocatedStorage(10)
            .engine("mysql")
            .engineVersion("5.6.21")
            .instanceClass("db.t2.micro")
            .dbName("baz")
            .password("barbarbarbar")
            .username("foo")
            .maintenanceWindow("Fri:09:00-Fri:09:30")
            .backupRetentionPeriod(0)
            .parameterGroupName("default.mysql5.6")
            .build());

        var test = new Snapshot("test", SnapshotArgs.builder()        
            .dbInstanceIdentifier(bar.identifier())
            .dbSnapshotIdentifier("testsnapshot1234")
            .build());

    }
}
```
```yaml
resources:
  bar:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 10
      engine: mysql
      engineVersion: 5.6.21
      instanceClass: db.t2.micro
      dbName: baz
      password: barbarbarbar
      username: foo
      maintenanceWindow: Fri:09:00-Fri:09:30
      backupRetentionPeriod: 0
      parameterGroupName: default.mysql5.6
  test:
    type: aws:rds:Snapshot
    properties:
      dbInstanceIdentifier: ${bar.identifier}
      dbSnapshotIdentifier: testsnapshot1234
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_db_snapshot` using the snapshot identifier. For example:

```sh
 $ pulumi import aws:rds/snapshot:Snapshot example my-snapshot
```
 