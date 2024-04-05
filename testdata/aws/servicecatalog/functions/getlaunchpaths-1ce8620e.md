Lists the paths to the specified product. A path is how the user has access to a specified product, and is necessary when provisioning a product. A path also determines the constraints put on the product.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.servicecatalog.getLaunchPaths({
    productId: "prod-yakog5pdriver",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.get_launch_paths(product_id="prod-yakog5pdriver")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ServiceCatalog.GetLaunchPaths.Invoke(new()
    {
        ProductId = "prod-yakog5pdriver",
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
		_, err := servicecatalog.GetLaunchPaths(ctx, &servicecatalog.GetLaunchPathsArgs{
			ProductId: "prod-yakog5pdriver",
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
import com.pulumi.aws.servicecatalog.inputs.GetLaunchPathsArgs;
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
        final var example = ServicecatalogFunctions.getLaunchPaths(GetLaunchPathsArgs.builder()
            .productId("prod-yakog5pdriver")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:servicecatalog:getLaunchPaths
      Arguments:
        productId: prod-yakog5pdriver
```
{{% /example %}}
{{% /examples %}}