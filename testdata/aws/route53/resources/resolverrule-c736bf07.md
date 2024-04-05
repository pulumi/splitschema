Provides a Route53 Resolver rule.

{{% examples %}}
## Example Usage
{{% example %}}
### System rule

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sys = new aws.route53.ResolverRule("sys", {
    domainName: "subdomain.example.com",
    ruleType: "SYSTEM",
});
```
```python
import pulumi
import pulumi_aws as aws

sys = aws.route53.ResolverRule("sys",
    domain_name="subdomain.example.com",
    rule_type="SYSTEM")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var sys = new Aws.Route53.ResolverRule("sys", new()
    {
        DomainName = "subdomain.example.com",
        RuleType = "SYSTEM",
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
		_, err := route53.NewResolverRule(ctx, "sys", &route53.ResolverRuleArgs{
			DomainName: pulumi.String("subdomain.example.com"),
			RuleType:   pulumi.String("SYSTEM"),
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
import com.pulumi.aws.route53.ResolverRule;
import com.pulumi.aws.route53.ResolverRuleArgs;
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
        var sys = new ResolverRule("sys", ResolverRuleArgs.builder()        
            .domainName("subdomain.example.com")
            .ruleType("SYSTEM")
            .build());

    }
}
```
```yaml
resources:
  sys:
    type: aws:route53:ResolverRule
    properties:
      domainName: subdomain.example.com
      ruleType: SYSTEM
```
{{% /example %}}
{{% example %}}
### Forward rule

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fwd = new aws.route53.ResolverRule("fwd", {
    domainName: "example.com",
    ruleType: "FORWARD",
    resolverEndpointId: aws_route53_resolver_endpoint.foo.id,
    targetIps: [{
        ip: "123.45.67.89",
    }],
    tags: {
        Environment: "Prod",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

fwd = aws.route53.ResolverRule("fwd",
    domain_name="example.com",
    rule_type="FORWARD",
    resolver_endpoint_id=aws_route53_resolver_endpoint["foo"]["id"],
    target_ips=[aws.route53.ResolverRuleTargetIpArgs(
        ip="123.45.67.89",
    )],
    tags={
        "Environment": "Prod",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fwd = new Aws.Route53.ResolverRule("fwd", new()
    {
        DomainName = "example.com",
        RuleType = "FORWARD",
        ResolverEndpointId = aws_route53_resolver_endpoint.Foo.Id,
        TargetIps = new[]
        {
            new Aws.Route53.Inputs.ResolverRuleTargetIpArgs
            {
                Ip = "123.45.67.89",
            },
        },
        Tags = 
        {
            { "Environment", "Prod" },
        },
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
		_, err := route53.NewResolverRule(ctx, "fwd", &route53.ResolverRuleArgs{
			DomainName:         pulumi.String("example.com"),
			RuleType:           pulumi.String("FORWARD"),
			ResolverEndpointId: pulumi.Any(aws_route53_resolver_endpoint.Foo.Id),
			TargetIps: route53.ResolverRuleTargetIpArray{
				&route53.ResolverRuleTargetIpArgs{
					Ip: pulumi.String("123.45.67.89"),
				},
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Prod"),
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
import com.pulumi.aws.route53.ResolverRule;
import com.pulumi.aws.route53.ResolverRuleArgs;
import com.pulumi.aws.route53.inputs.ResolverRuleTargetIpArgs;
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
        var fwd = new ResolverRule("fwd", ResolverRuleArgs.builder()        
            .domainName("example.com")
            .ruleType("FORWARD")
            .resolverEndpointId(aws_route53_resolver_endpoint.foo().id())
            .targetIps(ResolverRuleTargetIpArgs.builder()
                .ip("123.45.67.89")
                .build())
            .tags(Map.of("Environment", "Prod"))
            .build());

    }
}
```
```yaml
resources:
  fwd:
    type: aws:route53:ResolverRule
    properties:
      domainName: example.com
      ruleType: FORWARD
      resolverEndpointId: ${aws_route53_resolver_endpoint.foo.id}
      targetIps:
        - ip: 123.45.67.89
      tags:
        Environment: Prod
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 Resolver rules using the `id`. For example:

```sh
 $ pulumi import aws:route53/resolverRule:ResolverRule sys rslvr-rr-0123456789abcdef0
```
 