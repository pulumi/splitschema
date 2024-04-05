Provides a resource to accept a pending VPC Endpoint Connection accept request to VPC Endpoint Service.

{{% examples %}}
## Example Usage
{{% example %}}
### Accept cross-account request

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpcEndpointService = new aws.ec2.VpcEndpointService("exampleVpcEndpointService", {
    acceptanceRequired: false,
    networkLoadBalancerArns: [aws_lb.example.arn],
});
const exampleVpcEndpoint = new aws.ec2.VpcEndpoint("exampleVpcEndpoint", {
    vpcId: aws_vpc.test_alternate.id,
    serviceName: aws_vpc_endpoint_service.test.service_name,
    vpcEndpointType: "Interface",
    privateDnsEnabled: false,
    securityGroupIds: [aws_security_group.test.id],
}, {
    provider: aws.alternate,
});
const exampleVpcEndpointConnectionAccepter = new aws.ec2.VpcEndpointConnectionAccepter("exampleVpcEndpointConnectionAccepter", {
    vpcEndpointServiceId: exampleVpcEndpointService.id,
    vpcEndpointId: exampleVpcEndpoint.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_vpc_endpoint_service = aws.ec2.VpcEndpointService("exampleVpcEndpointService",
    acceptance_required=False,
    network_load_balancer_arns=[aws_lb["example"]["arn"]])
example_vpc_endpoint = aws.ec2.VpcEndpoint("exampleVpcEndpoint",
    vpc_id=aws_vpc["test_alternate"]["id"],
    service_name=aws_vpc_endpoint_service["test"]["service_name"],
    vpc_endpoint_type="Interface",
    private_dns_enabled=False,
    security_group_ids=[aws_security_group["test"]["id"]],
    opts=pulumi.ResourceOptions(provider=aws["alternate"]))
example_vpc_endpoint_connection_accepter = aws.ec2.VpcEndpointConnectionAccepter("exampleVpcEndpointConnectionAccepter",
    vpc_endpoint_service_id=example_vpc_endpoint_service.id,
    vpc_endpoint_id=example_vpc_endpoint.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleVpcEndpointService = new Aws.Ec2.VpcEndpointService("exampleVpcEndpointService", new()
    {
        AcceptanceRequired = false,
        NetworkLoadBalancerArns = new[]
        {
            aws_lb.Example.Arn,
        },
    });

    var exampleVpcEndpoint = new Aws.Ec2.VpcEndpoint("exampleVpcEndpoint", new()
    {
        VpcId = aws_vpc.Test_alternate.Id,
        ServiceName = aws_vpc_endpoint_service.Test.Service_name,
        VpcEndpointType = "Interface",
        PrivateDnsEnabled = false,
        SecurityGroupIds = new[]
        {
            aws_security_group.Test.Id,
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Alternate,
    });

    var exampleVpcEndpointConnectionAccepter = new Aws.Ec2.VpcEndpointConnectionAccepter("exampleVpcEndpointConnectionAccepter", new()
    {
        VpcEndpointServiceId = exampleVpcEndpointService.Id,
        VpcEndpointId = exampleVpcEndpoint.Id,
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
		exampleVpcEndpointService, err := ec2.NewVpcEndpointService(ctx, "exampleVpcEndpointService", &ec2.VpcEndpointServiceArgs{
			AcceptanceRequired: pulumi.Bool(false),
			NetworkLoadBalancerArns: pulumi.StringArray{
				aws_lb.Example.Arn,
			},
		})
		if err != nil {
			return err
		}
		exampleVpcEndpoint, err := ec2.NewVpcEndpoint(ctx, "exampleVpcEndpoint", &ec2.VpcEndpointArgs{
			VpcId:             pulumi.Any(aws_vpc.Test_alternate.Id),
			ServiceName:       pulumi.Any(aws_vpc_endpoint_service.Test.Service_name),
			VpcEndpointType:   pulumi.String("Interface"),
			PrivateDnsEnabled: pulumi.Bool(false),
			SecurityGroupIds: pulumi.StringArray{
				aws_security_group.Test.Id,
			},
		}, pulumi.Provider(aws.Alternate))
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcEndpointConnectionAccepter(ctx, "exampleVpcEndpointConnectionAccepter", &ec2.VpcEndpointConnectionAccepterArgs{
			VpcEndpointServiceId: exampleVpcEndpointService.ID(),
			VpcEndpointId:        exampleVpcEndpoint.ID(),
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
import com.pulumi.aws.ec2.VpcEndpointService;
import com.pulumi.aws.ec2.VpcEndpointServiceArgs;
import com.pulumi.aws.ec2.VpcEndpoint;
import com.pulumi.aws.ec2.VpcEndpointArgs;
import com.pulumi.aws.ec2.VpcEndpointConnectionAccepter;
import com.pulumi.aws.ec2.VpcEndpointConnectionAccepterArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleVpcEndpointService = new VpcEndpointService("exampleVpcEndpointService", VpcEndpointServiceArgs.builder()        
            .acceptanceRequired(false)
            .networkLoadBalancerArns(aws_lb.example().arn())
            .build());

        var exampleVpcEndpoint = new VpcEndpoint("exampleVpcEndpoint", VpcEndpointArgs.builder()        
            .vpcId(aws_vpc.test_alternate().id())
            .serviceName(aws_vpc_endpoint_service.test().service_name())
            .vpcEndpointType("Interface")
            .privateDnsEnabled(false)
            .securityGroupIds(aws_security_group.test().id())
            .build(), CustomResourceOptions.builder()
                .provider(aws.alternate())
                .build());

        var exampleVpcEndpointConnectionAccepter = new VpcEndpointConnectionAccepter("exampleVpcEndpointConnectionAccepter", VpcEndpointConnectionAccepterArgs.builder()        
            .vpcEndpointServiceId(exampleVpcEndpointService.id())
            .vpcEndpointId(exampleVpcEndpoint.id())
            .build());

    }
}
```
```yaml
resources:
  exampleVpcEndpointService:
    type: aws:ec2:VpcEndpointService
    properties:
      acceptanceRequired: false
      networkLoadBalancerArns:
        - ${aws_lb.example.arn}
  exampleVpcEndpoint:
    type: aws:ec2:VpcEndpoint
    properties:
      vpcId: ${aws_vpc.test_alternate.id}
      serviceName: ${aws_vpc_endpoint_service.test.service_name}
      vpcEndpointType: Interface
      privateDnsEnabled: false
      securityGroupIds:
        - ${aws_security_group.test.id}
    options:
      provider: ${aws.alternate}
  exampleVpcEndpointConnectionAccepter:
    type: aws:ec2:VpcEndpointConnectionAccepter
    properties:
      vpcEndpointServiceId: ${exampleVpcEndpointService.id}
      vpcEndpointId: ${exampleVpcEndpoint.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Endpoint Services using ID of the connection, which is the `VPC Endpoint Service ID` and `VPC Endpoint ID` separated by underscore (`_`).. For example:

```sh
 $ pulumi import aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter foo vpce-svc-0f97a19d3fa8220bc_vpce-010601a6db371e263
```
 