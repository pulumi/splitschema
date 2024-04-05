Provides an SSM Parameter resource.

> **Note:** `overwrite` also makes it possible to overwrite an existing SSM Parameter that's not created by the provider before. This argument has been deprecated and will be removed in v6.0.0 of the provider. For more information on how this affects the behavior of this resource, see this issue comment.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ssm.Parameter("foo", {
    type: "String",
    value: "bar",
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ssm.Parameter("foo",
    type="String",
    value="bar")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ssm.Parameter("foo", new()
    {
        Type = "String",
        Value = "bar",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewParameter(ctx, "foo", &ssm.ParameterArgs{
			Type:  pulumi.String("String"),
			Value: pulumi.String("bar"),
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
import com.pulumi.aws.ssm.Parameter;
import com.pulumi.aws.ssm.ParameterArgs;
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
        var foo = new Parameter("foo", ParameterArgs.builder()        
            .type("String")
            .value("bar")
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ssm:Parameter
    properties:
      type: String
      value: bar
```
{{% /example %}}
{{% example %}}
### Encrypted string using default SSM KMS key

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.rds.Instance("default", {
    allocatedStorage: 10,
    storageType: "gp2",
    engine: "mysql",
    engineVersion: "5.7.16",
    instanceClass: "db.t2.micro",
    dbName: "mydb",
    username: "foo",
    password: _var.database_master_password,
    dbSubnetGroupName: "my_database_subnet_group",
    parameterGroupName: "default.mysql5.7",
});
const secret = new aws.ssm.Parameter("secret", {
    description: "The parameter description",
    type: "SecureString",
    value: _var.database_master_password,
    tags: {
        environment: "production",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.Instance("default",
    allocated_storage=10,
    storage_type="gp2",
    engine="mysql",
    engine_version="5.7.16",
    instance_class="db.t2.micro",
    db_name="mydb",
    username="foo",
    password=var["database_master_password"],
    db_subnet_group_name="my_database_subnet_group",
    parameter_group_name="default.mysql5.7")
secret = aws.ssm.Parameter("secret",
    description="The parameter description",
    type="SecureString",
    value=var["database_master_password"],
    tags={
        "environment": "production",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Rds.Instance("default", new()
    {
        AllocatedStorage = 10,
        StorageType = "gp2",
        Engine = "mysql",
        EngineVersion = "5.7.16",
        InstanceClass = "db.t2.micro",
        DbName = "mydb",
        Username = "foo",
        Password = @var.Database_master_password,
        DbSubnetGroupName = "my_database_subnet_group",
        ParameterGroupName = "default.mysql5.7",
    });

    var secret = new Aws.Ssm.Parameter("secret", new()
    {
        Description = "The parameter description",
        Type = "SecureString",
        Value = @var.Database_master_password,
        Tags = 
        {
            { "environment", "production" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewInstance(ctx, "default", &rds.InstanceArgs{
			AllocatedStorage:   pulumi.Int(10),
			StorageType:        pulumi.String("gp2"),
			Engine:             pulumi.String("mysql"),
			EngineVersion:      pulumi.String("5.7.16"),
			InstanceClass:      pulumi.String("db.t2.micro"),
			DbName:             pulumi.String("mydb"),
			Username:           pulumi.String("foo"),
			Password:           pulumi.Any(_var.Database_master_password),
			DbSubnetGroupName:  pulumi.String("my_database_subnet_group"),
			ParameterGroupName: pulumi.String("default.mysql5.7"),
		})
		if err != nil {
			return err
		}
		_, err = ssm.NewParameter(ctx, "secret", &ssm.ParameterArgs{
			Description: pulumi.String("The parameter description"),
			Type:        pulumi.String("SecureString"),
			Value:       pulumi.Any(_var.Database_master_password),
			Tags: pulumi.StringMap{
				"environment": pulumi.String("production"),
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
import com.pulumi.aws.rds.Instance;
import com.pulumi.aws.rds.InstanceArgs;
import com.pulumi.aws.ssm.Parameter;
import com.pulumi.aws.ssm.ParameterArgs;
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
        var default_ = new Instance("default", InstanceArgs.builder()        
            .allocatedStorage(10)
            .storageType("gp2")
            .engine("mysql")
            .engineVersion("5.7.16")
            .instanceClass("db.t2.micro")
            .dbName("mydb")
            .username("foo")
            .password(var_.database_master_password())
            .dbSubnetGroupName("my_database_subnet_group")
            .parameterGroupName("default.mysql5.7")
            .build());

        var secret = new Parameter("secret", ParameterArgs.builder()        
            .description("The parameter description")
            .type("SecureString")
            .value(var_.database_master_password())
            .tags(Map.of("environment", "production"))
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 10
      storageType: gp2
      engine: mysql
      engineVersion: 5.7.16
      instanceClass: db.t2.micro
      dbName: mydb
      username: foo
      password: ${var.database_master_password}
      dbSubnetGroupName: my_database_subnet_group
      parameterGroupName: default.mysql5.7
  secret:
    type: aws:ssm:Parameter
    properties:
      description: The parameter description
      type: SecureString
      value: ${var.database_master_password}
      tags:
        environment: production
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SSM Parameters using the parameter store `name`. For example:

```sh
 $ pulumi import aws:ssm/parameter:Parameter my_param /my_path/my_paramname
```
 