Associates a customer gateway with a device and optionally, with a link.
If you specify a link, it must be associated with the specified device.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGlobalNetwork = new aws.networkmanager.GlobalNetwork("exampleGlobalNetwork", {description: "example"});
const exampleSite = new aws.networkmanager.Site("exampleSite", {globalNetworkId: exampleGlobalNetwork.id});
const exampleDevice = new aws.networkmanager.Device("exampleDevice", {
    globalNetworkId: exampleGlobalNetwork.id,
    siteId: exampleSite.id,
});
const exampleCustomerGateway = new aws.ec2.CustomerGateway("exampleCustomerGateway", {
    bgpAsn: "65000",
    ipAddress: "172.83.124.10",
    type: "ipsec.1",
});
const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("exampleTransitGateway", {});
const exampleVpnConnection = new aws.ec2.VpnConnection("exampleVpnConnection", {
    customerGatewayId: exampleCustomerGateway.id,
    transitGatewayId: exampleTransitGateway.id,
    type: exampleCustomerGateway.type,
    staticRoutesOnly: true,
});
const exampleTransitGatewayRegistration = new aws.networkmanager.TransitGatewayRegistration("exampleTransitGatewayRegistration", {
    globalNetworkId: exampleGlobalNetwork.id,
    transitGatewayArn: exampleTransitGateway.arn,
}, {
    dependsOn: [exampleVpnConnection],
});
const exampleCustomerGatewayAssociation = new aws.networkmanager.CustomerGatewayAssociation("exampleCustomerGatewayAssociation", {
    globalNetworkId: exampleGlobalNetwork.id,
    customerGatewayArn: exampleCustomerGateway.arn,
    deviceId: exampleDevice.id,
}, {
    dependsOn: [exampleTransitGatewayRegistration],
});
```
```python
import pulumi
import pulumi_aws as aws

example_global_network = aws.networkmanager.GlobalNetwork("exampleGlobalNetwork", description="example")
example_site = aws.networkmanager.Site("exampleSite", global_network_id=example_global_network.id)
example_device = aws.networkmanager.Device("exampleDevice",
    global_network_id=example_global_network.id,
    site_id=example_site.id)
example_customer_gateway = aws.ec2.CustomerGateway("exampleCustomerGateway",
    bgp_asn="65000",
    ip_address="172.83.124.10",
    type="ipsec.1")
example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway")
example_vpn_connection = aws.ec2.VpnConnection("exampleVpnConnection",
    customer_gateway_id=example_customer_gateway.id,
    transit_gateway_id=example_transit_gateway.id,
    type=example_customer_gateway.type,
    static_routes_only=True)
example_transit_gateway_registration = aws.networkmanager.TransitGatewayRegistration("exampleTransitGatewayRegistration",
    global_network_id=example_global_network.id,
    transit_gateway_arn=example_transit_gateway.arn,
    opts=pulumi.ResourceOptions(depends_on=[example_vpn_connection]))
