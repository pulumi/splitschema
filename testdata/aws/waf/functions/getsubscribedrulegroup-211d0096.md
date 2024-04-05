`aws.waf.getSubscribedRuleGroup` retrieves information about a Managed WAF Rule Group from AWS Marketplace (needs to be subscribed to first).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byName = aws.waf.getSubscribedRuleGroup({
    name: "F5 Bot Detection Signatures For AWS WAF",
});
const byMetricName = aws.waf.getSubscribedRuleGroup({
    metricName: "F5BotDetectionSignatures",
});
// ...
const acl = new aws.waf.WebAcl("acl", {rules: [
    {
        priority: 1,
        ruleId: byName.then(byName => byName.id),
        type: "GROUP",
    },
    {
        priority: 2,
        ruleId: byMetricName.then(byMetricName => byMetricName.id),
        type: "GROUP",
    },
]});
```
```python
import pulumi
import pulumi_aws as aws

by_name = aws.waf.get_subscribed_rule_group(name="F5 Bot Detection Signatures For AWS WAF")
by_metric_name = aws.waf.get_subscribed_rule_group(metric_name="F5BotDetectionSignatures")
# ...
acl = aws.waf.WebAcl("acl", rules=[
    aws.waf.WebAclRuleArgs(
        priority=1,
        rule_id=by_name.id,
        type="GROUP",
    ),
    aws.waf.WebAclRuleArgs(
        priority=2,
        rule_id=by_metric_name.id,
        type="GROUP",
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
    var byName = Aws.Waf.GetSubscribedRuleGroup.Invoke(new()
    {
        Name = "F5 Bot Detection Signatures For AWS WAF",
    });

    var byMetricName = Aws.Waf.GetSubscribedRuleGroup.Invoke(new()
    {
        MetricName = "F5BotDetectionSignatures",
    });

    // ...
    var acl = new Aws.Waf.WebAcl("acl", new()
    {
        Rules = new[]
        {
            new Aws.Waf.Inputs.WebAclRuleArgs
            {
                Priority = 1,
                RuleId = byName.Apply(getSubscribedRuleGroupResult => getSubscribedRuleGroupResult.Id),
                Type = "GROUP",
            },
            new Aws.Waf.Inputs.WebAclRuleArgs
            {
                Priority = 2,
                RuleId = byMetricName.Apply(getSubscribedRuleGroupResult => getSubscribedRuleGroupResult.Id),
                Type = "GROUP",
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
		byName, err := waf.GetSubscribedRuleGroup(ctx, &waf.GetSubscribedRuleGroupArgs{
			Name: pulumi.StringRef("F5 Bot Detection Signatures For AWS WAF"),
		}, nil)
		if err != nil {
			return err
		}
		byMetricName, err := waf.GetSubscribedRuleGroup(ctx, &waf.GetSubscribedRuleGroupArgs{
			MetricName: pulumi.StringRef("F5BotDetectionSignatures"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = waf.NewWebAcl(ctx, "acl", &waf.WebAclArgs{
			Rules: waf.WebAclRuleArray{
				&waf.WebAclRuleArgs{
					Priority: pulumi.Int(1),
					RuleId:   *pulumi.String(byName.Id),
					Type:     pulumi.String("GROUP"),
				},
				&waf.WebAclRuleArgs{
					Priority: pulumi.Int(2),
					RuleId:   *pulumi.String(byMetricName.Id),
					Type:     pulumi.String("GROUP"),
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
import com.pulumi.aws.waf.WafFunctions;
import com.pulumi.aws.waf.inputs.GetSubscribedRuleGroupArgs;
import com.pulumi.aws.waf.WebAcl;
import com.pulumi.aws.waf.WebAclArgs;
import com.pulumi.aws.waf.inputs.WebAclRuleArgs;
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
        final var byName = WafFunctions.getSubscribedRuleGroup(GetSubscribedRuleGroupArgs.builder()
            .name("F5 Bot Detection Signatures For AWS WAF")
            .build());

        final var byMetricName = WafFunctions.getSubscribedRuleGroup(GetSubscribedRuleGroupArgs.builder()
            .metricName("F5BotDetectionSignatures")
            .build());

        var acl = new WebAcl("acl", WebAclArgs.builder()        
            .rules(            
                WebAclRuleArgs.builder()
                    .priority(1)
                    .ruleId(byName.applyValue(getSubscribedRuleGroupResult -> getSubscribedRuleGroupResult.id()))
                    .type("GROUP")
                    .build(),
                WebAclRuleArgs.builder()
                    .priority(2)
                    .ruleId(byMetricName.applyValue(getSubscribedRuleGroupResult -> getSubscribedRuleGroupResult.id()))
                    .type("GROUP")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  acl:
    type: aws:waf:WebAcl
    properties:
      rules:
        - priority: 1
          ruleId: ${byName.id}
          type: GROUP
        - priority: 2
          ruleId: ${byMetricName.id}
          type: GROUP
variables:
  byName:
    fn::invoke:
      Function: aws:waf:getSubscribedRuleGroup
      Arguments:
        name: F5 Bot Detection Signatures For AWS WAF
  byMetricName:
    fn::invoke:
      Function: aws:waf:getSubscribedRuleGroup
      Arguments:
        metricName: F5BotDetectionSignatures
```
{{% /example %}}
{{% /examples %}}