Provides an AWS Config Configuration Recorder. Please note that this resource **does not start** the created recorder automatically.

> **Note:** _Starting_ the Configuration Recorder requires a delivery channel (while delivery channel creation requires Configuration Recorder). This is why `aws.cfg.RecorderStatus` is a separate resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

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
const foo = new aws.cfg.Recorder("foo", {roleArn: role.arn});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["config.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
role = aws.iam.Role("role", assume_role_policy=assume_role.json)
foo = aws.cfg.Recorder("foo", role_arn=role.arn)
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

    var foo = new Aws.Cfg.Recorder("foo", new()
    {
        RoleArn = role.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
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
		_, err = cfg.NewRecorder(ctx, "foo", &cfg.RecorderArgs{
			RoleArn: role.Arn,
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.cfg.Recorder;
import com.pulumi.aws.cfg.RecorderArgs;
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
                    .identifiers("config.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var role = new Role("role", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var foo = new Recorder("foo", RecorderArgs.builder()        
            .roleArn(role.arn())
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:cfg:Recorder
    properties:
      roleArn: ${role.arn}
  role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
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
```
{{% /example %}}
{{% example %}}
### Exclude Resources Types Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.cfg.Recorder("foo", {
    roleArn: aws_iam_role.r.arn,
    recordingGroup: {
        allSupported: false,
        exclusionByResourceTypes: [{
            resourceTypes: ["AWS::EC2::Instance"],
        }],
        recordingStrategies: [{
            useOnly: "EXCLUSION_BY_RESOURCE_TYPES",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.cfg.Recorder("foo",
    role_arn=aws_iam_role["r"]["arn"],
    recording_group=aws.cfg.RecorderRecordingGroupArgs(
        all_supported=False,
        exclusion_by_resource_types=[aws.cfg.RecorderRecordingGroupExclusionByResourceTypeArgs(
            resource_types=["AWS::EC2::Instance"],
        )],
        recording_strategies=[aws.cfg.RecorderRecordingGroupRecordingStrategyArgs(
            use_only="EXCLUSION_BY_RESOURCE_TYPES",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Cfg.Recorder("foo", new()
    {
        RoleArn = aws_iam_role.R.Arn,
        RecordingGroup = new Aws.Cfg.Inputs.RecorderRecordingGroupArgs
        {
            AllSupported = false,
            ExclusionByResourceTypes = new[]
            {
                new Aws.Cfg.Inputs.RecorderRecordingGroupExclusionByResourceTypeArgs
                {
                    ResourceTypes = new[]
                    {
                        "AWS::EC2::Instance",
                    },
                },
            },
            RecordingStrategies = new[]
            {
                new Aws.Cfg.Inputs.RecorderRecordingGroupRecordingStrategyArgs
                {
                    UseOnly = "EXCLUSION_BY_RESOURCE_TYPES",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cfg.NewRecorder(ctx, "foo", &cfg.RecorderArgs{
			RoleArn: pulumi.Any(aws_iam_role.R.Arn),
			RecordingGroup: &cfg.RecorderRecordingGroupArgs{
				AllSupported: pulumi.Bool(false),
				ExclusionByResourceTypes: cfg.RecorderRecordingGroupExclusionByResourceTypeArray{
					&cfg.RecorderRecordingGroupExclusionByResourceTypeArgs{
						ResourceTypes: pulumi.StringArray{
							pulumi.String("AWS::EC2::Instance"),
						},
					},
				},
				RecordingStrategies: cfg.RecorderRecordingGroupRecordingStrategyArray{
					&cfg.RecorderRecordingGroupRecordingStrategyArgs{
						UseOnly: pulumi.String("EXCLUSION_BY_RESOURCE_TYPES"),
					},
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
import com.pulumi.aws.cfg.Recorder;
import com.pulumi.aws.cfg.RecorderArgs;
import com.pulumi.aws.cfg.inputs.RecorderRecordingGroupArgs;
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
        var foo = new Recorder("foo", RecorderArgs.builder()        
            .roleArn(aws_iam_role.r().arn())
            .recordingGroup(RecorderRecordingGroupArgs.builder()
                .allSupported(false)
                .exclusionByResourceTypes(RecorderRecordingGroupExclusionByResourceTypeArgs.builder()
                    .resourceTypes("AWS::EC2::Instance")
                    .build())
                .recordingStrategies(RecorderRecordingGroupRecordingStrategyArgs.builder()
                    .useOnly("EXCLUSION_BY_RESOURCE_TYPES")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:cfg:Recorder
    properties:
      roleArn: ${aws_iam_role.r.arn}
      recordingGroup:
        allSupported: false
        exclusionByResourceTypes:
          - resourceTypes:
              - AWS::EC2::Instance
        recordingStrategies:
          - useOnly: EXCLUSION_BY_RESOURCE_TYPES
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Configuration Recorder using the name. For example:

```sh
 $ pulumi import aws:cfg/recorder:Recorder foo example
```
 