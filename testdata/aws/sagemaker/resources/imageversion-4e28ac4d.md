Provides a SageMaker Image Version resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.sagemaker.ImageVersion("test", {
    imageName: aws_sagemaker_image.test.id,
    baseImage: "012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.sagemaker.ImageVersion("test",
    image_name=aws_sagemaker_image["test"]["id"],
    base_image="012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Sagemaker.ImageVersion("test", new()
    {
        ImageName = aws_sagemaker_image.Test.Id,
        BaseImage = "012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewImageVersion(ctx, "test", &sagemaker.ImageVersionArgs{
			ImageName: pulumi.Any(aws_sagemaker_image.Test.Id),
			BaseImage: pulumi.String("012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest"),
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
import com.pulumi.aws.sagemaker.ImageVersion;
import com.pulumi.aws.sagemaker.ImageVersionArgs;
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
        var test = new ImageVersion("test", ImageVersionArgs.builder()        
            .imageName(aws_sagemaker_image.test().id())
            .baseImage("012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:sagemaker:ImageVersion
    properties:
      imageName: ${aws_sagemaker_image.test.id}
      baseImage: 012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Image Versions using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/imageVersion:ImageVersion test_image my-code-repo
```
 