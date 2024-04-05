Provides a proxy protocol policy, which allows an ELB to carry a client connection information to a backend.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lb = new aws.elb.LoadBalancer("lb", {
    availabilityZones: ["us-east-1a"],
    listeners: [
        {
            instancePort: 25,
            instanceProtocol: "tcp",
            lbPort: 25,
            lbProtocol: "tcp",
        },
        {
            instancePort: 587,
            instanceProtocol: "tcp",
            lbPort: 587,
            lbProtocol: "tcp",
        },
    ],
});
const smtp = new aws.ec2.ProxyProtocolPolicy("smtp", {
    loadBalancer: lb.name,
    instancePorts: [
        "25",
        "587",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

lb = aws.elb.LoadBalancer("lb",
    availability_zones=["us-east-1a"],
    listeners=[
        aws.elb.LoadBalancerListenerArgs(
            instance_port=25,
            instance_protocol="tcp",
            lb_port=25,
            lb_protocol="tcp",
        ),
        aws.elb.LoadBalancerListenerArgs(
            instance_port=587,
            instance_protocol="tcp",
            lb_port=587,
            lb_protocol="tcp",
        ),
    ])
smtp = aws.ec2.ProxyProtocolPolicy("smtp",
    load_balancer=lb.name,
    instance_ports=[
        "25",
        "587",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var lb = new Aws.Elb.LoadBalancer("lb", new()
    {
        AvailabilityZones = new[]
        {
            "us-east-1a",
        },
        Listeners = new[]
        {
            new Aws.Elb.Inputs.LoadBalancerListenerArgs
            {
                InstancePort = 25,
                InstanceProtocol = "tcp",
                LbPort = 25,
                LbProtocol = "tcp",
            },
            new Aws.Elb.Inputs.LoadBalancerListenerArgs
            {
                InstancePort = 587,
                InstanceProtocol = "tcp",
                LbPort = 587,
                LbProtocol = "tcp",
            },
        },
    });

    var smtp = new Aws.Ec2.ProxyProtocolPolicy("smtp", new()
    {
        LoadBalancer = lb.Name,
        InstancePorts = new[]
        {
            "25",
            "587",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		lb, err := elb.NewLoadBalancer(ctx, "lb", &elb.LoadBalancerArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(25),
					InstanceProtocol: pulumi.String("tcp"),
					LbPort:           pulumi.Int(25),
					LbProtocol:       pulumi.String("tcp"),
				},
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(587),
					InstanceProtocol: pulumi.String("tcp"),
					LbPort:           pulumi.Int(587),
					LbProtocol:       pulumi.String("tcp"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewProxyProtocolPolicy(ctx, "smtp", &ec2.ProxyProtocolPolicyArgs{
			LoadBalancer: lb.Name,
			InstancePorts: pulumi.StringArray{
				pulumi.String("25"),
				pulumi.String("587"),
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
import com.pulumi.aws.elb.LoadBalancer;
import com.pulumi.aws.elb.LoadBalancerArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerListenerArgs;
import com.pulumi.aws.ec2.ProxyProtocolPolicy;
import com.pulumi.aws.ec2.ProxyProtocolPolicyArgs;
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
        var lb = new LoadBalancer("lb", LoadBalancerArgs.builder()        
            .availabilityZones("us-east-1a")
            .listeners(            
                LoadBalancerListenerArgs.builder()
                    .instancePort(25)
                    .instanceProtocol("tcp")
                    .lbPort(25)
                    .lbProtocol("tcp")
                    .build(),
                LoadBalancerListenerArgs.builder()
                    .instancePort(587)
                    .instanceProtocol("tcp")
                    .lbPort(587)
                    .lbProtocol("tcp")
                    .build())
            .build());

        var smtp = new ProxyProtocolPolicy("smtp", ProxyProtocolPolicyArgs.builder()        
            .loadBalancer(lb.name())
            .instancePorts(            
                "25",
                "587")
            .build());

    }
}
```
```yaml
resources:
  lb:
    type: aws:elb:LoadBalancer
    properties:
      availabilityZones:
        - us-east-1a
      listeners:
        - instancePort: 25
          instanceProtocol: tcp
          lbPort: 25
          lbProtocol: tcp
        - instancePort: 587
          instanceProtocol: tcp
          lbPort: 587
          lbProtocol: tcp
  smtp:
    type: aws:ec2:ProxyProtocolPolicy
    properties:
      loadBalancer: ${lb.name}
      instancePorts:
        - '25'
        - '587'
```
{{% /example %}}
{{% /examples %}}