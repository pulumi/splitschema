Use this data source to get the default EBS encryption KMS key in the current region.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.ebs.getDefaultKmsKey({});
const example = new aws.ebs.Volume("example", {
    availabilityZone: "us-west-2a",
    encrypted: true,
    kmsKeyId: current.then(current => current.keyArn),
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.ebs.get_default_kms_key()
example = aws.ebs.Volume("example",
    availability_zone="us-west-2a",
    encrypted=True,
    kms_key_id=current.key_arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.Ebs.GetDefaultKmsKey.Invoke();

    var example = new Aws.Ebs.Volume("example", new()
    {
        AvailabilityZone = "us-west-2a",
        Encrypted = true,
        KmsKeyId = current.Apply(getDefaultKmsKeyResult => getDefaultKmsKeyResult.KeyArn),
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
		current, err := ebs.LookupDefaultKmsKey(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Encrypted:        pulumi.Bool(true),
			KmsKeyId:         *pulumi.String(current.KeyArn),
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
import com.pulumi.aws.ebs.EbsFunctions;
import com.pulumi.aws.ebs.Volume;
import com.pulumi.aws.ebs.VolumeArgs;
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
        final var current = EbsFunctions.getDefaultKmsKey();

        var example = new Volume("example", VolumeArgs.builder()        
            .availabilityZone("us-west-2a")
            .encrypted(true)
            .kmsKeyId(current.applyValue(getDefaultKmsKeyResult -> getDefaultKmsKeyResult.keyArn()))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ebs:Volume
    properties:
      availabilityZone: us-west-2a
      encrypted: true
      kmsKeyId: ${current.keyArn}
variables:
  current:
    fn::invoke:
      Function: aws:ebs:getDefaultKmsKey
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}