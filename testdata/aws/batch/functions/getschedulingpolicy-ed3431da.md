The Batch Scheduling Policy data source allows access to details of a specific Scheduling Policy within AWS Batch.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.batch.getSchedulingPolicy({
    arn: "arn:aws:batch:us-east-1:012345678910:scheduling-policy/example",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.batch.get_scheduling_policy(arn="arn:aws:batch:us-east-1:012345678910:scheduling-policy/example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Batch.GetSchedulingPolicy.Invoke(new()
    {
        Arn = "arn:aws:batch:us-east-1:012345678910:scheduling-policy/example",
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
		_, err := batch.LookupSchedulingPolicy(ctx, &batch.LookupSchedulingPolicyArgs{
			Arn: "arn:aws:batch:us-east-1:012345678910:scheduling-policy/example",
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
import com.pulumi.aws.batch.inputs.GetSchedulingPolicyArgs;
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
        final var test = BatchFunctions.getSchedulingPolicy(GetSchedulingPolicyArgs.builder()
            .arn("arn:aws:batch:us-east-1:012345678910:scheduling-policy/example")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:batch:getSchedulingPolicy
      Arguments:
        arn: arn:aws:batch:us-east-1:012345678910:scheduling-policy/example
```
{{% /example %}}
{{% /examples %}}