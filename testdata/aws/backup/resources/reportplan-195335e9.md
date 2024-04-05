Provides an AWS Backup Report Plan resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.ReportPlan("example", {
    description: "example description",
    reportDeliveryChannel: {
        formats: [
            "CSV",
            "JSON",
        ],
        s3BucketName: "example-bucket-name",
    },
    reportSetting: {
        reportTemplate: "RESTORE_JOB_REPORT",
    },
    tags: {
        Name: "Example Report Plan",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.backup.ReportPlan("example",
    description="example description",
    report_delivery_channel=aws.backup.ReportPlanReportDeliveryChannelArgs(
        formats=[
            "CSV",
            "JSON",
        ],
        s3_bucket_name="example-bucket-name",
    ),
    report_setting=aws.backup.ReportPlanReportSettingArgs(
        report_template="RESTORE_JOB_REPORT",
    ),
    tags={
        "Name": "Example Report Plan",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Backup.ReportPlan("example", new()
    {
        Description = "example description",
        ReportDeliveryChannel = new Aws.Backup.Inputs.ReportPlanReportDeliveryChannelArgs
        {
            Formats = new[]
            {
                "CSV",
                "JSON",
            },
            S3BucketName = "example-bucket-name",
        },
        ReportSetting = new Aws.Backup.Inputs.ReportPlanReportSettingArgs
        {
            ReportTemplate = "RESTORE_JOB_REPORT",
        },
        Tags = 
        {
            { "Name", "Example Report Plan" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewReportPlan(ctx, "example", &backup.ReportPlanArgs{
			Description: pulumi.String("example description"),
			ReportDeliveryChannel: &backup.ReportPlanReportDeliveryChannelArgs{
				Formats: pulumi.StringArray{
					pulumi.String("CSV"),
					pulumi.String("JSON"),
				},
				S3BucketName: pulumi.String("example-bucket-name"),
			},
			ReportSetting: &backup.ReportPlanReportSettingArgs{
				ReportTemplate: pulumi.String("RESTORE_JOB_REPORT"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Report Plan"),
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
import com.pulumi.aws.backup.ReportPlan;
import com.pulumi.aws.backup.ReportPlanArgs;
import com.pulumi.aws.backup.inputs.ReportPlanReportDeliveryChannelArgs;
import com.pulumi.aws.backup.inputs.ReportPlanReportSettingArgs;
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
        var example = new ReportPlan("example", ReportPlanArgs.builder()        
            .description("example description")
            .reportDeliveryChannel(ReportPlanReportDeliveryChannelArgs.builder()
                .formats(                
                    "CSV",
                    "JSON")
                .s3BucketName("example-bucket-name")
                .build())
            .reportSetting(ReportPlanReportSettingArgs.builder()
                .reportTemplate("RESTORE_JOB_REPORT")
                .build())
            .tags(Map.of("Name", "Example Report Plan"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:backup:ReportPlan
    properties:
      description: example description
      reportDeliveryChannel:
        formats:
          - CSV
          - JSON
        s3BucketName: example-bucket-name
      reportSetting:
        reportTemplate: RESTORE_JOB_REPORT
      tags:
        Name: Example Report Plan
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Backup Report Plan using the `id` which corresponds to the name of the Backup Report Plan. For example:

```sh
 $ pulumi import aws:backup/reportPlan:ReportPlan test <id>
```
 