Retrieve information about an EKS Cluster.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

export = async () => {
    const example = await aws.eks.getCluster({
        name: "example",
    });
    return {
        endpoint: example.endpoint,
        "kubeconfig-certificate-authority-data": example.certificateAuthorities?.[0]?.data,
        "identity-oidc-issuer": example.identities?.[0]?.oidcs?.[0]?.issuer,
    };
}
```
```python
import pulumi
import pulumi_aws as aws

example = aws.eks.get_cluster(name="example")
pulumi.export("endpoint", example.endpoint)
pulumi.export("kubeconfig-certificate-authority-data", example.certificate_authorities[0].data)
pulumi.export("identity-oidc-issuer", example.identities[0].oidcs[0].issuer)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Eks.GetCluster.Invoke(new()
    {
        Name = "example",
    });

    return new Dictionary<string, object?>
    {
        ["endpoint"] = example.Apply(getClusterResult => getClusterResult.Endpoint),
        ["kubeconfig-certificate-authority-data"] = example.Apply(getClusterResult => getClusterResult.CertificateAuthorities[0]?.Data),
        ["identity-oidc-issuer"] = example.Apply(getClusterResult => getClusterResult.Identities[0]?.Oidcs[0]?.Issuer),
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
		example, err := eks.LookupCluster(ctx, &eks.LookupClusterArgs{
			Name: "example",
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("endpoint", example.Endpoint)
		ctx.Export("kubeconfig-certificate-authority-data", example.CertificateAuthorities[0].Data)
		ctx.Export("identity-oidc-issuer", example.Identities[0].Oidcs[0].Issuer)
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
import com.pulumi.aws.eks.inputs.GetClusterArgs;
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
        final var example = EksFunctions.getCluster(GetClusterArgs.builder()
            .name("example")
            .build());

        ctx.export("endpoint", example.applyValue(getClusterResult -> getClusterResult.endpoint()));
        ctx.export("kubeconfig-certificate-authority-data", example.applyValue(getClusterResult -> getClusterResult.certificateAuthorities()[0].data()));
        ctx.export("identity-oidc-issuer", example.applyValue(getClusterResult -> getClusterResult.identities()[0].oidcs()[0].issuer()));
    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:eks:getCluster
      Arguments:
        name: example
outputs:
  endpoint: ${example.endpoint}
  kubeconfig-certificate-authority-data: ${example.certificateAuthorities[0].data}
  # Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.
  identity-oidc-issuer: ${example.identities[0].oidcs[0].issuer}
```
{{% /example %}}
{{% /examples %}}