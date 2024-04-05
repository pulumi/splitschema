Provides a CloudFront real-time log configuration resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["cloudfront.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: [
            "kinesis:DescribeStreamSummary",
            "kinesis:DescribeStream",
            "kinesis:PutRecord",
            "kinesis:PutRecords",
        ],
        resources: [aws_kinesis_stream.example.arn],
    }],
});
const exampleRolePolicy = new aws.iam.RolePolicy("exampleRolePolicy", {
    role: exampleRole.id,
    policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
const exampleRealtimeLogConfig = new aws.cloudfront.RealtimeLogConfig("exampleRealtimeLogConfig", {
    samplingRate: 75,
    fields: [
        "timestamp",
        "c-ip",
    ],
    endpoint: {
        streamType: "Kinesis",
        kinesisStreamConfig: {
            roleArn: exampleRole.arn,
            streamArn: aws_kinesis_stream.example.arn,
        },
    },
}, {
    dependsOn: [exampleRolePolicy],
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["cloudfront.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=[
        "kinesis:DescribeStreamSummary",
        "kinesis:DescribeStream",
        "kinesis:PutRecord",
        "kinesis:PutRecords",
    ],
    resources=[aws_kinesis_stream["example"]["arn"]],
)])
example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
    role=example_role.id,
    policy=example_policy_document.json)
example_realtime_log_config = aws.cloudfront.RealtimeLogConfig("exampleRealtimeLogConfig",
    sampling_rate=75,
    fields=[
        "timestamp",
        "c-ip",
    ],
    endpoint=aws.cloudfront.RealtimeLogConfigEndpointArgs(
        stream_type="Kinesis",
        kinesis_stream_config=aws.cloudfront.RealtimeLogConfigEndpointKinesisStreamConfigArgs(
            role_arn=example_role.arn,
            stream_arn=aws_kinesis_stream["example"]["arn"],
        ),
    ),
    opts=pulumi.ResourceOptions(depends_on=[example_role_policy]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
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
                            "cloudfront.amazonaws.com",
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

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "kinesis:DescribeStreamSummary",
                    "kinesis:DescribeStream",
                    "kinesis:PutRecord",
                    "kinesis:PutRecords",
                },
                Resources = new[]
                {
                    aws_kinesis_stream.Example.Arn,
                },
            },
        },
    });

    var exampleRolePolicy = new Aws.Iam.RolePolicy("exampleRolePolicy", new()
    {
        Role = exampleRole.Id,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleRealtimeLogConfig = new Aws.CloudFront.RealtimeLogConfig("exampleRealtimeLogConfig", new()
    {
        SamplingRate = 75,
        Fields = new[]
        {
            "timestamp",
            "c-ip",
        },
        Endpoint = new Aws.CloudFront.Inputs.RealtimeLogConfigEndpointArgs
        {
            StreamType = "Kinesis",
            KinesisStreamConfig = new Aws.CloudFront.Inputs.RealtimeLogConfigEndpointKinesisStreamConfigArgs
            {
                RoleArn = exampleRole.Arn,
                StreamArn = aws_kinesis_stream.Example.Arn,
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleRolePolicy,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Effect: pulumi.StringRef("Allow"),
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "Service",
Identifiers: []string{
"cloudfront.amazonaws.com",
},
},
},
Actions: []string{
"sts:AssumeRole",
},
},
},
}, nil);
if err != nil {
return err
}
exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
AssumeRolePolicy: *pulumi.String(assumeRole.Json),
})
if err != nil {
return err
}
examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Effect: pulumi.StringRef("Allow"),
Actions: []string{
"kinesis:DescribeStreamSummary",
"kinesis:DescribeStream",
"kinesis:PutRecord",
"kinesis:PutRecords",
},
Resources: interface{}{
aws_kinesis_stream.Example.Arn,
},
},
},
}, nil);
if err != nil {
return err
}
exampleRolePolicy, err := iam.NewRolePolicy(ctx, "exampleRolePolicy", &iam.RolePolicyArgs{
Role: exampleRole.ID(),
Policy: *pulumi.String(examplePolicyDocument.Json),
})
if err != nil {
return err
}
_, err = cloudfront.NewRealtimeLogConfig(ctx, "exampleRealtimeLogConfig", &cloudfront.RealtimeLogConfigArgs{
SamplingRate: pulumi.Int(75),
Fields: pulumi.StringArray{
pulumi.String("timestamp"),
pulumi.String("c-ip"),
},
Endpoint: &cloudfront.RealtimeLogConfigEndpointArgs{
StreamType: pulumi.String("Kinesis"),
KinesisStreamConfig: &cloudfront.RealtimeLogConfigEndpointKinesisStreamConfigArgs{
RoleArn: exampleRole.Arn,
StreamArn: pulumi.Any(aws_kinesis_stream.Example.Arn),
},
},
}, pulumi.DependsOn([]pulumi.Resource{
exampleRolePolicy,
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import com.pulumi.aws.cloudfront.RealtimeLogConfig;
import com.pulumi.aws.cloudfront.RealtimeLogConfigArgs;
import com.pulumi.aws.cloudfront.inputs.RealtimeLogConfigEndpointArgs;
import com.pulumi.aws.cloudfront.inputs.RealtimeLogConfigEndpointKinesisStreamConfigArgs;
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
        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("cloudfront.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions(                
                    "kinesis:DescribeStreamSummary",
                    "kinesis:DescribeStream",
                    "kinesis:PutRecord",
                    "kinesis:PutRecords")
                .resources(aws_kinesis_stream.example().arn())
                .build())
            .build());

        var exampleRolePolicy = new RolePolicy("exampleRolePolicy", RolePolicyArgs.builder()        
            .role(exampleRole.id())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleRealtimeLogConfig = new RealtimeLogConfig("exampleRealtimeLogConfig", RealtimeLogConfigArgs.builder()        
            .samplingRate(75)
            .fields(            
                "timestamp",
                "c-ip")
            .endpoint(RealtimeLogConfigEndpointArgs.builder()
                .streamType("Kinesis")
                .kinesisStreamConfig(RealtimeLogConfigEndpointKinesisStreamConfigArgs.builder()
                    .roleArn(exampleRole.arn())
                    .streamArn(aws_kinesis_stream.example().arn())
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleRolePolicy)
                .build());

    }
}
```
```yaml
resources:
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  exampleRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${exampleRole.id}
      policy: ${examplePolicyDocument.json}
  exampleRealtimeLogConfig:
    type: aws:cloudfront:RealtimeLogConfig
    properties:
      samplingRate: 75
      fields:
        - timestamp
        - c-ip
      endpoint:
        streamType: Kinesis
        kinesisStreamConfig:
          roleArn: ${exampleRole.arn}
          streamArn: ${aws_kinesis_stream.example.arn}
    options:
      dependson:
        - ${exampleRolePolicy}
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
                  - cloudfront.amazonaws.com
            actions:
              - sts:AssumeRole
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - kinesis:DescribeStreamSummary
              - kinesis:DescribeStream
              - kinesis:PutRecord
              - kinesis:PutRecords
            resources:
              - ${aws_kinesis_stream.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudFront real-time log configurations using the ARN. For example:

```sh
 $ pulumi import aws:cloudfront/realtimeLogConfig:RealtimeLogConfig example arn:aws:cloudfront::111122223333:realtime-log-config/ExampleNameForRealtimeLogConfig
```
 