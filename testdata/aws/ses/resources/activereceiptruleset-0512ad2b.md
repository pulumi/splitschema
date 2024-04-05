Provides a resource to designate the active SES receipt rule set

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ses.ActiveReceiptRuleSet("main", {ruleSetName: "primary-rules"});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ses.ActiveReceiptRuleSet("main", rule_set_name="primary-rules")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ses.ActiveReceiptRuleSet("main", new()
    {
        RuleSetName = "primary-rules",
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
		_, err := ses.NewActiveReceiptRuleSet(ctx, "main", &ses.ActiveReceiptRuleSetArgs{
			RuleSetName: pulumi.String("primary-rules"),
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
import com.pulumi.aws.ses.ActiveReceiptRuleSet;
import com.pulumi.aws.ses.ActiveReceiptRuleSetArgs;
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
        var main = new ActiveReceiptRuleSet("main", ActiveReceiptRuleSetArgs.builder()        
            .ruleSetName("primary-rules")
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:ses:ActiveReceiptRuleSet
    properties:
      ruleSetName: primary-rules
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import active SES receipt rule sets using the rule set name. For example:

```sh
 $ pulumi import aws:ses/activeReceiptRuleSet:ActiveReceiptRuleSet my_rule_set my_rule_set_name
```
 