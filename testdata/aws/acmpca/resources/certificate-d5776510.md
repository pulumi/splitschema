Provides a resource to issue a certificate using AWS Certificate Manager Private Certificate Authority (ACM PCA).

Certificates created using `aws.acmpca.Certificate` are not eligible for automatic renewal,
and must be replaced instead.
To issue a renewable certificate using an ACM PCA, create a `aws.acm.Certificate`
with the parameter `certificate_authority_arn`.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.acmpca.CertificateAuthority;
import com.pulumi.aws.acmpca.CertificateAuthorityArgs;
import com.pulumi.tls.PrivateKey;
import com.pulumi.tls.PrivateKeyArgs;
import com.pulumi.tls.CertRequest;
import com.pulumi.tls.CertRequestArgs;
import com.pulumi.tls.inputs.CertRequestSubjectArgs;
import com.pulumi.aws.acmpca.Certificate;
import com.pulumi.aws.acmpca.CertificateArgs;
import com.pulumi.aws.acmpca.inputs.CertificateValidityArgs;
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
            .privateCertificateConfiguration(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .permanentDeletionTimeInDays(7)
            .build());

        var key = new PrivateKey("key", PrivateKeyArgs.builder()        
            .algorithm("RSA")
            .build());

        var csr = new CertRequest("csr", CertRequestArgs.builder()        
            .keyAlgorithm("RSA")
            .privateKeyPem(key.privateKeyPem())
            .subject(CertRequestSubjectArgs.builder()
                .commonName("example")
                .build())
            .build());

        var exampleCertificate = new Certificate("exampleCertificate", CertificateArgs.builder()        
            .certificateAuthorityArn(exampleCertificateAuthority.arn())
            .certificateSigningRequest(csr.certRequestPem())
            .signingAlgorithm("SHA256WITHRSA")
            .validity(CertificateValidityArgs.builder()
                .type("YEARS")
                .value(1)
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleCertificate:
    type: aws:acmpca:Certificate
    properties:
      certificateAuthorityArn: ${exampleCertificateAuthority.arn}
      certificateSigningRequest: ${csr.certRequestPem}
      signingAlgorithm: SHA256WITHRSA
      validity:
        type: YEARS
        value: 1
  exampleCertificateAuthority:
    type: aws:acmpca:CertificateAuthority
    properties:
      privateCertificateConfiguration:
        - keyAlgorithm: RSA_4096
          signingAlgorithm: SHA512WITHRSA
          subject:
            - commonName: example.com
      permanentDeletionTimeInDays: 7
  key:
    type: tls:PrivateKey
    properties:
      algorithm: RSA
  csr:
    type: tls:CertRequest
    properties:
      keyAlgorithm: RSA
      privateKeyPem: ${key.privateKeyPem}
      subject:
        commonName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ACM PCA Certificates using their ARN. For example:

```sh
 $ pulumi import aws:acmpca/certificate:Certificate cert arn:aws:acm-pca:eu-west-1:675225743824:certificate-authority/08319ede-83g9-1400-8f21-c7d12b2b6edb/certificate/a4e9c2aa4bcfab625g1b9136464cd3a
```
 