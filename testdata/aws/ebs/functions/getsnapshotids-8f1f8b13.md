Use this data source to get a list of EBS Snapshot IDs matching the specified
criteria.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ebsVolumes = aws.ebs.getSnapshotIds({
    filters: [
        {
            name: "volume-size",
            values: ["40"],
        },
        {
            name: "tag:Name",
            values: ["Example"],
        },
    ],
    owners: ["self"],
});
```
```python
import pulumi
import pulumi_aws as aws

ebs_volumes = aws.ebs.get_snapshot_ids(filters=[
        aws.ebs.GetSnapshotIdsFilterArgs(
            name="volume-size",
            values=["40"],
        ),
        aws.ebs.GetSnapshotIdsFilterArgs(
            name="tag:Name",
            values=["Example"],
        ),
    ],
    owners=["self"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ebsVolumes = Aws.Ebs.GetSnapshotIds.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ebs.Inputs.GetSnapshotIdsFilterInputArgs
            {
                Name = "volume-size",
                Values = new[]
                {
                    "40",
                },
            },
            new Aws.Ebs.Inputs.GetSnapshotIdsFilterInputArgs
            {
                Name = "tag:Name",
                Values = new[]
                {
                    "Example",
                },
            },
        },
        Owners = new[]
        {
            "self",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ebs.GetSnapshotIds(ctx, &ebs.GetSnapshotIdsArgs{
			Filters: []ebs.GetSnapshotIdsFilter{
				{
					Name: "volume-size",
					Values: []string{
						"40",
					},
				},
				{
					Name: "tag:Name",
					Values: []string{
						"Example",
					},
				},
			},
			Owners: []string{
				"self",
			},
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
import com.pulumi.aws.ebs.EbsFunctions;
import com.pulumi.aws.ebs.inputs.GetSnapshotIdsArgs;
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
        final var ebsVolumes = EbsFunctions.getSnapshotIds(GetSnapshotIdsArgs.builder()
            .filters(            
                GetSnapshotIdsFilterArgs.builder()
                    .name("volume-size")
                    .values("40")
                    .build(),
                GetSnapshotIdsFilterArgs.builder()
                    .name("tag:Name")
                    .values("Example")
                    .build())
            .owners("self")
            .build());

    }
}
```
```yaml
variables:
  ebsVolumes:
    fn::invoke:
      Function: aws:ebs:getSnapshotIds
      Arguments:
        filters:
          - name: volume-size
            values:
              - '40'
          - name: tag:Name
            values:
              - Example
        owners:
          - self
```
{{% /example %}}
{{% /examples %}}