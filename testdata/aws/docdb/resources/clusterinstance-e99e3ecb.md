Provides an DocumentDB Cluster Resource Instance. A Cluster Instance Resource defines
attributes that are specific to a single instance in a [DocumentDB Cluster][1].

You do not designate a primary and subsequent replicas. Instead, you simply add DocumentDB
Instances and DocumentDB manages the replication. You can use the [count][3]
meta-parameter to make multiple instances and join them all to the same DocumentDB
Cluster, or you may specify different Cluster Instance resources with various
`instance_class` sizes.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.docdb.Cluster("default", {
    clusterIdentifier: "docdb-cluster-demo",
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    masterUsername: "foo",
    masterPassword: "barbut8chars",
});
const clusterInstances: aws.docdb.ClusterInstance[] = [];
for (const range = {value: 0}; range.value < 2; range.value++) {
    clusterInstances.push(new aws.docdb.ClusterInstance(`clusterInstances-${range.value}`, {
        identifier: `docdb-cluster-demo-${range.value}`,
        clusterIdentifier: _default.id,
        instanceClass: "db.r5.large",
    }));
}
```
```python
import pulumi
import pulumi_aws as aws

default = aws.docdb.Cluster("default",
    cluster_identifier="docdb-cluster-demo",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    master_username="foo",
    master_password="barbut8chars")
cluster_instances = []
for range in [{"value": i} for i in range(0, 2)]:
    cluster_instances.append(aws.docdb.ClusterInstance(f"clusterInstances-{range['value']}",
        identifier=f"docdb-cluster-demo-{range['value']}",
        cluster_identifier=default.id,
        instance_class="db.r5.large"))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.DocDB.Cluster("default", new()
    {
        ClusterIdentifier = "docdb-cluster-demo",
        AvailabilityZones = new[]
        {
            "us-west-2a",
            "us-west-2b",
            "us-west-2c",
        },
        MasterUsername = "foo",
        MasterPassword = "barbut8chars",
    });

    var clusterInstances = new List<Aws.DocDB.ClusterInstance>();
    for (var rangeIndex = 0; rangeIndex < 2; rangeIndex++)
    {
        var range = new { Value = rangeIndex };
        clusterInstances.Add(new Aws.DocDB.ClusterInstance($"clusterInstances-{range.Value}", new()
        {
            Identifier = $"docdb-cluster-demo-{range.Value}",
            ClusterIdentifier = @default.Id,
            InstanceClass = "db.r5.large",
        }));
    }
});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/docdb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := docdb.NewCluster(ctx, "default", &docdb.ClusterArgs{
			ClusterIdentifier: pulumi.String("docdb-cluster-demo"),
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
				pulumi.String("us-west-2c"),
			},
			MasterUsername: pulumi.String("foo"),
			MasterPassword: pulumi.String("barbut8chars"),
		})
		if err != nil {
			return err
		}
		var clusterInstances []*docdb.ClusterInstance
		for index := 0; index < 2; index++ {
			key0 := index
			val0 := index
			__res, err := docdb.NewClusterInstance(ctx, fmt.Sprintf("clusterInstances-%v", key0), &docdb.ClusterInstanceArgs{
				Identifier:        pulumi.String(fmt.Sprintf("docdb-cluster-demo-%v", val0)),
				ClusterIdentifier: _default.ID(),
				InstanceClass:     pulumi.String("db.r5.large"),
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
import com.pulumi.aws.docdb.Cluster;
import com.pulumi.aws.docdb.ClusterArgs;
import com.pulumi.aws.docdb.ClusterInstance;
import com.pulumi.aws.docdb.ClusterInstanceArgs;
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
            .clusterIdentifier("docdb-cluster-demo")
            .availabilityZones(            
                "us-west-2a",
                "us-west-2b",
                "us-west-2c")
            .masterUsername("foo")
            .masterPassword("barbut8chars")
            .build());

        for (var i = 0; i < 2; i++) {
            new ClusterInstance("clusterInstances-" + i, ClusterInstanceArgs.builder()            
                .identifier(String.format("docdb-cluster-demo-%s", range.value()))
                .clusterIdentifier(default_.id())
                .instanceClass("db.r5.large")
                .build());

        
}
    }
}
```
```yaml
resources:
  clusterInstances:
    type: aws:docdb:ClusterInstance
    properties:
      identifier: docdb-cluster-demo-${range.value}
      clusterIdentifier: ${default.id}
      instanceClass: db.r5.large
    options: {}
  default:
    type: aws:docdb:Cluster
    properties:
      clusterIdentifier: docdb-cluster-demo
      availabilityZones:
        - us-west-2a
        - us-west-2b
        - us-west-2c
      masterUsername: foo
      masterPassword: barbut8chars
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DocumentDB Cluster Instances using the `identifier`. For example:

```sh
 $ pulumi import aws:docdb/clusterInstance:ClusterInstance prod_instance_1 aurora-cluster-instance-1
```
 