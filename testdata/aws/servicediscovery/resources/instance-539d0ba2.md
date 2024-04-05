Provides a Service Discovery Instance resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpc = new aws.ec2.Vpc("exampleVpc", {
    cidrBlock: "10.0.0.0/16",
    enableDnsSupport: true,
    enableDnsHostnames: true,
});
const examplePrivateDnsNamespace = new aws.servicediscovery.PrivateDnsNamespace("examplePrivateDnsNamespace", {
    description: "example",
    vpc: exampleVpc.id,
});
const exampleService = new aws.servicediscovery.Service("exampleService", {
    dnsConfig: {
        namespaceId: examplePrivateDnsNamespace.id,
        dnsRecords: [{
            ttl: 10,
            type: "A",
        }],
        routingPolicy: "MULTIVALUE",
    },
    healthCheckCustomConfig: {
        failureThreshold: 1,
    },
});
const exampleInstance = new aws.servicediscovery.Instance("exampleInstance", {
    instanceId: "example-instance-id",
    serviceId: exampleService.id,
    attributes: {
        AWS_INSTANCE_IPV4: "172.18.0.1",
        custom_attribute: "custom",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc",
    cidr_block="10.0.0.0/16",
    enable_dns_support=True,
    enable_dns_hostnames=True)
example_private_dns_namespace = aws.servicediscovery.PrivateDnsNamespace("examplePrivateDnsNamespace",
    description="example",
    vpc=example_vpc.id)
example_service = aws.servicediscovery.Service("exampleService",
    dns_config=aws.servicediscovery.ServiceDnsConfigArgs(
        namespace_id=example_private_dns_namespace.id,
        dns_records=[aws.servicediscovery.ServiceDnsConfigDnsRecordArgs(
            ttl=10,
            type="A",
        )],
        routing_policy="MULTIVALUE",
    ),
    health_check_custom_config=aws.servicediscovery.ServiceHealthCheckCustomConfigArgs(
        failure_threshold=1,
    ))
example_instance = aws.servicediscovery.Instance("exampleInstance",
    instance_id="example-instance-id",
    service_id=example_service.id,
    attributes={
        "AWS_INSTANCE_IPV4": "172.18.0.1",
        "custom_attribute": "custom",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.0.0.0/16",
        EnableDnsSupport = true,
        EnableDnsHostnames = true,
    });

    var examplePrivateDnsNamespace = new Aws.ServiceDiscovery.PrivateDnsNamespace("examplePrivateDnsNamespace", new()
    {
        Description = "example",
        Vpc = exampleVpc.Id,
    });

    var exampleService = new Aws.ServiceDiscovery.Service("exampleService", new()
    {
        DnsConfig = new Aws.ServiceDiscovery.Inputs.ServiceDnsConfigArgs
        {
            NamespaceId = examplePrivateDnsNamespace.Id,
            DnsRecords = new[]
            {
                new Aws.ServiceDiscovery.Inputs.ServiceDnsConfigDnsRecordArgs
                {
                    Ttl = 10,
                    Type = "A",
                },
            },
            RoutingPolicy = "MULTIVALUE",
        },
        HealthCheckCustomConfig = new Aws.ServiceDiscovery.Inputs.ServiceHealthCheckCustomConfigArgs
        {
            FailureThreshold = 1,
        },
    });

    var exampleInstance = new Aws.ServiceDiscovery.Instance("exampleInstance", new()
    {
        InstanceId = "example-instance-id",
        ServiceId = exampleService.Id,
        Attributes = 
        {
            { "AWS_INSTANCE_IPV4", "172.18.0.1" },
            { "custom_attribute", "custom" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("10.0.0.0/16"),
			EnableDnsSupport:   pulumi.Bool(true),
			EnableDnsHostnames: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		examplePrivateDnsNamespace, err := servicediscovery.NewPrivateDnsNamespace(ctx, "examplePrivateDnsNamespace", &servicediscovery.PrivateDnsNamespaceArgs{
			Description: pulumi.String("example"),
			Vpc:         exampleVpc.ID(),
		})
		if err != nil {
			return err
		}
		exampleService, err := servicediscovery.NewService(ctx, "exampleService", &servicediscovery.ServiceArgs{
			DnsConfig: &servicediscovery.ServiceDnsConfigArgs{
				NamespaceId: examplePrivateDnsNamespace.ID(),
				DnsRecords: servicediscovery.ServiceDnsConfigDnsRecordArray{
					&servicediscovery.ServiceDnsConfigDnsRecordArgs{
						Ttl:  pulumi.Int(10),
						Type: pulumi.String("A"),
					},
				},
				RoutingPolicy: pulumi.String("MULTIVALUE"),
			},
			HealthCheckCustomConfig: &servicediscovery.ServiceHealthCheckCustomConfigArgs{
				FailureThreshold: pulumi.Int(1),
			},
		})
		if err != nil {
			return err
		}
		_, err = servicediscovery.NewInstance(ctx, "exampleInstance", &servicediscovery.InstanceArgs{
			InstanceId: pulumi.String("example-instance-id"),
			ServiceId:  exampleService.ID(),
			Attributes: pulumi.StringMap{
				"AWS_INSTANCE_IPV4": pulumi.String("172.18.0.1"),
				"custom_attribute":  pulumi.String("custom"),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.servicediscovery.PrivateDnsNamespace;
import com.pulumi.aws.servicediscovery.PrivateDnsNamespaceArgs;
import com.pulumi.aws.servicediscovery.Service;
import com.pulumi.aws.servicediscovery.ServiceArgs;
import com.pulumi.aws.servicediscovery.inputs.ServiceDnsConfigArgs;
import com.pulumi.aws.servicediscovery.inputs.ServiceHealthCheckCustomConfigArgs;
import com.pulumi.aws.servicediscovery.Instance;
import com.pulumi.aws.servicediscovery.InstanceArgs;
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
        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .enableDnsSupport(true)
            .enableDnsHostnames(true)
            .build());

        var examplePrivateDnsNamespace = new PrivateDnsNamespace("examplePrivateDnsNamespace", PrivateDnsNamespaceArgs.builder()        
            .description("example")
            .vpc(exampleVpc.id())
            .build());

        var exampleService = new Service("exampleService", ServiceArgs.builder()        
            .dnsConfig(ServiceDnsConfigArgs.builder()
                .namespaceId(examplePrivateDnsNamespace.id())
                .dnsRecords(ServiceDnsConfigDnsRecordArgs.builder()
                    .ttl(10)
                    .type("A")
                    .build())
                .routingPolicy("MULTIVALUE")
                .build())
            .healthCheckCustomConfig(ServiceHealthCheckCustomConfigArgs.builder()
                .failureThreshold(1)
                .build())
            .build());

        var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()        
            .instanceId("example-instance-id")
            .serviceId(exampleService.id())
            .attributes(Map.ofEntries(
                Map.entry("AWS_INSTANCE_IPV4", "172.18.0.1"),
                Map.entry("custom_attribute", "custom")
            ))
            .build());

    }
}
```
```yaml
resources:
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
      enableDnsSupport: true
      enableDnsHostnames: true
  examplePrivateDnsNamespace:
    type: aws:servicediscovery:PrivateDnsNamespace
    properties:
      description: example
      vpc: ${exampleVpc.id}
  exampleService:
    type: aws:servicediscovery:Service
    properties:
      dnsConfig:
        namespaceId: ${examplePrivateDnsNamespace.id}
        dnsRecords:
          - ttl: 10
            type: A
        routingPolicy: MULTIVALUE
      healthCheckCustomConfig:
        failureThreshold: 1
  exampleInstance:
    type: aws:servicediscovery:Instance
    properties:
      instanceId: example-instance-id
      serviceId: ${exampleService.id}
      attributes:
        AWS_INSTANCE_IPV4: 172.18.0.1
        custom_attribute: custom
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleHttpNamespace = new aws.servicediscovery.HttpNamespace("exampleHttpNamespace", {description: "example"});
const exampleService = new aws.servicediscovery.Service("exampleService", {namespaceId: exampleHttpNamespace.id});
const exampleInstance = new aws.servicediscovery.Instance("exampleInstance", {
    instanceId: "example-instance-id",
    serviceId: exampleService.id,
    attributes: {
        AWS_EC2_INSTANCE_ID: "i-0abdg374kd892cj6dl",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_http_namespace = aws.servicediscovery.HttpNamespace("exampleHttpNamespace", description="example")
example_service = aws.servicediscovery.Service("exampleService", namespace_id=example_http_namespace.id)
example_instance = aws.servicediscovery.Instance("exampleInstance",
    instance_id="example-instance-id",
    service_id=example_service.id,
    attributes={
        "AWS_EC2_INSTANCE_ID": "i-0abdg374kd892cj6dl",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleHttpNamespace = new Aws.ServiceDiscovery.HttpNamespace("exampleHttpNamespace", new()
    {
        Description = "example",
    });

    var exampleService = new Aws.ServiceDiscovery.Service("exampleService", new()
    {
        NamespaceId = exampleHttpNamespace.Id,
    });

    var exampleInstance = new Aws.ServiceDiscovery.Instance("exampleInstance", new()
    {
        InstanceId = "example-instance-id",
        ServiceId = exampleService.Id,
        Attributes = 
        {
            { "AWS_EC2_INSTANCE_ID", "i-0abdg374kd892cj6dl" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleHttpNamespace, err := servicediscovery.NewHttpNamespace(ctx, "exampleHttpNamespace", &servicediscovery.HttpNamespaceArgs{
			Description: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		exampleService, err := servicediscovery.NewService(ctx, "exampleService", &servicediscovery.ServiceArgs{
			NamespaceId: exampleHttpNamespace.ID(),
		})
		if err != nil {
			return err
		}
		_, err = servicediscovery.NewInstance(ctx, "exampleInstance", &servicediscovery.InstanceArgs{
			InstanceId: pulumi.String("example-instance-id"),
			ServiceId:  exampleService.ID(),
			Attributes: pulumi.StringMap{
				"AWS_EC2_INSTANCE_ID": pulumi.String("i-0abdg374kd892cj6dl"),
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
import com.pulumi.aws.servicediscovery.HttpNamespace;
import com.pulumi.aws.servicediscovery.HttpNamespaceArgs;
import com.pulumi.aws.servicediscovery.Service;
import com.pulumi.aws.servicediscovery.ServiceArgs;
import com.pulumi.aws.servicediscovery.Instance;
import com.pulumi.aws.servicediscovery.InstanceArgs;
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
        var exampleHttpNamespace = new HttpNamespace("exampleHttpNamespace", HttpNamespaceArgs.builder()        
            .description("example")
            .build());

        var exampleService = new Service("exampleService", ServiceArgs.builder()        
            .namespaceId(exampleHttpNamespace.id())
            .build());

        var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()        
            .instanceId("example-instance-id")
            .serviceId(exampleService.id())
            .attributes(Map.of("AWS_EC2_INSTANCE_ID", "i-0abdg374kd892cj6dl"))
            .build());

    }
}
```
```yaml
resources:
  exampleHttpNamespace:
    type: aws:servicediscovery:HttpNamespace
    properties:
      description: example
  exampleService:
    type: aws:servicediscovery:Service
    properties:
      namespaceId: ${exampleHttpNamespace.id}
  exampleInstance:
    type: aws:servicediscovery:Instance
    properties:
      instanceId: example-instance-id
      serviceId: ${exampleService.id}
      attributes:
        AWS_EC2_INSTANCE_ID: i-0abdg374kd892cj6dl
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Service Discovery Instance using the service ID and instance ID. For example:

```sh
 $ pulumi import aws:servicediscovery/instance:Instance example 0123456789/i-0123
```
 