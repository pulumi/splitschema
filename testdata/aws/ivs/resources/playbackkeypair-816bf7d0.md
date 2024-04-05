Resource for managing an AWS IVS (Interactive Video) Playback Key Pair.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const example = new aws.ivs.PlaybackKeyPair("example", {publicKey: fs.readFileSync("./public-key.pem", "utf8")});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ivs.PlaybackKeyPair("example", public_key=(lambda path: open(path).read())("./public-key.pem"))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ivs.PlaybackKeyPair("example", new()
    {
        PublicKey = File.ReadAllText("./public-key.pem"),
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ivs"
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
		_, err := ivs.NewPlaybackKeyPair(ctx, "example", &ivs.PlaybackKeyPairArgs{
			PublicKey: readFileOrPanic("./public-key.pem"),
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
import com.pulumi.aws.ivs.PlaybackKeyPair;
import com.pulumi.aws.ivs.PlaybackKeyPairArgs;
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
        var example = new PlaybackKeyPair("example", PlaybackKeyPairArgs.builder()        
            .publicKey(Files.readString(Paths.get("./public-key.pem")))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ivs:PlaybackKeyPair
    properties:
      # public-key.pem is a file that contains an ECDSA public key in PEM format.
      publicKey:
        fn::readFile: ./public-key.pem
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IVS (Interactive Video) Playback Key Pair using the ARN. For example:

```sh
 $ pulumi import aws:ivs/playbackKeyPair:PlaybackKeyPair example arn:aws:ivs:us-west-2:326937407773:playback-key/KDJRJNQhiQzA
```
 