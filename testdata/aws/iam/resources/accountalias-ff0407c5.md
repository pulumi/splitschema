> **Note:** There is only a single account alias per AWS account.

Manages the account alias for the AWS Account.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const alias = new aws.iam.AccountAlias("alias", {accountAlias: "my-account-alias"});
```
```python
import pulumi
import pulumi_aws as aws

alias = aws.iam.AccountAlias("alias", account_alias="my-account-alias")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @alias = new Aws.Iam.AccountAlias("alias", new()
    {
        Alias = "my-account-alias",
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
		_, err := iam.NewAccountAlias(ctx, "alias", &iam.AccountAliasArgs{
			AccountAlias: pulumi.String("my-account-alias"),
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
import com.pulumi.aws.iam.AccountAlias;
import com.pulumi.aws.iam.AccountAliasArgs;
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
        var alias = new AccountAlias("alias", AccountAliasArgs.builder()        
            .accountAlias("my-account-alias")
            .build());

    }
}
```
```yaml
resources:
  alias:
    type: aws:iam:AccountAlias
    properties:
      accountAlias: my-account-alias
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the current Account Alias using the `account_alias`. For example:

```sh
 $ pulumi import aws:iam/accountAlias:AccountAlias alias my-account-alias
```
 