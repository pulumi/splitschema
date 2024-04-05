Provides a FSx Backup resource.

## Lustre Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLustreFileSystem = new aws.fsx.LustreFileSystem("exampleLustreFileSystem", {
    storageCapacity: 1200,
    subnetIds: [aws_subnet.example.id],
    deploymentType: "PERSISTENT_1",
    perUnitStorageThroughput: 50,
});
const exampleBackup = new aws.fsx.Backup("exampleBackup", {fileSystemId: exampleLustreFileSystem.id});
```
```python
import pulumi
import pulumi_aws as aws

example_lustre_file_system = aws.fsx.LustreFileSystem("exampleLustreFileSystem",
    storage_capacity=1200,
    subnet_ids=[aws_subnet["example"]["id"]],
    deployment_type="PERSISTENT_1",
    per_unit_storage_throughput=50)
example_backup = aws.fsx.Backup("exampleBackup", file_system_id=example_lustre_file_system.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLustreFileSystem = new Aws.Fsx.LustreFileSystem("exampleLustreFileSystem", new()
    {
        StorageCapacity = 1200,
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        DeploymentType = "PERSISTENT_1",
        PerUnitStorageThroughput = 50,
    });

    var exampleBackup = new Aws.Fsx.Backup("exampleBackup", new()
    {
        FileSystemId = exampleLustreFileSystem.Id,
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
		exampleLustreFileSystem, err := fsx.NewLustreFileSystem(ctx, "exampleLustreFileSystem", &fsx.LustreFileSystemArgs{
			StorageCapacity: pulumi.Int(1200),
			SubnetIds: pulumi.String{
				aws_subnet.Example.Id,
			},
			DeploymentType:           pulumi.String("PERSISTENT_1"),
			PerUnitStorageThroughput: pulumi.Int(50),
		})
		if err != nil {
			return err
		}
		_, err = fsx.NewBackup(ctx, "exampleBackup", &fsx.BackupArgs{
			FileSystemId: exampleLustreFileSystem.ID(),
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
import com.pulumi.aws.fsx.LustreFileSystem;
import com.pulumi.aws.fsx.LustreFileSystemArgs;
import com.pulumi.aws.fsx.Backup;
import com.pulumi.aws.fsx.BackupArgs;
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
        var exampleLustreFileSystem = new LustreFileSystem("exampleLustreFileSystem", LustreFileSystemArgs.builder()        
            .storageCapacity(1200)
            .subnetIds(aws_subnet.example().id())
            .deploymentType("PERSISTENT_1")
            .perUnitStorageThroughput(50)
            .build());

        var exampleBackup = new Backup("exampleBackup", BackupArgs.builder()        
            .fileSystemId(exampleLustreFileSystem.id())
            .build());

    }
}
```
```yaml
resources:
  exampleBackup:
    type: aws:fsx:Backup
    properties:
      fileSystemId: ${exampleLustreFileSystem.id}
  exampleLustreFileSystem:
    type: aws:fsx:LustreFileSystem
    properties:
      storageCapacity: 1200
      subnetIds:
        - ${aws_subnet.example.id}
      deploymentType: PERSISTENT_1
      perUnitStorageThroughput: 50
```

## Windows Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleWindowsFileSystem = new aws.fsx.WindowsFileSystem("exampleWindowsFileSystem", {
    activeDirectoryId: aws_directory_service_directory.eample.id,
    skipFinalBackup: true,
    storageCapacity: 32,
    subnetIds: [aws_subnet.example1.id],
    throughputCapacity: 8,
});
const exampleBackup = new aws.fsx.Backup("exampleBackup", {fileSystemId: exampleWindowsFileSystem.id});
```
```python
import pulumi
import pulumi_aws as aws

example_windows_file_system = aws.fsx.WindowsFileSystem("exampleWindowsFileSystem",
    active_directory_id=aws_directory_service_directory["eample"]["id"],
    skip_final_backup=True,
    storage_capacity=32,
    subnet_ids=[aws_subnet["example1"]["id"]],
    throughput_capacity=8)
example_backup = aws.fsx.Backup("exampleBackup", file_system_id=example_windows_file_system.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleWindowsFileSystem = new Aws.Fsx.WindowsFileSystem("exampleWindowsFileSystem", new()
    {
        ActiveDirectoryId = aws_directory_service_directory.Eample.Id,
        SkipFinalBackup = true,
        StorageCapacity = 32,
        SubnetIds = new[]
        {
            aws_subnet.Example1.Id,
        },
        ThroughputCapacity = 8,
    });

    var exampleBackup = new Aws.Fsx.Backup("exampleBackup", new()
    {
        FileSystemId = exampleWindowsFileSystem.Id,
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
		exampleWindowsFileSystem, err := fsx.NewWindowsFileSystem(ctx, "exampleWindowsFileSystem", &fsx.WindowsFileSystemArgs{
			ActiveDirectoryId: pulumi.Any(aws_directory_service_directory.Eample.Id),
			SkipFinalBackup:   pulumi.Bool(true),
			StorageCapacity:   pulumi.Int(32),
			SubnetIds: pulumi.StringArray{
				aws_subnet.Example1.Id,
			},
			ThroughputCapacity: pulumi.Int(8),
		})
		if err != nil {
			return err
		}
		_, err = fsx.NewBackup(ctx, "exampleBackup", &fsx.BackupArgs{
			FileSystemId: exampleWindowsFileSystem.ID(),
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
import com.pulumi.aws.fsx.WindowsFileSystem;
import com.pulumi.aws.fsx.WindowsFileSystemArgs;
import com.pulumi.aws.fsx.Backup;
import com.pulumi.aws.fsx.BackupArgs;
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
        var exampleWindowsFileSystem = new WindowsFileSystem("exampleWindowsFileSystem", WindowsFileSystemArgs.builder()        
            .activeDirectoryId(aws_directory_service_directory.eample().id())
            .skipFinalBackup(true)
            .storageCapacity(32)
            .subnetIds(aws_subnet.example1().id())
            .throughputCapacity(8)
            .build());

        var exampleBackup = new Backup("exampleBackup", BackupArgs.builder()        
            .fileSystemId(exampleWindowsFileSystem.id())
            .build());

    }
}
```
```yaml
resources:
  exampleBackup:
    type: aws:fsx:Backup
    properties:
      fileSystemId: ${exampleWindowsFileSystem.id}
  exampleWindowsFileSystem:
    type: aws:fsx:WindowsFileSystem
    properties:
      activeDirectoryId: ${aws_directory_service_directory.eample.id}
      skipFinalBackup: true
      storageCapacity: 32
      subnetIds:
        - ${aws_subnet.example1.id}
      throughputCapacity: 8
```

## ONTAP Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOntapVolume = new aws.fsx.OntapVolume("exampleOntapVolume", {
    junctionPath: "/example",
    sizeInMegabytes: 1024,
    storageEfficiencyEnabled: true,
    storageVirtualMachineId: aws_fsx_ontap_storage_virtual_machine.test.id,
});
const exampleBackup = new aws.fsx.Backup("exampleBackup", {volumeId: exampleOntapVolume.id});
```
```python
import pulumi
import pulumi_aws as aws

example_ontap_volume = aws.fsx.OntapVolume("exampleOntapVolume",
    junction_path="/example",
    size_in_megabytes=1024,
    storage_efficiency_enabled=True,
    storage_virtual_machine_id=aws_fsx_ontap_storage_virtual_machine["test"]["id"])
example_backup = aws.fsx.Backup("exampleBackup", volume_id=example_ontap_volume.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleOntapVolume = new Aws.Fsx.OntapVolume("exampleOntapVolume", new()
    {
        JunctionPath = "/example",
        SizeInMegabytes = 1024,
        StorageEfficiencyEnabled = true,
        StorageVirtualMachineId = aws_fsx_ontap_storage_virtual_machine.Test.Id,
    });

    var exampleBackup = new Aws.Fsx.Backup("exampleBackup", new()
    {
        VolumeId = exampleOntapVolume.Id,
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
		exampleOntapVolume, err := fsx.NewOntapVolume(ctx, "exampleOntapVolume", &fsx.OntapVolumeArgs{
			JunctionPath:             pulumi.String("/example"),
			SizeInMegabytes:          pulumi.Int(1024),
			StorageEfficiencyEnabled: pulumi.Bool(true),
			StorageVirtualMachineId:  pulumi.Any(aws_fsx_ontap_storage_virtual_machine.Test.Id),
		})
		if err != nil {
			return err
		}
		_, err = fsx.NewBackup(ctx, "exampleBackup", &fsx.BackupArgs{
			VolumeId: exampleOntapVolume.ID(),
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
import com.pulumi.aws.fsx.OntapVolume;
import com.pulumi.aws.fsx.OntapVolumeArgs;
import com.pulumi.aws.fsx.Backup;
import com.pulumi.aws.fsx.BackupArgs;
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
        var exampleOntapVolume = new OntapVolume("exampleOntapVolume", OntapVolumeArgs.builder()        
            .junctionPath("/example")
            .sizeInMegabytes(1024)
            .storageEfficiencyEnabled(true)
            .storageVirtualMachineId(aws_fsx_ontap_storage_virtual_machine.test().id())
            .build());

        var exampleBackup = new Backup("exampleBackup", BackupArgs.builder()        
            .volumeId(exampleOntapVolume.id())
            .build());

    }
}
```
```yaml
resources:
  exampleBackup:
    type: aws:fsx:Backup
    properties:
      volumeId: ${exampleOntapVolume.id}
  exampleOntapVolume:
    type: aws:fsx:OntapVolume
    properties:
      junctionPath: /example
      sizeInMegabytes: 1024
      storageEfficiencyEnabled: true
      storageVirtualMachineId: ${aws_fsx_ontap_storage_virtual_machine.test.id}
```

## OpenZFS Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOpenZfsFileSystem = new aws.fsx.OpenZfsFileSystem("exampleOpenZfsFileSystem", {
    storageCapacity: 64,
    subnetIds: [aws_subnet.example.id],
    deploymentType: "SINGLE_AZ_1",
    throughputCapacity: 64,
});
const exampleBackup = new aws.fsx.Backup("exampleBackup", {fileSystemId: exampleOpenZfsFileSystem.id});
```
```python
import pulumi
import pulumi_aws as aws

example_open_zfs_file_system = aws.fsx.OpenZfsFileSystem("exampleOpenZfsFileSystem",
    storage_capacity=64,
    subnet_ids=[aws_subnet["example"]["id"]],
    deployment_type="SINGLE_AZ_1",
    throughput_capacity=64)
example_backup = aws.fsx.Backup("exampleBackup", file_system_id=example_open_zfs_file_system.id)
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

    var exampleBackup = new Aws.Fsx.Backup("exampleBackup", new()
    {
        FileSystemId = exampleOpenZfsFileSystem.Id,
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
		_, err = fsx.NewBackup(ctx, "exampleBackup", &fsx.BackupArgs{
			FileSystemId: exampleOpenZfsFileSystem.ID(),
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
import com.pulumi.aws.fsx.Backup;
import com.pulumi.aws.fsx.BackupArgs;
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

        var exampleBackup = new Backup("exampleBackup", BackupArgs.builder()        
            .fileSystemId(exampleOpenZfsFileSystem.id())
            .build());

    }
}
```
```yaml
resources:
  exampleBackup:
    type: aws:fsx:Backup
    properties:
      fileSystemId: ${exampleOpenZfsFileSystem.id}
  exampleOpenZfsFileSystem:
    type: aws:fsx:OpenZfsFileSystem
    properties:
      storageCapacity: 64
      subnetIds:
        - ${aws_subnet.example.id}
      deploymentType: SINGLE_AZ_1
      throughputCapacity: 64
```


## Import

Using `pulumi import`, import FSx Backups using the `id`. For example:

```sh
 $ pulumi import aws:fsx/backup:Backup example fs-543ab12b1ca672f33
```
 