Provides an DocumentDB subnet group resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.docdb.SubnetGroup("default", {
    subnetIds: [
        aws_subnet.frontend.id,
        aws_subnet.backend.id,
    ],
    tags: {
        Name: "My docdb subnet group",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.docdb.SubnetGroup("default",
    subnet_ids=[
        aws_subnet["frontend"]["id"],
        aws_subnet["backend"]["id"],
    ],
    tags={
        "Name": "My docdb subnet group",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.DocDB.SubnetGroup("default", new()
    {
        SubnetIds = new[]
        {
            aws_subnet.Frontend.Id,
            aws_subnet.Backend.Id,
        },
        Tags = 
        {
            { "Name", "My docdb subnet group" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/docdb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := docdb.NewSubnetGroup(ctx, "default", &docdb.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				aws_subnet.Frontend.Id,
				aws_subnet.Backend.Id,
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("My docdb subnet group"),
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
import com.pulumi.aws.docdb.SubnetGroup;
import com.pulumi.aws.docdb.SubnetGroupArgs;
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
            .tags(Map.of("Name", "My docdb subnet group"))
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:docdb:SubnetGroup
    properties:
      subnetIds:
        - ${aws_subnet.frontend.id}
        - ${aws_subnet.backend.id}
      tags:
        Name: My docdb subnet group
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DocumentDB Subnet groups using the `name`. For example:

```sh
 $ pulumi import aws:docdb/subnetGroup:SubnetGroup default production-subnet-group
```
 