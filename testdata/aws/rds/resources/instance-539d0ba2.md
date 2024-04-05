Provides an RDS instance resource.  A DB instance is an isolated database
environment in the cloud.  A DB instance can contain multiple user-created
databases.

Changes to a DB instance can occur when you manually change a parameter, such as
`allocated_storage`, and are reflected in the next maintenance window. Because
of this, this provider may report a difference in its planning phase because a
modification has not yet taken place. You can use the `apply_immediately` flag
to instruct the service to apply the change immediately (see documentation
below).

When upgrading the major version of an engine, `allow_major_version_upgrade` must be set to `true`.

> **Note:** using `apply_immediately` can result in a brief downtime as the server reboots.
See the AWS Docs on [RDS Instance Maintenance][instance-maintenance] for more information.

> **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
Read more about sensitive data instate.



## RDS Instance Class Types

Amazon RDS supports instance classes for the following use cases: General-purpose, Memory-optimized, Burstable Performance, and Optimized-reads.
For more information please read the AWS RDS documentation about [DB Instance Class Types](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)

## Low-Downtime Updates

By default, RDS applies updates to DB Instances in-place, which can lead to service interruptions.
Low-downtime updates minimize service interruptions by performing the updates with an [RDS Blue/Green deployment][blue-green] and switching over the instances when complete.

Low-downtime updates are only available for DB Instances using MySQL and MariaDB,
as other engines are not supported by RDS Blue/Green deployments.

Backups must be enabled to use low-downtime updates.

