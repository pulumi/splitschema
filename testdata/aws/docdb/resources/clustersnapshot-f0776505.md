Manages a DocumentDB database cluster snapshot for DocumentDB clusters.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.docdb.ClusterSnapshot("example", {
    dbClusterIdentifier: aws_docdb_cluster.example.id,
    dbClusterSnapshotIdentifier: "resourcetestsnapshot1234",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.docdb.ClusterSnapshot("example",
    db_cluster_identifier=aws_docdb_cluster["example"]["id"],
    db_cluster_snapshot_identifier="resourcetestsnapshot1234")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DocDB.ClusterSnapshot("example", new()
    {
        DbClusterIdentifier = aws_docdb_cluster.Example.Id,
        DbClusterSnapshotIdentifier = "resourcetestsnapshot1234",
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
		_, err := docdb.NewClusterSnapshot(ctx, "example", &docdb.ClusterSnapshotArgs{
			DbClusterIdentifier:         pulumi.Any(aws_docdb_cluster.Example.Id),
			DbClusterSnapshotIdentifier: pulumi.String("resourcetestsnapshot1234"),
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
import com.pulumi.aws.docdb.ClusterSnapshot;
import com.pulumi.aws.docdb.ClusterSnapshotArgs;
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
        var example = new ClusterSnapshot("example", ClusterSnapshotArgs.builder()        
            .dbClusterIdentifier(aws_docdb_cluster.example().id())
            .dbClusterSnapshotIdentifier("resourcetestsnapshot1234")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:docdb:ClusterSnapshot
    properties:
      dbClusterIdentifier: ${aws_docdb_cluster.example.id}
      dbClusterSnapshotIdentifier: resourcetestsnapshot1234
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_docdb_cluster_snapshot` using the cluster snapshot identifier. For example:

```sh
 $ pulumi import aws:docdb/clusterSnapshot:ClusterSnapshot example my-cluster-snapshot
```
 