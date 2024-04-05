Manages an attachment between one or more existing LF-tags and an existing Lake Formation resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Database Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLfTag = new aws.lakeformation.LfTag("exampleLfTag", {
    key: "right",
    values: [
        "abbey",
        "village",
        "luffield",
        "woodcote",
        "copse",
        "chapel",
        "stowe",
        "club",
    ],
});
const exampleResourceLfTags = new aws.lakeformation.ResourceLfTags("exampleResourceLfTags", {
    database: {
        name: aws_glue_catalog_database.example.name,
    },
    lfTags: [{
        key: exampleLfTag.key,
        value: "stowe",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_lf_tag = aws.lakeformation.LfTag("exampleLfTag",
    key="right",
    values=[
        "abbey",
        "village",
        "luffield",
        "woodcote",
        "copse",
        "chapel",
        "stowe",
        "club",
    ])
example_resource_lf_tags = aws.lakeformation.ResourceLfTags("exampleResourceLfTags",
    database=aws.lakeformation.ResourceLfTagsDatabaseArgs(
        name=aws_glue_catalog_database["example"]["name"],
    ),
    lf_tags=[aws.lakeformation.ResourceLfTagsLfTagArgs(
        key=example_lf_tag.key,
        value="stowe",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLfTag = new Aws.LakeFormation.LfTag("exampleLfTag", new()
    {
        Key = "right",
        Values = new[]
        {
            "abbey",
            "village",
            "luffield",
            "woodcote",
            "copse",
            "chapel",
            "stowe",
            "club",
        },
    });

    var exampleResourceLfTags = new Aws.LakeFormation.ResourceLfTags("exampleResourceLfTags", new()
    {
        Database = new Aws.LakeFormation.Inputs.ResourceLfTagsDatabaseArgs
        {
            Name = aws_glue_catalog_database.Example.Name,
        },
        LfTags = new[]
        {
            new Aws.LakeFormation.Inputs.ResourceLfTagsLfTagArgs
            {
                Key = exampleLfTag.Key,
                Value = "stowe",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLfTag, err := lakeformation.NewLfTag(ctx, "exampleLfTag", &lakeformation.LfTagArgs{
			Key: pulumi.String("right"),
			Values: pulumi.StringArray{
				pulumi.String("abbey"),
				pulumi.String("village"),
				pulumi.String("luffield"),
				pulumi.String("woodcote"),
				pulumi.String("copse"),
				pulumi.String("chapel"),
				pulumi.String("stowe"),
				pulumi.String("club"),
			},
		})
		if err != nil {
			return err
		}
		_, err = lakeformation.NewResourceLfTags(ctx, "exampleResourceLfTags", &lakeformation.ResourceLfTagsArgs{
			Database: &lakeformation.ResourceLfTagsDatabaseArgs{
				Name: pulumi.Any(aws_glue_catalog_database.Example.Name),
			},
			LfTags: lakeformation.ResourceLfTagsLfTagArray{
				&lakeformation.ResourceLfTagsLfTagArgs{
					Key:   exampleLfTag.Key,
					Value: pulumi.String("stowe"),
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
import com.pulumi.aws.lakeformation.LfTag;
import com.pulumi.aws.lakeformation.LfTagArgs;
import com.pulumi.aws.lakeformation.ResourceLfTags;
import com.pulumi.aws.lakeformation.ResourceLfTagsArgs;
import com.pulumi.aws.lakeformation.inputs.ResourceLfTagsDatabaseArgs;
import com.pulumi.aws.lakeformation.inputs.ResourceLfTagsLfTagArgs;
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
        var exampleLfTag = new LfTag("exampleLfTag", LfTagArgs.builder()        
            .key("right")
            .values(            
                "abbey",
                "village",
                "luffield",
                "woodcote",
                "copse",
                "chapel",
                "stowe",
                "club")
            .build());

        var exampleResourceLfTags = new ResourceLfTags("exampleResourceLfTags", ResourceLfTagsArgs.builder()        
            .database(ResourceLfTagsDatabaseArgs.builder()
                .name(aws_glue_catalog_database.example().name())
                .build())
            .lfTags(ResourceLfTagsLfTagArgs.builder()
                .key(exampleLfTag.key())
                .value("stowe")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleLfTag:
    type: aws:lakeformation:LfTag
    properties:
      key: right
      values:
        - abbey
        - village
        - luffield
        - woodcote
        - copse
        - chapel
        - stowe
        - club
  exampleResourceLfTags:
    type: aws:lakeformation:ResourceLfTags
    properties:
      database:
        name: ${aws_glue_catalog_database.example.name}
      lfTags:
        - key: ${exampleLfTag.key}
          value: stowe
```
{{% /example %}}
{{% example %}}
### Multiple Tags Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLfTag = new aws.lakeformation.LfTag("exampleLfTag", {
    key: "right",
    values: [
        "abbey",
        "village",
        "luffield",
        "woodcote",
        "copse",
        "chapel",
        "stowe",
        "club",
    ],
});
const example2 = new aws.lakeformation.LfTag("example2", {
    key: "left",
    values: [
        "farm",
        "theloop",
        "aintree",
        "brooklands",
        "maggotts",
        "becketts",
        "vale",
    ],
});
const exampleResourceLfTags = new aws.lakeformation.ResourceLfTags("exampleResourceLfTags", {
    database: {
        name: aws_glue_catalog_database.example.name,
    },
    lfTags: [
        {
            key: "right",
            value: "luffield",
        },
        {
            key: "left",
            value: "aintree",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_lf_tag = aws.lakeformation.LfTag("exampleLfTag",
    key="right",
    values=[
        "abbey",
        "village",
        "luffield",
        "woodcote",
        "copse",
        "chapel",
        "stowe",
        "club",
    ])
example2 = aws.lakeformation.LfTag("example2",
    key="left",
    values=[
        "farm",
        "theloop",
        "aintree",
        "brooklands",
        "maggotts",
        "becketts",
        "vale",
    ])
example_resource_lf_tags = aws.lakeformation.ResourceLfTags("exampleResourceLfTags",
    database=aws.lakeformation.ResourceLfTagsDatabaseArgs(
        name=aws_glue_catalog_database["example"]["name"],
    ),
    lf_tags=[
        aws.lakeformation.ResourceLfTagsLfTagArgs(
            key="right",
            value="luffield",
        ),
        aws.lakeformation.ResourceLfTagsLfTagArgs(
            key="left",
            value="aintree",
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
    var exampleLfTag = new Aws.LakeFormation.LfTag("exampleLfTag", new()
    {
        Key = "right",
        Values = new[]
        {
            "abbey",
            "village",
            "luffield",
            "woodcote",
            "copse",
            "chapel",
            "stowe",
            "club",
        },
    });

    var example2 = new Aws.LakeFormation.LfTag("example2", new()
    {
        Key = "left",
        Values = new[]
        {
            "farm",
            "theloop",
            "aintree",
            "brooklands",
            "maggotts",
            "becketts",
            "vale",
        },
    });

    var exampleResourceLfTags = new Aws.LakeFormation.ResourceLfTags("exampleResourceLfTags", new()
    {
        Database = new Aws.LakeFormation.Inputs.ResourceLfTagsDatabaseArgs
        {
            Name = aws_glue_catalog_database.Example.Name,
        },
        LfTags = new[]
        {
            new Aws.LakeFormation.Inputs.ResourceLfTagsLfTagArgs
            {
                Key = "right",
                Value = "luffield",
            },
            new Aws.LakeFormation.Inputs.ResourceLfTagsLfTagArgs
            {
                Key = "left",
                Value = "aintree",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lakeformation.NewLfTag(ctx, "exampleLfTag", &lakeformation.LfTagArgs{
			Key: pulumi.String("right"),
			Values: pulumi.StringArray{
				pulumi.String("abbey"),
				pulumi.String("village"),
				pulumi.String("luffield"),
				pulumi.String("woodcote"),
				pulumi.String("copse"),
				pulumi.String("chapel"),
				pulumi.String("stowe"),
				pulumi.String("club"),
			},
		})
		if err != nil {
			return err
		}
		_, err = lakeformation.NewLfTag(ctx, "example2", &lakeformation.LfTagArgs{
			Key: pulumi.String("left"),
			Values: pulumi.StringArray{
				pulumi.String("farm"),
				pulumi.String("theloop"),
				pulumi.String("aintree"),
				pulumi.String("brooklands"),
				pulumi.String("maggotts"),
				pulumi.String("becketts"),
				pulumi.String("vale"),
			},
		})
		if err != nil {
			return err
		}
		_, err = lakeformation.NewResourceLfTags(ctx, "exampleResourceLfTags", &lakeformation.ResourceLfTagsArgs{
			Database: &lakeformation.ResourceLfTagsDatabaseArgs{
				Name: pulumi.Any(aws_glue_catalog_database.Example.Name),
			},
			LfTags: lakeformation.ResourceLfTagsLfTagArray{
				&lakeformation.ResourceLfTagsLfTagArgs{
					Key:   pulumi.String("right"),
					Value: pulumi.String("luffield"),
				},
				&lakeformation.ResourceLfTagsLfTagArgs{
					Key:   pulumi.String("left"),
					Value: pulumi.String("aintree"),
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
import com.pulumi.aws.lakeformation.LfTag;
import com.pulumi.aws.lakeformation.LfTagArgs;
import com.pulumi.aws.lakeformation.ResourceLfTags;
import com.pulumi.aws.lakeformation.ResourceLfTagsArgs;
import com.pulumi.aws.lakeformation.inputs.ResourceLfTagsDatabaseArgs;
import com.pulumi.aws.lakeformation.inputs.ResourceLfTagsLfTagArgs;
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
        var exampleLfTag = new LfTag("exampleLfTag", LfTagArgs.builder()        
            .key("right")
            .values(            
                "abbey",
                "village",
                "luffield",
                "woodcote",
                "copse",
                "chapel",
                "stowe",
                "club")
            .build());

        var example2 = new LfTag("example2", LfTagArgs.builder()        
            .key("left")
            .values(            
                "farm",
                "theloop",
                "aintree",
                "brooklands",
                "maggotts",
                "becketts",
                "vale")
            .build());

        var exampleResourceLfTags = new ResourceLfTags("exampleResourceLfTags", ResourceLfTagsArgs.builder()        
            .database(ResourceLfTagsDatabaseArgs.builder()
                .name(aws_glue_catalog_database.example().name())
                .build())
            .lfTags(            
                ResourceLfTagsLfTagArgs.builder()
                    .key("right")
                    .value("luffield")
                    .build(),
                ResourceLfTagsLfTagArgs.builder()
                    .key("left")
                    .value("aintree")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  exampleLfTag:
    type: aws:lakeformation:LfTag
    properties:
      key: right
      values:
        - abbey
        - village
        - luffield
        - woodcote
        - copse
        - chapel
        - stowe
        - club
  example2:
    type: aws:lakeformation:LfTag
    properties:
      key: left
      values:
        - farm
        - theloop
        - aintree
        - brooklands
        - maggotts
        - becketts
        - vale
  exampleResourceLfTags:
    type: aws:lakeformation:ResourceLfTags
    properties:
      database:
        name: ${aws_glue_catalog_database.example.name}
      lfTags:
        - key: right
          value: luffield
        - key: left
          value: aintree
```
{{% /example %}}
{{% /examples %}}