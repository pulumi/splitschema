Provides a AWS Transfer AS2 Certificate resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const example = new aws.transfer.Certificate("example", {
    certificate: fs.readFileSync(`${path.module}/example.com/example.crt`, "utf8"),
    certificateChain: fs.readFileSync(`${path.module}/example.com/ca.crt`, "utf8"),
    privateKey: fs.readFileSync(`${path.module}/example.com/example.key`, "utf8"),
    description: "example",
    usage: "SIGNING",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Certificate("example",
    certificate=(lambda path: open(path).read())(f"{path['module']}/example.com/example.crt"),
    certificate_chain=(lambda path: open(path).read())(f"{path['module']}/example.com/ca.crt"),
    private_key=(lambda path: open(path).read())(f"{path['module']}/example.com/example.key"),
    description="example",
    usage="SIGNING")
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Certificate("example", new()
    {
        CertificateFile = File.ReadAllText($"{path.Module}/example.com/example.crt"),
        CertificateChain = File.ReadAllText($"{path.Module}/example.com/ca.crt"),
        PrivateKey = File.ReadAllText($"{path.Module}/example.com/example.key"),
        Description = "example",
        Usage = "SIGNING",
    });

});
```
```go
package main

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
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
		_, err := transfer.NewCertificate(ctx, "example", &transfer.CertificateArgs{
			Certificate:      readFileOrPanic(fmt.Sprintf("%v/example.com/example.crt", path.Module)),
			CertificateChain: readFileOrPanic(fmt.Sprintf("%v/example.com/ca.crt", path.Module)),
			PrivateKey:       readFileOrPanic(fmt.Sprintf("%v/example.com/example.key", path.Module)),
			Description:      pulumi.String("example"),
			Usage:            pulumi.String("SIGNING"),
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
import com.pulumi.aws.transfer.Certificate;
import com.pulumi.aws.transfer.CertificateArgs;
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
        var example = new Certificate("example", CertificateArgs.builder()        
            .certificate(Files.readString(Paths.get(String.format("%s/example.com/example.crt", path.module()))))
            .certificateChain(Files.readString(Paths.get(String.format("%s/example.com/ca.crt", path.module()))))
            .privateKey(Files.readString(Paths.get(String.format("%s/example.com/example.key", path.module()))))
            .description("example")
            .usage("SIGNING")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Certificate
    properties:
      certificate:
        fn::readFile: ${path.module}/example.com/example.crt
      certificateChain:
        fn::readFile: ${path.module}/example.com/ca.crt
      privateKey:
        fn::readFile: ${path.module}/example.com/example.key
      description: example
      usage: SIGNING
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transfer AS2 Certificate using the `certificate_id`. For example:

```sh
 $ pulumi import aws:transfer/certificate:Certificate example c-4221a88afd5f4362a
```
 