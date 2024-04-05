Provides a DataPipeline Pipeline resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.datapipeline.Pipeline("default", {});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.datapipeline.Pipeline("default")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.DataPipeline.Pipeline("default");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/datapipeline"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datapipeline.NewPipeline(ctx, "default", nil)
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
import com.pulumi.aws.datapipeline.Pipeline;
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
        var default_ = new Pipeline("default");

    }
}
```
```yaml
resources:
  default:
    type: aws:datapipeline:Pipeline
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datapipeline_pipeline` using the id (Pipeline ID). For example:

```sh
 $ pulumi import aws:datapipeline/pipeline:Pipeline default df-1234567890
```
 