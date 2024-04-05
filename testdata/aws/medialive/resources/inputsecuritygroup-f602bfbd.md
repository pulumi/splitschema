Resource for managing an AWS MediaLive InputSecurityGroup.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.medialive.InputSecurityGroup("example", {
    tags: {
        ENVIRONMENT: "prod",
    },
    whitelistRules: [{
        cidr: "10.0.0.8/32",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.medialive.InputSecurityGroup("example",
    tags={
        "ENVIRONMENT": "prod",
    },
    whitelist_rules=[aws.medialive.InputSecurityGroupWhitelistRuleArgs(
        cidr="10.0.0.8/32",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.MediaLive.InputSecurityGroup("example", new()
    {
        Tags = 
        {
            { "ENVIRONMENT", "prod" },
        },
        WhitelistRules = new[]
        {
            new Aws.MediaLive.Inputs.InputSecurityGroupWhitelistRuleArgs
            {
                Cidr = "10.0.0.8/32",
            },
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
		_, err := medialive.NewInputSecurityGroup(ctx, "example", &medialive.InputSecurityGroupArgs{
			Tags: pulumi.StringMap{
				"ENVIRONMENT": pulumi.String("prod"),
			},
			WhitelistRules: medialive.InputSecurityGroupWhitelistRuleArray{
				&medialive.InputSecurityGroupWhitelistRuleArgs{
					Cidr: pulumi.String("10.0.0.8/32"),
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
import com.pulumi.aws.medialive.InputSecurityGroup;
import com.pulumi.aws.medialive.InputSecurityGroupArgs;
import com.pulumi.aws.medialive.inputs.InputSecurityGroupWhitelistRuleArgs;
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
        var example = new InputSecurityGroup("example", InputSecurityGroupArgs.builder()        
            .tags(Map.of("ENVIRONMENT", "prod"))
            .whitelistRules(InputSecurityGroupWhitelistRuleArgs.builder()
                .cidr("10.0.0.8/32")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:medialive:InputSecurityGroup
    properties:
      tags:
        ENVIRONMENT: prod
      whitelistRules:
        - cidr: 10.0.0.8/32
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MediaLive InputSecurityGroup using the `id`. For example:

```sh
 $ pulumi import aws:medialive/inputSecurityGroup:InputSecurityGroup example 123456
```
 