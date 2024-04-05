Data source for managing an AWS SESv2 (Simple Email V2) Email Identity Mail From Attributes.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEmailIdentity = aws.sesv2.getEmailIdentity({
    emailIdentity: "example.com",
});
const exampleEmailIdentityMailFromAttributes = exampleEmailIdentity.then(exampleEmailIdentity => aws.sesv2.getEmailIdentityMailFromAttributes({
    emailIdentity: exampleEmailIdentity.emailIdentity,
}));
```
```python
import pulumi
import pulumi_aws as aws

example_email_identity = aws.sesv2.get_email_identity(email_identity="example.com")
example_email_identity_mail_from_attributes = aws.sesv2.get_email_identity_mail_from_attributes(email_identity=example_email_identity.email_identity)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleEmailIdentity = Aws.SesV2.GetEmailIdentity.Invoke(new()
    {
        EmailIdentity = "example.com",
    });

    var exampleEmailIdentityMailFromAttributes = Aws.SesV2.GetEmailIdentityMailFromAttributes.Invoke(new()
    {
        EmailIdentity = exampleEmailIdentity.Apply(getEmailIdentityResult => getEmailIdentityResult.EmailIdentity),
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
		exampleEmailIdentity, err := sesv2.LookupEmailIdentity(ctx, &sesv2.LookupEmailIdentityArgs{
			EmailIdentity: "example.com",
		}, nil)
		if err != nil {
			return err
		}
		_, err = sesv2.LookupEmailIdentityMailFromAttributes(ctx, &sesv2.LookupEmailIdentityMailFromAttributesArgs{
			EmailIdentity: exampleEmailIdentity.EmailIdentity,
		}, nil)
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
import com.pulumi.aws.sesv2.Sesv2Functions;
import com.pulumi.aws.sesv2.inputs.GetEmailIdentityArgs;
import com.pulumi.aws.sesv2.inputs.GetEmailIdentityMailFromAttributesArgs;
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
        final var exampleEmailIdentity = Sesv2Functions.getEmailIdentity(GetEmailIdentityArgs.builder()
            .emailIdentity("example.com")
            .build());

        final var exampleEmailIdentityMailFromAttributes = Sesv2Functions.getEmailIdentityMailFromAttributes(GetEmailIdentityMailFromAttributesArgs.builder()
            .emailIdentity(exampleEmailIdentity.applyValue(getEmailIdentityResult -> getEmailIdentityResult.emailIdentity()))
            .build());

    }
}
```
```yaml
variables:
  exampleEmailIdentity:
    fn::invoke:
      Function: aws:sesv2:getEmailIdentity
      Arguments:
        emailIdentity: example.com
  exampleEmailIdentityMailFromAttributes:
    fn::invoke:
      Function: aws:sesv2:getEmailIdentityMailFromAttributes
      Arguments:
        emailIdentity: ${exampleEmailIdentity.emailIdentity}
```
{{% /example %}}
{{% /examples %}}