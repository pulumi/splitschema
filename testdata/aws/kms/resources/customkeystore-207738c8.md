Resource for managing an AWS KMS (Key Management) Custom Key Store.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const test = new aws.kms.CustomKeyStore("test", {
    cloudHsmClusterId: _var.cloud_hsm_cluster_id,
    customKeyStoreName: "kms-custom-key-store-test",
    keyStorePassword: "noplaintextpasswords1",
    trustAnchorCertificate: fs.readFileSync("anchor-certificate.crt", "utf8"),
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.kms.CustomKeyStore("test",
    cloud_hsm_cluster_id=var["cloud_hsm_cluster_id"],
    custom_key_store_name="kms-custom-key-store-test",
    key_store_password="noplaintextpasswords1",
    trust_anchor_certificate=(lambda path: open(path).read())("anchor-certificate.crt"))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Kms.CustomKeyStore("test", new()
    {
        CloudHsmClusterId = @var.Cloud_hsm_cluster_id,
        CustomKeyStoreName = "kms-custom-key-store-test",
        KeyStorePassword = "noplaintextpasswords1",
        TrustAnchorCertificate = File.ReadAllText("anchor-certificate.crt"),
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
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
		_, err := kms.NewCustomKeyStore(ctx, "test", &kms.CustomKeyStoreArgs{
			CloudHsmClusterId:      pulumi.Any(_var.Cloud_hsm_cluster_id),
			CustomKeyStoreName:     pulumi.String("kms-custom-key-store-test"),
			KeyStorePassword:       pulumi.String("noplaintextpasswords1"),
			TrustAnchorCertificate: readFileOrPanic("anchor-certificate.crt"),
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
import com.pulumi.aws.kms.CustomKeyStore;
import com.pulumi.aws.kms.CustomKeyStoreArgs;
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
        var test = new CustomKeyStore("test", CustomKeyStoreArgs.builder()        
            .cloudHsmClusterId(var_.cloud_hsm_cluster_id())
            .customKeyStoreName("kms-custom-key-store-test")
            .keyStorePassword("noplaintextpasswords1")
            .trustAnchorCertificate(Files.readString(Paths.get("anchor-certificate.crt")))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:kms:CustomKeyStore
    properties:
      cloudHsmClusterId: ${var.cloud_hsm_cluster_id}
      customKeyStoreName: kms-custom-key-store-test
      keyStorePassword: noplaintextpasswords1
      trustAnchorCertificate:
        fn::readFile: anchor-certificate.crt
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import KMS (Key Management) Custom Key Store using the `id`. For example:

```sh
 $ pulumi import aws:kms/customKeyStore:CustomKeyStore example cks-5ebd4ef395a96288e
```
 