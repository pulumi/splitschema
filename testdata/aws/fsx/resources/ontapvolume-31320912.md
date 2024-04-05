Manages a FSx ONTAP Volume.
See the [FSx ONTAP User Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-volumes.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.fsx.OntapVolume("test", {
    junctionPath: "/test",
    sizeInMegabytes: 1024,
    storageEfficiencyEnabled: true,
    storageVirtualMachineId: aws_fsx_ontap_storage_virtual_machine.test.id,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.fsx.OntapVolume("test",
    junction_path="/test",
    size_in_megabytes=1024,
    storage_efficiency_enabled=True,
    storage_virtual_machine_id=aws_fsx_ontap_storage_virtual_machine["test"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Fsx.OntapVolume("test", new()
    {
        JunctionPath = "/test",
        SizeInMegabytes = 1024,
        StorageEfficiencyEnabled = true,
        StorageVirtualMachineId = aws_fsx_ontap_storage_virtual_machine.Test.Id,
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
		_, err := fsx.NewOntapVolume(ctx, "test", &fsx.OntapVolumeArgs{
			JunctionPath:             pulumi.String("/test"),
			SizeInMegabytes:          pulumi.Int(1024),
			StorageEfficiencyEnabled: pulumi.Bool(true),
			StorageVirtualMachineId:  pulumi.Any(aws_fsx_ontap_storage_virtual_machine.Test.Id),
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
        var test = new OntapVolume("test", OntapVolumeArgs.builder()        
            .junctionPath("/test")
            .sizeInMegabytes(1024)
            .storageEfficiencyEnabled(true)
            .storageVirtualMachineId(aws_fsx_ontap_storage_virtual_machine.test().id())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:fsx:OntapVolume
    properties:
      junctionPath: /test
      sizeInMegabytes: 1024
      storageEfficiencyEnabled: true
      storageVirtualMachineId: ${aws_fsx_ontap_storage_virtual_machine.test.id}
```
{{% /example %}}
{{% example %}}
### Using Tiering Policy

Additional information on tiering policy with ONTAP Volumes can be found in the [FSx ONTAP Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-volumes.html).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.fsx.OntapVolume("test", {
    junctionPath: "/test",
    sizeInMegabytes: 1024,
    storageEfficiencyEnabled: true,
    storageVirtualMachineId: aws_fsx_ontap_storage_virtual_machine.test.id,
    tieringPolicy: {
        name: "AUTO",
        coolingPeriod: 31,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.fsx.OntapVolume("test",
    junction_path="/test",
    size_in_megabytes=1024,
    storage_efficiency_enabled=True,
    storage_virtual_machine_id=aws_fsx_ontap_storage_virtual_machine["test"]["id"],
    tiering_policy=aws.fsx.OntapVolumeTieringPolicyArgs(
        name="AUTO",
        cooling_period=31,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Fsx.OntapVolume("test", new()
    {
        JunctionPath = "/test",
        SizeInMegabytes = 1024,
        StorageEfficiencyEnabled = true,
        StorageVirtualMachineId = aws_fsx_ontap_storage_virtual_machine.Test.Id,
        TieringPolicy = new Aws.Fsx.Inputs.OntapVolumeTieringPolicyArgs
        {
            Name = "AUTO",
            CoolingPeriod = 31,
        },
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
		_, err := fsx.NewOntapVolume(ctx, "test", &fsx.OntapVolumeArgs{
			JunctionPath:             pulumi.String("/test"),
			SizeInMegabytes:          pulumi.Int(1024),
			StorageEfficiencyEnabled: pulumi.Bool(true),
			StorageVirtualMachineId:  pulumi.Any(aws_fsx_ontap_storage_virtual_machine.Test.Id),
			TieringPolicy: &fsx.OntapVolumeTieringPolicyArgs{
				Name:          pulumi.String("AUTO"),
				CoolingPeriod: pulumi.Int(31),
			},
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
import com.pulumi.aws.fsx.inputs.OntapVolumeTieringPolicyArgs;
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
        var test = new OntapVolume("test", OntapVolumeArgs.builder()        
            .junctionPath("/test")
            .sizeInMegabytes(1024)
            .storageEfficiencyEnabled(true)
            .storageVirtualMachineId(aws_fsx_ontap_storage_virtual_machine.test().id())
            .tieringPolicy(OntapVolumeTieringPolicyArgs.builder()
                .name("AUTO")
                .coolingPeriod(31)
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:fsx:OntapVolume
    properties:
      junctionPath: /test
      sizeInMegabytes: 1024
      storageEfficiencyEnabled: true
      storageVirtualMachineId: ${aws_fsx_ontap_storage_virtual_machine.test.id}
      tieringPolicy:
        name: AUTO
        coolingPeriod: 31
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import FSx ONTAP volume using the `id`. For example:

```sh
 $ pulumi import aws:fsx/ontapVolume:OntapVolume example fsvol-12345678abcdef123
```
 