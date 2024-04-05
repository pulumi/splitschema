Disable/enable Security Hub standards control in the current region.

The `aws.securityhub.StandardsControl` behaves differently from normal resources, in that
Pulumi does not _create_ this resource, but instead "adopts" it
into management. When you _delete_ this resource configuration, Pulumi "abandons" resource as is and just removes it from the state.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.securityhub.Account("example", {});
const cisAwsFoundationsBenchmark = new aws.securityhub.StandardsSubscription("cisAwsFoundationsBenchmark", {standardsArn: "arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0"}, {
    dependsOn: [example],
});
const ensureIamPasswordPolicyPreventsPasswordReuse = new aws.securityhub.StandardsControl("ensureIamPasswordPolicyPreventsPasswordReuse", {
    standardsControlArn: "arn:aws:securityhub:us-east-1:111111111111:control/cis-aws-foundations-benchmark/v/1.2.0/1.10",
    controlStatus: "DISABLED",
    disabledReason: "We handle password policies within Okta",
}, {
    dependsOn: [cisAwsFoundationsBenchmark],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.securityhub.Account("example")
cis_aws_foundations_benchmark = aws.securityhub.StandardsSubscription("cisAwsFoundationsBenchmark", standards_arn="arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0",
opts=pulumi.ResourceOptions(depends_on=[example]))
ensure_iam_password_policy_prevents_password_reuse = aws.securityhub.StandardsControl("ensureIamPasswordPolicyPreventsPasswordReuse",
    standards_control_arn="arn:aws:securityhub:us-east-1:111111111111:control/cis-aws-foundations-benchmark/v/1.2.0/1.10",
    control_status="DISABLED",
    disabled_reason="We handle password policies within Okta",
    opts=pulumi.ResourceOptions(depends_on=[cis_aws_foundations_benchmark]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SecurityHub.Account("example");

    var cisAwsFoundationsBenchmark = new Aws.SecurityHub.StandardsSubscription("cisAwsFoundationsBenchmark", new()
    {
        StandardsArn = "arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            example,
        },
    });

    var ensureIamPasswordPolicyPreventsPasswordReuse = new Aws.SecurityHub.StandardsControl("ensureIamPasswordPolicyPreventsPasswordReuse", new()
    {
        StandardsControlArn = "arn:aws:securityhub:us-east-1:111111111111:control/cis-aws-foundations-benchmark/v/1.2.0/1.10",
        ControlStatus = "DISABLED",
        DisabledReason = "We handle password policies within Okta",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            cisAwsFoundationsBenchmark,
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
		example, err := securityhub.NewAccount(ctx, "example", nil)
		if err != nil {
			return err
		}
		cisAwsFoundationsBenchmark, err := securityhub.NewStandardsSubscription(ctx, "cisAwsFoundationsBenchmark", &securityhub.StandardsSubscriptionArgs{
			StandardsArn: pulumi.String("arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0"),
		}, pulumi.DependsOn([]pulumi.Resource{
			example,
		}))
		if err != nil {
			return err
		}
		_, err = securityhub.NewStandardsControl(ctx, "ensureIamPasswordPolicyPreventsPasswordReuse", &securityhub.StandardsControlArgs{
			StandardsControlArn: pulumi.String("arn:aws:securityhub:us-east-1:111111111111:control/cis-aws-foundations-benchmark/v/1.2.0/1.10"),
			ControlStatus:       pulumi.String("DISABLED"),
			DisabledReason:      pulumi.String("We handle password policies within Okta"),
		}, pulumi.DependsOn([]pulumi.Resource{
			cisAwsFoundationsBenchmark,
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
import com.pulumi.aws.securityhub.StandardsSubscription;
import com.pulumi.aws.securityhub.StandardsSubscriptionArgs;
import com.pulumi.aws.securityhub.StandardsControl;
import com.pulumi.aws.securityhub.StandardsControlArgs;
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

        var cisAwsFoundationsBenchmark = new StandardsSubscription("cisAwsFoundationsBenchmark", StandardsSubscriptionArgs.builder()        
            .standardsArn("arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0")
            .build(), CustomResourceOptions.builder()
                .dependsOn(example)
                .build());

        var ensureIamPasswordPolicyPreventsPasswordReuse = new StandardsControl("ensureIamPasswordPolicyPreventsPasswordReuse", StandardsControlArgs.builder()        
            .standardsControlArn("arn:aws:securityhub:us-east-1:111111111111:control/cis-aws-foundations-benchmark/v/1.2.0/1.10")
            .controlStatus("DISABLED")
            .disabledReason("We handle password policies within Okta")
            .build(), CustomResourceOptions.builder()
                .dependsOn(cisAwsFoundationsBenchmark)
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:securityhub:Account
  cisAwsFoundationsBenchmark:
    type: aws:securityhub:StandardsSubscription
    properties:
      standardsArn: arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0
    options:
      dependson:
        - ${example}
  ensureIamPasswordPolicyPreventsPasswordReuse:
    type: aws:securityhub:StandardsControl
    properties:
      standardsControlArn: arn:aws:securityhub:us-east-1:111111111111:control/cis-aws-foundations-benchmark/v/1.2.0/1.10
      controlStatus: DISABLED
      disabledReason: We handle password policies within Okta
    options:
      dependson:
        - ${cisAwsFoundationsBenchmark}
```
{{% /example %}}
{{% /examples %}}