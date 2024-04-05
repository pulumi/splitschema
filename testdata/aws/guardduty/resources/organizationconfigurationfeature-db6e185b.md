Provides a resource to manage a single Amazon GuardDuty [organization configuration feature](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty-features-activation-model.html#guardduty-features).

> **NOTE:** Deleting this resource does not disable the organization configuration feature, the resource in simply removed from state instead.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.guardduty.Detector("example", {enable: true});
const eksRuntimeMonitoring = new aws.guardduty.OrganizationConfigurationFeature("eksRuntimeMonitoring", {
    detectorId: example.id,
    autoEnable: "ALL",
    additionalConfigurations: [{
        name: "EKS_ADDON_MANAGEMENT",
        autoEnable: "NEW",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.guardduty.Detector("example", enable=True)
eks_runtime_monitoring = aws.guardduty.OrganizationConfigurationFeature("eksRuntimeMonitoring",
    detector_id=example.id,
    auto_enable="ALL",
    additional_configurations=[aws.guardduty.OrganizationConfigurationFeatureAdditionalConfigurationArgs(
        name="EKS_ADDON_MANAGEMENT",
        auto_enable="NEW",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.GuardDuty.Detector("example", new()
    {
        Enable = true,
    });

    var eksRuntimeMonitoring = new Aws.GuardDuty.OrganizationConfigurationFeature("eksRuntimeMonitoring", new()
    {
        DetectorId = example.Id,
        AutoEnable = "ALL",
        AdditionalConfigurations = new[]
        {
            new Aws.GuardDuty.Inputs.OrganizationConfigurationFeatureAdditionalConfigurationArgs
            {
                Name = "EKS_ADDON_MANAGEMENT",
                AutoEnable = "NEW",
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
		example, err := guardduty.NewDetector(ctx, "example", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = guardduty.NewOrganizationConfigurationFeature(ctx, "eksRuntimeMonitoring", &guardduty.OrganizationConfigurationFeatureArgs{
			DetectorId: example.ID(),
			AutoEnable: pulumi.String("ALL"),
			AdditionalConfigurations: guardduty.OrganizationConfigurationFeatureAdditionalConfigurationArray{
				&guardduty.OrganizationConfigurationFeatureAdditionalConfigurationArgs{
					Name:       pulumi.String("EKS_ADDON_MANAGEMENT"),
					AutoEnable: pulumi.String("NEW"),
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
import com.pulumi.aws.guardduty.OrganizationConfigurationFeature;
import com.pulumi.aws.guardduty.OrganizationConfigurationFeatureArgs;
import com.pulumi.aws.guardduty.inputs.OrganizationConfigurationFeatureAdditionalConfigurationArgs;
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
        var example = new Detector("example", DetectorArgs.builder()        
            .enable(true)
            .build());

        var eksRuntimeMonitoring = new OrganizationConfigurationFeature("eksRuntimeMonitoring", OrganizationConfigurationFeatureArgs.builder()        
            .detectorId(example.id())
            .autoEnable("ALL")
            .additionalConfigurations(OrganizationConfigurationFeatureAdditionalConfigurationArgs.builder()
                .name("EKS_ADDON_MANAGEMENT")
                .autoEnable("NEW")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:guardduty:Detector
    properties:
      enable: true
  eksRuntimeMonitoring:
    type: aws:guardduty:OrganizationConfigurationFeature
    properties:
      detectorId: ${example.id}
      autoEnable: ALL
      additionalConfigurations:
        - name: EKS_ADDON_MANAGEMENT
          autoEnable: NEW
```
{{% /example %}}
{{% /examples %}}