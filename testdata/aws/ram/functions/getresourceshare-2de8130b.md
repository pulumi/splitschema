`aws.ram.ResourceShare` Retrieve information about a RAM Resource Share.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ram.getResourceShare({
    name: "example",
    resourceOwner: "SELF",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ram.get_resource_share(name="example",
    resource_owner="SELF")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ram.GetResourceShare.Invoke(new()
    {
        Name = "example",
        ResourceOwner = "SELF",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ram.LookupResourceShare(ctx, &ram.LookupResourceShareArgs{
			Name:          "example",
			ResourceOwner: "SELF",
		}, nil)
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
import com.pulumi.aws.ram.RamFunctions;
import com.pulumi.aws.ram.inputs.GetResourceShareArgs;
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
        final var example = RamFunctions.getResourceShare(GetResourceShareArgs.builder()
            .name("example")
            .resourceOwner("SELF")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ram:getResourceShare
      Arguments:
        name: example
        resourceOwner: SELF
```
{{% /example %}}
{{% /examples %}}
## Search by filters

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const tagFilter = aws.ram.getResourceShare({
    filters: [{
        name: "NameOfTag",
        values: ["exampleNameTagValue"],
    }],
    name: "MyResourceName",
    resourceOwner: "SELF",
});
```
```python
import pulumi
import pulumi_aws as aws

tag_filter = aws.ram.get_resource_share(filters=[aws.ram.GetResourceShareFilterArgs(
        name="NameOfTag",
        values=["exampleNameTagValue"],
    )],
    name="MyResourceName",
    resource_owner="SELF")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var tagFilter = Aws.Ram.GetResourceShare.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ram.Inputs.GetResourceShareFilterInputArgs
            {
                Name = "NameOfTag",
                Values = new[]
                {
                    "exampleNameTagValue",
                },
            },
        },
        Name = "MyResourceName",
        ResourceOwner = "SELF",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ram.LookupResourceShare(ctx, &ram.LookupResourceShareArgs{
			Filters: []ram.GetResourceShareFilter{
				{
					Name: "NameOfTag",
					Values: []string{
						"exampleNameTagValue",
					},
				},
			},
			Name:          "MyResourceName",
			ResourceOwner: "SELF",
		}, nil)
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
import com.pulumi.aws.ram.RamFunctions;
import com.pulumi.aws.ram.inputs.GetResourceShareArgs;
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
        final var tagFilter = RamFunctions.getResourceShare(GetResourceShareArgs.builder()
            .filters(GetResourceShareFilterArgs.builder()
                .name("NameOfTag")
                .values("exampleNameTagValue")
                .build())
            .name("MyResourceName")
            .resourceOwner("SELF")
            .build());

    }
}
```
```yaml
variables:
  tagFilter:
    fn::invoke:
      Function: aws:ram:getResourceShare
      Arguments:
        filters:
          - name: NameOfTag
            values:
              - exampleNameTagValue
        name: MyResourceName
        resourceOwner: SELF
```
