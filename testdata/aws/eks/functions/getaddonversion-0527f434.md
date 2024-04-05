Retrieve information about a specific EKS add-on version compatible with an EKS cluster version.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

export = async () => {
    const defaultAddonVersion = await aws.eks.getAddonVersion({
        addonName: "vpc-cni",
        kubernetesVersion: aws_eks_cluster.example.version,
    });
    const latestAddonVersion = await aws.eks.getAddonVersion({
        addonName: "vpc-cni",
        kubernetesVersion: aws_eks_cluster.example.version,
        mostRecent: true,
    });
    const vpcCni = new aws.eks.Addon("vpcCni", {
        clusterName: aws_eks_cluster.example.name,
        addonName: "vpc-cni",
        addonVersion: latestAddonVersion.version,
    });
    return {
        "default": defaultAddonVersion.version,
        latest: latestAddonVersion.version,
    };
}
```
```python
import pulumi
import pulumi_aws as aws

default_addon_version = aws.eks.get_addon_version(addon_name="vpc-cni",
    kubernetes_version=aws_eks_cluster["example"]["version"])
latest_addon_version = aws.eks.get_addon_version(addon_name="vpc-cni",
    kubernetes_version=aws_eks_cluster["example"]["version"],
    most_recent=True)
vpc_cni = aws.eks.Addon("vpcCni",
    cluster_name=aws_eks_cluster["example"]["name"],
    addon_name="vpc-cni",
    addon_version=latest_addon_version.version)
pulumi.export("default", default_addon_version.version)
pulumi.export("latest", latest_addon_version.version)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultAddonVersion = Aws.Eks.GetAddonVersion.Invoke(new()
    {
        AddonName = "vpc-cni",
        KubernetesVersion = aws_eks_cluster.Example.Version,
    });

    var latestAddonVersion = Aws.Eks.GetAddonVersion.Invoke(new()
    {
        AddonName = "vpc-cni",
        KubernetesVersion = aws_eks_cluster.Example.Version,
        MostRecent = true,
    });

    var vpcCni = new Aws.Eks.Addon("vpcCni", new()
    {
        ClusterName = aws_eks_cluster.Example.Name,
        AddonName = "vpc-cni",
        AddonVersion = latestAddonVersion.Apply(getAddonVersionResult => getAddonVersionResult.Version),
    });

    return new Dictionary<string, object?>
    {
        ["default"] = defaultAddonVersion.Apply(getAddonVersionResult => getAddonVersionResult.Version),
        ["latest"] = latestAddonVersion.Apply(getAddonVersionResult => getAddonVersionResult.Version),
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultAddonVersion, err := eks.GetAddonVersion(ctx, &eks.GetAddonVersionArgs{
			AddonName:         "vpc-cni",
			KubernetesVersion: aws_eks_cluster.Example.Version,
		}, nil)
		if err != nil {
			return err
		}
		latestAddonVersion, err := eks.GetAddonVersion(ctx, &eks.GetAddonVersionArgs{
			AddonName:         "vpc-cni",
			KubernetesVersion: aws_eks_cluster.Example.Version,
			MostRecent:        pulumi.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		_, err = eks.NewAddon(ctx, "vpcCni", &eks.AddonArgs{
			ClusterName:  pulumi.Any(aws_eks_cluster.Example.Name),
			AddonName:    pulumi.String("vpc-cni"),
			AddonVersion: *pulumi.String(latestAddonVersion.Version),
		})
		if err != nil {
			return err
		}
		ctx.Export("default", defaultAddonVersion.Version)
		ctx.Export("latest", latestAddonVersion.Version)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.eks.EksFunctions;
import com.pulumi.aws.eks.inputs.GetAddonVersionArgs;
import com.pulumi.aws.eks.Addon;
import com.pulumi.aws.eks.AddonArgs;
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
        final var defaultAddonVersion = EksFunctions.getAddonVersion(GetAddonVersionArgs.builder()
            .addonName("vpc-cni")
            .kubernetesVersion(aws_eks_cluster.example().version())
            .build());

        final var latestAddonVersion = EksFunctions.getAddonVersion(GetAddonVersionArgs.builder()
            .addonName("vpc-cni")
            .kubernetesVersion(aws_eks_cluster.example().version())
            .mostRecent(true)
            .build());

        var vpcCni = new Addon("vpcCni", AddonArgs.builder()        
            .clusterName(aws_eks_cluster.example().name())
            .addonName("vpc-cni")
            .addonVersion(latestAddonVersion.applyValue(getAddonVersionResult -> getAddonVersionResult.version()))
            .build());

        ctx.export("default", defaultAddonVersion.applyValue(getAddonVersionResult -> getAddonVersionResult.version()));
        ctx.export("latest", latestAddonVersion.applyValue(getAddonVersionResult -> getAddonVersionResult.version()));
    }
}
```
```yaml
resources:
  vpcCni:
    type: aws:eks:Addon
    properties:
      clusterName: ${aws_eks_cluster.example.name}
      addonName: vpc-cni
      addonVersion: ${latestAddonVersion.version}
variables:
  defaultAddonVersion:
    fn::invoke:
      Function: aws:eks:getAddonVersion
      Arguments:
        addonName: vpc-cni
        kubernetesVersion: ${aws_eks_cluster.example.version}
  latestAddonVersion:
    fn::invoke:
      Function: aws:eks:getAddonVersion
      Arguments:
        addonName: vpc-cni
        kubernetesVersion: ${aws_eks_cluster.example.version}
        mostRecent: true
outputs:
  default: ${defaultAddonVersion.version}
  latest: ${latestAddonVersion.version}
```
{{% /example %}}
{{% /examples %}}