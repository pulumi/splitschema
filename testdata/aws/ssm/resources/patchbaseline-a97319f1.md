Provides an SSM Patch Baseline resource.

> **NOTE on Patch Baselines:** The `approved_patches` and `approval_rule` are
both marked as optional fields, but the Patch Baseline requires that at least one
of them is specified.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

Using `approved_patches` only.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const production = new aws.ssm.PatchBaseline("production", {approvedPatches: ["KB123456"]});
```
```python
import pulumi
import pulumi_aws as aws

production = aws.ssm.PatchBaseline("production", approved_patches=["KB123456"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var production = new Aws.Ssm.PatchBaseline("production", new()
    {
        ApprovedPatches = new[]
        {
            "KB123456",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewPatchBaseline(ctx, "production", &ssm.PatchBaselineArgs{
			ApprovedPatches: pulumi.StringArray{
				pulumi.String("KB123456"),
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
import com.pulumi.aws.ssm.PatchBaseline;
import com.pulumi.aws.ssm.PatchBaselineArgs;
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
        var production = new PatchBaseline("production", PatchBaselineArgs.builder()        
            .approvedPatches("KB123456")
            .build());

    }
}
```
```yaml
resources:
  production:
    type: aws:ssm:PatchBaseline
    properties:
      approvedPatches:
        - KB123456
```
{{% /example %}}
{{% example %}}
### Advanced Usage, specifying patch filters

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const production = new aws.ssm.PatchBaseline("production", {
    approvalRules: [
        {
            approveAfterDays: 7,
            complianceLevel: "HIGH",
            patchFilters: [
                {
                    key: "PRODUCT",
                    values: ["WindowsServer2016"],
                },
                {
                    key: "CLASSIFICATION",
                    values: [
                        "CriticalUpdates",
                        "SecurityUpdates",
                        "Updates",
                    ],
                },
                {
                    key: "MSRC_SEVERITY",
                    values: [
                        "Critical",
                        "Important",
                        "Moderate",
                    ],
                },
            ],
        },
        {
            approveAfterDays: 7,
            patchFilters: [{
                key: "PRODUCT",
                values: ["WindowsServer2012"],
            }],
        },
    ],
    approvedPatches: [
        "KB123456",
        "KB456789",
    ],
    description: "Patch Baseline Description",
    globalFilters: [
        {
            key: "PRODUCT",
            values: ["WindowsServer2008"],
        },
        {
            key: "CLASSIFICATION",
            values: ["ServicePacks"],
        },
        {
            key: "MSRC_SEVERITY",
            values: ["Low"],
        },
    ],
    rejectedPatches: ["KB987654"],
});
```
```python
import pulumi
import pulumi_aws as aws

production = aws.ssm.PatchBaseline("production",
    approval_rules=[
        aws.ssm.PatchBaselineApprovalRuleArgs(
            approve_after_days=7,
            compliance_level="HIGH",
            patch_filters=[
                aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                    key="PRODUCT",
                    values=["WindowsServer2016"],
                ),
                aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                    key="CLASSIFICATION",
                    values=[
                        "CriticalUpdates",
                        "SecurityUpdates",
                        "Updates",
                    ],
                ),
                aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                    key="MSRC_SEVERITY",
                    values=[
                        "Critical",
                        "Important",
                        "Moderate",
                    ],
                ),
            ],
        ),
        aws.ssm.PatchBaselineApprovalRuleArgs(
            approve_after_days=7,
            patch_filters=[aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                key="PRODUCT",
                values=["WindowsServer2012"],
            )],
        ),
    ],
    approved_patches=[
        "KB123456",
        "KB456789",
    ],
    description="Patch Baseline Description",
    global_filters=[
        aws.ssm.PatchBaselineGlobalFilterArgs(
            key="PRODUCT",
            values=["WindowsServer2008"],
        ),
        aws.ssm.PatchBaselineGlobalFilterArgs(
            key="CLASSIFICATION",
            values=["ServicePacks"],
        ),
        aws.ssm.PatchBaselineGlobalFilterArgs(
            key="MSRC_SEVERITY",
            values=["Low"],
        ),
    ],
    rejected_patches=["KB987654"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var production = new Aws.Ssm.PatchBaseline("production", new()
    {
        ApprovalRules = new[]
        {
            new Aws.Ssm.Inputs.PatchBaselineApprovalRuleArgs
            {
                ApproveAfterDays = 7,
                ComplianceLevel = "HIGH",
                PatchFilters = new[]
                {
                    new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
                    {
                        Key = "PRODUCT",
                        Values = new[]
                        {
                            "WindowsServer2016",
                        },
                    },
                    new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
                    {
                        Key = "CLASSIFICATION",
                        Values = new[]
                        {
                            "CriticalUpdates",
                            "SecurityUpdates",
                            "Updates",
                        },
                    },
                    new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
                    {
                        Key = "MSRC_SEVERITY",
                        Values = new[]
                        {
                            "Critical",
                            "Important",
                            "Moderate",
                        },
                    },
                },
            },
            new Aws.Ssm.Inputs.PatchBaselineApprovalRuleArgs
            {
                ApproveAfterDays = 7,
                PatchFilters = new[]
                {
                    new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
                    {
                        Key = "PRODUCT",
                        Values = new[]
                        {
                            "WindowsServer2012",
                        },
                    },
                },
            },
        },
        ApprovedPatches = new[]
        {
            "KB123456",
            "KB456789",
        },
        Description = "Patch Baseline Description",
        GlobalFilters = new[]
        {
            new Aws.Ssm.Inputs.PatchBaselineGlobalFilterArgs
            {
                Key = "PRODUCT",
                Values = new[]
                {
                    "WindowsServer2008",
                },
            },
            new Aws.Ssm.Inputs.PatchBaselineGlobalFilterArgs
            {
                Key = "CLASSIFICATION",
                Values = new[]
                {
                    "ServicePacks",
                },
            },
            new Aws.Ssm.Inputs.PatchBaselineGlobalFilterArgs
            {
                Key = "MSRC_SEVERITY",
                Values = new[]
                {
                    "Low",
                },
            },
        },
        RejectedPatches = new[]
        {
            "KB987654",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewPatchBaseline(ctx, "production", &ssm.PatchBaselineArgs{
			ApprovalRules: ssm.PatchBaselineApprovalRuleArray{
				&ssm.PatchBaselineApprovalRuleArgs{
					ApproveAfterDays: pulumi.Int(7),
					ComplianceLevel:  pulumi.String("HIGH"),
					PatchFilters: ssm.PatchBaselineApprovalRulePatchFilterArray{
						&ssm.PatchBaselineApprovalRulePatchFilterArgs{
							Key: pulumi.String("PRODUCT"),
							Values: pulumi.StringArray{
								pulumi.String("WindowsServer2016"),
							},
						},
						&ssm.PatchBaselineApprovalRulePatchFilterArgs{
							Key: pulumi.String("CLASSIFICATION"),
							Values: pulumi.StringArray{
								pulumi.String("CriticalUpdates"),
								pulumi.String("SecurityUpdates"),
								pulumi.String("Updates"),
							},
						},
						&ssm.PatchBaselineApprovalRulePatchFilterArgs{
							Key: pulumi.String("MSRC_SEVERITY"),
							Values: pulumi.StringArray{
								pulumi.String("Critical"),
								pulumi.String("Important"),
								pulumi.String("Moderate"),
							},
						},
					},
				},
				&ssm.PatchBaselineApprovalRuleArgs{
					ApproveAfterDays: pulumi.Int(7),
					PatchFilters: ssm.PatchBaselineApprovalRulePatchFilterArray{
						&ssm.PatchBaselineApprovalRulePatchFilterArgs{
							Key: pulumi.String("PRODUCT"),
							Values: pulumi.StringArray{
								pulumi.String("WindowsServer2012"),
							},
						},
					},
				},
			},
			ApprovedPatches: pulumi.StringArray{
				pulumi.String("KB123456"),
				pulumi.String("KB456789"),
			},
			Description: pulumi.String("Patch Baseline Description"),
			GlobalFilters: ssm.PatchBaselineGlobalFilterArray{
				&ssm.PatchBaselineGlobalFilterArgs{
					Key: pulumi.String("PRODUCT"),
					Values: pulumi.StringArray{
						pulumi.String("WindowsServer2008"),
					},
				},
				&ssm.PatchBaselineGlobalFilterArgs{
					Key: pulumi.String("CLASSIFICATION"),
					Values: pulumi.StringArray{
						pulumi.String("ServicePacks"),
					},
				},
				&ssm.PatchBaselineGlobalFilterArgs{
					Key: pulumi.String("MSRC_SEVERITY"),
					Values: pulumi.StringArray{
						pulumi.String("Low"),
					},
				},
			},
			RejectedPatches: pulumi.StringArray{
				pulumi.String("KB987654"),
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
import com.pulumi.aws.ssm.PatchBaseline;
import com.pulumi.aws.ssm.PatchBaselineArgs;
import com.pulumi.aws.ssm.inputs.PatchBaselineApprovalRuleArgs;
import com.pulumi.aws.ssm.inputs.PatchBaselineGlobalFilterArgs;
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
        var production = new PatchBaseline("production", PatchBaselineArgs.builder()        
            .approvalRules(            
                PatchBaselineApprovalRuleArgs.builder()
                    .approveAfterDays(7)
                    .complianceLevel("HIGH")
                    .patchFilters(                    
                        PatchBaselineApprovalRulePatchFilterArgs.builder()
                            .key("PRODUCT")
                            .values("WindowsServer2016")
                            .build(),
                        PatchBaselineApprovalRulePatchFilterArgs.builder()
                            .key("CLASSIFICATION")
                            .values(                            
                                "CriticalUpdates",
                                "SecurityUpdates",
                                "Updates")
                            .build(),
                        PatchBaselineApprovalRulePatchFilterArgs.builder()
                            .key("MSRC_SEVERITY")
                            .values(                            
                                "Critical",
                                "Important",
                                "Moderate")
                            .build())
                    .build(),
                PatchBaselineApprovalRuleArgs.builder()
                    .approveAfterDays(7)
                    .patchFilters(PatchBaselineApprovalRulePatchFilterArgs.builder()
                        .key("PRODUCT")
                        .values("WindowsServer2012")
                        .build())
                    .build())
            .approvedPatches(            
                "KB123456",
                "KB456789")
            .description("Patch Baseline Description")
            .globalFilters(            
                PatchBaselineGlobalFilterArgs.builder()
                    .key("PRODUCT")
                    .values("WindowsServer2008")
                    .build(),
                PatchBaselineGlobalFilterArgs.builder()
                    .key("CLASSIFICATION")
                    .values("ServicePacks")
                    .build(),
                PatchBaselineGlobalFilterArgs.builder()
                    .key("MSRC_SEVERITY")
                    .values("Low")
                    .build())
            .rejectedPatches("KB987654")
            .build());

    }
}
```
```yaml
resources:
  production:
    type: aws:ssm:PatchBaseline
    properties:
      approvalRules:
        - approveAfterDays: 7
          complianceLevel: HIGH
          patchFilters:
            - key: PRODUCT
              values:
                - WindowsServer2016
            - key: CLASSIFICATION
              values:
                - CriticalUpdates
                - SecurityUpdates
                - Updates
            - key: MSRC_SEVERITY
              values:
                - Critical
                - Important
                - Moderate
        - approveAfterDays: 7
          patchFilters:
            - key: PRODUCT
              values:
                - WindowsServer2012
      approvedPatches:
        - KB123456
        - KB456789
      description: Patch Baseline Description
      globalFilters:
        - key: PRODUCT
          values:
            - WindowsServer2008
        - key: CLASSIFICATION
          values:
            - ServicePacks
        - key: MSRC_SEVERITY
          values:
            - Low
      rejectedPatches:
        - KB987654
```
{{% /example %}}
{{% example %}}
### Advanced usage, specifying Microsoft application and Windows patch rules

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const windowsOsApps = new aws.ssm.PatchBaseline("windowsOsApps", {
    approvalRules: [
        {
            approveAfterDays: 7,
            patchFilters: [
                {
                    key: "CLASSIFICATION",
                    values: [
                        "CriticalUpdates",
                        "SecurityUpdates",
                    ],
                },
                {
                    key: "MSRC_SEVERITY",
                    values: [
                        "Critical",
                        "Important",
                    ],
                },
            ],
        },
        {
            approveAfterDays: 7,
            patchFilters: [
                {
                    key: "PATCH_SET",
                    values: ["APPLICATION"],
                },
                {
                    key: "PRODUCT",
                    values: [
                        "Office 2013",
                        "Office 2016",
                    ],
                },
            ],
        },
    ],
    description: "Patch both Windows and Microsoft apps",
    operatingSystem: "WINDOWS",
});
```
```python
import pulumi
import pulumi_aws as aws

windows_os_apps = aws.ssm.PatchBaseline("windowsOsApps",
    approval_rules=[
        aws.ssm.PatchBaselineApprovalRuleArgs(
            approve_after_days=7,
            patch_filters=[
                aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                    key="CLASSIFICATION",
                    values=[
                        "CriticalUpdates",
                        "SecurityUpdates",
                    ],
                ),
                aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                    key="MSRC_SEVERITY",
                    values=[
                        "Critical",
                        "Important",
                    ],
                ),
            ],
        ),
        aws.ssm.PatchBaselineApprovalRuleArgs(
            approve_after_days=7,
            patch_filters=[
                aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                    key="PATCH_SET",
                    values=["APPLICATION"],
                ),
                aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                    key="PRODUCT",
                    values=[
                        "Office 2013",
                        "Office 2016",
                    ],
                ),
            ],
        ),
    ],
    description="Patch both Windows and Microsoft apps",
    operating_system="WINDOWS")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var windowsOsApps = new Aws.Ssm.PatchBaseline("windowsOsApps", new()
    {
        ApprovalRules = new[]
        {
            new Aws.Ssm.Inputs.PatchBaselineApprovalRuleArgs
            {
                ApproveAfterDays = 7,
                PatchFilters = new[]
                {
                    new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
                    {
                        Key = "CLASSIFICATION",
                        Values = new[]
                        {
                            "CriticalUpdates",
                            "SecurityUpdates",
                        },
                    },
                    new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
                    {
                        Key = "MSRC_SEVERITY",
                        Values = new[]
                        {
                            "Critical",
                            "Important",
                        },
                    },
                },
            },
            new Aws.Ssm.Inputs.PatchBaselineApprovalRuleArgs
            {
                ApproveAfterDays = 7,
                PatchFilters = new[]
                {
                    new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
                    {
                        Key = "PATCH_SET",
                        Values = new[]
                        {
                            "APPLICATION",
                        },
                    },
                    new Aws.Ssm.Inputs.PatchBaselineApprovalRulePatchFilterArgs
                    {
                        Key = "PRODUCT",
                        Values = new[]
                        {
                            "Office 2013",
                            "Office 2016",
                        },
                    },
                },
            },
        },
        Description = "Patch both Windows and Microsoft apps",
        OperatingSystem = "WINDOWS",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewPatchBaseline(ctx, "windowsOsApps", &ssm.PatchBaselineArgs{
			ApprovalRules: ssm.PatchBaselineApprovalRuleArray{
				&ssm.PatchBaselineApprovalRuleArgs{
					ApproveAfterDays: pulumi.Int(7),
					PatchFilters: ssm.PatchBaselineApprovalRulePatchFilterArray{
						&ssm.PatchBaselineApprovalRulePatchFilterArgs{
							Key: pulumi.String("CLASSIFICATION"),
							Values: pulumi.StringArray{
								pulumi.String("CriticalUpdates"),
								pulumi.String("SecurityUpdates"),
							},
						},
						&ssm.PatchBaselineApprovalRulePatchFilterArgs{
							Key: pulumi.String("MSRC_SEVERITY"),
							Values: pulumi.StringArray{
								pulumi.String("Critical"),
								pulumi.String("Important"),
							},
						},
					},
				},
				&ssm.PatchBaselineApprovalRuleArgs{
					ApproveAfterDays: pulumi.Int(7),
					PatchFilters: ssm.PatchBaselineApprovalRulePatchFilterArray{
						&ssm.PatchBaselineApprovalRulePatchFilterArgs{
							Key: pulumi.String("PATCH_SET"),
							Values: pulumi.StringArray{
								pulumi.String("APPLICATION"),
							},
						},
						&ssm.PatchBaselineApprovalRulePatchFilterArgs{
							Key: pulumi.String("PRODUCT"),
							Values: pulumi.StringArray{
								pulumi.String("Office 2013"),
								pulumi.String("Office 2016"),
							},
						},
					},
				},
			},
			Description:     pulumi.String("Patch both Windows and Microsoft apps"),
			OperatingSystem: pulumi.String("WINDOWS"),
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
import com.pulumi.aws.ssm.PatchBaseline;
import com.pulumi.aws.ssm.PatchBaselineArgs;
import com.pulumi.aws.ssm.inputs.PatchBaselineApprovalRuleArgs;
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
        var windowsOsApps = new PatchBaseline("windowsOsApps", PatchBaselineArgs.builder()        
            .approvalRules(            
                PatchBaselineApprovalRuleArgs.builder()
                    .approveAfterDays(7)
                    .patchFilters(                    
                        PatchBaselineApprovalRulePatchFilterArgs.builder()
                            .key("CLASSIFICATION")
                            .values(                            
                                "CriticalUpdates",
                                "SecurityUpdates")
                            .build(),
                        PatchBaselineApprovalRulePatchFilterArgs.builder()
                            .key("MSRC_SEVERITY")
                            .values(                            
                                "Critical",
                                "Important")
                            .build())
                    .build(),
                PatchBaselineApprovalRuleArgs.builder()
                    .approveAfterDays(7)
                    .patchFilters(                    
                        PatchBaselineApprovalRulePatchFilterArgs.builder()
                            .key("PATCH_SET")
                            .values("APPLICATION")
                            .build(),
                        PatchBaselineApprovalRulePatchFilterArgs.builder()
                            .key("PRODUCT")
                            .values(                            
                                "Office 2013",
                                "Office 2016")
                            .build())
                    .build())
            .description("Patch both Windows and Microsoft apps")
            .operatingSystem("WINDOWS")
            .build());

    }
}
```
```yaml
resources:
  windowsOsApps:
    type: aws:ssm:PatchBaseline
    properties:
      approvalRules:
        - approveAfterDays: 7
          patchFilters:
            - key: CLASSIFICATION
              values:
                - CriticalUpdates
                - SecurityUpdates
            - key: MSRC_SEVERITY
              values:
                - Critical
                - Important
        - approveAfterDays: 7
          patchFilters:
            - key: PATCH_SET
              values:
                - APPLICATION
            - key: PRODUCT
              values:
                - Office 2013
                - Office 2016
      description: Patch both Windows and Microsoft apps
      operatingSystem: WINDOWS
```
{{% /example %}}
{{% example %}}
### Advanced usage, specifying alternate patch source repository

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const al201709 = new aws.ssm.PatchBaseline("al201709", {
    approvalRules: [{}],
    description: "My patch repository for Amazon Linux 2017.09",
    operatingSystem: "AMAZON_LINUX",
    sources: [{
        configuration: `[amzn-main]
name=amzn-main-Base
mirrorlist=http://repo./$awsregion./$awsdomain//$releasever/main/mirror.list
mirrorlist_expire=300
metadata_expire=300
priority=10
failovermethod=priority
fastestmirror_enabled=0
gpgcheck=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-amazon-ga
enabled=1
retries=3
timeout=5
report_instanceid=yes

`,
        name: "My-AL2017.09",
        products: ["AmazonLinux2017.09"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

al201709 = aws.ssm.PatchBaseline("al201709",
    approval_rules=[aws.ssm.PatchBaselineApprovalRuleArgs()],
    description="My patch repository for Amazon Linux 2017.09",
    operating_system="AMAZON_LINUX",
    sources=[aws.ssm.PatchBaselineSourceArgs(
        configuration="""[amzn-main]
name=amzn-main-Base
mirrorlist=http://repo./$awsregion./$awsdomain//$releasever/main/mirror.list
mirrorlist_expire=300
metadata_expire=300
priority=10
failovermethod=priority
fastestmirror_enabled=0
gpgcheck=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-amazon-ga
enabled=1
retries=3
timeout=5
report_instanceid=yes

""",
        name="My-AL2017.09",
        products=["AmazonLinux2017.09"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var al201709 = new Aws.Ssm.PatchBaseline("al201709", new()
    {
        ApprovalRules = new[]
        {
            null,
        },
        Description = "My patch repository for Amazon Linux 2017.09",
        OperatingSystem = "AMAZON_LINUX",
        Sources = new[]
        {
            new Aws.Ssm.Inputs.PatchBaselineSourceArgs
            {
                Configuration = @"[amzn-main]
name=amzn-main-Base
mirrorlist=http://repo./$awsregion./$awsdomain//$releasever/main/mirror.list
mirrorlist_expire=300
metadata_expire=300
priority=10
failovermethod=priority
fastestmirror_enabled=0
gpgcheck=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-amazon-ga
enabled=1
retries=3
timeout=5
report_instanceid=yes

",
                Name = "My-AL2017.09",
                Products = new[]
                {
                    "AmazonLinux2017.09",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewPatchBaseline(ctx, "al201709", &ssm.PatchBaselineArgs{
			ApprovalRules: ssm.PatchBaselineApprovalRuleArray{
				nil,
			},
			Description:     pulumi.String("My patch repository for Amazon Linux 2017.09"),
			OperatingSystem: pulumi.String("AMAZON_LINUX"),
			Sources: ssm.PatchBaselineSourceArray{
				&ssm.PatchBaselineSourceArgs{
					Configuration: pulumi.String(`[amzn-main]
name=amzn-main-Base
mirrorlist=http://repo./$awsregion./$awsdomain//$releasever/main/mirror.list
mirrorlist_expire=300
metadata_expire=300
priority=10
failovermethod=priority
fastestmirror_enabled=0
gpgcheck=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-amazon-ga
enabled=1
retries=3
timeout=5
report_instanceid=yes

`),
					Name: pulumi.String("My-AL2017.09"),
					Products: pulumi.StringArray{
						pulumi.String("AmazonLinux2017.09"),
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
import com.pulumi.aws.ssm.PatchBaseline;
import com.pulumi.aws.ssm.PatchBaselineArgs;
import com.pulumi.aws.ssm.inputs.PatchBaselineApprovalRuleArgs;
import com.pulumi.aws.ssm.inputs.PatchBaselineSourceArgs;
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
        var al201709 = new PatchBaseline("al201709", PatchBaselineArgs.builder()        
            .approvalRules()
            .description("My patch repository for Amazon Linux 2017.09")
            .operatingSystem("AMAZON_LINUX")
            .sources(PatchBaselineSourceArgs.builder()
                .configuration("""
[amzn-main]
name=amzn-main-Base
mirrorlist=http://repo./$awsregion./$awsdomain//$releasever/main/mirror.list
mirrorlist_expire=300
metadata_expire=300
priority=10
failovermethod=priority
fastestmirror_enabled=0
gpgcheck=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-amazon-ga
enabled=1
retries=3
timeout=5
report_instanceid=yes

                """)
                .name("My-AL2017.09")
                .products("AmazonLinux2017.09")
                .build())
            .build());

    }
}
```
```yaml
resources:
  al201709:
    type: aws:ssm:PatchBaseline
    properties:
      approvalRules:
        - {}
      description: My patch repository for Amazon Linux 2017.09
      operatingSystem: AMAZON_LINUX
      sources:
        - configuration: |+
            [amzn-main]
            name=amzn-main-Base
            mirrorlist=http://repo./$awsregion./$awsdomain//$releasever/main/mirror.list
            mirrorlist_expire=300
            metadata_expire=300
            priority=10
            failovermethod=priority
            fastestmirror_enabled=0
            gpgcheck=1
            gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-amazon-ga
            enabled=1
            retries=3
            timeout=5
            report_instanceid=yes

          name: My-AL2017.09
          products:
            - AmazonLinux2017.09
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SSM Patch Baselines using their baseline ID. For example:

```sh
 $ pulumi import aws:ssm/patchBaseline:PatchBaseline example pb-12345678
```
 