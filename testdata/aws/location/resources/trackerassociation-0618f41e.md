Resource for managing an AWS Location Tracker Association.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGeofenceCollection = new aws.location.GeofenceCollection("exampleGeofenceCollection", {collectionName: "example"});
const exampleTracker = new aws.location.Tracker("exampleTracker", {trackerName: "example"});
const exampleTrackerAssociation = new aws.location.TrackerAssociation("exampleTrackerAssociation", {
    consumerArn: exampleGeofenceCollection.collectionArn,
    trackerName: exampleTracker.trackerName,
});
```
```python
import pulumi
import pulumi_aws as aws

example_geofence_collection = aws.location.GeofenceCollection("exampleGeofenceCollection", collection_name="example")
example_tracker = aws.location.Tracker("exampleTracker", tracker_name="example")
example_tracker_association = aws.location.TrackerAssociation("exampleTrackerAssociation",
    consumer_arn=example_geofence_collection.collection_arn,
    tracker_name=example_tracker.tracker_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleGeofenceCollection = new Aws.Location.GeofenceCollection("exampleGeofenceCollection", new()
    {
        CollectionName = "example",
    });

    var exampleTracker = new Aws.Location.Tracker("exampleTracker", new()
    {
        TrackerName = "example",
    });

    var exampleTrackerAssociation = new Aws.Location.TrackerAssociation("exampleTrackerAssociation", new()
    {
        ConsumerArn = exampleGeofenceCollection.CollectionArn,
        TrackerName = exampleTracker.TrackerName,
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
		exampleGeofenceCollection, err := location.NewGeofenceCollection(ctx, "exampleGeofenceCollection", &location.GeofenceCollectionArgs{
			CollectionName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		exampleTracker, err := location.NewTracker(ctx, "exampleTracker", &location.TrackerArgs{
			TrackerName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = location.NewTrackerAssociation(ctx, "exampleTrackerAssociation", &location.TrackerAssociationArgs{
			ConsumerArn: exampleGeofenceCollection.CollectionArn,
			TrackerName: exampleTracker.TrackerName,
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
import com.pulumi.aws.location.Tracker;
import com.pulumi.aws.location.TrackerArgs;
import com.pulumi.aws.location.TrackerAssociation;
import com.pulumi.aws.location.TrackerAssociationArgs;
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
        var exampleGeofenceCollection = new GeofenceCollection("exampleGeofenceCollection", GeofenceCollectionArgs.builder()        
            .collectionName("example")
            .build());

        var exampleTracker = new Tracker("exampleTracker", TrackerArgs.builder()        
            .trackerName("example")
            .build());

        var exampleTrackerAssociation = new TrackerAssociation("exampleTrackerAssociation", TrackerAssociationArgs.builder()        
            .consumerArn(exampleGeofenceCollection.collectionArn())
            .trackerName(exampleTracker.trackerName())
            .build());

    }
}
```
```yaml
resources:
  exampleGeofenceCollection:
    type: aws:location:GeofenceCollection
    properties:
      collectionName: example
  exampleTracker:
    type: aws:location:Tracker
    properties:
      trackerName: example
  exampleTrackerAssociation:
    type: aws:location:TrackerAssociation
    properties:
      consumerArn: ${exampleGeofenceCollection.collectionArn}
      trackerName: ${exampleTracker.trackerName}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Location Tracker Association using the `tracker_name|consumer_arn`. For example:

```sh
 $ pulumi import aws:location/trackerAssociation:TrackerAssociation example "tracker_name|consumer_arn"
```
 