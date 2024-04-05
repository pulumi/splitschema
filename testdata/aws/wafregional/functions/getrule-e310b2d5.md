`aws.wafregional.Rule` Retrieves a WAF Regional Rule Resource Id.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.wafregional.getRule({
    name: "tfWAFRegionalRule",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.wafregional.get_rule(name="tfWAFRegionalRule")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.WafRegional.GetRule.Invoke(new()
    {
        Name = "tfWAFRegionalRule",
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
		_, err := wafregional.LookupRule(ctx, &wafregional.LookupRuleArgs{
			Name: "tfWAFRegionalRule",
		}, nil)
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
import com.pulumi.aws.wafregional.WafregionalFunctions;
import com.pulumi.aws.wafregional.inputs.GetRuleArgs;
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
        final var example = WafregionalFunctions.getRule(GetRuleArgs.builder()
            .name("tfWAFRegionalRule")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:wafregional:getRule
      Arguments:
        name: tfWAFRegionalRule
```
{{% /example %}}
{{% /examples %}}