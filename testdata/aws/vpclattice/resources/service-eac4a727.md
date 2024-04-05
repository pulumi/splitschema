Resource for managing an AWS VPC Lattice Service.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpclattice.Service("example", {
    authType: "AWS_IAM",
    customDomainName: "example.com",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.Service("example",
    auth_type="AWS_IAM",
    custom_domain_name="example.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VpcLattice.Service("example", new()
    {
        AuthType = "AWS_IAM",
        CustomDomainName = "example.com",
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
		_, err := vpclattice.NewService(ctx, "example", &vpclattice.ServiceArgs{
			AuthType:         pulumi.String("AWS_IAM"),
			CustomDomainName: pulumi.String("example.com"),
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
import com.pulumi.aws.vpclattice.Service;
import com.pulumi.aws.vpclattice.ServiceArgs;
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
        var example = new Service("example", ServiceArgs.builder()        
            .authType("AWS_IAM")
            .customDomainName("example.com")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:vpclattice:Service
    properties:
      authType: AWS_IAM
      customDomainName: example.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Lattice Service using the `id`. For example:

```sh
 $ pulumi import aws:vpclattice/service:Service example svc-06728e2357ea55f8a
```
 