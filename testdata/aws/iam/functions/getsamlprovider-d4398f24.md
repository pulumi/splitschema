This data source can be used to fetch information about a specific
IAM SAML provider. This will allow you to easily retrieve the metadata
document of an existing SAML provider.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.iam.getSamlProvider({
    arn: "arn:aws:iam::123456789:saml-provider/myprovider",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iam.get_saml_provider(arn="arn:aws:iam::123456789:saml-provider/myprovider")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Iam.GetSamlProvider.Invoke(new()
    {
        Arn = "arn:aws:iam::123456789:saml-provider/myprovider",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.LookupSamlProvider(ctx, &iam.LookupSamlProviderArgs{
			Arn: "arn:aws:iam::123456789:saml-provider/myprovider",
		}, nil)
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetSamlProviderArgs;
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
        final var example = IamFunctions.getSamlProvider(GetSamlProviderArgs.builder()
            .arn("arn:aws:iam::123456789:saml-provider/myprovider")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:iam:getSamlProvider
      Arguments:
        arn: arn:aws:iam::123456789:saml-provider/myprovider
```
{{% /example %}}
{{% /examples %}}