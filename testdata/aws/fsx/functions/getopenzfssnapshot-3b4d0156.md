Use this data source to get information about an Amazon FSx for OpenZFS Snapshot for use when provisioning new Volumes.

{{% examples %}}
## Example Usage
{{% example %}}
### Root volume Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.fsx.getOpenZfsSnapshot({
    filters: [{
        name: "volume-id",
        values: ["fsvol-073a32b6098a73feb"],
    }],
    mostRecent: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.fsx.get_open_zfs_snapshot(filters=[aws.fsx.GetOpenZfsSnapshotFilterArgs(
        name="volume-id",
        values=["fsvol-073a32b6098a73feb"],
    )],
    most_recent=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Fsx.GetOpenZfsSnapshot.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Fsx.Inputs.GetOpenZfsSnapshotFilterInputArgs
            {
                Name = "volume-id",
                Values = new[]
                {
                    "fsvol-073a32b6098a73feb",
                },
            },
        },
        MostRecent = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.LookupOpenZfsSnapshot(ctx, &fsx.LookupOpenZfsSnapshotArgs{
			Filters: []fsx.GetOpenZfsSnapshotFilter{
				{
					Name: "volume-id",
					Values: []string{
						"fsvol-073a32b6098a73feb",
					},
				},
			},
			MostRecent: pulumi.BoolRef(true),
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
import com.pulumi.aws.fsx.FsxFunctions;
import com.pulumi.aws.fsx.inputs.GetOpenZfsSnapshotArgs;
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
        final var example = FsxFunctions.getOpenZfsSnapshot(GetOpenZfsSnapshotArgs.builder()
            .filters(GetOpenZfsSnapshotFilterArgs.builder()
                .name("volume-id")
                .values("fsvol-073a32b6098a73feb")
                .build())
            .mostRecent(true)
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:fsx:getOpenZfsSnapshot
      Arguments:
        filters:
          - name: volume-id
            values:
              - fsvol-073a32b6098a73feb
        mostRecent: true
```
{{% /example %}}
{{% /examples %}}