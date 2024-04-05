Data source for managing an AWS VPC Lattice Auth Policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.vpclattice.getAuthPolicy({
    resourceIdentifier: aws_vpclattice_auth_policy.test.resource_identifier,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.vpclattice.get_auth_policy(resource_identifier=aws_vpclattice_auth_policy["test"]["resource_identifier"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.VpcLattice.GetAuthPolicy.Invoke(new()
    {
        ResourceIdentifier = aws_vpclattice_auth_policy.Test.Resource_identifier,
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
		_, err := vpclattice.LookupAuthPolicy(ctx, &vpclattice.LookupAuthPolicyArgs{
			ResourceIdentifier: aws_vpclattice_auth_policy.Test.Resource_identifier,
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
import com.pulumi.aws.vpclattice.inputs.GetAuthPolicyArgs;
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
        final var test = VpclatticeFunctions.getAuthPolicy(GetAuthPolicyArgs.builder()
            .resourceIdentifier(aws_vpclattice_auth_policy.test().resource_identifier())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:vpclattice:getAuthPolicy
      Arguments:
        resourceIdentifier: ${aws_vpclattice_auth_policy.test.resource_identifier}
```
{{% /example %}}
{{% /examples %}}