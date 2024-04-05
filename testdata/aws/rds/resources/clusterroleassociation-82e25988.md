Manages a RDS DB Cluster association with an IAM Role. Example use cases:

* [Creating an IAM Role to Allow Amazon Aurora to Access AWS Services](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.Authorizing.IAM.CreateRole.html)
* [Importing Amazon S3 Data into an RDS PostgreSQL DB Cluster](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.S3Import.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.ClusterRoleAssociation("example", {
    dbClusterIdentifier: aws_rds_cluster.example.id,
    featureName: "S3_INTEGRATION",
    roleArn: aws_iam_role.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.ClusterRoleAssociation("example",
    db_cluster_identifier=aws_rds_cluster["example"]["id"],
    feature_name="S3_INTEGRATION",
    role_arn=aws_iam_role["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rds.ClusterRoleAssociation("example", new()
    {
        DbClusterIdentifier = aws_rds_cluster.Example.Id,
        FeatureName = "S3_INTEGRATION",
        RoleArn = aws_iam_role.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewClusterRoleAssociation(ctx, "example", &rds.ClusterRoleAssociationArgs{
			DbClusterIdentifier: pulumi.Any(aws_rds_cluster.Example.Id),
			FeatureName:         pulumi.String("S3_INTEGRATION"),
			RoleArn:             pulumi.Any(aws_iam_role.Example.Arn),
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
import com.pulumi.aws.rds.ClusterRoleAssociation;
import com.pulumi.aws.rds.ClusterRoleAssociationArgs;
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
        var example = new ClusterRoleAssociation("example", ClusterRoleAssociationArgs.builder()        
            .dbClusterIdentifier(aws_rds_cluster.example().id())
            .featureName("S3_INTEGRATION")
            .roleArn(aws_iam_role.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rds:ClusterRoleAssociation
    properties:
      dbClusterIdentifier: ${aws_rds_cluster.example.id}
      featureName: S3_INTEGRATION
      roleArn: ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_rds_cluster_role_association` using the DB Cluster Identifier and IAM Role ARN separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:rds/clusterRoleAssociation:ClusterRoleAssociation example my-db-cluster,arn:aws:iam::123456789012:role/my-role
```
 