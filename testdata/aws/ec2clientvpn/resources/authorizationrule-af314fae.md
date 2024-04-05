Provides authorization rules for AWS Client VPN endpoints. For more information on usage, please see the
[AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2clientvpn.AuthorizationRule("example", {
    clientVpnEndpointId: aws_ec2_client_vpn_endpoint.example.id,
    targetNetworkCidr: aws_subnet.example.cidr_block,
    authorizeAllGroups: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2clientvpn.AuthorizationRule("example",
    client_vpn_endpoint_id=aws_ec2_client_vpn_endpoint["example"]["id"],
    target_network_cidr=aws_subnet["example"]["cidr_block"],
    authorize_all_groups=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2ClientVpn.AuthorizationRule("example", new()
    {
        ClientVpnEndpointId = aws_ec2_client_vpn_endpoint.Example.Id,
        TargetNetworkCidr = aws_subnet.Example.Cidr_block,
        AuthorizeAllGroups = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2clientvpn"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2clientvpn.NewAuthorizationRule(ctx, "example", &ec2clientvpn.AuthorizationRuleArgs{
			ClientVpnEndpointId: pulumi.Any(aws_ec2_client_vpn_endpoint.Example.Id),
			TargetNetworkCidr:   pulumi.Any(aws_subnet.Example.Cidr_block),
			AuthorizeAllGroups:  pulumi.Bool(true),
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
import com.pulumi.aws.ec2clientvpn.AuthorizationRule;
import com.pulumi.aws.ec2clientvpn.AuthorizationRuleArgs;
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
        var example = new AuthorizationRule("example", AuthorizationRuleArgs.builder()        
            .clientVpnEndpointId(aws_ec2_client_vpn_endpoint.example().id())
            .targetNetworkCidr(aws_subnet.example().cidr_block())
            .authorizeAllGroups(true)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2clientvpn:AuthorizationRule
    properties:
      clientVpnEndpointId: ${aws_ec2_client_vpn_endpoint.example.id}
      targetNetworkCidr: ${aws_subnet.example.cidr_block}
      authorizeAllGroups: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using the endpoint ID, target network CIDR, and group name:

__Using `pulumi import` to import__ AWS Client VPN authorization rules using the endpoint ID and target network CIDR. If there is a specific group name, include that also. All values are separated by a `,`. For example:

Using the endpoint ID and target network CIDR:

```sh
 $ pulumi import aws:ec2clientvpn/authorizationRule:AuthorizationRule example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24
```
 Using the endpoint ID, target network CIDR, and group name:

```sh
 $ pulumi import aws:ec2clientvpn/authorizationRule:AuthorizationRule example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24,team-a
```
 