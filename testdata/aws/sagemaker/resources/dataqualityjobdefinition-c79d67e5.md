Provides a SageMaker data quality job definition resource.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.sagemaker.DataQualityJobDefinition("test", {
    dataQualityAppSpecification: {
        imageUri: data.aws_sagemaker_prebuilt_ecr_image.monitor.registry_path,
    },
    dataQualityJobInput: {
        endpointInput: {
            endpointName: aws_sagemaker_endpoint.my_endpoint.name,
        },
    },
    dataQualityJobOutputConfig: {
        monitoringOutputs: {
            s3Output: {
                s3Uri: `https://${aws_s3_bucket.my_bucket.bucket_regional_domain_name}/output`,
            },
        },
    },
    jobResources: {
        clusterConfig: {
            instanceCount: 1,
            instanceType: "ml.t3.medium",
            volumeSizeInGb: 20,
        },
    },
    roleArn: aws_iam_role.my_role.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.sagemaker.DataQualityJobDefinition("test",
    data_quality_app_specification=aws.sagemaker.DataQualityJobDefinitionDataQualityAppSpecificationArgs(
        image_uri=data["aws_sagemaker_prebuilt_ecr_image"]["monitor"]["registry_path"],
    ),
    data_quality_job_input=aws.sagemaker.DataQualityJobDefinitionDataQualityJobInputArgs(
        endpoint_input=aws.sagemaker.DataQualityJobDefinitionDataQualityJobInputEndpointInputArgs(
            endpoint_name=aws_sagemaker_endpoint["my_endpoint"]["name"],
        ),
    ),
    data_quality_job_output_config=aws.sagemaker.DataQualityJobDefinitionDataQualityJobOutputConfigArgs(
        monitoring_outputs=aws.sagemaker.DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsArgs(
            s3_output=aws.sagemaker.DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3OutputArgs(
                s3_uri=f"https://{aws_s3_bucket['my_bucket']['bucket_regional_domain_name']}/output",
            ),
        ),
    ),
    job_resources=aws.sagemaker.DataQualityJobDefinitionJobResourcesArgs(
        cluster_config=aws.sagemaker.DataQualityJobDefinitionJobResourcesClusterConfigArgs(
            instance_count=1,
            instance_type="ml.t3.medium",
            volume_size_in_gb=20,
        ),
    ),
    role_arn=aws_iam_role["my_role"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Sagemaker.DataQualityJobDefinition("test", new()
    {
        DataQualityAppSpecification = new Aws.Sagemaker.Inputs.DataQualityJobDefinitionDataQualityAppSpecificationArgs
        {
            ImageUri = data.Aws_sagemaker_prebuilt_ecr_image.Monitor.Registry_path,
        },
        DataQualityJobInput = new Aws.Sagemaker.Inputs.DataQualityJobDefinitionDataQualityJobInputArgs
        {
            EndpointInput = new Aws.Sagemaker.Inputs.DataQualityJobDefinitionDataQualityJobInputEndpointInputArgs
            {
                EndpointName = aws_sagemaker_endpoint.My_endpoint.Name,
            },
        },
        DataQualityJobOutputConfig = new Aws.Sagemaker.Inputs.DataQualityJobDefinitionDataQualityJobOutputConfigArgs
        {
            MonitoringOutputs = new Aws.Sagemaker.Inputs.DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsArgs
            {
                S3Output = new Aws.Sagemaker.Inputs.DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3OutputArgs
                {
                    S3Uri = $"https://{aws_s3_bucket.My_bucket.Bucket_regional_domain_name}/output",
                },
            },
        },
        JobResources = new Aws.Sagemaker.Inputs.DataQualityJobDefinitionJobResourcesArgs
        {
            ClusterConfig = new Aws.Sagemaker.Inputs.DataQualityJobDefinitionJobResourcesClusterConfigArgs
            {
                InstanceCount = 1,
                InstanceType = "ml.t3.medium",
                VolumeSizeInGb = 20,
            },
        },
        RoleArn = aws_iam_role.My_role.Arn,
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewDataQualityJobDefinition(ctx, "test", &sagemaker.DataQualityJobDefinitionArgs{
			DataQualityAppSpecification: &sagemaker.DataQualityJobDefinitionDataQualityAppSpecificationArgs{
				ImageUri: pulumi.Any(data.Aws_sagemaker_prebuilt_ecr_image.Monitor.Registry_path),
			},
			DataQualityJobInput: &sagemaker.DataQualityJobDefinitionDataQualityJobInputArgs{
				EndpointInput: &sagemaker.DataQualityJobDefinitionDataQualityJobInputEndpointInputArgs{
					EndpointName: pulumi.Any(aws_sagemaker_endpoint.My_endpoint.Name),
				},
			},
			DataQualityJobOutputConfig: &sagemaker.DataQualityJobDefinitionDataQualityJobOutputConfigArgs{
				MonitoringOutputs: &sagemaker.DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsArgs{
					S3Output: sagemaker.DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3OutputArgs{
						S3Uri: pulumi.String(fmt.Sprintf("https://%v/output", aws_s3_bucket.My_bucket.Bucket_regional_domain_name)),
					},
				},
			},
			JobResources: &sagemaker.DataQualityJobDefinitionJobResourcesArgs{
				ClusterConfig: &sagemaker.DataQualityJobDefinitionJobResourcesClusterConfigArgs{
					InstanceCount:  pulumi.Int(1),
					InstanceType:   pulumi.String("ml.t3.medium"),
					VolumeSizeInGb: pulumi.Int(20),
				},
			},
			RoleArn: pulumi.Any(aws_iam_role.My_role.Arn),
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
import com.pulumi.aws.sagemaker.DataQualityJobDefinition;
import com.pulumi.aws.sagemaker.DataQualityJobDefinitionArgs;
import com.pulumi.aws.sagemaker.inputs.DataQualityJobDefinitionDataQualityAppSpecificationArgs;
import com.pulumi.aws.sagemaker.inputs.DataQualityJobDefinitionDataQualityJobInputArgs;
import com.pulumi.aws.sagemaker.inputs.DataQualityJobDefinitionDataQualityJobInputEndpointInputArgs;
import com.pulumi.aws.sagemaker.inputs.DataQualityJobDefinitionDataQualityJobOutputConfigArgs;
import com.pulumi.aws.sagemaker.inputs.DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsArgs;
import com.pulumi.aws.sagemaker.inputs.DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3OutputArgs;
import com.pulumi.aws.sagemaker.inputs.DataQualityJobDefinitionJobResourcesArgs;
import com.pulumi.aws.sagemaker.inputs.DataQualityJobDefinitionJobResourcesClusterConfigArgs;
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
        var test = new DataQualityJobDefinition("test", DataQualityJobDefinitionArgs.builder()        
            .dataQualityAppSpecification(DataQualityJobDefinitionDataQualityAppSpecificationArgs.builder()
                .imageUri(data.aws_sagemaker_prebuilt_ecr_image().monitor().registry_path())
                .build())
            .dataQualityJobInput(DataQualityJobDefinitionDataQualityJobInputArgs.builder()
                .endpointInput(DataQualityJobDefinitionDataQualityJobInputEndpointInputArgs.builder()
                    .endpointName(aws_sagemaker_endpoint.my_endpoint().name())
                    .build())
                .build())
            .dataQualityJobOutputConfig(DataQualityJobDefinitionDataQualityJobOutputConfigArgs.builder()
                .monitoringOutputs(DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsArgs.builder()
                    .s3Output(DataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3OutputArgs.builder()
                        .s3Uri(String.format("https://%s/output", aws_s3_bucket.my_bucket().bucket_regional_domain_name()))
                        .build())
                    .build())
                .build())
            .jobResources(DataQualityJobDefinitionJobResourcesArgs.builder()
                .clusterConfig(DataQualityJobDefinitionJobResourcesClusterConfigArgs.builder()
                    .instanceCount(1)
                    .instanceType("ml.t3.medium")
                    .volumeSizeInGb(20)
                    .build())
                .build())
            .roleArn(aws_iam_role.my_role().arn())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:sagemaker:DataQualityJobDefinition
    properties:
      dataQualityAppSpecification:
        imageUri: ${data.aws_sagemaker_prebuilt_ecr_image.monitor.registry_path}
      dataQualityJobInput:
        endpointInput:
          endpointName: ${aws_sagemaker_endpoint.my_endpoint.name}
      dataQualityJobOutputConfig:
        monitoringOutputs:
          s3Output:
            s3Uri: https://${aws_s3_bucket.my_bucket.bucket_regional_domain_name}/output
      jobResources:
        clusterConfig:
          instanceCount: 1
          instanceType: ml.t3.medium
          volumeSizeInGb: 20
      roleArn: ${aws_iam_role.my_role.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import data quality job definitions using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/dataQualityJobDefinition:DataQualityJobDefinition test_data_quality_job_definition data-quality-job-definition-foo
```
 