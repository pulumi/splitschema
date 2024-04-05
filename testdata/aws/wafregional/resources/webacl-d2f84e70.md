Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.

{{% examples %}}
## Example Usage
{{% example %}}
### Regular Rule

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ipset = new aws.wafregional.IpSet("ipset", {ipSetDescriptors: [{
    type: "IPV4",
    value: "192.0.7.0/24",
}]});
const wafrule = new aws.wafregional.Rule("wafrule", {
    metricName: "tfWAFRule",
    predicates: [{
        dataId: ipset.id,
        negated: false,
        type: "IPMatch",
    }],
});
const wafacl = new aws.wafregional.WebAcl("wafacl", {
    metricName: "tfWebACL",
    defaultAction: {
        type: "ALLOW",
    },
    rules: [{
        action: {
            type: "BLOCK",
        },
        priority: 1,
        ruleId: wafrule.id,
        type: "REGULAR",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[aws.wafregional.IpSetIpSetDescriptorArgs(
    type="IPV4",
    value="192.0.7.0/24",
)])
wafrule = aws.wafregional.Rule("wafrule",
    metric_name="tfWAFRule",
    predicates=[aws.wafregional.RulePredicateArgs(
        data_id=ipset.id,
        negated=False,
        type="IPMatch",
    )])
wafacl = aws.wafregional.WebAcl("wafacl",
    metric_name="tfWebACL",
    default_action=aws.wafregional.WebAclDefaultActionArgs(
        type="ALLOW",
    ),
    rules=[aws.wafregional.WebAclRuleArgs(
        action=aws.wafregional.WebAclRuleActionArgs(
            type="BLOCK",
        ),
        priority=1,
        rule_id=wafrule.id,
        type="REGULAR",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ipset = new Aws.WafRegional.IpSet("ipset", new()
    {
        IpSetDescriptors = new[]
        {
            new Aws.WafRegional.Inputs.IpSetIpSetDescriptorArgs
            {
                Type = "IPV4",
                Value = "192.0.7.0/24",
            },
        },
    });

    var wafrule = new Aws.WafRegional.Rule("wafrule", new()
    {
        MetricName = "tfWAFRule",
        Predicates = new[]
        {
            new Aws.WafRegional.Inputs.RulePredicateArgs
            {
                DataId = ipset.Id,
                Negated = false,
                Type = "IPMatch",
            },
        },
    });

    var wafacl = new Aws.WafRegional.WebAcl("wafacl", new()
    {
        MetricName = "tfWebACL",
        DefaultAction = new Aws.WafRegional.Inputs.WebAclDefaultActionArgs
        {
            Type = "ALLOW",
        },
        Rules = new[]
        {
            new Aws.WafRegional.Inputs.WebAclRuleArgs
            {
                Action = new Aws.WafRegional.Inputs.WebAclRuleActionArgs
                {
                    Type = "BLOCK",
                },
                Priority = 1,
                RuleId = wafrule.Id,
                Type = "REGULAR",
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
		ipset, err := wafregional.NewIpSet(ctx, "ipset", &wafregional.IpSetArgs{
			IpSetDescriptors: wafregional.IpSetIpSetDescriptorArray{
				&wafregional.IpSetIpSetDescriptorArgs{
					Type:  pulumi.String("IPV4"),
					Value: pulumi.String("192.0.7.0/24"),
				},
			},
		})
		if err != nil {
			return err
		}
		wafrule, err := wafregional.NewRule(ctx, "wafrule", &wafregional.RuleArgs{
			MetricName: pulumi.String("tfWAFRule"),
			Predicates: wafregional.RulePredicateArray{
				&wafregional.RulePredicateArgs{
					DataId:  ipset.ID(),
					Negated: pulumi.Bool(false),
					Type:    pulumi.String("IPMatch"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = wafregional.NewWebAcl(ctx, "wafacl", &wafregional.WebAclArgs{
			MetricName: pulumi.String("tfWebACL"),
			DefaultAction: &wafregional.WebAclDefaultActionArgs{
				Type: pulumi.String("ALLOW"),
			},
			Rules: wafregional.WebAclRuleArray{
				&wafregional.WebAclRuleArgs{
					Action: &wafregional.WebAclRuleActionArgs{
						Type: pulumi.String("BLOCK"),
					},
					Priority: pulumi.Int(1),
					RuleId:   wafrule.ID(),
					Type:     pulumi.String("REGULAR"),
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
import com.pulumi.aws.wafregional.IpSet;
import com.pulumi.aws.wafregional.IpSetArgs;
import com.pulumi.aws.wafregional.inputs.IpSetIpSetDescriptorArgs;
import com.pulumi.aws.wafregional.Rule;
import com.pulumi.aws.wafregional.RuleArgs;
import com.pulumi.aws.wafregional.inputs.RulePredicateArgs;
import com.pulumi.aws.wafregional.WebAcl;
import com.pulumi.aws.wafregional.WebAclArgs;
import com.pulumi.aws.wafregional.inputs.WebAclDefaultActionArgs;
import com.pulumi.aws.wafregional.inputs.WebAclRuleArgs;
import com.pulumi.aws.wafregional.inputs.WebAclRuleActionArgs;
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
        var ipset = new IpSet("ipset", IpSetArgs.builder()        
            .ipSetDescriptors(IpSetIpSetDescriptorArgs.builder()
                .type("IPV4")
                .value("192.0.7.0/24")
                .build())
            .build());

        var wafrule = new Rule("wafrule", RuleArgs.builder()        
            .metricName("tfWAFRule")
            .predicates(RulePredicateArgs.builder()
                .dataId(ipset.id())
                .negated(false)
                .type("IPMatch")
                .build())
            .build());

        var wafacl = new WebAcl("wafacl", WebAclArgs.builder()        
            .metricName("tfWebACL")
            .defaultAction(WebAclDefaultActionArgs.builder()
                .type("ALLOW")
                .build())
            .rules(WebAclRuleArgs.builder()
                .action(WebAclRuleActionArgs.builder()
                    .type("BLOCK")
                    .build())
                .priority(1)
                .ruleId(wafrule.id())
                .type("REGULAR")
                .build())
            .build());

    }
}
```
```yaml
resources:
  ipset:
    type: aws:wafregional:IpSet
    properties:
      ipSetDescriptors:
        - type: IPV4
          value: 192.0.7.0/24
  wafrule:
    type: aws:wafregional:Rule
    properties:
      metricName: tfWAFRule
      predicates:
        - dataId: ${ipset.id}
          negated: false
          type: IPMatch
  wafacl:
    type: aws:wafregional:WebAcl
    properties:
      metricName: tfWebACL
      defaultAction:
        type: ALLOW
      rules:
        - action:
            type: BLOCK
          priority: 1
          ruleId: ${wafrule.id}
          type: REGULAR
```
{{% /example %}}
{{% example %}}
### Group Rule

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafregional.WebAcl("example", {
    metricName: "example",
    defaultAction: {
        type: "ALLOW",
    },
    rules: [{
        priority: 1,
        ruleId: aws_wafregional_rule_group.example.id,
        type: "GROUP",
        overrideAction: {
            type: "NONE",
        },
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.wafregional.WebAcl("example",
    metric_name="example",
    default_action=aws.wafregional.WebAclDefaultActionArgs(
        type="ALLOW",
    ),
    rules=[aws.wafregional.WebAclRuleArgs(
        priority=1,
        rule_id=aws_wafregional_rule_group["example"]["id"],
        type="GROUP",
        override_action=aws.wafregional.WebAclRuleOverrideActionArgs(
            type="NONE",
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.WafRegional.WebAcl("example", new()
    {
        MetricName = "example",
        DefaultAction = new Aws.WafRegional.Inputs.WebAclDefaultActionArgs
        {
            Type = "ALLOW",
        },
        Rules = new[]
        {
            new Aws.WafRegional.Inputs.WebAclRuleArgs
            {
                Priority = 1,
                RuleId = aws_wafregional_rule_group.Example.Id,
                Type = "GROUP",
                OverrideAction = new Aws.WafRegional.Inputs.WebAclRuleOverrideActionArgs
                {
                    Type = "NONE",
                },
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
		_, err := wafregional.NewWebAcl(ctx, "example", &wafregional.WebAclArgs{
			MetricName: pulumi.String("example"),
			DefaultAction: &wafregional.WebAclDefaultActionArgs{
				Type: pulumi.String("ALLOW"),
			},
			Rules: wafregional.WebAclRuleArray{
				&wafregional.WebAclRuleArgs{
					Priority: pulumi.Int(1),
					RuleId:   pulumi.Any(aws_wafregional_rule_group.Example.Id),
					Type:     pulumi.String("GROUP"),
					OverrideAction: &wafregional.WebAclRuleOverrideActionArgs{
						Type: pulumi.String("NONE"),
					},
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
import com.pulumi.aws.wafregional.WebAcl;
import com.pulumi.aws.wafregional.WebAclArgs;
import com.pulumi.aws.wafregional.inputs.WebAclDefaultActionArgs;
import com.pulumi.aws.wafregional.inputs.WebAclRuleArgs;
import com.pulumi.aws.wafregional.inputs.WebAclRuleOverrideActionArgs;
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
        var example = new WebAcl("example", WebAclArgs.builder()        
            .metricName("example")
            .defaultAction(WebAclDefaultActionArgs.builder()
                .type("ALLOW")
                .build())
            .rules(WebAclRuleArgs.builder()
                .priority(1)
                .ruleId(aws_wafregional_rule_group.example().id())
                .type("GROUP")
                .overrideAction(WebAclRuleOverrideActionArgs.builder()
                    .type("NONE")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:wafregional:WebAcl
    properties:
      metricName: example
      defaultAction:
        type: ALLOW
      rules:
        - priority: 1
          ruleId: ${aws_wafregional_rule_group.example.id}
          type: GROUP
          overrideAction:
            type: NONE
```
{{% /example %}}
{{% example %}}
### Logging

> *NOTE:* The Kinesis Firehose Delivery Stream name must begin with `aws-waf-logs-`. See the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) for more information about enabling WAF logging.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ... other configuration ...
const example = new aws.wafregional.WebAcl("example", {loggingConfiguration: {
    logDestination: aws_kinesis_firehose_delivery_stream.example.arn,
    redactedFields: {
        fieldToMatches: [
            {
                type: "URI",
            },
            {
                data: "referer",
                type: "HEADER",
            },
        ],
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configuration ...
example = aws.wafregional.WebAcl("example", logging_configuration=aws.wafregional.WebAclLoggingConfigurationArgs(
    log_destination=aws_kinesis_firehose_delivery_stream["example"]["arn"],
    redacted_fields=aws.wafregional.WebAclLoggingConfigurationRedactedFieldsArgs(
        field_to_matches=[
            aws.wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs(
                type="URI",
            ),
            aws.wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs(
                data="referer",
                type="HEADER",
            ),
        ],
    ),
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ... other configuration ...
    var example = new Aws.WafRegional.WebAcl("example", new()
    {
        LoggingConfiguration = new Aws.WafRegional.Inputs.WebAclLoggingConfigurationArgs
        {
            LogDestination = aws_kinesis_firehose_delivery_stream.Example.Arn,
            RedactedFields = new Aws.WafRegional.Inputs.WebAclLoggingConfigurationRedactedFieldsArgs
            {
                FieldToMatches = new[]
                {
                    new Aws.WafRegional.Inputs.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs
                    {
                        Type = "URI",
                    },
                    new Aws.WafRegional.Inputs.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs
                    {
                        Data = "referer",
                        Type = "HEADER",
                    },
                },
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
		_, err := wafregional.NewWebAcl(ctx, "example", &wafregional.WebAclArgs{
			LoggingConfiguration: &wafregional.WebAclLoggingConfigurationArgs{
				LogDestination: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
				RedactedFields: &wafregional.WebAclLoggingConfigurationRedactedFieldsArgs{
					FieldToMatches: wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArray{
						&wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs{
							Type: pulumi.String("URI"),
						},
						&wafregional.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs{
							Data: pulumi.String("referer"),
							Type: pulumi.String("HEADER"),
						},
					},
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
import com.pulumi.aws.wafregional.WebAcl;
import com.pulumi.aws.wafregional.WebAclArgs;
import com.pulumi.aws.wafregional.inputs.WebAclLoggingConfigurationArgs;
import com.pulumi.aws.wafregional.inputs.WebAclLoggingConfigurationRedactedFieldsArgs;
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
        var example = new WebAcl("example", WebAclArgs.builder()        
            .loggingConfiguration(WebAclLoggingConfigurationArgs.builder()
                .logDestination(aws_kinesis_firehose_delivery_stream.example().arn())
                .redactedFields(WebAclLoggingConfigurationRedactedFieldsArgs.builder()
                    .fieldToMatches(                    
                        WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs.builder()
                            .type("URI")
                            .build(),
                        WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs.builder()
                            .data("referer")
                            .type("HEADER")
                            .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:wafregional:WebAcl
    properties:
      loggingConfiguration:
        logDestination: ${aws_kinesis_firehose_delivery_stream.example.arn}
        redactedFields:
          fieldToMatches:
            - type: URI
            - data: referer
              type: HEADER
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAF Regional Web ACL using the id. For example:

```sh
 $ pulumi import aws:wafregional/webAcl:WebAcl wafacl a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 