Resource for managing an AWS SESv2 (Simple Email V2) Email Identity Mail From Attributes.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEmailIdentity = new aws.sesv2.EmailIdentity("exampleEmailIdentity", {emailIdentity: "example.com"});
const exampleEmailIdentityMailFromAttributes = new aws.sesv2.EmailIdentityMailFromAttributes("exampleEmailIdentityMailFromAttributes", {
    emailIdentity: exampleEmailIdentity.emailIdentity,
    behaviorOnMxFailure: "REJECT_MESSAGE",
    mailFromDomain: pulumi.interpolate`subdomain.${exampleEmailIdentity.emailIdentity}`,
});
```
```python
import pulumi
import pulumi_aws as aws

example_email_identity = aws.sesv2.EmailIdentity("exampleEmailIdentity", email_identity="example.com")
example_email_identity_mail_from_attributes = aws.sesv2.EmailIdentityMailFromAttributes("exampleEmailIdentityMailFromAttributes",
    email_identity=example_email_identity.email_identity,
    behavior_on_mx_failure="REJECT_MESSAGE",
    mail_from_domain=example_email_identity.email_identity.apply(lambda email_identity: f"subdomain.{email_identity}"))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleEmailIdentity = new Aws.SesV2.EmailIdentity("exampleEmailIdentity", new()
    {
        EmailIdentityDetails = "example.com",
    });

    var exampleEmailIdentityMailFromAttributes = new Aws.SesV2.EmailIdentityMailFromAttributes("exampleEmailIdentityMailFromAttributes", new()
    {
        EmailIdentity = exampleEmailIdentity.EmailIdentityDetails,
        BehaviorOnMxFailure = "REJECT_MESSAGE",
        MailFromDomain = exampleEmailIdentity.EmailIdentityDetails.Apply(emailIdentity => $"subdomain.{emailIdentity}"),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleEmailIdentity, err := sesv2.NewEmailIdentity(ctx, "exampleEmailIdentity", &sesv2.EmailIdentityArgs{
			EmailIdentity: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		_, err = sesv2.NewEmailIdentityMailFromAttributes(ctx, "exampleEmailIdentityMailFromAttributes", &sesv2.EmailIdentityMailFromAttributesArgs{
			EmailIdentity:       exampleEmailIdentity.EmailIdentity,
			BehaviorOnMxFailure: pulumi.String("REJECT_MESSAGE"),
			MailFromDomain: exampleEmailIdentity.EmailIdentity.ApplyT(func(emailIdentity string) (string, error) {
				return fmt.Sprintf("subdomain.%v", emailIdentity), nil
			}).(pulumi.StringOutput),
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
import com.pulumi.aws.sesv2.EmailIdentity;
import com.pulumi.aws.sesv2.EmailIdentityArgs;
import com.pulumi.aws.sesv2.EmailIdentityMailFromAttributes;
import com.pulumi.aws.sesv2.EmailIdentityMailFromAttributesArgs;
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
        var exampleEmailIdentity = new EmailIdentity("exampleEmailIdentity", EmailIdentityArgs.builder()        
            .emailIdentity("example.com")
            .build());

        var exampleEmailIdentityMailFromAttributes = new EmailIdentityMailFromAttributes("exampleEmailIdentityMailFromAttributes", EmailIdentityMailFromAttributesArgs.builder()        
            .emailIdentity(exampleEmailIdentity.emailIdentity())
            .behaviorOnMxFailure("REJECT_MESSAGE")
            .mailFromDomain(exampleEmailIdentity.emailIdentity().applyValue(emailIdentity -> String.format("subdomain.%s", emailIdentity)))
            .build());

    }
}
```
```yaml
resources:
  exampleEmailIdentity:
    type: aws:sesv2:EmailIdentity
    properties:
      emailIdentity: example.com
  exampleEmailIdentityMailFromAttributes:
    type: aws:sesv2:EmailIdentityMailFromAttributes
    properties:
      emailIdentity: ${exampleEmailIdentity.emailIdentity}
      behaviorOnMxFailure: REJECT_MESSAGE
      mailFromDomain: subdomain.${exampleEmailIdentity.emailIdentity}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SESv2 (Simple Email V2) Email Identity Mail From Attributes using the `email_identity`. For example:

```sh
 $ pulumi import aws:sesv2/emailIdentityMailFromAttributes:EmailIdentityMailFromAttributes example example.com
```
 