Provides a DAX Cluster resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.dax.Cluster("bar", {
    clusterName: "cluster-example",
    iamRoleArn: data.aws_iam_role.example.arn,
    nodeType: "dax.r4.large",
    replicationFactor: 1,
});
```
```python
import pulumi
import pulumi_aws as aws

bar = aws.dax.Cluster("bar",
    cluster_name="cluster-example",
    iam_role_arn=data["aws_iam_role"]["example"]["arn"],
    node_type="dax.r4.large",
    replication_factor=1)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bar = new Aws.Dax.Cluster("bar", new()
    {
        ClusterName = "cluster-example",
        IamRoleArn = data.Aws_iam_role.Example.Arn,
        NodeType = "dax.r4.large",
        ReplicationFactor = 1,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dax"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dax.NewCluster(ctx, "bar", &dax.ClusterArgs{
			ClusterName:       pulumi.String("cluster-example"),
			IamRoleArn:        pulumi.Any(data.Aws_iam_role.Example.Arn),
			NodeType:          pulumi.String("dax.r4.large"),
			ReplicationFactor: pulumi.Int(1),
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
import com.pulumi.aws.dax.Cluster;
import com.pulumi.aws.dax.ClusterArgs;
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
        var bar = new Cluster("bar", ClusterArgs.builder()        
            .clusterName("cluster-example")
            .iamRoleArn(data.aws_iam_role().example().arn())
            .nodeType("dax.r4.large")
            .replicationFactor(1)
            .build());

    }
}
```
```yaml
resources:
  bar:
    type: aws:dax:Cluster
    properties:
      clusterName: cluster-example
      iamRoleArn: ${data.aws_iam_role.example.arn}
      nodeType: dax.r4.large
      replicationFactor: 1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DAX Clusters using the `cluster_name`. For example:

```sh
 $ pulumi import aws:dax/cluster:Cluster my_cluster my_cluster
```
 