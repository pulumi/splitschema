Manages a DocumentDB Cluster.

Changes to a DocumentDB Cluster can occur when you manually change a
parameter, such as `port`, and are reflected in the next maintenance
window. Because of this, this provider may report a difference in its planning
phase because a modification has not yet taken place. You can use the
`apply_immediately` flag to instruct the service to apply the change immediately
(see documentation below).

> **Note:** using `apply_immediately` can result in a brief downtime as the server reboots.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const docdb = new aws.docdb.Cluster("docdb", {
    backupRetentionPeriod: 5,
    clusterIdentifier: "my-docdb-cluster",
    engine: "docdb",
    masterPassword: "mustbeeightchars",
    masterUsername: "foo",
    preferredBackupWindow: "07:00-09:00",
    skipFinalSnapshot: true,
});
```
```python
import pulumi
import pulumi_aws as aws

docdb = aws.docdb.Cluster("docdb",
    backup_retention_period=5,
    cluster_identifier="my-docdb-cluster",
    engine="docdb",
    master_password="mustbeeightchars",
    master_username="foo",
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
    var docdb = new Aws.DocDB.Cluster("docdb", new()
    {
        BackupRetentionPeriod = 5,
        ClusterIdentifier = "my-docdb-cluster",
        Engine = "docdb",
        MasterPassword = "mustbeeightchars",
        MasterUsername = "foo",
        PreferredBackupWindow = "07:00-09:00",
        SkipFinalSnapshot = true,
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
		_, err := docdb.NewCluster(ctx, "docdb", &docdb.ClusterArgs{
			BackupRetentionPeriod: pulumi.Int(5),
			ClusterIdentifier:     pulumi.String("my-docdb-cluster"),
			Engine:                pulumi.String("docdb"),
			MasterPassword:        pulumi.String("mustbeeightchars"),
			MasterUsername:        pulumi.String("foo"),
			PreferredBackupWindow: pulumi.String("07:00-09:00"),
			SkipFinalSnapshot:     pulumi.Bool(true),
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
import com.pulumi.aws.docdb.ClusterArgs;
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
        var docdb = new Cluster("docdb", ClusterArgs.builder()        
            .backupRetentionPeriod(5)
            .clusterIdentifier("my-docdb-cluster")
            .engine("docdb")
            .masterPassword("mustbeeightchars")
            .masterUsername("foo")
            .preferredBackupWindow("07:00-09:00")
            .skipFinalSnapshot(true)
            .build());

    }
}
```
```yaml
resources:
  docdb:
    type: aws:docdb:Cluster
    properties:
      backupRetentionPeriod: 5
      clusterIdentifier: my-docdb-cluster
      engine: docdb
      masterPassword: mustbeeightchars
      masterUsername: foo
      preferredBackupWindow: 07:00-09:00
      skipFinalSnapshot: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DocumentDB Clusters using the `cluster_identifier`. For example:

```sh
 $ pulumi import aws:docdb/cluster:Cluster docdb_cluster docdb-prod-cluster
```
 