Provides an AWS Network Firewall Resource Policy Resource for a rule group or firewall policy.

{{% examples %}}
## Example Usage
{{% example %}}
### For a Firewall Policy resource

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkfirewall.ResourcePolicy("example", {
    resourceArn: aws_networkfirewall_firewall_policy.example.arn,
    policy: JSON.stringify({
        Statement: [{
            Action: [
                "network-firewall:ListFirewallPolicies",
                "network-firewall:CreateFirewall",
                "network-firewall:UpdateFirewall",
                "network-firewall:AssociateFirewallPolicy",
            ],
            Effect: "Allow",
            Resource: aws_networkfirewall_firewall_policy.example.arn,
            Principal: {
                AWS: "arn:aws:iam::123456789012:root",
            },
        }],
        Version: "2012-10-17",
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.networkfirewall.ResourcePolicy("example",
    resource_arn=aws_networkfirewall_firewall_policy["example"]["arn"],
    policy=json.dumps({
        "Statement": [{
            "Action": [
                "network-firewall:ListFirewallPolicies",
                "network-firewall:CreateFirewall",
                "network-firewall:UpdateFirewall",
                "network-firewall:AssociateFirewallPolicy",
            ],
            "Effect": "Allow",
            "Resource": aws_networkfirewall_firewall_policy["example"]["arn"],
            "Principal": {
                "AWS": "arn:aws:iam::123456789012:root",
            },
        }],
        "Version": "2012-10-17",
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkFirewall.ResourcePolicy("example", new()
    {
        ResourceArn = aws_networkfirewall_firewall_policy.Example.Arn,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[]
                    {
                        "network-firewall:ListFirewallPolicies",
                        "network-firewall:CreateFirewall",
                        "network-firewall:UpdateFirewall",
                        "network-firewall:AssociateFirewallPolicy",
                    },
                    ["Effect"] = "Allow",
                    ["Resource"] = aws_networkfirewall_firewall_policy.Example.Arn,
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["AWS"] = "arn:aws:iam::123456789012:root",
                    },
                },
            },
            ["Version"] = "2012-10-17",
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": []string{
						"network-firewall:ListFirewallPolicies",
						"network-firewall:CreateFirewall",
						"network-firewall:UpdateFirewall",
						"network-firewall:AssociateFirewallPolicy",
					},
					"Effect":   "Allow",
					"Resource": aws_networkfirewall_firewall_policy.Example.Arn,
					"Principal": map[string]interface{}{
						"AWS": "arn:aws:iam::123456789012:root",
					},
				},
			},
			"Version": "2012-10-17",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = networkfirewall.NewResourcePolicy(ctx, "example", &networkfirewall.ResourcePolicyArgs{
			ResourceArn: pulumi.Any(aws_networkfirewall_firewall_policy.Example.Arn),
			Policy:      pulumi.String(json0),
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
import com.pulumi.aws.networkfirewall.ResourcePolicy;
import com.pulumi.aws.networkfirewall.ResourcePolicyArgs;
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
        var example = new ResourcePolicy("example", ResourcePolicyArgs.builder()        
            .resourceArn(aws_networkfirewall_firewall_policy.example().arn())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", jsonArray(
                            "network-firewall:ListFirewallPolicies", 
                            "network-firewall:CreateFirewall", 
                            "network-firewall:UpdateFirewall", 
                            "network-firewall:AssociateFirewallPolicy"
                        )),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Resource", aws_networkfirewall_firewall_policy.example().arn()),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", "arn:aws:iam::123456789012:root")
                        ))
                    ))),
                    jsonProperty("Version", "2012-10-17")
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:ResourcePolicy
    properties:
      resourceArn: ${aws_networkfirewall_firewall_policy.example.arn}
      # policy's Action element must include all of the following operations
      policy:
        fn::toJSON:
          Statement:
            - Action:
                - network-firewall:ListFirewallPolicies
                - network-firewall:CreateFirewall
                - network-firewall:UpdateFirewall
                - network-firewall:AssociateFirewallPolicy
              Effect: Allow
              Resource: ${aws_networkfirewall_firewall_policy.example.arn}
              Principal:
                AWS: arn:aws:iam::123456789012:root
          Version: 2012-10-17
