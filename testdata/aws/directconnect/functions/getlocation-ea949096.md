Retrieve information about a specific AWS Direct Connect location in the current AWS Region.
These are the locations that can be specified when configuring `aws.directconnect.Connection` or `aws.directconnect.LinkAggregationGroup` resources.

> **Note:** This data source is different from the `aws.directconnect.getLocations` data source which retrieves information about all the AWS Direct Connect locations in the current AWS Region.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.directconnect.getLocation({
    locationCode: "CS32A-24FL",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directconnect.get_location(location_code="CS32A-24FL")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.DirectConnect.GetLocation.Invoke(new()
    {
        LocationCode = "CS32A-24FL",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.GetLocation(ctx, &directconnect.GetLocationArgs{
			LocationCode: "CS32A-24FL",
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
import com.pulumi.aws.directconnect.DirectconnectFunctions;
import com.pulumi.aws.directconnect.inputs.GetLocationArgs;
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
        final var example = DirectconnectFunctions.getLocation(GetLocationArgs.builder()
            .locationCode("CS32A-24FL")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:directconnect:getLocation
      Arguments:
        locationCode: CS32A-24FL
```
{{% /example %}}
{{% /examples %}}