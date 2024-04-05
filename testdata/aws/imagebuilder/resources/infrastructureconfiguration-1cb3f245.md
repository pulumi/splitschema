Manages an Image Builder Infrastructure Configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.imagebuilder.InfrastructureConfiguration("example", {
    description: "example description",
    instanceProfileName: aws_iam_instance_profile.example.name,
    instanceTypes: [
        "t2.nano",
        "t3.micro",
    ],
    keyPair: aws_key_pair.example.key_name,
    securityGroupIds: [aws_security_group.example.id],
    snsTopicArn: aws_sns_topic.example.arn,
    subnetId: aws_subnet.main.id,
    terminateInstanceOnFailure: true,
    logging: {
        s3Logs: {
            s3BucketName: aws_s3_bucket.example.bucket,
            s3KeyPrefix: "logs",
        },
    },
    tags: {
        foo: "bar",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.InfrastructureConfiguration("example",
    description="example description",
    instance_profile_name=aws_iam_instance_profile["example"]["name"],
    instance_types=[
        "t2.nano",
        "t3.micro",
    ],
    key_pair=aws_key_pair["example"]["key_name"],
    security_group_ids=[aws_security_group["example"]["id"]],
    sns_topic_arn=aws_sns_topic["example"]["arn"],
    subnet_id=aws_subnet["main"]["id"],
    terminate_instance_on_failure=True,
    logging=aws.imagebuilder.InfrastructureConfigurationLoggingArgs(
        s3_logs=aws.imagebuilder.InfrastructureConfigurationLoggingS3LogsArgs(
            s3_bucket_name=aws_s3_bucket["example"]["bucket"],
            s3_key_prefix="logs",
        ),
    ),
    tags={
        "foo": "bar",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ImageBuilder.InfrastructureConfiguration("example", new()
    {
        Description = "example description",
        InstanceProfileName = aws_iam_instance_profile.Example.Name,
        InstanceTypes = new[]
        {
            "t2.nano",
            "t3.micro",
        },
        KeyPair = aws_key_pair.Example.Key_name,
        SecurityGroupIds = new[]
        {
            aws_security_group.Example.Id,
        },
        SnsTopicArn = aws_sns_topic.Example.Arn,
        SubnetId = aws_subnet.Main.Id,
        TerminateInstanceOnFailure = true,
        Logging = new Aws.ImageBuilder.Inputs.InfrastructureConfigurationLoggingArgs
        {
            S3Logs = new Aws.ImageBuilder.Inputs.InfrastructureConfigurationLoggingS3LogsArgs
            {
                S3BucketName = aws_s3_bucket.Example.Bucket,
                S3KeyPrefix = "logs",
            },
        },
        Tags = 
        {
            { "foo", "bar" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/imagebuilder"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := imagebuilder.NewInfrastructureConfiguration(ctx, "example", &imagebuilder.InfrastructureConfigurationArgs{
			Description:         pulumi.String("example description"),
			InstanceProfileName: pulumi.Any(aws_iam_instance_profile.Example.Name),
			InstanceTypes: pulumi.StringArray{
				pulumi.String("t2.nano"),
				pulumi.String("t3.micro"),
			},
			KeyPair: pulumi.Any(aws_key_pair.Example.Key_name),
			SecurityGroupIds: pulumi.StringArray{
				aws_security_group.Example.Id,
			},
			SnsTopicArn:                pulumi.Any(aws_sns_topic.Example.Arn),
			SubnetId:                   pulumi.Any(aws_subnet.Main.Id),
			TerminateInstanceOnFailure: pulumi.Bool(true),
			Logging: &imagebuilder.InfrastructureConfigurationLoggingArgs{
				S3Logs: &imagebuilder.InfrastructureConfigurationLoggingS3LogsArgs{
					S3BucketName: pulumi.Any(aws_s3_bucket.Example.Bucket),
					S3KeyPrefix:  pulumi.String("logs"),
				},
			},
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
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
import com.pulumi.aws.imagebuilder.InfrastructureConfiguration;
import com.pulumi.aws.imagebuilder.InfrastructureConfigurationArgs;
import com.pulumi.aws.imagebuilder.inputs.InfrastructureConfigurationLoggingArgs;
import com.pulumi.aws.imagebuilder.inputs.InfrastructureConfigurationLoggingS3LogsArgs;
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
        var example = new InfrastructureConfiguration("example", InfrastructureConfigurationArgs.builder()        
            .description("example description")
            .instanceProfileName(aws_iam_instance_profile.example().name())
            .instanceTypes(            
                "t2.nano",
                "t3.micro")
            .keyPair(aws_key_pair.example().key_name())
            .securityGroupIds(aws_security_group.example().id())
            .snsTopicArn(aws_sns_topic.example().arn())
            .subnetId(aws_subnet.main().id())
            .terminateInstanceOnFailure(true)
            .logging(InfrastructureConfigurationLoggingArgs.builder()
                .s3Logs(InfrastructureConfigurationLoggingS3LogsArgs.builder()
                    .s3BucketName(aws_s3_bucket.example().bucket())
                    .s3KeyPrefix("logs")
                    .build())
                .build())
            .tags(Map.of("foo", "bar"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:imagebuilder:InfrastructureConfiguration
    properties:
      description: example description
      instanceProfileName: ${aws_iam_instance_profile.example.name}
      instanceTypes:
        - t2.nano
        - t3.micro
      keyPair: ${aws_key_pair.example.key_name}
      securityGroupIds:
        - ${aws_security_group.example.id}
      snsTopicArn: ${aws_sns_topic.example.arn}
      subnetId: ${aws_subnet.main.id}
      terminateInstanceOnFailure: true
      logging:
        s3Logs:
          s3BucketName: ${aws_s3_bucket.example.bucket}
          s3KeyPrefix: logs
      tags:
        foo: bar
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_imagebuilder_infrastructure_configuration` using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration example arn:aws:imagebuilder:us-east-1:123456789012:infrastructure-configuration/example
```
 