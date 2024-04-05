Manages an association with WAF Regional Web ACL.

> **Note:** An Application Load Balancer can only be associated with one WAF Regional WebACL.

{{% examples %}}
## Example Usage
{{% example %}}
### Application Load Balancer Association

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ipset = new aws.wafregional.IpSet("ipset", {ipSetDescriptors: [{
    type: "IPV4",
    value: "192.0.7.0/24",
}]});
const fooRule = new aws.wafregional.Rule("fooRule", {
    metricName: "tfWAFRule",
    predicates: [{
        dataId: ipset.id,
        negated: false,
        type: "IPMatch",
    }],
});
const fooWebAcl = new aws.wafregional.WebAcl("fooWebAcl", {
    metricName: "foo",
    defaultAction: {
        type: "ALLOW",
    },
    rules: [{
        action: {
            type: "BLOCK",
        },
        priority: 1,
        ruleId: fooRule.id,
    }],
});
const fooVpc = new aws.ec2.Vpc("fooVpc", {cidrBlock: "10.1.0.0/16"});
const available = aws.getAvailabilityZones({});
const fooSubnet = new aws.ec2.Subnet("fooSubnet", {
    vpcId: fooVpc.id,
    cidrBlock: "10.1.1.0/24",
    availabilityZone: available.then(available => available.names?.[0]),
});
const bar = new aws.ec2.Subnet("bar", {
    vpcId: fooVpc.id,
    cidrBlock: "10.1.2.0/24",
    availabilityZone: available.then(available => available.names?.[1]),
});
const fooLoadBalancer = new aws.alb.LoadBalancer("fooLoadBalancer", {
    internal: true,
    subnets: [
        fooSubnet.id,
        bar.id,
    ],
});
const fooWebAclAssociation = new aws.wafregional.WebAclAssociation("fooWebAclAssociation", {
    resourceArn: fooLoadBalancer.arn,
    webAclId: fooWebAcl.id,
});
```
```python
import pulumi
import pulumi_aws as aws

ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[aws.wafregional.IpSetIpSetDescriptorArgs(
    type="IPV4",
    value="192.0.7.0/24",
)])
foo_rule = aws.wafregional.Rule("fooRule",
    metric_name="tfWAFRule",
    predicates=[aws.wafregional.RulePredicateArgs(
        data_id=ipset.id,
        negated=False,
        type="IPMatch",
    )])
foo_web_acl = aws.wafregional.WebAcl("fooWebAcl",
    metric_name="foo",
    default_action=aws.wafregional.WebAclDefaultActionArgs(
        type="ALLOW",
    ),
    rules=[aws.wafregional.WebAclRuleArgs(
        action=aws.wafregional.WebAclRuleActionArgs(
            type="BLOCK",
        ),
        priority=1,
        rule_id=foo_rule.id,
    )])
foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.1.0.0/16")
available = aws.get_availability_zones()
foo_subnet = aws.ec2.Subnet("fooSubnet",
    vpc_id=foo_vpc.id,
    cidr_block="10.1.1.0/24",
    availability_zone=available.names[0])
bar = aws.ec2.Subnet("bar",
    vpc_id=foo_vpc.id,
    cidr_block="10.1.2.0/24",
    availability_zone=available.names[1])
foo_load_balancer = aws.alb.LoadBalancer("fooLoadBalancer",
    internal=True,
    subnets=[
        foo_subnet.id,
        bar.id,
    ])
