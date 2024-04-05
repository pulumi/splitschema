Provides a SageMaker Image resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.Image("example", {
    imageName: "example",
    roleArn: aws_iam_role.test.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.Image("example",
    image_name="example",
    role_arn=aws_iam_role["test"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.Image("example", new()
    {
        ImageName = "example",
        RoleArn = aws_iam_role.Test.Arn,
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
		_, err := sagemaker.NewImage(ctx, "example", &sagemaker.ImageArgs{
			ImageName: pulumi.String("example"),
			RoleArn:   pulumi.Any(aws_iam_role.Test.Arn),
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
import com.pulumi.aws.sagemaker.Image;
import com.pulumi.aws.sagemaker.ImageArgs;
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
        var example = new Image("example", ImageArgs.builder()        
            .imageName("example")
            .roleArn(aws_iam_role.test().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:Image
    properties:
      imageName: example
      roleArn: ${aws_iam_role.test.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Code Images using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/image:Image test_image my-code-repo
```
 