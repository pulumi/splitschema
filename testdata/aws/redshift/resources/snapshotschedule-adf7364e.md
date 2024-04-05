{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.redshift.SnapshotSchedule("default", {
    definitions: ["rate(12 hours)"],
    identifier: "tf-redshift-snapshot-schedule",
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.redshift.SnapshotSchedule("default",
    definitions=["rate(12 hours)"],
    identifier="tf-redshift-snapshot-schedule")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.RedShift.SnapshotSchedule("default", new()
    {
        Definitions = new[]
        {
            "rate(12 hours)",
        },
        Identifier = "tf-redshift-snapshot-schedule",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewSnapshotSchedule(ctx, "default", &redshift.SnapshotScheduleArgs{
			Definitions: pulumi.StringArray{
				pulumi.String("rate(12 hours)"),
			},
			Identifier: pulumi.String("tf-redshift-snapshot-schedule"),
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
import com.pulumi.aws.redshift.SnapshotSchedule;
import com.pulumi.aws.redshift.SnapshotScheduleArgs;
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
        var default_ = new SnapshotSchedule("default", SnapshotScheduleArgs.builder()        
            .definitions("rate(12 hours)")
            .identifier("tf-redshift-snapshot-schedule")
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:redshift:SnapshotSchedule
    properties:
      definitions:
        - rate(12 hours)
      identifier: tf-redshift-snapshot-schedule
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Snapshot Schedule using the `identifier`. For example:

```sh
 $ pulumi import aws:redshift/snapshotSchedule:SnapshotSchedule default tf-redshift-snapshot-schedule
```
 