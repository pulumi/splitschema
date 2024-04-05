Manages an Amazon FSx for OpenZFS file system.
See the [FSx OpenZFS User Guide](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/what-is-fsx.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.fsx.OpenZfsFileSystem("test", {
    storageCapacity: 64,
    subnetIds: [aws_subnet.test1.id],
    deploymentType: "SINGLE_AZ_1",
    throughputCapacity: 64,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.fsx.OpenZfsFileSystem("test",
    storage_capacity=64,
    subnet_ids=[aws_subnet["test1"]["id"]],
    deployment_type="SINGLE_AZ_1",
    throughput_capacity=64)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Fsx.OpenZfsFileSystem("test", new()
    {
        StorageCapacity = 64,
        SubnetIds = new[]
        {
            aws_subnet.Test1.Id,
        },
        DeploymentType = "SINGLE_AZ_1",
        ThroughputCapacity = 64,
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
		_, err := fsx.NewOpenZfsFileSystem(ctx, "test", &fsx.OpenZfsFileSystemArgs{
			StorageCapacity: pulumi.Int(64),
			SubnetIds: pulumi.String{
				aws_subnet.Test1.Id,
			},
			DeploymentType:     pulumi.String("SINGLE_AZ_1"),
			ThroughputCapacity: pulumi.Int(64),
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
import com.pulumi.aws.fsx.OpenZfsFileSystem;
import com.pulumi.aws.fsx.OpenZfsFileSystemArgs;
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
        var test = new OpenZfsFileSystem("test", OpenZfsFileSystemArgs.builder()        
            .storageCapacity(64)
            .subnetIds(aws_subnet.test1().id())
            .deploymentType("SINGLE_AZ_1")
            .throughputCapacity(64)
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:fsx:OpenZfsFileSystem
    properties:
      storageCapacity: 64
      subnetIds:
        - ${aws_subnet.test1.id}
      deploymentType: SINGLE_AZ_1
      throughputCapacity: 64
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import FSx File Systems using the `id`. For example:

```sh
 $ pulumi import aws:fsx/openZfsFileSystem:OpenZfsFileSystem example fs-543ab12b1ca672f33
```
 Certain resource arguments, like `security_group_ids`, do not have a FSx API method for reading the information after creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

