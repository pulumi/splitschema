Provides a SageMaker Space resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.Space("example", {
    domainId: aws_sagemaker_domain.test.id,
    spaceName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.Space("example",
    domain_id=aws_sagemaker_domain["test"]["id"],
    space_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.Space("example", new()
    {
        DomainId = aws_sagemaker_domain.Test.Id,
        SpaceName = "example",
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
		_, err := sagemaker.NewSpace(ctx, "example", &sagemaker.SpaceArgs{
			DomainId:  pulumi.Any(aws_sagemaker_domain.Test.Id),
			SpaceName: pulumi.String("example"),
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
import com.pulumi.aws.sagemaker.Space;
import com.pulumi.aws.sagemaker.SpaceArgs;
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
        var example = new Space("example", SpaceArgs.builder()        
            .domainId(aws_sagemaker_domain.test().id())
            .spaceName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:Space
    properties:
      domainId: ${aws_sagemaker_domain.test.id}
      spaceName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Spaces using the `id`. For example:

```sh
 $ pulumi import aws:sagemaker/space:Space test_space arn:aws:sagemaker:us-west-2:123456789012:space/domain-id/space-name
```
 