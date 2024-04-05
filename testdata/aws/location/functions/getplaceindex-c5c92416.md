Retrieve information about a Location Service Place Index.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.location.getPlaceIndex({
    indexName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.location.get_place_index(index_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Location.GetPlaceIndex.Invoke(new()
    {
        IndexName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/location"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := location.LookupPlaceIndex(ctx, &location.LookupPlaceIndexArgs{
			IndexName: "example",
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
import com.pulumi.aws.location.LocationFunctions;
import com.pulumi.aws.location.inputs.GetPlaceIndexArgs;
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
        final var example = LocationFunctions.getPlaceIndex(GetPlaceIndexArgs.builder()
            .indexName("example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:location:getPlaceIndex
      Arguments:
        indexName: example
```
{{% /example %}}
{{% /examples %}}