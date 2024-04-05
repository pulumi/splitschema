Provides a Route 53 Resolver DNSSEC config resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpc = new aws.ec2.Vpc("exampleVpc", {
    cidrBlock: "10.0.0.0/16",
    enableDnsSupport: true,
    enableDnsHostnames: true,
});
const exampleResolverDnsSecConfig = new aws.route53.ResolverDnsSecConfig("exampleResolverDnsSecConfig", {resourceId: exampleVpc.id});
```
```python
import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc",
    cidr_block="10.0.0.0/16",
    enable_dns_support=True,
    enable_dns_hostnames=True)
example_resolver_dns_sec_config = aws.route53.ResolverDnsSecConfig("exampleResolverDnsSecConfig", resource_id=example_vpc.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.0.0.0/16",
        EnableDnsSupport = true,
        EnableDnsHostnames = true,
    });

    var exampleResolverDnsSecConfig = new Aws.Route53.ResolverDnsSecConfig("exampleResolverDnsSecConfig", new()
    {
        ResourceId = exampleVpc.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("10.0.0.0/16"),
			EnableDnsSupport:   pulumi.Bool(true),
			EnableDnsHostnames: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewResolverDnsSecConfig(ctx, "exampleResolverDnsSecConfig", &route53.ResolverDnsSecConfigArgs{
			ResourceId: exampleVpc.ID(),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.route53.ResolverDnsSecConfig;
import com.pulumi.aws.route53.ResolverDnsSecConfigArgs;
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
        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .enableDnsSupport(true)
            .enableDnsHostnames(true)
            .build());

        var exampleResolverDnsSecConfig = new ResolverDnsSecConfig("exampleResolverDnsSecConfig", ResolverDnsSecConfigArgs.builder()        
            .resourceId(exampleVpc.id())
            .build());

    }
}
```
```yaml
resources:
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
      enableDnsSupport: true
      enableDnsHostnames: true
  exampleResolverDnsSecConfig:
    type: aws:route53:ResolverDnsSecConfig
    properties:
      resourceId: ${exampleVpc.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import

Route 53 Resolver DNSSEC configs using the Route 53 Resolver DNSSEC config ID. For example:

```sh
 $ pulumi import aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig example rdsc-be1866ecc1683e95
```
 