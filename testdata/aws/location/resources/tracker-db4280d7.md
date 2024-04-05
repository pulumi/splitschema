Provides a Location Service Tracker.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.location.Tracker("example", {trackerName: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.location.Tracker("example", tracker_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Location.Tracker("example", new()
    {
        TrackerName = "example",
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
		_, err := location.NewTracker(ctx, "example", &location.TrackerArgs{
			TrackerName: pulumi.String("example"),
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
import com.pulumi.aws.location.Tracker;
import com.pulumi.aws.location.TrackerArgs;
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
        var example = new Tracker("example", TrackerArgs.builder()        
            .trackerName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:location:Tracker
    properties:
      trackerName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_location_tracker` resources using the tracker name. For example:

```sh
 $ pulumi import aws:location/tracker:Tracker example example
```
 