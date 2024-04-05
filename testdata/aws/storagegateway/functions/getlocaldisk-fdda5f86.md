Retrieve information about a Storage Gateway local disk. The disk identifier is useful for adding the disk as a cache or upload buffer to a gateway.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.storagegateway.getLocalDisk({
    diskPath: aws_volume_attachment.test.device_name,
    gatewayArn: aws_storagegateway_gateway.test.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.storagegateway.get_local_disk(disk_path=aws_volume_attachment["test"]["device_name"],
    gateway_arn=aws_storagegateway_gateway["test"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.StorageGateway.GetLocalDisk.Invoke(new()
    {
        DiskPath = aws_volume_attachment.Test.Device_name,
        GatewayArn = aws_storagegateway_gateway.Test.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagegateway.GetLocalDisk(ctx, &storagegateway.GetLocalDiskArgs{
			DiskPath:   pulumi.StringRef(aws_volume_attachment.Test.Device_name),
			GatewayArn: aws_storagegateway_gateway.Test.Arn,
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
import com.pulumi.aws.storagegateway.StoragegatewayFunctions;
import com.pulumi.aws.storagegateway.inputs.GetLocalDiskArgs;
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
        final var test = StoragegatewayFunctions.getLocalDisk(GetLocalDiskArgs.builder()
            .diskPath(aws_volume_attachment.test().device_name())
            .gatewayArn(aws_storagegateway_gateway.test().arn())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:storagegateway:getLocalDisk
      Arguments:
        diskPath: ${aws_volume_attachment.test.device_name}
        gatewayArn: ${aws_storagegateway_gateway.test.arn}
```
{{% /example %}}
{{% /examples %}}