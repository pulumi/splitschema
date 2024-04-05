Provides an RDS DB parameter group resource. Documentation of the available parameters for various RDS engines can be found at:

* [Aurora MySQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Reference.html)
* [Aurora PostgreSQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraPostgreSQL.Reference.html)
* [MariaDB Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MariaDB.Parameters.html)
* [Oracle Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ModifyInstance.Oracle.html#USER_ModifyInstance.Oracle.sqlnet)
* [PostgreSQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.html#Appendix.PostgreSQL.CommonDBATasks.Parameters)



> **NOTE:** After applying your changes, you may encounter a perpetual diff in your pulumi preview
output for a `parameter` whose `value` remains unchanged but whose `apply_method` is changing
(e.g., from `immediate` to `pending-reboot`, or `pending-reboot` to `immediate`). If only the
apply method of a parameter is changing, the AWS API will not register this change. To change
the `apply_method` of a parameter, its value must also change.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.rds.ParameterGroup("default", {
    family: "mysql5.6",
    parameters: [
        {
            name: "character_set_server",
            value: "utf8",
        },
        {
            name: "character_set_client",
            value: "utf8",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.ParameterGroup("default",
    family="mysql5.6",
    parameters=[
        aws.rds.ParameterGroupParameterArgs(
            name="character_set_server",
            value="utf8",
        ),
        aws.rds.ParameterGroupParameterArgs(
            name="character_set_client",
            value="utf8",
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Rds.ParameterGroup("default", new()
    {
        Family = "mysql5.6",
        Parameters = new[]
        {
            new Aws.Rds.Inputs.ParameterGroupParameterArgs
            {
                Name = "character_set_server",
                Value = "utf8",
            },
            new Aws.Rds.Inputs.ParameterGroupParameterArgs
            {
                Name = "character_set_client",
                Value = "utf8",
            },
        },
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
		_, err := rds.NewParameterGroup(ctx, "default", &rds.ParameterGroupArgs{
			Family: pulumi.String("mysql5.6"),
			Parameters: rds.ParameterGroupParameterArray{
				&rds.ParameterGroupParameterArgs{
					Name:  pulumi.String("character_set_server"),
					Value: pulumi.String("utf8"),
				},
				&rds.ParameterGroupParameterArgs{
					Name:  pulumi.String("character_set_client"),
					Value: pulumi.String("utf8"),
				},
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
import com.pulumi.aws.rds.ParameterGroup;
import com.pulumi.aws.rds.ParameterGroupArgs;
import com.pulumi.aws.rds.inputs.ParameterGroupParameterArgs;
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
        var default_ = new ParameterGroup("default", ParameterGroupArgs.builder()        
            .family("mysql5.6")
            .parameters(            
                ParameterGroupParameterArgs.builder()
                    .name("character_set_server")
                    .value("utf8")
                    .build(),
                ParameterGroupParameterArgs.builder()
                    .name("character_set_client")
                    .value("utf8")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:rds:ParameterGroup
    properties:
      family: mysql5.6
      parameters:
        - name: character_set_server
          value: utf8
        - name: character_set_client
          value: utf8
```
{{% /example %}}
{{% example %}}
### `create_before_destroy` Lifecycle Configuration

The `create_before_destroy`
lifecycle configuration is necessary for modifications that force re-creation of an existing,
in-use parameter group. This includes common situations like changing the group `name` or
bumping the `family` version during a major version upgrade. This configuration will prevent destruction
of the deposed parameter group while still in use by the database during upgrade.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleParameterGroup = new aws.rds.ParameterGroup("exampleParameterGroup", {
    family: "postgres13",
    parameters: [{
        name: "log_connections",
        value: "1",
    }],
});
const exampleInstance = new aws.rds.Instance("exampleInstance", {
    parameterGroupName: exampleParameterGroup.name,
    applyImmediately: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example_parameter_group = aws.rds.ParameterGroup("exampleParameterGroup",
    family="postgres13",
    parameters=[aws.rds.ParameterGroupParameterArgs(
        name="log_connections",
        value="1",
    )])
example_instance = aws.rds.Instance("exampleInstance",
    parameter_group_name=example_parameter_group.name,
    apply_immediately=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleParameterGroup = new Aws.Rds.ParameterGroup("exampleParameterGroup", new()
    {
        Family = "postgres13",
        Parameters = new[]
        {
            new Aws.Rds.Inputs.ParameterGroupParameterArgs
            {
                Name = "log_connections",
                Value = "1",
            },
        },
    });

    var exampleInstance = new Aws.Rds.Instance("exampleInstance", new()
    {
        ParameterGroupName = exampleParameterGroup.Name,
        ApplyImmediately = true,
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
		exampleParameterGroup, err := rds.NewParameterGroup(ctx, "exampleParameterGroup", &rds.ParameterGroupArgs{
			Family: pulumi.String("postgres13"),
			Parameters: rds.ParameterGroupParameterArray{
				&rds.ParameterGroupParameterArgs{
					Name:  pulumi.String("log_connections"),
					Value: pulumi.String("1"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = rds.NewInstance(ctx, "exampleInstance", &rds.InstanceArgs{
			ParameterGroupName: exampleParameterGroup.Name,
			ApplyImmediately:   pulumi.Bool(true),
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
import com.pulumi.aws.rds.ParameterGroup;
import com.pulumi.aws.rds.ParameterGroupArgs;
import com.pulumi.aws.rds.inputs.ParameterGroupParameterArgs;
import com.pulumi.aws.rds.Instance;
import com.pulumi.aws.rds.InstanceArgs;
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
        var exampleParameterGroup = new ParameterGroup("exampleParameterGroup", ParameterGroupArgs.builder()        
            .family("postgres13")
            .parameters(ParameterGroupParameterArgs.builder()
                .name("log_connections")
                .value("1")
                .build())
            .build());

        var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()        
            .parameterGroupName(exampleParameterGroup.name())
            .applyImmediately(true)
            .build());

    }
}
```
```yaml
resources:
  exampleParameterGroup:
    type: aws:rds:ParameterGroup
    properties:
      family: postgres13
      parameters:
        - name: log_connections
          value: '1'
  exampleInstance:
    type: aws:rds:Instance
    properties:
      # other attributes
      parameterGroupName: ${exampleParameterGroup.name}
      applyImmediately: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DB Parameter groups using the `name`. For example:

```sh
 $ pulumi import aws:rds/parameterGroup:ParameterGroup rds_pg rds-pg
```
 