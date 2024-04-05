Provides information on Service Catalog Portfolio Constraints.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.servicecatalog.getPortfolioConstraints({
    portfolioId: "port-3lli3b3an",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.get_portfolio_constraints(portfolio_id="port-3lli3b3an")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ServiceCatalog.GetPortfolioConstraints.Invoke(new()
    {
        PortfolioId = "port-3lli3b3an",
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
		_, err := servicecatalog.GetPortfolioConstraints(ctx, &servicecatalog.GetPortfolioConstraintsArgs{
			PortfolioId: "port-3lli3b3an",
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
import com.pulumi.aws.servicecatalog.ServicecatalogFunctions;
import com.pulumi.aws.servicecatalog.inputs.GetPortfolioConstraintsArgs;
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
        final var example = ServicecatalogFunctions.getPortfolioConstraints(GetPortfolioConstraintsArgs.builder()
            .portfolioId("port-3lli3b3an")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:servicecatalog:getPortfolioConstraints
      Arguments:
        portfolioId: port-3lli3b3an
```
{{% /example %}}
{{% /examples %}}