Provides a Redshift Cluster IAM Roles resource.

> **NOTE:** A Redshift cluster's default IAM role can be managed both by this resource's `default_iam_role_arn` argument and the `aws.redshift.Cluster` resource's `default_iam_role_arn` argument. Do not configure different values for both arguments. Doing so will cause a conflict of default IAM roles.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.ClusterIamRoles("example", {
    clusterIdentifier: aws_redshift_cluster.example.cluster_identifier,
    iamRoleArns: [aws_iam_role.example.arn],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.ClusterIamRoles("example",
    cluster_identifier=aws_redshift_cluster["example"]["cluster_identifier"],
    iam_role_arns=[aws_iam_role["example"]["arn"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedShift.ClusterIamRoles("example", new()
    {
        ClusterIdentifier = aws_redshift_cluster.Example.Cluster_identifier,
        IamRoleArns = new[]
        {
            aws_iam_role.Example.Arn,
        },
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
		_, err := redshift.NewClusterIamRoles(ctx, "example", &redshift.ClusterIamRolesArgs{
			ClusterIdentifier: pulumi.Any(aws_redshift_cluster.Example.Cluster_identifier),
			IamRoleArns: pulumi.StringArray{
				aws_iam_role.Example.Arn,
			},
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
import com.pulumi.aws.redshift.ClusterIamRoles;
import com.pulumi.aws.redshift.ClusterIamRolesArgs;
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
        var example = new ClusterIamRoles("example", ClusterIamRolesArgs.builder()        
            .clusterIdentifier(aws_redshift_cluster.example().cluster_identifier())
            .iamRoleArns(aws_iam_role.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:ClusterIamRoles
    properties:
      clusterIdentifier: ${aws_redshift_cluster.example.cluster_identifier}
      iamRoleArns:
        - ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Cluster IAM Roless using the `cluster_identifier`. For example:

```sh
 $ pulumi import aws:redshift/clusterIamRoles:ClusterIamRoles examplegroup1 example
```
 