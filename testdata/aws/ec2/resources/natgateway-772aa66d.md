Provides a resource to create a VPC NAT Gateway.

{{% examples %}}
## Example Usage
{{% example %}}
### Public NAT

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.NatGateway("example", {
    allocationId: aws_eip.example.id,
    subnetId: aws_subnet.example.id,
    tags: {
        Name: "gw NAT",
    },
}, {
    dependsOn: [aws_internet_gateway.example],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.NatGateway("example",
    allocation_id=aws_eip["example"]["id"],
    subnet_id=aws_subnet["example"]["id"],
    tags={
        "Name": "gw NAT",
    },
    opts=pulumi.ResourceOptions(depends_on=[aws_internet_gateway["example"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.NatGateway("example", new()
    {
        AllocationId = aws_eip.Example.Id,
        SubnetId = aws_subnet.Example.Id,
        Tags = 
        {
            { "Name", "gw NAT" },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_internet_gateway.Example,
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
		_, err := ec2.NewNatGateway(ctx, "example", &ec2.NatGatewayArgs{
			AllocationId: pulumi.Any(aws_eip.Example.Id),
			SubnetId:     pulumi.Any(aws_subnet.Example.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("gw NAT"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_internet_gateway.Example,
		}))
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
import com.pulumi.aws.ec2.NatGateway;
import com.pulumi.aws.ec2.NatGatewayArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var example = new NatGateway("example", NatGatewayArgs.builder()        
            .allocationId(aws_eip.example().id())
            .subnetId(aws_subnet.example().id())
            .tags(Map.of("Name", "gw NAT"))
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_internet_gateway.example())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:NatGateway
    properties:
      allocationId: ${aws_eip.example.id}
      subnetId: ${aws_subnet.example.id}
      tags:
        Name: gw NAT
    options:
      dependson:
        - ${aws_internet_gateway.example}
```
{{% /example %}}
{{% example %}}
### Public NAT with Secondary Private IP Addresses

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.NatGateway("example", {
    allocationId: aws_eip.example.id,
    subnetId: aws_subnet.example.id,
    secondaryAllocationIds: [aws_eip.secondary.id],
    secondaryPrivateIpAddresses: ["10.0.1.5"],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.NatGateway("example",
    allocation_id=aws_eip["example"]["id"],
    subnet_id=aws_subnet["example"]["id"],
    secondary_allocation_ids=[aws_eip["secondary"]["id"]],
    secondary_private_ip_addresses=["10.0.1.5"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.NatGateway("example", new()
    {
        AllocationId = aws_eip.Example.Id,
        SubnetId = aws_subnet.Example.Id,
        SecondaryAllocationIds = new[]
        {
            aws_eip.Secondary.Id,
        },
        SecondaryPrivateIpAddresses = new[]
        {
            "10.0.1.5",
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
		_, err := ec2.NewNatGateway(ctx, "example", &ec2.NatGatewayArgs{
			AllocationId: pulumi.Any(aws_eip.Example.Id),
			SubnetId:     pulumi.Any(aws_subnet.Example.Id),
			SecondaryAllocationIds: pulumi.StringArray{
				aws_eip.Secondary.Id,
			},
			SecondaryPrivateIpAddresses: pulumi.StringArray{
				pulumi.String("10.0.1.5"),
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
import com.pulumi.aws.ec2.NatGateway;
import com.pulumi.aws.ec2.NatGatewayArgs;
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
        var example = new NatGateway("example", NatGatewayArgs.builder()        
            .allocationId(aws_eip.example().id())
            .subnetId(aws_subnet.example().id())
            .secondaryAllocationIds(aws_eip.secondary().id())
            .secondaryPrivateIpAddresses("10.0.1.5")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:NatGateway
    properties:
      allocationId: ${aws_eip.example.id}
      subnetId: ${aws_subnet.example.id}
      secondaryAllocationIds:
        - ${aws_eip.secondary.id}
      secondaryPrivateIpAddresses:
        - 10.0.1.5
```
{{% /example %}}
{{% example %}}
### Private NAT

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.NatGateway("example", {
    connectivityType: "private",
    subnetId: aws_subnet.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.NatGateway("example",
    connectivity_type="private",
    subnet_id=aws_subnet["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.NatGateway("example", new()
    {
        ConnectivityType = "private",
        SubnetId = aws_subnet.Example.Id,
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
		_, err := ec2.NewNatGateway(ctx, "example", &ec2.NatGatewayArgs{
			ConnectivityType: pulumi.String("private"),
			SubnetId:         pulumi.Any(aws_subnet.Example.Id),
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
import com.pulumi.aws.ec2.NatGateway;
import com.pulumi.aws.ec2.NatGatewayArgs;
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
        var example = new NatGateway("example", NatGatewayArgs.builder()        
            .connectivityType("private")
            .subnetId(aws_subnet.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:NatGateway
    properties:
      connectivityType: private
      subnetId: ${aws_subnet.example.id}
```
{{% /example %}}
{{% example %}}
### Private NAT with Secondary Private IP Addresses

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.NatGateway("example", {
    connectivityType: "private",
    subnetId: aws_subnet.example.id,
    secondaryPrivateIpAddressCount: 7,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.NatGateway("example",
    connectivity_type="private",
    subnet_id=aws_subnet["example"]["id"],
    secondary_private_ip_address_count=7)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.NatGateway("example", new()
    {
        ConnectivityType = "private",
        SubnetId = aws_subnet.Example.Id,
        SecondaryPrivateIpAddressCount = 7,
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
		_, err := ec2.NewNatGateway(ctx, "example", &ec2.NatGatewayArgs{
			ConnectivityType:               pulumi.String("private"),
			SubnetId:                       pulumi.Any(aws_subnet.Example.Id),
			SecondaryPrivateIpAddressCount: pulumi.Int(7),
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
import com.pulumi.aws.ec2.NatGateway;
import com.pulumi.aws.ec2.NatGatewayArgs;
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
        var example = new NatGateway("example", NatGatewayArgs.builder()        
            .connectivityType("private")
            .subnetId(aws_subnet.example().id())
            .secondaryPrivateIpAddressCount(7)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:NatGateway
    properties:
      connectivityType: private
      subnetId: ${aws_subnet.example.id}
      secondaryPrivateIpAddressCount: 7
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import NAT Gateways using the `id`. For example:

```sh
 $ pulumi import aws:ec2/natGateway:NatGateway private_gw nat-05dba92075d71c408
```
 