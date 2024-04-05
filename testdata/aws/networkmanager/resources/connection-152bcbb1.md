Creates a connection between two devices.
The devices can be a physical or virtual appliance that connects to a third-party appliance in a VPC, or a physical appliance that connects to another physical appliance in an on-premises network.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkmanager.Connection("example", {
    globalNetworkId: aws_networkmanager_global_network.example.id,
    deviceId: aws_networkmanager_device.example1.id,
    connectedDeviceId: aws_networkmanager_device.example2.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkmanager.Connection("example",
    global_network_id=aws_networkmanager_global_network["example"]["id"],
    device_id=aws_networkmanager_device["example1"]["id"],
    connected_device_id=aws_networkmanager_device["example2"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkManager.Connection("example", new()
    {
        GlobalNetworkId = aws_networkmanager_global_network.Example.Id,
        DeviceId = aws_networkmanager_device.Example1.Id,
        ConnectedDeviceId = aws_networkmanager_device.Example2.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkmanager.NewConnection(ctx, "example", &networkmanager.ConnectionArgs{
			GlobalNetworkId:   pulumi.Any(aws_networkmanager_global_network.Example.Id),
			DeviceId:          pulumi.Any(aws_networkmanager_device.Example1.Id),
			ConnectedDeviceId: pulumi.Any(aws_networkmanager_device.Example2.Id),
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
import com.pulumi.aws.networkmanager.Connection;
import com.pulumi.aws.networkmanager.ConnectionArgs;
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
        var example = new Connection("example", ConnectionArgs.builder()        
            .globalNetworkId(aws_networkmanager_global_network.example().id())
            .deviceId(aws_networkmanager_device.example1().id())
            .connectedDeviceId(aws_networkmanager_device.example2().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkmanager:Connection
    properties:
      globalNetworkId: ${aws_networkmanager_global_network.example.id}
      deviceId: ${aws_networkmanager_device.example1.id}
      connectedDeviceId: ${aws_networkmanager_device.example2.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_networkmanager_connection` using the connection ARN. For example:

```sh
 $ pulumi import aws:networkmanager/connection:Connection example arn:aws:networkmanager::123456789012:device/global-network-0d47f6t230mz46dy4/connection-07f6fd08867abc123
```
 