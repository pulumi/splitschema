Creates a Signer Signing Profile Permission. That is, a cross-account permission for a signing profile.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const prodSp = new aws.signer.SigningProfile("prodSp", {
    platformId: "AWSLambda-SHA384-ECDSA",
    namePrefix: "prod_sp_",
    signatureValidityPeriod: {
        value: 5,
        type: "YEARS",
    },
    tags: {
        tag1: "value1",
        tag2: "value2",
    },
});
const spPermission1 = new aws.signer.SigningProfilePermission("spPermission1", {
    profileName: prodSp.name,
    action: "signer:StartSigningJob",
    principal: _var.aws_account,
});
const spPermission2 = new aws.signer.SigningProfilePermission("spPermission2", {
    profileName: prodSp.name,
    action: "signer:GetSigningProfile",
    principal: _var.aws_team_role_arn,
    statementId: "ProdAccountStartSigningJob_StatementId",
});
const spPermission3 = new aws.signer.SigningProfilePermission("spPermission3", {
    profileName: prodSp.name,
    action: "signer:RevokeSignature",
    principal: "123456789012",
    profileVersion: prodSp.version,
    statementIdPrefix: "version-permission-",
});
```
```python
import pulumi
import pulumi_aws as aws

prod_sp = aws.signer.SigningProfile("prodSp",
    platform_id="AWSLambda-SHA384-ECDSA",
    name_prefix="prod_sp_",
    signature_validity_period=aws.signer.SigningProfileSignatureValidityPeriodArgs(
        value=5,
        type="YEARS",
    ),
    tags={
        "tag1": "value1",
        "tag2": "value2",
    })
sp_permission1 = aws.signer.SigningProfilePermission("spPermission1",
    profile_name=prod_sp.name,
    action="signer:StartSigningJob",
    principal=var["aws_account"])
sp_permission2 = aws.signer.SigningProfilePermission("spPermission2",
    profile_name=prod_sp.name,
    action="signer:GetSigningProfile",
    principal=var["aws_team_role_arn"],
    statement_id="ProdAccountStartSigningJob_StatementId")
