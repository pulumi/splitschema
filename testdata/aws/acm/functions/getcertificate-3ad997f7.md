Use this data source to get the ARN of a certificate in AWS Certificate
Manager (ACM), you can reference
it by domain without having to hard code the ARNs as input.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const issued = aws.acm.getCertificate({
    domain: "tf.example.com",
    statuses: ["ISSUED"],
});
const amazonIssued = aws.acm.getCertificate({
    domain: "tf.example.com",
    mostRecent: true,
    types: ["AMAZON_ISSUED"],
});
const rsa4096 = aws.acm.getCertificate({
    domain: "tf.example.com",
    keyTypes: ["RSA_4096"],
});
```
```python
import pulumi
import pulumi_aws as aws

issued = aws.acm.get_certificate(domain="tf.example.com",
    statuses=["ISSUED"])
amazon_issued = aws.acm.get_certificate(domain="tf.example.com",
    most_recent=True,
    types=["AMAZON_ISSUED"])
rsa4096 = aws.acm.get_certificate(domain="tf.example.com",
    key_types=["RSA_4096"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var issued = Aws.Acm.GetCertificate.Invoke(new()
    {
        Domain = "tf.example.com",
        Statuses = new[]
        {
            "ISSUED",
        },
    });

    var amazonIssued = Aws.Acm.GetCertificate.Invoke(new()
    {
        Domain = "tf.example.com",
        MostRecent = true,
        Types = new[]
        {
            "AMAZON_ISSUED",
        },
    });

    var rsa4096 = Aws.Acm.GetCertificate.Invoke(new()
    {
        Domain = "tf.example.com",
        KeyTypes = new[]
        {
            "RSA_4096",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := acm.LookupCertificate(ctx, &acm.LookupCertificateArgs{
			Domain: "tf.example.com",
			Statuses: []string{
				"ISSUED",
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = acm.LookupCertificate(ctx, &acm.LookupCertificateArgs{
			Domain:     "tf.example.com",
			MostRecent: pulumi.BoolRef(true),
			Types: []string{
				"AMAZON_ISSUED",
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = acm.LookupCertificate(ctx, &acm.LookupCertificateArgs{
			Domain: "tf.example.com",
			KeyTypes: []string{
				"RSA_4096",
			},
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
import com.pulumi.aws.acm.AcmFunctions;
import com.pulumi.aws.acm.inputs.GetCertificateArgs;
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
        final var issued = AcmFunctions.getCertificate(GetCertificateArgs.builder()
            .domain("tf.example.com")
            .statuses("ISSUED")
            .build());

        final var amazonIssued = AcmFunctions.getCertificate(GetCertificateArgs.builder()
            .domain("tf.example.com")
            .mostRecent(true)
            .types("AMAZON_ISSUED")
            .build());

        final var rsa4096 = AcmFunctions.getCertificate(GetCertificateArgs.builder()
            .domain("tf.example.com")
            .keyTypes("RSA_4096")
            .build());

    }
}
```
```yaml
variables:
  issued:
    fn::invoke:
      Function: aws:acm:getCertificate
      Arguments:
        domain: tf.example.com
        statuses:
          - ISSUED
  amazonIssued:
    fn::invoke:
      Function: aws:acm:getCertificate
      Arguments:
        domain: tf.example.com
        mostRecent: true
        types:
          - AMAZON_ISSUED
  rsa4096:
    fn::invoke:
      Function: aws:acm:getCertificate
      Arguments:
        domain: tf.example.com
        keyTypes:
          - RSA_4096
```
{{% /example %}}
{{% /examples %}}