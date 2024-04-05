Resource for managing an AWS Kendra Data Source.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    description: "example",
    languageCode: "en",
    type: "CUSTOM",
    tags: {
        hello: "world",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    description="example",
    language_code="en",
    type="CUSTOM",
    tags={
        "hello": "world",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Description = "example",
        LanguageCode = "en",
        Type = "CUSTOM",
        Tags = 
        {
            { "hello", "world" },
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId:      pulumi.Any(aws_kendra_index.Example.Id),
			Description:  pulumi.String("example"),
			LanguageCode: pulumi.String("en"),
			Type:         pulumi.String("CUSTOM"),
			Tags: pulumi.StringMap{
				"hello": pulumi.String("world"),
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .description("example")
            .languageCode("en")
            .type("CUSTOM")
            .tags(Map.of("hello", "world"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      description: example
      languageCode: en
      type: CUSTOM
      tags:
        hello: world
```
{{% /example %}}
### S3 Connector
{{% example %}}
### With Schedule

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "S3",
    roleArn: aws_iam_role.example.arn,
    schedule: "cron(9 10 1 * ? *)",
    configuration: {
        s3Configuration: {
            bucketName: aws_s3_bucket.example.id,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="S3",
    role_arn=aws_iam_role["example"]["arn"],
    schedule="cron(9 10 1 * ? *)",
    configuration=aws.kendra.DataSourceConfigurationArgs(
        s3_configuration=aws.kendra.DataSourceConfigurationS3ConfigurationArgs(
            bucket_name=aws_s3_bucket["example"]["id"],
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
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "S3",
        RoleArn = aws_iam_role.Example.Arn,
        Schedule = "cron(9 10 1 * ? *)",
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            S3Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationS3ConfigurationArgs
            {
                BucketName = aws_s3_bucket.Example.Id,
            },
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId:  pulumi.Any(aws_kendra_index.Example.Id),
			Type:     pulumi.String("S3"),
			RoleArn:  pulumi.Any(aws_iam_role.Example.Arn),
			Schedule: pulumi.String("cron(9 10 1 * ? *)"),
			Configuration: &kendra.DataSourceConfigurationArgs{
				S3Configuration: &kendra.DataSourceConfigurationS3ConfigurationArgs{
					BucketName: pulumi.Any(aws_s3_bucket.Example.Id),
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationS3ConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("S3")
            .roleArn(aws_iam_role.example().arn())
            .schedule("cron(9 10 1 * ? *)")
            .configuration(DataSourceConfigurationArgs.builder()
                .s3Configuration(DataSourceConfigurationS3ConfigurationArgs.builder()
                    .bucketName(aws_s3_bucket.example().id())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: S3
      roleArn: ${aws_iam_role.example.arn}
      schedule: cron(9 10 1 * ? *)
      configuration:
        s3Configuration:
          bucketName: ${aws_s3_bucket.example.id}
```
{{% /example %}}
{{% example %}}
### With Access Control List

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "S3",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        s3Configuration: {
            bucketName: aws_s3_bucket.example.id,
            accessControlListConfiguration: {
                keyPath: `s3://${aws_s3_bucket.example.id}/path-1`,
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="S3",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        s3_configuration=aws.kendra.DataSourceConfigurationS3ConfigurationArgs(
            bucket_name=aws_s3_bucket["example"]["id"],
            access_control_list_configuration=aws.kendra.DataSourceConfigurationS3ConfigurationAccessControlListConfigurationArgs(
                key_path=f"s3://{aws_s3_bucket['example']['id']}/path-1",
            ),
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
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "S3",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            S3Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationS3ConfigurationArgs
            {
                BucketName = aws_s3_bucket.Example.Id,
                AccessControlListConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationS3ConfigurationAccessControlListConfigurationArgs
                {
                    KeyPath = $"s3://{aws_s3_bucket.Example.Id}/path-1",
                },
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("S3"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				S3Configuration: &kendra.DataSourceConfigurationS3ConfigurationArgs{
					BucketName: pulumi.Any(aws_s3_bucket.Example.Id),
					AccessControlListConfiguration: &kendra.DataSourceConfigurationS3ConfigurationAccessControlListConfigurationArgs{
						KeyPath: pulumi.String(fmt.Sprintf("s3://%v/path-1", aws_s3_bucket.Example.Id)),
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationS3ConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationS3ConfigurationAccessControlListConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("S3")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .s3Configuration(DataSourceConfigurationS3ConfigurationArgs.builder()
                    .bucketName(aws_s3_bucket.example().id())
                    .accessControlListConfiguration(DataSourceConfigurationS3ConfigurationAccessControlListConfigurationArgs.builder()
                        .keyPath(String.format("s3://%s/path-1", aws_s3_bucket.example().id()))
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: S3
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        s3Configuration:
          bucketName: ${aws_s3_bucket.example.id}
          accessControlListConfiguration:
            keyPath: s3://${aws_s3_bucket.example.id}/path-1
```
{{% /example %}}
{{% example %}}
### With Documents Metadata Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "S3",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        s3Configuration: {
            bucketName: aws_s3_bucket.example.id,
            exclusionPatterns: ["example"],
            inclusionPatterns: ["hello"],
            inclusionPrefixes: ["world"],
            documentsMetadataConfiguration: {
                s3Prefix: "example",
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="S3",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        s3_configuration=aws.kendra.DataSourceConfigurationS3ConfigurationArgs(
            bucket_name=aws_s3_bucket["example"]["id"],
            exclusion_patterns=["example"],
            inclusion_patterns=["hello"],
            inclusion_prefixes=["world"],
            documents_metadata_configuration=aws.kendra.DataSourceConfigurationS3ConfigurationDocumentsMetadataConfigurationArgs(
                s3_prefix="example",
            ),
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
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "S3",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            S3Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationS3ConfigurationArgs
            {
                BucketName = aws_s3_bucket.Example.Id,
                ExclusionPatterns = new[]
                {
                    "example",
                },
                InclusionPatterns = new[]
                {
                    "hello",
                },
                InclusionPrefixes = new[]
                {
                    "world",
                },
                DocumentsMetadataConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationS3ConfigurationDocumentsMetadataConfigurationArgs
                {
                    S3Prefix = "example",
                },
            },
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("S3"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				S3Configuration: &kendra.DataSourceConfigurationS3ConfigurationArgs{
					BucketName: pulumi.Any(aws_s3_bucket.Example.Id),
					ExclusionPatterns: pulumi.StringArray{
						pulumi.String("example"),
					},
					InclusionPatterns: pulumi.StringArray{
						pulumi.String("hello"),
					},
					InclusionPrefixes: pulumi.StringArray{
						pulumi.String("world"),
					},
					DocumentsMetadataConfiguration: &kendra.DataSourceConfigurationS3ConfigurationDocumentsMetadataConfigurationArgs{
						S3Prefix: pulumi.String("example"),
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationS3ConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationS3ConfigurationDocumentsMetadataConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("S3")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .s3Configuration(DataSourceConfigurationS3ConfigurationArgs.builder()
                    .bucketName(aws_s3_bucket.example().id())
                    .exclusionPatterns("example")
                    .inclusionPatterns("hello")
                    .inclusionPrefixes("world")
                    .documentsMetadataConfiguration(DataSourceConfigurationS3ConfigurationDocumentsMetadataConfigurationArgs.builder()
                        .s3Prefix("example")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: S3
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        s3Configuration:
          bucketName: ${aws_s3_bucket.example.id}
          exclusionPatterns:
            - example
          inclusionPatterns:
            - hello
          inclusionPrefixes:
            - world
          documentsMetadataConfiguration:
            s3Prefix: example
```
{{% /example %}}
### Web Crawler Connector
{{% example %}}
### With Seed URLs

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "WEBCRAWLER",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        webCrawlerConfiguration: {
            urls: {
                seedUrlConfiguration: {
                    seedUrls: ["REPLACE_WITH_YOUR_URL"],
                },
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="WEBCRAWLER",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        web_crawler_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationArgs(
            urls=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs(
                seed_url_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs(
                    seed_urls=["REPLACE_WITH_YOUR_URL"],
                ),
            ),
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
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "WEBCRAWLER",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            WebCrawlerConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationArgs
            {
                Urls = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs
                {
                    SeedUrlConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs
                    {
                        SeedUrls = new[]
                        {
                            "REPLACE_WITH_YOUR_URL",
                        },
                    },
                },
            },
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("WEBCRAWLER"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				WebCrawlerConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationArgs{
					Urls: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs{
						SeedUrlConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs{
							SeedUrls: pulumi.StringArray{
								pulumi.String("REPLACE_WITH_YOUR_URL"),
							},
						},
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("WEBCRAWLER")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .webCrawlerConfiguration(DataSourceConfigurationWebCrawlerConfigurationArgs.builder()
                    .urls(DataSourceConfigurationWebCrawlerConfigurationUrlsArgs.builder()
                        .seedUrlConfiguration(DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs.builder()
                            .seedUrls("REPLACE_WITH_YOUR_URL")
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: WEBCRAWLER
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        webCrawlerConfiguration:
          urls:
            seedUrlConfiguration:
              seedUrls:
                - REPLACE_WITH_YOUR_URL
```
{{% /example %}}
{{% example %}}
### With Site Maps

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "WEBCRAWLER",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        webCrawlerConfiguration: {
            urls: {
                siteMapsConfiguration: {
                    siteMaps: ["REPLACE_WITH_YOUR_URL"],
                },
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="WEBCRAWLER",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        web_crawler_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationArgs(
            urls=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs(
                site_maps_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfigurationArgs(
                    site_maps=["REPLACE_WITH_YOUR_URL"],
                ),
            ),
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
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "WEBCRAWLER",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            WebCrawlerConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationArgs
            {
                Urls = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs
                {
                    SiteMapsConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfigurationArgs
                    {
                        SiteMaps = new[]
                        {
                            "REPLACE_WITH_YOUR_URL",
                        },
                    },
                },
            },
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("WEBCRAWLER"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				WebCrawlerConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationArgs{
					Urls: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs{
						SiteMapsConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfigurationArgs{
							SiteMaps: pulumi.StringArray{
								pulumi.String("REPLACE_WITH_YOUR_URL"),
							},
						},
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("WEBCRAWLER")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .webCrawlerConfiguration(DataSourceConfigurationWebCrawlerConfigurationArgs.builder()
                    .urls(DataSourceConfigurationWebCrawlerConfigurationUrlsArgs.builder()
                        .siteMapsConfiguration(DataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfigurationArgs.builder()
                            .siteMaps("REPLACE_WITH_YOUR_URL")
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: WEBCRAWLER
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        webCrawlerConfiguration:
          urls:
            siteMapsConfiguration:
              siteMaps:
                - REPLACE_WITH_YOUR_URL
```
{{% /example %}}
{{% example %}}
### With Web Crawler Mode

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "WEBCRAWLER",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        webCrawlerConfiguration: {
            urls: {
                seedUrlConfiguration: {
                    webCrawlerMode: "SUBDOMAINS",
                    seedUrls: ["REPLACE_WITH_YOUR_URL"],
                },
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="WEBCRAWLER",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        web_crawler_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationArgs(
            urls=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs(
                seed_url_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs(
                    web_crawler_mode="SUBDOMAINS",
                    seed_urls=["REPLACE_WITH_YOUR_URL"],
                ),
            ),
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
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "WEBCRAWLER",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            WebCrawlerConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationArgs
            {
                Urls = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs
                {
                    SeedUrlConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs
                    {
                        WebCrawlerMode = "SUBDOMAINS",
                        SeedUrls = new[]
                        {
                            "REPLACE_WITH_YOUR_URL",
                        },
                    },
                },
            },
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("WEBCRAWLER"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				WebCrawlerConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationArgs{
					Urls: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs{
						SeedUrlConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs{
							WebCrawlerMode: pulumi.String("SUBDOMAINS"),
							SeedUrls: pulumi.StringArray{
								pulumi.String("REPLACE_WITH_YOUR_URL"),
							},
						},
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("WEBCRAWLER")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .webCrawlerConfiguration(DataSourceConfigurationWebCrawlerConfigurationArgs.builder()
                    .urls(DataSourceConfigurationWebCrawlerConfigurationUrlsArgs.builder()
                        .seedUrlConfiguration(DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs.builder()
                            .webCrawlerMode("SUBDOMAINS")
                            .seedUrls("REPLACE_WITH_YOUR_URL")
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: WEBCRAWLER
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        webCrawlerConfiguration:
          urls:
            seedUrlConfiguration:
              webCrawlerMode: SUBDOMAINS
              seedUrls:
                - REPLACE_WITH_YOUR_URL
```
{{% /example %}}
{{% example %}}
### With Authentication Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "WEBCRAWLER",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        webCrawlerConfiguration: {
            authenticationConfiguration: {
                basicAuthentications: [{
                    credentials: aws_secretsmanager_secret.example.arn,
                    host: "a.example.com",
                    port: 443,
                }],
            },
            urls: {
                seedUrlConfiguration: {
                    seedUrls: ["REPLACE_WITH_YOUR_URL"],
                },
            },
        },
    },
}, {
    dependsOn: [aws_secretsmanager_secret_version.example],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="WEBCRAWLER",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        web_crawler_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationArgs(
            authentication_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationArgs(
                basic_authentications=[aws.kendra.DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationBasicAuthenticationArgs(
                    credentials=aws_secretsmanager_secret["example"]["arn"],
                    host="a.example.com",
                    port=443,
                )],
            ),
            urls=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs(
                seed_url_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs(
                    seed_urls=["REPLACE_WITH_YOUR_URL"],
                ),
            ),
        ),
    ),
    opts=pulumi.ResourceOptions(depends_on=[aws_secretsmanager_secret_version["example"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "WEBCRAWLER",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            WebCrawlerConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationArgs
            {
                AuthenticationConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationArgs
                {
                    BasicAuthentications = new[]
                    {
                        new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationBasicAuthenticationArgs
                        {
                            Credentials = aws_secretsmanager_secret.Example.Arn,
                            Host = "a.example.com",
                            Port = 443,
                        },
                    },
                },
                Urls = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs
                {
                    SeedUrlConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs
                    {
                        SeedUrls = new[]
                        {
                            "REPLACE_WITH_YOUR_URL",
                        },
                    },
                },
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_secretsmanager_secret_version.Example,
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("WEBCRAWLER"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				WebCrawlerConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationArgs{
					AuthenticationConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationArgs{
						BasicAuthentications: kendra.DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationBasicAuthenticationArray{
							&kendra.DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationBasicAuthenticationArgs{
								Credentials: pulumi.Any(aws_secretsmanager_secret.Example.Arn),
								Host:        pulumi.String("a.example.com"),
								Port:        pulumi.Int(443),
							},
						},
					},
					Urls: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs{
						SeedUrlConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs{
							SeedUrls: pulumi.StringArray{
								pulumi.String("REPLACE_WITH_YOUR_URL"),
							},
						},
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_secretsmanager_secret_version.Example,
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("WEBCRAWLER")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .webCrawlerConfiguration(DataSourceConfigurationWebCrawlerConfigurationArgs.builder()
                    .authenticationConfiguration(DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationArgs.builder()
                        .basicAuthentications(DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationBasicAuthenticationArgs.builder()
                            .credentials(aws_secretsmanager_secret.example().arn())
                            .host("a.example.com")
                            .port("443")
                            .build())
                        .build())
                    .urls(DataSourceConfigurationWebCrawlerConfigurationUrlsArgs.builder()
                        .seedUrlConfiguration(DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs.builder()
                            .seedUrls("REPLACE_WITH_YOUR_URL")
                            .build())
                        .build())
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_secretsmanager_secret_version.example())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: WEBCRAWLER
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        webCrawlerConfiguration:
          authenticationConfiguration:
            basicAuthentications:
              - credentials: ${aws_secretsmanager_secret.example.arn}
                host: a.example.com
                port: '443'
          urls:
            seedUrlConfiguration:
              seedUrls:
                - REPLACE_WITH_YOUR_URL
    options:
      dependson:
        - ${aws_secretsmanager_secret_version.example}
```
{{% /example %}}
{{% example %}}
### With Crawl Depth

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "WEBCRAWLER",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        webCrawlerConfiguration: {
            crawlDepth: 3,
            urls: {
                seedUrlConfiguration: {
                    seedUrls: ["REPLACE_WITH_YOUR_URL"],
                },
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="WEBCRAWLER",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        web_crawler_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationArgs(
            crawl_depth=3,
            urls=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs(
                seed_url_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs(
                    seed_urls=["REPLACE_WITH_YOUR_URL"],
                ),
            ),
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
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "WEBCRAWLER",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            WebCrawlerConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationArgs
            {
                CrawlDepth = 3,
                Urls = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs
                {
                    SeedUrlConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs
                    {
                        SeedUrls = new[]
                        {
                            "REPLACE_WITH_YOUR_URL",
                        },
                    },
                },
            },
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("WEBCRAWLER"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				WebCrawlerConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationArgs{
					CrawlDepth: pulumi.Int(3),
					Urls: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs{
						SeedUrlConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs{
							SeedUrls: pulumi.StringArray{
								pulumi.String("REPLACE_WITH_YOUR_URL"),
							},
						},
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("WEBCRAWLER")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .webCrawlerConfiguration(DataSourceConfigurationWebCrawlerConfigurationArgs.builder()
                    .crawlDepth(3)
                    .urls(DataSourceConfigurationWebCrawlerConfigurationUrlsArgs.builder()
                        .seedUrlConfiguration(DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs.builder()
                            .seedUrls("REPLACE_WITH_YOUR_URL")
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: WEBCRAWLER
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        webCrawlerConfiguration:
          crawlDepth: 3
          urls:
            seedUrlConfiguration:
              seedUrls:
                - REPLACE_WITH_YOUR_URL
```
{{% /example %}}
{{% example %}}
### With Max Links Per Page

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "WEBCRAWLER",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        webCrawlerConfiguration: {
            maxLinksPerPage: 100,
            urls: {
                seedUrlConfiguration: {
                    seedUrls: ["REPLACE_WITH_YOUR_URL"],
                },
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="WEBCRAWLER",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        web_crawler_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationArgs(
            max_links_per_page=100,
            urls=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs(
                seed_url_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs(
                    seed_urls=["REPLACE_WITH_YOUR_URL"],
                ),
            ),
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
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "WEBCRAWLER",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            WebCrawlerConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationArgs
            {
                MaxLinksPerPage = 100,
                Urls = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs
                {
                    SeedUrlConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs
                    {
                        SeedUrls = new[]
                        {
                            "REPLACE_WITH_YOUR_URL",
                        },
                    },
                },
            },
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("WEBCRAWLER"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				WebCrawlerConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationArgs{
					MaxLinksPerPage: pulumi.Int(100),
					Urls: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs{
						SeedUrlConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs{
							SeedUrls: pulumi.StringArray{
								pulumi.String("REPLACE_WITH_YOUR_URL"),
							},
						},
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("WEBCRAWLER")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .webCrawlerConfiguration(DataSourceConfigurationWebCrawlerConfigurationArgs.builder()
                    .maxLinksPerPage(100)
                    .urls(DataSourceConfigurationWebCrawlerConfigurationUrlsArgs.builder()
                        .seedUrlConfiguration(DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs.builder()
                            .seedUrls("REPLACE_WITH_YOUR_URL")
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: WEBCRAWLER
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        webCrawlerConfiguration:
          maxLinksPerPage: 100
          urls:
            seedUrlConfiguration:
              seedUrls:
                - REPLACE_WITH_YOUR_URL
```
{{% /example %}}
{{% example %}}
### With Max Urls Per Minute Crawl Rate

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "WEBCRAWLER",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        webCrawlerConfiguration: {
            maxUrlsPerMinuteCrawlRate: 300,
            urls: {
                seedUrlConfiguration: {
                    seedUrls: ["REPLACE_WITH_YOUR_URL"],
                },
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="WEBCRAWLER",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        web_crawler_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationArgs(
            max_urls_per_minute_crawl_rate=300,
            urls=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs(
                seed_url_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs(
                    seed_urls=["REPLACE_WITH_YOUR_URL"],
                ),
            ),
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
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "WEBCRAWLER",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            WebCrawlerConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationArgs
            {
                MaxUrlsPerMinuteCrawlRate = 300,
                Urls = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs
                {
                    SeedUrlConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs
                    {
                        SeedUrls = new[]
                        {
                            "REPLACE_WITH_YOUR_URL",
                        },
                    },
                },
            },
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("WEBCRAWLER"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				WebCrawlerConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationArgs{
					MaxUrlsPerMinuteCrawlRate: pulumi.Int(300),
					Urls: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs{
						SeedUrlConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs{
							SeedUrls: pulumi.StringArray{
								pulumi.String("REPLACE_WITH_YOUR_URL"),
							},
						},
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("WEBCRAWLER")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .webCrawlerConfiguration(DataSourceConfigurationWebCrawlerConfigurationArgs.builder()
                    .maxUrlsPerMinuteCrawlRate(300)
                    .urls(DataSourceConfigurationWebCrawlerConfigurationUrlsArgs.builder()
                        .seedUrlConfiguration(DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs.builder()
                            .seedUrls("REPLACE_WITH_YOUR_URL")
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: WEBCRAWLER
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        webCrawlerConfiguration:
          maxUrlsPerMinuteCrawlRate: 300
          urls:
            seedUrlConfiguration:
              seedUrls:
                - REPLACE_WITH_YOUR_URL
```
{{% /example %}}
{{% example %}}
### With Proxy Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "WEBCRAWLER",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        webCrawlerConfiguration: {
            proxyConfiguration: {
                credentials: aws_secretsmanager_secret.example.arn,
                host: "a.example.com",
                port: 443,
            },
            urls: {
                seedUrlConfiguration: {
                    seedUrls: ["REPLACE_WITH_YOUR_URL"],
                },
            },
        },
    },
}, {
    dependsOn: [aws_secretsmanager_secret_version.example],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="WEBCRAWLER",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        web_crawler_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationArgs(
            proxy_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationProxyConfigurationArgs(
                credentials=aws_secretsmanager_secret["example"]["arn"],
                host="a.example.com",
                port=443,
            ),
            urls=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs(
                seed_url_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs(
                    seed_urls=["REPLACE_WITH_YOUR_URL"],
                ),
            ),
        ),
    ),
    opts=pulumi.ResourceOptions(depends_on=[aws_secretsmanager_secret_version["example"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "WEBCRAWLER",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            WebCrawlerConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationArgs
            {
                ProxyConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationProxyConfigurationArgs
                {
                    Credentials = aws_secretsmanager_secret.Example.Arn,
                    Host = "a.example.com",
                    Port = 443,
                },
                Urls = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs
                {
                    SeedUrlConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs
                    {
                        SeedUrls = new[]
                        {
                            "REPLACE_WITH_YOUR_URL",
                        },
                    },
                },
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_secretsmanager_secret_version.Example,
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("WEBCRAWLER"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				WebCrawlerConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationArgs{
					ProxyConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationProxyConfigurationArgs{
						Credentials: pulumi.Any(aws_secretsmanager_secret.Example.Arn),
						Host:        pulumi.String("a.example.com"),
						Port:        pulumi.Int(443),
					},
					Urls: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs{
						SeedUrlConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs{
							SeedUrls: pulumi.StringArray{
								pulumi.String("REPLACE_WITH_YOUR_URL"),
							},
						},
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_secretsmanager_secret_version.Example,
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationProxyConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("WEBCRAWLER")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .webCrawlerConfiguration(DataSourceConfigurationWebCrawlerConfigurationArgs.builder()
                    .proxyConfiguration(DataSourceConfigurationWebCrawlerConfigurationProxyConfigurationArgs.builder()
                        .credentials(aws_secretsmanager_secret.example().arn())
                        .host("a.example.com")
                        .port("443")
                        .build())
                    .urls(DataSourceConfigurationWebCrawlerConfigurationUrlsArgs.builder()
                        .seedUrlConfiguration(DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs.builder()
                            .seedUrls("REPLACE_WITH_YOUR_URL")
                            .build())
                        .build())
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_secretsmanager_secret_version.example())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: WEBCRAWLER
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        webCrawlerConfiguration:
          proxyConfiguration:
            credentials: ${aws_secretsmanager_secret.example.arn}
            host: a.example.com
            port: '443'
          urls:
            seedUrlConfiguration:
              seedUrls:
                - REPLACE_WITH_YOUR_URL
    options:
      dependson:
        - ${aws_secretsmanager_secret_version.example}
```
{{% /example %}}
{{% example %}}
### With URL Exclusion and Inclusion Patterns

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.DataSource("example", {
    indexId: aws_kendra_index.example.id,
    type: "WEBCRAWLER",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        webCrawlerConfiguration: {
            urlExclusionPatterns: ["example"],
            urlInclusionPatterns: ["hello"],
            urls: {
                seedUrlConfiguration: {
                    seedUrls: ["REPLACE_WITH_YOUR_URL"],
                },
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.DataSource("example",
    index_id=aws_kendra_index["example"]["id"],
    type="WEBCRAWLER",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.DataSourceConfigurationArgs(
        web_crawler_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationArgs(
            url_exclusion_patterns=["example"],
            url_inclusion_patterns=["hello"],
            urls=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs(
                seed_url_configuration=aws.kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs(
                    seed_urls=["REPLACE_WITH_YOUR_URL"],
                ),
            ),
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
    var example = new Aws.Kendra.DataSource("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Type = "WEBCRAWLER",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.DataSourceConfigurationArgs
        {
            WebCrawlerConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationArgs
            {
                UrlExclusionPatterns = new[]
                {
                    "example",
                },
                UrlInclusionPatterns = new[]
                {
                    "hello",
                },
                Urls = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs
                {
                    SeedUrlConfiguration = new Aws.Kendra.Inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs
                    {
                        SeedUrls = new[]
                        {
                            "REPLACE_WITH_YOUR_URL",
                        },
                    },
                },
            },
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
		_, err := kendra.NewDataSource(ctx, "example", &kendra.DataSourceArgs{
			IndexId: pulumi.Any(aws_kendra_index.Example.Id),
			Type:    pulumi.String("WEBCRAWLER"),
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.DataSourceConfigurationArgs{
				WebCrawlerConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationArgs{
					UrlExclusionPatterns: pulumi.StringArray{
						pulumi.String("example"),
					},
					UrlInclusionPatterns: pulumi.StringArray{
						pulumi.String("hello"),
					},
					Urls: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs{
						SeedUrlConfiguration: &kendra.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs{
							SeedUrls: pulumi.StringArray{
								pulumi.String("REPLACE_WITH_YOUR_URL"),
							},
						},
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
import com.pulumi.aws.kendra.DataSource;
import com.pulumi.aws.kendra.DataSourceArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsArgs;
import com.pulumi.aws.kendra.inputs.DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs;
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
        var example = new DataSource("example", DataSourceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .type("WEBCRAWLER")
            .roleArn(aws_iam_role.example().arn())
            .configuration(DataSourceConfigurationArgs.builder()
                .webCrawlerConfiguration(DataSourceConfigurationWebCrawlerConfigurationArgs.builder()
                    .urlExclusionPatterns("example")
                    .urlInclusionPatterns("hello")
                    .urls(DataSourceConfigurationWebCrawlerConfigurationUrlsArgs.builder()
                        .seedUrlConfiguration(DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfigurationArgs.builder()
                            .seedUrls("REPLACE_WITH_YOUR_URL")
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:DataSource
    properties:
      indexId: ${aws_kendra_index.example.id}
      type: WEBCRAWLER
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        webCrawlerConfiguration:
          urlExclusionPatterns:
            - example
          urlInclusionPatterns:
            - hello
          urls:
            seedUrlConfiguration:
              seedUrls:
                - REPLACE_WITH_YOUR_URL
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Kendra Data Source using the unique identifiers of the data_source and index separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:kendra/dataSource:DataSource example 1045d08d-66ef-4882-b3ed-dfb7df183e90/b34dfdf7-1f2b-4704-9581-79e00296845f
```
 