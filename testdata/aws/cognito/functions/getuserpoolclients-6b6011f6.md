Use this data source to get a list of Cognito user pools clients for a Cognito IdP user pool.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = aws.cognito.getUserPoolClients({
    userPoolId: aws_cognito_user_pool.main.id,
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.cognito.get_user_pool_clients(user_pool_id=aws_cognito_user_pool["main"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = Aws.Cognito.GetUserPoolClients.Invoke(new()
    {
        UserPoolId = aws_cognito_user_pool.Main.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognito.GetUserPoolClients(ctx, &cognito.GetUserPoolClientsArgs{
			UserPoolId: aws_cognito_user_pool.Main.Id,
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
import com.pulumi.aws.cognito.CognitoFunctions;
import com.pulumi.aws.cognito.inputs.GetUserPoolClientsArgs;
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
        final var main = CognitoFunctions.getUserPoolClients(GetUserPoolClientsArgs.builder()
            .userPoolId(aws_cognito_user_pool.main().id())
            .build());

    }
}
```
```yaml
variables:
  main:
    fn::invoke:
      Function: aws:cognito:getUserPoolClients
      Arguments:
        userPoolId: ${aws_cognito_user_pool.main.id}
```
{{% /example %}}
{{% /examples %}}