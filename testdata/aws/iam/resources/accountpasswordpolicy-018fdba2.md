> **Note:** There is only a single policy allowed per AWS account. An existing policy will be lost when using this resource as an effect of this limitation.

Manages Password Policy for the AWS Account.
See more about [Account Password Policy](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html)
in the official AWS docs.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const strict = new aws.iam.AccountPasswordPolicy("strict", {
    allowUsersToChangePassword: true,
    minimumPasswordLength: 8,
    requireLowercaseCharacters: true,
    requireNumbers: true,
    requireSymbols: true,
    requireUppercaseCharacters: true,
});
```
```python
import pulumi
import pulumi_aws as aws

strict = aws.iam.AccountPasswordPolicy("strict",
    allow_users_to_change_password=True,
    minimum_password_length=8,
    require_lowercase_characters=True,
    require_numbers=True,
    require_symbols=True,
    require_uppercase_characters=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var strict = new Aws.Iam.AccountPasswordPolicy("strict", new()
    {
        AllowUsersToChangePassword = true,
        MinimumPasswordLength = 8,
        RequireLowercaseCharacters = true,
        RequireNumbers = true,
        RequireSymbols = true,
        RequireUppercaseCharacters = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewAccountPasswordPolicy(ctx, "strict", &iam.AccountPasswordPolicyArgs{
			AllowUsersToChangePassword: pulumi.Bool(true),
			MinimumPasswordLength:      pulumi.Int(8),
			RequireLowercaseCharacters: pulumi.Bool(true),
			RequireNumbers:             pulumi.Bool(true),
			RequireSymbols:             pulumi.Bool(true),
			RequireUppercaseCharacters: pulumi.Bool(true),
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
import com.pulumi.aws.iam.AccountPasswordPolicy;
import com.pulumi.aws.iam.AccountPasswordPolicyArgs;
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
        var strict = new AccountPasswordPolicy("strict", AccountPasswordPolicyArgs.builder()        
            .allowUsersToChangePassword(true)
            .minimumPasswordLength(8)
            .requireLowercaseCharacters(true)
            .requireNumbers(true)
            .requireSymbols(true)
            .requireUppercaseCharacters(true)
            .build());

    }
}
```
```yaml
resources:
  strict:
    type: aws:iam:AccountPasswordPolicy
    properties:
      allowUsersToChangePassword: true
      minimumPasswordLength: 8
      requireLowercaseCharacters: true
      requireNumbers: true
      requireSymbols: true
      requireUppercaseCharacters: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM Account Password Policy using the word `iam-account-password-policy`. For example:

```sh
 $ pulumi import aws:iam/accountPasswordPolicy:AccountPasswordPolicy strict iam-account-password-policy
```
 