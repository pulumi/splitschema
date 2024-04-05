Provides a Location Service Place Index.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.location.PlaceIndex("example", {
    dataSource: "Here",
    indexName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.location.PlaceIndex("example",
    data_source="Here",
    index_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Location.PlaceIndex("example", new()
    {
        DataSource = "Here",
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
		_, err := location.NewPlaceIndex(ctx, "example", &location.PlaceIndexArgs{
			DataSource: pulumi.String("Here"),
			IndexName:  pulumi.String("example"),
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
import com.pulumi.aws.location.PlaceIndex;
import com.pulumi.aws.location.PlaceIndexArgs;
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
        var example = new PlaceIndex("example", PlaceIndexArgs.builder()        
            .dataSource("Here")
            .indexName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:location:PlaceIndex
    properties:
      dataSource: Here
      indexName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_location_place_index` resources using the place index name. For example:

```sh
 $ pulumi import aws:location/placeIndex:PlaceIndex example example
```
 