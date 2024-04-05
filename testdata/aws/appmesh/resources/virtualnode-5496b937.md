Provides an AWS App Mesh virtual node resource.

## Breaking Changes

Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92)), `aws.appmesh.VirtualNode` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:

* Rename the `service_name` attribute of the `dns` object to `hostname`.

* Replace the `backends` attribute of the `spec` object with one or more `backend` configuration blocks,
setting `virtual_service_name` to the name of the service.

The state associated with existing resources will automatically be migrated.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb1 = new aws.appmesh.VirtualNode("serviceb1", {
    meshName: aws_appmesh_mesh.simple.id,
    spec: {
        backends: [{
            virtualService: {
                virtualServiceName: "servicea.simpleapp.local",
            },
        }],
        listeners: [{
            portMapping: {
                port: 8080,
                protocol: "http",
            },
        }],
        serviceDiscovery: {
            dns: {
                hostname: "serviceb.simpleapp.local",
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

serviceb1 = aws.appmesh.VirtualNode("serviceb1",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec=aws.appmesh.VirtualNodeSpecArgs(
        backends=[aws.appmesh.VirtualNodeSpecBackendArgs(
            virtual_service=aws.appmesh.VirtualNodeSpecBackendVirtualServiceArgs(
                virtual_service_name="servicea.simpleapp.local",
            ),
        )],
        listeners=[aws.appmesh.VirtualNodeSpecListenerArgs(
            port_mapping=aws.appmesh.VirtualNodeSpecListenerPortMappingArgs(
                port=8080,
                protocol="http",
            ),
        )],
        service_discovery=aws.appmesh.VirtualNodeSpecServiceDiscoveryArgs(
            dns=aws.appmesh.VirtualNodeSpecServiceDiscoveryDnsArgs(
                hostname="serviceb.simpleapp.local",
            ),
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var serviceb1 = new Aws.AppMesh.VirtualNode("serviceb1", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        Spec = new Aws.AppMesh.Inputs.VirtualNodeSpecArgs
        {
            Backends = new[]
            {
                new Aws.AppMesh.Inputs.VirtualNodeSpecBackendArgs
                {
                    VirtualService = new Aws.AppMesh.Inputs.VirtualNodeSpecBackendVirtualServiceArgs
                    {
                        VirtualServiceName = "servicea.simpleapp.local",
                    },
                },
            },
            Listeners = new[]
            {
                new Aws.AppMesh.Inputs.VirtualNodeSpecListenerArgs
                {
                    PortMapping = new Aws.AppMesh.Inputs.VirtualNodeSpecListenerPortMappingArgs
                    {
                        Port = 8080,
                        Protocol = "http",
                    },
                },
            },
            ServiceDiscovery = new Aws.AppMesh.Inputs.VirtualNodeSpecServiceDiscoveryArgs
            {
                Dns = new Aws.AppMesh.Inputs.VirtualNodeSpecServiceDiscoveryDnsArgs
                {
                    Hostname = "serviceb.simpleapp.local",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appmesh"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appmesh.NewVirtualNode(ctx, "serviceb1", &appmesh.VirtualNodeArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.VirtualNodeSpecArgs{
				Backends: appmesh.VirtualNodeSpecBackendArray{
					&appmesh.VirtualNodeSpecBackendArgs{
						VirtualService: &appmesh.VirtualNodeSpecBackendVirtualServiceArgs{
							VirtualServiceName: pulumi.String("servicea.simpleapp.local"),
						},
					},
				},
				Listeners: appmesh.VirtualNodeSpecListenerArray{
					&appmesh.VirtualNodeSpecListenerArgs{
						PortMapping: &appmesh.VirtualNodeSpecListenerPortMappingArgs{
							Port:     pulumi.Int(8080),
							Protocol: pulumi.String("http"),
						},
					},
				},
				ServiceDiscovery: &appmesh.VirtualNodeSpecServiceDiscoveryArgs{
					Dns: &appmesh.VirtualNodeSpecServiceDiscoveryDnsArgs{
						Hostname: pulumi.String("serviceb.simpleapp.local"),
					},
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
import com.pulumi.aws.appmesh.VirtualNode;
import com.pulumi.aws.appmesh.VirtualNodeArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecServiceDiscoveryArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecServiceDiscoveryDnsArgs;
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
        var serviceb1 = new VirtualNode("serviceb1", VirtualNodeArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .spec(VirtualNodeSpecArgs.builder()
                .backends(VirtualNodeSpecBackendArgs.builder()
                    .virtualService(VirtualNodeSpecBackendVirtualServiceArgs.builder()
                        .virtualServiceName("servicea.simpleapp.local")
                        .build())
                    .build())
                .listeners(VirtualNodeSpecListenerArgs.builder()
                    .portMapping(VirtualNodeSpecListenerPortMappingArgs.builder()
                        .port(8080)
                        .protocol("http")
                        .build())
                    .build())
                .serviceDiscovery(VirtualNodeSpecServiceDiscoveryArgs.builder()
                    .dns(VirtualNodeSpecServiceDiscoveryDnsArgs.builder()
                        .hostname("serviceb.simpleapp.local")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  serviceb1:
    type: aws:appmesh:VirtualNode
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      spec:
        backends:
          - virtualService:
              virtualServiceName: servicea.simpleapp.local
        listeners:
          - portMapping:
              port: 8080
              protocol: http
        serviceDiscovery:
          dns:
            hostname: serviceb.simpleapp.local
```
{{% /example %}}
{{% example %}}
### AWS Cloud Map Service Discovery

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicediscovery.HttpNamespace("example", {});
const serviceb1 = new aws.appmesh.VirtualNode("serviceb1", {
    meshName: aws_appmesh_mesh.simple.id,
    spec: {
        backends: [{
            virtualService: {
                virtualServiceName: "servicea.simpleapp.local",
            },
        }],
        listeners: [{
            portMapping: {
                port: 8080,
                protocol: "http",
            },
        }],
        serviceDiscovery: {
            awsCloudMap: {
                attributes: {
                    stack: "blue",
                },
                serviceName: "serviceb1",
                namespaceName: example.name,
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicediscovery.HttpNamespace("example")
serviceb1 = aws.appmesh.VirtualNode("serviceb1",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec=aws.appmesh.VirtualNodeSpecArgs(
        backends=[aws.appmesh.VirtualNodeSpecBackendArgs(
            virtual_service=aws.appmesh.VirtualNodeSpecBackendVirtualServiceArgs(
                virtual_service_name="servicea.simpleapp.local",
            ),
        )],
        listeners=[aws.appmesh.VirtualNodeSpecListenerArgs(
            port_mapping=aws.appmesh.VirtualNodeSpecListenerPortMappingArgs(
                port=8080,
                protocol="http",
            ),
        )],
        service_discovery=aws.appmesh.VirtualNodeSpecServiceDiscoveryArgs(
            aws_cloud_map=aws.appmesh.VirtualNodeSpecServiceDiscoveryAwsCloudMapArgs(
                attributes={
                    "stack": "blue",
                },
                service_name="serviceb1",
                namespace_name=example.name,
            ),
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceDiscovery.HttpNamespace("example");

    var serviceb1 = new Aws.AppMesh.VirtualNode("serviceb1", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        Spec = new Aws.AppMesh.Inputs.VirtualNodeSpecArgs
        {
            Backends = new[]
            {
                new Aws.AppMesh.Inputs.VirtualNodeSpecBackendArgs
                {
                    VirtualService = new Aws.AppMesh.Inputs.VirtualNodeSpecBackendVirtualServiceArgs
                    {
                        VirtualServiceName = "servicea.simpleapp.local",
                    },
                },
            },
            Listeners = new[]
            {
                new Aws.AppMesh.Inputs.VirtualNodeSpecListenerArgs
                {
                    PortMapping = new Aws.AppMesh.Inputs.VirtualNodeSpecListenerPortMappingArgs
                    {
                        Port = 8080,
                        Protocol = "http",
                    },
                },
            },
            ServiceDiscovery = new Aws.AppMesh.Inputs.VirtualNodeSpecServiceDiscoveryArgs
            {
                AwsCloudMap = new Aws.AppMesh.Inputs.VirtualNodeSpecServiceDiscoveryAwsCloudMapArgs
                {
                    Attributes = 
                    {
                        { "stack", "blue" },
                    },
                    ServiceName = "serviceb1",
                    NamespaceName = example.Name,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appmesh"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := servicediscovery.NewHttpNamespace(ctx, "example", nil)
		if err != nil {
			return err
		}
		_, err = appmesh.NewVirtualNode(ctx, "serviceb1", &appmesh.VirtualNodeArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.VirtualNodeSpecArgs{
				Backends: appmesh.VirtualNodeSpecBackendArray{
					&appmesh.VirtualNodeSpecBackendArgs{
						VirtualService: &appmesh.VirtualNodeSpecBackendVirtualServiceArgs{
							VirtualServiceName: pulumi.String("servicea.simpleapp.local"),
						},
					},
				},
				Listeners: appmesh.VirtualNodeSpecListenerArray{
					&appmesh.VirtualNodeSpecListenerArgs{
						PortMapping: &appmesh.VirtualNodeSpecListenerPortMappingArgs{
							Port:     pulumi.Int(8080),
							Protocol: pulumi.String("http"),
						},
					},
				},
				ServiceDiscovery: &appmesh.VirtualNodeSpecServiceDiscoveryArgs{
					AwsCloudMap: &appmesh.VirtualNodeSpecServiceDiscoveryAwsCloudMapArgs{
						Attributes: pulumi.StringMap{
							"stack": pulumi.String("blue"),
						},
						ServiceName:   pulumi.String("serviceb1"),
						NamespaceName: example.Name,
					},
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
import com.pulumi.aws.servicediscovery.HttpNamespace;
import com.pulumi.aws.appmesh.VirtualNode;
import com.pulumi.aws.appmesh.VirtualNodeArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecServiceDiscoveryArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecServiceDiscoveryAwsCloudMapArgs;
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
        var example = new HttpNamespace("example");

        var serviceb1 = new VirtualNode("serviceb1", VirtualNodeArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .spec(VirtualNodeSpecArgs.builder()
                .backends(VirtualNodeSpecBackendArgs.builder()
                    .virtualService(VirtualNodeSpecBackendVirtualServiceArgs.builder()
                        .virtualServiceName("servicea.simpleapp.local")
                        .build())
                    .build())
                .listeners(VirtualNodeSpecListenerArgs.builder()
                    .portMapping(VirtualNodeSpecListenerPortMappingArgs.builder()
                        .port(8080)
                        .protocol("http")
                        .build())
                    .build())
                .serviceDiscovery(VirtualNodeSpecServiceDiscoveryArgs.builder()
                    .awsCloudMap(VirtualNodeSpecServiceDiscoveryAwsCloudMapArgs.builder()
                        .attributes(Map.of("stack", "blue"))
                        .serviceName("serviceb1")
                        .namespaceName(example.name())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicediscovery:HttpNamespace
  serviceb1:
    type: aws:appmesh:VirtualNode
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      spec:
        backends:
          - virtualService:
              virtualServiceName: servicea.simpleapp.local
        listeners:
          - portMapping:
              port: 8080
              protocol: http
        serviceDiscovery:
          awsCloudMap:
            attributes:
              stack: blue
            serviceName: serviceb1
            namespaceName: ${example.name}
```
{{% /example %}}
{{% example %}}
### Listener Health Check

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb1 = new aws.appmesh.VirtualNode("serviceb1", {
    meshName: aws_appmesh_mesh.simple.id,
    spec: {
        backends: [{
            virtualService: {
                virtualServiceName: "servicea.simpleapp.local",
            },
        }],
        listeners: [{
            portMapping: {
                port: 8080,
                protocol: "http",
            },
            healthCheck: {
                protocol: "http",
                path: "/ping",
                healthyThreshold: 2,
                unhealthyThreshold: 2,
                timeoutMillis: 2000,
                intervalMillis: 5000,
            },
        }],
        serviceDiscovery: {
            dns: {
                hostname: "serviceb.simpleapp.local",
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

serviceb1 = aws.appmesh.VirtualNode("serviceb1",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec=aws.appmesh.VirtualNodeSpecArgs(
        backends=[aws.appmesh.VirtualNodeSpecBackendArgs(
            virtual_service=aws.appmesh.VirtualNodeSpecBackendVirtualServiceArgs(
                virtual_service_name="servicea.simpleapp.local",
            ),
        )],
        listeners=[aws.appmesh.VirtualNodeSpecListenerArgs(
            port_mapping=aws.appmesh.VirtualNodeSpecListenerPortMappingArgs(
                port=8080,
                protocol="http",
            ),
            health_check=aws.appmesh.VirtualNodeSpecListenerHealthCheckArgs(
                protocol="http",
                path="/ping",
                healthy_threshold=2,
                unhealthy_threshold=2,
                timeout_millis=2000,
                interval_millis=5000,
            ),
        )],
        service_discovery=aws.appmesh.VirtualNodeSpecServiceDiscoveryArgs(
            dns=aws.appmesh.VirtualNodeSpecServiceDiscoveryDnsArgs(
                hostname="serviceb.simpleapp.local",
            ),
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var serviceb1 = new Aws.AppMesh.VirtualNode("serviceb1", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        Spec = new Aws.AppMesh.Inputs.VirtualNodeSpecArgs
        {
            Backends = new[]
            {
                new Aws.AppMesh.Inputs.VirtualNodeSpecBackendArgs
                {
                    VirtualService = new Aws.AppMesh.Inputs.VirtualNodeSpecBackendVirtualServiceArgs
                    {
                        VirtualServiceName = "servicea.simpleapp.local",
                    },
                },
            },
            Listeners = new[]
            {
                new Aws.AppMesh.Inputs.VirtualNodeSpecListenerArgs
                {
                    PortMapping = new Aws.AppMesh.Inputs.VirtualNodeSpecListenerPortMappingArgs
                    {
                        Port = 8080,
                        Protocol = "http",
                    },
                    HealthCheck = new Aws.AppMesh.Inputs.VirtualNodeSpecListenerHealthCheckArgs
                    {
                        Protocol = "http",
                        Path = "/ping",
                        HealthyThreshold = 2,
                        UnhealthyThreshold = 2,
                        TimeoutMillis = 2000,
                        IntervalMillis = 5000,
                    },
                },
            },
            ServiceDiscovery = new Aws.AppMesh.Inputs.VirtualNodeSpecServiceDiscoveryArgs
            {
                Dns = new Aws.AppMesh.Inputs.VirtualNodeSpecServiceDiscoveryDnsArgs
                {
                    Hostname = "serviceb.simpleapp.local",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appmesh"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appmesh.NewVirtualNode(ctx, "serviceb1", &appmesh.VirtualNodeArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.VirtualNodeSpecArgs{
				Backends: appmesh.VirtualNodeSpecBackendArray{
					&appmesh.VirtualNodeSpecBackendArgs{
						VirtualService: &appmesh.VirtualNodeSpecBackendVirtualServiceArgs{
							VirtualServiceName: pulumi.String("servicea.simpleapp.local"),
						},
					},
				},
				Listeners: appmesh.VirtualNodeSpecListenerArray{
					&appmesh.VirtualNodeSpecListenerArgs{
						PortMapping: &appmesh.VirtualNodeSpecListenerPortMappingArgs{
							Port:     pulumi.Int(8080),
							Protocol: pulumi.String("http"),
						},
						HealthCheck: &appmesh.VirtualNodeSpecListenerHealthCheckArgs{
							Protocol:           pulumi.String("http"),
							Path:               pulumi.String("/ping"),
							HealthyThreshold:   pulumi.Int(2),
							UnhealthyThreshold: pulumi.Int(2),
							TimeoutMillis:      pulumi.Int(2000),
							IntervalMillis:     pulumi.Int(5000),
						},
					},
				},
				ServiceDiscovery: &appmesh.VirtualNodeSpecServiceDiscoveryArgs{
					Dns: &appmesh.VirtualNodeSpecServiceDiscoveryDnsArgs{
						Hostname: pulumi.String("serviceb.simpleapp.local"),
					},
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
import com.pulumi.aws.appmesh.VirtualNode;
import com.pulumi.aws.appmesh.VirtualNodeArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecServiceDiscoveryArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecServiceDiscoveryDnsArgs;
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
        var serviceb1 = new VirtualNode("serviceb1", VirtualNodeArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .spec(VirtualNodeSpecArgs.builder()
                .backends(VirtualNodeSpecBackendArgs.builder()
                    .virtualService(VirtualNodeSpecBackendVirtualServiceArgs.builder()
                        .virtualServiceName("servicea.simpleapp.local")
                        .build())
                    .build())
                .listeners(VirtualNodeSpecListenerArgs.builder()
                    .portMapping(VirtualNodeSpecListenerPortMappingArgs.builder()
                        .port(8080)
                        .protocol("http")
                        .build())
                    .healthCheck(VirtualNodeSpecListenerHealthCheckArgs.builder()
                        .protocol("http")
                        .path("/ping")
                        .healthyThreshold(2)
                        .unhealthyThreshold(2)
                        .timeoutMillis(2000)
                        .intervalMillis(5000)
                        .build())
                    .build())
                .serviceDiscovery(VirtualNodeSpecServiceDiscoveryArgs.builder()
                    .dns(VirtualNodeSpecServiceDiscoveryDnsArgs.builder()
                        .hostname("serviceb.simpleapp.local")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  serviceb1:
    type: aws:appmesh:VirtualNode
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      spec:
        backends:
          - virtualService:
              virtualServiceName: servicea.simpleapp.local
        listeners:
          - portMapping:
              port: 8080
              protocol: http
            healthCheck:
              protocol: http
              path: /ping
              healthyThreshold: 2
              unhealthyThreshold: 2
              timeoutMillis: 2000
              intervalMillis: 5000
        serviceDiscovery:
          dns:
            hostname: serviceb.simpleapp.local
```
{{% /example %}}
{{% example %}}
### Logging

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb1 = new aws.appmesh.VirtualNode("serviceb1", {
    meshName: aws_appmesh_mesh.simple.id,
    spec: {
        backends: [{
            virtualService: {
                virtualServiceName: "servicea.simpleapp.local",
            },
        }],
        listeners: [{
            portMapping: {
                port: 8080,
                protocol: "http",
            },
        }],
        serviceDiscovery: {
            dns: {
                hostname: "serviceb.simpleapp.local",
            },
        },
        logging: {
            accessLog: {
                file: {
                    path: "/dev/stdout",
                },
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

serviceb1 = aws.appmesh.VirtualNode("serviceb1",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    spec=aws.appmesh.VirtualNodeSpecArgs(
        backends=[aws.appmesh.VirtualNodeSpecBackendArgs(
            virtual_service=aws.appmesh.VirtualNodeSpecBackendVirtualServiceArgs(
                virtual_service_name="servicea.simpleapp.local",
            ),
        )],
        listeners=[aws.appmesh.VirtualNodeSpecListenerArgs(
            port_mapping=aws.appmesh.VirtualNodeSpecListenerPortMappingArgs(
                port=8080,
                protocol="http",
            ),
        )],
        service_discovery=aws.appmesh.VirtualNodeSpecServiceDiscoveryArgs(
            dns=aws.appmesh.VirtualNodeSpecServiceDiscoveryDnsArgs(
                hostname="serviceb.simpleapp.local",
            ),
        ),
        logging=aws.appmesh.VirtualNodeSpecLoggingArgs(
            access_log=aws.appmesh.VirtualNodeSpecLoggingAccessLogArgs(
                file=aws.appmesh.VirtualNodeSpecLoggingAccessLogFileArgs(
                    path="/dev/stdout",
                ),
            ),
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var serviceb1 = new Aws.AppMesh.VirtualNode("serviceb1", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        Spec = new Aws.AppMesh.Inputs.VirtualNodeSpecArgs
        {
            Backends = new[]
            {
                new Aws.AppMesh.Inputs.VirtualNodeSpecBackendArgs
                {
                    VirtualService = new Aws.AppMesh.Inputs.VirtualNodeSpecBackendVirtualServiceArgs
                    {
                        VirtualServiceName = "servicea.simpleapp.local",
                    },
                },
            },
            Listeners = new[]
            {
                new Aws.AppMesh.Inputs.VirtualNodeSpecListenerArgs
                {
                    PortMapping = new Aws.AppMesh.Inputs.VirtualNodeSpecListenerPortMappingArgs
                    {
                        Port = 8080,
                        Protocol = "http",
                    },
                },
            },
            ServiceDiscovery = new Aws.AppMesh.Inputs.VirtualNodeSpecServiceDiscoveryArgs
            {
                Dns = new Aws.AppMesh.Inputs.VirtualNodeSpecServiceDiscoveryDnsArgs
                {
                    Hostname = "serviceb.simpleapp.local",
                },
            },
            Logging = new Aws.AppMesh.Inputs.VirtualNodeSpecLoggingArgs
            {
                AccessLog = new Aws.AppMesh.Inputs.VirtualNodeSpecLoggingAccessLogArgs
                {
                    File = new Aws.AppMesh.Inputs.VirtualNodeSpecLoggingAccessLogFileArgs
                    {
                        Path = "/dev/stdout",
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appmesh"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appmesh.NewVirtualNode(ctx, "serviceb1", &appmesh.VirtualNodeArgs{
			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
			Spec: &appmesh.VirtualNodeSpecArgs{
				Backends: appmesh.VirtualNodeSpecBackendArray{
					&appmesh.VirtualNodeSpecBackendArgs{
						VirtualService: &appmesh.VirtualNodeSpecBackendVirtualServiceArgs{
							VirtualServiceName: pulumi.String("servicea.simpleapp.local"),
						},
					},
				},
				Listeners: appmesh.VirtualNodeSpecListenerArray{
					&appmesh.VirtualNodeSpecListenerArgs{
						PortMapping: &appmesh.VirtualNodeSpecListenerPortMappingArgs{
							Port:     pulumi.Int(8080),
							Protocol: pulumi.String("http"),
						},
					},
				},
				ServiceDiscovery: &appmesh.VirtualNodeSpecServiceDiscoveryArgs{
					Dns: &appmesh.VirtualNodeSpecServiceDiscoveryDnsArgs{
						Hostname: pulumi.String("serviceb.simpleapp.local"),
					},
				},
				Logging: &appmesh.VirtualNodeSpecLoggingArgs{
					AccessLog: &appmesh.VirtualNodeSpecLoggingAccessLogArgs{
						File: &appmesh.VirtualNodeSpecLoggingAccessLogFileArgs{
							Path: pulumi.String("/dev/stdout"),
						},
					},
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
import com.pulumi.aws.appmesh.VirtualNode;
import com.pulumi.aws.appmesh.VirtualNodeArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecServiceDiscoveryArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecServiceDiscoveryDnsArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecLoggingArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecLoggingAccessLogArgs;
import com.pulumi.aws.appmesh.inputs.VirtualNodeSpecLoggingAccessLogFileArgs;
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
        var serviceb1 = new VirtualNode("serviceb1", VirtualNodeArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .spec(VirtualNodeSpecArgs.builder()
                .backends(VirtualNodeSpecBackendArgs.builder()
                    .virtualService(VirtualNodeSpecBackendVirtualServiceArgs.builder()
                        .virtualServiceName("servicea.simpleapp.local")
                        .build())
                    .build())
                .listeners(VirtualNodeSpecListenerArgs.builder()
                    .portMapping(VirtualNodeSpecListenerPortMappingArgs.builder()
                        .port(8080)
                        .protocol("http")
                        .build())
                    .build())
                .serviceDiscovery(VirtualNodeSpecServiceDiscoveryArgs.builder()
                    .dns(VirtualNodeSpecServiceDiscoveryDnsArgs.builder()
                        .hostname("serviceb.simpleapp.local")
                        .build())
                    .build())
                .logging(VirtualNodeSpecLoggingArgs.builder()
                    .accessLog(VirtualNodeSpecLoggingAccessLogArgs.builder()
                        .file(VirtualNodeSpecLoggingAccessLogFileArgs.builder()
                            .path("/dev/stdout")
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  serviceb1:
    type: aws:appmesh:VirtualNode
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      spec:
        backends:
          - virtualService:
              virtualServiceName: servicea.simpleapp.local
        listeners:
          - portMapping:
              port: 8080
              protocol: http
        serviceDiscovery:
          dns:
            hostname: serviceb.simpleapp.local
        logging:
          accessLog:
            file:
              path: /dev/stdout
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Mesh virtual nodes using `mesh_name` together with the virtual node's `name`. For example:

```sh
 $ pulumi import aws:appmesh/virtualNode:VirtualNode serviceb1 simpleapp/serviceBv1
```
 