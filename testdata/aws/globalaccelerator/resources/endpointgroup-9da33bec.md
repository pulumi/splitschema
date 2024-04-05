Provides a Global Accelerator endpoint group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.globalaccelerator.EndpointGroup("example", {
    listenerArn: aws_globalaccelerator_listener.example.id,
    endpointConfigurations: [{
        endpointId: aws_lb.example.arn,
        weight: 100,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.globalaccelerator.EndpointGroup("example",
    listener_arn=aws_globalaccelerator_listener["example"]["id"],
    endpoint_configurations=[aws.globalaccelerator.EndpointGroupEndpointConfigurationArgs(
        endpoint_id=aws_lb["example"]["arn"],
        weight=100,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.GlobalAccelerator.EndpointGroup("example", new()
    {
        ListenerArn = aws_globalaccelerator_listener.Example.Id,
        EndpointConfigurations = new[]
        {
            new Aws.GlobalAccelerator.Inputs.EndpointGroupEndpointConfigurationArgs
            {
                EndpointId = aws_lb.Example.Arn,
                Weight = 100,
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
		_, err := globalaccelerator.NewEndpointGroup(ctx, "example", &globalaccelerator.EndpointGroupArgs{
			ListenerArn: pulumi.Any(aws_globalaccelerator_listener.Example.Id),
			EndpointConfigurations: globalaccelerator.EndpointGroupEndpointConfigurationArray{
				&globalaccelerator.EndpointGroupEndpointConfigurationArgs{
					EndpointId: pulumi.Any(aws_lb.Example.Arn),
					Weight:     pulumi.Int(100),
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
import com.pulumi.aws.globalaccelerator.EndpointGroup;
import com.pulumi.aws.globalaccelerator.EndpointGroupArgs;
import com.pulumi.aws.globalaccelerator.inputs.EndpointGroupEndpointConfigurationArgs;
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
        var example = new EndpointGroup("example", EndpointGroupArgs.builder()        
            .listenerArn(aws_globalaccelerator_listener.example().id())
            .endpointConfigurations(EndpointGroupEndpointConfigurationArgs.builder()
                .endpointId(aws_lb.example().arn())
                .weight(100)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:globalaccelerator:EndpointGroup
    properties:
      listenerArn: ${aws_globalaccelerator_listener.example.id}
      endpointConfigurations:
        - endpointId: ${aws_lb.example.arn}
          weight: 100
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Global Accelerator endpoint groups using the `id`. For example:

```sh
 $ pulumi import aws:globalaccelerator/endpointGroup:EndpointGroup example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxx/endpoint-group/xxxxxxxx
```
 