Provides a resource to associate additional IPv6 CIDR blocks with a VPC.

The `aws.ec2.VpcIpv6CidrBlockAssociation` resource allows IPv6 CIDR blocks to be added to the VPC.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testVpc = new aws.ec2.Vpc("testVpc", {cidrBlock: "10.0.0.0/16"});
const testVpcIpv6CidrBlockAssociation = new aws.ec2.VpcIpv6CidrBlockAssociation("testVpcIpv6CidrBlockAssociation", {
    ipv6IpamPoolId: aws_vpc_ipam_pool.test.id,
    vpcId: testVpc.id,
});
```
```python
import pulumi
import pulumi_aws as aws

test_vpc = aws.ec2.Vpc("testVpc", cidr_block="10.0.0.0/16")
test_vpc_ipv6_cidr_block_association = aws.ec2.VpcIpv6CidrBlockAssociation("testVpcIpv6CidrBlockAssociation",
    ipv6_ipam_pool_id=aws_vpc_ipam_pool["test"]["id"],
    vpc_id=test_vpc.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testVpc = new Aws.Ec2.Vpc("testVpc", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var testVpcIpv6CidrBlockAssociation = new Aws.Ec2.VpcIpv6CidrBlockAssociation("testVpcIpv6CidrBlockAssociation", new()
    {
        Ipv6IpamPoolId = aws_vpc_ipam_pool.Test.Id,
        VpcId = testVpc.Id,
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
		testVpc, err := ec2.NewVpc(ctx, "testVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcIpv6CidrBlockAssociation(ctx, "testVpcIpv6CidrBlockAssociation", &ec2.VpcIpv6CidrBlockAssociationArgs{
			Ipv6IpamPoolId: pulumi.Any(aws_vpc_ipam_pool.Test.Id),
			VpcId:          testVpc.ID(),
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
import com.pulumi.aws.ec2.VpcIpv6CidrBlockAssociation;
import com.pulumi.aws.ec2.VpcIpv6CidrBlockAssociationArgs;
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
        var testVpc = new Vpc("testVpc", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var testVpcIpv6CidrBlockAssociation = new VpcIpv6CidrBlockAssociation("testVpcIpv6CidrBlockAssociation", VpcIpv6CidrBlockAssociationArgs.builder()        
            .ipv6IpamPoolId(aws_vpc_ipam_pool.test().id())
            .vpcId(testVpc.id())
            .build());

    }
}
```
```yaml
resources:
  testVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  testVpcIpv6CidrBlockAssociation:
    type: aws:ec2:VpcIpv6CidrBlockAssociation
    properties:
      ipv6IpamPoolId: ${aws_vpc_ipam_pool.test.id}
      vpcId: ${testVpc.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_vpc_ipv6_cidr_block_association` using the VPC CIDR Association ID. For example:

```sh
 $ pulumi import aws:ec2/vpcIpv6CidrBlockAssociation:VpcIpv6CidrBlockAssociation example vpc-cidr-assoc-xxxxxxxx
```
 