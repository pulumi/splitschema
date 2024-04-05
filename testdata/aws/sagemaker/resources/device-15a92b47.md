Provides a SageMaker Device resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.Device("example", {
    deviceFleetName: aws_sagemaker_device_fleet.example.device_fleet_name,
    device: {
        deviceName: "example",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.Device("example",
    device_fleet_name=aws_sagemaker_device_fleet["example"]["device_fleet_name"],
    device=aws.sagemaker.DeviceDeviceArgs(
        device_name="example",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.Device("example", new()
    {
        DeviceFleetName = aws_sagemaker_device_fleet.Example.Device_fleet_name,
        DeviceDetails = new Aws.Sagemaker.Inputs.DeviceDeviceArgs
        {
            DeviceName = "example",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewDevice(ctx, "example", &sagemaker.DeviceArgs{
			DeviceFleetName: pulumi.Any(aws_sagemaker_device_fleet.Example.Device_fleet_name),
			Device: &sagemaker.DeviceDeviceArgs{
				DeviceName: pulumi.String("example"),
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
import com.pulumi.aws.sagemaker.Device;
import com.pulumi.aws.sagemaker.DeviceArgs;
import com.pulumi.aws.sagemaker.inputs.DeviceDeviceArgs;
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
        var example = new Device("example", DeviceArgs.builder()        
            .deviceFleetName(aws_sagemaker_device_fleet.example().device_fleet_name())
            .device(DeviceDeviceArgs.builder()
                .deviceName("example")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:Device
    properties:
      deviceFleetName: ${aws_sagemaker_device_fleet.example.device_fleet_name}
      device:
        deviceName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Devices using the `device-fleet-name/device-name`. For example:

```sh
 $ pulumi import aws:sagemaker/device:Device example my-fleet/my-device
```
 