Manages an EC2 Transit Gateway Multicast Domain.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const available = aws.getAvailabilityZones({
    state: "available",
});
const amazonLinux = aws.ec2.getAmi({
    mostRecent: true,
    owners: ["amazon"],
    filters: [
        {
            name: "name",
            values: ["amzn-ami-hvm-*-x86_64-gp2"],
        },
        {
            name: "owner-alias",
            values: ["amazon"],
        },
    ],
});
const vpc1 = new aws.ec2.Vpc("vpc1", {cidrBlock: "10.0.0.0/16"});
const vpc2 = new aws.ec2.Vpc("vpc2", {cidrBlock: "10.1.0.0/16"});
const subnet1 = new aws.ec2.Subnet("subnet1", {
    vpcId: vpc1.id,
    cidrBlock: "10.0.1.0/24",
    availabilityZone: available.then(available => available.names?.[0]),
});
const subnet2 = new aws.ec2.Subnet("subnet2", {
    vpcId: vpc1.id,
    cidrBlock: "10.0.2.0/24",
    availabilityZone: available.then(available => available.names?.[1]),
});
const subnet3 = new aws.ec2.Subnet("subnet3", {
    vpcId: vpc2.id,
    cidrBlock: "10.1.1.0/24",
    availabilityZone: available.then(available => available.names?.[0]),
});
const instance1 = new aws.ec2.Instance("instance1", {
    ami: amazonLinux.then(amazonLinux => amazonLinux.id),
    instanceType: "t2.micro",
    subnetId: subnet1.id,
});
const instance2 = new aws.ec2.Instance("instance2", {
    ami: amazonLinux.then(amazonLinux => amazonLinux.id),
    instanceType: "t2.micro",
    subnetId: subnet2.id,
});
const instance3 = new aws.ec2.Instance("instance3", {
    ami: amazonLinux.then(amazonLinux => amazonLinux.id),
    instanceType: "t2.micro",
    subnetId: subnet3.id,
});
const tgw = new aws.ec2transitgateway.TransitGateway("tgw", {multicastSupport: "enable"});
const attachment1 = new aws.ec2transitgateway.VpcAttachment("attachment1", {
    subnetIds: [
        subnet1.id,
        subnet2.id,
    ],
    transitGatewayId: tgw.id,
    vpcId: vpc1.id,
});
const attachment2 = new aws.ec2transitgateway.VpcAttachment("attachment2", {
    subnetIds: [subnet3.id],
    transitGatewayId: tgw.id,
    vpcId: vpc2.id,
});
const domain = new aws.ec2transitgateway.MulticastDomain("domain", {
    transitGatewayId: tgw.id,
    staticSourcesSupport: "enable",
    tags: {
        Name: "Transit_Gateway_Multicast_Domain_Example",
    },
});
const association3 = new aws.ec2transitgateway.MulticastDomainAssociation("association3", {
    subnetId: subnet3.id,
    transitGatewayAttachmentId: attachment2.id,
    transitGatewayMulticastDomainId: domain.id,
});
const source = new aws.ec2transitgateway.MulticastGroupSource("source", {
    groupIpAddress: "224.0.0.1",
    networkInterfaceId: instance3.primaryNetworkInterfaceId,
    transitGatewayMulticastDomainId: association3.transitGatewayMulticastDomainId,
});
const association1 = new aws.ec2transitgateway.MulticastDomainAssociation("association1", {
    subnetId: subnet1.id,
    transitGatewayAttachmentId: attachment1.id,
    transitGatewayMulticastDomainId: domain.id,
});
const association2 = new aws.ec2transitgateway.MulticastDomainAssociation("association2", {
    subnetId: subnet2.id,
    transitGatewayAttachmentId: attachment2.id,
    transitGatewayMulticastDomainId: domain.id,
});
const member1 = new aws.ec2transitgateway.MulticastGroupMember("member1", {
    groupIpAddress: "224.0.0.1",
    networkInterfaceId: instance1.primaryNetworkInterfaceId,
    transitGatewayMulticastDomainId: association1.transitGatewayMulticastDomainId,
});
const member2 = new aws.ec2transitgateway.MulticastGroupMember("member2", {
    groupIpAddress: "224.0.0.1",
    networkInterfaceId: instance2.primaryNetworkInterfaceId,
    transitGatewayMulticastDomainId: association1.transitGatewayMulticastDomainId,
});
```
```python
import pulumi
import pulumi_aws as aws

