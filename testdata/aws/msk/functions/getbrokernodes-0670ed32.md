Get information on an Amazon MSK Broker Nodes.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.msk.getBrokerNodes({
    clusterArn: aws_msk_cluster.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.msk.get_broker_nodes(cluster_arn=aws_msk_cluster["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Msk.GetBrokerNodes.Invoke(new()
    {
        ClusterArn = aws_msk_cluster.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/msk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := msk.GetBrokerNodes(ctx, &msk.GetBrokerNodesArgs{
			ClusterArn: aws_msk_cluster.Example.Arn,
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
import com.pulumi.aws.msk.MskFunctions;
import com.pulumi.aws.msk.inputs.GetBrokerNodesArgs;
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
        final var example = MskFunctions.getBrokerNodes(GetBrokerNodesArgs.builder()
            .clusterArn(aws_msk_cluster.example().arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:msk:getBrokerNodes
      Arguments:
        clusterArn: ${aws_msk_cluster.example.arn}
```
{{% /example %}}
{{% /examples %}}