Provides a Step Function Activity resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sfnActivity = new aws.sfn.Activity("sfnActivity", {});
```
```python
import pulumi
import pulumi_aws as aws

sfn_activity = aws.sfn.Activity("sfnActivity")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var sfnActivity = new Aws.Sfn.Activity("sfnActivity");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sfn"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sfn.NewActivity(ctx, "sfnActivity", nil)
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
import com.pulumi.aws.sfn.Activity;
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
        var sfnActivity = new Activity("sfnActivity");

    }
}
```
```yaml
resources:
  sfnActivity:
    type: aws:sfn:Activity
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import activities using the `arn`. For example:

```sh
 $ pulumi import aws:sfn/activity:Activity foo arn:aws:states:eu-west-1:123456789098:activity:bar
```
 