Data source for managing an AWS OpenSearch Serverless VPC Endpoint.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.opensearch.getServerlessVpcEndpoint({
    vpcEndpointId: "vpce-829a4487959e2a839",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.opensearch.get_serverless_vpc_endpoint(vpc_endpoint_id="vpce-829a4487959e2a839")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.OpenSearch.GetServerlessVpcEndpoint.Invoke(new()
    {
        VpcEndpointId = "vpce-829a4487959e2a839",
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
		_, err := opensearch.LookupServerlessVpcEndpoint(ctx, &opensearch.LookupServerlessVpcEndpointArgs{
			VpcEndpointId: "vpce-829a4487959e2a839",
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
import com.pulumi.aws.opensearch.OpensearchFunctions;
import com.pulumi.aws.opensearch.inputs.GetServerlessVpcEndpointArgs;
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
        final var example = OpensearchFunctions.getServerlessVpcEndpoint(GetServerlessVpcEndpointArgs.builder()
            .vpcEndpointId("vpce-829a4487959e2a839")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:opensearch:getServerlessVpcEndpoint
      Arguments:
        vpcEndpointId: vpce-829a4487959e2a839
```
{{% /example %}}
{{% /examples %}}