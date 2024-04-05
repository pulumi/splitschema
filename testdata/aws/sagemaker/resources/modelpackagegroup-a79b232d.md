Provides a SageMaker Model Package Group resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.ModelPackageGroup("example", {modelPackageGroupName: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.ModelPackageGroup("example", model_package_group_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.ModelPackageGroup("example", new()
    {
        ModelPackageGroupName = "example",
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
		_, err := sagemaker.NewModelPackageGroup(ctx, "example", &sagemaker.ModelPackageGroupArgs{
			ModelPackageGroupName: pulumi.String("example"),
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
import com.pulumi.aws.sagemaker.ModelPackageGroup;
import com.pulumi.aws.sagemaker.ModelPackageGroupArgs;
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
        var example = new ModelPackageGroup("example", ModelPackageGroupArgs.builder()        
            .modelPackageGroupName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:ModelPackageGroup
    properties:
      modelPackageGroupName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Model Package Groups using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/modelPackageGroup:ModelPackageGroup test_model_package_group my-code-repo
```
 