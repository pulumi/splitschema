Provides a VPC DHCP Options Association resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dnsResolver = new aws.ec2.VpcDhcpOptionsAssociation("dnsResolver", {
    vpcId: aws_vpc.foo.id,
    dhcpOptionsId: aws_vpc_dhcp_options.foo.id,
});
```
```python
import pulumi
import pulumi_aws as aws

dns_resolver = aws.ec2.VpcDhcpOptionsAssociation("dnsResolver",
    vpc_id=aws_vpc["foo"]["id"],
    dhcp_options_id=aws_vpc_dhcp_options["foo"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var dnsResolver = new Aws.Ec2.VpcDhcpOptionsAssociation("dnsResolver", new()
    {
        VpcId = aws_vpc.Foo.Id,
        DhcpOptionsId = aws_vpc_dhcp_options.Foo.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewVpcDhcpOptionsAssociation(ctx, "dnsResolver", &ec2.VpcDhcpOptionsAssociationArgs{
			VpcId:         pulumi.Any(aws_vpc.Foo.Id),
			DhcpOptionsId: pulumi.Any(aws_vpc_dhcp_options.Foo.Id),
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
import com.pulumi.aws.ec2.VpcDhcpOptionsAssociation;
import com.pulumi.aws.ec2.VpcDhcpOptionsAssociationArgs;
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
        var dnsResolver = new VpcDhcpOptionsAssociation("dnsResolver", VpcDhcpOptionsAssociationArgs.builder()        
            .vpcId(aws_vpc.foo().id())
            .dhcpOptionsId(aws_vpc_dhcp_options.foo().id())
            .build());

    }
}
```
```yaml
resources:
  dnsResolver:
    type: aws:ec2:VpcDhcpOptionsAssociation
    properties:
      vpcId: ${aws_vpc.foo.id}
      dhcpOptionsId: ${aws_vpc_dhcp_options.foo.id}
```
{{% /example %}}
{{% /examples %}}
## Remarks

* You can only associate one DHCP Options Set to a given VPC ID.
* Removing the DHCP Options Association automatically sets AWS's `default` DHCP Options Set to the VPC.


## Import

Using `pulumi import`, import DHCP associations using the VPC ID associated with the options. For example:

```sh
 $ pulumi import aws:ec2/vpcDhcpOptionsAssociation:VpcDhcpOptionsAssociation imported vpc-0f001273ec18911b1
```
 