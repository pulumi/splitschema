Provides a resource to manage an S3 Object Lambda Access Point resource policy.

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
const exampleObjectLambdaAccessPointPolicy = new aws.s3control.ObjectLambdaAccessPointPolicy("exampleObjectLambdaAccessPointPolicy", {policy: exampleObjectLambdaAccessPoint.arn.apply(arn => JSON.stringify({
    Version: "2008-10-17",
    Statement: [{
        Effect: "Allow",
        Action: "s3-object-lambda:GetObject",
        Principal: {
            AWS: data.aws_caller_identity.current.account_id,
        },
        Resource: arn,
    }],
}))});
```
```python
import pulumi
import json
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
example_object_lambda_access_point_policy = aws.s3control.ObjectLambdaAccessPointPolicy("exampleObjectLambdaAccessPointPolicy", policy=example_object_lambda_access_point.arn.apply(lambda arn: json.dumps({
    "Version": "2008-10-17",
    "Statement": [{
        "Effect": "Allow",
        "Action": "s3-object-lambda:GetObject",
        "Principal": {
            "AWS": data["aws_caller_identity"]["current"]["account_id"],
        },
        "Resource": arn,
    }],
})))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
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

    var exampleObjectLambdaAccessPointPolicy = new Aws.S3Control.ObjectLambdaAccessPointPolicy("exampleObjectLambdaAccessPointPolicy", new()
    {
        Policy = exampleObjectLambdaAccessPoint.Arn.Apply(arn => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2008-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Effect"] = "Allow",
                    ["Action"] = "s3-object-lambda:GetObject",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["AWS"] = data.Aws_caller_identity.Current.Account_id,
                    },
                    ["Resource"] = arn,
                },
            },
        })),
    });

});
```
```go
package main

import (
	"encoding/json"

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
		exampleObjectLambdaAccessPoint, err := s3control.NewObjectLambdaAccessPoint(ctx, "exampleObjectLambdaAccessPoint", &s3control.ObjectLambdaAccessPointArgs{
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
		_, err = s3control.NewObjectLambdaAccessPointPolicy(ctx, "exampleObjectLambdaAccessPointPolicy", &s3control.ObjectLambdaAccessPointPolicyArgs{
			Policy: exampleObjectLambdaAccessPoint.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"Version": "2008-10-17",
					"Statement": []map[string]interface{}{
						map[string]interface{}{
							"Effect": "Allow",
							"Action": "s3-object-lambda:GetObject",
							"Principal": map[string]interface{}{
								"AWS": data.Aws_caller_identity.Current.Account_id,
							},
							"Resource": arn,
						},
					},
				})
				if err != nil {
					return _zero, err
				}
				json0 := string(tmpJSON0)
				return pulumi.String(json0), nil
			}).(pulumi.StringOutput),
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
import com.pulumi.aws.s3control.ObjectLambdaAccessPointPolicy;
import com.pulumi.aws.s3control.ObjectLambdaAccessPointPolicyArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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

        var exampleObjectLambdaAccessPointPolicy = new ObjectLambdaAccessPointPolicy("exampleObjectLambdaAccessPointPolicy", ObjectLambdaAccessPointPolicyArgs.builder()        
            .policy(exampleObjectLambdaAccessPoint.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("Version", "2008-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Action", "s3-object-lambda:GetObject"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", data.aws_caller_identity().current().account_id())
                        )),
                        jsonProperty("Resource", arn)
                    )))
                ))))
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
  exampleObjectLambdaAccessPointPolicy:
    type: aws:s3control:ObjectLambdaAccessPointPolicy
    properties:
      policy:
        fn::toJSON:
          Version: 2008-10-17
          Statement:
            - Effect: Allow
              Action: s3-object-lambda:GetObject
              Principal:
                AWS: ${data.aws_caller_identity.current.account_id}
              Resource: ${exampleObjectLambdaAccessPoint.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Object Lambda Access Point policies using the `account_id` and `name`, separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:s3control/objectLambdaAccessPointPolicy:ObjectLambdaAccessPointPolicy example 123456789012:example
```
 