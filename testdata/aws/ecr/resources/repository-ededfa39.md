Provides an Elastic Container Registry Repository.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ecr.Repository("foo", {
    imageScanningConfiguration: {
        scanOnPush: true,
    },
    imageTagMutability: "MUTABLE",
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ecr.Repository("foo",
    image_scanning_configuration=aws.ecr.RepositoryImageScanningConfigurationArgs(
        scan_on_push=True,
    ),
    image_tag_mutability="MUTABLE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ecr.Repository("foo", new()
    {
        ImageScanningConfiguration = new Aws.Ecr.Inputs.RepositoryImageScanningConfigurationArgs
        {
            ScanOnPush = true,
        },
        ImageTagMutability = "MUTABLE",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecr.NewRepository(ctx, "foo", &ecr.RepositoryArgs{
			ImageScanningConfiguration: &ecr.RepositoryImageScanningConfigurationArgs{
				ScanOnPush: pulumi.Bool(true),
			},
			ImageTagMutability: pulumi.String("MUTABLE"),
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
import com.pulumi.aws.ecr.Repository;
import com.pulumi.aws.ecr.RepositoryArgs;
import com.pulumi.aws.ecr.inputs.RepositoryImageScanningConfigurationArgs;
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
        var foo = new Repository("foo", RepositoryArgs.builder()        
            .imageScanningConfiguration(RepositoryImageScanningConfigurationArgs.builder()
                .scanOnPush(true)
                .build())
            .imageTagMutability("MUTABLE")
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ecr:Repository
    properties:
      imageScanningConfiguration:
        scanOnPush: true
      imageTagMutability: MUTABLE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECR Repositories using the `name`. For example:

```sh
 $ pulumi import aws:ecr/repository:Repository service test-service
```
 