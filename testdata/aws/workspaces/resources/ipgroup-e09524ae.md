Provides an IP access control group in AWS WorkSpaces Service

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const contractors = new aws.workspaces.IpGroup("contractors", {
    description: "Contractors IP access control group",
    rules: [
        {
            description: "NY",
            source: "150.24.14.0/24",
        },
        {
            description: "LA",
            source: "125.191.14.85/32",
        },
        {
            description: "STL",
            source: "44.98.100.0/24",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

contractors = aws.workspaces.IpGroup("contractors",
    description="Contractors IP access control group",
    rules=[
        aws.workspaces.IpGroupRuleArgs(
            description="NY",
            source="150.24.14.0/24",
        ),
        aws.workspaces.IpGroupRuleArgs(
            description="LA",
            source="125.191.14.85/32",
        ),
        aws.workspaces.IpGroupRuleArgs(
            description="STL",
            source="44.98.100.0/24",
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
    var contractors = new Aws.Workspaces.IpGroup("contractors", new()
    {
        Description = "Contractors IP access control group",
        Rules = new[]
        {
            new Aws.Workspaces.Inputs.IpGroupRuleArgs
            {
                Description = "NY",
                Source = "150.24.14.0/24",
            },
            new Aws.Workspaces.Inputs.IpGroupRuleArgs
            {
                Description = "LA",
                Source = "125.191.14.85/32",
            },
            new Aws.Workspaces.Inputs.IpGroupRuleArgs
            {
                Description = "STL",
                Source = "44.98.100.0/24",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/workspaces"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workspaces.NewIpGroup(ctx, "contractors", &workspaces.IpGroupArgs{
			Description: pulumi.String("Contractors IP access control group"),
			Rules: workspaces.IpGroupRuleArray{
				&workspaces.IpGroupRuleArgs{
					Description: pulumi.String("NY"),
					Source:      pulumi.String("150.24.14.0/24"),
				},
				&workspaces.IpGroupRuleArgs{
					Description: pulumi.String("LA"),
					Source:      pulumi.String("125.191.14.85/32"),
				},
				&workspaces.IpGroupRuleArgs{
					Description: pulumi.String("STL"),
					Source:      pulumi.String("44.98.100.0/24"),
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
import com.pulumi.aws.workspaces.IpGroup;
import com.pulumi.aws.workspaces.IpGroupArgs;
import com.pulumi.aws.workspaces.inputs.IpGroupRuleArgs;
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
        var contractors = new IpGroup("contractors", IpGroupArgs.builder()        
            .description("Contractors IP access control group")
            .rules(            
                IpGroupRuleArgs.builder()
                    .description("NY")
                    .source("150.24.14.0/24")
                    .build(),
                IpGroupRuleArgs.builder()
                    .description("LA")
                    .source("125.191.14.85/32")
                    .build(),
                IpGroupRuleArgs.builder()
                    .description("STL")
                    .source("44.98.100.0/24")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  contractors:
    type: aws:workspaces:IpGroup
    properties:
      description: Contractors IP access control group
      rules:
        - description: NY
          source: 150.24.14.0/24
        - description: LA
          source: 125.191.14.85/32
        - description: STL
          source: 44.98.100.0/24
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WorkSpaces IP groups using their GroupID. For example:

```sh
 $ pulumi import aws:workspaces/ipGroup:IpGroup example wsipg-488lrtl3k
```
 