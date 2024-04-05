Provides a resource to manage whether default EBS encryption is enabled for your AWS account in the current AWS region. To manage the default KMS key for the region, see the `aws.ebs.DefaultKmsKey` resource.

> **NOTE:** Removing this resource disables default EBS encryption.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ebs.EncryptionByDefault("example", {enabled: true});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ebs.EncryptionByDefault("example", enabled=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ebs.EncryptionByDefault("example", new()
    {
        Enabled = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ebs.NewEncryptionByDefault(ctx, "example", &ebs.EncryptionByDefaultArgs{
			Enabled: pulumi.Bool(true),
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
import com.pulumi.aws.ebs.EncryptionByDefault;
import com.pulumi.aws.ebs.EncryptionByDefaultArgs;
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
        var example = new EncryptionByDefault("example", EncryptionByDefaultArgs.builder()        
            .enabled(true)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ebs:EncryptionByDefault
    properties:
      enabled: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the default EBS encryption state. For example:

```sh
 $ pulumi import aws:ebs/encryptionByDefault:EncryptionByDefault example default
```
 