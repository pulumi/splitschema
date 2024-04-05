Provides a resource to create a Service Catalog Portfolio.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const portfolio = new aws.servicecatalog.Portfolio("portfolio", {
    description: "List of my organizations apps",
    providerName: "Brett",
});
```
```python
import pulumi
import pulumi_aws as aws

portfolio = aws.servicecatalog.Portfolio("portfolio",
    description="List of my organizations apps",
    provider_name="Brett")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var portfolio = new Aws.ServiceCatalog.Portfolio("portfolio", new()
    {
        Description = "List of my organizations apps",
        ProviderName = "Brett",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicecatalog"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicecatalog.NewPortfolio(ctx, "portfolio", &servicecatalog.PortfolioArgs{
			Description:  pulumi.String("List of my organizations apps"),
			ProviderName: pulumi.String("Brett"),
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
import com.pulumi.aws.servicecatalog.Portfolio;
import com.pulumi.aws.servicecatalog.PortfolioArgs;
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
        var portfolio = new Portfolio("portfolio", PortfolioArgs.builder()        
            .description("List of my organizations apps")
            .providerName("Brett")
            .build());

    }
}
```
```yaml
resources:
  portfolio:
    type: aws:servicecatalog:Portfolio
    properties:
      description: List of my organizations apps
      providerName: Brett
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Service Catalog Portfolios using the Service Catalog Portfolio `id`. For example:

```sh
 $ pulumi import aws:servicecatalog/portfolio:Portfolio testfolio port-12344321
```
 