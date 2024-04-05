Provides a Cloud9 EC2 Development Environment.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloud9.EnvironmentEC2("example", {instanceType: "t2.micro"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloud9.EnvironmentEC2("example", instance_type="t2.micro")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Cloud9.EnvironmentEC2("example", new()
    {
        InstanceType = "t2.micro",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloud9"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloud9.NewEnvironmentEC2(ctx, "example", &cloud9.EnvironmentEC2Args{
			InstanceType: pulumi.String("t2.micro"),
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
import com.pulumi.aws.cloud9.EnvironmentEC2;
import com.pulumi.aws.cloud9.EnvironmentEC2Args;
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
        var example = new EnvironmentEC2("example", EnvironmentEC2Args.builder()        
            .instanceType("t2.micro")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloud9:EnvironmentEC2
    properties:
      instanceType: t2.micro
```

Get the URL of the Cloud9 environment after creation:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloud9.EnvironmentEC2("example", {instanceType: "t2.micro"});
const cloud9Instance = aws.ec2.getInstanceOutput({
    filters: [{
        name: "tag:aws:cloud9:environment",
        values: [example.id],
    }],
});
export const cloud9Url = pulumi.interpolate`https://${_var.region}.console.aws.amazon.com/cloud9/ide/${example.id}`;
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloud9.EnvironmentEC2("example", instance_type="t2.micro")
cloud9_instance = aws.ec2.get_instance_output(filters=[aws.ec2.GetInstanceFilterArgs(
    name="tag:aws:cloud9:environment",
    values=[example.id],
)])
pulumi.export("cloud9Url", example.id.apply(lambda id: f"https://{var['region']}.console.aws.amazon.com/cloud9/ide/{id}"))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Cloud9.EnvironmentEC2("example", new()
    {
        InstanceType = "t2.micro",
    });

    var cloud9Instance = Aws.Ec2.GetInstance.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetInstanceFilterInputArgs
            {
                Name = "tag:aws:cloud9:environment",
                Values = new[]
                {
                    example.Id,
                },
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["cloud9Url"] = example.Id.Apply(id => $"https://{@var.Region}.console.aws.amazon.com/cloud9/ide/{id}"),
    };
});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloud9"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := cloud9.NewEnvironmentEC2(ctx, "example", &cloud9.EnvironmentEC2Args{
			InstanceType: pulumi.String("t2.micro"),
		})
		if err != nil {
			return err
		}
		_ = ec2.LookupInstanceOutput(ctx, ec2.GetInstanceOutputArgs{
			Filters: ec2.GetInstanceFilterArray{
				&ec2.GetInstanceFilterArgs{
					Name: pulumi.String("tag:aws:cloud9:environment"),
					Values: pulumi.StringArray{
						example.ID(),
					},
				},
			},
		}, nil)
		ctx.Export("cloud9Url", example.ID().ApplyT(func(id string) (string, error) {
			return fmt.Sprintf("https://%v.console.aws.amazon.com/cloud9/ide/%v", _var.Region, id), nil
		}).(pulumi.StringOutput))
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.cloud9.EnvironmentEC2;
import com.pulumi.aws.cloud9.EnvironmentEC2Args;
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetInstanceArgs;
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
        var example = new EnvironmentEC2("example", EnvironmentEC2Args.builder()        
            .instanceType("t2.micro")
            .build());

        final var cloud9Instance = Ec2Functions.getInstance(GetInstanceArgs.builder()
            .filters(GetInstanceFilterArgs.builder()
                .name("tag:aws:cloud9:environment")
                .values(example.id())
                .build())
            .build());

        ctx.export("cloud9Url", example.id().applyValue(id -> String.format("https://%s.console.aws.amazon.com/cloud9/ide/%s", var_.region(),id)));
    }
}
```
```yaml
resources:
  example:
    type: aws:cloud9:EnvironmentEC2
    properties:
      instanceType: t2.micro
variables:
  cloud9Instance:
    fn::invoke:
      Function: aws:ec2:getInstance
      Arguments:
        filters:
          - name: tag:aws:cloud9:environment
            values:
              - ${example.id}
outputs:
  cloud9Url: https://${var.region}.console.aws.amazon.com/cloud9/ide/${example.id}
