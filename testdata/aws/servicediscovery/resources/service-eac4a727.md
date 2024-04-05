Provides a Service Discovery Service resource.

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
		_, err = servicediscovery.NewService(ctx, "exampleService", &servicediscovery.ServiceArgs{
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
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePublicDnsNamespace = new aws.servicediscovery.PublicDnsNamespace("examplePublicDnsNamespace", {description: "example"});
const exampleService = new aws.servicediscovery.Service("exampleService", {
    dnsConfig: {
        namespaceId: examplePublicDnsNamespace.id,
        dnsRecords: [{
            ttl: 10,
            type: "A",
        }],
    },
    healthCheckConfig: {
        failureThreshold: 10,
        resourcePath: "path",
        type: "HTTP",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_public_dns_namespace = aws.servicediscovery.PublicDnsNamespace("examplePublicDnsNamespace", description="example")
example_service = aws.servicediscovery.Service("exampleService",
    dns_config=aws.servicediscovery.ServiceDnsConfigArgs(
        namespace_id=example_public_dns_namespace.id,
        dns_records=[aws.servicediscovery.ServiceDnsConfigDnsRecordArgs(
            ttl=10,
            type="A",
        )],
    ),
    health_check_config=aws.servicediscovery.ServiceHealthCheckConfigArgs(
        failure_threshold=10,
        resource_path="path",
        type="HTTP",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplePublicDnsNamespace = new Aws.ServiceDiscovery.PublicDnsNamespace("examplePublicDnsNamespace", new()
    {
        Description = "example",
    });

    var exampleService = new Aws.ServiceDiscovery.Service("exampleService", new()
    {
        DnsConfig = new Aws.ServiceDiscovery.Inputs.ServiceDnsConfigArgs
        {
            NamespaceId = examplePublicDnsNamespace.Id,
            DnsRecords = new[]
            {
                new Aws.ServiceDiscovery.Inputs.ServiceDnsConfigDnsRecordArgs
                {
                    Ttl = 10,
                    Type = "A",
                },
            },
        },
        HealthCheckConfig = new Aws.ServiceDiscovery.Inputs.ServiceHealthCheckConfigArgs
        {
            FailureThreshold = 10,
            ResourcePath = "path",
            Type = "HTTP",
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
		examplePublicDnsNamespace, err := servicediscovery.NewPublicDnsNamespace(ctx, "examplePublicDnsNamespace", &servicediscovery.PublicDnsNamespaceArgs{
			Description: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = servicediscovery.NewService(ctx, "exampleService", &servicediscovery.ServiceArgs{
			DnsConfig: &servicediscovery.ServiceDnsConfigArgs{
				NamespaceId: examplePublicDnsNamespace.ID(),
				DnsRecords: servicediscovery.ServiceDnsConfigDnsRecordArray{
					&servicediscovery.ServiceDnsConfigDnsRecordArgs{
						Ttl:  pulumi.Int(10),
						Type: pulumi.String("A"),
					},
				},
			},
			HealthCheckConfig: &servicediscovery.ServiceHealthCheckConfigArgs{
				FailureThreshold: pulumi.Int(10),
				ResourcePath:     pulumi.String("path"),
				Type:             pulumi.String("HTTP"),
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
import com.pulumi.aws.servicediscovery.PublicDnsNamespace;
import com.pulumi.aws.servicediscovery.PublicDnsNamespaceArgs;
import com.pulumi.aws.servicediscovery.Service;
import com.pulumi.aws.servicediscovery.ServiceArgs;
import com.pulumi.aws.servicediscovery.inputs.ServiceDnsConfigArgs;
import com.pulumi.aws.servicediscovery.inputs.ServiceHealthCheckConfigArgs;
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
        var examplePublicDnsNamespace = new PublicDnsNamespace("examplePublicDnsNamespace", PublicDnsNamespaceArgs.builder()        
            .description("example")
            .build());

        var exampleService = new Service("exampleService", ServiceArgs.builder()        
            .dnsConfig(ServiceDnsConfigArgs.builder()
                .namespaceId(examplePublicDnsNamespace.id())
                .dnsRecords(ServiceDnsConfigDnsRecordArgs.builder()
                    .ttl(10)
                    .type("A")
                    .build())
                .build())
            .healthCheckConfig(ServiceHealthCheckConfigArgs.builder()
                .failureThreshold(10)
                .resourcePath("path")
                .type("HTTP")
                .build())
            .build());

    }
}
```
```yaml
resources:
  examplePublicDnsNamespace:
    type: aws:servicediscovery:PublicDnsNamespace
    properties:
      description: example
  exampleService:
    type: aws:servicediscovery:Service
    properties:
      dnsConfig:
        namespaceId: ${examplePublicDnsNamespace.id}
        dnsRecords:
          - ttl: 10
            type: A
      healthCheckConfig:
        failureThreshold: 10
        resourcePath: path
        type: HTTP
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Service Discovery Service using the service ID. For example:

```sh
 $ pulumi import aws:servicediscovery/service:Service example 0123456789
```
 