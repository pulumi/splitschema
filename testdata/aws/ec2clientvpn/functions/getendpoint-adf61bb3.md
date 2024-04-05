Get information on an EC2 Client VPN endpoint.

{{% examples %}}
## Example Usage
{{% example %}}
### By Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2clientvpn.getEndpoint({
    filters: [{
        name: "tag:Name",
        values: ["ExampleVpn"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2clientvpn.get_endpoint(filters=[aws.ec2clientvpn.GetEndpointFilterArgs(
    name="tag:Name",
    values=["ExampleVpn"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2ClientVpn.GetEndpoint.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2ClientVpn.Inputs.GetEndpointFilterInputArgs
            {
                Name = "tag:Name",
                Values = new[]
                {
                    "ExampleVpn",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2clientvpn"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2clientvpn.LookupEndpoint(ctx, &ec2clientvpn.LookupEndpointArgs{
			Filters: []ec2clientvpn.GetEndpointFilter{
				{
					Name: "tag:Name",
					Values: []string{
						"ExampleVpn",
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
import com.pulumi.aws.ec2clientvpn.Ec2clientvpnFunctions;
import com.pulumi.aws.ec2clientvpn.inputs.GetEndpointArgs;
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
        final var example = Ec2clientvpnFunctions.getEndpoint(GetEndpointArgs.builder()
            .filters(GetEndpointFilterArgs.builder()
                .name("tag:Name")
                .values("ExampleVpn")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2clientvpn:getEndpoint
      Arguments:
        filters:
          - name: tag:Name
            values:
              - ExampleVpn
```
{{% /example %}}
{{% example %}}
### By Identifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2clientvpn.getEndpoint({
    clientVpnEndpointId: "cvpn-endpoint-083cf50d6eb314f21",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2clientvpn.get_endpoint(client_vpn_endpoint_id="cvpn-endpoint-083cf50d6eb314f21")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2ClientVpn.GetEndpoint.Invoke(new()
    {
        ClientVpnEndpointId = "cvpn-endpoint-083cf50d6eb314f21",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2clientvpn"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2clientvpn.LookupEndpoint(ctx, &ec2clientvpn.LookupEndpointArgs{
			ClientVpnEndpointId: pulumi.StringRef("cvpn-endpoint-083cf50d6eb314f21"),
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
import com.pulumi.aws.ec2clientvpn.Ec2clientvpnFunctions;
import com.pulumi.aws.ec2clientvpn.inputs.GetEndpointArgs;
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
        final var example = Ec2clientvpnFunctions.getEndpoint(GetEndpointArgs.builder()
            .clientVpnEndpointId("cvpn-endpoint-083cf50d6eb314f21")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2clientvpn:getEndpoint
      Arguments:
        clientVpnEndpointId: cvpn-endpoint-083cf50d6eb314f21
```
{{% /example %}}
{{% /examples %}}