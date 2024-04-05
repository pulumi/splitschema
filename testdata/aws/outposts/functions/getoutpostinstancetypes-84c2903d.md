Information about Outposts Instance Types.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.outposts.getOutpostInstanceTypes({
    arn: data.aws_outposts_outpost.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.outposts.get_outpost_instance_types(arn=data["aws_outposts_outpost"]["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Outposts.GetOutpostInstanceTypes.Invoke(new()
    {
        Arn = data.Aws_outposts_outpost.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/outposts"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := outposts.GetOutpostInstanceTypes(ctx, &outposts.GetOutpostInstanceTypesArgs{
			Arn: data.Aws_outposts_outpost.Example.Arn,
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
import com.pulumi.aws.outposts.OutpostsFunctions;
import com.pulumi.aws.outposts.inputs.GetOutpostInstanceTypesArgs;
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
        final var example = OutpostsFunctions.getOutpostInstanceTypes(GetOutpostInstanceTypesArgs.builder()
            .arn(data.aws_outposts_outpost().example().arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:outposts:getOutpostInstanceTypes
      Arguments:
        arn: ${data.aws_outposts_outpost.example.arn}
```
{{% /example %}}
{{% /examples %}}