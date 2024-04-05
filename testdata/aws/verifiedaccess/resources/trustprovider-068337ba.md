Resource for managing a Verified Access Trust Provider.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.TrustProvider("example", {
    policyReferenceName: "example",
    trustProviderType: "user",
    userTrustProviderType: "iam-identity-center",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.TrustProvider("example",
    policy_reference_name="example",
    trust_provider_type="user",
    user_trust_provider_type="iam-identity-center")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.TrustProvider("example", new()
    {
        PolicyReferenceName = "example",
        TrustProviderType = "user",
        UserTrustProviderType = "iam-identity-center",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewTrustProvider(ctx, "example", &verifiedaccess.TrustProviderArgs{
			PolicyReferenceName:   pulumi.String("example"),
			TrustProviderType:     pulumi.String("user"),
			UserTrustProviderType: pulumi.String("iam-identity-center"),
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
import com.pulumi.aws.verifiedaccess.TrustProvider;
import com.pulumi.aws.verifiedaccess.TrustProviderArgs;
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
        var example = new TrustProvider("example", TrustProviderArgs.builder()        
            .policyReferenceName("example")
            .trustProviderType("user")
            .userTrustProviderType("iam-identity-center")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:TrustProvider
    properties:
      policyReferenceName: example
      trustProviderType: user
      userTrustProviderType: iam-identity-center
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transfer Workflows using the

`id`. For example:

```sh
 $ pulumi import aws:verifiedaccess/trustProvider:TrustProvider example vatp-8012925589
```
 