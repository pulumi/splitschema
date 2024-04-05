Resource for managing a QuickSight Analysis.

{{% examples %}}
## Example Usage
{{% example %}}
### From Source Template

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.Analysis("example", {
    analysisId: "example-id",
    sourceEntity: {
        sourceTemplate: {
            arn: aws_quicksight_template.source.arn,
            dataSetReferences: [{
                dataSetArn: aws_quicksight_data_set.dataset.arn,
                dataSetPlaceholder: "1",
            }],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.Analysis("example",
    analysis_id="example-id",
    source_entity=aws.quicksight.AnalysisSourceEntityArgs(
        source_template=aws.quicksight.AnalysisSourceEntitySourceTemplateArgs(
            arn=aws_quicksight_template["source"]["arn"],
            data_set_references=[aws.quicksight.AnalysisSourceEntitySourceTemplateDataSetReferenceArgs(
                data_set_arn=aws_quicksight_data_set["dataset"]["arn"],
                data_set_placeholder="1",
            )],
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
    var example = new Aws.Quicksight.Analysis("example", new()
    {
        AnalysisId = "example-id",
        SourceEntity = new Aws.Quicksight.Inputs.AnalysisSourceEntityArgs
        {
            SourceTemplate = new Aws.Quicksight.Inputs.AnalysisSourceEntitySourceTemplateArgs
            {
                Arn = aws_quicksight_template.Source.Arn,
                DataSetReferences = new[]
                {
                    new Aws.Quicksight.Inputs.AnalysisSourceEntitySourceTemplateDataSetReferenceArgs
                    {
                        DataSetArn = aws_quicksight_data_set.Dataset.Arn,
                        DataSetPlaceholder = "1",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewAnalysis(ctx, "example", &quicksight.AnalysisArgs{
			AnalysisId: pulumi.String("example-id"),
			SourceEntity: &quicksight.AnalysisSourceEntityArgs{
				SourceTemplate: &quicksight.AnalysisSourceEntitySourceTemplateArgs{
					Arn: pulumi.Any(aws_quicksight_template.Source.Arn),
					DataSetReferences: quicksight.AnalysisSourceEntitySourceTemplateDataSetReferenceArray{
						&quicksight.AnalysisSourceEntitySourceTemplateDataSetReferenceArgs{
							DataSetArn:         pulumi.Any(aws_quicksight_data_set.Dataset.Arn),
							DataSetPlaceholder: pulumi.String("1"),
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
import com.pulumi.aws.quicksight.Analysis;
import com.pulumi.aws.quicksight.AnalysisArgs;
import com.pulumi.aws.quicksight.inputs.AnalysisSourceEntityArgs;
import com.pulumi.aws.quicksight.inputs.AnalysisSourceEntitySourceTemplateArgs;
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
        var example = new Analysis("example", AnalysisArgs.builder()        
            .analysisId("example-id")
            .sourceEntity(AnalysisSourceEntityArgs.builder()
                .sourceTemplate(AnalysisSourceEntitySourceTemplateArgs.builder()
                    .arn(aws_quicksight_template.source().arn())
                    .dataSetReferences(AnalysisSourceEntitySourceTemplateDataSetReferenceArgs.builder()
                        .dataSetArn(aws_quicksight_data_set.dataset().arn())
                        .dataSetPlaceholder("1")
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
    type: aws:quicksight:Analysis
    properties:
      analysisId: example-id
      sourceEntity:
        sourceTemplate:
          arn: ${aws_quicksight_template.source.arn}
          dataSetReferences:
            - dataSetArn: ${aws_quicksight_data_set.dataset.arn}
              dataSetPlaceholder: '1'
```
{{% /example %}}
{{% example %}}
### With Definition

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.quicksight.Analysis;
import com.pulumi.aws.quicksight.AnalysisArgs;
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
        var example = new Analysis("example", AnalysisArgs.builder()        
            .analysisId("example-id")
            .definition(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:Analysis
    properties:
      analysisId: example-id
      definition:
        dataSetIdentifiersDeclarations:
          - dataSetArn: ${aws_quicksight_data_set.dataset.arn}
            identifier: '1'
        sheets:
          - title: Example
            sheetId: Example1
            visuals:
              - lineChartVisual:
                  visualId: LineChart
                  title:
                    formatText:
                      plainText: Line Chart Example
                  chartConfiguration:
                    fieldWells:
                      lineChartAggregatedFieldWells:
                        categories:
                          - categoricalDimensionField:
                              fieldId: '1'
                              column:
                                dataSetIdentifier: '1'
                                columnName: Column1
                        values:
                          - categoricalMeasureField:
                              fieldId: '2'
                              column:
                                dataSetIdentifier: '1'
                                columnName: Column1
                              aggregationFunction: COUNT
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a QuickSight Analysis using the AWS account ID and analysis ID separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:quicksight/analysis:Analysis example 123456789012,example-id
```
 