Provides a Route 53 Resolver config resource.

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
const exampleResolverConfig = new aws.route53.ResolverConfig("exampleResolverConfig", {
    resourceId: exampleVpc.id,
    autodefinedReverseFlag: "DISABLE",
});
```
```python
import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc",
    cidr_block="10.0.0.0/16",
    enable_dns_support=True,
    enable_dns_hostnames=True)
example_resolver_config = aws.route53.ResolverConfig("exampleResolverConfig",
    resource_id=example_vpc.id,
    autodefined_reverse_flag="DISABLE")
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

    var exampleResolverConfig = new Aws.Route53.ResolverConfig("exampleResolverConfig", new()
    {
        ResourceId = exampleVpc.Id,
        AutodefinedReverseFlag = "DISABLE",
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
		_, err = route53.NewResolverConfig(ctx, "exampleResolverConfig", &route53.ResolverConfigArgs{
			ResourceId:             exampleVpc.ID(),
			AutodefinedReverseFlag: pulumi.String("DISABLE"),
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
import com.pulumi.aws.route53.ResolverConfig;
import com.pulumi.aws.route53.ResolverConfigArgs;
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

        var exampleResolverConfig = new ResolverConfig("exampleResolverConfig", ResolverConfigArgs.builder()        
            .resourceId(exampleVpc.id())
            .autodefinedReverseFlag("DISABLE")
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
  exampleResolverConfig:
    type: aws:route53:ResolverConfig
    properties:
      resourceId: ${exampleVpc.id}
      autodefinedReverseFlag: DISABLE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route 53 Resolver configs using the Route 53 Resolver config ID. For example:

```sh
 $ pulumi import aws:route53/resolverConfig:ResolverConfig example rslvr-rc-715aa20c73a23da7
```
 