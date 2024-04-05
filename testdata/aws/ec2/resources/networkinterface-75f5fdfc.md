Provides an Elastic network interface (ENI) resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ec2.NetworkInterface("test", {
    subnetId: aws_subnet.public_a.id,
    privateIps: ["10.0.0.50"],
    securityGroups: [aws_security_group.web.id],
    attachments: [{
        instance: aws_instance.test.id,
        deviceIndex: 1,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ec2.NetworkInterface("test",
    subnet_id=aws_subnet["public_a"]["id"],
    private_ips=["10.0.0.50"],
    security_groups=[aws_security_group["web"]["id"]],
    attachments=[aws.ec2.NetworkInterfaceAttachmentArgs(
        instance=aws_instance["test"]["id"],
        device_index=1,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ec2.NetworkInterface("test", new()
    {
        SubnetId = aws_subnet.Public_a.Id,
        PrivateIps = new[]
        {
            "10.0.0.50",
        },
        SecurityGroups = new[]
        {
            aws_security_group.Web.Id,
        },
        Attachments = new[]
        {
            new Aws.Ec2.Inputs.NetworkInterfaceAttachmentArgs
            {
                Instance = aws_instance.Test.Id,
                DeviceIndex = 1,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewNetworkInterface(ctx, "test", &ec2.NetworkInterfaceArgs{
			SubnetId: pulumi.Any(aws_subnet.Public_a.Id),
			PrivateIps: pulumi.StringArray{
				pulumi.String("10.0.0.50"),
			},
			SecurityGroups: pulumi.StringArray{
				aws_security_group.Web.Id,
			},
			Attachments: ec2.NetworkInterfaceAttachmentTypeArray{
				&ec2.NetworkInterfaceAttachmentTypeArgs{
					Instance:    pulumi.Any(aws_instance.Test.Id),
					DeviceIndex: pulumi.Int(1),
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
import com.pulumi.aws.ec2.NetworkInterface;
import com.pulumi.aws.ec2.NetworkInterfaceArgs;
import com.pulumi.aws.ec2.inputs.NetworkInterfaceAttachmentArgs;
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
        var test = new NetworkInterface("test", NetworkInterfaceArgs.builder()        
            .subnetId(aws_subnet.public_a().id())
            .privateIps("10.0.0.50")
            .securityGroups(aws_security_group.web().id())
            .attachments(NetworkInterfaceAttachmentArgs.builder()
                .instance(aws_instance.test().id())
                .deviceIndex(1)
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ec2:NetworkInterface
    properties:
      subnetId: ${aws_subnet.public_a.id}
      privateIps:
        - 10.0.0.50
      securityGroups:
        - ${aws_security_group.web.id}
      attachments:
        - instance: ${aws_instance.test.id}
          deviceIndex: 1
```
{{% /example %}}
### Example of Managing Multiple IPs on a Network Interface

By default, private IPs are managed through the `private_ips` and `private_ips_count` arguments which manage IPs as a set of IPs that are configured without regard to order. For a new network interface, the same primary IP address is consistently selected from a given set of addresses, regardless of the order provided. However, modifications of the set of addresses of an existing interface will not alter the current primary IP address unless it has been removed from the set.

In order to manage the private IPs as a sequentially ordered list, configure `private_ip_list_enabled` to `true` and use `private_ip_list` to manage the IPs. This will disable the `private_ips` and `private_ips_count` settings, which must be removed from the config file but are still exported. Note that changing the first address of `private_ip_list`, which is the primary, always requires a new interface.

If you are managing a specific set or list of IPs, instead of just using `private_ips_count`, this is a potential workflow for also leveraging `private_ips_count` to have AWS automatically assign additional IP addresses:

1. Comment out `private_ips`, `private_ip_list`, `private_ip_list_enabled` in your configuration
2. Set the desired `private_ips_count` (count of the number of secondaries, the primary is not included)
3. Apply to assign the extra IPs
4. Remove `private_ips_count` and restore your settings from the first step
5. Add the new IPs to your current settings
6. Apply again to update the stored state

This process can also be used to remove IP addresses in addition to the option of manually removing them. Adding IP addresses in a manually is more difficult because it requires knowledge of which addresses are available.
{{% /examples %}}

## Import

Using `pulumi import`, import Network Interfaces using the `id`. For example:

```sh
 $ pulumi import aws:ec2/networkInterface:NetworkInterface test eni-e5aa89a3
```
 