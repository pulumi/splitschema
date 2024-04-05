Data source for managing an AWS Web Services Budgets Budget.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.budgets.getBudget({
    name: aws_budgets_budget.test.name,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.budgets.get_budget(name=aws_budgets_budget["test"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Budgets.GetBudget.Invoke(new()
    {
        Name = aws_budgets_budget.Test.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/budgets"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := budgets.LookupBudget(ctx, &budgets.LookupBudgetArgs{
			Name: aws_budgets_budget.Test.Name,
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
import com.pulumi.aws.budgets.BudgetsFunctions;
import com.pulumi.aws.budgets.inputs.GetBudgetArgs;
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
        final var test = BudgetsFunctions.getBudget(GetBudgetArgs.builder()
            .name(aws_budgets_budget.test().name())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:budgets:getBudget
      Arguments:
        name: ${aws_budgets_budget.test.name}
```
{{% /example %}}
{{% /examples %}}