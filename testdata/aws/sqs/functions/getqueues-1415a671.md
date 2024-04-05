Data source for managing an AWS SQS (Simple Queue) Queues.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.sqs.getQueues({
    queueNamePrefix: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sqs.get_queues(queue_name_prefix="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Sqs.GetQueues.Invoke(new()
    {
        QueueNamePrefix = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sqs.GetQueues(ctx, &sqs.GetQueuesArgs{
			QueueNamePrefix: pulumi.StringRef("example"),
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
import com.pulumi.aws.sqs.SqsFunctions;
import com.pulumi.aws.sqs.inputs.GetQueuesArgs;
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
        final var example = SqsFunctions.getQueues(GetQueuesArgs.builder()
            .queueNamePrefix("example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:sqs:getQueues
      Arguments:
        queueNamePrefix: example
```
{{% /example %}}
{{% /examples %}}