foo_web_acl_association = aws.wafregional.WebAclAssociation("fooWebAclAssociation",
    resource_arn=foo_load_balancer.arn,
    web_acl_id=foo_web_acl.id)
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

    var fooRule = new Aws.WafRegional.Rule("fooRule", new()
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

    var fooWebAcl = new Aws.WafRegional.WebAcl("fooWebAcl", new()
    {
        MetricName = "foo",
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
                RuleId = fooRule.Id,
            },
        },
    });

    var fooVpc = new Aws.Ec2.Vpc("fooVpc", new()
    {
        CidrBlock = "10.1.0.0/16",
    });

    var available = Aws.GetAvailabilityZones.Invoke();

    var fooSubnet = new Aws.Ec2.Subnet("fooSubnet", new()
    {
        VpcId = fooVpc.Id,
        CidrBlock = "10.1.1.0/24",
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
    });

    var bar = new Aws.Ec2.Subnet("bar", new()
    {
        VpcId = fooVpc.Id,
        CidrBlock = "10.1.2.0/24",
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[1]),
    });

    var fooLoadBalancer = new Aws.Alb.LoadBalancer("fooLoadBalancer", new()
    {
        Internal = true,
        Subnets = new[]
        {
            fooSubnet.Id,
            bar.Id,
        },
    });

    var fooWebAclAssociation = new Aws.WafRegional.WebAclAssociation("fooWebAclAssociation", new()
    {
        ResourceArn = fooLoadBalancer.Arn,
        WebAclId = fooWebAcl.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/alb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
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
		fooRule, err := wafregional.NewRule(ctx, "fooRule", &wafregional.RuleArgs{
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
		fooWebAcl, err := wafregional.NewWebAcl(ctx, "fooWebAcl", &wafregional.WebAclArgs{
			MetricName: pulumi.String("foo"),
			DefaultAction: &wafregional.WebAclDefaultActionArgs{
				Type: pulumi.String("ALLOW"),
			},
			Rules: wafregional.WebAclRuleArray{
				&wafregional.WebAclRuleArgs{
					Action: &wafregional.WebAclRuleActionArgs{
						Type: pulumi.String("BLOCK"),
					},
					Priority: pulumi.Int(1),
					RuleId:   fooRule.ID(),
				},
			},
		})
		if err != nil {
			return err
		}
		fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		available, err := aws.GetAvailabilityZones(ctx, nil, nil)
		if err != nil {
			return err
		}
		fooSubnet, err := ec2.NewSubnet(ctx, "fooSubnet", &ec2.SubnetArgs{
			VpcId:            fooVpc.ID(),
			CidrBlock:        pulumi.String("10.1.1.0/24"),
			AvailabilityZone: *pulumi.String(available.Names[0]),
		})
		if err != nil {
			return err
		}
		bar, err := ec2.NewSubnet(ctx, "bar", &ec2.SubnetArgs{
			VpcId:            fooVpc.ID(),
			CidrBlock:        pulumi.String("10.1.2.0/24"),
			AvailabilityZone: *pulumi.String(available.Names[1]),
		})
		if err != nil {
			return err
		}
		fooLoadBalancer, err := alb.NewLoadBalancer(ctx, "fooLoadBalancer", &alb.LoadBalancerArgs{
			Internal: pulumi.Bool(true),
			Subnets: pulumi.StringArray{
				fooSubnet.ID(),
				bar.ID(),
			},
		})
		if err != nil {
			return err
		}
		_, err = wafregional.NewWebAclAssociation(ctx, "fooWebAclAssociation", &wafregional.WebAclAssociationArgs{
			ResourceArn: fooLoadBalancer.Arn,
			WebAclId:    fooWebAcl.ID(),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.alb.LoadBalancer;
import com.pulumi.aws.alb.LoadBalancerArgs;
import com.pulumi.aws.wafregional.WebAclAssociation;
import com.pulumi.aws.wafregional.WebAclAssociationArgs;
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

        var fooRule = new Rule("fooRule", RuleArgs.builder()        
            .metricName("tfWAFRule")
            .predicates(RulePredicateArgs.builder()
                .dataId(ipset.id())
                .negated(false)
                .type("IPMatch")
                .build())
            .build());

        var fooWebAcl = new WebAcl("fooWebAcl", WebAclArgs.builder()        
            .metricName("foo")
            .defaultAction(WebAclDefaultActionArgs.builder()
                .type("ALLOW")
                .build())
            .rules(WebAclRuleArgs.builder()
                .action(WebAclRuleActionArgs.builder()
                    .type("BLOCK")
                    .build())
                .priority(1)
                .ruleId(fooRule.id())
                .build())
            .build());

        var fooVpc = new Vpc("fooVpc", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .build());

        final var available = AwsFunctions.getAvailabilityZones();

        var fooSubnet = new Subnet("fooSubnet", SubnetArgs.builder()        
            .vpcId(fooVpc.id())
            .cidrBlock("10.1.1.0/24")
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]))
            .build());

        var bar = new Subnet("bar", SubnetArgs.builder()        
            .vpcId(fooVpc.id())
            .cidrBlock("10.1.2.0/24")
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[1]))
            .build());

        var fooLoadBalancer = new LoadBalancer("fooLoadBalancer", LoadBalancerArgs.builder()        
            .internal(true)
            .subnets(            
                fooSubnet.id(),
                bar.id())
            .build());

        var fooWebAclAssociation = new WebAclAssociation("fooWebAclAssociation", WebAclAssociationArgs.builder()        
            .resourceArn(fooLoadBalancer.arn())
            .webAclId(fooWebAcl.id())
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
  fooRule:
    type: aws:wafregional:Rule
    properties:
      metricName: tfWAFRule
      predicates:
        - dataId: ${ipset.id}
          negated: false
          type: IPMatch
  fooWebAcl:
    type: aws:wafregional:WebAcl
    properties:
      metricName: foo
      defaultAction:
        type: ALLOW
      rules:
        - action:
            type: BLOCK
          priority: 1
          ruleId: ${fooRule.id}
  fooVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
  fooSubnet:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${fooVpc.id}
      cidrBlock: 10.1.1.0/24
      availabilityZone: ${available.names[0]}
  bar:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${fooVpc.id}
      cidrBlock: 10.1.2.0/24
      availabilityZone: ${available.names[1]}
  fooLoadBalancer:
    type: aws:alb:LoadBalancer
    properties:
      internal: true
      subnets:
        - ${fooSubnet.id}
        - ${bar.id}
  fooWebAclAssociation:
    type: aws:wafregional:WebAclAssociation
    properties:
      resourceArn: ${fooLoadBalancer.arn}
      webAclId: ${fooWebAcl.id}
variables:
  available:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### API Gateway Association

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as crypto from "crypto";

const ipset = new aws.wafregional.IpSet("ipset", {ipSetDescriptors: [{
    type: "IPV4",
    value: "192.0.7.0/24",
}]});
const fooRule = new aws.wafregional.Rule("fooRule", {
    metricName: "tfWAFRule",
    predicates: [{
        dataId: ipset.id,
        negated: false,
        type: "IPMatch",
    }],
});
const fooWebAcl = new aws.wafregional.WebAcl("fooWebAcl", {
    metricName: "foo",
    defaultAction: {
        type: "ALLOW",
    },
    rules: [{
        action: {
            type: "BLOCK",
        },
        priority: 1,
        ruleId: fooRule.id,
    }],
});
const exampleRestApi = new aws.apigateway.RestApi("exampleRestApi", {body: JSON.stringify({
    openapi: "3.0.1",
    info: {
        title: "example",
        version: "1.0",
    },
    paths: {
        "/path1": {
            get: {
                "x-amazon-apigateway-integration": {
                    httpMethod: "GET",
                    payloadFormatVersion: "1.0",
                    type: "HTTP_PROXY",
                    uri: "https://ip-ranges.amazonaws.com/ip-ranges.json",
                },
            },
        },
    },
})});
const exampleDeployment = new aws.apigateway.Deployment("exampleDeployment", {
    restApi: exampleRestApi.id,
    triggers: {
        redeployment: exampleRestApi.body.apply(body => JSON.stringify(body)).apply(toJSON => crypto.createHash('sha1').update(toJSON).digest('hex')),
    },
});
const exampleStage = new aws.apigateway.Stage("exampleStage", {
    deployment: exampleDeployment.id,
    restApi: exampleRestApi.id,
    stageName: "example",
});
const association = new aws.wafregional.WebAclAssociation("association", {
    resourceArn: exampleStage.arn,
    webAclId: fooWebAcl.id,
});
```
```python
import pulumi
import hashlib
import json
import pulumi_aws as aws

ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[aws.wafregional.IpSetIpSetDescriptorArgs(
    type="IPV4",
    value="192.0.7.0/24",
)])
foo_rule = aws.wafregional.Rule("fooRule",
    metric_name="tfWAFRule",
    predicates=[aws.wafregional.RulePredicateArgs(
        data_id=ipset.id,
        negated=False,
        type="IPMatch",
    )])
foo_web_acl = aws.wafregional.WebAcl("fooWebAcl",
    metric_name="foo",
    default_action=aws.wafregional.WebAclDefaultActionArgs(
        type="ALLOW",
    ),
    rules=[aws.wafregional.WebAclRuleArgs(
        action=aws.wafregional.WebAclRuleActionArgs(
            type="BLOCK",
        ),
        priority=1,
        rule_id=foo_rule.id,
    )])
example_rest_api = aws.apigateway.RestApi("exampleRestApi", body=json.dumps({
    "openapi": "3.0.1",
    "info": {
        "title": "example",
        "version": "1.0",
    },
    "paths": {
        "/path1": {
            "get": {
                "x-amazon-apigateway-integration": {
                    "httpMethod": "GET",
                    "payloadFormatVersion": "1.0",
                    "type": "HTTP_PROXY",
                    "uri": "https://ip-ranges.amazonaws.com/ip-ranges.json",
                },
            },
        },
    },
}))
example_deployment = aws.apigateway.Deployment("exampleDeployment",
    rest_api=example_rest_api.id,
    triggers={
        "redeployment": example_rest_api.body.apply(lambda body: json.dumps(body)).apply(lambda to_json: hashlib.sha1(to_json.encode()).hexdigest()),
    })
example_stage = aws.apigateway.Stage("exampleStage",
    deployment=example_deployment.id,
    rest_api=example_rest_api.id,
    stage_name="example")
association = aws.wafregional.WebAclAssociation("association",
    resource_arn=example_stage.arn,
    web_acl_id=foo_web_acl.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

	
string ComputeSHA1(string input) 
{
    var hash = SHA1.Create().ComputeHash(Encoding.UTF8.GetBytes(input));
    return BitConverter.ToString(hash).Replace("-","").ToLowerInvariant();
}

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

    var fooRule = new Aws.WafRegional.Rule("fooRule", new()
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

    var fooWebAcl = new Aws.WafRegional.WebAcl("fooWebAcl", new()
    {
        MetricName = "foo",
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
                RuleId = fooRule.Id,
            },
        },
    });

    var exampleRestApi = new Aws.ApiGateway.RestApi("exampleRestApi", new()
    {
        Body = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["openapi"] = "3.0.1",
            ["info"] = new Dictionary<string, object?>
            {
                ["title"] = "example",
                ["version"] = "1.0",
            },
            ["paths"] = new Dictionary<string, object?>
            {
                ["/path1"] = new Dictionary<string, object?>
                {
                    ["get"] = new Dictionary<string, object?>
                    {
                        ["x-amazon-apigateway-integration"] = new Dictionary<string, object?>
                        {
                            ["httpMethod"] = "GET",
                            ["payloadFormatVersion"] = "1.0",
                            ["type"] = "HTTP_PROXY",
                            ["uri"] = "https://ip-ranges.amazonaws.com/ip-ranges.json",
                        },
                    },
                },
            },
        }),
    });

    var exampleDeployment = new Aws.ApiGateway.Deployment("exampleDeployment", new()
    {
        RestApi = exampleRestApi.Id,
        Triggers = 
        {
            { "redeployment", exampleRestApi.Body.Apply(body => JsonSerializer.Serialize(body)).Apply(toJSON => ComputeSHA1(toJSON)) },
        },
    });

    var exampleStage = new Aws.ApiGateway.Stage("exampleStage", new()
    {
        Deployment = exampleDeployment.Id,
        RestApi = exampleRestApi.Id,
        StageName = "example",
    });

    var association = new Aws.WafRegional.WebAclAssociation("association", new()
    {
        ResourceArn = exampleStage.Arn,
        WebAclId = fooWebAcl.Id,
    });

});
```
```go
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func sha1Hash(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

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
		fooRule, err := wafregional.NewRule(ctx, "fooRule", &wafregional.RuleArgs{
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
		fooWebAcl, err := wafregional.NewWebAcl(ctx, "fooWebAcl", &wafregional.WebAclArgs{
			MetricName: pulumi.String("foo"),
			DefaultAction: &wafregional.WebAclDefaultActionArgs{
				Type: pulumi.String("ALLOW"),
			},
			Rules: wafregional.WebAclRuleArray{
				&wafregional.WebAclRuleArgs{
					Action: &wafregional.WebAclRuleActionArgs{
						Type: pulumi.String("BLOCK"),
					},
					Priority: pulumi.Int(1),
					RuleId:   fooRule.ID(),
				},
			},
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"openapi": "3.0.1",
			"info": map[string]interface{}{
				"title":   "example",
				"version": "1.0",
			},
			"paths": map[string]interface{}{
				"/path1": map[string]interface{}{
					"get": map[string]interface{}{
						"x-amazon-apigateway-integration": map[string]interface{}{
							"httpMethod":           "GET",
							"payloadFormatVersion": "1.0",
							"type":                 "HTTP_PROXY",
							"uri":                  "https://ip-ranges.amazonaws.com/ip-ranges.json",
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleRestApi, err := apigateway.NewRestApi(ctx, "exampleRestApi", &apigateway.RestApiArgs{
			Body: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		exampleDeployment, err := apigateway.NewDeployment(ctx, "exampleDeployment", &apigateway.DeploymentArgs{
			RestApi: exampleRestApi.ID(),
			Triggers: pulumi.StringMap{
				"redeployment": exampleRestApi.Body.ApplyT(func(body *string) (pulumi.String, error) {
					var _zero pulumi.String
					tmpJSON1, err := json.Marshal(body)
					if err != nil {
						return _zero, err
					}
					json1 := string(tmpJSON1)
					return pulumi.String(json1), nil
				}).(pulumi.StringOutput).ApplyT(func(toJSON string) (pulumi.String, error) {
					return pulumi.String(sha1Hash(toJSON)), nil
				}).(pulumi.StringOutput),
			},
		})
		if err != nil {
			return err
		}
		exampleStage, err := apigateway.NewStage(ctx, "exampleStage", &apigateway.StageArgs{
			Deployment: exampleDeployment.ID(),
			RestApi:    exampleRestApi.ID(),
			StageName:  pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = wafregional.NewWebAclAssociation(ctx, "association", &wafregional.WebAclAssociationArgs{
			ResourceArn: exampleStage.Arn,
			WebAclId:    fooWebAcl.ID(),
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
import com.pulumi.aws.apigateway.RestApi;
import com.pulumi.aws.apigateway.RestApiArgs;
import com.pulumi.aws.apigateway.Deployment;
import com.pulumi.aws.apigateway.DeploymentArgs;
import com.pulumi.aws.apigateway.Stage;
import com.pulumi.aws.apigateway.StageArgs;
import com.pulumi.aws.wafregional.WebAclAssociation;
import com.pulumi.aws.wafregional.WebAclAssociationArgs;
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
        var ipset = new IpSet("ipset", IpSetArgs.builder()        
            .ipSetDescriptors(IpSetIpSetDescriptorArgs.builder()
                .type("IPV4")
                .value("192.0.7.0/24")
                .build())
            .build());

        var fooRule = new Rule("fooRule", RuleArgs.builder()        
            .metricName("tfWAFRule")
            .predicates(RulePredicateArgs.builder()
                .dataId(ipset.id())
                .negated(false)
                .type("IPMatch")
                .build())
            .build());

        var fooWebAcl = new WebAcl("fooWebAcl", WebAclArgs.builder()        
            .metricName("foo")
            .defaultAction(WebAclDefaultActionArgs.builder()
                .type("ALLOW")
                .build())
            .rules(WebAclRuleArgs.builder()
                .action(WebAclRuleActionArgs.builder()
                    .type("BLOCK")
                    .build())
                .priority(1)
                .ruleId(fooRule.id())
                .build())
            .build());

        var exampleRestApi = new RestApi("exampleRestApi", RestApiArgs.builder()        
            .body(serializeJson(
                jsonObject(
                    jsonProperty("openapi", "3.0.1"),
                    jsonProperty("info", jsonObject(
                        jsonProperty("title", "example"),
                        jsonProperty("version", "1.0")
                    )),
                    jsonProperty("paths", jsonObject(
                        jsonProperty("/path1", jsonObject(
                            jsonProperty("get", jsonObject(
                                jsonProperty("x-amazon-apigateway-integration", jsonObject(
                                    jsonProperty("httpMethod", "GET"),
                                    jsonProperty("payloadFormatVersion", "1.0"),
                                    jsonProperty("type", "HTTP_PROXY"),
                                    jsonProperty("uri", "https://ip-ranges.amazonaws.com/ip-ranges.json")
                                ))
                            ))
                        ))
                    ))
                )))
            .build());

        var exampleDeployment = new Deployment("exampleDeployment", DeploymentArgs.builder()        
            .restApi(exampleRestApi.id())
            .triggers(Map.of("redeployment", exampleRestApi.body().applyValue(body -> serializeJson(
                body)).applyValue(toJSON -> computeSHA1(toJSON))))
            .build());

        var exampleStage = new Stage("exampleStage", StageArgs.builder()        
            .deployment(exampleDeployment.id())
            .restApi(exampleRestApi.id())
            .stageName("example")
            .build());

        var association = new WebAclAssociation("association", WebAclAssociationArgs.builder()        
            .resourceArn(exampleStage.arn())
            .webAclId(fooWebAcl.id())
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAF Regional Web ACL Association using their `web_acl_id:resource_arn`. For example:

```sh
 $ pulumi import aws:wafregional/webAclAssociation:WebAclAssociation foo web_acl_id:resource_arn
```
 