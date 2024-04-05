Subscribes to a Security Hub product.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const current = aws.getRegion({});
const exampleProductSubscription = new aws.securityhub.ProductSubscription("exampleProductSubscription", {productArn: current.then(current => `arn:aws:securityhub:${current.name}:733251395267:product/alertlogic/althreatmanagement`)}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
current = aws.get_region()
example_product_subscription = aws.securityhub.ProductSubscription("exampleProductSubscription", product_arn=f"arn:aws:securityhub:{current.name}:733251395267:product/alertlogic/althreatmanagement",
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

    var current = Aws.GetRegion.Invoke();

    var exampleProductSubscription = new Aws.SecurityHub.ProductSubscription("exampleProductSubscription", new()
    {
        ProductArn = $"arn:aws:securityhub:{current.Apply(getRegionResult => getRegionResult.Name)}:733251395267:product/alertlogic/althreatmanagement",
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
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		current, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewProductSubscription(ctx, "exampleProductSubscription", &securityhub.ProductSubscriptionArgs{
			ProductArn: pulumi.String(fmt.Sprintf("arn:aws:securityhub:%v:733251395267:product/alertlogic/althreatmanagement", current.Name)),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.securityhub.ProductSubscription;
import com.pulumi.aws.securityhub.ProductSubscriptionArgs;
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

        final var current = AwsFunctions.getRegion();

        var exampleProductSubscription = new ProductSubscription("exampleProductSubscription", ProductSubscriptionArgs.builder()        
            .productArn(String.format("arn:aws:securityhub:%s:733251395267:product/alertlogic/althreatmanagement", current.applyValue(getRegionResult -> getRegionResult.name())))
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
  exampleProductSubscription:
    type: aws:securityhub:ProductSubscription
    properties:
      productArn: arn:aws:securityhub:${current.name}:733251395267:product/alertlogic/althreatmanagement
    options:
      dependson:
        - ${exampleAccount}
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Security Hub product subscriptions using `product_arn,arn`. For example:

```sh
 $ pulumi import aws:securityhub/productSubscription:ProductSubscription example arn:aws:securityhub:eu-west-1:733251395267:product/alertlogic/althreatmanagement,arn:aws:securityhub:eu-west-1:123456789012:product-subscription/alertlogic/althreatmanagement
```
 