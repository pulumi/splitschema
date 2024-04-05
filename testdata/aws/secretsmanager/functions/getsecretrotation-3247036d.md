Retrieve information about a Secrets Manager secret rotation. To retrieve secret metadata, see the `aws.secretsmanager.Secret` data source. To retrieve a secret value, see the `aws.secretsmanager.SecretVersion` data source.

{{% examples %}}
## Example Usage
{{% example %}}
### Retrieve Secret Rotation Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.secretsmanager.getSecretRotation({
    secretId: data.aws_secretsmanager_secret.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.secretsmanager.get_secret_rotation(secret_id=data["aws_secretsmanager_secret"]["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.SecretsManager.GetSecretRotation.Invoke(new()
    {
        SecretId = data.Aws_secretsmanager_secret.Example.Id,
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
		_, err := secretsmanager.LookupSecretRotation(ctx, &secretsmanager.LookupSecretRotationArgs{
			SecretId: data.Aws_secretsmanager_secret.Example.Id,
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
import com.pulumi.aws.secretsmanager.inputs.GetSecretRotationArgs;
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
        final var example = SecretsmanagerFunctions.getSecretRotation(GetSecretRotationArgs.builder()
            .secretId(data.aws_secretsmanager_secret().example().id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:secretsmanager:getSecretRotation
      Arguments:
        secretId: ${data.aws_secretsmanager_secret.example.id}
```
{{% /example %}}
{{% /examples %}}