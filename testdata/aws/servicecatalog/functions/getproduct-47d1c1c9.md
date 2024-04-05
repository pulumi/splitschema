Use this data source to retrieve information about a Service Catalog product.

> **NOTE:** A "provisioning artifact" is also known as a "version," and a "distributor" is also known as a "vendor."

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.servicecatalog.getProduct({
    id: "prod-dnigbtea24ste",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.get_product(id="prod-dnigbtea24ste")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ServiceCatalog.GetProduct.Invoke(new()
    {
        Id = "prod-dnigbtea24ste",
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
		_, err := servicecatalog.LookupProduct(ctx, &servicecatalog.LookupProductArgs{
			Id: "prod-dnigbtea24ste",
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
import com.pulumi.aws.servicecatalog.inputs.GetProductArgs;
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
        final var example = ServicecatalogFunctions.getProduct(GetProductArgs.builder()
            .id("prod-dnigbtea24ste")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:servicecatalog:getProduct
      Arguments:
        id: prod-dnigbtea24ste
```
{{% /example %}}
{{% /examples %}}