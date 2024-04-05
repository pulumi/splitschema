Provides an AWS App Mesh route resource.

{{% examples %}}
## Example Usage
{{% example %}}
### HTTP Routing

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb = new aws.appmesh.Route("serviceb", {
    meshName: aws_appmesh_mesh.simple.id,
    virtualRouterName: aws_appmesh_virtual_router.serviceb.name,
    spec: {
        httpRoute: {
            match: {
                prefix: "/",
            },
            action: {
                weightedTargets: [
                    {
                        virtualNode: aws_appmesh_virtual_node.serviceb1.name,
                        weight: 90,
                    },
                    {
                        virtualNode: aws_appmesh_virtual_node.serviceb2.name,
                        weight: 10,
                    },
                ],
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

serviceb = aws.appmesh.Route("serviceb",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    virtual_router_name=aws_appmesh_virtual_router["serviceb"]["name"],
    spec=aws.appmesh.RouteSpecArgs(
        http_route=aws.appmesh.RouteSpecHttpRouteArgs(
            match=aws.appmesh.RouteSpecHttpRouteMatchArgs(
                prefix="/",
            ),
            action=aws.appmesh.RouteSpecHttpRouteActionArgs(
                weighted_targets=[
                    aws.appmesh.RouteSpecHttpRouteActionWeightedTargetArgs(
                        virtual_node=aws_appmesh_virtual_node["serviceb1"]["name"],
                        weight=90,
                    ),
                    aws.appmesh.RouteSpecHttpRouteActionWeightedTargetArgs(
                        virtual_node=aws_appmesh_virtual_node["serviceb2"]["name"],
                        weight=10,
                    ),
                ],
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
    var serviceb = new Aws.AppMesh.Route("serviceb", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
        Spec = new Aws.AppMesh.Inputs.RouteSpecArgs
        {
            HttpRoute = new Aws.AppMesh.Inputs.RouteSpecHttpRouteArgs
            {
                Match = new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchArgs
                {
                    Prefix = "/",
                },
                Action = new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionArgs
                {
                    WeightedTargets = new[]
                    {
                        new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionWeightedTargetArgs
                        {
                            VirtualNode = aws_appmesh_virtual_node.Serviceb1.Name,
                            Weight = 90,
                        },
                        new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionWeightedTargetArgs
                        {
                            VirtualNode = aws_appmesh_virtual_node.Serviceb2.Name,
                            Weight = 10,
                        },
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
		_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
			MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
			VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
			Spec: &appmesh.RouteSpecArgs{
				HttpRoute: &appmesh.RouteSpecHttpRouteArgs{
					Match: &appmesh.RouteSpecHttpRouteMatchArgs{
						Prefix: pulumi.String("/"),
					},
					Action: &appmesh.RouteSpecHttpRouteActionArgs{
						WeightedTargets: appmesh.RouteSpecHttpRouteActionWeightedTargetArray{
							&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
								VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
								Weight:      pulumi.Int(90),
							},
							&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
								VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb2.Name),
								Weight:      pulumi.Int(10),
							},
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
import com.pulumi.aws.appmesh.Route;
import com.pulumi.aws.appmesh.RouteArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteMatchArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteActionArgs;
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
        var serviceb = new Route("serviceb", RouteArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .virtualRouterName(aws_appmesh_virtual_router.serviceb().name())
            .spec(RouteSpecArgs.builder()
                .httpRoute(RouteSpecHttpRouteArgs.builder()
                    .match(RouteSpecHttpRouteMatchArgs.builder()
                        .prefix("/")
                        .build())
                    .action(RouteSpecHttpRouteActionArgs.builder()
                        .weightedTargets(                        
                            RouteSpecHttpRouteActionWeightedTargetArgs.builder()
                                .virtualNode(aws_appmesh_virtual_node.serviceb1().name())
                                .weight(90)
                                .build(),
                            RouteSpecHttpRouteActionWeightedTargetArgs.builder()
                                .virtualNode(aws_appmesh_virtual_node.serviceb2().name())
                                .weight(10)
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
  serviceb:
    type: aws:appmesh:Route
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      virtualRouterName: ${aws_appmesh_virtual_router.serviceb.name}
      spec:
        httpRoute:
          match:
            prefix: /
          action:
            weightedTargets:
              - virtualNode: ${aws_appmesh_virtual_node.serviceb1.name}
                weight: 90
              - virtualNode: ${aws_appmesh_virtual_node.serviceb2.name}
                weight: 10
```
{{% /example %}}
{{% example %}}
### HTTP Header Routing

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb = new aws.appmesh.Route("serviceb", {
    meshName: aws_appmesh_mesh.simple.id,
    virtualRouterName: aws_appmesh_virtual_router.serviceb.name,
    spec: {
        httpRoute: {
            match: {
                method: "POST",
                prefix: "/",
                scheme: "https",
                headers: [{
                    name: "clientRequestId",
                    match: {
                        prefix: "123",
                    },
                }],
            },
            action: {
                weightedTargets: [{
                    virtualNode: aws_appmesh_virtual_node.serviceb.name,
                    weight: 100,
                }],
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

serviceb = aws.appmesh.Route("serviceb",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    virtual_router_name=aws_appmesh_virtual_router["serviceb"]["name"],
    spec=aws.appmesh.RouteSpecArgs(
        http_route=aws.appmesh.RouteSpecHttpRouteArgs(
            match=aws.appmesh.RouteSpecHttpRouteMatchArgs(
                method="POST",
                prefix="/",
                scheme="https",
                headers=[aws.appmesh.RouteSpecHttpRouteMatchHeaderArgs(
                    name="clientRequestId",
                    match=aws.appmesh.RouteSpecHttpRouteMatchHeaderMatchArgs(
                        prefix="123",
                    ),
                )],
            ),
            action=aws.appmesh.RouteSpecHttpRouteActionArgs(
                weighted_targets=[aws.appmesh.RouteSpecHttpRouteActionWeightedTargetArgs(
                    virtual_node=aws_appmesh_virtual_node["serviceb"]["name"],
                    weight=100,
                )],
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
    var serviceb = new Aws.AppMesh.Route("serviceb", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
        Spec = new Aws.AppMesh.Inputs.RouteSpecArgs
        {
            HttpRoute = new Aws.AppMesh.Inputs.RouteSpecHttpRouteArgs
            {
                Match = new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchArgs
                {
                    Method = "POST",
                    Prefix = "/",
                    Scheme = "https",
                    Headers = new[]
                    {
                        new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchHeaderArgs
                        {
                            Name = "clientRequestId",
                            Match = new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchHeaderMatchArgs
                            {
                                Prefix = "123",
                            },
                        },
                    },
                },
                Action = new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionArgs
                {
                    WeightedTargets = new[]
                    {
                        new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionWeightedTargetArgs
                        {
                            VirtualNode = aws_appmesh_virtual_node.Serviceb.Name,
                            Weight = 100,
                        },
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
		_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
			MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
			VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
			Spec: &appmesh.RouteSpecArgs{
				HttpRoute: &appmesh.RouteSpecHttpRouteArgs{
					Match: &appmesh.RouteSpecHttpRouteMatchArgs{
						Method: pulumi.String("POST"),
						Prefix: pulumi.String("/"),
						Scheme: pulumi.String("https"),
						Headers: appmesh.RouteSpecHttpRouteMatchHeaderArray{
							&appmesh.RouteSpecHttpRouteMatchHeaderArgs{
								Name: pulumi.String("clientRequestId"),
								Match: &appmesh.RouteSpecHttpRouteMatchHeaderMatchArgs{
									Prefix: pulumi.String("123"),
								},
							},
						},
					},
					Action: &appmesh.RouteSpecHttpRouteActionArgs{
						WeightedTargets: appmesh.RouteSpecHttpRouteActionWeightedTargetArray{
							&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
								VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb.Name),
								Weight:      pulumi.Int(100),
							},
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
import com.pulumi.aws.appmesh.Route;
import com.pulumi.aws.appmesh.RouteArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteMatchArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteActionArgs;
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
        var serviceb = new Route("serviceb", RouteArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .virtualRouterName(aws_appmesh_virtual_router.serviceb().name())
            .spec(RouteSpecArgs.builder()
                .httpRoute(RouteSpecHttpRouteArgs.builder()
                    .match(RouteSpecHttpRouteMatchArgs.builder()
                        .method("POST")
                        .prefix("/")
                        .scheme("https")
                        .headers(RouteSpecHttpRouteMatchHeaderArgs.builder()
                            .name("clientRequestId")
                            .match(RouteSpecHttpRouteMatchHeaderMatchArgs.builder()
                                .prefix("123")
                                .build())
                            .build())
                        .build())
                    .action(RouteSpecHttpRouteActionArgs.builder()
                        .weightedTargets(RouteSpecHttpRouteActionWeightedTargetArgs.builder()
                            .virtualNode(aws_appmesh_virtual_node.serviceb().name())
                            .weight(100)
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
  serviceb:
    type: aws:appmesh:Route
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      virtualRouterName: ${aws_appmesh_virtual_router.serviceb.name}
      spec:
        httpRoute:
          match:
            method: POST
            prefix: /
            scheme: https
            headers:
              - name: clientRequestId
                match:
                  prefix: '123'
          action:
            weightedTargets:
              - virtualNode: ${aws_appmesh_virtual_node.serviceb.name}
                weight: 100
```
{{% /example %}}
{{% example %}}
### Retry Policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb = new aws.appmesh.Route("serviceb", {
    meshName: aws_appmesh_mesh.simple.id,
    virtualRouterName: aws_appmesh_virtual_router.serviceb.name,
    spec: {
        httpRoute: {
            match: {
                prefix: "/",
            },
            retryPolicy: {
                httpRetryEvents: ["server-error"],
                maxRetries: 1,
                perRetryTimeout: {
                    unit: "s",
                    value: 15,
                },
            },
            action: {
                weightedTargets: [{
                    virtualNode: aws_appmesh_virtual_node.serviceb.name,
                    weight: 100,
                }],
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

serviceb = aws.appmesh.Route("serviceb",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    virtual_router_name=aws_appmesh_virtual_router["serviceb"]["name"],
    spec=aws.appmesh.RouteSpecArgs(
        http_route=aws.appmesh.RouteSpecHttpRouteArgs(
            match=aws.appmesh.RouteSpecHttpRouteMatchArgs(
                prefix="/",
            ),
            retry_policy=aws.appmesh.RouteSpecHttpRouteRetryPolicyArgs(
                http_retry_events=["server-error"],
                max_retries=1,
                per_retry_timeout=aws.appmesh.RouteSpecHttpRouteRetryPolicyPerRetryTimeoutArgs(
                    unit="s",
                    value=15,
                ),
            ),
            action=aws.appmesh.RouteSpecHttpRouteActionArgs(
                weighted_targets=[aws.appmesh.RouteSpecHttpRouteActionWeightedTargetArgs(
                    virtual_node=aws_appmesh_virtual_node["serviceb"]["name"],
                    weight=100,
                )],
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
    var serviceb = new Aws.AppMesh.Route("serviceb", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
        Spec = new Aws.AppMesh.Inputs.RouteSpecArgs
        {
            HttpRoute = new Aws.AppMesh.Inputs.RouteSpecHttpRouteArgs
            {
                Match = new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchArgs
                {
                    Prefix = "/",
                },
                RetryPolicy = new Aws.AppMesh.Inputs.RouteSpecHttpRouteRetryPolicyArgs
                {
                    HttpRetryEvents = new[]
                    {
                        "server-error",
                    },
                    MaxRetries = 1,
                    PerRetryTimeout = new Aws.AppMesh.Inputs.RouteSpecHttpRouteRetryPolicyPerRetryTimeoutArgs
                    {
                        Unit = "s",
                        Value = 15,
                    },
                },
                Action = new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionArgs
                {
                    WeightedTargets = new[]
                    {
                        new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionWeightedTargetArgs
                        {
                            VirtualNode = aws_appmesh_virtual_node.Serviceb.Name,
                            Weight = 100,
                        },
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
		_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
			MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
			VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
			Spec: &appmesh.RouteSpecArgs{
				HttpRoute: &appmesh.RouteSpecHttpRouteArgs{
					Match: &appmesh.RouteSpecHttpRouteMatchArgs{
						Prefix: pulumi.String("/"),
					},
					RetryPolicy: &appmesh.RouteSpecHttpRouteRetryPolicyArgs{
						HttpRetryEvents: pulumi.StringArray{
							pulumi.String("server-error"),
						},
						MaxRetries: pulumi.Int(1),
						PerRetryTimeout: &appmesh.RouteSpecHttpRouteRetryPolicyPerRetryTimeoutArgs{
							Unit:  pulumi.String("s"),
							Value: pulumi.Int(15),
						},
					},
					Action: &appmesh.RouteSpecHttpRouteActionArgs{
						WeightedTargets: appmesh.RouteSpecHttpRouteActionWeightedTargetArray{
							&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
								VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb.Name),
								Weight:      pulumi.Int(100),
							},
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
import com.pulumi.aws.appmesh.Route;
import com.pulumi.aws.appmesh.RouteArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteMatchArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteRetryPolicyArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteRetryPolicyPerRetryTimeoutArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecHttpRouteActionArgs;
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
        var serviceb = new Route("serviceb", RouteArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .virtualRouterName(aws_appmesh_virtual_router.serviceb().name())
            .spec(RouteSpecArgs.builder()
                .httpRoute(RouteSpecHttpRouteArgs.builder()
                    .match(RouteSpecHttpRouteMatchArgs.builder()
                        .prefix("/")
                        .build())
                    .retryPolicy(RouteSpecHttpRouteRetryPolicyArgs.builder()
                        .httpRetryEvents("server-error")
                        .maxRetries(1)
                        .perRetryTimeout(RouteSpecHttpRouteRetryPolicyPerRetryTimeoutArgs.builder()
                            .unit("s")
                            .value(15)
                            .build())
                        .build())
                    .action(RouteSpecHttpRouteActionArgs.builder()
                        .weightedTargets(RouteSpecHttpRouteActionWeightedTargetArgs.builder()
                            .virtualNode(aws_appmesh_virtual_node.serviceb().name())
                            .weight(100)
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
  serviceb:
    type: aws:appmesh:Route
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      virtualRouterName: ${aws_appmesh_virtual_router.serviceb.name}
      spec:
        httpRoute:
          match:
            prefix: /
          retryPolicy:
            httpRetryEvents:
              - server-error
            maxRetries: 1
            perRetryTimeout:
              unit: s
              value: 15
          action:
            weightedTargets:
              - virtualNode: ${aws_appmesh_virtual_node.serviceb.name}
                weight: 100
```
{{% /example %}}
{{% example %}}
### TCP Routing

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceb = new aws.appmesh.Route("serviceb", {
    meshName: aws_appmesh_mesh.simple.id,
    virtualRouterName: aws_appmesh_virtual_router.serviceb.name,
    spec: {
        tcpRoute: {
            action: {
                weightedTargets: [{
                    virtualNode: aws_appmesh_virtual_node.serviceb1.name,
                    weight: 100,
                }],
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

serviceb = aws.appmesh.Route("serviceb",
    mesh_name=aws_appmesh_mesh["simple"]["id"],
    virtual_router_name=aws_appmesh_virtual_router["serviceb"]["name"],
    spec=aws.appmesh.RouteSpecArgs(
        tcp_route=aws.appmesh.RouteSpecTcpRouteArgs(
            action=aws.appmesh.RouteSpecTcpRouteActionArgs(
                weighted_targets=[aws.appmesh.RouteSpecTcpRouteActionWeightedTargetArgs(
                    virtual_node=aws_appmesh_virtual_node["serviceb1"]["name"],
                    weight=100,
                )],
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
    var serviceb = new Aws.AppMesh.Route("serviceb", new()
    {
        MeshName = aws_appmesh_mesh.Simple.Id,
        VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
        Spec = new Aws.AppMesh.Inputs.RouteSpecArgs
        {
            TcpRoute = new Aws.AppMesh.Inputs.RouteSpecTcpRouteArgs
            {
                Action = new Aws.AppMesh.Inputs.RouteSpecTcpRouteActionArgs
                {
                    WeightedTargets = new[]
                    {
                        new Aws.AppMesh.Inputs.RouteSpecTcpRouteActionWeightedTargetArgs
                        {
                            VirtualNode = aws_appmesh_virtual_node.Serviceb1.Name,
                            Weight = 100,
                        },
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
		_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
			MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
			VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
			Spec: &appmesh.RouteSpecArgs{
				TcpRoute: &appmesh.RouteSpecTcpRouteArgs{
					Action: &appmesh.RouteSpecTcpRouteActionArgs{
						WeightedTargets: appmesh.RouteSpecTcpRouteActionWeightedTargetArray{
							&appmesh.RouteSpecTcpRouteActionWeightedTargetArgs{
								VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
								Weight:      pulumi.Int(100),
							},
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
import com.pulumi.aws.appmesh.Route;
import com.pulumi.aws.appmesh.RouteArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecTcpRouteArgs;
import com.pulumi.aws.appmesh.inputs.RouteSpecTcpRouteActionArgs;
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
        var serviceb = new Route("serviceb", RouteArgs.builder()        
            .meshName(aws_appmesh_mesh.simple().id())
            .virtualRouterName(aws_appmesh_virtual_router.serviceb().name())
            .spec(RouteSpecArgs.builder()
                .tcpRoute(RouteSpecTcpRouteArgs.builder()
                    .action(RouteSpecTcpRouteActionArgs.builder()
                        .weightedTargets(RouteSpecTcpRouteActionWeightedTargetArgs.builder()
                            .virtualNode(aws_appmesh_virtual_node.serviceb1().name())
                            .weight(100)
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
  serviceb:
    type: aws:appmesh:Route
    properties:
      meshName: ${aws_appmesh_mesh.simple.id}
      virtualRouterName: ${aws_appmesh_virtual_router.serviceb.name}
      spec:
        tcpRoute:
          action:
            weightedTargets:
              - virtualNode: ${aws_appmesh_virtual_node.serviceb1.name}
                weight: 100
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Mesh virtual routes using `mesh_name` and `virtual_router_name` together with the route's `name`. For example:

```sh
 $ pulumi import aws:appmesh/route:Route serviceb simpleapp/serviceB/serviceB-route
```
 