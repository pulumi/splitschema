Provides a VPC/Subnet/ENI/Transit Gateway/Transit Gateway Attachment Flow Log to capture IP traffic for a specific network
interface, subnet, or VPC. Logs are sent to a CloudWatch Log Group, a S3 Bucket, or Amazon Kinesis Data Firehose

{{% examples %}}
## Example Usage
{{% example %}}
### CloudWatch Logging

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["vpc-flow-logs.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const exampleFlowLog = new aws.ec2.FlowLog("exampleFlowLog", {
    iamRoleArn: exampleRole.arn,
    logDestination: exampleLogGroup.arn,
    trafficType: "ALL",
    vpcId: aws_vpc.example.id,
});
const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: [
            "logs:CreateLogGroup",
            "logs:CreateLogStream",
            "logs:PutLogEvents",
            "logs:DescribeLogGroups",
            "logs:DescribeLogStreams",
        ],
        resources: ["*"],
    }],
});
const exampleRolePolicy = new aws.iam.RolePolicy("exampleRolePolicy", {
    role: exampleRole.id,
    policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["vpc-flow-logs.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
example_flow_log = aws.ec2.FlowLog("exampleFlowLog",
    iam_role_arn=example_role.arn,
    log_destination=example_log_group.arn,
    traffic_type="ALL",
    vpc_id=aws_vpc["example"]["id"])
example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=[
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:DescribeLogGroups",
        "logs:DescribeLogStreams",
    ],
    resources=["*"],
)])
example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
    role=example_role.id,
    policy=example_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");

    var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            "vpc-flow-logs.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleFlowLog = new Aws.Ec2.FlowLog("exampleFlowLog", new()
    {
        IamRoleArn = exampleRole.Arn,
        LogDestination = exampleLogGroup.Arn,
        TrafficType = "ALL",
        VpcId = aws_vpc.Example.Id,
    });

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:PutLogEvents",
                    "logs:DescribeLogGroups",
                    "logs:DescribeLogStreams",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var exampleRolePolicy = new Aws.Iam.RolePolicy("exampleRolePolicy", new()
    {
        Role = exampleRole.Id,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
		if err != nil {
			return err
		}
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"vpc-flow-logs.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewFlowLog(ctx, "exampleFlowLog", &ec2.FlowLogArgs{
			IamRoleArn:     exampleRole.Arn,
			LogDestination: exampleLogGroup.Arn,
			TrafficType:    pulumi.String("ALL"),
			VpcId:          pulumi.Any(aws_vpc.Example.Id),
		})
		if err != nil {
			return err
		}
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"logs:CreateLogGroup",
						"logs:CreateLogStream",
						"logs:PutLogEvents",
						"logs:DescribeLogGroups",
						"logs:DescribeLogStreams",
					},
					Resources: []string{
						"*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "exampleRolePolicy", &iam.RolePolicyArgs{
			Role:   exampleRole.ID(),
			Policy: *pulumi.String(examplePolicyDocument.Json),
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.ec2.FlowLog;
import com.pulumi.aws.ec2.FlowLogArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
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

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("vpc-flow-logs.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleFlowLog = new FlowLog("exampleFlowLog", FlowLogArgs.builder()        
            .iamRoleArn(exampleRole.arn())
            .logDestination(exampleLogGroup.arn())
            .trafficType("ALL")
            .vpcId(aws_vpc.example().id())
            .build());

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions(                
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:PutLogEvents",
                    "logs:DescribeLogGroups",
                    "logs:DescribeLogStreams")
                .resources("*")
                .build())
            .build());

        var exampleRolePolicy = new RolePolicy("exampleRolePolicy", RolePolicyArgs.builder()        
            .role(exampleRole.id())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  exampleFlowLog:
    type: aws:ec2:FlowLog
    properties:
      iamRoleArn: ${exampleRole.arn}
      logDestination: ${exampleLogGroup.arn}
      trafficType: ALL
      vpcId: ${aws_vpc.example.id}
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  exampleRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${exampleRole.id}
      policy: ${examplePolicyDocument.json}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - vpc-flow-logs.amazonaws.com
            actions:
              - sts:AssumeRole
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - logs:CreateLogGroup
              - logs:CreateLogStream
              - logs:PutLogEvents
              - logs:DescribeLogGroups
              - logs:DescribeLogStreams
            resources:
              - '*'
```
{{% /example %}}
{{% example %}}
### Amazon Kinesis Data Firehose logging

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs;
import com.pulumi.aws.ec2.FlowLog;
import com.pulumi.aws.ec2.FlowLogArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("firehose.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleFirehoseDeliveryStream = new FirehoseDeliveryStream("exampleFirehoseDeliveryStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("extended_s3")
            .extendedS3Configuration(FirehoseDeliveryStreamExtendedS3ConfigurationArgs.builder()
                .roleArn(exampleRole.arn())
                .bucketArn(exampleBucketV2.arn())
                .build())
            .tags(Map.of("LogDeliveryEnabled", "true"))
            .build());

        var exampleFlowLog = new FlowLog("exampleFlowLog", FlowLogArgs.builder()        
            .logDestination(exampleFirehoseDeliveryStream.arn())
            .logDestinationType("kinesis-data-firehose")
            .trafficType("ALL")
            .vpcId(aws_vpc.example().id())
            .build());

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .acl("private")
            .build());

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .effect("Allow")
            .actions(            
                "logs:CreateLogDelivery",
                "logs:DeleteLogDelivery",
                "logs:ListLogDeliveries",
                "logs:GetLogDelivery",
                "firehose:TagDeliveryStream")
            .resources("*")
            .build());

        var exampleRolePolicy = new RolePolicy("exampleRolePolicy", RolePolicyArgs.builder()        
            .role(exampleRole.id())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  exampleFlowLog:
    type: aws:ec2:FlowLog
    properties:
      logDestination: ${exampleFirehoseDeliveryStream.arn}
      logDestinationType: kinesis-data-firehose
      trafficType: ALL
      vpcId: ${aws_vpc.example.id}
  exampleFirehoseDeliveryStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: extended_s3
      extendedS3Configuration:
        roleArn: ${exampleRole.arn}
        bucketArn: ${exampleBucketV2.arn}
      tags:
        LogDeliveryEnabled: 'true'
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${exampleBucketV2.id}
      acl: private
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  exampleRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${exampleRole.id}
      policy: ${examplePolicyDocument.json}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - firehose.amazonaws.com
            actions:
              - sts:AssumeRole
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        effect: Allow
        actions:
          - logs:CreateLogDelivery
          - logs:DeleteLogDelivery
          - logs:ListLogDeliveries
          - logs:GetLogDelivery
          - firehose:TagDeliveryStream
        resources:
          - '*'
```
{{% /example %}}
{{% example %}}
### S3 Logging

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleFlowLog = new aws.ec2.FlowLog("exampleFlowLog", {
    logDestination: exampleBucketV2.arn,
    logDestinationType: "s3",
    trafficType: "ALL",
    vpcId: aws_vpc.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_flow_log = aws.ec2.FlowLog("exampleFlowLog",
    log_destination=example_bucket_v2.arn,
    log_destination_type="s3",
    traffic_type="ALL",
    vpc_id=aws_vpc["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleFlowLog = new Aws.Ec2.FlowLog("exampleFlowLog", new()
    {
        LogDestination = exampleBucketV2.Arn,
        LogDestinationType = "s3",
        TrafficType = "ALL",
        VpcId = aws_vpc.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewFlowLog(ctx, "exampleFlowLog", &ec2.FlowLogArgs{
			LogDestination:     exampleBucketV2.Arn,
			LogDestinationType: pulumi.String("s3"),
			TrafficType:        pulumi.String("ALL"),
			VpcId:              pulumi.Any(aws_vpc.Example.Id),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.ec2.FlowLog;
import com.pulumi.aws.ec2.FlowLogArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleFlowLog = new FlowLog("exampleFlowLog", FlowLogArgs.builder()        
            .logDestination(exampleBucketV2.arn())
            .logDestinationType("s3")
            .trafficType("ALL")
            .vpcId(aws_vpc.example().id())
            .build());

    }
}
```
```yaml
resources:
  exampleFlowLog:
    type: aws:ec2:FlowLog
    properties:
      logDestination: ${exampleBucketV2.arn}
      logDestinationType: s3
      trafficType: ALL
      vpcId: ${aws_vpc.example.id}
  exampleBucketV2:
    type: aws:s3:BucketV2
```
{{% /example %}}
{{% example %}}
### S3 Logging in Apache Parquet format with per-hour partitions

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleFlowLog = new aws.ec2.FlowLog("exampleFlowLog", {
    logDestination: exampleBucketV2.arn,
    logDestinationType: "s3",
    trafficType: "ALL",
    vpcId: aws_vpc.example.id,
    destinationOptions: {
        fileFormat: "parquet",
        perHourPartition: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_flow_log = aws.ec2.FlowLog("exampleFlowLog",
    log_destination=example_bucket_v2.arn,
    log_destination_type="s3",
    traffic_type="ALL",
    vpc_id=aws_vpc["example"]["id"],
    destination_options=aws.ec2.FlowLogDestinationOptionsArgs(
        file_format="parquet",
        per_hour_partition=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleFlowLog = new Aws.Ec2.FlowLog("exampleFlowLog", new()
    {
        LogDestination = exampleBucketV2.Arn,
        LogDestinationType = "s3",
        TrafficType = "ALL",
        VpcId = aws_vpc.Example.Id,
        DestinationOptions = new Aws.Ec2.Inputs.FlowLogDestinationOptionsArgs
        {
            FileFormat = "parquet",
            PerHourPartition = true,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewFlowLog(ctx, "exampleFlowLog", &ec2.FlowLogArgs{
			LogDestination:     exampleBucketV2.Arn,
			LogDestinationType: pulumi.String("s3"),
			TrafficType:        pulumi.String("ALL"),
			VpcId:              pulumi.Any(aws_vpc.Example.Id),
			DestinationOptions: &ec2.FlowLogDestinationOptionsArgs{
				FileFormat:       pulumi.String("parquet"),
				PerHourPartition: pulumi.Bool(true),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.ec2.FlowLog;
import com.pulumi.aws.ec2.FlowLogArgs;
import com.pulumi.aws.ec2.inputs.FlowLogDestinationOptionsArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleFlowLog = new FlowLog("exampleFlowLog", FlowLogArgs.builder()        
            .logDestination(exampleBucketV2.arn())
            .logDestinationType("s3")
            .trafficType("ALL")
            .vpcId(aws_vpc.example().id())
            .destinationOptions(FlowLogDestinationOptionsArgs.builder()
                .fileFormat("parquet")
                .perHourPartition(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleFlowLog:
    type: aws:ec2:FlowLog
    properties:
      logDestination: ${exampleBucketV2.arn}
      logDestinationType: s3
      trafficType: ALL
      vpcId: ${aws_vpc.example.id}
      destinationOptions:
        fileFormat: parquet
        perHourPartition: true
  exampleBucketV2:
    type: aws:s3:BucketV2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Flow Logs using the `id`. For example:

```sh
 $ pulumi import aws:ec2/flowLog:FlowLog test_flow_log fl-1a2b3c4d
```
 