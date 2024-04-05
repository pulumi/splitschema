`aws.route53.ResolverFirewallConfig` provides details about a specific a Route 53 Resolver DNS Firewall config.

This data source allows to find a details about a specific a Route 53 Resolver DNS Firewall config.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows how to get a firewall config using the VPC ID.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.route53.getResolverFirewallConfig({
    resourceId: "vpc-exampleid",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_firewall_config(resource_id="vpc-exampleid")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Route53.GetResolverFirewallConfig.Invoke(new()
    {
        ResourceId = "vpc-exampleid",
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
		_, err := route53.LookupResolverFirewallConfig(ctx, &route53.LookupResolverFirewallConfigArgs{
			ResourceId: "vpc-exampleid",
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
import com.pulumi.aws.route53.inputs.GetResolverFirewallConfigArgs;
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
        final var example = Route53Functions.getResolverFirewallConfig(GetResolverFirewallConfigArgs.builder()
            .resourceId("vpc-exampleid")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:route53:getResolverFirewallConfig
      Arguments:
        resourceId: vpc-exampleid
```
{{% /example %}}
{{% /examples %}}