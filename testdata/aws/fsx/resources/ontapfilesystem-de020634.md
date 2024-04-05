Manages an Amazon FSx for NetApp ONTAP file system.
See the [FSx ONTAP User Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/what-is-fsx-ontap.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.fsx.OntapFileSystem("test", {
    storageCapacity: 1024,
    subnetIds: [
        aws_subnet.test1.id,
        aws_subnet.test2.id,
    ],
    deploymentType: "MULTI_AZ_1",
    throughputCapacity: 512,
    preferredSubnetId: aws_subnet.test1.id,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.fsx.OntapFileSystem("test",
    storage_capacity=1024,
    subnet_ids=[
        aws_subnet["test1"]["id"],
        aws_subnet["test2"]["id"],
    ],
    deployment_type="MULTI_AZ_1",
    throughput_capacity=512,
    preferred_subnet_id=aws_subnet["test1"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Fsx.OntapFileSystem("test", new()
    {
        StorageCapacity = 1024,
        SubnetIds = new[]
        {
            aws_subnet.Test1.Id,
            aws_subnet.Test2.Id,
        },
        DeploymentType = "MULTI_AZ_1",
        ThroughputCapacity = 512,
        PreferredSubnetId = aws_subnet.Test1.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.NewOntapFileSystem(ctx, "test", &fsx.OntapFileSystemArgs{
			StorageCapacity: pulumi.Int(1024),
			SubnetIds: pulumi.StringArray{
				aws_subnet.Test1.Id,
				aws_subnet.Test2.Id,
			},
			DeploymentType:     pulumi.String("MULTI_AZ_1"),
			ThroughputCapacity: pulumi.Int(512),
			PreferredSubnetId:  pulumi.Any(aws_subnet.Test1.Id),
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
import com.pulumi.aws.fsx.OntapFileSystem;
import com.pulumi.aws.fsx.OntapFileSystemArgs;
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
        var test = new OntapFileSystem("test", OntapFileSystemArgs.builder()        
            .storageCapacity(1024)
            .subnetIds(            
                aws_subnet.test1().id(),
                aws_subnet.test2().id())
            .deploymentType("MULTI_AZ_1")
            .throughputCapacity(512)
            .preferredSubnetId(aws_subnet.test1().id())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:fsx:OntapFileSystem
    properties:
      storageCapacity: 1024
      subnetIds:
        - ${aws_subnet.test1.id}
        - ${aws_subnet.test2.id}
      deploymentType: MULTI_AZ_1
      throughputCapacity: 512
      preferredSubnetId: ${aws_subnet.test1.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import FSx File Systems using the `id`. For example:

```sh
 $ pulumi import aws:fsx/ontapFileSystem:OntapFileSystem example fs-543ab12b1ca672f33
```
 Certain resource arguments, like `security_group_ids`, do not have a FSx API method for reading the information after creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

