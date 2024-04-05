Provides an RDS Cluster Instance Resource. A Cluster Instance Resource defines
attributes that are specific to a single instance in a RDS Cluster,
specifically running Amazon Aurora.

Unlike other RDS resources that support replication, with Amazon Aurora you do
not designate a primary and subsequent replicas. Instead, you simply add RDS
Instances and Aurora manages the replication. You can use the [count][5]
meta-parameter to make multiple instances and join them all to the same RDS
Cluster, or you may specify different Cluster Instance resources with various
`instance_class` sizes.

For more information on Amazon Aurora, see [Aurora on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Aurora.html) in the Amazon RDS User Guide.

> **NOTE:** Deletion Protection from the RDS service can only be enabled at the cluster level, not for individual cluster instances. You can still add the [`protect` CustomResourceOption](https://www.pulumi.com/docs/intro/concepts/programming-model/#protect) to this resource configuration if you desire protection from accidental deletion.

> **NOTE:** `aurora` is no longer a valid `engine` because of [Amazon Aurora's MySQL-Compatible Edition version 1 end of life](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.MySQL56.EOL.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.rds.Cluster("default", {
    clusterIdentifier: "aurora-cluster-demo",
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    databaseName: "mydb",
    masterUsername: "foo",
    masterPassword: "barbut8chars",
});
const clusterInstances: aws.rds.ClusterInstance[] = [];
for (const range = {value: 0}; range.value < 2; range.value++) {
    clusterInstances.push(new aws.rds.ClusterInstance(`clusterInstances-${range.value}`, {
        identifier: `aurora-cluster-demo-${range.value}`,
        clusterIdentifier: _default.id,
        instanceClass: "db.r4.large",
        engine: _default.engine,
        engineVersion: _default.engineVersion,
    }));
}
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.Cluster("default",
    cluster_identifier="aurora-cluster-demo",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    database_name="mydb",
    master_username="foo",
    master_password="barbut8chars")
cluster_instances = []
for range in [{"value": i} for i in range(0, 2)]:
    cluster_instances.append(aws.rds.ClusterInstance(f"clusterInstances-{range['value']}",
        identifier=f"aurora-cluster-demo-{range['value']}",
        cluster_identifier=default.id,
        instance_class="db.r4.large",
        engine=default.engine,
        engine_version=default.engine_version))
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
        ClusterIdentifier = "aurora-cluster-demo",
        AvailabilityZones = new[]
        {
            "us-west-2a",
            "us-west-2b",
            "us-west-2c",
        },
        DatabaseName = "mydb",
        MasterUsername = "foo",
        MasterPassword = "barbut8chars",
    });

    var clusterInstances = new List<Aws.Rds.ClusterInstance>();
    for (var rangeIndex = 0; rangeIndex < 2; rangeIndex++)
    {
        var range = new { Value = rangeIndex };
        clusterInstances.Add(new Aws.Rds.ClusterInstance($"clusterInstances-{range.Value}", new()
        {
            Identifier = $"aurora-cluster-demo-{range.Value}",
            ClusterIdentifier = @default.Id,
            InstanceClass = "db.r4.large",
            Engine = @default.Engine,
            EngineVersion = @default.EngineVersion,
        }));
    }
});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "default", &rds.ClusterArgs{
			ClusterIdentifier: pulumi.String("aurora-cluster-demo"),
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
				pulumi.String("us-west-2c"),
			},
			DatabaseName:   pulumi.String("mydb"),
			MasterUsername: pulumi.String("foo"),
			MasterPassword: pulumi.String("barbut8chars"),
		})
		if err != nil {
			return err
		}
		var clusterInstances []*rds.ClusterInstance
		for index := 0; index < 2; index++ {
			key0 := index
			val0 := index
			__res, err := rds.NewClusterInstance(ctx, fmt.Sprintf("clusterInstances-%v", key0), &rds.ClusterInstanceArgs{
				Identifier:        pulumi.String(fmt.Sprintf("aurora-cluster-demo-%v", val0)),
				ClusterIdentifier: _default.ID(),
				InstanceClass:     pulumi.String("db.r4.large"),
				Engine:            _default.Engine,
				EngineVersion:     _default.EngineVersion,
			})
			if err != nil {
				return err
			}
			clusterInstances = append(clusterInstances, __res)
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
import com.pulumi.aws.rds.ClusterInstance;
import com.pulumi.aws.rds.ClusterInstanceArgs;
import com.pulumi.codegen.internal.KeyedValue;
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
            .clusterIdentifier("aurora-cluster-demo")
            .availabilityZones(            
                "us-west-2a",
                "us-west-2b",
                "us-west-2c")
            .databaseName("mydb")
            .masterUsername("foo")
            .masterPassword("barbut8chars")
            .build());

        for (var i = 0; i < 2; i++) {
            new ClusterInstance("clusterInstances-" + i, ClusterInstanceArgs.builder()            
                .identifier(String.format("aurora-cluster-demo-%s", range.value()))
                .clusterIdentifier(default_.id())
                .instanceClass("db.r4.large")
                .engine(default_.engine())
                .engineVersion(default_.engineVersion())
                .build());

        
}
    }
}
```
```yaml
resources:
  clusterInstances:
    type: aws:rds:ClusterInstance
    properties:
      identifier: aurora-cluster-demo-${range.value}
      clusterIdentifier: ${default.id}
      instanceClass: db.r4.large
      engine: ${default.engine}
      engineVersion: ${default.engineVersion}
    options: {}
  default:
    type: aws:rds:Cluster
    properties:
      clusterIdentifier: aurora-cluster-demo
      availabilityZones:
        - us-west-2a
        - us-west-2b
        - us-west-2c
      databaseName: mydb
      masterUsername: foo
      masterPassword: barbut8chars
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import RDS Cluster Instances using the `identifier`. For example:

```sh
 $ pulumi import aws:rds/clusterInstance:ClusterInstance prod_instance_1 aurora-cluster-instance-1
```
 