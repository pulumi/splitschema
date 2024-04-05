Resource for managing an AWS FinSpace Kx Database.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {
    description: "Example KMS Key",
    deletionWindowInDays: 7,
});
const exampleKxEnvironment = new aws.finspace.KxEnvironment("exampleKxEnvironment", {kmsKeyId: exampleKey.arn});
const exampleKxDatabase = new aws.finspace.KxDatabase("exampleKxDatabase", {
    environmentId: exampleKxEnvironment.id,
    description: "Example database description",
});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey",
    description="Example KMS Key",
    deletion_window_in_days=7)
example_kx_environment = aws.finspace.KxEnvironment("exampleKxEnvironment", kms_key_id=example_key.arn)
example_kx_database = aws.finspace.KxDatabase("exampleKxDatabase",
    environment_id=example_kx_environment.id,
    description="Example database description")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "Example KMS Key",
        DeletionWindowInDays = 7,
    });

    var exampleKxEnvironment = new Aws.FinSpace.KxEnvironment("exampleKxEnvironment", new()
    {
        KmsKeyId = exampleKey.Arn,
    });

    var exampleKxDatabase = new Aws.FinSpace.KxDatabase("exampleKxDatabase", new()
    {
        EnvironmentId = exampleKxEnvironment.Id,
        Description = "Example database description",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/finspace"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description:          pulumi.String("Example KMS Key"),
			DeletionWindowInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		exampleKxEnvironment, err := finspace.NewKxEnvironment(ctx, "exampleKxEnvironment", &finspace.KxEnvironmentArgs{
			KmsKeyId: exampleKey.Arn,
		})
		if err != nil {
			return err
		}
		_, err = finspace.NewKxDatabase(ctx, "exampleKxDatabase", &finspace.KxDatabaseArgs{
			EnvironmentId: exampleKxEnvironment.ID(),
			Description:   pulumi.String("Example database description"),
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.finspace.KxEnvironment;
import com.pulumi.aws.finspace.KxEnvironmentArgs;
import com.pulumi.aws.finspace.KxDatabase;
import com.pulumi.aws.finspace.KxDatabaseArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("Example KMS Key")
            .deletionWindowInDays(7)
            .build());

        var exampleKxEnvironment = new KxEnvironment("exampleKxEnvironment", KxEnvironmentArgs.builder()        
            .kmsKeyId(exampleKey.arn())
            .build());

        var exampleKxDatabase = new KxDatabase("exampleKxDatabase", KxDatabaseArgs.builder()        
            .environmentId(exampleKxEnvironment.id())
            .description("Example database description")
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: Example KMS Key
      deletionWindowInDays: 7
  exampleKxEnvironment:
    type: aws:finspace:KxEnvironment
    properties:
      kmsKeyId: ${exampleKey.arn}
  exampleKxDatabase:
    type: aws:finspace:KxDatabase
    properties:
      environmentId: ${exampleKxEnvironment.id}
      description: Example database description
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an AWS FinSpace Kx Database using the `id` (environment ID and database name, comma-delimited). For example:

```sh
 $ pulumi import aws:finspace/kxDatabase:KxDatabase example n3ceo7wqxoxcti5tujqwzs,my-tf-kx-database
```
 