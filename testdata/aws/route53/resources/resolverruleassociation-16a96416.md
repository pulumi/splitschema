Provides a Route53 Resolver rule association.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53.ResolverRuleAssociation("example", {
    resolverRuleId: aws_route53_resolver_rule.sys.id,
    vpcId: aws_vpc.foo.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.ResolverRuleAssociation("example",
    resolver_rule_id=aws_route53_resolver_rule["sys"]["id"],
    vpc_id=aws_vpc["foo"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53.ResolverRuleAssociation("example", new()
    {
        ResolverRuleId = aws_route53_resolver_rule.Sys.Id,
        VpcId = aws_vpc.Foo.Id,
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
		_, err := route53.NewResolverRuleAssociation(ctx, "example", &route53.ResolverRuleAssociationArgs{
			ResolverRuleId: pulumi.Any(aws_route53_resolver_rule.Sys.Id),
			VpcId:          pulumi.Any(aws_vpc.Foo.Id),
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
import com.pulumi.aws.route53.ResolverRuleAssociation;
import com.pulumi.aws.route53.ResolverRuleAssociationArgs;
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
        var example = new ResolverRuleAssociation("example", ResolverRuleAssociationArgs.builder()        
            .resolverRuleId(aws_route53_resolver_rule.sys().id())
            .vpcId(aws_vpc.foo().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53:ResolverRuleAssociation
    properties:
      resolverRuleId: ${aws_route53_resolver_rule.sys.id}
      vpcId: ${aws_vpc.foo.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 Resolver rule associations using the `id`. For example:

```sh
 $ pulumi import aws:route53/resolverRuleAssociation:ResolverRuleAssociation example rslvr-rrassoc-97242eaf88example
```
 