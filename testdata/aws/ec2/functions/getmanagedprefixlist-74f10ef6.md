`aws.ec2.ManagedPrefixList` provides details about a specific AWS prefix list or
customer-managed prefix list in the current region.

{{% examples %}}
## Example Usage
{{% example %}}
### Find the regional DynamoDB prefix list

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegion({});
const example = current.then(current => aws.ec2.getManagedPrefixList({
    name: `com.amazonaws.${current.name}.dynamodb`,
}));
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_region()
example = aws.ec2.get_managed_prefix_list(name=f"com.amazonaws.{current.name}.dynamodb")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegion.Invoke();

    var example = Aws.Ec2.GetManagedPrefixList.Invoke(new()
    {
        Name = $"com.amazonaws.{current.Apply(getRegionResult => getRegionResult.Name)}.dynamodb",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = ec2.LookupManagedPrefixList(ctx, &ec2.LookupManagedPrefixListArgs{
			Name: pulumi.StringRef(fmt.Sprintf("com.amazonaws.%v.dynamodb", current.Name)),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetManagedPrefixListArgs;
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
        final var current = AwsFunctions.getRegion();

        final var example = Ec2Functions.getManagedPrefixList(GetManagedPrefixListArgs.builder()
            .name(String.format("com.amazonaws.%s.dynamodb", current.applyValue(getRegionResult -> getRegionResult.name())))
            .build());

    }
}
```
```yaml
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  example:
    fn::invoke:
      Function: aws:ec2:getManagedPrefixList
      Arguments:
        name: com.amazonaws.${current.name}.dynamodb
```
{{% /example %}}
{{% example %}}
### Find a managed prefix list using filters

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getManagedPrefixList({
    filters: [{
        name: "prefix-list-name",
        values: ["my-prefix-list"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_managed_prefix_list(filters=[aws.ec2.GetManagedPrefixListFilterArgs(
    name="prefix-list-name",
    values=["my-prefix-list"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetManagedPrefixList.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetManagedPrefixListFilterInputArgs
            {
                Name = "prefix-list-name",
                Values = new[]
                {
                    "my-prefix-list",
                },
            },
        },
    });

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
		_, err := ec2.LookupManagedPrefixList(ctx, &ec2.LookupManagedPrefixListArgs{
			Filters: []ec2.GetManagedPrefixListFilter{
				{
					Name: "prefix-list-name",
					Values: []string{
						"my-prefix-list",
					},
				},
			},
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetManagedPrefixListArgs;
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
        final var example = Ec2Functions.getManagedPrefixList(GetManagedPrefixListArgs.builder()
            .filters(GetManagedPrefixListFilterArgs.builder()
                .name("prefix-list-name")
                .values("my-prefix-list")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getManagedPrefixList
      Arguments:
        filters:
          - name: prefix-list-name
            values:
              - my-prefix-list
```
{{% /example %}}
{{% /examples %}}