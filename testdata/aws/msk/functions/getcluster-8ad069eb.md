Get information on an Amazon MSK Cluster.

> **Note:** This data sources returns information on _provisioned_ clusters.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.msk.getCluster({
    clusterName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.msk.get_cluster(cluster_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Msk.GetCluster.Invoke(new()
    {
        ClusterName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/msk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := msk.LookupCluster(ctx, &msk.LookupClusterArgs{
			ClusterName: "example",
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
import com.pulumi.aws.msk.MskFunctions;
import com.pulumi.aws.msk.inputs.GetClusterArgs;
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
        final var example = MskFunctions.getCluster(GetClusterArgs.builder()
            .clusterName("example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:msk:getCluster
      Arguments:
        clusterName: example
```
{{% /example %}}
{{% /examples %}}