Provides an WAF Regional Rule Resource for use with Application Load Balancer.

{{% examples %}}
## Example Usage
{{% example %}}

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
        type: "IPMatch",
        dataId: ipset.id,
        negated: false,
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
        type="IPMatch",
        data_id=ipset.id,
        negated=False,
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
                Type = "IPMatch",
                DataId = ipset.Id,
                Negated = false,
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
		_, err = wafregional.NewRule(ctx, "wafrule", &wafregional.RuleArgs{
			MetricName: pulumi.String("tfWAFRule"),
			Predicates: wafregional.RulePredicateArray{
				&wafregional.RulePredicateArgs{
					Type:    pulumi.String("IPMatch"),
					DataId:  ipset.ID(),
					Negated: pulumi.Bool(false),
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
                .type("IPMatch")
                .dataId(ipset.id())
                .negated(false)
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
        - type: IPMatch
          dataId: ${ipset.id}
          negated: false
```
{{% /example %}}
{{% /examples %}}
## Nested Fields

### `predicate`

See the [WAF Documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_Predicate.html) for more information.

#### Arguments

* `type` - (Required) The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`
* `data_id` - (Required) The unique identifier of a predicate, such as the ID of a `ByteMatchSet` or `IPSet`.
* `negated` - (Required) Whether to use the settings or the negated settings that you specified in the objects.


## Import

Using `pulumi import`, import WAF Regional Rule using the id. For example:

```sh
 $ pulumi import aws:wafregional/rule:Rule wafrule a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 