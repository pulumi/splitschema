Manages an EC2 Instance Connect Endpoint.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.InstanceConnectEndpoint("example", {subnetId: aws_subnet.example.id});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.InstanceConnectEndpoint("example", subnet_id=aws_subnet["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2TransitGateway.InstanceConnectEndpoint("example", new()
    {
        SubnetId = aws_subnet.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.NewInstanceConnectEndpoint(ctx, "example", &ec2transitgateway.InstanceConnectEndpointArgs{
			SubnetId: pulumi.Any(aws_subnet.Example.Id),
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
import com.pulumi.aws.ec2transitgateway.InstanceConnectEndpoint;
import com.pulumi.aws.ec2transitgateway.InstanceConnectEndpointArgs;
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
        var example = new InstanceConnectEndpoint("example", InstanceConnectEndpointArgs.builder()        
            .subnetId(aws_subnet.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:InstanceConnectEndpoint
    properties:
      subnetId: ${aws_subnet.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EC2 Instance Connect Endpoints using the `id`. For example:

```sh
 $ pulumi import aws:ec2transitgateway/instanceConnectEndpoint:InstanceConnectEndpoint example eice-012345678
```
 