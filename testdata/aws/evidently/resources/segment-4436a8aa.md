Provides a CloudWatch Evidently Segment resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Segment("example", {
    pattern: "{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}",
    tags: {
        Key1: "example Segment",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Segment("example",
    pattern="{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}",
    tags={
        "Key1": "example Segment",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Segment("example", new()
    {
        Pattern = "{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}",
        Tags = 
        {
            { "Key1", "example Segment" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewSegment(ctx, "example", &evidently.SegmentArgs{
			Pattern: pulumi.String("{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}"),
			Tags: pulumi.StringMap{
				"Key1": pulumi.String("example Segment"),
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
import com.pulumi.aws.evidently.Segment;
import com.pulumi.aws.evidently.SegmentArgs;
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
        var example = new Segment("example", SegmentArgs.builder()        
            .pattern("{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}")
            .tags(Map.of("Key1", "example Segment"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Segment
    properties:
      pattern: '{"Price":[{"numeric":[">",10,"<=",20]}]}'
      tags:
        Key1: example Segment
```
{{% /example %}}
{{% example %}}
### With JSON object in pattern

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Segment("example", {
    pattern: `  {
    "Price": [
      {
        "numeric": [">",10,"<=",20]
      }
    ]
  }
  
`,
    tags: {
        Key1: "example Segment",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Segment("example",
    pattern="""  {
    "Price": [
      {
        "numeric": [">",10,"<=",20]
      }
    ]
  }
  
""",
    tags={
        "Key1": "example Segment",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Segment("example", new()
    {
        Pattern = @"  {
    ""Price"": [
      {
        ""numeric"": ["">"",10,""<="",20]
      }
    ]
  }
  
",
        Tags = 
        {
            { "Key1", "example Segment" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewSegment(ctx, "example", &evidently.SegmentArgs{
			Pattern: pulumi.String(`  {
    "Price": [
      {
        "numeric": [">",10,"<=",20]
      }
    ]
  }
  
`),
			Tags: pulumi.StringMap{
				"Key1": pulumi.String("example Segment"),
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
import com.pulumi.aws.evidently.Segment;
import com.pulumi.aws.evidently.SegmentArgs;
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
        var example = new Segment("example", SegmentArgs.builder()        
            .pattern("""
  {
    "Price": [
      {
        "numeric": [">",10,"<=",20]
      }
    ]
  }
  
            """)
            .tags(Map.of("Key1", "example Segment"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Segment
    properties:
      pattern: "  {\n    \"Price\": [\n      {\n        \"numeric\": [\">\",10,\"<=\",20]\n      }\n    ]\n  }\n  \n"
      tags:
        Key1: example Segment
```
{{% /example %}}
{{% example %}}
### With Description

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Segment("example", {
    description: "example",
    pattern: "{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Segment("example",
    description="example",
    pattern="{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Segment("example", new()
    {
        Description = "example",
        Pattern = "{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewSegment(ctx, "example", &evidently.SegmentArgs{
			Description: pulumi.String("example"),
			Pattern:     pulumi.String("{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}"),
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
import com.pulumi.aws.evidently.Segment;
import com.pulumi.aws.evidently.SegmentArgs;
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
        var example = new Segment("example", SegmentArgs.builder()        
            .description("example")
            .pattern("{\"Price\":[{\"numeric\":[\">\",10,\"<=\",20]}]}")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Segment
    properties:
      description: example
      pattern: '{"Price":[{"numeric":[">",10,"<=",20]}]}'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch Evidently Segment using the `arn`. For example:

```sh
 $ pulumi import aws:evidently/segment:Segment example arn:aws:evidently:us-west-2:123456789012:segment/example
```
 