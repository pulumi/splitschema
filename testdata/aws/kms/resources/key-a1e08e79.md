Manages a single-Region or multi-Region primary KMS key.

> **NOTE on KMS Key Policy:** KMS Key Policy can be configured in either the standalone resource `aws.kms.KeyPolicy`
or with the parameter `policy` in this resource.
Configuring with both will cause inconsistencies and may overwrite configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const key = new aws.kms.Key("key", {
    deletionWindowInDays: 10,
    description: "KMS key 1",
});
```
```python
import pulumi
import pulumi_aws as aws

key = aws.kms.Key("key",
    deletion_window_in_days=10,
    description="KMS key 1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var key = new Aws.Kms.Key("key", new()
    {
        DeletionWindowInDays = 10,
        Description = "KMS key 1",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kms.NewKey(ctx, "key", &kms.KeyArgs{
			DeletionWindowInDays: pulumi.Int(10),
			Description:          pulumi.String("KMS key 1"),
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
        var key = new Key("key", KeyArgs.builder()        
            .deletionWindowInDays(10)
            .description("KMS key 1")
            .build());

    }
}
```
```yaml
resources:
  key:
    type: aws:kms:Key
    properties:
      deletionWindowInDays: 10
      description: KMS key 1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import KMS Keys using the `id`. For example:

```sh
 $ pulumi import aws:kms/key:Key a 1234abcd-12ab-34cd-56ef-1234567890ab
```
 