Use this data source to get information about an EBS Snapshot for use when provisioning EBS Volumes

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ebsVolume = aws.ebs.getSnapshot({
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
    mostRecent: true,
    owners: ["self"],
});
```
```python
import pulumi
import pulumi_aws as aws

ebs_volume = aws.ebs.get_snapshot(filters=[
        aws.ebs.GetSnapshotFilterArgs(
            name="volume-size",
            values=["40"],
        ),
        aws.ebs.GetSnapshotFilterArgs(
            name="tag:Name",
            values=["Example"],
        ),
    ],
    most_recent=True,
    owners=["self"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ebsVolume = Aws.Ebs.GetSnapshot.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ebs.Inputs.GetSnapshotFilterInputArgs
            {
                Name = "volume-size",
                Values = new[]
                {
                    "40",
                },
            },
            new Aws.Ebs.Inputs.GetSnapshotFilterInputArgs
            {
                Name = "tag:Name",
                Values = new[]
                {
                    "Example",
                },
            },
        },
        MostRecent = true,
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
		_, err := ebs.LookupSnapshot(ctx, &ebs.LookupSnapshotArgs{
			Filters: []ebs.GetSnapshotFilter{
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
			MostRecent: pulumi.BoolRef(true),
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
import com.pulumi.aws.ebs.inputs.GetSnapshotArgs;
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
        final var ebsVolume = EbsFunctions.getSnapshot(GetSnapshotArgs.builder()
            .filters(            
                GetSnapshotFilterArgs.builder()
                    .name("volume-size")
                    .values("40")
                    .build(),
                GetSnapshotFilterArgs.builder()
                    .name("tag:Name")
                    .values("Example")
                    .build())
            .mostRecent(true)
            .owners("self")
            .build());

    }
}
```
```yaml
variables:
  ebsVolume:
    fn::invoke:
      Function: aws:ebs:getSnapshot
      Arguments:
        filters:
          - name: volume-size
            values:
              - '40'
          - name: tag:Name
            values:
              - Example
        mostRecent: true
        owners:
          - self
```
{{% /example %}}
{{% /examples %}}