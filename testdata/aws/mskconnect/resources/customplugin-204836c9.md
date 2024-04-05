Provides an Amazon MSK Connect Custom Plugin Resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketObjectv2 = new aws.s3.BucketObjectv2("exampleBucketObjectv2", {
    bucket: exampleBucketV2.id,
    key: "debezium.zip",
    source: new pulumi.asset.FileAsset("debezium.zip"),
});
const exampleCustomPlugin = new aws.mskconnect.CustomPlugin("exampleCustomPlugin", {
    contentType: "ZIP",
    location: {
        s3: {
            bucketArn: exampleBucketV2.arn,
            fileKey: exampleBucketObjectv2.key,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_objectv2 = aws.s3.BucketObjectv2("exampleBucketObjectv2",
    bucket=example_bucket_v2.id,
    key="debezium.zip",
    source=pulumi.FileAsset("debezium.zip"))
example_custom_plugin = aws.mskconnect.CustomPlugin("exampleCustomPlugin",
    content_type="ZIP",
    location=aws.mskconnect.CustomPluginLocationArgs(
        s3=aws.mskconnect.CustomPluginLocationS3Args(
            bucket_arn=example_bucket_v2.arn,
            file_key=example_bucket_objectv2.key,
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
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketObjectv2 = new Aws.S3.BucketObjectv2("exampleBucketObjectv2", new()
    {
        Bucket = exampleBucketV2.Id,
        Key = "debezium.zip",
        Source = new FileAsset("debezium.zip"),
    });

    var exampleCustomPlugin = new Aws.MskConnect.CustomPlugin("exampleCustomPlugin", new()
    {
        ContentType = "ZIP",
        Location = new Aws.MskConnect.Inputs.CustomPluginLocationArgs
        {
            S3 = new Aws.MskConnect.Inputs.CustomPluginLocationS3Args
            {
                BucketArn = exampleBucketV2.Arn,
                FileKey = exampleBucketObjectv2.Key,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mskconnect"
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
			Key:    pulumi.String("debezium.zip"),
			Source: pulumi.NewFileAsset("debezium.zip"),
		})
		if err != nil {
			return err
		}
		_, err = mskconnect.NewCustomPlugin(ctx, "exampleCustomPlugin", &mskconnect.CustomPluginArgs{
			ContentType: pulumi.String("ZIP"),
			Location: &mskconnect.CustomPluginLocationArgs{
				S3: &mskconnect.CustomPluginLocationS3Args{
					BucketArn: exampleBucketV2.Arn,
					FileKey:   exampleBucketObjectv2.Key,
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
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.s3.BucketObjectv2Args;
import com.pulumi.aws.mskconnect.CustomPlugin;
import com.pulumi.aws.mskconnect.CustomPluginArgs;
import com.pulumi.aws.mskconnect.inputs.CustomPluginLocationArgs;
import com.pulumi.aws.mskconnect.inputs.CustomPluginLocationS3Args;
import com.pulumi.asset.FileAsset;
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
            .key("debezium.zip")
            .source(new FileAsset("debezium.zip"))
            .build());

        var exampleCustomPlugin = new CustomPlugin("exampleCustomPlugin", CustomPluginArgs.builder()        
            .contentType("ZIP")
            .location(CustomPluginLocationArgs.builder()
                .s3(CustomPluginLocationS3Args.builder()
                    .bucketArn(exampleBucketV2.arn())
                    .fileKey(exampleBucketObjectv2.key())
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
  exampleBucketObjectv2:
    type: aws:s3:BucketObjectv2
    properties:
      bucket: ${exampleBucketV2.id}
      key: debezium.zip
      source:
        fn::FileAsset: debezium.zip
  exampleCustomPlugin:
    type: aws:mskconnect:CustomPlugin
    properties:
      contentType: ZIP
      location:
        s3:
          bucketArn: ${exampleBucketV2.arn}
          fileKey: ${exampleBucketObjectv2.key}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MSK Connect Custom Plugin using the plugin's `arn`. For example:

```sh
 $ pulumi import aws:mskconnect/customPlugin:CustomPlugin example 'arn:aws:kafkaconnect:eu-central-1:123456789012:custom-plugin/debezium-example/abcdefgh-1234-5678-9abc-defghijklmno-4'
```
 