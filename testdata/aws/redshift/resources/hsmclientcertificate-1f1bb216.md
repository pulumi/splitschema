Creates an HSM client certificate that an Amazon Redshift cluster will use to connect to the client's HSM in order to store and retrieve the keys used to encrypt the cluster databases.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.HsmClientCertificate("example", {hsmClientCertificateIdentifier: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.HsmClientCertificate("example", hsm_client_certificate_identifier="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedShift.HsmClientCertificate("example", new()
    {
        HsmClientCertificateIdentifier = "example",
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
		_, err := redshift.NewHsmClientCertificate(ctx, "example", &redshift.HsmClientCertificateArgs{
			HsmClientCertificateIdentifier: pulumi.String("example"),
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
import com.pulumi.aws.redshift.HsmClientCertificate;
import com.pulumi.aws.redshift.HsmClientCertificateArgs;
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
        var example = new HsmClientCertificate("example", HsmClientCertificateArgs.builder()        
            .hsmClientCertificateIdentifier("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:HsmClientCertificate
    properties:
      hsmClientCertificateIdentifier: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift HSM Client Certificates using `hsm_client_certificate_identifier`. For example:

```sh
 $ pulumi import aws:redshift/hsmClientCertificate:HsmClientCertificate test example
```
 