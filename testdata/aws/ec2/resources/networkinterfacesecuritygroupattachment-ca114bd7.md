This resource attaches a security group to an Elastic Network Interface (ENI).
It can be used to attach a security group to any existing ENI, be it a
secondary ENI or one attached as the primary interface on an instance.

> **NOTE on instances, interfaces, and security groups:** This provider currently
provides the capability to assign security groups via the [`aws.ec2.Instance`][1]
and the [`aws.ec2.NetworkInterface`][2] resources. Using this resource in
conjunction with security groups provided in-line in those resources will cause
conflicts, and will lead to spurious diffs and undefined behavior - please use
one or the other.

{{% examples %}}
## Example Usage
{{% example %}}

The following provides a very basic example of setting up an instance (provided
by `instance`) in the default security group, creating a security group
(provided by `sg`) and then attaching the security group to the instance's
primary network interface via the `aws.ec2.NetworkInterfaceSecurityGroupAttachment` resource,
named `sg_attachment`:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ami = aws.ec2.getAmi({
    mostRecent: true,
    filters: [{
        name: "name",
        values: ["amzn-ami-hvm-*"],
    }],
    owners: ["amazon"],
});
const instance = new aws.ec2.Instance("instance", {
    instanceType: "t2.micro",
    ami: ami.then(ami => ami.id),
    tags: {
        type: "test-instance",
    },
});
const sg = new aws.ec2.SecurityGroup("sg", {tags: {
    type: "test-security-group",
}});
const sgAttachment = new aws.ec2.NetworkInterfaceSecurityGroupAttachment("sgAttachment", {
    securityGroupId: sg.id,
    networkInterfaceId: instance.primaryNetworkInterfaceId,
});
```
```python
import pulumi
import pulumi_aws as aws

ami = aws.ec2.get_ami(most_recent=True,
    filters=[aws.ec2.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-hvm-*"],
    )],
    owners=["amazon"])
instance = aws.ec2.Instance("instance",
    instance_type="t2.micro",
    ami=ami.id,
    tags={
        "type": "test-instance",
    })
