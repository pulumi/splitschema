Gets a registration code used to register a CA certificate with AWS IoT.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";

const example = aws.iot.getRegistrationCode({});
const verificationPrivateKey = new tls.PrivateKey("verificationPrivateKey", {algorithm: "RSA"});
const verificationCertRequest = new tls.CertRequest("verificationCertRequest", {
    keyAlgorithm: "RSA",
    privateKeyPem: verificationPrivateKey.privateKeyPem,
    subject: {
        commonName: example.then(example => example.registrationCode),
    },
});
```
```python
import pulumi
import pulumi_aws as aws
import pulumi_tls as tls

example = aws.iot.get_registration_code()
verification_private_key = tls.PrivateKey("verificationPrivateKey", algorithm="RSA")
verification_cert_request = tls.CertRequest("verificationCertRequest",
    key_algorithm="RSA",
    private_key_pem=verification_private_key.private_key_pem,
    subject=tls.CertRequestSubjectArgs(
        common_name=example.registration_code,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Tls = Pulumi.Tls;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Iot.GetRegistrationCode.Invoke();

    var verificationPrivateKey = new Tls.PrivateKey("verificationPrivateKey", new()
    {
        Algorithm = "RSA",
    });

    var verificationCertRequest = new Tls.CertRequest("verificationCertRequest", new()
    {
        KeyAlgorithm = "RSA",
        PrivateKeyPem = verificationPrivateKey.PrivateKeyPem,
        Subject = new Tls.Inputs.CertRequestSubjectArgs
        {
            CommonName = example.Apply(getRegistrationCodeResult => getRegistrationCodeResult.RegistrationCode),
        },
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
		example, err := iot.GetRegistrationCode(ctx, nil, nil)
		if err != nil {
			return err
		}
		verificationPrivateKey, err := tls.NewPrivateKey(ctx, "verificationPrivateKey", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA"),
		})
		if err != nil {
			return err
		}
		_, err = tls.NewCertRequest(ctx, "verificationCertRequest", &tls.CertRequestArgs{
			KeyAlgorithm:  pulumi.String("RSA"),
			PrivateKeyPem: verificationPrivateKey.PrivateKeyPem,
			Subject: &tls.CertRequestSubjectArgs{
				CommonName: *pulumi.String(example.RegistrationCode),
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
import com.pulumi.aws.iot.IotFunctions;
import com.pulumi.tls.PrivateKey;
import com.pulumi.tls.PrivateKeyArgs;
import com.pulumi.tls.CertRequest;
import com.pulumi.tls.CertRequestArgs;
import com.pulumi.tls.inputs.CertRequestSubjectArgs;
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
        final var example = IotFunctions.getRegistrationCode();

        var verificationPrivateKey = new PrivateKey("verificationPrivateKey", PrivateKeyArgs.builder()        
            .algorithm("RSA")
            .build());

        var verificationCertRequest = new CertRequest("verificationCertRequest", CertRequestArgs.builder()        
            .keyAlgorithm("RSA")
            .privateKeyPem(verificationPrivateKey.privateKeyPem())
            .subject(CertRequestSubjectArgs.builder()
                .commonName(example.applyValue(getRegistrationCodeResult -> getRegistrationCodeResult.registrationCode()))
                .build())
            .build());

    }
}
```
```yaml
resources:
  verificationPrivateKey:
    type: tls:PrivateKey
    properties:
      algorithm: RSA
  verificationCertRequest:
    type: tls:CertRequest
    properties:
      keyAlgorithm: RSA
      privateKeyPem: ${verificationPrivateKey.privateKeyPem}
      subject:
        commonName: ${example.registrationCode}
variables:
  example:
    fn::invoke:
      Function: aws:iot:getRegistrationCode
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}