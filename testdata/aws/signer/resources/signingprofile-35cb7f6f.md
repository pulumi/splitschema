Creates a Signer Signing Profile. A signing profile contains information about the code signing configuration parameters that can be used by a given code signing user.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testSp = new aws.signer.SigningProfile("testSp", {platformId: "AWSLambda-SHA384-ECDSA"});
const prodSp = new aws.signer.SigningProfile("prodSp", {
    namePrefix: "prod_sp_",
    platformId: "AWSLambda-SHA384-ECDSA",
    signatureValidityPeriod: {
        type: "YEARS",
        value: 5,
    },
    tags: {
        tag1: "value1",
        tag2: "value2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_sp = aws.signer.SigningProfile("testSp", platform_id="AWSLambda-SHA384-ECDSA")
prod_sp = aws.signer.SigningProfile("prodSp",
    name_prefix="prod_sp_",
    platform_id="AWSLambda-SHA384-ECDSA",
    signature_validity_period=aws.signer.SigningProfileSignatureValidityPeriodArgs(
        type="YEARS",
        value=5,
    ),
    tags={
        "tag1": "value1",
        "tag2": "value2",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testSp = new Aws.Signer.SigningProfile("testSp", new()
    {
        PlatformId = "AWSLambda-SHA384-ECDSA",
    });

    var prodSp = new Aws.Signer.SigningProfile("prodSp", new()
    {
        NamePrefix = "prod_sp_",
        PlatformId = "AWSLambda-SHA384-ECDSA",
        SignatureValidityPeriod = new Aws.Signer.Inputs.SigningProfileSignatureValidityPeriodArgs
        {
            Type = "YEARS",
            Value = 5,
        },
        Tags = 
        {
            { "tag1", "value1" },
            { "tag2", "value2" },
        },
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
		_, err := signer.NewSigningProfile(ctx, "testSp", &signer.SigningProfileArgs{
			PlatformId: pulumi.String("AWSLambda-SHA384-ECDSA"),
		})
		if err != nil {
			return err
		}
		_, err = signer.NewSigningProfile(ctx, "prodSp", &signer.SigningProfileArgs{
			NamePrefix: pulumi.String("prod_sp_"),
			PlatformId: pulumi.String("AWSLambda-SHA384-ECDSA"),
			SignatureValidityPeriod: &signer.SigningProfileSignatureValidityPeriodArgs{
				Type:  pulumi.String("YEARS"),
				Value: pulumi.Int(5),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
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
import com.pulumi.aws.signer.SigningProfile;
import com.pulumi.aws.signer.SigningProfileArgs;
import com.pulumi.aws.signer.inputs.SigningProfileSignatureValidityPeriodArgs;
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
        var testSp = new SigningProfile("testSp", SigningProfileArgs.builder()        
            .platformId("AWSLambda-SHA384-ECDSA")
            .build());

        var prodSp = new SigningProfile("prodSp", SigningProfileArgs.builder()        
            .namePrefix("prod_sp_")
            .platformId("AWSLambda-SHA384-ECDSA")
            .signatureValidityPeriod(SigningProfileSignatureValidityPeriodArgs.builder()
                .type("YEARS")
                .value(5)
                .build())
            .tags(Map.ofEntries(
                Map.entry("tag1", "value1"),
                Map.entry("tag2", "value2")
            ))
            .build());

    }
}
```
```yaml
resources:
  testSp:
    type: aws:signer:SigningProfile
    properties:
      platformId: AWSLambda-SHA384-ECDSA
  prodSp:
    type: aws:signer:SigningProfile
    properties:
      namePrefix: prod_sp_
      platformId: AWSLambda-SHA384-ECDSA
      signatureValidityPeriod:
        type: YEARS
        value: 5
      tags:
        tag1: value1
        tag2: value2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Signer signing profiles using the `name`. For example:

```sh
 $ pulumi import aws:signer/signingProfile:SigningProfile test_signer_signing_profile test_sp_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK
```
 