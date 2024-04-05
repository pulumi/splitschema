Provides an IAM SAML provider.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const _default = new aws.iam.SamlProvider("default", {samlMetadataDocument: fs.readFileSync("saml-metadata.xml", "utf8")});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.iam.SamlProvider("default", saml_metadata_document=(lambda path: open(path).read())("saml-metadata.xml"))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Iam.SamlProvider("default", new()
    {
        SamlMetadataDocument = File.ReadAllText("saml-metadata.xml"),
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
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
		_, err := iam.NewSamlProvider(ctx, "default", &iam.SamlProviderArgs{
			SamlMetadataDocument: readFileOrPanic("saml-metadata.xml"),
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
import com.pulumi.aws.iam.SamlProvider;
import com.pulumi.aws.iam.SamlProviderArgs;
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
        var default_ = new SamlProvider("default", SamlProviderArgs.builder()        
            .samlMetadataDocument(Files.readString(Paths.get("saml-metadata.xml")))
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:iam:SamlProvider
    properties:
      samlMetadataDocument:
        fn::readFile: saml-metadata.xml
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM SAML Providers using the `arn`. For example:

```sh
 $ pulumi import aws:iam/samlProvider:SamlProvider default arn:aws:iam::123456789012:saml-provider/SAMLADFS
```
 