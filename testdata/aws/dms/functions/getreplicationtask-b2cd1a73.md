Data source for managing an AWS DMS (Database Migration) Replication Task.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.dms.getReplicationTask({
    replicationTaskId: aws_dms_replication_task.test.replication_task_id,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.dms.get_replication_task(replication_task_id=aws_dms_replication_task["test"]["replication_task_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Dms.GetReplicationTask.Invoke(new()
    {
        ReplicationTaskId = aws_dms_replication_task.Test.Replication_task_id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.LookupReplicationTask(ctx, &dms.LookupReplicationTaskArgs{
			ReplicationTaskId: aws_dms_replication_task.Test.Replication_task_id,
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
import com.pulumi.aws.dms.DmsFunctions;
import com.pulumi.aws.dms.inputs.GetReplicationTaskArgs;
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
        final var test = DmsFunctions.getReplicationTask(GetReplicationTaskArgs.builder()
            .replicationTaskId(aws_dms_replication_task.test().replication_task_id())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:dms:getReplicationTask
      Arguments:
        replicationTaskId: ${aws_dms_replication_task.test.replication_task_id}
```
{{% /example %}}
{{% /examples %}}