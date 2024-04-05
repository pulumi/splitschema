Manages an EMR Containers (EMR on EKS) Virtual Cluster.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.emrcontainers.VirtualCluster("example", {containerProvider: {
    id: aws_eks_cluster.example.name,
    type: "EKS",
    info: {
        eksInfo: {
            namespace: "default",
        },
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.emrcontainers.VirtualCluster("example", container_provider=aws.emrcontainers.VirtualClusterContainerProviderArgs(
    id=aws_eks_cluster["example"]["name"],
    type="EKS",
    info=aws.emrcontainers.VirtualClusterContainerProviderInfoArgs(
        eks_info=aws.emrcontainers.VirtualClusterContainerProviderInfoEksInfoArgs(
            namespace="default",
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
    var example = new Aws.EmrContainers.VirtualCluster("example", new()
    {
        ContainerProvider = new Aws.EmrContainers.Inputs.VirtualClusterContainerProviderArgs
        {
            Id = aws_eks_cluster.Example.Name,
            Type = "EKS",
            Info = new Aws.EmrContainers.Inputs.VirtualClusterContainerProviderInfoArgs
            {
                EksInfo = new Aws.EmrContainers.Inputs.VirtualClusterContainerProviderInfoEksInfoArgs
                {
                    Namespace = "default",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emrcontainers"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emrcontainers.NewVirtualCluster(ctx, "example", &emrcontainers.VirtualClusterArgs{
			ContainerProvider: &emrcontainers.VirtualClusterContainerProviderArgs{
				Id:   pulumi.Any(aws_eks_cluster.Example.Name),
				Type: pulumi.String("EKS"),
				Info: &emrcontainers.VirtualClusterContainerProviderInfoArgs{
					EksInfo: &emrcontainers.VirtualClusterContainerProviderInfoEksInfoArgs{
						Namespace: pulumi.String("default"),
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
import com.pulumi.aws.emrcontainers.VirtualCluster;
import com.pulumi.aws.emrcontainers.VirtualClusterArgs;
import com.pulumi.aws.emrcontainers.inputs.VirtualClusterContainerProviderArgs;
import com.pulumi.aws.emrcontainers.inputs.VirtualClusterContainerProviderInfoArgs;
import com.pulumi.aws.emrcontainers.inputs.VirtualClusterContainerProviderInfoEksInfoArgs;
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
        var example = new VirtualCluster("example", VirtualClusterArgs.builder()        
            .containerProvider(VirtualClusterContainerProviderArgs.builder()
                .id(aws_eks_cluster.example().name())
                .type("EKS")
                .info(VirtualClusterContainerProviderInfoArgs.builder()
                    .eksInfo(VirtualClusterContainerProviderInfoEksInfoArgs.builder()
                        .namespace("default")
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
    type: aws:emrcontainers:VirtualCluster
    properties:
      containerProvider:
        id: ${aws_eks_cluster.example.name}
        type: EKS
        info:
          eksInfo:
            namespace: default
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EKS Clusters using the `id`. For example:

```sh
 $ pulumi import aws:emrcontainers/virtualCluster:VirtualCluster example a1b2c3d4e5f6g7h8i9j10k11l
```
 