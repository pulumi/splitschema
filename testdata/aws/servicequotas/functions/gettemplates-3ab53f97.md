Data source for managing an AWS Service Quotas Templates.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.servicequotas.getTemplates({
    region: "us-east-1",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicequotas.get_templates(region="us-east-1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ServiceQuotas.GetTemplates.Invoke(new()
    {
        Region = "us-east-1",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicequotas"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicequotas.GetTemplates(ctx, &servicequotas.GetTemplatesArgs{
			Region: "us-east-1",
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
import com.pulumi.aws.servicequotas.ServicequotasFunctions;
import com.pulumi.aws.servicequotas.inputs.GetTemplatesArgs;
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
        final var example = ServicequotasFunctions.getTemplates(GetTemplatesArgs.builder()
            .region("us-east-1")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:servicequotas:getTemplates
      Arguments:
        region: us-east-1
```
{{% /example %}}
{{% /examples %}}