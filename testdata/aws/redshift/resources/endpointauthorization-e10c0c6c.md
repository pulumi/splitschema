Creates a new Amazon Redshift endpoint authorization.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.EndpointAuthorization("example", {
    account: "01234567910",
    clusterIdentifier: aws_redshift_cluster.example.cluster_identifier,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.EndpointAuthorization("example",
    account="01234567910",
    cluster_identifier=aws_redshift_cluster["example"]["cluster_identifier"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedShift.EndpointAuthorization("example", new()
    {
        Account = "01234567910",
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
		_, err := redshift.NewEndpointAuthorization(ctx, "example", &redshift.EndpointAuthorizationArgs{
			Account:           pulumi.String("01234567910"),
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
import com.pulumi.aws.redshift.EndpointAuthorization;
import com.pulumi.aws.redshift.EndpointAuthorizationArgs;
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
        var example = new EndpointAuthorization("example", EndpointAuthorizationArgs.builder()        
            .account("01234567910")
            .clusterIdentifier(aws_redshift_cluster.example().cluster_identifier())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:EndpointAuthorization
    properties:
      account: '01234567910'
      clusterIdentifier: ${aws_redshift_cluster.example.cluster_identifier}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift endpoint authorization using the `id`. For example:

```sh
 $ pulumi import aws:redshift/endpointAuthorization:EndpointAuthorization example 01234567910:cluster-example-id
```
 