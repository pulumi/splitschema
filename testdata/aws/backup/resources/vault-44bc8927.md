Provides an AWS Backup vault resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.Vault("example", {kmsKeyArn: aws_kms_key.example.arn});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.backup.Vault("example", kms_key_arn=aws_kms_key["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Backup.Vault("example", new()
    {
        KmsKeyArn = aws_kms_key.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewVault(ctx, "example", &backup.VaultArgs{
			KmsKeyArn: pulumi.Any(aws_kms_key.Example.Arn),
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
import com.pulumi.aws.backup.Vault;
import com.pulumi.aws.backup.VaultArgs;
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
        var example = new Vault("example", VaultArgs.builder()        
            .kmsKeyArn(aws_kms_key.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:backup:Vault
    properties:
      kmsKeyArn: ${aws_kms_key.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Backup vault using the `name`. For example:

```sh
 $ pulumi import aws:backup/vault:Vault test-vault TestVault
```
 