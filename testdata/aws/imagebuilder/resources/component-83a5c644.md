Manages an Image Builder Component.

{{% examples %}}
## Example Usage
{{% example %}}
### URI Document

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.imagebuilder.Component("example", {
    platform: "Linux",
    uri: `s3://${aws_s3_object.example.bucket}/${aws_s3_object.example.key}`,
    version: "1.0.0",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.Component("example",
    platform="Linux",
    uri=f"s3://{aws_s3_object['example']['bucket']}/{aws_s3_object['example']['key']}",
    version="1.0.0")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ImageBuilder.Component("example", new()
    {
        Platform = "Linux",
        Uri = $"s3://{aws_s3_object.Example.Bucket}/{aws_s3_object.Example.Key}",
        Version = "1.0.0",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/imagebuilder"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := imagebuilder.NewComponent(ctx, "example", &imagebuilder.ComponentArgs{
			Platform: pulumi.String("Linux"),
			Uri:      pulumi.String(fmt.Sprintf("s3://%v/%v", aws_s3_object.Example.Bucket, aws_s3_object.Example.Key)),
			Version:  pulumi.String("1.0.0"),
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
import com.pulumi.aws.imagebuilder.Component;
import com.pulumi.aws.imagebuilder.ComponentArgs;
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
        var example = new Component("example", ComponentArgs.builder()        
            .platform("Linux")
            .uri(String.format("s3://%s/%s", aws_s3_object.example().bucket(),aws_s3_object.example().key()))
            .version("1.0.0")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:imagebuilder:Component
    properties:
      platform: Linux
      uri: s3://${aws_s3_object.example.bucket}/${aws_s3_object.example.key}
      version: 1.0.0
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_imagebuilder_components` resources using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:imagebuilder/component:Component example arn:aws:imagebuilder:us-east-1:123456789012:component/example/1.0.0/1
```
 Certain resource arguments, such as `uri`, cannot be read via the API and imported into the provider. The provider will display a difference for these arguments the first run after import if declared in the the provider configuration for an imported resource.

