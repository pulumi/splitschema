Provides a CloudWatch Evidently Feature resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Feature("example", {
    project: aws_evidently_project.example.name,
    description: "example description",
    variations: [{
        name: "Variation1",
        value: {
            stringValue: "example",
        },
    }],
    tags: {
        Key1: "example Feature",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Feature("example",
    project=aws_evidently_project["example"]["name"],
    description="example description",
    variations=[aws.evidently.FeatureVariationArgs(
        name="Variation1",
        value=aws.evidently.FeatureVariationValueArgs(
            string_value="example",
        ),
    )],
    tags={
        "Key1": "example Feature",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Feature("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        Description = "example description",
        Variations = new[]
        {
            new Aws.Evidently.Inputs.FeatureVariationArgs
            {
                Name = "Variation1",
                Value = new Aws.Evidently.Inputs.FeatureVariationValueArgs
                {
                    StringValue = "example",
                },
            },
        },
        Tags = 
        {
            { "Key1", "example Feature" },
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
		_, err := evidently.NewFeature(ctx, "example", &evidently.FeatureArgs{
			Project:     pulumi.Any(aws_evidently_project.Example.Name),
			Description: pulumi.String("example description"),
			Variations: evidently.FeatureVariationArray{
				&evidently.FeatureVariationArgs{
					Name: pulumi.String("Variation1"),
					Value: &evidently.FeatureVariationValueArgs{
						StringValue: pulumi.String("example"),
					},
				},
			},
			Tags: pulumi.StringMap{
				"Key1": pulumi.String("example Feature"),
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
import com.pulumi.aws.evidently.Feature;
import com.pulumi.aws.evidently.FeatureArgs;
import com.pulumi.aws.evidently.inputs.FeatureVariationArgs;
import com.pulumi.aws.evidently.inputs.FeatureVariationValueArgs;
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
        var example = new Feature("example", FeatureArgs.builder()        
            .project(aws_evidently_project.example().name())
            .description("example description")
            .variations(FeatureVariationArgs.builder()
                .name("Variation1")
                .value(FeatureVariationValueArgs.builder()
                    .stringValue("example")
                    .build())
                .build())
            .tags(Map.of("Key1", "example Feature"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Feature
    properties:
      project: ${aws_evidently_project.example.name}
      description: example description
      variations:
        - name: Variation1
          value:
            stringValue: example
      tags:
        Key1: example Feature
```
{{% /example %}}
{{% example %}}
### With default variation

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Feature("example", {
    project: aws_evidently_project.example.name,
    defaultVariation: "Variation2",
    variations: [
        {
            name: "Variation1",
            value: {
                stringValue: "exampleval1",
            },
        },
        {
            name: "Variation2",
            value: {
                stringValue: "exampleval2",
            },
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Feature("example",
    project=aws_evidently_project["example"]["name"],
    default_variation="Variation2",
    variations=[
        aws.evidently.FeatureVariationArgs(
            name="Variation1",
            value=aws.evidently.FeatureVariationValueArgs(
                string_value="exampleval1",
            ),
        ),
        aws.evidently.FeatureVariationArgs(
            name="Variation2",
            value=aws.evidently.FeatureVariationValueArgs(
                string_value="exampleval2",
            ),
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
    var example = new Aws.Evidently.Feature("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        DefaultVariation = "Variation2",
        Variations = new[]
        {
            new Aws.Evidently.Inputs.FeatureVariationArgs
            {
                Name = "Variation1",
                Value = new Aws.Evidently.Inputs.FeatureVariationValueArgs
                {
                    StringValue = "exampleval1",
                },
            },
            new Aws.Evidently.Inputs.FeatureVariationArgs
            {
                Name = "Variation2",
                Value = new Aws.Evidently.Inputs.FeatureVariationValueArgs
                {
                    StringValue = "exampleval2",
                },
            },
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
		_, err := evidently.NewFeature(ctx, "example", &evidently.FeatureArgs{
			Project:          pulumi.Any(aws_evidently_project.Example.Name),
			DefaultVariation: pulumi.String("Variation2"),
			Variations: evidently.FeatureVariationArray{
				&evidently.FeatureVariationArgs{
					Name: pulumi.String("Variation1"),
					Value: &evidently.FeatureVariationValueArgs{
						StringValue: pulumi.String("exampleval1"),
					},
				},
				&evidently.FeatureVariationArgs{
					Name: pulumi.String("Variation2"),
					Value: &evidently.FeatureVariationValueArgs{
						StringValue: pulumi.String("exampleval2"),
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
import com.pulumi.aws.evidently.Feature;
import com.pulumi.aws.evidently.FeatureArgs;
import com.pulumi.aws.evidently.inputs.FeatureVariationArgs;
import com.pulumi.aws.evidently.inputs.FeatureVariationValueArgs;
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
        var example = new Feature("example", FeatureArgs.builder()        
            .project(aws_evidently_project.example().name())
            .defaultVariation("Variation2")
            .variations(            
                FeatureVariationArgs.builder()
                    .name("Variation1")
                    .value(FeatureVariationValueArgs.builder()
                        .stringValue("exampleval1")
                        .build())
                    .build(),
                FeatureVariationArgs.builder()
                    .name("Variation2")
                    .value(FeatureVariationValueArgs.builder()
                        .stringValue("exampleval2")
                        .build())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Feature
    properties:
      project: ${aws_evidently_project.example.name}
      defaultVariation: Variation2
      variations:
        - name: Variation1
          value:
            stringValue: exampleval1
        - name: Variation2
          value:
            stringValue: exampleval2
```
{{% /example %}}
{{% example %}}
### With entity overrides

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Feature("example", {
    project: aws_evidently_project.example.name,
    entityOverrides: {
        test1: "Variation1",
    },
    variations: [
        {
            name: "Variation1",
            value: {
                stringValue: "exampleval1",
            },
        },
        {
            name: "Variation2",
            value: {
                stringValue: "exampleval2",
            },
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Feature("example",
    project=aws_evidently_project["example"]["name"],
    entity_overrides={
        "test1": "Variation1",
    },
    variations=[
        aws.evidently.FeatureVariationArgs(
            name="Variation1",
            value=aws.evidently.FeatureVariationValueArgs(
                string_value="exampleval1",
            ),
        ),
        aws.evidently.FeatureVariationArgs(
            name="Variation2",
            value=aws.evidently.FeatureVariationValueArgs(
                string_value="exampleval2",
            ),
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
    var example = new Aws.Evidently.Feature("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        EntityOverrides = 
        {
            { "test1", "Variation1" },
        },
        Variations = new[]
        {
            new Aws.Evidently.Inputs.FeatureVariationArgs
            {
                Name = "Variation1",
                Value = new Aws.Evidently.Inputs.FeatureVariationValueArgs
                {
                    StringValue = "exampleval1",
                },
            },
            new Aws.Evidently.Inputs.FeatureVariationArgs
            {
                Name = "Variation2",
                Value = new Aws.Evidently.Inputs.FeatureVariationValueArgs
                {
                    StringValue = "exampleval2",
                },
            },
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
		_, err := evidently.NewFeature(ctx, "example", &evidently.FeatureArgs{
			Project: pulumi.Any(aws_evidently_project.Example.Name),
			EntityOverrides: pulumi.StringMap{
				"test1": pulumi.String("Variation1"),
			},
			Variations: evidently.FeatureVariationArray{
				&evidently.FeatureVariationArgs{
					Name: pulumi.String("Variation1"),
					Value: &evidently.FeatureVariationValueArgs{
						StringValue: pulumi.String("exampleval1"),
					},
				},
				&evidently.FeatureVariationArgs{
					Name: pulumi.String("Variation2"),
					Value: &evidently.FeatureVariationValueArgs{
						StringValue: pulumi.String("exampleval2"),
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
import com.pulumi.aws.evidently.Feature;
import com.pulumi.aws.evidently.FeatureArgs;
import com.pulumi.aws.evidently.inputs.FeatureVariationArgs;
import com.pulumi.aws.evidently.inputs.FeatureVariationValueArgs;
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
        var example = new Feature("example", FeatureArgs.builder()        
            .project(aws_evidently_project.example().name())
            .entityOverrides(Map.of("test1", "Variation1"))
            .variations(            
                FeatureVariationArgs.builder()
                    .name("Variation1")
                    .value(FeatureVariationValueArgs.builder()
                        .stringValue("exampleval1")
                        .build())
                    .build(),
                FeatureVariationArgs.builder()
                    .name("Variation2")
                    .value(FeatureVariationValueArgs.builder()
                        .stringValue("exampleval2")
                        .build())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Feature
    properties:
      project: ${aws_evidently_project.example.name}
      entityOverrides:
        test1: Variation1
      variations:
        - name: Variation1
          value:
            stringValue: exampleval1
        - name: Variation2
          value:
            stringValue: exampleval2
```
{{% /example %}}
{{% example %}}
### With evaluation strategy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Feature("example", {
    project: aws_evidently_project.example.name,
    evaluationStrategy: "ALL_RULES",
    entityOverrides: {
        test1: "Variation1",
    },
    variations: [{
        name: "Variation1",
        value: {
            stringValue: "exampleval1",
        },
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Feature("example",
    project=aws_evidently_project["example"]["name"],
    evaluation_strategy="ALL_RULES",
    entity_overrides={
        "test1": "Variation1",
    },
    variations=[aws.evidently.FeatureVariationArgs(
        name="Variation1",
        value=aws.evidently.FeatureVariationValueArgs(
            string_value="exampleval1",
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Feature("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        EvaluationStrategy = "ALL_RULES",
        EntityOverrides = 
        {
            { "test1", "Variation1" },
        },
        Variations = new[]
        {
            new Aws.Evidently.Inputs.FeatureVariationArgs
            {
                Name = "Variation1",
                Value = new Aws.Evidently.Inputs.FeatureVariationValueArgs
                {
                    StringValue = "exampleval1",
                },
            },
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
		_, err := evidently.NewFeature(ctx, "example", &evidently.FeatureArgs{
			Project:            pulumi.Any(aws_evidently_project.Example.Name),
			EvaluationStrategy: pulumi.String("ALL_RULES"),
			EntityOverrides: pulumi.StringMap{
				"test1": pulumi.String("Variation1"),
			},
			Variations: evidently.FeatureVariationArray{
				&evidently.FeatureVariationArgs{
					Name: pulumi.String("Variation1"),
					Value: &evidently.FeatureVariationValueArgs{
						StringValue: pulumi.String("exampleval1"),
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
import com.pulumi.aws.evidently.Feature;
import com.pulumi.aws.evidently.FeatureArgs;
import com.pulumi.aws.evidently.inputs.FeatureVariationArgs;
import com.pulumi.aws.evidently.inputs.FeatureVariationValueArgs;
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
        var example = new Feature("example", FeatureArgs.builder()        
            .project(aws_evidently_project.example().name())
            .evaluationStrategy("ALL_RULES")
            .entityOverrides(Map.of("test1", "Variation1"))
            .variations(FeatureVariationArgs.builder()
                .name("Variation1")
                .value(FeatureVariationValueArgs.builder()
                    .stringValue("exampleval1")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Feature
    properties:
      project: ${aws_evidently_project.example.name}
      evaluationStrategy: ALL_RULES
      entityOverrides:
        test1: Variation1
      variations:
        - name: Variation1
          value:
            stringValue: exampleval1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch Evidently Feature using the feature `name` and `name` or `arn` of the hosting CloudWatch Evidently Project separated by a `:`. For example:

```sh
 $ pulumi import aws:evidently/feature:Feature example exampleFeatureName:arn:aws:evidently:us-east-1:123456789012:project/example
```
 