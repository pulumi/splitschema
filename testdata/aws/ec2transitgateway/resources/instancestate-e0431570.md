Provides an EC2 instance state resource. This allows managing an instance power state.

> **NOTE on Instance State Management:** AWS does not currently have an EC2 API operation to determine an instance has finished processing user data. As a result, this resource can interfere with user data processing. For example, this resource may stop an instance while the user data script is in mid run.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ubuntu = aws.ec2.getAmi({
    mostRecent: true,
    filters: [
        {
            name: "name",
            values: ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"],
        },
        {
            name: "virtualization-type",
            values: ["hvm"],
        },
    ],
    owners: ["099720109477"],
});
const testInstance = new aws.ec2.Instance("testInstance", {
    ami: ubuntu.then(ubuntu => ubuntu.id),
    instanceType: "t3.micro",
    tags: {
        Name: "HelloWorld",
    },
});
const testInstanceState = new aws.ec2transitgateway.InstanceState("testInstanceState", {
    instanceId: testInstance.id,
    state: "stopped",
});
```
```python
import pulumi
import pulumi_aws as aws

ubuntu = aws.ec2.get_ami(most_recent=True,
    filters=[
        aws.ec2.GetAmiFilterArgs(
            name="name",
            values=["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"],
        ),
        aws.ec2.GetAmiFilterArgs(
            name="virtualization-type",
            values=["hvm"],
        ),
    ],
    owners=["099720109477"])
test_instance = aws.ec2.Instance("testInstance",
    ami=ubuntu.id,
    instance_type="t3.micro",
    tags={
        "Name": "HelloWorld",
    })
test_instance_state = aws.ec2transitgateway.InstanceState("testInstanceState",
    instance_id=test_instance.id,
    state="stopped")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ubuntu = Aws.Ec2.GetAmi.Invoke(new()
    {
        MostRecent = true,
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*",
                },
            },
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "virtualization-type",
                Values = new[]
                {
                    "hvm",
                },
            },
        },
        Owners = new[]
        {
            "099720109477",
        },
    });

    var testInstance = new Aws.Ec2.Instance("testInstance", new()
    {
        Ami = ubuntu.Apply(getAmiResult => getAmiResult.Id),
        InstanceType = "t3.micro",
        Tags = 
        {
            { "Name", "HelloWorld" },
        },
    });

    var testInstanceState = new Aws.Ec2TransitGateway.InstanceState("testInstanceState", new()
    {
        InstanceId = testInstance.Id,
        State = "stopped",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ubuntu, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*",
					},
				},
				{
					Name: "virtualization-type",
					Values: []string{
						"hvm",
					},
				},
			},
			Owners: []string{
				"099720109477",
			},
		}, nil)
		if err != nil {
			return err
		}
		testInstance, err := ec2.NewInstance(ctx, "testInstance", &ec2.InstanceArgs{
			Ami:          *pulumi.String(ubuntu.Id),
			InstanceType: pulumi.String("t3.micro"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2transitgateway.NewInstanceState(ctx, "testInstanceState", &ec2transitgateway.InstanceStateArgs{
			InstanceId: testInstance.ID(),
			State:      pulumi.String("stopped"),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import com.pulumi.aws.ec2transitgateway.InstanceState;
import com.pulumi.aws.ec2transitgateway.InstanceStateArgs;
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
        final var ubuntu = Ec2Functions.getAmi(GetAmiArgs.builder()
            .mostRecent(true)
            .filters(            
                GetAmiFilterArgs.builder()
                    .name("name")
                    .values("ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*")
                    .build(),
                GetAmiFilterArgs.builder()
                    .name("virtualization-type")
                    .values("hvm")
                    .build())
            .owners("099720109477")
            .build());

        var testInstance = new Instance("testInstance", InstanceArgs.builder()        
            .ami(ubuntu.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceType("t3.micro")
            .tags(Map.of("Name", "HelloWorld"))
            .build());

        var testInstanceState = new InstanceState("testInstanceState", InstanceStateArgs.builder()        
            .instanceId(testInstance.id())
            .state("stopped")
            .build());

    }
}
```
```yaml
resources:
  testInstance:
    type: aws:ec2:Instance
    properties:
      ami: ${ubuntu.id}
      instanceType: t3.micro
      tags:
        Name: HelloWorld
  testInstanceState:
    type: aws:ec2transitgateway:InstanceState
    properties:
      instanceId: ${testInstance.id}
      state: stopped
variables:
  ubuntu:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        mostRecent: true
        filters:
          - name: name
            values:
              - ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*
          - name: virtualization-type
            values:
              - hvm
        owners:
          - '099720109477'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_instance_state` using the `instance_id` attribute. For example:

```sh
 $ pulumi import aws:ec2transitgateway/instanceState:InstanceState test i-02cae6557dfcf2f96
```
 