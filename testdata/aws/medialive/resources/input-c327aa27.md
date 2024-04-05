Resource for managing an AWS MediaLive Input.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleInputSecurityGroup = new aws.medialive.InputSecurityGroup("exampleInputSecurityGroup", {
    whitelistRules: [{
        cidr: "10.0.0.8/32",
    }],
    tags: {
        ENVIRONMENT: "prod",
    },
});
const exampleInput = new aws.medialive.Input("exampleInput", {
    inputSecurityGroups: [exampleInputSecurityGroup.id],
    type: "UDP_PUSH",
    tags: {
        ENVIRONMENT: "prod",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_input_security_group = aws.medialive.InputSecurityGroup("exampleInputSecurityGroup",
    whitelist_rules=[aws.medialive.InputSecurityGroupWhitelistRuleArgs(
        cidr="10.0.0.8/32",
    )],
    tags={
        "ENVIRONMENT": "prod",
    })
example_input = aws.medialive.Input("exampleInput",
    input_security_groups=[example_input_security_group.id],
    type="UDP_PUSH",
    tags={
        "ENVIRONMENT": "prod",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleInputSecurityGroup = new Aws.MediaLive.InputSecurityGroup("exampleInputSecurityGroup", new()
    {
        WhitelistRules = new[]
        {
            new Aws.MediaLive.Inputs.InputSecurityGroupWhitelistRuleArgs
            {
                Cidr = "10.0.0.8/32",
            },
        },
        Tags = 
        {
            { "ENVIRONMENT", "prod" },
        },
    });

    var exampleInput = new Aws.MediaLive.Input("exampleInput", new()
    {
        InputSecurityGroups = new[]
        {
            exampleInputSecurityGroup.Id,
        },
        Type = "UDP_PUSH",
        Tags = 
        {
            { "ENVIRONMENT", "prod" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/medialive"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleInputSecurityGroup, err := medialive.NewInputSecurityGroup(ctx, "exampleInputSecurityGroup", &medialive.InputSecurityGroupArgs{
			WhitelistRules: medialive.InputSecurityGroupWhitelistRuleArray{
				&medialive.InputSecurityGroupWhitelistRuleArgs{
					Cidr: pulumi.String("10.0.0.8/32"),
				},
			},
			Tags: pulumi.StringMap{
				"ENVIRONMENT": pulumi.String("prod"),
			},
		})
		if err != nil {
			return err
		}
		_, err = medialive.NewInput(ctx, "exampleInput", &medialive.InputArgs{
			InputSecurityGroups: pulumi.StringArray{
				exampleInputSecurityGroup.ID(),
			},
			Type: pulumi.String("UDP_PUSH"),
			Tags: pulumi.StringMap{
				"ENVIRONMENT": pulumi.String("prod"),
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
import com.pulumi.aws.medialive.InputSecurityGroup;
import com.pulumi.aws.medialive.InputSecurityGroupArgs;
import com.pulumi.aws.medialive.inputs.InputSecurityGroupWhitelistRuleArgs;
import com.pulumi.aws.medialive.Input;
import com.pulumi.aws.medialive.InputArgs;
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
        var exampleInputSecurityGroup = new InputSecurityGroup("exampleInputSecurityGroup", InputSecurityGroupArgs.builder()        
            .whitelistRules(InputSecurityGroupWhitelistRuleArgs.builder()
                .cidr("10.0.0.8/32")
                .build())
            .tags(Map.of("ENVIRONMENT", "prod"))
            .build());

        var exampleInput = new Input("exampleInput", InputArgs.builder()        
            .inputSecurityGroups(exampleInputSecurityGroup.id())
            .type("UDP_PUSH")
            .tags(Map.of("ENVIRONMENT", "prod"))
            .build());

    }
}
```
```yaml
resources:
  exampleInputSecurityGroup:
    type: aws:medialive:InputSecurityGroup
    properties:
      whitelistRules:
        - cidr: 10.0.0.8/32
      tags:
        ENVIRONMENT: prod
  exampleInput:
    type: aws:medialive:Input
    properties:
      inputSecurityGroups:
        - ${exampleInputSecurityGroup.id}
      type: UDP_PUSH
      tags:
        ENVIRONMENT: prod
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MediaLive Input using the `id`. For example:

```sh
 $ pulumi import aws:medialive/input:Input example 12345678
```
 