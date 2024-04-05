The Availability Zones data source allows access to the list of AWS
Availability Zones which can be accessed by an AWS account within the region
configured in the provider.

This is different from the `aws.getAvailabilityZone` (singular) data source,
which provides some details about a specific availability zone.

> When [Local Zones](https://aws.amazon.com/about-aws/global-infrastructure/localzones/) are enabled in a region, by default the API and this data source include both Local Zones and Availability Zones. To return only Availability Zones, see the example section below.

{{% examples %}}
## Example Usage
{{% example %}}
### By State

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const available = aws.getAvailabilityZones({
    state: "available",
});
const primary = new aws.ec2.Subnet("primary", {availabilityZone: available.then(available => available.names?.[0])});
// ...
const secondary = new aws.ec2.Subnet("secondary", {availabilityZone: available.then(available => available.names?.[1])});
// ...
```
```python
import pulumi
import pulumi_aws as aws

available = aws.get_availability_zones(state="available")
primary = aws.ec2.Subnet("primary", availability_zone=available.names[0])
# ...
secondary = aws.ec2.Subnet("secondary", availability_zone=available.names[1])
# ...
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var available = Aws.GetAvailabilityZones.Invoke(new()
    {
        State = "available",
    });

    var primary = new Aws.Ec2.Subnet("primary", new()
    {
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
    });

    // ...
    var secondary = new Aws.Ec2.Subnet("secondary", new()
    {
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[1]),
    });

    // ...
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		available, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			State: pulumi.StringRef("available"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSubnet(ctx, "primary", &ec2.SubnetArgs{
			AvailabilityZone: *pulumi.String(available.Names[0]),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewSubnet(ctx, "secondary", &ec2.SubnetArgs{
			AvailabilityZone: *pulumi.String(available.Names[1]),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
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
        final var available = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
            .state("available")
            .build());

        var primary = new Subnet("primary", SubnetArgs.builder()        
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]))
            .build());

        var secondary = new Subnet("secondary", SubnetArgs.builder()        
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[1]))
            .build());

    }
}
```
```yaml
resources:
  primary:
    type: aws:ec2:Subnet
    properties:
      availabilityZone: ${available.names[0]}
  secondary:
    type: aws:ec2:Subnet
    properties:
      availabilityZone: ${available.names[1]}
variables:
  available: # e.g., Create subnets in the first two available availability zones
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments:
        state: available
```
{{% /example %}}
{{% example %}}
### By Filter

All Local Zones (regardless of opt-in status):

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.getAvailabilityZones({
    allAvailabilityZones: true,
    filters: [{
        name: "opt-in-status",
        values: [
            "not-opted-in",
            "opted-in",
        ],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.get_availability_zones(all_availability_zones=True,
    filters=[aws.GetAvailabilityZonesFilterArgs(
        name="opt-in-status",
        values=[
            "not-opted-in",
            "opted-in",
        ],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.GetAvailabilityZones.Invoke(new()
    {
        AllAvailabilityZones = true,
        Filters = new[]
        {
            new Aws.Inputs.GetAvailabilityZonesFilterInputArgs
            {
                Name = "opt-in-status",
                Values = new[]
                {
                    "not-opted-in",
                    "opted-in",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			AllAvailabilityZones: pulumi.BoolRef(true),
			Filters: []aws.GetAvailabilityZonesFilter{
				{
					Name: "opt-in-status",
					Values: []string{
						"not-opted-in",
						"opted-in",
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
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
        final var example = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
            .allAvailabilityZones(true)
            .filters(GetAvailabilityZonesFilterArgs.builder()
                .name("opt-in-status")
                .values(                
                    "not-opted-in",
                    "opted-in")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments:
        allAvailabilityZones: true
        filters:
          - name: opt-in-status
            values:
              - not-opted-in
              - opted-in
```

Only Availability Zones (no Local Zones):

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.getAvailabilityZones({
    filters: [{
        name: "opt-in-status",
        values: ["opt-in-not-required"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.get_availability_zones(filters=[aws.GetAvailabilityZonesFilterArgs(
    name="opt-in-status",
    values=["opt-in-not-required"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.GetAvailabilityZones.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Inputs.GetAvailabilityZonesFilterInputArgs
            {
                Name = "opt-in-status",
                Values = new[]
                {
                    "opt-in-not-required",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			Filters: []aws.GetAvailabilityZonesFilter{
				{
					Name: "opt-in-status",
					Values: []string{
						"opt-in-not-required",
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
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
        final var example = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
            .filters(GetAvailabilityZonesFilterArgs.builder()
                .name("opt-in-status")
                .values("opt-in-not-required")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments:
        filters:
          - name: opt-in-status
            values:
              - opt-in-not-required
```
{{% /example %}}
{{% /examples %}}