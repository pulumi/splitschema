Provides a WAF Regional XSS Match Set Resource for use with Application Load Balancer.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const xssMatchSet = new aws.wafregional.XssMatchSet("xssMatchSet", {xssMatchTuples: [
    {
        fieldToMatch: {
            type: "URI",
        },
        textTransformation: "NONE",
    },
    {
        fieldToMatch: {
            type: "QUERY_STRING",
        },
        textTransformation: "NONE",
    },
]});
```
```python
import pulumi
import pulumi_aws as aws

xss_match_set = aws.wafregional.XssMatchSet("xssMatchSet", xss_match_tuples=[
    aws.wafregional.XssMatchSetXssMatchTupleArgs(
        field_to_match=aws.wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs(
            type="URI",
        ),
        text_transformation="NONE",
    ),
    aws.wafregional.XssMatchSetXssMatchTupleArgs(
        field_to_match=aws.wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs(
            type="QUERY_STRING",
        ),
        text_transformation="NONE",
    ),
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var xssMatchSet = new Aws.WafRegional.XssMatchSet("xssMatchSet", new()
    {
        XssMatchTuples = new[]
        {
            new Aws.WafRegional.Inputs.XssMatchSetXssMatchTupleArgs
            {
                FieldToMatch = new Aws.WafRegional.Inputs.XssMatchSetXssMatchTupleFieldToMatchArgs
                {
                    Type = "URI",
                },
                TextTransformation = "NONE",
            },
            new Aws.WafRegional.Inputs.XssMatchSetXssMatchTupleArgs
            {
                FieldToMatch = new Aws.WafRegional.Inputs.XssMatchSetXssMatchTupleFieldToMatchArgs
                {
                    Type = "QUERY_STRING",
                },
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
		_, err := wafregional.NewXssMatchSet(ctx, "xssMatchSet", &wafregional.XssMatchSetArgs{
			XssMatchTuples: wafregional.XssMatchSetXssMatchTupleArray{
				&wafregional.XssMatchSetXssMatchTupleArgs{
					FieldToMatch: &wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs{
						Type: pulumi.String("URI"),
					},
					TextTransformation: pulumi.String("NONE"),
				},
				&wafregional.XssMatchSetXssMatchTupleArgs{
					FieldToMatch: &wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs{
						Type: pulumi.String("QUERY_STRING"),
					},
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
import com.pulumi.aws.wafregional.XssMatchSet;
import com.pulumi.aws.wafregional.XssMatchSetArgs;
import com.pulumi.aws.wafregional.inputs.XssMatchSetXssMatchTupleArgs;
import com.pulumi.aws.wafregional.inputs.XssMatchSetXssMatchTupleFieldToMatchArgs;
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
        var xssMatchSet = new XssMatchSet("xssMatchSet", XssMatchSetArgs.builder()        
            .xssMatchTuples(            
                XssMatchSetXssMatchTupleArgs.builder()
                    .fieldToMatch(XssMatchSetXssMatchTupleFieldToMatchArgs.builder()
                        .type("URI")
                        .build())
                    .textTransformation("NONE")
                    .build(),
                XssMatchSetXssMatchTupleArgs.builder()
                    .fieldToMatch(XssMatchSetXssMatchTupleFieldToMatchArgs.builder()
                        .type("QUERY_STRING")
                        .build())
                    .textTransformation("NONE")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  xssMatchSet:
    type: aws:wafregional:XssMatchSet
    properties:
      xssMatchTuples:
        - fieldToMatch:
            type: URI
          textTransformation: NONE
        - fieldToMatch:
            type: QUERY_STRING
          textTransformation: NONE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS WAF Regional XSS Match using the `id`. For example:

```sh
 $ pulumi import aws:wafregional/xssMatchSet:XssMatchSet example 12345abcde
```
 