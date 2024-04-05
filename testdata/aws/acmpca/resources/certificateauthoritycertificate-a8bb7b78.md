Associates a certificate with an AWS Certificate Manager Private Certificate Authority (ACM PCA Certificate Authority). An ACM PCA Certificate Authority is unable to issue certificates until it has a certificate associated with it. A root level ACM PCA Certificate Authority is able to self-sign its own root certificate.

{{% examples %}}
## Example Usage
{{% example %}}
### Self-Signed Root Certificate Authority Certificate

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCertificateAuthority = new aws.acmpca.CertificateAuthority("exampleCertificateAuthority", {
    type: "ROOT",
    certificateAuthorityConfiguration: {
        keyAlgorithm: "RSA_4096",
        signingAlgorithm: "SHA512WITHRSA",
        subject: {
            commonName: "example.com",
        },
    },
});
const current = aws.getPartition({});
const exampleCertificate = new aws.acmpca.Certificate("exampleCertificate", {
    certificateAuthorityArn: exampleCertificateAuthority.arn,
    certificateSigningRequest: exampleCertificateAuthority.certificateSigningRequest,
    signingAlgorithm: "SHA512WITHRSA",
    templateArn: current.then(current => `arn:${current.partition}:acm-pca:::template/RootCACertificate/V1`),
    validity: {
        type: "YEARS",
        value: "1",
    },
});
const exampleCertificateAuthorityCertificate = new aws.acmpca.CertificateAuthorityCertificate("exampleCertificateAuthorityCertificate", {
    certificateAuthorityArn: exampleCertificateAuthority.arn,
    certificate: exampleCertificate.certificate,
    certificateChain: exampleCertificate.certificateChain,
});
```
```python
import pulumi
import pulumi_aws as aws

example_certificate_authority = aws.acmpca.CertificateAuthority("exampleCertificateAuthority",
    type="ROOT",
    certificate_authority_configuration=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs(
        key_algorithm="RSA_4096",
        signing_algorithm="SHA512WITHRSA",
        subject=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs(
            common_name="example.com",
        ),
    ))
current = aws.get_partition()
example_certificate = aws.acmpca.Certificate("exampleCertificate",
    certificate_authority_arn=example_certificate_authority.arn,
    certificate_signing_request=example_certificate_authority.certificate_signing_request,
    signing_algorithm="SHA512WITHRSA",
    template_arn=f"arn:{current.partition}:acm-pca:::template/RootCACertificate/V1",
    validity=aws.acmpca.CertificateValidityArgs(
        type="YEARS",
        value="1",
    ))
