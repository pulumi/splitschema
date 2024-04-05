Provides a DAX Parameter Group resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dax.ParameterGroup("example", {parameters: [
    {
        name: "query-ttl-millis",
        value: "100000",
    },
    {
        name: "record-ttl-millis",
        value: "100000",
    },
]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.dax.ParameterGroup("example", parameters=[
    aws.dax.ParameterGroupParameterArgs(
        name="query-ttl-millis",
        value="100000",
    ),
    aws.dax.ParameterGroupParameterArgs(
        name="record-ttl-millis",
        value="100000",
    ),
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Dax.ParameterGroup("example", new()
    {
        Parameters = new[]
        {
            new Aws.Dax.Inputs.ParameterGroupParameterArgs
            {
                Name = "query-ttl-millis",
                Value = "100000",
            },
            new Aws.Dax.Inputs.ParameterGroupParameterArgs
            {
                Name = "record-ttl-millis",
                Value = "100000",
            },
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
		_, err := dax.NewParameterGroup(ctx, "example", &dax.ParameterGroupArgs{
			Parameters: dax.ParameterGroupParameterArray{
				&dax.ParameterGroupParameterArgs{
					Name:  pulumi.String("query-ttl-millis"),
					Value: pulumi.String("100000"),
				},
				&dax.ParameterGroupParameterArgs{
					Name:  pulumi.String("record-ttl-millis"),
					Value: pulumi.String("100000"),
				},
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
import com.pulumi.aws.dax.ParameterGroup;
import com.pulumi.aws.dax.ParameterGroupArgs;
import com.pulumi.aws.dax.inputs.ParameterGroupParameterArgs;
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
        var example = new ParameterGroup("example", ParameterGroupArgs.builder()        
            .parameters(            
                ParameterGroupParameterArgs.builder()
                    .name("query-ttl-millis")
                    .value("100000")
                    .build(),
                ParameterGroupParameterArgs.builder()
                    .name("record-ttl-millis")
                    .value("100000")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:dax:ParameterGroup
    properties:
      parameters:
        - name: query-ttl-millis
          value: '100000'
        - name: record-ttl-millis
          value: '100000'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DAX Parameter Group using the `name`. For example:

```sh
 $ pulumi import aws:dax/parameterGroup:ParameterGroup example my_dax_pg
```
 