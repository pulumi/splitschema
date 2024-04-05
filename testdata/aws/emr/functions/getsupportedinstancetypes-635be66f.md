Data source for managing AWS EMR Supported Instance Types.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.emr.getSupportedInstanceTypes({
    releaseLabel: "ebs-6.15.0",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.emr.get_supported_instance_types(release_label="ebs-6.15.0")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Emr.GetSupportedInstanceTypes.Invoke(new()
    {
        ReleaseLabel = "ebs-6.15.0",
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
		_, err := emr.GetSupportedInstanceTypes(ctx, &emr.GetSupportedInstanceTypesArgs{
			ReleaseLabel: "ebs-6.15.0",
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
import com.pulumi.aws.emr.inputs.GetSupportedInstanceTypesArgs;
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
        final var example = EmrFunctions.getSupportedInstanceTypes(GetSupportedInstanceTypesArgs.builder()
            .releaseLabel("ebs-6.15.0")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:emr:getSupportedInstanceTypes
      Arguments:
        releaseLabel: ebs-6.15.0
```
{{% /example %}}
{{% /examples %}}