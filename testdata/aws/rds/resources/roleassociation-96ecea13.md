Manages an RDS DB Instance association with an IAM Role. Example use cases:

* [Amazon RDS Oracle integration with Amazon S3](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html)
* [Importing Amazon S3 Data into an RDS PostgreSQL DB Instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.S3Import.html)

> To manage the RDS DB Instance IAM Role for [Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html), see the `aws.rds.Instance` resource `monitoring_role_arn` argument instead.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.RoleAssociation("example", {
    dbInstanceIdentifier: aws_db_instance.example.identifier,
    featureName: "S3_INTEGRATION",
    roleArn: aws_iam_role.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.RoleAssociation("example",
    db_instance_identifier=aws_db_instance["example"]["identifier"],
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
    var example = new Aws.Rds.RoleAssociation("example", new()
    {
        DbInstanceIdentifier = aws_db_instance.Example.Identifier,
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
		_, err := rds.NewRoleAssociation(ctx, "example", &rds.RoleAssociationArgs{
			DbInstanceIdentifier: pulumi.Any(aws_db_instance.Example.Identifier),
			FeatureName:          pulumi.String("S3_INTEGRATION"),
			RoleArn:              pulumi.Any(aws_iam_role.Example.Arn),
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
import com.pulumi.aws.rds.RoleAssociation;
import com.pulumi.aws.rds.RoleAssociationArgs;
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
        var example = new RoleAssociation("example", RoleAssociationArgs.builder()        
            .dbInstanceIdentifier(aws_db_instance.example().identifier())
            .featureName("S3_INTEGRATION")
            .roleArn(aws_iam_role.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rds:RoleAssociation
    properties:
      dbInstanceIdentifier: ${aws_db_instance.example.identifier}
      featureName: S3_INTEGRATION
      roleArn: ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_db_instance_role_association` using the DB Instance Identifier and IAM Role ARN separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:rds/roleAssociation:RoleAssociation example my-db-instance,arn:aws:iam::123456789012:role/my-role
```
 