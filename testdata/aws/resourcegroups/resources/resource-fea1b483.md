Resource for managing an AWS Resource Groups Resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDedicatedHost = new aws.ec2.DedicatedHost("exampleDedicatedHost", {
    instanceFamily: "t3",
    availabilityZone: "us-east-1a",
    hostRecovery: "off",
    autoPlacement: "on",
});
const exampleGroup = new aws.resourcegroups.Group("exampleGroup", {});
const exampleResource = new aws.resourcegroups.Resource("exampleResource", {
    groupArn: exampleGroup.arn,
    resourceArn: exampleDedicatedHost.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example_dedicated_host = aws.ec2.DedicatedHost("exampleDedicatedHost",
    instance_family="t3",
    availability_zone="us-east-1a",
    host_recovery="off",
    auto_placement="on")
example_group = aws.resourcegroups.Group("exampleGroup")
example_resource = aws.resourcegroups.Resource("exampleResource",
    group_arn=example_group.arn,
    resource_arn=example_dedicated_host.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleDedicatedHost = new Aws.Ec2.DedicatedHost("exampleDedicatedHost", new()
    {
        InstanceFamily = "t3",
        AvailabilityZone = "us-east-1a",
        HostRecovery = "off",
        AutoPlacement = "on",
    });

    var exampleGroup = new Aws.ResourceGroups.Group("exampleGroup");

    var exampleResource = new Aws.ResourceGroups.Resource("exampleResource", new()
    {
        GroupArn = exampleGroup.Arn,
        ResourceArn = exampleDedicatedHost.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/resourcegroups"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDedicatedHost, err := ec2.NewDedicatedHost(ctx, "exampleDedicatedHost", &ec2.DedicatedHostArgs{
			InstanceFamily:   pulumi.String("t3"),
			AvailabilityZone: pulumi.String("us-east-1a"),
			HostRecovery:     pulumi.String("off"),
			AutoPlacement:    pulumi.String("on"),
		})
		if err != nil {
			return err
		}
		exampleGroup, err := resourcegroups.NewGroup(ctx, "exampleGroup", nil)
		if err != nil {
			return err
		}
		_, err = resourcegroups.NewResource(ctx, "exampleResource", &resourcegroups.ResourceArgs{
			GroupArn:    exampleGroup.Arn,
			ResourceArn: exampleDedicatedHost.Arn,
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
import com.pulumi.aws.ec2.DedicatedHost;
import com.pulumi.aws.ec2.DedicatedHostArgs;
import com.pulumi.aws.resourcegroups.Group;
import com.pulumi.aws.resourcegroups.Resource;
import com.pulumi.aws.resourcegroups.ResourceArgs;
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
        var exampleDedicatedHost = new DedicatedHost("exampleDedicatedHost", DedicatedHostArgs.builder()        
            .instanceFamily("t3")
            .availabilityZone("us-east-1a")
            .hostRecovery("off")
            .autoPlacement("on")
            .build());

        var exampleGroup = new Group("exampleGroup");

        var exampleResource = new Resource("exampleResource", ResourceArgs.builder()        
            .groupArn(exampleGroup.arn())
            .resourceArn(exampleDedicatedHost.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleDedicatedHost:
    type: aws:ec2:DedicatedHost
    properties:
      instanceFamily: t3
      availabilityZone: us-east-1a
      hostRecovery: off
      autoPlacement: on
  exampleGroup:
    type: aws:resourcegroups:Group
  exampleResource:
    type: aws:resourcegroups:Resource
    properties:
      groupArn: ${exampleGroup.arn}
      resourceArn: ${exampleDedicatedHost.arn}
```
{{% /example %}}
{{% /examples %}}