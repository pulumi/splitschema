Resource for managing a QuickSight Template.

{{% examples %}}
## Example Usage
{{% example %}}
### From Source Template

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.Template("example", {
    templateId: "example-id",
    versionDescription: "version",
    sourceEntity: {
        sourceTemplate: {
            arn: aws_quicksight_template.source.arn,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.Template("example",
    template_id="example-id",
    version_description="version",
    source_entity=aws.quicksight.TemplateSourceEntityArgs(
        source_template=aws.quicksight.TemplateSourceEntitySourceTemplateArgs(
            arn=aws_quicksight_template["source"]["arn"],
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
    var example = new Aws.Quicksight.Template("example", new()
    {
        TemplateId = "example-id",
        VersionDescription = "version",
        SourceEntity = new Aws.Quicksight.Inputs.TemplateSourceEntityArgs
        {
            SourceTemplate = new Aws.Quicksight.Inputs.TemplateSourceEntitySourceTemplateArgs
            {
                Arn = aws_quicksight_template.Source.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewTemplate(ctx, "example", &quicksight.TemplateArgs{
			TemplateId:         pulumi.String("example-id"),
			VersionDescription: pulumi.String("version"),
			SourceEntity: &quicksight.TemplateSourceEntityArgs{
				SourceTemplate: &quicksight.TemplateSourceEntitySourceTemplateArgs{
					Arn: pulumi.Any(aws_quicksight_template.Source.Arn),
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
import com.pulumi.aws.quicksight.Template;
import com.pulumi.aws.quicksight.TemplateArgs;
import com.pulumi.aws.quicksight.inputs.TemplateSourceEntityArgs;
import com.pulumi.aws.quicksight.inputs.TemplateSourceEntitySourceTemplateArgs;
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
        var example = new Template("example", TemplateArgs.builder()        
            .templateId("example-id")
            .versionDescription("version")
            .sourceEntity(TemplateSourceEntityArgs.builder()
                .sourceTemplate(TemplateSourceEntitySourceTemplateArgs.builder()
                    .arn(aws_quicksight_template.source().arn())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:Template
    properties:
      templateId: example-id
      versionDescription: version
      sourceEntity:
        sourceTemplate:
          arn: ${aws_quicksight_template.source.arn}
```
{{% /example %}}
{{% example %}}
### With Definition

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.quicksight.Template;
import com.pulumi.aws.quicksight.TemplateArgs;
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
        var example = new Template("example", TemplateArgs.builder()        
            .definition(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .templateId("example-id")
            .versionDescription("version")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:Template
    properties:
      definition:
        dataSetConfigurations:
          - dataSetSchema:
              columnSchemaList:
                - dataType: STRING
                  name: Column1
                - dataType: INTEGER
                  name: Column2
            placeholder: '1'
        sheets:
          - sheetId: Test1
            title: Test
            visuals:
              - barChartVisual:
                  chartConfiguration:
                    fieldWells:
                      barChartAggregatedFieldWells:
                        category:
                          - categoricalDimensionField:
                              column:
                                columnName: Column1
                                dataSetIdentifier: '1'
                              fieldId: '1'
                        values:
                          - numericalMeasureField:
                              aggregationFunction:
                                simpleNumericalAggregation: SUM
                              column:
                                columnName: Column2
                                dataSetIdentifier: '1'
                              fieldId: '2'
                  visualId: BarChart
      templateId: example-id
      versionDescription: version
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a QuickSight Template using the AWS account ID and template ID separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:quicksight/template:Template example 123456789012,example-id
```
 