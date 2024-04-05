Manages Cost and Usage Report Definitions.

> *NOTE:* The AWS Cost and Usage Report service is only available in `us-east-1` currently.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCurReportDefinition = new aws.cur.ReportDefinition("exampleCurReportDefinition", {
    additionalArtifacts: [
        "REDSHIFT",
        "QUICKSIGHT",
    ],
    additionalSchemaElements: [
        "RESOURCES",
        "SPLIT_COST_ALLOCATION_DATA",
    ],
    compression: "GZIP",
    format: "textORcsv",
    reportName: "example-cur-report-definition",
    s3Bucket: "example-bucket-name",
    s3Region: "us-east-1",
    timeUnit: "HOURLY",
});
```
```python
import pulumi
import pulumi_aws as aws

example_cur_report_definition = aws.cur.ReportDefinition("exampleCurReportDefinition",
    additional_artifacts=[
        "REDSHIFT",
        "QUICKSIGHT",
    ],
    additional_schema_elements=[
        "RESOURCES",
        "SPLIT_COST_ALLOCATION_DATA",
    ],
    compression="GZIP",
    format="textORcsv",
    report_name="example-cur-report-definition",
    s3_bucket="example-bucket-name",
    s3_region="us-east-1",
    time_unit="HOURLY")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleCurReportDefinition = new Aws.Cur.ReportDefinition("exampleCurReportDefinition", new()
    {
        AdditionalArtifacts = new[]
        {
            "REDSHIFT",
            "QUICKSIGHT",
        },
        AdditionalSchemaElements = new[]
        {
            "RESOURCES",
            "SPLIT_COST_ALLOCATION_DATA",
        },
        Compression = "GZIP",
        Format = "textORcsv",
        ReportName = "example-cur-report-definition",
        S3Bucket = "example-bucket-name",
        S3Region = "us-east-1",
        TimeUnit = "HOURLY",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cur"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cur.NewReportDefinition(ctx, "exampleCurReportDefinition", &cur.ReportDefinitionArgs{
			AdditionalArtifacts: pulumi.StringArray{
				pulumi.String("REDSHIFT"),
				pulumi.String("QUICKSIGHT"),
			},
			AdditionalSchemaElements: pulumi.StringArray{
				pulumi.String("RESOURCES"),
				pulumi.String("SPLIT_COST_ALLOCATION_DATA"),
			},
			Compression: pulumi.String("GZIP"),
			Format:      pulumi.String("textORcsv"),
			ReportName:  pulumi.String("example-cur-report-definition"),
			S3Bucket:    pulumi.String("example-bucket-name"),
			S3Region:    pulumi.String("us-east-1"),
			TimeUnit:    pulumi.String("HOURLY"),
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
import com.pulumi.aws.cur.ReportDefinition;
import com.pulumi.aws.cur.ReportDefinitionArgs;
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
        var exampleCurReportDefinition = new ReportDefinition("exampleCurReportDefinition", ReportDefinitionArgs.builder()        
            .additionalArtifacts(            
                "REDSHIFT",
                "QUICKSIGHT")
            .additionalSchemaElements(            
                "RESOURCES",
                "SPLIT_COST_ALLOCATION_DATA")
            .compression("GZIP")
            .format("textORcsv")
            .reportName("example-cur-report-definition")
            .s3Bucket("example-bucket-name")
            .s3Region("us-east-1")
            .timeUnit("HOURLY")
            .build());

    }
}
```
```yaml
resources:
  exampleCurReportDefinition:
    type: aws:cur:ReportDefinition
    properties:
      additionalArtifacts:
        - REDSHIFT
        - QUICKSIGHT
      additionalSchemaElements:
        - RESOURCES
        - SPLIT_COST_ALLOCATION_DATA
      compression: GZIP
      format: textORcsv
      reportName: example-cur-report-definition
      s3Bucket: example-bucket-name
      s3Region: us-east-1
      timeUnit: HOURLY
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Report Definitions using the `report_name`. For example:

```sh
 $ pulumi import aws:cur/reportDefinition:ReportDefinition example_cur_report_definition example-cur-report-definition
```
 