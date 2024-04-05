Information about Redshift Orderable Clusters and valid parameter combinations.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.redshift.getOrderableCluster({
    clusterType: "multi-node",
    preferredNodeTypes: [
        "dc2.large",
        "ds2.xlarge",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.redshift.get_orderable_cluster(cluster_type="multi-node",
    preferred_node_types=[
        "dc2.large",
        "ds2.xlarge",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.RedShift.GetOrderableCluster.Invoke(new()
    {
        ClusterType = "multi-node",
        PreferredNodeTypes = new[]
        {
            "dc2.large",
            "ds2.xlarge",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.GetOrderableCluster(ctx, &redshift.GetOrderableClusterArgs{
			ClusterType: pulumi.StringRef("multi-node"),
			PreferredNodeTypes: []string{
				"dc2.large",
				"ds2.xlarge",
			},
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
import com.pulumi.aws.redshift.RedshiftFunctions;
import com.pulumi.aws.redshift.inputs.GetOrderableClusterArgs;
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
        final var test = RedshiftFunctions.getOrderableCluster(GetOrderableClusterArgs.builder()
            .clusterType("multi-node")
            .preferredNodeTypes(            
                "dc2.large",
                "ds2.xlarge")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:redshift:getOrderableCluster
      Arguments:
        clusterType: multi-node
        preferredNodeTypes:
          - dc2.large
          - ds2.xlarge
```
{{% /example %}}
{{% /examples %}}