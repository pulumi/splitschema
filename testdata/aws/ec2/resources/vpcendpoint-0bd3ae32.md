Provides a VPC Endpoint resource.

> **NOTE on VPC Endpoints and VPC Endpoint Associations:** The provider provides both standalone VPC Endpoint Associations for
Route Tables - (an association between a VPC endpoint and a single `route_table_id`),
Security Groups - (an association between a VPC endpoint and a single `security_group_id`),
and Subnets - (an association between a VPC endpoint and a single `subnet_id`) and
a VPC Endpoint resource with `route_table_ids` and `subnet_ids` attributes.
Do not use the same resource ID in both a VPC Endpoint resource and a VPC Endpoint Association resource.
Doing so will cause a conflict of associations and will overwrite the association.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const s3 = new aws.ec2.VpcEndpoint("s3", {
    vpcId: aws_vpc.main.id,
    serviceName: "com.amazonaws.us-west-2.s3",
});
```
```python
import pulumi
import pulumi_aws as aws

s3 = aws.ec2.VpcEndpoint("s3",
    vpc_id=aws_vpc["main"]["id"],
    service_name="com.amazonaws.us-west-2.s3")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var s3 = new Aws.Ec2.VpcEndpoint("s3", new()
    {
        VpcId = aws_vpc.Main.Id,
        ServiceName = "com.amazonaws.us-west-2.s3",
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
		_, err := ec2.NewVpcEndpoint(ctx, "s3", &ec2.VpcEndpointArgs{
			VpcId:       pulumi.Any(aws_vpc.Main.Id),
			ServiceName: pulumi.String("com.amazonaws.us-west-2.s3"),
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
import com.pulumi.aws.ec2.VpcEndpoint;
import com.pulumi.aws.ec2.VpcEndpointArgs;
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
        var s3 = new VpcEndpoint("s3", VpcEndpointArgs.builder()        
            .vpcId(aws_vpc.main().id())
            .serviceName("com.amazonaws.us-west-2.s3")
            .build());

    }
}
```
```yaml
resources:
  s3:
    type: aws:ec2:VpcEndpoint
    properties:
      vpcId: ${aws_vpc.main.id}
      serviceName: com.amazonaws.us-west-2.s3
```
{{% /example %}}
{{% example %}}
### Basic w/ Tags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const s3 = new aws.ec2.VpcEndpoint("s3", {
    vpcId: aws_vpc.main.id,
    serviceName: "com.amazonaws.us-west-2.s3",
    tags: {
        Environment: "test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

s3 = aws.ec2.VpcEndpoint("s3",
    vpc_id=aws_vpc["main"]["id"],
    service_name="com.amazonaws.us-west-2.s3",
    tags={
        "Environment": "test",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var s3 = new Aws.Ec2.VpcEndpoint("s3", new()
    {
        VpcId = aws_vpc.Main.Id,
        ServiceName = "com.amazonaws.us-west-2.s3",
        Tags = 
        {
            { "Environment", "test" },
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
		_, err := ec2.NewVpcEndpoint(ctx, "s3", &ec2.VpcEndpointArgs{
			VpcId:       pulumi.Any(aws_vpc.Main.Id),
			ServiceName: pulumi.String("com.amazonaws.us-west-2.s3"),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("test"),
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
import com.pulumi.aws.ec2.VpcEndpoint;
import com.pulumi.aws.ec2.VpcEndpointArgs;
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
        var s3 = new VpcEndpoint("s3", VpcEndpointArgs.builder()        
            .vpcId(aws_vpc.main().id())
            .serviceName("com.amazonaws.us-west-2.s3")
            .tags(Map.of("Environment", "test"))
            .build());

    }
}
```
```yaml
resources:
  s3:
    type: aws:ec2:VpcEndpoint
    properties:
      vpcId: ${aws_vpc.main.id}
      serviceName: com.amazonaws.us-west-2.s3
      tags:
        Environment: test
```
{{% /example %}}
{{% example %}}
### Interface Endpoint Type

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ec2 = new aws.ec2.VpcEndpoint("ec2", {
    vpcId: aws_vpc.main.id,
    serviceName: "com.amazonaws.us-west-2.ec2",
    vpcEndpointType: "Interface",
    securityGroupIds: [aws_security_group.sg1.id],
    privateDnsEnabled: true,
});
```
```python
import pulumi
import pulumi_aws as aws

ec2 = aws.ec2.VpcEndpoint("ec2",
    vpc_id=aws_vpc["main"]["id"],
    service_name="com.amazonaws.us-west-2.ec2",
    vpc_endpoint_type="Interface",
    security_group_ids=[aws_security_group["sg1"]["id"]],
    private_dns_enabled=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ec2 = new Aws.Ec2.VpcEndpoint("ec2", new()
    {
        VpcId = aws_vpc.Main.Id,
        ServiceName = "com.amazonaws.us-west-2.ec2",
        VpcEndpointType = "Interface",
        SecurityGroupIds = new[]
        {
            aws_security_group.Sg1.Id,
        },
        PrivateDnsEnabled = true,
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
		_, err := ec2.NewVpcEndpoint(ctx, "ec2", &ec2.VpcEndpointArgs{
			VpcId:           pulumi.Any(aws_vpc.Main.Id),
			ServiceName:     pulumi.String("com.amazonaws.us-west-2.ec2"),
			VpcEndpointType: pulumi.String("Interface"),
			SecurityGroupIds: pulumi.StringArray{
				aws_security_group.Sg1.Id,
			},
			PrivateDnsEnabled: pulumi.Bool(true),
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
import com.pulumi.aws.ec2.VpcEndpoint;
import com.pulumi.aws.ec2.VpcEndpointArgs;
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
        var ec2 = new VpcEndpoint("ec2", VpcEndpointArgs.builder()        
            .vpcId(aws_vpc.main().id())
            .serviceName("com.amazonaws.us-west-2.ec2")
            .vpcEndpointType("Interface")
            .securityGroupIds(aws_security_group.sg1().id())
            .privateDnsEnabled(true)
            .build());

    }
}
```
```yaml
resources:
  ec2:
    type: aws:ec2:VpcEndpoint
    properties:
      vpcId: ${aws_vpc.main.id}
      serviceName: com.amazonaws.us-west-2.ec2
      vpcEndpointType: Interface
      securityGroupIds:
        - ${aws_security_group.sg1.id}
      privateDnsEnabled: true
```
{{% /example %}}
{{% example %}}
### Gateway Load Balancer Endpoint Type

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const exampleVpcEndpointService = new aws.ec2.VpcEndpointService("exampleVpcEndpointService", {
    acceptanceRequired: false,
    allowedPrincipals: [current.then(current => current.arn)],
    gatewayLoadBalancerArns: [aws_lb.example.arn],
});
const exampleVpcEndpoint = new aws.ec2.VpcEndpoint("exampleVpcEndpoint", {
    serviceName: exampleVpcEndpointService.serviceName,
    subnetIds: [aws_subnet.example.id],
    vpcEndpointType: exampleVpcEndpointService.serviceType,
    vpcId: aws_vpc.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
example_vpc_endpoint_service = aws.ec2.VpcEndpointService("exampleVpcEndpointService",
    acceptance_required=False,
    allowed_principals=[current.arn],
    gateway_load_balancer_arns=[aws_lb["example"]["arn"]])
example_vpc_endpoint = aws.ec2.VpcEndpoint("exampleVpcEndpoint",
    service_name=example_vpc_endpoint_service.service_name,
    subnet_ids=[aws_subnet["example"]["id"]],
    vpc_endpoint_type=example_vpc_endpoint_service.service_type,
    vpc_id=aws_vpc["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var exampleVpcEndpointService = new Aws.Ec2.VpcEndpointService("exampleVpcEndpointService", new()
    {
        AcceptanceRequired = false,
        AllowedPrincipals = new[]
        {
            current.Apply(getCallerIdentityResult => getCallerIdentityResult.Arn),
        },
        GatewayLoadBalancerArns = new[]
        {
            aws_lb.Example.Arn,
        },
    });

    var exampleVpcEndpoint = new Aws.Ec2.VpcEndpoint("exampleVpcEndpoint", new()
    {
        ServiceName = exampleVpcEndpointService.ServiceName,
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        VpcEndpointType = exampleVpcEndpointService.ServiceType,
        VpcId = aws_vpc.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleVpcEndpointService, err := ec2.NewVpcEndpointService(ctx, "exampleVpcEndpointService", &ec2.VpcEndpointServiceArgs{
			AcceptanceRequired: pulumi.Bool(false),
			AllowedPrincipals: pulumi.StringArray{
				*pulumi.String(current.Arn),
			},
			GatewayLoadBalancerArns: pulumi.StringArray{
				aws_lb.Example.Arn,
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcEndpoint(ctx, "exampleVpcEndpoint", &ec2.VpcEndpointArgs{
			ServiceName: exampleVpcEndpointService.ServiceName,
			SubnetIds: pulumi.StringArray{
				aws_subnet.Example.Id,
			},
			VpcEndpointType: exampleVpcEndpointService.ServiceType,
			VpcId:           pulumi.Any(aws_vpc.Example.Id),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.ec2.VpcEndpointService;
import com.pulumi.aws.ec2.VpcEndpointServiceArgs;
import com.pulumi.aws.ec2.VpcEndpoint;
import com.pulumi.aws.ec2.VpcEndpointArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        var exampleVpcEndpointService = new VpcEndpointService("exampleVpcEndpointService", VpcEndpointServiceArgs.builder()        
            .acceptanceRequired(false)
            .allowedPrincipals(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.arn()))
            .gatewayLoadBalancerArns(aws_lb.example().arn())
            .build());

        var exampleVpcEndpoint = new VpcEndpoint("exampleVpcEndpoint", VpcEndpointArgs.builder()        
            .serviceName(exampleVpcEndpointService.serviceName())
            .subnetIds(aws_subnet.example().id())
            .vpcEndpointType(exampleVpcEndpointService.serviceType())
            .vpcId(aws_vpc.example().id())
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
      allowedPrincipals:
        - ${current.arn}
      gatewayLoadBalancerArns:
        - ${aws_lb.example.arn}
  exampleVpcEndpoint:
    type: aws:ec2:VpcEndpoint
    properties:
      serviceName: ${exampleVpcEndpointService.serviceName}
      subnetIds:
        - ${aws_subnet.example.id}
      vpcEndpointType: ${exampleVpcEndpointService.serviceType}
      vpcId: ${aws_vpc.example.id}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Endpoints using the VPC endpoint `id`. For example:

```sh
 $ pulumi import aws:ec2/vpcEndpoint:VpcEndpoint endpoint1 vpce-3ecf2a57
```
 