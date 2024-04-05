Provides a License Manager license configuration resource.

> **Note:** Removing the `license_count` attribute is not supported by the License Manager API.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.licensemanager.LicenseConfiguration("example", {
    description: "Example",
    licenseCount: 10,
    licenseCountHardLimit: true,
    licenseCountingType: "Socket",
    licenseRules: ["#minimumSockets=2"],
    tags: {
        foo: "barr",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.licensemanager.LicenseConfiguration("example",
    description="Example",
    license_count=10,
    license_count_hard_limit=True,
    license_counting_type="Socket",
    license_rules=["#minimumSockets=2"],
    tags={
        "foo": "barr",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.LicenseManager.LicenseConfiguration("example", new()
    {
        Description = "Example",
        LicenseCount = 10,
        LicenseCountHardLimit = true,
        LicenseCountingType = "Socket",
        LicenseRules = new[]
        {
            "#minimumSockets=2",
        },
        Tags = 
        {
            { "foo", "barr" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/licensemanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := licensemanager.NewLicenseConfiguration(ctx, "example", &licensemanager.LicenseConfigurationArgs{
			Description:           pulumi.String("Example"),
			LicenseCount:          pulumi.Int(10),
			LicenseCountHardLimit: pulumi.Bool(true),
			LicenseCountingType:   pulumi.String("Socket"),
			LicenseRules: pulumi.StringArray{
				pulumi.String("#minimumSockets=2"),
			},
			Tags: pulumi.StringMap{
				"foo": pulumi.String("barr"),
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
import com.pulumi.aws.licensemanager.LicenseConfiguration;
import com.pulumi.aws.licensemanager.LicenseConfigurationArgs;
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
        var example = new LicenseConfiguration("example", LicenseConfigurationArgs.builder()        
            .description("Example")
            .licenseCount(10)
            .licenseCountHardLimit(true)
            .licenseCountingType("Socket")
            .licenseRules("#minimumSockets=2")
            .tags(Map.of("foo", "barr"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:licensemanager:LicenseConfiguration
    properties:
      description: Example
      licenseCount: 10
      licenseCountHardLimit: true
      licenseCountingType: Socket
      licenseRules:
        - '#minimumSockets=2'
      tags:
        foo: barr
```
{{% /example %}}
{{% /examples %}}
## Rules

License rules should be in the format of `#RuleType=RuleValue`. Supported rule types:

* `minimumVcpus` - Resource must have minimum vCPU count in order to use the license. Default: 1
* `maximumVcpus` - Resource must have maximum vCPU count in order to use the license. Default: unbounded, limit: 10000
* `minimumCores` - Resource must have minimum core count in order to use the license. Default: 1
* `maximumCores` - Resource must have maximum core count in order to use the license. Default: unbounded, limit: 10000
* `minimumSockets` - Resource must have minimum socket count in order to use the license. Default: 1
* `maximumSockets` - Resource must have maximum socket count in order to use the license. Default: unbounded, limit: 10000
* `allowedTenancy` - Defines where the license can be used. If set, restricts license usage to selected tenancies. Specify a comma delimited list of `EC2-Default`, `EC2-DedicatedHost`, `EC2-DedicatedInstance`


## Import

Using `pulumi import`, import license configurations using the `id`. For example:

```sh
 $ pulumi import aws:licensemanager/licenseConfiguration:LicenseConfiguration example arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef
```
 