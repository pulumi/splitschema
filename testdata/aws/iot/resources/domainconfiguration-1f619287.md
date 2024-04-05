Creates and manages an AWS IoT domain configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const iot = new aws.iot.DomainConfiguration("iot", {
    domainName: "iot.example.com",
    serviceType: "DATA",
    serverCertificateArns: [aws_acm_certificate.cert.arn],
});
```
```python
import pulumi
import pulumi_aws as aws

iot = aws.iot.DomainConfiguration("iot",
    domain_name="iot.example.com",
    service_type="DATA",
    server_certificate_arns=[aws_acm_certificate["cert"]["arn"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var iot = new Aws.Iot.DomainConfiguration("iot", new()
    {
        DomainName = "iot.example.com",
        ServiceType = "DATA",
        ServerCertificateArns = new[]
        {
            aws_acm_certificate.Cert.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iot.NewDomainConfiguration(ctx, "iot", &iot.DomainConfigurationArgs{
			DomainName:  pulumi.String("iot.example.com"),
			ServiceType: pulumi.String("DATA"),
			ServerCertificateArns: pulumi.StringArray{
				aws_acm_certificate.Cert.Arn,
			},
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
import com.pulumi.aws.iot.DomainConfiguration;
import com.pulumi.aws.iot.DomainConfigurationArgs;
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
        var iot = new DomainConfiguration("iot", DomainConfigurationArgs.builder()        
            .domainName("iot.example.com")
            .serviceType("DATA")
            .serverCertificateArns(aws_acm_certificate.cert().arn())
            .build());

    }
}
```
```yaml
resources:
  iot:
    type: aws:iot:DomainConfiguration
    properties:
      domainName: iot.example.com
      serviceType: DATA
      serverCertificateArns:
        - ${aws_acm_certificate.cert.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import domain configurations using the name. For example:

```sh
 $ pulumi import aws:iot/domainConfiguration:DomainConfiguration example example
```
 