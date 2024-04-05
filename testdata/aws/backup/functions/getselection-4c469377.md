Use this data source to get information on an existing backup selection.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.backup.getSelection({
    planId: data.aws_backup_plan.example.id,
    selectionId: "selection-id-example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.backup.get_selection(plan_id=data["aws_backup_plan"]["example"]["id"],
    selection_id="selection-id-example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Backup.GetSelection.Invoke(new()
    {
        PlanId = data.Aws_backup_plan.Example.Id,
        SelectionId = "selection-id-example",
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
		_, err := backup.LookupSelection(ctx, &backup.LookupSelectionArgs{
			PlanId:      data.Aws_backup_plan.Example.Id,
			SelectionId: "selection-id-example",
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
import com.pulumi.aws.backup.inputs.GetSelectionArgs;
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
        final var example = BackupFunctions.getSelection(GetSelectionArgs.builder()
            .planId(data.aws_backup_plan().example().id())
            .selectionId("selection-id-example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:backup:getSelection
      Arguments:
        planId: ${data.aws_backup_plan.example.id}
        selectionId: selection-id-example
```
{{% /example %}}
{{% /examples %}}