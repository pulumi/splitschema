Provides a resource to create an AWS Firewall Manager policy. You need to be using AWS organizations and have enabled the Firewall Manager administrator account.

> **NOTE:** Due to limitations with testing, we provide it as best effort. If you find it useful, and have the ability to help test or notice issues, consider reaching out to us on GitHub.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRuleGroup = new aws.wafregional.RuleGroup("exampleRuleGroup", {metricName: "WAFRuleGroupExample"});
const examplePolicy = new aws.fms.Policy("examplePolicy", {
    excludeResourceTags: false,
    remediationEnabled: false,
    resourceType: "AWS::ElasticLoadBalancingV2::LoadBalancer",
    securityServicePolicyData: {
        type: "WAF",
        managedServiceData: exampleRuleGroup.id.apply(id => JSON.stringify({
            type: "WAF",
            ruleGroups: [{
                id: id,
                overrideAction: {
                    type: "COUNT",
                },
            }],
            defaultAction: {
                type: "BLOCK",
            },
            overrideCustomerWebACLAssociation: false,
        })),
    },
    tags: {
        Name: "example-fms-policy",
    },
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_rule_group = aws.wafregional.RuleGroup("exampleRuleGroup", metric_name="WAFRuleGroupExample")
example_policy = aws.fms.Policy("examplePolicy",
    exclude_resource_tags=False,
    remediation_enabled=False,
    resource_type="AWS::ElasticLoadBalancingV2::LoadBalancer",
    security_service_policy_data=aws.fms.PolicySecurityServicePolicyDataArgs(
        type="WAF",
        managed_service_data=example_rule_group.id.apply(lambda id: json.dumps({
            "type": "WAF",
            "ruleGroups": [{
                "id": id,
                "overrideAction": {
                    "type": "COUNT",
                },
            }],
            "defaultAction": {
                "type": "BLOCK",
            },
            "overrideCustomerWebACLAssociation": False,
        })),
    ),
    tags={
        "Name": "example-fms-policy",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleRuleGroup = new Aws.WafRegional.RuleGroup("exampleRuleGroup", new()
    {
        MetricName = "WAFRuleGroupExample",
    });

    var examplePolicy = new Aws.Fms.Policy("examplePolicy", new()
    {
        ExcludeResourceTags = false,
        RemediationEnabled = false,
        ResourceType = "AWS::ElasticLoadBalancingV2::LoadBalancer",
        SecurityServicePolicyData = new Aws.Fms.Inputs.PolicySecurityServicePolicyDataArgs
        {
            Type = "WAF",
            ManagedServiceData = exampleRuleGroup.Id.Apply(id => JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["type"] = "WAF",
                ["ruleGroups"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["id"] = id,
                        ["overrideAction"] = new Dictionary<string, object?>
                        {
                            ["type"] = "COUNT",
                        },
                    },
                },
                ["defaultAction"] = new Dictionary<string, object?>
                {
                    ["type"] = "BLOCK",
                },
                ["overrideCustomerWebACLAssociation"] = false,
            })),
        },
        Tags = 
        {
            { "Name", "example-fms-policy" },
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRuleGroup, err := wafregional.NewRuleGroup(ctx, "exampleRuleGroup", &wafregional.RuleGroupArgs{
			MetricName: pulumi.String("WAFRuleGroupExample"),
		})
		if err != nil {
			return err
		}
		_, err = fms.NewPolicy(ctx, "examplePolicy", &fms.PolicyArgs{
			ExcludeResourceTags: pulumi.Bool(false),
			RemediationEnabled:  pulumi.Bool(false),
			ResourceType:        pulumi.String("AWS::ElasticLoadBalancingV2::LoadBalancer"),
			SecurityServicePolicyData: &fms.PolicySecurityServicePolicyDataArgs{
				Type: pulumi.String("WAF"),
				ManagedServiceData: exampleRuleGroup.ID().ApplyT(func(id string) (pulumi.String, error) {
					var _zero pulumi.String
					tmpJSON0, err := json.Marshal(map[string]interface{}{
						"type": "WAF",
						"ruleGroups": []map[string]interface{}{
							map[string]interface{}{
								"id": id,
								"overrideAction": map[string]interface{}{
									"type": "COUNT",
								},
							},
						},
						"defaultAction": map[string]interface{}{
							"type": "BLOCK",
						},
						"overrideCustomerWebACLAssociation": false,
					})
					if err != nil {
						return _zero, err
					}
					json0 := string(tmpJSON0)
					return pulumi.String(json0), nil
				}).(pulumi.StringOutput),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example-fms-policy"),
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
import com.pulumi.aws.wafregional.RuleGroup;
import com.pulumi.aws.wafregional.RuleGroupArgs;
import com.pulumi.aws.fms.Policy;
import com.pulumi.aws.fms.PolicyArgs;
import com.pulumi.aws.fms.inputs.PolicySecurityServicePolicyDataArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var exampleRuleGroup = new RuleGroup("exampleRuleGroup", RuleGroupArgs.builder()        
            .metricName("WAFRuleGroupExample")
            .build());

        var examplePolicy = new Policy("examplePolicy", PolicyArgs.builder()        
            .excludeResourceTags(false)
            .remediationEnabled(false)
            .resourceType("AWS::ElasticLoadBalancingV2::LoadBalancer")
            .securityServicePolicyData(PolicySecurityServicePolicyDataArgs.builder()
                .type("WAF")
                .managedServiceData(exampleRuleGroup.id().applyValue(id -> serializeJson(
                    jsonObject(
                        jsonProperty("type", "WAF"),
                        jsonProperty("ruleGroups", jsonArray(jsonObject(
                            jsonProperty("id", id),
                            jsonProperty("overrideAction", jsonObject(
                                jsonProperty("type", "COUNT")
                            ))
                        ))),
                        jsonProperty("defaultAction", jsonObject(
                            jsonProperty("type", "BLOCK")
                        )),
                        jsonProperty("overrideCustomerWebACLAssociation", false)
                    ))))
                .build())
            .tags(Map.of("Name", "example-fms-policy"))
            .build());

    }
}
```
```yaml
resources:
  examplePolicy:
    type: aws:fms:Policy
    properties:
      excludeResourceTags: false
      remediationEnabled: false
      resourceType: AWS::ElasticLoadBalancingV2::LoadBalancer
      securityServicePolicyData:
        type: WAF
        managedServiceData:
          fn::toJSON:
            type: WAF
            ruleGroups:
              - id: ${exampleRuleGroup.id}
                overrideAction:
                  type: COUNT
            defaultAction:
              type: BLOCK
            overrideCustomerWebACLAssociation: false
      tags:
        Name: example-fms-policy
  exampleRuleGroup:
    type: aws:wafregional:RuleGroup
    properties:
      metricName: WAFRuleGroupExample
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Firewall Manager policies using the policy ID. For example:

```sh
 $ pulumi import aws:fms/policy:Policy example 5be49585-a7e3-4c49-dde1-a179fe4a619a
```
 