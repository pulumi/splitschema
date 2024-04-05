Use this data source to get information about an EventBridge Partner Event Source. This data source will only return one partner event source. An error will be returned if multiple sources match the same name prefix.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplepartner = aws.cloudwatch.getEventSource({
    namePrefix: "aws.partner/examplepartner.com",
});
```
```python
import pulumi
import pulumi_aws as aws

examplepartner = aws.cloudwatch.get_event_source(name_prefix="aws.partner/examplepartner.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplepartner = Aws.CloudWatch.GetEventSource.Invoke(new()
    {
        NamePrefix = "aws.partner/examplepartner.com",
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
		_, err := cloudwatch.GetEventSource(ctx, &cloudwatch.GetEventSourceArgs{
			NamePrefix: pulumi.StringRef("aws.partner/examplepartner.com"),
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
import com.pulumi.aws.cloudwatch.inputs.GetEventSourceArgs;
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
        final var examplepartner = CloudwatchFunctions.getEventSource(GetEventSourceArgs.builder()
            .namePrefix("aws.partner/examplepartner.com")
            .build());

    }
}
```
```yaml
variables:
  examplepartner:
    fn::invoke:
      Function: aws:cloudwatch:getEventSource
      Arguments:
        namePrefix: aws.partner/examplepartner.com
```
{{% /example %}}
{{% /examples %}}