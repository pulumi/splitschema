Manages an Amazon FSx for OpenZFS volume.
See the [FSx OpenZFS User Guide](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/what-is-fsx.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}
### Root volume Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOpenZfsFileSystem = new aws.fsx.OpenZfsFileSystem("exampleOpenZfsFileSystem", {
    storageCapacity: 64,
    subnetIds: [aws_subnet.example.id],
    deploymentType: "SINGLE_AZ_1",
    throughputCapacity: 64,
});
const exampleOpenZfsSnapshot = new aws.fsx.OpenZfsSnapshot("exampleOpenZfsSnapshot", {volumeId: exampleOpenZfsFileSystem.rootVolumeId});
```
```python
import pulumi
import pulumi_aws as aws

example_open_zfs_file_system = aws.fsx.OpenZfsFileSystem("exampleOpenZfsFileSystem",
    storage_capacity=64,
    subnet_ids=[aws_subnet["example"]["id"]],
    deployment_type="SINGLE_AZ_1",
    throughput_capacity=64)
example_open_zfs_snapshot = aws.fsx.OpenZfsSnapshot("exampleOpenZfsSnapshot", volume_id=example_open_zfs_file_system.root_volume_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleOpenZfsFileSystem = new Aws.Fsx.OpenZfsFileSystem("exampleOpenZfsFileSystem", new()
    {
        StorageCapacity = 64,
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        DeploymentType = "SINGLE_AZ_1",
        ThroughputCapacity = 64,
    });

    var exampleOpenZfsSnapshot = new Aws.Fsx.OpenZfsSnapshot("exampleOpenZfsSnapshot", new()
    {
        VolumeId = exampleOpenZfsFileSystem.RootVolumeId,
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
		exampleOpenZfsFileSystem, err := fsx.NewOpenZfsFileSystem(ctx, "exampleOpenZfsFileSystem", &fsx.OpenZfsFileSystemArgs{
			StorageCapacity: pulumi.Int(64),
			SubnetIds: pulumi.String{
				aws_subnet.Example.Id,
			},
			DeploymentType:     pulumi.String("SINGLE_AZ_1"),
			ThroughputCapacity: pulumi.Int(64),
		})
		if err != nil {
			return err
		}
		_, err = fsx.NewOpenZfsSnapshot(ctx, "exampleOpenZfsSnapshot", &fsx.OpenZfsSnapshotArgs{
			VolumeId: exampleOpenZfsFileSystem.RootVolumeId,
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
import com.pulumi.aws.fsx.OpenZfsSnapshot;
import com.pulumi.aws.fsx.OpenZfsSnapshotArgs;
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
        var exampleOpenZfsFileSystem = new OpenZfsFileSystem("exampleOpenZfsFileSystem", OpenZfsFileSystemArgs.builder()        
            .storageCapacity(64)
            .subnetIds(aws_subnet.example().id())
            .deploymentType("SINGLE_AZ_1")
            .throughputCapacity(64)
            .build());

        var exampleOpenZfsSnapshot = new OpenZfsSnapshot("exampleOpenZfsSnapshot", OpenZfsSnapshotArgs.builder()        
            .volumeId(exampleOpenZfsFileSystem.rootVolumeId())
            .build());

    }
}
```
```yaml
resources:
  exampleOpenZfsSnapshot:
    type: aws:fsx:OpenZfsSnapshot
    properties:
      volumeId: ${exampleOpenZfsFileSystem.rootVolumeId}
  exampleOpenZfsFileSystem:
    type: aws:fsx:OpenZfsFileSystem
    properties:
      storageCapacity: 64
      subnetIds:
        - ${aws_subnet.example.id}
      deploymentType: SINGLE_AZ_1
      throughputCapacity: 64
```
{{% /example %}}
{{% example %}}
### Child volume Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOpenZfsFileSystem = new aws.fsx.OpenZfsFileSystem("exampleOpenZfsFileSystem", {
    storageCapacity: 64,
    subnetIds: [aws_subnet.example.id],
    deploymentType: "SINGLE_AZ_1",
    throughputCapacity: 64,
});
const exampleOpenZfsVolume = new aws.fsx.OpenZfsVolume("exampleOpenZfsVolume", {parentVolumeId: exampleOpenZfsFileSystem.rootVolumeId});
const exampleOpenZfsSnapshot = new aws.fsx.OpenZfsSnapshot("exampleOpenZfsSnapshot", {volumeId: exampleOpenZfsVolume.id});
```
```python
import pulumi
import pulumi_aws as aws

example_open_zfs_file_system = aws.fsx.OpenZfsFileSystem("exampleOpenZfsFileSystem",
    storage_capacity=64,
    subnet_ids=[aws_subnet["example"]["id"]],
    deployment_type="SINGLE_AZ_1",
    throughput_capacity=64)
example_open_zfs_volume = aws.fsx.OpenZfsVolume("exampleOpenZfsVolume", parent_volume_id=example_open_zfs_file_system.root_volume_id)
example_open_zfs_snapshot = aws.fsx.OpenZfsSnapshot("exampleOpenZfsSnapshot", volume_id=example_open_zfs_volume.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleOpenZfsFileSystem = new Aws.Fsx.OpenZfsFileSystem("exampleOpenZfsFileSystem", new()
    {
        StorageCapacity = 64,
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        DeploymentType = "SINGLE_AZ_1",
        ThroughputCapacity = 64,
    });

    var exampleOpenZfsVolume = new Aws.Fsx.OpenZfsVolume("exampleOpenZfsVolume", new()
    {
        ParentVolumeId = exampleOpenZfsFileSystem.RootVolumeId,
    });

    var exampleOpenZfsSnapshot = new Aws.Fsx.OpenZfsSnapshot("exampleOpenZfsSnapshot", new()
    {
        VolumeId = exampleOpenZfsVolume.Id,
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
		exampleOpenZfsFileSystem, err := fsx.NewOpenZfsFileSystem(ctx, "exampleOpenZfsFileSystem", &fsx.OpenZfsFileSystemArgs{
			StorageCapacity: pulumi.Int(64),
			SubnetIds: pulumi.String{
				aws_subnet.Example.Id,
			},
			DeploymentType:     pulumi.String("SINGLE_AZ_1"),
			ThroughputCapacity: pulumi.Int(64),
		})
		if err != nil {
			return err
		}
		exampleOpenZfsVolume, err := fsx.NewOpenZfsVolume(ctx, "exampleOpenZfsVolume", &fsx.OpenZfsVolumeArgs{
			ParentVolumeId: exampleOpenZfsFileSystem.RootVolumeId,
		})
		if err != nil {
			return err
		}
		_, err = fsx.NewOpenZfsSnapshot(ctx, "exampleOpenZfsSnapshot", &fsx.OpenZfsSnapshotArgs{
			VolumeId: exampleOpenZfsVolume.ID(),
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
import com.pulumi.aws.fsx.OpenZfsVolume;
import com.pulumi.aws.fsx.OpenZfsVolumeArgs;
import com.pulumi.aws.fsx.OpenZfsSnapshot;
import com.pulumi.aws.fsx.OpenZfsSnapshotArgs;
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
        var exampleOpenZfsFileSystem = new OpenZfsFileSystem("exampleOpenZfsFileSystem", OpenZfsFileSystemArgs.builder()        
            .storageCapacity(64)
            .subnetIds(aws_subnet.example().id())
            .deploymentType("SINGLE_AZ_1")
            .throughputCapacity(64)
            .build());

        var exampleOpenZfsVolume = new OpenZfsVolume("exampleOpenZfsVolume", OpenZfsVolumeArgs.builder()        
            .parentVolumeId(exampleOpenZfsFileSystem.rootVolumeId())
            .build());

        var exampleOpenZfsSnapshot = new OpenZfsSnapshot("exampleOpenZfsSnapshot", OpenZfsSnapshotArgs.builder()        
            .volumeId(exampleOpenZfsVolume.id())
            .build());

    }
}
```
```yaml
resources:
  exampleOpenZfsSnapshot:
    type: aws:fsx:OpenZfsSnapshot
    properties:
      volumeId: ${exampleOpenZfsVolume.id}
  exampleOpenZfsVolume:
    type: aws:fsx:OpenZfsVolume
    properties:
      parentVolumeId: ${exampleOpenZfsFileSystem.rootVolumeId}
  exampleOpenZfsFileSystem:
    type: aws:fsx:OpenZfsFileSystem
    properties:
      storageCapacity: 64
      subnetIds:
        - ${aws_subnet.example.id}
      deploymentType: SINGLE_AZ_1
      throughputCapacity: 64
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import FSx OpenZFS snapshot using the `id`. For example:

```sh
 $ pulumi import aws:fsx/openZfsSnapshot:OpenZfsSnapshot example fs-543ab12b1ca672f33
```
 