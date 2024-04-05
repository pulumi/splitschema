`aws.ec2.getVpcIpamPools` provides details about IPAM pools.

This resource can prove useful when IPAM pools are created in another root
module and you need the pool ids as input variables. For example, pools
can be shared via RAM and used to create vpcs with CIDRs from that pool.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.ec2.getVpcIpamPools({
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
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ec2.get_vpc_ipam_pools(filters=[
    aws.ec2.GetVpcIpamPoolsFilterArgs(
        name="description",
        values=["*test*"],
    ),
    aws.ec2.GetVpcIpamPoolsFilterArgs(
        name="address-family",
        values=["ipv4"],
    ),
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Ec2.GetVpcIpamPools.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetVpcIpamPoolsFilterInputArgs
            {
                Name = "description",
                Values = new[]
                {
                    "*test*",
                },
            },
            new Aws.Ec2.Inputs.GetVpcIpamPoolsFilterInputArgs
            {
                Name = "address-family",
                Values = new[]
                {
                    "ipv4",
                },
            },
        },
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
		_, err := ec2.GetVpcIpamPools(ctx, &ec2.GetVpcIpamPoolsArgs{
			Filters: []ec2.GetVpcIpamPoolsFilter{
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
import com.pulumi.aws.ec2.inputs.GetVpcIpamPoolsArgs;
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
        final var test = Ec2Functions.getVpcIpamPools(GetVpcIpamPoolsArgs.builder()
            .filters(            
                GetVpcIpamPoolsFilterArgs.builder()
                    .name("description")
                    .values("*test*")
                    .build(),
                GetVpcIpamPoolsFilterArgs.builder()
                    .name("address-family")
                    .values("ipv4")
                    .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:ec2:getVpcIpamPools
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