Data source for getting information about AWS EC2 Public IPv4 Pools.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getPublicIpv4Pools({});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_public_ipv4_pools()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetPublicIpv4Pools.Invoke();

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
		_, err := ec2.GetPublicIpv4Pools(ctx, nil, nil)
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
import com.pulumi.aws.ec2.inputs.GetPublicIpv4PoolsArgs;
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
        final var example = Ec2Functions.getPublicIpv4Pools();

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getPublicIpv4Pools
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### Usage with Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getPublicIpv4Pools({
    filters: [{
        name: "tag-key",
        values: ["ExampleTagKey"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_public_ipv4_pools(filters=[aws.ec2.GetPublicIpv4PoolsFilterArgs(
    name="tag-key",
    values=["ExampleTagKey"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetPublicIpv4Pools.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetPublicIpv4PoolsFilterInputArgs
            {
                Name = "tag-key",
                Values = new[]
                {
                    "ExampleTagKey",
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
		_, err := ec2.GetPublicIpv4Pools(ctx, &ec2.GetPublicIpv4PoolsArgs{
			Filters: []ec2.GetPublicIpv4PoolsFilter{
				{
					Name: "tag-key",
					Values: []string{
						"ExampleTagKey",
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
import com.pulumi.aws.ec2.inputs.GetPublicIpv4PoolsArgs;
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
        final var example = Ec2Functions.getPublicIpv4Pools(GetPublicIpv4PoolsArgs.builder()
            .filters(GetPublicIpv4PoolsFilterArgs.builder()
                .name("tag-key")
                .values("ExampleTagKey")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getPublicIpv4Pools
      Arguments:
        filters:
          - name: tag-key
            values:
              - ExampleTagKey
```
{{% /example %}}
{{% /examples %}}