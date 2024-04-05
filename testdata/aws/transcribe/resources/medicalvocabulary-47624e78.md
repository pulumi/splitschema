Resource for managing an AWS Transcribe MedicalVocabulary.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {forceDestroy: true});
const object = new aws.s3.BucketObjectv2("object", {
    bucket: exampleBucketV2.id,
    key: "transcribe/test1.txt",
    source: new pulumi.asset.FileAsset("test.txt"),
});
const exampleMedicalVocabulary = new aws.transcribe.MedicalVocabulary("exampleMedicalVocabulary", {
    vocabularyName: "example",
    languageCode: "en-US",
    vocabularyFileUri: pulumi.interpolate`s3://${exampleBucketV2.id}/${object.key}`,
    tags: {
        tag1: "value1",
        tag2: "value3",
    },
}, {
    dependsOn: [object],
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2", force_destroy=True)
object = aws.s3.BucketObjectv2("object",
    bucket=example_bucket_v2.id,
    key="transcribe/test1.txt",
    source=pulumi.FileAsset("test.txt"))
example_medical_vocabulary = aws.transcribe.MedicalVocabulary("exampleMedicalVocabulary",
    vocabulary_name="example",
    language_code="en-US",
    vocabulary_file_uri=pulumi.Output.all(example_bucket_v2.id, object.key).apply(lambda id, key: f"s3://{id}/{key}"),
    tags={
        "tag1": "value1",
        "tag2": "value3",
    },
    opts=pulumi.ResourceOptions(depends_on=[object]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2", new()
    {
        ForceDestroy = true,
    });

    var @object = new Aws.S3.BucketObjectv2("object", new()
    {
        Bucket = exampleBucketV2.Id,
        Key = "transcribe/test1.txt",
        Source = new FileAsset("test.txt"),
    });

    var exampleMedicalVocabulary = new Aws.Transcribe.MedicalVocabulary("exampleMedicalVocabulary", new()
    {
        VocabularyName = "example",
        LanguageCode = "en-US",
        VocabularyFileUri = Output.Tuple(exampleBucketV2.Id, @object.Key).Apply(values =>
        {
            var id = values.Item1;
            var key = values.Item2;
            return $"s3://{id}/{key}";
        }),
        Tags = 
        {
            { "tag1", "value1" },
            { "tag2", "value3" },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            @object,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transcribe"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", &s3.BucketV2Args{
			ForceDestroy: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		object, err := s3.NewBucketObjectv2(ctx, "object", &s3.BucketObjectv2Args{
			Bucket: exampleBucketV2.ID(),
			Key:    pulumi.String("transcribe/test1.txt"),
			Source: pulumi.NewFileAsset("test.txt"),
		})
		if err != nil {
			return err
		}
		_, err = transcribe.NewMedicalVocabulary(ctx, "exampleMedicalVocabulary", &transcribe.MedicalVocabularyArgs{
			VocabularyName: pulumi.String("example"),
			LanguageCode:   pulumi.String("en-US"),
			VocabularyFileUri: pulumi.All(exampleBucketV2.ID(), object.Key).ApplyT(func(_args []interface{}) (string, error) {
				id := _args[0].(string)
				key := _args[1].(string)
				return fmt.Sprintf("s3://%v/%v", id, key), nil
			}).(pulumi.StringOutput),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value3"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			object,
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
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.s3.BucketObjectv2Args;
import com.pulumi.aws.transcribe.MedicalVocabulary;
import com.pulumi.aws.transcribe.MedicalVocabularyArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2", BucketV2Args.builder()        
            .forceDestroy(true)
            .build());

        var object = new BucketObjectv2("object", BucketObjectv2Args.builder()        
            .bucket(exampleBucketV2.id())
            .key("transcribe/test1.txt")
            .source(new FileAsset("test.txt"))
            .build());

        var exampleMedicalVocabulary = new MedicalVocabulary("exampleMedicalVocabulary", MedicalVocabularyArgs.builder()        
            .vocabularyName("example")
            .languageCode("en-US")
            .vocabularyFileUri(Output.tuple(exampleBucketV2.id(), object.key()).applyValue(values -> {
                var id = values.t1;
                var key = values.t2;
                return String.format("s3://%s/%s", id,key);
            }))
            .tags(Map.ofEntries(
                Map.entry("tag1", "value1"),
                Map.entry("tag2", "value3")
            ))
            .build(), CustomResourceOptions.builder()
                .dependsOn(object)
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
  object:
    type: aws:s3:BucketObjectv2
    properties:
      bucket: ${exampleBucketV2.id}
      key: transcribe/test1.txt
      source:
        fn::FileAsset: test.txt
  exampleMedicalVocabulary:
    type: aws:transcribe:MedicalVocabulary
    properties:
      vocabularyName: example
      languageCode: en-US
      vocabularyFileUri: s3://${exampleBucketV2.id}/${object.key}
      tags:
        tag1: value1
        tag2: value3
    options:
      dependson:
        - ${object}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transcribe MedicalVocabulary using the `vocabulary_name`. For example:

```sh
 $ pulumi import aws:transcribe/medicalVocabulary:MedicalVocabulary example example-name
```
 