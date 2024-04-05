Generate a random password.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.secretsmanager.getRandomPassword({
    excludeNumbers: true,
    passwordLength: 50,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.secretsmanager.get_random_password(exclude_numbers=True,
    password_length=50)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.SecretsManager.GetRandomPassword.Invoke(new()
    {
        ExcludeNumbers = true,
        PasswordLength = 50,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := secretsmanager.GetRandomPassword(ctx, &secretsmanager.GetRandomPasswordArgs{
			ExcludeNumbers: pulumi.BoolRef(true),
			PasswordLength: pulumi.IntRef(50),
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
import com.pulumi.aws.secretsmanager.SecretsmanagerFunctions;
import com.pulumi.aws.secretsmanager.inputs.GetRandomPasswordArgs;
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
        final var test = SecretsmanagerFunctions.getRandomPassword(GetRandomPasswordArgs.builder()
            .excludeNumbers(true)
            .passwordLength(50)
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:secretsmanager:getRandomPassword
      Arguments:
        excludeNumbers: true
        passwordLength: 50
```
{{% /example %}}
{{% /examples %}}