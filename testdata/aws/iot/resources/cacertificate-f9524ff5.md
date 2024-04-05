Creates and manages an AWS IoT CA Certificate.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";

const caPrivateKey = new tls.PrivateKey("caPrivateKey", {algorithm: "RSA"});
const caSelfSignedCert = new tls.SelfSignedCert("caSelfSignedCert", {
    privateKeyPem: caPrivateKey.privateKeyPem,
    subject: {
        commonName: "example.com",
        organization: "ACME Examples, Inc",
    },
    validityPeriodHours: 12,
    allowedUses: [
        "key_encipherment",
        "digital_signature",
        "server_auth",
    ],
    isCaCertificate: true,
});
const verificationPrivateKey = new tls.PrivateKey("verificationPrivateKey", {algorithm: "RSA"});
const exampleRegistrationCode = aws.iot.getRegistrationCode({});
const verificationCertRequest = new tls.CertRequest("verificationCertRequest", {
    privateKeyPem: verificationPrivateKey.privateKeyPem,
    subject: {
        commonName: exampleRegistrationCode.then(exampleRegistrationCode => exampleRegistrationCode.registrationCode),
    },
});
const verificationLocallySignedCert = new tls.LocallySignedCert("verificationLocallySignedCert", {
    certRequestPem: verificationCertRequest.certRequestPem,
    caPrivateKeyPem: caPrivateKey.privateKeyPem,
    caCertPem: caSelfSignedCert.certPem,
    validityPeriodHours: 12,
    allowedUses: [
        "key_encipherment",
        "digital_signature",
        "server_auth",
    ],
});
const exampleCaCertificate = new aws.iot.CaCertificate("exampleCaCertificate", {
    active: true,
    caCertificatePem: caSelfSignedCert.certPem,
    verificationCertificatePem: verificationLocallySignedCert.certPem,
    allowAutoRegistration: true,
});
```
```python
import pulumi
import pulumi_aws as aws
import pulumi_tls as tls

ca_private_key = tls.PrivateKey("caPrivateKey", algorithm="RSA")
ca_self_signed_cert = tls.SelfSignedCert("caSelfSignedCert",
    private_key_pem=ca_private_key.private_key_pem,
    subject=tls.SelfSignedCertSubjectArgs(
        common_name="example.com",
        organization="ACME Examples, Inc",
    ),
    validity_period_hours=12,
    allowed_uses=[
        "key_encipherment",
        "digital_signature",
        "server_auth",
    ],
    is_ca_certificate=True)
verification_private_key = tls.PrivateKey("verificationPrivateKey", algorithm="RSA")
example_registration_code = aws.iot.get_registration_code()
verification_cert_request = tls.CertRequest("verificationCertRequest",
    private_key_pem=verification_private_key.private_key_pem,
    subject=tls.CertRequestSubjectArgs(
        common_name=example_registration_code.registration_code,
    ))
verification_locally_signed_cert = tls.LocallySignedCert("verificationLocallySignedCert",
    cert_request_pem=verification_cert_request.cert_request_pem,
    ca_private_key_pem=ca_private_key.private_key_pem,
    ca_cert_pem=ca_self_signed_cert.cert_pem,
    validity_period_hours=12,
    allowed_uses=[
        "key_encipherment",
        "digital_signature",
        "server_auth",
    ])
