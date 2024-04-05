Provides a MemoryDB Snapshot.

More information about snapshot and restore can be found in the [MemoryDB User Guide](https://docs.aws.amazon.com/memorydb/latest/devguide/snapshots.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.memorydb.Snapshot("example", {clusterName: aws_memorydb_cluster.example.name});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.memorydb.Snapshot("example", cluster_name=aws_memorydb_cluster["example"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.MemoryDb.Snapshot("example", new()
    {
        ClusterName = aws_memorydb_cluster.Example.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/memorydb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := memorydb.NewSnapshot(ctx, "example", &memorydb.SnapshotArgs{
			ClusterName: pulumi.Any(aws_memorydb_cluster.Example.Name),
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
import com.pulumi.aws.memorydb.Snapshot;
import com.pulumi.aws.memorydb.SnapshotArgs;
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
        var example = new Snapshot("example", SnapshotArgs.builder()        
            .clusterName(aws_memorydb_cluster.example().name())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:memorydb:Snapshot
    properties:
      clusterName: ${aws_memorydb_cluster.example.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a snapshot using the `name`. For example:

```sh
 $ pulumi import aws:memorydb/snapshot:Snapshot example my-snapshot
```
 