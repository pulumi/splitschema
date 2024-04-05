Manages an Amazon API Gateway Version 2 VPC Link.

> **Note:** Amazon API Gateway Version 2 VPC Links enable private integrations that connect HTTP APIs to private resources in a VPC.
To enable private integration for REST APIs, use the Amazon API Gateway Version 1 VPC Link resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.VpcLink("example", {
    securityGroupIds: [data.aws_security_group.example.id],
    subnetIds: data.aws_subnets.example.ids,
    tags: {
        Usage: "example",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.VpcLink("example",
    security_group_ids=[data["aws_security_group"]["example"]["id"]],
    subnet_ids=data["aws_subnets"]["example"]["ids"],
    tags={
        "Usage": "example",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.VpcLink("example", new()
    {
        SecurityGroupIds = new[]
        {
            data.Aws_security_group.Example.Id,
        },
        SubnetIds = data.Aws_subnets.Example.Ids,
        Tags = 
        {
            { "Usage", "example" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewVpcLink(ctx, "example", &apigatewayv2.VpcLinkArgs{
			SecurityGroupIds: pulumi.StringArray{
				data.Aws_security_group.Example.Id,
			},
			SubnetIds: pulumi.Any(data.Aws_subnets.Example.Ids),
			Tags: pulumi.StringMap{
				"Usage": pulumi.String("example"),
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
import com.pulumi.aws.apigatewayv2.VpcLink;
import com.pulumi.aws.apigatewayv2.VpcLinkArgs;
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
        var example = new VpcLink("example", VpcLinkArgs.builder()        
            .securityGroupIds(data.aws_security_group().example().id())
            .subnetIds(data.aws_subnets().example().ids())
            .tags(Map.of("Usage", "example"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:VpcLink
    properties:
      securityGroupIds:
        - ${data.aws_security_group.example.id}
      subnetIds: ${data.aws_subnets.example.ids}
      tags:
        Usage: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_apigatewayv2_vpc_link` using the VPC Link identifier. For example:

```sh
 $ pulumi import aws:apigatewayv2/vpcLink:VpcLink example aabbccddee
```
 