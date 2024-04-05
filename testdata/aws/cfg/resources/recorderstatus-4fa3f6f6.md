Manages status (recording / stopped) of an AWS Config Configuration Recorder.

> **Note:** Starting Configuration Recorder requires a Delivery Channel to be present. Use of `depends_on` (as shown below) is recommended to avoid race conditions.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucketV2 = new aws.s3.BucketV2("bucketV2", {});
const fooDeliveryChannel = new aws.cfg.DeliveryChannel("fooDeliveryChannel", {s3BucketName: bucketV2.bucket});
const fooRecorderStatus = new aws.cfg.RecorderStatus("fooRecorderStatus", {isEnabled: true}, {
    dependsOn: [fooDeliveryChannel],
});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["config.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const role = new aws.iam.Role("role", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const rolePolicyAttachment = new aws.iam.RolePolicyAttachment("rolePolicyAttachment", {
    role: role.name,
    policyArn: "arn:aws:iam::aws:policy/service-role/AWS_ConfigRole",
});
const fooRecorder = new aws.cfg.Recorder("fooRecorder", {roleArn: role.arn});
const policyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        actions: ["s3:*"],
        resources: [
            bucketV2.arn,
            pulumi.interpolate`${bucketV2.arn}/*`,
        ],
    }],
});
const rolePolicy = new aws.iam.RolePolicy("rolePolicy", {
    role: role.id,
    policy: policyDocument.apply(policyDocument => policyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

bucket_v2 = aws.s3.BucketV2("bucketV2")
foo_delivery_channel = aws.cfg.DeliveryChannel("fooDeliveryChannel", s3_bucket_name=bucket_v2.bucket)
foo_recorder_status = aws.cfg.RecorderStatus("fooRecorderStatus", is_enabled=True,
opts=pulumi.ResourceOptions(depends_on=[foo_delivery_channel]))
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["config.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
role = aws.iam.Role("role", assume_role_policy=assume_role.json)
role_policy_attachment = aws.iam.RolePolicyAttachment("rolePolicyAttachment",
    role=role.name,
    policy_arn="arn:aws:iam::aws:policy/service-role/AWS_ConfigRole")
foo_recorder = aws.cfg.Recorder("fooRecorder", role_arn=role.arn)
policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["s3:*"],
    resources=[
        bucket_v2.arn,
        bucket_v2.arn.apply(lambda arn: f"{arn}/*"),
    ],
)])
role_policy = aws.iam.RolePolicy("rolePolicy",
    role=role.id,
    policy=policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bucketV2 = new Aws.S3.BucketV2("bucketV2");

    var fooDeliveryChannel = new Aws.Cfg.DeliveryChannel("fooDeliveryChannel", new()
    {
        S3BucketName = bucketV2.Bucket,
    });

    var fooRecorderStatus = new Aws.Cfg.RecorderStatus("fooRecorderStatus", new()
    {
        IsEnabled = true,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            fooDeliveryChannel,
        },
    });

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
                            "config.amazonaws.com",
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

    var role = new Aws.Iam.Role("role", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var rolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("rolePolicyAttachment", new()
    {
        Role = role.Name,
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AWS_ConfigRole",
    });

    var fooRecorder = new Aws.Cfg.Recorder("fooRecorder", new()
    {
        RoleArn = role.Arn,
    });

    var policyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "s3:*",
                },
                Resources = new[]
                {
                    bucketV2.Arn,
                    $"{bucketV2.Arn}/*",
                },
            },
        },
    });

    var rolePolicy = new Aws.Iam.RolePolicy("rolePolicy", new()
    {
        Role = role.Id,
        Policy = policyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucketV2, err := s3.NewBucketV2(ctx, "bucketV2", nil)
		if err != nil {
			return err
		}
		fooDeliveryChannel, err := cfg.NewDeliveryChannel(ctx, "fooDeliveryChannel", &cfg.DeliveryChannelArgs{
			S3BucketName: bucketV2.Bucket,
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewRecorderStatus(ctx, "fooRecorderStatus", &cfg.RecorderStatusArgs{
			IsEnabled: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			fooDeliveryChannel,
		}))
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
								"config.amazonaws.com",
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
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "rolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      role.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWS_ConfigRole"),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewRecorder(ctx, "fooRecorder", &cfg.RecorderArgs{
			RoleArn: role.Arn,
		})
		if err != nil {
			return err
		}
		policyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:*"),
					},
					Resources: pulumi.StringArray{
						bucketV2.Arn,
						bucketV2.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
				},
			},
		}, nil)
		_, err = iam.NewRolePolicy(ctx, "rolePolicy", &iam.RolePolicyArgs{
			Role: role.ID(),
			Policy: policyDocument.ApplyT(func(policyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &policyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
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
import com.pulumi.aws.cfg.DeliveryChannel;
import com.pulumi.aws.cfg.DeliveryChannelArgs;
import com.pulumi.aws.cfg.RecorderStatus;
import com.pulumi.aws.cfg.RecorderStatusArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.cfg.Recorder;
import com.pulumi.aws.cfg.RecorderArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
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
        var bucketV2 = new BucketV2("bucketV2");

        var fooDeliveryChannel = new DeliveryChannel("fooDeliveryChannel", DeliveryChannelArgs.builder()        
            .s3BucketName(bucketV2.bucket())
            .build());

        var fooRecorderStatus = new RecorderStatus("fooRecorderStatus", RecorderStatusArgs.builder()        
            .isEnabled(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(fooDeliveryChannel)
                .build());

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("config.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var role = new Role("role", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var rolePolicyAttachment = new RolePolicyAttachment("rolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .role(role.name())
            .policyArn("arn:aws:iam::aws:policy/service-role/AWS_ConfigRole")
            .build());

        var fooRecorder = new Recorder("fooRecorder", RecorderArgs.builder()        
            .roleArn(role.arn())
            .build());

        final var policyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("s3:*")
                .resources(                
                    bucketV2.arn(),
                    bucketV2.arn().applyValue(arn -> String.format("%s/*", arn)))
                .build())
            .build());

        var rolePolicy = new RolePolicy("rolePolicy", RolePolicyArgs.builder()        
            .role(role.id())
            .policy(policyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(policyDocument -> policyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  fooRecorderStatus:
    type: aws:cfg:RecorderStatus
    properties:
      isEnabled: true
    options:
      dependson:
        - ${fooDeliveryChannel}
  rolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${role.name}
      policyArn: arn:aws:iam::aws:policy/service-role/AWS_ConfigRole
  bucketV2:
    type: aws:s3:BucketV2
  fooDeliveryChannel:
    type: aws:cfg:DeliveryChannel
    properties:
      s3BucketName: ${bucketV2.bucket}
  fooRecorder:
    type: aws:cfg:Recorder
    properties:
      roleArn: ${role.arn}
  role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  rolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${role.id}
      policy: ${policyDocument.json}
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
                  - config.amazonaws.com
            actions:
              - sts:AssumeRole
  policyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - s3:*
            resources:
              - ${bucketV2.arn}
              - ${bucketV2.arn}/*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Configuration Recorder Status using the name of the Configuration Recorder. For example:

```sh
 $ pulumi import aws:cfg/recorderStatus:RecorderStatus foo example
```
 