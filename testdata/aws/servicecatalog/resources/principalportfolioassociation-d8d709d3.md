Manages a Service Catalog Principal Portfolio Association.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicecatalog.PrincipalPortfolioAssociation("example", {
    portfolioId: "port-68656c6c6f",
    principalArn: "arn:aws:iam::123456789012:user/Eleanor",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.PrincipalPortfolioAssociation("example",
    portfolio_id="port-68656c6c6f",
    principal_arn="arn:aws:iam::123456789012:user/Eleanor")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceCatalog.PrincipalPortfolioAssociation("example", new()
    {
        PortfolioId = "port-68656c6c6f",
        PrincipalArn = "arn:aws:iam::123456789012:user/Eleanor",
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
		_, err := servicecatalog.NewPrincipalPortfolioAssociation(ctx, "example", &servicecatalog.PrincipalPortfolioAssociationArgs{
			PortfolioId:  pulumi.String("port-68656c6c6f"),
			PrincipalArn: pulumi.String("arn:aws:iam::123456789012:user/Eleanor"),
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
import com.pulumi.aws.servicecatalog.PrincipalPortfolioAssociation;
import com.pulumi.aws.servicecatalog.PrincipalPortfolioAssociationArgs;
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
        var example = new PrincipalPortfolioAssociation("example", PrincipalPortfolioAssociationArgs.builder()        
            .portfolioId("port-68656c6c6f")
            .principalArn("arn:aws:iam::123456789012:user/Eleanor")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicecatalog:PrincipalPortfolioAssociation
    properties:
      portfolioId: port-68656c6c6f
      principalArn: arn:aws:iam::123456789012:user/Eleanor
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicecatalog_principal_portfolio_association` using `accept_language`, `principal_arn`, `portfolio_id`, and `principal_type` separated by a comma. For example:

```sh
 $ pulumi import aws:servicecatalog/principalPortfolioAssociation:PrincipalPortfolioAssociation example en,arn:aws:iam::123456789012:user/Eleanor,port-68656c6c6f,IAM
```
 