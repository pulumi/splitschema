Manages a [RDS Aurora Cluster](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Aurora.html). To manage cluster instances that inherit configuration from the cluster (when not running the cluster in `serverless` engine mode), see the `aws.rds.ClusterInstance` resource. To manage non-Aurora databases (e.g., MySQL, PostgreSQL, SQL Server, etc.), see the `aws.rds.Instance` resource.

For information on the difference between the available Aurora MySQL engines
see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
in the Amazon RDS User Guide.

Changes to an RDS Cluster can occur when you manually change a
parameter, such as `port`, and are reflected in the next maintenance
window. Because of this, this provider may report a difference in its planning
phase because a modification has not yet taken place. You can use the
`apply_immediately` flag to instruct the service to apply the change immediately
(see documentation below).

> **Note:** using `apply_immediately` can result in a
brief downtime as the server reboots. See the AWS Docs on [RDS Maintenance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html)
for more information.

{{% examples %}}
## Example Usage
{{% example %}}
### Aurora MySQL 2.x (MySQL 5.7)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.rds.Cluster("default", {
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    backupRetentionPeriod: 5,
    clusterIdentifier: "aurora-cluster-demo",
    databaseName: "mydb",
    engine: "aurora-mysql",
    engineVersion: "5.7.mysql_aurora.2.03.2",
    masterPassword: "bar",
    masterUsername: "foo",
    preferredBackupWindow: "07:00-09:00",
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.Cluster("default",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    backup_retention_period=5,
    cluster_identifier="aurora-cluster-demo",
    database_name="mydb",
    engine="aurora-mysql",
    engine_version="5.7.mysql_aurora.2.03.2",
    master_password="bar",
    master_username="foo",
    preferred_backup_window="07:00-09:00")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Rds.Cluster("default", new()
    {
        AvailabilityZones = new[]
        {
            "us-west-2a",
            "us-west-2b",
            "us-west-2c",
        },
        BackupRetentionPeriod = 5,
        ClusterIdentifier = "aurora-cluster-demo",
        DatabaseName = "mydb",
        Engine = "aurora-mysql",
        EngineVersion = "5.7.mysql_aurora.2.03.2",
        MasterPassword = "bar",
        MasterUsername = "foo",
        PreferredBackupWindow = "07:00-09:00",
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
		_, err := rds.NewCluster(ctx, "default", &rds.ClusterArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
				pulumi.String("us-west-2c"),
			},
			BackupRetentionPeriod: pulumi.Int(5),
			ClusterIdentifier:     pulumi.String("aurora-cluster-demo"),
			DatabaseName:          pulumi.String("mydb"),
			Engine:                pulumi.String("aurora-mysql"),
			EngineVersion:         pulumi.String("5.7.mysql_aurora.2.03.2"),
			MasterPassword:        pulumi.String("bar"),
			MasterUsername:        pulumi.String("foo"),
			PreferredBackupWindow: pulumi.String("07:00-09:00"),
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
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
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
        var default_ = new Cluster("default", ClusterArgs.builder()        
            .availabilityZones(            
                "us-west-2a",
                "us-west-2b",
                "us-west-2c")
            .backupRetentionPeriod(5)
            .clusterIdentifier("aurora-cluster-demo")
            .databaseName("mydb")
            .engine("aurora-mysql")
            .engineVersion("5.7.mysql_aurora.2.03.2")
            .masterPassword("bar")
            .masterUsername("foo")
            .preferredBackupWindow("07:00-09:00")
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:rds:Cluster
    properties:
      availabilityZones:
        - us-west-2a
        - us-west-2b
        - us-west-2c
      backupRetentionPeriod: 5
      clusterIdentifier: aurora-cluster-demo
      databaseName: mydb
      engine: aurora-mysql
      engineVersion: 5.7.mysql_aurora.2.03.2
      masterPassword: bar
      masterUsername: foo
      preferredBackupWindow: 07:00-09:00
```
{{% /example %}}
{{% example %}}
### Aurora MySQL 1.x (MySQL 5.6)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.rds.Cluster("default", {
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    backupRetentionPeriod: 5,
    clusterIdentifier: "aurora-cluster-demo",
    databaseName: "mydb",
    masterPassword: "bar",
    masterUsername: "foo",
    preferredBackupWindow: "07:00-09:00",
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.Cluster("default",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    backup_retention_period=5,
    cluster_identifier="aurora-cluster-demo",
    database_name="mydb",
    master_password="bar",
    master_username="foo",
    preferred_backup_window="07:00-09:00")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Rds.Cluster("default", new()
    {
        AvailabilityZones = new[]
        {
            "us-west-2a",
            "us-west-2b",
            "us-west-2c",
        },
        BackupRetentionPeriod = 5,
        ClusterIdentifier = "aurora-cluster-demo",
        DatabaseName = "mydb",
        MasterPassword = "bar",
        MasterUsername = "foo",
        PreferredBackupWindow = "07:00-09:00",
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
		_, err := rds.NewCluster(ctx, "default", &rds.ClusterArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
				pulumi.String("us-west-2c"),
			},
			BackupRetentionPeriod: pulumi.Int(5),
			ClusterIdentifier:     pulumi.String("aurora-cluster-demo"),
			DatabaseName:          pulumi.String("mydb"),
			MasterPassword:        pulumi.String("bar"),
			MasterUsername:        pulumi.String("foo"),
			PreferredBackupWindow: pulumi.String("07:00-09:00"),
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
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
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
        var default_ = new Cluster("default", ClusterArgs.builder()        
            .availabilityZones(            
                "us-west-2a",
                "us-west-2b",
                "us-west-2c")
            .backupRetentionPeriod(5)
            .clusterIdentifier("aurora-cluster-demo")
            .databaseName("mydb")
            .masterPassword("bar")
            .masterUsername("foo")
            .preferredBackupWindow("07:00-09:00")
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:rds:Cluster
    properties:
      availabilityZones:
        - us-west-2a
        - us-west-2b
        - us-west-2c
      backupRetentionPeriod: 5
      clusterIdentifier: aurora-cluster-demo
      databaseName: mydb
      masterPassword: bar
      masterUsername: foo
      preferredBackupWindow: 07:00-09:00
```
{{% /example %}}
{{% example %}}
### Aurora with PostgreSQL engine

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const postgresql = new aws.rds.Cluster("postgresql", {
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    backupRetentionPeriod: 5,
    clusterIdentifier: "aurora-cluster-demo",
    databaseName: "mydb",
    engine: "aurora-postgresql",
    masterPassword: "bar",
    masterUsername: "foo",
    preferredBackupWindow: "07:00-09:00",
});
```
```python
import pulumi
import pulumi_aws as aws

postgresql = aws.rds.Cluster("postgresql",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    backup_retention_period=5,
    cluster_identifier="aurora-cluster-demo",
    database_name="mydb",
    engine="aurora-postgresql",
    master_password="bar",
    master_username="foo",
    preferred_backup_window="07:00-09:00")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var postgresql = new Aws.Rds.Cluster("postgresql", new()
    {
        AvailabilityZones = new[]
        {
            "us-west-2a",
            "us-west-2b",
            "us-west-2c",
        },
        BackupRetentionPeriod = 5,
        ClusterIdentifier = "aurora-cluster-demo",
        DatabaseName = "mydb",
        Engine = "aurora-postgresql",
        MasterPassword = "bar",
        MasterUsername = "foo",
        PreferredBackupWindow = "07:00-09:00",
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
		_, err := rds.NewCluster(ctx, "postgresql", &rds.ClusterArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
				pulumi.String("us-west-2c"),
			},
			BackupRetentionPeriod: pulumi.Int(5),
			ClusterIdentifier:     pulumi.String("aurora-cluster-demo"),
			DatabaseName:          pulumi.String("mydb"),
			Engine:                pulumi.String("aurora-postgresql"),
			MasterPassword:        pulumi.String("bar"),
			MasterUsername:        pulumi.String("foo"),
			PreferredBackupWindow: pulumi.String("07:00-09:00"),
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
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
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
        var postgresql = new Cluster("postgresql", ClusterArgs.builder()        
            .availabilityZones(            
                "us-west-2a",
                "us-west-2b",
                "us-west-2c")
            .backupRetentionPeriod(5)
            .clusterIdentifier("aurora-cluster-demo")
            .databaseName("mydb")
            .engine("aurora-postgresql")
            .masterPassword("bar")
            .masterUsername("foo")
            .preferredBackupWindow("07:00-09:00")
            .build());

    }
}
```
```yaml
resources:
  postgresql:
    type: aws:rds:Cluster
    properties:
      availabilityZones:
        - us-west-2a
        - us-west-2b
        - us-west-2c
      backupRetentionPeriod: 5
      clusterIdentifier: aurora-cluster-demo
      databaseName: mydb
      engine: aurora-postgresql
      masterPassword: bar
      masterUsername: foo
      preferredBackupWindow: 07:00-09:00
```
{{% /example %}}
{{% example %}}
### RDS Multi-AZ Cluster

> More information about RDS Multi-AZ Clusters can be found in the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html).

To create a Multi-AZ RDS cluster, you must additionally specify the `engine`, `storage_type`, `allocated_storage`, `iops` and `db_cluster_instance_class` attributes.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.Cluster("example", {
    allocatedStorage: 100,
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    clusterIdentifier: "example",
    dbClusterInstanceClass: "db.r6gd.xlarge",
    engine: "mysql",
    iops: 1000,
    masterPassword: "mustbeeightcharaters",
    masterUsername: "test",
    storageType: "io1",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.Cluster("example",
    allocated_storage=100,
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    cluster_identifier="example",
    db_cluster_instance_class="db.r6gd.xlarge",
    engine="mysql",
    iops=1000,
    master_password="mustbeeightcharaters",
    master_username="test",
    storage_type="io1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rds.Cluster("example", new()
    {
        AllocatedStorage = 100,
        AvailabilityZones = new[]
        {
            "us-west-2a",
            "us-west-2b",
            "us-west-2c",
        },
        ClusterIdentifier = "example",
        DbClusterInstanceClass = "db.r6gd.xlarge",
        Engine = "mysql",
        Iops = 1000,
        MasterPassword = "mustbeeightcharaters",
        MasterUsername = "test",
        StorageType = "io1",
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
		_, err := rds.NewCluster(ctx, "example", &rds.ClusterArgs{
			AllocatedStorage: pulumi.Int(100),
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
				pulumi.String("us-west-2c"),
			},
			ClusterIdentifier:      pulumi.String("example"),
			DbClusterInstanceClass: pulumi.String("db.r6gd.xlarge"),
			Engine:                 pulumi.String("mysql"),
			Iops:                   pulumi.Int(1000),
			MasterPassword:         pulumi.String("mustbeeightcharaters"),
			MasterUsername:         pulumi.String("test"),
			StorageType:            pulumi.String("io1"),
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
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
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
        var example = new Cluster("example", ClusterArgs.builder()        
            .allocatedStorage(100)
            .availabilityZones(            
                "us-west-2a",
                "us-west-2b",
                "us-west-2c")
            .clusterIdentifier("example")
            .dbClusterInstanceClass("db.r6gd.xlarge")
            .engine("mysql")
            .iops(1000)
            .masterPassword("mustbeeightcharaters")
            .masterUsername("test")
            .storageType("io1")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rds:Cluster
    properties:
      allocatedStorage: 100
      availabilityZones:
        - us-west-2a
        - us-west-2b
        - us-west-2c
      clusterIdentifier: example
      dbClusterInstanceClass: db.r6gd.xlarge
      engine: mysql
      iops: 1000
      masterPassword: mustbeeightcharaters
      masterUsername: test
      storageType: io1
```
{{% /example %}}
{{% example %}}
### RDS Serverless v2 Cluster

> More information about RDS Serverless v2 Clusters can be found in the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html).

To create a Serverless v2 RDS cluster, you must additionally specify the `engine_mode` and `serverlessv2_scaling_configuration` attributes. An `aws.rds.ClusterInstance` resource must also be added to the cluster with the `instance_class` attribute specified.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCluster = new aws.rds.Cluster("exampleCluster", {
    clusterIdentifier: "example",
    engine: "aurora-postgresql",
    engineMode: "provisioned",
    engineVersion: "13.6",
    databaseName: "test",
    masterUsername: "test",
    masterPassword: "must_be_eight_characters",
    serverlessv2ScalingConfiguration: {
        maxCapacity: 1,
        minCapacity: 0.5,
    },
});
const exampleClusterInstance = new aws.rds.ClusterInstance("exampleClusterInstance", {
    clusterIdentifier: exampleCluster.id,
    instanceClass: "db.serverless",
    engine: exampleCluster.engine,
    engineVersion: exampleCluster.engineVersion,
});
```
```python
import pulumi
import pulumi_aws as aws

example_cluster = aws.rds.Cluster("exampleCluster",
    cluster_identifier="example",
    engine="aurora-postgresql",
    engine_mode="provisioned",
    engine_version="13.6",
    database_name="test",
    master_username="test",
    master_password="must_be_eight_characters",
    serverlessv2_scaling_configuration=aws.rds.ClusterServerlessv2ScalingConfigurationArgs(
        max_capacity=1,
        min_capacity=0.5,
    ))
example_cluster_instance = aws.rds.ClusterInstance("exampleClusterInstance",
    cluster_identifier=example_cluster.id,
    instance_class="db.serverless",
    engine=example_cluster.engine,
    engine_version=example_cluster.engine_version)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleCluster = new Aws.Rds.Cluster("exampleCluster", new()
    {
        ClusterIdentifier = "example",
        Engine = "aurora-postgresql",
        EngineMode = "provisioned",
        EngineVersion = "13.6",
        DatabaseName = "test",
        MasterUsername = "test",
        MasterPassword = "must_be_eight_characters",
        Serverlessv2ScalingConfiguration = new Aws.Rds.Inputs.ClusterServerlessv2ScalingConfigurationArgs
        {
            MaxCapacity = 1,
            MinCapacity = 0.5,
        },
    });

    var exampleClusterInstance = new Aws.Rds.ClusterInstance("exampleClusterInstance", new()
    {
        ClusterIdentifier = exampleCluster.Id,
        InstanceClass = "db.serverless",
        Engine = exampleCluster.Engine,
        EngineVersion = exampleCluster.EngineVersion,
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
		exampleCluster, err := rds.NewCluster(ctx, "exampleCluster", &rds.ClusterArgs{
			ClusterIdentifier: pulumi.String("example"),
			Engine:            pulumi.String("aurora-postgresql"),
			EngineMode:        pulumi.String("provisioned"),
			EngineVersion:     pulumi.String("13.6"),
			DatabaseName:      pulumi.String("test"),
			MasterUsername:    pulumi.String("test"),
			MasterPassword:    pulumi.String("must_be_eight_characters"),
			Serverlessv2ScalingConfiguration: &rds.ClusterServerlessv2ScalingConfigurationArgs{
				MaxCapacity: pulumi.Float64(1),
				MinCapacity: pulumi.Float64(0.5),
			},
		})
		if err != nil {
			return err
		}
		_, err = rds.NewClusterInstance(ctx, "exampleClusterInstance", &rds.ClusterInstanceArgs{
			ClusterIdentifier: exampleCluster.ID(),
			InstanceClass:     pulumi.String("db.serverless"),
			Engine:            exampleCluster.Engine,
			EngineVersion:     exampleCluster.EngineVersion,
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
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
import com.pulumi.aws.rds.inputs.ClusterServerlessv2ScalingConfigurationArgs;
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
        var exampleCluster = new Cluster("exampleCluster", ClusterArgs.builder()        
            .clusterIdentifier("example")
            .engine("aurora-postgresql")
            .engineMode("provisioned")
            .engineVersion("13.6")
            .databaseName("test")
            .masterUsername("test")
            .masterPassword("must_be_eight_characters")
            .serverlessv2ScalingConfiguration(ClusterServerlessv2ScalingConfigurationArgs.builder()
                .maxCapacity(1)
                .minCapacity(0.5)
                .build())
            .build());

        var exampleClusterInstance = new ClusterInstance("exampleClusterInstance", ClusterInstanceArgs.builder()        
            .clusterIdentifier(exampleCluster.id())
            .instanceClass("db.serverless")
            .engine(exampleCluster.engine())
            .engineVersion(exampleCluster.engineVersion())
            .build());

    }
}
```
```yaml
resources:
  exampleCluster:
    type: aws:rds:Cluster
    properties:
      clusterIdentifier: example
      engine: aurora-postgresql
      engineMode: provisioned
      engineVersion: '13.6'
      databaseName: test
      masterUsername: test
      masterPassword: must_be_eight_characters
      serverlessv2ScalingConfiguration:
        maxCapacity: 1
        minCapacity: 0.5
  exampleClusterInstance:
    type: aws:rds:ClusterInstance
    properties:
      clusterIdentifier: ${exampleCluster.id}
      instanceClass: db.serverless
      engine: ${exampleCluster.engine}
      engineVersion: ${exampleCluster.engineVersion}
```
{{% /example %}}
{{% example %}}
### RDS/Aurora Managed Master Passwords via Secrets Manager, default KMS Key

> More information about RDS/Aurora Aurora integrates with Secrets Manager to manage master user passwords for your DB clusters can be found in the [RDS User Guide](https://aws.amazon.com/about-aws/whats-new/2022/12/amazon-rds-integration-aws-secrets-manager/) and [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html).

You can specify the `manage_master_user_password` attribute to enable managing the master password with Secrets Manager. You can also update an existing cluster to use Secrets Manager by specify the `manage_master_user_password` attribute and removing the `master_password` attribute (removal is required).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.rds.Cluster("test", {
    clusterIdentifier: "example",
    databaseName: "test",
    manageMasterUserPassword: true,
    masterUsername: "test",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.rds.Cluster("test",
    cluster_identifier="example",
    database_name="test",
    manage_master_user_password=True,
    master_username="test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Rds.Cluster("test", new()
    {
        ClusterIdentifier = "example",
        DatabaseName = "test",
        ManageMasterUserPassword = true,
        MasterUsername = "test",
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
		_, err := rds.NewCluster(ctx, "test", &rds.ClusterArgs{
			ClusterIdentifier:        pulumi.String("example"),
			DatabaseName:             pulumi.String("test"),
			ManageMasterUserPassword: pulumi.Bool(true),
			MasterUsername:           pulumi.String("test"),
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
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
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
        var test = new Cluster("test", ClusterArgs.builder()        
            .clusterIdentifier("example")
            .databaseName("test")
            .manageMasterUserPassword(true)
            .masterUsername("test")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:rds:Cluster
    properties:
      clusterIdentifier: example
      databaseName: test
      manageMasterUserPassword: true
      masterUsername: test
```
{{% /example %}}
{{% example %}}
### RDS/Aurora Managed Master Passwords via Secrets Manager, specific KMS Key

> More information about RDS/Aurora Aurora integrates with Secrets Manager to manage master user passwords for your DB clusters can be found in the [RDS User Guide](https://aws.amazon.com/about-aws/whats-new/2022/12/amazon-rds-integration-aws-secrets-manager/) and [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html).

You can specify the `master_user_secret_kms_key_id` attribute to specify a specific KMS Key.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kms.Key("example", {description: "Example KMS Key"});
const test = new aws.rds.Cluster("test", {
    clusterIdentifier: "example",
    databaseName: "test",
    manageMasterUserPassword: true,
    masterUsername: "test",
    masterUserSecretKmsKeyId: example.keyId,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kms.Key("example", description="Example KMS Key")
test = aws.rds.Cluster("test",
    cluster_identifier="example",
    database_name="test",
    manage_master_user_password=True,
    master_username="test",
    master_user_secret_kms_key_id=example.key_id)
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

    var test = new Aws.Rds.Cluster("test", new()
    {
        ClusterIdentifier = "example",
        DatabaseName = "test",
        ManageMasterUserPassword = true,
        MasterUsername = "test",
        MasterUserSecretKmsKeyId = example.KeyId,
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
		_, err = rds.NewCluster(ctx, "test", &rds.ClusterArgs{
			ClusterIdentifier:        pulumi.String("example"),
			DatabaseName:             pulumi.String("test"),
			ManageMasterUserPassword: pulumi.Bool(true),
			MasterUsername:           pulumi.String("test"),
			MasterUserSecretKmsKeyId: example.KeyId,
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
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
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

        var test = new Cluster("test", ClusterArgs.builder()        
            .clusterIdentifier("example")
            .databaseName("test")
            .manageMasterUserPassword(true)
            .masterUsername("test")
            .masterUserSecretKmsKeyId(example.keyId())
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
  test:
    type: aws:rds:Cluster
    properties:
      clusterIdentifier: example
      databaseName: test
      manageMasterUserPassword: true
      masterUsername: test
      masterUserSecretKmsKeyId: ${example.keyId}
```
{{% /example %}}
{{% example %}}
### Global Cluster Restored From Snapshot

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleClusterSnapshot = aws.rds.getClusterSnapshot({
    dbClusterIdentifier: "example-original-cluster",
    mostRecent: true,
});
const exampleCluster = new aws.rds.Cluster("exampleCluster", {
    engine: "aurora",
    engineVersion: "5.6.mysql_aurora.1.22.4",
    clusterIdentifier: "example",
    snapshotIdentifier: exampleClusterSnapshot.then(exampleClusterSnapshot => exampleClusterSnapshot.id),
});
const exampleGlobalCluster = new aws.rds.GlobalCluster("exampleGlobalCluster", {
    globalClusterIdentifier: "example",
    sourceDbClusterIdentifier: exampleCluster.arn,
    forceDestroy: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example_cluster_snapshot = aws.rds.get_cluster_snapshot(db_cluster_identifier="example-original-cluster",
    most_recent=True)
example_cluster = aws.rds.Cluster("exampleCluster",
    engine="aurora",
    engine_version="5.6.mysql_aurora.1.22.4",
    cluster_identifier="example",
    snapshot_identifier=example_cluster_snapshot.id)
example_global_cluster = aws.rds.GlobalCluster("exampleGlobalCluster",
    global_cluster_identifier="example",
    source_db_cluster_identifier=example_cluster.arn,
    force_destroy=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleClusterSnapshot = Aws.Rds.GetClusterSnapshot.Invoke(new()
    {
        DbClusterIdentifier = "example-original-cluster",
        MostRecent = true,
    });

    var exampleCluster = new Aws.Rds.Cluster("exampleCluster", new()
    {
        Engine = "aurora",
        EngineVersion = "5.6.mysql_aurora.1.22.4",
        ClusterIdentifier = "example",
        SnapshotIdentifier = exampleClusterSnapshot.Apply(getClusterSnapshotResult => getClusterSnapshotResult.Id),
    });

    var exampleGlobalCluster = new Aws.Rds.GlobalCluster("exampleGlobalCluster", new()
    {
        GlobalClusterIdentifier = "example",
        SourceDbClusterIdentifier = exampleCluster.Arn,
        ForceDestroy = true,
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
		exampleClusterSnapshot, err := rds.LookupClusterSnapshot(ctx, &rds.LookupClusterSnapshotArgs{
			DbClusterIdentifier: pulumi.StringRef("example-original-cluster"),
			MostRecent:          pulumi.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		exampleCluster, err := rds.NewCluster(ctx, "exampleCluster", &rds.ClusterArgs{
			Engine:             pulumi.String("aurora"),
			EngineVersion:      pulumi.String("5.6.mysql_aurora.1.22.4"),
			ClusterIdentifier:  pulumi.String("example"),
			SnapshotIdentifier: *pulumi.String(exampleClusterSnapshot.Id),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewGlobalCluster(ctx, "exampleGlobalCluster", &rds.GlobalClusterArgs{
			GlobalClusterIdentifier:   pulumi.String("example"),
			SourceDbClusterIdentifier: exampleCluster.Arn,
			ForceDestroy:              pulumi.Bool(true),
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
import com.pulumi.aws.rds.GlobalCluster;
import com.pulumi.aws.rds.GlobalClusterArgs;
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
        final var exampleClusterSnapshot = RdsFunctions.getClusterSnapshot(GetClusterSnapshotArgs.builder()
            .dbClusterIdentifier("example-original-cluster")
            .mostRecent(true)
            .build());

        var exampleCluster = new Cluster("exampleCluster", ClusterArgs.builder()        
            .engine("aurora")
            .engineVersion("5.6.mysql_aurora.1.22.4")
            .clusterIdentifier("example")
            .snapshotIdentifier(exampleClusterSnapshot.applyValue(getClusterSnapshotResult -> getClusterSnapshotResult.id()))
            .build());

        var exampleGlobalCluster = new GlobalCluster("exampleGlobalCluster", GlobalClusterArgs.builder()        
            .globalClusterIdentifier("example")
            .sourceDbClusterIdentifier(exampleCluster.arn())
            .forceDestroy(true)
            .build());

    }
}
```
```yaml
resources:
  exampleCluster:
    type: aws:rds:Cluster
    properties:
      # Because the global cluster is sourced from this cluster, the initial 
      #   # engine and engine_version values are defined here and automatically
      #   # inherited by the global cluster.
      engine: aurora
      engineVersion: 5.6.mysql_aurora.1.22.4
      clusterIdentifier: example
      snapshotIdentifier: ${exampleClusterSnapshot.id}
  exampleGlobalCluster:
    type: aws:rds:GlobalCluster
    properties:
      globalClusterIdentifier: example
      sourceDbClusterIdentifier: ${exampleCluster.arn}
      forceDestroy: true
variables:
  exampleClusterSnapshot:
    fn::invoke:
      Function: aws:rds:getClusterSnapshot
      Arguments:
        dbClusterIdentifier: example-original-cluster
        mostRecent: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import RDS Clusters using the `cluster_identifier`. For example:

```sh
 $ pulumi import aws:rds/cluster:Cluster aurora_cluster aurora-prod-cluster
```
 