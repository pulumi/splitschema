Resource for managing an AWS VPC Lattice Service Network.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpclattice.ServiceNetwork("example", {authType: "AWS_IAM"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.ServiceNetwork("example", auth_type="AWS_IAM")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VpcLattice.ServiceNetwork("example", new()
    {
        AuthType = "AWS_IAM",
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
		_, err := vpclattice.NewServiceNetwork(ctx, "example", &vpclattice.ServiceNetworkArgs{
			AuthType: pulumi.String("AWS_IAM"),
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
import com.pulumi.aws.vpclattice.ServiceNetwork;
import com.pulumi.aws.vpclattice.ServiceNetworkArgs;
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
        var example = new ServiceNetwork("example", ServiceNetworkArgs.builder()        
            .authType("AWS_IAM")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:vpclattice:ServiceNetwork
    properties:
      authType: AWS_IAM
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Lattice Service Network using the `id`. For example:

```sh
 $ pulumi import aws:vpclattice/serviceNetwork:ServiceNetwork example sn-0158f91c1e3358dba
```
 