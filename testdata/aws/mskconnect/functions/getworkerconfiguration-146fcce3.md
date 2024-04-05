Get information on an Amazon MSK Connect Worker Configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.mskconnect.getWorkerConfiguration({
    name: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.mskconnect.get_worker_configuration(name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.MskConnect.GetWorkerConfiguration.Invoke(new()
    {
        Name = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mskconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mskconnect.LookupWorkerConfiguration(ctx, &mskconnect.LookupWorkerConfigurationArgs{
			Name: "example",
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
import com.pulumi.aws.mskconnect.MskconnectFunctions;
import com.pulumi.aws.mskconnect.inputs.GetWorkerConfigurationArgs;
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
        final var example = MskconnectFunctions.getWorkerConfiguration(GetWorkerConfigurationArgs.builder()
            .name("example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:mskconnect:getWorkerConfiguration
      Arguments:
        name: example
```
{{% /example %}}
{{% /examples %}}