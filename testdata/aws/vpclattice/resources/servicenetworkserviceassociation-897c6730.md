Resource for managing an AWS VPC Lattice Service Network Service Association.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpclattice.ServiceNetworkServiceAssociation("example", {
    serviceIdentifier: aws_vpclattice_service.example.id,
    serviceNetworkIdentifier: aws_vpclattice_service_network.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.ServiceNetworkServiceAssociation("example",
    service_identifier=aws_vpclattice_service["example"]["id"],
    service_network_identifier=aws_vpclattice_service_network["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VpcLattice.ServiceNetworkServiceAssociation("example", new()
    {
        ServiceIdentifier = aws_vpclattice_service.Example.Id,
        ServiceNetworkIdentifier = aws_vpclattice_service_network.Example.Id,
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
		_, err := vpclattice.NewServiceNetworkServiceAssociation(ctx, "example", &vpclattice.ServiceNetworkServiceAssociationArgs{
			ServiceIdentifier:        pulumi.Any(aws_vpclattice_service.Example.Id),
			ServiceNetworkIdentifier: pulumi.Any(aws_vpclattice_service_network.Example.Id),
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
import com.pulumi.aws.vpclattice.ServiceNetworkServiceAssociation;
import com.pulumi.aws.vpclattice.ServiceNetworkServiceAssociationArgs;
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
        var example = new ServiceNetworkServiceAssociation("example", ServiceNetworkServiceAssociationArgs.builder()        
            .serviceIdentifier(aws_vpclattice_service.example().id())
            .serviceNetworkIdentifier(aws_vpclattice_service_network.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:vpclattice:ServiceNetworkServiceAssociation
    properties:
      serviceIdentifier: ${aws_vpclattice_service.example.id}
      serviceNetworkIdentifier: ${aws_vpclattice_service_network.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Lattice Service Network Service Association using the `id`. For example:

```sh
 $ pulumi import aws:vpclattice/serviceNetworkServiceAssociation:ServiceNetworkServiceAssociation example snsa-05e2474658a88f6ba
```
 