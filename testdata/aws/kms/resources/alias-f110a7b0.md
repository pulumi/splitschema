Provides an alias for a KMS customer master key. AWS Console enforces 1-to-1 mapping between aliases & keys,
but API (hence this provider too) allows you to create as many aliases as
the [account limits](http://docs.aws.amazon.com/kms/latest/developerguide/limits.html) allow you.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const key = new aws.kms.Key("key", {});
const alias = new aws.kms.Alias("alias", {targetKeyId: key.keyId});
```
```python
import pulumi
import pulumi_aws as aws

key = aws.kms.Key("key")
alias = aws.kms.Alias("alias", target_key_id=key.key_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var key = new Aws.Kms.Key("key");

    var @alias = new Aws.Kms.Alias("alias", new()
    {
        TargetKeyId = key.KeyId,
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
		key, err := kms.NewKey(ctx, "key", nil)
		if err != nil {
			return err
		}
		_, err = kms.NewAlias(ctx, "alias", &kms.AliasArgs{
			TargetKeyId: key.KeyId,
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.Alias;
import com.pulumi.aws.kms.AliasArgs;
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
        var key = new Key("key");

        var alias = new Alias("alias", AliasArgs.builder()        
            .targetKeyId(key.keyId())
            .build());

    }
}
```
```yaml
resources:
  key:
    type: aws:kms:Key
  alias:
    type: aws:kms:Alias
    properties:
      targetKeyId: ${key.keyId}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import KMS aliases using the `name`. For example:

```sh
 $ pulumi import aws:kms/alias:Alias a alias/my-key-alias
```
 