Resource for managing an AWS OpenSearchServerless VPC Endpoint.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.opensearch.ServerlessVpcEndpoint("example", {
    subnetIds: [aws_subnet.example.id],
    vpcId: aws_vpc.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.opensearch.ServerlessVpcEndpoint("example",
    subnet_ids=[aws_subnet["example"]["id"]],
    vpc_id=aws_vpc["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.OpenSearch.ServerlessVpcEndpoint("example", new()
    {
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        VpcId = aws_vpc.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opensearch.NewServerlessVpcEndpoint(ctx, "example", &opensearch.ServerlessVpcEndpointArgs{
			SubnetIds: pulumi.StringArray{
				aws_subnet.Example.Id,
			},
			VpcId: pulumi.Any(aws_vpc.Example.Id),
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
import com.pulumi.aws.opensearch.ServerlessVpcEndpoint;
import com.pulumi.aws.opensearch.ServerlessVpcEndpointArgs;
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
        var example = new ServerlessVpcEndpoint("example", ServerlessVpcEndpointArgs.builder()        
            .subnetIds(aws_subnet.example().id())
            .vpcId(aws_vpc.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:opensearch:ServerlessVpcEndpoint
    properties:
      subnetIds:
        - ${aws_subnet.example.id}
      vpcId: ${aws_vpc.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import OpenSearchServerless Vpc Endpointa using the `id`. For example:

```sh
 $ pulumi import aws:opensearch/serverlessVpcEndpoint:ServerlessVpcEndpoint example vpce-8012925589
```
 