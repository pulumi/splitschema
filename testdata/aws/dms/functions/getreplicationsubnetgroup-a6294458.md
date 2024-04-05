Data source for managing an AWS DMS (Database Migration) Replication Subnet Group.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.dms.getReplicationSubnetGroup({
    replicationSubnetGroupId: aws_dms_replication_subnet_group.test.replication_subnet_group_id,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.dms.get_replication_subnet_group(replication_subnet_group_id=aws_dms_replication_subnet_group["test"]["replication_subnet_group_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Dms.GetReplicationSubnetGroup.Invoke(new()
    {
        ReplicationSubnetGroupId = aws_dms_replication_subnet_group.Test.Replication_subnet_group_id,
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
		_, err := dms.LookupReplicationSubnetGroup(ctx, &dms.LookupReplicationSubnetGroupArgs{
			ReplicationSubnetGroupId: aws_dms_replication_subnet_group.Test.Replication_subnet_group_id,
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
import com.pulumi.aws.dms.inputs.GetReplicationSubnetGroupArgs;
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
        final var test = DmsFunctions.getReplicationSubnetGroup(GetReplicationSubnetGroupArgs.builder()
            .replicationSubnetGroupId(aws_dms_replication_subnet_group.test().replication_subnet_group_id())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:dms:getReplicationSubnetGroup
      Arguments:
        replicationSubnetGroupId: ${aws_dms_replication_subnet_group.test.replication_subnet_group_id}
```
{{% /example %}}
{{% /examples %}}