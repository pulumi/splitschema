Retrieve information about EMR Release Labels.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.emr.getReleaseLabels({
    filters: {
        application: "spark@2.1.0",
        prefix: "emr-5",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.emr.get_release_labels(filters=aws.emr.GetReleaseLabelsFiltersArgs(
    application="spark@2.1.0",
    prefix="emr-5",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Emr.GetReleaseLabels.Invoke(new()
    {
        Filters = new Aws.Emr.Inputs.GetReleaseLabelsFiltersInputArgs
        {
            Application = "spark@2.1.0",
            Prefix = "emr-5",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.GetReleaseLabels(ctx, &emr.GetReleaseLabelsArgs{
			Filters: emr.GetReleaseLabelsFilters{
				Application: pulumi.StringRef("spark@2.1.0"),
				Prefix:      pulumi.StringRef("emr-5"),
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
import com.pulumi.aws.emr.EmrFunctions;
import com.pulumi.aws.emr.inputs.GetReleaseLabelsArgs;
import com.pulumi.aws.emr.inputs.GetReleaseLabelsFiltersArgs;
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
        final var example = EmrFunctions.getReleaseLabels(GetReleaseLabelsArgs.builder()
            .filters(GetReleaseLabelsFiltersArgs.builder()
                .application("spark@2.1.0")
                .prefix("emr-5")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:emr:getReleaseLabels
      Arguments:
        filters:
          application: spark@2.1.0
          prefix: emr-5
```
{{% /example %}}
{{% /examples %}}