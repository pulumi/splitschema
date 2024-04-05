Provides a resource to manage an S3 Object Lambda Access Point.
An Object Lambda access point is associated with exactly one standard access point and thus one Amazon S3 bucket.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleAccessPoint = new aws.s3.AccessPoint("exampleAccessPoint", {bucket: exampleBucketV2.id});
const exampleObjectLambdaAccessPoint = new aws.s3control.ObjectLambdaAccessPoint("exampleObjectLambdaAccessPoint", {configuration: {
    supportingAccessPoint: exampleAccessPoint.arn,
    transformationConfigurations: [{
        actions: ["GetObject"],
        contentTransformation: {
            awsLambda: {
                functionArn: aws_lambda_function.example.arn,
            },
        },
    }],
}});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_access_point = aws.s3.AccessPoint("exampleAccessPoint", bucket=example_bucket_v2.id)
example_object_lambda_access_point = aws.s3control.ObjectLambdaAccessPoint("exampleObjectLambdaAccessPoint", configuration=aws.s3control.ObjectLambdaAccessPointConfigurationArgs(
    supporting_access_point=example_access_point.arn,
    transformation_configurations=[aws.s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationArgs(
        actions=["GetObject"],
        content_transformation=aws.s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationArgs(
            aws_lambda=aws.s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambdaArgs(
                function_arn=aws_lambda_function["example"]["arn"],
            ),
        ),
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
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleAccessPoint = new Aws.S3.AccessPoint("exampleAccessPoint", new()
    {
        Bucket = exampleBucketV2.Id,
    });

    var exampleObjectLambdaAccessPoint = new Aws.S3Control.ObjectLambdaAccessPoint("exampleObjectLambdaAccessPoint", new()
    {
        Configuration = new Aws.S3Control.Inputs.ObjectLambdaAccessPointConfigurationArgs
        {
            SupportingAccessPoint = exampleAccessPoint.Arn,
            TransformationConfigurations = new[]
            {
                new Aws.S3Control.Inputs.ObjectLambdaAccessPointConfigurationTransformationConfigurationArgs
                {
                    Actions = new[]
                    {
                        "GetObject",
                    },
                    ContentTransformation = new Aws.S3Control.Inputs.ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationArgs
                    {
                        AwsLambda = new Aws.S3Control.Inputs.ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambdaArgs
                        {
                            FunctionArn = aws_lambda_function.Example.Arn,
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		exampleAccessPoint, err := s3.NewAccessPoint(ctx, "exampleAccessPoint", &s3.AccessPointArgs{
			Bucket: exampleBucketV2.ID(),
		})
		if err != nil {
			return err
		}
		_, err = s3control.NewObjectLambdaAccessPoint(ctx, "exampleObjectLambdaAccessPoint", &s3control.ObjectLambdaAccessPointArgs{
			Configuration: &s3control.ObjectLambdaAccessPointConfigurationArgs{
				SupportingAccessPoint: exampleAccessPoint.Arn,
				TransformationConfigurations: s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationArray{
					&s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationArgs{
						Actions: pulumi.StringArray{
							pulumi.String("GetObject"),
						},
						ContentTransformation: &s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationArgs{
							AwsLambda: &s3control.ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambdaArgs{
								FunctionArn: pulumi.Any(aws_lambda_function.Example.Arn),
							},
						},
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.AccessPoint;
import com.pulumi.aws.s3.AccessPointArgs;
import com.pulumi.aws.s3control.ObjectLambdaAccessPoint;
import com.pulumi.aws.s3control.ObjectLambdaAccessPointArgs;
import com.pulumi.aws.s3control.inputs.ObjectLambdaAccessPointConfigurationArgs;
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

        var exampleAccessPoint = new AccessPoint("exampleAccessPoint", AccessPointArgs.builder()        
            .bucket(exampleBucketV2.id())
            .build());

        var exampleObjectLambdaAccessPoint = new ObjectLambdaAccessPoint("exampleObjectLambdaAccessPoint", ObjectLambdaAccessPointArgs.builder()        
            .configuration(ObjectLambdaAccessPointConfigurationArgs.builder()
                .supportingAccessPoint(exampleAccessPoint.arn())
                .transformationConfigurations(ObjectLambdaAccessPointConfigurationTransformationConfigurationArgs.builder()
                    .actions("GetObject")
                    .contentTransformation(ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationArgs.builder()
                        .awsLambda(ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambdaArgs.builder()
                            .functionArn(aws_lambda_function.example().arn())
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleAccessPoint:
    type: aws:s3:AccessPoint
    properties:
      bucket: ${exampleBucketV2.id}
  exampleObjectLambdaAccessPoint:
    type: aws:s3control:ObjectLambdaAccessPoint
    properties:
      configuration:
        supportingAccessPoint: ${exampleAccessPoint.arn}
        transformationConfigurations:
          - actions:
              - GetObject
            contentTransformation:
              awsLambda:
                functionArn: ${aws_lambda_function.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Object Lambda Access Points using the `account_id` and `name`, separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:s3control/objectLambdaAccessPoint:ObjectLambdaAccessPoint example 123456789012:example
```
 