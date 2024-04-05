Provides a WAF Rate Based Rule Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ipset = new aws.waf.IpSet("ipset", {ipSetDescriptors: [{
    type: "IPV4",
    value: "192.0.7.0/24",
}]});
const wafrule = new aws.waf.RateBasedRule("wafrule", {
    metricName: "tfWAFRule",
    rateKey: "IP",
    rateLimit: 100,
    predicates: [{
        dataId: ipset.id,
        negated: false,
        type: "IPMatch",
    }],
}, {
    dependsOn: [ipset],
});
```
```python
import pulumi
import pulumi_aws as aws

ipset = aws.waf.IpSet("ipset", ip_set_descriptors=[aws.waf.IpSetIpSetDescriptorArgs(
    type="IPV4",
    value="192.0.7.0/24",
)])
wafrule = aws.waf.RateBasedRule("wafrule",
    metric_name="tfWAFRule",
    rate_key="IP",
    rate_limit=100,
    predicates=[aws.waf.RateBasedRulePredicateArgs(
        data_id=ipset.id,
        negated=False,
        type="IPMatch",
    )],
    opts=pulumi.ResourceOptions(depends_on=[ipset]))
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

    var wafrule = new Aws.Waf.RateBasedRule("wafrule", new()
    {
        MetricName = "tfWAFRule",
        RateKey = "IP",
        RateLimit = 100,
        Predicates = new[]
        {
            new Aws.Waf.Inputs.RateBasedRulePredicateArgs
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
		_, err = waf.NewRateBasedRule(ctx, "wafrule", &waf.RateBasedRuleArgs{
			MetricName: pulumi.String("tfWAFRule"),
			RateKey:    pulumi.String("IP"),
			RateLimit:  pulumi.Int(100),
			Predicates: waf.RateBasedRulePredicateArray{
				&waf.RateBasedRulePredicateArgs{
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
import com.pulumi.aws.waf.RateBasedRule;
import com.pulumi.aws.waf.RateBasedRuleArgs;
import com.pulumi.aws.waf.inputs.RateBasedRulePredicateArgs;
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

        var wafrule = new RateBasedRule("wafrule", RateBasedRuleArgs.builder()        
            .metricName("tfWAFRule")
            .rateKey("IP")
            .rateLimit(100)
            .predicates(RateBasedRulePredicateArgs.builder()
                .dataId(ipset.id())
                .negated(false)
                .type("IPMatch")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(ipset)
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
    type: aws:waf:RateBasedRule
    properties:
      metricName: tfWAFRule
      rateKey: IP
      rateLimit: 100
      predicates:
        - dataId: ${ipset.id}
          negated: false
          type: IPMatch
    options:
      dependson:
        - ${ipset}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAF Rated Based Rule using the id. For example:

```sh
 $ pulumi import aws:waf/rateBasedRule:RateBasedRule wafrule a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 