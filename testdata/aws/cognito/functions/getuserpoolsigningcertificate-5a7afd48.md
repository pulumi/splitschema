Use this data source to get the signing certificate for a Cognito IdP user pool.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sc = aws.cognito.getUserPoolSigningCertificate({
    userPoolId: aws_cognito_user_pool.my_pool.id,
});
```
```python
import pulumi
import pulumi_aws as aws

sc = aws.cognito.get_user_pool_signing_certificate(user_pool_id=aws_cognito_user_pool["my_pool"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var sc = Aws.Cognito.GetUserPoolSigningCertificate.Invoke(new()
    {
        UserPoolId = aws_cognito_user_pool.My_pool.Id,
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
		_, err := cognito.GetUserPoolSigningCertificate(ctx, &cognito.GetUserPoolSigningCertificateArgs{
			UserPoolId: aws_cognito_user_pool.My_pool.Id,
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
import com.pulumi.aws.cognito.inputs.GetUserPoolSigningCertificateArgs;
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
        final var sc = CognitoFunctions.getUserPoolSigningCertificate(GetUserPoolSigningCertificateArgs.builder()
            .userPoolId(aws_cognito_user_pool.my_pool().id())
            .build());

    }
}
```
```yaml
variables:
  sc:
    fn::invoke:
      Function: aws:cognito:getUserPoolSigningCertificate
      Arguments:
        userPoolId: ${aws_cognito_user_pool.my_pool.id}
```
{{% /example %}}
{{% /examples %}}