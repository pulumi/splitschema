Provides a WAF Rule Group Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRule = new aws.waf.Rule("exampleRule", {metricName: "example"});
const exampleRuleGroup = new aws.waf.RuleGroup("exampleRuleGroup", {
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

example_rule = aws.waf.Rule("exampleRule", metric_name="example")
example_rule_group = aws.waf.RuleGroup("exampleRuleGroup",
    metric_name="example",
    activated_rules=[aws.waf.RuleGroupActivatedRuleArgs(
        action=aws.waf.RuleGroupActivatedRuleActionArgs(
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
    var exampleRule = new Aws.Waf.Rule("exampleRule", new()
    {
        MetricName = "example",
    });

    var exampleRuleGroup = new Aws.Waf.RuleGroup("exampleRuleGroup", new()
    {
        MetricName = "example",
        ActivatedRules = new[]
        {
            new Aws.Waf.Inputs.RuleGroupActivatedRuleArgs
            {
                Action = new Aws.Waf.Inputs.RuleGroupActivatedRuleActionArgs
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRule, err := waf.NewRule(ctx, "exampleRule", &waf.RuleArgs{
			MetricName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = waf.NewRuleGroup(ctx, "exampleRuleGroup", &waf.RuleGroupArgs{
			MetricName: pulumi.String("example"),
			ActivatedRules: waf.RuleGroupActivatedRuleArray{
				&waf.RuleGroupActivatedRuleArgs{
					Action: &waf.RuleGroupActivatedRuleActionArgs{
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
import com.pulumi.aws.waf.Rule;
import com.pulumi.aws.waf.RuleArgs;
import com.pulumi.aws.waf.RuleGroup;
import com.pulumi.aws.waf.RuleGroupArgs;
import com.pulumi.aws.waf.inputs.RuleGroupActivatedRuleArgs;
import com.pulumi.aws.waf.inputs.RuleGroupActivatedRuleActionArgs;
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
    type: aws:waf:Rule
    properties:
      metricName: example
  exampleRuleGroup:
    type: aws:waf:RuleGroup
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

Using `pulumi import`, import WAF Rule Group using the id. For example:

```sh
 $ pulumi import aws:waf/ruleGroup:RuleGroup example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 