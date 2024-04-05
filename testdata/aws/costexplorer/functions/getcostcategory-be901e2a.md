Provides details about a specific CostExplorer Cost Category.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.costexplorer.getCostCategory({
    costCategoryArn: "costCategoryARN",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.costexplorer.get_cost_category(cost_category_arn="costCategoryARN")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.CostExplorer.GetCostCategory.Invoke(new()
    {
        CostCategoryArn = "costCategoryARN",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/costexplorer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costexplorer.LookupCostCategory(ctx, &costexplorer.LookupCostCategoryArgs{
			CostCategoryArn: "costCategoryARN",
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
import com.pulumi.aws.costexplorer.CostexplorerFunctions;
import com.pulumi.aws.costexplorer.inputs.GetCostCategoryArgs;
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
        final var example = CostexplorerFunctions.getCostCategory(GetCostCategoryArgs.builder()
            .costCategoryArn("costCategoryARN")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:costexplorer:getCostCategory
      Arguments:
        costCategoryArn: costCategoryARN
```
{{% /example %}}
{{% /examples %}}