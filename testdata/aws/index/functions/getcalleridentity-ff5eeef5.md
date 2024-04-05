Use this data source to get the access to the effective Account ID, User ID, and ARN in
which this provider is authorized.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
export const accountId = current.then(current => current.accountId);
export const callerArn = current.then(current => current.arn);
export const callerUser = current.then(current => current.userId);
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
pulumi.export("accountId", current.account_id)
pulumi.export("callerArn", current.arn)
pulumi.export("callerUser", current.user_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    return new Dictionary<string, object?>
    {
        ["accountId"] = current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
        ["callerArn"] = current.Apply(getCallerIdentityResult => getCallerIdentityResult.Arn),
        ["callerUser"] = current.Apply(getCallerIdentityResult => getCallerIdentityResult.UserId),
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		ctx.Export("accountId", current.AccountId)
		ctx.Export("callerArn", current.Arn)
		ctx.Export("callerUser", current.UserId)
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

        ctx.export("accountId", current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()));
        ctx.export("callerArn", current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.arn()));
        ctx.export("callerUser", current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.userId()));
    }
}
```
```yaml
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
outputs:
  accountId: ${current.accountId}
  callerArn: ${current.arn}
  callerUser: ${current.userId}
```
{{% /example %}}
{{% /examples %}}