Resource for managing an AWS Shield Application Layer Automatic Response for automatic DDoS mitigation.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentRegion = aws.getRegion({});
const currentCallerIdentity = aws.getCallerIdentity({});
const currentPartition = aws.getPartition({});
const example = new aws.shield.ApplicationLayerAutomaticResponse("example", {
    action: "COUNT",
    resourceArn: Promise.all([currentPartition, currentCallerIdentity]).then(([currentPartition, currentCallerIdentity]) => `arn:${currentPartition.partition}:cloudfront:${currentCallerIdentity.accountId}:distribution/${_var.distribution_id}`),
});
```
```python
import pulumi
import pulumi_aws as aws

current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
current_partition = aws.get_partition()
example = aws.shield.ApplicationLayerAutomaticResponse("example",
    action="COUNT",
    resource_arn=f"arn:{current_partition.partition}:cloudfront:{current_caller_identity.account_id}:distribution/{var['distribution_id']}")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentRegion = Aws.GetRegion.Invoke();

    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var currentPartition = Aws.GetPartition.Invoke();

    var example = new Aws.Shield.ApplicationLayerAutomaticResponse("example", new()
    {
        Action = "COUNT",
        ResourceArn = Output.Tuple(currentPartition, currentCallerIdentity).Apply(values =>
        {
            var currentPartition = values.Item1;
            var currentCallerIdentity = values.Item2;
            return $"arn:{currentPartition.Apply(getPartitionResult => getPartitionResult.Partition)}:cloudfront:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:distribution/{@var.Distribution_id}";
        }),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/shield"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentPartition, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = shield.NewApplicationLayerAutomaticResponse(ctx, "example", &shield.ApplicationLayerAutomaticResponseArgs{
			Action:      pulumi.String("COUNT"),
			ResourceArn: pulumi.String(fmt.Sprintf("arn:%v:cloudfront:%v:distribution/%v", currentPartition.Partition, currentCallerIdentity.AccountId, _var.Distribution_id)),
		})
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.shield.ApplicationLayerAutomaticResponse;
import com.pulumi.aws.shield.ApplicationLayerAutomaticResponseArgs;
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
        final var currentRegion = AwsFunctions.getRegion();

        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        final var currentPartition = AwsFunctions.getPartition();

        var example = new ApplicationLayerAutomaticResponse("example", ApplicationLayerAutomaticResponseArgs.builder()        
            .action("COUNT")
            .resourceArn(String.format("arn:%s:cloudfront:%s:distribution/%s", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),var_.distribution_id()))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:shield:ApplicationLayerAutomaticResponse
    properties:
      action: COUNT
      resourceArn: arn:${currentPartition.partition}:cloudfront:${currentCallerIdentity.accountId}:distribution/${var.distribution_id}
variables:
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  currentPartition:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}