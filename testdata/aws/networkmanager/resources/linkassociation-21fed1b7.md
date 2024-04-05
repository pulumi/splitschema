Associates a link to a device.
A device can be associated to multiple links and a link can be associated to multiple devices.
The device and link must be in the same global network and the same site.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkmanager.LinkAssociation("example", {
    globalNetworkId: aws_networkmanager_global_network.example.id,
    linkId: aws_networkmanager_link.example.id,
    deviceId: aws_networkmanager_device.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkmanager.LinkAssociation("example",
    global_network_id=aws_networkmanager_global_network["example"]["id"],
    link_id=aws_networkmanager_link["example"]["id"],
    device_id=aws_networkmanager_device["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkManager.LinkAssociation("example", new()
    {
        GlobalNetworkId = aws_networkmanager_global_network.Example.Id,
        LinkId = aws_networkmanager_link.Example.Id,
        DeviceId = aws_networkmanager_device.Example.Id,
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
		_, err := networkmanager.NewLinkAssociation(ctx, "example", &networkmanager.LinkAssociationArgs{
			GlobalNetworkId: pulumi.Any(aws_networkmanager_global_network.Example.Id),
			LinkId:          pulumi.Any(aws_networkmanager_link.Example.Id),
			DeviceId:        pulumi.Any(aws_networkmanager_device.Example.Id),
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
import com.pulumi.aws.networkmanager.LinkAssociation;
import com.pulumi.aws.networkmanager.LinkAssociationArgs;
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
        var example = new LinkAssociation("example", LinkAssociationArgs.builder()        
            .globalNetworkId(aws_networkmanager_global_network.example().id())
            .linkId(aws_networkmanager_link.example().id())
            .deviceId(aws_networkmanager_device.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkmanager:LinkAssociation
    properties:
      globalNetworkId: ${aws_networkmanager_global_network.example.id}
      linkId: ${aws_networkmanager_link.example.id}
      deviceId: ${aws_networkmanager_device.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_networkmanager_link_association` using the global network ID, link ID and device ID. For example:

```sh
 $ pulumi import aws:networkmanager/linkAssociation:LinkAssociation example global-network-0d47f6t230mz46dy4,link-444555aaabbb11223,device-07f6fd08867abc123
```
 