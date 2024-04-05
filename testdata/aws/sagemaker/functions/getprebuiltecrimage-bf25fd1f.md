Get information about prebuilt Amazon SageMaker Docker images.

> **NOTE:** The AWS provider creates a validly constructed `registry_path` but does not verify that the `registry_path` corresponds to an existing image. For example, using a `registry_path` containing an `image_tag` that does not correspond to a Docker image in the ECR repository, will result in an error.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.sagemaker.getPrebuiltEcrImage({
    imageTag: "2.2-1.0.11.0",
    repositoryName: "sagemaker-scikit-learn",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.sagemaker.get_prebuilt_ecr_image(image_tag="2.2-1.0.11.0",
    repository_name="sagemaker-scikit-learn")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Sagemaker.GetPrebuiltEcrImage.Invoke(new()
    {
        ImageTag = "2.2-1.0.11.0",
        RepositoryName = "sagemaker-scikit-learn",
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
		_, err := sagemaker.GetPrebuiltEcrImage(ctx, &sagemaker.GetPrebuiltEcrImageArgs{
			ImageTag:       pulumi.StringRef("2.2-1.0.11.0"),
			RepositoryName: "sagemaker-scikit-learn",
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
import com.pulumi.aws.sagemaker.SagemakerFunctions;
import com.pulumi.aws.sagemaker.inputs.GetPrebuiltEcrImageArgs;
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
        final var test = SagemakerFunctions.getPrebuiltEcrImage(GetPrebuiltEcrImageArgs.builder()
            .imageTag("2.2-1.0.11.0")
            .repositoryName("sagemaker-scikit-learn")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:sagemaker:getPrebuiltEcrImage
      Arguments:
        imageTag: 2.2-1.0.11.0
        repositoryName: sagemaker-scikit-learn
```
{{% /example %}}
{{% /examples %}}