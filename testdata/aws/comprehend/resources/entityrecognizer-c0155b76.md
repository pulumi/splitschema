Resource for managing an AWS Comprehend Entity Recognizer.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const documents = new aws.s3.BucketObjectv2("documents", {});
// ...
const entities = new aws.s3.BucketObjectv2("entities", {});
// ...
const example = new aws.comprehend.EntityRecognizer("example", {
    dataAccessRoleArn: aws_iam_role.example.arn,
    languageCode: "en",
    inputDataConfig: {
        entityTypes: [
            {
                type: "ENTITY_1",
            },
            {
                type: "ENTITY_2",
            },
        ],
        documents: {
            s3Uri: pulumi.interpolate`s3://${aws_s3_bucket.documents.bucket}/${documents.id}`,
        },
        entityList: {
            s3Uri: pulumi.interpolate`s3://${aws_s3_bucket.entities.bucket}/${entities.id}`,
        },
    },
}, {
    dependsOn: [aws_iam_role_policy.example],
});
```
```python
import pulumi
import pulumi_aws as aws

documents = aws.s3.BucketObjectv2("documents")
# ...
entities = aws.s3.BucketObjectv2("entities")
# ...
example = aws.comprehend.EntityRecognizer("example",
    data_access_role_arn=aws_iam_role["example"]["arn"],
    language_code="en",
    input_data_config=aws.comprehend.EntityRecognizerInputDataConfigArgs(
        entity_types=[
            aws.comprehend.EntityRecognizerInputDataConfigEntityTypeArgs(
                type="ENTITY_1",
            ),
            aws.comprehend.EntityRecognizerInputDataConfigEntityTypeArgs(
                type="ENTITY_2",
            ),
        ],
        documents=aws.comprehend.EntityRecognizerInputDataConfigDocumentsArgs(
            s3_uri=documents.id.apply(lambda id: f"s3://{aws_s3_bucket['documents']['bucket']}/{id}"),
        ),
        entity_list=aws.comprehend.EntityRecognizerInputDataConfigEntityListArgs(
            s3_uri=entities.id.apply(lambda id: f"s3://{aws_s3_bucket['entities']['bucket']}/{id}"),
        ),
    ),
    opts=pulumi.ResourceOptions(depends_on=[aws_iam_role_policy["example"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var documents = new Aws.S3.BucketObjectv2("documents");

    // ...
    var entities = new Aws.S3.BucketObjectv2("entities");

    // ...
    var example = new Aws.Comprehend.EntityRecognizer("example", new()
    {
        DataAccessRoleArn = aws_iam_role.Example.Arn,
        LanguageCode = "en",
        InputDataConfig = new Aws.Comprehend.Inputs.EntityRecognizerInputDataConfigArgs
        {
            EntityTypes = new[]
            {
                new Aws.Comprehend.Inputs.EntityRecognizerInputDataConfigEntityTypeArgs
                {
                    Type = "ENTITY_1",
                },
                new Aws.Comprehend.Inputs.EntityRecognizerInputDataConfigEntityTypeArgs
                {
                    Type = "ENTITY_2",
                },
            },
            Documents = new Aws.Comprehend.Inputs.EntityRecognizerInputDataConfigDocumentsArgs
            {
                S3Uri = documents.Id.Apply(id => $"s3://{aws_s3_bucket.Documents.Bucket}/{id}"),
            },
            EntityList = new Aws.Comprehend.Inputs.EntityRecognizerInputDataConfigEntityListArgs
            {
                S3Uri = entities.Id.Apply(id => $"s3://{aws_s3_bucket.Entities.Bucket}/{id}"),
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_role_policy.Example,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/comprehend"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		documents, err := s3.NewBucketObjectv2(ctx, "documents", nil)
		if err != nil {
			return err
		}
		entities, err := s3.NewBucketObjectv2(ctx, "entities", nil)
		if err != nil {
			return err
		}
		_, err = comprehend.NewEntityRecognizer(ctx, "example", &comprehend.EntityRecognizerArgs{
			DataAccessRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			LanguageCode:      pulumi.String("en"),
			InputDataConfig: &comprehend.EntityRecognizerInputDataConfigArgs{
				EntityTypes: comprehend.EntityRecognizerInputDataConfigEntityTypeArray{
					&comprehend.EntityRecognizerInputDataConfigEntityTypeArgs{
						Type: pulumi.String("ENTITY_1"),
					},
					&comprehend.EntityRecognizerInputDataConfigEntityTypeArgs{
						Type: pulumi.String("ENTITY_2"),
					},
				},
				Documents: &comprehend.EntityRecognizerInputDataConfigDocumentsArgs{
					S3Uri: documents.ID().ApplyT(func(id string) (string, error) {
						return fmt.Sprintf("s3://%v/%v", aws_s3_bucket.Documents.Bucket, id), nil
					}).(pulumi.StringOutput),
				},
				EntityList: &comprehend.EntityRecognizerInputDataConfigEntityListArgs{
					S3Uri: entities.ID().ApplyT(func(id string) (string, error) {
						return fmt.Sprintf("s3://%v/%v", aws_s3_bucket.Entities.Bucket, id), nil
					}).(pulumi.StringOutput),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_iam_role_policy.Example,
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
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.comprehend.EntityRecognizer;
import com.pulumi.aws.comprehend.EntityRecognizerArgs;
import com.pulumi.aws.comprehend.inputs.EntityRecognizerInputDataConfigArgs;
import com.pulumi.aws.comprehend.inputs.EntityRecognizerInputDataConfigDocumentsArgs;
import com.pulumi.aws.comprehend.inputs.EntityRecognizerInputDataConfigEntityListArgs;
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
        var documents = new BucketObjectv2("documents");

        var entities = new BucketObjectv2("entities");

        var example = new EntityRecognizer("example", EntityRecognizerArgs.builder()        
            .dataAccessRoleArn(aws_iam_role.example().arn())
            .languageCode("en")
            .inputDataConfig(EntityRecognizerInputDataConfigArgs.builder()
                .entityTypes(                
                    EntityRecognizerInputDataConfigEntityTypeArgs.builder()
                        .type("ENTITY_1")
                        .build(),
                    EntityRecognizerInputDataConfigEntityTypeArgs.builder()
                        .type("ENTITY_2")
                        .build())
                .documents(EntityRecognizerInputDataConfigDocumentsArgs.builder()
                    .s3Uri(documents.id().applyValue(id -> String.format("s3://%s/%s", aws_s3_bucket.documents().bucket(),id)))
                    .build())
                .entityList(EntityRecognizerInputDataConfigEntityListArgs.builder()
                    .s3Uri(entities.id().applyValue(id -> String.format("s3://%s/%s", aws_s3_bucket.entities().bucket(),id)))
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_iam_role_policy.example())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:comprehend:EntityRecognizer
    properties:
      dataAccessRoleArn: ${aws_iam_role.example.arn}
      languageCode: en
      inputDataConfig:
        entityTypes:
          - type: ENTITY_1
          - type: ENTITY_2
        documents:
          s3Uri: s3://${aws_s3_bucket.documents.bucket}/${documents.id}
        entityList:
          s3Uri: s3://${aws_s3_bucket.entities.bucket}/${entities.id}
    options:
      dependson:
        - ${aws_iam_role_policy.example}
  documents:
    type: aws:s3:BucketObjectv2
  entities:
    type: aws:s3:BucketObjectv2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Comprehend Entity Recognizer using the ARN. For example:

```sh
 $ pulumi import aws:comprehend/entityRecognizer:EntityRecognizer example arn:aws:comprehend:us-west-2:123456789012:entity-recognizer/example
```
 