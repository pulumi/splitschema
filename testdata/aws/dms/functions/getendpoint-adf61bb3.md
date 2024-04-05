Data source for managing an AWS DMS (Database Migration) Endpoint.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.dms.getEndpoint({
    endpointId: "test_id",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.dms.get_endpoint(endpoint_id="test_id")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Dms.GetEndpoint.Invoke(new()
    {
        EndpointId = "test_id",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.LookupEndpoint(ctx, &dms.LookupEndpointArgs{
			EndpointId: "test_id",
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
import com.pulumi.aws.dms.DmsFunctions;
import com.pulumi.aws.dms.inputs.GetEndpointArgs;
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
        final var test = DmsFunctions.getEndpoint(GetEndpointArgs.builder()
            .endpointId("test_id")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:dms:getEndpoint
      Arguments:
        endpointId: test_id
```
{{% /example %}}
{{% /examples %}}