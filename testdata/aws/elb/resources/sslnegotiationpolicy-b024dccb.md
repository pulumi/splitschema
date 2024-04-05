Provides a load balancer SSL negotiation policy, which allows an ELB to control the ciphers and protocols that are supported during SSL negotiations between a client and a load balancer.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lb = new aws.elb.LoadBalancer("lb", {
    availabilityZones: ["us-east-1a"],
    listeners: [{
        instancePort: 8000,
        instanceProtocol: "https",
        lbPort: 443,
        lbProtocol: "https",
        sslCertificateId: "arn:aws:iam::123456789012:server-certificate/certName",
    }],
});
const foo = new aws.elb.SslNegotiationPolicy("foo", {
    loadBalancer: lb.id,
    lbPort: 443,
    attributes: [
        {
            name: "Protocol-TLSv1",
            value: "false",
        },
        {
            name: "Protocol-TLSv1.1",
            value: "false",
        },
        {
            name: "Protocol-TLSv1.2",
            value: "true",
        },
        {
            name: "Server-Defined-Cipher-Order",
            value: "true",
        },
        {
            name: "ECDHE-RSA-AES128-GCM-SHA256",
            value: "true",
        },
        {
            name: "AES128-GCM-SHA256",
            value: "true",
        },
        {
            name: "EDH-RSA-DES-CBC3-SHA",
            value: "false",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

lb = aws.elb.LoadBalancer("lb",
    availability_zones=["us-east-1a"],
    listeners=[aws.elb.LoadBalancerListenerArgs(
        instance_port=8000,
        instance_protocol="https",
        lb_port=443,
        lb_protocol="https",
        ssl_certificate_id="arn:aws:iam::123456789012:server-certificate/certName",
    )])
foo = aws.elb.SslNegotiationPolicy("foo",
    load_balancer=lb.id,
    lb_port=443,
    attributes=[
        aws.elb.SslNegotiationPolicyAttributeArgs(
            name="Protocol-TLSv1",
            value="false",
        ),
        aws.elb.SslNegotiationPolicyAttributeArgs(
            name="Protocol-TLSv1.1",
            value="false",
        ),
        aws.elb.SslNegotiationPolicyAttributeArgs(
            name="Protocol-TLSv1.2",
            value="true",
        ),
        aws.elb.SslNegotiationPolicyAttributeArgs(
            name="Server-Defined-Cipher-Order",
            value="true",
        ),
        aws.elb.SslNegotiationPolicyAttributeArgs(
            name="ECDHE-RSA-AES128-GCM-SHA256",
            value="true",
        ),
        aws.elb.SslNegotiationPolicyAttributeArgs(
            name="AES128-GCM-SHA256",
            value="true",
        ),
        aws.elb.SslNegotiationPolicyAttributeArgs(
            name="EDH-RSA-DES-CBC3-SHA",
            value="false",
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
                InstancePort = 8000,
                InstanceProtocol = "https",
                LbPort = 443,
                LbProtocol = "https",
                SslCertificateId = "arn:aws:iam::123456789012:server-certificate/certName",
            },
        },
    });

    var foo = new Aws.Elb.SslNegotiationPolicy("foo", new()
    {
        LoadBalancer = lb.Id,
        LbPort = 443,
        Attributes = new[]
        {
            new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
            {
                Name = "Protocol-TLSv1",
                Value = "false",
            },
            new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
            {
                Name = "Protocol-TLSv1.1",
                Value = "false",
            },
            new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
            {
                Name = "Protocol-TLSv1.2",
                Value = "true",
            },
            new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
            {
                Name = "Server-Defined-Cipher-Order",
                Value = "true",
            },
            new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
            {
                Name = "ECDHE-RSA-AES128-GCM-SHA256",
                Value = "true",
            },
            new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
            {
                Name = "AES128-GCM-SHA256",
                Value = "true",
            },
            new Aws.Elb.Inputs.SslNegotiationPolicyAttributeArgs
            {
                Name = "EDH-RSA-DES-CBC3-SHA",
                Value = "false",
            },
        },
    });

});
```
```go
package main

import (
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
					InstancePort:     pulumi.Int(8000),
					InstanceProtocol: pulumi.String("https"),
					LbPort:           pulumi.Int(443),
					LbProtocol:       pulumi.String("https"),
					SslCertificateId: pulumi.String("arn:aws:iam::123456789012:server-certificate/certName"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = elb.NewSslNegotiationPolicy(ctx, "foo", &elb.SslNegotiationPolicyArgs{
			LoadBalancer: lb.ID(),
			LbPort:       pulumi.Int(443),
			Attributes: elb.SslNegotiationPolicyAttributeArray{
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("Protocol-TLSv1"),
					Value: pulumi.String("false"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("Protocol-TLSv1.1"),
					Value: pulumi.String("false"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("Protocol-TLSv1.2"),
					Value: pulumi.String("true"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("Server-Defined-Cipher-Order"),
					Value: pulumi.String("true"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("ECDHE-RSA-AES128-GCM-SHA256"),
					Value: pulumi.String("true"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("AES128-GCM-SHA256"),
					Value: pulumi.String("true"),
				},
				&elb.SslNegotiationPolicyAttributeArgs{
					Name:  pulumi.String("EDH-RSA-DES-CBC3-SHA"),
					Value: pulumi.String("false"),
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
import com.pulumi.aws.elb.LoadBalancer;
import com.pulumi.aws.elb.LoadBalancerArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerListenerArgs;
import com.pulumi.aws.elb.SslNegotiationPolicy;
import com.pulumi.aws.elb.SslNegotiationPolicyArgs;
import com.pulumi.aws.elb.inputs.SslNegotiationPolicyAttributeArgs;
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
            .listeners(LoadBalancerListenerArgs.builder()
                .instancePort(8000)
                .instanceProtocol("https")
                .lbPort(443)
                .lbProtocol("https")
                .sslCertificateId("arn:aws:iam::123456789012:server-certificate/certName")
                .build())
            .build());

        var foo = new SslNegotiationPolicy("foo", SslNegotiationPolicyArgs.builder()        
            .loadBalancer(lb.id())
            .lbPort(443)
            .attributes(            
                SslNegotiationPolicyAttributeArgs.builder()
                    .name("Protocol-TLSv1")
                    .value("false")
                    .build(),
                SslNegotiationPolicyAttributeArgs.builder()
                    .name("Protocol-TLSv1.1")
                    .value("false")
                    .build(),
                SslNegotiationPolicyAttributeArgs.builder()
                    .name("Protocol-TLSv1.2")
                    .value("true")
                    .build(),
                SslNegotiationPolicyAttributeArgs.builder()
                    .name("Server-Defined-Cipher-Order")
                    .value("true")
                    .build(),
                SslNegotiationPolicyAttributeArgs.builder()
                    .name("ECDHE-RSA-AES128-GCM-SHA256")
                    .value("true")
                    .build(),
                SslNegotiationPolicyAttributeArgs.builder()
                    .name("AES128-GCM-SHA256")
                    .value("true")
                    .build(),
                SslNegotiationPolicyAttributeArgs.builder()
                    .name("EDH-RSA-DES-CBC3-SHA")
                    .value("false")
                    .build())
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
        - instancePort: 8000
          instanceProtocol: https
          lbPort: 443
          lbProtocol: https
          sslCertificateId: arn:aws:iam::123456789012:server-certificate/certName
  foo:
    type: aws:elb:SslNegotiationPolicy
    properties:
      loadBalancer: ${lb.id}
      lbPort: 443
      attributes:
        - name: Protocol-TLSv1
          value: 'false'
        - name: Protocol-TLSv1.1
          value: 'false'
        - name: Protocol-TLSv1.2
          value: 'true'
        - name: Server-Defined-Cipher-Order
          value: 'true'
        - name: ECDHE-RSA-AES128-GCM-SHA256
          value: 'true'
        - name: AES128-GCM-SHA256
          value: 'true'
        - name: EDH-RSA-DES-CBC3-SHA
          value: 'false'
```
{{% /example %}}
{{% /examples %}}