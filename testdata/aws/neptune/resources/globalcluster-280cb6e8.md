Manages a Neptune Global Cluster. A global cluster consists of one primary region and up to five read-only secondary regions. You issue write operations directly to the primary cluster in the primary region and Amazon Neptune automatically replicates the data to the secondary regions using dedicated infrastructure.

More information about Neptune Global Clusters can be found in the [Neptune User Guide](https://docs.aws.amazon.com/neptune/latest/userguide/neptune-global-database.html).

{{% examples %}}
## Example Usage
{{% example %}}
### New Neptune Global Cluster

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.Provider("primary", {region: "us-east-2"});
const secondary = new aws.Provider("secondary", {region: "us-east-1"});
const example = new aws.neptune.GlobalCluster("example", {
    globalClusterIdentifier: "global-test",
    engine: "neptune",
    engineVersion: "1.2.0.0",
});
const primaryCluster = new aws.neptune.Cluster("primaryCluster", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    clusterIdentifier: "test-primary-cluster",
    globalClusterIdentifier: example.id,
    neptuneSubnetGroupName: "default",
}, {
    provider: aws.primary,
});
const primaryClusterInstance = new aws.neptune.ClusterInstance("primaryClusterInstance", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    identifier: "test-primary-cluster-instance",
    clusterIdentifier: primaryCluster.id,
    instanceClass: "db.r5.large",
    neptuneSubnetGroupName: "default",
}, {
    provider: aws.primary,
});
const secondaryCluster = new aws.neptune.Cluster("secondaryCluster", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    clusterIdentifier: "test-secondary-cluster",
    globalClusterIdentifier: example.id,
    neptuneSubnetGroupName: "default",
}, {
    provider: aws.secondary,
});
const secondaryClusterInstance = new aws.neptune.ClusterInstance("secondaryClusterInstance", {
    engine: example.engine,
    engineVersion: example.engineVersion,
    identifier: "test-secondary-cluster-instance",
    clusterIdentifier: secondaryCluster.id,
    instanceClass: "db.r5.large",
    neptuneSubnetGroupName: "default",
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
example = aws.neptune.GlobalCluster("example",
    global_cluster_identifier="global-test",
    engine="neptune",
    engine_version="1.2.0.0")
primary_cluster = aws.neptune.Cluster("primaryCluster",
    engine=example.engine,
    engine_version=example.engine_version,
    cluster_identifier="test-primary-cluster",
    global_cluster_identifier=example.id,
    neptune_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
primary_cluster_instance = aws.neptune.ClusterInstance("primaryClusterInstance",
    engine=example.engine,
    engine_version=example.engine_version,
    identifier="test-primary-cluster-instance",
    cluster_identifier=primary_cluster.id,
    instance_class="db.r5.large",
    neptune_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
secondary_cluster = aws.neptune.Cluster("secondaryCluster",
    engine=example.engine,
    engine_version=example.engine_version,
    cluster_identifier="test-secondary-cluster",
    global_cluster_identifier=example.id,
    neptune_subnet_group_name="default",
    opts=pulumi.ResourceOptions(provider=aws["secondary"]))
secondary_cluster_instance = aws.neptune.ClusterInstance("secondaryClusterInstance",
    engine=example.engine,
    engine_version=example.engine_version,
    identifier="test-secondary-cluster-instance",
    cluster_identifier=secondary_cluster.id,
    instance_class="db.r5.large",
    neptune_subnet_group_name="default",
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

    var example = new Aws.Neptune.GlobalCluster("example", new()
    {
        GlobalClusterIdentifier = "global-test",
        Engine = "neptune",
        EngineVersion = "1.2.0.0",
    });

    var primaryCluster = new Aws.Neptune.Cluster("primaryCluster", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        ClusterIdentifier = "test-primary-cluster",
        GlobalClusterIdentifier = example.Id,
        NeptuneSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var primaryClusterInstance = new Aws.Neptune.ClusterInstance("primaryClusterInstance", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        Identifier = "test-primary-cluster-instance",
        ClusterIdentifier = primaryCluster.Id,
        InstanceClass = "db.r5.large",
        NeptuneSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var secondaryCluster = new Aws.Neptune.Cluster("secondaryCluster", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        ClusterIdentifier = "test-secondary-cluster",
        GlobalClusterIdentifier = example.Id,
        NeptuneSubnetGroupName = "default",
    }, new CustomResourceOptions
    {
        Provider = aws.Secondary,
    });

    var secondaryClusterInstance = new Aws.Neptune.ClusterInstance("secondaryClusterInstance", new()
    {
        Engine = example.Engine,
        EngineVersion = example.EngineVersion,
        Identifier = "test-secondary-cluster-instance",
        ClusterIdentifier = secondaryCluster.Id,
        InstanceClass = "db.r5.large",
        NeptuneSubnetGroupName = "default",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/neptune"
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
		example, err := neptune.NewGlobalCluster(ctx, "example", &neptune.GlobalClusterArgs{
			GlobalClusterIdentifier: pulumi.String("global-test"),
			Engine:                  pulumi.String("neptune"),
			EngineVersion:           pulumi.String("1.2.0.0"),
		})
		if err != nil {
			return err
		}
		primaryCluster, err := neptune.NewCluster(ctx, "primaryCluster", &neptune.ClusterArgs{
			Engine:                  example.Engine,
			EngineVersion:           example.EngineVersion,
			ClusterIdentifier:       pulumi.String("test-primary-cluster"),
			GlobalClusterIdentifier: example.ID(),
			NeptuneSubnetGroupName:  pulumi.String("default"),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		primaryClusterInstance, err := neptune.NewClusterInstance(ctx, "primaryClusterInstance", &neptune.ClusterInstanceArgs{
			Engine:                 example.Engine,
			EngineVersion:          example.EngineVersion,
			Identifier:             pulumi.String("test-primary-cluster-instance"),
			ClusterIdentifier:      primaryCluster.ID(),
			InstanceClass:          pulumi.String("db.r5.large"),
			NeptuneSubnetGroupName: pulumi.String("default"),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		secondaryCluster, err := neptune.NewCluster(ctx, "secondaryCluster", &neptune.ClusterArgs{
			Engine:                  example.Engine,
			EngineVersion:           example.EngineVersion,
			ClusterIdentifier:       pulumi.String("test-secondary-cluster"),
			GlobalClusterIdentifier: example.ID(),
			NeptuneSubnetGroupName:  pulumi.String("default"),
		}, pulumi.Provider(aws.Secondary))
		if err != nil {
			return err
		}
		_, err = neptune.NewClusterInstance(ctx, "secondaryClusterInstance", &neptune.ClusterInstanceArgs{
			Engine:                 example.Engine,
			EngineVersion:          example.EngineVersion,
			Identifier:             pulumi.String("test-secondary-cluster-instance"),
			ClusterIdentifier:      secondaryCluster.ID(),
			InstanceClass:          pulumi.String("db.r5.large"),
			NeptuneSubnetGroupName: pulumi.String("default"),
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
import com.pulumi.aws.neptune.GlobalCluster;
import com.pulumi.aws.neptune.GlobalClusterArgs;
import com.pulumi.aws.neptune.Cluster;
import com.pulumi.aws.neptune.ClusterArgs;
import com.pulumi.aws.neptune.ClusterInstance;
import com.pulumi.aws.neptune.ClusterInstanceArgs;
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
            .engine("neptune")
            .engineVersion("1.2.0.0")
            .build());

        var primaryCluster = new Cluster("primaryCluster", ClusterArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .clusterIdentifier("test-primary-cluster")
            .globalClusterIdentifier(example.id())
            .neptuneSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.primary())
                .build());

        var primaryClusterInstance = new ClusterInstance("primaryClusterInstance", ClusterInstanceArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .identifier("test-primary-cluster-instance")
            .clusterIdentifier(primaryCluster.id())
            .instanceClass("db.r5.large")
            .neptuneSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.primary())
                .build());

        var secondaryCluster = new Cluster("secondaryCluster", ClusterArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .clusterIdentifier("test-secondary-cluster")
            .globalClusterIdentifier(example.id())
            .neptuneSubnetGroupName("default")
            .build(), CustomResourceOptions.builder()
                .provider(aws.secondary())
                .build());

        var secondaryClusterInstance = new ClusterInstance("secondaryClusterInstance", ClusterInstanceArgs.builder()        
            .engine(example.engine())
            .engineVersion(example.engineVersion())
            .identifier("test-secondary-cluster-instance")
            .clusterIdentifier(secondaryCluster.id())
            .instanceClass("db.r5.large")
            .neptuneSubnetGroupName("default")
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
    type: aws:neptune:GlobalCluster
    properties:
      globalClusterIdentifier: global-test
      engine: neptune
      engineVersion: 1.2.0.0
  primaryCluster:
    type: aws:neptune:Cluster
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      clusterIdentifier: test-primary-cluster
      globalClusterIdentifier: ${example.id}
      neptuneSubnetGroupName: default
    options:
      provider: ${aws.primary}
  primaryClusterInstance:
    type: aws:neptune:ClusterInstance
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      identifier: test-primary-cluster-instance
      clusterIdentifier: ${primaryCluster.id}
      instanceClass: db.r5.large
      neptuneSubnetGroupName: default
    options:
      provider: ${aws.primary}
  secondaryCluster:
    type: aws:neptune:Cluster
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      clusterIdentifier: test-secondary-cluster
      globalClusterIdentifier: ${example.id}
      neptuneSubnetGroupName: default
    options:
      provider: ${aws.secondary}
  secondaryClusterInstance:
    type: aws:neptune:ClusterInstance
    properties:
      engine: ${example.engine}
      engineVersion: ${example.engineVersion}
      identifier: test-secondary-cluster-instance
      clusterIdentifier: ${secondaryCluster.id}
      instanceClass: db.r5.large
      neptuneSubnetGroupName: default
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
const exampleCluster = new aws.neptune.Cluster("exampleCluster", {});
const exampleGlobalCluster = new aws.neptune.GlobalCluster("exampleGlobalCluster", {
    globalClusterIdentifier: "example",
    sourceDbClusterIdentifier: exampleCluster.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configuration ...
example_cluster = aws.neptune.Cluster("exampleCluster")
example_global_cluster = aws.neptune.GlobalCluster("exampleGlobalCluster",
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
    var exampleCluster = new Aws.Neptune.Cluster("exampleCluster");

    var exampleGlobalCluster = new Aws.Neptune.GlobalCluster("exampleGlobalCluster", new()
    {
        GlobalClusterIdentifier = "example",
        SourceDbClusterIdentifier = exampleCluster.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/neptune"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleCluster, err := neptune.NewCluster(ctx, "exampleCluster", nil)
		if err != nil {
			return err
		}
		_, err = neptune.NewGlobalCluster(ctx, "exampleGlobalCluster", &neptune.GlobalClusterArgs{
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
import com.pulumi.aws.neptune.Cluster;
import com.pulumi.aws.neptune.GlobalCluster;
import com.pulumi.aws.neptune.GlobalClusterArgs;
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
    type: aws:neptune:Cluster
  exampleGlobalCluster:
    type: aws:neptune:GlobalCluster
    properties:
      globalClusterIdentifier: example
      sourceDbClusterIdentifier: ${exampleCluster.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_neptune_global_cluster` using the Global Cluster identifier. For example:

```sh
 $ pulumi import aws:neptune/globalCluster:GlobalCluster example example
```
 Certain resource arguments, like `source_db_cluster_identifier`, do not have an API method for reading the information after creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