available = aws.get_availability_zones(state="available")
amazon_linux = aws.ec2.get_ami(most_recent=True,
    owners=["amazon"],
    filters=[
        aws.ec2.GetAmiFilterArgs(
            name="name",
            values=["amzn-ami-hvm-*-x86_64-gp2"],
        ),
        aws.ec2.GetAmiFilterArgs(
            name="owner-alias",
            values=["amazon"],
        ),
    ])
vpc1 = aws.ec2.Vpc("vpc1", cidr_block="10.0.0.0/16")
vpc2 = aws.ec2.Vpc("vpc2", cidr_block="10.1.0.0/16")
subnet1 = aws.ec2.Subnet("subnet1",
    vpc_id=vpc1.id,
    cidr_block="10.0.1.0/24",
    availability_zone=available.names[0])
subnet2 = aws.ec2.Subnet("subnet2",
    vpc_id=vpc1.id,
    cidr_block="10.0.2.0/24",
    availability_zone=available.names[1])
subnet3 = aws.ec2.Subnet("subnet3",
    vpc_id=vpc2.id,
    cidr_block="10.1.1.0/24",
    availability_zone=available.names[0])
instance1 = aws.ec2.Instance("instance1",
    ami=amazon_linux.id,
    instance_type="t2.micro",
    subnet_id=subnet1.id)
instance2 = aws.ec2.Instance("instance2",
    ami=amazon_linux.id,
    instance_type="t2.micro",
    subnet_id=subnet2.id)
instance3 = aws.ec2.Instance("instance3",
    ami=amazon_linux.id,
    instance_type="t2.micro",
    subnet_id=subnet3.id)
tgw = aws.ec2transitgateway.TransitGateway("tgw", multicast_support="enable")
attachment1 = aws.ec2transitgateway.VpcAttachment("attachment1",
    subnet_ids=[
        subnet1.id,
        subnet2.id,
    ],
    transit_gateway_id=tgw.id,
    vpc_id=vpc1.id)
attachment2 = aws.ec2transitgateway.VpcAttachment("attachment2",
    subnet_ids=[subnet3.id],
    transit_gateway_id=tgw.id,
    vpc_id=vpc2.id)
domain = aws.ec2transitgateway.MulticastDomain("domain",
    transit_gateway_id=tgw.id,
    static_sources_support="enable",
    tags={
        "Name": "Transit_Gateway_Multicast_Domain_Example",
    })
association3 = aws.ec2transitgateway.MulticastDomainAssociation("association3",
    subnet_id=subnet3.id,
    transit_gateway_attachment_id=attachment2.id,
    transit_gateway_multicast_domain_id=domain.id)
source = aws.ec2transitgateway.MulticastGroupSource("source",
    group_ip_address="224.0.0.1",
    network_interface_id=instance3.primary_network_interface_id,
    transit_gateway_multicast_domain_id=association3.transit_gateway_multicast_domain_id)
association1 = aws.ec2transitgateway.MulticastDomainAssociation("association1",
    subnet_id=subnet1.id,
    transit_gateway_attachment_id=attachment1.id,
    transit_gateway_multicast_domain_id=domain.id)
association2 = aws.ec2transitgateway.MulticastDomainAssociation("association2",
    subnet_id=subnet2.id,
    transit_gateway_attachment_id=attachment2.id,
    transit_gateway_multicast_domain_id=domain.id)
member1 = aws.ec2transitgateway.MulticastGroupMember("member1",
    group_ip_address="224.0.0.1",
    network_interface_id=instance1.primary_network_interface_id,
    transit_gateway_multicast_domain_id=association1.transit_gateway_multicast_domain_id)
