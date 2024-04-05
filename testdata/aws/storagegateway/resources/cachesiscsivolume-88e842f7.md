Manages an AWS Storage Gateway cached iSCSI volume.

> **NOTE:** The gateway must have cache added (e.g., via the `aws.storagegateway.Cache` resource) before creating volumes otherwise the Storage Gateway API will return an error.

> **NOTE:** The gateway must have an upload buffer added (e.g., via the `aws.storagegateway.UploadBuffer` resource) before the volume is operational to clients, however the Storage Gateway API will allow volume creation without error in that case and return volume status as `UPLOAD BUFFER NOT CONFIGURED`.

{{% examples %}}
## Example Usage

> **NOTE:** These examples are referencing the `aws.storagegateway.Cache` resource `gateway_arn` attribute to ensure this provider properly adds cache before creating the volume. If you are not using this method, you may need to declare an expicit dependency (e.g. via `depends_on = [aws_storagegateway_cache.example]`) to ensure proper ordering.
{{% example %}}
### Create Empty Cached iSCSI Volume

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.CachesIscsiVolume("example", {
    gatewayArn: aws_storagegateway_cache.example.gateway_arn,
    networkInterfaceId: aws_instance.example.private_ip,
    targetName: "example",
    volumeSizeInBytes: 5368709120,
});
// 5 GB
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.CachesIscsiVolume("example",
    gateway_arn=aws_storagegateway_cache["example"]["gateway_arn"],
    network_interface_id=aws_instance["example"]["private_ip"],
    target_name="example",
    volume_size_in_bytes=5368709120)
