{{% examples %}}
## Example Usage
{{% example %}}

The following example below creates a CloudFront public key.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const example = new aws.cloudfront.PublicKey("example", {
    comment: "test public key",
    encodedKey: fs.readFileSync("public_key.pem", "utf8"),
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.PublicKey("example",
    comment="test public key",
    encoded_key=(lambda path: open(path).read())("public_key.pem"))
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
		_, err := cloudfront.NewPublicKey(ctx, "example", &cloudfront.PublicKeyArgs{
			Comment:    pulumi.String("test public key"),
			EncodedKey: readFileOrPanic("public_key.pem"),
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
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudFront Public Key using the `id`. For example:

```sh
 $ pulumi import aws:cloudfront/publicKey:PublicKey example K3D5EWEUDCCXON
```
 