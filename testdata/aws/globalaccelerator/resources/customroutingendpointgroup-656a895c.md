Provides a Global Accelerator custom routing endpoint group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.globalaccelerator.CustomRoutingEndpointGroup("example", {
    listenerArn: aws_globalaccelerator_custom_routing_listener.example.id,
    destinationConfigurations: [{
        fromPort: 80,
        toPort: 8080,
        protocols: ["TCP"],
    }],
    endpointConfigurations: [{
        endpointId: aws_subnet.example.id,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.globalaccelerator.CustomRoutingEndpointGroup("example",
    listener_arn=aws_globalaccelerator_custom_routing_listener["example"]["id"],
    destination_configurations=[aws.globalaccelerator.CustomRoutingEndpointGroupDestinationConfigurationArgs(
        from_port=80,
        to_port=8080,
        protocols=["TCP"],
    )],
    endpoint_configurations=[aws.globalaccelerator.CustomRoutingEndpointGroupEndpointConfigurationArgs(
        endpoint_id=aws_subnet["example"]["id"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.GlobalAccelerator.CustomRoutingEndpointGroup("example", new()
    {
        ListenerArn = aws_globalaccelerator_custom_routing_listener.Example.Id,
        DestinationConfigurations = new[]
        {
            new Aws.GlobalAccelerator.Inputs.CustomRoutingEndpointGroupDestinationConfigurationArgs
            {
                FromPort = 80,
                ToPort = 8080,
                Protocols = new[]
                {
                    "TCP",
                },
            },
        },
        EndpointConfigurations = new[]
        {
            new Aws.GlobalAccelerator.Inputs.CustomRoutingEndpointGroupEndpointConfigurationArgs
            {
                EndpointId = aws_subnet.Example.Id,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/globalaccelerator"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := globalaccelerator.NewCustomRoutingEndpointGroup(ctx, "example", &globalaccelerator.CustomRoutingEndpointGroupArgs{
			ListenerArn: pulumi.Any(aws_globalaccelerator_custom_routing_listener.Example.Id),
			DestinationConfigurations: globalaccelerator.CustomRoutingEndpointGroupDestinationConfigurationArray{
				&globalaccelerator.CustomRoutingEndpointGroupDestinationConfigurationArgs{
					FromPort: pulumi.Int(80),
					ToPort:   pulumi.Int(8080),
					Protocols: pulumi.StringArray{
						pulumi.String("TCP"),
					},
				},
			},
			EndpointConfigurations: globalaccelerator.CustomRoutingEndpointGroupEndpointConfigurationArray{
				&globalaccelerator.CustomRoutingEndpointGroupEndpointConfigurationArgs{
					EndpointId: pulumi.Any(aws_subnet.Example.Id),
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
import com.pulumi.aws.globalaccelerator.CustomRoutingEndpointGroup;
import com.pulumi.aws.globalaccelerator.CustomRoutingEndpointGroupArgs;
import com.pulumi.aws.globalaccelerator.inputs.CustomRoutingEndpointGroupDestinationConfigurationArgs;
import com.pulumi.aws.globalaccelerator.inputs.CustomRoutingEndpointGroupEndpointConfigurationArgs;
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
        var example = new CustomRoutingEndpointGroup("example", CustomRoutingEndpointGroupArgs.builder()        
            .listenerArn(aws_globalaccelerator_custom_routing_listener.example().id())
            .destinationConfigurations(CustomRoutingEndpointGroupDestinationConfigurationArgs.builder()
                .fromPort(80)
                .toPort(8080)
                .protocols("TCP")
                .build())
            .endpointConfigurations(CustomRoutingEndpointGroupEndpointConfigurationArgs.builder()
                .endpointId(aws_subnet.example().id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:globalaccelerator:CustomRoutingEndpointGroup
    properties:
      listenerArn: ${aws_globalaccelerator_custom_routing_listener.example.id}
      destinationConfigurations:
        - fromPort: 80
          toPort: 8080
          protocols:
            - TCP
      endpointConfigurations:
        - endpointId: ${aws_subnet.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Global Accelerator custom routing endpoint groups using the `id`. For example:

```sh
 $ pulumi import aws:globalaccelerator/customRoutingEndpointGroup:CustomRoutingEndpointGroup example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxx/endpoint-group/xxxxxxxx
```
 