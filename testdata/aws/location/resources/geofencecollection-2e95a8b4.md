Resource for managing an AWS Location Geofence Collection.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.location.GeofenceCollection("example", {collectionName: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.location.GeofenceCollection("example", collection_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Location.GeofenceCollection("example", new()
    {
        CollectionName = "example",
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
		_, err := location.NewGeofenceCollection(ctx, "example", &location.GeofenceCollectionArgs{
			CollectionName: pulumi.String("example"),
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
import com.pulumi.aws.location.GeofenceCollection;
import com.pulumi.aws.location.GeofenceCollectionArgs;
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
        var example = new GeofenceCollection("example", GeofenceCollectionArgs.builder()        
            .collectionName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:location:GeofenceCollection
    properties:
      collectionName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Location Geofence Collection using the `collection_name`. For example:

```sh
 $ pulumi import aws:location/geofenceCollection:GeofenceCollection example example
```
 