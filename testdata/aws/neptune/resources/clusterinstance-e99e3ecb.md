A Cluster Instance Resource defines attributes that are specific to a single instance in a Neptune Cluster.

You can simply add neptune instances and Neptune manages the replication. You can use the count
meta-parameter to make multiple instances and join them all to the same Neptune Cluster, or you may specify different Cluster Instance resources with various `instance_class` sizes.

{{% examples %}}
## Example Usage
{{% example %}}

The following example will create a neptune cluster with two neptune instances(one writer and one reader).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.neptune.Cluster("default", {
    clusterIdentifier: "neptune-cluster-demo",
    engine: "neptune",
    backupRetentionPeriod: 5,
    preferredBackupWindow: "07:00-09:00",
    skipFinalSnapshot: true,
    iamDatabaseAuthenticationEnabled: true,
    applyImmediately: true,
});
const example: aws.neptune.ClusterInstance[] = [];
for (const range = {value: 0}; range.value < 2; range.value++) {
    example.push(new aws.neptune.ClusterInstance(`example-${range.value}`, {
        clusterIdentifier: _default.id,
        engine: "neptune",
        instanceClass: "db.r4.large",
        applyImmediately: true,
    }));
}
```
```python
import pulumi
import pulumi_aws as aws

default = aws.neptune.Cluster("default",
    cluster_identifier="neptune-cluster-demo",
    engine="neptune",
    backup_retention_period=5,
    preferred_backup_window="07:00-09:00",
    skip_final_snapshot=True,
    iam_database_authentication_enabled=True,
    apply_immediately=True)
example = []
for range in [{"value": i} for i in range(0, 2)]:
    example.append(aws.neptune.ClusterInstance(f"example-{range['value']}",
        cluster_identifier=default.id,
        engine="neptune",
        instance_class="db.r4.large",
        apply_immediately=True))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Neptune.Cluster("default", new()
    {
        ClusterIdentifier = "neptune-cluster-demo",
        Engine = "neptune",
        BackupRetentionPeriod = 5,
        PreferredBackupWindow = "07:00-09:00",
        SkipFinalSnapshot = true,
        IamDatabaseAuthenticationEnabled = true,
        ApplyImmediately = true,
    });

    var example = new List<Aws.Neptune.ClusterInstance>();
    for (var rangeIndex = 0; rangeIndex < 2; rangeIndex++)
    {
        var range = new { Value = rangeIndex };
        example.Add(new Aws.Neptune.ClusterInstance($"example-{range.Value}", new()
        {
            ClusterIdentifier = @default.Id,
            Engine = "neptune",
            InstanceClass = "db.r4.large",
            ApplyImmediately = true,
        }));
    }
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
		_, err := neptune.NewCluster(ctx, "default", &neptune.ClusterArgs{
			ClusterIdentifier:                pulumi.String("neptune-cluster-demo"),
			Engine:                           pulumi.String("neptune"),
			BackupRetentionPeriod:            pulumi.Int(5),
			PreferredBackupWindow:            pulumi.String("07:00-09:00"),
			SkipFinalSnapshot:                pulumi.Bool(true),
			IamDatabaseAuthenticationEnabled: pulumi.Bool(true),
			ApplyImmediately:                 pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		var example []*neptune.ClusterInstance
		for index := 0; index < 2; index++ {
			key0 := index
			_ := index
			__res, err := neptune.NewClusterInstance(ctx, fmt.Sprintf("example-%v", key0), &neptune.ClusterInstanceArgs{
				ClusterIdentifier: _default.ID(),
				Engine:            pulumi.String("neptune"),
				InstanceClass:     pulumi.String("db.r4.large"),
				ApplyImmediately:  pulumi.Bool(true),
			})
			if err != nil {
				return err
			}
			example = append(example, __res)
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
import com.pulumi.aws.neptune.ClusterArgs;
import com.pulumi.aws.neptune.ClusterInstance;
import com.pulumi.aws.neptune.ClusterInstanceArgs;
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
            .clusterIdentifier("neptune-cluster-demo")
            .engine("neptune")
            .backupRetentionPeriod(5)
            .preferredBackupWindow("07:00-09:00")
            .skipFinalSnapshot(true)
            .iamDatabaseAuthenticationEnabled(true)
            .applyImmediately(true)
            .build());

        for (var i = 0; i < 2; i++) {
            new ClusterInstance("example-" + i, ClusterInstanceArgs.builder()            
                .clusterIdentifier(default_.id())
                .engine("neptune")
                .instanceClass("db.r4.large")
                .applyImmediately(true)
                .build());

        
}
    }
}
```
```yaml
resources:
  default:
    type: aws:neptune:Cluster
    properties:
      clusterIdentifier: neptune-cluster-demo
      engine: neptune
      backupRetentionPeriod: 5
      preferredBackupWindow: 07:00-09:00
      skipFinalSnapshot: true
      iamDatabaseAuthenticationEnabled: true
      applyImmediately: true
  example:
    type: aws:neptune:ClusterInstance
    properties:
      clusterIdentifier: ${default.id}
      engine: neptune
      instanceClass: db.r4.large
      applyImmediately: true
    options: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_neptune_cluster_instance` using the instance identifier. For example:

```sh
 $ pulumi import aws:neptune/clusterInstance:ClusterInstance example my-instance
```
 