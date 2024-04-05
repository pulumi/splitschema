Provides a resource to manage an AWS Certificate Manager Private Certificate Authorities Permission.
Currently, this is only required in order to allow the ACM service to automatically renew certificates issued by a PCA.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCertificateAuthority = new aws.acmpca.CertificateAuthority("exampleCertificateAuthority", {certificateAuthorityConfiguration: {
    keyAlgorithm: "RSA_4096",
    signingAlgorithm: "SHA512WITHRSA",
    subject: {
        commonName: "example.com",
    },
}});
const examplePermission = new aws.acmpca.Permission("examplePermission", {
    certificateAuthorityArn: exampleCertificateAuthority.arn,
    actions: [
        "IssueCertificate",
        "GetCertificate",
        "ListPermissions",
    ],
    principal: "acm.amazonaws.com",
});
```
```python
import pulumi
import pulumi_aws as aws

example_certificate_authority = aws.acmpca.CertificateAuthority("exampleCertificateAuthority", certificate_authority_configuration=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs(
    key_algorithm="RSA_4096",
    signing_algorithm="SHA512WITHRSA",
    subject=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs(
        common_name="example.com",
    ),
))
example_permission = aws.acmpca.Permission("examplePermission",
    certificate_authority_arn=example_certificate_authority.arn,
    actions=[
        "IssueCertificate",
        "GetCertificate",
        "ListPermissions",
    ],
    principal="acm.amazonaws.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleCertificateAuthority = new Aws.Acmpca.CertificateAuthority("exampleCertificateAuthority", new()
    {
        CertificateAuthorityConfiguration = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationArgs
        {
            KeyAlgorithm = "RSA_4096",
            SigningAlgorithm = "SHA512WITHRSA",
            Subject = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs
            {
                CommonName = "example.com",
            },
        },
    });

    var examplePermission = new Aws.Acmpca.Permission("examplePermission", new()
    {
        CertificateAuthorityArn = exampleCertificateAuthority.Arn,
        Actions = new[]
        {
            "IssueCertificate",
            "GetCertificate",
            "ListPermissions",
        },
        Principal = "acm.amazonaws.com",
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
		exampleCertificateAuthority, err := acmpca.NewCertificateAuthority(ctx, "exampleCertificateAuthority", &acmpca.CertificateAuthorityArgs{
			CertificateAuthorityConfiguration: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs{
				KeyAlgorithm:     pulumi.String("RSA_4096"),
				SigningAlgorithm: pulumi.String("SHA512WITHRSA"),
				Subject: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs{
					CommonName: pulumi.String("example.com"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = acmpca.NewPermission(ctx, "examplePermission", &acmpca.PermissionArgs{
			CertificateAuthorityArn: exampleCertificateAuthority.Arn,
			Actions: pulumi.StringArray{
				pulumi.String("IssueCertificate"),
				pulumi.String("GetCertificate"),
				pulumi.String("ListPermissions"),
			},
			Principal: pulumi.String("acm.amazonaws.com"),
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
import com.pulumi.aws.acmpca.CertificateAuthority;
import com.pulumi.aws.acmpca.CertificateAuthorityArgs;
import com.pulumi.aws.acmpca.inputs.CertificateAuthorityCertificateAuthorityConfigurationArgs;
import com.pulumi.aws.acmpca.inputs.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs;
import com.pulumi.aws.acmpca.Permission;
import com.pulumi.aws.acmpca.PermissionArgs;
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
        var exampleCertificateAuthority = new CertificateAuthority("exampleCertificateAuthority", CertificateAuthorityArgs.builder()        
            .certificateAuthorityConfiguration(CertificateAuthorityCertificateAuthorityConfigurationArgs.builder()
                .keyAlgorithm("RSA_4096")
                .signingAlgorithm("SHA512WITHRSA")
                .subject(CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs.builder()
                    .commonName("example.com")
                    .build())
                .build())
            .build());

        var examplePermission = new Permission("examplePermission", PermissionArgs.builder()        
            .certificateAuthorityArn(exampleCertificateAuthority.arn())
            .actions(            
                "IssueCertificate",
                "GetCertificate",
                "ListPermissions")
            .principal("acm.amazonaws.com")
            .build());

    }
}
```
```yaml
resources:
  examplePermission:
    type: aws:acmpca:Permission
    properties:
      certificateAuthorityArn: ${exampleCertificateAuthority.arn}
      actions:
        - IssueCertificate
        - GetCertificate
        - ListPermissions
      principal: acm.amazonaws.com
  exampleCertificateAuthority:
    type: aws:acmpca:CertificateAuthority
    properties:
      certificateAuthorityConfiguration:
        keyAlgorithm: RSA_4096
        signingAlgorithm: SHA512WITHRSA
        subject:
          commonName: example.com
```
{{% /example %}}
{{% /examples %}}