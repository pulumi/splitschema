Manages a Kinesis Analytics v2 Application Snapshot.
Snapshots are the AWS implementation of [Flink Savepoints](https://ci.apache.org/projects/flink/flink-docs-release-1.11/ops/state/savepoints.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kinesisanalyticsv2.ApplicationSnapshot("example", {
    applicationName: aws_kinesisanalyticsv2_application.example.name,
    snapshotName: "example-snapshot",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kinesisanalyticsv2.ApplicationSnapshot("example",
    application_name=aws_kinesisanalyticsv2_application["example"]["name"],
    snapshot_name="example-snapshot")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.KinesisAnalyticsV2.ApplicationSnapshot("example", new()
    {
        ApplicationName = aws_kinesisanalyticsv2_application.Example.Name,
        SnapshotName = "example-snapshot",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesisanalyticsv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kinesisanalyticsv2.NewApplicationSnapshot(ctx, "example", &kinesisanalyticsv2.ApplicationSnapshotArgs{
			ApplicationName: pulumi.Any(aws_kinesisanalyticsv2_application.Example.Name),
			SnapshotName:    pulumi.String("example-snapshot"),
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
import com.pulumi.aws.kinesisanalyticsv2.ApplicationSnapshot;
import com.pulumi.aws.kinesisanalyticsv2.ApplicationSnapshotArgs;
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
        var example = new ApplicationSnapshot("example", ApplicationSnapshotArgs.builder()        
            .applicationName(aws_kinesisanalyticsv2_application.example().name())
            .snapshotName("example-snapshot")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kinesisanalyticsv2:ApplicationSnapshot
    properties:
      applicationName: ${aws_kinesisanalyticsv2_application.example.name}
      snapshotName: example-snapshot
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_kinesisanalyticsv2_application` using `application_name` together with `snapshot_name`. For example:

```sh
 $ pulumi import aws:kinesisanalyticsv2/applicationSnapshot:ApplicationSnapshot example example-application/example-snapshot
```
 