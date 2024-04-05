Get permissions for a principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3. Permissions are granted to a principal, in a Data Catalog, relative to a Lake Formation resource, which includes the Data Catalog, databases, tables, LF-tags, and LF-tag policies. For more information, see [Security and Access Control to Metadata and Data in Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html).

> **NOTE:** This data source deals with explicitly granted permissions. Lake Formation grants implicit permissions to data lake administrators, database creators, and table creators. For more information, see [Implicit Lake Formation Permissions](https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Permissions For A Lake Formation S3 Resource

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.lakeformation.getPermissions({
    principal: aws_iam_role.workflow_role.arn,
    dataLocation: {
        arn: aws_lakeformation_resource.test.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lakeformation.get_permissions(principal=aws_iam_role["workflow_role"]["arn"],
    data_location=aws.lakeformation.GetPermissionsDataLocationArgs(
        arn=aws_lakeformation_resource["test"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.LakeFormation.GetPermissions.Invoke(new()
    {
        Principal = aws_iam_role.Workflow_role.Arn,
        DataLocation = new Aws.LakeFormation.Inputs.GetPermissionsDataLocationInputArgs
        {
            Arn = aws_lakeformation_resource.Test.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lakeformation.LookupPermissions(ctx, &lakeformation.LookupPermissionsArgs{
			Principal: aws_iam_role.Workflow_role.Arn,
			DataLocation: lakeformation.GetPermissionsDataLocation{
				Arn: aws_lakeformation_resource.Test.Arn,
			},
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
import com.pulumi.aws.lakeformation.LakeformationFunctions;
import com.pulumi.aws.lakeformation.inputs.GetPermissionsArgs;
import com.pulumi.aws.lakeformation.inputs.GetPermissionsDataLocationArgs;
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
        final var test = LakeformationFunctions.getPermissions(GetPermissionsArgs.builder()
            .principal(aws_iam_role.workflow_role().arn())
            .dataLocation(GetPermissionsDataLocationArgs.builder()
                .arn(aws_lakeformation_resource.test().arn())
                .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:lakeformation:getPermissions
      Arguments:
        principal: ${aws_iam_role.workflow_role.arn}
        dataLocation:
          arn: ${aws_lakeformation_resource.test.arn}
```
{{% /example %}}
{{% example %}}
### Permissions For A Glue Catalog Database

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.lakeformation.getPermissions({
    principal: aws_iam_role.workflow_role.arn,
    database: {
        name: aws_glue_catalog_database.test.name,
        catalogId: "110376042874",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lakeformation.get_permissions(principal=aws_iam_role["workflow_role"]["arn"],
    database=aws.lakeformation.GetPermissionsDatabaseArgs(
        name=aws_glue_catalog_database["test"]["name"],
        catalog_id="110376042874",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.LakeFormation.GetPermissions.Invoke(new()
    {
        Principal = aws_iam_role.Workflow_role.Arn,
        Database = new Aws.LakeFormation.Inputs.GetPermissionsDatabaseInputArgs
        {
            Name = aws_glue_catalog_database.Test.Name,
            CatalogId = "110376042874",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lakeformation.LookupPermissions(ctx, &lakeformation.LookupPermissionsArgs{
			Principal: aws_iam_role.Workflow_role.Arn,
			Database: lakeformation.GetPermissionsDatabase{
				Name:      aws_glue_catalog_database.Test.Name,
				CatalogId: "110376042874",
			},
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
import com.pulumi.aws.lakeformation.LakeformationFunctions;
import com.pulumi.aws.lakeformation.inputs.GetPermissionsArgs;
import com.pulumi.aws.lakeformation.inputs.GetPermissionsDatabaseArgs;
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
        final var test = LakeformationFunctions.getPermissions(GetPermissionsArgs.builder()
            .principal(aws_iam_role.workflow_role().arn())
            .database(GetPermissionsDatabaseArgs.builder()
                .name(aws_glue_catalog_database.test().name())
                .catalogId("110376042874")
                .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:lakeformation:getPermissions
      Arguments:
        principal: ${aws_iam_role.workflow_role.arn}
        database:
          name: ${aws_glue_catalog_database.test.name}
          catalogId: '110376042874'
```
{{% /example %}}
{{% example %}}
### Permissions For Tag-Based Access Control

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.lakeformation.getPermissions({
    principal: aws_iam_role.workflow_role.arn,
    lfTagPolicy: {
        resourceType: "DATABASE",
        expressions: [
            {
                key: "Team",
                values: ["Sales"],
            },
            {
                key: "Environment",
                values: [
                    "Dev",
                    "Production",
                ],
            },
        ],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lakeformation.get_permissions(principal=aws_iam_role["workflow_role"]["arn"],
    lf_tag_policy=aws.lakeformation.GetPermissionsLfTagPolicyArgs(
        resource_type="DATABASE",
        expressions=[
            aws.lakeformation.GetPermissionsLfTagPolicyExpressionArgs(
                key="Team",
                values=["Sales"],
            ),
            aws.lakeformation.GetPermissionsLfTagPolicyExpressionArgs(
                key="Environment",
                values=[
                    "Dev",
                    "Production",
                ],
            ),
        ],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.LakeFormation.GetPermissions.Invoke(new()
    {
        Principal = aws_iam_role.Workflow_role.Arn,
        LfTagPolicy = new Aws.LakeFormation.Inputs.GetPermissionsLfTagPolicyInputArgs
        {
            ResourceType = "DATABASE",
            Expressions = new[]
            {
                new Aws.LakeFormation.Inputs.GetPermissionsLfTagPolicyExpressionInputArgs
                {
                    Key = "Team",
                    Values = new[]
                    {
                        "Sales",
                    },
                },
                new Aws.LakeFormation.Inputs.GetPermissionsLfTagPolicyExpressionInputArgs
                {
                    Key = "Environment",
                    Values = new[]
                    {
                        "Dev",
                        "Production",
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lakeformation.LookupPermissions(ctx, &lakeformation.LookupPermissionsArgs{
			Principal: aws_iam_role.Workflow_role.Arn,
			LfTagPolicy: lakeformation.GetPermissionsLfTagPolicy{
				ResourceType: "DATABASE",
				Expressions: []lakeformation.GetPermissionsLfTagPolicyExpression{
					{
						Key: "Team",
						Values: []string{
							"Sales",
						},
					},
					{
						Key: "Environment",
						Values: []string{
							"Dev",
							"Production",
						},
					},
				},
			},
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
import com.pulumi.aws.lakeformation.LakeformationFunctions;
import com.pulumi.aws.lakeformation.inputs.GetPermissionsArgs;
import com.pulumi.aws.lakeformation.inputs.GetPermissionsLfTagPolicyArgs;
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
        final var test = LakeformationFunctions.getPermissions(GetPermissionsArgs.builder()
            .principal(aws_iam_role.workflow_role().arn())
            .lfTagPolicy(GetPermissionsLfTagPolicyArgs.builder()
                .resourceType("DATABASE")
                .expressions(                
                    GetPermissionsLfTagPolicyExpressionArgs.builder()
                        .key("Team")
                        .values("Sales")
                        .build(),
                    GetPermissionsLfTagPolicyExpressionArgs.builder()
                        .key("Environment")
                        .values(                        
                            "Dev",
                            "Production")
                        .build())
                .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:lakeformation:getPermissions
      Arguments:
        principal: ${aws_iam_role.workflow_role.arn}
        lfTagPolicy:
          resourceType: DATABASE
          expressions:
            - key: Team
              values:
                - Sales
            - key: Environment
              values:
                - Dev
                - Production
```
{{% /example %}}
{{% /examples %}}