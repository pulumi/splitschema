Provides an Amazon Connect User Hierarchy Structure resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.UserHierarchyStructure("example", {
    hierarchyStructure: {
        levelOne: {
            name: "levelone",
        },
    },
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.UserHierarchyStructure("example",
    hierarchy_structure=aws.connect.UserHierarchyStructureHierarchyStructureArgs(
        level_one=aws.connect.UserHierarchyStructureHierarchyStructureLevelOneArgs(
            name="levelone",
        ),
    ),
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.UserHierarchyStructure("example", new()
    {
        HierarchyStructure = new Aws.Connect.Inputs.UserHierarchyStructureHierarchyStructureArgs
        {
            LevelOne = new Aws.Connect.Inputs.UserHierarchyStructureHierarchyStructureLevelOneArgs
            {
                Name = "levelone",
            },
        },
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewUserHierarchyStructure(ctx, "example", &connect.UserHierarchyStructureArgs{
			HierarchyStructure: &connect.UserHierarchyStructureHierarchyStructureArgs{
				LevelOne: &connect.UserHierarchyStructureHierarchyStructureLevelOneArgs{
					Name: pulumi.String("levelone"),
				},
			},
			InstanceId: pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
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
import com.pulumi.aws.connect.UserHierarchyStructure;
import com.pulumi.aws.connect.UserHierarchyStructureArgs;
import com.pulumi.aws.connect.inputs.UserHierarchyStructureHierarchyStructureArgs;
import com.pulumi.aws.connect.inputs.UserHierarchyStructureHierarchyStructureLevelOneArgs;
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
        var example = new UserHierarchyStructure("example", UserHierarchyStructureArgs.builder()        
            .hierarchyStructure(UserHierarchyStructureHierarchyStructureArgs.builder()
                .levelOne(UserHierarchyStructureHierarchyStructureLevelOneArgs.builder()
                    .name("levelone")
                    .build())
                .build())
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:UserHierarchyStructure
    properties:
      hierarchyStructure:
        levelOne:
          name: levelone
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
```
{{% /example %}}
{{% example %}}
### With Five Levels

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.UserHierarchyStructure("example", {
    hierarchyStructure: {
        levelFive: {
            name: "levelfive",
        },
        levelFour: {
            name: "levelfour",
        },
        levelOne: {
            name: "levelone",
        },
        levelThree: {
            name: "levelthree",
        },
        levelTwo: {
            name: "leveltwo",
        },
    },
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.UserHierarchyStructure("example",
    hierarchy_structure=aws.connect.UserHierarchyStructureHierarchyStructureArgs(
        level_five=aws.connect.UserHierarchyStructureHierarchyStructureLevelFiveArgs(
            name="levelfive",
        ),
        level_four=aws.connect.UserHierarchyStructureHierarchyStructureLevelFourArgs(
            name="levelfour",
        ),
        level_one=aws.connect.UserHierarchyStructureHierarchyStructureLevelOneArgs(
            name="levelone",
        ),
        level_three=aws.connect.UserHierarchyStructureHierarchyStructureLevelThreeArgs(
            name="levelthree",
        ),
        level_two=aws.connect.UserHierarchyStructureHierarchyStructureLevelTwoArgs(
            name="leveltwo",
        ),
    ),
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.UserHierarchyStructure("example", new()
    {
        HierarchyStructure = new Aws.Connect.Inputs.UserHierarchyStructureHierarchyStructureArgs
        {
            LevelFive = new Aws.Connect.Inputs.UserHierarchyStructureHierarchyStructureLevelFiveArgs
            {
                Name = "levelfive",
            },
            LevelFour = new Aws.Connect.Inputs.UserHierarchyStructureHierarchyStructureLevelFourArgs
            {
                Name = "levelfour",
            },
            LevelOne = new Aws.Connect.Inputs.UserHierarchyStructureHierarchyStructureLevelOneArgs
            {
                Name = "levelone",
            },
            LevelThree = new Aws.Connect.Inputs.UserHierarchyStructureHierarchyStructureLevelThreeArgs
            {
                Name = "levelthree",
            },
            LevelTwo = new Aws.Connect.Inputs.UserHierarchyStructureHierarchyStructureLevelTwoArgs
            {
                Name = "leveltwo",
            },
        },
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewUserHierarchyStructure(ctx, "example", &connect.UserHierarchyStructureArgs{
			HierarchyStructure: &connect.UserHierarchyStructureHierarchyStructureArgs{
				LevelFive: &connect.UserHierarchyStructureHierarchyStructureLevelFiveArgs{
					Name: pulumi.String("levelfive"),
				},
				LevelFour: &connect.UserHierarchyStructureHierarchyStructureLevelFourArgs{
					Name: pulumi.String("levelfour"),
				},
				LevelOne: &connect.UserHierarchyStructureHierarchyStructureLevelOneArgs{
					Name: pulumi.String("levelone"),
				},
				LevelThree: &connect.UserHierarchyStructureHierarchyStructureLevelThreeArgs{
					Name: pulumi.String("levelthree"),
				},
				LevelTwo: &connect.UserHierarchyStructureHierarchyStructureLevelTwoArgs{
					Name: pulumi.String("leveltwo"),
				},
			},
			InstanceId: pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
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
import com.pulumi.aws.connect.UserHierarchyStructure;
import com.pulumi.aws.connect.UserHierarchyStructureArgs;
import com.pulumi.aws.connect.inputs.UserHierarchyStructureHierarchyStructureArgs;
import com.pulumi.aws.connect.inputs.UserHierarchyStructureHierarchyStructureLevelFiveArgs;
import com.pulumi.aws.connect.inputs.UserHierarchyStructureHierarchyStructureLevelFourArgs;
import com.pulumi.aws.connect.inputs.UserHierarchyStructureHierarchyStructureLevelOneArgs;
import com.pulumi.aws.connect.inputs.UserHierarchyStructureHierarchyStructureLevelThreeArgs;
import com.pulumi.aws.connect.inputs.UserHierarchyStructureHierarchyStructureLevelTwoArgs;
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
        var example = new UserHierarchyStructure("example", UserHierarchyStructureArgs.builder()        
            .hierarchyStructure(UserHierarchyStructureHierarchyStructureArgs.builder()
                .levelFive(UserHierarchyStructureHierarchyStructureLevelFiveArgs.builder()
                    .name("levelfive")
                    .build())
                .levelFour(UserHierarchyStructureHierarchyStructureLevelFourArgs.builder()
                    .name("levelfour")
                    .build())
                .levelOne(UserHierarchyStructureHierarchyStructureLevelOneArgs.builder()
                    .name("levelone")
                    .build())
                .levelThree(UserHierarchyStructureHierarchyStructureLevelThreeArgs.builder()
                    .name("levelthree")
                    .build())
                .levelTwo(UserHierarchyStructureHierarchyStructureLevelTwoArgs.builder()
                    .name("leveltwo")
                    .build())
                .build())
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:UserHierarchyStructure
    properties:
      hierarchyStructure:
        levelFive:
          name: levelfive
        levelFour:
          name: levelfour
        levelOne:
          name: levelone
        levelThree:
          name: levelthree
        levelTwo:
          name: leveltwo
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect User Hierarchy Structures using the `instance_id`. For example:

```sh
 $ pulumi import aws:connect/userHierarchyStructure:UserHierarchyStructure example f1288a1f-6193-445a-b47e-af739b2
```
 