Manages an EMR Containers (EMR on EKS) Job Template.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.emrcontainers.JobTemplate("example", {jobTemplateData: {
    executionRoleArn: aws_iam_role.example.arn,
    releaseLabel: "emr-6.10.0-latest",
    jobDriver: {
        sparkSqlJobDriver: {
            entryPoint: "default",
        },
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.emrcontainers.JobTemplate("example", job_template_data=aws.emrcontainers.JobTemplateJobTemplateDataArgs(
    execution_role_arn=aws_iam_role["example"]["arn"],
    release_label="emr-6.10.0-latest",
    job_driver=aws.emrcontainers.JobTemplateJobTemplateDataJobDriverArgs(
        spark_sql_job_driver=aws.emrcontainers.JobTemplateJobTemplateDataJobDriverSparkSqlJobDriverArgs(
            entry_point="default",
        ),
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
    var example = new Aws.EmrContainers.JobTemplate("example", new()
    {
        JobTemplateData = new Aws.EmrContainers.Inputs.JobTemplateJobTemplateDataArgs
        {
            ExecutionRoleArn = aws_iam_role.Example.Arn,
            ReleaseLabel = "emr-6.10.0-latest",
            JobDriver = new Aws.EmrContainers.Inputs.JobTemplateJobTemplateDataJobDriverArgs
            {
                SparkSqlJobDriver = new Aws.EmrContainers.Inputs.JobTemplateJobTemplateDataJobDriverSparkSqlJobDriverArgs
                {
                    EntryPoint = "default",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emrcontainers"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emrcontainers.NewJobTemplate(ctx, "example", &emrcontainers.JobTemplateArgs{
			JobTemplateData: &emrcontainers.JobTemplateJobTemplateDataArgs{
				ExecutionRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
				ReleaseLabel:     pulumi.String("emr-6.10.0-latest"),
				JobDriver: &emrcontainers.JobTemplateJobTemplateDataJobDriverArgs{
					SparkSqlJobDriver: &emrcontainers.JobTemplateJobTemplateDataJobDriverSparkSqlJobDriverArgs{
						EntryPoint: pulumi.String("default"),
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
import com.pulumi.aws.emrcontainers.JobTemplate;
import com.pulumi.aws.emrcontainers.JobTemplateArgs;
import com.pulumi.aws.emrcontainers.inputs.JobTemplateJobTemplateDataArgs;
import com.pulumi.aws.emrcontainers.inputs.JobTemplateJobTemplateDataJobDriverArgs;
import com.pulumi.aws.emrcontainers.inputs.JobTemplateJobTemplateDataJobDriverSparkSqlJobDriverArgs;
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
        var example = new JobTemplate("example", JobTemplateArgs.builder()        
            .jobTemplateData(JobTemplateJobTemplateDataArgs.builder()
                .executionRoleArn(aws_iam_role.example().arn())
                .releaseLabel("emr-6.10.0-latest")
                .jobDriver(JobTemplateJobTemplateDataJobDriverArgs.builder()
                    .sparkSqlJobDriver(JobTemplateJobTemplateDataJobDriverSparkSqlJobDriverArgs.builder()
                        .entryPoint("default")
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
    type: aws:emrcontainers:JobTemplate
    properties:
      jobTemplateData:
        executionRoleArn: ${aws_iam_role.example.arn}
        releaseLabel: emr-6.10.0-latest
        jobDriver:
          sparkSqlJobDriver:
            entryPoint: default
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EKS job templates using the `id`. For example:

```sh
 $ pulumi import aws:emrcontainers/jobTemplate:JobTemplate example a1b2c3d4e5f6g7h8i9j10k11l
```
 