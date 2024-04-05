Provides a resource to manage an Amazon GuardDuty detector.

> **NOTE:** Deleting this resource is equivalent to "disabling" GuardDuty for an AWS region, which removes all existing findings. You can set the `enable` attribute to `false` to instead "suspend" monitoring and feedback reporting while keeping existing data. See the [Suspending or Disabling Amazon GuardDuty documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_suspend-disable.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDetector = new aws.guardduty.Detector("myDetector", {
    datasources: {
        kubernetes: {
            auditLogs: {
                enable: false,
            },
        },
        malwareProtection: {
            scanEc2InstanceWithFindings: {
                ebsVolumes: {
                    enable: true,
                },
            },
        },
        s3Logs: {
            enable: true,
        },
    },
    enable: true,
});
```
```python
import pulumi
import pulumi_aws as aws

my_detector = aws.guardduty.Detector("myDetector",
    datasources=aws.guardduty.DetectorDatasourcesArgs(
        kubernetes=aws.guardduty.DetectorDatasourcesKubernetesArgs(
            audit_logs=aws.guardduty.DetectorDatasourcesKubernetesAuditLogsArgs(
                enable=False,
            ),
        ),
        malware_protection=aws.guardduty.DetectorDatasourcesMalwareProtectionArgs(
            scan_ec2_instance_with_findings=aws.guardduty.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs(
                ebs_volumes=aws.guardduty.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs(
                    enable=True,
                ),
            ),
        ),
        s3_logs=aws.guardduty.DetectorDatasourcesS3LogsArgs(
            enable=True,
        ),
    ),
    enable=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myDetector = new Aws.GuardDuty.Detector("myDetector", new()
    {
        Datasources = new Aws.GuardDuty.Inputs.DetectorDatasourcesArgs
        {
            Kubernetes = new Aws.GuardDuty.Inputs.DetectorDatasourcesKubernetesArgs
            {
                AuditLogs = new Aws.GuardDuty.Inputs.DetectorDatasourcesKubernetesAuditLogsArgs
                {
                    Enable = false,
                },
            },
            MalwareProtection = new Aws.GuardDuty.Inputs.DetectorDatasourcesMalwareProtectionArgs
            {
                ScanEc2InstanceWithFindings = new Aws.GuardDuty.Inputs.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs
                {
                    EbsVolumes = new Aws.GuardDuty.Inputs.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs
                    {
                        Enable = true,
                    },
                },
            },
            S3Logs = new Aws.GuardDuty.Inputs.DetectorDatasourcesS3LogsArgs
            {
                Enable = true,
            },
        },
        Enable = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/guardduty"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := guardduty.NewDetector(ctx, "myDetector", &guardduty.DetectorArgs{
			Datasources: &guardduty.DetectorDatasourcesArgs{
				Kubernetes: &guardduty.DetectorDatasourcesKubernetesArgs{
					AuditLogs: &guardduty.DetectorDatasourcesKubernetesAuditLogsArgs{
						Enable: pulumi.Bool(false),
					},
				},
				MalwareProtection: &guardduty.DetectorDatasourcesMalwareProtectionArgs{
					ScanEc2InstanceWithFindings: &guardduty.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs{
						EbsVolumes: &guardduty.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs{
							Enable: pulumi.Bool(true),
						},
					},
				},
				S3Logs: &guardduty.DetectorDatasourcesS3LogsArgs{
					Enable: pulumi.Bool(true),
				},
			},
			Enable: pulumi.Bool(true),
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
import com.pulumi.aws.guardduty.Detector;
import com.pulumi.aws.guardduty.DetectorArgs;
import com.pulumi.aws.guardduty.inputs.DetectorDatasourcesArgs;
import com.pulumi.aws.guardduty.inputs.DetectorDatasourcesKubernetesArgs;
import com.pulumi.aws.guardduty.inputs.DetectorDatasourcesKubernetesAuditLogsArgs;
import com.pulumi.aws.guardduty.inputs.DetectorDatasourcesMalwareProtectionArgs;
import com.pulumi.aws.guardduty.inputs.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs;
import com.pulumi.aws.guardduty.inputs.DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs;
import com.pulumi.aws.guardduty.inputs.DetectorDatasourcesS3LogsArgs;
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
        var myDetector = new Detector("myDetector", DetectorArgs.builder()        
            .datasources(DetectorDatasourcesArgs.builder()
                .kubernetes(DetectorDatasourcesKubernetesArgs.builder()
                    .auditLogs(DetectorDatasourcesKubernetesAuditLogsArgs.builder()
                        .enable(false)
                        .build())
                    .build())
                .malwareProtection(DetectorDatasourcesMalwareProtectionArgs.builder()
                    .scanEc2InstanceWithFindings(DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs.builder()
                        .ebsVolumes(DetectorDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs.builder()
                            .enable(true)
                            .build())
                        .build())
                    .build())
                .s3Logs(DetectorDatasourcesS3LogsArgs.builder()
                    .enable(true)
                    .build())
                .build())
            .enable(true)
            .build());

    }
}
```
```yaml
resources:
  myDetector:
    type: aws:guardduty:Detector
    properties:
      datasources:
        kubernetes:
          auditLogs:
            enable: false
        malwareProtection:
          scanEc2InstanceWithFindings:
            ebsVolumes:
              enable: true
        s3Logs:
          enable: true
      enable: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GuardDuty detectors using the detector ID. For example:

```sh
 $ pulumi import aws:guardduty/detector:Detector MyDetector 00b00fd5aecc0ab60a708659477e9617
```
 The ID of the detector can be retrieved via the [AWS CLI](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/guardduty/list-detectors.html) using `aws guardduty list-detectors`.

