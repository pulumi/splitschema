Resource for managing an AWS Kendra Thesaurus.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Thesaurus("example", {
    indexId: aws_kendra_index.example.id,
    roleArn: aws_iam_role.example.arn,
    sourceS3Path: {
        bucket: aws_s3_bucket.example.id,
        key: aws_s3_object.example.key,
    },
    tags: {
        Name: "Example Kendra Thesaurus",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Thesaurus("example",
    index_id=aws_kendra_index["example"]["id"],
    role_arn=aws_iam_role["example"]["arn"],
    source_s3_path=aws.kendra.ThesaurusSourceS3PathArgs(
        bucket=aws_s3_bucket["example"]["id"],
        key=aws_s3_object["example"]["key"],
    ),
    tags={
        "Name": "Example Kendra Thesaurus",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Thesaurus("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        RoleArn = aws_iam_role.Example.Arn,
        SourceS3Path = new Aws.Kendra.Inputs.ThesaurusSourceS3PathArgs
        {
            Bucket = aws_s3_bucket.Example.Id,
            Key = aws_s3_object.Example.Key,
        },
        Tags = 
        {
            { "Name", "Example Kendra Thesaurus" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.NewThesaurus(ctx, "example", &kendra.ThesaurusArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			SourceS3Path: &kendra.ThesaurusSourceS3PathArgs{
				Bucket: pulumi.Any(aws_s3_bucket.Example.Id),
				Key:    pulumi.Any(aws_s3_object.Example.Key),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Kendra Thesaurus"),
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
import com.pulumi.aws.kendra.Thesaurus;
import com.pulumi.aws.kendra.ThesaurusArgs;
import com.pulumi.aws.kendra.inputs.ThesaurusSourceS3PathArgs;
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
        var example = new Thesaurus("example", ThesaurusArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .roleArn(aws_iam_role.example().arn())
            .sourceS3Path(ThesaurusSourceS3PathArgs.builder()
                .bucket(aws_s3_bucket.example().id())
                .key(aws_s3_object.example().key())
                .build())
            .tags(Map.of("Name", "Example Kendra Thesaurus"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Thesaurus
    properties:
      indexId: ${aws_kendra_index.example.id}
      roleArn: ${aws_iam_role.example.arn}
      sourceS3Path:
        bucket: ${aws_s3_bucket.example.id}
        key: ${aws_s3_object.example.key}
      tags:
        Name: Example Kendra Thesaurus
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_kendra_thesaurus` using the unique identifiers of the thesaurus and index separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:kendra/thesaurus:Thesaurus example thesaurus-123456780/idx-8012925589
```
 