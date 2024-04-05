Retrieve information about a AWS Elemental MediaConvert Queue.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.mediaconvert.getQueue({
    id: "tf-example-queue",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.mediaconvert.get_queue(id="tf-example-queue")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.MediaConvert.GetQueue.Invoke(new()
    {
        Id = "tf-example-queue",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mediaconvert"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mediaconvert.LookupQueue(ctx, &mediaconvert.LookupQueueArgs{
			Id: "tf-example-queue",
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
import com.pulumi.aws.mediaconvert.MediaconvertFunctions;
import com.pulumi.aws.mediaconvert.inputs.GetQueueArgs;
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
        final var example = MediaconvertFunctions.getQueue(GetQueueArgs.builder()
            .id("tf-example-queue")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:mediaconvert:getQueue
      Arguments:
        id: tf-example-queue
```
{{% /example %}}
{{% /examples %}}