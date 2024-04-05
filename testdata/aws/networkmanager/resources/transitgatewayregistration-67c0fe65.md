Registers a transit gateway to a global network. The transit gateway can be in any AWS Region,
but it must be owned by the same AWS account that owns the global network.
You cannot register a transit gateway in more than one global network.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGlobalNetwork = new aws.networkmanager.GlobalNetwork("exampleGlobalNetwork", {description: "example"});
const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("exampleTransitGateway", {});
const exampleTransitGatewayRegistration = new aws.networkmanager.TransitGatewayRegistration("exampleTransitGatewayRegistration", {
    globalNetworkId: exampleGlobalNetwork.id,
    transitGatewayArn: exampleTransitGateway.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example_global_network = aws.networkmanager.GlobalNetwork("exampleGlobalNetwork", description="example")
example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway")
example_transit_gateway_registration = aws.networkmanager.TransitGatewayRegistration("exampleTransitGatewayRegistration",
    global_network_id=example_global_network.id,
    transit_gateway_arn=example_transit_gateway.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleGlobalNetwork = new Aws.NetworkManager.GlobalNetwork("exampleGlobalNetwork", new()
    {
        Description = "example",
    });

    var exampleTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("exampleTransitGateway");

    var exampleTransitGatewayRegistration = new Aws.NetworkManager.TransitGatewayRegistration("exampleTransitGatewayRegistration", new()
    {
        GlobalNetworkId = exampleGlobalNetwork.Id,
        TransitGatewayArn = exampleTransitGateway.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleGlobalNetwork, err := networkmanager.NewGlobalNetwork(ctx, "exampleGlobalNetwork", &networkmanager.GlobalNetworkArgs{
			Description: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", nil)
		if err != nil {
			return err
		}
		_, err = networkmanager.NewTransitGatewayRegistration(ctx, "exampleTransitGatewayRegistration", &networkmanager.TransitGatewayRegistrationArgs{
			GlobalNetworkId:   exampleGlobalNetwork.ID(),
			TransitGatewayArn: exampleTransitGateway.Arn,
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
import com.pulumi.aws.networkmanager.GlobalNetwork;
import com.pulumi.aws.networkmanager.GlobalNetworkArgs;
import com.pulumi.aws.ec2transitgateway.TransitGateway;
import com.pulumi.aws.networkmanager.TransitGatewayRegistration;
import com.pulumi.aws.networkmanager.TransitGatewayRegistrationArgs;
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
        var exampleGlobalNetwork = new GlobalNetwork("exampleGlobalNetwork", GlobalNetworkArgs.builder()        
            .description("example")
            .build());

        var exampleTransitGateway = new TransitGateway("exampleTransitGateway");

        var exampleTransitGatewayRegistration = new TransitGatewayRegistration("exampleTransitGatewayRegistration", TransitGatewayRegistrationArgs.builder()        
            .globalNetworkId(exampleGlobalNetwork.id())
            .transitGatewayArn(exampleTransitGateway.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleGlobalNetwork:
    type: aws:networkmanager:GlobalNetwork
    properties:
      description: example
  exampleTransitGateway:
    type: aws:ec2transitgateway:TransitGateway
  exampleTransitGatewayRegistration:
    type: aws:networkmanager:TransitGatewayRegistration
    properties:
      globalNetworkId: ${exampleGlobalNetwork.id}
      transitGatewayArn: ${exampleTransitGateway.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_networkmanager_transit_gateway_registration` using the global network ID and transit gateway ARN. For example:

```sh
 $ pulumi import aws:networkmanager/transitGatewayRegistration:TransitGatewayRegistration example global-network-0d47f6t230mz46dy4,arn:aws:ec2:us-west-2:123456789012:transit-gateway/tgw-123abc05e04123abc
```
 