# 5 GB
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.CachesIscsiVolume("example", new()
    {
        GatewayArn = aws_storagegateway_cache.Example.Gateway_arn,
        NetworkInterfaceId = aws_instance.Example.Private_ip,
        TargetName = "example",
        VolumeSizeInBytes = 5368709120,
    });

    // 5 GB
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
		_, err := storagegateway.NewCachesIscsiVolume(ctx, "example", &storagegateway.CachesIscsiVolumeArgs{
			GatewayArn:         pulumi.Any(aws_storagegateway_cache.Example.Gateway_arn),
			NetworkInterfaceId: pulumi.Any(aws_instance.Example.Private_ip),
			TargetName:         pulumi.String("example"),
			VolumeSizeInBytes:  pulumi.Int(5368709120),
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
import com.pulumi.aws.storagegateway.CachesIscsiVolume;
import com.pulumi.aws.storagegateway.CachesIscsiVolumeArgs;
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
        var example = new CachesIscsiVolume("example", CachesIscsiVolumeArgs.builder()        
            .gatewayArn(aws_storagegateway_cache.example().gateway_arn())
            .networkInterfaceId(aws_instance.example().private_ip())
            .targetName("example")
            .volumeSizeInBytes(5368709120)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:CachesIscsiVolume
    properties:
      gatewayArn: ${aws_storagegateway_cache.example.gateway_arn}
      networkInterfaceId: ${aws_instance.example.private_ip}
      targetName: example
      volumeSizeInBytes: 5.36870912e+09
```
{{% /example %}}
{{% example %}}
### Create Cached iSCSI Volume From Snapshot

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.CachesIscsiVolume("example", {
    gatewayArn: aws_storagegateway_cache.example.gateway_arn,
    networkInterfaceId: aws_instance.example.private_ip,
    snapshotId: aws_ebs_snapshot.example.id,
    targetName: "example",
    volumeSizeInBytes: aws_ebs_snapshot.example.volume_size * 1024 * 1024 * 1024,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.CachesIscsiVolume("example",
    gateway_arn=aws_storagegateway_cache["example"]["gateway_arn"],
    network_interface_id=aws_instance["example"]["private_ip"],
    snapshot_id=aws_ebs_snapshot["example"]["id"],
    target_name="example",
    volume_size_in_bytes=aws_ebs_snapshot["example"]["volume_size"] * 1024 * 1024 * 1024)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.CachesIscsiVolume("example", new()
    {
        GatewayArn = aws_storagegateway_cache.Example.Gateway_arn,
        NetworkInterfaceId = aws_instance.Example.Private_ip,
        SnapshotId = aws_ebs_snapshot.Example.Id,
        TargetName = "example",
        VolumeSizeInBytes = aws_ebs_snapshot.Example.Volume_size * 1024 * 1024 * 1024,
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
		_, err := storagegateway.NewCachesIscsiVolume(ctx, "example", &storagegateway.CachesIscsiVolumeArgs{
			GatewayArn:         pulumi.Any(aws_storagegateway_cache.Example.Gateway_arn),
			NetworkInterfaceId: pulumi.Any(aws_instance.Example.Private_ip),
			SnapshotId:         pulumi.Any(aws_ebs_snapshot.Example.Id),
			TargetName:         pulumi.String("example"),
			VolumeSizeInBytes:  aws_ebs_snapshot.Example.Volume_size * 1024 * 1024 * 1024,
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
import com.pulumi.aws.storagegateway.CachesIscsiVolume;
import com.pulumi.aws.storagegateway.CachesIscsiVolumeArgs;
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
        var example = new CachesIscsiVolume("example", CachesIscsiVolumeArgs.builder()        
            .gatewayArn(aws_storagegateway_cache.example().gateway_arn())
            .networkInterfaceId(aws_instance.example().private_ip())
            .snapshotId(aws_ebs_snapshot.example().id())
            .targetName("example")
            .volumeSizeInBytes(aws_ebs_snapshot.example().volume_size() * 1024 * 1024 * 1024)
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Create Cached iSCSI Volume From Source Volume

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.CachesIscsiVolume("example", {
    gatewayArn: aws_storagegateway_cache.example.gateway_arn,
    networkInterfaceId: aws_instance.example.private_ip,
    sourceVolumeArn: aws_storagegateway_cached_iscsi_volume.existing.arn,
    targetName: "example",
    volumeSizeInBytes: aws_storagegateway_cached_iscsi_volume.existing.volume_size_in_bytes,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.CachesIscsiVolume("example",
    gateway_arn=aws_storagegateway_cache["example"]["gateway_arn"],
    network_interface_id=aws_instance["example"]["private_ip"],
    source_volume_arn=aws_storagegateway_cached_iscsi_volume["existing"]["arn"],
    target_name="example",
    volume_size_in_bytes=aws_storagegateway_cached_iscsi_volume["existing"]["volume_size_in_bytes"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.CachesIscsiVolume("example", new()
    {
        GatewayArn = aws_storagegateway_cache.Example.Gateway_arn,
        NetworkInterfaceId = aws_instance.Example.Private_ip,
        SourceVolumeArn = aws_storagegateway_cached_iscsi_volume.Existing.Arn,
        TargetName = "example",
        VolumeSizeInBytes = aws_storagegateway_cached_iscsi_volume.Existing.Volume_size_in_bytes,
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
		_, err := storagegateway.NewCachesIscsiVolume(ctx, "example", &storagegateway.CachesIscsiVolumeArgs{
			GatewayArn:         pulumi.Any(aws_storagegateway_cache.Example.Gateway_arn),
			NetworkInterfaceId: pulumi.Any(aws_instance.Example.Private_ip),
			SourceVolumeArn:    pulumi.Any(aws_storagegateway_cached_iscsi_volume.Existing.Arn),
			TargetName:         pulumi.String("example"),
			VolumeSizeInBytes:  pulumi.Any(aws_storagegateway_cached_iscsi_volume.Existing.Volume_size_in_bytes),
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
import com.pulumi.aws.storagegateway.CachesIscsiVolume;
import com.pulumi.aws.storagegateway.CachesIscsiVolumeArgs;
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
        var example = new CachesIscsiVolume("example", CachesIscsiVolumeArgs.builder()        
            .gatewayArn(aws_storagegateway_cache.example().gateway_arn())
            .networkInterfaceId(aws_instance.example().private_ip())
            .sourceVolumeArn(aws_storagegateway_cached_iscsi_volume.existing().arn())
            .targetName("example")
            .volumeSizeInBytes(aws_storagegateway_cached_iscsi_volume.existing().volume_size_in_bytes())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:CachesIscsiVolume
    properties:
      gatewayArn: ${aws_storagegateway_cache.example.gateway_arn}
      networkInterfaceId: ${aws_instance.example.private_ip}
      sourceVolumeArn: ${aws_storagegateway_cached_iscsi_volume.existing.arn}
      targetName: example
      volumeSizeInBytes: ${aws_storagegateway_cached_iscsi_volume.existing.volume_size_in_bytes}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_storagegateway_cached_iscsi_volume` using the volume Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678
```
 