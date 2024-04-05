Provides an AWS EIP Association as a top level resource, to associate and
disassociate Elastic IPs from AWS Instances and Network Interfaces.

> **NOTE:** Do not use this resource to associate an EIP to `aws.lb.LoadBalancer` or `aws.ec2.NatGateway` resources. Instead use the `allocation_id` available in those resources to allow AWS to manage the association, otherwise you will see `AuthFailure` errors.

> **NOTE:** `aws.ec2.EipAssociation` is useful in scenarios where EIPs are either
pre-existing or distributed to customers or users and therefore cannot be changed.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const web = new aws.ec2.Instance("web", {
    ami: "ami-21f78e11",
    availabilityZone: "us-west-2a",
    instanceType: "t2.micro",
    tags: {
        Name: "HelloWorld",
    },
});
const example = new aws.ec2.Eip("example", {domain: "vpc"});
const eipAssoc = new aws.ec2.EipAssociation("eipAssoc", {
    instanceId: web.id,
    allocationId: example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

web = aws.ec2.Instance("web",
    ami="ami-21f78e11",
    availability_zone="us-west-2a",
    instance_type="t2.micro",
    tags={
        "Name": "HelloWorld",
    })
example = aws.ec2.Eip("example", domain="vpc")
eip_assoc = aws.ec2.EipAssociation("eipAssoc",
    instance_id=web.id,
    allocation_id=example.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var web = new Aws.Ec2.Instance("web", new()
    {
        Ami = "ami-21f78e11",
        AvailabilityZone = "us-west-2a",
        InstanceType = "t2.micro",
        Tags = 
        {
            { "Name", "HelloWorld" },
        },
    });

    var example = new Aws.Ec2.Eip("example", new()
    {
        Domain = "vpc",
    });

    var eipAssoc = new Aws.Ec2.EipAssociation("eipAssoc", new()
    {
        InstanceId = web.Id,
        AllocationId = example.Id,
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
		web, err := ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
			Ami:              pulumi.String("ami-21f78e11"),
			AvailabilityZone: pulumi.String("us-west-2a"),
			InstanceType:     pulumi.String("t2.micro"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
			},
		})
		if err != nil {
			return err
		}
		example, err := ec2.NewEip(ctx, "example", &ec2.EipArgs{
			Domain: pulumi.String("vpc"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewEipAssociation(ctx, "eipAssoc", &ec2.EipAssociationArgs{
			InstanceId:   web.ID(),
			AllocationId: example.ID(),
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
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import com.pulumi.aws.ec2.Eip;
import com.pulumi.aws.ec2.EipArgs;
import com.pulumi.aws.ec2.EipAssociation;
import com.pulumi.aws.ec2.EipAssociationArgs;
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
        var web = new Instance("web", InstanceArgs.builder()        
            .ami("ami-21f78e11")
            .availabilityZone("us-west-2a")
            .instanceType("t2.micro")
            .tags(Map.of("Name", "HelloWorld"))
            .build());

        var example = new Eip("example", EipArgs.builder()        
            .domain("vpc")
            .build());

        var eipAssoc = new EipAssociation("eipAssoc", EipAssociationArgs.builder()        
            .instanceId(web.id())
            .allocationId(example.id())
            .build());

    }
}
```
```yaml
resources:
  eipAssoc:
    type: aws:ec2:EipAssociation
    properties:
      instanceId: ${web.id}
      allocationId: ${example.id}
  web:
    type: aws:ec2:Instance
    properties:
      ami: ami-21f78e11
      availabilityZone: us-west-2a
      instanceType: t2.micro
      tags:
        Name: HelloWorld
  example:
    type: aws:ec2:Eip
    properties:
      domain: vpc
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EIP Assocations using their association IDs. For example:

```sh
 $ pulumi import aws:ec2/eipAssociation:EipAssociation test eipassoc-ab12c345
```
 