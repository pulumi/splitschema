Manages a version of a CloudFormation Type.



{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudformation.CloudFormationType("example", {
    schemaHandlerPackage: `s3://${aws_s3_object.example.bucket}/${aws_s3_object.example.key}`,
    type: "RESOURCE",
    typeName: "ExampleCompany::ExampleService::ExampleResource",
    loggingConfig: {
        logGroupName: aws_cloudwatch_log_group.example.name,
        logRoleArn: aws_iam_role.example.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudformation.CloudFormationType("example",
    schema_handler_package=f"s3://{aws_s3_object['example']['bucket']}/{aws_s3_object['example']['key']}",
    type="RESOURCE",
    type_name="ExampleCompany::ExampleService::ExampleResource",
    logging_config=aws.cloudformation.CloudFormationTypeLoggingConfigArgs(
        log_group_name=aws_cloudwatch_log_group["example"]["name"],
        log_role_arn=aws_iam_role["example"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudFormation.CloudFormationType("example", new()
    {
        SchemaHandlerPackage = $"s3://{aws_s3_object.Example.Bucket}/{aws_s3_object.Example.Key}",
        Type = "RESOURCE",
        TypeName = "ExampleCompany::ExampleService::ExampleResource",
        LoggingConfig = new Aws.CloudFormation.Inputs.CloudFormationTypeLoggingConfigArgs
        {
            LogGroupName = aws_cloudwatch_log_group.Example.Name,
            LogRoleArn = aws_iam_role.Example.Arn,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudformation.NewCloudFormationType(ctx, "example", &cloudformation.CloudFormationTypeArgs{
			SchemaHandlerPackage: pulumi.String(fmt.Sprintf("s3://%v/%v", aws_s3_object.Example.Bucket, aws_s3_object.Example.Key)),
			Type:                 pulumi.String("RESOURCE"),
			TypeName:             pulumi.String("ExampleCompany::ExampleService::ExampleResource"),
			LoggingConfig: &cloudformation.CloudFormationTypeLoggingConfigArgs{
				LogGroupName: pulumi.Any(aws_cloudwatch_log_group.Example.Name),
				LogRoleArn:   pulumi.Any(aws_iam_role.Example.Arn),
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
import com.pulumi.aws.cloudformation.CloudFormationType;
import com.pulumi.aws.cloudformation.CloudFormationTypeArgs;
import com.pulumi.aws.cloudformation.inputs.CloudFormationTypeLoggingConfigArgs;
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
        var example = new CloudFormationType("example", CloudFormationTypeArgs.builder()        
            .schemaHandlerPackage(String.format("s3://%s/%s", aws_s3_object.example().bucket(),aws_s3_object.example().key()))
            .type("RESOURCE")
            .typeName("ExampleCompany::ExampleService::ExampleResource")
            .loggingConfig(CloudFormationTypeLoggingConfigArgs.builder()
                .logGroupName(aws_cloudwatch_log_group.example().name())
                .logRoleArn(aws_iam_role.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudformation:CloudFormationType
    properties:
      schemaHandlerPackage: s3://${aws_s3_object.example.bucket}/${aws_s3_object.example.key}
      type: RESOURCE
      typeName: ExampleCompany::ExampleService::ExampleResource
      loggingConfig:
        logGroupName: ${aws_cloudwatch_log_group.example.name}
        logRoleArn: ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_cloudformation_type` using the type version Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:cloudformation/cloudFormationType:CloudFormationType example arn:aws:cloudformation:us-east-1:123456789012:type/resource/ExampleCompany-ExampleService-ExampleType/1
```
 