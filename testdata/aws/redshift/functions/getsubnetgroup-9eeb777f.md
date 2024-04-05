Provides details about a specific redshift subnet group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.redshift.getSubnetGroup({
    name: aws_redshift_subnet_group.example.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.get_subnet_group(name=aws_redshift_subnet_group["example"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.RedShift.GetSubnetGroup.Invoke(new()
    {
        Name = aws_redshift_subnet_group.Example.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.LookupSubnetGroup(ctx, &redshift.LookupSubnetGroupArgs{
			Name: aws_redshift_subnet_group.Example.Name,
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
import com.pulumi.aws.redshift.RedshiftFunctions;
import com.pulumi.aws.redshift.inputs.GetSubnetGroupArgs;
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
        final var example = RedshiftFunctions.getSubnetGroup(GetSubnetGroupArgs.builder()
            .name(aws_redshift_subnet_group.example().name())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:redshift:getSubnetGroup
      Arguments:
        name: ${aws_redshift_subnet_group.example.name}
```
{{% /example %}}
{{% /examples %}}