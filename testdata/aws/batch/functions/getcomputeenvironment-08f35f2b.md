The Batch Compute Environment data source allows access to details of a specific
compute environment within AWS Batch.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const batch-mongo = aws.batch.getComputeEnvironment({
    computeEnvironmentName: "batch-mongo-production",
});
```
```python
import pulumi
import pulumi_aws as aws

batch_mongo = aws.batch.get_compute_environment(compute_environment_name="batch-mongo-production")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var batch_mongo = Aws.Batch.GetComputeEnvironment.Invoke(new()
    {
        ComputeEnvironmentName = "batch-mongo-production",
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
		_, err := batch.LookupComputeEnvironment(ctx, &batch.LookupComputeEnvironmentArgs{
			ComputeEnvironmentName: "batch-mongo-production",
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
import com.pulumi.aws.batch.inputs.GetComputeEnvironmentArgs;
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
        final var batch-mongo = BatchFunctions.getComputeEnvironment(GetComputeEnvironmentArgs.builder()
            .computeEnvironmentName("batch-mongo-production")
            .build());

    }
}
```
```yaml
variables:
  batch-mongo:
    fn::invoke:
      Function: aws:batch:getComputeEnvironment
      Arguments:
        computeEnvironmentName: batch-mongo-production
```
{{% /example %}}
{{% /examples %}}