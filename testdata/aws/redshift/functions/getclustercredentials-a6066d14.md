Provides redshift cluster temporary credentials.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.redshift.getClusterCredentials({
    clusterIdentifier: aws_redshift_cluster.example.cluster_identifier,
    dbUser: aws_redshift_cluster.example.master_username,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.get_cluster_credentials(cluster_identifier=aws_redshift_cluster["example"]["cluster_identifier"],
    db_user=aws_redshift_cluster["example"]["master_username"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.RedShift.GetClusterCredentials.Invoke(new()
    {
        ClusterIdentifier = aws_redshift_cluster.Example.Cluster_identifier,
        DbUser = aws_redshift_cluster.Example.Master_username,
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
		_, err := redshift.GetClusterCredentials(ctx, &redshift.GetClusterCredentialsArgs{
			ClusterIdentifier: aws_redshift_cluster.Example.Cluster_identifier,
			DbUser:            aws_redshift_cluster.Example.Master_username,
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
import com.pulumi.aws.redshift.RedshiftFunctions;
import com.pulumi.aws.redshift.inputs.GetClusterCredentialsArgs;
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
        final var example = RedshiftFunctions.getClusterCredentials(GetClusterCredentialsArgs.builder()
            .clusterIdentifier(aws_redshift_cluster.example().cluster_identifier())
            .dbUser(aws_redshift_cluster.example().master_username())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:redshift:getClusterCredentials
      Arguments:
        clusterIdentifier: ${aws_redshift_cluster.example.cluster_identifier}
        dbUser: ${aws_redshift_cluster.example.master_username}
```
{{% /example %}}
{{% /examples %}}