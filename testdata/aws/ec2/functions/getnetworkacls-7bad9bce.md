{{% examples %}}
## Example Usage
{{% example %}}

The following shows outputting all network ACL ids in a vpc.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleNetworkAcls = aws.ec2.getNetworkAcls({
    vpcId: _var.vpc_id,
});
export const example = exampleNetworkAcls.then(exampleNetworkAcls => exampleNetworkAcls.ids);
```
```python
import pulumi
import pulumi_aws as aws

example_network_acls = aws.ec2.get_network_acls(vpc_id=var["vpc_id"])
pulumi.export("example", example_network_acls.ids)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleNetworkAcls = Aws.Ec2.GetNetworkAcls.Invoke(new()
    {
        VpcId = @var.Vpc_id,
    });

    return new Dictionary<string, object?>
    {
        ["example"] = exampleNetworkAcls.Apply(getNetworkAclsResult => getNetworkAclsResult.Ids),
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
		exampleNetworkAcls, err := ec2.GetNetworkAcls(ctx, &ec2.GetNetworkAclsArgs{
			VpcId: pulumi.StringRef(_var.Vpc_id),
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("example", exampleNetworkAcls.Ids)
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
import com.pulumi.aws.ec2.inputs.GetNetworkAclsArgs;
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
        final var exampleNetworkAcls = Ec2Functions.getNetworkAcls(GetNetworkAclsArgs.builder()
            .vpcId(var_.vpc_id())
            .build());

        ctx.export("example", exampleNetworkAcls.applyValue(getNetworkAclsResult -> getNetworkAclsResult.ids()));
    }
}
```
```yaml
variables:
  exampleNetworkAcls:
    fn::invoke:
      Function: aws:ec2:getNetworkAcls
      Arguments:
        vpcId: ${var.vpc_id}
outputs:
  example: ${exampleNetworkAcls.ids}
```

The following example retrieves a list of all network ACL ids in a VPC with a custom
tag of `Tier` set to a value of "Private".

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getNetworkAcls({
    vpcId: _var.vpc_id,
    tags: {
        Tier: "Private",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_network_acls(vpc_id=var["vpc_id"],
    tags={
        "Tier": "Private",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetNetworkAcls.Invoke(new()
    {
        VpcId = @var.Vpc_id,
        Tags = 
        {
            { "Tier", "Private" },
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
		_, err := ec2.GetNetworkAcls(ctx, &ec2.GetNetworkAclsArgs{
			VpcId: pulumi.StringRef(_var.Vpc_id),
			Tags: map[string]interface{}{
				"Tier": "Private",
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
import com.pulumi.aws.ec2.inputs.GetNetworkAclsArgs;
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
        final var example = Ec2Functions.getNetworkAcls(GetNetworkAclsArgs.builder()
            .vpcId(var_.vpc_id())
            .tags(Map.of("Tier", "Private"))
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getNetworkAcls
      Arguments:
        vpcId: ${var.vpc_id}
        tags:
          Tier: Private
```

The following example retrieves a network ACL id in a VPC which associated
with specific subnet.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getNetworkAcls({
    vpcId: _var.vpc_id,
    filters: [{
        name: "association.subnet-id",
        values: [aws_subnet.test.id],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_network_acls(vpc_id=var["vpc_id"],
    filters=[aws.ec2.GetNetworkAclsFilterArgs(
        name="association.subnet-id",
        values=[aws_subnet["test"]["id"]],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetNetworkAcls.Invoke(new()
    {
        VpcId = @var.Vpc_id,
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetNetworkAclsFilterInputArgs
            {
                Name = "association.subnet-id",
                Values = new[]
                {
                    aws_subnet.Test.Id,
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
_, err := ec2.GetNetworkAcls(ctx, &ec2.GetNetworkAclsArgs{
VpcId: pulumi.StringRef(_var.Vpc_id),
Filters: []ec2.GetNetworkAclsFilter{
{
Name: "association.subnet-id",
Values: interface{}{
aws_subnet.Test.Id,
},
},
},
}, nil);
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
import com.pulumi.aws.ec2.inputs.GetNetworkAclsArgs;
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
        final var example = Ec2Functions.getNetworkAcls(GetNetworkAclsArgs.builder()
            .vpcId(var_.vpc_id())
            .filters(GetNetworkAclsFilterArgs.builder()
                .name("association.subnet-id")
                .values(aws_subnet.test().id())
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getNetworkAcls
      Arguments:
        vpcId: ${var.vpc_id}
        filters:
          - name: association.subnet-id
            values:
              - ${aws_subnet.test.id}
```
{{% /example %}}
{{% /examples %}}