Creates an entry (a rule) in a network ACL with the specified rule number.

> **NOTE on Network ACLs and Network ACL Rules:** This provider currently
provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
defined in-line. At this time you cannot use a Network ACL with in-line rules
in conjunction with any Network ACL Rule resources. Doing so will cause
a conflict of rule settings and will overwrite rules.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const barNetworkAcl = new aws.ec2.NetworkAcl("barNetworkAcl", {vpcId: aws_vpc.foo.id});
const barNetworkAclRule = new aws.ec2.NetworkAclRule("barNetworkAclRule", {
    networkAclId: barNetworkAcl.id,
    ruleNumber: 200,
    egress: false,
    protocol: "tcp",
    ruleAction: "allow",
    cidrBlock: aws_vpc.foo.cidr_block,
    fromPort: 22,
    toPort: 22,
});
```
```python
import pulumi
import pulumi_aws as aws

bar_network_acl = aws.ec2.NetworkAcl("barNetworkAcl", vpc_id=aws_vpc["foo"]["id"])
bar_network_acl_rule = aws.ec2.NetworkAclRule("barNetworkAclRule",
    network_acl_id=bar_network_acl.id,
    rule_number=200,
    egress=False,
    protocol="tcp",
    rule_action="allow",
    cidr_block=aws_vpc["foo"]["cidr_block"],
    from_port=22,
    to_port=22)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var barNetworkAcl = new Aws.Ec2.NetworkAcl("barNetworkAcl", new()
    {
        VpcId = aws_vpc.Foo.Id,
    });

    var barNetworkAclRule = new Aws.Ec2.NetworkAclRule("barNetworkAclRule", new()
    {
        NetworkAclId = barNetworkAcl.Id,
        RuleNumber = 200,
        Egress = false,
        Protocol = "tcp",
        RuleAction = "allow",
        CidrBlock = aws_vpc.Foo.Cidr_block,
        FromPort = 22,
        ToPort = 22,
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
		barNetworkAcl, err := ec2.NewNetworkAcl(ctx, "barNetworkAcl", &ec2.NetworkAclArgs{
			VpcId: pulumi.Any(aws_vpc.Foo.Id),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewNetworkAclRule(ctx, "barNetworkAclRule", &ec2.NetworkAclRuleArgs{
			NetworkAclId: barNetworkAcl.ID(),
			RuleNumber:   pulumi.Int(200),
			Egress:       pulumi.Bool(false),
			Protocol:     pulumi.String("tcp"),
			RuleAction:   pulumi.String("allow"),
			CidrBlock:    pulumi.Any(aws_vpc.Foo.Cidr_block),
			FromPort:     pulumi.Int(22),
			ToPort:       pulumi.Int(22),
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
import com.pulumi.aws.ec2.NetworkAcl;
import com.pulumi.aws.ec2.NetworkAclArgs;
import com.pulumi.aws.ec2.NetworkAclRule;
import com.pulumi.aws.ec2.NetworkAclRuleArgs;
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
        var barNetworkAcl = new NetworkAcl("barNetworkAcl", NetworkAclArgs.builder()        
            .vpcId(aws_vpc.foo().id())
            .build());

        var barNetworkAclRule = new NetworkAclRule("barNetworkAclRule", NetworkAclRuleArgs.builder()        
            .networkAclId(barNetworkAcl.id())
            .ruleNumber(200)
            .egress(false)
            .protocol("tcp")
            .ruleAction("allow")
            .cidrBlock(aws_vpc.foo().cidr_block())
            .fromPort(22)
            .toPort(22)
            .build());

    }
}
```
```yaml
resources:
  barNetworkAcl:
    type: aws:ec2:NetworkAcl
    properties:
      vpcId: ${aws_vpc.foo.id}
  barNetworkAclRule:
    type: aws:ec2:NetworkAclRule
    properties:
      networkAclId: ${barNetworkAcl.id}
      ruleNumber: 200
      egress: false
      protocol: tcp
      ruleAction: allow
      cidrBlock: ${aws_vpc.foo.cidr_block}
      fromPort: 22
      toPort: 22
```

> **Note:** One of either `cidr_block` or `ipv6_cidr_block` is required.
{{% /example %}}
{{% /examples %}}

## Import

Using the procotol's decimal value:

__Using `pulumi import` to import__ individual rules using `NETWORK_ACL_ID:RULE_NUMBER:PROTOCOL:EGRESS`, where `PROTOCOL` can be a decimal (such as "6") or string (such as "tcp") value. For example:

Using the procotol's string value:

```sh
 $ pulumi import aws:ec2/networkAclRule:NetworkAclRule my_rule acl-7aaabd18:100:tcp:false
```
 Using the procotol's decimal value:

```sh
 $ pulumi import aws:ec2/networkAclRule:NetworkAclRule my_rule acl-7aaabd18:100:6:false
```
 