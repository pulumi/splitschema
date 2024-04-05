Manages an AWS Storage Gateway stored iSCSI volume.

> **NOTE:** The gateway must have a working storage added (e.g., via the `aws.storagegateway.WorkingStorage` resource) before the volume is operational to clients, however the Storage Gateway API will allow volume creation without error in that case and return volume status as `WORKING STORAGE NOT CONFIGURED`.

{{% examples %}}
## Example Usage
{{% example %}}
### Create Empty Stored iSCSI Volume

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.StoredIscsiVolume("example", {
    gatewayArn: aws_storagegateway_cache.example.gateway_arn,
    networkInterfaceId: aws_instance.example.private_ip,
    targetName: "example",
    preserveExistingData: false,
    diskId: data.aws_storagegateway_local_disk.test.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.StoredIscsiVolume("example",
    gateway_arn=aws_storagegateway_cache["example"]["gateway_arn"],
    network_interface_id=aws_instance["example"]["private_ip"],
    target_name="example",
    preserve_existing_data=False,
    disk_id=data["aws_storagegateway_local_disk"]["test"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.StoredIscsiVolume("example", new()
    {
        GatewayArn = aws_storagegateway_cache.Example.Gateway_arn,
        NetworkInterfaceId = aws_instance.Example.Private_ip,
        TargetName = "example",
        PreserveExistingData = false,
        DiskId = data.Aws_storagegateway_local_disk.Test.Id,
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
		_, err := storagegateway.NewStoredIscsiVolume(ctx, "example", &storagegateway.StoredIscsiVolumeArgs{
			GatewayArn:           pulumi.Any(aws_storagegateway_cache.Example.Gateway_arn),
			NetworkInterfaceId:   pulumi.Any(aws_instance.Example.Private_ip),
			TargetName:           pulumi.String("example"),
			PreserveExistingData: pulumi.Bool(false),
			DiskId:               pulumi.Any(data.Aws_storagegateway_local_disk.Test.Id),
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
import com.pulumi.aws.storagegateway.StoredIscsiVolume;
import com.pulumi.aws.storagegateway.StoredIscsiVolumeArgs;
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
        var example = new StoredIscsiVolume("example", StoredIscsiVolumeArgs.builder()        
            .gatewayArn(aws_storagegateway_cache.example().gateway_arn())
            .networkInterfaceId(aws_instance.example().private_ip())
            .targetName("example")
            .preserveExistingData(false)
            .diskId(data.aws_storagegateway_local_disk().test().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:StoredIscsiVolume
    properties:
      gatewayArn: ${aws_storagegateway_cache.example.gateway_arn}
      networkInterfaceId: ${aws_instance.example.private_ip}
      targetName: example
      preserveExistingData: false
      diskId: ${data.aws_storagegateway_local_disk.test.id}
```
{{% /example %}}
{{% example %}}
### Create Stored iSCSI Volume From Snapshot

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.StoredIscsiVolume("example", {
    gatewayArn: aws_storagegateway_cache.example.gateway_arn,
    networkInterfaceId: aws_instance.example.private_ip,
    snapshotId: aws_ebs_snapshot.example.id,
    targetName: "example",
    preserveExistingData: false,
    diskId: data.aws_storagegateway_local_disk.test.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.StoredIscsiVolume("example",
    gateway_arn=aws_storagegateway_cache["example"]["gateway_arn"],
    network_interface_id=aws_instance["example"]["private_ip"],
    snapshot_id=aws_ebs_snapshot["example"]["id"],
    target_name="example",
    preserve_existing_data=False,
    disk_id=data["aws_storagegateway_local_disk"]["test"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.StoredIscsiVolume("example", new()
    {
        GatewayArn = aws_storagegateway_cache.Example.Gateway_arn,
        NetworkInterfaceId = aws_instance.Example.Private_ip,
        SnapshotId = aws_ebs_snapshot.Example.Id,
        TargetName = "example",
        PreserveExistingData = false,
        DiskId = data.Aws_storagegateway_local_disk.Test.Id,
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
		_, err := storagegateway.NewStoredIscsiVolume(ctx, "example", &storagegateway.StoredIscsiVolumeArgs{
			GatewayArn:           pulumi.Any(aws_storagegateway_cache.Example.Gateway_arn),
			NetworkInterfaceId:   pulumi.Any(aws_instance.Example.Private_ip),
			SnapshotId:           pulumi.Any(aws_ebs_snapshot.Example.Id),
			TargetName:           pulumi.String("example"),
			PreserveExistingData: pulumi.Bool(false),
			DiskId:               pulumi.Any(data.Aws_storagegateway_local_disk.Test.Id),
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
import com.pulumi.aws.storagegateway.StoredIscsiVolume;
import com.pulumi.aws.storagegateway.StoredIscsiVolumeArgs;
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
        var example = new StoredIscsiVolume("example", StoredIscsiVolumeArgs.builder()        
            .gatewayArn(aws_storagegateway_cache.example().gateway_arn())
            .networkInterfaceId(aws_instance.example().private_ip())
            .snapshotId(aws_ebs_snapshot.example().id())
            .targetName("example")
            .preserveExistingData(false)
            .diskId(data.aws_storagegateway_local_disk().test().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:StoredIscsiVolume
    properties:
      gatewayArn: ${aws_storagegateway_cache.example.gateway_arn}
      networkInterfaceId: ${aws_instance.example.private_ip}
      snapshotId: ${aws_ebs_snapshot.example.id}
      targetName: example
      preserveExistingData: false
      diskId: ${data.aws_storagegateway_local_disk.test.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_storagegateway_stored_iscsi_volume` using the volume Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:storagegateway/storedIscsiVolume:StoredIscsiVolume example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678
```
 