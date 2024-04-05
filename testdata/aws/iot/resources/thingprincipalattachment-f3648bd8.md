Attaches Principal to AWS IoT Thing.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const example = new aws.iot.Thing("example", {});
const cert = new aws.iot.Certificate("cert", {
    csr: fs.readFileSync("csr.pem", "utf8"),
    active: true,
});
const att = new aws.iot.ThingPrincipalAttachment("att", {
    principal: cert.arn,
    thing: example.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iot.Thing("example")
cert = aws.iot.Certificate("cert",
    csr=(lambda path: open(path).read())("csr.pem"),
    active=True)
att = aws.iot.ThingPrincipalAttachment("att",
    principal=cert.arn,
    thing=example.name)
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Iot.Thing("example");

    var cert = new Aws.Iot.Certificate("cert", new()
    {
        Csr = File.ReadAllText("csr.pem"),
        Active = true,
    });

    var att = new Aws.Iot.ThingPrincipalAttachment("att", new()
    {
        Principal = cert.Arn,
        Thing = example.Name,
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
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
		example, err := iot.NewThing(ctx, "example", nil)
		if err != nil {
			return err
		}
		cert, err := iot.NewCertificate(ctx, "cert", &iot.CertificateArgs{
			Csr:    readFileOrPanic("csr.pem"),
			Active: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = iot.NewThingPrincipalAttachment(ctx, "att", &iot.ThingPrincipalAttachmentArgs{
			Principal: cert.Arn,
			Thing:     example.Name,
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
import com.pulumi.aws.iot.Thing;
import com.pulumi.aws.iot.Certificate;
import com.pulumi.aws.iot.CertificateArgs;
import com.pulumi.aws.iot.ThingPrincipalAttachment;
import com.pulumi.aws.iot.ThingPrincipalAttachmentArgs;
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
        var example = new Thing("example");

        var cert = new Certificate("cert", CertificateArgs.builder()        
            .csr(Files.readString(Paths.get("csr.pem")))
            .active(true)
            .build());

        var att = new ThingPrincipalAttachment("att", ThingPrincipalAttachmentArgs.builder()        
            .principal(cert.arn())
            .thing(example.name())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:iot:Thing
  cert:
    type: aws:iot:Certificate
    properties:
      csr:
        fn::readFile: csr.pem
      active: true
  att:
    type: aws:iot:ThingPrincipalAttachment
    properties:
      principal: ${cert.arn}
      thing: ${example.name}
```
{{% /example %}}
{{% /examples %}}