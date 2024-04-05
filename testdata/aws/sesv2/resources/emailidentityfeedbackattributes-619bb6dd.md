Resource for managing an AWS SESv2 (Simple Email V2) Email Identity Feedback Attributes.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEmailIdentity = new aws.sesv2.EmailIdentity("exampleEmailIdentity", {emailIdentity: "example.com"});
const exampleEmailIdentityFeedbackAttributes = new aws.sesv2.EmailIdentityFeedbackAttributes("exampleEmailIdentityFeedbackAttributes", {
    emailIdentity: exampleEmailIdentity.emailIdentity,
    emailForwardingEnabled: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example_email_identity = aws.sesv2.EmailIdentity("exampleEmailIdentity", email_identity="example.com")
example_email_identity_feedback_attributes = aws.sesv2.EmailIdentityFeedbackAttributes("exampleEmailIdentityFeedbackAttributes",
    email_identity=example_email_identity.email_identity,
    email_forwarding_enabled=True)
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

    var exampleEmailIdentityFeedbackAttributes = new Aws.SesV2.EmailIdentityFeedbackAttributes("exampleEmailIdentityFeedbackAttributes", new()
    {
        EmailIdentity = exampleEmailIdentity.EmailIdentityDetails,
        EmailForwardingEnabled = true,
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
		exampleEmailIdentity, err := sesv2.NewEmailIdentity(ctx, "exampleEmailIdentity", &sesv2.EmailIdentityArgs{
			EmailIdentity: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		_, err = sesv2.NewEmailIdentityFeedbackAttributes(ctx, "exampleEmailIdentityFeedbackAttributes", &sesv2.EmailIdentityFeedbackAttributesArgs{
			EmailIdentity:          exampleEmailIdentity.EmailIdentity,
			EmailForwardingEnabled: pulumi.Bool(true),
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
import com.pulumi.aws.sesv2.EmailIdentityFeedbackAttributes;
import com.pulumi.aws.sesv2.EmailIdentityFeedbackAttributesArgs;
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

        var exampleEmailIdentityFeedbackAttributes = new EmailIdentityFeedbackAttributes("exampleEmailIdentityFeedbackAttributes", EmailIdentityFeedbackAttributesArgs.builder()        
            .emailIdentity(exampleEmailIdentity.emailIdentity())
            .emailForwardingEnabled(true)
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
  exampleEmailIdentityFeedbackAttributes:
    type: aws:sesv2:EmailIdentityFeedbackAttributes
    properties:
      emailIdentity: ${exampleEmailIdentity.emailIdentity}
      emailForwardingEnabled: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SESv2 (Simple Email V2) Email Identity Feedback Attributes using the `email_identity`. For example:

```sh
 $ pulumi import aws:sesv2/emailIdentityFeedbackAttributes:EmailIdentityFeedbackAttributes example example.com
```
 