sp_permission3 = aws.signer.SigningProfilePermission("spPermission3",
    profile_name=prod_sp.name,
    action="signer:RevokeSignature",
    principal="123456789012",
    profile_version=prod_sp.version,
    statement_id_prefix="version-permission-")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var prodSp = new Aws.Signer.SigningProfile("prodSp", new()
    {
        PlatformId = "AWSLambda-SHA384-ECDSA",
        NamePrefix = "prod_sp_",
        SignatureValidityPeriod = new Aws.Signer.Inputs.SigningProfileSignatureValidityPeriodArgs
        {
            Value = 5,
            Type = "YEARS",
        },
        Tags = 
        {
            { "tag1", "value1" },
            { "tag2", "value2" },
        },
    });

    var spPermission1 = new Aws.Signer.SigningProfilePermission("spPermission1", new()
    {
        ProfileName = prodSp.Name,
        Action = "signer:StartSigningJob",
        Principal = @var.Aws_account,
    });

    var spPermission2 = new Aws.Signer.SigningProfilePermission("spPermission2", new()
    {
        ProfileName = prodSp.Name,
        Action = "signer:GetSigningProfile",
        Principal = @var.Aws_team_role_arn,
        StatementId = "ProdAccountStartSigningJob_StatementId",
    });

    var spPermission3 = new Aws.Signer.SigningProfilePermission("spPermission3", new()
    {
        ProfileName = prodSp.Name,
        Action = "signer:RevokeSignature",
        Principal = "123456789012",
        ProfileVersion = prodSp.Version,
        StatementIdPrefix = "version-permission-",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/signer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		prodSp, err := signer.NewSigningProfile(ctx, "prodSp", &signer.SigningProfileArgs{
			PlatformId: pulumi.String("AWSLambda-SHA384-ECDSA"),
			NamePrefix: pulumi.String("prod_sp_"),
			SignatureValidityPeriod: &signer.SigningProfileSignatureValidityPeriodArgs{
				Value: pulumi.Int(5),
				Type:  pulumi.String("YEARS"),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		_, err = signer.NewSigningProfilePermission(ctx, "spPermission1", &signer.SigningProfilePermissionArgs{
			ProfileName: prodSp.Name,
			Action:      pulumi.String("signer:StartSigningJob"),
			Principal:   pulumi.Any(_var.Aws_account),
		})
		if err != nil {
			return err
		}
		_, err = signer.NewSigningProfilePermission(ctx, "spPermission2", &signer.SigningProfilePermissionArgs{
			ProfileName: prodSp.Name,
			Action:      pulumi.String("signer:GetSigningProfile"),
			Principal:   pulumi.Any(_var.Aws_team_role_arn),
			StatementId: pulumi.String("ProdAccountStartSigningJob_StatementId"),
		})
		if err != nil {
			return err
		}
		_, err = signer.NewSigningProfilePermission(ctx, "spPermission3", &signer.SigningProfilePermissionArgs{
			ProfileName:       prodSp.Name,
			Action:            pulumi.String("signer:RevokeSignature"),
			Principal:         pulumi.String("123456789012"),
			ProfileVersion:    prodSp.Version,
			StatementIdPrefix: pulumi.String("version-permission-"),
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
import com.pulumi.aws.signer.SigningProfile;
import com.pulumi.aws.signer.SigningProfileArgs;
import com.pulumi.aws.signer.inputs.SigningProfileSignatureValidityPeriodArgs;
import com.pulumi.aws.signer.SigningProfilePermission;
import com.pulumi.aws.signer.SigningProfilePermissionArgs;
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
        var prodSp = new SigningProfile("prodSp", SigningProfileArgs.builder()        
            .platformId("AWSLambda-SHA384-ECDSA")
            .namePrefix("prod_sp_")
            .signatureValidityPeriod(SigningProfileSignatureValidityPeriodArgs.builder()
                .value(5)
                .type("YEARS")
                .build())
            .tags(Map.ofEntries(
                Map.entry("tag1", "value1"),
                Map.entry("tag2", "value2")
            ))
            .build());

        var spPermission1 = new SigningProfilePermission("spPermission1", SigningProfilePermissionArgs.builder()        
            .profileName(prodSp.name())
            .action("signer:StartSigningJob")
            .principal(var_.aws_account())
            .build());

        var spPermission2 = new SigningProfilePermission("spPermission2", SigningProfilePermissionArgs.builder()        
            .profileName(prodSp.name())
            .action("signer:GetSigningProfile")
            .principal(var_.aws_team_role_arn())
            .statementId("ProdAccountStartSigningJob_StatementId")
            .build());

        var spPermission3 = new SigningProfilePermission("spPermission3", SigningProfilePermissionArgs.builder()        
            .profileName(prodSp.name())
            .action("signer:RevokeSignature")
            .principal("123456789012")
            .profileVersion(prodSp.version())
            .statementIdPrefix("version-permission-")
            .build());

    }
}
```
```yaml
resources:
  prodSp:
    type: aws:signer:SigningProfile
    properties:
      platformId: AWSLambda-SHA384-ECDSA
      namePrefix: prod_sp_
      signatureValidityPeriod:
        value: 5
        type: YEARS
      tags:
        tag1: value1
        tag2: value2
  spPermission1:
    type: aws:signer:SigningProfilePermission
    properties:
      profileName: ${prodSp.name}
      action: signer:StartSigningJob
      principal: ${var.aws_account}
  spPermission2:
    type: aws:signer:SigningProfilePermission
    properties:
      profileName: ${prodSp.name}
      action: signer:GetSigningProfile
      principal: ${var.aws_team_role_arn}
      statementId: ProdAccountStartSigningJob_StatementId
  spPermission3:
    type: aws:signer:SigningProfilePermission
    properties:
      profileName: ${prodSp.name}
      action: signer:RevokeSignature
      principal: '123456789012'
      profileVersion: ${prodSp.version}
      statementIdPrefix: version-permission-
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Signer signing profile permission statements using profile_name/statement_id. For example:

```sh
 $ pulumi import aws:signer/signingProfilePermission:SigningProfilePermission test_signer_signing_profile_permission prod_profile_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK/ProdAccountStartSigningJobStatementId
```
 