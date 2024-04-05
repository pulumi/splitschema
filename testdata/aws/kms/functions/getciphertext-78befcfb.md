The KMS ciphertext data source allows you to encrypt plaintext into ciphertext
by using an AWS KMS customer master key. The value returned by this data source
changes every apply. For a stable ciphertext value, see the `aws.kms.Ciphertext`
resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const oauthConfig = new aws.kms.Key("oauthConfig", {
    description: "oauth config",
    isEnabled: true,
});
const oauth = aws.kms.getCipherTextOutput({
    keyId: oauthConfig.keyId,
    plaintext: `{
  "client_id": "e587dbae22222f55da22",
  "client_secret": "8289575d00000ace55e1815ec13673955721b8a5"
}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

oauth_config = aws.kms.Key("oauthConfig",
    description="oauth config",
    is_enabled=True)
oauth = aws.kms.get_cipher_text_output(key_id=oauth_config.key_id,
    plaintext="""{
  "client_id": "e587dbae22222f55da22",
  "client_secret": "8289575d00000ace55e1815ec13673955721b8a5"
}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var oauthConfig = new Aws.Kms.Key("oauthConfig", new()
    {
        Description = "oauth config",
        IsEnabled = true,
    });

    var oauth = Aws.Kms.GetCipherText.Invoke(new()
    {
        KeyId = oauthConfig.KeyId,
        Plaintext = @"{
  ""client_id"": ""e587dbae22222f55da22"",
  ""client_secret"": ""8289575d00000ace55e1815ec13673955721b8a5""
}
",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		oauthConfig, err := kms.NewKey(ctx, "oauthConfig", &kms.KeyArgs{
			Description: pulumi.String("oauth config"),
			IsEnabled:   pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_ = kms.GetCipherTextOutput(ctx, kms.GetCipherTextOutputArgs{
			KeyId:     oauthConfig.KeyId,
			Plaintext: pulumi.String("{\n  \"client_id\": \"e587dbae22222f55da22\",\n  \"client_secret\": \"8289575d00000ace55e1815ec13673955721b8a5\"\n}\n"),
		}, nil)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.kms.KmsFunctions;
import com.pulumi.aws.kms.inputs.GetCipherTextArgs;
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
        var oauthConfig = new Key("oauthConfig", KeyArgs.builder()        
            .description("oauth config")
            .isEnabled(true)
            .build());

        final var oauth = KmsFunctions.getCipherText(GetCipherTextArgs.builder()
            .keyId(oauthConfig.keyId())
            .plaintext("""
{
  "client_id": "e587dbae22222f55da22",
  "client_secret": "8289575d00000ace55e1815ec13673955721b8a5"
}
            """)
            .build());

    }
}
```
```yaml
resources:
  oauthConfig:
    type: aws:kms:Key
    properties:
      description: oauth config
      isEnabled: true
variables:
  oauth:
    fn::invoke:
      Function: aws:kms:getCipherText
      Arguments:
        keyId: ${oauthConfig.keyId}
        plaintext: |
          {
            "client_id": "e587dbae22222f55da22",
            "client_secret": "8289575d00000ace55e1815ec13673955721b8a5"
          }
```
{{% /example %}}
{{% /examples %}}