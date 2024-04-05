Provides a CloudFront Field-level Encryption Profile resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const example = new aws.cloudfront.PublicKey("example", {
    comment: "test public key",
    encodedKey: fs.readFileSync("public_key.pem", "utf8"),
});
const test = new aws.cloudfront.FieldLevelEncryptionProfile("test", {
    comment: "test comment",
    encryptionEntities: {
        items: [{
            publicKeyId: example.id,
            providerId: "test provider",
            fieldPatterns: {
                items: ["DateOfBirth"],
            },
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.PublicKey("example",
    comment="test public key",
    encoded_key=(lambda path: open(path).read())("public_key.pem"))
test = aws.cloudfront.FieldLevelEncryptionProfile("test",
    comment="test comment",
    encryption_entities=aws.cloudfront.FieldLevelEncryptionProfileEncryptionEntitiesArgs(
        items=[aws.cloudfront.FieldLevelEncryptionProfileEncryptionEntitiesItemArgs(
            public_key_id=example.id,
            provider_id="test provider",
            field_patterns=aws.cloudfront.FieldLevelEncryptionProfileEncryptionEntitiesItemFieldPatternsArgs(
                items=["DateOfBirth"],
            ),
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudFront.PublicKey("example", new()
    {
        Comment = "test public key",
        EncodedKey = File.ReadAllText("public_key.pem"),
    });

    var test = new Aws.CloudFront.FieldLevelEncryptionProfile("test", new()
    {
        Comment = "test comment",
        EncryptionEntities = new Aws.CloudFront.Inputs.FieldLevelEncryptionProfileEncryptionEntitiesArgs
        {
            Items = new[]
            {
                new Aws.CloudFront.Inputs.FieldLevelEncryptionProfileEncryptionEntitiesItemArgs
                {
                    PublicKeyId = example.Id,
                    ProviderId = "test provider",
                    FieldPatterns = new Aws.CloudFront.Inputs.FieldLevelEncryptionProfileEncryptionEntitiesItemFieldPatternsArgs
                    {
                        Items = new[]
                        {
                            "DateOfBirth",
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
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
		example, err := cloudfront.NewPublicKey(ctx, "example", &cloudfront.PublicKeyArgs{
			Comment:    pulumi.String("test public key"),
			EncodedKey: readFileOrPanic("public_key.pem"),
		})
		if err != nil {
			return err
		}
		_, err = cloudfront.NewFieldLevelEncryptionProfile(ctx, "test", &cloudfront.FieldLevelEncryptionProfileArgs{
			Comment: pulumi.String("test comment"),
			EncryptionEntities: &cloudfront.FieldLevelEncryptionProfileEncryptionEntitiesArgs{
				Items: cloudfront.FieldLevelEncryptionProfileEncryptionEntitiesItemArray{
					&cloudfront.FieldLevelEncryptionProfileEncryptionEntitiesItemArgs{
						PublicKeyId: example.ID(),
						ProviderId:  pulumi.String("test provider"),
						FieldPatterns: &cloudfront.FieldLevelEncryptionProfileEncryptionEntitiesItemFieldPatternsArgs{
							Items: pulumi.StringArray{
								pulumi.String("DateOfBirth"),
							},
						},
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
import com.pulumi.aws.cloudfront.PublicKey;
import com.pulumi.aws.cloudfront.PublicKeyArgs;
import com.pulumi.aws.cloudfront.FieldLevelEncryptionProfile;
import com.pulumi.aws.cloudfront.FieldLevelEncryptionProfileArgs;
import com.pulumi.aws.cloudfront.inputs.FieldLevelEncryptionProfileEncryptionEntitiesArgs;
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
        var example = new PublicKey("example", PublicKeyArgs.builder()        
            .comment("test public key")
            .encodedKey(Files.readString(Paths.get("public_key.pem")))
            .build());

        var test = new FieldLevelEncryptionProfile("test", FieldLevelEncryptionProfileArgs.builder()        
            .comment("test comment")
            .encryptionEntities(FieldLevelEncryptionProfileEncryptionEntitiesArgs.builder()
                .items(FieldLevelEncryptionProfileEncryptionEntitiesItemArgs.builder()
                    .publicKeyId(example.id())
                    .providerId("test provider")
                    .fieldPatterns(FieldLevelEncryptionProfileEncryptionEntitiesItemFieldPatternsArgs.builder()
                        .items("DateOfBirth")
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
    type: aws:cloudfront:PublicKey
    properties:
      comment: test public key
      encodedKey:
        fn::readFile: public_key.pem
  test:
    type: aws:cloudfront:FieldLevelEncryptionProfile
    properties:
      comment: test comment
      encryptionEntities:
        items:
          - publicKeyId: ${example.id}
            providerId: test provider
            fieldPatterns:
              items:
                - DateOfBirth
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cloudfront Field Level Encryption Profile using the `id`. For example:

```sh
 $ pulumi import aws:cloudfront/fieldLevelEncryptionProfile:FieldLevelEncryptionProfile profile K3D5EWEUDCCXON
```
 