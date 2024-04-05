Associates a Direct Connect Gateway with a VGW or transit gateway.

To create a cross-account association, create an `aws.directconnect.GatewayAssociationProposal` resource
in the AWS account that owns the VGW or transit gateway and then accept the proposal in the AWS account that owns the Direct Connect Gateway
by creating an `aws.directconnect.GatewayAssociation` resource with the `proposal_id` and `associated_gateway_owner_account_id` attributes set.

{{% examples %}}
## Example Usage
{{% example %}}
### VPN Gateway Association

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGateway = new aws.directconnect.Gateway("exampleGateway", {amazonSideAsn: "64512"});
const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.255.255.0/28"});
const exampleVpnGateway = new aws.ec2.VpnGateway("exampleVpnGateway", {vpcId: exampleVpc.id});
const exampleGatewayAssociation = new aws.directconnect.GatewayAssociation("exampleGatewayAssociation", {
    dxGatewayId: exampleGateway.id,
    associatedGatewayId: exampleVpnGateway.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn="64512")
example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.255.255.0/28")
example_vpn_gateway = aws.ec2.VpnGateway("exampleVpnGateway", vpc_id=example_vpc.id)
example_gateway_association = aws.directconnect.GatewayAssociation("exampleGatewayAssociation",
    dx_gateway_id=example_gateway.id,
    associated_gateway_id=example_vpn_gateway.id)
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

    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.255.255.0/28",
    });

    var exampleVpnGateway = new Aws.Ec2.VpnGateway("exampleVpnGateway", new()
    {
        VpcId = exampleVpc.Id,
    });

    var exampleGatewayAssociation = new Aws.DirectConnect.GatewayAssociation("exampleGatewayAssociation", new()
    {
        DxGatewayId = exampleGateway.Id,
        AssociatedGatewayId = exampleVpnGateway.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
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
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.255.255.0/28"),
		})
		if err != nil {
			return err
		}
		exampleVpnGateway, err := ec2.NewVpnGateway(ctx, "exampleVpnGateway", &ec2.VpnGatewayArgs{
			VpcId: exampleVpc.ID(),
		})
		if err != nil {
			return err
		}
		_, err = directconnect.NewGatewayAssociation(ctx, "exampleGatewayAssociation", &directconnect.GatewayAssociationArgs{
			DxGatewayId:         exampleGateway.ID(),
			AssociatedGatewayId: exampleVpnGateway.ID(),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.VpnGateway;
import com.pulumi.aws.ec2.VpnGatewayArgs;
import com.pulumi.aws.directconnect.GatewayAssociation;
import com.pulumi.aws.directconnect.GatewayAssociationArgs;
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

        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.255.255.0/28")
            .build());

        var exampleVpnGateway = new VpnGateway("exampleVpnGateway", VpnGatewayArgs.builder()        
            .vpcId(exampleVpc.id())
            .build());

        var exampleGatewayAssociation = new GatewayAssociation("exampleGatewayAssociation", GatewayAssociationArgs.builder()        
            .dxGatewayId(exampleGateway.id())
            .associatedGatewayId(exampleVpnGateway.id())
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
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.255.255.0/28
  exampleVpnGateway:
    type: aws:ec2:VpnGateway
    properties:
      vpcId: ${exampleVpc.id}
  exampleGatewayAssociation:
    type: aws:directconnect:GatewayAssociation
    properties:
      dxGatewayId: ${exampleGateway.id}
      associatedGatewayId: ${exampleVpnGateway.id}
```
{{% /example %}}
{{% example %}}
### Transit Gateway Association

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGateway = new aws.directconnect.Gateway("exampleGateway", {amazonSideAsn: "64512"});
const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("exampleTransitGateway", {});
const exampleGatewayAssociation = new aws.directconnect.GatewayAssociation("exampleGatewayAssociation", {
    dxGatewayId: exampleGateway.id,
    associatedGatewayId: exampleTransitGateway.id,
    allowedPrefixes: [
        "10.255.255.0/30",
        "10.255.255.8/30",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn="64512")
example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway")
example_gateway_association = aws.directconnect.GatewayAssociation("exampleGatewayAssociation",
    dx_gateway_id=example_gateway.id,
    associated_gateway_id=example_transit_gateway.id,
    allowed_prefixes=[
        "10.255.255.0/30",
        "10.255.255.8/30",
    ])
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

    var exampleTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("exampleTransitGateway");

    var exampleGatewayAssociation = new Aws.DirectConnect.GatewayAssociation("exampleGatewayAssociation", new()
    {
        DxGatewayId = exampleGateway.Id,
        AssociatedGatewayId = exampleTransitGateway.Id,
        AllowedPrefixes = new[]
        {
            "10.255.255.0/30",
            "10.255.255.8/30",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
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
		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", nil)
		if err != nil {
			return err
		}
		_, err = directconnect.NewGatewayAssociation(ctx, "exampleGatewayAssociation", &directconnect.GatewayAssociationArgs{
			DxGatewayId:         exampleGateway.ID(),
			AssociatedGatewayId: exampleTransitGateway.ID(),
			AllowedPrefixes: pulumi.StringArray{
				pulumi.String("10.255.255.0/30"),
				pulumi.String("10.255.255.8/30"),
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
import com.pulumi.aws.directconnect.GatewayAssociation;
import com.pulumi.aws.directconnect.GatewayAssociationArgs;
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

        var exampleTransitGateway = new TransitGateway("exampleTransitGateway");

        var exampleGatewayAssociation = new GatewayAssociation("exampleGatewayAssociation", GatewayAssociationArgs.builder()        
            .dxGatewayId(exampleGateway.id())
            .associatedGatewayId(exampleTransitGateway.id())
            .allowedPrefixes(            
                "10.255.255.0/30",
                "10.255.255.8/30")
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
  exampleGatewayAssociation:
    type: aws:directconnect:GatewayAssociation
    properties:
      dxGatewayId: ${exampleGateway.id}
      associatedGatewayId: ${exampleTransitGateway.id}
      allowedPrefixes:
        - 10.255.255.0/30
        - 10.255.255.8/30
```
{{% /example %}}
{{% example %}}
### Allowed Prefixes

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGateway = new aws.directconnect.Gateway("exampleGateway", {amazonSideAsn: "64512"});
const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.255.255.0/28"});
const exampleVpnGateway = new aws.ec2.VpnGateway("exampleVpnGateway", {vpcId: exampleVpc.id});
const exampleGatewayAssociation = new aws.directconnect.GatewayAssociation("exampleGatewayAssociation", {
    dxGatewayId: exampleGateway.id,
    associatedGatewayId: exampleVpnGateway.id,
    allowedPrefixes: [
        "210.52.109.0/24",
        "175.45.176.0/22",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn="64512")
example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.255.255.0/28")
example_vpn_gateway = aws.ec2.VpnGateway("exampleVpnGateway", vpc_id=example_vpc.id)
example_gateway_association = aws.directconnect.GatewayAssociation("exampleGatewayAssociation",
    dx_gateway_id=example_gateway.id,
    associated_gateway_id=example_vpn_gateway.id,
    allowed_prefixes=[
        "210.52.109.0/24",
        "175.45.176.0/22",
    ])
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

    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.255.255.0/28",
    });

    var exampleVpnGateway = new Aws.Ec2.VpnGateway("exampleVpnGateway", new()
    {
        VpcId = exampleVpc.Id,
    });

    var exampleGatewayAssociation = new Aws.DirectConnect.GatewayAssociation("exampleGatewayAssociation", new()
    {
        DxGatewayId = exampleGateway.Id,
        AssociatedGatewayId = exampleVpnGateway.Id,
        AllowedPrefixes = new[]
        {
            "210.52.109.0/24",
            "175.45.176.0/22",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
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
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.255.255.0/28"),
		})
		if err != nil {
			return err
		}
		exampleVpnGateway, err := ec2.NewVpnGateway(ctx, "exampleVpnGateway", &ec2.VpnGatewayArgs{
			VpcId: exampleVpc.ID(),
		})
		if err != nil {
			return err
		}
		_, err = directconnect.NewGatewayAssociation(ctx, "exampleGatewayAssociation", &directconnect.GatewayAssociationArgs{
			DxGatewayId:         exampleGateway.ID(),
			AssociatedGatewayId: exampleVpnGateway.ID(),
			AllowedPrefixes: pulumi.StringArray{
				pulumi.String("210.52.109.0/24"),
				pulumi.String("175.45.176.0/22"),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.VpnGateway;
import com.pulumi.aws.ec2.VpnGatewayArgs;
import com.pulumi.aws.directconnect.GatewayAssociation;
import com.pulumi.aws.directconnect.GatewayAssociationArgs;
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

        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.255.255.0/28")
            .build());

        var exampleVpnGateway = new VpnGateway("exampleVpnGateway", VpnGatewayArgs.builder()        
            .vpcId(exampleVpc.id())
            .build());

        var exampleGatewayAssociation = new GatewayAssociation("exampleGatewayAssociation", GatewayAssociationArgs.builder()        
            .dxGatewayId(exampleGateway.id())
            .associatedGatewayId(exampleVpnGateway.id())
            .allowedPrefixes(            
                "210.52.109.0/24",
                "175.45.176.0/22")
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
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.255.255.0/28
  exampleVpnGateway:
    type: aws:ec2:VpnGateway
    properties:
      vpcId: ${exampleVpc.id}
  exampleGatewayAssociation:
    type: aws:directconnect:GatewayAssociation
    properties:
      dxGatewayId: ${exampleGateway.id}
      associatedGatewayId: ${exampleVpnGateway.id}
      allowedPrefixes:
        - 210.52.109.0/24
        - 175.45.176.0/22
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Direct Connect gateway associations using `dx_gateway_id` together with `associated_gateway_id`. For example:

```sh
 $ pulumi import aws:directconnect/gatewayAssociation:GatewayAssociation example 345508c3-7215-4aef-9832-07c125d5bd0f/vgw-98765432
```
 