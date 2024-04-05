Provides a Cognito User Pool Client resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const client = aws.cognito.getUserPoolClient({
    clientId: "38fjsnc484p94kpqsnet7mpld0",
    userPoolId: "us-west-2_aaaaaaaaa",
});
```
```python
import pulumi
import pulumi_aws as aws

client = aws.cognito.get_user_pool_client(client_id="38fjsnc484p94kpqsnet7mpld0",
    user_pool_id="us-west-2_aaaaaaaaa")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var client = Aws.Cognito.GetUserPoolClient.Invoke(new()
    {
        ClientId = "38fjsnc484p94kpqsnet7mpld0",
        UserPoolId = "us-west-2_aaaaaaaaa",
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
		_, err := cognito.LookupUserPoolClient(ctx, &cognito.LookupUserPoolClientArgs{
			ClientId:   "38fjsnc484p94kpqsnet7mpld0",
			UserPoolId: "us-west-2_aaaaaaaaa",
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
import com.pulumi.aws.cognito.inputs.GetUserPoolClientArgs;
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
        final var client = CognitoFunctions.getUserPoolClient(GetUserPoolClientArgs.builder()
            .clientId("38fjsnc484p94kpqsnet7mpld0")
            .userPoolId("us-west-2_aaaaaaaaa")
            .build());

    }
}
```
```yaml
variables:
  client:
    fn::invoke:
      Function: aws:cognito:getUserPoolClient
      Arguments:
        clientId: 38fjsnc484p94kpqsnet7mpld0
        userPoolId: us-west-2_aaaaaaaaa
```
{{% /example %}}
{{% /examples %}}