```
{{% /example %}}
{{% example %}}
### For a Rule Group resource

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkfirewall.ResourcePolicy("example", {
    resourceArn: aws_networkfirewall_rule_group.example.arn,
    policy: JSON.stringify({
        Statement: [{
            Action: [
                "network-firewall:ListRuleGroups",
                "network-firewall:CreateFirewallPolicy",
                "network-firewall:UpdateFirewallPolicy",
            ],
            Effect: "Allow",
            Resource: aws_networkfirewall_rule_group.example.arn,
            Principal: {
                AWS: "arn:aws:iam::123456789012:root",
            },
        }],
        Version: "2012-10-17",
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.networkfirewall.ResourcePolicy("example",
    resource_arn=aws_networkfirewall_rule_group["example"]["arn"],
    policy=json.dumps({
        "Statement": [{
            "Action": [
                "network-firewall:ListRuleGroups",
                "network-firewall:CreateFirewallPolicy",
                "network-firewall:UpdateFirewallPolicy",
            ],
            "Effect": "Allow",
            "Resource": aws_networkfirewall_rule_group["example"]["arn"],
            "Principal": {
                "AWS": "arn:aws:iam::123456789012:root",
            },
        }],
        "Version": "2012-10-17",
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkFirewall.ResourcePolicy("example", new()
    {
        ResourceArn = aws_networkfirewall_rule_group.Example.Arn,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[]
                    {
                        "network-firewall:ListRuleGroups",
                        "network-firewall:CreateFirewallPolicy",
                        "network-firewall:UpdateFirewallPolicy",
                    },
                    ["Effect"] = "Allow",
                    ["Resource"] = aws_networkfirewall_rule_group.Example.Arn,
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["AWS"] = "arn:aws:iam::123456789012:root",
                    },
                },
            },
            ["Version"] = "2012-10-17",
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": []string{
						"network-firewall:ListRuleGroups",
						"network-firewall:CreateFirewallPolicy",
						"network-firewall:UpdateFirewallPolicy",
					},
					"Effect":   "Allow",
					"Resource": aws_networkfirewall_rule_group.Example.Arn,
					"Principal": map[string]interface{}{
						"AWS": "arn:aws:iam::123456789012:root",
					},
				},
			},
			"Version": "2012-10-17",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = networkfirewall.NewResourcePolicy(ctx, "example", &networkfirewall.ResourcePolicyArgs{
			ResourceArn: pulumi.Any(aws_networkfirewall_rule_group.Example.Arn),
			Policy:      pulumi.String(json0),
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
import com.pulumi.aws.networkfirewall.ResourcePolicy;
import com.pulumi.aws.networkfirewall.ResourcePolicyArgs;
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
        var example = new ResourcePolicy("example", ResourcePolicyArgs.builder()        
            .resourceArn(aws_networkfirewall_rule_group.example().arn())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", jsonArray(
                            "network-firewall:ListRuleGroups", 
                            "network-firewall:CreateFirewallPolicy", 
                            "network-firewall:UpdateFirewallPolicy"
                        )),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Resource", aws_networkfirewall_rule_group.example().arn()),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", "arn:aws:iam::123456789012:root")
                        ))
                    ))),
                    jsonProperty("Version", "2012-10-17")
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:ResourcePolicy
    properties:
      resourceArn: ${aws_networkfirewall_rule_group.example.arn}
      # policy's Action element must include all of the following operations
      policy:
        fn::toJSON:
          Statement:
            - Action:
                - network-firewall:ListRuleGroups
                - network-firewall:CreateFirewallPolicy
                - network-firewall:UpdateFirewallPolicy
              Effect: Allow
              Resource: ${aws_networkfirewall_rule_group.example.arn}
              Principal:
                AWS: arn:aws:iam::123456789012:root
          Version: 2012-10-17
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Network Firewall Resource Policies using the `resource_arn`. For example:


```sh
 $ pulumi import aws:networkfirewall/resourcePolicy:ResourcePolicy example arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
```
 