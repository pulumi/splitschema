Provides a Synthetics Group resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.synthetics.Group("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.synthetics.Group("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Synthetics.Group("example");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/synthetics"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synthetics.NewGroup(ctx, "example", nil)
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
import com.pulumi.aws.synthetics.Group;
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
        var example = new Group("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:synthetics:Group
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch Synthetics Group using the `name`. For example:

```sh
 $ pulumi import aws:synthetics/group:Group example example
```
 