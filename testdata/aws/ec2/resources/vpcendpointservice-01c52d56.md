Provides a VPC Endpoint Service resource.
Service consumers can create an _Interface_ VPC Endpoint to connect to the service.

> **NOTE on VPC Endpoint Services and VPC Endpoint Service Allowed Principals:** This provider provides
both a standalone VPC Endpoint Service Allowed Principal resource
and a VPC Endpoint Service resource with an `allowed_principals` attribute. Do not use the same principal ARN in both
a VPC Endpoint Service resource and a VPC Endpoint Service Allowed Principal resource. Doing so will cause a conflict
and will overwrite the association.

{{% examples %}}
## Example Usage
{{% example %}}
### Network Load Balancers

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.VpcEndpointService("example", {
    acceptanceRequired: false,
    networkLoadBalancerArns: [aws_lb.example.arn],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.VpcEndpointService("example",
    acceptance_required=False,
    network_load_balancer_arns=[aws_lb["example"]["arn"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.VpcEndpointService("example", new()
    {
        AcceptanceRequired = false,
        NetworkLoadBalancerArns = new[]
        {
            aws_lb.Example.Arn,
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
		_, err := ec2.NewVpcEndpointService(ctx, "example", &ec2.VpcEndpointServiceArgs{
			AcceptanceRequired: pulumi.Bool(false),
			NetworkLoadBalancerArns: pulumi.StringArray{
				aws_lb.Example.Arn,
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
import com.pulumi.aws.ec2.VpcEndpointService;
import com.pulumi.aws.ec2.VpcEndpointServiceArgs;
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
        var example = new VpcEndpointService("example", VpcEndpointServiceArgs.builder()        
            .acceptanceRequired(false)
            .networkLoadBalancerArns(aws_lb.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:VpcEndpointService
    properties:
      acceptanceRequired: false
      networkLoadBalancerArns:
        - ${aws_lb.example.arn}
```
{{% /example %}}
{{% example %}}
### Gateway Load Balancers

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.VpcEndpointService("example", {
    acceptanceRequired: false,
    gatewayLoadBalancerArns: [aws_lb.example.arn],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.VpcEndpointService("example",
    acceptance_required=False,
    gateway_load_balancer_arns=[aws_lb["example"]["arn"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.VpcEndpointService("example", new()
    {
        AcceptanceRequired = false,
        GatewayLoadBalancerArns = new[]
        {
            aws_lb.Example.Arn,
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
		_, err := ec2.NewVpcEndpointService(ctx, "example", &ec2.VpcEndpointServiceArgs{
			AcceptanceRequired: pulumi.Bool(false),
			GatewayLoadBalancerArns: pulumi.StringArray{
				aws_lb.Example.Arn,
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
import com.pulumi.aws.ec2.VpcEndpointService;
import com.pulumi.aws.ec2.VpcEndpointServiceArgs;
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
        var example = new VpcEndpointService("example", VpcEndpointServiceArgs.builder()        
            .acceptanceRequired(false)
            .gatewayLoadBalancerArns(aws_lb.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:VpcEndpointService
    properties:
      acceptanceRequired: false
      gatewayLoadBalancerArns:
        - ${aws_lb.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Endpoint Services using the VPC endpoint service `id`. For example:

```sh
 $ pulumi import aws:ec2/vpcEndpointService:VpcEndpointService foo vpce-svc-0f97a19d3fa8220bc
```
 