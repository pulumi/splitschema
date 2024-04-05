Provides an AWS Route 53 Recovery Readiness Cell.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53recoveryreadiness.Cell("example", {cellName: "us-west-2-failover-cell"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53recoveryreadiness.Cell("example", cell_name="us-west-2-failover-cell")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53RecoveryReadiness.Cell("example", new()
    {
        CellName = "us-west-2-failover-cell",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53recoveryreadiness"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53recoveryreadiness.NewCell(ctx, "example", &route53recoveryreadiness.CellArgs{
			CellName: pulumi.String("us-west-2-failover-cell"),
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
import com.pulumi.aws.route53recoveryreadiness.Cell;
import com.pulumi.aws.route53recoveryreadiness.CellArgs;
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
        var example = new Cell("example", CellArgs.builder()        
            .cellName("us-west-2-failover-cell")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53recoveryreadiness:Cell
    properties:
      cellName: us-west-2-failover-cell
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 Recovery Readiness cells using the cell name. For example:

```sh
 $ pulumi import aws:route53recoveryreadiness/cell:Cell us-west-2-failover-cell us-west-2-failover-cell
```
 