example_customer_gateway_association = aws.networkmanager.CustomerGatewayAssociation("exampleCustomerGatewayAssociation",
    global_network_id=example_global_network.id,
    customer_gateway_arn=example_customer_gateway.arn,
    device_id=example_device.id,
    opts=pulumi.ResourceOptions(depends_on=[example_transit_gateway_registration]))
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

    var exampleSite = new Aws.NetworkManager.Site("exampleSite", new()
    {
        GlobalNetworkId = exampleGlobalNetwork.Id,
    });

    var exampleDevice = new Aws.NetworkManager.Device("exampleDevice", new()
    {
        GlobalNetworkId = exampleGlobalNetwork.Id,
        SiteId = exampleSite.Id,
    });

    var exampleCustomerGateway = new Aws.Ec2.CustomerGateway("exampleCustomerGateway", new()
    {
        BgpAsn = "65000",
        IpAddress = "172.83.124.10",
        Type = "ipsec.1",
    });

    var exampleTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("exampleTransitGateway");

    var exampleVpnConnection = new Aws.Ec2.VpnConnection("exampleVpnConnection", new()
    {
        CustomerGatewayId = exampleCustomerGateway.Id,
        TransitGatewayId = exampleTransitGateway.Id,
        Type = exampleCustomerGateway.Type,
        StaticRoutesOnly = true,
    });

    var exampleTransitGatewayRegistration = new Aws.NetworkManager.TransitGatewayRegistration("exampleTransitGatewayRegistration", new()
    {
        GlobalNetworkId = exampleGlobalNetwork.Id,
        TransitGatewayArn = exampleTransitGateway.Arn,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleVpnConnection,
        },
    });

    var exampleCustomerGatewayAssociation = new Aws.NetworkManager.CustomerGatewayAssociation("exampleCustomerGatewayAssociation", new()
    {
        GlobalNetworkId = exampleGlobalNetwork.Id,
        CustomerGatewayArn = exampleCustomerGateway.Arn,
        DeviceId = exampleDevice.Id,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleTransitGatewayRegistration,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
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
		exampleSite, err := networkmanager.NewSite(ctx, "exampleSite", &networkmanager.SiteArgs{
			GlobalNetworkId: exampleGlobalNetwork.ID(),
		})
		if err != nil {
			return err
		}
		exampleDevice, err := networkmanager.NewDevice(ctx, "exampleDevice", &networkmanager.DeviceArgs{
			GlobalNetworkId: exampleGlobalNetwork.ID(),
			SiteId:          exampleSite.ID(),
		})
		if err != nil {
			return err
		}
		exampleCustomerGateway, err := ec2.NewCustomerGateway(ctx, "exampleCustomerGateway", &ec2.CustomerGatewayArgs{
			BgpAsn:    pulumi.String("65000"),
			IpAddress: pulumi.String("172.83.124.10"),
			Type:      pulumi.String("ipsec.1"),
		})
		if err != nil {
			return err
		}
		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", nil)
		if err != nil {
			return err
		}
		exampleVpnConnection, err := ec2.NewVpnConnection(ctx, "exampleVpnConnection", &ec2.VpnConnectionArgs{
			CustomerGatewayId: exampleCustomerGateway.ID(),
			TransitGatewayId:  exampleTransitGateway.ID(),
			Type:              exampleCustomerGateway.Type,
			StaticRoutesOnly:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		exampleTransitGatewayRegistration, err := networkmanager.NewTransitGatewayRegistration(ctx, "exampleTransitGatewayRegistration", &networkmanager.TransitGatewayRegistrationArgs{
			GlobalNetworkId:   exampleGlobalNetwork.ID(),
			TransitGatewayArn: exampleTransitGateway.Arn,
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleVpnConnection,
		}))
		if err != nil {
			return err
		}
		_, err = networkmanager.NewCustomerGatewayAssociation(ctx, "exampleCustomerGatewayAssociation", &networkmanager.CustomerGatewayAssociationArgs{
			GlobalNetworkId:    exampleGlobalNetwork.ID(),
			CustomerGatewayArn: exampleCustomerGateway.Arn,
			DeviceId:           exampleDevice.ID(),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleTransitGatewayRegistration,
		}))
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
import com.pulumi.aws.networkmanager.Site;
import com.pulumi.aws.networkmanager.SiteArgs;
import com.pulumi.aws.networkmanager.Device;
import com.pulumi.aws.networkmanager.DeviceArgs;
import com.pulumi.aws.ec2.CustomerGateway;
import com.pulumi.aws.ec2.CustomerGatewayArgs;
import com.pulumi.aws.ec2transitgateway.TransitGateway;
import com.pulumi.aws.ec2.VpnConnection;
import com.pulumi.aws.ec2.VpnConnectionArgs;
import com.pulumi.aws.networkmanager.TransitGatewayRegistration;
import com.pulumi.aws.networkmanager.TransitGatewayRegistrationArgs;
import com.pulumi.aws.networkmanager.CustomerGatewayAssociation;
import com.pulumi.aws.networkmanager.CustomerGatewayAssociationArgs;
import com.pulumi.resources.CustomResourceOptions;
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

        var exampleSite = new Site("exampleSite", SiteArgs.builder()        
            .globalNetworkId(exampleGlobalNetwork.id())
            .build());

        var exampleDevice = new Device("exampleDevice", DeviceArgs.builder()        
            .globalNetworkId(exampleGlobalNetwork.id())
            .siteId(exampleSite.id())
            .build());

        var exampleCustomerGateway = new CustomerGateway("exampleCustomerGateway", CustomerGatewayArgs.builder()        
            .bgpAsn(65000)
            .ipAddress("172.83.124.10")
            .type("ipsec.1")
            .build());

        var exampleTransitGateway = new TransitGateway("exampleTransitGateway");

        var exampleVpnConnection = new VpnConnection("exampleVpnConnection", VpnConnectionArgs.builder()        
            .customerGatewayId(exampleCustomerGateway.id())
            .transitGatewayId(exampleTransitGateway.id())
            .type(exampleCustomerGateway.type())
            .staticRoutesOnly(true)
            .build());

        var exampleTransitGatewayRegistration = new TransitGatewayRegistration("exampleTransitGatewayRegistration", TransitGatewayRegistrationArgs.builder()        
            .globalNetworkId(exampleGlobalNetwork.id())
            .transitGatewayArn(exampleTransitGateway.arn())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleVpnConnection)
                .build());

        var exampleCustomerGatewayAssociation = new CustomerGatewayAssociation("exampleCustomerGatewayAssociation", CustomerGatewayAssociationArgs.builder()        
            .globalNetworkId(exampleGlobalNetwork.id())
            .customerGatewayArn(exampleCustomerGateway.arn())
            .deviceId(exampleDevice.id())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleTransitGatewayRegistration)
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
  exampleSite:
    type: aws:networkmanager:Site
    properties:
      globalNetworkId: ${exampleGlobalNetwork.id}
  exampleDevice:
    type: aws:networkmanager:Device
    properties:
      globalNetworkId: ${exampleGlobalNetwork.id}
      siteId: ${exampleSite.id}
  exampleCustomerGateway:
    type: aws:ec2:CustomerGateway
    properties:
      bgpAsn: 65000
      ipAddress: 172.83.124.10
      type: ipsec.1
  exampleTransitGateway:
    type: aws:ec2transitgateway:TransitGateway
  exampleVpnConnection:
    type: aws:ec2:VpnConnection
    properties:
      customerGatewayId: ${exampleCustomerGateway.id}
      transitGatewayId: ${exampleTransitGateway.id}
      type: ${exampleCustomerGateway.type}
      staticRoutesOnly: true
  exampleTransitGatewayRegistration:
    type: aws:networkmanager:TransitGatewayRegistration
    properties:
      globalNetworkId: ${exampleGlobalNetwork.id}
      transitGatewayArn: ${exampleTransitGateway.arn}
    options:
      dependson:
        - ${exampleVpnConnection}
  exampleCustomerGatewayAssociation:
    type: aws:networkmanager:CustomerGatewayAssociation
    properties:
      globalNetworkId: ${exampleGlobalNetwork.id}
      customerGatewayArn: ${exampleCustomerGateway.arn}
      deviceId: ${exampleDevice.id}
    options:
      dependson:
        - ${exampleTransitGatewayRegistration}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_networkmanager_customer_gateway_association` using the global network ID and customer gateway ARN. For example:

```sh
 $ pulumi import aws:networkmanager/customerGatewayAssociation:CustomerGatewayAssociation example global-network-0d47f6t230mz46dy4,arn:aws:ec2:us-west-2:123456789012:customer-gateway/cgw-123abc05e04123abc
```
 