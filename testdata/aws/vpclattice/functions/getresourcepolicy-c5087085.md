Data source for managing an AWS VPC Lattice Resource Policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.vpclattice.getResourcePolicy({
    resourceArn: aws_vpclattice_service_network.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.get_resource_policy(resource_arn=aws_vpclattice_service_network["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.VpcLattice.GetResourcePolicy.Invoke(new()
    {
        ResourceArn = aws_vpclattice_service_network.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.LookupResourcePolicy(ctx, &vpclattice.LookupResourcePolicyArgs{
			ResourceArn: aws_vpclattice_service_network.Example.Arn,
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
import com.pulumi.aws.vpclattice.VpclatticeFunctions;
import com.pulumi.aws.vpclattice.inputs.GetResourcePolicyArgs;
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
        final var example = VpclatticeFunctions.getResourcePolicy(GetResourcePolicyArgs.builder()
            .resourceArn(aws_vpclattice_service_network.example().arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:vpclattice:getResourcePolicy
      Arguments:
        resourceArn: ${aws_vpclattice_service_network.example.arn}
```
{{% /example %}}
{{% /examples %}}