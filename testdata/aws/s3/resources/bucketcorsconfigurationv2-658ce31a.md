Provides an S3 bucket CORS configuration resource. For more information about CORS, go to [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/userguide/cors.html) in the Amazon S3 User Guide.

> **NOTE:** S3 Buckets only support a single CORS configuration. Declaring multiple `aws.s3.BucketCorsConfigurationV2` resources to the same S3 Bucket will cause a perpetual difference in configuration.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketCorsConfigurationV2 = new aws.s3.BucketCorsConfigurationV2("exampleBucketCorsConfigurationV2", {
    bucket: exampleBucketV2.id,
    corsRules: [
        {
            allowedHeaders: ["*"],
            allowedMethods: [
                "PUT",
                "POST",
            ],
            allowedOrigins: ["https://s3-website-test.domain.example"],
            exposeHeaders: ["ETag"],
            maxAgeSeconds: 3000,
        },
        {
            allowedMethods: ["GET"],
            allowedOrigins: ["*"],
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_cors_configuration_v2 = aws.s3.BucketCorsConfigurationV2("exampleBucketCorsConfigurationV2",
    bucket=example_bucket_v2.id,
    cors_rules=[
        aws.s3.BucketCorsConfigurationV2CorsRuleArgs(
            allowed_headers=["*"],
            allowed_methods=[
                "PUT",
                "POST",
            ],
            allowed_origins=["https://s3-website-test.domain.example"],
            expose_headers=["ETag"],
            max_age_seconds=3000,
        ),
        aws.s3.BucketCorsConfigurationV2CorsRuleArgs(
            allowed_methods=["GET"],
            allowed_origins=["*"],
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketCorsConfigurationV2 = new Aws.S3.BucketCorsConfigurationV2("exampleBucketCorsConfigurationV2", new()
    {
        Bucket = exampleBucketV2.Id,
        CorsRules = new[]
        {
            new Aws.S3.Inputs.BucketCorsConfigurationV2CorsRuleArgs
            {
                AllowedHeaders = new[]
                {
                    "*",
                },
                AllowedMethods = new[]
                {
                    "PUT",
                    "POST",
                },
                AllowedOrigins = new[]
                {
                    "https://s3-website-test.domain.example",
                },
                ExposeHeaders = new[]
                {
                    "ETag",
                },
                MaxAgeSeconds = 3000,
            },
            new Aws.S3.Inputs.BucketCorsConfigurationV2CorsRuleArgs
            {
                AllowedMethods = new[]
                {
                    "GET",
                },
                AllowedOrigins = new[]
                {
                    "*",
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
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketCorsConfigurationV2(ctx, "exampleBucketCorsConfigurationV2", &s3.BucketCorsConfigurationV2Args{
			Bucket: exampleBucketV2.ID(),
			CorsRules: s3.BucketCorsConfigurationV2CorsRuleArray{
				&s3.BucketCorsConfigurationV2CorsRuleArgs{
					AllowedHeaders: pulumi.StringArray{
						pulumi.String("*"),
					},
					AllowedMethods: pulumi.StringArray{
						pulumi.String("PUT"),
						pulumi.String("POST"),
					},
					AllowedOrigins: pulumi.StringArray{
						pulumi.String("https://s3-website-test.domain.example"),
					},
					ExposeHeaders: pulumi.StringArray{
						pulumi.String("ETag"),
					},
					MaxAgeSeconds: pulumi.Int(3000),
				},
				&s3.BucketCorsConfigurationV2CorsRuleArgs{
					AllowedMethods: pulumi.StringArray{
						pulumi.String("GET"),
					},
					AllowedOrigins: pulumi.StringArray{
						pulumi.String("*"),
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
import com.pulumi.aws.s3.BucketCorsConfigurationV2;
import com.pulumi.aws.s3.BucketCorsConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketCorsConfigurationV2CorsRuleArgs;
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

        var exampleBucketCorsConfigurationV2 = new BucketCorsConfigurationV2("exampleBucketCorsConfigurationV2", BucketCorsConfigurationV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .corsRules(            
                BucketCorsConfigurationV2CorsRuleArgs.builder()
                    .allowedHeaders("*")
                    .allowedMethods(                    
                        "PUT",
                        "POST")
                    .allowedOrigins("https://s3-website-test.domain.example")
                    .exposeHeaders("ETag")
                    .maxAgeSeconds(3000)
                    .build(),
                BucketCorsConfigurationV2CorsRuleArgs.builder()
                    .allowedMethods("GET")
                    .allowedOrigins("*")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketCorsConfigurationV2:
    type: aws:s3:BucketCorsConfigurationV2
    properties:
      bucket: ${exampleBucketV2.id}
      corsRules:
        - allowedHeaders:
            - '*'
          allowedMethods:
            - PUT
            - POST
          allowedOrigins:
            - https://s3-website-test.domain.example
          exposeHeaders:
            - ETag
          maxAgeSeconds: 3000
        - allowedMethods:
            - GET
          allowedOrigins:
            - '*'
```
{{% /example %}}
{{% /examples %}}

## Import

If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

__Using `pulumi import` to import__ S3 bucket CORS configuration using the `bucket` or using the `bucket` and `expected_bucket_owner` separated by a comma (`,`). For example:

If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`:

```sh
 $ pulumi import aws:s3/bucketCorsConfigurationV2:BucketCorsConfigurationV2 example bucket-name
```
 If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

```sh
 $ pulumi import aws:s3/bucketCorsConfigurationV2:BucketCorsConfigurationV2 example bucket-name,123456789012
```
 