example_certificate_authority_certificate = aws.acmpca.CertificateAuthorityCertificate("exampleCertificateAuthorityCertificate",
    certificate_authority_arn=example_certificate_authority.arn,
    certificate=example_certificate.certificate,
    certificate_chain=example_certificate.certificate_chain)
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
        Type = "ROOT",
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

    var current = Aws.GetPartition.Invoke();

    var exampleCertificate = new Aws.Acmpca.Certificate("exampleCertificate", new()
    {
        CertificateAuthorityArn = exampleCertificateAuthority.Arn,
        CertificateSigningRequest = exampleCertificateAuthority.CertificateSigningRequest,
        SigningAlgorithm = "SHA512WITHRSA",
        TemplateArn = $"arn:{current.Apply(getPartitionResult => getPartitionResult.Partition)}:acm-pca:::template/RootCACertificate/V1",
        Validity = new Aws.Acmpca.Inputs.CertificateValidityArgs
        {
            Type = "YEARS",
            Value = "1",
        },
    });

    var exampleCertificateAuthorityCertificate = new Aws.Acmpca.CertificateAuthorityCertificate("exampleCertificateAuthorityCertificate", new()
    {
        CertificateAuthorityArn = exampleCertificateAuthority.Arn,
        Certificate = exampleCertificate.Certificate,
        CertificateChain = exampleCertificate.CertificateChain,
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acmpca"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleCertificateAuthority, err := acmpca.NewCertificateAuthority(ctx, "exampleCertificateAuthority", &acmpca.CertificateAuthorityArgs{
			Type: pulumi.String("ROOT"),
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
		current, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleCertificate, err := acmpca.NewCertificate(ctx, "exampleCertificate", &acmpca.CertificateArgs{
			CertificateAuthorityArn:   exampleCertificateAuthority.Arn,
			CertificateSigningRequest: exampleCertificateAuthority.CertificateSigningRequest,
			SigningAlgorithm:          pulumi.String("SHA512WITHRSA"),
			TemplateArn:               pulumi.String(fmt.Sprintf("arn:%v:acm-pca:::template/RootCACertificate/V1", current.Partition)),
			Validity: &acmpca.CertificateValidityArgs{
				Type:  pulumi.String("YEARS"),
				Value: pulumi.String("1"),
			},
		})
		if err != nil {
			return err
		}
		_, err = acmpca.NewCertificateAuthorityCertificate(ctx, "exampleCertificateAuthorityCertificate", &acmpca.CertificateAuthorityCertificateArgs{
			CertificateAuthorityArn: exampleCertificateAuthority.Arn,
			Certificate:             exampleCertificate.Certificate,
			CertificateChain:        exampleCertificate.CertificateChain,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.acmpca.Certificate;
import com.pulumi.aws.acmpca.CertificateArgs;
import com.pulumi.aws.acmpca.inputs.CertificateValidityArgs;
import com.pulumi.aws.acmpca.CertificateAuthorityCertificate;
import com.pulumi.aws.acmpca.CertificateAuthorityCertificateArgs;
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
            .type("ROOT")
            .certificateAuthorityConfiguration(CertificateAuthorityCertificateAuthorityConfigurationArgs.builder()
                .keyAlgorithm("RSA_4096")
                .signingAlgorithm("SHA512WITHRSA")
                .subject(CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs.builder()
                    .commonName("example.com")
                    .build())
                .build())
            .build());

        final var current = AwsFunctions.getPartition();

        var exampleCertificate = new Certificate("exampleCertificate", CertificateArgs.builder()        
            .certificateAuthorityArn(exampleCertificateAuthority.arn())
            .certificateSigningRequest(exampleCertificateAuthority.certificateSigningRequest())
            .signingAlgorithm("SHA512WITHRSA")
            .templateArn(String.format("arn:%s:acm-pca:::template/RootCACertificate/V1", current.applyValue(getPartitionResult -> getPartitionResult.partition())))
            .validity(CertificateValidityArgs.builder()
                .type("YEARS")
                .value(1)
                .build())
            .build());

        var exampleCertificateAuthorityCertificate = new CertificateAuthorityCertificate("exampleCertificateAuthorityCertificate", CertificateAuthorityCertificateArgs.builder()        
            .certificateAuthorityArn(exampleCertificateAuthority.arn())
            .certificate(exampleCertificate.certificate())
            .certificateChain(exampleCertificate.certificateChain())
            .build());

    }
}
```
```yaml
resources:
  exampleCertificateAuthorityCertificate:
    type: aws:acmpca:CertificateAuthorityCertificate
    properties:
      certificateAuthorityArn: ${exampleCertificateAuthority.arn}
      certificate: ${exampleCertificate.certificate}
      certificateChain: ${exampleCertificate.certificateChain}
  exampleCertificate:
    type: aws:acmpca:Certificate
    properties:
      certificateAuthorityArn: ${exampleCertificateAuthority.arn}
      certificateSigningRequest: ${exampleCertificateAuthority.certificateSigningRequest}
      signingAlgorithm: SHA512WITHRSA
      templateArn: arn:${current.partition}:acm-pca:::template/RootCACertificate/V1
      validity:
        type: YEARS
        value: 1
  exampleCertificateAuthority:
    type: aws:acmpca:CertificateAuthority
    properties:
      type: ROOT
      certificateAuthorityConfiguration:
        keyAlgorithm: RSA_4096
        signingAlgorithm: SHA512WITHRSA
        subject:
          commonName: example.com
variables:
  current:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### Certificate for Subordinate Certificate Authority

Note that the certificate for the subordinate certificate authority must be issued by the root certificate authority using a signing request from the subordinate certificate authority.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const subordinateCertificateAuthority = new aws.acmpca.CertificateAuthority("subordinateCertificateAuthority", {
    type: "SUBORDINATE",
    certificateAuthorityConfiguration: {
        keyAlgorithm: "RSA_2048",
        signingAlgorithm: "SHA512WITHRSA",
        subject: {
            commonName: "sub.example.com",
        },
    },
});
const rootCertificateAuthority = new aws.acmpca.CertificateAuthority("rootCertificateAuthority", {});
// ...
const current = aws.getPartition({});
const subordinateCertificate = new aws.acmpca.Certificate("subordinateCertificate", {
    certificateAuthorityArn: rootCertificateAuthority.arn,
    certificateSigningRequest: subordinateCertificateAuthority.certificateSigningRequest,
    signingAlgorithm: "SHA512WITHRSA",
    templateArn: current.then(current => `arn:${current.partition}:acm-pca:::template/SubordinateCACertificate_PathLen0/V1`),
    validity: {
        type: "YEARS",
        value: "1",
    },
});
const subordinateCertificateAuthorityCertificate = new aws.acmpca.CertificateAuthorityCertificate("subordinateCertificateAuthorityCertificate", {
    certificateAuthorityArn: subordinateCertificateAuthority.arn,
    certificate: subordinateCertificate.certificate,
    certificateChain: subordinateCertificate.certificateChain,
});
const rootCertificateAuthorityCertificate = new aws.acmpca.CertificateAuthorityCertificate("rootCertificateAuthorityCertificate", {});
// ...
const rootCertificate = new aws.acmpca.Certificate("rootCertificate", {});
// ...
```
```python
import pulumi
import pulumi_aws as aws

subordinate_certificate_authority = aws.acmpca.CertificateAuthority("subordinateCertificateAuthority",
    type="SUBORDINATE",
    certificate_authority_configuration=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs(
        key_algorithm="RSA_2048",
        signing_algorithm="SHA512WITHRSA",
        subject=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs(
            common_name="sub.example.com",
        ),
    ))
root_certificate_authority = aws.acmpca.CertificateAuthority("rootCertificateAuthority")
# ...
current = aws.get_partition()
subordinate_certificate = aws.acmpca.Certificate("subordinateCertificate",
    certificate_authority_arn=root_certificate_authority.arn,
    certificate_signing_request=subordinate_certificate_authority.certificate_signing_request,
    signing_algorithm="SHA512WITHRSA",
    template_arn=f"arn:{current.partition}:acm-pca:::template/SubordinateCACertificate_PathLen0/V1",
    validity=aws.acmpca.CertificateValidityArgs(
        type="YEARS",
        value="1",
    ))
subordinate_certificate_authority_certificate = aws.acmpca.CertificateAuthorityCertificate("subordinateCertificateAuthorityCertificate",
    certificate_authority_arn=subordinate_certificate_authority.arn,
    certificate=subordinate_certificate.certificate,
    certificate_chain=subordinate_certificate.certificate_chain)
root_certificate_authority_certificate = aws.acmpca.CertificateAuthorityCertificate("rootCertificateAuthorityCertificate")
# ...
root_certificate = aws.acmpca.Certificate("rootCertificate")
# ...
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var subordinateCertificateAuthority = new Aws.Acmpca.CertificateAuthority("subordinateCertificateAuthority", new()
    {
        Type = "SUBORDINATE",
        CertificateAuthorityConfiguration = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationArgs
        {
            KeyAlgorithm = "RSA_2048",
            SigningAlgorithm = "SHA512WITHRSA",
            Subject = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs
            {
                CommonName = "sub.example.com",
            },
        },
    });

    var rootCertificateAuthority = new Aws.Acmpca.CertificateAuthority("rootCertificateAuthority");

    // ...
    var current = Aws.GetPartition.Invoke();

    var subordinateCertificate = new Aws.Acmpca.Certificate("subordinateCertificate", new()
    {
        CertificateAuthorityArn = rootCertificateAuthority.Arn,
        CertificateSigningRequest = subordinateCertificateAuthority.CertificateSigningRequest,
        SigningAlgorithm = "SHA512WITHRSA",
        TemplateArn = $"arn:{current.Apply(getPartitionResult => getPartitionResult.Partition)}:acm-pca:::template/SubordinateCACertificate_PathLen0/V1",
        Validity = new Aws.Acmpca.Inputs.CertificateValidityArgs
        {
            Type = "YEARS",
            Value = "1",
        },
    });

    var subordinateCertificateAuthorityCertificate = new Aws.Acmpca.CertificateAuthorityCertificate("subordinateCertificateAuthorityCertificate", new()
    {
        CertificateAuthorityArn = subordinateCertificateAuthority.Arn,
        Certificate = subordinateCertificate.Certificate,
        CertificateChain = subordinateCertificate.CertificateChain,
    });

    var rootCertificateAuthorityCertificate = new Aws.Acmpca.CertificateAuthorityCertificate("rootCertificateAuthorityCertificate");

    // ...
    var rootCertificate = new Aws.Acmpca.Certificate("rootCertificate");

    // ...
});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acmpca"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		subordinateCertificateAuthority, err := acmpca.NewCertificateAuthority(ctx, "subordinateCertificateAuthority", &acmpca.CertificateAuthorityArgs{
			Type: pulumi.String("SUBORDINATE"),
			CertificateAuthorityConfiguration: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs{
				KeyAlgorithm:     pulumi.String("RSA_2048"),
				SigningAlgorithm: pulumi.String("SHA512WITHRSA"),
				Subject: &acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs{
					CommonName: pulumi.String("sub.example.com"),
				},
			},
		})
		if err != nil {
			return err
		}
		rootCertificateAuthority, err := acmpca.NewCertificateAuthority(ctx, "rootCertificateAuthority", nil)
		if err != nil {
			return err
		}
		current, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		subordinateCertificate, err := acmpca.NewCertificate(ctx, "subordinateCertificate", &acmpca.CertificateArgs{
			CertificateAuthorityArn:   rootCertificateAuthority.Arn,
			CertificateSigningRequest: subordinateCertificateAuthority.CertificateSigningRequest,
			SigningAlgorithm:          pulumi.String("SHA512WITHRSA"),
			TemplateArn:               pulumi.String(fmt.Sprintf("arn:%v:acm-pca:::template/SubordinateCACertificate_PathLen0/V1", current.Partition)),
			Validity: &acmpca.CertificateValidityArgs{
				Type:  pulumi.String("YEARS"),
				Value: pulumi.String("1"),
			},
		})
		if err != nil {
			return err
		}
		_, err = acmpca.NewCertificateAuthorityCertificate(ctx, "subordinateCertificateAuthorityCertificate", &acmpca.CertificateAuthorityCertificateArgs{
			CertificateAuthorityArn: subordinateCertificateAuthority.Arn,
			Certificate:             subordinateCertificate.Certificate,
			CertificateChain:        subordinateCertificate.CertificateChain,
		})
		if err != nil {
			return err
		}
		_, err = acmpca.NewCertificateAuthorityCertificate(ctx, "rootCertificateAuthorityCertificate", nil)
		if err != nil {
			return err
		}
		_, err = acmpca.NewCertificate(ctx, "rootCertificate", nil)
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.acmpca.Certificate;
import com.pulumi.aws.acmpca.CertificateArgs;
import com.pulumi.aws.acmpca.inputs.CertificateValidityArgs;
import com.pulumi.aws.acmpca.CertificateAuthorityCertificate;
import com.pulumi.aws.acmpca.CertificateAuthorityCertificateArgs;
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
        var subordinateCertificateAuthority = new CertificateAuthority("subordinateCertificateAuthority", CertificateAuthorityArgs.builder()        
            .type("SUBORDINATE")
            .certificateAuthorityConfiguration(CertificateAuthorityCertificateAuthorityConfigurationArgs.builder()
                .keyAlgorithm("RSA_2048")
                .signingAlgorithm("SHA512WITHRSA")
                .subject(CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs.builder()
                    .commonName("sub.example.com")
                    .build())
                .build())
            .build());

        var rootCertificateAuthority = new CertificateAuthority("rootCertificateAuthority");

        final var current = AwsFunctions.getPartition();

        var subordinateCertificate = new Certificate("subordinateCertificate", CertificateArgs.builder()        
            .certificateAuthorityArn(rootCertificateAuthority.arn())
            .certificateSigningRequest(subordinateCertificateAuthority.certificateSigningRequest())
            .signingAlgorithm("SHA512WITHRSA")
            .templateArn(String.format("arn:%s:acm-pca:::template/SubordinateCACertificate_PathLen0/V1", current.applyValue(getPartitionResult -> getPartitionResult.partition())))
            .validity(CertificateValidityArgs.builder()
                .type("YEARS")
                .value(1)
                .build())
            .build());

        var subordinateCertificateAuthorityCertificate = new CertificateAuthorityCertificate("subordinateCertificateAuthorityCertificate", CertificateAuthorityCertificateArgs.builder()        
            .certificateAuthorityArn(subordinateCertificateAuthority.arn())
            .certificate(subordinateCertificate.certificate())
            .certificateChain(subordinateCertificate.certificateChain())
            .build());

        var rootCertificateAuthorityCertificate = new CertificateAuthorityCertificate("rootCertificateAuthorityCertificate");

        var rootCertificate = new Certificate("rootCertificate");

    }
}
```
```yaml
resources:
  subordinateCertificateAuthorityCertificate:
    type: aws:acmpca:CertificateAuthorityCertificate
    properties:
      certificateAuthorityArn: ${subordinateCertificateAuthority.arn}
      certificate: ${subordinateCertificate.certificate}
      certificateChain: ${subordinateCertificate.certificateChain}
  subordinateCertificate:
    type: aws:acmpca:Certificate
    properties:
      certificateAuthorityArn: ${rootCertificateAuthority.arn}
      certificateSigningRequest: ${subordinateCertificateAuthority.certificateSigningRequest}
      signingAlgorithm: SHA512WITHRSA
      templateArn: arn:${current.partition}:acm-pca:::template/SubordinateCACertificate_PathLen0/V1
      validity:
        type: YEARS
        value: 1
  subordinateCertificateAuthority:
    type: aws:acmpca:CertificateAuthority
    properties:
      type: SUBORDINATE
      certificateAuthorityConfiguration:
        keyAlgorithm: RSA_2048
        signingAlgorithm: SHA512WITHRSA
        subject:
          commonName: sub.example.com
  rootCertificateAuthority:
    type: aws:acmpca:CertificateAuthority
  rootCertificateAuthorityCertificate:
    type: aws:acmpca:CertificateAuthorityCertificate
  rootCertificate:
    type: aws:acmpca:Certificate
variables:
  current:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}