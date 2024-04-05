`aws.ec2.VpcIpamPool` provides details about an IPAM pool.

This resource can prove useful when an ipam pool was created in another root
module and you need the pool's id as an input variable. For example, pools
can be shared via RAM and used to create vpcs with CIDRs from that pool.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows an account that has only 1 pool, perhaps shared
via RAM, and using that pool id to create a VPC with a CIDR derived from
AWS IPAM.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testVpcIpamPool = aws.ec2.getVpcIpamPool({
    filters: [
        {
            name: "description",
            values: ["*test*"],
        },
        {
            name: "address-family",
            values: ["ipv4"],
        },
    ],
});
const testVpc = new aws.ec2.Vpc("testVpc", {
    ipv4IpamPoolId: testVpcIpamPool.then(testVpcIpamPool => testVpcIpamPool.id),
    ipv4NetmaskLength: 28,
});
```
```python
import pulumi
import pulumi_aws as aws

test_vpc_ipam_pool = aws.ec2.get_vpc_ipam_pool(filters=[
    aws.ec2.GetVpcIpamPoolFilterArgs(
        name="description",
        values=["*test*"],
    ),
    aws.ec2.GetVpcIpamPoolFilterArgs(
        name="address-family",
        values=["ipv4"],
    ),
])
test_vpc = aws.ec2.Vpc("testVpc",
    ipv4_ipam_pool_id=test_vpc_ipam_pool.id,
    ipv4_netmask_length=28)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testVpcIpamPool = Aws.Ec2.GetVpcIpamPool.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetVpcIpamPoolFilterInputArgs
            {
                Name = "description",
                Values = new[]
                {
                    "*test*",
                },
            },
            new Aws.Ec2.Inputs.GetVpcIpamPoolFilterInputArgs
            {
                Name = "address-family",
                Values = new[]
                {
                    "ipv4",
                },
            },
        },
    });

    var testVpc = new Aws.Ec2.Vpc("testVpc", new()
    {
        Ipv4IpamPoolId = testVpcIpamPool.Apply(getVpcIpamPoolResult => getVpcIpamPoolResult.Id),
        Ipv4NetmaskLength = 28,
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
		testVpcIpamPool, err := ec2.LookupVpcIpamPool(ctx, &ec2.LookupVpcIpamPoolArgs{
			Filters: []ec2.GetVpcIpamPoolFilter{
				{
					Name: "description",
					Values: []string{
						"*test*",
					},
				},
				{
					Name: "address-family",
					Values: []string{
						"ipv4",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewVpc(ctx, "testVpc", &ec2.VpcArgs{
			Ipv4IpamPoolId:    *pulumi.String(testVpcIpamPool.Id),
			Ipv4NetmaskLength: pulumi.Int(28),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetVpcIpamPoolArgs;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
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
        final var testVpcIpamPool = Ec2Functions.getVpcIpamPool(GetVpcIpamPoolArgs.builder()
            .filters(            
                GetVpcIpamPoolFilterArgs.builder()
                    .name("description")
                    .values("*test*")
                    .build(),
                GetVpcIpamPoolFilterArgs.builder()
                    .name("address-family")
                    .values("ipv4")
                    .build())
            .build());

        var testVpc = new Vpc("testVpc", VpcArgs.builder()        
            .ipv4IpamPoolId(testVpcIpamPool.applyValue(getVpcIpamPoolResult -> getVpcIpamPoolResult.id()))
            .ipv4NetmaskLength(28)
            .build());

    }
}
```
```yaml
resources:
  testVpc:
    type: aws:ec2:Vpc
    properties:
      ipv4IpamPoolId: ${testVpcIpamPool.id}
      ipv4NetmaskLength: 28
variables:
  testVpcIpamPool:
    fn::invoke:
      Function: aws:ec2:getVpcIpamPool
      Arguments:
        filters:
          - name: description
            values:
              - '*test*'
          - name: address-family
            values:
              - ipv4
```
{{% /example %}}
{{% /examples %}}