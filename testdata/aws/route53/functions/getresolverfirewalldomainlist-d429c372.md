`aws.route53.ResolverFirewallDomainList` Retrieves the specified firewall domain list.

This data source allows to retrieve details about a specific a Route 53 Resolver DNS Firewall domain list.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows how to get a firewall domain list from its ID.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.route53.getResolverFirewallDomainList({
    firewallDomainListId: "rslvr-fdl-example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_firewall_domain_list(firewall_domain_list_id="rslvr-fdl-example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Route53.GetResolverFirewallDomainList.Invoke(new()
    {
        FirewallDomainListId = "rslvr-fdl-example",
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
		_, err := route53.LookupResolverFirewallDomainList(ctx, &route53.LookupResolverFirewallDomainListArgs{
			FirewallDomainListId: "rslvr-fdl-example",
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
import com.pulumi.aws.route53.inputs.GetResolverFirewallDomainListArgs;
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
        final var example = Route53Functions.getResolverFirewallDomainList(GetResolverFirewallDomainListArgs.builder()
            .firewallDomainListId("rslvr-fdl-example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:route53:getResolverFirewallDomainList
      Arguments:
        firewallDomainListId: rslvr-fdl-example
```
{{% /example %}}
{{% /examples %}}