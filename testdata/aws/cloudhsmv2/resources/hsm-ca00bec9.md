Creates an HSM module in Amazon CloudHSM v2 cluster.

{{% examples %}}
## Example Usage
{{% example %}}

The following example below creates an HSM module in CloudHSM cluster.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cluster = aws.cloudhsmv2.getCluster({
    clusterId: _var.cloudhsm_cluster_id,
});
const cloudhsmV2Hsm = new aws.cloudhsmv2.Hsm("cloudhsmV2Hsm", {
    subnetId: cluster.then(cluster => cluster.subnetIds?.[0]),
    clusterId: cluster.then(cluster => cluster.clusterId),
});
```
```python
import pulumi
import pulumi_aws as aws

cluster = aws.cloudhsmv2.get_cluster(cluster_id=var["cloudhsm_cluster_id"])
cloudhsm_v2_hsm = aws.cloudhsmv2.Hsm("cloudhsmV2Hsm",
    subnet_id=cluster.subnet_ids[0],
    cluster_id=cluster.cluster_id)
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
        ClusterId = @var.Cloudhsm_cluster_id,
    });

    var cloudhsmV2Hsm = new Aws.CloudHsmV2.Hsm("cloudhsmV2Hsm", new()
    {
        SubnetId = cluster.Apply(getClusterResult => getClusterResult.SubnetIds[0]),
        ClusterId = cluster.Apply(getClusterResult => getClusterResult.ClusterId),
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
		cluster, err := cloudhsmv2.LookupCluster(ctx, &cloudhsmv2.LookupClusterArgs{
			ClusterId: _var.Cloudhsm_cluster_id,
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudhsmv2.NewHsm(ctx, "cloudhsmV2Hsm", &cloudhsmv2.HsmArgs{
			SubnetId:  *pulumi.String(cluster.SubnetIds[0]),
			ClusterId: *pulumi.String(cluster.ClusterId),
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
import com.pulumi.aws.cloudhsmv2.Cloudhsmv2Functions;
import com.pulumi.aws.cloudhsmv2.inputs.GetClusterArgs;
import com.pulumi.aws.cloudhsmv2.Hsm;
import com.pulumi.aws.cloudhsmv2.HsmArgs;
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
            .clusterId(var_.cloudhsm_cluster_id())
            .build());

        var cloudhsmV2Hsm = new Hsm("cloudhsmV2Hsm", HsmArgs.builder()        
            .subnetId(cluster.applyValue(getClusterResult -> getClusterResult.subnetIds()[0]))
            .clusterId(cluster.applyValue(getClusterResult -> getClusterResult.clusterId()))
            .build());

    }
}
```
```yaml
resources:
  cloudhsmV2Hsm:
    type: aws:cloudhsmv2:Hsm
    properties:
      subnetId: ${cluster.subnetIds[0]}
      clusterId: ${cluster.clusterId}
variables:
  cluster:
    fn::invoke:
      Function: aws:cloudhsmv2:getCluster
      Arguments:
        clusterId: ${var.cloudhsm_cluster_id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import HSM modules using their HSM ID. For example:

```sh
 $ pulumi import aws:cloudhsmv2/hsm:Hsm bar hsm-quo8dahtaca
```
 