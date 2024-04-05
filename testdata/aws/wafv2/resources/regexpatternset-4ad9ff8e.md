Provides an AWS WAFv2 Regex Pattern Set Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafv2.RegexPatternSet("example", {
    description: "Example regex pattern set",
    regularExpressions: [
        {
            regexString: "one",
        },
        {
            regexString: "two",
        },
    ],
    scope: "REGIONAL",
    tags: {
        Tag1: "Value1",
        Tag2: "Value2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.wafv2.RegexPatternSet("example",
    description="Example regex pattern set",
    regular_expressions=[
        aws.wafv2.RegexPatternSetRegularExpressionArgs(
            regex_string="one",
        ),
        aws.wafv2.RegexPatternSetRegularExpressionArgs(
            regex_string="two",
        ),
    ],
    scope="REGIONAL",
    tags={
        "Tag1": "Value1",
        "Tag2": "Value2",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.WafV2.RegexPatternSet("example", new()
    {
        Description = "Example regex pattern set",
        RegularExpressions = new[]
        {
            new Aws.WafV2.Inputs.RegexPatternSetRegularExpressionArgs
            {
                RegexString = "one",
            },
            new Aws.WafV2.Inputs.RegexPatternSetRegularExpressionArgs
            {
                RegexString = "two",
            },
        },
        Scope = "REGIONAL",
        Tags = 
        {
            { "Tag1", "Value1" },
            { "Tag2", "Value2" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafv2.NewRegexPatternSet(ctx, "example", &wafv2.RegexPatternSetArgs{
			Description: pulumi.String("Example regex pattern set"),
			RegularExpressions: wafv2.RegexPatternSetRegularExpressionArray{
				&wafv2.RegexPatternSetRegularExpressionArgs{
					RegexString: pulumi.String("one"),
				},
				&wafv2.RegexPatternSetRegularExpressionArgs{
					RegexString: pulumi.String("two"),
				},
			},
			Scope: pulumi.String("REGIONAL"),
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
				"Tag2": pulumi.String("Value2"),
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
import com.pulumi.aws.wafv2.RegexPatternSet;
import com.pulumi.aws.wafv2.RegexPatternSetArgs;
import com.pulumi.aws.wafv2.inputs.RegexPatternSetRegularExpressionArgs;
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
        var example = new RegexPatternSet("example", RegexPatternSetArgs.builder()        
            .description("Example regex pattern set")
            .regularExpressions(            
                RegexPatternSetRegularExpressionArgs.builder()
                    .regexString("one")
                    .build(),
                RegexPatternSetRegularExpressionArgs.builder()
                    .regexString("two")
                    .build())
            .scope("REGIONAL")
            .tags(Map.ofEntries(
                Map.entry("Tag1", "Value1"),
                Map.entry("Tag2", "Value2")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:wafv2:RegexPatternSet
    properties:
      description: Example regex pattern set
      regularExpressions:
        - regexString: one
        - regexString: two
      scope: REGIONAL
      tags:
        Tag1: Value1
        Tag2: Value2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAFv2 Regex Pattern Sets using `ID/name/scope`. For example:

```sh
 $ pulumi import aws:wafv2/regexPatternSet:RegexPatternSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
```
 