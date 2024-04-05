Use this data source to retrieve information about an EventBridge connection.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.cloudwatch.getEventConnection({
    name: "test",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.cloudwatch.get_event_connection(name="test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.CloudWatch.GetEventConnection.Invoke(new()
    {
        Name = "test",
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
		_, err := cloudwatch.LookupEventConnection(ctx, &cloudwatch.LookupEventConnectionArgs{
			Name: "test",
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
import com.pulumi.aws.cloudwatch.CloudwatchFunctions;
import com.pulumi.aws.cloudwatch.inputs.GetEventConnectionArgs;
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
        final var test = CloudwatchFunctions.getEventConnection(GetEventConnectionArgs.builder()
            .name("test")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:cloudwatch:getEventConnection
      Arguments:
        name: test
```
{{% /example %}}
{{% /examples %}}