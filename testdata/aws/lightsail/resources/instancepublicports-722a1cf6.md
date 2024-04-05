Opens ports for a specific Amazon Lightsail instance, and specifies the IP addresses allowed to connect to the instance through the ports, and the protocol.

> See [What is Amazon Lightsail?](https://lightsail.aws.amazon.com/ls/docs/getting-started/article/what-is-amazon-lightsail) for more information.

> **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testInstance = new aws.lightsail.Instance("testInstance", {
    availabilityZone: data.aws_availability_zones.available.names[0],
    blueprintId: "amazon_linux_2",
    bundleId: "nano_1_0",
});
const testInstancePublicPorts = new aws.lightsail.InstancePublicPorts("testInstancePublicPorts", {
    instanceName: testInstance.name,
    portInfos: [{
        protocol: "tcp",
        fromPort: 80,
        toPort: 80,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

test_instance = aws.lightsail.Instance("testInstance",
    availability_zone=data["aws_availability_zones"]["available"]["names"],
    blueprint_id="amazon_linux_2",
    bundle_id="nano_1_0")
test_instance_public_ports = aws.lightsail.InstancePublicPorts("testInstancePublicPorts",
    instance_name=test_instance.name,
    port_infos=[aws.lightsail.InstancePublicPortsPortInfoArgs(
        protocol="tcp",
        from_port=80,
        to_port=80,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testInstance = new Aws.LightSail.Instance("testInstance", new()
    {
        AvailabilityZone = data.Aws_availability_zones.Available.Names[0],
        BlueprintId = "amazon_linux_2",
        BundleId = "nano_1_0",
    });

    var testInstancePublicPorts = new Aws.LightSail.InstancePublicPorts("testInstancePublicPorts", new()
    {
        InstanceName = testInstance.Name,
        PortInfos = new[]
        {
            new Aws.LightSail.Inputs.InstancePublicPortsPortInfoArgs
            {
                Protocol = "tcp",
                FromPort = 80,
                ToPort = 80,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testInstance, err := lightsail.NewInstance(ctx, "testInstance", &lightsail.InstanceArgs{
			AvailabilityZone: pulumi.Any(data.Aws_availability_zones.Available.Names[0]),
			BlueprintId:      pulumi.String("amazon_linux_2"),
			BundleId:         pulumi.String("nano_1_0"),
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewInstancePublicPorts(ctx, "testInstancePublicPorts", &lightsail.InstancePublicPortsArgs{
			InstanceName: testInstance.Name,
			PortInfos: lightsail.InstancePublicPortsPortInfoArray{
				&lightsail.InstancePublicPortsPortInfoArgs{
					Protocol: pulumi.String("tcp"),
					FromPort: pulumi.Int(80),
					ToPort:   pulumi.Int(80),
				},
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
import com.pulumi.aws.lightsail.Instance;
import com.pulumi.aws.lightsail.InstanceArgs;
import com.pulumi.aws.lightsail.InstancePublicPorts;
import com.pulumi.aws.lightsail.InstancePublicPortsArgs;
import com.pulumi.aws.lightsail.inputs.InstancePublicPortsPortInfoArgs;
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
        var testInstance = new Instance("testInstance", InstanceArgs.builder()        
            .availabilityZone(data.aws_availability_zones().available().names()[0])
            .blueprintId("amazon_linux_2")
            .bundleId("nano_1_0")
            .build());

        var testInstancePublicPorts = new InstancePublicPorts("testInstancePublicPorts", InstancePublicPortsArgs.builder()        
            .instanceName(testInstance.name())
            .portInfos(InstancePublicPortsPortInfoArgs.builder()
                .protocol("tcp")
                .fromPort(80)
                .toPort(80)
                .build())
            .build());

    }
}
```
```yaml
resources:
  testInstance:
    type: aws:lightsail:Instance
    properties:
      availabilityZone: ${data.aws_availability_zones.available.names[0]}
      blueprintId: amazon_linux_2
      bundleId: nano_1_0
  testInstancePublicPorts:
    type: aws:lightsail:InstancePublicPorts
    properties:
      instanceName: ${testInstance.name}
      portInfos:
        - protocol: tcp
          fromPort: 80
          toPort: 80
```
{{% /example %}}
{{% /examples %}}