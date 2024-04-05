{{% examples %}}
## Example Usage
{{% example %}}

The following example below creates a CloudFront key group.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const examplePublicKey = new aws.cloudfront.PublicKey("examplePublicKey", {
    comment: "example public key",
    encodedKey: fs.readFileSync("public_key.pem", "utf8"),
});
const exampleKeyGroup = new aws.cloudfront.KeyGroup("exampleKeyGroup", {
    comment: "example key group",
    items: [examplePublicKey.id],
});
```
```python
import pulumi
import pulumi_aws as aws

example_public_key = aws.cloudfront.PublicKey("examplePublicKey",
    comment="example public key",
    encoded_key=(lambda path: open(path).read())("public_key.pem"))
example_key_group = aws.cloudfront.KeyGroup("exampleKeyGroup",
    comment="example key group",
    items=[example_public_key.id])
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplePublicKey = new Aws.CloudFront.PublicKey("examplePublicKey", new()
    {
        Comment = "example public key",
        EncodedKey = File.ReadAllText("public_key.pem"),
    });

    var exampleKeyGroup = new Aws.CloudFront.KeyGroup("exampleKeyGroup", new()
    {
        Comment = "example key group",
        Items = new[]
        {
            examplePublicKey.Id,
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
		examplePublicKey, err := cloudfront.NewPublicKey(ctx, "examplePublicKey", &cloudfront.PublicKeyArgs{
			Comment:    pulumi.String("example public key"),
			EncodedKey: readFileOrPanic("public_key.pem"),
		})
		if err != nil {
			return err
		}
		_, err = cloudfront.NewKeyGroup(ctx, "exampleKeyGroup", &cloudfront.KeyGroupArgs{
			Comment: pulumi.String("example key group"),
			Items: pulumi.StringArray{
				examplePublicKey.ID(),
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
import com.pulumi.aws.cloudfront.KeyGroup;
import com.pulumi.aws.cloudfront.KeyGroupArgs;
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
        var examplePublicKey = new PublicKey("examplePublicKey", PublicKeyArgs.builder()        
            .comment("example public key")
            .encodedKey(Files.readString(Paths.get("public_key.pem")))
            .build());

        var exampleKeyGroup = new KeyGroup("exampleKeyGroup", KeyGroupArgs.builder()        
            .comment("example key group")
            .items(examplePublicKey.id())
            .build());

    }
}
```
```yaml
resources:
  examplePublicKey:
    type: aws:cloudfront:PublicKey
    properties:
      comment: example public key
      encodedKey:
        fn::readFile: public_key.pem
  exampleKeyGroup:
    type: aws:cloudfront:KeyGroup
    properties:
      comment: example key group
      items:
        - ${examplePublicKey.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudFront Key Group using the `id`. For example:

```sh
 $ pulumi import aws:cloudfront/keyGroup:KeyGroup example 4b4f2r1c-315d-5c2e-f093-216t50jed10f
```
 