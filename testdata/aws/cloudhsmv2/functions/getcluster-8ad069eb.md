Use this data source to get information about a CloudHSM v2 cluster

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cluster = aws.cloudhsmv2.getCluster({
    clusterId: "cluster-testclusterid",
});
```
```python
import pulumi
import pulumi_aws as aws

cluster = aws.cloudhsmv2.get_cluster(cluster_id="cluster-testclusterid")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var cluster = Aws.CloudHsmV2.GetCluster.Invoke(new()
    {
        ClusterId = "cluster-testclusterid",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudhsmv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudhsmv2.LookupCluster(ctx, &cloudhsmv2.LookupClusterArgs{
			ClusterId: "cluster-testclusterid",
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
import com.pulumi.aws.cloudhsmv2.Cloudhsmv2Functions;
import com.pulumi.aws.cloudhsmv2.inputs.GetClusterArgs;
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
        final var cluster = Cloudhsmv2Functions.getCluster(GetClusterArgs.builder()
            .clusterId("cluster-testclusterid")
            .build());

    }
}
```
```yaml
variables:
  cluster:
    fn::invoke:
      Function: aws:cloudhsmv2:getCluster
      Arguments:
        clusterId: cluster-testclusterid
```
{{% /example %}}
{{% /examples %}}