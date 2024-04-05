The ECS Service data source allows access to details of a specific
Service within a AWS ECS Cluster.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ecs.getService({
    serviceName: "example",
    clusterArn: data.aws_ecs_cluster.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ecs.get_service(service_name="example",
    cluster_arn=data["aws_ecs_cluster"]["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ecs.GetService.Invoke(new()
    {
        ServiceName = "example",
        ClusterArn = data.Aws_ecs_cluster.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.LookupService(ctx, &ecs.LookupServiceArgs{
			ServiceName: "example",
			ClusterArn:  data.Aws_ecs_cluster.Example.Arn,
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
import com.pulumi.aws.ecs.EcsFunctions;
import com.pulumi.aws.ecs.inputs.GetServiceArgs;
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
        final var example = EcsFunctions.getService(GetServiceArgs.builder()
            .serviceName("example")
            .clusterArn(data.aws_ecs_cluster().example().arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ecs:getService
      Arguments:
        serviceName: example
        clusterArn: ${data.aws_ecs_cluster.example.arn}
```
{{% /example %}}
{{% /examples %}}