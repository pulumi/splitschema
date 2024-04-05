Provides a resource to manage AWS Secrets Manager secret version including its secret value. To manage secret metadata, see the `aws.secretsmanager.Secret` resource.

> **NOTE:** If the `AWSCURRENT` staging label is present on this version during resource deletion, that label cannot be removed and will be skipped to prevent errors when fully deleting the secret. That label will leave this secret version active even after the resource is deleted from this provider unless the secret itself is deleted. Move the `AWSCURRENT` staging label before or after deleting this resource from this provider to fully trigger version deprecation if necessary.

{{% examples %}}
## Example Usage
{{% example %}}
### Simple String Value

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.secretsmanager.SecretVersion("example", {
    secretId: aws_secretsmanager_secret.example.id,
    secretString: "example-string-to-protect",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.secretsmanager.SecretVersion("example",
    secret_id=aws_secretsmanager_secret["example"]["id"],
    secret_string="example-string-to-protect")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SecretsManager.SecretVersion("example", new()
    {
        SecretId = aws_secretsmanager_secret.Example.Id,
        SecretString = "example-string-to-protect",
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
		_, err := secretsmanager.NewSecretVersion(ctx, "example", &secretsmanager.SecretVersionArgs{
			SecretId:     pulumi.Any(aws_secretsmanager_secret.Example.Id),
			SecretString: pulumi.String("example-string-to-protect"),
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
import com.pulumi.aws.secretsmanager.SecretVersion;
import com.pulumi.aws.secretsmanager.SecretVersionArgs;
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
        var example = new SecretVersion("example", SecretVersionArgs.builder()        
            .secretId(aws_secretsmanager_secret.example().id())
            .secretString("example-string-to-protect")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:secretsmanager:SecretVersion
    properties:
      secretId: ${aws_secretsmanager_secret.example.id}
      secretString: example-string-to-protect
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_secretsmanager_secret_version` using the secret ID and version ID. For example:

```sh
 $ pulumi import aws:secretsmanager/secretVersion:SecretVersion example 'arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456|xxxxx-xxxxxxx-xxxxxxx-xxxxx'
```
 