Use this data source to get information about an EBS volume for use in other
resources.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ebsVolume = aws.ebs.getVolume({
    filters: [
        {
            name: "volume-type",
            values: ["gp2"],
        },
        {
            name: "tag:Name",
            values: ["Example"],
        },
    ],
    mostRecent: true,
});
```
```python
import pulumi
import pulumi_aws as aws

ebs_volume = aws.ebs.get_volume(filters=[
        aws.ebs.GetVolumeFilterArgs(
            name="volume-type",
            values=["gp2"],
        ),
        aws.ebs.GetVolumeFilterArgs(
            name="tag:Name",
            values=["Example"],
        ),
    ],
    most_recent=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ebsVolume = Aws.Ebs.GetVolume.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ebs.Inputs.GetVolumeFilterInputArgs
            {
                Name = "volume-type",
                Values = new[]
                {
                    "gp2",
                },
            },
            new Aws.Ebs.Inputs.GetVolumeFilterInputArgs
            {
                Name = "tag:Name",
                Values = new[]
                {
                    "Example",
                },
            },
        },
        MostRecent = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ebs.LookupVolume(ctx, &ebs.LookupVolumeArgs{
			Filters: []ebs.GetVolumeFilter{
				{
					Name: "volume-type",
					Values: []string{
						"gp2",
					},
				},
				{
					Name: "tag:Name",
					Values: []string{
						"Example",
					},
				},
			},
			MostRecent: pulumi.BoolRef(true),
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
import com.pulumi.aws.ebs.EbsFunctions;
import com.pulumi.aws.ebs.inputs.GetVolumeArgs;
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
        final var ebsVolume = EbsFunctions.getVolume(GetVolumeArgs.builder()
            .filters(            
                GetVolumeFilterArgs.builder()
                    .name("volume-type")
                    .values("gp2")
                    .build(),
                GetVolumeFilterArgs.builder()
                    .name("tag:Name")
                    .values("Example")
                    .build())
            .mostRecent(true)
            .build());

    }
}
```
```yaml
variables:
  ebsVolume:
    fn::invoke:
      Function: aws:ebs:getVolume
      Arguments:
        filters:
          - name: volume-type
            values:
              - gp2
          - name: tag:Name
            values:
              - Example
        mostRecent: true
```
{{% /example %}}
{{% /examples %}}