Provides a resource to manage AWS Secrets Manager secret rotation. To manage a secret, see the `aws.secretsmanager.Secret` resource. To manage a secret value, see the `aws.secretsmanager.SecretVersion` resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.secretsmanager.SecretRotation("example", {
    secretId: aws_secretsmanager_secret.example.id,
    rotationLambdaArn: aws_lambda_function.example.arn,
    rotationRules: {
        automaticallyAfterDays: 30,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.secretsmanager.SecretRotation("example",
    secret_id=aws_secretsmanager_secret["example"]["id"],
    rotation_lambda_arn=aws_lambda_function["example"]["arn"],
    rotation_rules=aws.secretsmanager.SecretRotationRotationRulesArgs(
        automatically_after_days=30,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SecretsManager.SecretRotation("example", new()
    {
        SecretId = aws_secretsmanager_secret.Example.Id,
        RotationLambdaArn = aws_lambda_function.Example.Arn,
        RotationRules = new Aws.SecretsManager.Inputs.SecretRotationRotationRulesArgs
        {
            AutomaticallyAfterDays = 30,
        },
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
		_, err := secretsmanager.NewSecretRotation(ctx, "example", &secretsmanager.SecretRotationArgs{
			SecretId:          pulumi.Any(aws_secretsmanager_secret.Example.Id),
			RotationLambdaArn: pulumi.Any(aws_lambda_function.Example.Arn),
			RotationRules: &secretsmanager.SecretRotationRotationRulesArgs{
				AutomaticallyAfterDays: pulumi.Int(30),
			},
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
import com.pulumi.aws.secretsmanager.SecretRotation;
import com.pulumi.aws.secretsmanager.SecretRotationArgs;
import com.pulumi.aws.secretsmanager.inputs.SecretRotationRotationRulesArgs;
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
        var example = new SecretRotation("example", SecretRotationArgs.builder()        
            .secretId(aws_secretsmanager_secret.example().id())
            .rotationLambdaArn(aws_lambda_function.example().arn())
            .rotationRules(SecretRotationRotationRulesArgs.builder()
                .automaticallyAfterDays(30)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:secretsmanager:SecretRotation
    properties:
      secretId: ${aws_secretsmanager_secret.example.id}
      rotationLambdaArn: ${aws_lambda_function.example.arn}
      rotationRules:
        automaticallyAfterDays: 30
```
{{% /example %}}
### Rotation Configuration

To enable automatic secret rotation, the Secrets Manager service requires usage of a Lambda function. The [Rotate Secrets section in the Secrets Manager User Guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html) provides additional information about deploying a prebuilt Lambda functions for supported credential rotation (e.g., RDS) or deploying a custom Lambda function.

> **NOTE:** Configuring rotation causes the secret to rotate once as soon as you enable rotation. Before you do this, you must ensure that all of your applications that use the credentials stored in the secret are updated to retrieve the secret from AWS Secrets Manager. The old credentials might no longer be usable after the initial rotation and any applications that you fail to update will break as soon as the old credentials are no longer valid.

> **NOTE:** If you cancel a rotation that is in progress (by removing the `rotation` configuration), it can leave the VersionStage labels in an unexpected state. Depending on what step of the rotation was in progress, you might need to remove the staging label AWSPENDING from the partially created version, specified by the SecretVersionId response value. You should also evaluate the partially rotated new version to see if it should be deleted, which you can do by removing all staging labels from the new version's VersionStage field.
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_secretsmanager_secret_rotation` using the secret Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:secretsmanager/secretRotation:SecretRotation example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
```
 