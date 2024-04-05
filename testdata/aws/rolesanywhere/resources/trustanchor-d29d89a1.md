Resource for managing a Roles Anywhere Trust Anchor.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCertificateAuthority = new aws.acmpca.CertificateAuthority("exampleCertificateAuthority", {
    permanentDeletionTimeInDays: 7,
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
const testCertificate = new aws.acmpca.Certificate("testCertificate", {
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
    certificate: aws_acmpca_certificate.example.certificate,
    certificateChain: aws_acmpca_certificate.example.certificate_chain,
});
const testTrustAnchor = new aws.rolesanywhere.TrustAnchor("testTrustAnchor", {source: {
    sourceData: {
        acmPcaArn: exampleCertificateAuthority.arn,
    },
    sourceType: "AWS_ACM_PCA",
}}, {
    dependsOn: [exampleCertificateAuthorityCertificate],
});
```
```python
import pulumi
import pulumi_aws as aws

example_certificate_authority = aws.acmpca.CertificateAuthority("exampleCertificateAuthority",
    permanent_deletion_time_in_days=7,
    type="ROOT",
    certificate_authority_configuration=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs(
        key_algorithm="RSA_4096",
        signing_algorithm="SHA512WITHRSA",
        subject=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs(
            common_name="example.com",
        ),
    ))
current = aws.get_partition()
test_certificate = aws.acmpca.Certificate("testCertificate",
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
    certificate=aws_acmpca_certificate["example"]["certificate"],
    certificate_chain=aws_acmpca_certificate["example"]["certificate_chain"])
test_trust_anchor = aws.rolesanywhere.TrustAnchor("testTrustAnchor", source=aws.rolesanywhere.TrustAnchorSourceArgs(
    source_data=aws.rolesanywhere.TrustAnchorSourceSourceDataArgs(
        acm_pca_arn=example_certificate_authority.arn,
    ),
    source_type="AWS_ACM_PCA",
),
opts=pulumi.ResourceOptions(depends_on=[example_certificate_authority_certificate]))
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
        PermanentDeletionTimeInDays = 7,
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

    var testCertificate = new Aws.Acmpca.Certificate("testCertificate", new()
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
        Certificate = aws_acmpca_certificate.Example.Certificate,
        CertificateChain = aws_acmpca_certificate.Example.Certificate_chain,
    });

    var testTrustAnchor = new Aws.RolesAnywhere.TrustAnchor("testTrustAnchor", new()
    {
        Source = new Aws.RolesAnywhere.Inputs.TrustAnchorSourceArgs
        {
            SourceData = new Aws.RolesAnywhere.Inputs.TrustAnchorSourceSourceDataArgs
            {
                AcmPcaArn = exampleCertificateAuthority.Arn,
            },
            SourceType = "AWS_ACM_PCA",
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleCertificateAuthorityCertificate,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acmpca"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rolesanywhere"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleCertificateAuthority, err := acmpca.NewCertificateAuthority(ctx, "exampleCertificateAuthority", &acmpca.CertificateAuthorityArgs{
			PermanentDeletionTimeInDays: pulumi.Int(7),
			Type:                        pulumi.String("ROOT"),
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
		_, err = acmpca.NewCertificate(ctx, "testCertificate", &acmpca.CertificateArgs{
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
		exampleCertificateAuthorityCertificate, err := acmpca.NewCertificateAuthorityCertificate(ctx, "exampleCertificateAuthorityCertificate", &acmpca.CertificateAuthorityCertificateArgs{
			CertificateAuthorityArn: exampleCertificateAuthority.Arn,
			Certificate:             pulumi.Any(aws_acmpca_certificate.Example.Certificate),
			CertificateChain:        pulumi.Any(aws_acmpca_certificate.Example.Certificate_chain),
		})
		if err != nil {
			return err
		}
		_, err = rolesanywhere.NewTrustAnchor(ctx, "testTrustAnchor", &rolesanywhere.TrustAnchorArgs{
			Source: &rolesanywhere.TrustAnchorSourceArgs{
				SourceData: &rolesanywhere.TrustAnchorSourceSourceDataArgs{
					AcmPcaArn: exampleCertificateAuthority.Arn,
				},
				SourceType: pulumi.String("AWS_ACM_PCA"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleCertificateAuthorityCertificate,
		}))
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
import com.pulumi.aws.rolesanywhere.TrustAnchor;
import com.pulumi.aws.rolesanywhere.TrustAnchorArgs;
import com.pulumi.aws.rolesanywhere.inputs.TrustAnchorSourceArgs;
import com.pulumi.aws.rolesanywhere.inputs.TrustAnchorSourceSourceDataArgs;
import com.pulumi.resources.CustomResourceOptions;
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
            .permanentDeletionTimeInDays(7)
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

        var testCertificate = new Certificate("testCertificate", CertificateArgs.builder()        
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
            .certificate(aws_acmpca_certificate.example().certificate())
            .certificateChain(aws_acmpca_certificate.example().certificate_chain())
            .build());

        var testTrustAnchor = new TrustAnchor("testTrustAnchor", TrustAnchorArgs.builder()        
            .source(TrustAnchorSourceArgs.builder()
                .sourceData(TrustAnchorSourceSourceDataArgs.builder()
                    .acmPcaArn(exampleCertificateAuthority.arn())
                    .build())
                .sourceType("AWS_ACM_PCA")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleCertificateAuthorityCertificate)
                .build());

    }
}
```
```yaml
resources:
  exampleCertificateAuthority:
    type: aws:acmpca:CertificateAuthority
    properties:
      permanentDeletionTimeInDays: 7
      type: ROOT
      certificateAuthorityConfiguration:
        keyAlgorithm: RSA_4096
        signingAlgorithm: SHA512WITHRSA
        subject:
          commonName: example.com
  testCertificate:
    type: aws:acmpca:Certificate
    properties:
      certificateAuthorityArn: ${exampleCertificateAuthority.arn}
      certificateSigningRequest: ${exampleCertificateAuthority.certificateSigningRequest}
      signingAlgorithm: SHA512WITHRSA
      templateArn: arn:${current.partition}:acm-pca:::template/RootCACertificate/V1
      validity:
        type: YEARS
        value: 1
  exampleCertificateAuthorityCertificate:
    type: aws:acmpca:CertificateAuthorityCertificate
    properties:
      certificateAuthorityArn: ${exampleCertificateAuthority.arn}
      certificate: ${aws_acmpca_certificate.example.certificate}
      certificateChain: ${aws_acmpca_certificate.example.certificate_chain}
  testTrustAnchor:
    type: aws:rolesanywhere:TrustAnchor
    properties:
      source:
        sourceData:
          acmPcaArn: ${exampleCertificateAuthority.arn}
        sourceType: AWS_ACM_PCA
    options:
      dependson:
        - ${exampleCertificateAuthorityCertificate}
variables:
  current:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_rolesanywhere_trust_anchor` using its `id`. For example:

```sh
 $ pulumi import aws:rolesanywhere/trustAnchor:TrustAnchor example 92b2fbbb-984d-41a3-a765-e3cbdb69ebb1
```
 