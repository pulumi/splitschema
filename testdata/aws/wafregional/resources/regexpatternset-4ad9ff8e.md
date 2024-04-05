Provides a WAF Regional Regex Pattern Set Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafregional.RegexPatternSet("example", {regexPatternStrings: [
    "one",
    "two",
]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.wafregional.RegexPatternSet("example", regex_pattern_strings=[
    "one",
    "two",
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.WafRegional.RegexPatternSet("example", new()
    {
        RegexPatternStrings = new[]
        {
            "one",
            "two",
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
		_, err := wafregional.NewRegexPatternSet(ctx, "example", &wafregional.RegexPatternSetArgs{
			RegexPatternStrings: pulumi.StringArray{
				pulumi.String("one"),
				pulumi.String("two"),
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
import com.pulumi.aws.wafregional.RegexPatternSet;
import com.pulumi.aws.wafregional.RegexPatternSetArgs;
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
            .regexPatternStrings(            
                "one",
                "two")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:wafregional:RegexPatternSet
    properties:
      regexPatternStrings:
        - one
        - two
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAF Regional Regex Pattern Set using the id. For example:

```sh
 $ pulumi import aws:wafregional/regexPatternSet:RegexPatternSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 