member2 = aws.ec2transitgateway.MulticastGroupMember("member2",
    group_ip_address="224.0.0.1",
    network_interface_id=instance2.primary_network_interface_id,
    transit_gateway_multicast_domain_id=association1.transit_gateway_multicast_domain_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var available = Aws.GetAvailabilityZones.Invoke(new()
    {
        State = "available",
    });

    var amazonLinux = Aws.Ec2.GetAmi.Invoke(new()
    {
        MostRecent = true,
        Owners = new[]
        {
            "amazon",
        },
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "amzn-ami-hvm-*-x86_64-gp2",
                },
            },
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "owner-alias",
                Values = new[]
                {
                    "amazon",
                },
            },
        },
    });

    var vpc1 = new Aws.Ec2.Vpc("vpc1", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var vpc2 = new Aws.Ec2.Vpc("vpc2", new()
    {
        CidrBlock = "10.1.0.0/16",
    });

    var subnet1 = new Aws.Ec2.Subnet("subnet1", new()
    {
        VpcId = vpc1.Id,
        CidrBlock = "10.0.1.0/24",
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
    });

    var subnet2 = new Aws.Ec2.Subnet("subnet2", new()
    {
        VpcId = vpc1.Id,
        CidrBlock = "10.0.2.0/24",
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[1]),
    });

    var subnet3 = new Aws.Ec2.Subnet("subnet3", new()
    {
        VpcId = vpc2.Id,
        CidrBlock = "10.1.1.0/24",
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
    });

    var instance1 = new Aws.Ec2.Instance("instance1", new()
    {
        Ami = amazonLinux.Apply(getAmiResult => getAmiResult.Id),
        InstanceType = "t2.micro",
        SubnetId = subnet1.Id,
    });

    var instance2 = new Aws.Ec2.Instance("instance2", new()
    {
        Ami = amazonLinux.Apply(getAmiResult => getAmiResult.Id),
        InstanceType = "t2.micro",
        SubnetId = subnet2.Id,
    });

    var instance3 = new Aws.Ec2.Instance("instance3", new()
    {
        Ami = amazonLinux.Apply(getAmiResult => getAmiResult.Id),
        InstanceType = "t2.micro",
        SubnetId = subnet3.Id,
    });

    var tgw = new Aws.Ec2TransitGateway.TransitGateway("tgw", new()
    {
        MulticastSupport = "enable",
    });

    var attachment1 = new Aws.Ec2TransitGateway.VpcAttachment("attachment1", new()
    {
        SubnetIds = new[]
        {
            subnet1.Id,
            subnet2.Id,
        },
        TransitGatewayId = tgw.Id,
        VpcId = vpc1.Id,
    });

    var attachment2 = new Aws.Ec2TransitGateway.VpcAttachment("attachment2", new()
    {
        SubnetIds = new[]
        {
            subnet3.Id,
        },
        TransitGatewayId = tgw.Id,
        VpcId = vpc2.Id,
    });

    var domain = new Aws.Ec2TransitGateway.MulticastDomain("domain", new()
    {
        TransitGatewayId = tgw.Id,
        StaticSourcesSupport = "enable",
        Tags = 
        {
            { "Name", "Transit_Gateway_Multicast_Domain_Example" },
        },
    });

    var association3 = new Aws.Ec2TransitGateway.MulticastDomainAssociation("association3", new()
    {
        SubnetId = subnet3.Id,
        TransitGatewayAttachmentId = attachment2.Id,
        TransitGatewayMulticastDomainId = domain.Id,
    });

    var source = new Aws.Ec2TransitGateway.MulticastGroupSource("source", new()
    {
        GroupIpAddress = "224.0.0.1",
        NetworkInterfaceId = instance3.PrimaryNetworkInterfaceId,
        TransitGatewayMulticastDomainId = association3.TransitGatewayMulticastDomainId,
    });

    var association1 = new Aws.Ec2TransitGateway.MulticastDomainAssociation("association1", new()
    {
        SubnetId = subnet1.Id,
        TransitGatewayAttachmentId = attachment1.Id,
        TransitGatewayMulticastDomainId = domain.Id,
    });

    var association2 = new Aws.Ec2TransitGateway.MulticastDomainAssociation("association2", new()
    {
        SubnetId = subnet2.Id,
        TransitGatewayAttachmentId = attachment2.Id,
        TransitGatewayMulticastDomainId = domain.Id,
    });

    var member1 = new Aws.Ec2TransitGateway.MulticastGroupMember("member1", new()
    {
        GroupIpAddress = "224.0.0.1",
        NetworkInterfaceId = instance1.PrimaryNetworkInterfaceId,
        TransitGatewayMulticastDomainId = association1.TransitGatewayMulticastDomainId,
    });

    var member2 = new Aws.Ec2TransitGateway.MulticastGroupMember("member2", new()
    {
        GroupIpAddress = "224.0.0.1",
        NetworkInterfaceId = instance2.PrimaryNetworkInterfaceId,
        TransitGatewayMulticastDomainId = association1.TransitGatewayMulticastDomainId,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		available, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			State: pulumi.StringRef("available"),
		}, nil)
		if err != nil {
			return err
		}
		amazonLinux, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Owners: []string{
				"amazon",
			},
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"amzn-ami-hvm-*-x86_64-gp2",
					},
				},
				{
					Name: "owner-alias",
					Values: []string{
						"amazon",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		vpc1, err := ec2.NewVpc(ctx, "vpc1", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		vpc2, err := ec2.NewVpc(ctx, "vpc2", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		subnet1, err := ec2.NewSubnet(ctx, "subnet1", &ec2.SubnetArgs{
			VpcId:            vpc1.ID(),
			CidrBlock:        pulumi.String("10.0.1.0/24"),
			AvailabilityZone: *pulumi.String(available.Names[0]),
		})
		if err != nil {
			return err
		}
		subnet2, err := ec2.NewSubnet(ctx, "subnet2", &ec2.SubnetArgs{
			VpcId:            vpc1.ID(),
			CidrBlock:        pulumi.String("10.0.2.0/24"),
			AvailabilityZone: *pulumi.String(available.Names[1]),
		})
		if err != nil {
			return err
		}
		subnet3, err := ec2.NewSubnet(ctx, "subnet3", &ec2.SubnetArgs{
			VpcId:            vpc2.ID(),
			CidrBlock:        pulumi.String("10.1.1.0/24"),
			AvailabilityZone: *pulumi.String(available.Names[0]),
		})
		if err != nil {
			return err
		}
		instance1, err := ec2.NewInstance(ctx, "instance1", &ec2.InstanceArgs{
			Ami:          *pulumi.String(amazonLinux.Id),
			InstanceType: pulumi.String("t2.micro"),
			SubnetId:     subnet1.ID(),
		})
		if err != nil {
			return err
		}
		instance2, err := ec2.NewInstance(ctx, "instance2", &ec2.InstanceArgs{
			Ami:          *pulumi.String(amazonLinux.Id),
			InstanceType: pulumi.String("t2.micro"),
			SubnetId:     subnet2.ID(),
		})
		if err != nil {
			return err
		}
		instance3, err := ec2.NewInstance(ctx, "instance3", &ec2.InstanceArgs{
			Ami:          *pulumi.String(amazonLinux.Id),
			InstanceType: pulumi.String("t2.micro"),
			SubnetId:     subnet3.ID(),
		})
		if err != nil {
			return err
		}
		tgw, err := ec2transitgateway.NewTransitGateway(ctx, "tgw", &ec2transitgateway.TransitGatewayArgs{
			MulticastSupport: pulumi.String("enable"),
		})
		if err != nil {
			return err
		}
		attachment1, err := ec2transitgateway.NewVpcAttachment(ctx, "attachment1", &ec2transitgateway.VpcAttachmentArgs{
			SubnetIds: pulumi.StringArray{
				subnet1.ID(),
				subnet2.ID(),
			},
			TransitGatewayId: tgw.ID(),
			VpcId:            vpc1.ID(),
		})
		if err != nil {
			return err
		}
		attachment2, err := ec2transitgateway.NewVpcAttachment(ctx, "attachment2", &ec2transitgateway.VpcAttachmentArgs{
			SubnetIds: pulumi.StringArray{
				subnet3.ID(),
			},
			TransitGatewayId: tgw.ID(),
			VpcId:            vpc2.ID(),
		})
		if err != nil {
			return err
		}
		domain, err := ec2transitgateway.NewMulticastDomain(ctx, "domain", &ec2transitgateway.MulticastDomainArgs{
			TransitGatewayId:     tgw.ID(),
			StaticSourcesSupport: pulumi.String("enable"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Transit_Gateway_Multicast_Domain_Example"),
			},
		})
		if err != nil {
			return err
		}
		association3, err := ec2transitgateway.NewMulticastDomainAssociation(ctx, "association3", &ec2transitgateway.MulticastDomainAssociationArgs{
			SubnetId:                        subnet3.ID(),
			TransitGatewayAttachmentId:      attachment2.ID(),
			TransitGatewayMulticastDomainId: domain.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2transitgateway.NewMulticastGroupSource(ctx, "source", &ec2transitgateway.MulticastGroupSourceArgs{
			GroupIpAddress:                  pulumi.String("224.0.0.1"),
			NetworkInterfaceId:              instance3.PrimaryNetworkInterfaceId,
			TransitGatewayMulticastDomainId: association3.TransitGatewayMulticastDomainId,
		})
		if err != nil {
			return err
		}
		association1, err := ec2transitgateway.NewMulticastDomainAssociation(ctx, "association1", &ec2transitgateway.MulticastDomainAssociationArgs{
			SubnetId:                        subnet1.ID(),
			TransitGatewayAttachmentId:      attachment1.ID(),
			TransitGatewayMulticastDomainId: domain.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2transitgateway.NewMulticastDomainAssociation(ctx, "association2", &ec2transitgateway.MulticastDomainAssociationArgs{
			SubnetId:                        subnet2.ID(),
			TransitGatewayAttachmentId:      attachment2.ID(),
			TransitGatewayMulticastDomainId: domain.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2transitgateway.NewMulticastGroupMember(ctx, "member1", &ec2transitgateway.MulticastGroupMemberArgs{
			GroupIpAddress:                  pulumi.String("224.0.0.1"),
			NetworkInterfaceId:              instance1.PrimaryNetworkInterfaceId,
			TransitGatewayMulticastDomainId: association1.TransitGatewayMulticastDomainId,
		})
		if err != nil {
			return err
		}
		_, err = ec2transitgateway.NewMulticastGroupMember(ctx, "member2", &ec2transitgateway.MulticastGroupMemberArgs{
			GroupIpAddress:                  pulumi.String("224.0.0.1"),
			NetworkInterfaceId:              instance2.PrimaryNetworkInterfaceId,
			TransitGatewayMulticastDomainId: association1.TransitGatewayMulticastDomainId,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import com.pulumi.aws.ec2transitgateway.TransitGateway;
import com.pulumi.aws.ec2transitgateway.TransitGatewayArgs;
import com.pulumi.aws.ec2transitgateway.VpcAttachment;
import com.pulumi.aws.ec2transitgateway.VpcAttachmentArgs;
import com.pulumi.aws.ec2transitgateway.MulticastDomain;
import com.pulumi.aws.ec2transitgateway.MulticastDomainArgs;
import com.pulumi.aws.ec2transitgateway.MulticastDomainAssociation;
import com.pulumi.aws.ec2transitgateway.MulticastDomainAssociationArgs;
import com.pulumi.aws.ec2transitgateway.MulticastGroupSource;
import com.pulumi.aws.ec2transitgateway.MulticastGroupSourceArgs;
import com.pulumi.aws.ec2transitgateway.MulticastGroupMember;
import com.pulumi.aws.ec2transitgateway.MulticastGroupMemberArgs;
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
        final var available = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
            .state("available")
            .build());

        final var amazonLinux = Ec2Functions.getAmi(GetAmiArgs.builder()
            .mostRecent(true)
            .owners("amazon")
            .filters(            
                GetAmiFilterArgs.builder()
                    .name("name")
                    .values("amzn-ami-hvm-*-x86_64-gp2")
                    .build(),
                GetAmiFilterArgs.builder()
                    .name("owner-alias")
                    .values("amazon")
                    .build())
            .build());

        var vpc1 = new Vpc("vpc1", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var vpc2 = new Vpc("vpc2", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .build());

        var subnet1 = new Subnet("subnet1", SubnetArgs.builder()        
            .vpcId(vpc1.id())
            .cidrBlock("10.0.1.0/24")
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]))
            .build());

        var subnet2 = new Subnet("subnet2", SubnetArgs.builder()        
            .vpcId(vpc1.id())
            .cidrBlock("10.0.2.0/24")
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[1]))
            .build());

        var subnet3 = new Subnet("subnet3", SubnetArgs.builder()        
            .vpcId(vpc2.id())
            .cidrBlock("10.1.1.0/24")
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]))
            .build());

        var instance1 = new Instance("instance1", InstanceArgs.builder()        
            .ami(amazonLinux.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceType("t2.micro")
            .subnetId(subnet1.id())
            .build());

        var instance2 = new Instance("instance2", InstanceArgs.builder()        
            .ami(amazonLinux.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceType("t2.micro")
            .subnetId(subnet2.id())
            .build());

        var instance3 = new Instance("instance3", InstanceArgs.builder()        
            .ami(amazonLinux.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceType("t2.micro")
            .subnetId(subnet3.id())
            .build());

        var tgw = new TransitGateway("tgw", TransitGatewayArgs.builder()        
            .multicastSupport("enable")
            .build());

        var attachment1 = new VpcAttachment("attachment1", VpcAttachmentArgs.builder()        
            .subnetIds(            
                subnet1.id(),
                subnet2.id())
            .transitGatewayId(tgw.id())
            .vpcId(vpc1.id())
            .build());

        var attachment2 = new VpcAttachment("attachment2", VpcAttachmentArgs.builder()        
            .subnetIds(subnet3.id())
            .transitGatewayId(tgw.id())
            .vpcId(vpc2.id())
            .build());

        var domain = new MulticastDomain("domain", MulticastDomainArgs.builder()        
            .transitGatewayId(tgw.id())
            .staticSourcesSupport("enable")
            .tags(Map.of("Name", "Transit_Gateway_Multicast_Domain_Example"))
            .build());

        var association3 = new MulticastDomainAssociation("association3", MulticastDomainAssociationArgs.builder()        
            .subnetId(subnet3.id())
            .transitGatewayAttachmentId(attachment2.id())
            .transitGatewayMulticastDomainId(domain.id())
            .build());

        var source = new MulticastGroupSource("source", MulticastGroupSourceArgs.builder()        
            .groupIpAddress("224.0.0.1")
            .networkInterfaceId(instance3.primaryNetworkInterfaceId())
            .transitGatewayMulticastDomainId(association3.transitGatewayMulticastDomainId())
            .build());

        var association1 = new MulticastDomainAssociation("association1", MulticastDomainAssociationArgs.builder()        
            .subnetId(subnet1.id())
            .transitGatewayAttachmentId(attachment1.id())
            .transitGatewayMulticastDomainId(domain.id())
            .build());

        var association2 = new MulticastDomainAssociation("association2", MulticastDomainAssociationArgs.builder()        
            .subnetId(subnet2.id())
            .transitGatewayAttachmentId(attachment2.id())
            .transitGatewayMulticastDomainId(domain.id())
            .build());

        var member1 = new MulticastGroupMember("member1", MulticastGroupMemberArgs.builder()        
            .groupIpAddress("224.0.0.1")
            .networkInterfaceId(instance1.primaryNetworkInterfaceId())
            .transitGatewayMulticastDomainId(association1.transitGatewayMulticastDomainId())
            .build());

        var member2 = new MulticastGroupMember("member2", MulticastGroupMemberArgs.builder()        
            .groupIpAddress("224.0.0.1")
            .networkInterfaceId(instance2.primaryNetworkInterfaceId())
            .transitGatewayMulticastDomainId(association1.transitGatewayMulticastDomainId())
            .build());

    }
}
```
```yaml
resources:
  vpc1:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  vpc2:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
  subnet1:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${vpc1.id}
      cidrBlock: 10.0.1.0/24
      availabilityZone: ${available.names[0]}
  subnet2:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${vpc1.id}
      cidrBlock: 10.0.2.0/24
      availabilityZone: ${available.names[1]}
  subnet3:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${vpc2.id}
      cidrBlock: 10.1.1.0/24
      availabilityZone: ${available.names[0]}
  instance1:
    type: aws:ec2:Instance
    properties:
      ami: ${amazonLinux.id}
      instanceType: t2.micro
      subnetId: ${subnet1.id}
  instance2:
    type: aws:ec2:Instance
    properties:
      ami: ${amazonLinux.id}
      instanceType: t2.micro
      subnetId: ${subnet2.id}
  instance3:
    type: aws:ec2:Instance
    properties:
      ami: ${amazonLinux.id}
      instanceType: t2.micro
      subnetId: ${subnet3.id}
  tgw:
    type: aws:ec2transitgateway:TransitGateway
    properties:
      multicastSupport: enable
  attachment1:
    type: aws:ec2transitgateway:VpcAttachment
    properties:
      subnetIds:
        - ${subnet1.id}
        - ${subnet2.id}
      transitGatewayId: ${tgw.id}
      vpcId: ${vpc1.id}
  attachment2:
    type: aws:ec2transitgateway:VpcAttachment
    properties:
      subnetIds:
        - ${subnet3.id}
      transitGatewayId: ${tgw.id}
      vpcId: ${vpc2.id}
  domain:
    type: aws:ec2transitgateway:MulticastDomain
    properties:
      transitGatewayId: ${tgw.id}
      staticSourcesSupport: enable
      tags:
        Name: Transit_Gateway_Multicast_Domain_Example
  association3:
    type: aws:ec2transitgateway:MulticastDomainAssociation
    properties:
      subnetId: ${subnet3.id}
      transitGatewayAttachmentId: ${attachment2.id}
      transitGatewayMulticastDomainId: ${domain.id}
  source:
    type: aws:ec2transitgateway:MulticastGroupSource
    properties:
      groupIpAddress: 224.0.0.1
      networkInterfaceId: ${instance3.primaryNetworkInterfaceId}
      transitGatewayMulticastDomainId: ${association3.transitGatewayMulticastDomainId}
  association1:
    type: aws:ec2transitgateway:MulticastDomainAssociation
    properties:
      subnetId: ${subnet1.id}
      transitGatewayAttachmentId: ${attachment1.id}
      transitGatewayMulticastDomainId: ${domain.id}
  association2:
    type: aws:ec2transitgateway:MulticastDomainAssociation
    properties:
      subnetId: ${subnet2.id}
      transitGatewayAttachmentId: ${attachment2.id}
      transitGatewayMulticastDomainId: ${domain.id}
  member1:
    type: aws:ec2transitgateway:MulticastGroupMember
    properties:
      groupIpAddress: 224.0.0.1
      networkInterfaceId: ${instance1.primaryNetworkInterfaceId}
      transitGatewayMulticastDomainId: ${association1.transitGatewayMulticastDomainId}
  member2:
    type: aws:ec2transitgateway:MulticastGroupMember
    properties:
      groupIpAddress: 224.0.0.1
      networkInterfaceId: ${instance2.primaryNetworkInterfaceId}
      transitGatewayMulticastDomainId: ${association1.transitGatewayMulticastDomainId}
variables:
  available:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments:
        state: available
  amazonLinux:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        mostRecent: true
        owners:
          - amazon
        filters:
          - name: name
            values:
              - amzn-ami-hvm-*-x86_64-gp2
          - name: owner-alias
            values:
              - amazon
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_transit_gateway_multicast_domain` using the EC2 Transit Gateway Multicast Domain identifier. For example:

```sh
 $ pulumi import aws:ec2transitgateway/multicastDomain:MulticastDomain example tgw-mcast-domain-12345
```
 