Provides an Elastic Load Balancer resource, also known as a "Classic
Load Balancer" after the release of
Application/Network Load Balancers.

> **NOTE on ELB Instances and ELB Attachments:** This provider currently
provides both a standalone ELB Attachment resource
(describing an instance attached to an ELB), and an ELB resource with
`instances` defined in-line. At this time you cannot use an ELB with in-line
instances in conjunction with a ELB Attachment resources. Doing so will cause a
conflict and will overwrite attachments.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new load balancer
const bar = new aws.elb.LoadBalancer("bar", {
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    accessLogs: {
        bucket: "foo",
        bucketPrefix: "bar",
        interval: 60,
    },
    listeners: [
        {
            instancePort: 8000,
            instanceProtocol: "http",
            lbPort: 80,
            lbProtocol: "http",
        },
        {
            instancePort: 8000,
            instanceProtocol: "http",
            lbPort: 443,
            lbProtocol: "https",
            sslCertificateId: "arn:aws:iam::123456789012:server-certificate/certName",
        },
    ],
    healthCheck: {
        healthyThreshold: 2,
        unhealthyThreshold: 2,
        timeout: 3,
        target: "HTTP:8000/",
        interval: 30,
    },
    instances: [aws_instance.foo.id],
    crossZoneLoadBalancing: true,
    idleTimeout: 400,
    connectionDraining: true,
    connectionDrainingTimeout: 400,
    tags: {
        Name: "foobar-elb",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new load balancer
bar = aws.elb.LoadBalancer("bar",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    access_logs=aws.elb.LoadBalancerAccessLogsArgs(
        bucket="foo",
        bucket_prefix="bar",
        interval=60,
    ),
    listeners=[
        aws.elb.LoadBalancerListenerArgs(
            instance_port=8000,
            instance_protocol="http",
            lb_port=80,
            lb_protocol="http",
        ),
        aws.elb.LoadBalancerListenerArgs(
            instance_port=8000,
            instance_protocol="http",
            lb_port=443,
            lb_protocol="https",
            ssl_certificate_id="arn:aws:iam::123456789012:server-certificate/certName",
        ),
    ],
    health_check=aws.elb.LoadBalancerHealthCheckArgs(
        healthy_threshold=2,
        unhealthy_threshold=2,
        timeout=3,
        target="HTTP:8000/",
        interval=30,
    ),
    instances=[aws_instance["foo"]["id"]],
    cross_zone_load_balancing=True,
    idle_timeout=400,
    connection_draining=True,
    connection_draining_timeout=400,
    tags={
        "Name": "foobar-elb",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new load balancer
    var bar = new Aws.Elb.LoadBalancer("bar", new()
    {
        AvailabilityZones = new[]
        {
            "us-west-2a",
            "us-west-2b",
            "us-west-2c",
        },
        AccessLogs = new Aws.Elb.Inputs.LoadBalancerAccessLogsArgs
        {
            Bucket = "foo",
            BucketPrefix = "bar",
            Interval = 60,
        },
        Listeners = new[]
        {
            new Aws.Elb.Inputs.LoadBalancerListenerArgs
            {
                InstancePort = 8000,
                InstanceProtocol = "http",
                LbPort = 80,
                LbProtocol = "http",
            },
            new Aws.Elb.Inputs.LoadBalancerListenerArgs
            {
                InstancePort = 8000,
                InstanceProtocol = "http",
                LbPort = 443,
                LbProtocol = "https",
                SslCertificateId = "arn:aws:iam::123456789012:server-certificate/certName",
            },
        },
        HealthCheck = new Aws.Elb.Inputs.LoadBalancerHealthCheckArgs
        {
            HealthyThreshold = 2,
            UnhealthyThreshold = 2,
            Timeout = 3,
            Target = "HTTP:8000/",
            Interval = 30,
        },
        Instances = new[]
        {
            aws_instance.Foo.Id,
        },
        CrossZoneLoadBalancing = true,
        IdleTimeout = 400,
        ConnectionDraining = true,
        ConnectionDrainingTimeout = 400,
        Tags = 
        {
            { "Name", "foobar-elb" },
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
		_, err := elb.NewLoadBalancer(ctx, "bar", &elb.LoadBalancerArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
				pulumi.String("us-west-2c"),
			},
			AccessLogs: &elb.LoadBalancerAccessLogsArgs{
				Bucket:       pulumi.String("foo"),
				BucketPrefix: pulumi.String("bar"),
				Interval:     pulumi.Int(60),
			},
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(8000),
					InstanceProtocol: pulumi.String("http"),
					LbPort:           pulumi.Int(80),
					LbProtocol:       pulumi.String("http"),
				},
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(8000),
					InstanceProtocol: pulumi.String("http"),
					LbPort:           pulumi.Int(443),
					LbProtocol:       pulumi.String("https"),
					SslCertificateId: pulumi.String("arn:aws:iam::123456789012:server-certificate/certName"),
				},
			},
			HealthCheck: &elb.LoadBalancerHealthCheckArgs{
				HealthyThreshold:   pulumi.Int(2),
				UnhealthyThreshold: pulumi.Int(2),
				Timeout:            pulumi.Int(3),
				Target:             pulumi.String("HTTP:8000/"),
				Interval:           pulumi.Int(30),
			},
			Instances: pulumi.StringArray{
				aws_instance.Foo.Id,
			},
			CrossZoneLoadBalancing:    pulumi.Bool(true),
			IdleTimeout:               pulumi.Int(400),
			ConnectionDraining:        pulumi.Bool(true),
			ConnectionDrainingTimeout: pulumi.Int(400),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("foobar-elb"),
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
import com.pulumi.aws.elb.inputs.LoadBalancerAccessLogsArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerListenerArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerHealthCheckArgs;
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
        var bar = new LoadBalancer("bar", LoadBalancerArgs.builder()        
            .availabilityZones(            
                "us-west-2a",
                "us-west-2b",
                "us-west-2c")
            .accessLogs(LoadBalancerAccessLogsArgs.builder()
                .bucket("foo")
                .bucketPrefix("bar")
                .interval(60)
                .build())
            .listeners(            
                LoadBalancerListenerArgs.builder()
                    .instancePort(8000)
                    .instanceProtocol("http")
                    .lbPort(80)
                    .lbProtocol("http")
                    .build(),
                LoadBalancerListenerArgs.builder()
                    .instancePort(8000)
                    .instanceProtocol("http")
                    .lbPort(443)
                    .lbProtocol("https")
                    .sslCertificateId("arn:aws:iam::123456789012:server-certificate/certName")
                    .build())
            .healthCheck(LoadBalancerHealthCheckArgs.builder()
                .healthyThreshold(2)
                .unhealthyThreshold(2)
                .timeout(3)
                .target("HTTP:8000/")
                .interval(30)
                .build())
            .instances(aws_instance.foo().id())
            .crossZoneLoadBalancing(true)
            .idleTimeout(400)
            .connectionDraining(true)
            .connectionDrainingTimeout(400)
            .tags(Map.of("Name", "foobar-elb"))
            .build());

    }
}
```
```yaml
resources:
  # Create a new load balancer
  bar:
    type: aws:elb:LoadBalancer
    properties:
      availabilityZones:
        - us-west-2a
        - us-west-2b
        - us-west-2c
      accessLogs:
        bucket: foo
        bucketPrefix: bar
        interval: 60
      listeners:
        - instancePort: 8000
          instanceProtocol: http
          lbPort: 80
          lbProtocol: http
        - instancePort: 8000
          instanceProtocol: http
          lbPort: 443
          lbProtocol: https
          sslCertificateId: arn:aws:iam::123456789012:server-certificate/certName
      healthCheck:
        healthyThreshold: 2
        unhealthyThreshold: 2
        timeout: 3
        target: HTTP:8000/
        interval: 30
      instances:
        - ${aws_instance.foo.id}
      crossZoneLoadBalancing: true
      idleTimeout: 400
      connectionDraining: true
      connectionDrainingTimeout: 400
      tags:
        Name: foobar-elb
```
{{% /example %}}
{{% /examples %}}
## Note on ECDSA Key Algorithm

If the ARN of the `ssl_certificate_id` that is pointed to references a
certificate that was signed by an ECDSA key, note that ELB only supports the
P256 and P384 curves.  Using a certificate signed by a key using a different
curve could produce the error `ERR_SSL_VERSION_OR_CIPHER_MISMATCH` in your
browser.


## Import

Using `pulumi import`, import ELBs using the `name`. For example:

```sh
 $ pulumi import aws:elb/loadBalancer:LoadBalancer bar elb-production-12345
```
 