Resource for managing an Amazon Inspector Delegated Admin Account.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const example = new aws.inspector2.DelegatedAdminAccount("example", {accountId: current.then(current => current.accountId)});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
example = aws.inspector2.DelegatedAdminAccount("example", account_id=current.account_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var example = new Aws.Inspector2.DelegatedAdminAccount("example", new()
    {
        AccountId = current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/inspector2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = inspector2.NewDelegatedAdminAccount(ctx, "example", &inspector2.DelegatedAdminAccountArgs{
			AccountId: *pulumi.String(current.AccountId),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inspector2.DelegatedAdminAccount;
import com.pulumi.aws.inspector2.DelegatedAdminAccountArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        var example = new DelegatedAdminAccount("example", DelegatedAdminAccountArgs.builder()        
            .accountId(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:inspector2:DelegatedAdminAccount
    properties:
      accountId: ${current.accountId}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Inspector Delegated Admin Account using the `account_id`. For example:

```sh
 $ pulumi import aws:inspector2/delegatedAdminAccount:DelegatedAdminAccount example 012345678901
```
 