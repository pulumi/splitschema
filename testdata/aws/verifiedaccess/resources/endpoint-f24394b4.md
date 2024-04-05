Resource for managing an AWS EC2 (Elastic Compute Cloud) Verified Access Endpoint.

{{% examples %}}
## Example Usage
{{% example %}}
### ALB Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.Endpoint("example", {
    applicationDomain: "example.com",
    attachmentType: "vpc",
    description: "example",
    domainCertificateArn: aws_acm_certificate.example.arn,
    endpointDomainPrefix: "example",
    endpointType: "load-balancer",
    loadBalancerOptions: {
        loadBalancerArn: aws_lb.example.arn,
        port: 443,
        protocol: "https",
        subnetIds: .map(subnet => (subnet.id)),
    },
    securityGroupIds: [aws_security_group.example.id],
    verifiedAccessGroupId: aws_verifiedaccess_group.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.Endpoint("example",
    application_domain="example.com",
    attachment_type="vpc",
    description="example",
    domain_certificate_arn=aws_acm_certificate["example"]["arn"],
    endpoint_domain_prefix="example",
    endpoint_type="load-balancer",
    load_balancer_options=aws.verifiedaccess.EndpointLoadBalancerOptionsArgs(
        load_balancer_arn=aws_lb["example"]["arn"],
        port=443,
        protocol="https",
        subnet_ids=[subnet["id"] for subnet in aws_subnet["public"]],
    ),
    security_group_ids=[aws_security_group["example"]["id"]],
    verified_access_group_id=aws_verifiedaccess_group["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.Endpoint("example", new()
    {
        ApplicationDomain = "example.com",
        AttachmentType = "vpc",
        Description = "example",
        DomainCertificateArn = aws_acm_certificate.Example.Arn,
        EndpointDomainPrefix = "example",
        EndpointType = "load-balancer",
        LoadBalancerOptions = new Aws.VerifiedAccess.Inputs.EndpointLoadBalancerOptionsArgs
        {
            LoadBalancerArn = aws_lb.Example.Arn,
            Port = 443,
            Protocol = "https",
            SubnetIds = .Select(subnet => 
            {
                return subnet.Id;
            }).ToList(),
        },
        SecurityGroupIds = new[]
        {
            aws_security_group.Example.Id,
        },
        VerifiedAccessGroupId = aws_verifiedaccess_group.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewEndpoint(ctx, "example", &verifiedaccess.EndpointArgs{
			ApplicationDomain:    pulumi.String("example.com"),
			AttachmentType:       pulumi.String("vpc"),
			Description:          pulumi.String("example"),
			DomainCertificateArn: pulumi.Any(aws_acm_certificate.Example.Arn),
			EndpointDomainPrefix: pulumi.String("example"),
			EndpointType:         pulumi.String("load-balancer"),
			LoadBalancerOptions: &verifiedaccess.EndpointLoadBalancerOptionsArgs{
				LoadBalancerArn: pulumi.Any(aws_lb.Example.Arn),
				Port:            pulumi.Int(443),
				Protocol:        pulumi.String("https"),
				SubnetIds:       "TODO: For expression",
			},
			SecurityGroupIds: pulumi.StringArray{
				aws_security_group.Example.Id,
			},
			VerifiedAccessGroupId: pulumi.Any(aws_verifiedaccess_group.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /example %}}
{{% example %}}
### Network Interface Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.Endpoint("example", {
    applicationDomain: "example.com",
    attachmentType: "vpc",
    description: "example",
    domainCertificateArn: aws_acm_certificate.example.arn,
    endpointDomainPrefix: "example",
    endpointType: "network-interface",
    networkInterfaceOptions: {
        networkInterfaceId: aws_network_interface.example.id,
        port: 443,
        protocol: "https",
    },
    securityGroupIds: [aws_security_group.example.id],
    verifiedAccessGroupId: aws_verifiedaccess_group.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.Endpoint("example",
    application_domain="example.com",
    attachment_type="vpc",
    description="example",
    domain_certificate_arn=aws_acm_certificate["example"]["arn"],
    endpoint_domain_prefix="example",
    endpoint_type="network-interface",
    network_interface_options=aws.verifiedaccess.EndpointNetworkInterfaceOptionsArgs(
        network_interface_id=aws_network_interface["example"]["id"],
        port=443,
        protocol="https",
    ),
    security_group_ids=[aws_security_group["example"]["id"]],
    verified_access_group_id=aws_verifiedaccess_group["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.Endpoint("example", new()
    {
        ApplicationDomain = "example.com",
        AttachmentType = "vpc",
        Description = "example",
        DomainCertificateArn = aws_acm_certificate.Example.Arn,
        EndpointDomainPrefix = "example",
        EndpointType = "network-interface",
        NetworkInterfaceOptions = new Aws.VerifiedAccess.Inputs.EndpointNetworkInterfaceOptionsArgs
        {
            NetworkInterfaceId = aws_network_interface.Example.Id,
            Port = 443,
            Protocol = "https",
        },
        SecurityGroupIds = new[]
        {
            aws_security_group.Example.Id,
        },
        VerifiedAccessGroupId = aws_verifiedaccess_group.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewEndpoint(ctx, "example", &verifiedaccess.EndpointArgs{
			ApplicationDomain:    pulumi.String("example.com"),
			AttachmentType:       pulumi.String("vpc"),
			Description:          pulumi.String("example"),
			DomainCertificateArn: pulumi.Any(aws_acm_certificate.Example.Arn),
			EndpointDomainPrefix: pulumi.String("example"),
			EndpointType:         pulumi.String("network-interface"),
			NetworkInterfaceOptions: &verifiedaccess.EndpointNetworkInterfaceOptionsArgs{
				NetworkInterfaceId: pulumi.Any(aws_network_interface.Example.Id),
				Port:               pulumi.Int(443),
				Protocol:           pulumi.String("https"),
			},
			SecurityGroupIds: pulumi.StringArray{
				aws_security_group.Example.Id,
			},
			VerifiedAccessGroupId: pulumi.Any(aws_verifiedaccess_group.Example.Id),
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
import com.pulumi.aws.verifiedaccess.Endpoint;
import com.pulumi.aws.verifiedaccess.EndpointArgs;
import com.pulumi.aws.verifiedaccess.inputs.EndpointNetworkInterfaceOptionsArgs;
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
            .applicationDomain("example.com")
            .attachmentType("vpc")
            .description("example")
            .domainCertificateArn(aws_acm_certificate.example().arn())
            .endpointDomainPrefix("example")
            .endpointType("network-interface")
            .networkInterfaceOptions(EndpointNetworkInterfaceOptionsArgs.builder()
                .networkInterfaceId(aws_network_interface.example().id())
                .port(443)
                .protocol("https")
                .build())
            .securityGroupIds(aws_security_group.example().id())
            .verifiedAccessGroupId(aws_verifiedaccess_group.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:Endpoint
    properties:
      applicationDomain: example.com
      attachmentType: vpc
      description: example
      domainCertificateArn: ${aws_acm_certificate.example.arn}
      endpointDomainPrefix: example
      endpointType: network-interface
      networkInterfaceOptions:
        networkInterfaceId: ${aws_network_interface.example.id}
        port: 443
        protocol: https
      securityGroupIds:
        - ${aws_security_group.example.id}
      verifiedAccessGroupId: ${aws_verifiedaccess_group.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Verified Access Instances using the

`id`. For example:

```sh
 $ pulumi import aws:verifiedaccess/endpoint:Endpoint example vae-8012925589
```
 