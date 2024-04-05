The Batch Job Queue data source allows access to details of a specific
job queue within AWS Batch.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test-queue = aws.batch.getJobQueue({
    name: "tf-test-batch-job-queue",
});
```
```python
import pulumi
import pulumi_aws as aws

test_queue = aws.batch.get_job_queue(name="tf-test-batch-job-queue")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test_queue = Aws.Batch.GetJobQueue.Invoke(new()
    {
        Name = "tf-test-batch-job-queue",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/batch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.LookupJobQueue(ctx, &batch.LookupJobQueueArgs{
			Name: "tf-test-batch-job-queue",
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
import com.pulumi.aws.batch.BatchFunctions;
import com.pulumi.aws.batch.inputs.GetJobQueueArgs;
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
        final var test-queue = BatchFunctions.getJobQueue(GetJobQueueArgs.builder()
            .name("tf-test-batch-job-queue")
            .build());

    }
}
```
```yaml
variables:
  test-queue:
    fn::invoke:
      Function: aws:batch:getJobQueue
      Arguments:
        name: tf-test-batch-job-queue
```
{{% /example %}}
{{% /examples %}}