Provides a lightsail certificate.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lightsail.Certificate("test", {
    domainName: "testdomain.com",
    subjectAlternativeNames: ["www.testdomain.com"],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lightsail.Certificate("test",
    domain_name="testdomain.com",
    subject_alternative_names=["www.testdomain.com"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LightSail.Certificate("test", new()
    {
        DomainName = "testdomain.com",
        SubjectAlternativeNames = new[]
        {
            "www.testdomain.com",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewCertificate(ctx, "test", &lightsail.CertificateArgs{
			DomainName: pulumi.String("testdomain.com"),
			SubjectAlternativeNames: pulumi.StringArray{
				pulumi.String("www.testdomain.com"),
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
import com.pulumi.aws.lightsail.Certificate;
import com.pulumi.aws.lightsail.CertificateArgs;
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
        var test = new Certificate("test", CertificateArgs.builder()        
            .domainName("testdomain.com")
            .subjectAlternativeNames("www.testdomain.com")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lightsail:Certificate
    properties:
      domainName: testdomain.com
      subjectAlternativeNames:
        - www.testdomain.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_lightsail_certificate` using the certificate name. For example:

```sh
 $ pulumi import aws:lightsail/certificate:Certificate test CertificateName
```
 