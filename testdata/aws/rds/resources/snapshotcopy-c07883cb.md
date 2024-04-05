Manages an RDS database instance snapshot copy. For managing RDS database cluster snapshots, see the `aws.rds.ClusterSnapshot` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleInstance = new aws.rds.Instance("exampleInstance", {
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
const exampleSnapshot = new aws.rds.Snapshot("exampleSnapshot", {
    dbInstanceIdentifier: exampleInstance.identifier,
    dbSnapshotIdentifier: "testsnapshot1234",
});
const exampleSnapshotCopy = new aws.rds.SnapshotCopy("exampleSnapshotCopy", {
    sourceDbSnapshotIdentifier: exampleSnapshot.dbSnapshotArn,
    targetDbSnapshotIdentifier: "testsnapshot1234-copy",
});
```
```python
import pulumi
import pulumi_aws as aws

example_instance = aws.rds.Instance("exampleInstance",
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
example_snapshot = aws.rds.Snapshot("exampleSnapshot",
    db_instance_identifier=example_instance.identifier,
    db_snapshot_identifier="testsnapshot1234")
example_snapshot_copy = aws.rds.SnapshotCopy("exampleSnapshotCopy",
    source_db_snapshot_identifier=example_snapshot.db_snapshot_arn,
    target_db_snapshot_identifier="testsnapshot1234-copy")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleInstance = new Aws.Rds.Instance("exampleInstance", new()
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

    var exampleSnapshot = new Aws.Rds.Snapshot("exampleSnapshot", new()
    {
        DbInstanceIdentifier = exampleInstance.Identifier,
        DbSnapshotIdentifier = "testsnapshot1234",
    });

    var exampleSnapshotCopy = new Aws.Rds.SnapshotCopy("exampleSnapshotCopy", new()
    {
        SourceDbSnapshotIdentifier = exampleSnapshot.DbSnapshotArn,
        TargetDbSnapshotIdentifier = "testsnapshot1234-copy",
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
		exampleInstance, err := rds.NewInstance(ctx, "exampleInstance", &rds.InstanceArgs{
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
		exampleSnapshot, err := rds.NewSnapshot(ctx, "exampleSnapshot", &rds.SnapshotArgs{
			DbInstanceIdentifier: exampleInstance.Identifier,
			DbSnapshotIdentifier: pulumi.String("testsnapshot1234"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewSnapshotCopy(ctx, "exampleSnapshotCopy", &rds.SnapshotCopyArgs{
			SourceDbSnapshotIdentifier: exampleSnapshot.DbSnapshotArn,
			TargetDbSnapshotIdentifier: pulumi.String("testsnapshot1234-copy"),
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
import com.pulumi.aws.rds.SnapshotCopy;
import com.pulumi.aws.rds.SnapshotCopyArgs;
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
        var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()        
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

        var exampleSnapshot = new Snapshot("exampleSnapshot", SnapshotArgs.builder()        
            .dbInstanceIdentifier(exampleInstance.identifier())
            .dbSnapshotIdentifier("testsnapshot1234")
            .build());

        var exampleSnapshotCopy = new SnapshotCopy("exampleSnapshotCopy", SnapshotCopyArgs.builder()        
            .sourceDbSnapshotIdentifier(exampleSnapshot.dbSnapshotArn())
            .targetDbSnapshotIdentifier("testsnapshot1234-copy")
            .build());

    }
}
```
```yaml
resources:
  exampleInstance:
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
  exampleSnapshot:
    type: aws:rds:Snapshot
    properties:
      dbInstanceIdentifier: ${exampleInstance.identifier}
      dbSnapshotIdentifier: testsnapshot1234
  exampleSnapshotCopy:
    type: aws:rds:SnapshotCopy
    properties:
      sourceDbSnapshotIdentifier: ${exampleSnapshot.dbSnapshotArn}
      targetDbSnapshotIdentifier: testsnapshot1234-copy
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_db_snapshot_copy` using the snapshot identifier. For example:

```sh
 $ pulumi import aws:rds/snapshotCopy:SnapshotCopy example my-snapshot
```
 