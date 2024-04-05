Provides a resource to manage an S3 Outposts Endpoint.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3outposts.Endpoint("example", {
    outpostId: data.aws_outposts_outpost.example.id,
    securityGroupId: aws_security_group.example.id,
    subnetId: aws_subnet.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3outposts.Endpoint("example",
    outpost_id=data["aws_outposts_outpost"]["example"]["id"],
    security_group_id=aws_security_group["example"]["id"],
    subnet_id=aws_subnet["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3Outposts.Endpoint("example", new()
    {
        OutpostId = data.Aws_outposts_outpost.Example.Id,
        SecurityGroupId = aws_security_group.Example.Id,
        SubnetId = aws_subnet.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3outposts"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3outposts.NewEndpoint(ctx, "example", &s3outposts.EndpointArgs{
			OutpostId:       pulumi.Any(data.Aws_outposts_outpost.Example.Id),
			SecurityGroupId: pulumi.Any(aws_security_group.Example.Id),
			SubnetId:        pulumi.Any(aws_subnet.Example.Id),
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
import com.pulumi.aws.s3outposts.Endpoint;
import com.pulumi.aws.s3outposts.EndpointArgs;
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
        var example = new Endpoint("example", EndpointArgs.builder()        
            .outpostId(data.aws_outposts_outpost().example().id())
            .securityGroupId(aws_security_group.example().id())
            .subnetId(aws_subnet.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3outposts:Endpoint
    properties:
      outpostId: ${data.aws_outposts_outpost.example.id}
      securityGroupId: ${aws_security_group.example.id}
      subnetId: ${aws_subnet.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 Outposts Endpoints using Amazon Resource Name (ARN), EC2 Security Group identifier, and EC2 Subnet identifier, separated by commas (`,`). For example:

```sh
 $ pulumi import aws:s3outposts/endpoint:Endpoint example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/endpoint/0123456789abcdef,sg-12345678,subnet-12345678
```
 