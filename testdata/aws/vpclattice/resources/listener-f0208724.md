Resource for managing an AWS VPC Lattice Listener.

{{% examples %}}
## Example Usage
{{% example %}}
### Forward action

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.vpclattice.Service("test", {});
const exampleTargetGroup = new aws.vpclattice.TargetGroup("exampleTargetGroup", {
    type: "INSTANCE",
    config: {
        port: 80,
        protocol: "HTTP",
        vpcIdentifier: aws_vpc.test.id,
    },
});
const exampleListener = new aws.vpclattice.Listener("exampleListener", {
    protocol: "HTTP",
    serviceIdentifier: aws_vpclattice_service.example.id,
    defaultAction: {
        forwards: [{
            targetGroups: [{
                targetGroupIdentifier: exampleTargetGroup.id,
            }],
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.vpclattice.Service("test")
example_target_group = aws.vpclattice.TargetGroup("exampleTargetGroup",
    type="INSTANCE",
    config=aws.vpclattice.TargetGroupConfigArgs(
        port=80,
        protocol="HTTP",
        vpc_identifier=aws_vpc["test"]["id"],
    ))
example_listener = aws.vpclattice.Listener("exampleListener",
    protocol="HTTP",
    service_identifier=aws_vpclattice_service["example"]["id"],
    default_action=aws.vpclattice.ListenerDefaultActionArgs(
        forwards=[aws.vpclattice.ListenerDefaultActionForwardArgs(
            target_groups=[aws.vpclattice.ListenerDefaultActionForwardTargetGroupArgs(
                target_group_identifier=example_target_group.id,
            )],
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.VpcLattice.Service("test");

    var exampleTargetGroup = new Aws.VpcLattice.TargetGroup("exampleTargetGroup", new()
    {
        Type = "INSTANCE",
        Config = new Aws.VpcLattice.Inputs.TargetGroupConfigArgs
        {
            Port = 80,
            Protocol = "HTTP",
            VpcIdentifier = aws_vpc.Test.Id,
        },
    });

    var exampleListener = new Aws.VpcLattice.Listener("exampleListener", new()
    {
        Protocol = "HTTP",
        ServiceIdentifier = aws_vpclattice_service.Example.Id,
        DefaultAction = new Aws.VpcLattice.Inputs.ListenerDefaultActionArgs
        {
            Forwards = new[]
            {
                new Aws.VpcLattice.Inputs.ListenerDefaultActionForwardArgs
                {
                    TargetGroups = new[]
                    {
                        new Aws.VpcLattice.Inputs.ListenerDefaultActionForwardTargetGroupArgs
                        {
                            TargetGroupIdentifier = exampleTargetGroup.Id,
                        },
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.NewService(ctx, "test", nil)
		if err != nil {
			return err
		}
		exampleTargetGroup, err := vpclattice.NewTargetGroup(ctx, "exampleTargetGroup", &vpclattice.TargetGroupArgs{
			Type: pulumi.String("INSTANCE"),
			Config: &vpclattice.TargetGroupConfigArgs{
				Port:          pulumi.Int(80),
				Protocol:      pulumi.String("HTTP"),
				VpcIdentifier: pulumi.Any(aws_vpc.Test.Id),
			},
		})
		if err != nil {
			return err
		}
		_, err = vpclattice.NewListener(ctx, "exampleListener", &vpclattice.ListenerArgs{
			Protocol:          pulumi.String("HTTP"),
			ServiceIdentifier: pulumi.Any(aws_vpclattice_service.Example.Id),
			DefaultAction: &vpclattice.ListenerDefaultActionArgs{
				Forwards: vpclattice.ListenerDefaultActionForwardArray{
					&vpclattice.ListenerDefaultActionForwardArgs{
						TargetGroups: vpclattice.ListenerDefaultActionForwardTargetGroupArray{
							&vpclattice.ListenerDefaultActionForwardTargetGroupArgs{
								TargetGroupIdentifier: exampleTargetGroup.ID(),
							},
						},
					},
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
import com.pulumi.aws.vpclattice.Service;
import com.pulumi.aws.vpclattice.TargetGroup;
import com.pulumi.aws.vpclattice.TargetGroupArgs;
import com.pulumi.aws.vpclattice.inputs.TargetGroupConfigArgs;
import com.pulumi.aws.vpclattice.Listener;
import com.pulumi.aws.vpclattice.ListenerArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerDefaultActionArgs;
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
        var test = new Service("test");

        var exampleTargetGroup = new TargetGroup("exampleTargetGroup", TargetGroupArgs.builder()        
            .type("INSTANCE")
            .config(TargetGroupConfigArgs.builder()
                .port(80)
                .protocol("HTTP")
                .vpcIdentifier(aws_vpc.test().id())
                .build())
            .build());

        var exampleListener = new Listener("exampleListener", ListenerArgs.builder()        
            .protocol("HTTP")
            .serviceIdentifier(aws_vpclattice_service.example().id())
            .defaultAction(ListenerDefaultActionArgs.builder()
                .forwards(ListenerDefaultActionForwardArgs.builder()
                    .targetGroups(ListenerDefaultActionForwardTargetGroupArgs.builder()
                        .targetGroupIdentifier(exampleTargetGroup.id())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:vpclattice:Service
  exampleTargetGroup:
    type: aws:vpclattice:TargetGroup
    properties:
      type: INSTANCE
      config:
        port: 80
        protocol: HTTP
        vpcIdentifier: ${aws_vpc.test.id}
  exampleListener:
    type: aws:vpclattice:Listener
    properties:
      protocol: HTTP
      serviceIdentifier: ${aws_vpclattice_service.example.id}
      defaultAction:
        forwards:
          - targetGroups:
              - targetGroupIdentifier: ${exampleTargetGroup.id}
```
{{% /example %}}
{{% example %}}
### Forward action with weighted target groups

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.vpclattice.Service("test", {});
const example1 = new aws.vpclattice.TargetGroup("example1", {
    type: "INSTANCE",
    config: {
        port: 80,
        protocol: "HTTP",
        vpcIdentifier: aws_vpc.test.id,
    },
});
const example2 = new aws.vpclattice.TargetGroup("example2", {
    type: "INSTANCE",
    config: {
        port: 8080,
        protocol: "HTTP",
        vpcIdentifier: aws_vpc.test.id,
    },
});
const example = new aws.vpclattice.Listener("example", {
    protocol: "HTTP",
    serviceIdentifier: aws_vpclattice_service.example.id,
    defaultAction: {
        forwards: [{
            targetGroups: [
                {
                    targetGroupIdentifier: example1.id,
                    weight: 80,
                },
                {
                    targetGroupIdentifier: example2.id,
                    weight: 20,
                },
            ],
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.vpclattice.Service("test")
example1 = aws.vpclattice.TargetGroup("example1",
    type="INSTANCE",
    config=aws.vpclattice.TargetGroupConfigArgs(
        port=80,
        protocol="HTTP",
        vpc_identifier=aws_vpc["test"]["id"],
    ))
example2 = aws.vpclattice.TargetGroup("example2",
    type="INSTANCE",
    config=aws.vpclattice.TargetGroupConfigArgs(
        port=8080,
        protocol="HTTP",
        vpc_identifier=aws_vpc["test"]["id"],
    ))
example = aws.vpclattice.Listener("example",
    protocol="HTTP",
    service_identifier=aws_vpclattice_service["example"]["id"],
    default_action=aws.vpclattice.ListenerDefaultActionArgs(
        forwards=[aws.vpclattice.ListenerDefaultActionForwardArgs(
            target_groups=[
                aws.vpclattice.ListenerDefaultActionForwardTargetGroupArgs(
                    target_group_identifier=example1.id,
                    weight=80,
                ),
                aws.vpclattice.ListenerDefaultActionForwardTargetGroupArgs(
                    target_group_identifier=example2.id,
                    weight=20,
                ),
            ],
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.VpcLattice.Service("test");

    var example1 = new Aws.VpcLattice.TargetGroup("example1", new()
    {
        Type = "INSTANCE",
        Config = new Aws.VpcLattice.Inputs.TargetGroupConfigArgs
        {
            Port = 80,
            Protocol = "HTTP",
            VpcIdentifier = aws_vpc.Test.Id,
        },
    });

    var example2 = new Aws.VpcLattice.TargetGroup("example2", new()
    {
        Type = "INSTANCE",
        Config = new Aws.VpcLattice.Inputs.TargetGroupConfigArgs
        {
            Port = 8080,
            Protocol = "HTTP",
            VpcIdentifier = aws_vpc.Test.Id,
        },
    });

    var example = new Aws.VpcLattice.Listener("example", new()
    {
        Protocol = "HTTP",
        ServiceIdentifier = aws_vpclattice_service.Example.Id,
        DefaultAction = new Aws.VpcLattice.Inputs.ListenerDefaultActionArgs
        {
            Forwards = new[]
            {
                new Aws.VpcLattice.Inputs.ListenerDefaultActionForwardArgs
                {
                    TargetGroups = new[]
                    {
                        new Aws.VpcLattice.Inputs.ListenerDefaultActionForwardTargetGroupArgs
                        {
                            TargetGroupIdentifier = example1.Id,
                            Weight = 80,
                        },
                        new Aws.VpcLattice.Inputs.ListenerDefaultActionForwardTargetGroupArgs
                        {
                            TargetGroupIdentifier = example2.Id,
                            Weight = 20,
                        },
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.NewService(ctx, "test", nil)
		if err != nil {
			return err
		}
		example1, err := vpclattice.NewTargetGroup(ctx, "example1", &vpclattice.TargetGroupArgs{
			Type: pulumi.String("INSTANCE"),
			Config: &vpclattice.TargetGroupConfigArgs{
				Port:          pulumi.Int(80),
				Protocol:      pulumi.String("HTTP"),
				VpcIdentifier: pulumi.Any(aws_vpc.Test.Id),
			},
		})
		if err != nil {
			return err
		}
		example2, err := vpclattice.NewTargetGroup(ctx, "example2", &vpclattice.TargetGroupArgs{
			Type: pulumi.String("INSTANCE"),
			Config: &vpclattice.TargetGroupConfigArgs{
				Port:          pulumi.Int(8080),
				Protocol:      pulumi.String("HTTP"),
				VpcIdentifier: pulumi.Any(aws_vpc.Test.Id),
			},
		})
		if err != nil {
			return err
		}
		_, err = vpclattice.NewListener(ctx, "example", &vpclattice.ListenerArgs{
			Protocol:          pulumi.String("HTTP"),
			ServiceIdentifier: pulumi.Any(aws_vpclattice_service.Example.Id),
			DefaultAction: &vpclattice.ListenerDefaultActionArgs{
				Forwards: vpclattice.ListenerDefaultActionForwardArray{
					&vpclattice.ListenerDefaultActionForwardArgs{
						TargetGroups: vpclattice.ListenerDefaultActionForwardTargetGroupArray{
							&vpclattice.ListenerDefaultActionForwardTargetGroupArgs{
								TargetGroupIdentifier: example1.ID(),
								Weight:                pulumi.Int(80),
							},
							&vpclattice.ListenerDefaultActionForwardTargetGroupArgs{
								TargetGroupIdentifier: example2.ID(),
								Weight:                pulumi.Int(20),
							},
						},
					},
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
import com.pulumi.aws.vpclattice.Service;
import com.pulumi.aws.vpclattice.TargetGroup;
import com.pulumi.aws.vpclattice.TargetGroupArgs;
import com.pulumi.aws.vpclattice.inputs.TargetGroupConfigArgs;
import com.pulumi.aws.vpclattice.Listener;
import com.pulumi.aws.vpclattice.ListenerArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerDefaultActionArgs;
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
        var test = new Service("test");

        var example1 = new TargetGroup("example1", TargetGroupArgs.builder()        
            .type("INSTANCE")
            .config(TargetGroupConfigArgs.builder()
                .port(80)
                .protocol("HTTP")
                .vpcIdentifier(aws_vpc.test().id())
                .build())
            .build());

        var example2 = new TargetGroup("example2", TargetGroupArgs.builder()        
            .type("INSTANCE")
            .config(TargetGroupConfigArgs.builder()
                .port(8080)
                .protocol("HTTP")
                .vpcIdentifier(aws_vpc.test().id())
                .build())
            .build());

        var example = new Listener("example", ListenerArgs.builder()        
            .protocol("HTTP")
            .serviceIdentifier(aws_vpclattice_service.example().id())
            .defaultAction(ListenerDefaultActionArgs.builder()
                .forwards(ListenerDefaultActionForwardArgs.builder()
                    .targetGroups(                    
                        ListenerDefaultActionForwardTargetGroupArgs.builder()
                            .targetGroupIdentifier(example1.id())
                            .weight(80)
                            .build(),
                        ListenerDefaultActionForwardTargetGroupArgs.builder()
                            .targetGroupIdentifier(example2.id())
                            .weight(20)
                            .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:vpclattice:Service
  example1:
    type: aws:vpclattice:TargetGroup
    properties:
      type: INSTANCE
      config:
        port: 80
        protocol: HTTP
        vpcIdentifier: ${aws_vpc.test.id}
  example2:
    type: aws:vpclattice:TargetGroup
    properties:
      type: INSTANCE
      config:
        port: 8080
        protocol: HTTP
        vpcIdentifier: ${aws_vpc.test.id}
  example:
    type: aws:vpclattice:Listener
    properties:
      protocol: HTTP
      serviceIdentifier: ${aws_vpclattice_service.example.id}
      defaultAction:
        forwards:
          - targetGroups:
              - targetGroupIdentifier: ${example1.id}
                weight: 80
              - targetGroupIdentifier: ${example2.id}
                weight: 20
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Lattice Listener using the `listener_id` of the listener and the `id` of the VPC Lattice service combined with a `/` character. For example:

```sh
 $ pulumi import aws:vpclattice/listener:Listener example svc-1a2b3c4d/listener-987654321
```
 