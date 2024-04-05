Provides an Neptune Cluster Resource. A Cluster Resource defines attributes that are
applied to the entire cluster of Neptune Cluster Instances.

Changes to a Neptune Cluster can occur when you manually change a
parameter, such as `backup_retention_period`, and are reflected in the next maintenance
window. Because of this, this provider may report a difference in its planning
phase because a modification has not yet taken place. You can use the
`apply_immediately` flag to instruct the service to apply the change immediately
(see documentation below).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.neptune.Cluster("default", {
    applyImmediately: true,
    backupRetentionPeriod: 5,
    clusterIdentifier: "neptune-cluster-demo",
    engine: "neptune",
    iamDatabaseAuthenticationEnabled: true,
    preferredBackupWindow: "07:00-09:00",
    skipFinalSnapshot: true,
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.neptune.Cluster("default",
    apply_immediately=True,
    backup_retention_period=5,
    cluster_identifier="neptune-cluster-demo",
    engine="neptune",
    iam_database_authentication_enabled=True,
    preferred_backup_window="07:00-09:00",
    skip_final_snapshot=True)
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
        ApplyImmediately = true,
        BackupRetentionPeriod = 5,
        ClusterIdentifier = "neptune-cluster-demo",
        Engine = "neptune",
        IamDatabaseAuthenticationEnabled = true,
        PreferredBackupWindow = "07:00-09:00",
        SkipFinalSnapshot = true,
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
		_, err := neptune.NewCluster(ctx, "default", &neptune.ClusterArgs{
			ApplyImmediately:                 pulumi.Bool(true),
			BackupRetentionPeriod:            pulumi.Int(5),
			ClusterIdentifier:                pulumi.String("neptune-cluster-demo"),
			Engine:                           pulumi.String("neptune"),
			IamDatabaseAuthenticationEnabled: pulumi.Bool(true),
			PreferredBackupWindow:            pulumi.String("07:00-09:00"),
			SkipFinalSnapshot:                pulumi.Bool(true),
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
import com.pulumi.aws.neptune.ClusterArgs;
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
            .applyImmediately(true)
            .backupRetentionPeriod(5)
            .clusterIdentifier("neptune-cluster-demo")
            .engine("neptune")
            .iamDatabaseAuthenticationEnabled(true)
            .preferredBackupWindow("07:00-09:00")
            .skipFinalSnapshot(true)
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:neptune:Cluster
    properties:
      applyImmediately: true
      backupRetentionPeriod: 5
      clusterIdentifier: neptune-cluster-demo
      engine: neptune
      iamDatabaseAuthenticationEnabled: true
      preferredBackupWindow: 07:00-09:00
      skipFinalSnapshot: true
```

> **Note:** AWS Neptune does not support user name/password–based access control.
See the AWS [Docs](https://docs.aws.amazon.com/neptune/latest/userguide/limits.html) for more information.
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_neptune_cluster` using the cluster identifier. For example:

```sh
 $ pulumi import aws:neptune/cluster:Cluster example my-cluster
```
 