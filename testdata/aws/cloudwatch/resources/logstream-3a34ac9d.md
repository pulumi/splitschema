Provides a CloudWatch Log Stream resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const yada = new aws.cloudwatch.LogGroup("yada", {});
const foo = new aws.cloudwatch.LogStream("foo", {logGroupName: yada.name});
```
```python
import pulumi
import pulumi_aws as aws

yada = aws.cloudwatch.LogGroup("yada")
foo = aws.cloudwatch.LogStream("foo", log_group_name=yada.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var yada = new Aws.CloudWatch.LogGroup("yada");

    var foo = new Aws.CloudWatch.LogStream("foo", new()
    {
        LogGroupName = yada.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		yada, err := cloudwatch.NewLogGroup(ctx, "yada", nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogStream(ctx, "foo", &cloudwatch.LogStreamArgs{
			LogGroupName: yada.Name,
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.cloudwatch.LogStream;
import com.pulumi.aws.cloudwatch.LogStreamArgs;
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
        var yada = new LogGroup("yada");

        var foo = new LogStream("foo", LogStreamArgs.builder()        
            .logGroupName(yada.name())
            .build());

    }
}
```
```yaml
resources:
  yada:
    type: aws:cloudwatch:LogGroup
  foo:
    type: aws:cloudwatch:LogStream
    properties:
      logGroupName: ${yada.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cloudwatch Log Stream using the stream's `log_group_name` and `name`. For example:

```sh
 $ pulumi import aws:cloudwatch/logStream:LogStream foo Yada:SampleLogStream1234
```
 