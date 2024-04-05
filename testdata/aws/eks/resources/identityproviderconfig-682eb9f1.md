Manages an EKS Identity Provider Configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.eks.IdentityProviderConfig("example", {
    clusterName: aws_eks_cluster.example.name,
    oidc: {
        clientId: "your client_id",
        identityProviderConfigName: "example",
        issuerUrl: "your issuer_url",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.eks.IdentityProviderConfig("example",
    cluster_name=aws_eks_cluster["example"]["name"],
    oidc=aws.eks.IdentityProviderConfigOidcArgs(
        client_id="your client_id",
        identity_provider_config_name="example",
        issuer_url="your issuer_url",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Eks.IdentityProviderConfig("example", new()
    {
        ClusterName = aws_eks_cluster.Example.Name,
        Oidc = new Aws.Eks.Inputs.IdentityProviderConfigOidcArgs
        {
            ClientId = "your client_id",
            IdentityProviderConfigName = "example",
            IssuerUrl = "your issuer_url",
        },
    });

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
		_, err := eks.NewIdentityProviderConfig(ctx, "example", &eks.IdentityProviderConfigArgs{
			ClusterName: pulumi.Any(aws_eks_cluster.Example.Name),
			Oidc: &eks.IdentityProviderConfigOidcArgs{
				ClientId:                   pulumi.String("your client_id"),
				IdentityProviderConfigName: pulumi.String("example"),
				IssuerUrl:                  pulumi.String("your issuer_url"),
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
import com.pulumi.aws.eks.IdentityProviderConfig;
import com.pulumi.aws.eks.IdentityProviderConfigArgs;
import com.pulumi.aws.eks.inputs.IdentityProviderConfigOidcArgs;
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
        var example = new IdentityProviderConfig("example", IdentityProviderConfigArgs.builder()        
            .clusterName(aws_eks_cluster.example().name())
            .oidc(IdentityProviderConfigOidcArgs.builder()
                .clientId("your client_id")
                .identityProviderConfigName("example")
                .issuerUrl("your issuer_url")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:eks:IdentityProviderConfig
    properties:
      clusterName: ${aws_eks_cluster.example.name}
      oidc:
        clientId: your client_id
        identityProviderConfigName: example
        issuerUrl: your issuer_url
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EKS Identity Provider Configurations using the `cluster_name` and `identity_provider_config_name` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:eks/identityProviderConfig:IdentityProviderConfig my_identity_provider_config my_cluster:my_identity_provider_config
```
 