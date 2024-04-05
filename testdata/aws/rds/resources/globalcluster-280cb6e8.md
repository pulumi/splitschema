Manages an RDS Global Cluster, which is an Aurora global database spread across multiple regions. The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.

More information about Aurora global databases can be found in the [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database-creating).

{{% examples %}}
## Example Usage
{{% example %}}
### New MySQL Global Cluster

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.GlobalCluster("example", {
    globalClusterIdentifier: "global-test",
    engine: "aurora",
    engineVersion: "5.6.mysql_aurora.1.22.2",
    databaseName: "example_db",
});
const primaryCluster = new aws.rds.Cluster("primaryCluster", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    clusterIdentifier: "test-primary-cluster",
    masterUsername: "username",
    masterPassword: "somepass123",
    databaseName: "example_db",
    globalClusterIdentifier: example.id,
    dbSubnetGroupName: "default",
}, {
    provider: aws.primary,
});
const primaryClusterInstance = new aws.rds.ClusterInstance("primaryClusterInstance", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    identifier: "test-primary-cluster-instance",
    clusterIdentifier: primaryCluster.id,
    instanceClass: "db.r4.large",
    dbSubnetGroupName: "default",
}, {
    provider: aws.primary,
});
const secondaryCluster = new aws.rds.Cluster("secondaryCluster", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    clusterIdentifier: "test-secondary-cluster",
    globalClusterIdentifier: example.id,
    dbSubnetGroupName: "default",
}, {
    provider: aws.secondary,
    dependsOn: [primaryClusterInstance],
});
const secondaryClusterInstance = new aws.rds.ClusterInstance("secondaryClusterInstance", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    identifier: "test-secondary-cluster-instance",
    clusterIdentifier: secondaryCluster.id,
    instanceClass: "db.r4.large",
    dbSubnetGroupName: "default",
}, {
    provider: aws.secondary,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.GlobalCluster("example",
    global_cluster_identifier="global-test",
    engine="aurora",
    engine_version="5.6.mysql_aurora.1.22.2",
    database_name="example_db")
primary_cluster = aws.rds.Cluster("primaryCluster",
    engine=example.engine,
    engine_version=example.engine_version,
    cluster_identifier="test-primary-cluster",
    master_username="username",
    master_password="somepass123",
    database_name="example_db",
    global_cluster_identifier=example.id,
    db_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
primary_cluster_instance = aws.rds.ClusterInstance("primaryClusterInstance",
    engine=example.engine,
    engine_version=example.engine_version,
    identifier="test-primary-cluster-instance",
    cluster_identifier=primary_cluster.id,
    instance_class="db.r4.large",
    db_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
secondary_cluster = aws.rds.Cluster("secondaryCluster",
    engine=example.engine,
    engine_version=example.engine_version,
    cluster_identifier="test-secondary-cluster",
    global_cluster_identifier=example.id,
    db_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["secondary"],
        depends_on=[primary_cluster_instance]))
secondary_cluster_instance = aws.rds.ClusterInstance("secondaryClusterInstance",
    engine=example.engine,
    engine_version=example.engine_version,
    identifier="test-secondary-cluster-instance",
    cluster_identifier=secondary_cluster.id,
    instance_class="db.r4.large",
    db_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["secondary"]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rds.GlobalCluster("example", new()
    {
        GlobalClusterIdentifier = "global-test",
        Engine = "aurora",
        EngineVersion = "5.6.mysql_aurora.1.22.2",
        DatabaseName = "example_db",
    });

    var primaryCluster = new Aws.Rds.Cluster("primaryCluster", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        ClusterIdentifier = "test-primary-cluster",
        MasterUsername = "username",
        MasterPassword = "somepass123",
        DatabaseName = "example_db",
        GlobalClusterIdentifier = example.Id,
        DbSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var primaryClusterInstance = new Aws.Rds.ClusterInstance("primaryClusterInstance", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        Identifier = "test-primary-cluster-instance",
        ClusterIdentifier = primaryCluster.Id,
        InstanceClass = "db.r4.large",
        DbSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var secondaryCluster = new Aws.Rds.Cluster("secondaryCluster", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        ClusterIdentifier = "test-secondary-cluster",
        GlobalClusterIdentifier = example.Id,
        DbSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Secondary,
        DependsOn = new[]
        {
            primaryClusterInstance,
        },
    });

    var secondaryClusterInstance = new Aws.Rds.ClusterInstance("secondaryClusterInstance", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        Identifier = "test-secondary-cluster-instance",
        ClusterIdentifier = secondaryCluster.Id,
        InstanceClass = "db.r4.large",
        DbSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Secondary,
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
		example, err := rds.NewGlobalCluster(ctx, "example", &rds.GlobalClusterArgs{
			GlobalClusterIdentifier: pulumi.String("global-test"),
			Engine:                  pulumi.String("aurora"),
			EngineVersion:           pulumi.String("5.6.mysql_aurora.1.22.2"),
			DatabaseName:            pulumi.String("example_db"),
		})
		if err != nil {
			return err
		}
		primaryCluster, err := rds.NewCluster(ctx, "primaryCluster", &rds.ClusterArgs{
			Engine:                  example.Engine,
			EngineVersion:           example.EngineVersion,
			ClusterIdentifier:       pulumi.String("test-primary-cluster"),
			MasterUsername:          pulumi.String("username"),
			MasterPassword:          pulumi.String("somepass123"),
			DatabaseName:            pulumi.String("example_db"),
			GlobalClusterIdentifier: example.ID(),
			DbSubnetGroupName:       pulumi.String("default"),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		primaryClusterInstance, err := rds.NewClusterInstance(ctx, "primaryClusterInstance", &rds.ClusterInstanceArgs{
			Engine:            example.Engine,
			EngineVersion:     example.EngineVersion,
			Identifier:        pulumi.String("test-primary-cluster-instance"),
			ClusterIdentifier: primaryCluster.ID(),
			InstanceClass:     pulumi.String("db.r4.large"),
			DbSubnetGroupName: pulumi.String("default"),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		secondaryCluster, err := rds.NewCluster(ctx, "secondaryCluster", &rds.ClusterArgs{
			Engine:                  example.Engine,
			EngineVersion:           example.EngineVersion,
			ClusterIdentifier:       pulumi.String("test-secondary-cluster"),
			GlobalClusterIdentifier: example.ID(),
			DbSubnetGroupName:       pulumi.String("default"),
		}, pulumi.Provider(aws.Secondary), pulumi.DependsOn([]pulumi.Resource{
			primaryClusterInstance,
		}))
		if err != nil {
			return err
		}
		_, err = rds.NewClusterInstance(ctx, "secondaryClusterInstance", &rds.ClusterInstanceArgs{
			Engine:            example.Engine,
			EngineVersion:     example.EngineVersion,
			Identifier:        pulumi.String("test-secondary-cluster-instance"),
			ClusterIdentifier: secondaryCluster.ID(),
			InstanceClass:     pulumi.String("db.r4.large"),
			DbSubnetGroupName: pulumi.String("default"),
		}, pulumi.Provider(aws.Secondary))
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
import com.pulumi.aws.rds.GlobalCluster;
import com.pulumi.aws.rds.GlobalClusterArgs;
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
import com.pulumi.aws.rds.ClusterInstance;
import com.pulumi.aws.rds.ClusterInstanceArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var example = new GlobalCluster("example", GlobalClusterArgs.builder()        
            .globalClusterIdentifier("global-test")
            .engine("aurora")
            .engineVersion("5.6.mysql_aurora.1.22.2")
            .databaseName("example_db")
            .build());

        var primaryCluster = new Cluster("primaryCluster", ClusterArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .clusterIdentifier("test-primary-cluster")
            .masterUsername("username")
            .masterPassword("somepass123")
            .databaseName("example_db")
            .globalClusterIdentifier(example.id())
            .dbSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.primary())
                .build());

        var primaryClusterInstance = new ClusterInstance("primaryClusterInstance", ClusterInstanceArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .identifier("test-primary-cluster-instance")
            .clusterIdentifier(primaryCluster.id())
            .instanceClass("db.r4.large")
            .dbSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.primary())
                .build());

        var secondaryCluster = new Cluster("secondaryCluster", ClusterArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .clusterIdentifier("test-secondary-cluster")
            .globalClusterIdentifier(example.id())
            .dbSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.secondary())
                .dependsOn(primaryClusterInstance)
                .build());

        var secondaryClusterInstance = new ClusterInstance("secondaryClusterInstance", ClusterInstanceArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .identifier("test-secondary-cluster-instance")
            .clusterIdentifier(secondaryCluster.id())
            .instanceClass("db.r4.large")
            .dbSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.secondary())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rds:GlobalCluster
    properties:
      globalClusterIdentifier: global-test
      engine: aurora
      engineVersion: 5.6.mysql_aurora.1.22.2
      databaseName: example_db
  primaryCluster:
    type: aws:rds:Cluster
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      clusterIdentifier: test-primary-cluster
      masterUsername: username
      masterPassword: somepass123
      databaseName: example_db
      globalClusterIdentifier: ${example.id}
      dbSubnetGroupName: default
    options:
      provider: ${aws.primary}
  primaryClusterInstance:
    type: aws:rds:ClusterInstance
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      identifier: test-primary-cluster-instance
      clusterIdentifier: ${primaryCluster.id}
      instanceClass: db.r4.large
      dbSubnetGroupName: default
    options:
      provider: ${aws.primary}
  secondaryCluster:
    type: aws:rds:Cluster
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      clusterIdentifier: test-secondary-cluster
      globalClusterIdentifier: ${example.id}
      dbSubnetGroupName: default
    options:
      provider: ${aws.secondary}
      dependson:
        - ${primaryClusterInstance}
  secondaryClusterInstance:
    type: aws:rds:ClusterInstance
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      identifier: test-secondary-cluster-instance
      clusterIdentifier: ${secondaryCluster.id}
      instanceClass: db.r4.large
      dbSubnetGroupName: default
    options:
      provider: ${aws.secondary}
```
{{% /example %}}
{{% example %}}
### New PostgreSQL Global Cluster

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.Provider("primary", {region: "us-east-2"});
const secondary = new aws.Provider("secondary", {region: "us-east-1"});
const example = new aws.rds.GlobalCluster("example", {
    globalClusterIdentifier: "global-test",
    engine: "aurora-postgresql",
    engineVersion: "11.9",
    databaseName: "example_db",
});
const primaryCluster = new aws.rds.Cluster("primaryCluster", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    clusterIdentifier: "test-primary-cluster",
    masterUsername: "username",
    masterPassword: "somepass123",
    databaseName: "example_db",
    globalClusterIdentifier: example.id,
    dbSubnetGroupName: "default",
}, {
    provider: aws.primary,
});
const primaryClusterInstance = new aws.rds.ClusterInstance("primaryClusterInstance", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    identifier: "test-primary-cluster-instance",
    clusterIdentifier: primaryCluster.id,
    instanceClass: "db.r4.large",
    dbSubnetGroupName: "default",
}, {
    provider: aws.primary,
});
const secondaryCluster = new aws.rds.Cluster("secondaryCluster", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    clusterIdentifier: "test-secondary-cluster",
    globalClusterIdentifier: example.id,
    skipFinalSnapshot: true,
    dbSubnetGroupName: "default",
}, {
    provider: aws.secondary,
    dependsOn: [primaryClusterInstance],
});
const secondaryClusterInstance = new aws.rds.ClusterInstance("secondaryClusterInstance", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    identifier: "test-secondary-cluster-instance",
    clusterIdentifier: secondaryCluster.id,
    instanceClass: "db.r4.large",
    dbSubnetGroupName: "default",
}, {
    provider: aws.secondary,
});
```
```python
import pulumi
import pulumi_aws as aws

primary = aws.Provider("primary", region="us-east-2")
secondary = aws.Provider("secondary", region="us-east-1")
example = aws.rds.GlobalCluster("example",
    global_cluster_identifier="global-test",
    engine="aurora-postgresql",
    engine_version="11.9",
    database_name="example_db")
primary_cluster = aws.rds.Cluster("primaryCluster",
    engine=example.engine,
    engine_version=example.engine_version,
    cluster_identifier="test-primary-cluster",
    master_username="username",
    master_password="somepass123",
    database_name="example_db",
    global_cluster_identifier=example.id,
    db_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
primary_cluster_instance = aws.rds.ClusterInstance("primaryClusterInstance",
    engine=example.engine,
    engine_version=example.engine_version,
    identifier="test-primary-cluster-instance",
    cluster_identifier=primary_cluster.id,
    instance_class="db.r4.large",
    db_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
secondary_cluster = aws.rds.Cluster("secondaryCluster",
    engine=example.engine,
    engine_version=example.engine_version,
    cluster_identifier="test-secondary-cluster",
    global_cluster_identifier=example.id,
    skip_final_snapshot=True,
    db_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["secondary"],
        depends_on=[primary_cluster_instance]))
secondary_cluster_instance = aws.rds.ClusterInstance("secondaryClusterInstance",
    engine=example.engine,
    engine_version=example.engine_version,
    identifier="test-secondary-cluster-instance",
    cluster_identifier=secondary_cluster.id,
    instance_class="db.r4.large",
    db_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["secondary"]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var primary = new Aws.Provider("primary", new()
    {
        Region = "us-east-2",
    });

    var secondary = new Aws.Provider("secondary", new()
    {
        Region = "us-east-1",
    });

    var example = new Aws.Rds.GlobalCluster("example", new()
    {
        GlobalClusterIdentifier = "global-test",
        Engine = "aurora-postgresql",
        EngineVersion = "11.9",
        DatabaseName = "example_db",
    });

    var primaryCluster = new Aws.Rds.Cluster("primaryCluster", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        ClusterIdentifier = "test-primary-cluster",
        MasterUsername = "username",
        MasterPassword = "somepass123",
        DatabaseName = "example_db",
        GlobalClusterIdentifier = example.Id,
        DbSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var primaryClusterInstance = new Aws.Rds.ClusterInstance("primaryClusterInstance", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        Identifier = "test-primary-cluster-instance",
        ClusterIdentifier = primaryCluster.Id,
        InstanceClass = "db.r4.large",
        DbSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var secondaryCluster = new Aws.Rds.Cluster("secondaryCluster", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        ClusterIdentifier = "test-secondary-cluster",
        GlobalClusterIdentifier = example.Id,
        SkipFinalSnapshot = true,
        DbSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Secondary,
        DependsOn = new[]
        {
            primaryClusterInstance,
        },
    });

    var secondaryClusterInstance = new Aws.Rds.ClusterInstance("secondaryClusterInstance", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        Identifier = "test-secondary-cluster-instance",
        ClusterIdentifier = secondaryCluster.Id,
        InstanceClass = "db.r4.large",
        DbSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Secondary,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "primary", &aws.ProviderArgs{
			Region: pulumi.String("us-east-2"),
		})
		if err != nil {
			return err
		}
		_, err = aws.NewProvider(ctx, "secondary", &aws.ProviderArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		example, err := rds.NewGlobalCluster(ctx, "example", &rds.GlobalClusterArgs{
			GlobalClusterIdentifier: pulumi.String("global-test"),
			Engine:                  pulumi.String("aurora-postgresql"),
			EngineVersion:           pulumi.String("11.9"),
			DatabaseName:            pulumi.String("example_db"),
		})
		if err != nil {
			return err
		}
		primaryCluster, err := rds.NewCluster(ctx, "primaryCluster", &rds.ClusterArgs{
			Engine:                  example.Engine,
			EngineVersion:           example.EngineVersion,
			ClusterIdentifier:       pulumi.String("test-primary-cluster"),
			MasterUsername:          pulumi.String("username"),
			MasterPassword:          pulumi.String("somepass123"),
			DatabaseName:            pulumi.String("example_db"),
			GlobalClusterIdentifier: example.ID(),
			DbSubnetGroupName:       pulumi.String("default"),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		primaryClusterInstance, err := rds.NewClusterInstance(ctx, "primaryClusterInstance", &rds.ClusterInstanceArgs{
			Engine:            example.Engine,
			EngineVersion:     example.EngineVersion,
			Identifier:        pulumi.String("test-primary-cluster-instance"),
			ClusterIdentifier: primaryCluster.ID(),
			InstanceClass:     pulumi.String("db.r4.large"),
			DbSubnetGroupName: pulumi.String("default"),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		secondaryCluster, err := rds.NewCluster(ctx, "secondaryCluster", &rds.ClusterArgs{
			Engine:                  example.Engine,
			EngineVersion:           example.EngineVersion,
			ClusterIdentifier:       pulumi.String("test-secondary-cluster"),
			GlobalClusterIdentifier: example.ID(),
			SkipFinalSnapshot:       pulumi.Bool(true),
			DbSubnetGroupName:       pulumi.String("default"),
		}, pulumi.Provider(aws.Secondary), pulumi.DependsOn([]pulumi.Resource{
			primaryClusterInstance,
		}))
		if err != nil {
			return err
		}
		_, err = rds.NewClusterInstance(ctx, "secondaryClusterInstance", &rds.ClusterInstanceArgs{
			Engine:            example.Engine,
			EngineVersion:     example.EngineVersion,
			Identifier:        pulumi.String("test-secondary-cluster-instance"),
			ClusterIdentifier: secondaryCluster.ID(),
			InstanceClass:     pulumi.String("db.r4.large"),
			DbSubnetGroupName: pulumi.String("default"),
		}, pulumi.Provider(aws.Secondary))
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
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.rds.GlobalCluster;
import com.pulumi.aws.rds.GlobalClusterArgs;
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
import com.pulumi.aws.rds.ClusterInstance;
import com.pulumi.aws.rds.ClusterInstanceArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var primary = new Provider("primary", ProviderArgs.builder()        
            .region("us-east-2")
            .build());

        var secondary = new Provider("secondary", ProviderArgs.builder()        
            .region("us-east-1")
            .build());

        var example = new GlobalCluster("example", GlobalClusterArgs.builder()        
            .globalClusterIdentifier("global-test")
            .engine("aurora-postgresql")
            .engineVersion("11.9")
            .databaseName("example_db")
            .build());

        var primaryCluster = new Cluster("primaryCluster", ClusterArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .clusterIdentifier("test-primary-cluster")
            .masterUsername("username")
            .masterPassword("somepass123")
            .databaseName("example_db")
            .globalClusterIdentifier(example.id())
            .dbSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.primary())
                .build());

        var primaryClusterInstance = new ClusterInstance("primaryClusterInstance", ClusterInstanceArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .identifier("test-primary-cluster-instance")
            .clusterIdentifier(primaryCluster.id())
            .instanceClass("db.r4.large")
            .dbSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.primary())
                .build());

        var secondaryCluster = new Cluster("secondaryCluster", ClusterArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .clusterIdentifier("test-secondary-cluster")
            .globalClusterIdentifier(example.id())
            .skipFinalSnapshot(true)
            .dbSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.secondary())
                .dependsOn(primaryClusterInstance)
                .build());

        var secondaryClusterInstance = new ClusterInstance("secondaryClusterInstance", ClusterInstanceArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .identifier("test-secondary-cluster-instance")
            .clusterIdentifier(secondaryCluster.id())
            .instanceClass("db.r4.large")
            .dbSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.secondary())
                .build());

    }
}
```
```yaml
resources:
  primary:
    type: pulumi:providers:aws
    properties:
      region: us-east-2
  secondary:
    type: pulumi:providers:aws
    properties:
      region: us-east-1
  example:
    type: aws:rds:GlobalCluster
    properties:
      globalClusterIdentifier: global-test
      engine: aurora-postgresql
      engineVersion: '11.9'
      databaseName: example_db
  primaryCluster:
    type: aws:rds:Cluster
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      clusterIdentifier: test-primary-cluster
      masterUsername: username
      masterPassword: somepass123
      databaseName: example_db
      globalClusterIdentifier: ${example.id}
      dbSubnetGroupName: default
    options:
      provider: ${aws.primary}
  primaryClusterInstance:
    type: aws:rds:ClusterInstance
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      identifier: test-primary-cluster-instance
      clusterIdentifier: ${primaryCluster.id}
      instanceClass: db.r4.large
      dbSubnetGroupName: default
    options:
      provider: ${aws.primary}
  secondaryCluster:
    type: aws:rds:Cluster
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      clusterIdentifier: test-secondary-cluster
      globalClusterIdentifier: ${example.id}
      skipFinalSnapshot: true
      dbSubnetGroupName: default
    options:
      provider: ${aws.secondary}
      dependson:
        - ${primaryClusterInstance}
  secondaryClusterInstance:
    type: aws:rds:ClusterInstance
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      identifier: test-secondary-cluster-instance
      clusterIdentifier: ${secondaryCluster.id}
      instanceClass: db.r4.large
      dbSubnetGroupName: default
    options:
      provider: ${aws.secondary}
```
{{% /example %}}
{{% example %}}
### New Global Cluster From Existing DB Cluster

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ... other configuration ...
const exampleCluster = new aws.rds.Cluster("exampleCluster", {});
const exampleGlobalCluster = new aws.rds.GlobalCluster("exampleGlobalCluster", {
    forceDestroy: true,
    globalClusterIdentifier: "example",
    sourceDbClusterIdentifier: exampleCluster.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configuration ...
example_cluster = aws.rds.Cluster("exampleCluster")
example_global_cluster = aws.rds.GlobalCluster("exampleGlobalCluster",
    force_destroy=True,
    global_cluster_identifier="example",
    source_db_cluster_identifier=example_cluster.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ... other configuration ...
    var exampleCluster = new Aws.Rds.Cluster("exampleCluster");

    var exampleGlobalCluster = new Aws.Rds.GlobalCluster("exampleGlobalCluster", new()
    {
        ForceDestroy = true,
        GlobalClusterIdentifier = "example",
        SourceDbClusterIdentifier = exampleCluster.Arn,
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
		exampleCluster, err := rds.NewCluster(ctx, "exampleCluster", nil)
		if err != nil {
			return err
		}
		_, err = rds.NewGlobalCluster(ctx, "exampleGlobalCluster", &rds.GlobalClusterArgs{
			ForceDestroy:              pulumi.Bool(true),
			GlobalClusterIdentifier:   pulumi.String("example"),
			SourceDbClusterIdentifier: exampleCluster.Arn,
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
        var exampleCluster = new Cluster("exampleCluster");

        var exampleGlobalCluster = new GlobalCluster("exampleGlobalCluster", GlobalClusterArgs.builder()        
            .forceDestroy(true)
            .globalClusterIdentifier("example")
            .sourceDbClusterIdentifier(exampleCluster.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleCluster:
    type: aws:rds:Cluster
  exampleGlobalCluster:
    type: aws:rds:GlobalCluster
    properties:
      forceDestroy: true
      globalClusterIdentifier: example
      sourceDbClusterIdentifier: ${exampleCluster.arn}
```
{{% /example %}}
{{% example %}}
### Upgrading Engine Versions

When you upgrade the version of an `aws.rds.GlobalCluster`, the provider will attempt to in-place upgrade the engine versions of all associated clusters. Since the `aws.rds.Cluster` resource is being updated through the `aws.rds.GlobalCluster`, you are likely to get an error (`Provider produced inconsistent final plan`). To avoid this, use the `lifecycle` `ignore_changes` meta argument as shown below on the `aws.rds.Cluster`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.GlobalCluster("example", {
    globalClusterIdentifier: "kyivkharkiv",
    engine: "aurora-mysql",
    engineVersion: "5.7.mysql_aurora.2.07.5",
});
const primaryCluster = new aws.rds.Cluster("primaryCluster", {
    allowMajorVersionUpgrade: true,
    applyImmediately: true,
    clusterIdentifier: "odessadnipro",
    databaseName: "totoro",
    engine: example.engine,
    engineVersion: example.engineVersion,
    globalClusterIdentifier: example.id,
    masterPassword: "satsukimae",
    masterUsername: "maesatsuki",
    skipFinalSnapshot: true,
});
const primaryClusterInstance = new aws.rds.ClusterInstance("primaryClusterInstance", {
    applyImmediately: true,
    clusterIdentifier: primaryCluster.id,
    engine: primaryCluster.engine,
    engineVersion: primaryCluster.engineVersion,
    identifier: "donetsklviv",
    instanceClass: "db.r4.large",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.GlobalCluster("example",
    global_cluster_identifier="kyivkharkiv",
    engine="aurora-mysql",
    engine_version="5.7.mysql_aurora.2.07.5")
primary_cluster = aws.rds.Cluster("primaryCluster",
    allow_major_version_upgrade=True,
    apply_immediately=True,
    cluster_identifier="odessadnipro",
    database_name="totoro",
    engine=example.engine,
    engine_version=example.engine_version,
    global_cluster_identifier=example.id,
    master_password="satsukimae",
    master_username="maesatsuki",
    skip_final_snapshot=True)
primary_cluster_instance = aws.rds.ClusterInstance("primaryClusterInstance",
    apply_immediately=True,
    cluster_identifier=primary_cluster.id,
    engine=primary_cluster.engine,
    engine_version=primary_cluster.engine_version,
    identifier="donetsklviv",
    instance_class="db.r4.large")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rds.GlobalCluster("example", new()
    {
        GlobalClusterIdentifier = "kyivkharkiv",
        Engine = "aurora-mysql",
        EngineVersion = "5.7.mysql_aurora.2.07.5",
    });

    var primaryCluster = new Aws.Rds.Cluster("primaryCluster", new()
    {
        AllowMajorVersionUpgrade = true,
        ApplyImmediately = true,
        ClusterIdentifier = "odessadnipro",
        DatabaseName = "totoro",
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        GlobalClusterIdentifier = example.Id,
        MasterPassword = "satsukimae",
        MasterUsername = "maesatsuki",
        SkipFinalSnapshot = true,
    });

    var primaryClusterInstance = new Aws.Rds.ClusterInstance("primaryClusterInstance", new()
    {
        ApplyImmediately = true,
        ClusterIdentifier = primaryCluster.Id,
        Engine = primaryCluster.Engine,
        EngineVersion = primaryCluster.EngineVersion,
        Identifier = "donetsklviv",
        InstanceClass = "db.r4.large",
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
		example, err := rds.NewGlobalCluster(ctx, "example", &rds.GlobalClusterArgs{
			GlobalClusterIdentifier: pulumi.String("kyivkharkiv"),
			Engine:                  pulumi.String("aurora-mysql"),
			EngineVersion:           pulumi.String("5.7.mysql_aurora.2.07.5"),
		})
		if err != nil {
			return err
		}
		primaryCluster, err := rds.NewCluster(ctx, "primaryCluster", &rds.ClusterArgs{
			AllowMajorVersionUpgrade: pulumi.Bool(true),
			ApplyImmediately:         pulumi.Bool(true),
			ClusterIdentifier:        pulumi.String("odessadnipro"),
			DatabaseName:             pulumi.String("totoro"),
			Engine:                   example.Engine,
			EngineVersion:            example.EngineVersion,
			GlobalClusterIdentifier:  example.ID(),
			MasterPassword:           pulumi.String("satsukimae"),
			MasterUsername:           pulumi.String("maesatsuki"),
			SkipFinalSnapshot:        pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewClusterInstance(ctx, "primaryClusterInstance", &rds.ClusterInstanceArgs{
			ApplyImmediately:  pulumi.Bool(true),
			ClusterIdentifier: primaryCluster.ID(),
			Engine:            primaryCluster.Engine,
			EngineVersion:     primaryCluster.EngineVersion,
			Identifier:        pulumi.String("donetsklviv"),
			InstanceClass:     pulumi.String("db.r4.large"),
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
import com.pulumi.aws.rds.GlobalCluster;
import com.pulumi.aws.rds.GlobalClusterArgs;
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
        var example = new GlobalCluster("example", GlobalClusterArgs.builder()        
            .globalClusterIdentifier("kyivkharkiv")
            .engine("aurora-mysql")
            .engineVersion("5.7.mysql_aurora.2.07.5")
            .build());

        var primaryCluster = new Cluster("primaryCluster", ClusterArgs.builder()        
            .allowMajorVersionUpgrade(true)
            .applyImmediately(true)
            .clusterIdentifier("odessadnipro")
            .databaseName("totoro")
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .globalClusterIdentifier(example.id())
            .masterPassword("satsukimae")
            .masterUsername("maesatsuki")
            .skipFinalSnapshot(true)
            .build());

        var primaryClusterInstance = new ClusterInstance("primaryClusterInstance", ClusterInstanceArgs.builder()        
            .applyImmediately(true)
            .clusterIdentifier(primaryCluster.id())
            .engine(primaryCluster.engine())
            .engineVersion(primaryCluster.engineVersion())
            .identifier("donetsklviv")
            .instanceClass("db.r4.large")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rds:GlobalCluster
    properties:
      globalClusterIdentifier: kyivkharkiv
      engine: aurora-mysql
      engineVersion: 5.7.mysql_aurora.2.07.5
  primaryCluster:
    type: aws:rds:Cluster
    properties:
      allowMajorVersionUpgrade: true
      applyImmediately: true
      clusterIdentifier: odessadnipro
      databaseName: totoro
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      globalClusterIdentifier: ${example.id}
      masterPassword: satsukimae
      masterUsername: maesatsuki
      skipFinalSnapshot: true
  primaryClusterInstance:
    type: aws:rds:ClusterInstance
    properties:
      applyImmediately: true
      clusterIdentifier: ${primaryCluster.id}
      engine: ${primaryCluster.engine}
      engineVersion: ${primaryCluster.engineVersion}
      identifier: donetsklviv
      instanceClass: db.r4.large
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_rds_global_cluster` using the RDS Global Cluster identifier. For example:

```sh
 $ pulumi import aws:rds/globalCluster:GlobalCluster example example
```
 Certain resource arguments, like `force_destroy`, only exist within this provider. If the argument is set in the the provider configuration on an imported resource, This provider will show a difference on the first plan after import to update the state value. This change is safe to apply immediately so the state matches the desired configuration.

Certain resource arguments, like `source_db_cluster_identifier`, do not have an API method for reading the information after creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

