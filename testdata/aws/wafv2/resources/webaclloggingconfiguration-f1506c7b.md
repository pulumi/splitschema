This resource creates a WAFv2 Web ACL Logging Configuration.

!> **WARNING:** When logging from a WAFv2 Web ACL to a CloudWatch Log Group, the WAFv2 service tries to create or update a generic Log Resource Policy named `AWSWAF-LOGS`. However, if there are a large number of Web ACLs or if the account frequently creates and deletes Web ACLs, this policy may exceed the maximum policy size. As a result, this resource type will fail to be created. More details about this issue can be found in this issue. To prevent this issue, you can manage a specific resource policy. Please refer to the example below for managing a CloudWatch Log Group with a managed CloudWatch Log Resource Policy.

{{% examples %}}
## Example Usage
{{% example %}}
### With Redacted Fields

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafv2.WebAclLoggingConfiguration("example", {
    logDestinationConfigs: [aws_kinesis_firehose_delivery_stream.example.arn],
    resourceArn: aws_wafv2_web_acl.example.arn,
    redactedFields: [{
        singleHeader: {
            name: "user-agent",
        },
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.wafv2.WebAclLoggingConfiguration("example",
    log_destination_configs=[aws_kinesis_firehose_delivery_stream["example"]["arn"]],
    resource_arn=aws_wafv2_web_acl["example"]["arn"],
    redacted_fields=[aws.wafv2.WebAclLoggingConfigurationRedactedFieldArgs(
        single_header=aws.wafv2.WebAclLoggingConfigurationRedactedFieldSingleHeaderArgs(
            name="user-agent",
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.WafV2.WebAclLoggingConfiguration("example", new()
    {
        LogDestinationConfigs = new[]
        {
            aws_kinesis_firehose_delivery_stream.Example.Arn,
        },
        ResourceArn = aws_wafv2_web_acl.Example.Arn,
        RedactedFields = new[]
        {
            new Aws.WafV2.Inputs.WebAclLoggingConfigurationRedactedFieldArgs
            {
                SingleHeader = new Aws.WafV2.Inputs.WebAclLoggingConfigurationRedactedFieldSingleHeaderArgs
                {
                    Name = "user-agent",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafv2.NewWebAclLoggingConfiguration(ctx, "example", &wafv2.WebAclLoggingConfigurationArgs{
			LogDestinationConfigs: pulumi.StringArray{
				aws_kinesis_firehose_delivery_stream.Example.Arn,
			},
			ResourceArn: pulumi.Any(aws_wafv2_web_acl.Example.Arn),
			RedactedFields: wafv2.WebAclLoggingConfigurationRedactedFieldArray{
				&wafv2.WebAclLoggingConfigurationRedactedFieldArgs{
					SingleHeader: &wafv2.WebAclLoggingConfigurationRedactedFieldSingleHeaderArgs{
						Name: pulumi.String("user-agent"),
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
import com.pulumi.aws.wafv2.WebAclLoggingConfiguration;
import com.pulumi.aws.wafv2.WebAclLoggingConfigurationArgs;
import com.pulumi.aws.wafv2.inputs.WebAclLoggingConfigurationRedactedFieldArgs;
import com.pulumi.aws.wafv2.inputs.WebAclLoggingConfigurationRedactedFieldSingleHeaderArgs;
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
        var example = new WebAclLoggingConfiguration("example", WebAclLoggingConfigurationArgs.builder()        
            .logDestinationConfigs(aws_kinesis_firehose_delivery_stream.example().arn())
            .resourceArn(aws_wafv2_web_acl.example().arn())
            .redactedFields(WebAclLoggingConfigurationRedactedFieldArgs.builder()
                .singleHeader(WebAclLoggingConfigurationRedactedFieldSingleHeaderArgs.builder()
                    .name("user-agent")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:wafv2:WebAclLoggingConfiguration
    properties:
      logDestinationConfigs:
        - ${aws_kinesis_firehose_delivery_stream.example.arn}
      resourceArn: ${aws_wafv2_web_acl.example.arn}
      redactedFields:
        - singleHeader:
            name: user-agent
```
{{% /example %}}
{{% example %}}
### With Logging Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafv2.WebAclLoggingConfiguration("example", {
    logDestinationConfigs: [aws_kinesis_firehose_delivery_stream.example.arn],
    resourceArn: aws_wafv2_web_acl.example.arn,
    loggingFilter: {
        defaultBehavior: "KEEP",
        filters: [
            {
                behavior: "DROP",
                conditions: [
                    {
                        actionCondition: {
                            action: "COUNT",
                        },
                    },
                    {
                        labelNameCondition: {
                            labelName: "awswaf:111122223333:rulegroup:testRules:LabelNameZ",
                        },
                    },
                ],
                requirement: "MEETS_ALL",
            },
            {
                behavior: "KEEP",
                conditions: [{
                    actionCondition: {
                        action: "ALLOW",
                    },
                }],
                requirement: "MEETS_ANY",
            },
        ],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.wafv2.WebAclLoggingConfiguration("example",
    log_destination_configs=[aws_kinesis_firehose_delivery_stream["example"]["arn"]],
    resource_arn=aws_wafv2_web_acl["example"]["arn"],
    logging_filter=aws.wafv2.WebAclLoggingConfigurationLoggingFilterArgs(
        default_behavior="KEEP",
        filters=[
            aws.wafv2.WebAclLoggingConfigurationLoggingFilterFilterArgs(
                behavior="DROP",
                conditions=[
                    aws.wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs(
                        action_condition=aws.wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionArgs(
                            action="COUNT",
                        ),
                    ),
                    aws.wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs(
                        label_name_condition=aws.wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameConditionArgs(
                            label_name="awswaf:111122223333:rulegroup:testRules:LabelNameZ",
                        ),
                    ),
                ],
                requirement="MEETS_ALL",
            ),
            aws.wafv2.WebAclLoggingConfigurationLoggingFilterFilterArgs(
                behavior="KEEP",
                conditions=[aws.wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs(
                    action_condition=aws.wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionArgs(
                        action="ALLOW",
                    ),
                )],
                requirement="MEETS_ANY",
            ),
        ],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.WafV2.WebAclLoggingConfiguration("example", new()
    {
        LogDestinationConfigs = new[]
        {
            aws_kinesis_firehose_delivery_stream.Example.Arn,
        },
        ResourceArn = aws_wafv2_web_acl.Example.Arn,
        LoggingFilter = new Aws.WafV2.Inputs.WebAclLoggingConfigurationLoggingFilterArgs
        {
            DefaultBehavior = "KEEP",
            Filters = new[]
            {
                new Aws.WafV2.Inputs.WebAclLoggingConfigurationLoggingFilterFilterArgs
                {
                    Behavior = "DROP",
                    Conditions = new[]
                    {
                        new Aws.WafV2.Inputs.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs
                        {
                            ActionCondition = new Aws.WafV2.Inputs.WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionArgs
                            {
                                Action = "COUNT",
                            },
                        },
                        new Aws.WafV2.Inputs.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs
                        {
                            LabelNameCondition = new Aws.WafV2.Inputs.WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameConditionArgs
                            {
                                LabelName = "awswaf:111122223333:rulegroup:testRules:LabelNameZ",
                            },
                        },
                    },
                    Requirement = "MEETS_ALL",
                },
                new Aws.WafV2.Inputs.WebAclLoggingConfigurationLoggingFilterFilterArgs
                {
                    Behavior = "KEEP",
                    Conditions = new[]
                    {
                        new Aws.WafV2.Inputs.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs
                        {
                            ActionCondition = new Aws.WafV2.Inputs.WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionArgs
                            {
                                Action = "ALLOW",
                            },
                        },
                    },
                    Requirement = "MEETS_ANY",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafv2.NewWebAclLoggingConfiguration(ctx, "example", &wafv2.WebAclLoggingConfigurationArgs{
			LogDestinationConfigs: pulumi.StringArray{
				aws_kinesis_firehose_delivery_stream.Example.Arn,
			},
			ResourceArn: pulumi.Any(aws_wafv2_web_acl.Example.Arn),
			LoggingFilter: &wafv2.WebAclLoggingConfigurationLoggingFilterArgs{
				DefaultBehavior: pulumi.String("KEEP"),
				Filters: wafv2.WebAclLoggingConfigurationLoggingFilterFilterArray{
					&wafv2.WebAclLoggingConfigurationLoggingFilterFilterArgs{
						Behavior: pulumi.String("DROP"),
						Conditions: wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArray{
							&wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs{
								ActionCondition: &wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionArgs{
									Action: pulumi.String("COUNT"),
								},
							},
							&wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs{
								LabelNameCondition: &wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameConditionArgs{
									LabelName: pulumi.String("awswaf:111122223333:rulegroup:testRules:LabelNameZ"),
								},
							},
						},
						Requirement: pulumi.String("MEETS_ALL"),
					},
					&wafv2.WebAclLoggingConfigurationLoggingFilterFilterArgs{
						Behavior: pulumi.String("KEEP"),
						Conditions: wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArray{
							&wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs{
								ActionCondition: &wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionArgs{
									Action: pulumi.String("ALLOW"),
								},
							},
						},
						Requirement: pulumi.String("MEETS_ANY"),
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
import com.pulumi.aws.wafv2.WebAclLoggingConfiguration;
import com.pulumi.aws.wafv2.WebAclLoggingConfigurationArgs;
import com.pulumi.aws.wafv2.inputs.WebAclLoggingConfigurationLoggingFilterArgs;
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
        var example = new WebAclLoggingConfiguration("example", WebAclLoggingConfigurationArgs.builder()        
            .logDestinationConfigs(aws_kinesis_firehose_delivery_stream.example().arn())
            .resourceArn(aws_wafv2_web_acl.example().arn())
            .loggingFilter(WebAclLoggingConfigurationLoggingFilterArgs.builder()
                .defaultBehavior("KEEP")
                .filters(                
                    WebAclLoggingConfigurationLoggingFilterFilterArgs.builder()
                        .behavior("DROP")
                        .conditions(                        
                            WebAclLoggingConfigurationLoggingFilterFilterConditionArgs.builder()
                                .actionCondition(WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionArgs.builder()
                                    .action("COUNT")
                                    .build())
                                .build(),
                            WebAclLoggingConfigurationLoggingFilterFilterConditionArgs.builder()
                                .labelNameCondition(WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameConditionArgs.builder()
                                    .labelName("awswaf:111122223333:rulegroup:testRules:LabelNameZ")
                                    .build())
                                .build())
                        .requirement("MEETS_ALL")
                        .build(),
                    WebAclLoggingConfigurationLoggingFilterFilterArgs.builder()
                        .behavior("KEEP")
                        .conditions(WebAclLoggingConfigurationLoggingFilterFilterConditionArgs.builder()
                            .actionCondition(WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionArgs.builder()
                                .action("ALLOW")
                                .build())
                            .build())
                        .requirement("MEETS_ANY")
                        .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:wafv2:WebAclLoggingConfiguration
    properties:
      logDestinationConfigs:
        - ${aws_kinesis_firehose_delivery_stream.example.arn}
      resourceArn: ${aws_wafv2_web_acl.example.arn}
      loggingFilter:
        defaultBehavior: KEEP
        filters:
          - behavior: DROP
            conditions:
              - actionCondition:
                  action: COUNT
              - labelNameCondition:
                  labelName: awswaf:111122223333:rulegroup:testRules:LabelNameZ
            requirement: MEETS_ALL
          - behavior: KEEP
            conditions:
              - actionCondition:
                  action: ALLOW
            requirement: MEETS_ANY
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAFv2 Web ACL Logging Configurations using the ARN of the WAFv2 Web ACL. For example:

```sh
 $ pulumi import aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration example arn:aws:wafv2:us-west-2:123456789012:regional/webacl/test-logs/a1b2c3d4-5678-90ab-cdef
```
 