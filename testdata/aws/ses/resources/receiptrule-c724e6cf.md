Provides an SES receipt rule resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Add a header to the email and store it in S3
const store = new aws.ses.ReceiptRule("store", {
    addHeaderActions: [{
        headerName: "Custom-Header",
        headerValue: "Added by SES",
        position: 1,
    }],
    enabled: true,
    recipients: ["karen@example.com"],
    ruleSetName: "default-rule-set",
    s3Actions: [{
        bucketName: "emails",
        position: 2,
    }],
    scanEnabled: true,
});
```
```python
import pulumi
import pulumi_aws as aws

# Add a header to the email and store it in S3
store = aws.ses.ReceiptRule("store",
    add_header_actions=[aws.ses.ReceiptRuleAddHeaderActionArgs(
        header_name="Custom-Header",
        header_value="Added by SES",
        position=1,
    )],
    enabled=True,
    recipients=["karen@example.com"],
    rule_set_name="default-rule-set",
    s3_actions=[aws.ses.ReceiptRuleS3ActionArgs(
        bucket_name="emails",
        position=2,
    )],
    scan_enabled=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Add a header to the email and store it in S3
    var store = new Aws.Ses.ReceiptRule("store", new()
    {
        AddHeaderActions = new[]
        {
            new Aws.Ses.Inputs.ReceiptRuleAddHeaderActionArgs
            {
                HeaderName = "Custom-Header",
                HeaderValue = "Added by SES",
                Position = 1,
            },
        },
        Enabled = true,
        Recipients = new[]
        {
            "karen@example.com",
        },
        RuleSetName = "default-rule-set",
        S3Actions = new[]
        {
            new Aws.Ses.Inputs.ReceiptRuleS3ActionArgs
            {
                BucketName = "emails",
                Position = 2,
            },
        },
        ScanEnabled = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewReceiptRule(ctx, "store", &ses.ReceiptRuleArgs{
			AddHeaderActions: ses.ReceiptRuleAddHeaderActionArray{
				&ses.ReceiptRuleAddHeaderActionArgs{
					HeaderName:  pulumi.String("Custom-Header"),
					HeaderValue: pulumi.String("Added by SES"),
					Position:    pulumi.Int(1),
				},
			},
			Enabled: pulumi.Bool(true),
			Recipients: pulumi.StringArray{
				pulumi.String("karen@example.com"),
			},
			RuleSetName: pulumi.String("default-rule-set"),
			S3Actions: ses.ReceiptRuleS3ActionArray{
				&ses.ReceiptRuleS3ActionArgs{
					BucketName: pulumi.String("emails"),
					Position:   pulumi.Int(2),
				},
			},
			ScanEnabled: pulumi.Bool(true),
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
import com.pulumi.aws.ses.ReceiptRule;
import com.pulumi.aws.ses.ReceiptRuleArgs;
import com.pulumi.aws.ses.inputs.ReceiptRuleAddHeaderActionArgs;
import com.pulumi.aws.ses.inputs.ReceiptRuleS3ActionArgs;
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
        var store = new ReceiptRule("store", ReceiptRuleArgs.builder()        
            .addHeaderActions(ReceiptRuleAddHeaderActionArgs.builder()
                .headerName("Custom-Header")
                .headerValue("Added by SES")
                .position(1)
                .build())
            .enabled(true)
            .recipients("karen@example.com")
            .ruleSetName("default-rule-set")
            .s3Actions(ReceiptRuleS3ActionArgs.builder()
                .bucketName("emails")
                .position(2)
                .build())
            .scanEnabled(true)
            .build());

    }
}
```
```yaml
resources:
  # Add a header to the email and store it in S3
  store:
    type: aws:ses:ReceiptRule
    properties:
      addHeaderActions:
        - headerName: Custom-Header
          headerValue: Added by SES
          position: 1
      enabled: true
      recipients:
        - karen@example.com
      ruleSetName: default-rule-set
      s3Actions:
        - bucketName: emails
          position: 2
      scanEnabled: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SES receipt rules using the ruleset name and rule name separated by `:`. For example:

```sh
 $ pulumi import aws:ses/receiptRule:ReceiptRule my_rule my_rule_set:my_rule
```
 