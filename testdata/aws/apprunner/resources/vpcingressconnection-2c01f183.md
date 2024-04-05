Manages an App Runner VPC Ingress Connection.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apprunner.VpcIngressConnection("example", {
    serviceArn: aws_apprunner_service.example.arn,
    ingressVpcConfiguration: {
        vpcId: aws_default_vpc["default"].id,
        vpcEndpointId: aws_vpc_endpoint.apprunner.id,
    },
    tags: {
        foo: "bar",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apprunner.VpcIngressConnection("example",
    service_arn=aws_apprunner_service["example"]["arn"],
    ingress_vpc_configuration=aws.apprunner.VpcIngressConnectionIngressVpcConfigurationArgs(
        vpc_id=aws_default_vpc["default"]["id"],
        vpc_endpoint_id=aws_vpc_endpoint["apprunner"]["id"],
    ),
    tags={
        "foo": "bar",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppRunner.VpcIngressConnection("example", new()
    {
        ServiceArn = aws_apprunner_service.Example.Arn,
        IngressVpcConfiguration = new Aws.AppRunner.Inputs.VpcIngressConnectionIngressVpcConfigurationArgs
        {
            VpcId = aws_default_vpc.Default.Id,
            VpcEndpointId = aws_vpc_endpoint.Apprunner.Id,
        },
        Tags = 
        {
            { "foo", "bar" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apprunner"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apprunner.NewVpcIngressConnection(ctx, "example", &apprunner.VpcIngressConnectionArgs{
			ServiceArn: pulumi.Any(aws_apprunner_service.Example.Arn),
			IngressVpcConfiguration: &apprunner.VpcIngressConnectionIngressVpcConfigurationArgs{
				VpcId:         pulumi.Any(aws_default_vpc.Default.Id),
				VpcEndpointId: pulumi.Any(aws_vpc_endpoint.Apprunner.Id),
			},
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
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
import com.pulumi.aws.apprunner.VpcIngressConnection;
import com.pulumi.aws.apprunner.VpcIngressConnectionArgs;
import com.pulumi.aws.apprunner.inputs.VpcIngressConnectionIngressVpcConfigurationArgs;
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
        var example = new VpcIngressConnection("example", VpcIngressConnectionArgs.builder()        
            .serviceArn(aws_apprunner_service.example().arn())
            .ingressVpcConfiguration(VpcIngressConnectionIngressVpcConfigurationArgs.builder()
                .vpcId(aws_default_vpc.default().id())
                .vpcEndpointId(aws_vpc_endpoint.apprunner().id())
                .build())
            .tags(Map.of("foo", "bar"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apprunner:VpcIngressConnection
    properties:
      serviceArn: ${aws_apprunner_service.example.arn}
      ingressVpcConfiguration:
        vpcId: ${aws_default_vpc.default.id}
        vpcEndpointId: ${aws_vpc_endpoint.apprunner.id}
      tags:
        foo: bar
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Runner VPC Ingress Connection using the `arn`. For example:

```sh
 $ pulumi import aws:apprunner/vpcIngressConnection:VpcIngressConnection example "arn:aws:apprunner:us-west-2:837424938642:vpcingressconnection/example/b379f86381d74825832c2e82080342fa"
```
 