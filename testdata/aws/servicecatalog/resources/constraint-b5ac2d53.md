Manages a Service Catalog Constraint.

> **NOTE:** This resource does not associate a Service Catalog product and portfolio. However, the product and portfolio must be associated (see the `aws.servicecatalog.ProductPortfolioAssociation` resource) prior to creating a constraint or you will receive an error.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicecatalog.Constraint("example", {
    description: "Back off, man. I'm a scientist.",
    portfolioId: aws_servicecatalog_portfolio.example.id,
    productId: aws_servicecatalog_product.example.id,
    type: "LAUNCH",
    parameters: JSON.stringify({
        RoleArn: "arn:aws:iam::123456789012:role/LaunchRole",
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.servicecatalog.Constraint("example",
    description="Back off, man. I'm a scientist.",
    portfolio_id=aws_servicecatalog_portfolio["example"]["id"],
    product_id=aws_servicecatalog_product["example"]["id"],
    type="LAUNCH",
    parameters=json.dumps({
        "RoleArn": "arn:aws:iam::123456789012:role/LaunchRole",
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceCatalog.Constraint("example", new()
    {
        Description = "Back off, man. I'm a scientist.",
        PortfolioId = aws_servicecatalog_portfolio.Example.Id,
        ProductId = aws_servicecatalog_product.Example.Id,
        Type = "LAUNCH",
        Parameters = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["RoleArn"] = "arn:aws:iam::123456789012:role/LaunchRole",
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicecatalog"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"RoleArn": "arn:aws:iam::123456789012:role/LaunchRole",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = servicecatalog.NewConstraint(ctx, "example", &servicecatalog.ConstraintArgs{
			Description: pulumi.String("Back off, man. I'm a scientist."),
			PortfolioId: pulumi.Any(aws_servicecatalog_portfolio.Example.Id),
			ProductId:   pulumi.Any(aws_servicecatalog_product.Example.Id),
			Type:        pulumi.String("LAUNCH"),
			Parameters:  pulumi.String(json0),
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
import com.pulumi.aws.servicecatalog.Constraint;
import com.pulumi.aws.servicecatalog.ConstraintArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var example = new Constraint("example", ConstraintArgs.builder()        
            .description("Back off, man. I'm a scientist.")
            .portfolioId(aws_servicecatalog_portfolio.example().id())
            .productId(aws_servicecatalog_product.example().id())
            .type("LAUNCH")
            .parameters(serializeJson(
                jsonObject(
                    jsonProperty("RoleArn", "arn:aws:iam::123456789012:role/LaunchRole")
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicecatalog:Constraint
    properties:
      description: Back off, man. I'm a scientist.
      portfolioId: ${aws_servicecatalog_portfolio.example.id}
      productId: ${aws_servicecatalog_product.example.id}
      type: LAUNCH
      parameters:
        fn::toJSON:
          RoleArn: arn:aws:iam::123456789012:role/LaunchRole
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicecatalog_constraint` using the constraint ID. For example:

```sh
 $ pulumi import aws:servicecatalog/constraint:Constraint example cons-nmdkb6cgxfcrs
```
 