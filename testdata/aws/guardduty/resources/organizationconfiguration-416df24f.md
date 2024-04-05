Manages the GuardDuty Organization Configuration in the current AWS Region. The AWS account utilizing this resource must have been assigned as a delegated Organization administrator account, e.g., via the `aws.guardduty.OrganizationAdminAccount` resource. More information about Organizations support in GuardDuty can be found in the [GuardDuty User Guide](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html).

> **NOTE:** This is an advanced resource. The provider will automatically assume management of the GuardDuty Organization Configuration without import and perform no actions on removal from the resource configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDetector = new aws.guardduty.Detector("exampleDetector", {enable: true});
const exampleOrganizationConfiguration = new aws.guardduty.OrganizationConfiguration("exampleOrganizationConfiguration", {
    autoEnableOrganizationMembers: "ALL",
    detectorId: exampleDetector.id,
    datasources: {
        s3Logs: {
            autoEnable: true,
        },
        kubernetes: {
            auditLogs: {
                enable: true,
            },
        },
        malwareProtection: {
            scanEc2InstanceWithFindings: {
                ebsVolumes: {
                    autoEnable: true,
                },
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_detector = aws.guardduty.Detector("exampleDetector", enable=True)
example_organization_configuration = aws.guardduty.OrganizationConfiguration("exampleOrganizationConfiguration",
    auto_enable_organization_members="ALL",
    detector_id=example_detector.id,
    datasources=aws.guardduty.OrganizationConfigurationDatasourcesArgs(
        s3_logs=aws.guardduty.OrganizationConfigurationDatasourcesS3LogsArgs(
            auto_enable=True,
        ),
        kubernetes=aws.guardduty.OrganizationConfigurationDatasourcesKubernetesArgs(
            audit_logs=aws.guardduty.OrganizationConfigurationDatasourcesKubernetesAuditLogsArgs(
                enable=True,
            ),
        ),
        malware_protection=aws.guardduty.OrganizationConfigurationDatasourcesMalwareProtectionArgs(
            scan_ec2_instance_with_findings=aws.guardduty.OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs(
                ebs_volumes=aws.guardduty.OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs(
                    auto_enable=True,
                ),
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
    var exampleDetector = new Aws.GuardDuty.Detector("exampleDetector", new()
    {
        Enable = true,
    });

    var exampleOrganizationConfiguration = new Aws.GuardDuty.OrganizationConfiguration("exampleOrganizationConfiguration", new()
    {
        AutoEnableOrganizationMembers = "ALL",
        DetectorId = exampleDetector.Id,
        Datasources = new Aws.GuardDuty.Inputs.OrganizationConfigurationDatasourcesArgs
        {
            S3Logs = new Aws.GuardDuty.Inputs.OrganizationConfigurationDatasourcesS3LogsArgs
            {
                AutoEnable = true,
            },
            Kubernetes = new Aws.GuardDuty.Inputs.OrganizationConfigurationDatasourcesKubernetesArgs
            {
                AuditLogs = new Aws.GuardDuty.Inputs.OrganizationConfigurationDatasourcesKubernetesAuditLogsArgs
                {
                    Enable = true,
                },
            },
            MalwareProtection = new Aws.GuardDuty.Inputs.OrganizationConfigurationDatasourcesMalwareProtectionArgs
            {
                ScanEc2InstanceWithFindings = new Aws.GuardDuty.Inputs.OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs
                {
                    EbsVolumes = new Aws.GuardDuty.Inputs.OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs
                    {
                        AutoEnable = true,
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/guardduty"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDetector, err := guardduty.NewDetector(ctx, "exampleDetector", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = guardduty.NewOrganizationConfiguration(ctx, "exampleOrganizationConfiguration", &guardduty.OrganizationConfigurationArgs{
			AutoEnableOrganizationMembers: pulumi.String("ALL"),
			DetectorId:                    exampleDetector.ID(),
			Datasources: &guardduty.OrganizationConfigurationDatasourcesArgs{
				S3Logs: &guardduty.OrganizationConfigurationDatasourcesS3LogsArgs{
					AutoEnable: pulumi.Bool(true),
				},
				Kubernetes: &guardduty.OrganizationConfigurationDatasourcesKubernetesArgs{
					AuditLogs: &guardduty.OrganizationConfigurationDatasourcesKubernetesAuditLogsArgs{
						Enable: pulumi.Bool(true),
					},
				},
				MalwareProtection: &guardduty.OrganizationConfigurationDatasourcesMalwareProtectionArgs{
					ScanEc2InstanceWithFindings: &guardduty.OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs{
						EbsVolumes: &guardduty.OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs{
							AutoEnable: pulumi.Bool(true),
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
import com.pulumi.aws.guardduty.Detector;
import com.pulumi.aws.guardduty.DetectorArgs;
import com.pulumi.aws.guardduty.OrganizationConfiguration;
import com.pulumi.aws.guardduty.OrganizationConfigurationArgs;
import com.pulumi.aws.guardduty.inputs.OrganizationConfigurationDatasourcesArgs;
import com.pulumi.aws.guardduty.inputs.OrganizationConfigurationDatasourcesS3LogsArgs;
import com.pulumi.aws.guardduty.inputs.OrganizationConfigurationDatasourcesKubernetesArgs;
import com.pulumi.aws.guardduty.inputs.OrganizationConfigurationDatasourcesKubernetesAuditLogsArgs;
import com.pulumi.aws.guardduty.inputs.OrganizationConfigurationDatasourcesMalwareProtectionArgs;
import com.pulumi.aws.guardduty.inputs.OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs;
import com.pulumi.aws.guardduty.inputs.OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs;
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
        var exampleDetector = new Detector("exampleDetector", DetectorArgs.builder()        
            .enable(true)
            .build());

        var exampleOrganizationConfiguration = new OrganizationConfiguration("exampleOrganizationConfiguration", OrganizationConfigurationArgs.builder()        
            .autoEnableOrganizationMembers("ALL")
            .detectorId(exampleDetector.id())
            .datasources(OrganizationConfigurationDatasourcesArgs.builder()
                .s3Logs(OrganizationConfigurationDatasourcesS3LogsArgs.builder()
                    .autoEnable(true)
                    .build())
                .kubernetes(OrganizationConfigurationDatasourcesKubernetesArgs.builder()
                    .auditLogs(OrganizationConfigurationDatasourcesKubernetesAuditLogsArgs.builder()
                        .enable(true)
                        .build())
                    .build())
                .malwareProtection(OrganizationConfigurationDatasourcesMalwareProtectionArgs.builder()
                    .scanEc2InstanceWithFindings(OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs.builder()
                        .ebsVolumes(OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs.builder()
                            .autoEnable(true)
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleDetector:
    type: aws:guardduty:Detector
    properties:
      enable: true
  exampleOrganizationConfiguration:
    type: aws:guardduty:OrganizationConfiguration
    properties:
      autoEnableOrganizationMembers: ALL
      detectorId: ${exampleDetector.id}
      datasources:
        s3Logs:
          autoEnable: true
        kubernetes:
          auditLogs:
            enable: true
        malwareProtection:
          scanEc2InstanceWithFindings:
            ebsVolumes:
              autoEnable: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GuardDuty Organization Configurations using the GuardDuty Detector ID. For example:

```sh
 $ pulumi import aws:guardduty/organizationConfiguration:OrganizationConfiguration example 00b00fd5aecc0ab60a708659477e9617
```
 