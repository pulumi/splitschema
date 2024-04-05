Provides a Load Balancer resource.

> **Note:** `aws.alb.LoadBalancer` is known as `aws.lb.LoadBalancer`. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}
### Application Load Balancer

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lb.LoadBalancer("test", {
    internal: false,
    loadBalancerType: "application",
    securityGroups: [aws_security_group.lb_sg.id],
    subnets: .map(subnet => (subnet.id)),
    enableDeletionProtection: true,
    accessLogs: {
        bucket: aws_s3_bucket.lb_logs.id,
        prefix: "test-lb",
        enabled: true,
    },
    tags: {
        Environment: "production",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lb.LoadBalancer("test",
    internal=False,
    load_balancer_type="application",
    security_groups=[aws_security_group["lb_sg"]["id"]],
    subnets=[subnet["id"] for subnet in aws_subnet["public"]],
    enable_deletion_protection=True,
    access_logs=aws.lb.LoadBalancerAccessLogsArgs(
        bucket=aws_s3_bucket["lb_logs"]["id"],
        prefix="test-lb",
        enabled=True,
    ),
    tags={
        "Environment": "production",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LB.LoadBalancer("test", new()
    {
        Internal = false,
        LoadBalancerType = "application",
        SecurityGroups = new[]
        {
            aws_security_group.Lb_sg.Id,
        },
        Subnets = .Select(subnet => 
        {
            return subnet.Id;
        }).ToList(),
        EnableDeletionProtection = true,
        AccessLogs = new Aws.LB.Inputs.LoadBalancerAccessLogsArgs
        {
            Bucket = aws_s3_bucket.Lb_logs.Id,
            Prefix = "test-lb",
            Enabled = true,
        },
        Tags = 
        {
            { "Environment", "production" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewLoadBalancer(ctx, "test", &lb.LoadBalancerArgs{
			Internal:         pulumi.Bool(false),
			LoadBalancerType: pulumi.String("application"),
			SecurityGroups: pulumi.StringArray{
				aws_security_group.Lb_sg.Id,
			},
			Subnets:                  "TODO: For expression",
			EnableDeletionProtection: pulumi.Bool(true),
			AccessLogs: &lb.LoadBalancerAccessLogsArgs{
				Bucket:  pulumi.Any(aws_s3_bucket.Lb_logs.Id),
				Prefix:  pulumi.String("test-lb"),
				Enabled: pulumi.Bool(true),
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
			},
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
### Network Load Balancer

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lb.LoadBalancer("test", {
    internal: false,
    loadBalancerType: "network",
    subnets: .map(subnet => (subnet.id)),
    enableDeletionProtection: true,
    tags: {
        Environment: "production",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lb.LoadBalancer("test",
    internal=False,
    load_balancer_type="network",
    subnets=[subnet["id"] for subnet in aws_subnet["public"]],
    enable_deletion_protection=True,
    tags={
        "Environment": "production",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LB.LoadBalancer("test", new()
    {
        Internal = false,
        LoadBalancerType = "network",
        Subnets = .Select(subnet => 
        {
            return subnet.Id;
        }).ToList(),
        EnableDeletionProtection = true,
        Tags = 
        {
            { "Environment", "production" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewLoadBalancer(ctx, "test", &lb.LoadBalancerArgs{
			Internal:                 pulumi.Bool(false),
			LoadBalancerType:         pulumi.String("network"),
			Subnets:                  "TODO: For expression",
			EnableDeletionProtection: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
			},
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
### Specifying Elastic IPs

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lb.LoadBalancer("example", {
    loadBalancerType: "network",
    subnetMappings: [
        {
            subnetId: aws_subnet.example1.id,
            allocationId: aws_eip.example1.id,
        },
        {
            subnetId: aws_subnet.example2.id,
            allocationId: aws_eip.example2.id,
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lb.LoadBalancer("example",
    load_balancer_type="network",
    subnet_mappings=[
        aws.lb.LoadBalancerSubnetMappingArgs(
            subnet_id=aws_subnet["example1"]["id"],
            allocation_id=aws_eip["example1"]["id"],
        ),
        aws.lb.LoadBalancerSubnetMappingArgs(
            subnet_id=aws_subnet["example2"]["id"],
            allocation_id=aws_eip["example2"]["id"],
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.LB.LoadBalancer("example", new()
    {
        LoadBalancerType = "network",
        SubnetMappings = new[]
        {
            new Aws.LB.Inputs.LoadBalancerSubnetMappingArgs
            {
                SubnetId = aws_subnet.Example1.Id,
                AllocationId = aws_eip.Example1.Id,
            },
            new Aws.LB.Inputs.LoadBalancerSubnetMappingArgs
            {
                SubnetId = aws_subnet.Example2.Id,
                AllocationId = aws_eip.Example2.Id,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewLoadBalancer(ctx, "example", &lb.LoadBalancerArgs{
			LoadBalancerType: pulumi.String("network"),
			SubnetMappings: lb.LoadBalancerSubnetMappingArray{
				&lb.LoadBalancerSubnetMappingArgs{
					SubnetId:     pulumi.Any(aws_subnet.Example1.Id),
					AllocationId: pulumi.Any(aws_eip.Example1.Id),
				},
				&lb.LoadBalancerSubnetMappingArgs{
					SubnetId:     pulumi.Any(aws_subnet.Example2.Id),
					AllocationId: pulumi.Any(aws_eip.Example2.Id),
				},
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
import com.pulumi.aws.lb.LoadBalancer;
import com.pulumi.aws.lb.LoadBalancerArgs;
import com.pulumi.aws.lb.inputs.LoadBalancerSubnetMappingArgs;
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
        var example = new LoadBalancer("example", LoadBalancerArgs.builder()        
            .loadBalancerType("network")
            .subnetMappings(            
                LoadBalancerSubnetMappingArgs.builder()
                    .subnetId(aws_subnet.example1().id())
                    .allocationId(aws_eip.example1().id())
                    .build(),
                LoadBalancerSubnetMappingArgs.builder()
                    .subnetId(aws_subnet.example2().id())
                    .allocationId(aws_eip.example2().id())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lb:LoadBalancer
    properties:
      loadBalancerType: network
      subnetMappings:
        - subnetId: ${aws_subnet.example1.id}
          allocationId: ${aws_eip.example1.id}
        - subnetId: ${aws_subnet.example2.id}
          allocationId: ${aws_eip.example2.id}
```
{{% /example %}}
{{% example %}}
### Specifying private IP addresses for an internal-facing load balancer

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lb.LoadBalancer("example", {
    loadBalancerType: "network",
    subnetMappings: [
        {
            subnetId: aws_subnet.example1.id,
            privateIpv4Address: "10.0.1.15",
        },
        {
            subnetId: aws_subnet.example2.id,
            privateIpv4Address: "10.0.2.15",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lb.LoadBalancer("example",
    load_balancer_type="network",
    subnet_mappings=[
        aws.lb.LoadBalancerSubnetMappingArgs(
            subnet_id=aws_subnet["example1"]["id"],
            private_ipv4_address="10.0.1.15",
        ),
        aws.lb.LoadBalancerSubnetMappingArgs(
            subnet_id=aws_subnet["example2"]["id"],
            private_ipv4_address="10.0.2.15",
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.LB.LoadBalancer("example", new()
    {
        LoadBalancerType = "network",
        SubnetMappings = new[]
        {
            new Aws.LB.Inputs.LoadBalancerSubnetMappingArgs
            {
                SubnetId = aws_subnet.Example1.Id,
                PrivateIpv4Address = "10.0.1.15",
            },
            new Aws.LB.Inputs.LoadBalancerSubnetMappingArgs
            {
                SubnetId = aws_subnet.Example2.Id,
                PrivateIpv4Address = "10.0.2.15",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewLoadBalancer(ctx, "example", &lb.LoadBalancerArgs{
			LoadBalancerType: pulumi.String("network"),
			SubnetMappings: lb.LoadBalancerSubnetMappingArray{
				&lb.LoadBalancerSubnetMappingArgs{
					SubnetId:           pulumi.Any(aws_subnet.Example1.Id),
					PrivateIpv4Address: pulumi.String("10.0.1.15"),
				},
				&lb.LoadBalancerSubnetMappingArgs{
					SubnetId:           pulumi.Any(aws_subnet.Example2.Id),
					PrivateIpv4Address: pulumi.String("10.0.2.15"),
				},
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
import com.pulumi.aws.lb.LoadBalancer;
import com.pulumi.aws.lb.LoadBalancerArgs;
import com.pulumi.aws.lb.inputs.LoadBalancerSubnetMappingArgs;
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
        var example = new LoadBalancer("example", LoadBalancerArgs.builder()        
            .loadBalancerType("network")
            .subnetMappings(            
                LoadBalancerSubnetMappingArgs.builder()
                    .subnetId(aws_subnet.example1().id())
                    .privateIpv4Address("10.0.1.15")
                    .build(),
                LoadBalancerSubnetMappingArgs.builder()
                    .subnetId(aws_subnet.example2().id())
                    .privateIpv4Address("10.0.2.15")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lb:LoadBalancer
    properties:
      loadBalancerType: network
      subnetMappings:
        - subnetId: ${aws_subnet.example1.id}
          privateIpv4Address: 10.0.1.15
        - subnetId: ${aws_subnet.example2.id}
          privateIpv4Address: 10.0.2.15
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import LBs using their ARN. For example:

```sh
 $ pulumi import aws:lb/loadBalancer:LoadBalancer bar arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-load-balancer/50dc6c495c0c9188
```
 