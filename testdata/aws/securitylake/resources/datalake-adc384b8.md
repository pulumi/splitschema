Resource for managing an AWS Security Lake Data Lake.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.securitylake.DataLake("example", {
    metaStoreManagerRoleArn: aws_iam_role.meta_store_manager.arn,
    configuration: {
        region: "eu-west-1",
        encryptionConfigurations: [{
            kmsKeyId: "S3_MANAGED_KEY",
        }],
        lifecycleConfiguration: {
            transitions: [
                {
                    days: 31,
                    storageClass: "STANDARD_IA",
                },
                {
                    days: 80,
                    storageClass: "ONEZONE_IA",
                },
            ],
            expiration: {
                days: 300,
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.securitylake.DataLake("example",
    meta_store_manager_role_arn=aws_iam_role["meta_store_manager"]["arn"],
    configuration=aws.securitylake.DataLakeConfigurationArgs(
        region="eu-west-1",
        encryption_configurations=[{
            "kmsKeyId": "S3_MANAGED_KEY",
        }],
        lifecycle_configuration=aws.securitylake.DataLakeConfigurationLifecycleConfigurationArgs(
            transitions=[
                aws.securitylake.DataLakeConfigurationLifecycleConfigurationTransitionArgs(
                    days=31,
                    storage_class="STANDARD_IA",
                ),
                aws.securitylake.DataLakeConfigurationLifecycleConfigurationTransitionArgs(
                    days=80,
                    storage_class="ONEZONE_IA",
                ),
            ],
            expiration=aws.securitylake.DataLakeConfigurationLifecycleConfigurationExpirationArgs(
                days=300,
            ),
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SecurityLake.DataLake("example", new()
    {
        MetaStoreManagerRoleArn = aws_iam_role.Meta_store_manager.Arn,
        Configuration = new Aws.SecurityLake.Inputs.DataLakeConfigurationArgs
        {
            Region = "eu-west-1",
            EncryptionConfigurations = new[]
            {
                
                {
                    { "kmsKeyId", "S3_MANAGED_KEY" },
                },
            },
            LifecycleConfiguration = new Aws.SecurityLake.Inputs.DataLakeConfigurationLifecycleConfigurationArgs
            {
                Transitions = new[]
                {
                    new Aws.SecurityLake.Inputs.DataLakeConfigurationLifecycleConfigurationTransitionArgs
                    {
                        Days = 31,
                        StorageClass = "STANDARD_IA",
                    },
                    new Aws.SecurityLake.Inputs.DataLakeConfigurationLifecycleConfigurationTransitionArgs
                    {
                        Days = 80,
                        StorageClass = "ONEZONE_IA",
                    },
                },
                Expiration = new Aws.SecurityLake.Inputs.DataLakeConfigurationLifecycleConfigurationExpirationArgs
                {
                    Days = 300,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securitylake"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securitylake.NewDataLake(ctx, "example", &securitylake.DataLakeArgs{
			MetaStoreManagerRoleArn: pulumi.Any(aws_iam_role.Meta_store_manager.Arn),
			Configuration: &securitylake.DataLakeConfigurationArgs{
				Region: pulumi.String("eu-west-1"),
				EncryptionConfigurations: pulumi.MapArray{
					pulumi.Map{
						"kmsKeyId": pulumi.Any("S3_MANAGED_KEY"),
					},
				},
				LifecycleConfiguration: &securitylake.DataLakeConfigurationLifecycleConfigurationArgs{
					Transitions: securitylake.DataLakeConfigurationLifecycleConfigurationTransitionArray{
						&securitylake.DataLakeConfigurationLifecycleConfigurationTransitionArgs{
							Days:         pulumi.Int(31),
							StorageClass: pulumi.String("STANDARD_IA"),
						},
						&securitylake.DataLakeConfigurationLifecycleConfigurationTransitionArgs{
							Days:         pulumi.Int(80),
							StorageClass: pulumi.String("ONEZONE_IA"),
						},
					},
					Expiration: &securitylake.DataLakeConfigurationLifecycleConfigurationExpirationArgs{
						Days: pulumi.Int(300),
					},
				},
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
import com.pulumi.aws.securitylake.DataLake;
import com.pulumi.aws.securitylake.DataLakeArgs;
import com.pulumi.aws.securitylake.inputs.DataLakeConfigurationArgs;
import com.pulumi.aws.securitylake.inputs.DataLakeConfigurationLifecycleConfigurationArgs;
import com.pulumi.aws.securitylake.inputs.DataLakeConfigurationLifecycleConfigurationExpirationArgs;
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
        var example = new DataLake("example", DataLakeArgs.builder()        
            .metaStoreManagerRoleArn(aws_iam_role.meta_store_manager().arn())
            .configuration(DataLakeConfigurationArgs.builder()
                .region("eu-west-1")
                .encryptionConfigurations(Map.of("kmsKeyId", "S3_MANAGED_KEY"))
                .lifecycleConfiguration(DataLakeConfigurationLifecycleConfigurationArgs.builder()
                    .transitions(                    
                        DataLakeConfigurationLifecycleConfigurationTransitionArgs.builder()
                            .days(31)
                            .storageClass("STANDARD_IA")
                            .build(),
                        DataLakeConfigurationLifecycleConfigurationTransitionArgs.builder()
                            .days(80)
                            .storageClass("ONEZONE_IA")
                            .build())
                    .expiration(DataLakeConfigurationLifecycleConfigurationExpirationArgs.builder()
                        .days(300)
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:securitylake:DataLake
    properties:
      metaStoreManagerRoleArn: ${aws_iam_role.meta_store_manager.arn}
      configuration:
        region: eu-west-1
        encryptionConfigurations:
          - kmsKeyId: S3_MANAGED_KEY
        lifecycleConfiguration:
          transitions:
            - days: 31
              storageClass: STANDARD_IA
            - days: 80
              storageClass: ONEZONE_IA
          expiration:
            days: 300
```
{{% /example %}}
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.securitylake.DataLake("example", {
    metaStoreManagerRoleArn: aws_iam_role.meta_store_manager.arn,
    configuration: {
        region: "eu-west-1",
        encryptionConfigurations: [{
            kmsKeyId: "S3_MANAGED_KEY",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.securitylake.DataLake("example",
    meta_store_manager_role_arn=aws_iam_role["meta_store_manager"]["arn"],
    configuration=aws.securitylake.DataLakeConfigurationArgs(
        region="eu-west-1",
        encryption_configurations=[{
            "kmsKeyId": "S3_MANAGED_KEY",
        }],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SecurityLake.DataLake("example", new()
    {
        MetaStoreManagerRoleArn = aws_iam_role.Meta_store_manager.Arn,
        Configuration = new Aws.SecurityLake.Inputs.DataLakeConfigurationArgs
        {
            Region = "eu-west-1",
            EncryptionConfigurations = new[]
            {
                
                {
                    { "kmsKeyId", "S3_MANAGED_KEY" },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securitylake"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securitylake.NewDataLake(ctx, "example", &securitylake.DataLakeArgs{
			MetaStoreManagerRoleArn: pulumi.Any(aws_iam_role.Meta_store_manager.Arn),
			Configuration: &securitylake.DataLakeConfigurationArgs{
				Region: pulumi.String("eu-west-1"),
				EncryptionConfigurations: pulumi.MapArray{
					pulumi.Map{
						"kmsKeyId": pulumi.Any("S3_MANAGED_KEY"),
					},
				},
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
import com.pulumi.aws.securitylake.DataLake;
import com.pulumi.aws.securitylake.DataLakeArgs;
import com.pulumi.aws.securitylake.inputs.DataLakeConfigurationArgs;
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
        var example = new DataLake("example", DataLakeArgs.builder()        
            .metaStoreManagerRoleArn(aws_iam_role.meta_store_manager().arn())
            .configuration(DataLakeConfigurationArgs.builder()
                .region("eu-west-1")
                .encryptionConfigurations(Map.of("kmsKeyId", "S3_MANAGED_KEY"))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:securitylake:DataLake
    properties:
      metaStoreManagerRoleArn: ${aws_iam_role.meta_store_manager.arn}
      configuration:
        region: eu-west-1
        encryptionConfigurations:
          - kmsKeyId: S3_MANAGED_KEY
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Security Hub standards subscriptions using the standards subscription ARN. For example:

```sh
 $ pulumi import aws:securitylake/dataLake:DataLake example arn:aws:securitylake:eu-west-1:123456789012:data-lake/default
```
 