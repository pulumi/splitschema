Use this data source to get information on an existing autoscaling group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = aws.autoscaling.getGroup({
    name: "foo",
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.autoscaling.get_group(name="foo")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = Aws.AutoScaling.GetGroup.Invoke(new()
    {
        Name = "foo",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := autoscaling.LookupGroup(ctx, &autoscaling.LookupGroupArgs{
			Name: "foo",
		}, nil)
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
import com.pulumi.aws.autoscaling.AutoscalingFunctions;
import com.pulumi.aws.autoscaling.inputs.GetGroupArgs;
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
        final var foo = AutoscalingFunctions.getGroup(GetGroupArgs.builder()
            .name("foo")
            .build());

    }
}
```
```yaml
variables:
  foo:
    fn::invoke:
      Function: aws:autoscaling:getGroup
      Arguments:
        name: foo
```
{{% /example %}}
{{% /examples %}}