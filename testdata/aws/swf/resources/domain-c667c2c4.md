Provides an SWF Domain resource.

{{% examples %}}
## Example Usage
{{% example %}}

To register a basic SWF domain:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.swf.Domain("foo", {
    description: "SWF Domain",
    workflowExecutionRetentionPeriodInDays: "30",
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.swf.Domain("foo",
    description="SWF Domain",
    workflow_execution_retention_period_in_days="30")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Swf.Domain("foo", new()
    {
        Description = "SWF Domain",
        WorkflowExecutionRetentionPeriodInDays = "30",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/swf"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := swf.NewDomain(ctx, "foo", &swf.DomainArgs{
			Description:                            pulumi.String("SWF Domain"),
			WorkflowExecutionRetentionPeriodInDays: pulumi.String("30"),
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
import com.pulumi.aws.swf.Domain;
import com.pulumi.aws.swf.DomainArgs;
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
        var foo = new Domain("foo", DomainArgs.builder()        
            .description("SWF Domain")
            .workflowExecutionRetentionPeriodInDays(30)
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:swf:Domain
    properties:
      description: SWF Domain
      workflowExecutionRetentionPeriodInDays: 30
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SWF Domains using the `name`. For example:

```sh
 $ pulumi import aws:swf/domain:Domain foo test-domain
```
 