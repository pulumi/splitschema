Provides a resource to manage a Resource Explorer index in the current AWS Region.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.resourceexplorer.Index("example", {type: "LOCAL"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.resourceexplorer.Index("example", type="LOCAL")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ResourceExplorer.Index("example", new()
    {
        Type = "LOCAL",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/resourceexplorer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resourceexplorer.NewIndex(ctx, "example", &resourceexplorer.IndexArgs{
			Type: pulumi.String("LOCAL"),
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
import com.pulumi.aws.resourceexplorer.Index;
import com.pulumi.aws.resourceexplorer.IndexArgs;
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
        var example = new Index("example", IndexArgs.builder()        
            .type("LOCAL")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:resourceexplorer:Index
    properties:
      type: LOCAL
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Resource Explorer indexes using the `arn`. For example:

```sh
 $ pulumi import aws:resourceexplorer/index:Index example arn:aws:resource-explorer-2:us-east-1:123456789012:index/6047ac4e-207e-4487-9bcf-cb53bb0ff5cc
```
 