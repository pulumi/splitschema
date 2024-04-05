Resource for managing an AWS Comprehend Document Classifier.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const documents = new aws.s3.BucketObjectv2("documents", {});
// ...
const example = new aws.comprehend.DocumentClassifier("example", {
    dataAccessRoleArn: aws_iam_role.example.arn,
    languageCode: "en",
    inputDataConfig: {
        s3Uri: pulumi.interpolate`s3://${aws_s3_bucket.test.bucket}/${documents.id}`,
    },
}, {
    dependsOn: [aws_iam_role_policy.example],
});
const entities = new aws.s3.BucketObjectv2("entities", {});
// ...
```
```python
import pulumi
import pulumi_aws as aws

documents = aws.s3.BucketObjectv2("documents")
# ...
example = aws.comprehend.DocumentClassifier("example",
    data_access_role_arn=aws_iam_role["example"]["arn"],
    language_code="en",
    input_data_config=aws.comprehend.DocumentClassifierInputDataConfigArgs(
        s3_uri=documents.id.apply(lambda id: f"s3://{aws_s3_bucket['test']['bucket']}/{id}"),
    ),
    opts=pulumi.ResourceOptions(depends_on=[aws_iam_role_policy["example"]]))
entities = aws.s3.BucketObjectv2("entities")
# ...
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
    var example = new Aws.Comprehend.DocumentClassifier("example", new()
    {
        DataAccessRoleArn = aws_iam_role.Example.Arn,
        LanguageCode = "en",
        InputDataConfig = new Aws.Comprehend.Inputs.DocumentClassifierInputDataConfigArgs
        {
            S3Uri = documents.Id.Apply(id => $"s3://{aws_s3_bucket.Test.Bucket}/{id}"),
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_role_policy.Example,
        },
    });

    var entities = new Aws.S3.BucketObjectv2("entities");

    // ...
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
		_, err = comprehend.NewDocumentClassifier(ctx, "example", &comprehend.DocumentClassifierArgs{
			DataAccessRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			LanguageCode:      pulumi.String("en"),
			InputDataConfig: &comprehend.DocumentClassifierInputDataConfigArgs{
				S3Uri: documents.ID().ApplyT(func(id string) (string, error) {
					return fmt.Sprintf("s3://%v/%v", aws_s3_bucket.Test.Bucket, id), nil
				}).(pulumi.StringOutput),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_iam_role_policy.Example,
		}))
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObjectv2(ctx, "entities", nil)
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
import com.pulumi.aws.comprehend.DocumentClassifier;
import com.pulumi.aws.comprehend.DocumentClassifierArgs;
import com.pulumi.aws.comprehend.inputs.DocumentClassifierInputDataConfigArgs;
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

        var example = new DocumentClassifier("example", DocumentClassifierArgs.builder()        
            .dataAccessRoleArn(aws_iam_role.example().arn())
            .languageCode("en")
            .inputDataConfig(DocumentClassifierInputDataConfigArgs.builder()
                .s3Uri(documents.id().applyValue(id -> String.format("s3://%s/%s", aws_s3_bucket.test().bucket(),id)))
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_iam_role_policy.example())
                .build());

        var entities = new BucketObjectv2("entities");

    }
}
```
```yaml
resources:
  example:
    type: aws:comprehend:DocumentClassifier
    properties:
      dataAccessRoleArn: ${aws_iam_role.example.arn}
      languageCode: en
      inputDataConfig:
        s3Uri: s3://${aws_s3_bucket.test.bucket}/${documents.id}
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

Using `pulumi import`, import Comprehend Document Classifier using the ARN. For example:

```sh
 $ pulumi import aws:comprehend/documentClassifier:DocumentClassifier example arn:aws:comprehend:us-west-2:123456789012:document_classifier/example
```
 