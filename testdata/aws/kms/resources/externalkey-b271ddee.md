Manages a single-Region or multi-Region primary KMS key that uses external key material.
To instead manage a single-Region or multi-Region primary KMS key where AWS automatically generates and potentially rotates key material, see the `aws.kms.Key` resource.


{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kms.ExternalKey("example", {description: "KMS EXTERNAL for AMI encryption"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kms.ExternalKey("example", description="KMS EXTERNAL for AMI encryption")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kms.ExternalKey("example", new()
    {
        Description = "KMS EXTERNAL for AMI encryption",
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
		_, err := kms.NewExternalKey(ctx, "example", &kms.ExternalKeyArgs{
			Description: pulumi.String("KMS EXTERNAL for AMI encryption"),
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
import com.pulumi.aws.kms.ExternalKey;
import com.pulumi.aws.kms.ExternalKeyArgs;
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
        var example = new ExternalKey("example", ExternalKeyArgs.builder()        
            .description("KMS EXTERNAL for AMI encryption")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kms:ExternalKey
    properties:
      description: KMS EXTERNAL for AMI encryption
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import KMS External Keys using the `id`. For example:

```sh
 $ pulumi import aws:kms/externalKey:ExternalKey a arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
```
 