```

Allocate a static IP to the Cloud9 environment:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloud9.EnvironmentEC2("example", {instanceType: "t2.micro"});
const cloud9Instance = aws.ec2.getInstanceOutput({
    filters: [{
        name: "tag:aws:cloud9:environment",
        values: [example.id],
    }],
});
const cloud9Eip = new aws.ec2.Eip("cloud9Eip", {
    instance: cloud9Instance.apply(cloud9Instance => cloud9Instance.id),
    domain: "vpc",
});
export const cloud9PublicIp = cloud9Eip.publicIp;
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloud9.EnvironmentEC2("example", instance_type="t2.micro")
cloud9_instance = aws.ec2.get_instance_output(filters=[aws.ec2.GetInstanceFilterArgs(
    name="tag:aws:cloud9:environment",
    values=[example.id],
)])
cloud9_eip = aws.ec2.Eip("cloud9Eip",
    instance=cloud9_instance.id,
    domain="vpc")
pulumi.export("cloud9PublicIp", cloud9_eip.public_ip)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Cloud9.EnvironmentEC2("example", new()
    {
        InstanceType = "t2.micro",
    });

    var cloud9Instance = Aws.Ec2.GetInstance.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetInstanceFilterInputArgs
            {
                Name = "tag:aws:cloud9:environment",
                Values = new[]
                {
                    example.Id,
                },
            },
        },
    });

    var cloud9Eip = new Aws.Ec2.Eip("cloud9Eip", new()
    {
        Instance = cloud9Instance.Apply(getInstanceResult => getInstanceResult.Id),
        Domain = "vpc",
    });

    return new Dictionary<string, object?>
    {
        ["cloud9PublicIp"] = cloud9Eip.PublicIp,
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloud9"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := cloud9.NewEnvironmentEC2(ctx, "example", &cloud9.EnvironmentEC2Args{
			InstanceType: pulumi.String("t2.micro"),
		})
		if err != nil {
			return err
		}
		cloud9Instance := ec2.LookupInstanceOutput(ctx, ec2.GetInstanceOutputArgs{
			Filters: ec2.GetInstanceFilterArray{
				&ec2.GetInstanceFilterArgs{
					Name: pulumi.String("tag:aws:cloud9:environment"),
					Values: pulumi.StringArray{
						example.ID(),
					},
				},
			},
		}, nil)
		cloud9Eip, err := ec2.NewEip(ctx, "cloud9Eip", &ec2.EipArgs{
			Instance: cloud9Instance.ApplyT(func(cloud9Instance ec2.GetInstanceResult) (*string, error) {
				return &cloud9Instance.Id, nil
			}).(pulumi.StringPtrOutput),
			Domain: pulumi.String("vpc"),
		})
		if err != nil {
			return err
		}
		ctx.Export("cloud9PublicIp", cloud9Eip.PublicIp)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.cloud9.EnvironmentEC2;
import com.pulumi.aws.cloud9.EnvironmentEC2Args;
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetInstanceArgs;
import com.pulumi.aws.ec2.Eip;
import com.pulumi.aws.ec2.EipArgs;
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
        var example = new EnvironmentEC2("example", EnvironmentEC2Args.builder()        
            .instanceType("t2.micro")
            .build());

        final var cloud9Instance = Ec2Functions.getInstance(GetInstanceArgs.builder()
            .filters(GetInstanceFilterArgs.builder()
                .name("tag:aws:cloud9:environment")
                .values(example.id())
                .build())
            .build());

        var cloud9Eip = new Eip("cloud9Eip", EipArgs.builder()        
            .instance(cloud9Instance.applyValue(getInstanceResult -> getInstanceResult).applyValue(cloud9Instance -> cloud9Instance.applyValue(getInstanceResult -> getInstanceResult.id())))
            .domain("vpc")
            .build());

        ctx.export("cloud9PublicIp", cloud9Eip.publicIp());
    }
}
```
```yaml
resources:
  example:
    type: aws:cloud9:EnvironmentEC2
    properties:
      instanceType: t2.micro
  cloud9Eip:
    type: aws:ec2:Eip
    properties:
      instance: ${cloud9Instance.id}
      domain: vpc
variables:
  cloud9Instance:
    fn::invoke:
      Function: aws:ec2:getInstance
      Arguments:
        filters:
          - name: tag:aws:cloud9:environment
            values:
              - ${example.id}
outputs:
  cloud9PublicIp: ${cloud9Eip.publicIp}
```
{{% /example %}}
{{% /examples %}}