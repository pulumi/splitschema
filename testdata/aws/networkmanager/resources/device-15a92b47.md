Creates a device in a global network. If you specify both a site ID and a location,
the location of the site is used for visualization in the Network Manager console.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkmanager.Device("example", {
    globalNetworkId: aws_networkmanager_global_network.example.id,
    siteId: aws_networkmanager_site.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkmanager.Device("example",
    global_network_id=aws_networkmanager_global_network["example"]["id"],
    site_id=aws_networkmanager_site["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkManager.Device("example", new()
    {
        GlobalNetworkId = aws_networkmanager_global_network.Example.Id,
        SiteId = aws_networkmanager_site.Example.Id,
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
		_, err := networkmanager.NewDevice(ctx, "example", &networkmanager.DeviceArgs{
			GlobalNetworkId: pulumi.Any(aws_networkmanager_global_network.Example.Id),
			SiteId:          pulumi.Any(aws_networkmanager_site.Example.Id),
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
import com.pulumi.aws.networkmanager.Device;
import com.pulumi.aws.networkmanager.DeviceArgs;
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
            .globalNetworkId(aws_networkmanager_global_network.example().id())
            .siteId(aws_networkmanager_site.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkmanager:Device
    properties:
      globalNetworkId: ${aws_networkmanager_global_network.example.id}
      siteId: ${aws_networkmanager_site.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_networkmanager_device` using the device ARN. For example:

```sh
 $ pulumi import aws:networkmanager/device:Device example arn:aws:networkmanager::123456789012:device/global-network-0d47f6t230mz46dy4/device-07f6fd08867abc123
```
 