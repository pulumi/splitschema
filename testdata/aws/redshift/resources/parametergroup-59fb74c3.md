Provides a Redshift Cluster parameter group resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.redshift.ParameterGroup("bar", {
    family: "redshift-1.0",
    parameters: [
        {
            name: "require_ssl",
            value: "true",
        },
        {
            name: "query_group",
            value: "example",
        },
        {
            name: "enable_user_activity_logging",
            value: "true",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

bar = aws.redshift.ParameterGroup("bar",
    family="redshift-1.0",
    parameters=[
        aws.redshift.ParameterGroupParameterArgs(
            name="require_ssl",
            value="true",
        ),
        aws.redshift.ParameterGroupParameterArgs(
            name="query_group",
            value="example",
        ),
        aws.redshift.ParameterGroupParameterArgs(
            name="enable_user_activity_logging",
            value="true",
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
    var bar = new Aws.RedShift.ParameterGroup("bar", new()
    {
        Family = "redshift-1.0",
        Parameters = new[]
        {
            new Aws.RedShift.Inputs.ParameterGroupParameterArgs
            {
                Name = "require_ssl",
                Value = "true",
            },
            new Aws.RedShift.Inputs.ParameterGroupParameterArgs
            {
                Name = "query_group",
                Value = "example",
            },
            new Aws.RedShift.Inputs.ParameterGroupParameterArgs
            {
                Name = "enable_user_activity_logging",
                Value = "true",
            },
        },
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
		_, err := redshift.NewParameterGroup(ctx, "bar", &redshift.ParameterGroupArgs{
			Family: pulumi.String("redshift-1.0"),
			Parameters: redshift.ParameterGroupParameterArray{
				&redshift.ParameterGroupParameterArgs{
					Name:  pulumi.String("require_ssl"),
					Value: pulumi.String("true"),
				},
				&redshift.ParameterGroupParameterArgs{
					Name:  pulumi.String("query_group"),
					Value: pulumi.String("example"),
				},
				&redshift.ParameterGroupParameterArgs{
					Name:  pulumi.String("enable_user_activity_logging"),
					Value: pulumi.String("true"),
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
import com.pulumi.aws.redshift.ParameterGroup;
import com.pulumi.aws.redshift.ParameterGroupArgs;
import com.pulumi.aws.redshift.inputs.ParameterGroupParameterArgs;
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
        var bar = new ParameterGroup("bar", ParameterGroupArgs.builder()        
            .family("redshift-1.0")
            .parameters(            
                ParameterGroupParameterArgs.builder()
                    .name("require_ssl")
                    .value("true")
                    .build(),
                ParameterGroupParameterArgs.builder()
                    .name("query_group")
                    .value("example")
                    .build(),
                ParameterGroupParameterArgs.builder()
                    .name("enable_user_activity_logging")
                    .value("true")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  bar:
    type: aws:redshift:ParameterGroup
    properties:
      family: redshift-1.0
      parameters:
        - name: require_ssl
          value: 'true'
        - name: query_group
          value: example
        - name: enable_user_activity_logging
          value: 'true'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Parameter Groups using the `name`. For example:

```sh
 $ pulumi import aws:redshift/parameterGroup:ParameterGroup paramgroup1 parameter-group-test-pulumi
```
 