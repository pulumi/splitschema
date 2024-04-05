Provides a Route 53 Resolver DNS Firewall rule resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleResolverFirewallDomainList = new aws.route53.ResolverFirewallDomainList("exampleResolverFirewallDomainList", {
    domains: ["example.com"],
    tags: {},
});
const exampleResolverFirewallRuleGroup = new aws.route53.ResolverFirewallRuleGroup("exampleResolverFirewallRuleGroup", {tags: {}});
const exampleResolverFirewallRule = new aws.route53.ResolverFirewallRule("exampleResolverFirewallRule", {
    action: "BLOCK",
    blockOverrideDnsType: "CNAME",
    blockOverrideDomain: "example.com",
    blockOverrideTtl: 1,
    blockResponse: "OVERRIDE",
    firewallDomainListId: exampleResolverFirewallDomainList.id,
    firewallRuleGroupId: exampleResolverFirewallRuleGroup.id,
    priority: 100,
});
```
```python
import pulumi
import pulumi_aws as aws

example_resolver_firewall_domain_list = aws.route53.ResolverFirewallDomainList("exampleResolverFirewallDomainList",
    domains=["example.com"],
    tags={})
example_resolver_firewall_rule_group = aws.route53.ResolverFirewallRuleGroup("exampleResolverFirewallRuleGroup", tags={})
example_resolver_firewall_rule = aws.route53.ResolverFirewallRule("exampleResolverFirewallRule",
    action="BLOCK",
    block_override_dns_type="CNAME",
    block_override_domain="example.com",
    block_override_ttl=1,
    block_response="OVERRIDE",
    firewall_domain_list_id=example_resolver_firewall_domain_list.id,
    firewall_rule_group_id=example_resolver_firewall_rule_group.id,
    priority=100)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleResolverFirewallDomainList = new Aws.Route53.ResolverFirewallDomainList("exampleResolverFirewallDomainList", new()
    {
        Domains = new[]
        {
            "example.com",
        },
        Tags = null,
    });

    var exampleResolverFirewallRuleGroup = new Aws.Route53.ResolverFirewallRuleGroup("exampleResolverFirewallRuleGroup", new()
    {
        Tags = null,
    });

    var exampleResolverFirewallRule = new Aws.Route53.ResolverFirewallRule("exampleResolverFirewallRule", new()
    {
        Action = "BLOCK",
        BlockOverrideDnsType = "CNAME",
        BlockOverrideDomain = "example.com",
        BlockOverrideTtl = 1,
        BlockResponse = "OVERRIDE",
        FirewallDomainListId = exampleResolverFirewallDomainList.Id,
        FirewallRuleGroupId = exampleResolverFirewallRuleGroup.Id,
        Priority = 100,
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
		exampleResolverFirewallDomainList, err := route53.NewResolverFirewallDomainList(ctx, "exampleResolverFirewallDomainList", &route53.ResolverFirewallDomainListArgs{
			Domains: pulumi.StringArray{
				pulumi.String("example.com"),
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		exampleResolverFirewallRuleGroup, err := route53.NewResolverFirewallRuleGroup(ctx, "exampleResolverFirewallRuleGroup", &route53.ResolverFirewallRuleGroupArgs{
			Tags: nil,
		})
		if err != nil {
			return err
		}
		_, err = route53.NewResolverFirewallRule(ctx, "exampleResolverFirewallRule", &route53.ResolverFirewallRuleArgs{
			Action:               pulumi.String("BLOCK"),
			BlockOverrideDnsType: pulumi.String("CNAME"),
			BlockOverrideDomain:  pulumi.String("example.com"),
			BlockOverrideTtl:     pulumi.Int(1),
			BlockResponse:        pulumi.String("OVERRIDE"),
			FirewallDomainListId: exampleResolverFirewallDomainList.ID(),
			FirewallRuleGroupId:  exampleResolverFirewallRuleGroup.ID(),
			Priority:             pulumi.Int(100),
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
import com.pulumi.aws.route53.ResolverFirewallDomainList;
import com.pulumi.aws.route53.ResolverFirewallDomainListArgs;
import com.pulumi.aws.route53.ResolverFirewallRuleGroup;
import com.pulumi.aws.route53.ResolverFirewallRuleGroupArgs;
import com.pulumi.aws.route53.ResolverFirewallRule;
import com.pulumi.aws.route53.ResolverFirewallRuleArgs;
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
        var exampleResolverFirewallDomainList = new ResolverFirewallDomainList("exampleResolverFirewallDomainList", ResolverFirewallDomainListArgs.builder()        
            .domains("example.com")
            .tags()
            .build());

        var exampleResolverFirewallRuleGroup = new ResolverFirewallRuleGroup("exampleResolverFirewallRuleGroup", ResolverFirewallRuleGroupArgs.builder()        
            .tags()
            .build());

        var exampleResolverFirewallRule = new ResolverFirewallRule("exampleResolverFirewallRule", ResolverFirewallRuleArgs.builder()        
            .action("BLOCK")
            .blockOverrideDnsType("CNAME")
            .blockOverrideDomain("example.com")
            .blockOverrideTtl(1)
            .blockResponse("OVERRIDE")
            .firewallDomainListId(exampleResolverFirewallDomainList.id())
            .firewallRuleGroupId(exampleResolverFirewallRuleGroup.id())
            .priority(100)
            .build());

    }
}
```
```yaml
resources:
  exampleResolverFirewallDomainList:
    type: aws:route53:ResolverFirewallDomainList
    properties:
      domains:
        - example.com
      tags: {}
  exampleResolverFirewallRuleGroup:
    type: aws:route53:ResolverFirewallRuleGroup
    properties:
      tags: {}
  exampleResolverFirewallRule:
    type: aws:route53:ResolverFirewallRule
    properties:
      action: BLOCK
      blockOverrideDnsType: CNAME
      blockOverrideDomain: example.com
      blockOverrideTtl: 1
      blockResponse: OVERRIDE
      firewallDomainListId: ${exampleResolverFirewallDomainList.id}
      firewallRuleGroupId: ${exampleResolverFirewallRuleGroup.id}
      priority: 100
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import

Route 53 Resolver DNS Firewall rules using the Route 53 Resolver DNS Firewall rule group ID and domain list ID separated by ':'. For example:

```sh
 $ pulumi import aws:route53/resolverFirewallRule:ResolverFirewallRule example rslvr-frg-0123456789abcdef:rslvr-fdl-0123456789abcdef
```
 