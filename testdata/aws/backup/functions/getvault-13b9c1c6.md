Use this data source to get information on an existing backup vault.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.backup.getVault({
    name: "example_backup_vault",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.backup.get_vault(name="example_backup_vault")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Backup.GetVault.Invoke(new()
    {
        Name = "example_backup_vault",
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
		_, err := backup.LookupVault(ctx, &backup.LookupVaultArgs{
			Name: "example_backup_vault",
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
import com.pulumi.aws.backup.BackupFunctions;
import com.pulumi.aws.backup.inputs.GetVaultArgs;
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
        final var example = BackupFunctions.getVault(GetVaultArgs.builder()
            .name("example_backup_vault")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:backup:getVault
      Arguments:
        name: example_backup_vault
```
{{% /example %}}
{{% /examples %}}