example_ca_certificate = aws.iot.CaCertificate("exampleCaCertificate",
    active=True,
    ca_certificate_pem=ca_self_signed_cert.cert_pem,
    verification_certificate_pem=verification_locally_signed_cert.cert_pem,
    allow_auto_registration=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Tls = Pulumi.Tls;

return await Deployment.RunAsync(() => 
{
    var caPrivateKey = new Tls.PrivateKey("caPrivateKey", new()
    {
        Algorithm = "RSA",
    });

    var caSelfSignedCert = new Tls.SelfSignedCert("caSelfSignedCert", new()
    {
        PrivateKeyPem = caPrivateKey.PrivateKeyPem,
        Subject = new Tls.Inputs.SelfSignedCertSubjectArgs
        {
            CommonName = "example.com",
            Organization = "ACME Examples, Inc",
        },
        ValidityPeriodHours = 12,
        AllowedUses = new[]
        {
            "key_encipherment",
            "digital_signature",
            "server_auth",
        },
        IsCaCertificate = true,
    });

    var verificationPrivateKey = new Tls.PrivateKey("verificationPrivateKey", new()
    {
        Algorithm = "RSA",
    });

    var exampleRegistrationCode = Aws.Iot.GetRegistrationCode.Invoke();

    var verificationCertRequest = new Tls.CertRequest("verificationCertRequest", new()
    {
        PrivateKeyPem = verificationPrivateKey.PrivateKeyPem,
        Subject = new Tls.Inputs.CertRequestSubjectArgs
        {
            CommonName = exampleRegistrationCode.Apply(getRegistrationCodeResult => getRegistrationCodeResult.RegistrationCode),
        },
    });

    var verificationLocallySignedCert = new Tls.LocallySignedCert("verificationLocallySignedCert", new()
    {
        CertRequestPem = verificationCertRequest.CertRequestPem,
        CaPrivateKeyPem = caPrivateKey.PrivateKeyPem,
        CaCertPem = caSelfSignedCert.CertPem,
        ValidityPeriodHours = 12,
        AllowedUses = new[]
        {
            "key_encipherment",
            "digital_signature",
            "server_auth",
        },
    });

    var exampleCaCertificate = new Aws.Iot.CaCertificate("exampleCaCertificate", new()
    {
        Active = true,
        CaCertificatePem = caSelfSignedCert.CertPem,
        VerificationCertificatePem = verificationLocallySignedCert.CertPem,
        AllowAutoRegistration = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		caPrivateKey, err := tls.NewPrivateKey(ctx, "caPrivateKey", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA"),
		})
		if err != nil {
			return err
		}
		caSelfSignedCert, err := tls.NewSelfSignedCert(ctx, "caSelfSignedCert", &tls.SelfSignedCertArgs{
			PrivateKeyPem: caPrivateKey.PrivateKeyPem,
			Subject: &tls.SelfSignedCertSubjectArgs{
				CommonName:   pulumi.String("example.com"),
				Organization: pulumi.String("ACME Examples, Inc"),
			},
			ValidityPeriodHours: pulumi.Int(12),
			AllowedUses: pulumi.StringArray{
				pulumi.String("key_encipherment"),
				pulumi.String("digital_signature"),
				pulumi.String("server_auth"),
			},
			IsCaCertificate: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		verificationPrivateKey, err := tls.NewPrivateKey(ctx, "verificationPrivateKey", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA"),
		})
		if err != nil {
			return err
		}
		exampleRegistrationCode, err := iot.GetRegistrationCode(ctx, nil, nil)
		if err != nil {
			return err
		}
		verificationCertRequest, err := tls.NewCertRequest(ctx, "verificationCertRequest", &tls.CertRequestArgs{
			PrivateKeyPem: verificationPrivateKey.PrivateKeyPem,
			Subject: &tls.CertRequestSubjectArgs{
				CommonName: *pulumi.String(exampleRegistrationCode.RegistrationCode),
			},
		})
		if err != nil {
			return err
		}
		verificationLocallySignedCert, err := tls.NewLocallySignedCert(ctx, "verificationLocallySignedCert", &tls.LocallySignedCertArgs{
			CertRequestPem:      verificationCertRequest.CertRequestPem,
			CaPrivateKeyPem:     caPrivateKey.PrivateKeyPem,
			CaCertPem:           caSelfSignedCert.CertPem,
			ValidityPeriodHours: pulumi.Int(12),
			AllowedUses: pulumi.StringArray{
				pulumi.String("key_encipherment"),
				pulumi.String("digital_signature"),
				pulumi.String("server_auth"),
			},
		})
		if err != nil {
			return err
		}
		_, err = iot.NewCaCertificate(ctx, "exampleCaCertificate", &iot.CaCertificateArgs{
			Active:                     pulumi.Bool(true),
			CaCertificatePem:           caSelfSignedCert.CertPem,
			VerificationCertificatePem: verificationLocallySignedCert.CertPem,
			AllowAutoRegistration:      pulumi.Bool(true),
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
import com.pulumi.tls.PrivateKey;
import com.pulumi.tls.PrivateKeyArgs;
import com.pulumi.tls.SelfSignedCert;
import com.pulumi.tls.SelfSignedCertArgs;
import com.pulumi.tls.inputs.SelfSignedCertSubjectArgs;
import com.pulumi.aws.iot.IotFunctions;
import com.pulumi.tls.CertRequest;
import com.pulumi.tls.CertRequestArgs;
import com.pulumi.tls.inputs.CertRequestSubjectArgs;
import com.pulumi.tls.LocallySignedCert;
import com.pulumi.tls.LocallySignedCertArgs;
import com.pulumi.aws.iot.CaCertificate;
import com.pulumi.aws.iot.CaCertificateArgs;
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
        var caPrivateKey = new PrivateKey("caPrivateKey", PrivateKeyArgs.builder()        
            .algorithm("RSA")
            .build());

        var caSelfSignedCert = new SelfSignedCert("caSelfSignedCert", SelfSignedCertArgs.builder()        
            .privateKeyPem(caPrivateKey.privateKeyPem())
            .subject(SelfSignedCertSubjectArgs.builder()
                .commonName("example.com")
                .organization("ACME Examples, Inc")
                .build())
            .validityPeriodHours(12)
            .allowedUses(            
                "key_encipherment",
                "digital_signature",
                "server_auth")
            .isCaCertificate(true)
            .build());

        var verificationPrivateKey = new PrivateKey("verificationPrivateKey", PrivateKeyArgs.builder()        
            .algorithm("RSA")
            .build());

        final var exampleRegistrationCode = IotFunctions.getRegistrationCode();

        var verificationCertRequest = new CertRequest("verificationCertRequest", CertRequestArgs.builder()        
            .privateKeyPem(verificationPrivateKey.privateKeyPem())
            .subject(CertRequestSubjectArgs.builder()
                .commonName(exampleRegistrationCode.applyValue(getRegistrationCodeResult -> getRegistrationCodeResult.registrationCode()))
                .build())
            .build());

        var verificationLocallySignedCert = new LocallySignedCert("verificationLocallySignedCert", LocallySignedCertArgs.builder()        
            .certRequestPem(verificationCertRequest.certRequestPem())
            .caPrivateKeyPem(caPrivateKey.privateKeyPem())
            .caCertPem(caSelfSignedCert.certPem())
            .validityPeriodHours(12)
            .allowedUses(            
                "key_encipherment",
                "digital_signature",
                "server_auth")
            .build());

        var exampleCaCertificate = new CaCertificate("exampleCaCertificate", CaCertificateArgs.builder()        
            .active(true)
            .caCertificatePem(caSelfSignedCert.certPem())
            .verificationCertificatePem(verificationLocallySignedCert.certPem())
            .allowAutoRegistration(true)
            .build());

    }
}
```
```yaml
resources:
  caSelfSignedCert:
    type: tls:SelfSignedCert
    properties:
      privateKeyPem: ${caPrivateKey.privateKeyPem}
      subject:
        commonName: example.com
        organization: ACME Examples, Inc
      validityPeriodHours: 12
      allowedUses:
        - key_encipherment
        - digital_signature
        - server_auth
      isCaCertificate: true
  caPrivateKey:
    type: tls:PrivateKey
    properties:
      algorithm: RSA
  verificationCertRequest:
    type: tls:CertRequest
    properties:
      privateKeyPem: ${verificationPrivateKey.privateKeyPem}
      subject:
        commonName: ${exampleRegistrationCode.registrationCode}
  verificationPrivateKey:
    type: tls:PrivateKey
    properties:
      algorithm: RSA
  verificationLocallySignedCert:
    type: tls:LocallySignedCert
    properties:
      certRequestPem: ${verificationCertRequest.certRequestPem}
      caPrivateKeyPem: ${caPrivateKey.privateKeyPem}
      caCertPem: ${caSelfSignedCert.certPem}
      validityPeriodHours: 12
      allowedUses:
        - key_encipherment
        - digital_signature
        - server_auth
  exampleCaCertificate:
    type: aws:iot:CaCertificate
    properties:
      active: true
      caCertificatePem: ${caSelfSignedCert.certPem}
      verificationCertificatePem: ${verificationLocallySignedCert.certPem}
      allowAutoRegistration: true
variables:
  exampleRegistrationCode:
    fn::invoke:
      Function: aws:iot:getRegistrationCode
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}