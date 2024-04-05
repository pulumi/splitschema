Get information on a Certificate issued by a AWS Certificate Manager Private Certificate Authority.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.acmpca.getCertificate({
    arn: "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012/certificate/1234b4a0d73e2056789bdbe77d5b1a23",
    certificateAuthorityArn: "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.acmpca.get_certificate(arn="arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012/certificate/1234b4a0d73e2056789bdbe77d5b1a23",
    certificate_authority_arn="arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Acmpca.GetCertificate.Invoke(new()
    {
        Arn = "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012/certificate/1234b4a0d73e2056789bdbe77d5b1a23",
        CertificateAuthorityArn = "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acmpca"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := acmpca.LookupCertificate(ctx, &acmpca.LookupCertificateArgs{
			Arn:                     "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012/certificate/1234b4a0d73e2056789bdbe77d5b1a23",
			CertificateAuthorityArn: "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012",
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
import com.pulumi.aws.acmpca.AcmpcaFunctions;
import com.pulumi.aws.acmpca.inputs.GetCertificateArgs;
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
        final var example = AcmpcaFunctions.getCertificate(GetCertificateArgs.builder()
            .arn("arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012/certificate/1234b4a0d73e2056789bdbe77d5b1a23")
            .certificateAuthorityArn("arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:acmpca:getCertificate
      Arguments:
        arn: arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012/certificate/1234b4a0d73e2056789bdbe77d5b1a23
        certificateAuthorityArn: arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
```
{{% /example %}}
{{% /examples %}}