Provides an IAM Signing Certificate resource to upload Signing Certificates.

> **Note:** All arguments including the certificate body will be stored in the raw state as plain-text.
{{% examples %}}
## Example Usage
{{% example %}}

**Using certs on file:**

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const testCert = new aws.iam.SigningCertificate("testCert", {
    username: "some_test_cert",
    certificateBody: fs.readFileSync("self-ca-cert.pem", "utf8"),
});
```
```python
import pulumi
import pulumi_aws as aws

test_cert = aws.iam.SigningCertificate("testCert",
    username="some_test_cert",
    certificate_body=(lambda path: open(path).read())("self-ca-cert.pem"))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCert = new Aws.Iam.SigningCertificate("testCert", new()
    {
        Username = "some_test_cert",
        CertificateBody = File.ReadAllText("self-ca-cert.pem"),
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewSigningCertificate(ctx, "testCert", &iam.SigningCertificateArgs{
			Username:        pulumi.String("some_test_cert"),
			CertificateBody: readFileOrPanic("self-ca-cert.pem"),
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
import com.pulumi.aws.iam.SigningCertificate;
import com.pulumi.aws.iam.SigningCertificateArgs;
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
        var testCert = new SigningCertificate("testCert", SigningCertificateArgs.builder()        
            .username("some_test_cert")
            .certificateBody(Files.readString(Paths.get("self-ca-cert.pem")))
            .build());

    }
}
```
```yaml
resources:
  testCert:
    type: aws:iam:SigningCertificate
    properties:
      username: some_test_cert
      certificateBody:
        fn::readFile: self-ca-cert.pem
```

**Example with cert in-line:**

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCertAlt = new aws.iam.SigningCertificate("testCertAlt", {
    certificateBody: `-----BEGIN CERTIFICATE-----
[......] # cert contents
-----END CERTIFICATE-----

`,
    username: "some_test_cert",
});
```
```python
import pulumi
import pulumi_aws as aws

test_cert_alt = aws.iam.SigningCertificate("testCertAlt",
    certificate_body="""-----BEGIN CERTIFICATE-----
[......] # cert contents
-----END CERTIFICATE-----

""",
    username="some_test_cert")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCertAlt = new Aws.Iam.SigningCertificate("testCertAlt", new()
    {
        CertificateBody = @"-----BEGIN CERTIFICATE-----
[......] # cert contents
-----END CERTIFICATE-----

",
        Username = "some_test_cert",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewSigningCertificate(ctx, "testCertAlt", &iam.SigningCertificateArgs{
			CertificateBody: pulumi.String("-----BEGIN CERTIFICATE-----\n[......] # cert contents\n-----END CERTIFICATE-----\n\n"),
			Username:        pulumi.String("some_test_cert"),
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
import com.pulumi.aws.iam.SigningCertificate;
import com.pulumi.aws.iam.SigningCertificateArgs;
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
        var testCertAlt = new SigningCertificate("testCertAlt", SigningCertificateArgs.builder()        
            .certificateBody("""
-----BEGIN CERTIFICATE-----
[......] # cert contents
-----END CERTIFICATE-----

            """)
            .username("some_test_cert")
            .build());

    }
}
```
```yaml
resources:
  testCertAlt:
    type: aws:iam:SigningCertificate
    properties:
      certificateBody: |+
        -----BEGIN CERTIFICATE-----
        [......] # cert contents
        -----END CERTIFICATE-----

      username: some_test_cert
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM Signing Certificates using the `id`. For example:

```sh
 $ pulumi import aws:iam/signingCertificate:SigningCertificate certificate IDIDIDIDID:user-name
```
 