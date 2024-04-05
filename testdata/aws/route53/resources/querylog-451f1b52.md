Provides a Route53 query logging configuration resource.

> **NOTE:** There are restrictions on the configuration of query logging. Notably,
the CloudWatch log group must be in the `us-east-1` region,
a permissive CloudWatch log resource policy must be in place, and
the Route53 hosted zone must be public.
See [Configuring Logging for DNS Queries](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html?console_help=true#query-logs-configuring) for additional details.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Example CloudWatch log group in us-east-1
const us_east_1 = new aws.Provider("us-east-1", {region: "us-east-1"});
const awsRoute53ExampleCom = new aws.cloudwatch.LogGroup("awsRoute53ExampleCom", {retentionInDays: 30}, {
    provider: aws["us-east-1"],
});
// Example CloudWatch log resource policy to allow Route53 to write logs
// to any log group under /aws/route53/*
const route53-query-logging-policyPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        actions: [
            "logs:CreateLogStream",
            "logs:PutLogEvents",
        ],
        resources: ["arn:aws:logs:*:*:log-group:/aws/route53/*"],
        principals: [{
            identifiers: ["route53.amazonaws.com"],
            type: "Service",
        }],
    }],
});
const route53_query_logging_policyLogResourcePolicy = new aws.cloudwatch.LogResourcePolicy("route53-query-logging-policyLogResourcePolicy", {
    policyDocument: route53_query_logging_policyPolicyDocument.then(route53_query_logging_policyPolicyDocument => route53_query_logging_policyPolicyDocument.json),
    policyName: "route53-query-logging-policy",
}, {
    provider: aws["us-east-1"],
});
// Example Route53 zone with query logging
const exampleComZone = new aws.route53.Zone("exampleComZone", {});
const exampleComQueryLog = new aws.route53.QueryLog("exampleComQueryLog", {
    cloudwatchLogGroupArn: awsRoute53ExampleCom.arn,
    zoneId: exampleComZone.zoneId,
}, {
    dependsOn: [route53_query_logging_policyLogResourcePolicy],
});
```
```python
import pulumi
import pulumi_aws as aws

# Example CloudWatch log group in us-east-1
us_east_1 = aws.Provider("us-east-1", region="us-east-1")
aws_route53_example_com = aws.cloudwatch.LogGroup("awsRoute53ExampleCom", retention_in_days=30,
opts=pulumi.ResourceOptions(provider=aws["us-east-1"]))
# Example CloudWatch log resource policy to allow Route53 to write logs
# to any log group under /aws/route53/*
route53_query_logging_policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=[
        "logs:CreateLogStream",
        "logs:PutLogEvents",
    ],
    resources=["arn:aws:logs:*:*:log-group:/aws/route53/*"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        identifiers=["route53.amazonaws.com"],
        type="Service",
    )],
)])
route53_query_logging_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("route53-query-logging-policyLogResourcePolicy",
    policy_document=route53_query_logging_policy_policy_document.json,
    policy_name="route53-query-logging-policy",
    opts=pulumi.ResourceOptions(provider=aws["us-east-1"]))
