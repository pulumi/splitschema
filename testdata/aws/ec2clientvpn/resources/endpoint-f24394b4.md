Provides an AWS Client VPN endpoint for OpenVPN clients. For more information on usage, please see the
[AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2clientvpn.Endpoint("example", {
    description: "clientvpn-example",
    serverCertificateArn: aws_acm_certificate.cert.arn,
    clientCidrBlock: "10.0.0.0/16",
    authenticationOptions: [{
        type: "certificate-authentication",
        rootCertificateChainArn: aws_acm_certificate.root_cert.arn,
    }],
    connectionLogOptions: {
        enabled: true,
        cloudwatchLogGroup: aws_cloudwatch_log_group.lg.name,
        cloudwatchLogStream: aws_cloudwatch_log_stream.ls.name,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2clientvpn.Endpoint("example",
    description="clientvpn-example",
    server_certificate_arn=aws_acm_certificate["cert"]["arn"],
    client_cidr_block="10.0.0.0/16",
    authentication_options=[aws.ec2clientvpn.EndpointAuthenticationOptionArgs(
        type="certificate-authentication",
        root_certificate_chain_arn=aws_acm_certificate["root_cert"]["arn"],
    )],
    connection_log_options=aws.ec2clientvpn.EndpointConnectionLogOptionsArgs(
        enabled=True,
        cloudwatch_log_group=aws_cloudwatch_log_group["lg"]["name"],
        cloudwatch_log_stream=aws_cloudwatch_log_stream["ls"]["name"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2ClientVpn.Endpoint("example", new()
    {
        Description = "clientvpn-example",
        ServerCertificateArn = aws_acm_certificate.Cert.Arn,
        ClientCidrBlock = "10.0.0.0/16",
        AuthenticationOptions = new[]
        {
            new Aws.Ec2ClientVpn.Inputs.EndpointAuthenticationOptionArgs
            {
                Type = "certificate-authentication",
                RootCertificateChainArn = aws_acm_certificate.Root_cert.Arn,
            },
        },
        ConnectionLogOptions = new Aws.Ec2ClientVpn.Inputs.EndpointConnectionLogOptionsArgs
        {
            Enabled = true,
            CloudwatchLogGroup = aws_cloudwatch_log_group.Lg.Name,
            CloudwatchLogStream = aws_cloudwatch_log_stream.Ls.Name,
        },
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
		_, err := ec2clientvpn.NewEndpoint(ctx, "example", &ec2clientvpn.EndpointArgs{
			Description:          pulumi.String("clientvpn-example"),
			ServerCertificateArn: pulumi.Any(aws_acm_certificate.Cert.Arn),
			ClientCidrBlock:      pulumi.String("10.0.0.0/16"),
			AuthenticationOptions: ec2clientvpn.EndpointAuthenticationOptionArray{
				&ec2clientvpn.EndpointAuthenticationOptionArgs{
					Type:                    pulumi.String("certificate-authentication"),
					RootCertificateChainArn: pulumi.Any(aws_acm_certificate.Root_cert.Arn),
				},
			},
			ConnectionLogOptions: &ec2clientvpn.EndpointConnectionLogOptionsArgs{
				Enabled:             pulumi.Bool(true),
				CloudwatchLogGroup:  pulumi.Any(aws_cloudwatch_log_group.Lg.Name),
				CloudwatchLogStream: pulumi.Any(aws_cloudwatch_log_stream.Ls.Name),
			},
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
        var example = new Endpoint("example", EndpointArgs.builder()        
            .description("clientvpn-example")
            .serverCertificateArn(aws_acm_certificate.cert().arn())
            .clientCidrBlock("10.0.0.0/16")
            .authenticationOptions(EndpointAuthenticationOptionArgs.builder()
                .type("certificate-authentication")
                .rootCertificateChainArn(aws_acm_certificate.root_cert().arn())
                .build())
            .connectionLogOptions(EndpointConnectionLogOptionsArgs.builder()
                .enabled(true)
                .cloudwatchLogGroup(aws_cloudwatch_log_group.lg().name())
                .cloudwatchLogStream(aws_cloudwatch_log_stream.ls().name())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2clientvpn:Endpoint
    properties:
      description: clientvpn-example
      serverCertificateArn: ${aws_acm_certificate.cert.arn}
      clientCidrBlock: 10.0.0.0/16
      authenticationOptions:
        - type: certificate-authentication
          rootCertificateChainArn: ${aws_acm_certificate.root_cert.arn}
      connectionLogOptions:
        enabled: true
        cloudwatchLogGroup: ${aws_cloudwatch_log_group.lg.name}
        cloudwatchLogStream: ${aws_cloudwatch_log_stream.ls.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS Client VPN endpoints using the `id` value found via `aws ec2 describe-client-vpn-endpoints`. For example:

```sh
 $ pulumi import aws:ec2clientvpn/endpoint:Endpoint example cvpn-endpoint-0ac3a1abbccddd666
```
 