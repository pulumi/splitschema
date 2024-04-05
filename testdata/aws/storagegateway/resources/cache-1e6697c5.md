Manages an AWS Storage Gateway cache.

> **NOTE:** The Storage Gateway API provides no method to remove a cache disk. Destroying this resource does not perform any Storage Gateway actions.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.Cache("example", {
    diskId: data.aws_storagegateway_local_disk.example.id,
    gatewayArn: aws_storagegateway_gateway.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.Cache("example",
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
    var example = new Aws.StorageGateway.Cache("example", new()
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
		_, err := storagegateway.NewCache(ctx, "example", &storagegateway.CacheArgs{
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
import com.pulumi.aws.storagegateway.Cache;
import com.pulumi.aws.storagegateway.CacheArgs;
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
        var example = new Cache("example", CacheArgs.builder()        
            .diskId(data.aws_storagegateway_local_disk().example().id())
            .gatewayArn(aws_storagegateway_gateway.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:Cache
    properties:
      diskId: ${data.aws_storagegateway_local_disk.example.id}
      gatewayArn: ${aws_storagegateway_gateway.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_storagegateway_cache` using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`). For example:

```sh
 $ pulumi import aws:storagegateway/cache:Cache example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
```
 