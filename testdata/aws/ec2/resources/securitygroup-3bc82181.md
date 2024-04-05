Provides a security group resource.

> **NOTE on Security Groups and Security Group Rules:** This provider currently provides a Security Group resource with `ingress` and `egress` rules defined in-line and a Security Group Rule resource which manages one or more `ingress` or `egress` rules. Both of these resource were added before AWS assigned a [security group rule unique ID](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules.html), and they do not work well in all scenarios using the`description` and `tags` attributes, which rely on the unique ID. The `aws.vpc.SecurityGroupEgressRule` and `aws.vpc.SecurityGroupIngressRule` resources have been added to address these limitations and should be used for all new security group rules. You should not use the `aws.vpc.SecurityGroupEgressRule` and `aws.vpc.SecurityGroupIngressRule` resources in conjunction with an `aws.ec2.SecurityGroup` resource with in-line rules or with `aws.ec2.SecurityGroupRule` resources defined for the same Security Group, as rule conflicts may occur and rules will be overwritten.

> **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).

> **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), security groups associated with Lambda Functions can take up to 45 minutes to successfully delete.

> **NOTE:** The `cidr_blocks` and `ipv6_cidr_blocks` parameters are optional in the `ingress` and `egress` blocks. If nothing is specified, traffic will be blocked as described in _NOTE on Egress rules_ later.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const allowTls = new aws.ec2.SecurityGroup("allowTls", {
    description: "Allow TLS inbound traffic",
    vpcId: aws_vpc.main.id,
    ingress: [{
        description: "TLS from VPC",
        fromPort: 443,
        toPort: 443,
        protocol: "tcp",
        cidrBlocks: [aws_vpc.main.cidr_block],
        ipv6CidrBlocks: [aws_vpc.main.ipv6_cidr_block],
    }],
    egress: [{
        fromPort: 0,
        toPort: 0,
        protocol: "-1",
        cidrBlocks: ["0.0.0.0/0"],
        ipv6CidrBlocks: ["::/0"],
    }],
    tags: {
        Name: "allow_tls",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

allow_tls = aws.ec2.SecurityGroup("allowTls",
    description="Allow TLS inbound traffic",
    vpc_id=aws_vpc["main"]["id"],
    ingress=[aws.ec2.SecurityGroupIngressArgs(
        description="TLS from VPC",
        from_port=443,
        to_port=443,
        protocol="tcp",
        cidr_blocks=[aws_vpc["main"]["cidr_block"]],
        ipv6_cidr_blocks=[aws_vpc["main"]["ipv6_cidr_block"]],
    )],
    egress=[aws.ec2.SecurityGroupEgressArgs(
        from_port=0,
        to_port=0,
        protocol="-1",
        cidr_blocks=["0.0.0.0/0"],
        ipv6_cidr_blocks=["::/0"],
    )],
    tags={
        "Name": "allow_tls",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var allowTls = new Aws.Ec2.SecurityGroup("allowTls", new()
    {
        Description = "Allow TLS inbound traffic",
        VpcId = aws_vpc.Main.Id,
        Ingress = new[]
        {
            new Aws.Ec2.Inputs.SecurityGroupIngressArgs
            {
                Description = "TLS from VPC",
                FromPort = 443,
                ToPort = 443,
                Protocol = "tcp",
                CidrBlocks = new[]
                {
                    aws_vpc.Main.Cidr_block,
                },
                Ipv6CidrBlocks = new[]
                {
                    aws_vpc.Main.Ipv6_cidr_block,
                },
            },
        },
        Egress = new[]
        {
            new Aws.Ec2.Inputs.SecurityGroupEgressArgs
            {
                FromPort = 0,
                ToPort = 0,
                Protocol = "-1",
                CidrBlocks = new[]
                {
                    "0.0.0.0/0",
                },
                Ipv6CidrBlocks = new[]
                {
                    "::/0",
                },
            },
        },
        Tags = 
        {
            { "Name", "allow_tls" },
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
		_, err := ec2.NewSecurityGroup(ctx, "allowTls", &ec2.SecurityGroupArgs{
			Description: pulumi.String("Allow TLS inbound traffic"),
			VpcId:       pulumi.Any(aws_vpc.Main.Id),
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					Description: pulumi.String("TLS from VPC"),
					FromPort:    pulumi.Int(443),
					ToPort:      pulumi.Int(443),
					Protocol:    pulumi.String("tcp"),
					CidrBlocks: pulumi.StringArray{
						aws_vpc.Main.Cidr_block,
					},
					Ipv6CidrBlocks: pulumi.StringArray{
						aws_vpc.Main.Ipv6_cidr_block,
					},
				},
			},
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
					Protocol: pulumi.String("-1"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
					Ipv6CidrBlocks: pulumi.StringArray{
						pulumi.String("::/0"),
					},
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("allow_tls"),
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
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
import com.pulumi.aws.ec2.inputs.SecurityGroupIngressArgs;
import com.pulumi.aws.ec2.inputs.SecurityGroupEgressArgs;
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
        var allowTls = new SecurityGroup("allowTls", SecurityGroupArgs.builder()        
            .description("Allow TLS inbound traffic")
            .vpcId(aws_vpc.main().id())
            .ingress(SecurityGroupIngressArgs.builder()
                .description("TLS from VPC")
                .fromPort(443)
                .toPort(443)
                .protocol("tcp")
                .cidrBlocks(aws_vpc.main().cidr_block())
                .ipv6CidrBlocks(aws_vpc.main().ipv6_cidr_block())
                .build())
            .egress(SecurityGroupEgressArgs.builder()
                .fromPort(0)
                .toPort(0)
                .protocol("-1")
                .cidrBlocks("0.0.0.0/0")
                .ipv6CidrBlocks("::/0")
                .build())
            .tags(Map.of("Name", "allow_tls"))
            .build());

    }
}
```
```yaml
resources:
  allowTls:
    type: aws:ec2:SecurityGroup
    properties:
      description: Allow TLS inbound traffic
      vpcId: ${aws_vpc.main.id}
      ingress:
        - description: TLS from VPC
          fromPort: 443
          toPort: 443
          protocol: tcp
          cidrBlocks:
            - ${aws_vpc.main.cidr_block}
          ipv6CidrBlocks:
            - ${aws_vpc.main.ipv6_cidr_block}
      egress:
        - fromPort: 0
          toPort: 0
          protocol: '-1'
          cidrBlocks:
            - 0.0.0.0/0
          ipv6CidrBlocks:
            - ::/0
      tags:
        Name: allow_tls
```

> **NOTE on Egress rules:** By default, AWS creates an `ALLOW ALL` egress rule when creating a new Security Group inside of a VPC. When creating a new Security Group inside a VPC, **this provider will remove this default rule**, and require you specifically re-create it if you desire that rule. We feel this leads to fewer surprises in terms of controlling your egress rules. If you desire this rule to be in place, you can use this `egress` block:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.SecurityGroup("example", {egress: [{
    cidrBlocks: ["0.0.0.0/0"],
    fromPort: 0,
    ipv6CidrBlocks: ["::/0"],
    protocol: "-1",
    toPort: 0,
}]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.SecurityGroup("example", egress=[aws.ec2.SecurityGroupEgressArgs(
    cidr_blocks=["0.0.0.0/0"],
    from_port=0,
    ipv6_cidr_blocks=["::/0"],
    protocol="-1",
    to_port=0,
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.SecurityGroup("example", new()
    {
        Egress = new[]
        {
            new Aws.Ec2.Inputs.SecurityGroupEgressArgs
            {
                CidrBlocks = new[]
                {
                    "0.0.0.0/0",
                },
                FromPort = 0,
                Ipv6CidrBlocks = new[]
                {
                    "::/0",
                },
                Protocol = "-1",
                ToPort = 0,
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
		_, err := ec2.NewSecurityGroup(ctx, "example", &ec2.SecurityGroupArgs{
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
					FromPort: pulumi.Int(0),
					Ipv6CidrBlocks: pulumi.StringArray{
						pulumi.String("::/0"),
					},
					Protocol: pulumi.String("-1"),
					ToPort:   pulumi.Int(0),
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
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
import com.pulumi.aws.ec2.inputs.SecurityGroupEgressArgs;
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
        var example = new SecurityGroup("example", SecurityGroupArgs.builder()        
            .egress(SecurityGroupEgressArgs.builder()
                .cidrBlocks("0.0.0.0/0")
                .fromPort(0)
                .ipv6CidrBlocks("::/0")
                .protocol("-1")
                .toPort(0)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:SecurityGroup
    properties:
      egress:
        - cidrBlocks:
            - 0.0.0.0/0
          fromPort: 0
          ipv6CidrBlocks:
            - ::/0
          protocol: '-1'
          toPort: 0
```
{{% /example %}}
{{% example %}}
### Usage With Prefix List IDs

Prefix Lists are either managed by AWS internally, or created by the customer using a
Prefix List resource. Prefix Lists provided by
AWS are associated with a prefix list name, or service name, that is linked to a specific region.
Prefix list IDs are exported on VPC Endpoints, so you can use this format:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myEndpoint = new aws.ec2.VpcEndpoint("myEndpoint", {});
// ... other configuration ...
// ... other configuration ...
const example = new aws.ec2.SecurityGroup("example", {egress: [{
    fromPort: 0,
    toPort: 0,
    protocol: "-1",
    prefixListIds: [myEndpoint.prefixListId],
}]});
```
```python
import pulumi
import pulumi_aws as aws

my_endpoint = aws.ec2.VpcEndpoint("myEndpoint")
# ... other configuration ...
# ... other configuration ...
example = aws.ec2.SecurityGroup("example", egress=[aws.ec2.SecurityGroupEgressArgs(
    from_port=0,
    to_port=0,
    protocol="-1",
    prefix_list_ids=[my_endpoint.prefix_list_id],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myEndpoint = new Aws.Ec2.VpcEndpoint("myEndpoint");

    // ... other configuration ...
    // ... other configuration ...
    var example = new Aws.Ec2.SecurityGroup("example", new()
    {
        Egress = new[]
        {
            new Aws.Ec2.Inputs.SecurityGroupEgressArgs
            {
                FromPort = 0,
                ToPort = 0,
                Protocol = "-1",
                PrefixListIds = new[]
                {
                    myEndpoint.PrefixListId,
                },
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
		myEndpoint, err := ec2.NewVpcEndpoint(ctx, "myEndpoint", nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSecurityGroup(ctx, "example", &ec2.SecurityGroupArgs{
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
					Protocol: pulumi.String("-1"),
					PrefixListIds: pulumi.StringArray{
						myEndpoint.PrefixListId,
					},
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
import com.pulumi.aws.ec2.VpcEndpoint;
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
import com.pulumi.aws.ec2.inputs.SecurityGroupEgressArgs;
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
        var myEndpoint = new VpcEndpoint("myEndpoint");

        var example = new SecurityGroup("example", SecurityGroupArgs.builder()        
            .egress(SecurityGroupEgressArgs.builder()
                .fromPort(0)
                .toPort(0)
                .protocol("-1")
                .prefixListIds(myEndpoint.prefixListId())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:SecurityGroup
    properties:
      egress:
        - fromPort: 0
          toPort: 0
          protocol: '-1'
          prefixListIds:
            - ${myEndpoint.prefixListId}
  myEndpoint:
    type: aws:ec2:VpcEndpoint
```

You can also find a specific Prefix List using the `aws.ec2.getPrefixList` data source.
{{% /example %}}
{{% example %}}
### Removing All Ingress and Egress Rules

The `ingress` and `egress` arguments are processed in attributes-as-blocks mode. Due to this, removing these arguments from the configuration will **not** cause the provider to destroy the managed rules. To subsequently remove all managed ingress and egress rules:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.SecurityGroup("example", {
    vpcId: aws_vpc.example.id,
    ingress: [],
    egress: [],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.SecurityGroup("example",
    vpc_id=aws_vpc["example"]["id"],
    ingress=[],
    egress=[])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.SecurityGroup("example", new()
    {
        VpcId = aws_vpc.Example.Id,
        Ingress = new[] {},
        Egress = new[] {},
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
		_, err := ec2.NewSecurityGroup(ctx, "example", &ec2.SecurityGroupArgs{
			VpcId:   pulumi.Any(aws_vpc.Example.Id),
			Ingress: ec2.SecurityGroupIngressArray{},
			Egress:  ec2.SecurityGroupEgressArray{},
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
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
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
        var example = new SecurityGroup("example", SecurityGroupArgs.builder()        
            .vpcId(aws_vpc.example().id())
            .ingress()
            .egress()
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:SecurityGroup
    properties:
      vpcId: ${aws_vpc.example.id}
      ingress: []
      egress: []
```
{{% /example %}}
### Recreating a Security Group

A simple security group `name` change "forces new" the security group--the provider destroys the security group and creates a new one. (Likewise, `description`, `name_prefix`, or `vpc_id` [cannot be changed](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/working-with-security-groups.html#creating-security-group).) Attempting to recreate the security group leads to a variety of complications depending on how it is used.

Security groups are generally associated with other resources--**more than 100** AWS Provider resources reference security groups. Referencing a resource from another resource creates a one-way dependency. For example, if you create an EC2 `aws.ec2.Instance` that has a `vpc_security_group_ids` argument that refers to an `aws.ec2.SecurityGroup` resource, the `aws.ec2.SecurityGroup` is a dependent of the `aws.ec2.Instance`. Because of this, the provider will create the security group first so that it can then be associated with the EC2 instance.

However, the dependency relationship actually goes both directions causing the _Security Group Deletion Problem_. AWS does not allow you to delete the security group associated with another resource (_e.g._, the `aws.ec2.Instance`).

The provider does not model bi-directional dependencies like this, but, even if it did, simply knowing the dependency situation would not be enough to solve it. For example, some resources must always have an associated security group while others don't need to. In addition, when the `aws.ec2.SecurityGroup` resource attempts to recreate, it receives a dependent object error, which does not provide information on whether the dependent object is a security group rule or, for example, an associated EC2 instance. Within the provider, the associated resource (_e.g._, `aws.ec2.Instance`) does not receive an error when the `aws.ec2.SecurityGroup` is trying to recreate even though that is where changes to the associated resource would need to take place (_e.g._, removing the security group association).

Despite these sticky problems, below are some ways to improve your experience when you find it necessary to recreate a security group.
{{% example %}}
### `create_before_destroy`

(This example is one approach to recreating security groups. For more information on the challenges and the _Security Group Deletion Problem_, see the section above.)

Normally, the provider first deletes the existing security group resource and then creates a new one. When a security group is associated with a resource, the delete won't succeed. You can invert the default behavior using the `create_before_destroy` meta argument:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.SecurityGroup("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.SecurityGroup("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.SecurityGroup("example");

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
		_, err := ec2.NewSecurityGroup(ctx, "example", nil)
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
import com.pulumi.aws.ec2.SecurityGroup;
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
        var example = new SecurityGroup("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:SecurityGroup
```
{{% /example %}}
{{% example %}}
### `replace_triggered_by`

(This example is one approach to recreating security groups. For more information on the challenges and the _Security Group Deletion Problem_, see the section above.)

To replace a resource when a security group changes, use the `replace_triggered_by` meta argument. Note that in this example, the `aws.ec2.Instance` will be destroyed and created again when the `aws.ec2.SecurityGroup` changes.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleSecurityGroup = new aws.ec2.SecurityGroup("exampleSecurityGroup", {});
// ... other configuration ...
const exampleInstance = new aws.ec2.Instance("exampleInstance", {
    instanceType: "t3.small",
    vpcSecurityGroupIds: [aws_security_group.test.id],
});
```
```python
import pulumi
import pulumi_aws as aws

example_security_group = aws.ec2.SecurityGroup("exampleSecurityGroup")
# ... other configuration ...
example_instance = aws.ec2.Instance("exampleInstance",
    instance_type="t3.small",
    vpc_security_group_ids=[aws_security_group["test"]["id"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleSecurityGroup = new Aws.Ec2.SecurityGroup("exampleSecurityGroup");

    // ... other configuration ...
    var exampleInstance = new Aws.Ec2.Instance("exampleInstance", new()
    {
        InstanceType = "t3.small",
        VpcSecurityGroupIds = new[]
        {
            aws_security_group.Test.Id,
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
		_, err := ec2.NewSecurityGroup(ctx, "exampleSecurityGroup", nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewInstance(ctx, "exampleInstance", &ec2.InstanceArgs{
			InstanceType: pulumi.String("t3.small"),
			VpcSecurityGroupIds: pulumi.StringArray{
				aws_security_group.Test.Id,
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
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
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
        var exampleSecurityGroup = new SecurityGroup("exampleSecurityGroup");

        var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()        
            .instanceType("t3.small")
            .vpcSecurityGroupIds(aws_security_group.test().id())
            .build());

    }
}
```
```yaml
resources:
  exampleSecurityGroup:
    type: aws:ec2:SecurityGroup
  exampleInstance:
    type: aws:ec2:Instance
    properties:
      instanceType: t3.small # ... other configuration ...
      vpcSecurityGroupIds:
        - ${aws_security_group.test.id}
```
{{% /example %}}
{{% example %}}
### Shorter timeout

(This example is one approach to recreating security groups. For more information on the challenges and the _Security Group Deletion Problem_, see the section above.)

If destroying a security group takes a long time, it may be because the provider cannot distinguish between a dependent object (_e.g._, a security group rule or EC2 instance) that is _in the process of being deleted_ and one that is not. In other words, it may be waiting for a train that isn't scheduled to arrive. To fail faster, shorten the `delete` timeout from the default timeout:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.SecurityGroup("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.SecurityGroup("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.SecurityGroup("example");

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
		_, err := ec2.NewSecurityGroup(ctx, "example", nil)
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
import com.pulumi.aws.ec2.SecurityGroup;
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
        var example = new SecurityGroup("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:SecurityGroup
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Security Groups using the security group `id`. For example:

```sh
 $ pulumi import aws:ec2/securityGroup:SecurityGroup elb_sg sg-903004f8
```
 