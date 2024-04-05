Provides an Neptune Cluster Endpoint Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.neptune.ClusterEndpoint("example", {
    clusterIdentifier: aws_neptune_cluster.test.cluster_identifier,
    clusterEndpointIdentifier: "example",
    endpointType: "READER",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.neptune.ClusterEndpoint("example",
    cluster_identifier=aws_neptune_cluster["test"]["cluster_identifier"],
    cluster_endpoint_identifier="example",
    endpoint_type="READER")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Neptune.ClusterEndpoint("example", new()
    {
        ClusterIdentifier = aws_neptune_cluster.Test.Cluster_identifier,
        ClusterEndpointIdentifier = "example",
        EndpointType = "READER",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/neptune"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := neptune.NewClusterEndpoint(ctx, "example", &neptune.ClusterEndpointArgs{
			ClusterIdentifier:         pulumi.Any(aws_neptune_cluster.Test.Cluster_identifier),
			ClusterEndpointIdentifier: pulumi.String("example"),
			EndpointType:              pulumi.String("READER"),
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
import com.pulumi.aws.neptune.ClusterEndpoint;
import com.pulumi.aws.neptune.ClusterEndpointArgs;
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
        var example = new ClusterEndpoint("example", ClusterEndpointArgs.builder()        
            .clusterIdentifier(aws_neptune_cluster.test().cluster_identifier())
            .clusterEndpointIdentifier("example")
            .endpointType("READER")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:neptune:ClusterEndpoint
    properties:
      clusterIdentifier: ${aws_neptune_cluster.test.cluster_identifier}
      clusterEndpointIdentifier: example
      endpointType: READER
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_neptune_cluster_endpoint` using the `cluster-identifier:endpoint-identfier`. For example:

```sh
 $ pulumi import aws:neptune/clusterEndpoint:ClusterEndpoint example my-cluster:my-endpoint
```
 