# Example Route53 zone with query logging
example_com_zone = aws.route53.Zone("exampleComZone")
example_com_query_log = aws.route53.QueryLog("exampleComQueryLog",
    cloudwatch_log_group_arn=aws_route53_example_com.arn,
    zone_id=example_com_zone.zone_id,
    opts=pulumi.ResourceOptions(depends_on=[route53_query_logging_policy_log_resource_policy]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Example CloudWatch log group in us-east-1
    var us_east_1 = new Aws.Provider("us-east-1", new()
    {
        Region = "us-east-1",
    });

    var awsRoute53ExampleCom = new Aws.CloudWatch.LogGroup("awsRoute53ExampleCom", new()
    {
        RetentionInDays = 30,
    }, new CustomResourceOptions
    {
        Provider = aws.Us_east_1,
    });

    // Example CloudWatch log resource policy to allow Route53 to write logs
    // to any log group under /aws/route53/*
    var route53_query_logging_policyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "logs:CreateLogStream",
                    "logs:PutLogEvents",
                },
                Resources = new[]
                {
                    "arn:aws:logs:*:*:log-group:/aws/route53/*",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Identifiers = new[]
                        {
                            "route53.amazonaws.com",
                        },
                        Type = "Service",
                    },
                },
            },
        },
    });

    var route53_query_logging_policyLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("route53-query-logging-policyLogResourcePolicy", new()
    {
        PolicyDocument = route53_query_logging_policyPolicyDocument.Apply(route53_query_logging_policyPolicyDocument => route53_query_logging_policyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json)),
        PolicyName = "route53-query-logging-policy",
    }, new CustomResourceOptions
    {
        Provider = aws.Us_east_1,
    });

    // Example Route53 zone with query logging
    var exampleComZone = new Aws.Route53.Zone("exampleComZone");

    var exampleComQueryLog = new Aws.Route53.QueryLog("exampleComQueryLog", new()
    {
        CloudwatchLogGroupArn = awsRoute53ExampleCom.Arn,
        ZoneId = exampleComZone.ZoneId,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            route53_query_logging_policyLogResourcePolicy,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "us-east-1", &aws.ProviderArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		awsRoute53ExampleCom, err := cloudwatch.NewLogGroup(ctx, "awsRoute53ExampleCom", &cloudwatch.LogGroupArgs{
			RetentionInDays: pulumi.Int(30),
		}, pulumi.Provider(aws.UsEast1))
		if err != nil {
			return err
		}
		route53_query_logging_policyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"logs:CreateLogStream",
						"logs:PutLogEvents",
					},
					Resources: []string{
						"arn:aws:logs:*:*:log-group:/aws/route53/*",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Identifiers: []string{
								"route53.amazonaws.com",
							},
							Type: "Service",
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogResourcePolicy(ctx, "route53-query-logging-policyLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
			PolicyDocument: *pulumi.String(route53_query_logging_policyPolicyDocument.Json),
			PolicyName:     pulumi.String("route53-query-logging-policy"),
		}, pulumi.Provider(aws.UsEast1))
		if err != nil {
			return err
		}
		exampleComZone, err := route53.NewZone(ctx, "exampleComZone", nil)
		if err != nil {
			return err
		}
		_, err = route53.NewQueryLog(ctx, "exampleComQueryLog", &route53.QueryLogArgs{
			CloudwatchLogGroupArn: awsRoute53ExampleCom.Arn,
			ZoneId:                exampleComZone.ZoneId,
		}, pulumi.DependsOn([]pulumi.Resource{
			route53_query_logging_policyLogResourcePolicy,
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
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.cloudwatch.LogGroupArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.cloudwatch.LogResourcePolicy;
import com.pulumi.aws.cloudwatch.LogResourcePolicyArgs;
import com.pulumi.aws.route53.Zone;
import com.pulumi.aws.route53.QueryLog;
import com.pulumi.aws.route53.QueryLogArgs;
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
        var us_east_1 = new Provider("us-east-1", ProviderArgs.builder()        
            .region("us-east-1")
            .build());

        var awsRoute53ExampleCom = new LogGroup("awsRoute53ExampleCom", LogGroupArgs.builder()        
            .retentionInDays(30)
            .build(), CustomResourceOptions.builder()
                .provider(aws.us-east-1())
                .build());

        final var route53-query-logging-policyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions(                
                    "logs:CreateLogStream",
                    "logs:PutLogEvents")
                .resources("arn:aws:logs:*:*:log-group:/aws/route53/*")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .identifiers("route53.amazonaws.com")
                    .type("Service")
                    .build())
                .build())
            .build());

        var route53_query_logging_policyLogResourcePolicy = new LogResourcePolicy("route53-query-logging-policyLogResourcePolicy", LogResourcePolicyArgs.builder()        
            .policyDocument(route53_query_logging_policyPolicyDocument.json())
            .policyName("route53-query-logging-policy")
            .build(), CustomResourceOptions.builder()
                .provider(aws.us-east-1())
                .build());

        var exampleComZone = new Zone("exampleComZone");

        var exampleComQueryLog = new QueryLog("exampleComQueryLog", QueryLogArgs.builder()        
            .cloudwatchLogGroupArn(awsRoute53ExampleCom.arn())
            .zoneId(exampleComZone.zoneId())
            .build(), CustomResourceOptions.builder()
                .dependsOn(route53_query_logging_policyLogResourcePolicy)
                .build());

    }
}
```
```yaml
resources:
  # Example CloudWatch log group in us-east-1
  us-east-1:
    type: pulumi:providers:aws
    properties:
      region: us-east-1
  awsRoute53ExampleCom: # Example CloudWatch log resource policy to allow Route53 to write logs
  # to any log group under /aws/route53/*
    type: aws:cloudwatch:LogGroup
    properties:
      retentionInDays: 30
    options:
      provider: ${aws"us-east-1"[%!s(MISSING)]}
  route53-query-logging-policyLogResourcePolicy: # Example Route53 zone with query logging
    type: aws:cloudwatch:LogResourcePolicy
    properties:
      policyDocument: ${["route53-query-logging-policyPolicyDocument"].json}
      policyName: route53-query-logging-policy
    options:
      provider: ${aws"us-east-1"[%!s(MISSING)]}
  exampleComZone:
    type: aws:route53:Zone
  exampleComQueryLog:
    type: aws:route53:QueryLog
    properties:
      cloudwatchLogGroupArn: ${awsRoute53ExampleCom.arn}
      zoneId: ${exampleComZone.zoneId}
    options:
      dependson:
        - ${["route53-query-logging-policyLogResourcePolicy"]}
variables:
  route53-query-logging-policyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - logs:CreateLogStream
              - logs:PutLogEvents
            resources:
              - arn:aws:logs:*:*:log-group:/aws/route53/*
            principals:
              - identifiers:
                  - route53.amazonaws.com
                type: Service
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 query logging configurations using their ID. For example:

```sh
 $ pulumi import aws:route53/queryLog:QueryLog example_com xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```
 