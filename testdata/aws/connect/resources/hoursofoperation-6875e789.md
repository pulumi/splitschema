Provides an Amazon Connect Hours of Operation resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.connect.HoursOfOperation("test", {
    configs: [
        {
            day: "MONDAY",
            endTime: {
                hours: 23,
                minutes: 8,
            },
            startTime: {
                hours: 8,
                minutes: 0,
            },
        },
        {
            day: "TUESDAY",
            endTime: {
                hours: 21,
                minutes: 0,
            },
            startTime: {
                hours: 9,
                minutes: 0,
            },
        },
    ],
    description: "Monday office hours",
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    tags: {
        Name: "Example Hours of Operation",
    },
    timeZone: "EST",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.connect.HoursOfOperation("test",
    configs=[
        aws.connect.HoursOfOperationConfigArgs(
            day="MONDAY",
            end_time=aws.connect.HoursOfOperationConfigEndTimeArgs(
                hours=23,
                minutes=8,
            ),
            start_time=aws.connect.HoursOfOperationConfigStartTimeArgs(
                hours=8,
                minutes=0,
            ),
        ),
        aws.connect.HoursOfOperationConfigArgs(
            day="TUESDAY",
            end_time=aws.connect.HoursOfOperationConfigEndTimeArgs(
                hours=21,
                minutes=0,
            ),
            start_time=aws.connect.HoursOfOperationConfigStartTimeArgs(
                hours=9,
                minutes=0,
            ),
        ),
    ],
    description="Monday office hours",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    tags={
        "Name": "Example Hours of Operation",
    },
    time_zone="EST")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Connect.HoursOfOperation("test", new()
    {
        Configs = new[]
        {
            new Aws.Connect.Inputs.HoursOfOperationConfigArgs
            {
                Day = "MONDAY",
                EndTime = new Aws.Connect.Inputs.HoursOfOperationConfigEndTimeArgs
                {
                    Hours = 23,
                    Minutes = 8,
                },
                StartTime = new Aws.Connect.Inputs.HoursOfOperationConfigStartTimeArgs
                {
                    Hours = 8,
                    Minutes = 0,
                },
            },
            new Aws.Connect.Inputs.HoursOfOperationConfigArgs
            {
                Day = "TUESDAY",
                EndTime = new Aws.Connect.Inputs.HoursOfOperationConfigEndTimeArgs
                {
                    Hours = 21,
                    Minutes = 0,
                },
                StartTime = new Aws.Connect.Inputs.HoursOfOperationConfigStartTimeArgs
                {
                    Hours = 9,
                    Minutes = 0,
                },
            },
        },
        Description = "Monday office hours",
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        Tags = 
        {
            { "Name", "Example Hours of Operation" },
        },
        TimeZone = "EST",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewHoursOfOperation(ctx, "test", &connect.HoursOfOperationArgs{
			Configs: connect.HoursOfOperationConfigArray{
				&connect.HoursOfOperationConfigArgs{
					Day: pulumi.String("MONDAY"),
					EndTime: &connect.HoursOfOperationConfigEndTimeArgs{
						Hours:   pulumi.Int(23),
						Minutes: pulumi.Int(8),
					},
					StartTime: &connect.HoursOfOperationConfigStartTimeArgs{
						Hours:   pulumi.Int(8),
						Minutes: pulumi.Int(0),
					},
				},
				&connect.HoursOfOperationConfigArgs{
					Day: pulumi.String("TUESDAY"),
					EndTime: &connect.HoursOfOperationConfigEndTimeArgs{
						Hours:   pulumi.Int(21),
						Minutes: pulumi.Int(0),
					},
					StartTime: &connect.HoursOfOperationConfigStartTimeArgs{
						Hours:   pulumi.Int(9),
						Minutes: pulumi.Int(0),
					},
				},
			},
			Description: pulumi.String("Monday office hours"),
			InstanceId:  pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Hours of Operation"),
			},
			TimeZone: pulumi.String("EST"),
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
import com.pulumi.aws.connect.HoursOfOperation;
import com.pulumi.aws.connect.HoursOfOperationArgs;
import com.pulumi.aws.connect.inputs.HoursOfOperationConfigArgs;
import com.pulumi.aws.connect.inputs.HoursOfOperationConfigEndTimeArgs;
import com.pulumi.aws.connect.inputs.HoursOfOperationConfigStartTimeArgs;
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
        var test = new HoursOfOperation("test", HoursOfOperationArgs.builder()        
            .configs(            
                HoursOfOperationConfigArgs.builder()
                    .day("MONDAY")
                    .endTime(HoursOfOperationConfigEndTimeArgs.builder()
                        .hours(23)
                        .minutes(8)
                        .build())
                    .startTime(HoursOfOperationConfigStartTimeArgs.builder()
                        .hours(8)
                        .minutes(0)
                        .build())
                    .build(),
                HoursOfOperationConfigArgs.builder()
                    .day("TUESDAY")
                    .endTime(HoursOfOperationConfigEndTimeArgs.builder()
                        .hours(21)
                        .minutes(0)
                        .build())
                    .startTime(HoursOfOperationConfigStartTimeArgs.builder()
                        .hours(9)
                        .minutes(0)
                        .build())
                    .build())
            .description("Monday office hours")
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .tags(Map.of("Name", "Example Hours of Operation"))
            .timeZone("EST")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:connect:HoursOfOperation
    properties:
      configs:
        - day: MONDAY
          endTime:
            hours: 23
            minutes: 8
          startTime:
            hours: 8
            minutes: 0
        - day: TUESDAY
          endTime:
            hours: 21
            minutes: 0
          startTime:
            hours: 9
            minutes: 0
      description: Monday office hours
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
      tags:
        Name: Example Hours of Operation
      timeZone: EST
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect Hours of Operations using the `instance_id` and `hours_of_operation_id` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:connect/hoursOfOperation:HoursOfOperation example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
```
 