Manages a KMS multi-Region replica key that uses external key material.
See the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-import.html) for more information on importing key material into multi-Region keys.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.Provider("primary", {region: "us-east-1"});
const primaryExternalKey = new aws.kms.ExternalKey("primaryExternalKey", {
    description: "Multi-Region primary key",
    deletionWindowInDays: 30,
    multiRegion: true,
    enabled: true,
    keyMaterialBase64: "...",
}, {
    provider: aws.primary,
});
const replica = new aws.kms.ReplicaExternalKey("replica", {
    description: "Multi-Region replica key",
    deletionWindowInDays: 7,
    primaryKeyArn: aws_kms_external.primary.arn,
    keyMaterialBase64: "...",
});
// Must be the same key material as the primary's.
```
```python
import pulumi
import pulumi_aws as aws

primary = aws.Provider("primary", region="us-east-1")
primary_external_key = aws.kms.ExternalKey("primaryExternalKey",
    description="Multi-Region primary key",
    deletion_window_in_days=30,
    multi_region=True,
    enabled=True,
    key_material_base64="...",
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
replica = aws.kms.ReplicaExternalKey("replica",
    description="Multi-Region replica key",
    deletion_window_in_days=7,
    primary_key_arn=aws_kms_external["primary"]["arn"],
    key_material_base64="...")
# Must be the same key material as the primary's.
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var primary = new Aws.Provider("primary", new()
    {
        Region = "us-east-1",
    });

    var primaryExternalKey = new Aws.Kms.ExternalKey("primaryExternalKey", new()
    {
        Description = "Multi-Region primary key",
        DeletionWindowInDays = 30,
        MultiRegion = true,
        Enabled = true,
        KeyMaterialBase64 = "...",
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var replica = new Aws.Kms.ReplicaExternalKey("replica", new()
    {
        Description = "Multi-Region replica key",
        DeletionWindowInDays = 7,
        PrimaryKeyArn = aws_kms_external.Primary.Arn,
        KeyMaterialBase64 = "...",
    });

    // Must be the same key material as the primary's.
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "primary", &aws.ProviderArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		_, err = kms.NewExternalKey(ctx, "primaryExternalKey", &kms.ExternalKeyArgs{
			Description:          pulumi.String("Multi-Region primary key"),
			DeletionWindowInDays: pulumi.Int(30),
			MultiRegion:          pulumi.Bool(true),
			Enabled:              pulumi.Bool(true),
			KeyMaterialBase64:    pulumi.String("..."),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		_, err = kms.NewReplicaExternalKey(ctx, "replica", &kms.ReplicaExternalKeyArgs{
			Description:          pulumi.String("Multi-Region replica key"),
			DeletionWindowInDays: pulumi.Int(7),
			PrimaryKeyArn:        pulumi.Any(aws_kms_external.Primary.Arn),
			KeyMaterialBase64:    pulumi.String("..."),
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
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.kms.ExternalKey;
import com.pulumi.aws.kms.ExternalKeyArgs;
import com.pulumi.aws.kms.ReplicaExternalKey;
import com.pulumi.aws.kms.ReplicaExternalKeyArgs;
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
        var primary = new Provider("primary", ProviderArgs.builder()        
            .region("us-east-1")
            .build());

        var primaryExternalKey = new ExternalKey("primaryExternalKey", ExternalKeyArgs.builder()        
            .description("Multi-Region primary key")
            .deletionWindowInDays(30)
            .multiRegion(true)
            .enabled(true)
            .keyMaterialBase64("...")
            .build(), CustomResourceOptions.builder()
                .provider(aws.primary())
                .build());

        var replica = new ReplicaExternalKey("replica", ReplicaExternalKeyArgs.builder()        
            .description("Multi-Region replica key")
            .deletionWindowInDays(7)
            .primaryKeyArn(aws_kms_external.primary().arn())
            .keyMaterialBase64("...")
            .build());

    }
}
```
```yaml
resources:
  primary:
    type: pulumi:providers:aws
    properties:
      region: us-east-1
  primaryExternalKey:
    type: aws:kms:ExternalKey
    properties:
      description: Multi-Region primary key
      deletionWindowInDays: 30
      multiRegion: true
      enabled: true
      keyMaterialBase64: '...'
    options:
      provider: ${aws.primary}
  replica:
    type: aws:kms:ReplicaExternalKey
    properties:
      description: Multi-Region replica key
      deletionWindowInDays: 7
      primaryKeyArn: ${aws_kms_external.primary.arn}
      keyMaterialBase64: '...'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import KMS multi-Region replica keys using the `id`. For example:

```sh
 $ pulumi import aws:kms/replicaExternalKey:ReplicaExternalKey example 1234abcd-12ab-34cd-56ef-1234567890ab
```
 