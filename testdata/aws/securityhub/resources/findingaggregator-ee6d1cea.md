Manages a Security Hub finding aggregator. Security Hub needs to be enabled in a region in order for the aggregator to pull through findings.

{{% examples %}}
## Example Usage
{{% example %}}
### All Regions Usage

The following example will enable the aggregator for every region.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleFindingAggregator = new aws.securityhub.FindingAggregator("exampleFindingAggregator", {linkingMode: "ALL_REGIONS"}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_finding_aggregator = aws.securityhub.FindingAggregator("exampleFindingAggregator", linking_mode="ALL_REGIONS",
opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleFindingAggregator = new Aws.SecurityHub.FindingAggregator("exampleFindingAggregator", new()
    {
        LinkingMode = "ALL_REGIONS",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewFindingAggregator(ctx, "exampleFindingAggregator", &securityhub.FindingAggregatorArgs{
			LinkingMode: pulumi.String("ALL_REGIONS"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
		}))
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.FindingAggregator;
import com.pulumi.aws.securityhub.FindingAggregatorArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleFindingAggregator = new FindingAggregator("exampleFindingAggregator", FindingAggregatorArgs.builder()        
            .linkingMode("ALL_REGIONS")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:securityhub:Account
  exampleFindingAggregator:
    type: aws:securityhub:FindingAggregator
    properties:
      linkingMode: ALL_REGIONS
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% example %}}
### All Regions Except Specified Regions Usage

The following example will enable the aggregator for every region except those specified in `specified_regions`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleFindingAggregator = new aws.securityhub.FindingAggregator("exampleFindingAggregator", {
    linkingMode: "ALL_REGIONS_EXCEPT_SPECIFIED",
    specifiedRegions: [
        "eu-west-1",
        "eu-west-2",
    ],
}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_finding_aggregator = aws.securityhub.FindingAggregator("exampleFindingAggregator",
    linking_mode="ALL_REGIONS_EXCEPT_SPECIFIED",
    specified_regions=[
        "eu-west-1",
        "eu-west-2",
    ],
    opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleFindingAggregator = new Aws.SecurityHub.FindingAggregator("exampleFindingAggregator", new()
    {
        LinkingMode = "ALL_REGIONS_EXCEPT_SPECIFIED",
        SpecifiedRegions = new[]
        {
            "eu-west-1",
            "eu-west-2",
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewFindingAggregator(ctx, "exampleFindingAggregator", &securityhub.FindingAggregatorArgs{
			LinkingMode: pulumi.String("ALL_REGIONS_EXCEPT_SPECIFIED"),
			SpecifiedRegions: pulumi.StringArray{
				pulumi.String("eu-west-1"),
				pulumi.String("eu-west-2"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
		}))
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.FindingAggregator;
import com.pulumi.aws.securityhub.FindingAggregatorArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleFindingAggregator = new FindingAggregator("exampleFindingAggregator", FindingAggregatorArgs.builder()        
            .linkingMode("ALL_REGIONS_EXCEPT_SPECIFIED")
            .specifiedRegions(            
                "eu-west-1",
                "eu-west-2")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:securityhub:Account
  exampleFindingAggregator:
    type: aws:securityhub:FindingAggregator
    properties:
      linkingMode: ALL_REGIONS_EXCEPT_SPECIFIED
      specifiedRegions:
        - eu-west-1
        - eu-west-2
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% example %}}
### Specified Regions Usage

The following example will enable the aggregator for every region specified in `specified_regions`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleFindingAggregator = new aws.securityhub.FindingAggregator("exampleFindingAggregator", {
    linkingMode: "SPECIFIED_REGIONS",
    specifiedRegions: [
        "eu-west-1",
        "eu-west-2",
    ],
}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_finding_aggregator = aws.securityhub.FindingAggregator("exampleFindingAggregator",
    linking_mode="SPECIFIED_REGIONS",
    specified_regions=[
        "eu-west-1",
        "eu-west-2",
    ],
    opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleFindingAggregator = new Aws.SecurityHub.FindingAggregator("exampleFindingAggregator", new()
    {
        LinkingMode = "SPECIFIED_REGIONS",
        SpecifiedRegions = new[]
        {
            "eu-west-1",
            "eu-west-2",
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewFindingAggregator(ctx, "exampleFindingAggregator", &securityhub.FindingAggregatorArgs{
			LinkingMode: pulumi.String("SPECIFIED_REGIONS"),
			SpecifiedRegions: pulumi.StringArray{
				pulumi.String("eu-west-1"),
				pulumi.String("eu-west-2"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
		}))
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.FindingAggregator;
import com.pulumi.aws.securityhub.FindingAggregatorArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleFindingAggregator = new FindingAggregator("exampleFindingAggregator", FindingAggregatorArgs.builder()        
            .linkingMode("SPECIFIED_REGIONS")
            .specifiedRegions(            
                "eu-west-1",
                "eu-west-2")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:securityhub:Account
  exampleFindingAggregator:
    type: aws:securityhub:FindingAggregator
    properties:
      linkingMode: SPECIFIED_REGIONS
      specifiedRegions:
        - eu-west-1
        - eu-west-2
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an existing Security Hub finding aggregator using the `arn`. For example:

```sh
 $ pulumi import aws:securityhub/findingAggregator:FindingAggregator example arn:aws:securityhub:eu-west-1:123456789098:finding-aggregator/abcd1234-abcd-1234-1234-abcdef123456
```
 