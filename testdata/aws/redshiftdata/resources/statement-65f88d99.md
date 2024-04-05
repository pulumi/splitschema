Executes a Redshift Data Statement.

{{% examples %}}
## Example Usage
{{% example %}}
### cluster_identifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshiftdata.Statement("example", {
    clusterIdentifier: aws_redshift_cluster.example.cluster_identifier,
    database: aws_redshift_cluster.example.database_name,
    dbUser: aws_redshift_cluster.example.master_username,
    sql: "CREATE GROUP group_name;",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshiftdata.Statement("example",
    cluster_identifier=aws_redshift_cluster["example"]["cluster_identifier"],
    database=aws_redshift_cluster["example"]["database_name"],
    db_user=aws_redshift_cluster["example"]["master_username"],
    sql="CREATE GROUP group_name;")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedshiftData.Statement("example", new()
    {
        ClusterIdentifier = aws_redshift_cluster.Example.Cluster_identifier,
        Database = aws_redshift_cluster.Example.Database_name,
        DbUser = aws_redshift_cluster.Example.Master_username,
        Sql = "CREATE GROUP group_name;",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshiftdata"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshiftdata.NewStatement(ctx, "example", &redshiftdata.StatementArgs{
			ClusterIdentifier: pulumi.Any(aws_redshift_cluster.Example.Cluster_identifier),
			Database:          pulumi.Any(aws_redshift_cluster.Example.Database_name),
			DbUser:            pulumi.Any(aws_redshift_cluster.Example.Master_username),
			Sql:               pulumi.String("CREATE GROUP group_name;"),
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
import com.pulumi.aws.redshiftdata.Statement;
import com.pulumi.aws.redshiftdata.StatementArgs;
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
        var example = new Statement("example", StatementArgs.builder()        
            .clusterIdentifier(aws_redshift_cluster.example().cluster_identifier())
            .database(aws_redshift_cluster.example().database_name())
            .dbUser(aws_redshift_cluster.example().master_username())
            .sql("CREATE GROUP group_name;")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshiftdata:Statement
    properties:
      clusterIdentifier: ${aws_redshift_cluster.example.cluster_identifier}
      database: ${aws_redshift_cluster.example.database_name}
      dbUser: ${aws_redshift_cluster.example.master_username}
      sql: CREATE GROUP group_name;
```
{{% /example %}}
{{% example %}}
### workgroup_name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshiftdata.Statement("example", {
    workgroupName: aws_redshiftserverless_workgroup.example.workgroup_name,
    database: "dev",
    sql: "CREATE GROUP group_name;",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshiftdata.Statement("example",
    workgroup_name=aws_redshiftserverless_workgroup["example"]["workgroup_name"],
    database="dev",
    sql="CREATE GROUP group_name;")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedshiftData.Statement("example", new()
    {
        WorkgroupName = aws_redshiftserverless_workgroup.Example.Workgroup_name,
        Database = "dev",
        Sql = "CREATE GROUP group_name;",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshiftdata"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshiftdata.NewStatement(ctx, "example", &redshiftdata.StatementArgs{
			WorkgroupName: pulumi.Any(aws_redshiftserverless_workgroup.Example.Workgroup_name),
			Database:      pulumi.String("dev"),
			Sql:           pulumi.String("CREATE GROUP group_name;"),
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
import com.pulumi.aws.redshiftdata.Statement;
import com.pulumi.aws.redshiftdata.StatementArgs;
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
        var example = new Statement("example", StatementArgs.builder()        
            .workgroupName(aws_redshiftserverless_workgroup.example().workgroup_name())
            .database("dev")
            .sql("CREATE GROUP group_name;")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshiftdata:Statement
    properties:
      workgroupName: ${aws_redshiftserverless_workgroup.example.workgroup_name}
      database: dev
      sql: CREATE GROUP group_name;
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Data Statements using the `id`. For example:

```sh
 $ pulumi import aws:redshiftdata/statement:Statement example example
```
 