Subscribes to a Security Hub standard.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.securityhub.Account("example", {});
const current = aws.getRegion({});
const cis = new aws.securityhub.StandardsSubscription("cis", {standardsArn: "arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0"}, {
    dependsOn: [example],
});
const pci321 = new aws.securityhub.StandardsSubscription("pci321", {standardsArn: current.then(current => `arn:aws:securityhub:${current.name}::standards/pci-dss/v/3.2.1`)}, {
    dependsOn: [example],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.securityhub.Account("example")
current = aws.get_region()
cis = aws.securityhub.StandardsSubscription("cis", standards_arn="arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0",
opts=pulumi.ResourceOptions(depends_on=[example]))
pci321 = aws.securityhub.StandardsSubscription("pci321", standards_arn=f"arn:aws:securityhub:{current.name}::standards/pci-dss/v/3.2.1",
opts=pulumi.ResourceOptions(depends_on=[example]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SecurityHub.Account("example");

    var current = Aws.GetRegion.Invoke();

    var cis = new Aws.SecurityHub.StandardsSubscription("cis", new()
    {
        StandardsArn = "arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            example,
        },
    });

    var pci321 = new Aws.SecurityHub.StandardsSubscription("pci321", new()
    {
        StandardsArn = $"arn:aws:securityhub:{current.Apply(getRegionResult => getRegionResult.Name)}::standards/pci-dss/v/3.2.1",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            example,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := securityhub.NewAccount(ctx, "example", nil)
		if err != nil {
			return err
		}
		current, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewStandardsSubscription(ctx, "cis", &securityhub.StandardsSubscriptionArgs{
			StandardsArn: pulumi.String("arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0"),
		}, pulumi.DependsOn([]pulumi.Resource{
			example,
		}))
		if err != nil {
			return err
		}
		_, err = securityhub.NewStandardsSubscription(ctx, "pci321", &securityhub.StandardsSubscriptionArgs{
			StandardsArn: pulumi.String(fmt.Sprintf("arn:aws:securityhub:%v::standards/pci-dss/v/3.2.1", current.Name)),
		}, pulumi.DependsOn([]pulumi.Resource{
			example,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.securityhub.StandardsSubscription;
import com.pulumi.aws.securityhub.StandardsSubscriptionArgs;
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
        var example = new Account("example");

        final var current = AwsFunctions.getRegion();

        var cis = new StandardsSubscription("cis", StandardsSubscriptionArgs.builder()        
            .standardsArn("arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0")
            .build(), CustomResourceOptions.builder()
                .dependsOn(example)
                .build());

        var pci321 = new StandardsSubscription("pci321", StandardsSubscriptionArgs.builder()        
            .standardsArn(String.format("arn:aws:securityhub:%s::standards/pci-dss/v/3.2.1", current.applyValue(getRegionResult -> getRegionResult.name())))
            .build(), CustomResourceOptions.builder()
                .dependsOn(example)
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:securityhub:Account
  cis:
    type: aws:securityhub:StandardsSubscription
    properties:
      standardsArn: arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0
    options:
      dependson:
        - ${example}
  pci321:
    type: aws:securityhub:StandardsSubscription
    properties:
      standardsArn: arn:aws:securityhub:${current.name}::standards/pci-dss/v/3.2.1
    options:
      dependson:
        - ${example}
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Security Hub standards subscriptions using the standards subscription ARN. For example:

```sh
 $ pulumi import aws:securityhub/standardsSubscription:StandardsSubscription cis arn:aws:securityhub:eu-west-1:123456789012:subscription/cis-aws-foundations-benchmark/v/1.2.0
```
 ```sh
 $ pulumi import aws:securityhub/standardsSubscription:StandardsSubscription pci_321 arn:aws:securityhub:eu-west-1:123456789012:subscription/pci-dss/v/3.2.1
```
 ```sh
 $ pulumi import aws:securityhub/standardsSubscription:StandardsSubscription nist_800_53_rev_5 arn:aws:securityhub:eu-west-1:123456789012:subscription/nist-800-53/v/5.0.0
```
 