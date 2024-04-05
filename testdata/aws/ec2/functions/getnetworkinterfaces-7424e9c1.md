{{% examples %}}
## Example Usage
{{% example %}}

The following shows outputting all network interface ids in a region.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleNetworkInterfaces = aws.ec2.getNetworkInterfaces({});
export const example = exampleNetworkInterfaces.then(exampleNetworkInterfaces => exampleNetworkInterfaces.ids);
```
```python
import pulumi
import pulumi_aws as aws

example_network_interfaces = aws.ec2.get_network_interfaces()
pulumi.export("example", example_network_interfaces.ids)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleNetworkInterfaces = Aws.Ec2.GetNetworkInterfaces.Invoke();

    return new Dictionary<string, object?>
    {
        ["example"] = exampleNetworkInterfaces.Apply(getNetworkInterfacesResult => getNetworkInterfacesResult.Ids),
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleNetworkInterfaces, err := ec2.GetNetworkInterfaces(ctx, nil, nil)
		if err != nil {
			return err
		}
		ctx.Export("example", exampleNetworkInterfaces.Ids)
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
import com.pulumi.aws.ec2.inputs.GetNetworkInterfacesArgs;
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
        final var exampleNetworkInterfaces = Ec2Functions.getNetworkInterfaces();

        ctx.export("example", exampleNetworkInterfaces.applyValue(getNetworkInterfacesResult -> getNetworkInterfacesResult.ids()));
    }
}
```
```yaml
variables:
  exampleNetworkInterfaces:
    fn::invoke:
      Function: aws:ec2:getNetworkInterfaces
      Arguments: {}
outputs:
  example: ${exampleNetworkInterfaces.ids}
```

The following example retrieves a list of all network interface ids with a custom tag of `Name` set to a value of `test`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getNetworkInterfaces({
    tags: {
        Name: "test",
    },
});
export const example1 = example.then(example => example.ids);
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_network_interfaces(tags={
    "Name": "test",
})
pulumi.export("example1", example.ids)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetNetworkInterfaces.Invoke(new()
    {
        Tags = 
        {
            { "Name", "test" },
        },
    });

    return new Dictionary<string, object?>
    {
        ["example1"] = example.Apply(getNetworkInterfacesResult => getNetworkInterfacesResult.Ids),
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := ec2.GetNetworkInterfaces(ctx, &ec2.GetNetworkInterfacesArgs{
			Tags: map[string]interface{}{
				"Name": "test",
			},
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("example1", example.Ids)
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
import com.pulumi.aws.ec2.inputs.GetNetworkInterfacesArgs;
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
        final var example = Ec2Functions.getNetworkInterfaces(GetNetworkInterfacesArgs.builder()
            .tags(Map.of("Name", "test"))
            .build());

        ctx.export("example1", example.applyValue(getNetworkInterfacesResult -> getNetworkInterfacesResult.ids()));
    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getNetworkInterfaces
      Arguments:
        tags:
          Name: test
outputs:
  example1: ${example.ids}
```

The following example retrieves a network interface ids which associated
with specific subnet.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleNetworkInterfaces = aws.ec2.getNetworkInterfaces({
    filters: [{
        name: "subnet-id",
        values: [aws_subnet.test.id],
    }],
});
export const example = exampleNetworkInterfaces.then(exampleNetworkInterfaces => exampleNetworkInterfaces.ids);
```
```python
import pulumi
import pulumi_aws as aws

example_network_interfaces = aws.ec2.get_network_interfaces(filters=[aws.ec2.GetNetworkInterfacesFilterArgs(
    name="subnet-id",
    values=[aws_subnet["test"]["id"]],
)])
pulumi.export("example", example_network_interfaces.ids)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleNetworkInterfaces = Aws.Ec2.GetNetworkInterfaces.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetNetworkInterfacesFilterInputArgs
            {
                Name = "subnet-id",
                Values = new[]
                {
                    aws_subnet.Test.Id,
                },
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["example"] = exampleNetworkInterfaces.Apply(getNetworkInterfacesResult => getNetworkInterfacesResult.Ids),
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
exampleNetworkInterfaces, err := ec2.GetNetworkInterfaces(ctx, &ec2.GetNetworkInterfacesArgs{
Filters: []ec2.GetNetworkInterfacesFilter{
{
Name: "subnet-id",
Values: interface{}{
aws_subnet.Test.Id,
},
},
},
}, nil);
if err != nil {
return err
}
ctx.Export("example", exampleNetworkInterfaces.Ids)
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
import com.pulumi.aws.ec2.inputs.GetNetworkInterfacesArgs;
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
        final var exampleNetworkInterfaces = Ec2Functions.getNetworkInterfaces(GetNetworkInterfacesArgs.builder()
            .filters(GetNetworkInterfacesFilterArgs.builder()
                .name("subnet-id")
                .values(aws_subnet.test().id())
                .build())
            .build());

        ctx.export("example", exampleNetworkInterfaces.applyValue(getNetworkInterfacesResult -> getNetworkInterfacesResult.ids()));
    }
}
```
```yaml
variables:
  exampleNetworkInterfaces:
    fn::invoke:
      Function: aws:ec2:getNetworkInterfaces
      Arguments:
        filters:
          - name: subnet-id
            values:
              - ${aws_subnet.test.id}
outputs:
  example: ${exampleNetworkInterfaces.ids}
```
{{% /example %}}
{{% /examples %}}