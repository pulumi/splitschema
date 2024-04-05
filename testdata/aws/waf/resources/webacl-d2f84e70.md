Provides a WAF Web ACL Resource

{{% examples %}}
## Example Usage
{{% example %}}

This example blocks requests coming from `192.0.7.0/24` and allows everything else.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ipset = new aws.waf.IpSet("ipset", {ipSetDescriptors: [{
    type: "IPV4",
    value: "192.0.7.0/24",
}]});
const wafrule = new aws.waf.Rule("wafrule", {
    metricName: "tfWAFRule",
    predicates: [{
        dataId: ipset.id,
        negated: false,
        type: "IPMatch",
    }],
}, {
    dependsOn: [ipset],
});
const wafAcl = new aws.waf.WebAcl("wafAcl", {
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
}, {
    dependsOn: [
        ipset,
        wafrule,
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

ipset = aws.waf.IpSet("ipset", ip_set_descriptors=[aws.waf.IpSetIpSetDescriptorArgs(
    type="IPV4",
    value="192.0.7.0/24",
)])
wafrule = aws.waf.Rule("wafrule",
    metric_name="tfWAFRule",
    predicates=[aws.waf.RulePredicateArgs(
        data_id=ipset.id,
        negated=False,
        type="IPMatch",
    )],
    opts=pulumi.ResourceOptions(depends_on=[ipset]))
waf_acl = aws.waf.WebAcl("wafAcl",
    metric_name="tfWebACL",
    default_action=aws.waf.WebAclDefaultActionArgs(
        type="ALLOW",
    ),
    rules=[aws.waf.WebAclRuleArgs(
        action=aws.waf.WebAclRuleActionArgs(
            type="BLOCK",
        ),
        priority=1,
        rule_id=wafrule.id,
        type="REGULAR",
    )],
    opts=pulumi.ResourceOptions(depends_on=[
            ipset,
            wafrule,
        ]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ipset = new Aws.Waf.IpSet("ipset", new()
    {
        IpSetDescriptors = new[]
        {
            new Aws.Waf.Inputs.IpSetIpSetDescriptorArgs
            {
                Type = "IPV4",
                Value = "192.0.7.0/24",
            },
        },
    });

    var wafrule = new Aws.Waf.Rule("wafrule", new()
    {
        MetricName = "tfWAFRule",
        Predicates = new[]
        {
            new Aws.Waf.Inputs.RulePredicateArgs
            {
                DataId = ipset.Id,
                Negated = false,
                Type = "IPMatch",
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            ipset,
        },
    });

    var wafAcl = new Aws.Waf.WebAcl("wafAcl", new()
    {
        MetricName = "tfWebACL",
        DefaultAction = new Aws.Waf.Inputs.WebAclDefaultActionArgs
        {
            Type = "ALLOW",
        },
        Rules = new[]
        {
            new Aws.Waf.Inputs.WebAclRuleArgs
            {
                Action = new Aws.Waf.Inputs.WebAclRuleActionArgs
                {
                    Type = "BLOCK",
                },
                Priority = 1,
                RuleId = wafrule.Id,
                Type = "REGULAR",
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            ipset,
            wafrule,
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
		ipset, err := waf.NewIpSet(ctx, "ipset", &waf.IpSetArgs{
			IpSetDescriptors: waf.IpSetIpSetDescriptorArray{
				&waf.IpSetIpSetDescriptorArgs{
					Type:  pulumi.String("IPV4"),
					Value: pulumi.String("192.0.7.0/24"),
				},
			},
		})
		if err != nil {
			return err
		}
		wafrule, err := waf.NewRule(ctx, "wafrule", &waf.RuleArgs{
			MetricName: pulumi.String("tfWAFRule"),
			Predicates: waf.RulePredicateArray{
				&waf.RulePredicateArgs{
					DataId:  ipset.ID(),
					Negated: pulumi.Bool(false),
					Type:    pulumi.String("IPMatch"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			ipset,
		}))
		if err != nil {
			return err
		}
		_, err = waf.NewWebAcl(ctx, "wafAcl", &waf.WebAclArgs{
			MetricName: pulumi.String("tfWebACL"),
			DefaultAction: &waf.WebAclDefaultActionArgs{
				Type: pulumi.String("ALLOW"),
			},
			Rules: waf.WebAclRuleArray{
				&waf.WebAclRuleArgs{
					Action: &waf.WebAclRuleActionArgs{
						Type: pulumi.String("BLOCK"),
					},
					Priority: pulumi.Int(1),
					RuleId:   wafrule.ID(),
					Type:     pulumi.String("REGULAR"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			ipset,
			wafrule,
		}))
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
import com.pulumi.aws.waf.IpSet;
import com.pulumi.aws.waf.IpSetArgs;
import com.pulumi.aws.waf.inputs.IpSetIpSetDescriptorArgs;
import com.pulumi.aws.waf.Rule;
import com.pulumi.aws.waf.RuleArgs;
import com.pulumi.aws.waf.inputs.RulePredicateArgs;
import com.pulumi.aws.waf.WebAcl;
import com.pulumi.aws.waf.WebAclArgs;
import com.pulumi.aws.waf.inputs.WebAclDefaultActionArgs;
import com.pulumi.aws.waf.inputs.WebAclRuleArgs;
import com.pulumi.aws.waf.inputs.WebAclRuleActionArgs;
import com.pulumi.resources.CustomResourceOptions;
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
            .build(), CustomResourceOptions.builder()
                .dependsOn(ipset)
                .build());

        var wafAcl = new WebAcl("wafAcl", WebAclArgs.builder()        
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
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    ipset,
                    wafrule)
                .build());

    }
}
```
```yaml
resources:
  ipset:
    type: aws:waf:IpSet
    properties:
      ipSetDescriptors:
        - type: IPV4
          value: 192.0.7.0/24
  wafrule:
    type: aws:waf:Rule
    properties:
      metricName: tfWAFRule
      predicates:
        - dataId: ${ipset.id}
          negated: false
          type: IPMatch
    options:
      dependson:
        - ${ipset}
  wafAcl:
    type: aws:waf:WebAcl
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
    options:
      dependson:
        - ${ipset}
        - ${wafrule}
```
{{% /example %}}
{{% example %}}
### Logging

> *NOTE:* The Kinesis Firehose Delivery Stream name must begin with `aws-waf-logs-` and be located in `us-east-1` region. See the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) for more information about enabling WAF logging.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.waf.WebAcl("example", {loggingConfiguration: {
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

example = aws.waf.WebAcl("example", logging_configuration=aws.waf.WebAclLoggingConfigurationArgs(
    log_destination=aws_kinesis_firehose_delivery_stream["example"]["arn"],
    redacted_fields=aws.waf.WebAclLoggingConfigurationRedactedFieldsArgs(
        field_to_matches=[
            aws.waf.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs(
                type="URI",
            ),
            aws.waf.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs(
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
    var example = new Aws.Waf.WebAcl("example", new()
    {
        LoggingConfiguration = new Aws.Waf.Inputs.WebAclLoggingConfigurationArgs
        {
            LogDestination = aws_kinesis_firehose_delivery_stream.Example.Arn,
            RedactedFields = new Aws.Waf.Inputs.WebAclLoggingConfigurationRedactedFieldsArgs
            {
                FieldToMatches = new[]
                {
                    new Aws.Waf.Inputs.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs
                    {
                        Type = "URI",
                    },
                    new Aws.Waf.Inputs.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewWebAcl(ctx, "example", &waf.WebAclArgs{
			LoggingConfiguration: &waf.WebAclLoggingConfigurationArgs{
				LogDestination: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
				RedactedFields: &waf.WebAclLoggingConfigurationRedactedFieldsArgs{
					FieldToMatches: waf.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArray{
						&waf.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs{
							Type: pulumi.String("URI"),
						},
						&waf.WebAclLoggingConfigurationRedactedFieldsFieldToMatchArgs{
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
import com.pulumi.aws.waf.WebAcl;
import com.pulumi.aws.waf.WebAclArgs;
import com.pulumi.aws.waf.inputs.WebAclLoggingConfigurationArgs;
import com.pulumi.aws.waf.inputs.WebAclLoggingConfigurationRedactedFieldsArgs;
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
    type: aws:waf:WebAcl
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

Using `pulumi import`, import WAF Web ACL using the `id`. For example:

```sh
 $ pulumi import aws:waf/webAcl:WebAcl main 0c8e583e-18f3-4c13-9e2a-67c4805d2f94
```
 