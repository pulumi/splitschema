Attaches a Lightsail disk to a Lightsail Instance

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const available = aws.getAvailabilityZones({
    state: "available",
    filters: [{
        name: "opt-in-status",
        values: ["opt-in-not-required"],
    }],
});
const testDisk = new aws.lightsail.Disk("testDisk", {
    sizeInGb: 8,
    availabilityZone: available.then(available => available.names?.[0]),
});
const testInstance = new aws.lightsail.Instance("testInstance", {
    availabilityZone: available.then(available => available.names?.[0]),
    blueprintId: "amazon_linux_2",
    bundleId: "nano_1_0",
});
const testDisk_attachment = new aws.lightsail.Disk_attachment("testDisk_attachment", {
    diskName: testDisk.name,
    instanceName: testInstance.name,
    diskPath: "/dev/xvdf",
});
```
```python
import pulumi
import pulumi_aws as aws

available = aws.get_availability_zones(state="available",
    filters=[aws.GetAvailabilityZonesFilterArgs(
        name="opt-in-status",
        values=["opt-in-not-required"],
    )])
test_disk = aws.lightsail.Disk("testDisk",
    size_in_gb=8,
    availability_zone=available.names[0])
test_instance = aws.lightsail.Instance("testInstance",
    availability_zone=available.names[0],
    blueprint_id="amazon_linux_2",
    bundle_id="nano_1_0")
test_disk_attachment = aws.lightsail.Disk_attachment("testDisk_attachment",
    disk_name=test_disk.name,
    instance_name=test_instance.name,
    disk_path="/dev/xvdf")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var available = Aws.GetAvailabilityZones.Invoke(new()
    {
        State = "available",
        Filters = new[]
        {
            new Aws.Inputs.GetAvailabilityZonesFilterInputArgs
            {
                Name = "opt-in-status",
                Values = new[]
                {
                    "opt-in-not-required",
                },
            },
        },
    });

    var testDisk = new Aws.LightSail.Disk("testDisk", new()
    {
        SizeInGb = 8,
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
    });

    var testInstance = new Aws.LightSail.Instance("testInstance", new()
    {
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
        BlueprintId = "amazon_linux_2",
        BundleId = "nano_1_0",
    });

    var testDisk_attachment = new Aws.LightSail.Disk_attachment("testDisk_attachment", new()
    {
        DiskName = testDisk.Name,
        InstanceName = testInstance.Name,
        DiskPath = "/dev/xvdf",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		available, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			State: pulumi.StringRef("available"),
			Filters: []aws.GetAvailabilityZonesFilter{
				{
					Name: "opt-in-status",
					Values: []string{
						"opt-in-not-required",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		testDisk, err := lightsail.NewDisk(ctx, "testDisk", &lightsail.DiskArgs{
			SizeInGb:         pulumi.Int(8),
			AvailabilityZone: *pulumi.String(available.Names[0]),
		})
		if err != nil {
			return err
		}
		testInstance, err := lightsail.NewInstance(ctx, "testInstance", &lightsail.InstanceArgs{
			AvailabilityZone: *pulumi.String(available.Names[0]),
			BlueprintId:      pulumi.String("amazon_linux_2"),
			BundleId:         pulumi.String("nano_1_0"),
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewDisk_attachment(ctx, "testDisk_attachment", &lightsail.Disk_attachmentArgs{
			DiskName:     testDisk.Name,
			InstanceName: testInstance.Name,
			DiskPath:     pulumi.String("/dev/xvdf"),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
import com.pulumi.aws.lightsail.Disk;
import com.pulumi.aws.lightsail.DiskArgs;
import com.pulumi.aws.lightsail.Instance;
import com.pulumi.aws.lightsail.InstanceArgs;
import com.pulumi.aws.lightsail.Disk_attachment;
import com.pulumi.aws.lightsail.Disk_attachmentArgs;
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
        final var available = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
            .state("available")
            .filters(GetAvailabilityZonesFilterArgs.builder()
                .name("opt-in-status")
                .values("opt-in-not-required")
                .build())
            .build());

        var testDisk = new Disk("testDisk", DiskArgs.builder()        
            .sizeInGb(8)
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]))
            .build());

        var testInstance = new Instance("testInstance", InstanceArgs.builder()        
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]))
            .blueprintId("amazon_linux_2")
            .bundleId("nano_1_0")
            .build());

        var testDisk_attachment = new Disk_attachment("testDisk_attachment", Disk_attachmentArgs.builder()        
            .diskName(testDisk.name())
            .instanceName(testInstance.name())
            .diskPath("/dev/xvdf")
            .build());

    }
}
```
```yaml
resources:
  testDisk:
    type: aws:lightsail:Disk
    properties:
      sizeInGb: 8
      availabilityZone: ${available.names[0]}
  testInstance:
    type: aws:lightsail:Instance
    properties:
      availabilityZone: ${available.names[0]}
      blueprintId: amazon_linux_2
      bundleId: nano_1_0
  testDisk_attachment:
    type: aws:lightsail:Disk_attachment
    properties:
      diskName: ${testDisk.name}
      instanceName: ${testInstance.name}
      diskPath: /dev/xvdf
variables:
  available:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments:
        state: available
        filters:
          - name: opt-in-status
            values:
              - opt-in-not-required
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_lightsail_disk` using the id attribute. For example:

```sh
 $ pulumi import aws:lightsail/disk_attachment:Disk_attachment test test-disk,test-instance
```
 