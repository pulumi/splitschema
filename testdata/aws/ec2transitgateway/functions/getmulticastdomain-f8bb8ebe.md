Get information on an EC2 Transit Gateway Multicast Domain.

{{% examples %}}
## Example Usage
{{% example %}}
### By Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getMulticastDomain({
    filters: [{
        name: "transit-gateway-multicast-domain-id",
        values: ["tgw-mcast-domain-12345678"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_multicast_domain(filters=[aws.ec2transitgateway.GetMulticastDomainFilterArgs(
    name="transit-gateway-multicast-domain-id",
    values=["tgw-mcast-domain-12345678"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetMulticastDomain.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2TransitGateway.Inputs.GetMulticastDomainFilterInputArgs
            {
                Name = "transit-gateway-multicast-domain-id",
                Values = new[]
                {
                    "tgw-mcast-domain-12345678",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.LookupMulticastDomain(ctx, &ec2transitgateway.LookupMulticastDomainArgs{
			Filters: []ec2transitgateway.GetMulticastDomainFilter{
				{
					Name: "transit-gateway-multicast-domain-id",
					Values: []string{
						"tgw-mcast-domain-12345678",
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
import com.pulumi.aws.ec2transitgateway.Ec2transitgatewayFunctions;
import com.pulumi.aws.ec2transitgateway.inputs.GetMulticastDomainArgs;
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
        final var example = Ec2transitgatewayFunctions.getMulticastDomain(GetMulticastDomainArgs.builder()
            .filters(GetMulticastDomainFilterArgs.builder()
                .name("transit-gateway-multicast-domain-id")
                .values("tgw-mcast-domain-12345678")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getMulticastDomain
      Arguments:
        filters:
          - name: transit-gateway-multicast-domain-id
            values:
              - tgw-mcast-domain-12345678
```
{{% /example %}}
{{% example %}}
### By Identifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getMulticastDomain({
    transitGatewayMulticastDomainId: "tgw-mcast-domain-12345678",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_multicast_domain(transit_gateway_multicast_domain_id="tgw-mcast-domain-12345678")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetMulticastDomain.Invoke(new()
    {
        TransitGatewayMulticastDomainId = "tgw-mcast-domain-12345678",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.LookupMulticastDomain(ctx, &ec2transitgateway.LookupMulticastDomainArgs{
			TransitGatewayMulticastDomainId: pulumi.StringRef("tgw-mcast-domain-12345678"),
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
import com.pulumi.aws.ec2transitgateway.Ec2transitgatewayFunctions;
import com.pulumi.aws.ec2transitgateway.inputs.GetMulticastDomainArgs;
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
        final var example = Ec2transitgatewayFunctions.getMulticastDomain(GetMulticastDomainArgs.builder()
            .transitGatewayMulticastDomainId("tgw-mcast-domain-12345678")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getMulticastDomain
      Arguments:
        transitGatewayMulticastDomainId: tgw-mcast-domain-12345678
```
{{% /example %}}
{{% /examples %}}