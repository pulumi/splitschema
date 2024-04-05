Creates a new Amazon Redshift Serverless Workgroup.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshiftserverless.Workgroup("example", {
    namespaceName: "concurrency-scaling",
    workgroupName: "concurrency-scaling",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshiftserverless.Workgroup("example",
    namespace_name="concurrency-scaling",
    workgroup_name="concurrency-scaling")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedshiftServerless.Workgroup("example", new()
    {
        NamespaceName = "concurrency-scaling",
        WorkgroupName = "concurrency-scaling",
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
		_, err := redshiftserverless.NewWorkgroup(ctx, "example", &redshiftserverless.WorkgroupArgs{
			NamespaceName: pulumi.String("concurrency-scaling"),
			WorkgroupName: pulumi.String("concurrency-scaling"),
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
import com.pulumi.aws.redshiftserverless.Workgroup;
import com.pulumi.aws.redshiftserverless.WorkgroupArgs;
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
        var example = new Workgroup("example", WorkgroupArgs.builder()        
            .namespaceName("concurrency-scaling")
            .workgroupName("concurrency-scaling")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshiftserverless:Workgroup
    properties:
      namespaceName: concurrency-scaling
      workgroupName: concurrency-scaling
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Serverless Workgroups using the `workgroup_name`. For example:

```sh
 $ pulumi import aws:redshiftserverless/workgroup:Workgroup example example
```
 