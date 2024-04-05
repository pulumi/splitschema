Provides a resource to associate additional IPv4 CIDR blocks with a VPC.

When a VPC is created, a primary IPv4 CIDR block for the VPC must be specified.
The `aws.ec2.VpcIpv4CidrBlockAssociation` resource allows further IPv4 CIDR blocks to be added to the VPC.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
const secondaryCidr = new aws.ec2.VpcIpv4CidrBlockAssociation("secondaryCidr", {
    vpcId: main.id,
    cidrBlock: "172.20.0.0/16",
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
secondary_cidr = aws.ec2.VpcIpv4CidrBlockAssociation("secondaryCidr",
    vpc_id=main.id,
    cidr_block="172.20.0.0/16")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var secondaryCidr = new Aws.Ec2.VpcIpv4CidrBlockAssociation("secondaryCidr", new()
    {
        VpcId = main.Id,
        CidrBlock = "172.20.0.0/16",
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
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcIpv4CidrBlockAssociation(ctx, "secondaryCidr", &ec2.VpcIpv4CidrBlockAssociationArgs{
			VpcId:     main.ID(),
			CidrBlock: pulumi.String("172.20.0.0/16"),
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
import com.pulumi.aws.ec2.VpcIpv4CidrBlockAssociation;
import com.pulumi.aws.ec2.VpcIpv4CidrBlockAssociationArgs;
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
        var main = new Vpc("main", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var secondaryCidr = new VpcIpv4CidrBlockAssociation("secondaryCidr", VpcIpv4CidrBlockAssociationArgs.builder()        
            .vpcId(main.id())
            .cidrBlock("172.20.0.0/16")
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  secondaryCidr:
    type: aws:ec2:VpcIpv4CidrBlockAssociation
    properties:
      vpcId: ${main.id}
      cidrBlock: 172.20.0.0/16
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_vpc_ipv4_cidr_block_association` using the VPC CIDR Association ID. For example:

```sh
 $ pulumi import aws:ec2/vpcIpv4CidrBlockAssociation:VpcIpv4CidrBlockAssociation example vpc-cidr-assoc-xxxxxxxx
```
 