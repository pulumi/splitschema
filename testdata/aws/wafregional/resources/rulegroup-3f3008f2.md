Provides a WAF Regional Rule Group Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRule = new aws.wafregional.Rule("exampleRule", {metricName: "example"});
const exampleRuleGroup = new aws.wafregional.RuleGroup("exampleRuleGroup", {
    metricName: "example",
    activatedRules: [{
        action: {
            type: "COUNT",
        },
        priority: 50,
        ruleId: exampleRule.id,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_rule = aws.wafregional.Rule("exampleRule", metric_name="example")
example_rule_group = aws.wafregional.RuleGroup("exampleRuleGroup",
    metric_name="example",
    activated_rules=[aws.wafregional.RuleGroupActivatedRuleArgs(
        action=aws.wafregional.RuleGroupActivatedRuleActionArgs(
            type="COUNT",
        ),
        priority=50,
        rule_id=example_rule.id,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleRule = new Aws.WafRegional.Rule("exampleRule", new()
    {
        MetricName = "example",
    });

    var exampleRuleGroup = new Aws.WafRegional.RuleGroup("exampleRuleGroup", new()
    {
        MetricName = "example",
        ActivatedRules = new[]
        {
            new Aws.WafRegional.Inputs.RuleGroupActivatedRuleArgs
            {
                Action = new Aws.WafRegional.Inputs.RuleGroupActivatedRuleActionArgs
                {
                    Type = "COUNT",
                },
                Priority = 50,
                RuleId = exampleRule.Id,
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
		exampleRule, err := wafregional.NewRule(ctx, "exampleRule", &wafregional.RuleArgs{
			MetricName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = wafregional.NewRuleGroup(ctx, "exampleRuleGroup", &wafregional.RuleGroupArgs{
			MetricName: pulumi.String("example"),
			ActivatedRules: wafregional.RuleGroupActivatedRuleArray{
				&wafregional.RuleGroupActivatedRuleArgs{
					Action: &wafregional.RuleGroupActivatedRuleActionArgs{
						Type: pulumi.String("COUNT"),
					},
					Priority: pulumi.Int(50),
					RuleId:   exampleRule.ID(),
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
import com.pulumi.aws.wafregional.Rule;
import com.pulumi.aws.wafregional.RuleArgs;
import com.pulumi.aws.wafregional.RuleGroup;
import com.pulumi.aws.wafregional.RuleGroupArgs;
import com.pulumi.aws.wafregional.inputs.RuleGroupActivatedRuleArgs;
import com.pulumi.aws.wafregional.inputs.RuleGroupActivatedRuleActionArgs;
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
        var exampleRule = new Rule("exampleRule", RuleArgs.builder()        
            .metricName("example")
            .build());

        var exampleRuleGroup = new RuleGroup("exampleRuleGroup", RuleGroupArgs.builder()        
            .metricName("example")
            .activatedRules(RuleGroupActivatedRuleArgs.builder()
                .action(RuleGroupActivatedRuleActionArgs.builder()
                    .type("COUNT")
                    .build())
                .priority(50)
                .ruleId(exampleRule.id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleRule:
    type: aws:wafregional:Rule
    properties:
      metricName: example
  exampleRuleGroup:
    type: aws:wafregional:RuleGroup
    properties:
      metricName: example
      activatedRules:
        - action:
            type: COUNT
          priority: 50
          ruleId: ${exampleRule.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAF Regional Rule Group using the id. For example:

```sh
 $ pulumi import aws:wafregional/ruleGroup:RuleGroup example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 