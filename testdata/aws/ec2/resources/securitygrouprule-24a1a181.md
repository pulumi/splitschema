Provides a security group rule resource. Represents a single `ingress` or
`egress` group rule, which can be added to external Security Groups.

> **NOTE on Security Groups and Security Group Rules:** This provider currently provides a Security Group resource with `ingress` and `egress` rules defined in-line and a Security Group Rule resource which manages one or more `ingress` or
`egress` rules. Both of these resource were added before AWS assigned a [security group rule unique ID](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules.html), and they do not work well in all scenarios using the`description` and `tags` attributes, which rely on the unique ID.
The `aws.vpc.SecurityGroupEgressRule` and `aws.vpc.SecurityGroupIngressRule` resources have been added to address these limitations and should be used for all new security group rules.
You should not use the `aws.vpc.SecurityGroupEgressRule` and `aws.vpc.SecurityGroupIngressRule` resources in conjunction with an `aws.ec2.SecurityGroup` resource with in-line rules or with `aws.ec2.SecurityGroupRule` resources defined for the same Security Group, as rule conflicts may occur and rules will be overwritten.

> **NOTE:** Setting `protocol = "all"` or `protocol = -1` with `from_port` and `to_port` will result in the EC2 API creating a security group rule with all ports open. This API behavior cannot be controlled by this provider and may generate warnings in the future.

