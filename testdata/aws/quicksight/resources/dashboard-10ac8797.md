Resource for managing a QuickSight Dashboard.

{{% examples %}}
## Example Usage
{{% example %}}
### From Source Template

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.Dashboard("example", {
    dashboardId: "example-id",
    versionDescription: "version",
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

example = aws.quicksight.Dashboard("example",
    dashboard_id="example-id",
    version_description="version",
    source_entity=aws.quicksight.DashboardSourceEntityArgs(
        source_template=aws.quicksight.DashboardSourceEntitySourceTemplateArgs(
            arn=aws_quicksight_template["source"]["arn"],
            data_set_references=[aws.quicksight.DashboardSourceEntitySourceTemplateDataSetReferenceArgs(
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
    var example = new Aws.Quicksight.Dashboard("example", new()
    {
        DashboardId = "example-id",
        VersionDescription = "version",
        SourceEntity = new Aws.Quicksight.Inputs.DashboardSourceEntityArgs
        {
            SourceTemplate = new Aws.Quicksight.Inputs.DashboardSourceEntitySourceTemplateArgs
            {
                Arn = aws_quicksight_template.Source.Arn,
                DataSetReferences = new[]
                {
                    new Aws.Quicksight.Inputs.DashboardSourceEntitySourceTemplateDataSetReferenceArgs
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
		_, err := quicksight.NewDashboard(ctx, "example", &quicksight.DashboardArgs{
			DashboardId:        pulumi.String("example-id"),
			VersionDescription: pulumi.String("version"),
			SourceEntity: &quicksight.DashboardSourceEntityArgs{
				SourceTemplate: &quicksight.DashboardSourceEntitySourceTemplateArgs{
					Arn: pulumi.Any(aws_quicksight_template.Source.Arn),
					DataSetReferences: quicksight.DashboardSourceEntitySourceTemplateDataSetReferenceArray{
						&quicksight.DashboardSourceEntitySourceTemplateDataSetReferenceArgs{
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
import com.pulumi.aws.quicksight.Dashboard;
import com.pulumi.aws.quicksight.DashboardArgs;
import com.pulumi.aws.quicksight.inputs.DashboardSourceEntityArgs;
import com.pulumi.aws.quicksight.inputs.DashboardSourceEntitySourceTemplateArgs;
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
        var example = new Dashboard("example", DashboardArgs.builder()        
            .dashboardId("example-id")
            .versionDescription("version")
            .sourceEntity(DashboardSourceEntityArgs.builder()
                .sourceTemplate(DashboardSourceEntitySourceTemplateArgs.builder()
                    .arn(aws_quicksight_template.source().arn())
                    .dataSetReferences(DashboardSourceEntitySourceTemplateDataSetReferenceArgs.builder()
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
    type: aws:quicksight:Dashboard
    properties:
      dashboardId: example-id
      versionDescription: version
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
import com.pulumi.aws.quicksight.Dashboard;
import com.pulumi.aws.quicksight.DashboardArgs;
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
        var example = new Dashboard("example", DashboardArgs.builder()        
            .dashboardId("example-id")
            .versionDescription("version")
            .definition(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:Dashboard
    properties:
      dashboardId: example-id
      versionDescription: version
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

Using `pulumi import`, import a QuickSight Dashboard using the AWS account ID and dashboard ID separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:quicksight/dashboard:Dashboard example 123456789012,example-id
```
 