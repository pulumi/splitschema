Manages a Site-to-Site VPN connection. A Site-to-Site VPN connection is an Internet Protocol security (IPsec) VPN connection between a VPC and an on-premises network.
Any new Site-to-Site VPN connection that you create is an [AWS VPN connection](https://docs.aws.amazon.com/vpn/latest/s2svpn/vpn-categories.html).

> **Note:** The CIDR blocks in the arguments `tunnel1_inside_cidr` and `tunnel2_inside_cidr` must have a prefix of /30 and be a part of a specific range.
[Read more about this in the AWS documentation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpnTunnelOptionsSpecification.html).

{{% examples %}}
## Example Usage
{{% example %}}
### EC2 Transit Gateway

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("exampleTransitGateway", {});
const exampleCustomerGateway = new aws.ec2.CustomerGateway("exampleCustomerGateway", {
    bgpAsn: "65000",
    ipAddress: "172.0.0.1",
    type: "ipsec.1",
});
const exampleVpnConnection = new aws.ec2.VpnConnection("exampleVpnConnection", {
    customerGatewayId: exampleCustomerGateway.id,
    transitGatewayId: exampleTransitGateway.id,
    type: exampleCustomerGateway.type,
});
```
```python
import pulumi
import pulumi_aws as aws

example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway")
example_customer_gateway = aws.ec2.CustomerGateway("exampleCustomerGateway",
    bgp_asn="65000",
    ip_address="172.0.0.1",
    type="ipsec.1")
example_vpn_connection = aws.ec2.VpnConnection("exampleVpnConnection",
    customer_gateway_id=example_customer_gateway.id,
    transit_gateway_id=example_transit_gateway.id,
    type=example_customer_gateway.type)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("exampleTransitGateway");

    var exampleCustomerGateway = new Aws.Ec2.CustomerGateway("exampleCustomerGateway", new()
    {
        BgpAsn = "65000",
        IpAddress = "172.0.0.1",
        Type = "ipsec.1",
    });

    var exampleVpnConnection = new Aws.Ec2.VpnConnection("exampleVpnConnection", new()
    {
        CustomerGatewayId = exampleCustomerGateway.Id,
        TransitGatewayId = exampleTransitGateway.Id,
        Type = exampleCustomerGateway.Type,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", nil)
		if err != nil {
			return err
		}
		exampleCustomerGateway, err := ec2.NewCustomerGateway(ctx, "exampleCustomerGateway", &ec2.CustomerGatewayArgs{
			BgpAsn:    pulumi.String("65000"),
			IpAddress: pulumi.String("172.0.0.1"),
			Type:      pulumi.String("ipsec.1"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpnConnection(ctx, "exampleVpnConnection", &ec2.VpnConnectionArgs{
			CustomerGatewayId: exampleCustomerGateway.ID(),
			TransitGatewayId:  exampleTransitGateway.ID(),
			Type:              exampleCustomerGateway.Type,
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
import com.pulumi.aws.ec2transitgateway.TransitGateway;
import com.pulumi.aws.ec2.CustomerGateway;
import com.pulumi.aws.ec2.CustomerGatewayArgs;
import com.pulumi.aws.ec2.VpnConnection;
import com.pulumi.aws.ec2.VpnConnectionArgs;
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
        var exampleTransitGateway = new TransitGateway("exampleTransitGateway");

        var exampleCustomerGateway = new CustomerGateway("exampleCustomerGateway", CustomerGatewayArgs.builder()        
            .bgpAsn(65000)
            .ipAddress("172.0.0.1")
            .type("ipsec.1")
            .build());

        var exampleVpnConnection = new VpnConnection("exampleVpnConnection", VpnConnectionArgs.builder()        
            .customerGatewayId(exampleCustomerGateway.id())
            .transitGatewayId(exampleTransitGateway.id())
            .type(exampleCustomerGateway.type())
            .build());

    }
}
```
```yaml
resources:
  exampleTransitGateway:
    type: aws:ec2transitgateway:TransitGateway
  exampleCustomerGateway:
    type: aws:ec2:CustomerGateway
    properties:
      bgpAsn: 65000
      ipAddress: 172.0.0.1
      type: ipsec.1
  exampleVpnConnection:
    type: aws:ec2:VpnConnection
    properties:
      customerGatewayId: ${exampleCustomerGateway.id}
      transitGatewayId: ${exampleTransitGateway.id}
      type: ${exampleCustomerGateway.type}
```
{{% /example %}}
{{% example %}}
### Virtual Private Gateway

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const vpc = new aws.ec2.Vpc("vpc", {cidrBlock: "10.0.0.0/16"});
const vpnGateway = new aws.ec2.VpnGateway("vpnGateway", {vpcId: vpc.id});
const customerGateway = new aws.ec2.CustomerGateway("customerGateway", {
    bgpAsn: "65000",
    ipAddress: "172.0.0.1",
    type: "ipsec.1",
});
const main = new aws.ec2.VpnConnection("main", {
    vpnGatewayId: vpnGateway.id,
    customerGatewayId: customerGateway.id,
    type: "ipsec.1",
    staticRoutesOnly: true,
});
```
```python
import pulumi
import pulumi_aws as aws

vpc = aws.ec2.Vpc("vpc", cidr_block="10.0.0.0/16")
vpn_gateway = aws.ec2.VpnGateway("vpnGateway", vpc_id=vpc.id)
customer_gateway = aws.ec2.CustomerGateway("customerGateway",
    bgp_asn="65000",
    ip_address="172.0.0.1",
    type="ipsec.1")
main = aws.ec2.VpnConnection("main",
    vpn_gateway_id=vpn_gateway.id,
    customer_gateway_id=customer_gateway.id,
    type="ipsec.1",
    static_routes_only=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var vpc = new Aws.Ec2.Vpc("vpc", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var vpnGateway = new Aws.Ec2.VpnGateway("vpnGateway", new()
    {
        VpcId = vpc.Id,
    });

    var customerGateway = new Aws.Ec2.CustomerGateway("customerGateway", new()
    {
        BgpAsn = "65000",
        IpAddress = "172.0.0.1",
        Type = "ipsec.1",
    });

    var main = new Aws.Ec2.VpnConnection("main", new()
    {
        VpnGatewayId = vpnGateway.Id,
        CustomerGatewayId = customerGateway.Id,
        Type = "ipsec.1",
        StaticRoutesOnly = true,
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
		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		vpnGateway, err := ec2.NewVpnGateway(ctx, "vpnGateway", &ec2.VpnGatewayArgs{
			VpcId: vpc.ID(),
		})
		if err != nil {
			return err
		}
		customerGateway, err := ec2.NewCustomerGateway(ctx, "customerGateway", &ec2.CustomerGatewayArgs{
			BgpAsn:    pulumi.String("65000"),
			IpAddress: pulumi.String("172.0.0.1"),
			Type:      pulumi.String("ipsec.1"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpnConnection(ctx, "main", &ec2.VpnConnectionArgs{
			VpnGatewayId:      vpnGateway.ID(),
			CustomerGatewayId: customerGateway.ID(),
			Type:              pulumi.String("ipsec.1"),
			StaticRoutesOnly:  pulumi.Bool(true),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.VpnGateway;
import com.pulumi.aws.ec2.VpnGatewayArgs;
import com.pulumi.aws.ec2.CustomerGateway;
import com.pulumi.aws.ec2.CustomerGatewayArgs;
import com.pulumi.aws.ec2.VpnConnection;
import com.pulumi.aws.ec2.VpnConnectionArgs;
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
        var vpc = new Vpc("vpc", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var vpnGateway = new VpnGateway("vpnGateway", VpnGatewayArgs.builder()        
            .vpcId(vpc.id())
            .build());

        var customerGateway = new CustomerGateway("customerGateway", CustomerGatewayArgs.builder()        
            .bgpAsn(65000)
            .ipAddress("172.0.0.1")
            .type("ipsec.1")
            .build());

        var main = new VpnConnection("main", VpnConnectionArgs.builder()        
            .vpnGatewayId(vpnGateway.id())
            .customerGatewayId(customerGateway.id())
            .type("ipsec.1")
            .staticRoutesOnly(true)
            .build());

    }
}
```
```yaml
resources:
  vpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  vpnGateway:
    type: aws:ec2:VpnGateway
    properties:
      vpcId: ${vpc.id}
  customerGateway:
    type: aws:ec2:CustomerGateway
    properties:
      bgpAsn: 65000
      ipAddress: 172.0.0.1
      type: ipsec.1
  main:
    type: aws:ec2:VpnConnection
    properties:
      vpnGatewayId: ${vpnGateway.id}
      customerGatewayId: ${customerGateway.id}
      type: ipsec.1
      staticRoutesOnly: true
```
{{% /example %}}
{{% example %}}
### AWS Site to Site Private VPN

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGateway = new aws.directconnect.Gateway("exampleGateway", {amazonSideAsn: "64512"});
const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("exampleTransitGateway", {
    amazonSideAsn: 64513,
    description: "example_ipsec_vpn_example",
    transitGatewayCidrBlocks: ["10.0.0.0/24"],
});
const exampleCustomerGateway = new aws.ec2.CustomerGateway("exampleCustomerGateway", {
    bgpAsn: "64514",
    ipAddress: "10.0.0.1",
    type: "ipsec.1",
    tags: {
        Name: "example_ipsec_vpn_example",
    },
});
const exampleGatewayAssociation = new aws.directconnect.GatewayAssociation("exampleGatewayAssociation", {
    dxGatewayId: exampleGateway.id,
    associatedGatewayId: exampleTransitGateway.id,
    allowedPrefixes: ["10.0.0.0/8"],
});
const exampleDirectConnectGatewayAttachment = aws.ec2transitgateway.getDirectConnectGatewayAttachmentOutput({
    transitGatewayId: exampleTransitGateway.id,
    dxGatewayId: exampleGateway.id,
});
const exampleVpnConnection = new aws.ec2.VpnConnection("exampleVpnConnection", {
    customerGatewayId: exampleCustomerGateway.id,
    outsideIpAddressType: "PrivateIpv4",
    transitGatewayId: exampleTransitGateway.id,
    transportTransitGatewayAttachmentId: exampleDirectConnectGatewayAttachment.apply(exampleDirectConnectGatewayAttachment => exampleDirectConnectGatewayAttachment.id),
    type: "ipsec.1",
    tags: {
        Name: "example_ipsec_vpn_example",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn="64512")
example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway",
    amazon_side_asn=64513,
    description="example_ipsec_vpn_example",
    transit_gateway_cidr_blocks=["10.0.0.0/24"])
example_customer_gateway = aws.ec2.CustomerGateway("exampleCustomerGateway",
    bgp_asn="64514",
    ip_address="10.0.0.1",
    type="ipsec.1",
    tags={
        "Name": "example_ipsec_vpn_example",
    })
example_gateway_association = aws.directconnect.GatewayAssociation("exampleGatewayAssociation",
    dx_gateway_id=example_gateway.id,
    associated_gateway_id=example_transit_gateway.id,
    allowed_prefixes=["10.0.0.0/8"])
example_direct_connect_gateway_attachment = aws.ec2transitgateway.get_direct_connect_gateway_attachment_output(transit_gateway_id=example_transit_gateway.id,
    dx_gateway_id=example_gateway.id)
example_vpn_connection = aws.ec2.VpnConnection("exampleVpnConnection",
    customer_gateway_id=example_customer_gateway.id,
    outside_ip_address_type="PrivateIpv4",
    transit_gateway_id=example_transit_gateway.id,
    transport_transit_gateway_attachment_id=example_direct_connect_gateway_attachment.id,
    type="ipsec.1",
    tags={
        "Name": "example_ipsec_vpn_example",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleGateway = new Aws.DirectConnect.Gateway("exampleGateway", new()
    {
        AmazonSideAsn = "64512",
    });

    var exampleTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("exampleTransitGateway", new()
    {
        AmazonSideAsn = 64513,
        Description = "example_ipsec_vpn_example",
        TransitGatewayCidrBlocks = new[]
        {
            "10.0.0.0/24",
        },
    });

    var exampleCustomerGateway = new Aws.Ec2.CustomerGateway("exampleCustomerGateway", new()
    {
        BgpAsn = "64514",
        IpAddress = "10.0.0.1",
        Type = "ipsec.1",
        Tags = 
        {
            { "Name", "example_ipsec_vpn_example" },
        },
    });

    var exampleGatewayAssociation = new Aws.DirectConnect.GatewayAssociation("exampleGatewayAssociation", new()
    {
        DxGatewayId = exampleGateway.Id,
        AssociatedGatewayId = exampleTransitGateway.Id,
        AllowedPrefixes = new[]
        {
            "10.0.0.0/8",
        },
    });

    var exampleDirectConnectGatewayAttachment = Aws.Ec2TransitGateway.GetDirectConnectGatewayAttachment.Invoke(new()
    {
        TransitGatewayId = exampleTransitGateway.Id,
        DxGatewayId = exampleGateway.Id,
    });

    var exampleVpnConnection = new Aws.Ec2.VpnConnection("exampleVpnConnection", new()
    {
        CustomerGatewayId = exampleCustomerGateway.Id,
        OutsideIpAddressType = "PrivateIpv4",
        TransitGatewayId = exampleTransitGateway.Id,
        TransportTransitGatewayAttachmentId = exampleDirectConnectGatewayAttachment.Apply(getDirectConnectGatewayAttachmentResult => getDirectConnectGatewayAttachmentResult.Id),
        Type = "ipsec.1",
        Tags = 
        {
            { "Name", "example_ipsec_vpn_example" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleGateway, err := directconnect.NewGateway(ctx, "exampleGateway", &directconnect.GatewayArgs{
			AmazonSideAsn: pulumi.String("64512"),
		})
		if err != nil {
			return err
		}
		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", &ec2transitgateway.TransitGatewayArgs{
			AmazonSideAsn: pulumi.Int(64513),
			Description:   pulumi.String("example_ipsec_vpn_example"),
			TransitGatewayCidrBlocks: pulumi.StringArray{
				pulumi.String("10.0.0.0/24"),
			},
		})
		if err != nil {
			return err
		}
		exampleCustomerGateway, err := ec2.NewCustomerGateway(ctx, "exampleCustomerGateway", &ec2.CustomerGatewayArgs{
			BgpAsn:    pulumi.String("64514"),
			IpAddress: pulumi.String("10.0.0.1"),
			Type:      pulumi.String("ipsec.1"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example_ipsec_vpn_example"),
			},
		})
		if err != nil {
			return err
		}
		_, err = directconnect.NewGatewayAssociation(ctx, "exampleGatewayAssociation", &directconnect.GatewayAssociationArgs{
			DxGatewayId:         exampleGateway.ID(),
			AssociatedGatewayId: exampleTransitGateway.ID(),
			AllowedPrefixes: pulumi.StringArray{
				pulumi.String("10.0.0.0/8"),
			},
		})
		if err != nil {
			return err
		}
		exampleDirectConnectGatewayAttachment := ec2transitgateway.GetDirectConnectGatewayAttachmentOutput(ctx, ec2transitgateway.GetDirectConnectGatewayAttachmentOutputArgs{
			TransitGatewayId: exampleTransitGateway.ID(),
			DxGatewayId:      exampleGateway.ID(),
		}, nil)
		_, err = ec2.NewVpnConnection(ctx, "exampleVpnConnection", &ec2.VpnConnectionArgs{
			CustomerGatewayId:    exampleCustomerGateway.ID(),
			OutsideIpAddressType: pulumi.String("PrivateIpv4"),
			TransitGatewayId:     exampleTransitGateway.ID(),
			TransportTransitGatewayAttachmentId: exampleDirectConnectGatewayAttachment.ApplyT(func(exampleDirectConnectGatewayAttachment ec2transitgateway.GetDirectConnectGatewayAttachmentResult) (*string, error) {
				return &exampleDirectConnectGatewayAttachment.Id, nil
			}).(pulumi.StringPtrOutput),
			Type: pulumi.String("ipsec.1"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example_ipsec_vpn_example"),
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
import com.pulumi.aws.directconnect.Gateway;
import com.pulumi.aws.directconnect.GatewayArgs;
import com.pulumi.aws.ec2transitgateway.TransitGateway;
import com.pulumi.aws.ec2transitgateway.TransitGatewayArgs;
import com.pulumi.aws.ec2.CustomerGateway;
import com.pulumi.aws.ec2.CustomerGatewayArgs;
import com.pulumi.aws.directconnect.GatewayAssociation;
import com.pulumi.aws.directconnect.GatewayAssociationArgs;
import com.pulumi.aws.ec2transitgateway.Ec2transitgatewayFunctions;
import com.pulumi.aws.ec2transitgateway.inputs.GetDirectConnectGatewayAttachmentArgs;
import com.pulumi.aws.ec2.VpnConnection;
import com.pulumi.aws.ec2.VpnConnectionArgs;
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
        var exampleGateway = new Gateway("exampleGateway", GatewayArgs.builder()        
            .amazonSideAsn("64512")
            .build());

        var exampleTransitGateway = new TransitGateway("exampleTransitGateway", TransitGatewayArgs.builder()        
            .amazonSideAsn("64513")
            .description("example_ipsec_vpn_example")
            .transitGatewayCidrBlocks("10.0.0.0/24")
            .build());

        var exampleCustomerGateway = new CustomerGateway("exampleCustomerGateway", CustomerGatewayArgs.builder()        
            .bgpAsn(64514)
            .ipAddress("10.0.0.1")
            .type("ipsec.1")
            .tags(Map.of("Name", "example_ipsec_vpn_example"))
            .build());

        var exampleGatewayAssociation = new GatewayAssociation("exampleGatewayAssociation", GatewayAssociationArgs.builder()        
            .dxGatewayId(exampleGateway.id())
            .associatedGatewayId(exampleTransitGateway.id())
            .allowedPrefixes("10.0.0.0/8")
            .build());

        final var exampleDirectConnectGatewayAttachment = Ec2transitgatewayFunctions.getDirectConnectGatewayAttachment(GetDirectConnectGatewayAttachmentArgs.builder()
            .transitGatewayId(exampleTransitGateway.id())
            .dxGatewayId(exampleGateway.id())
            .build());

        var exampleVpnConnection = new VpnConnection("exampleVpnConnection", VpnConnectionArgs.builder()        
            .customerGatewayId(exampleCustomerGateway.id())
            .outsideIpAddressType("PrivateIpv4")
            .transitGatewayId(exampleTransitGateway.id())
            .transportTransitGatewayAttachmentId(exampleDirectConnectGatewayAttachment.applyValue(getDirectConnectGatewayAttachmentResult -> getDirectConnectGatewayAttachmentResult).applyValue(exampleDirectConnectGatewayAttachment -> exampleDirectConnectGatewayAttachment.applyValue(getDirectConnectGatewayAttachmentResult -> getDirectConnectGatewayAttachmentResult.id())))
            .type("ipsec.1")
            .tags(Map.of("Name", "example_ipsec_vpn_example"))
            .build());

    }
}
```
```yaml
resources:
  exampleGateway:
    type: aws:directconnect:Gateway
    properties:
      amazonSideAsn: '64512'
  exampleTransitGateway:
    type: aws:ec2transitgateway:TransitGateway
    properties:
      amazonSideAsn: '64513'
      description: example_ipsec_vpn_example
      transitGatewayCidrBlocks:
        - 10.0.0.0/24
  exampleCustomerGateway:
    type: aws:ec2:CustomerGateway
    properties:
      bgpAsn: 64514
      ipAddress: 10.0.0.1
      type: ipsec.1
      tags:
        Name: example_ipsec_vpn_example
  exampleGatewayAssociation:
    type: aws:directconnect:GatewayAssociation
    properties:
      dxGatewayId: ${exampleGateway.id}
      associatedGatewayId: ${exampleTransitGateway.id}
      allowedPrefixes:
        - 10.0.0.0/8
  exampleVpnConnection:
    type: aws:ec2:VpnConnection
    properties:
      customerGatewayId: ${exampleCustomerGateway.id}
      outsideIpAddressType: PrivateIpv4
      transitGatewayId: ${exampleTransitGateway.id}
      transportTransitGatewayAttachmentId: ${exampleDirectConnectGatewayAttachment.id}
      type: ipsec.1
      tags:
        Name: example_ipsec_vpn_example
variables:
  exampleDirectConnectGatewayAttachment:
    fn::invoke:
      Function: aws:ec2transitgateway:getDirectConnectGatewayAttachment
      Arguments:
        transitGatewayId: ${exampleTransitGateway.id}
        dxGatewayId: ${exampleGateway.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPN Connections using the VPN connection `id`. For example:

```sh
 $ pulumi import aws:ec2/vpnConnection:VpnConnection testvpnconnection vpn-40f41529
```
 