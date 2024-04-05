Provides a Route 53 Resolver DNS Firewall rule group association resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleResolverFirewallRuleGroup = new aws.route53.ResolverFirewallRuleGroup("exampleResolverFirewallRuleGroup", {});
const exampleResolverFirewallRuleGroupAssociation = new aws.route53.ResolverFirewallRuleGroupAssociation("exampleResolverFirewallRuleGroupAssociation", {
    firewallRuleGroupId: exampleResolverFirewallRuleGroup.id,
    priority: 100,
    vpcId: aws_vpc.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_resolver_firewall_rule_group = aws.route53.ResolverFirewallRuleGroup("exampleResolverFirewallRuleGroup")
example_resolver_firewall_rule_group_association = aws.route53.ResolverFirewallRuleGroupAssociation("exampleResolverFirewallRuleGroupAssociation",
    firewall_rule_group_id=example_resolver_firewall_rule_group.id,
    priority=100,
    vpc_id=aws_vpc["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleResolverFirewallRuleGroup = new Aws.Route53.ResolverFirewallRuleGroup("exampleResolverFirewallRuleGroup");

    var exampleResolverFirewallRuleGroupAssociation = new Aws.Route53.ResolverFirewallRuleGroupAssociation("exampleResolverFirewallRuleGroupAssociation", new()
    {
        FirewallRuleGroupId = exampleResolverFirewallRuleGroup.Id,
        Priority = 100,
        VpcId = aws_vpc.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleResolverFirewallRuleGroup, err := route53.NewResolverFirewallRuleGroup(ctx, "exampleResolverFirewallRuleGroup", nil)
		if err != nil {
			return err
		}
		_, err = route53.NewResolverFirewallRuleGroupAssociation(ctx, "exampleResolverFirewallRuleGroupAssociation", &route53.ResolverFirewallRuleGroupAssociationArgs{
			FirewallRuleGroupId: exampleResolverFirewallRuleGroup.ID(),
			Priority:            pulumi.Int(100),
			VpcId:               pulumi.Any(aws_vpc.Example.Id),
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
import com.pulumi.aws.route53.ResolverFirewallRuleGroup;
import com.pulumi.aws.route53.ResolverFirewallRuleGroupAssociation;
import com.pulumi.aws.route53.ResolverFirewallRuleGroupAssociationArgs;
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
        var exampleResolverFirewallRuleGroup = new ResolverFirewallRuleGroup("exampleResolverFirewallRuleGroup");

        var exampleResolverFirewallRuleGroupAssociation = new ResolverFirewallRuleGroupAssociation("exampleResolverFirewallRuleGroupAssociation", ResolverFirewallRuleGroupAssociationArgs.builder()        
            .firewallRuleGroupId(exampleResolverFirewallRuleGroup.id())
            .priority(100)
            .vpcId(aws_vpc.example().id())
            .build());

    }
}
```
```yaml
resources:
  exampleResolverFirewallRuleGroup:
    type: aws:route53:ResolverFirewallRuleGroup
  exampleResolverFirewallRuleGroupAssociation:
    type: aws:route53:ResolverFirewallRuleGroupAssociation
    properties:
      firewallRuleGroupId: ${exampleResolverFirewallRuleGroup.id}
      priority: 100
      vpcId: ${aws_vpc.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route 53 Resolver DNS Firewall rule group associations using the Route 53 Resolver DNS Firewall rule group association ID. For example:

```sh
 $ pulumi import aws:route53/resolverFirewallRuleGroupAssociation:ResolverFirewallRuleGroupAssociation example rslvr-frgassoc-0123456789abcdef
```
 