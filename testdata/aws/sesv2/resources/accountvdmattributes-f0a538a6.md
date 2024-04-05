Resource for managing an AWS SESv2 (Simple Email V2) Account VDM Attributes.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sesv2.AccountVdmAttributes("example", {
    dashboardAttributes: {
        engagementMetrics: "ENABLED",
    },
    guardianAttributes: {
        optimizedSharedDelivery: "ENABLED",
    },
    vdmEnabled: "ENABLED",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.AccountVdmAttributes("example",
    dashboard_attributes=aws.sesv2.AccountVdmAttributesDashboardAttributesArgs(
        engagement_metrics="ENABLED",
    ),
    guardian_attributes=aws.sesv2.AccountVdmAttributesGuardianAttributesArgs(
        optimized_shared_delivery="ENABLED",
    ),
    vdm_enabled="ENABLED")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SesV2.AccountVdmAttributes("example", new()
    {
        DashboardAttributes = new Aws.SesV2.Inputs.AccountVdmAttributesDashboardAttributesArgs
        {
            EngagementMetrics = "ENABLED",
        },
        GuardianAttributes = new Aws.SesV2.Inputs.AccountVdmAttributesGuardianAttributesArgs
        {
            OptimizedSharedDelivery = "ENABLED",
        },
        VdmEnabled = "ENABLED",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sesv2.NewAccountVdmAttributes(ctx, "example", &sesv2.AccountVdmAttributesArgs{
			DashboardAttributes: &sesv2.AccountVdmAttributesDashboardAttributesArgs{
				EngagementMetrics: pulumi.String("ENABLED"),
			},
			GuardianAttributes: &sesv2.AccountVdmAttributesGuardianAttributesArgs{
				OptimizedSharedDelivery: pulumi.String("ENABLED"),
			},
			VdmEnabled: pulumi.String("ENABLED"),
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
import com.pulumi.aws.sesv2.AccountVdmAttributes;
import com.pulumi.aws.sesv2.AccountVdmAttributesArgs;
import com.pulumi.aws.sesv2.inputs.AccountVdmAttributesDashboardAttributesArgs;
import com.pulumi.aws.sesv2.inputs.AccountVdmAttributesGuardianAttributesArgs;
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
        var example = new AccountVdmAttributes("example", AccountVdmAttributesArgs.builder()        
            .dashboardAttributes(AccountVdmAttributesDashboardAttributesArgs.builder()
                .engagementMetrics("ENABLED")
                .build())
            .guardianAttributes(AccountVdmAttributesGuardianAttributesArgs.builder()
                .optimizedSharedDelivery("ENABLED")
                .build())
            .vdmEnabled("ENABLED")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sesv2:AccountVdmAttributes
    properties:
      dashboardAttributes:
        engagementMetrics: ENABLED
      guardianAttributes:
        optimizedSharedDelivery: ENABLED
      vdmEnabled: ENABLED
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SESv2 (Simple Email V2) Account VDM Attributes using the word `ses-account-vdm-attributes`. For example:

```sh
 $ pulumi import aws:sesv2/accountVdmAttributes:AccountVdmAttributes example ses-account-vdm-attributes
```
 