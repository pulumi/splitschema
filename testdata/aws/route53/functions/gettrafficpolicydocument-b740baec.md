Generates an Route53 traffic policy document in JSON format for use with resources that expect policy documents such as `aws.route53.TrafficPolicy`.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegion({});
const exampleTrafficPolicyDocument = Promise.all([current, current]).then(([current, current1]) => aws.route53.getTrafficPolicyDocument({
    recordType: "A",
    startRule: "site_switch",
    endpoints: [
        {
            id: "my_elb",
            type: "elastic-load-balancer",
            value: `elb-111111.${current.name}.elb.amazonaws.com`,
        },
        {
            id: "site_down_banner",
            type: "s3-website",
            region: current1.name,
            value: "www.example.com",
        },
    ],
    rules: [{
        id: "site_switch",
        type: "failover",
        primary: {
            endpointReference: "my_elb",
        },
        secondary: {
            endpointReference: "site_down_banner",
        },
    }],
}));
const exampleTrafficPolicy = new aws.route53.TrafficPolicy("exampleTrafficPolicy", {
    comment: "example comment",
    document: exampleTrafficPolicyDocument.then(exampleTrafficPolicyDocument => exampleTrafficPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_region()
example_traffic_policy_document = aws.route53.get_traffic_policy_document(record_type="A",
    start_rule="site_switch",
    endpoints=[
        aws.route53.GetTrafficPolicyDocumentEndpointArgs(
            id="my_elb",
            type="elastic-load-balancer",
            value=f"elb-111111.{current.name}.elb.amazonaws.com",
        ),
        aws.route53.GetTrafficPolicyDocumentEndpointArgs(
            id="site_down_banner",
            type="s3-website",
            region=current.name,
            value="www.example.com",
        ),
    ],
    rules=[aws.route53.GetTrafficPolicyDocumentRuleArgs(
        id="site_switch",
        type="failover",
        primary=aws.route53.GetTrafficPolicyDocumentRulePrimaryArgs(
            endpoint_reference="my_elb",
        ),
        secondary=aws.route53.GetTrafficPolicyDocumentRuleSecondaryArgs(
            endpoint_reference="site_down_banner",
        ),
    )])
example_traffic_policy = aws.route53.TrafficPolicy("exampleTrafficPolicy",
    comment="example comment",
    document=example_traffic_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegion.Invoke();

    var exampleTrafficPolicyDocument = Aws.Route53.GetTrafficPolicyDocument.Invoke(new()
    {
        RecordType = "A",
        StartRule = "site_switch",
        Endpoints = new[]
        {
            new Aws.Route53.Inputs.GetTrafficPolicyDocumentEndpointInputArgs
            {
                Id = "my_elb",
                Type = "elastic-load-balancer",
                Value = $"elb-111111.{current.Apply(getRegionResult => getRegionResult.Name)}.elb.amazonaws.com",
            },
            new Aws.Route53.Inputs.GetTrafficPolicyDocumentEndpointInputArgs
            {
                Id = "site_down_banner",
                Type = "s3-website",
                Region = current.Apply(getRegionResult => getRegionResult.Name),
                Value = "www.example.com",
            },
        },
        Rules = new[]
        {
            new Aws.Route53.Inputs.GetTrafficPolicyDocumentRuleInputArgs
            {
                Id = "site_switch",
                Type = "failover",
                Primary = new Aws.Route53.Inputs.GetTrafficPolicyDocumentRulePrimaryInputArgs
                {
                    EndpointReference = "my_elb",
                },
                Secondary = new Aws.Route53.Inputs.GetTrafficPolicyDocumentRuleSecondaryInputArgs
                {
                    EndpointReference = "site_down_banner",
                },
            },
        },
    });

    var exampleTrafficPolicy = new Aws.Route53.TrafficPolicy("exampleTrafficPolicy", new()
    {
        Comment = "example comment",
        Document = exampleTrafficPolicyDocument.Apply(getTrafficPolicyDocumentResult => getTrafficPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleTrafficPolicyDocument, err := route53.GetTrafficPolicyDocument(ctx, &route53.GetTrafficPolicyDocumentArgs{
			RecordType: pulumi.StringRef("A"),
			StartRule:  pulumi.StringRef("site_switch"),
			Endpoints: []route53.GetTrafficPolicyDocumentEndpoint{
				{
					Id:    "my_elb",
					Type:  pulumi.StringRef("elastic-load-balancer"),
					Value: pulumi.StringRef(fmt.Sprintf("elb-111111.%v.elb.amazonaws.com", current.Name)),
				},
				{
					Id:     "site_down_banner",
					Type:   pulumi.StringRef("s3-website"),
					Region: pulumi.StringRef(current.Name),
					Value:  pulumi.StringRef("www.example.com"),
				},
			},
			Rules: []route53.GetTrafficPolicyDocumentRule{
				{
					Id:   "site_switch",
					Type: pulumi.StringRef("failover"),
					Primary: {
						EndpointReference: pulumi.StringRef("my_elb"),
					},
					Secondary: {
						EndpointReference: pulumi.StringRef("site_down_banner"),
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = route53.NewTrafficPolicy(ctx, "exampleTrafficPolicy", &route53.TrafficPolicyArgs{
			Comment:  pulumi.String("example comment"),
			Document: *pulumi.String(exampleTrafficPolicyDocument.Json),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetTrafficPolicyDocumentArgs;
import com.pulumi.aws.route53.TrafficPolicy;
import com.pulumi.aws.route53.TrafficPolicyArgs;
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
        final var current = AwsFunctions.getRegion();

        final var exampleTrafficPolicyDocument = Route53Functions.getTrafficPolicyDocument(GetTrafficPolicyDocumentArgs.builder()
            .recordType("A")
            .startRule("site_switch")
            .endpoints(            
                GetTrafficPolicyDocumentEndpointArgs.builder()
                    .id("my_elb")
                    .type("elastic-load-balancer")
                    .value(String.format("elb-111111.%s.elb.amazonaws.com", current.applyValue(getRegionResult -> getRegionResult.name())))
                    .build(),
                GetTrafficPolicyDocumentEndpointArgs.builder()
                    .id("site_down_banner")
                    .type("s3-website")
                    .region(current.applyValue(getRegionResult -> getRegionResult.name()))
                    .value("www.example.com")
                    .build())
            .rules(GetTrafficPolicyDocumentRuleArgs.builder()
                .id("site_switch")
                .type("failover")
                .primary(GetTrafficPolicyDocumentRulePrimaryArgs.builder()
                    .endpointReference("my_elb")
                    .build())
                .secondary(GetTrafficPolicyDocumentRuleSecondaryArgs.builder()
                    .endpointReference("site_down_banner")
                    .build())
                .build())
            .build());

        var exampleTrafficPolicy = new TrafficPolicy("exampleTrafficPolicy", TrafficPolicyArgs.builder()        
            .comment("example comment")
            .document(exampleTrafficPolicyDocument.applyValue(getTrafficPolicyDocumentResult -> getTrafficPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  exampleTrafficPolicy:
    type: aws:route53:TrafficPolicy
    properties:
      comment: example comment
      document: ${exampleTrafficPolicyDocument.json}
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  exampleTrafficPolicyDocument:
    fn::invoke:
      Function: aws:route53:getTrafficPolicyDocument
      Arguments:
        recordType: A
        startRule: site_switch
        endpoints:
          - id: my_elb
            type: elastic-load-balancer
            value: elb-111111.${current.name}.elb.amazonaws.com
          - id: site_down_banner
            type: s3-website
            region: ${current.name}
            value: www.example.com
        rules:
          - id: site_switch
            type: failover
            primary:
              endpointReference: my_elb
            secondary:
              endpointReference: site_down_banner
```
{{% /example %}}
{{% example %}}
### Complex Example

The following example showcases the use of nested rules within the traffic policy document and introduces the `geoproximity` rule type.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleTrafficPolicyDocument = aws.route53.getTrafficPolicyDocument({
    recordType: "A",
    startRule: "geoproximity_rule",
    endpoints: [
        {
            id: "na_endpoint_a",
            type: "elastic-load-balancer",
            value: "elb-111111.us-west-1.elb.amazonaws.com",
        },
        {
            id: "na_endpoint_b",
            type: "elastic-load-balancer",
            value: "elb-222222.us-west-1.elb.amazonaws.com",
        },
        {
            id: "eu_endpoint",
            type: "elastic-load-balancer",
            value: "elb-333333.eu-west-1.elb.amazonaws.com",
        },
        {
            id: "ap_endpoint",
            type: "elastic-load-balancer",
            value: "elb-444444.ap-northeast-2.elb.amazonaws.com",
        },
    ],
    rules: [
        {
            id: "na_rule",
            type: "failover",
            primary: {
                endpointReference: "na_endpoint_a",
            },
            secondary: {
                endpointReference: "na_endpoint_b",
            },
        },
        {
            id: "geoproximity_rule",
            type: "geoproximity",
            geoProximityLocations: [
                {
                    region: "aws:route53:us-west-1",
                    bias: "10",
                    evaluateTargetHealth: true,
                    ruleReference: "na_rule",
                },
                {
                    region: "aws:route53:eu-west-1",
                    bias: "10",
                    evaluateTargetHealth: true,
                    endpointReference: "eu_endpoint",
                },
                {
                    region: "aws:route53:ap-northeast-2",
                    bias: "0",
                    evaluateTargetHealth: true,
                    endpointReference: "ap_endpoint",
                },
            ],
        },
    ],
});
const exampleTrafficPolicy = new aws.route53.TrafficPolicy("exampleTrafficPolicy", {
    comment: "example comment",
    document: exampleTrafficPolicyDocument.then(exampleTrafficPolicyDocument => exampleTrafficPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_traffic_policy_document = aws.route53.get_traffic_policy_document(record_type="A",
    start_rule="geoproximity_rule",
    endpoints=[
        aws.route53.GetTrafficPolicyDocumentEndpointArgs(
            id="na_endpoint_a",
            type="elastic-load-balancer",
            value="elb-111111.us-west-1.elb.amazonaws.com",
        ),
        aws.route53.GetTrafficPolicyDocumentEndpointArgs(
            id="na_endpoint_b",
            type="elastic-load-balancer",
            value="elb-222222.us-west-1.elb.amazonaws.com",
        ),
        aws.route53.GetTrafficPolicyDocumentEndpointArgs(
            id="eu_endpoint",
            type="elastic-load-balancer",
            value="elb-333333.eu-west-1.elb.amazonaws.com",
        ),
        aws.route53.GetTrafficPolicyDocumentEndpointArgs(
            id="ap_endpoint",
            type="elastic-load-balancer",
            value="elb-444444.ap-northeast-2.elb.amazonaws.com",
        ),
    ],
    rules=[
        aws.route53.GetTrafficPolicyDocumentRuleArgs(
            id="na_rule",
            type="failover",
            primary=aws.route53.GetTrafficPolicyDocumentRulePrimaryArgs(
                endpoint_reference="na_endpoint_a",
            ),
            secondary=aws.route53.GetTrafficPolicyDocumentRuleSecondaryArgs(
                endpoint_reference="na_endpoint_b",
            ),
        ),
        aws.route53.GetTrafficPolicyDocumentRuleArgs(
            id="geoproximity_rule",
            type="geoproximity",
            geo_proximity_locations=[
                aws.route53.GetTrafficPolicyDocumentRuleGeoProximityLocationArgs(
                    region="aws:route53:us-west-1",
                    bias="10",
                    evaluate_target_health=True,
                    rule_reference="na_rule",
                ),
                aws.route53.GetTrafficPolicyDocumentRuleGeoProximityLocationArgs(
                    region="aws:route53:eu-west-1",
                    bias="10",
                    evaluate_target_health=True,
                    endpoint_reference="eu_endpoint",
                ),
                aws.route53.GetTrafficPolicyDocumentRuleGeoProximityLocationArgs(
                    region="aws:route53:ap-northeast-2",
                    bias="0",
                    evaluate_target_health=True,
                    endpoint_reference="ap_endpoint",
                ),
            ],
        ),
    ])
example_traffic_policy = aws.route53.TrafficPolicy("exampleTrafficPolicy",
    comment="example comment",
    document=example_traffic_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleTrafficPolicyDocument = Aws.Route53.GetTrafficPolicyDocument.Invoke(new()
    {
        RecordType = "A",
        StartRule = "geoproximity_rule",
        Endpoints = new[]
        {
            new Aws.Route53.Inputs.GetTrafficPolicyDocumentEndpointInputArgs
            {
                Id = "na_endpoint_a",
                Type = "elastic-load-balancer",
                Value = "elb-111111.us-west-1.elb.amazonaws.com",
            },
            new Aws.Route53.Inputs.GetTrafficPolicyDocumentEndpointInputArgs
            {
                Id = "na_endpoint_b",
                Type = "elastic-load-balancer",
                Value = "elb-222222.us-west-1.elb.amazonaws.com",
            },
            new Aws.Route53.Inputs.GetTrafficPolicyDocumentEndpointInputArgs
            {
                Id = "eu_endpoint",
                Type = "elastic-load-balancer",
                Value = "elb-333333.eu-west-1.elb.amazonaws.com",
            },
            new Aws.Route53.Inputs.GetTrafficPolicyDocumentEndpointInputArgs
            {
                Id = "ap_endpoint",
                Type = "elastic-load-balancer",
                Value = "elb-444444.ap-northeast-2.elb.amazonaws.com",
            },
        },
        Rules = new[]
        {
            new Aws.Route53.Inputs.GetTrafficPolicyDocumentRuleInputArgs
            {
                Id = "na_rule",
                Type = "failover",
                Primary = new Aws.Route53.Inputs.GetTrafficPolicyDocumentRulePrimaryInputArgs
                {
                    EndpointReference = "na_endpoint_a",
                },
                Secondary = new Aws.Route53.Inputs.GetTrafficPolicyDocumentRuleSecondaryInputArgs
                {
                    EndpointReference = "na_endpoint_b",
                },
            },
            new Aws.Route53.Inputs.GetTrafficPolicyDocumentRuleInputArgs
            {
                Id = "geoproximity_rule",
                Type = "geoproximity",
                GeoProximityLocations = new[]
                {
                    new Aws.Route53.Inputs.GetTrafficPolicyDocumentRuleGeoProximityLocationInputArgs
                    {
                        Region = "aws:route53:us-west-1",
                        Bias = "10",
                        EvaluateTargetHealth = true,
                        RuleReference = "na_rule",
                    },
                    new Aws.Route53.Inputs.GetTrafficPolicyDocumentRuleGeoProximityLocationInputArgs
                    {
                        Region = "aws:route53:eu-west-1",
                        Bias = "10",
                        EvaluateTargetHealth = true,
                        EndpointReference = "eu_endpoint",
                    },
                    new Aws.Route53.Inputs.GetTrafficPolicyDocumentRuleGeoProximityLocationInputArgs
                    {
                        Region = "aws:route53:ap-northeast-2",
                        Bias = "0",
                        EvaluateTargetHealth = true,
                        EndpointReference = "ap_endpoint",
                    },
                },
            },
        },
    });

    var exampleTrafficPolicy = new Aws.Route53.TrafficPolicy("exampleTrafficPolicy", new()
    {
        Comment = "example comment",
        Document = exampleTrafficPolicyDocument.Apply(getTrafficPolicyDocumentResult => getTrafficPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleTrafficPolicyDocument, err := route53.GetTrafficPolicyDocument(ctx, &route53.GetTrafficPolicyDocumentArgs{
			RecordType: pulumi.StringRef("A"),
			StartRule:  pulumi.StringRef("geoproximity_rule"),
			Endpoints: []route53.GetTrafficPolicyDocumentEndpoint{
				{
					Id:    "na_endpoint_a",
					Type:  pulumi.StringRef("elastic-load-balancer"),
					Value: pulumi.StringRef("elb-111111.us-west-1.elb.amazonaws.com"),
				},
				{
					Id:    "na_endpoint_b",
					Type:  pulumi.StringRef("elastic-load-balancer"),
					Value: pulumi.StringRef("elb-222222.us-west-1.elb.amazonaws.com"),
				},
				{
					Id:    "eu_endpoint",
					Type:  pulumi.StringRef("elastic-load-balancer"),
					Value: pulumi.StringRef("elb-333333.eu-west-1.elb.amazonaws.com"),
				},
				{
					Id:    "ap_endpoint",
					Type:  pulumi.StringRef("elastic-load-balancer"),
					Value: pulumi.StringRef("elb-444444.ap-northeast-2.elb.amazonaws.com"),
				},
			},
			Rules: pulumi.Array{
				route53.GetTrafficPolicyDocumentRule{
					Id:   "na_rule",
					Type: pulumi.StringRef("failover"),
					Primary: route53.GetTrafficPolicyDocumentRulePrimary{
						EndpointReference: pulumi.StringRef("na_endpoint_a"),
					},
					Secondary: route53.GetTrafficPolicyDocumentRuleSecondary{
						EndpointReference: pulumi.StringRef("na_endpoint_b"),
					},
				},
				route53.GetTrafficPolicyDocumentRule{
					Id:   "geoproximity_rule",
					Type: pulumi.StringRef("geoproximity"),
					GeoProximityLocations: []route53.GetTrafficPolicyDocumentRuleGeoProximityLocation{
						{
							Region:               pulumi.StringRef("aws:route53:us-west-1"),
							Bias:                 pulumi.StringRef("10"),
							EvaluateTargetHealth: pulumi.BoolRef(true),
							RuleReference:        pulumi.StringRef("na_rule"),
						},
						{
							Region:               pulumi.StringRef("aws:route53:eu-west-1"),
							Bias:                 pulumi.StringRef("10"),
							EvaluateTargetHealth: pulumi.BoolRef(true),
							EndpointReference:    pulumi.StringRef("eu_endpoint"),
						},
						{
							Region:               pulumi.StringRef("aws:route53:ap-northeast-2"),
							Bias:                 pulumi.StringRef("0"),
							EvaluateTargetHealth: pulumi.BoolRef(true),
							EndpointReference:    pulumi.StringRef("ap_endpoint"),
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = route53.NewTrafficPolicy(ctx, "exampleTrafficPolicy", &route53.TrafficPolicyArgs{
			Comment:  pulumi.String("example comment"),
			Document: *pulumi.String(exampleTrafficPolicyDocument.Json),
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
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetTrafficPolicyDocumentArgs;
import com.pulumi.aws.route53.TrafficPolicy;
import com.pulumi.aws.route53.TrafficPolicyArgs;
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
        final var exampleTrafficPolicyDocument = Route53Functions.getTrafficPolicyDocument(GetTrafficPolicyDocumentArgs.builder()
            .recordType("A")
            .startRule("geoproximity_rule")
            .endpoints(            
                GetTrafficPolicyDocumentEndpointArgs.builder()
                    .id("na_endpoint_a")
                    .type("elastic-load-balancer")
                    .value("elb-111111.us-west-1.elb.amazonaws.com")
                    .build(),
                GetTrafficPolicyDocumentEndpointArgs.builder()
                    .id("na_endpoint_b")
                    .type("elastic-load-balancer")
                    .value("elb-222222.us-west-1.elb.amazonaws.com")
                    .build(),
                GetTrafficPolicyDocumentEndpointArgs.builder()
                    .id("eu_endpoint")
                    .type("elastic-load-balancer")
                    .value("elb-333333.eu-west-1.elb.amazonaws.com")
                    .build(),
                GetTrafficPolicyDocumentEndpointArgs.builder()
                    .id("ap_endpoint")
                    .type("elastic-load-balancer")
                    .value("elb-444444.ap-northeast-2.elb.amazonaws.com")
                    .build())
            .rules(            
                GetTrafficPolicyDocumentRuleArgs.builder()
                    .id("na_rule")
                    .type("failover")
                    .primary(GetTrafficPolicyDocumentRulePrimaryArgs.builder()
                        .endpointReference("na_endpoint_a")
                        .build())
                    .secondary(GetTrafficPolicyDocumentRuleSecondaryArgs.builder()
                        .endpointReference("na_endpoint_b")
                        .build())
                    .build(),
                GetTrafficPolicyDocumentRuleArgs.builder()
                    .id("geoproximity_rule")
                    .type("geoproximity")
                    .geoProximityLocations(                    
                        GetTrafficPolicyDocumentRuleGeoProximityLocationArgs.builder()
                            .region("aws:route53:us-west-1")
                            .bias(10)
                            .evaluateTargetHealth(true)
                            .ruleReference("na_rule")
                            .build(),
                        GetTrafficPolicyDocumentRuleGeoProximityLocationArgs.builder()
                            .region("aws:route53:eu-west-1")
                            .bias(10)
                            .evaluateTargetHealth(true)
                            .endpointReference("eu_endpoint")
                            .build(),
                        GetTrafficPolicyDocumentRuleGeoProximityLocationArgs.builder()
                            .region("aws:route53:ap-northeast-2")
                            .bias(0)
                            .evaluateTargetHealth(true)
                            .endpointReference("ap_endpoint")
                            .build())
                    .build())
            .build());

        var exampleTrafficPolicy = new TrafficPolicy("exampleTrafficPolicy", TrafficPolicyArgs.builder()        
            .comment("example comment")
            .document(exampleTrafficPolicyDocument.applyValue(getTrafficPolicyDocumentResult -> getTrafficPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  exampleTrafficPolicy:
    type: aws:route53:TrafficPolicy
    properties:
      comment: example comment
      document: ${exampleTrafficPolicyDocument.json}
variables:
  exampleTrafficPolicyDocument:
    fn::invoke:
      Function: aws:route53:getTrafficPolicyDocument
      Arguments:
        recordType: A
        startRule: geoproximity_rule
        endpoints:
          - id: na_endpoint_a
            type: elastic-load-balancer
            value: elb-111111.us-west-1.elb.amazonaws.com
          - id: na_endpoint_b
            type: elastic-load-balancer
            value: elb-222222.us-west-1.elb.amazonaws.com
          - id: eu_endpoint
            type: elastic-load-balancer
            value: elb-333333.eu-west-1.elb.amazonaws.com
          - id: ap_endpoint
            type: elastic-load-balancer
            value: elb-444444.ap-northeast-2.elb.amazonaws.com
        rules:
          - id: na_rule
            type: failover
            primary:
              endpointReference: na_endpoint_a
            secondary:
              endpointReference: na_endpoint_b
          - id: geoproximity_rule
            type: geoproximity
            geoProximityLocations:
              - region: aws:route53:us-west-1
                bias: 10
                evaluateTargetHealth: true
                ruleReference: na_rule
              - region: aws:route53:eu-west-1
                bias: 10
                evaluateTargetHealth: true
                endpointReference: eu_endpoint
              - region: aws:route53:ap-northeast-2
                bias: 0
                evaluateTargetHealth: true
                endpointReference: ap_endpoint
```
{{% /example %}}
{{% /examples %}}