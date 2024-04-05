Use this data source to get information on an existing backup framework.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.backup.getFramework({
    name: "my_example_backup_framework_name",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.backup.get_framework(name="my_example_backup_framework_name")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Backup.GetFramework.Invoke(new()
    {
        Name = "my_example_backup_framework_name",
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
		_, err := backup.LookupFramework(ctx, &backup.LookupFrameworkArgs{
			Name: "my_example_backup_framework_name",
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
import com.pulumi.aws.backup.inputs.GetFrameworkArgs;
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
        final var example = BackupFunctions.getFramework(GetFrameworkArgs.builder()
            .name("my_example_backup_framework_name")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:backup:getFramework
      Arguments:
        name: my_example_backup_framework_name
```
{{% /example %}}
{{% /examples %}}