`aws.route53.ResolverFirewallRuleGroup` Retrieves the specified firewall rule group.

This data source allows to retrieve details about a specific a Route 53 Resolver DNS Firewall rule group.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows how to get a firewall rule group from its ID.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.route53.getResolverFirewallRuleGroup({
    firewallRuleGroupId: "rslvr-frg-example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_firewall_rule_group(firewall_rule_group_id="rslvr-frg-example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Route53.GetResolverFirewallRuleGroup.Invoke(new()
    {
        FirewallRuleGroupId = "rslvr-frg-example",
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
		_, err := route53.LookupResolverFirewallRuleGroup(ctx, &route53.LookupResolverFirewallRuleGroupArgs{
			FirewallRuleGroupId: "rslvr-frg-example",
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
import com.pulumi.aws.route53.inputs.GetResolverFirewallRuleGroupArgs;
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
        final var example = Route53Functions.getResolverFirewallRuleGroup(GetResolverFirewallRuleGroupArgs.builder()
            .firewallRuleGroupId("rslvr-frg-example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:route53:getResolverFirewallRuleGroup
      Arguments:
        firewallRuleGroupId: rslvr-frg-example
```
{{% /example %}}
{{% /examples %}}