Manages an AWS Storage Gateway upload buffer.

> **NOTE:** The Storage Gateway API provides no method to remove an upload buffer disk. Destroying this resource does not perform any Storage Gateway actions.

{{% examples %}}
## Example Usage
{{% example %}}
### Cached and VTL Gateway Type

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testLocalDisk = aws.storagegateway.getLocalDisk({
    diskNode: aws_volume_attachment.test.device_name,
    gatewayArn: aws_storagegateway_gateway.test.arn,
});
const testUploadBuffer = new aws.storagegateway.UploadBuffer("testUploadBuffer", {
    diskPath: testLocalDisk.then(testLocalDisk => testLocalDisk.diskPath),
    gatewayArn: aws_storagegateway_gateway.test.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

test_local_disk = aws.storagegateway.get_local_disk(disk_node=aws_volume_attachment["test"]["device_name"],
    gateway_arn=aws_storagegateway_gateway["test"]["arn"])
test_upload_buffer = aws.storagegateway.UploadBuffer("testUploadBuffer",
    disk_path=test_local_disk.disk_path,
    gateway_arn=aws_storagegateway_gateway["test"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testLocalDisk = Aws.StorageGateway.GetLocalDisk.Invoke(new()
    {
        DiskNode = aws_volume_attachment.Test.Device_name,
        GatewayArn = aws_storagegateway_gateway.Test.Arn,
    });

    var testUploadBuffer = new Aws.StorageGateway.UploadBuffer("testUploadBuffer", new()
    {
        DiskPath = testLocalDisk.Apply(getLocalDiskResult => getLocalDiskResult.DiskPath),
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
		testLocalDisk, err := storagegateway.GetLocalDisk(ctx, &storagegateway.GetLocalDiskArgs{
			DiskNode:   pulumi.StringRef(aws_volume_attachment.Test.Device_name),
			GatewayArn: aws_storagegateway_gateway.Test.Arn,
		}, nil)
		if err != nil {
			return err
		}
		_, err = storagegateway.NewUploadBuffer(ctx, "testUploadBuffer", &storagegateway.UploadBufferArgs{
			DiskPath:   *pulumi.String(testLocalDisk.DiskPath),
			GatewayArn: pulumi.Any(aws_storagegateway_gateway.Test.Arn),
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
import com.pulumi.aws.storagegateway.StoragegatewayFunctions;
import com.pulumi.aws.storagegateway.inputs.GetLocalDiskArgs;
import com.pulumi.aws.storagegateway.UploadBuffer;
import com.pulumi.aws.storagegateway.UploadBufferArgs;
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
        final var testLocalDisk = StoragegatewayFunctions.getLocalDisk(GetLocalDiskArgs.builder()
            .diskNode(aws_volume_attachment.test().device_name())
            .gatewayArn(aws_storagegateway_gateway.test().arn())
            .build());

        var testUploadBuffer = new UploadBuffer("testUploadBuffer", UploadBufferArgs.builder()        
            .diskPath(testLocalDisk.applyValue(getLocalDiskResult -> getLocalDiskResult.diskPath()))
            .gatewayArn(aws_storagegateway_gateway.test().arn())
            .build());

    }
}
```
```yaml
resources:
  testUploadBuffer:
    type: aws:storagegateway:UploadBuffer
    properties:
      diskPath: ${testLocalDisk.diskPath}
      gatewayArn: ${aws_storagegateway_gateway.test.arn}
variables:
  testLocalDisk:
    fn::invoke:
      Function: aws:storagegateway:getLocalDisk
      Arguments:
        diskNode: ${aws_volume_attachment.test.device_name}
        gatewayArn: ${aws_storagegateway_gateway.test.arn}
```
{{% /example %}}
{{% example %}}
### Stored Gateway Type

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.storagegateway.getLocalDisk({
    diskNode: aws_volume_attachment.test.device_name,
    gatewayArn: aws_storagegateway_gateway.test.arn,
});
const example = new aws.storagegateway.UploadBuffer("example", {
    diskId: data.aws_storagegateway_local_disk.example.id,
    gatewayArn: aws_storagegateway_gateway.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.storagegateway.get_local_disk(disk_node=aws_volume_attachment["test"]["device_name"],
    gateway_arn=aws_storagegateway_gateway["test"]["arn"])
example = aws.storagegateway.UploadBuffer("example",
    disk_id=data["aws_storagegateway_local_disk"]["example"]["id"],
    gateway_arn=aws_storagegateway_gateway["example"]["arn"])
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
        DiskNode = aws_volume_attachment.Test.Device_name,
        GatewayArn = aws_storagegateway_gateway.Test.Arn,
    });

    var example = new Aws.StorageGateway.UploadBuffer("example", new()
    {
        DiskId = data.Aws_storagegateway_local_disk.Example.Id,
        GatewayArn = aws_storagegateway_gateway.Example.Arn,
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
			DiskNode:   pulumi.StringRef(aws_volume_attachment.Test.Device_name),
			GatewayArn: aws_storagegateway_gateway.Test.Arn,
		}, nil)
		if err != nil {
			return err
		}
		_, err = storagegateway.NewUploadBuffer(ctx, "example", &storagegateway.UploadBufferArgs{
			DiskId:     pulumi.Any(data.Aws_storagegateway_local_disk.Example.Id),
			GatewayArn: pulumi.Any(aws_storagegateway_gateway.Example.Arn),
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
import com.pulumi.aws.storagegateway.StoragegatewayFunctions;
import com.pulumi.aws.storagegateway.inputs.GetLocalDiskArgs;
import com.pulumi.aws.storagegateway.UploadBuffer;
import com.pulumi.aws.storagegateway.UploadBufferArgs;
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
            .diskNode(aws_volume_attachment.test().device_name())
            .gatewayArn(aws_storagegateway_gateway.test().arn())
            .build());

        var example = new UploadBuffer("example", UploadBufferArgs.builder()        
            .diskId(data.aws_storagegateway_local_disk().example().id())
            .gatewayArn(aws_storagegateway_gateway.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:UploadBuffer
    properties:
      diskId: ${data.aws_storagegateway_local_disk.example.id}
      gatewayArn: ${aws_storagegateway_gateway.example.arn}
variables:
  test:
    fn::invoke:
      Function: aws:storagegateway:getLocalDisk
      Arguments:
        diskNode: ${aws_volume_attachment.test.device_name}
        gatewayArn: ${aws_storagegateway_gateway.test.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_storagegateway_upload_buffer` using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`). For example:

```sh
 $ pulumi import aws:storagegateway/uploadBuffer:UploadBuffer example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
```
 