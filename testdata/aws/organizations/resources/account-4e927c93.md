Provides a resource to create a member account in the current organization.

> **Note:** Account management must be done from the organization's root account.

> **Note:** By default, deleting this resource will only remove an AWS account from an organization. You must set the `close_on_deletion` flag to true to close the account. It is worth noting that quotas are enforced when using the `close_on_deletion` argument, which can produce a [CLOSE_ACCOUNT_QUOTA_EXCEEDED](https://docs.aws.amazon.com/organizations/latest/APIReference/API_CloseAccount.html) error, and require you to close the account manually.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const account = new aws.organizations.Account("account", {email: "john@doe.org"});
```
```python
import pulumi
import pulumi_aws as aws

account = aws.organizations.Account("account", email="john@doe.org")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var account = new Aws.Organizations.Account("account", new()
    {
        Email = "john@doe.org",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewAccount(ctx, "account", &organizations.AccountArgs{
			Email: pulumi.String("john@doe.org"),
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
import com.pulumi.aws.organizations.Account;
import com.pulumi.aws.organizations.AccountArgs;
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
        var account = new Account("account", AccountArgs.builder()        
            .email("john@doe.org")
            .build());

    }
}
```
```yaml
resources:
  account:
    type: aws:organizations:Account
    properties:
      email: john@doe.org
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the AWS member account using the `account_id`. For example:

```sh
 $ pulumi import aws:organizations/account:Account my_account 111111111111
```
 Certain resource arguments, like `role_name`, do not have an Organizations API method for reading the information after account creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

