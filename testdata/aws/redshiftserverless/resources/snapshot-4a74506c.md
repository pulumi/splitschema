Creates a new Amazon Redshift Serverless Snapshot.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshiftserverless.Snapshot("example", {
    namespaceName: aws_redshiftserverless_workgroup.example.namespace_name,
    snapshotName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshiftserverless.Snapshot("example",
    namespace_name=aws_redshiftserverless_workgroup["example"]["namespace_name"],
    snapshot_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedshiftServerless.Snapshot("example", new()
    {
        NamespaceName = aws_redshiftserverless_workgroup.Example.Namespace_name,
        SnapshotName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshiftserverless"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshiftserverless.NewSnapshot(ctx, "example", &redshiftserverless.SnapshotArgs{
			NamespaceName: pulumi.Any(aws_redshiftserverless_workgroup.Example.Namespace_name),
			SnapshotName:  pulumi.String("example"),
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
import com.pulumi.aws.redshiftserverless.Snapshot;
import com.pulumi.aws.redshiftserverless.SnapshotArgs;
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
        var example = new Snapshot("example", SnapshotArgs.builder()        
            .namespaceName(aws_redshiftserverless_workgroup.example().namespace_name())
            .snapshotName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshiftserverless:Snapshot
    properties:
      namespaceName: ${aws_redshiftserverless_workgroup.example.namespace_name}
      snapshotName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Serverless Snapshots using the `snapshot_name`. For example:

```sh
 $ pulumi import aws:redshiftserverless/snapshot:Snapshot example example
```
 