Manages IoT event configurations.

> **NOTE:** Deleting this resource does not disable the event configurations, the resource in simply removed from state instead.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iot.EventConfigurations("example", {eventConfigurations: {
    CA_CERTIFICATE: false,
    CERTIFICATE: true,
    JOB: false,
    JOB_EXECUTION: false,
    POLICY: false,
    THING: true,
    THING_GROUP: false,
    THING_GROUP_HIERARCHY: false,
    THING_GROUP_MEMBERSHIP: false,
    THING_TYPE: false,
    THING_TYPE_ASSOCIATION: false,
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iot.EventConfigurations("example", event_configurations={
    "CA_CERTIFICATE": False,
    "CERTIFICATE": True,
    "JOB": False,
    "JOB_EXECUTION": False,
    "POLICY": False,
    "THING": True,
    "THING_GROUP": False,
    "THING_GROUP_HIERARCHY": False,
    "THING_GROUP_MEMBERSHIP": False,
    "THING_TYPE": False,
    "THING_TYPE_ASSOCIATION": False,
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Iot.EventConfigurations("example", new()
    {
        Configurations = 
        {
            { "CA_CERTIFICATE", false },
            { "CERTIFICATE", true },
            { "JOB", false },
            { "JOB_EXECUTION", false },
            { "POLICY", false },
            { "THING", true },
            { "THING_GROUP", false },
            { "THING_GROUP_HIERARCHY", false },
            { "THING_GROUP_MEMBERSHIP", false },
            { "THING_TYPE", false },
            { "THING_TYPE_ASSOCIATION", false },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iot.NewEventConfigurations(ctx, "example", &iot.EventConfigurationsArgs{
			EventConfigurations: pulumi.BoolMap{
				"CA_CERTIFICATE":         pulumi.Bool(false),
				"CERTIFICATE":            pulumi.Bool(true),
				"JOB":                    pulumi.Bool(false),
				"JOB_EXECUTION":          pulumi.Bool(false),
				"POLICY":                 pulumi.Bool(false),
				"THING":                  pulumi.Bool(true),
				"THING_GROUP":            pulumi.Bool(false),
				"THING_GROUP_HIERARCHY":  pulumi.Bool(false),
				"THING_GROUP_MEMBERSHIP": pulumi.Bool(false),
				"THING_TYPE":             pulumi.Bool(false),
				"THING_TYPE_ASSOCIATION": pulumi.Bool(false),
			},
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
import com.pulumi.aws.iot.EventConfigurations;
import com.pulumi.aws.iot.EventConfigurationsArgs;
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
        var example = new EventConfigurations("example", EventConfigurationsArgs.builder()        
            .eventConfigurations(Map.ofEntries(
                Map.entry("CA_CERTIFICATE", false),
                Map.entry("CERTIFICATE", true),
                Map.entry("JOB", false),
                Map.entry("JOB_EXECUTION", false),
                Map.entry("POLICY", false),
                Map.entry("THING", true),
                Map.entry("THING_GROUP", false),
                Map.entry("THING_GROUP_HIERARCHY", false),
                Map.entry("THING_GROUP_MEMBERSHIP", false),
                Map.entry("THING_TYPE", false),
                Map.entry("THING_TYPE_ASSOCIATION", false)
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:iot:EventConfigurations
    properties:
      eventConfigurations:
        CA_CERTIFICATE: false
        CERTIFICATE: true
        JOB: false
        JOB_EXECUTION: false
        POLICY: false
        THING: true
        THING_GROUP: false
        THING_GROUP_HIERARCHY: false
        THING_GROUP_MEMBERSHIP: false
        THING_TYPE: false
        THING_TYPE_ASSOCIATION: false
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IoT Event Configurations using the AWS Region. For example:

```sh
 $ pulumi import aws:iot/eventConfigurations:EventConfigurations example us-west-2
```
 