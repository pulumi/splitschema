Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
By using this data source, you can reference SQS queues without having to hardcode
the ARNs as input.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.sqs.getQueue({
    name: "queue",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sqs.get_queue(name="queue")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Sqs.GetQueue.Invoke(new()
    {
        Name = "queue",
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
		_, err := sqs.LookupQueue(ctx, &sqs.LookupQueueArgs{
			Name: "queue",
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
import com.pulumi.aws.sqs.inputs.GetQueueArgs;
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
        final var example = SqsFunctions.getQueue(GetQueueArgs.builder()
            .name("queue")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:sqs:getQueue
      Arguments:
        name: queue
```
{{% /example %}}
{{% /examples %}}