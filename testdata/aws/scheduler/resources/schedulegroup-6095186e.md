Provides an EventBridge Scheduler Schedule Group resource.

You can find out more about EventBridge Scheduler in the [User Guide](https://docs.aws.amazon.com/scheduler/latest/UserGuide/what-is-scheduler.html).

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.scheduler.ScheduleGroup("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.scheduler.ScheduleGroup("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Scheduler.ScheduleGroup("example");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/scheduler"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := scheduler.NewScheduleGroup(ctx, "example", nil)
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
import com.pulumi.aws.scheduler.ScheduleGroup;
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
        var example = new ScheduleGroup("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:scheduler:ScheduleGroup
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import schedule groups using the `name`. For example:

```sh
 $ pulumi import aws:scheduler/scheduleGroup:ScheduleGroup example my-schedule-group
```
 