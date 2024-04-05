Manages a Service Catalog Portfolio Share. Shares the specified portfolio with the specified account or organization node. You can share portfolios to an organization, an organizational unit, or a specific account.

If the portfolio share with the specified account or organization node already exists, using this resource to re-create the share will have no effect and will not return an error. You can then use this resource to update the share.

> **NOTE:** Shares to an organization node can only be created by the management account of an organization or by a delegated administrator. If a delegated admin is de-registered, they can no longer create portfolio shares.

> **NOTE:** AWSOrganizationsAccess must be enabled in order to create a portfolio share to an organization node.

> **NOTE:** You can't share a shared resource, including portfolios that contain a shared product.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicecatalog.PortfolioShare("example", {
    principalId: "012128675309",
    portfolioId: aws_servicecatalog_portfolio.example.id,
    type: "ACCOUNT",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.PortfolioShare("example",
    principal_id="012128675309",
    portfolio_id=aws_servicecatalog_portfolio["example"]["id"],
    type="ACCOUNT")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceCatalog.PortfolioShare("example", new()
    {
        PrincipalId = "012128675309",
        PortfolioId = aws_servicecatalog_portfolio.Example.Id,
        Type = "ACCOUNT",
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
		_, err := servicecatalog.NewPortfolioShare(ctx, "example", &servicecatalog.PortfolioShareArgs{
			PrincipalId: pulumi.String("012128675309"),
			PortfolioId: pulumi.Any(aws_servicecatalog_portfolio.Example.Id),
			Type:        pulumi.String("ACCOUNT"),
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
import com.pulumi.aws.servicecatalog.PortfolioShare;
import com.pulumi.aws.servicecatalog.PortfolioShareArgs;
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
        var example = new PortfolioShare("example", PortfolioShareArgs.builder()        
            .principalId("012128675309")
            .portfolioId(aws_servicecatalog_portfolio.example().id())
            .type("ACCOUNT")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicecatalog:PortfolioShare
    properties:
      principalId: '012128675309'
      portfolioId: ${aws_servicecatalog_portfolio.example.id}
      type: ACCOUNT
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicecatalog_portfolio_share` using the portfolio share ID. For example:

```sh
 $ pulumi import aws:servicecatalog/portfolioShare:PortfolioShare example port-12344321:ACCOUNT:123456789012
```
 