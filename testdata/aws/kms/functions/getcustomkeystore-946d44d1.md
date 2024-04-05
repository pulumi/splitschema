Use this data source to get the metadata KMS custom key store.
By using this data source, you can reference KMS custom key store
without having to hard code the ID as input.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const keystore = aws.kms.getCustomKeyStore({
    customKeyStoreName: "my_cloudhsm",
});
```
```python
import pulumi
import pulumi_aws as aws

keystore = aws.kms.get_custom_key_store(custom_key_store_name="my_cloudhsm")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var keystore = Aws.Kms.GetCustomKeyStore.Invoke(new()
    {
        CustomKeyStoreName = "my_cloudhsm",
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
		_, err := kms.LookupCustomKeyStore(ctx, &kms.LookupCustomKeyStoreArgs{
			CustomKeyStoreName: pulumi.StringRef("my_cloudhsm"),
		}, nil)
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
import com.pulumi.aws.kms.KmsFunctions;
import com.pulumi.aws.kms.inputs.GetCustomKeyStoreArgs;
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
        final var keystore = KmsFunctions.getCustomKeyStore(GetCustomKeyStoreArgs.builder()
            .customKeyStoreName("my_cloudhsm")
            .build());

    }
}
```
```yaml
variables:
  keystore:
    fn::invoke:
      Function: aws:kms:getCustomKeyStore
      Arguments:
        customKeyStoreName: my_cloudhsm
```
{{% /example %}}
{{% /examples %}}