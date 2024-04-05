Resource for managing QuickSight Group Membership

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.GroupMembership("example", {
    groupName: "all-access-users",
    memberName: "john_smith",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.GroupMembership("example",
    group_name="all-access-users",
    member_name="john_smith")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.GroupMembership("example", new()
    {
        GroupName = "all-access-users",
        MemberName = "john_smith",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewGroupMembership(ctx, "example", &quicksight.GroupMembershipArgs{
			GroupName:  pulumi.String("all-access-users"),
			MemberName: pulumi.String("john_smith"),
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
import com.pulumi.aws.quicksight.GroupMembership;
import com.pulumi.aws.quicksight.GroupMembershipArgs;
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
        var example = new GroupMembership("example", GroupMembershipArgs.builder()        
            .groupName("all-access-users")
            .memberName("john_smith")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:GroupMembership
    properties:
      groupName: all-access-users
      memberName: john_smith
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import QuickSight Group membership using the AWS account ID, namespace, group name and member name separated by `/`. For example:

```sh
 $ pulumi import aws:quicksight/groupMembership:GroupMembership example 123456789123/default/all-access-users/john_smith
```
 