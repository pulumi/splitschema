The IAM Account Alias data source allows access to the account alias
for the effective account in which this provider is working.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.iam.getAccountAlias({});
export const accountId = current.then(current => current.accountAlias);
```
```python
import pulumi
import pulumi_aws as aws

current = aws.iam.get_account_alias()
pulumi.export("accountId", current.account_alias)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.Iam.GetAccountAlias.Invoke();

    return new Dictionary<string, object?>
    {
        ["accountId"] = current.Apply(getAccountAliasResult => getAccountAliasResult.AccountAlias),
    };
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
		current, err := iam.LookupAccountAlias(ctx, nil, nil)
		if err != nil {
			return err
		}
		ctx.Export("accountId", current.AccountAlias)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.iam.IamFunctions;
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
        final var current = IamFunctions.getAccountAlias();

        ctx.export("accountId", current.applyValue(getAccountAliasResult -> getAccountAliasResult.accountAlias()));
    }
}
```
```yaml
variables:
  current:
    fn::invoke:
      Function: aws:iam:getAccountAlias
      Arguments: {}
outputs:
  accountId: ${current.accountAlias}
```
{{% /example %}}
{{% /examples %}}