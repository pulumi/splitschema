Manages a Config Conformance Pack. More information about this collection of Config rules and remediation actions can be found in the
[Conformance Packs](https://docs.aws.amazon.com/config/latest/developerguide/conformance-packs.html) documentation.
Sample Conformance Pack templates may be found in the
[AWS Config Rules Repository](https://github.com/awslabs/aws-config-rules/tree/master/aws-config-conformance-packs).

> **NOTE:** The account must have a Configuration Recorder with proper IAM permissions before the Conformance Pack will
successfully create or update. See also the
`aws.cfg.Recorder` resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Template Body

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cfg.ConformancePack("example", {
    inputParameters: [{
        parameterName: "AccessKeysRotatedParameterMaxAccessKeyAge",
        parameterValue: "90",
    }],
    templateBody: `Parameters:
  AccessKeysRotatedParameterMaxAccessKeyAge:
    Type: String
Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
`,
}, {
    dependsOn: [aws_config_configuration_recorder.example],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cfg.ConformancePack("example",
    input_parameters=[aws.cfg.ConformancePackInputParameterArgs(
        parameter_name="AccessKeysRotatedParameterMaxAccessKeyAge",
        parameter_value="90",
    )],
    template_body="""Parameters:
  AccessKeysRotatedParameterMaxAccessKeyAge:
    Type: String
Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
""",
    opts=pulumi.ResourceOptions(depends_on=[aws_config_configuration_recorder["example"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Cfg.ConformancePack("example", new()
    {
        InputParameters = new[]
        {
            new Aws.Cfg.Inputs.ConformancePackInputParameterArgs
            {
                ParameterName = "AccessKeysRotatedParameterMaxAccessKeyAge",
                ParameterValue = "90",
            },
        },
        TemplateBody = @"Parameters:
  AccessKeysRotatedParameterMaxAccessKeyAge:
    Type: String
Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_config_configuration_recorder.Example,
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
		_, err := cfg.NewConformancePack(ctx, "example", &cfg.ConformancePackArgs{
			InputParameters: cfg.ConformancePackInputParameterArray{
				&cfg.ConformancePackInputParameterArgs{
					ParameterName:  pulumi.String("AccessKeysRotatedParameterMaxAccessKeyAge"),
					ParameterValue: pulumi.String("90"),
				},
			},
			TemplateBody: pulumi.String(`Parameters:
  AccessKeysRotatedParameterMaxAccessKeyAge:
    Type: String
Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
`),
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_config_configuration_recorder.Example,
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
import com.pulumi.aws.cfg.ConformancePack;
import com.pulumi.aws.cfg.ConformancePackArgs;
import com.pulumi.aws.cfg.inputs.ConformancePackInputParameterArgs;
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
        var example = new ConformancePack("example", ConformancePackArgs.builder()        
            .inputParameters(ConformancePackInputParameterArgs.builder()
                .parameterName("AccessKeysRotatedParameterMaxAccessKeyAge")
                .parameterValue("90")
                .build())
            .templateBody("""
Parameters:
  AccessKeysRotatedParameterMaxAccessKeyAge:
    Type: String
Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
            """)
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_config_configuration_recorder.example())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cfg:ConformancePack
    properties:
      inputParameters:
        - parameterName: AccessKeysRotatedParameterMaxAccessKeyAge
          parameterValue: '90'
      templateBody: |
        Parameters:
          AccessKeysRotatedParameterMaxAccessKeyAge:
            Type: String
        Resources:
          IAMPasswordPolicy:
            Properties:
              ConfigRuleName: IAMPasswordPolicy
              Source:
                Owner: AWS
                SourceIdentifier: IAM_PASSWORD_POLICY
            Type: AWS::Config::ConfigRule
    options:
      dependson:
        - ${aws_config_configuration_recorder.example}
```
{{% /example %}}
{{% example %}}
### Template S3 URI

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketObjectv2 = new aws.s3.BucketObjectv2("exampleBucketObjectv2", {
    bucket: exampleBucketV2.id,
    key: "example-key",
    content: `Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
`,
});
const exampleConformancePack = new aws.cfg.ConformancePack("exampleConformancePack", {templateS3Uri: pulumi.interpolate`s3://${exampleBucketV2.bucket}/${exampleBucketObjectv2.key}`}, {
    dependsOn: [aws_config_configuration_recorder.example],
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_objectv2 = aws.s3.BucketObjectv2("exampleBucketObjectv2",
    bucket=example_bucket_v2.id,
    key="example-key",
    content="""Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
""")
example_conformance_pack = aws.cfg.ConformancePack("exampleConformancePack", template_s3_uri=pulumi.Output.all(example_bucket_v2.bucket, example_bucket_objectv2.key).apply(lambda bucket, key: f"s3://{bucket}/{key}"),
opts=pulumi.ResourceOptions(depends_on=[aws_config_configuration_recorder["example"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketObjectv2 = new Aws.S3.BucketObjectv2("exampleBucketObjectv2", new()
    {
        Bucket = exampleBucketV2.Id,
        Key = "example-key",
        Content = @"Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
",
    });

    var exampleConformancePack = new Aws.Cfg.ConformancePack("exampleConformancePack", new()
    {
        TemplateS3Uri = Output.Tuple(exampleBucketV2.Bucket, exampleBucketObjectv2.Key).Apply(values =>
        {
            var bucket = values.Item1;
            var key = values.Item2;
            return $"s3://{bucket}/{key}";
        }),
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_config_configuration_recorder.Example,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		exampleBucketObjectv2, err := s3.NewBucketObjectv2(ctx, "exampleBucketObjectv2", &s3.BucketObjectv2Args{
			Bucket: exampleBucketV2.ID(),
			Key:    pulumi.String("example-key"),
			Content: pulumi.String(`Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
`),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewConformancePack(ctx, "exampleConformancePack", &cfg.ConformancePackArgs{
			TemplateS3Uri: pulumi.All(exampleBucketV2.Bucket, exampleBucketObjectv2.Key).ApplyT(func(_args []interface{}) (string, error) {
				bucket := _args[0].(string)
				key := _args[1].(string)
				return fmt.Sprintf("s3://%v/%v", bucket, key), nil
			}).(pulumi.StringOutput),
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_config_configuration_recorder.Example,
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.s3.BucketObjectv2Args;
import com.pulumi.aws.cfg.ConformancePack;
import com.pulumi.aws.cfg.ConformancePackArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleBucketObjectv2 = new BucketObjectv2("exampleBucketObjectv2", BucketObjectv2Args.builder()        
            .bucket(exampleBucketV2.id())
            .key("example-key")
            .content("""
Resources:
  IAMPasswordPolicy:
    Properties:
      ConfigRuleName: IAMPasswordPolicy
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
    Type: AWS::Config::ConfigRule
            """)
            .build());

        var exampleConformancePack = new ConformancePack("exampleConformancePack", ConformancePackArgs.builder()        
            .templateS3Uri(Output.tuple(exampleBucketV2.bucket(), exampleBucketObjectv2.key()).applyValue(values -> {
                var bucket = values.t1;
                var key = values.t2;
                return String.format("s3://%s/%s", bucket,key);
            }))
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_config_configuration_recorder.example())
                .build());

    }
}
```
```yaml
resources:
  exampleConformancePack:
    type: aws:cfg:ConformancePack
    properties:
      templateS3Uri: s3://${exampleBucketV2.bucket}/${exampleBucketObjectv2.key}
    options:
      dependson:
        - ${aws_config_configuration_recorder.example}
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketObjectv2:
    type: aws:s3:BucketObjectv2
    properties:
      bucket: ${exampleBucketV2.id}
      key: example-key
      content: |
        Resources:
          IAMPasswordPolicy:
            Properties:
              ConfigRuleName: IAMPasswordPolicy
              Source:
                Owner: AWS
                SourceIdentifier: IAM_PASSWORD_POLICY
            Type: AWS::Config::ConfigRule
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Config Conformance Packs using the `name`. For example:

```sh
 $ pulumi import aws:cfg/conformancePack:ConformancePack example example
```
 