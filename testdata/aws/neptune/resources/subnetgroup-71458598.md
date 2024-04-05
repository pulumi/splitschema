Provides an Neptune subnet group resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.neptune.SubnetGroup("default", {
    subnetIds: [
        aws_subnet.frontend.id,
        aws_subnet.backend.id,
    ],
    tags: {
        Name: "My neptune subnet group",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.neptune.SubnetGroup("default",
    subnet_ids=[
        aws_subnet["frontend"]["id"],
        aws_subnet["backend"]["id"],
    ],
    tags={
        "Name": "My neptune subnet group",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Neptune.SubnetGroup("default", new()
    {
        SubnetIds = new[]
        {
            aws_subnet.Frontend.Id,
            aws_subnet.Backend.Id,
        },
        Tags = 
        {
            { "Name", "My neptune subnet group" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/neptune"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := neptune.NewSubnetGroup(ctx, "default", &neptune.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				aws_subnet.Frontend.Id,
				aws_subnet.Backend.Id,
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("My neptune subnet group"),
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
import com.pulumi.aws.neptune.SubnetGroup;
import com.pulumi.aws.neptune.SubnetGroupArgs;
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
        var default_ = new SubnetGroup("default", SubnetGroupArgs.builder()        
            .subnetIds(            
                aws_subnet.frontend().id(),
                aws_subnet.backend().id())
            .tags(Map.of("Name", "My neptune subnet group"))
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:neptune:SubnetGroup
    properties:
      subnetIds:
        - ${aws_subnet.frontend.id}
        - ${aws_subnet.backend.id}
      tags:
        Name: My neptune subnet group
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Neptune Subnet groups using the `name`. For example:

```sh
 $ pulumi import aws:neptune/subnetGroup:SubnetGroup default production-subnet-group
```
 