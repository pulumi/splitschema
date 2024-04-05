The CloudFormation Stack data source allows access to stack
outputs and other useful data including the template body.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const network = aws.cloudformation.getStack({
    name: "my-network-stack",
});
const web = new aws.ec2.Instance("web", {
    ami: "ami-abb07bcb",
    instanceType: "t2.micro",
    subnetId: network.then(network => network.outputs?.SubnetId),
    tags: {
        Name: "HelloWorld",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

network = aws.cloudformation.get_stack(name="my-network-stack")
web = aws.ec2.Instance("web",
    ami="ami-abb07bcb",
    instance_type="t2.micro",
    subnet_id=network.outputs["SubnetId"],
    tags={
        "Name": "HelloWorld",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var network = Aws.CloudFormation.GetStack.Invoke(new()
    {
        Name = "my-network-stack",
    });

    var web = new Aws.Ec2.Instance("web", new()
    {
        Ami = "ami-abb07bcb",
        InstanceType = "t2.micro",
        SubnetId = network.Apply(getStackResult => getStackResult.Outputs?.SubnetId),
        Tags = 
        {
            { "Name", "HelloWorld" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudformation"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		network, err := cloudformation.LookupStack(ctx, &cloudformation.LookupStackArgs{
			Name: "my-network-stack",
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
			Ami:          pulumi.String("ami-abb07bcb"),
			InstanceType: pulumi.String("t2.micro"),
			SubnetId:     *pulumi.String(network.Outputs.SubnetId),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
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
import com.pulumi.aws.cloudformation.CloudformationFunctions;
import com.pulumi.aws.cloudformation.inputs.GetStackArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
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
        final var network = CloudformationFunctions.getStack(GetStackArgs.builder()
            .name("my-network-stack")
            .build());

        var web = new Instance("web", InstanceArgs.builder()        
            .ami("ami-abb07bcb")
            .instanceType("t2.micro")
            .subnetId(network.applyValue(getStackResult -> getStackResult.outputs().SubnetId()))
            .tags(Map.of("Name", "HelloWorld"))
            .build());

    }
}
```
```yaml
resources:
  web:
    type: aws:ec2:Instance
    properties:
      ami: ami-abb07bcb
      instanceType: t2.micro
      subnetId: ${network.outputs.SubnetId}
      tags:
        Name: HelloWorld
variables:
  network:
    fn::invoke:
      Function: aws:cloudformation:getStack
      Arguments:
        name: my-network-stack
```
{{% /example %}}
{{% /examples %}}