sg = aws.ec2.SecurityGroup("sg", tags={
    "type": "test-security-group",
})
sg_attachment = aws.ec2.NetworkInterfaceSecurityGroupAttachment("sgAttachment",
    security_group_id=sg.id,
    network_interface_id=instance.primary_network_interface_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ami = Aws.Ec2.GetAmi.Invoke(new()
    {
        MostRecent = true,
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "amzn-ami-hvm-*",
                },
            },
        },
        Owners = new[]
        {
            "amazon",
        },
    });

    var instance = new Aws.Ec2.Instance("instance", new()
    {
        InstanceType = "t2.micro",
        Ami = ami.Apply(getAmiResult => getAmiResult.Id),
        Tags = 
        {
            { "type", "test-instance" },
        },
    });

    var sg = new Aws.Ec2.SecurityGroup("sg", new()
    {
        Tags = 
        {
            { "type", "test-security-group" },
        },
    });

    var sgAttachment = new Aws.Ec2.NetworkInterfaceSecurityGroupAttachment("sgAttachment", new()
    {
        SecurityGroupId = sg.Id,
        NetworkInterfaceId = instance.PrimaryNetworkInterfaceId,
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
		ami, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"amzn-ami-hvm-*",
					},
				},
			},
			Owners: []string{
				"amazon",
			},
		}, nil)
		if err != nil {
			return err
		}
		instance, err := ec2.NewInstance(ctx, "instance", &ec2.InstanceArgs{
			InstanceType: pulumi.String("t2.micro"),
			Ami:          *pulumi.String(ami.Id),
			Tags: pulumi.StringMap{
				"type": pulumi.String("test-instance"),
			},
		})
		if err != nil {
			return err
		}
		sg, err := ec2.NewSecurityGroup(ctx, "sg", &ec2.SecurityGroupArgs{
			Tags: pulumi.StringMap{
				"type": pulumi.String("test-security-group"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewNetworkInterfaceSecurityGroupAttachment(ctx, "sgAttachment", &ec2.NetworkInterfaceSecurityGroupAttachmentArgs{
			SecurityGroupId:    sg.ID(),
			NetworkInterfaceId: instance.PrimaryNetworkInterfaceId,
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
import com.pulumi.aws.ec2.NetworkInterfaceSecurityGroupAttachment;
import com.pulumi.aws.ec2.NetworkInterfaceSecurityGroupAttachmentArgs;
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
        final var ami = Ec2Functions.getAmi(GetAmiArgs.builder()
            .mostRecent(true)
            .filters(GetAmiFilterArgs.builder()
                .name("name")
                .values("amzn-ami-hvm-*")
                .build())
            .owners("amazon")
            .build());

        var instance = new Instance("instance", InstanceArgs.builder()        
            .instanceType("t2.micro")
            .ami(ami.applyValue(getAmiResult -> getAmiResult.id()))
            .tags(Map.of("type", "test-instance"))
            .build());

        var sg = new SecurityGroup("sg", SecurityGroupArgs.builder()        
            .tags(Map.of("type", "test-security-group"))
            .build());

        var sgAttachment = new NetworkInterfaceSecurityGroupAttachment("sgAttachment", NetworkInterfaceSecurityGroupAttachmentArgs.builder()        
            .securityGroupId(sg.id())
            .networkInterfaceId(instance.primaryNetworkInterfaceId())
            .build());

    }
}
```
```yaml
resources:
  instance:
    type: aws:ec2:Instance
    properties:
      instanceType: t2.micro
      ami: ${ami.id}
      tags:
        type: test-instance
  sg:
    type: aws:ec2:SecurityGroup
    properties:
      tags:
        type: test-security-group
  sgAttachment:
    type: aws:ec2:NetworkInterfaceSecurityGroupAttachment
    properties:
      securityGroupId: ${sg.id}
      networkInterfaceId: ${instance.primaryNetworkInterfaceId}
variables:
  ami:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        mostRecent: true
        filters:
          - name: name
            values:
              - amzn-ami-hvm-*
        owners:
          - amazon
```

In this example, `instance` is provided by the `aws.ec2.Instance` data source,
fetching an external instance, possibly not managed by this provider.
`sg_attachment` then attaches to the output instance's `network_interface_id`:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const instance = aws.ec2.getInstance({
    instanceId: "i-1234567890abcdef0",
});
const sg = new aws.ec2.SecurityGroup("sg", {tags: {
    type: "test-security-group",
}});
const sgAttachment = new aws.ec2.NetworkInterfaceSecurityGroupAttachment("sgAttachment", {
    securityGroupId: sg.id,
    networkInterfaceId: instance.then(instance => instance.networkInterfaceId),
});
```
```python
import pulumi
import pulumi_aws as aws

instance = aws.ec2.get_instance(instance_id="i-1234567890abcdef0")
sg = aws.ec2.SecurityGroup("sg", tags={
    "type": "test-security-group",
})
sg_attachment = aws.ec2.NetworkInterfaceSecurityGroupAttachment("sgAttachment",
    security_group_id=sg.id,
    network_interface_id=instance.network_interface_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var instance = Aws.Ec2.GetInstance.Invoke(new()
    {
        InstanceId = "i-1234567890abcdef0",
    });

    var sg = new Aws.Ec2.SecurityGroup("sg", new()
    {
        Tags = 
        {
            { "type", "test-security-group" },
        },
    });

    var sgAttachment = new Aws.Ec2.NetworkInterfaceSecurityGroupAttachment("sgAttachment", new()
    {
        SecurityGroupId = sg.Id,
        NetworkInterfaceId = instance.Apply(getInstanceResult => getInstanceResult.NetworkInterfaceId),
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
		instance, err := ec2.LookupInstance(ctx, &ec2.LookupInstanceArgs{
			InstanceId: pulumi.StringRef("i-1234567890abcdef0"),
		}, nil)
		if err != nil {
			return err
		}
		sg, err := ec2.NewSecurityGroup(ctx, "sg", &ec2.SecurityGroupArgs{
			Tags: pulumi.StringMap{
				"type": pulumi.String("test-security-group"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewNetworkInterfaceSecurityGroupAttachment(ctx, "sgAttachment", &ec2.NetworkInterfaceSecurityGroupAttachmentArgs{
			SecurityGroupId:    sg.ID(),
			NetworkInterfaceId: *pulumi.String(instance.NetworkInterfaceId),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetInstanceArgs;
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
import com.pulumi.aws.ec2.NetworkInterfaceSecurityGroupAttachment;
import com.pulumi.aws.ec2.NetworkInterfaceSecurityGroupAttachmentArgs;
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
        final var instance = Ec2Functions.getInstance(GetInstanceArgs.builder()
            .instanceId("i-1234567890abcdef0")
            .build());

        var sg = new SecurityGroup("sg", SecurityGroupArgs.builder()        
            .tags(Map.of("type", "test-security-group"))
            .build());

        var sgAttachment = new NetworkInterfaceSecurityGroupAttachment("sgAttachment", NetworkInterfaceSecurityGroupAttachmentArgs.builder()        
            .securityGroupId(sg.id())
            .networkInterfaceId(instance.applyValue(getInstanceResult -> getInstanceResult.networkInterfaceId()))
            .build());

    }
}
```
```yaml
resources:
  sg:
    type: aws:ec2:SecurityGroup
    properties:
      tags:
        type: test-security-group
  sgAttachment:
    type: aws:ec2:NetworkInterfaceSecurityGroupAttachment
    properties:
      securityGroupId: ${sg.id}
      networkInterfaceId: ${instance.networkInterfaceId}
variables:
  instance:
    fn::invoke:
      Function: aws:ec2:getInstance
      Arguments:
        instanceId: i-1234567890abcdef0
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Network Interface Security Group attachments using the associated network interface ID and security group ID, separated by an underscore (`_`). For example:

```sh
 $ pulumi import aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment sg_attachment eni-1234567890abcdef0_sg-1234567890abcdef0
```
 