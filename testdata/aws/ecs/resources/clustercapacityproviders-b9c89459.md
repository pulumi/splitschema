Manages the capacity providers of an ECS Cluster.

More information about capacity providers can be found in the [ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCluster = new aws.ecs.Cluster("exampleCluster", {});
const exampleClusterCapacityProviders = new aws.ecs.ClusterCapacityProviders("exampleClusterCapacityProviders", {
    clusterName: exampleCluster.name,
    capacityProviders: ["FARGATE"],
    defaultCapacityProviderStrategies: [{
        base: 1,
        weight: 100,
        capacityProvider: "FARGATE",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_cluster = aws.ecs.Cluster("exampleCluster")
example_cluster_capacity_providers = aws.ecs.ClusterCapacityProviders("exampleClusterCapacityProviders",
    cluster_name=example_cluster.name,
    capacity_providers=["FARGATE"],
    default_capacity_provider_strategies=[aws.ecs.ClusterCapacityProvidersDefaultCapacityProviderStrategyArgs(
        base=1,
        weight=100,
        capacity_provider="FARGATE",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleCluster = new Aws.Ecs.Cluster("exampleCluster");

    var exampleClusterCapacityProviders = new Aws.Ecs.ClusterCapacityProviders("exampleClusterCapacityProviders", new()
    {
        ClusterName = exampleCluster.Name,
        CapacityProviders = new[]
        {
            "FARGATE",
        },
        DefaultCapacityProviderStrategies = new[]
        {
            new Aws.Ecs.Inputs.ClusterCapacityProvidersDefaultCapacityProviderStrategyArgs
            {
                Base = 1,
                Weight = 100,
                CapacityProvider = "FARGATE",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleCluster, err := ecs.NewCluster(ctx, "exampleCluster", nil)
		if err != nil {
			return err
		}
		_, err = ecs.NewClusterCapacityProviders(ctx, "exampleClusterCapacityProviders", &ecs.ClusterCapacityProvidersArgs{
			ClusterName: exampleCluster.Name,
			CapacityProviders: pulumi.StringArray{
				pulumi.String("FARGATE"),
			},
			DefaultCapacityProviderStrategies: ecs.ClusterCapacityProvidersDefaultCapacityProviderStrategyArray{
				&ecs.ClusterCapacityProvidersDefaultCapacityProviderStrategyArgs{
					Base:             pulumi.Int(1),
					Weight:           pulumi.Int(100),
					CapacityProvider: pulumi.String("FARGATE"),
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
import com.pulumi.aws.ecs.Cluster;
import com.pulumi.aws.ecs.ClusterCapacityProviders;
import com.pulumi.aws.ecs.ClusterCapacityProvidersArgs;
import com.pulumi.aws.ecs.inputs.ClusterCapacityProvidersDefaultCapacityProviderStrategyArgs;
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
        var exampleCluster = new Cluster("exampleCluster");

        var exampleClusterCapacityProviders = new ClusterCapacityProviders("exampleClusterCapacityProviders", ClusterCapacityProvidersArgs.builder()        
            .clusterName(exampleCluster.name())
            .capacityProviders("FARGATE")
            .defaultCapacityProviderStrategies(ClusterCapacityProvidersDefaultCapacityProviderStrategyArgs.builder()
                .base(1)
                .weight(100)
                .capacityProvider("FARGATE")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleCluster:
    type: aws:ecs:Cluster
  exampleClusterCapacityProviders:
    type: aws:ecs:ClusterCapacityProviders
    properties:
      clusterName: ${exampleCluster.name}
      capacityProviders:
        - FARGATE
      defaultCapacityProviderStrategies:
        - base: 1
          weight: 100
          capacityProvider: FARGATE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECS cluster capacity providers using the `cluster_name` attribute. For example:

```sh
 $ pulumi import aws:ecs/clusterCapacityProviders:ClusterCapacityProviders example my-cluster
```
 