Use this data source to get information about an EC2 Dedicated Host.

{{% examples %}}
## Example Usage
{{% example %}}
### Filter Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.ec2.getDedicatedHost({
    filters: [{
        name: "instance-type",
        values: ["c5.18xlarge"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ec2.get_dedicated_host(filters=[aws.ec2.GetDedicatedHostFilterArgs(
    name="instance-type",
    values=["c5.18xlarge"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Ec2.GetDedicatedHost.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetDedicatedHostFilterInputArgs
            {
                Name = "instance-type",
                Values = new[]
                {
                    "c5.18xlarge",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.LookupDedicatedHost(ctx, &ec2.LookupDedicatedHostArgs{
			Filters: []ec2.GetDedicatedHostFilter{
				{
					Name: "instance-type",
					Values: []string{
						"c5.18xlarge",
					},
				},
			},
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetDedicatedHostArgs;
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
        final var test = Ec2Functions.getDedicatedHost(GetDedicatedHostArgs.builder()
            .filters(GetDedicatedHostFilterArgs.builder()
                .name("instance-type")
                .values("c5.18xlarge")
                .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:ec2:getDedicatedHost
      Arguments:
        filters:
          - name: instance-type
            values:
              - c5.18xlarge
```
{{% /example %}}
{{% /examples %}}