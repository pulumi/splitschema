Use the `aws_kendra_index_block_list` resource to manage an AWS Kendra block list used for query suggestions for an index.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.QuerySuggestionsBlockList("example", {
    indexId: aws_kendra_index.example.id,
    roleArn: aws_iam_role.example.arn,
    sourceS3Path: {
        bucket: aws_s3_bucket.example.id,
        key: "example/suggestions.txt",
    },
    tags: {
        Name: "Example Kendra Index",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.QuerySuggestionsBlockList("example",
    index_id=aws_kendra_index["example"]["id"],
    role_arn=aws_iam_role["example"]["arn"],
    source_s3_path=aws.kendra.QuerySuggestionsBlockListSourceS3PathArgs(
        bucket=aws_s3_bucket["example"]["id"],
        key="example/suggestions.txt",
    ),
    tags={
        "Name": "Example Kendra Index",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.QuerySuggestionsBlockList("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        RoleArn = aws_iam_role.Example.Arn,
        SourceS3Path = new Aws.Kendra.Inputs.QuerySuggestionsBlockListSourceS3PathArgs
        {
            Bucket = aws_s3_bucket.Example.Id,
            Key = "example/suggestions.txt",
        },
        Tags = 
        {
            { "Name", "Example Kendra Index" },
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
		_, err := kendra.NewQuerySuggestionsBlockList(ctx, "example", &kendra.QuerySuggestionsBlockListArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			SourceS3Path: &kendra.QuerySuggestionsBlockListSourceS3PathArgs{
				Bucket: pulumi.Any(aws_s3_bucket.Example.Id),
				Key:    pulumi.String("example/suggestions.txt"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Kendra Index"),
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
import com.pulumi.aws.kendra.QuerySuggestionsBlockList;
import com.pulumi.aws.kendra.QuerySuggestionsBlockListArgs;
import com.pulumi.aws.kendra.inputs.QuerySuggestionsBlockListSourceS3PathArgs;
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
        var example = new QuerySuggestionsBlockList("example", QuerySuggestionsBlockListArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .roleArn(aws_iam_role.example().arn())
            .sourceS3Path(QuerySuggestionsBlockListSourceS3PathArgs.builder()
                .bucket(aws_s3_bucket.example().id())
                .key("example/suggestions.txt")
                .build())
            .tags(Map.of("Name", "Example Kendra Index"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:QuerySuggestionsBlockList
    properties:
      indexId: ${aws_kendra_index.example.id}
      roleArn: ${aws_iam_role.example.arn}
      sourceS3Path:
        bucket: ${aws_s3_bucket.example.id}
        key: example/suggestions.txt
      tags:
        Name: Example Kendra Index
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the `aws_kendra_query_suggestions_block_list` resource using the unique identifiers of the block list and index separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:kendra/querySuggestionsBlockList:QuerySuggestionsBlockList example blocklist-123456780/idx-8012925589
```
 