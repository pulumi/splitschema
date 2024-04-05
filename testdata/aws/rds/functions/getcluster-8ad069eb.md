Provides information about an RDS cluster.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const clusterName = aws.rds.getCluster({
    clusterIdentifier: "clusterName",
});
```
```python
import pulumi
import pulumi_aws as aws

cluster_name = aws.rds.get_cluster(cluster_identifier="clusterName")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var clusterName = Aws.Rds.GetCluster.Invoke(new()
    {
        ClusterIdentifier = "clusterName",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.LookupCluster(ctx, &rds.LookupClusterArgs{
			ClusterIdentifier: "clusterName",
		}, nil)
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
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetClusterArgs;
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
        final var clusterName = RdsFunctions.getCluster(GetClusterArgs.builder()
            .clusterIdentifier("clusterName")
            .build());

    }
}
```
```yaml
variables:
  clusterName:
    fn::invoke:
      Function: aws:rds:getCluster
      Arguments:
        clusterIdentifier: clusterName
```
{{% /example %}}
{{% /examples %}}