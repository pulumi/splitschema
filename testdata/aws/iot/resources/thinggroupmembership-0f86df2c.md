Adds an IoT Thing to an IoT Thing Group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iot.ThingGroupMembership("example", {
    overrideDynamicGroup: true,
    thingGroupName: "example-group",
    thingName: "example-thing",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iot.ThingGroupMembership("example",
    override_dynamic_group=True,
    thing_group_name="example-group",
    thing_name="example-thing")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Iot.ThingGroupMembership("example", new()
    {
        OverrideDynamicGroup = true,
        ThingGroupName = "example-group",
        ThingName = "example-thing",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iot.NewThingGroupMembership(ctx, "example", &iot.ThingGroupMembershipArgs{
			OverrideDynamicGroup: pulumi.Bool(true),
			ThingGroupName:       pulumi.String("example-group"),
			ThingName:            pulumi.String("example-thing"),
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
import com.pulumi.aws.iot.ThingGroupMembership;
import com.pulumi.aws.iot.ThingGroupMembershipArgs;
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
        var example = new ThingGroupMembership("example", ThingGroupMembershipArgs.builder()        
            .overrideDynamicGroup(true)
            .thingGroupName("example-group")
            .thingName("example-thing")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:iot:ThingGroupMembership
    properties:
      overrideDynamicGroup: true
      thingGroupName: example-group
      thingName: example-thing
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IoT Thing Group Membership using the thing group name and thing name. For example:

```sh
 $ pulumi import aws:iot/thingGroupMembership:ThingGroupMembership example thing_group_name/thing_name
```
 