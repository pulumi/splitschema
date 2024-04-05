Provides an OpsWorks instance resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const my_instance = new aws.opsworks.Instance("my-instance", {
    stackId: aws_opsworks_stack.main.id,
    layerIds: [aws_opsworks_custom_layer["my-layer"].id],
    instanceType: "t2.micro",
    os: "Amazon Linux 2015.09",
    state: "stopped",
});
```
```python
import pulumi
import pulumi_aws as aws

my_instance = aws.opsworks.Instance("my-instance",
    stack_id=aws_opsworks_stack["main"]["id"],
    layer_ids=[aws_opsworks_custom_layer["my-layer"]["id"]],
    instance_type="t2.micro",
    os="Amazon Linux 2015.09",
    state="stopped")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var my_instance = new Aws.OpsWorks.Instance("my-instance", new()
    {
        StackId = aws_opsworks_stack.Main.Id,
        LayerIds = new[]
        {
            aws_opsworks_custom_layer.My_layer.Id,
        },
        InstanceType = "t2.micro",
        Os = "Amazon Linux 2015.09",
        State = "stopped",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewInstance(ctx, "my-instance", &opsworks.InstanceArgs{
			StackId: pulumi.Any(aws_opsworks_stack.Main.Id),
			LayerIds: pulumi.StringArray{
				aws_opsworks_custom_layer.MyLayer.Id,
			},
			InstanceType: pulumi.String("t2.micro"),
			Os:           pulumi.String("Amazon Linux 2015.09"),
			State:        pulumi.String("stopped"),
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
import com.pulumi.aws.opsworks.Instance;
import com.pulumi.aws.opsworks.InstanceArgs;
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
        var my_instance = new Instance("my-instance", InstanceArgs.builder()        
            .stackId(aws_opsworks_stack.main().id())
            .layerIds(aws_opsworks_custom_layer.my-layer().id())
            .instanceType("t2.micro")
            .os("Amazon Linux 2015.09")
            .state("stopped")
            .build());

    }
}
```
```yaml
resources:
  my-instance:
    type: aws:opsworks:Instance
    properties:
      stackId: ${aws_opsworks_stack.main.id}
      layerIds:
        - ${aws_opsworks_custom_layer"my-layer"[%!s(MISSING)].id}
      instanceType: t2.micro
      os: Amazon Linux 2015.09
      state: stopped
```
{{% /example %}}
{{% /examples %}}
## Block devices

Each of the `*_block_device` attributes controls a portion of the AWS
Instance's "Block Device Mapping". It's a good idea to familiarize yourself with [AWS's Block Device
Mapping docs](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
to understand the implications of using these attributes.

### `ebs_block_device`

* `delete_on_termination` - (Optional) Whether the volume should be destroyed on instance termination. Default is `true`.
* `device_name` - (Required) Name of the device to mount.
* `iops` - (Optional) Amount of provisioned [IOPS](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). This must be set with a `volume_type` of `io1`.
* `snapshot_id` - (Optional) Snapshot ID to mount.
* `volume_size` - (Optional) Size of the volume in gigabytes.
* `volume_type` - (Optional) Type of volume. Valid values are `standard`, `gp2`, or `io1`. Default is `standard`.

Modifying any `ebs_block_device` currently requires resource replacement.

### `ephemeral_block_device`

* `device_name` - Name of the block device to mount on the instance.
* `virtual_name` - The [Instance Store Device Name](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames) (e.g., `ephemeral0`).

Each AWS Instance type has a different set of Instance Store block devices
available for attachment. AWS [publishes a
list](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes)
of which ephemeral devices are available on each type. The devices are always
identified by the `virtual_name` in the format `ephemeral{0..N}`.

### `root_block_device`

* `delete_on_termination` - (Optional) Whether the volume should be destroyed on instance termination. Default is `true`.
* `iops` - (Optional) Amount of provisioned [IOPS](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). This must be set with a `volume_type` of `io1`.
* `volume_size` - (Optional) Size of the volume in gigabytes.
* `volume_type` - (Optional) Type of volume. Valid values are `standard`, `gp2`, or `io1`. Default is `standard`.

Modifying any of the `root_block_device` settings requires resource
replacement.

> **NOTE:** Currently, changes to `*_block_device` configuration of _existing_
resources cannot be automatically detected by this provider. After making updates
to block device configuration, resource recreation can be manually triggered by
using the [`up` command with the --replace argument](https://www.pulumi.com/docs/reference/cli/pulumi_up/).


## Import

Using `pulumi import`, import Opsworks Instances using the instance `id`. For example:

```sh
 $ pulumi import aws:opsworks/instance:Instance my_instance 4d6d1710-ded9-42a1-b08e-b043ad7af1e2
```
 