Manages a Glue Crawler. More information can be found in the [AWS Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html)

{{% examples %}}
## Example Usage
{{% example %}}
### DynamoDB Target Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Crawler("example", {
    databaseName: aws_glue_catalog_database.example.name,
    role: aws_iam_role.example.arn,
    dynamodbTargets: [{
        path: "table-name",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Crawler("example",
    database_name=aws_glue_catalog_database["example"]["name"],
    role=aws_iam_role["example"]["arn"],
    dynamodb_targets=[aws.glue.CrawlerDynamodbTargetArgs(
        path="table-name",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Crawler("example", new()
    {
        DatabaseName = aws_glue_catalog_database.Example.Name,
        Role = aws_iam_role.Example.Arn,
        DynamodbTargets = new[]
        {
            new Aws.Glue.Inputs.CrawlerDynamodbTargetArgs
            {
                Path = "table-name",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCrawler(ctx, "example", &glue.CrawlerArgs{
			DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
			Role:         pulumi.Any(aws_iam_role.Example.Arn),
			DynamodbTargets: glue.CrawlerDynamodbTargetArray{
				&glue.CrawlerDynamodbTargetArgs{
					Path: pulumi.String("table-name"),
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
import com.pulumi.aws.glue.Crawler;
import com.pulumi.aws.glue.CrawlerArgs;
import com.pulumi.aws.glue.inputs.CrawlerDynamodbTargetArgs;
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
        var example = new Crawler("example", CrawlerArgs.builder()        
            .databaseName(aws_glue_catalog_database.example().name())
            .role(aws_iam_role.example().arn())
            .dynamodbTargets(CrawlerDynamodbTargetArgs.builder()
                .path("table-name")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Crawler
    properties:
      databaseName: ${aws_glue_catalog_database.example.name}
      role: ${aws_iam_role.example.arn}
      dynamodbTargets:
        - path: table-name
```
{{% /example %}}
{{% example %}}
### JDBC Target Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Crawler("example", {
    databaseName: aws_glue_catalog_database.example.name,
    role: aws_iam_role.example.arn,
    jdbcTargets: [{
        connectionName: aws_glue_connection.example.name,
        path: "database-name/%",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Crawler("example",
    database_name=aws_glue_catalog_database["example"]["name"],
    role=aws_iam_role["example"]["arn"],
    jdbc_targets=[aws.glue.CrawlerJdbcTargetArgs(
        connection_name=aws_glue_connection["example"]["name"],
        path="database-name/%",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Crawler("example", new()
    {
        DatabaseName = aws_glue_catalog_database.Example.Name,
        Role = aws_iam_role.Example.Arn,
        JdbcTargets = new[]
        {
            new Aws.Glue.Inputs.CrawlerJdbcTargetArgs
            {
                ConnectionName = aws_glue_connection.Example.Name,
                Path = "database-name/%",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCrawler(ctx, "example", &glue.CrawlerArgs{
			DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
			Role:         pulumi.Any(aws_iam_role.Example.Arn),
			JdbcTargets: glue.CrawlerJdbcTargetArray{
				&glue.CrawlerJdbcTargetArgs{
					ConnectionName: pulumi.Any(aws_glue_connection.Example.Name),
					Path:           pulumi.String("database-name/%"),
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
import com.pulumi.aws.glue.Crawler;
import com.pulumi.aws.glue.CrawlerArgs;
import com.pulumi.aws.glue.inputs.CrawlerJdbcTargetArgs;
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
        var example = new Crawler("example", CrawlerArgs.builder()        
            .databaseName(aws_glue_catalog_database.example().name())
            .role(aws_iam_role.example().arn())
            .jdbcTargets(CrawlerJdbcTargetArgs.builder()
                .connectionName(aws_glue_connection.example().name())
                .path("database-name/%")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Crawler
    properties:
      databaseName: ${aws_glue_catalog_database.example.name}
      role: ${aws_iam_role.example.arn}
      jdbcTargets:
        - connectionName: ${aws_glue_connection.example.name}
          path: database-name/%
```
{{% /example %}}
{{% example %}}
### S3 Target Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Crawler("example", {
    databaseName: aws_glue_catalog_database.example.name,
    role: aws_iam_role.example.arn,
    s3Targets: [{
        path: `s3://${aws_s3_bucket.example.bucket}`,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Crawler("example",
    database_name=aws_glue_catalog_database["example"]["name"],
    role=aws_iam_role["example"]["arn"],
    s3_targets=[aws.glue.CrawlerS3TargetArgs(
        path=f"s3://{aws_s3_bucket['example']['bucket']}",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Crawler("example", new()
    {
        DatabaseName = aws_glue_catalog_database.Example.Name,
        Role = aws_iam_role.Example.Arn,
        S3Targets = new[]
        {
            new Aws.Glue.Inputs.CrawlerS3TargetArgs
            {
                Path = $"s3://{aws_s3_bucket.Example.Bucket}",
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCrawler(ctx, "example", &glue.CrawlerArgs{
			DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
			Role:         pulumi.Any(aws_iam_role.Example.Arn),
			S3Targets: glue.CrawlerS3TargetArray{
				&glue.CrawlerS3TargetArgs{
					Path: pulumi.String(fmt.Sprintf("s3://%v", aws_s3_bucket.Example.Bucket)),
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
import com.pulumi.aws.glue.Crawler;
import com.pulumi.aws.glue.CrawlerArgs;
import com.pulumi.aws.glue.inputs.CrawlerS3TargetArgs;
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
        var example = new Crawler("example", CrawlerArgs.builder()        
            .databaseName(aws_glue_catalog_database.example().name())
            .role(aws_iam_role.example().arn())
            .s3Targets(CrawlerS3TargetArgs.builder()
                .path(String.format("s3://%s", aws_s3_bucket.example().bucket()))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Crawler
    properties:
      databaseName: ${aws_glue_catalog_database.example.name}
      role: ${aws_iam_role.example.arn}
      s3Targets:
        - path: s3://${aws_s3_bucket.example.bucket}
```
{{% /example %}}
{{% example %}}
### Catalog Target Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Crawler("example", {
    databaseName: aws_glue_catalog_database.example.name,
    role: aws_iam_role.example.arn,
    catalogTargets: [{
        databaseName: aws_glue_catalog_database.example.name,
        tables: [aws_glue_catalog_table.example.name],
    }],
    schemaChangePolicy: {
        deleteBehavior: "LOG",
    },
    configuration: `{
  "Version":1.0,
  "Grouping": {
    "TableGroupingPolicy": "CombineCompatibleSchemas"
  }
}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Crawler("example",
    database_name=aws_glue_catalog_database["example"]["name"],
    role=aws_iam_role["example"]["arn"],
    catalog_targets=[aws.glue.CrawlerCatalogTargetArgs(
        database_name=aws_glue_catalog_database["example"]["name"],
        tables=[aws_glue_catalog_table["example"]["name"]],
    )],
    schema_change_policy=aws.glue.CrawlerSchemaChangePolicyArgs(
        delete_behavior="LOG",
    ),
    configuration="""{
  "Version":1.0,
  "Grouping": {
    "TableGroupingPolicy": "CombineCompatibleSchemas"
  }
}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Crawler("example", new()
    {
        DatabaseName = aws_glue_catalog_database.Example.Name,
        Role = aws_iam_role.Example.Arn,
        CatalogTargets = new[]
        {
            new Aws.Glue.Inputs.CrawlerCatalogTargetArgs
            {
                DatabaseName = aws_glue_catalog_database.Example.Name,
                Tables = new[]
                {
                    aws_glue_catalog_table.Example.Name,
                },
            },
        },
        SchemaChangePolicy = new Aws.Glue.Inputs.CrawlerSchemaChangePolicyArgs
        {
            DeleteBehavior = "LOG",
        },
        Configuration = @"{
  ""Version"":1.0,
  ""Grouping"": {
    ""TableGroupingPolicy"": ""CombineCompatibleSchemas""
  }
}
",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCrawler(ctx, "example", &glue.CrawlerArgs{
			DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
			Role:         pulumi.Any(aws_iam_role.Example.Arn),
			CatalogTargets: glue.CrawlerCatalogTargetArray{
				&glue.CrawlerCatalogTargetArgs{
					DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
					Tables: pulumi.StringArray{
						aws_glue_catalog_table.Example.Name,
					},
				},
			},
			SchemaChangePolicy: &glue.CrawlerSchemaChangePolicyArgs{
				DeleteBehavior: pulumi.String("LOG"),
			},
			Configuration: pulumi.String(`{
  "Version":1.0,
  "Grouping": {
    "TableGroupingPolicy": "CombineCompatibleSchemas"
  }
}
`),
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
import com.pulumi.aws.glue.Crawler;
import com.pulumi.aws.glue.CrawlerArgs;
import com.pulumi.aws.glue.inputs.CrawlerCatalogTargetArgs;
import com.pulumi.aws.glue.inputs.CrawlerSchemaChangePolicyArgs;
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
        var example = new Crawler("example", CrawlerArgs.builder()        
            .databaseName(aws_glue_catalog_database.example().name())
            .role(aws_iam_role.example().arn())
            .catalogTargets(CrawlerCatalogTargetArgs.builder()
                .databaseName(aws_glue_catalog_database.example().name())
                .tables(aws_glue_catalog_table.example().name())
                .build())
            .schemaChangePolicy(CrawlerSchemaChangePolicyArgs.builder()
                .deleteBehavior("LOG")
                .build())
            .configuration("""
{
  "Version":1.0,
  "Grouping": {
    "TableGroupingPolicy": "CombineCompatibleSchemas"
  }
}
            """)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Crawler
    properties:
      databaseName: ${aws_glue_catalog_database.example.name}
      role: ${aws_iam_role.example.arn}
      catalogTargets:
        - databaseName: ${aws_glue_catalog_database.example.name}
          tables:
            - ${aws_glue_catalog_table.example.name}
      schemaChangePolicy:
        deleteBehavior: LOG
      configuration: |
        {
          "Version":1.0,
          "Grouping": {
            "TableGroupingPolicy": "CombineCompatibleSchemas"
          }
        }
```
{{% /example %}}
{{% example %}}
### MongoDB Target Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Crawler("example", {
    databaseName: aws_glue_catalog_database.example.name,
    role: aws_iam_role.example.arn,
    mongodbTargets: [{
        connectionName: aws_glue_connection.example.name,
        path: "database-name/%",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Crawler("example",
    database_name=aws_glue_catalog_database["example"]["name"],
    role=aws_iam_role["example"]["arn"],
    mongodb_targets=[aws.glue.CrawlerMongodbTargetArgs(
        connection_name=aws_glue_connection["example"]["name"],
        path="database-name/%",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Crawler("example", new()
    {
        DatabaseName = aws_glue_catalog_database.Example.Name,
        Role = aws_iam_role.Example.Arn,
        MongodbTargets = new[]
        {
            new Aws.Glue.Inputs.CrawlerMongodbTargetArgs
            {
                ConnectionName = aws_glue_connection.Example.Name,
                Path = "database-name/%",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCrawler(ctx, "example", &glue.CrawlerArgs{
			DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
			Role:         pulumi.Any(aws_iam_role.Example.Arn),
			MongodbTargets: glue.CrawlerMongodbTargetArray{
				&glue.CrawlerMongodbTargetArgs{
					ConnectionName: pulumi.Any(aws_glue_connection.Example.Name),
					Path:           pulumi.String("database-name/%"),
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
import com.pulumi.aws.glue.Crawler;
import com.pulumi.aws.glue.CrawlerArgs;
import com.pulumi.aws.glue.inputs.CrawlerMongodbTargetArgs;
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
        var example = new Crawler("example", CrawlerArgs.builder()        
            .databaseName(aws_glue_catalog_database.example().name())
            .role(aws_iam_role.example().arn())
            .mongodbTargets(CrawlerMongodbTargetArgs.builder()
                .connectionName(aws_glue_connection.example().name())
                .path("database-name/%")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Crawler
    properties:
      databaseName: ${aws_glue_catalog_database.example.name}
      role: ${aws_iam_role.example.arn}
      mongodbTargets:
        - connectionName: ${aws_glue_connection.example.name}
          path: database-name/%
```
{{% /example %}}
{{% example %}}
### Configuration Settings Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const eventsCrawler = new aws.glue.Crawler("eventsCrawler", {
    databaseName: aws_glue_catalog_database.glue_database.name,
    schedule: "cron(0 1 * * ? *)",
    role: aws_iam_role.glue_role.arn,
    tags: _var.tags,
    configuration: JSON.stringify({
        Grouping: {
            TableGroupingPolicy: "CombineCompatibleSchemas",
        },
        CrawlerOutput: {
            Partitions: {
                AddOrUpdateBehavior: "InheritFromTable",
            },
        },
        Version: 1,
    }),
    s3Targets: [{
        path: `s3://${aws_s3_bucket.data_lake_bucket.bucket}`,
    }],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

events_crawler = aws.glue.Crawler("eventsCrawler",
    database_name=aws_glue_catalog_database["glue_database"]["name"],
    schedule="cron(0 1 * * ? *)",
    role=aws_iam_role["glue_role"]["arn"],
    tags=var["tags"],
    configuration=json.dumps({
        "Grouping": {
            "TableGroupingPolicy": "CombineCompatibleSchemas",
        },
        "CrawlerOutput": {
            "Partitions": {
                "AddOrUpdateBehavior": "InheritFromTable",
            },
        },
        "Version": 1,
    }),
    s3_targets=[aws.glue.CrawlerS3TargetArgs(
        path=f"s3://{aws_s3_bucket['data_lake_bucket']['bucket']}",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var eventsCrawler = new Aws.Glue.Crawler("eventsCrawler", new()
    {
        DatabaseName = aws_glue_catalog_database.Glue_database.Name,
        Schedule = "cron(0 1 * * ? *)",
        Role = aws_iam_role.Glue_role.Arn,
        Tags = @var.Tags,
        Configuration = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Grouping"] = new Dictionary<string, object?>
            {
                ["TableGroupingPolicy"] = "CombineCompatibleSchemas",
            },
            ["CrawlerOutput"] = new Dictionary<string, object?>
            {
                ["Partitions"] = new Dictionary<string, object?>
                {
                    ["AddOrUpdateBehavior"] = "InheritFromTable",
                },
            },
            ["Version"] = 1,
        }),
        S3Targets = new[]
        {
            new Aws.Glue.Inputs.CrawlerS3TargetArgs
            {
                Path = $"s3://{aws_s3_bucket.Data_lake_bucket.Bucket}",
            },
        },
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Grouping": map[string]interface{}{
				"TableGroupingPolicy": "CombineCompatibleSchemas",
			},
			"CrawlerOutput": map[string]interface{}{
				"Partitions": map[string]interface{}{
					"AddOrUpdateBehavior": "InheritFromTable",
				},
			},
			"Version": 1,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = glue.NewCrawler(ctx, "eventsCrawler", &glue.CrawlerArgs{
			DatabaseName:  pulumi.Any(aws_glue_catalog_database.Glue_database.Name),
			Schedule:      pulumi.String("cron(0 1 * * ? *)"),
			Role:          pulumi.Any(aws_iam_role.Glue_role.Arn),
			Tags:          pulumi.Any(_var.Tags),
			Configuration: pulumi.String(json0),
			S3Targets: glue.CrawlerS3TargetArray{
				&glue.CrawlerS3TargetArgs{
					Path: pulumi.String(fmt.Sprintf("s3://%v", aws_s3_bucket.Data_lake_bucket.Bucket)),
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
import com.pulumi.aws.glue.Crawler;
import com.pulumi.aws.glue.CrawlerArgs;
import com.pulumi.aws.glue.inputs.CrawlerS3TargetArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var eventsCrawler = new Crawler("eventsCrawler", CrawlerArgs.builder()        
            .databaseName(aws_glue_catalog_database.glue_database().name())
            .schedule("cron(0 1 * * ? *)")
            .role(aws_iam_role.glue_role().arn())
            .tags(var_.tags())
            .configuration(serializeJson(
                jsonObject(
                    jsonProperty("Grouping", jsonObject(
                        jsonProperty("TableGroupingPolicy", "CombineCompatibleSchemas")
                    )),
                    jsonProperty("CrawlerOutput", jsonObject(
                        jsonProperty("Partitions", jsonObject(
                            jsonProperty("AddOrUpdateBehavior", "InheritFromTable")
                        ))
                    )),
                    jsonProperty("Version", 1)
                )))
            .s3Targets(CrawlerS3TargetArgs.builder()
                .path(String.format("s3://%s", aws_s3_bucket.data_lake_bucket().bucket()))
                .build())
            .build());

    }
}
```
```yaml
resources:
  eventsCrawler:
    type: aws:glue:Crawler
    properties:
      databaseName: ${aws_glue_catalog_database.glue_database.name}
      schedule: cron(0 1 * * ? *)
      role: ${aws_iam_role.glue_role.arn}
      tags: ${var.tags}
      configuration:
        fn::toJSON:
          Grouping:
            TableGroupingPolicy: CombineCompatibleSchemas
          CrawlerOutput:
            Partitions:
              AddOrUpdateBehavior: InheritFromTable
          Version: 1
      s3Targets:
        - path: s3://${aws_s3_bucket.data_lake_bucket.bucket}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Crawlers using `name`. For example:

```sh
 $ pulumi import aws:glue/crawler:Crawler MyJob MyJob
```
 