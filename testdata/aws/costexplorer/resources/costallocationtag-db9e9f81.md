Provides a CE Cost Allocation Tag.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.costexplorer.CostAllocationTag("example", {
    status: "Active",
    tagKey: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.costexplorer.CostAllocationTag("example",
    status="Active",
    tag_key="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CostExplorer.CostAllocationTag("example", new()
    {
        Status = "Active",
        TagKey = "example",
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
		_, err := costexplorer.NewCostAllocationTag(ctx, "example", &costexplorer.CostAllocationTagArgs{
			Status: pulumi.String("Active"),
			TagKey: pulumi.String("example"),
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
import com.pulumi.aws.costexplorer.CostAllocationTag;
import com.pulumi.aws.costexplorer.CostAllocationTagArgs;
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
        var example = new CostAllocationTag("example", CostAllocationTagArgs.builder()        
            .status("Active")
            .tagKey("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:costexplorer:CostAllocationTag
    properties:
      status: Active
      tagKey: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ce_cost_allocation_tag` using the `id`. For example:

```sh
 $ pulumi import aws:costexplorer/costAllocationTag:CostAllocationTag example key
```
 