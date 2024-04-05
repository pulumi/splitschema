Manages Bedrock model invocation logging configuration.

> Model invocation logging is configured per AWS region. To avoid overwriting settings, this resource should not be defined in multiple configurations.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
import com.pulumi.aws.bedrockmodel.InvocationLoggingConfiguration;
import com.pulumi.aws.bedrockmodel.InvocationLoggingConfigurationArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        var exampleBucketV2 = new BucketV2("exampleBucketV2", BucketV2Args.builder()        
            .forceDestroy(true)
            .build());

        var exampleBucketPolicy = new BucketPolicy("exampleBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(exampleBucketV2.bucket())
            .policy(exampleBucketV2.arn().applyValue(arn -> """
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "bedrock.amazonaws.com"
      },
      "Action": [
        "s3:*"
      ],
      "Resource": [
        "%s/*"
      ],
      "Condition": {
        "StringEquals": {
          "aws:SourceAccount": "%s"
        },
        "ArnLike": {
          "aws:SourceArn": "arn:aws:bedrock:us-east-1:%s:*"
        }
      }
    }
  ]
}
", arn,current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))))
            .build());

        var exampleInvocationLoggingConfiguration = new InvocationLoggingConfiguration("exampleInvocationLoggingConfiguration", InvocationLoggingConfigurationArgs.builder()        
            .loggingConfig(InvocationLoggingConfigurationLoggingConfigArgs.builder()
                .embeddingDataDeliveryEnabled(true)
                .imageDataDeliveryEnabled(true)
                .textDataDeliveryEnabled(true)
                .s3Config(InvocationLoggingConfigurationLoggingConfigS3ConfigArgs.builder()
                    .bucketName(exampleBucketV2.id())
                    .keyPrefix("bedrock")
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleBucketPolicy)
                .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
    properties:
      forceDestroy: true
  exampleBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${exampleBucketV2.bucket}
      policy: |
        {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": "bedrock.amazonaws.com"
              },
              "Action": [
                "s3:*"
              ],
              "Resource": [
                "${exampleBucketV2.arn}/*"
              ],
              "Condition": {
                "StringEquals": {
                  "aws:SourceAccount": "${current.accountId}"
                },
                "ArnLike": {
                  "aws:SourceArn": "arn:aws:bedrock:us-east-1:${current.accountId}:*"
                }
              }
            }
          ]
        }
  exampleInvocationLoggingConfiguration:
    type: aws:bedrockmodel:InvocationLoggingConfiguration
    properties:
      loggingConfig:
        - embeddingDataDeliveryEnabled: true
          imageDataDeliveryEnabled: true
          textDataDeliveryEnabled: true
          s3Config:
            - bucketName: ${exampleBucketV2.id}
              keyPrefix: bedrock
    options:
      dependson:
        - ${exampleBucketPolicy}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Bedrock custom model using the `id` set to the AWS region. For example:

```sh
 $ pulumi import aws:bedrockmodel/invocationLoggingConfiguration:InvocationLoggingConfiguration my_config us-east-1
```
 