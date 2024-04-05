Data source for managing an AWS DMS (Database Migration) Replication Instance.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.dms.getReplicationInstance({
    replicationInstanceId: aws_dms_replication_instance.test.replication_instance_id,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.dms.get_replication_instance(replication_instance_id=aws_dms_replication_instance["test"]["replication_instance_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Dms.GetReplicationInstance.Invoke(new()
    {
        ReplicationInstanceId = aws_dms_replication_instance.Test.Replication_instance_id,
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
		_, err := dms.LookupReplicationInstance(ctx, &dms.LookupReplicationInstanceArgs{
			ReplicationInstanceId: aws_dms_replication_instance.Test.Replication_instance_id,
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
import com.pulumi.aws.dms.inputs.GetReplicationInstanceArgs;
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
        final var test = DmsFunctions.getReplicationInstance(GetReplicationInstanceArgs.builder()
            .replicationInstanceId(aws_dms_replication_instance.test().replication_instance_id())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:dms:getReplicationInstance
      Arguments:
        replicationInstanceId: ${aws_dms_replication_instance.test.replication_instance_id}
```
{{% /example %}}
{{% /examples %}}