The ECR Image data source allows the details of an image with a particular tag or digest to be retrieved.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceImage = aws.ecr.getImage({
    imageTag: "latest",
    repositoryName: "my/service",
});
```
```python
import pulumi
import pulumi_aws as aws

service_image = aws.ecr.get_image(image_tag="latest",
    repository_name="my/service")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var serviceImage = Aws.Ecr.GetImage.Invoke(new()
    {
        ImageTag = "latest",
        RepositoryName = "my/service",
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
		_, err := ecr.GetImage(ctx, &ecr.GetImageArgs{
			ImageTag:       pulumi.StringRef("latest"),
			RepositoryName: "my/service",
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
import com.pulumi.aws.ecr.EcrFunctions;
import com.pulumi.aws.ecr.inputs.GetImageArgs;
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
        final var serviceImage = EcrFunctions.getImage(GetImageArgs.builder()
            .imageTag("latest")
            .repositoryName("my/service")
            .build());

    }
}
```
```yaml
variables:
  serviceImage:
    fn::invoke:
      Function: aws:ecr:getImage
      Arguments:
        imageTag: latest
        repositoryName: my/service
```
{{% /example %}}
{{% /examples %}}