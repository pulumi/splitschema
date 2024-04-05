Provides an Elastic MapReduce Cluster Instance Group configuration.
See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/emr/) for more information.

> **NOTE:** At this time, Instance Groups cannot be destroyed through the API nor
web interface. Instance Groups are destroyed when the EMR Cluster is destroyed.
this provider will resize any Instance Group to zero when destroying the resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const task = new aws.emr.InstanceGroup("task", {
    clusterId: aws_emr_cluster["tf-test-cluster"].id,
    instanceCount: 1,
    instanceType: "m5.xlarge",
});
```
```python
import pulumi
import pulumi_aws as aws

task = aws.emr.InstanceGroup("task",
    cluster_id=aws_emr_cluster["tf-test-cluster"]["id"],
    instance_count=1,
    instance_type="m5.xlarge")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var task = new Aws.Emr.InstanceGroup("task", new()
    {
        ClusterId = aws_emr_cluster.Tf_test_cluster.Id,
        InstanceCount = 1,
        InstanceType = "m5.xlarge",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.NewInstanceGroup(ctx, "task", &emr.InstanceGroupArgs{
			ClusterId:     pulumi.Any(aws_emr_cluster.TfTestCluster.Id),
			InstanceCount: pulumi.Int(1),
			InstanceType:  pulumi.String("m5.xlarge"),
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
import com.pulumi.aws.emr.InstanceGroup;
import com.pulumi.aws.emr.InstanceGroupArgs;
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
        var task = new InstanceGroup("task", InstanceGroupArgs.builder()        
            .clusterId(aws_emr_cluster.tf-test-cluster().id())
            .instanceCount(1)
            .instanceType("m5.xlarge")
            .build());

    }
}
```
```yaml
resources:
  task:
    type: aws:emr:InstanceGroup
    properties:
      clusterId: ${aws_emr_cluster"tf-test-cluster"[%!s(MISSING)].id}
      instanceCount: 1
      instanceType: m5.xlarge
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EMR task instance group using their EMR Cluster id and Instance Group id separated by a forward-slash `/`. For example:

```sh
 $ pulumi import aws:emr/instanceGroup:InstanceGroup task_group j-123456ABCDEF/ig-15EK4O09RZLNR
```
 