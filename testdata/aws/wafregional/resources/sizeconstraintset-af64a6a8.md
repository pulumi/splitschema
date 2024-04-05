Provides a WAF Regional Size Constraint Set Resource for use with Application Load Balancer.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sizeConstraintSet = new aws.wafregional.SizeConstraintSet("sizeConstraintSet", {sizeConstraints: [{
    comparisonOperator: "EQ",
    fieldToMatch: {
        type: "BODY",
    },
    size: 4096,
    textTransformation: "NONE",
}]});
```
```python
import pulumi
import pulumi_aws as aws

size_constraint_set = aws.wafregional.SizeConstraintSet("sizeConstraintSet", size_constraints=[aws.wafregional.SizeConstraintSetSizeConstraintArgs(
    comparison_operator="EQ",
    field_to_match=aws.wafregional.SizeConstraintSetSizeConstraintFieldToMatchArgs(
        type="BODY",
    ),
    size=4096,
    text_transformation="NONE",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var sizeConstraintSet = new Aws.WafRegional.SizeConstraintSet("sizeConstraintSet", new()
    {
        SizeConstraints = new[]
        {
            new Aws.WafRegional.Inputs.SizeConstraintSetSizeConstraintArgs
            {
                ComparisonOperator = "EQ",
                FieldToMatch = new Aws.WafRegional.Inputs.SizeConstraintSetSizeConstraintFieldToMatchArgs
                {
                    Type = "BODY",
                },
                Size = 4096,
                TextTransformation = "NONE",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.NewSizeConstraintSet(ctx, "sizeConstraintSet", &wafregional.SizeConstraintSetArgs{
			SizeConstraints: wafregional.SizeConstraintSetSizeConstraintArray{
				&wafregional.SizeConstraintSetSizeConstraintArgs{
					ComparisonOperator: pulumi.String("EQ"),
					FieldToMatch: &wafregional.SizeConstraintSetSizeConstraintFieldToMatchArgs{
						Type: pulumi.String("BODY"),
					},
					Size:               pulumi.Int(4096),
					TextTransformation: pulumi.String("NONE"),
				},
			},
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
import com.pulumi.aws.wafregional.SizeConstraintSet;
import com.pulumi.aws.wafregional.SizeConstraintSetArgs;
import com.pulumi.aws.wafregional.inputs.SizeConstraintSetSizeConstraintArgs;
import com.pulumi.aws.wafregional.inputs.SizeConstraintSetSizeConstraintFieldToMatchArgs;
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
        var sizeConstraintSet = new SizeConstraintSet("sizeConstraintSet", SizeConstraintSetArgs.builder()        
            .sizeConstraints(SizeConstraintSetSizeConstraintArgs.builder()
                .comparisonOperator("EQ")
                .fieldToMatch(SizeConstraintSetSizeConstraintFieldToMatchArgs.builder()
                    .type("BODY")
                    .build())
                .size("4096")
                .textTransformation("NONE")
                .build())
            .build());

    }
}
```
```yaml
resources:
  sizeConstraintSet:
    type: aws:wafregional:SizeConstraintSet
    properties:
      sizeConstraints:
        - comparisonOperator: EQ
          fieldToMatch:
            type: BODY
          size: '4096'
          textTransformation: NONE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAF Size Constraint Set using the id. For example:

```sh
 $ pulumi import aws:wafregional/sizeConstraintSet:SizeConstraintSet size_constraint_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 