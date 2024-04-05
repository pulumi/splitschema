Provides an Amazon Connect User Hierarchy Group resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

> **NOTE:** The User Hierarchy Structure must be created before creating a User Hierarchy Group.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.UserHierarchyGroup("example", {
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    tags: {
        Name: "Example User Hierarchy Group",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.UserHierarchyGroup("example",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    tags={
        "Name": "Example User Hierarchy Group",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.UserHierarchyGroup("example", new()
    {
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        Tags = 
        {
            { "Name", "Example User Hierarchy Group" },
        },
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
		_, err := connect.NewUserHierarchyGroup(ctx, "example", &connect.UserHierarchyGroupArgs{
			InstanceId: pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example User Hierarchy Group"),
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
import com.pulumi.aws.connect.UserHierarchyGroup;
import com.pulumi.aws.connect.UserHierarchyGroupArgs;
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
        var example = new UserHierarchyGroup("example", UserHierarchyGroupArgs.builder()        
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .tags(Map.of("Name", "Example User Hierarchy Group"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:UserHierarchyGroup
    properties:
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
      tags:
        Name: Example User Hierarchy Group
```
{{% /example %}}
{{% example %}}
### With a parent group

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const parent = new aws.connect.UserHierarchyGroup("parent", {
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    tags: {
        Name: "Example User Hierarchy Group Parent",
    },
});
const child = new aws.connect.UserHierarchyGroup("child", {
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    parentGroupId: parent.hierarchyGroupId,
    tags: {
        Name: "Example User Hierarchy Group Child",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

parent = aws.connect.UserHierarchyGroup("parent",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    tags={
        "Name": "Example User Hierarchy Group Parent",
    })
child = aws.connect.UserHierarchyGroup("child",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    parent_group_id=parent.hierarchy_group_id,
    tags={
        "Name": "Example User Hierarchy Group Child",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var parent = new Aws.Connect.UserHierarchyGroup("parent", new()
    {
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        Tags = 
        {
            { "Name", "Example User Hierarchy Group Parent" },
        },
    });

    var child = new Aws.Connect.UserHierarchyGroup("child", new()
    {
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        ParentGroupId = parent.HierarchyGroupId,
        Tags = 
        {
            { "Name", "Example User Hierarchy Group Child" },
        },
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
		parent, err := connect.NewUserHierarchyGroup(ctx, "parent", &connect.UserHierarchyGroupArgs{
			InstanceId: pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example User Hierarchy Group Parent"),
			},
		})
		if err != nil {
			return err
		}
		_, err = connect.NewUserHierarchyGroup(ctx, "child", &connect.UserHierarchyGroupArgs{
			InstanceId:    pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
			ParentGroupId: parent.HierarchyGroupId,
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example User Hierarchy Group Child"),
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
import com.pulumi.aws.connect.UserHierarchyGroup;
import com.pulumi.aws.connect.UserHierarchyGroupArgs;
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
        var parent = new UserHierarchyGroup("parent", UserHierarchyGroupArgs.builder()        
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .tags(Map.of("Name", "Example User Hierarchy Group Parent"))
            .build());

        var child = new UserHierarchyGroup("child", UserHierarchyGroupArgs.builder()        
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .parentGroupId(parent.hierarchyGroupId())
            .tags(Map.of("Name", "Example User Hierarchy Group Child"))
            .build());

    }
}
```
```yaml
resources:
  parent:
    type: aws:connect:UserHierarchyGroup
    properties:
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
      tags:
        Name: Example User Hierarchy Group Parent
  child:
    type: aws:connect:UserHierarchyGroup
    properties:
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
      parentGroupId: ${parent.hierarchyGroupId}
      tags:
        Name: Example User Hierarchy Group Child
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect User Hierarchy Groups using the `instance_id` and `hierarchy_group_id` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:connect/userHierarchyGroup:UserHierarchyGroup example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
```
 