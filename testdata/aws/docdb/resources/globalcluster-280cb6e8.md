Manages an DocumentDB Global Cluster. A global cluster consists of one primary region and up to five read-only secondary regions. You issue write operations directly to the primary cluster in the primary region and Amazon DocumentDB automatically replicates the data to the secondary regions using dedicated infrastructure.

More information about DocumentDB Global Clusters can be found in the [DocumentDB Developer Guide](https://docs.aws.amazon.com/documentdb/latest/developerguide/global-clusters.html).

{{% examples %}}
## Example Usage
{{% example %}}
### New DocumentDB Global Cluster

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.Provider("primary", {region: "us-east-2"});
const secondary = new aws.Provider("secondary", {region: "us-east-1"});
const example = new aws.docdb.GlobalCluster("example", {
    globalClusterIdentifier: "global-test",
    engine: "docdb",
    engineVersion: "4.0.0",
});
const primaryCluster = new aws.docdb.Cluster("primaryCluster", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    clusterIdentifier: "test-primary-cluster",
    masterUsername: "username",
    masterPassword: "somepass123",
    globalClusterIdentifier: example.id,
    dbSubnetGroupName: "default",
}, {
    provider: aws.primary,
});
const primaryClusterInstance = new aws.docdb.ClusterInstance("primaryClusterInstance", {
    engine: example.engine,
    identifier: "test-primary-cluster-instance",
    clusterIdentifier: primaryCluster.id,
    instanceClass: "db.r5.large",
}, {
    provider: aws.primary,
});
const secondaryCluster = new aws.docdb.Cluster("secondaryCluster", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    clusterIdentifier: "test-secondary-cluster",
    globalClusterIdentifier: example.id,
    dbSubnetGroupName: "default",
}, {
    provider: aws.secondary,
    dependsOn: [primaryCluster],
});
const secondaryClusterInstance = new aws.docdb.ClusterInstance("secondaryClusterInstance", {
    engine: example.engine,
    identifier: "test-secondary-cluster-instance",
    clusterIdentifier: secondaryCluster.id,
    instanceClass: "db.r5.large",
}, {
    provider: aws.secondary,
    dependsOn: [primaryClusterInstance],
});
```
```python
import pulumi
import pulumi_aws as aws

primary = aws.Provider("primary", region="us-east-2")
secondary = aws.Provider("secondary", region="us-east-1")
example = aws.docdb.GlobalCluster("example",
    global_cluster_identifier="global-test",
    engine="docdb",
    engine_version="4.0.0")
