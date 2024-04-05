Information about a Neptune engine version.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.neptune.getEngineVersion({
    preferredVersions: [
        "1.0.3.0",
        "1.0.2.2",
        "1.0.2.1",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.neptune.get_engine_version(preferred_versions=[
    "1.0.3.0",
    "1.0.2.2",
    "1.0.2.1",
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Neptune.GetEngineVersion.Invoke(new()
    {
        PreferredVersions = new[]
        {
            "1.0.3.0",
            "1.0.2.2",
            "1.0.2.1",
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
		_, err := neptune.GetEngineVersion(ctx, &neptune.GetEngineVersionArgs{
			PreferredVersions: []string{
				"1.0.3.0",
				"1.0.2.2",
				"1.0.2.1",
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
import com.pulumi.aws.neptune.NeptuneFunctions;
import com.pulumi.aws.neptune.inputs.GetEngineVersionArgs;
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
        final var test = NeptuneFunctions.getEngineVersion(GetEngineVersionArgs.builder()
            .preferredVersions(            
                "1.0.3.0",
                "1.0.2.2",
                "1.0.2.1")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:neptune:getEngineVersion
      Arguments:
        preferredVersions:
          - 1.0.3.0
          - 1.0.2.2
          - 1.0.2.1
```
{{% /example %}}
{{% /examples %}}