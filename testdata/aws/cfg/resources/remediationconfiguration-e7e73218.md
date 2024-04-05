Provides an AWS Config Remediation Configuration.

> **Note:** Config Remediation Configuration requires an existing Config Rule to be present.

{{% examples %}}
## Example Usage
{{% example %}}

AWS managed rules can be used by setting the source owner to `AWS` and the source identifier to the name of the managed rule. More information about AWS managed rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const thisRule = new aws.cfg.Rule("thisRule", {source: {
    owner: "AWS",
    sourceIdentifier: "S3_BUCKET_VERSIONING_ENABLED",
}});
const thisRemediationConfiguration = new aws.cfg.RemediationConfiguration("thisRemediationConfiguration", {
    configRuleName: thisRule.name,
    resourceType: "AWS::S3::Bucket",
    targetType: "SSM_DOCUMENT",
    targetId: "AWS-EnableS3BucketEncryption",
    targetVersion: "1",
    parameters: [
        {
            name: "AutomationAssumeRole",
            staticValue: "arn:aws:iam::875924563244:role/security_config",
        },
        {
            name: "BucketName",
            resourceValue: "RESOURCE_ID",
        },
        {
            name: "SSEAlgorithm",
            staticValue: "AES256",
        },
    ],
    automatic: true,
    maximumAutomaticAttempts: 10,
    retryAttemptSeconds: 600,
    executionControls: {
        ssmControls: {
            concurrentExecutionRatePercentage: 25,
            errorPercentage: 20,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

this_rule = aws.cfg.Rule("thisRule", source=aws.cfg.RuleSourceArgs(
    owner="AWS",
    source_identifier="S3_BUCKET_VERSIONING_ENABLED",
))
this_remediation_configuration = aws.cfg.RemediationConfiguration("thisRemediationConfiguration",
    config_rule_name=this_rule.name,
    resource_type="AWS::S3::Bucket",
    target_type="SSM_DOCUMENT",
    target_id="AWS-EnableS3BucketEncryption",
    target_version="1",
    parameters=[
        aws.cfg.RemediationConfigurationParameterArgs(
            name="AutomationAssumeRole",
            static_value="arn:aws:iam::875924563244:role/security_config",
        ),
        aws.cfg.RemediationConfigurationParameterArgs(
            name="BucketName",
            resource_value="RESOURCE_ID",
        ),
        aws.cfg.RemediationConfigurationParameterArgs(
            name="SSEAlgorithm",
            static_value="AES256",
        ),
    ],
    automatic=True,
    maximum_automatic_attempts=10,
    retry_attempt_seconds=600,
    execution_controls=aws.cfg.RemediationConfigurationExecutionControlsArgs(
        ssm_controls=aws.cfg.RemediationConfigurationExecutionControlsSsmControlsArgs(
            concurrent_execution_rate_percentage=25,
            error_percentage=20,
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var thisRule = new Aws.Cfg.Rule("thisRule", new()
    {
        Source = new Aws.Cfg.Inputs.RuleSourceArgs
        {
            Owner = "AWS",
            SourceIdentifier = "S3_BUCKET_VERSIONING_ENABLED",
        },
    });

    var thisRemediationConfiguration = new Aws.Cfg.RemediationConfiguration("thisRemediationConfiguration", new()
    {
        ConfigRuleName = thisRule.Name,
        ResourceType = "AWS::S3::Bucket",
        TargetType = "SSM_DOCUMENT",
        TargetId = "AWS-EnableS3BucketEncryption",
        TargetVersion = "1",
        Parameters = new[]
        {
            new Aws.Cfg.Inputs.RemediationConfigurationParameterArgs
            {
                Name = "AutomationAssumeRole",
                StaticValue = "arn:aws:iam::875924563244:role/security_config",
            },
            new Aws.Cfg.Inputs.RemediationConfigurationParameterArgs
            {
                Name = "BucketName",
                ResourceValue = "RESOURCE_ID",
            },
            new Aws.Cfg.Inputs.RemediationConfigurationParameterArgs
            {
                Name = "SSEAlgorithm",
                StaticValue = "AES256",
            },
        },
        Automatic = true,
        MaximumAutomaticAttempts = 10,
        RetryAttemptSeconds = 600,
        ExecutionControls = new Aws.Cfg.Inputs.RemediationConfigurationExecutionControlsArgs
        {
            SsmControls = new Aws.Cfg.Inputs.RemediationConfigurationExecutionControlsSsmControlsArgs
            {
                ConcurrentExecutionRatePercentage = 25,
                ErrorPercentage = 20,
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
		thisRule, err := cfg.NewRule(ctx, "thisRule", &cfg.RuleArgs{
			Source: &cfg.RuleSourceArgs{
				Owner:            pulumi.String("AWS"),
				SourceIdentifier: pulumi.String("S3_BUCKET_VERSIONING_ENABLED"),
			},
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewRemediationConfiguration(ctx, "thisRemediationConfiguration", &cfg.RemediationConfigurationArgs{
			ConfigRuleName: thisRule.Name,
			ResourceType:   pulumi.String("AWS::S3::Bucket"),
			TargetType:     pulumi.String("SSM_DOCUMENT"),
			TargetId:       pulumi.String("AWS-EnableS3BucketEncryption"),
			TargetVersion:  pulumi.String("1"),
			Parameters: cfg.RemediationConfigurationParameterArray{
				&cfg.RemediationConfigurationParameterArgs{
					Name:        pulumi.String("AutomationAssumeRole"),
					StaticValue: pulumi.String("arn:aws:iam::875924563244:role/security_config"),
				},
				&cfg.RemediationConfigurationParameterArgs{
					Name:          pulumi.String("BucketName"),
					ResourceValue: pulumi.String("RESOURCE_ID"),
				},
				&cfg.RemediationConfigurationParameterArgs{
					Name:        pulumi.String("SSEAlgorithm"),
					StaticValue: pulumi.String("AES256"),
				},
			},
			Automatic:                pulumi.Bool(true),
			MaximumAutomaticAttempts: pulumi.Int(10),
			RetryAttemptSeconds:      pulumi.Int(600),
			ExecutionControls: &cfg.RemediationConfigurationExecutionControlsArgs{
				SsmControls: &cfg.RemediationConfigurationExecutionControlsSsmControlsArgs{
					ConcurrentExecutionRatePercentage: pulumi.Int(25),
					ErrorPercentage:                   pulumi.Int(20),
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
import com.pulumi.aws.cfg.Rule;
import com.pulumi.aws.cfg.RuleArgs;
import com.pulumi.aws.cfg.inputs.RuleSourceArgs;
import com.pulumi.aws.cfg.RemediationConfiguration;
import com.pulumi.aws.cfg.RemediationConfigurationArgs;
import com.pulumi.aws.cfg.inputs.RemediationConfigurationParameterArgs;
import com.pulumi.aws.cfg.inputs.RemediationConfigurationExecutionControlsArgs;
import com.pulumi.aws.cfg.inputs.RemediationConfigurationExecutionControlsSsmControlsArgs;
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
        var thisRule = new Rule("thisRule", RuleArgs.builder()        
            .source(RuleSourceArgs.builder()
                .owner("AWS")
                .sourceIdentifier("S3_BUCKET_VERSIONING_ENABLED")
                .build())
            .build());

        var thisRemediationConfiguration = new RemediationConfiguration("thisRemediationConfiguration", RemediationConfigurationArgs.builder()        
            .configRuleName(thisRule.name())
            .resourceType("AWS::S3::Bucket")
            .targetType("SSM_DOCUMENT")
            .targetId("AWS-EnableS3BucketEncryption")
            .targetVersion("1")
            .parameters(            
                RemediationConfigurationParameterArgs.builder()
                    .name("AutomationAssumeRole")
                    .staticValue("arn:aws:iam::875924563244:role/security_config")
                    .build(),
                RemediationConfigurationParameterArgs.builder()
                    .name("BucketName")
                    .resourceValue("RESOURCE_ID")
                    .build(),
                RemediationConfigurationParameterArgs.builder()
                    .name("SSEAlgorithm")
                    .staticValue("AES256")
                    .build())
            .automatic(true)
            .maximumAutomaticAttempts(10)
            .retryAttemptSeconds(600)
            .executionControls(RemediationConfigurationExecutionControlsArgs.builder()
                .ssmControls(RemediationConfigurationExecutionControlsSsmControlsArgs.builder()
                    .concurrentExecutionRatePercentage(25)
                    .errorPercentage(20)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  thisRule:
    type: aws:cfg:Rule
    properties:
      source:
        owner: AWS
        sourceIdentifier: S3_BUCKET_VERSIONING_ENABLED
  thisRemediationConfiguration:
    type: aws:cfg:RemediationConfiguration
    properties:
      configRuleName: ${thisRule.name}
      resourceType: AWS::S3::Bucket
      targetType: SSM_DOCUMENT
      targetId: AWS-EnableS3BucketEncryption
      targetVersion: '1'
      parameters:
        - name: AutomationAssumeRole
          staticValue: arn:aws:iam::875924563244:role/security_config
        - name: BucketName
          resourceValue: RESOURCE_ID
        - name: SSEAlgorithm
          staticValue: AES256
      automatic: true
      maximumAutomaticAttempts: 10
      retryAttemptSeconds: 600
      executionControls:
        ssmControls:
          concurrentExecutionRatePercentage: 25
          errorPercentage: 20
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Remediation Configurations using the name config_rule_name. For example:

```sh
 $ pulumi import aws:cfg/remediationConfiguration:RemediationConfiguration this example
```
 