primary_cluster = aws.docdb.Cluster("primaryCluster",
    engine=example.engine,
    engine_version=example.engine_version,
    cluster_identifier="test-primary-cluster",
    master_username="username",
    master_password="somepass123",
    global_cluster_identifier=example.id,
    db_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
primary_cluster_instance = aws.docdb.ClusterInstance("primaryClusterInstance",
    engine=example.engine,
    identifier="test-primary-cluster-instance",
    cluster_identifier=primary_cluster.id,
    instance_class="db.r5.large",
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
secondary_cluster = aws.docdb.Cluster("secondaryCluster",
    engine=example.engine,
    engine_version=example.engine_version,
    cluster_identifier="test-secondary-cluster",
    global_cluster_identifier=example.id,
    db_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["secondary"],
        depends_on=[primary_cluster]))
secondary_cluster_instance = aws.docdb.ClusterInstance("secondaryClusterInstance",
    engine=example.engine,
    identifier="test-secondary-cluster-instance",
    cluster_identifier=secondary_cluster.id,
    instance_class="db.r5.large",
    opts=pulumi.ResourceOptions(provider=aws["secondary"],
        depends_on=[primary_cluster_instance]))
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

    var example = new Aws.DocDB.GlobalCluster("example", new()
    {
        GlobalClusterIdentifier = "global-test",
        Engine = "docdb",
        EngineVersion = "4.0.0",
    });

    var primaryCluster = new Aws.DocDB.Cluster("primaryCluster", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        ClusterIdentifier = "test-primary-cluster",
        MasterUsername = "username",
        MasterPassword = "somepass123",
        GlobalClusterIdentifier = example.Id,
        DbSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var primaryClusterInstance = new Aws.DocDB.ClusterInstance("primaryClusterInstance", new()
    {
        Engine = example.Engine,
        Identifier = "test-primary-cluster-instance",
        ClusterIdentifier = primaryCluster.Id,
        InstanceClass = "db.r5.large",
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var secondaryCluster = new Aws.DocDB.Cluster("secondaryCluster", new()
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
            primaryCluster,
        },
    });

    var secondaryClusterInstance = new Aws.DocDB.ClusterInstance("secondaryClusterInstance", new()
    {
        Engine = example.Engine,
        Identifier = "test-secondary-cluster-instance",
        ClusterIdentifier = secondaryCluster.Id,
        InstanceClass = "db.r5.large",
    }, new CustomResourceOptions
    {
        Provider = aws.Secondary,
        DependsOn = new[]
        {
            primaryClusterInstance,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/docdb"
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
		example, err := docdb.NewGlobalCluster(ctx, "example", &docdb.GlobalClusterArgs{
			GlobalClusterIdentifier: pulumi.String("global-test"),
			Engine:                  pulumi.String("docdb"),
			EngineVersion:           pulumi.String("4.0.0"),
		})
		if err != nil {
			return err
		}
		primaryCluster, err := docdb.NewCluster(ctx, "primaryCluster", &docdb.ClusterArgs{
			Engine:                  example.Engine,
			EngineVersion:           example.EngineVersion,
			ClusterIdentifier:       pulumi.String("test-primary-cluster"),
			MasterUsername:          pulumi.String("username"),
			MasterPassword:          pulumi.String("somepass123"),
			GlobalClusterIdentifier: example.ID(),
			DbSubnetGroupName:       pulumi.String("default"),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		primaryClusterInstance, err := docdb.NewClusterInstance(ctx, "primaryClusterInstance", &docdb.ClusterInstanceArgs{
			Engine:            example.Engine,
			Identifier:        pulumi.String("test-primary-cluster-instance"),
			ClusterIdentifier: primaryCluster.ID(),
			InstanceClass:     pulumi.String("db.r5.large"),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		secondaryCluster, err := docdb.NewCluster(ctx, "secondaryCluster", &docdb.ClusterArgs{
			Engine:                  example.Engine,
			EngineVersion:           example.EngineVersion,
			ClusterIdentifier:       pulumi.String("test-secondary-cluster"),
			GlobalClusterIdentifier: example.ID(),
			DbSubnetGroupName:       pulumi.String("default"),
		}, pulumi.Provider(aws.Secondary), pulumi.DependsOn([]pulumi.Resource{
			primaryCluster,
		}))
		if err != nil {
			return err
		}
		_, err = docdb.NewClusterInstance(ctx, "secondaryClusterInstance", &docdb.ClusterInstanceArgs{
			Engine:            example.Engine,
			Identifier:        pulumi.String("test-secondary-cluster-instance"),
			ClusterIdentifier: secondaryCluster.ID(),
			InstanceClass:     pulumi.String("db.r5.large"),
		}, pulumi.Provider(aws.Secondary), pulumi.DependsOn([]pulumi.Resource{
			primaryClusterInstance,
		}))
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
import com.pulumi.aws.docdb.GlobalCluster;
import com.pulumi.aws.docdb.GlobalClusterArgs;
import com.pulumi.aws.docdb.Cluster;
import com.pulumi.aws.docdb.ClusterArgs;
import com.pulumi.aws.docdb.ClusterInstance;
import com.pulumi.aws.docdb.ClusterInstanceArgs;
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
            .engine("docdb")
            .engineVersion("4.0.0")
            .build());

        var primaryCluster = new Cluster("primaryCluster", ClusterArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .clusterIdentifier("test-primary-cluster")
            .masterUsername("username")
            .masterPassword("somepass123")
            .globalClusterIdentifier(example.id())
            .dbSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.primary())
                .build());

        var primaryClusterInstance = new ClusterInstance("primaryClusterInstance", ClusterInstanceArgs.builder()        
            .engine(example.engine())
            .identifier("test-primary-cluster-instance")
            .clusterIdentifier(primaryCluster.id())
            .instanceClass("db.r5.large")
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
                .dependsOn(primaryCluster)
                .build());

        var secondaryClusterInstance = new ClusterInstance("secondaryClusterInstance", ClusterInstanceArgs.builder()        
            .engine(example.engine())
            .identifier("test-secondary-cluster-instance")
            .clusterIdentifier(secondaryCluster.id())
            .instanceClass("db.r5.large")
            .build(), CustomResourceOptions.builder()
                .provider(aws.secondary())
                .dependsOn(primaryClusterInstance)
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
    type: aws:docdb:GlobalCluster
    properties:
      globalClusterIdentifier: global-test
      engine: docdb
      engineVersion: 4.0.0
  primaryCluster:
    type: aws:docdb:Cluster
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      clusterIdentifier: test-primary-cluster
      masterUsername: username
      masterPassword: somepass123
      globalClusterIdentifier: ${example.id}
      dbSubnetGroupName: default
    options:
      provider: ${aws.primary}
  primaryClusterInstance:
    type: aws:docdb:ClusterInstance
    properties:
      engine: ${example.engine}
      identifier: test-primary-cluster-instance
      clusterIdentifier: ${primaryCluster.id}
      instanceClass: db.r5.large
    options:
      provider: ${aws.primary}
  secondaryCluster:
    type: aws:docdb:Cluster
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      clusterIdentifier: test-secondary-cluster
      globalClusterIdentifier: ${example.id}
      dbSubnetGroupName: default
    options:
      provider: ${aws.secondary}
      dependson:
        - ${primaryCluster}
  secondaryClusterInstance:
    type: aws:docdb:ClusterInstance
    properties:
      engine: ${example.engine}
      identifier: test-secondary-cluster-instance
      clusterIdentifier: ${secondaryCluster.id}
      instanceClass: db.r5.large
    options:
      provider: ${aws.secondary}
      dependson:
        - ${primaryClusterInstance}
```
{{% /example %}}
{{% example %}}
### New Global Cluster From Existing DB Cluster

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ... other configuration ...
const exampleCluster = new aws.docdb.Cluster("exampleCluster", {});
const exampleGlobalCluster = new aws.docdb.GlobalCluster("exampleGlobalCluster", {
    globalClusterIdentifier: "example",
    sourceDbClusterIdentifier: exampleCluster.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configuration ...
example_cluster = aws.docdb.Cluster("exampleCluster")
example_global_cluster = aws.docdb.GlobalCluster("exampleGlobalCluster",
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
    var exampleCluster = new Aws.DocDB.Cluster("exampleCluster");

    var exampleGlobalCluster = new Aws.DocDB.GlobalCluster("exampleGlobalCluster", new()
    {
        GlobalClusterIdentifier = "example",
        SourceDbClusterIdentifier = exampleCluster.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/docdb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleCluster, err := docdb.NewCluster(ctx, "exampleCluster", nil)
		if err != nil {
			return err
		}
		_, err = docdb.NewGlobalCluster(ctx, "exampleGlobalCluster", &docdb.GlobalClusterArgs{
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
import com.pulumi.aws.docdb.Cluster;
import com.pulumi.aws.docdb.GlobalCluster;
import com.pulumi.aws.docdb.GlobalClusterArgs;
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
            .globalClusterIdentifier("example")
            .sourceDbClusterIdentifier(exampleCluster.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleCluster:
    type: aws:docdb:Cluster
  exampleGlobalCluster:
    type: aws:docdb:GlobalCluster
    properties:
      globalClusterIdentifier: example
      sourceDbClusterIdentifier: ${exampleCluster.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_docdb_global_cluster` using the Global Cluster identifier. For example:

```sh
 $ pulumi import aws:docdb/globalCluster:GlobalCluster example example
```
 Certain resource arguments, like `source_db_cluster_identifier`, do not have an API method for reading the information after creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

