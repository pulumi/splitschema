Provides a Step Function State Machine Alias.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sfnAlias = new aws.sfn.Alias("sfnAlias", {routingConfigurations: [{
    stateMachineVersionArn: aws_sfn_state_machine.sfn_test.state_machine_version_arn,
    weight: 100,
}]});
const mySfnAlias = new aws.sfn.Alias("mySfnAlias", {routingConfigurations: [
    {
        stateMachineVersionArn: "arn:aws:states:us-east-1:12345:stateMachine:demo:3",
        weight: 50,
    },
    {
        stateMachineVersionArn: "arn:aws:states:us-east-1:12345:stateMachine:demo:2",
        weight: 50,
    },
]});
```
```python
import pulumi
import pulumi_aws as aws

sfn_alias = aws.sfn.Alias("sfnAlias", routing_configurations=[aws.sfn.AliasRoutingConfigurationArgs(
    state_machine_version_arn=aws_sfn_state_machine["sfn_test"]["state_machine_version_arn"],
    weight=100,
)])
my_sfn_alias = aws.sfn.Alias("mySfnAlias", routing_configurations=[
    aws.sfn.AliasRoutingConfigurationArgs(
        state_machine_version_arn="arn:aws:states:us-east-1:12345:stateMachine:demo:3",
        weight=50,
    ),
    aws.sfn.AliasRoutingConfigurationArgs(
        state_machine_version_arn="arn:aws:states:us-east-1:12345:stateMachine:demo:2",
        weight=50,
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
    var sfnAlias = new Aws.Sfn.Alias("sfnAlias", new()
    {
        RoutingConfigurations = new[]
        {
            new Aws.Sfn.Inputs.AliasRoutingConfigurationArgs
            {
                StateMachineVersionArn = aws_sfn_state_machine.Sfn_test.State_machine_version_arn,
                Weight = 100,
            },
        },
    });

    var mySfnAlias = new Aws.Sfn.Alias("mySfnAlias", new()
    {
        RoutingConfigurations = new[]
        {
            new Aws.Sfn.Inputs.AliasRoutingConfigurationArgs
            {
                StateMachineVersionArn = "arn:aws:states:us-east-1:12345:stateMachine:demo:3",
                Weight = 50,
            },
            new Aws.Sfn.Inputs.AliasRoutingConfigurationArgs
            {
                StateMachineVersionArn = "arn:aws:states:us-east-1:12345:stateMachine:demo:2",
                Weight = 50,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sfn"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sfn.NewAlias(ctx, "sfnAlias", &sfn.AliasArgs{
			RoutingConfigurations: sfn.AliasRoutingConfigurationArray{
				&sfn.AliasRoutingConfigurationArgs{
					StateMachineVersionArn: pulumi.Any(aws_sfn_state_machine.Sfn_test.State_machine_version_arn),
					Weight:                 pulumi.Int(100),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = sfn.NewAlias(ctx, "mySfnAlias", &sfn.AliasArgs{
			RoutingConfigurations: sfn.AliasRoutingConfigurationArray{
				&sfn.AliasRoutingConfigurationArgs{
					StateMachineVersionArn: pulumi.String("arn:aws:states:us-east-1:12345:stateMachine:demo:3"),
					Weight:                 pulumi.Int(50),
				},
				&sfn.AliasRoutingConfigurationArgs{
					StateMachineVersionArn: pulumi.String("arn:aws:states:us-east-1:12345:stateMachine:demo:2"),
					Weight:                 pulumi.Int(50),
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
import com.pulumi.aws.sfn.Alias;
import com.pulumi.aws.sfn.AliasArgs;
import com.pulumi.aws.sfn.inputs.AliasRoutingConfigurationArgs;
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
        var sfnAlias = new Alias("sfnAlias", AliasArgs.builder()        
            .routingConfigurations(AliasRoutingConfigurationArgs.builder()
                .stateMachineVersionArn(aws_sfn_state_machine.sfn_test().state_machine_version_arn())
                .weight(100)
                .build())
            .build());

        var mySfnAlias = new Alias("mySfnAlias", AliasArgs.builder()        
            .routingConfigurations(            
                AliasRoutingConfigurationArgs.builder()
                    .stateMachineVersionArn("arn:aws:states:us-east-1:12345:stateMachine:demo:3")
                    .weight(50)
                    .build(),
                AliasRoutingConfigurationArgs.builder()
                    .stateMachineVersionArn("arn:aws:states:us-east-1:12345:stateMachine:demo:2")
                    .weight(50)
                    .build())
            .build());

    }
}
```
```yaml
resources:
  sfnAlias:
    type: aws:sfn:Alias
    properties:
      routingConfigurations:
        - stateMachineVersionArn: ${aws_sfn_state_machine.sfn_test.state_machine_version_arn}
          weight: 100
  mySfnAlias:
    type: aws:sfn:Alias
    properties:
      routingConfigurations:
        - stateMachineVersionArn: arn:aws:states:us-east-1:12345:stateMachine:demo:3
          weight: 50
        - stateMachineVersionArn: arn:aws:states:us-east-1:12345:stateMachine:demo:2
          weight: 50
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SFN (Step Functions) Alias using the `arn`. For example:

```sh
 $ pulumi import aws:sfn/alias:Alias foo arn:aws:states:us-east-1:123456789098:stateMachine:myStateMachine:foo
```
 