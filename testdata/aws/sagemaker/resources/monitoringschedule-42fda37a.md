Provides a SageMaker monitoring schedule resource.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.sagemaker.MonitoringSchedule("test", {monitoringScheduleConfig: {
    monitoringJobDefinitionName: aws_sagemaker_data_quality_job_definition.test.name,
    monitoringType: "DataQuality",
}});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.sagemaker.MonitoringSchedule("test", monitoring_schedule_config=aws.sagemaker.MonitoringScheduleMonitoringScheduleConfigArgs(
    monitoring_job_definition_name=aws_sagemaker_data_quality_job_definition["test"]["name"],
    monitoring_type="DataQuality",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Sagemaker.MonitoringSchedule("test", new()
    {
        MonitoringScheduleConfig = new Aws.Sagemaker.Inputs.MonitoringScheduleMonitoringScheduleConfigArgs
        {
            MonitoringJobDefinitionName = aws_sagemaker_data_quality_job_definition.Test.Name,
            MonitoringType = "DataQuality",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewMonitoringSchedule(ctx, "test", &sagemaker.MonitoringScheduleArgs{
			MonitoringScheduleConfig: &sagemaker.MonitoringScheduleMonitoringScheduleConfigArgs{
				MonitoringJobDefinitionName: pulumi.Any(aws_sagemaker_data_quality_job_definition.Test.Name),
				MonitoringType:              pulumi.String("DataQuality"),
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
import com.pulumi.aws.sagemaker.MonitoringSchedule;
import com.pulumi.aws.sagemaker.MonitoringScheduleArgs;
import com.pulumi.aws.sagemaker.inputs.MonitoringScheduleMonitoringScheduleConfigArgs;
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
        var test = new MonitoringSchedule("test", MonitoringScheduleArgs.builder()        
            .monitoringScheduleConfig(MonitoringScheduleMonitoringScheduleConfigArgs.builder()
                .monitoringJobDefinitionName(aws_sagemaker_data_quality_job_definition.test().name())
                .monitoringType("DataQuality")
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:sagemaker:MonitoringSchedule
    properties:
      monitoringScheduleConfig:
        monitoringJobDefinitionName: ${aws_sagemaker_data_quality_job_definition.test.name}
        monitoringType: DataQuality
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import monitoring schedules using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/monitoringSchedule:MonitoringSchedule test_monitoring_schedule monitoring-schedule-foo
```
 