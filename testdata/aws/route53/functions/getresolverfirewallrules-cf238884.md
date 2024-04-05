`aws.route53.getResolverFirewallRules` Provides details about rules in a specific Route53 Resolver Firewall rule group.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows how to get Route53 Resolver Firewall rules based on its associated firewall group id.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.route53.getResolverFirewallRules({
    firewallRuleGroupId: aws_route53_resolver_firewall_rule_group.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_firewall_rules(firewall_rule_group_id=aws_route53_resolver_firewall_rule_group["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Route53.GetResolverFirewallRules.Invoke(new()
    {
        FirewallRuleGroupId = aws_route53_resolver_firewall_rule_group.Example.Id,
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
		_, err := route53.GetResolverFirewallRules(ctx, &route53.GetResolverFirewallRulesArgs{
			FirewallRuleGroupId: aws_route53_resolver_firewall_rule_group.Example.Id,
		}, nil)
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
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetResolverFirewallRulesArgs;
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
        final var example = Route53Functions.getResolverFirewallRules(GetResolverFirewallRulesArgs.builder()
            .firewallRuleGroupId(aws_route53_resolver_firewall_rule_group.example().id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:route53:getResolverFirewallRules
      Arguments:
        firewallRuleGroupId: ${aws_route53_resolver_firewall_rule_group.example.id}
```
{{% /example %}}
{{% /examples %}}