Enable low-downtime updates by setting `blue_green_update.enabled` to `true`.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.rds.Instance("default", {
    allocatedStorage: 10,
    dbName: "mydb",
    engine: "mysql",
    engineVersion: "5.7",
    instanceClass: "db.t3.micro",
    parameterGroupName: "default.mysql5.7",
    password: "foobarbaz",
    skipFinalSnapshot: true,
    username: "foo",
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.Instance("default",
    allocated_storage=10,
    db_name="mydb",
    engine="mysql",
    engine_version="5.7",
    instance_class="db.t3.micro",
    parameter_group_name="default.mysql5.7",
    password="foobarbaz",
    skip_final_snapshot=True,
    username="foo")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Rds.Instance("default", new()
    {
        AllocatedStorage = 10,
        DbName = "mydb",
        Engine = "mysql",
        EngineVersion = "5.7",
        InstanceClass = "db.t3.micro",
        ParameterGroupName = "default.mysql5.7",
        Password = "foobarbaz",
        SkipFinalSnapshot = true,
        Username = "foo",
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
		_, err := rds.NewInstance(ctx, "default", &rds.InstanceArgs{
			AllocatedStorage:   pulumi.Int(10),
			DbName:             pulumi.String("mydb"),
			Engine:             pulumi.String("mysql"),
			EngineVersion:      pulumi.String("5.7"),
			InstanceClass:      pulumi.String("db.t3.micro"),
			ParameterGroupName: pulumi.String("default.mysql5.7"),
			Password:           pulumi.String("foobarbaz"),
			SkipFinalSnapshot:  pulumi.Bool(true),
			Username:           pulumi.String("foo"),
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
        var default_ = new Instance("default", InstanceArgs.builder()        
            .allocatedStorage(10)
            .dbName("mydb")
            .engine("mysql")
            .engineVersion("5.7")
            .instanceClass("db.t3.micro")
            .parameterGroupName("default.mysql5.7")
            .password("foobarbaz")
            .skipFinalSnapshot(true)
            .username("foo")
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 10
      dbName: mydb
      engine: mysql
      engineVersion: '5.7'
      instanceClass: db.t3.micro
      parameterGroupName: default.mysql5.7
      password: foobarbaz
      skipFinalSnapshot: true
      username: foo
```
{{% /example %}}
{{% example %}}
### RDS Custom for Oracle Usage with Replica

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetOrderableDbInstanceArgs;
import com.pulumi.aws.kms.KmsFunctions;
import com.pulumi.aws.kms.inputs.GetKeyArgs;
import com.pulumi.aws.rds.Instance;
import com.pulumi.aws.rds.InstanceArgs;
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
        final var custom-oracle = RdsFunctions.getOrderableDbInstance(GetOrderableDbInstanceArgs.builder()
            .engine("custom-oracle-ee")
            .engineVersion("19.c.ee.002")
            .licenseModel("bring-your-own-license")
            .storageType("gp3")
            .preferredInstanceClasses(            
                "db.r5.xlarge",
                "db.r5.2xlarge",
                "db.r5.4xlarge")
            .build());

        final var byId = KmsFunctions.getKey(GetKeyArgs.builder()
            .keyId("example-ef278353ceba4a5a97de6784565b9f78")
            .build());

        var default_ = new Instance("default", InstanceArgs.builder()        
            .allocatedStorage(50)
            .autoMinorVersionUpgrade(false)
            .customIamInstanceProfile("AWSRDSCustomInstanceProfile")
            .backupRetentionPeriod(7)
            .dbSubnetGroupName(local.db_subnet_group_name())
            .engine(custom_oracle.engine())
            .engineVersion(custom_oracle.engineVersion())
            .identifier("ee-instance-demo")
            .instanceClass(custom_oracle.instanceClass())
            .kmsKeyId(byId.applyValue(getKeyResult -> getKeyResult.arn()))
            .licenseModel(custom_oracle.licenseModel())
            .multiAz(false)
            .password("avoid-plaintext-passwords")
            .username("test")
            .storageEncrypted(true)
            .timeouts(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

        var test_replica = new Instance("test-replica", InstanceArgs.builder()        
            .replicateSourceDb(default_.identifier())
            .replicaMode("mounted")
            .autoMinorVersionUpgrade(false)
            .customIamInstanceProfile("AWSRDSCustomInstanceProfile")
            .backupRetentionPeriod(7)
            .identifier("ee-instance-replica")
            .instanceClass(custom_oracle.instanceClass())
            .kmsKeyId(byId.applyValue(getKeyResult -> getKeyResult.arn()))
            .multiAz(false)
            .skipFinalSnapshot(true)
            .storageEncrypted(true)
            .timeouts(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 50
      autoMinorVersionUpgrade: false
      # Custom for Oracle does not support minor version upgrades
      customIamInstanceProfile: AWSRDSCustomInstanceProfile
      # Instance profile is required for Custom for Oracle. See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc
      backupRetentionPeriod: 7
      dbSubnetGroupName: ${local.db_subnet_group_name}
      engine: ${["custom-oracle"].engine}
      engineVersion: ${["custom-oracle"].engineVersion}
      identifier: ee-instance-demo
      instanceClass: ${["custom-oracle"].instanceClass}
      kmsKeyId: ${byId.arn}
      licenseModel: ${["custom-oracle"].licenseModel}
      multiAz: false
      # Custom for Oracle does not support multi-az
      password: avoid-plaintext-passwords
      username: test
      storageEncrypted: true
      timeouts:
        - create: 3h
          delete: 3h
          update: 3h
  test-replica:
    type: aws:rds:Instance
    properties:
      replicateSourceDb: ${default.identifier}
      replicaMode: mounted
      autoMinorVersionUpgrade: false
      customIamInstanceProfile: AWSRDSCustomInstanceProfile
      # Instance profile is required for Custom for Oracle. See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc
      backupRetentionPeriod: 7
      identifier: ee-instance-replica
      instanceClass: ${["custom-oracle"].instanceClass}
      kmsKeyId: ${byId.arn}
      multiAz: false
      # Custom for Oracle does not support multi-az
      skipFinalSnapshot: true
      storageEncrypted: true
      timeouts:
        - create: 3h
          delete: 3h
          update: 3h
variables:
  custom-oracle:
    fn::invoke:
      Function: aws:rds:getOrderableDbInstance
      Arguments:
        engine: custom-oracle-ee
        engineVersion: 19.c.ee.002
        licenseModel: bring-your-own-license
        storageType: gp3
        preferredInstanceClasses:
          - db.r5.xlarge
          - db.r5.2xlarge
          - db.r5.4xlarge
  byId:
    fn::invoke:
      Function: aws:kms:getKey
      Arguments:
        keyId: example-ef278353ceba4a5a97de6784565b9f78
```
{{% /example %}}
{{% example %}}
### RDS Custom for SQL Server

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetOrderableDbInstanceArgs;
import com.pulumi.aws.kms.KmsFunctions;
import com.pulumi.aws.kms.inputs.GetKeyArgs;
import com.pulumi.aws.rds.Instance;
import com.pulumi.aws.rds.InstanceArgs;
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
        final var custom-sqlserver = RdsFunctions.getOrderableDbInstance(GetOrderableDbInstanceArgs.builder()
            .engine("custom-sqlserver-se")
            .engineVersion("15.00.4249.2.v1")
            .storageType("gp3")
            .preferredInstanceClasses(            
                "db.r5.xlarge",
                "db.r5.2xlarge",
                "db.r5.4xlarge")
            .build());

        final var byId = KmsFunctions.getKey(GetKeyArgs.builder()
            .keyId("example-ef278353ceba4a5a97de6784565b9f78")
            .build());

        var example = new Instance("example", InstanceArgs.builder()        
            .allocatedStorage(500)
            .autoMinorVersionUpgrade(false)
            .customIamInstanceProfile("AWSRDSCustomSQLServerInstanceRole")
            .backupRetentionPeriod(7)
            .dbSubnetGroupName(local.db_subnet_group_name())
            .engine(custom_sqlserver.engine())
            .engineVersion(custom_sqlserver.engineVersion())
            .identifier("sql-instance-demo")
            .instanceClass(custom_sqlserver.instanceClass())
            .kmsKeyId(byId.applyValue(getKeyResult -> getKeyResult.arn()))
            .multiAz(false)
            .password("avoid-plaintext-passwords")
            .storageEncrypted(true)
            .username("test")
            .timeouts(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 500
      autoMinorVersionUpgrade: false
      # Custom for SQL Server does not support minor version upgrades
      customIamInstanceProfile: AWSRDSCustomSQLServerInstanceRole
      # Instance profile is required for Custom for SQL Server. See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-sqlserver.html#custom-setup-sqlserver.iam
      backupRetentionPeriod: 7
      dbSubnetGroupName: ${local.db_subnet_group_name}
      # Copy the subnet group from the RDS Console
      engine: ${["custom-sqlserver"].engine}
      engineVersion: ${["custom-sqlserver"].engineVersion}
      identifier: sql-instance-demo
      instanceClass: ${["custom-sqlserver"].instanceClass}
      kmsKeyId: ${byId.arn}
      multiAz: false
      # Custom for SQL Server does support multi-az
      password: avoid-plaintext-passwords
      storageEncrypted: true
      username: test
      timeouts:
        - create: 3h
          delete: 3h
          update: 3h
variables:
  custom-sqlserver:
    fn::invoke:
      Function: aws:rds:getOrderableDbInstance
      Arguments:
        engine: custom-sqlserver-se
        engineVersion: 15.00.4249.2.v1
        storageType: gp3
        preferredInstanceClasses:
          - db.r5.xlarge
          - db.r5.2xlarge
          - db.r5.4xlarge
  byId:
    fn::invoke:
      Function: aws:kms:getKey
      Arguments:
        keyId: example-ef278353ceba4a5a97de6784565b9f78
```
{{% /example %}}
{{% example %}}
### RDS Db2 Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const default = aws.rds.getEngineVersion({
    engine: "db2-se",
});
const exampleOrderableDbInstance = Promise.all([_default, _default]).then(([_default, _default1]) => aws.rds.getOrderableDbInstance({
    engine: _default.engine,
    engineVersion: _default1.version,
    licenseModel: "bring-your-own-license",
    storageType: "gp3",
    preferredInstanceClasses: [
        "db.t3.small",
        "db.r6i.large",
        "db.m6i.large",
    ],
}));
// The RDS Db2 instance resource requires licensing information. Create a new parameter group using the default paramater group as a source, and set license information.
const exampleParameterGroup = new aws.rds.ParameterGroup("exampleParameterGroup", {
    family: _default.then(_default => _default.parameterGroupFamily),
    parameters: [
        {
            applyMethod: "immediate",
            name: "rds.ibm_customer_id",
            value: "0",
        },
        {
            applyMethod: "immediate",
            name: "rds.ibm_site_id",
            value: "0",
        },
    ],
});
// Create the RDS Db2 instance, use the data sources defined to set attributes
const exampleInstance = new aws.rds.Instance("exampleInstance", {
    allocatedStorage: 100,
    backupRetentionPeriod: 7,
    dbName: "test",
    engine: exampleOrderableDbInstance.then(exampleOrderableDbInstance => exampleOrderableDbInstance.engine),
    engineVersion: exampleOrderableDbInstance.then(exampleOrderableDbInstance => exampleOrderableDbInstance.engineVersion),
    identifier: "db2-instance-demo",
    instanceClass: exampleOrderableDbInstance.then(exampleOrderableDbInstance => exampleOrderableDbInstance.instanceClass).apply((x) => aws.rds.instancetype.InstanceType[x]),
    parameterGroupName: exampleParameterGroup.name,
    password: "avoid-plaintext-passwords",
    username: "test",
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.get_engine_version(engine="db2-se")
example_orderable_db_instance = aws.rds.get_orderable_db_instance(engine=default.engine,
    engine_version=default.version,
    license_model="bring-your-own-license",
    storage_type="gp3",
    preferred_instance_classes=[
        "db.t3.small",
        "db.r6i.large",
        "db.m6i.large",
    ])
# The RDS Db2 instance resource requires licensing information. Create a new parameter group using the default paramater group as a source, and set license information.
example_parameter_group = aws.rds.ParameterGroup("exampleParameterGroup",
    family=default.parameter_group_family,
    parameters=[
        aws.rds.ParameterGroupParameterArgs(
            apply_method="immediate",
            name="rds.ibm_customer_id",
            value="0",
        ),
        aws.rds.ParameterGroupParameterArgs(
            apply_method="immediate",
            name="rds.ibm_site_id",
            value="0",
        ),
    ])
# Create the RDS Db2 instance, use the data sources defined to set attributes
example_instance = aws.rds.Instance("exampleInstance",
    allocated_storage=100,
    backup_retention_period=7,
    db_name="test",
    engine=example_orderable_db_instance.engine,
    engine_version=example_orderable_db_instance.engine_version,
    identifier="db2-instance-demo",
    instance_class=example_orderable_db_instance.instance_class.apply(lambda x: aws.rds/instancetype.InstanceType(x)),
    parameter_group_name=example_parameter_group.name,
    password="avoid-plaintext-passwords",
    username="test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = Aws.Rds.GetEngineVersion.Invoke(new()
    {
        Engine = "db2-se",
    });

    var exampleOrderableDbInstance = Aws.Rds.GetOrderableDbInstance.Invoke(new()
    {
        Engine = @default.Apply(getEngineVersionResult => getEngineVersionResult.Engine),
        EngineVersion = @default.Apply(getEngineVersionResult => getEngineVersionResult.Version),
        LicenseModel = "bring-your-own-license",
        StorageType = "gp3",
        PreferredInstanceClasses = new[]
        {
            "db.t3.small",
            "db.r6i.large",
            "db.m6i.large",
        },
    });

    // The RDS Db2 instance resource requires licensing information. Create a new parameter group using the default paramater group as a source, and set license information.
    var exampleParameterGroup = new Aws.Rds.ParameterGroup("exampleParameterGroup", new()
    {
        Family = @default.Apply(@default => @default.Apply(getEngineVersionResult => getEngineVersionResult.ParameterGroupFamily)),
        Parameters = new[]
        {
            new Aws.Rds.Inputs.ParameterGroupParameterArgs
            {
                ApplyMethod = "immediate",
                Name = "rds.ibm_customer_id",
                Value = "0",
            },
            new Aws.Rds.Inputs.ParameterGroupParameterArgs
            {
                ApplyMethod = "immediate",
                Name = "rds.ibm_site_id",
                Value = "0",
            },
        },
    });

    // Create the RDS Db2 instance, use the data sources defined to set attributes
    var exampleInstance = new Aws.Rds.Instance("exampleInstance", new()
    {
        AllocatedStorage = 100,
        BackupRetentionPeriod = 7,
        DbName = "test",
        Engine = exampleOrderableDbInstance.Apply(getOrderableDbInstanceResult => getOrderableDbInstanceResult.Engine),
        EngineVersion = exampleOrderableDbInstance.Apply(getOrderableDbInstanceResult => getOrderableDbInstanceResult.EngineVersion),
        Identifier = "db2-instance-demo",
        InstanceClass = exampleOrderableDbInstance.Apply(getOrderableDbInstanceResult => getOrderableDbInstanceResult.InstanceClass).Apply(System.Enum.Parse<Aws.Rds.InstanceType.InstanceType>),
        ParameterGroupName = exampleParameterGroup.Name,
        Password = "avoid-plaintext-passwords",
        Username = "test",
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
		_default, err := rds.GetEngineVersion(ctx, &rds.GetEngineVersionArgs{
			Engine: "db2-se",
		}, nil)
		if err != nil {
			return err
		}
		exampleOrderableDbInstance, err := rds.GetOrderableDbInstance(ctx, &rds.GetOrderableDbInstanceArgs{
			Engine:        _default.Engine,
			EngineVersion: pulumi.StringRef(_default.Version),
			LicenseModel:  pulumi.StringRef("bring-your-own-license"),
			StorageType:   pulumi.StringRef("gp3"),
			PreferredInstanceClasses: []string{
				"db.t3.small",
				"db.r6i.large",
				"db.m6i.large",
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleParameterGroup, err := rds.NewParameterGroup(ctx, "exampleParameterGroup", &rds.ParameterGroupArgs{
			Family: *pulumi.String(_default.ParameterGroupFamily),
			Parameters: rds.ParameterGroupParameterArray{
				&rds.ParameterGroupParameterArgs{
					ApplyMethod: pulumi.String("immediate"),
					Name:        pulumi.String("rds.ibm_customer_id"),
					Value:       pulumi.String("0"),
				},
				&rds.ParameterGroupParameterArgs{
					ApplyMethod: pulumi.String("immediate"),
					Name:        pulumi.String("rds.ibm_site_id"),
					Value:       pulumi.String("0"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = rds.NewInstance(ctx, "exampleInstance", &rds.InstanceArgs{
			AllocatedStorage:      pulumi.Int(100),
			BackupRetentionPeriod: pulumi.Int(7),
			DbName:                pulumi.String("test"),
			Engine:                *pulumi.String(exampleOrderableDbInstance.Engine),
			EngineVersion:         *pulumi.String(exampleOrderableDbInstance.EngineVersion),
			Identifier:            pulumi.String("db2-instance-demo"),
			InstanceClass:         exampleOrderableDbInstance.InstanceClass.ApplyT(func(x *string) rds.InstanceType { return rds.InstanceType(*x) }).(rds.InstanceTypeOutput),
			ParameterGroupName:    exampleParameterGroup.Name,
			Password:              pulumi.String("avoid-plaintext-passwords"),
			Username:              pulumi.String("test"),
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
import com.pulumi.aws.rds.inputs.GetEngineVersionArgs;
import com.pulumi.aws.rds.inputs.GetOrderableDbInstanceArgs;
import com.pulumi.aws.rds.ParameterGroup;
import com.pulumi.aws.rds.ParameterGroupArgs;
import com.pulumi.aws.rds.inputs.ParameterGroupParameterArgs;
import com.pulumi.aws.rds.Instance;
import com.pulumi.aws.rds.InstanceArgs;
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
        final var default = RdsFunctions.getEngineVersion(GetEngineVersionArgs.builder()
            .engine("db2-se")
            .build());

        final var exampleOrderableDbInstance = RdsFunctions.getOrderableDbInstance(GetOrderableDbInstanceArgs.builder()
            .engine(default_.engine())
            .engineVersion(default_.version())
            .licenseModel("bring-your-own-license")
            .storageType("gp3")
            .preferredInstanceClasses(            
                "db.t3.small",
                "db.r6i.large",
                "db.m6i.large")
            .build());

        var exampleParameterGroup = new ParameterGroup("exampleParameterGroup", ParameterGroupArgs.builder()        
            .family(default_.parameterGroupFamily())
            .parameters(            
                ParameterGroupParameterArgs.builder()
                    .applyMethod("immediate")
                    .name("rds.ibm_customer_id")
                    .value(0)
                    .build(),
                ParameterGroupParameterArgs.builder()
                    .applyMethod("immediate")
                    .name("rds.ibm_site_id")
                    .value(0)
                    .build())
            .build());

        var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()        
            .allocatedStorage(100)
            .backupRetentionPeriod(7)
            .dbName("test")
            .engine(exampleOrderableDbInstance.applyValue(getOrderableDbInstanceResult -> getOrderableDbInstanceResult.engine()))
            .engineVersion(exampleOrderableDbInstance.applyValue(getOrderableDbInstanceResult -> getOrderableDbInstanceResult.engineVersion()))
            .identifier("db2-instance-demo")
            .instanceClass(exampleOrderableDbInstance.applyValue(getOrderableDbInstanceResult -> getOrderableDbInstanceResult.instanceClass()))
            .parameterGroupName(exampleParameterGroup.name())
            .password("avoid-plaintext-passwords")
            .username("test")
            .build());

    }
}
```
```yaml
resources:
  # The RDS Db2 instance resource requires licensing information. Create a new parameter group using the default paramater group as a source, and set license information.
  exampleParameterGroup:
    type: aws:rds:ParameterGroup
    properties:
      family: ${default.parameterGroupFamily}
      parameters:
        - applyMethod: immediate
          name: rds.ibm_customer_id
          value: 0
        - applyMethod: immediate
          name: rds.ibm_site_id
          value: 0
  # Create the RDS Db2 instance, use the data sources defined to set attributes
  exampleInstance:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 100
      backupRetentionPeriod: 7
      dbName: test
      engine: ${exampleOrderableDbInstance.engine}
      engineVersion: ${exampleOrderableDbInstance.engineVersion}
      identifier: db2-instance-demo
      instanceClass: ${exampleOrderableDbInstance.instanceClass}
      parameterGroupName: ${exampleParameterGroup.name}
      password: avoid-plaintext-passwords
      username: test
variables:
  default:
    fn::invoke:
      Function: aws:rds:getEngineVersion
      Arguments:
        engine: db2-se
  exampleOrderableDbInstance:
    fn::invoke:
      Function: aws:rds:getOrderableDbInstance
      Arguments:
        engine: ${default.engine}
        engineVersion: ${default.version}
        licenseModel: bring-your-own-license
        storageType: gp3
        preferredInstanceClasses:
          - db.t3.small
          - db.r6i.large
          - db.m6i.large
```
{{% /example %}}
{{% example %}}
### Storage Autoscaling

To enable Storage Autoscaling with instances that support the feature, define the `max_allocated_storage` argument higher than the `allocated_storage` argument. This provider will automatically hide differences with the `allocated_storage` argument value if autoscaling occurs.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.Instance("example", {
    allocatedStorage: 50,
    maxAllocatedStorage: 100,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.Instance("example",
    allocated_storage=50,
    max_allocated_storage=100)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rds.Instance("example", new()
    {
        AllocatedStorage = 50,
        MaxAllocatedStorage = 100,
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
		_, err := rds.NewInstance(ctx, "example", &rds.InstanceArgs{
			AllocatedStorage:    pulumi.Int(50),
			MaxAllocatedStorage: pulumi.Int(100),
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
        var example = new Instance("example", InstanceArgs.builder()        
            .allocatedStorage(50)
            .maxAllocatedStorage(100)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 50
      maxAllocatedStorage: 100
```
{{% /example %}}
{{% example %}}
### Managed Master Passwords via Secrets Manager, default KMS Key

> More information about RDS/Aurora Aurora integrates with Secrets Manager to manage master user passwords for your DB clusters can be found in the [RDS User Guide](https://aws.amazon.com/about-aws/whats-new/2022/12/amazon-rds-integration-aws-secrets-manager/) and [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html).

You can specify the `manage_master_user_password` attribute to enable managing the master password with Secrets Manager. You can also update an existing cluster to use Secrets Manager by specify the `manage_master_user_password` attribute and removing the `password` attribute (removal is required).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.rds.Instance("default", {
    allocatedStorage: 10,
    dbName: "mydb",
    engine: "mysql",
    engineVersion: "5.7",
    instanceClass: "db.t3.micro",
    manageMasterUserPassword: true,
    parameterGroupName: "default.mysql5.7",
    username: "foo",
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.Instance("default",
    allocated_storage=10,
    db_name="mydb",
    engine="mysql",
    engine_version="5.7",
    instance_class="db.t3.micro",
    manage_master_user_password=True,
    parameter_group_name="default.mysql5.7",
    username="foo")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Rds.Instance("default", new()
    {
        AllocatedStorage = 10,
        DbName = "mydb",
        Engine = "mysql",
        EngineVersion = "5.7",
        InstanceClass = "db.t3.micro",
        ManageMasterUserPassword = true,
        ParameterGroupName = "default.mysql5.7",
        Username = "foo",
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
		_, err := rds.NewInstance(ctx, "default", &rds.InstanceArgs{
			AllocatedStorage:         pulumi.Int(10),
			DbName:                   pulumi.String("mydb"),
			Engine:                   pulumi.String("mysql"),
			EngineVersion:            pulumi.String("5.7"),
			InstanceClass:            pulumi.String("db.t3.micro"),
			ManageMasterUserPassword: pulumi.Bool(true),
			ParameterGroupName:       pulumi.String("default.mysql5.7"),
			Username:                 pulumi.String("foo"),
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
        var default_ = new Instance("default", InstanceArgs.builder()        
            .allocatedStorage(10)
            .dbName("mydb")
            .engine("mysql")
            .engineVersion("5.7")
            .instanceClass("db.t3.micro")
            .manageMasterUserPassword(true)
            .parameterGroupName("default.mysql5.7")
            .username("foo")
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 10
      dbName: mydb
      engine: mysql
      engineVersion: '5.7'
      instanceClass: db.t3.micro
      manageMasterUserPassword: true
      parameterGroupName: default.mysql5.7
      username: foo
```
{{% /example %}}
{{% example %}}
### Managed Master Passwords via Secrets Manager, specific KMS Key

> More information about RDS/Aurora Aurora integrates with Secrets Manager to manage master user passwords for your DB clusters can be found in the [RDS User Guide](https://aws.amazon.com/about-aws/whats-new/2022/12/amazon-rds-integration-aws-secrets-manager/) and [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html).

You can specify the `master_user_secret_kms_key_id` attribute to specify a specific KMS Key.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kms.Key("example", {description: "Example KMS Key"});
const _default = new aws.rds.Instance("default", {
    allocatedStorage: 10,
    dbName: "mydb",
    engine: "mysql",
    engineVersion: "5.7",
    instanceClass: "db.t3.micro",
    manageMasterUserPassword: true,
    masterUserSecretKmsKeyId: example.keyId,
    username: "foo",
    parameterGroupName: "default.mysql5.7",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kms.Key("example", description="Example KMS Key")
default = aws.rds.Instance("default",
    allocated_storage=10,
    db_name="mydb",
    engine="mysql",
    engine_version="5.7",
    instance_class="db.t3.micro",
    manage_master_user_password=True,
    master_user_secret_kms_key_id=example.key_id,
    username="foo",
    parameter_group_name="default.mysql5.7")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kms.Key("example", new()
    {
        Description = "Example KMS Key",
    });

    var @default = new Aws.Rds.Instance("default", new()
    {
        AllocatedStorage = 10,
        DbName = "mydb",
        Engine = "mysql",
        EngineVersion = "5.7",
        InstanceClass = "db.t3.micro",
        ManageMasterUserPassword = true,
        MasterUserSecretKmsKeyId = example.KeyId,
        Username = "foo",
        ParameterGroupName = "default.mysql5.7",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := kms.NewKey(ctx, "example", &kms.KeyArgs{
			Description: pulumi.String("Example KMS Key"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewInstance(ctx, "default", &rds.InstanceArgs{
			AllocatedStorage:         pulumi.Int(10),
			DbName:                   pulumi.String("mydb"),
			Engine:                   pulumi.String("mysql"),
			EngineVersion:            pulumi.String("5.7"),
			InstanceClass:            pulumi.String("db.t3.micro"),
			ManageMasterUserPassword: pulumi.Bool(true),
			MasterUserSecretKmsKeyId: example.KeyId,
			Username:                 pulumi.String("foo"),
			ParameterGroupName:       pulumi.String("default.mysql5.7"),
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.rds.Instance;
import com.pulumi.aws.rds.InstanceArgs;
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
        var example = new Key("example", KeyArgs.builder()        
            .description("Example KMS Key")
            .build());

        var default_ = new Instance("default", InstanceArgs.builder()        
            .allocatedStorage(10)
            .dbName("mydb")
            .engine("mysql")
            .engineVersion("5.7")
            .instanceClass("db.t3.micro")
            .manageMasterUserPassword(true)
            .masterUserSecretKmsKeyId(example.keyId())
            .username("foo")
            .parameterGroupName("default.mysql5.7")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kms:Key
    properties:
      description: Example KMS Key
  default:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 10
      dbName: mydb
      engine: mysql
      engineVersion: '5.7'
      instanceClass: db.t3.micro
      manageMasterUserPassword: true
      masterUserSecretKmsKeyId: ${example.keyId}
      username: foo
      parameterGroupName: default.mysql5.7
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DB Instances using the `identifier`. For example:

```sh
 $ pulumi import aws:rds/instance:Instance default mydb-rds-instance
```
 