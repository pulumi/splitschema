Provides an SES receipt filter resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const filter = new aws.ses.ReceiptFilter("filter", {
    cidr: "10.10.10.10",
    policy: "Block",
});
```
```python
import pulumi
import pulumi_aws as aws

filter = aws.ses.ReceiptFilter("filter",
    cidr="10.10.10.10",
    policy="Block")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var filter = new Aws.Ses.ReceiptFilter("filter", new()
    {
        Cidr = "10.10.10.10",
        Policy = "Block",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewReceiptFilter(ctx, "filter", &ses.ReceiptFilterArgs{
			Cidr:   pulumi.String("10.10.10.10"),
			Policy: pulumi.String("Block"),
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
import com.pulumi.aws.ses.ReceiptFilter;
import com.pulumi.aws.ses.ReceiptFilterArgs;
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
        var filter = new ReceiptFilter("filter", ReceiptFilterArgs.builder()        
            .cidr("10.10.10.10")
            .policy("Block")
            .build());

    }
}
```
```yaml
resources:
  filter:
    type: aws:ses:ReceiptFilter
    properties:
      cidr: 10.10.10.10
      policy: Block
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SES Receipt Filter using their `name`. For example:

```sh
 $ pulumi import aws:ses/receiptFilter:ReceiptFilter test some-filter
```
 