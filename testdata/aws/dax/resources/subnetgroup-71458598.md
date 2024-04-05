Provides a DAX Subnet Group resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dax.SubnetGroup("example", {subnetIds: [
    aws_subnet.example1.id,
    aws_subnet.example2.id,
]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.dax.SubnetGroup("example", subnet_ids=[
    aws_subnet["example1"]["id"],
    aws_subnet["example2"]["id"],
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Dax.SubnetGroup("example", new()
    {
        SubnetIds = new[]
        {
            aws_subnet.Example1.Id,
            aws_subnet.Example2.Id,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dax"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dax.NewSubnetGroup(ctx, "example", &dax.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				aws_subnet.Example1.Id,
				aws_subnet.Example2.Id,
			},
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
import com.pulumi.aws.dax.SubnetGroup;
import com.pulumi.aws.dax.SubnetGroupArgs;
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
        var example = new SubnetGroup("example", SubnetGroupArgs.builder()        
            .subnetIds(            
                aws_subnet.example1().id(),
                aws_subnet.example2().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:dax:SubnetGroup
    properties:
      subnetIds:
        - ${aws_subnet.example1.id}
        - ${aws_subnet.example2.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DAX Subnet Group using the `name`. For example:

```sh
 $ pulumi import aws:dax/subnetGroup:SubnetGroup example my_dax_sg
```
 