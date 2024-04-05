Provides a MemoryDB Cluster.

More information about MemoryDB can be found in the [Developer Guide](https://docs.aws.amazon.com/memorydb/latest/devguide/what-is-memorydb-for-redis.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.memorydb.Cluster("example", {
    aclName: "open-access",
    nodeType: "db.t4g.small",
    numShards: 2,
    securityGroupIds: [aws_security_group.example.id],
    snapshotRetentionLimit: 7,
    subnetGroupName: aws_memorydb_subnet_group.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.memorydb.Cluster("example",
    acl_name="open-access",
    node_type="db.t4g.small",
    num_shards=2,
    security_group_ids=[aws_security_group["example"]["id"]],
    snapshot_retention_limit=7,
    subnet_group_name=aws_memorydb_subnet_group["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.MemoryDb.Cluster("example", new()
    {
        AclName = "open-access",
        NodeType = "db.t4g.small",
        NumShards = 2,
        SecurityGroupIds = new[]
        {
            aws_security_group.Example.Id,
        },
        SnapshotRetentionLimit = 7,
        SubnetGroupName = aws_memorydb_subnet_group.Example.Id,
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
		_, err := memorydb.NewCluster(ctx, "example", &memorydb.ClusterArgs{
			AclName:   pulumi.String("open-access"),
			NodeType:  pulumi.String("db.t4g.small"),
			NumShards: pulumi.Int(2),
			SecurityGroupIds: pulumi.StringArray{
				aws_security_group.Example.Id,
			},
			SnapshotRetentionLimit: pulumi.Int(7),
			SubnetGroupName:        pulumi.Any(aws_memorydb_subnet_group.Example.Id),
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
import com.pulumi.aws.memorydb.Cluster;
import com.pulumi.aws.memorydb.ClusterArgs;
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
        var example = new Cluster("example", ClusterArgs.builder()        
            .aclName("open-access")
            .nodeType("db.t4g.small")
            .numShards(2)
            .securityGroupIds(aws_security_group.example().id())
            .snapshotRetentionLimit(7)
            .subnetGroupName(aws_memorydb_subnet_group.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:memorydb:Cluster
    properties:
      aclName: open-access
      nodeType: db.t4g.small
      numShards: 2
      securityGroupIds:
        - ${aws_security_group.example.id}
      snapshotRetentionLimit: 7
      subnetGroupName: ${aws_memorydb_subnet_group.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a cluster using the `name`. For example:

```sh
 $ pulumi import aws:memorydb/cluster:Cluster example my-cluster
```
 