Provides a WAF Regex Match Set Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRegexPatternSet = new aws.waf.RegexPatternSet("exampleRegexPatternSet", {regexPatternStrings: [
    "one",
    "two",
]});
const exampleRegexMatchSet = new aws.waf.RegexMatchSet("exampleRegexMatchSet", {regexMatchTuples: [{
    fieldToMatch: {
        data: "User-Agent",
        type: "HEADER",
    },
    regexPatternSetId: exampleRegexPatternSet.id,
    textTransformation: "NONE",
}]});
```
```python
import pulumi
import pulumi_aws as aws

example_regex_pattern_set = aws.waf.RegexPatternSet("exampleRegexPatternSet", regex_pattern_strings=[
    "one",
    "two",
])
example_regex_match_set = aws.waf.RegexMatchSet("exampleRegexMatchSet", regex_match_tuples=[aws.waf.RegexMatchSetRegexMatchTupleArgs(
    field_to_match=aws.waf.RegexMatchSetRegexMatchTupleFieldToMatchArgs(
        data="User-Agent",
        type="HEADER",
    ),
    regex_pattern_set_id=example_regex_pattern_set.id,
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
    var exampleRegexPatternSet = new Aws.Waf.RegexPatternSet("exampleRegexPatternSet", new()
    {
        RegexPatternStrings = new[]
        {
            "one",
            "two",
        },
    });

    var exampleRegexMatchSet = new Aws.Waf.RegexMatchSet("exampleRegexMatchSet", new()
    {
        RegexMatchTuples = new[]
        {
            new Aws.Waf.Inputs.RegexMatchSetRegexMatchTupleArgs
            {
                FieldToMatch = new Aws.Waf.Inputs.RegexMatchSetRegexMatchTupleFieldToMatchArgs
                {
                    Data = "User-Agent",
                    Type = "HEADER",
                },
                RegexPatternSetId = exampleRegexPatternSet.Id,
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
		exampleRegexPatternSet, err := waf.NewRegexPatternSet(ctx, "exampleRegexPatternSet", &waf.RegexPatternSetArgs{
			RegexPatternStrings: pulumi.StringArray{
				pulumi.String("one"),
				pulumi.String("two"),
			},
		})
		if err != nil {
			return err
		}
		_, err = waf.NewRegexMatchSet(ctx, "exampleRegexMatchSet", &waf.RegexMatchSetArgs{
			RegexMatchTuples: waf.RegexMatchSetRegexMatchTupleArray{
				&waf.RegexMatchSetRegexMatchTupleArgs{
					FieldToMatch: &waf.RegexMatchSetRegexMatchTupleFieldToMatchArgs{
						Data: pulumi.String("User-Agent"),
						Type: pulumi.String("HEADER"),
					},
					RegexPatternSetId:  exampleRegexPatternSet.ID(),
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
import com.pulumi.aws.waf.RegexPatternSet;
import com.pulumi.aws.waf.RegexPatternSetArgs;
import com.pulumi.aws.waf.RegexMatchSet;
import com.pulumi.aws.waf.RegexMatchSetArgs;
import com.pulumi.aws.waf.inputs.RegexMatchSetRegexMatchTupleArgs;
import com.pulumi.aws.waf.inputs.RegexMatchSetRegexMatchTupleFieldToMatchArgs;
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
        var exampleRegexPatternSet = new RegexPatternSet("exampleRegexPatternSet", RegexPatternSetArgs.builder()        
            .regexPatternStrings(            
                "one",
                "two")
            .build());

        var exampleRegexMatchSet = new RegexMatchSet("exampleRegexMatchSet", RegexMatchSetArgs.builder()        
            .regexMatchTuples(RegexMatchSetRegexMatchTupleArgs.builder()
                .fieldToMatch(RegexMatchSetRegexMatchTupleFieldToMatchArgs.builder()
                    .data("User-Agent")
                    .type("HEADER")
                    .build())
                .regexPatternSetId(exampleRegexPatternSet.id())
                .textTransformation("NONE")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleRegexMatchSet:
    type: aws:waf:RegexMatchSet
    properties:
      regexMatchTuples:
        - fieldToMatch:
            data: User-Agent
            type: HEADER
          regexPatternSetId: ${exampleRegexPatternSet.id}
          textTransformation: NONE
  exampleRegexPatternSet:
    type: aws:waf:RegexPatternSet
    properties:
      regexPatternStrings:
        - one
        - two
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAF Regex Match Set using their ID. For example:

```sh
 $ pulumi import aws:waf/regexMatchSet:RegexMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 