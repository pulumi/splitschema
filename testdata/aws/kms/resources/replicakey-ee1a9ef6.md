Manages a KMS multi-Region replica key.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.Provider("primary", {region: "us-east-1"});
const primaryKey = new aws.kms.Key("primaryKey", {
    description: "Multi-Region primary key",
    deletionWindowInDays: 30,
    multiRegion: true,
}, {
    provider: aws.primary,
});
const replica = new aws.kms.ReplicaKey("replica", {
    description: "Multi-Region replica key",
    deletionWindowInDays: 7,
    primaryKeyArn: primaryKey.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

primary = aws.Provider("primary", region="us-east-1")
primary_key = aws.kms.Key("primaryKey",
    description="Multi-Region primary key",
    deletion_window_in_days=30,
    multi_region=True,
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
replica = aws.kms.ReplicaKey("replica",
    description="Multi-Region replica key",
    deletion_window_in_days=7,
    primary_key_arn=primary_key.arn)
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

    var primaryKey = new Aws.Kms.Key("primaryKey", new()
    {
        Description = "Multi-Region primary key",
        DeletionWindowInDays = 30,
        MultiRegion = true,
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var replica = new Aws.Kms.ReplicaKey("replica", new()
    {
        Description = "Multi-Region replica key",
        DeletionWindowInDays = 7,
        PrimaryKeyArn = primaryKey.Arn,
    });

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
		primaryKey, err := kms.NewKey(ctx, "primaryKey", &kms.KeyArgs{
			Description:          pulumi.String("Multi-Region primary key"),
			DeletionWindowInDays: pulumi.Int(30),
			MultiRegion:          pulumi.Bool(true),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		_, err = kms.NewReplicaKey(ctx, "replica", &kms.ReplicaKeyArgs{
			Description:          pulumi.String("Multi-Region replica key"),
			DeletionWindowInDays: pulumi.Int(7),
			PrimaryKeyArn:        primaryKey.Arn,
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.kms.ReplicaKey;
import com.pulumi.aws.kms.ReplicaKeyArgs;
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

        var primaryKey = new Key("primaryKey", KeyArgs.builder()        
            .description("Multi-Region primary key")
            .deletionWindowInDays(30)
            .multiRegion(true)
            .build(), CustomResourceOptions.builder()
                .provider(aws.primary())
                .build());

        var replica = new ReplicaKey("replica", ReplicaKeyArgs.builder()        
            .description("Multi-Region replica key")
            .deletionWindowInDays(7)
            .primaryKeyArn(primaryKey.arn())
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
  primaryKey:
    type: aws:kms:Key
    properties:
      description: Multi-Region primary key
      deletionWindowInDays: 30
      multiRegion: true
    options:
      provider: ${aws.primary}
  replica:
    type: aws:kms:ReplicaKey
    properties:
      description: Multi-Region replica key
      deletionWindowInDays: 7
      primaryKeyArn: ${primaryKey.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import KMS multi-Region replica keys using the `id`. For example:

```sh
 $ pulumi import aws:kms/replicaKey:ReplicaKey example 1234abcd-12ab-34cd-56ef-1234567890ab
```
 