Get an existing AWS Customer Gateway.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = aws.ec2.getCustomerGateway({
    filters: [{
        name: "tag:Name",
        values: ["foo-prod"],
    }],
});
const main = new aws.ec2.VpnGateway("main", {
    vpcId: aws_vpc.main.id,
    amazonSideAsn: "7224",
});
const transit = new aws.ec2.VpnConnection("transit", {
    vpnGatewayId: main.id,
    customerGatewayId: foo.then(foo => foo.id),
    type: foo.then(foo => foo.type),
    staticRoutesOnly: false,
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ec2.get_customer_gateway(filters=[aws.ec2.GetCustomerGatewayFilterArgs(
    name="tag:Name",
    values=["foo-prod"],
)])
main = aws.ec2.VpnGateway("main",
    vpc_id=aws_vpc["main"]["id"],
    amazon_side_asn="7224")
transit = aws.ec2.VpnConnection("transit",
    vpn_gateway_id=main.id,
    customer_gateway_id=foo.id,
    type=foo.type,
    static_routes_only=False)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = Aws.Ec2.GetCustomerGateway.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetCustomerGatewayFilterInputArgs
            {
                Name = "tag:Name",
                Values = new[]
                {
                    "foo-prod",
                },
            },
        },
    });

    var main = new Aws.Ec2.VpnGateway("main", new()
    {
        VpcId = aws_vpc.Main.Id,
        AmazonSideAsn = "7224",
    });

    var transit = new Aws.Ec2.VpnConnection("transit", new()
    {
        VpnGatewayId = main.Id,
        CustomerGatewayId = foo.Apply(getCustomerGatewayResult => getCustomerGatewayResult.Id),
        Type = foo.Apply(getCustomerGatewayResult => getCustomerGatewayResult.Type),
        StaticRoutesOnly = false,
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
		foo, err := ec2.LookupCustomerGateway(ctx, &ec2.LookupCustomerGatewayArgs{
			Filters: []ec2.GetCustomerGatewayFilter{
				{
					Name: "tag:Name",
					Values: []string{
						"foo-prod",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		main, err := ec2.NewVpnGateway(ctx, "main", &ec2.VpnGatewayArgs{
			VpcId:         pulumi.Any(aws_vpc.Main.Id),
			AmazonSideAsn: pulumi.String("7224"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpnConnection(ctx, "transit", &ec2.VpnConnectionArgs{
			VpnGatewayId:      main.ID(),
			CustomerGatewayId: *pulumi.String(foo.Id),
			Type:              *pulumi.String(foo.Type),
			StaticRoutesOnly:  pulumi.Bool(false),
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
import com.pulumi.aws.ec2.inputs.GetCustomerGatewayArgs;
import com.pulumi.aws.ec2.VpnGateway;
import com.pulumi.aws.ec2.VpnGatewayArgs;
import com.pulumi.aws.ec2.VpnConnection;
import com.pulumi.aws.ec2.VpnConnectionArgs;
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
        final var foo = Ec2Functions.getCustomerGateway(GetCustomerGatewayArgs.builder()
            .filters(GetCustomerGatewayFilterArgs.builder()
                .name("tag:Name")
                .values("foo-prod")
                .build())
            .build());

        var main = new VpnGateway("main", VpnGatewayArgs.builder()        
            .vpcId(aws_vpc.main().id())
            .amazonSideAsn(7224)
            .build());

        var transit = new VpnConnection("transit", VpnConnectionArgs.builder()        
            .vpnGatewayId(main.id())
            .customerGatewayId(foo.applyValue(getCustomerGatewayResult -> getCustomerGatewayResult.id()))
            .type(foo.applyValue(getCustomerGatewayResult -> getCustomerGatewayResult.type()))
            .staticRoutesOnly(false)
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:ec2:VpnGateway
    properties:
      vpcId: ${aws_vpc.main.id}
      amazonSideAsn: 7224
  transit:
    type: aws:ec2:VpnConnection
    properties:
      vpnGatewayId: ${main.id}
      customerGatewayId: ${foo.id}
      type: ${foo.type}
      staticRoutesOnly: false
variables:
  foo:
    fn::invoke:
      Function: aws:ec2:getCustomerGateway
      Arguments:
        filters:
          - name: tag:Name
            values:
              - foo-prod
```
{{% /example %}}
{{% /examples %}}