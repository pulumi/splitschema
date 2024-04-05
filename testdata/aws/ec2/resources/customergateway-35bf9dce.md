Provides a customer gateway inside a VPC. These objects can be connected to VPN gateways via VPN connections, and allow you to establish tunnels between your network and the VPC.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.CustomerGateway("main", {
    bgpAsn: "65000",
    ipAddress: "172.83.124.10",
    tags: {
        Name: "main-customer-gateway",
    },
    type: "ipsec.1",
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.CustomerGateway("main",
    bgp_asn="65000",
    ip_address="172.83.124.10",
    tags={
        "Name": "main-customer-gateway",
    },
    type="ipsec.1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.CustomerGateway("main", new()
    {
        BgpAsn = "65000",
        IpAddress = "172.83.124.10",
        Tags = 
        {
            { "Name", "main-customer-gateway" },
        },
        Type = "ipsec.1",
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
		_, err := ec2.NewCustomerGateway(ctx, "main", &ec2.CustomerGatewayArgs{
			BgpAsn:    pulumi.String("65000"),
			IpAddress: pulumi.String("172.83.124.10"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main-customer-gateway"),
			},
			Type: pulumi.String("ipsec.1"),
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
import com.pulumi.aws.ec2.CustomerGateway;
import com.pulumi.aws.ec2.CustomerGatewayArgs;
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
        var main = new CustomerGateway("main", CustomerGatewayArgs.builder()        
            .bgpAsn(65000)
            .ipAddress("172.83.124.10")
            .tags(Map.of("Name", "main-customer-gateway"))
            .type("ipsec.1")
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:ec2:CustomerGateway
    properties:
      bgpAsn: 65000
      ipAddress: 172.83.124.10
      tags:
        Name: main-customer-gateway
      type: ipsec.1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Customer Gateways using the `id`. For example:

```sh
 $ pulumi import aws:ec2/customerGateway:CustomerGateway main cgw-b4dc3961
```
 