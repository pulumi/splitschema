Provides additional routes for AWS Client VPN endpoints. For more information on usage, please see the
[AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEndpoint = new aws.ec2clientvpn.Endpoint("exampleEndpoint", {
    description: "Example Client VPN endpoint",
    serverCertificateArn: aws_acm_certificate.example.arn,
    clientCidrBlock: "10.0.0.0/16",
    authenticationOptions: [{
        type: "certificate-authentication",
        rootCertificateChainArn: aws_acm_certificate.example.arn,
    }],
    connectionLogOptions: {
        enabled: false,
    },
});
const exampleNetworkAssociation = new aws.ec2clientvpn.NetworkAssociation("exampleNetworkAssociation", {
    clientVpnEndpointId: exampleEndpoint.id,
    subnetId: aws_subnet.example.id,
});
const exampleRoute = new aws.ec2clientvpn.Route("exampleRoute", {
    clientVpnEndpointId: exampleEndpoint.id,
    destinationCidrBlock: "0.0.0.0/0",
    targetVpcSubnetId: exampleNetworkAssociation.subnetId,
});
```
```python
import pulumi
import pulumi_aws as aws

example_endpoint = aws.ec2clientvpn.Endpoint("exampleEndpoint",
    description="Example Client VPN endpoint",
    server_certificate_arn=aws_acm_certificate["example"]["arn"],
    client_cidr_block="10.0.0.0/16",
    authentication_options=[aws.ec2clientvpn.EndpointAuthenticationOptionArgs(
        type="certificate-authentication",
        root_certificate_chain_arn=aws_acm_certificate["example"]["arn"],
    )],
    connection_log_options=aws.ec2clientvpn.EndpointConnectionLogOptionsArgs(
        enabled=False,
    ))
example_network_association = aws.ec2clientvpn.NetworkAssociation("exampleNetworkAssociation",
    client_vpn_endpoint_id=example_endpoint.id,
    subnet_id=aws_subnet["example"]["id"])
example_route = aws.ec2clientvpn.Route("exampleRoute",
    client_vpn_endpoint_id=example_endpoint.id,
    destination_cidr_block="0.0.0.0/0",
    target_vpc_subnet_id=example_network_association.subnet_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleEndpoint = new Aws.Ec2ClientVpn.Endpoint("exampleEndpoint", new()
    {
        Description = "Example Client VPN endpoint",
        ServerCertificateArn = aws_acm_certificate.Example.Arn,
        ClientCidrBlock = "10.0.0.0/16",
        AuthenticationOptions = new[]
        {
            new Aws.Ec2ClientVpn.Inputs.EndpointAuthenticationOptionArgs
            {
                Type = "certificate-authentication",
                RootCertificateChainArn = aws_acm_certificate.Example.Arn,
            },
        },
        ConnectionLogOptions = new Aws.Ec2ClientVpn.Inputs.EndpointConnectionLogOptionsArgs
        {
            Enabled = false,
        },
    });

    var exampleNetworkAssociation = new Aws.Ec2ClientVpn.NetworkAssociation("exampleNetworkAssociation", new()
    {
        ClientVpnEndpointId = exampleEndpoint.Id,
        SubnetId = aws_subnet.Example.Id,
    });

    var exampleRoute = new Aws.Ec2ClientVpn.Route("exampleRoute", new()
    {
        ClientVpnEndpointId = exampleEndpoint.Id,
        DestinationCidrBlock = "0.0.0.0/0",
        TargetVpcSubnetId = exampleNetworkAssociation.SubnetId,
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
		exampleEndpoint, err := ec2clientvpn.NewEndpoint(ctx, "exampleEndpoint", &ec2clientvpn.EndpointArgs{
			Description:          pulumi.String("Example Client VPN endpoint"),
			ServerCertificateArn: pulumi.Any(aws_acm_certificate.Example.Arn),
			ClientCidrBlock:      pulumi.String("10.0.0.0/16"),
			AuthenticationOptions: ec2clientvpn.EndpointAuthenticationOptionArray{
				&ec2clientvpn.EndpointAuthenticationOptionArgs{
					Type:                    pulumi.String("certificate-authentication"),
					RootCertificateChainArn: pulumi.Any(aws_acm_certificate.Example.Arn),
				},
			},
			ConnectionLogOptions: &ec2clientvpn.EndpointConnectionLogOptionsArgs{
				Enabled: pulumi.Bool(false),
			},
		})
		if err != nil {
			return err
		}
		exampleNetworkAssociation, err := ec2clientvpn.NewNetworkAssociation(ctx, "exampleNetworkAssociation", &ec2clientvpn.NetworkAssociationArgs{
			ClientVpnEndpointId: exampleEndpoint.ID(),
			SubnetId:            pulumi.Any(aws_subnet.Example.Id),
		})
		if err != nil {
			return err
		}
		_, err = ec2clientvpn.NewRoute(ctx, "exampleRoute", &ec2clientvpn.RouteArgs{
			ClientVpnEndpointId:  exampleEndpoint.ID(),
			DestinationCidrBlock: pulumi.String("0.0.0.0/0"),
			TargetVpcSubnetId:    exampleNetworkAssociation.SubnetId,
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
import com.pulumi.aws.ec2clientvpn.Endpoint;
import com.pulumi.aws.ec2clientvpn.EndpointArgs;
import com.pulumi.aws.ec2clientvpn.inputs.EndpointAuthenticationOptionArgs;
import com.pulumi.aws.ec2clientvpn.inputs.EndpointConnectionLogOptionsArgs;
import com.pulumi.aws.ec2clientvpn.NetworkAssociation;
import com.pulumi.aws.ec2clientvpn.NetworkAssociationArgs;
import com.pulumi.aws.ec2clientvpn.Route;
import com.pulumi.aws.ec2clientvpn.RouteArgs;
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
        var exampleEndpoint = new Endpoint("exampleEndpoint", EndpointArgs.builder()        
            .description("Example Client VPN endpoint")
            .serverCertificateArn(aws_acm_certificate.example().arn())
            .clientCidrBlock("10.0.0.0/16")
            .authenticationOptions(EndpointAuthenticationOptionArgs.builder()
                .type("certificate-authentication")
                .rootCertificateChainArn(aws_acm_certificate.example().arn())
                .build())
            .connectionLogOptions(EndpointConnectionLogOptionsArgs.builder()
                .enabled(false)
                .build())
            .build());

        var exampleNetworkAssociation = new NetworkAssociation("exampleNetworkAssociation", NetworkAssociationArgs.builder()        
            .clientVpnEndpointId(exampleEndpoint.id())
            .subnetId(aws_subnet.example().id())
            .build());

        var exampleRoute = new Route("exampleRoute", RouteArgs.builder()        
            .clientVpnEndpointId(exampleEndpoint.id())
            .destinationCidrBlock("0.0.0.0/0")
            .targetVpcSubnetId(exampleNetworkAssociation.subnetId())
            .build());

    }
}
```
```yaml
resources:
  exampleRoute:
    type: aws:ec2clientvpn:Route
    properties:
      clientVpnEndpointId: ${exampleEndpoint.id}
      destinationCidrBlock: 0.0.0.0/0
      targetVpcSubnetId: ${exampleNetworkAssociation.subnetId}
  exampleNetworkAssociation:
    type: aws:ec2clientvpn:NetworkAssociation
    properties:
      clientVpnEndpointId: ${exampleEndpoint.id}
      subnetId: ${aws_subnet.example.id}
  exampleEndpoint:
    type: aws:ec2clientvpn:Endpoint
    properties:
      description: Example Client VPN endpoint
      serverCertificateArn: ${aws_acm_certificate.example.arn}
      clientCidrBlock: 10.0.0.0/16
      authenticationOptions:
        - type: certificate-authentication
          rootCertificateChainArn: ${aws_acm_certificate.example.arn}
      connectionLogOptions:
        enabled: false
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS Client VPN routes using the endpoint ID, target subnet ID, and destination CIDR block. All values are separated by a `,`. For example:

```sh
 $ pulumi import aws:ec2clientvpn/route:Route example cvpn-endpoint-1234567890abcdef,subnet-9876543210fedcba,10.1.0.0/24
```
 