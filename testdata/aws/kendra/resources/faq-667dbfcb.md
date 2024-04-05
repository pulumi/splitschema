Resource for managing an AWS Kendra FAQ.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Faq("example", {
    indexId: aws_kendra_index.example.id,
    roleArn: aws_iam_role.example.arn,
    s3Path: {
        bucket: aws_s3_bucket.example.id,
        key: aws_s3_object.example.key,
    },
    tags: {
        Name: "Example Kendra Faq",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Faq("example",
    index_id=aws_kendra_index["example"]["id"],
    role_arn=aws_iam_role["example"]["arn"],
    s3_path=aws.kendra.FaqS3PathArgs(
        bucket=aws_s3_bucket["example"]["id"],
        key=aws_s3_object["example"]["key"],
    ),
    tags={
        "Name": "Example Kendra Faq",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Faq("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        RoleArn = aws_iam_role.Example.Arn,
        S3Path = new Aws.Kendra.Inputs.FaqS3PathArgs
        {
            Bucket = aws_s3_bucket.Example.Id,
            Key = aws_s3_object.Example.Key,
        },
        Tags = 
        {
            { "Name", "Example Kendra Faq" },
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
		_, err := kendra.NewFaq(ctx, "example", &kendra.FaqArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			S3Path: &kendra.FaqS3PathArgs{
				Bucket: pulumi.Any(aws_s3_bucket.Example.Id),
				Key:    pulumi.Any(aws_s3_object.Example.Key),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Kendra Faq"),
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
import com.pulumi.aws.kendra.Faq;
import com.pulumi.aws.kendra.FaqArgs;
import com.pulumi.aws.kendra.inputs.FaqS3PathArgs;
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
        var example = new Faq("example", FaqArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .roleArn(aws_iam_role.example().arn())
            .s3Path(FaqS3PathArgs.builder()
                .bucket(aws_s3_bucket.example().id())
                .key(aws_s3_object.example().key())
                .build())
            .tags(Map.of("Name", "Example Kendra Faq"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Faq
    properties:
      indexId: ${aws_kendra_index.example.id}
      roleArn: ${aws_iam_role.example.arn}
      s3Path:
        bucket: ${aws_s3_bucket.example.id}
        key: ${aws_s3_object.example.key}
      tags:
        Name: Example Kendra Faq
```
{{% /example %}}
{{% example %}}
### With File Format

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Faq("example", {
    indexId: aws_kendra_index.example.id,
    fileFormat: "CSV",
    roleArn: aws_iam_role.example.arn,
    s3Path: {
        bucket: aws_s3_bucket.example.id,
        key: aws_s3_object.example.key,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Faq("example",
    index_id=aws_kendra_index["example"]["id"],
    file_format="CSV",
    role_arn=aws_iam_role["example"]["arn"],
    s3_path=aws.kendra.FaqS3PathArgs(
        bucket=aws_s3_bucket["example"]["id"],
        key=aws_s3_object["example"]["key"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Faq("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        FileFormat = "CSV",
        RoleArn = aws_iam_role.Example.Arn,
        S3Path = new Aws.Kendra.Inputs.FaqS3PathArgs
        {
            Bucket = aws_s3_bucket.Example.Id,
            Key = aws_s3_object.Example.Key,
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
		_, err := kendra.NewFaq(ctx, "example", &kendra.FaqArgs{
			IndexId:    pulumi.Any(aws_kendra_index.Example.Id),
			FileFormat: pulumi.String("CSV"),
			RoleArn:    pulumi.Any(aws_iam_role.Example.Arn),
			S3Path: &kendra.FaqS3PathArgs{
				Bucket: pulumi.Any(aws_s3_bucket.Example.Id),
				Key:    pulumi.Any(aws_s3_object.Example.Key),
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
import com.pulumi.aws.kendra.Faq;
import com.pulumi.aws.kendra.FaqArgs;
import com.pulumi.aws.kendra.inputs.FaqS3PathArgs;
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
        var example = new Faq("example", FaqArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .fileFormat("CSV")
            .roleArn(aws_iam_role.example().arn())
            .s3Path(FaqS3PathArgs.builder()
                .bucket(aws_s3_bucket.example().id())
                .key(aws_s3_object.example().key())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Faq
    properties:
      indexId: ${aws_kendra_index.example.id}
      fileFormat: CSV
      roleArn: ${aws_iam_role.example.arn}
      s3Path:
        bucket: ${aws_s3_bucket.example.id}
        key: ${aws_s3_object.example.key}
```
{{% /example %}}
{{% example %}}
### With Language Code

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Faq("example", {
    indexId: aws_kendra_index.example.id,
    languageCode: "en",
    roleArn: aws_iam_role.example.arn,
    s3Path: {
        bucket: aws_s3_bucket.example.id,
        key: aws_s3_object.example.key,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Faq("example",
    index_id=aws_kendra_index["example"]["id"],
    language_code="en",
    role_arn=aws_iam_role["example"]["arn"],
    s3_path=aws.kendra.FaqS3PathArgs(
        bucket=aws_s3_bucket["example"]["id"],
        key=aws_s3_object["example"]["key"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Faq("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        LanguageCode = "en",
        RoleArn = aws_iam_role.Example.Arn,
        S3Path = new Aws.Kendra.Inputs.FaqS3PathArgs
        {
            Bucket = aws_s3_bucket.Example.Id,
            Key = aws_s3_object.Example.Key,
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
		_, err := kendra.NewFaq(ctx, "example", &kendra.FaqArgs{
			IndexId:      pulumi.Any(aws_kendra_index.Example.Id),
			LanguageCode: pulumi.String("en"),
			RoleArn:      pulumi.Any(aws_iam_role.Example.Arn),
			S3Path: &kendra.FaqS3PathArgs{
				Bucket: pulumi.Any(aws_s3_bucket.Example.Id),
				Key:    pulumi.Any(aws_s3_object.Example.Key),
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
import com.pulumi.aws.kendra.Faq;
import com.pulumi.aws.kendra.FaqArgs;
import com.pulumi.aws.kendra.inputs.FaqS3PathArgs;
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
        var example = new Faq("example", FaqArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .languageCode("en")
            .roleArn(aws_iam_role.example().arn())
            .s3Path(FaqS3PathArgs.builder()
                .bucket(aws_s3_bucket.example().id())
                .key(aws_s3_object.example().key())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Faq
    properties:
      indexId: ${aws_kendra_index.example.id}
      languageCode: en
      roleArn: ${aws_iam_role.example.arn}
      s3Path:
        bucket: ${aws_s3_bucket.example.id}
        key: ${aws_s3_object.example.key}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_kendra_faq` using the unique identifiers of the FAQ and index separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:kendra/faq:Faq example faq-123456780/idx-8012925589
```
 