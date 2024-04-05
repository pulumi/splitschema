Registers sources (network interfaces) with the transit gateway multicast group.
A multicast source is a network interface attached to a supported instance that sends multicast traffic.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.MulticastGroupSource("example", {
    groupIpAddress: "224.0.0.1",
    networkInterfaceId: aws_network_interface.example.id,
    transitGatewayMulticastDomainId: aws_ec2_transit_gateway_multicast_domain.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.MulticastGroupSource("example",
    group_ip_address="224.0.0.1",
    network_interface_id=aws_network_interface["example"]["id"],
    transit_gateway_multicast_domain_id=aws_ec2_transit_gateway_multicast_domain["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2TransitGateway.MulticastGroupSource("example", new()
    {
        GroupIpAddress = "224.0.0.1",
        NetworkInterfaceId = aws_network_interface.Example.Id,
        TransitGatewayMulticastDomainId = aws_ec2_transit_gateway_multicast_domain.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.NewMulticastGroupSource(ctx, "example", &ec2transitgateway.MulticastGroupSourceArgs{
			GroupIpAddress:                  pulumi.String("224.0.0.1"),
			NetworkInterfaceId:              pulumi.Any(aws_network_interface.Example.Id),
			TransitGatewayMulticastDomainId: pulumi.Any(aws_ec2_transit_gateway_multicast_domain.Example.Id),
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
import com.pulumi.aws.ec2transitgateway.MulticastGroupSource;
import com.pulumi.aws.ec2transitgateway.MulticastGroupSourceArgs;
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
        var example = new MulticastGroupSource("example", MulticastGroupSourceArgs.builder()        
            .groupIpAddress("224.0.0.1")
            .networkInterfaceId(aws_network_interface.example().id())
            .transitGatewayMulticastDomainId(aws_ec2_transit_gateway_multicast_domain.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:MulticastGroupSource
    properties:
      groupIpAddress: 224.0.0.1
      networkInterfaceId: ${aws_network_interface.example.id}
      transitGatewayMulticastDomainId: ${aws_ec2_transit_gateway_multicast_domain.example.id}
```
{{% /example %}}
{{% /examples %}}