> **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.SecurityGroupRule("example", {
    type: "ingress",
    fromPort: 0,
    toPort: 65535,
    protocol: "tcp",
    cidrBlocks: [aws_vpc.example.cidr_block],
    ipv6CidrBlocks: [aws_vpc.example.ipv6_cidr_block],
    securityGroupId: "sg-123456",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.SecurityGroupRule("example",
    type="ingress",
    from_port=0,
    to_port=65535,
    protocol="tcp",
    cidr_blocks=[aws_vpc["example"]["cidr_block"]],
    ipv6_cidr_blocks=[aws_vpc["example"]["ipv6_cidr_block"]],
    security_group_id="sg-123456")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.SecurityGroupRule("example", new()
    {
        Type = "ingress",
        FromPort = 0,
        ToPort = 65535,
        Protocol = "tcp",
        CidrBlocks = new[]
        {
            aws_vpc.Example.Cidr_block,
        },
        Ipv6CidrBlocks = new[]
        {
            aws_vpc.Example.Ipv6_cidr_block,
        },
        SecurityGroupId = "sg-123456",
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
		_, err := ec2.NewSecurityGroupRule(ctx, "example", &ec2.SecurityGroupRuleArgs{
			Type:     pulumi.String("ingress"),
			FromPort: pulumi.Int(0),
			ToPort:   pulumi.Int(65535),
			Protocol: pulumi.String("tcp"),
			CidrBlocks: pulumi.StringArray{
				aws_vpc.Example.Cidr_block,
			},
			Ipv6CidrBlocks: pulumi.StringArray{
				aws_vpc.Example.Ipv6_cidr_block,
			},
			SecurityGroupId: pulumi.String("sg-123456"),
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
import com.pulumi.aws.ec2.SecurityGroupRule;
import com.pulumi.aws.ec2.SecurityGroupRuleArgs;
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
        var example = new SecurityGroupRule("example", SecurityGroupRuleArgs.builder()        
            .type("ingress")
            .fromPort(0)
            .toPort(65535)
            .protocol("tcp")
            .cidrBlocks(aws_vpc.example().cidr_block())
            .ipv6CidrBlocks(aws_vpc.example().ipv6_cidr_block())
            .securityGroupId("sg-123456")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:SecurityGroupRule
    properties:
      type: ingress
      fromPort: 0
      toPort: 65535
      protocol: tcp
      cidrBlocks:
        - ${aws_vpc.example.cidr_block}
      ipv6CidrBlocks:
        - ${aws_vpc.example.ipv6_cidr_block}
      securityGroupId: sg-123456
```
{{% /example %}}
{{% example %}}
### Usage With Prefix List IDs

Prefix Lists are either managed by AWS internally, or created by the customer using a
Managed Prefix List resource. Prefix Lists provided by
AWS are associated with a prefix list name, or service name, that is linked to a specific region.

Prefix list IDs are exported on VPC Endpoints, so you can use this format:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ...
const myEndpoint = new aws.ec2.VpcEndpoint("myEndpoint", {});
// ...
const allowAll = new aws.ec2.SecurityGroupRule("allowAll", {
    type: "egress",
    toPort: 0,
    protocol: "-1",
    prefixListIds: [myEndpoint.prefixListId],
    fromPort: 0,
    securityGroupId: "sg-123456",
});
```
```python
import pulumi
import pulumi_aws as aws

# ...
my_endpoint = aws.ec2.VpcEndpoint("myEndpoint")
# ...
allow_all = aws.ec2.SecurityGroupRule("allowAll",
    type="egress",
    to_port=0,
    protocol="-1",
    prefix_list_ids=[my_endpoint.prefix_list_id],
    from_port=0,
    security_group_id="sg-123456")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ...
    var myEndpoint = new Aws.Ec2.VpcEndpoint("myEndpoint");

    // ...
    var allowAll = new Aws.Ec2.SecurityGroupRule("allowAll", new()
    {
        Type = "egress",
        ToPort = 0,
        Protocol = "-1",
        PrefixListIds = new[]
        {
            myEndpoint.PrefixListId,
        },
        FromPort = 0,
        SecurityGroupId = "sg-123456",
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
		_, err = ec2.NewSecurityGroupRule(ctx, "allowAll", &ec2.SecurityGroupRuleArgs{
			Type:     pulumi.String("egress"),
			ToPort:   pulumi.Int(0),
			Protocol: pulumi.String("-1"),
			PrefixListIds: pulumi.StringArray{
				myEndpoint.PrefixListId,
			},
			FromPort:        pulumi.Int(0),
			SecurityGroupId: pulumi.String("sg-123456"),
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
import com.pulumi.aws.ec2.SecurityGroupRule;
import com.pulumi.aws.ec2.SecurityGroupRuleArgs;
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

        var allowAll = new SecurityGroupRule("allowAll", SecurityGroupRuleArgs.builder()        
            .type("egress")
            .toPort(0)
            .protocol("-1")
            .prefixListIds(myEndpoint.prefixListId())
            .fromPort(0)
            .securityGroupId("sg-123456")
            .build());

    }
}
```
```yaml
resources:
  allowAll:
    type: aws:ec2:SecurityGroupRule
    properties:
      type: egress
      toPort: 0
      protocol: '-1'
      prefixListIds:
        - ${myEndpoint.prefixListId}
      fromPort: 0
      securityGroupId: sg-123456
  # ...
  myEndpoint:
    type: aws:ec2:VpcEndpoint
```

You can also find a specific Prefix List using the `aws.ec2.getPrefixList`
or `ec2_managed_prefix_list` data sources:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegion({});
const s3 = current.then(current => aws.ec2.getPrefixList({
    name: `com.amazonaws.${current.name}.s3`,
}));
const s3GatewayEgress = new aws.ec2.SecurityGroupRule("s3GatewayEgress", {
    description: "S3 Gateway Egress",
    type: "egress",
    securityGroupId: "sg-123456",
    fromPort: 443,
    toPort: 443,
    protocol: "tcp",
    prefixListIds: [s3.then(s3 => s3.id)],
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_region()
s3 = aws.ec2.get_prefix_list(name=f"com.amazonaws.{current.name}.s3")
s3_gateway_egress = aws.ec2.SecurityGroupRule("s3GatewayEgress",
    description="S3 Gateway Egress",
    type="egress",
    security_group_id="sg-123456",
    from_port=443,
    to_port=443,
    protocol="tcp",
    prefix_list_ids=[s3.id])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegion.Invoke();

    var s3 = Aws.Ec2.GetPrefixList.Invoke(new()
    {
        Name = $"com.amazonaws.{current.Apply(getRegionResult => getRegionResult.Name)}.s3",
    });

    var s3GatewayEgress = new Aws.Ec2.SecurityGroupRule("s3GatewayEgress", new()
    {
        Description = "S3 Gateway Egress",
        Type = "egress",
        SecurityGroupId = "sg-123456",
        FromPort = 443,
        ToPort = 443,
        Protocol = "tcp",
        PrefixListIds = new[]
        {
            s3.Apply(getPrefixListResult => getPrefixListResult.Id),
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		s3, err := ec2.GetPrefixList(ctx, &ec2.GetPrefixListArgs{
			Name: pulumi.StringRef(fmt.Sprintf("com.amazonaws.%v.s3", current.Name)),
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSecurityGroupRule(ctx, "s3GatewayEgress", &ec2.SecurityGroupRuleArgs{
			Description:     pulumi.String("S3 Gateway Egress"),
			Type:            pulumi.String("egress"),
			SecurityGroupId: pulumi.String("sg-123456"),
			FromPort:        pulumi.Int(443),
			ToPort:          pulumi.Int(443),
			Protocol:        pulumi.String("tcp"),
			PrefixListIds: pulumi.StringArray{
				*pulumi.String(s3.Id),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetPrefixListArgs;
import com.pulumi.aws.ec2.SecurityGroupRule;
import com.pulumi.aws.ec2.SecurityGroupRuleArgs;
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
        final var current = AwsFunctions.getRegion();

        final var s3 = Ec2Functions.getPrefixList(GetPrefixListArgs.builder()
            .name(String.format("com.amazonaws.%s.s3", current.applyValue(getRegionResult -> getRegionResult.name())))
            .build());

        var s3GatewayEgress = new SecurityGroupRule("s3GatewayEgress", SecurityGroupRuleArgs.builder()        
            .description("S3 Gateway Egress")
            .type("egress")
            .securityGroupId("sg-123456")
            .fromPort(443)
            .toPort(443)
            .protocol("tcp")
            .prefixListIds(s3.applyValue(getPrefixListResult -> getPrefixListResult.id()))
            .build());

    }
}
```
```yaml
resources:
  s3GatewayEgress:
    type: aws:ec2:SecurityGroupRule
    properties:
      # S3 Gateway interfaces are implemented at the routing level which means we
      #   # can avoid the metered billing of a VPC endpoint interface by allowing
      #   # outbound traffic to the public IP ranges, which will be routed through
      #   # the Gateway interface:
      #   # https://docs.aws.amazon.com/AmazonS3/latest/userguide/privatelink-interface-endpoints.html
      description: S3 Gateway Egress
      type: egress
      securityGroupId: sg-123456
      fromPort: 443
      toPort: 443
      protocol: tcp
      prefixListIds:
        - ${s3.id}
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  s3:
    fn::invoke:
      Function: aws:ec2:getPrefixList
      Arguments:
        name: com.amazonaws.${current.name}.s3
```
{{% /example %}}
{{% /examples %}}

## Import

Import a rule with various IPv4 and IPv6 source CIDR blocks:

Import a rule, applicable to all ports, with a protocol other than TCP/UDP/ICMP/ICMPV6/ALL, e.g., Multicast Transport Protocol (MTP), using the IANA protocol number. For example92.

Import a default any/any egress rule to 0.0.0.0/0:

Import an egress rule with a prefix list ID destination:

Import a rule applicable to all protocols and ports with a security group source:

Import a rule that has itself and an IPv6 CIDR block as sources:

__Using `pulumi import` to import__ Security Group Rules using the `security_group_id`, `type`, `protocol`, `from_port`, `to_port`, and source(s)/destination(s) (such as a `cidr_block`) separated by underscores (`_`). All parts are required. For example:

__NOTE:__ Not all rule permissions (e.g., not all of a rule's CIDR blocks) need to be imported for this provider to manage rule permissions. However, importing some of a rule's permissions but not others, and then making changes to the rule will result in the creation of an additional rule to capture the updated permissions. Rule permissions that were not imported are left intact in the original rule.

Import an ingress rule in security group `sg-6e616f6d69` for TCP port 8000 with an IPv4 destination CIDR of `10.0.3.0/24`:

```sh
 $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress sg-6e616f6d69_ingress_tcp_8000_8000_10.0.3.0/24
```
 Import a rule with various IPv4 and IPv6 source CIDR blocks:

```sh
 $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress sg-4973616163_ingress_tcp_100_121_10.1.0.0/16_2001:db8::/48_10.2.0.0/16_2002:db8::/48
```
 Import a rule, applicable to all ports, with a protocol other than TCP/UDP/ICMP/ICMPV6/ALL, e.g., Multicast Transport Protocol (MTP), using the IANA protocol number. For example92.

```sh
 $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress sg-6777656e646f6c796e_ingress_92_0_65536_10.0.3.0/24_10.0.4.0/24
```
 Import a default any/any egress rule to 0.0.0.0/0:

```sh
 $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule default_egress sg-6777656e646f6c796e_egress_all_0_0_0.0.0.0/0
```
 Import an egress rule with a prefix list ID destination:

```sh
 $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule egress sg-62726f6479_egress_tcp_8000_8000_pl-6469726b
```
 Import a rule applicable to all protocols and ports with a security group source:

```sh
 $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule ingress_rule sg-7472697374616e_ingress_all_0_65536_sg-6176657279
```
 Import a rule that has itself and an IPv6 CIDR block as sources:

```sh
 $ pulumi import aws:ec2/securityGroupRule:SecurityGroupRule rule_name sg-656c65616e6f72_ingress_tcp_80_80_self_2001:db8::/48
```
 