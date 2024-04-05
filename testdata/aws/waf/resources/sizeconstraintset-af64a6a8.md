Use the `aws.waf.SizeConstraintSet` resource to manage WAF size constraint sets.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sizeConstraintSet = new aws.waf.SizeConstraintSet("sizeConstraintSet", {sizeConstraints: [{
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

size_constraint_set = aws.waf.SizeConstraintSet("sizeConstraintSet", size_constraints=[aws.waf.SizeConstraintSetSizeConstraintArgs(
    comparison_operator="EQ",
    field_to_match=aws.waf.SizeConstraintSetSizeConstraintFieldToMatchArgs(
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
    var sizeConstraintSet = new Aws.Waf.SizeConstraintSet("sizeConstraintSet", new()
    {
        SizeConstraints = new[]
        {
            new Aws.Waf.Inputs.SizeConstraintSetSizeConstraintArgs
            {
                ComparisonOperator = "EQ",
                FieldToMatch = new Aws.Waf.Inputs.SizeConstraintSetSizeConstraintFieldToMatchArgs
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewSizeConstraintSet(ctx, "sizeConstraintSet", &waf.SizeConstraintSetArgs{
			SizeConstraints: waf.SizeConstraintSetSizeConstraintArray{
				&waf.SizeConstraintSetSizeConstraintArgs{
					ComparisonOperator: pulumi.String("EQ"),
					FieldToMatch: &waf.SizeConstraintSetSizeConstraintFieldToMatchArgs{
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
import com.pulumi.aws.waf.SizeConstraintSet;
import com.pulumi.aws.waf.SizeConstraintSetArgs;
import com.pulumi.aws.waf.inputs.SizeConstraintSetSizeConstraintArgs;
import com.pulumi.aws.waf.inputs.SizeConstraintSetSizeConstraintFieldToMatchArgs;
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
    type: aws:waf:SizeConstraintSet
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

Using `pulumi import`, import AWS WAF Size Constraint Set using their ID. For example:

```sh
 $ pulumi import aws:waf/sizeConstraintSet:SizeConstraintSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 