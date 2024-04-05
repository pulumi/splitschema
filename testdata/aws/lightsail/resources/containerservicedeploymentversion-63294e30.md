Provides a resource to manage a deployment version for your Amazon Lightsail container service.

> **NOTE:** The Amazon Lightsail container service must be enabled to create a deployment.

> **NOTE:** This resource allows you to manage an Amazon Lightsail container service deployment version but the provider cannot destroy it. Removing this resource from your configuration will remove it from your statefile.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lightsail.ContainerServiceDeploymentVersion("example", {
    containers: [{
        containerName: "hello-world",
        image: "amazon/amazon-lightsail:hello-world",
        commands: [],
        environment: {
            MY_ENVIRONMENT_VARIABLE: "my_value",
        },
        ports: {
            "80": "HTTP",
        },
    }],
    publicEndpoint: {
        containerName: "hello-world",
        containerPort: 80,
        healthCheck: {
            healthyThreshold: 2,
            unhealthyThreshold: 2,
            timeoutSeconds: 2,
            intervalSeconds: 5,
            path: "/",
            successCodes: "200-499",
        },
    },
    serviceName: aws_lightsail_container_service.example.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lightsail.ContainerServiceDeploymentVersion("example",
    containers=[aws.lightsail.ContainerServiceDeploymentVersionContainerArgs(
        container_name="hello-world",
        image="amazon/amazon-lightsail:hello-world",
        commands=[],
        environment={
            "MY_ENVIRONMENT_VARIABLE": "my_value",
        },
        ports={
            "80": "HTTP",
        },
    )],
    public_endpoint=aws.lightsail.ContainerServiceDeploymentVersionPublicEndpointArgs(
        container_name="hello-world",
        container_port=80,
        health_check=aws.lightsail.ContainerServiceDeploymentVersionPublicEndpointHealthCheckArgs(
            healthy_threshold=2,
            unhealthy_threshold=2,
            timeout_seconds=2,
            interval_seconds=5,
            path="/",
            success_codes="200-499",
        ),
    ),
    service_name=aws_lightsail_container_service["example"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.LightSail.ContainerServiceDeploymentVersion("example", new()
    {
        Containers = new[]
        {
            new Aws.LightSail.Inputs.ContainerServiceDeploymentVersionContainerArgs
            {
                ContainerName = "hello-world",
                Image = "amazon/amazon-lightsail:hello-world",
                Commands = new() { },
                Environment = 
                {
                    { "MY_ENVIRONMENT_VARIABLE", "my_value" },
                },
                Ports = 
                {
                    { "80", "HTTP" },
                },
            },
        },
        PublicEndpoint = new Aws.LightSail.Inputs.ContainerServiceDeploymentVersionPublicEndpointArgs
        {
            ContainerName = "hello-world",
            ContainerPort = 80,
            HealthCheck = new Aws.LightSail.Inputs.ContainerServiceDeploymentVersionPublicEndpointHealthCheckArgs
            {
                HealthyThreshold = 2,
                UnhealthyThreshold = 2,
                TimeoutSeconds = 2,
                IntervalSeconds = 5,
                Path = "/",
                SuccessCodes = "200-499",
            },
        },
        ServiceName = aws_lightsail_container_service.Example.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewContainerServiceDeploymentVersion(ctx, "example", &lightsail.ContainerServiceDeploymentVersionArgs{
			Containers: lightsail.ContainerServiceDeploymentVersionContainerArray{
				&lightsail.ContainerServiceDeploymentVersionContainerArgs{
					ContainerName: pulumi.String("hello-world"),
					Image:         pulumi.String("amazon/amazon-lightsail:hello-world"),
					Commands:      pulumi.StringArray{},
					Environment: pulumi.StringMap{
						"MY_ENVIRONMENT_VARIABLE": pulumi.String("my_value"),
					},
					Ports: pulumi.StringMap{
						"80": pulumi.String("HTTP"),
					},
				},
			},
			PublicEndpoint: &lightsail.ContainerServiceDeploymentVersionPublicEndpointArgs{
				ContainerName: pulumi.String("hello-world"),
				ContainerPort: pulumi.Int(80),
				HealthCheck: &lightsail.ContainerServiceDeploymentVersionPublicEndpointHealthCheckArgs{
					HealthyThreshold:   pulumi.Int(2),
					UnhealthyThreshold: pulumi.Int(2),
					TimeoutSeconds:     pulumi.Int(2),
					IntervalSeconds:    pulumi.Int(5),
					Path:               pulumi.String("/"),
					SuccessCodes:       pulumi.String("200-499"),
				},
			},
			ServiceName: pulumi.Any(aws_lightsail_container_service.Example.Name),
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
import com.pulumi.aws.lightsail.ContainerServiceDeploymentVersion;
import com.pulumi.aws.lightsail.ContainerServiceDeploymentVersionArgs;
import com.pulumi.aws.lightsail.inputs.ContainerServiceDeploymentVersionContainerArgs;
import com.pulumi.aws.lightsail.inputs.ContainerServiceDeploymentVersionPublicEndpointArgs;
import com.pulumi.aws.lightsail.inputs.ContainerServiceDeploymentVersionPublicEndpointHealthCheckArgs;
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
        var example = new ContainerServiceDeploymentVersion("example", ContainerServiceDeploymentVersionArgs.builder()        
            .containers(ContainerServiceDeploymentVersionContainerArgs.builder()
                .containerName("hello-world")
                .image("amazon/amazon-lightsail:hello-world")
                .commands()
                .environment(Map.of("MY_ENVIRONMENT_VARIABLE", "my_value"))
                .ports(Map.of("80", "HTTP"))
                .build())
            .publicEndpoint(ContainerServiceDeploymentVersionPublicEndpointArgs.builder()
                .containerName("hello-world")
                .containerPort(80)
                .healthCheck(ContainerServiceDeploymentVersionPublicEndpointHealthCheckArgs.builder()
                    .healthyThreshold(2)
                    .unhealthyThreshold(2)
                    .timeoutSeconds(2)
                    .intervalSeconds(5)
                    .path("/")
                    .successCodes("200-499")
                    .build())
                .build())
            .serviceName(aws_lightsail_container_service.example().name())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lightsail:ContainerServiceDeploymentVersion
    properties:
      containers:
        - containerName: hello-world
          image: amazon/amazon-lightsail:hello-world
          commands: []
          environment:
            MY_ENVIRONMENT_VARIABLE: my_value
          ports:
            '80': HTTP
      publicEndpoint:
        containerName: hello-world
        containerPort: 80
        healthCheck:
          healthyThreshold: 2
          unhealthyThreshold: 2
          timeoutSeconds: 2
          intervalSeconds: 5
          path: /
          successCodes: 200-499
      serviceName: ${aws_lightsail_container_service.example.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lightsail Container Service Deployment Version using the `service_name` and `version` separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:lightsail/containerServiceDeploymentVersion:ContainerServiceDeploymentVersion example container-service-1/1
```
 