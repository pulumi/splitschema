Creates a new Amazon Redshift endpoint access.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.EndpointAccess("example", {
    endpointName: "example",
    subnetGroupName: aws_redshift_subnet_group.example.id,
    clusterIdentifier: aws_redshift_cluster.example.cluster_identifier,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.EndpointAccess("example",
    endpoint_name="example",
    subnet_group_name=aws_redshift_subnet_group["example"]["id"],
    cluster_identifier=aws_redshift_cluster["example"]["cluster_identifier"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedShift.EndpointAccess("example", new()
    {
        EndpointName = "example",
        SubnetGroupName = aws_redshift_subnet_group.Example.Id,
        ClusterIdentifier = aws_redshift_cluster.Example.Cluster_identifier,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewEndpointAccess(ctx, "example", &redshift.EndpointAccessArgs{
			EndpointName:      pulumi.String("example"),
			SubnetGroupName:   pulumi.Any(aws_redshift_subnet_group.Example.Id),
			ClusterIdentifier: pulumi.Any(aws_redshift_cluster.Example.Cluster_identifier),
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
import com.pulumi.aws.redshift.EndpointAccess;
import com.pulumi.aws.redshift.EndpointAccessArgs;
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
        var example = new EndpointAccess("example", EndpointAccessArgs.builder()        
            .endpointName("example")
            .subnetGroupName(aws_redshift_subnet_group.example().id())
            .clusterIdentifier(aws_redshift_cluster.example().cluster_identifier())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:EndpointAccess
    properties:
      endpointName: example
      subnetGroupName: ${aws_redshift_subnet_group.example.id}
      clusterIdentifier: ${aws_redshift_cluster.example.cluster_identifier}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift endpoint access using the `name`. For example:

```sh
 $ pulumi import aws:redshift/endpointAccess:EndpointAccess example example
```
 