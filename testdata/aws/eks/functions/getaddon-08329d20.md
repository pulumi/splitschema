Retrieve information about an EKS add-on.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.eks.getAddon({
    addonName: "vpc-cni",
    clusterName: aws_eks_cluster.example.name,
});
export const eksAddonOutputs = aws_eks_addon.example;
```
```python
import pulumi
import pulumi_aws as aws

example = aws.eks.get_addon(addon_name="vpc-cni",
    cluster_name=aws_eks_cluster["example"]["name"])
pulumi.export("eksAddonOutputs", aws_eks_addon["example"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Eks.GetAddon.Invoke(new()
    {
        AddonName = "vpc-cni",
        ClusterName = aws_eks_cluster.Example.Name,
    });

    return new Dictionary<string, object?>
    {
        ["eksAddonOutputs"] = aws_eks_addon.Example,
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
		_, err := eks.LookupAddon(ctx, &eks.LookupAddonArgs{
			AddonName:   "vpc-cni",
			ClusterName: aws_eks_cluster.Example.Name,
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("eksAddonOutputs", aws_eks_addon.Example)
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
import com.pulumi.aws.eks.inputs.GetAddonArgs;
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
        final var example = EksFunctions.getAddon(GetAddonArgs.builder()
            .addonName("vpc-cni")
            .clusterName(aws_eks_cluster.example().name())
            .build());

        ctx.export("eksAddonOutputs", aws_eks_addon.example());
    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:eks:getAddon
      Arguments:
        addonName: vpc-cni
        clusterName: ${aws_eks_cluster.example.name}
outputs:
  eksAddonOutputs: ${aws_eks_addon.example}
```
{{% /example %}}
{{% /examples %}}