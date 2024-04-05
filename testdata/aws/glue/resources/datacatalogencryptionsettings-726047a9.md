Provides a Glue Data Catalog Encryption Settings resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.DataCatalogEncryptionSettings("example", {dataCatalogEncryptionSettings: {
    connectionPasswordEncryption: {
        awsKmsKeyId: aws_kms_key.test.arn,
        returnConnectionPasswordEncrypted: true,
    },
    encryptionAtRest: {
        catalogEncryptionMode: "SSE-KMS",
        sseAwsKmsKeyId: aws_kms_key.test.arn,
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.DataCatalogEncryptionSettings("example", data_catalog_encryption_settings=aws.glue.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs(
    connection_password_encryption=aws.glue.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionArgs(
        aws_kms_key_id=aws_kms_key["test"]["arn"],
        return_connection_password_encrypted=True,
    ),
    encryption_at_rest=aws.glue.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestArgs(
        catalog_encryption_mode="SSE-KMS",
        sse_aws_kms_key_id=aws_kms_key["test"]["arn"],
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
    var example = new Aws.Glue.DataCatalogEncryptionSettings("example", new()
    {
        DataCatalogEncryptionSettingsConfig = new Aws.Glue.Inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs
        {
            ConnectionPasswordEncryption = new Aws.Glue.Inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionArgs
            {
                AwsKmsKeyId = aws_kms_key.Test.Arn,
                ReturnConnectionPasswordEncrypted = true,
            },
            EncryptionAtRest = new Aws.Glue.Inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestArgs
            {
                CatalogEncryptionMode = "SSE-KMS",
                SseAwsKmsKeyId = aws_kms_key.Test.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewDataCatalogEncryptionSettings(ctx, "example", &glue.DataCatalogEncryptionSettingsArgs{
			DataCatalogEncryptionSettings: &glue.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs{
				ConnectionPasswordEncryption: &glue.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionArgs{
					AwsKmsKeyId:                       pulumi.Any(aws_kms_key.Test.Arn),
					ReturnConnectionPasswordEncrypted: pulumi.Bool(true),
				},
				EncryptionAtRest: &glue.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestArgs{
					CatalogEncryptionMode: pulumi.String("SSE-KMS"),
					SseAwsKmsKeyId:        pulumi.Any(aws_kms_key.Test.Arn),
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
import com.pulumi.aws.glue.DataCatalogEncryptionSettings;
import com.pulumi.aws.glue.DataCatalogEncryptionSettingsArgs;
import com.pulumi.aws.glue.inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs;
import com.pulumi.aws.glue.inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionArgs;
import com.pulumi.aws.glue.inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestArgs;
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
        var example = new DataCatalogEncryptionSettings("example", DataCatalogEncryptionSettingsArgs.builder()        
            .dataCatalogEncryptionSettings(DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs.builder()
                .connectionPasswordEncryption(DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionArgs.builder()
                    .awsKmsKeyId(aws_kms_key.test().arn())
                    .returnConnectionPasswordEncrypted(true)
                    .build())
                .encryptionAtRest(DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestArgs.builder()
                    .catalogEncryptionMode("SSE-KMS")
                    .sseAwsKmsKeyId(aws_kms_key.test().arn())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:DataCatalogEncryptionSettings
    properties:
      dataCatalogEncryptionSettings:
        connectionPasswordEncryption:
          awsKmsKeyId: ${aws_kms_key.test.arn}
          returnConnectionPasswordEncrypted: true
        encryptionAtRest:
          catalogEncryptionMode: SSE-KMS
          sseAwsKmsKeyId: ${aws_kms_key.test.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Data Catalog Encryption Settings using `CATALOG-ID` (AWS account ID if not custom). For example:

```sh
 $ pulumi import aws:glue/dataCatalogEncryptionSettings:DataCatalogEncryptionSettings example 123456789012
```
 