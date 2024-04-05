Creates an HSM configuration that contains the information required by an Amazon Redshift cluster to store and use database encryption keys in a Hardware Security Module (HSM).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.HsmConfiguration("example", {
    description: "example",
    hsmConfigurationIdentifier: "example",
    hsmIpAddress: "10.0.0.1",
    hsmPartitionName: "aws",
    hsmPartitionPassword: "example",
    hsmServerPublicCertificate: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.HsmConfiguration("example",
    description="example",
    hsm_configuration_identifier="example",
    hsm_ip_address="10.0.0.1",
    hsm_partition_name="aws",
    hsm_partition_password="example",
    hsm_server_public_certificate="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedShift.HsmConfiguration("example", new()
    {
        Description = "example",
        HsmConfigurationIdentifier = "example",
        HsmIpAddress = "10.0.0.1",
        HsmPartitionName = "aws",
        HsmPartitionPassword = "example",
        HsmServerPublicCertificate = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewHsmConfiguration(ctx, "example", &redshift.HsmConfigurationArgs{
			Description:                pulumi.String("example"),
			HsmConfigurationIdentifier: pulumi.String("example"),
			HsmIpAddress:               pulumi.String("10.0.0.1"),
			HsmPartitionName:           pulumi.String("aws"),
			HsmPartitionPassword:       pulumi.String("example"),
			HsmServerPublicCertificate: pulumi.String("example"),
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
import com.pulumi.aws.redshift.HsmConfiguration;
import com.pulumi.aws.redshift.HsmConfigurationArgs;
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
        var example = new HsmConfiguration("example", HsmConfigurationArgs.builder()        
            .description("example")
            .hsmConfigurationIdentifier("example")
            .hsmIpAddress("10.0.0.1")
            .hsmPartitionName("aws")
            .hsmPartitionPassword("example")
            .hsmServerPublicCertificate("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:HsmConfiguration
    properties:
      description: example
      hsmConfigurationIdentifier: example
      hsmIpAddress: 10.0.0.1
      hsmPartitionName: aws
      hsmPartitionPassword: example
      hsmServerPublicCertificate: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift HSM Client Certificates using `hsm_configuration_identifier`. For example:

```sh
 $ pulumi import aws:redshift/hsmConfiguration:HsmConfiguration example example
```
 