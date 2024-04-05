Manages an AWS Elasticsearch Domain.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticsearch.Domain("example", {
    clusterConfig: {
        instanceType: "r4.large.elasticsearch",
    },
    elasticsearchVersion: "7.10",
    tags: {
        Domain: "TestDomain",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.elasticsearch.Domain("example",
    cluster_config=aws.elasticsearch.DomainClusterConfigArgs(
        instance_type="r4.large.elasticsearch",
    ),
    elasticsearch_version="7.10",
    tags={
        "Domain": "TestDomain",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ElasticSearch.Domain("example", new()
    {
        ClusterConfig = new Aws.ElasticSearch.Inputs.DomainClusterConfigArgs
        {
            InstanceType = "r4.large.elasticsearch",
        },
        ElasticsearchVersion = "7.10",
        Tags = 
        {
            { "Domain", "TestDomain" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticsearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticsearch.NewDomain(ctx, "example", &elasticsearch.DomainArgs{
			ClusterConfig: &elasticsearch.DomainClusterConfigArgs{
				InstanceType: pulumi.String("r4.large.elasticsearch"),
			},
			ElasticsearchVersion: pulumi.String("7.10"),
			Tags: pulumi.StringMap{
				"Domain": pulumi.String("TestDomain"),
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
import com.pulumi.aws.elasticsearch.Domain;
import com.pulumi.aws.elasticsearch.DomainArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainClusterConfigArgs;
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
        var example = new Domain("example", DomainArgs.builder()        
            .clusterConfig(DomainClusterConfigArgs.builder()
                .instanceType("r4.large.elasticsearch")
                .build())
            .elasticsearchVersion("7.10")
            .tags(Map.of("Domain", "TestDomain"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:elasticsearch:Domain
    properties:
      clusterConfig:
        instanceType: r4.large.elasticsearch
      elasticsearchVersion: '7.10'
      tags:
        Domain: TestDomain
```
{{% /example %}}
{{% example %}}
### Access Policy

> See also: `aws.elasticsearch.DomainPolicy` resource

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const domain = config.get("domain") || "tf-test";
const currentRegion = aws.getRegion({});
const currentCallerIdentity = aws.getCallerIdentity({});
const example = new aws.elasticsearch.Domain("example", {accessPolicies: Promise.all([currentRegion, currentCallerIdentity]).then(([currentRegion, currentCallerIdentity]) => `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "es:*",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "arn:aws:es:${currentRegion.name}:${currentCallerIdentity.accountId}:domain/${domain}/*",
      "Condition": {
        "IpAddress": {"aws:SourceIp": ["66.193.100.22/32"]}
      }
    }
  ]
}
`)});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
domain = config.get("domain")
if domain is None:
    domain = "tf-test"
current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
example = aws.elasticsearch.Domain("example", access_policies=f"""{{
  "Version": "2012-10-17",
  "Statement": [
    {{
      "Action": "es:*",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "arn:aws:es:{current_region.name}:{current_caller_identity.account_id}:domain/{domain}/*",
      "Condition": {{
        "IpAddress": {{"aws:SourceIp": ["66.193.100.22/32"]}}
      }}
    }}
  ]
}}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var domain = config.Get("domain") ?? "tf-test";
    var currentRegion = Aws.GetRegion.Invoke();

    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var example = new Aws.ElasticSearch.Domain("example", new()
    {
        AccessPolicies = Output.Tuple(currentRegion, currentCallerIdentity).Apply(values =>
        {
            var currentRegion = values.Item1;
            var currentCallerIdentity = values.Item2;
            return @$"{{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {{
      ""Action"": ""es:*"",
      ""Principal"": ""*"",
      ""Effect"": ""Allow"",
      ""Resource"": ""arn:aws:es:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:domain/{domain}/*"",
      ""Condition"": {{
        ""IpAddress"": {{""aws:SourceIp"": [""66.193.100.22/32""]}}
      }}
    }}
  ]
}}
";
        }),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticsearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		domain := "tf-test"
		if param := cfg.Get("domain"); param != "" {
			domain = param
		}
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = elasticsearch.NewDomain(ctx, "example", &elasticsearch.DomainArgs{
			AccessPolicies: pulumi.Any(fmt.Sprintf(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "es:*",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "arn:aws:es:%v:%v:domain/%v/*",
      "Condition": {
        "IpAddress": {"aws:SourceIp": ["66.193.100.22/32"]}
      }
    }
  ]
}
`, currentRegion.Name, currentCallerIdentity.AccountId, domain)),
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
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.elasticsearch.Domain;
import com.pulumi.aws.elasticsearch.DomainArgs;
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
        final var config = ctx.config();
        final var domain = config.get("domain").orElse("tf-test");
        final var currentRegion = AwsFunctions.getRegion();

        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        var example = new Domain("example", DomainArgs.builder()        
            .accessPolicies("""
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "es:*",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "arn:aws:es:%s:%s:domain/%s/*",
      "Condition": {
        "IpAddress": {"aws:SourceIp": ["66.193.100.22/32"]}
      }
    }
  ]
}
", currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),domain))
            .build());

    }
}
```
```yaml
configuration:
  domain:
    type: string
    default: tf-test
resources:
  example:
    type: aws:elasticsearch:Domain
    properties:
      accessPolicies: |
        {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "es:*",
              "Principal": "*",
              "Effect": "Allow",
              "Resource": "arn:aws:es:${currentRegion.name}:${currentCallerIdentity.accountId}:domain/${domain}/*",
              "Condition": {
                "IpAddress": {"aws:SourceIp": ["66.193.100.22/32"]}
              }
            }
          ]
        }
variables:
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### Log Publishing to CloudWatch Logs

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["es.amazonaws.com"],
        }],
        actions: [
            "logs:PutLogEvents",
            "logs:PutLogEventsBatch",
            "logs:CreateLogStream",
        ],
        resources: ["arn:aws:logs:*"],
    }],
});
const exampleLogResourcePolicy = new aws.cloudwatch.LogResourcePolicy("exampleLogResourcePolicy", {
    policyName: "example",
    policyDocument: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
// .. other configuration ...
const exampleDomain = new aws.elasticsearch.Domain("exampleDomain", {logPublishingOptions: [{
    cloudwatchLogGroupArn: exampleLogGroup.arn,
    logType: "INDEX_SLOW_LOGS",
}]});
```
```python
import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["es.amazonaws.com"],
    )],
    actions=[
        "logs:PutLogEvents",
        "logs:PutLogEventsBatch",
        "logs:CreateLogStream",
    ],
    resources=["arn:aws:logs:*"],
)])
example_log_resource_policy = aws.cloudwatch.LogResourcePolicy("exampleLogResourcePolicy",
    policy_name="example",
    policy_document=example_policy_document.json)
# .. other configuration ...
example_domain = aws.elasticsearch.Domain("exampleDomain", log_publishing_options=[aws.elasticsearch.DomainLogPublishingOptionArgs(
    cloudwatch_log_group_arn=example_log_group.arn,
    log_type="INDEX_SLOW_LOGS",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "es.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "logs:PutLogEvents",
                    "logs:PutLogEventsBatch",
                    "logs:CreateLogStream",
                },
                Resources = new[]
                {
                    "arn:aws:logs:*",
                },
            },
        },
    });

    var exampleLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("exampleLogResourcePolicy", new()
    {
        PolicyName = "example",
        PolicyDocument = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    // .. other configuration ...
    var exampleDomain = new Aws.ElasticSearch.Domain("exampleDomain", new()
    {
        LogPublishingOptions = new[]
        {
            new Aws.ElasticSearch.Inputs.DomainLogPublishingOptionArgs
            {
                CloudwatchLogGroupArn = exampleLogGroup.Arn,
                LogType = "INDEX_SLOW_LOGS",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticsearch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
		if err != nil {
			return err
		}
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"es.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"logs:PutLogEvents",
						"logs:PutLogEventsBatch",
						"logs:CreateLogStream",
					},
					Resources: []string{
						"arn:aws:logs:*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogResourcePolicy(ctx, "exampleLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
			PolicyName:     pulumi.String("example"),
			PolicyDocument: *pulumi.String(examplePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = elasticsearch.NewDomain(ctx, "exampleDomain", &elasticsearch.DomainArgs{
			LogPublishingOptions: elasticsearch.DomainLogPublishingOptionArray{
				&elasticsearch.DomainLogPublishingOptionArgs{
					CloudwatchLogGroupArn: exampleLogGroup.Arn,
					LogType:               pulumi.String("INDEX_SLOW_LOGS"),
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.cloudwatch.LogResourcePolicy;
import com.pulumi.aws.cloudwatch.LogResourcePolicyArgs;
import com.pulumi.aws.elasticsearch.Domain;
import com.pulumi.aws.elasticsearch.DomainArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainLogPublishingOptionArgs;
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
        var exampleLogGroup = new LogGroup("exampleLogGroup");

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("es.amazonaws.com")
                    .build())
                .actions(                
                    "logs:PutLogEvents",
                    "logs:PutLogEventsBatch",
                    "logs:CreateLogStream")
                .resources("arn:aws:logs:*")
                .build())
            .build());

        var exampleLogResourcePolicy = new LogResourcePolicy("exampleLogResourcePolicy", LogResourcePolicyArgs.builder()        
            .policyName("example")
            .policyDocument(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleDomain = new Domain("exampleDomain", DomainArgs.builder()        
            .logPublishingOptions(DomainLogPublishingOptionArgs.builder()
                .cloudwatchLogGroupArn(exampleLogGroup.arn())
                .logType("INDEX_SLOW_LOGS")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
  exampleLogResourcePolicy:
    type: aws:cloudwatch:LogResourcePolicy
    properties:
      policyName: example
      policyDocument: ${examplePolicyDocument.json}
  exampleDomain:
    type: aws:elasticsearch:Domain
    properties:
      logPublishingOptions:
        - cloudwatchLogGroupArn: ${exampleLogGroup.arn}
          logType: INDEX_SLOW_LOGS
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - es.amazonaws.com
            actions:
              - logs:PutLogEvents
              - logs:PutLogEventsBatch
              - logs:CreateLogStream
            resources:
              - arn:aws:logs:*
```
{{% /example %}}
{{% example %}}
### VPC based ES

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const vpc = config.requireObject("vpc");
const domain = config.get("domain") || "tf-test";
const selectedVpc = aws.ec2.getVpc({
    tags: {
        Name: vpc,
    },
});
const selectedSubnets = selectedVpc.then(selectedVpc => aws.ec2.getSubnets({
    filters: [{
        name: "vpc-id",
        values: [selectedVpc.id],
    }],
    tags: {
        Tier: "private",
    },
}));
const currentRegion = aws.getRegion({});
const currentCallerIdentity = aws.getCallerIdentity({});
const esSecurityGroup = new aws.ec2.SecurityGroup("esSecurityGroup", {
    description: "Managed by Pulumi",
    vpcId: selectedVpc.then(selectedVpc => selectedVpc.id),
    ingress: [{
        fromPort: 443,
        toPort: 443,
        protocol: "tcp",
        cidrBlocks: [selectedVpc.then(selectedVpc => selectedVpc.cidrBlock)],
    }],
});
const esServiceLinkedRole = new aws.iam.ServiceLinkedRole("esServiceLinkedRole", {awsServiceName: "opensearchservice.amazonaws.com"});
const esDomain = new aws.elasticsearch.Domain("esDomain", {
    elasticsearchVersion: "6.3",
    clusterConfig: {
        instanceType: "m4.large.elasticsearch",
        zoneAwarenessEnabled: true,
    },
    vpcOptions: {
        subnetIds: [
            selectedSubnets.then(selectedSubnets => selectedSubnets.ids?.[0]),
            selectedSubnets.then(selectedSubnets => selectedSubnets.ids?.[1]),
        ],
        securityGroupIds: [esSecurityGroup.id],
    },
    advancedOptions: {
        "rest.action.multi.allow_explicit_index": "true",
    },
    accessPolicies: Promise.all([currentRegion, currentCallerIdentity]).then(([currentRegion, currentCallerIdentity]) => `{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Action": "es:*",
			"Principal": "*",
			"Effect": "Allow",
			"Resource": "arn:aws:es:${currentRegion.name}:${currentCallerIdentity.accountId}:domain/${domain}/*"
		}
	]
}
`),
    tags: {
        Domain: "TestDomain",
    },
}, {
    dependsOn: [esServiceLinkedRole],
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
vpc = config.require_object("vpc")
domain = config.get("domain")
if domain is None:
    domain = "tf-test"
selected_vpc = aws.ec2.get_vpc(tags={
    "Name": vpc,
})
selected_subnets = aws.ec2.get_subnets(filters=[aws.ec2.GetSubnetsFilterArgs(
        name="vpc-id",
        values=[selected_vpc.id],
    )],
    tags={
        "Tier": "private",
    })
current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
es_security_group = aws.ec2.SecurityGroup("esSecurityGroup",
    description="Managed by Pulumi",
    vpc_id=selected_vpc.id,
    ingress=[aws.ec2.SecurityGroupIngressArgs(
        from_port=443,
        to_port=443,
        protocol="tcp",
        cidr_blocks=[selected_vpc.cidr_block],
    )])
es_service_linked_role = aws.iam.ServiceLinkedRole("esServiceLinkedRole", aws_service_name="opensearchservice.amazonaws.com")
es_domain = aws.elasticsearch.Domain("esDomain",
    elasticsearch_version="6.3",
    cluster_config=aws.elasticsearch.DomainClusterConfigArgs(
        instance_type="m4.large.elasticsearch",
        zone_awareness_enabled=True,
    ),
    vpc_options=aws.elasticsearch.DomainVpcOptionsArgs(
        subnet_ids=[
            selected_subnets.ids[0],
            selected_subnets.ids[1],
        ],
        security_group_ids=[es_security_group.id],
    ),
    advanced_options={
        "rest.action.multi.allow_explicit_index": "true",
    },
    access_policies=f"""{{
	"Version": "2012-10-17",
	"Statement": [
		{{
			"Action": "es:*",
			"Principal": "*",
			"Effect": "Allow",
			"Resource": "arn:aws:es:{current_region.name}:{current_caller_identity.account_id}:domain/{domain}/*"
		}}
	]
}}
""",
    tags={
        "Domain": "TestDomain",
    },
    opts=pulumi.ResourceOptions(depends_on=[es_service_linked_role]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var vpc = config.RequireObject<dynamic>("vpc");
    var domain = config.Get("domain") ?? "tf-test";
    var selectedVpc = Aws.Ec2.GetVpc.Invoke(new()
    {
        Tags = 
        {
            { "Name", vpc },
        },
    });

    var selectedSubnets = Aws.Ec2.GetSubnets.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetSubnetsFilterInputArgs
            {
                Name = "vpc-id",
                Values = new[]
                {
                    selectedVpc.Apply(getVpcResult => getVpcResult.Id),
                },
            },
        },
        Tags = 
        {
            { "Tier", "private" },
        },
    });

    var currentRegion = Aws.GetRegion.Invoke();

    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var esSecurityGroup = new Aws.Ec2.SecurityGroup("esSecurityGroup", new()
    {
        Description = "Managed by Pulumi",
        VpcId = selectedVpc.Apply(getVpcResult => getVpcResult.Id),
        Ingress = new[]
        {
            new Aws.Ec2.Inputs.SecurityGroupIngressArgs
            {
                FromPort = 443,
                ToPort = 443,
                Protocol = "tcp",
                CidrBlocks = new[]
                {
                    selectedVpc.Apply(getVpcResult => getVpcResult.CidrBlock),
                },
            },
        },
    });

    var esServiceLinkedRole = new Aws.Iam.ServiceLinkedRole("esServiceLinkedRole", new()
    {
        AwsServiceName = "opensearchservice.amazonaws.com",
    });

    var esDomain = new Aws.ElasticSearch.Domain("esDomain", new()
    {
        ElasticsearchVersion = "6.3",
        ClusterConfig = new Aws.ElasticSearch.Inputs.DomainClusterConfigArgs
        {
            InstanceType = "m4.large.elasticsearch",
            ZoneAwarenessEnabled = true,
        },
        VpcOptions = new Aws.ElasticSearch.Inputs.DomainVpcOptionsArgs
        {
            SubnetIds = new[]
            {
                selectedSubnets.Apply(getSubnetsResult => getSubnetsResult.Ids[0]),
                selectedSubnets.Apply(getSubnetsResult => getSubnetsResult.Ids[1]),
            },
            SecurityGroupIds = new[]
            {
                esSecurityGroup.Id,
            },
        },
        AdvancedOptions = 
        {
            { "rest.action.multi.allow_explicit_index", "true" },
        },
        AccessPolicies = Output.Tuple(currentRegion, currentCallerIdentity).Apply(values =>
        {
            var currentRegion = values.Item1;
            var currentCallerIdentity = values.Item2;
            return @$"{{
	""Version"": ""2012-10-17"",
	""Statement"": [
		{{
			""Action"": ""es:*"",
			""Principal"": ""*"",
			""Effect"": ""Allow"",
			""Resource"": ""arn:aws:es:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:domain/{domain}/*""
		}}
	]
}}
";
        }),
        Tags = 
        {
            { "Domain", "TestDomain" },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            esServiceLinkedRole,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticsearch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
cfg := config.New(ctx, "")
vpc := cfg.RequireObject("vpc")
domain := "tf-test";
if param := cfg.Get("domain"); param != ""{
domain = param
}
selectedVpc, err := ec2.LookupVpc(ctx, &ec2.LookupVpcArgs{
Tags: interface{}{
Name: vpc,
},
}, nil);
if err != nil {
return err
}
selectedSubnets, err := ec2.GetSubnets(ctx, &ec2.GetSubnetsArgs{
Filters: []ec2.GetSubnetsFilter{
{
Name: "vpc-id",
Values: interface{}{
selectedVpc.Id,
},
},
},
Tags: map[string]interface{}{
"Tier": "private",
},
}, nil);
if err != nil {
return err
}
currentRegion, err := aws.GetRegion(ctx, nil, nil);
if err != nil {
return err
}
currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil);
if err != nil {
return err
}
esSecurityGroup, err := ec2.NewSecurityGroup(ctx, "esSecurityGroup", &ec2.SecurityGroupArgs{
Description: pulumi.String("Managed by Pulumi"),
VpcId: *pulumi.String(selectedVpc.Id),
Ingress: ec2.SecurityGroupIngressArray{
&ec2.SecurityGroupIngressArgs{
FromPort: pulumi.Int(443),
ToPort: pulumi.Int(443),
Protocol: pulumi.String("tcp"),
CidrBlocks: pulumi.StringArray{
*pulumi.String(selectedVpc.CidrBlock),
},
},
},
})
if err != nil {
return err
}
esServiceLinkedRole, err := iam.NewServiceLinkedRole(ctx, "esServiceLinkedRole", &iam.ServiceLinkedRoleArgs{
AwsServiceName: pulumi.String("opensearchservice.amazonaws.com"),
})
if err != nil {
return err
}
_, err = elasticsearch.NewDomain(ctx, "esDomain", &elasticsearch.DomainArgs{
ElasticsearchVersion: pulumi.String("6.3"),
ClusterConfig: &elasticsearch.DomainClusterConfigArgs{
InstanceType: pulumi.String("m4.large.elasticsearch"),
ZoneAwarenessEnabled: pulumi.Bool(true),
},
VpcOptions: &elasticsearch.DomainVpcOptionsArgs{
SubnetIds: pulumi.StringArray{
*pulumi.String(selectedSubnets.Ids[0]),
*pulumi.String(selectedSubnets.Ids[1]),
},
SecurityGroupIds: pulumi.StringArray{
esSecurityGroup.ID(),
},
},
AdvancedOptions: pulumi.StringMap{
"rest.action.multi.allow_explicit_index": pulumi.String("true"),
},
AccessPolicies: pulumi.Any(fmt.Sprintf(`{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Action": "es:*",
			"Principal": "*",
			"Effect": "Allow",
			"Resource": "arn:aws:es:%v:%v:domain/%v/*"
		}
	]
}
`, currentRegion.Name, currentCallerIdentity.AccountId, domain)),
Tags: pulumi.StringMap{
"Domain": pulumi.String("TestDomain"),
},
}, pulumi.DependsOn([]pulumi.Resource{
esServiceLinkedRole,
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetVpcArgs;
import com.pulumi.aws.ec2.inputs.GetSubnetsArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
import com.pulumi.aws.ec2.inputs.SecurityGroupIngressArgs;
import com.pulumi.aws.iam.ServiceLinkedRole;
import com.pulumi.aws.iam.ServiceLinkedRoleArgs;
import com.pulumi.aws.elasticsearch.Domain;
import com.pulumi.aws.elasticsearch.DomainArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainClusterConfigArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainVpcOptionsArgs;
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
        final var config = ctx.config();
        final var vpc = config.get("vpc");
        final var domain = config.get("domain").orElse("tf-test");
        final var selectedVpc = Ec2Functions.getVpc(GetVpcArgs.builder()
            .tags(Map.of("Name", vpc))
            .build());

        final var selectedSubnets = Ec2Functions.getSubnets(GetSubnetsArgs.builder()
            .filters(GetSubnetsFilterArgs.builder()
                .name("vpc-id")
                .values(selectedVpc.applyValue(getVpcResult -> getVpcResult.id()))
                .build())
            .tags(Map.of("Tier", "private"))
            .build());

        final var currentRegion = AwsFunctions.getRegion();

        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        var esSecurityGroup = new SecurityGroup("esSecurityGroup", SecurityGroupArgs.builder()        
            .description("Managed by Pulumi")
            .vpcId(selectedVpc.applyValue(getVpcResult -> getVpcResult.id()))
            .ingress(SecurityGroupIngressArgs.builder()
                .fromPort(443)
                .toPort(443)
                .protocol("tcp")
                .cidrBlocks(selectedVpc.applyValue(getVpcResult -> getVpcResult.cidrBlock()))
                .build())
            .build());

        var esServiceLinkedRole = new ServiceLinkedRole("esServiceLinkedRole", ServiceLinkedRoleArgs.builder()        
            .awsServiceName("opensearchservice.amazonaws.com")
            .build());

        var esDomain = new Domain("esDomain", DomainArgs.builder()        
            .elasticsearchVersion("6.3")
            .clusterConfig(DomainClusterConfigArgs.builder()
                .instanceType("m4.large.elasticsearch")
                .zoneAwarenessEnabled(true)
                .build())
            .vpcOptions(DomainVpcOptionsArgs.builder()
                .subnetIds(                
                    selectedSubnets.applyValue(getSubnetsResult -> getSubnetsResult.ids()[0]),
                    selectedSubnets.applyValue(getSubnetsResult -> getSubnetsResult.ids()[1]))
                .securityGroupIds(esSecurityGroup.id())
                .build())
            .advancedOptions(Map.of("rest.action.multi.allow_explicit_index", "true"))
            .accessPolicies("""
{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Action": "es:*",
			"Principal": "*",
			"Effect": "Allow",
			"Resource": "arn:aws:es:%s:%s:domain/%s/*"
		}
	]
}
", currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),domain))
            .tags(Map.of("Domain", "TestDomain"))
            .build(), CustomResourceOptions.builder()
                .dependsOn(esServiceLinkedRole)
                .build());

    }
}
```
```yaml
configuration:
  vpc:
    type: dynamic
  domain:
    type: string
    default: tf-test
resources:
  esSecurityGroup:
    type: aws:ec2:SecurityGroup
    properties:
      description: Managed by Pulumi
      vpcId: ${selectedVpc.id}
      ingress:
        - fromPort: 443
          toPort: 443
          protocol: tcp
          cidrBlocks:
            - ${selectedVpc.cidrBlock}
  esServiceLinkedRole:
    type: aws:iam:ServiceLinkedRole
    properties:
      awsServiceName: opensearchservice.amazonaws.com
  esDomain:
    type: aws:elasticsearch:Domain
    properties:
      elasticsearchVersion: '6.3'
      clusterConfig:
        instanceType: m4.large.elasticsearch
        zoneAwarenessEnabled: true
      vpcOptions:
        subnetIds:
          - ${selectedSubnets.ids[0]}
          - ${selectedSubnets.ids[1]}
        securityGroupIds:
          - ${esSecurityGroup.id}
      advancedOptions:
        rest.action.multi.allow_explicit_index: 'true'
      accessPolicies: |
        {
        	"Version": "2012-10-17",
        	"Statement": [
        		{
        			"Action": "es:*",
        			"Principal": "*",
        			"Effect": "Allow",
        			"Resource": "arn:aws:es:${currentRegion.name}:${currentCallerIdentity.accountId}:domain/${domain}/*"
        		}
        	]
        }
      tags:
        Domain: TestDomain
    options:
      dependson:
        - ${esServiceLinkedRole}
variables:
  selectedVpc:
    fn::invoke:
      Function: aws:ec2:getVpc
      Arguments:
        tags:
          Name: ${vpc}
  selectedSubnets:
    fn::invoke:
      Function: aws:ec2:getSubnets
      Arguments:
        filters:
          - name: vpc-id
            values:
              - ${selectedVpc.id}
        tags:
          Tier: private
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Elasticsearch domains using the `domain_name`. For example:

```sh
 $ pulumi import aws:elasticsearch/domain:Domain example domain_name
```
 