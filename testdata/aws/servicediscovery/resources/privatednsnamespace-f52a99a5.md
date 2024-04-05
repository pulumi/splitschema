Provides a Service Discovery Private DNS Namespace resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.0.0.0/16"});
const examplePrivateDnsNamespace = new aws.servicediscovery.PrivateDnsNamespace("examplePrivateDnsNamespace", {
    description: "example",
    vpc: exampleVpc.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
example_private_dns_namespace = aws.servicediscovery.PrivateDnsNamespace("examplePrivateDnsNamespace",
    description="example",
    vpc=example_vpc.id)
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
    });

    var examplePrivateDnsNamespace = new Aws.ServiceDiscovery.PrivateDnsNamespace("examplePrivateDnsNamespace", new()
    {
        Description = "example",
        Vpc = exampleVpc.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = servicediscovery.NewPrivateDnsNamespace(ctx, "examplePrivateDnsNamespace", &servicediscovery.PrivateDnsNamespaceArgs{
			Description: pulumi.String("example"),
			Vpc:         exampleVpc.ID(),
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
import com.pulumi.aws.servicediscovery.PrivateDnsNamespace;
import com.pulumi.aws.servicediscovery.PrivateDnsNamespaceArgs;
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
            .build());

        var examplePrivateDnsNamespace = new PrivateDnsNamespace("examplePrivateDnsNamespace", PrivateDnsNamespaceArgs.builder()        
            .description("example")
            .vpc(exampleVpc.id())
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
  examplePrivateDnsNamespace:
    type: aws:servicediscovery:PrivateDnsNamespace
    properties:
      description: example
      vpc: ${exampleVpc.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Service Discovery Private DNS Namespace using the namespace ID and VPC ID. For example:

```sh
 $ pulumi import aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace example 0123456789:vpc-123345
```
 