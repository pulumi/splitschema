Provides an SES email identity resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ses.EmailIdentity("example", {email: "email@example.com"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ses.EmailIdentity("example", email="email@example.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ses.EmailIdentity("example", new()
    {
        Email = "email@example.com",
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
		_, err := ses.NewEmailIdentity(ctx, "example", &ses.EmailIdentityArgs{
			Email: pulumi.String("email@example.com"),
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
import com.pulumi.aws.ses.EmailIdentity;
import com.pulumi.aws.ses.EmailIdentityArgs;
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
        var example = new EmailIdentity("example", EmailIdentityArgs.builder()        
            .email("email@example.com")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ses:EmailIdentity
    properties:
      email: email@example.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SES email identities using the email address. For example:

```sh
 $ pulumi import aws:ses/emailIdentity:EmailIdentity example email@example.com
```
 