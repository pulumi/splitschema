Manages a Service Catalog Product Portfolio Association.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicecatalog.ProductPortfolioAssociation("example", {
    portfolioId: "port-68656c6c6f",
    productId: "prod-dnigbtea24ste",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.ProductPortfolioAssociation("example",
    portfolio_id="port-68656c6c6f",
    product_id="prod-dnigbtea24ste")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceCatalog.ProductPortfolioAssociation("example", new()
    {
        PortfolioId = "port-68656c6c6f",
        ProductId = "prod-dnigbtea24ste",
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
		_, err := servicecatalog.NewProductPortfolioAssociation(ctx, "example", &servicecatalog.ProductPortfolioAssociationArgs{
			PortfolioId: pulumi.String("port-68656c6c6f"),
			ProductId:   pulumi.String("prod-dnigbtea24ste"),
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
import com.pulumi.aws.servicecatalog.ProductPortfolioAssociation;
import com.pulumi.aws.servicecatalog.ProductPortfolioAssociationArgs;
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
        var example = new ProductPortfolioAssociation("example", ProductPortfolioAssociationArgs.builder()        
            .portfolioId("port-68656c6c6f")
            .productId("prod-dnigbtea24ste")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicecatalog:ProductPortfolioAssociation
    properties:
      portfolioId: port-68656c6c6f
      productId: prod-dnigbtea24ste
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicecatalog_product_portfolio_association` using the accept language, portfolio ID, and product ID. For example:

```sh
 $ pulumi import aws:servicecatalog/productPortfolioAssociation:ProductPortfolioAssociation example en:port-68656c6c6f:prod-dnigbtea24ste
```
 