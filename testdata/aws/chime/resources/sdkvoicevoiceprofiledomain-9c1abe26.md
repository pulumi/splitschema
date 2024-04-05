Resource for managing an AWS Chime SDK Voice Profile Domain.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {
    description: "KMS Key for Voice Profile Domain",
    deletionWindowInDays: 7,
});
const exampleSdkvoiceVoiceProfileDomain = new aws.chime.SdkvoiceVoiceProfileDomain("exampleSdkvoiceVoiceProfileDomain", {
    serverSideEncryptionConfiguration: {
        kmsKeyArn: exampleKey.arn,
    },
    description: "My Voice Profile Domain",
    tags: {
        key1: "value1",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey",
    description="KMS Key for Voice Profile Domain",
    deletion_window_in_days=7)
example_sdkvoice_voice_profile_domain = aws.chime.SdkvoiceVoiceProfileDomain("exampleSdkvoiceVoiceProfileDomain",
    server_side_encryption_configuration=aws.chime.SdkvoiceVoiceProfileDomainServerSideEncryptionConfigurationArgs(
        kms_key_arn=example_key.arn,
    ),
    description="My Voice Profile Domain",
    tags={
        "key1": "value1",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "KMS Key for Voice Profile Domain",
        DeletionWindowInDays = 7,
    });

    var exampleSdkvoiceVoiceProfileDomain = new Aws.Chime.SdkvoiceVoiceProfileDomain("exampleSdkvoiceVoiceProfileDomain", new()
    {
        ServerSideEncryptionConfiguration = new Aws.Chime.Inputs.SdkvoiceVoiceProfileDomainServerSideEncryptionConfigurationArgs
        {
            KmsKeyArn = exampleKey.Arn,
        },
        Description = "My Voice Profile Domain",
        Tags = 
        {
            { "key1", "value1" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chime"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description:          pulumi.String("KMS Key for Voice Profile Domain"),
			DeletionWindowInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		_, err = chime.NewSdkvoiceVoiceProfileDomain(ctx, "exampleSdkvoiceVoiceProfileDomain", &chime.SdkvoiceVoiceProfileDomainArgs{
			ServerSideEncryptionConfiguration: &chime.SdkvoiceVoiceProfileDomainServerSideEncryptionConfigurationArgs{
				KmsKeyArn: exampleKey.Arn,
			},
			Description: pulumi.String("My Voice Profile Domain"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.chime.SdkvoiceVoiceProfileDomain;
import com.pulumi.aws.chime.SdkvoiceVoiceProfileDomainArgs;
import com.pulumi.aws.chime.inputs.SdkvoiceVoiceProfileDomainServerSideEncryptionConfigurationArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("KMS Key for Voice Profile Domain")
            .deletionWindowInDays(7)
            .build());

        var exampleSdkvoiceVoiceProfileDomain = new SdkvoiceVoiceProfileDomain("exampleSdkvoiceVoiceProfileDomain", SdkvoiceVoiceProfileDomainArgs.builder()        
            .serverSideEncryptionConfiguration(SdkvoiceVoiceProfileDomainServerSideEncryptionConfigurationArgs.builder()
                .kmsKeyArn(exampleKey.arn())
                .build())
            .description("My Voice Profile Domain")
            .tags(Map.of("key1", "value1"))
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: KMS Key for Voice Profile Domain
      deletionWindowInDays: 7
  exampleSdkvoiceVoiceProfileDomain:
    type: aws:chime:SdkvoiceVoiceProfileDomain
    properties:
      serverSideEncryptionConfiguration:
        kmsKeyArn: ${exampleKey.arn}
      description: My Voice Profile Domain
      tags:
        key1: value1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS Chime SDK Voice Profile Domain using the `id`. For example:

```sh
 $ pulumi import aws:chime/sdkvoiceVoiceProfileDomain:SdkvoiceVoiceProfileDomain example abcdef123456
```
 