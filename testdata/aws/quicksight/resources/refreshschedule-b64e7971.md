Resource for managing a QuickSight Refresh Schedule.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.RefreshSchedule("example", {
    dataSetId: "dataset-id",
    schedule: {
        refreshType: "FULL_REFRESH",
        scheduleFrequency: {
            interval: "HOURLY",
        },
    },
    scheduleId: "schedule-id",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.RefreshSchedule("example",
    data_set_id="dataset-id",
    schedule=aws.quicksight.RefreshScheduleScheduleArgs(
        refresh_type="FULL_REFRESH",
        schedule_frequency=aws.quicksight.RefreshScheduleScheduleScheduleFrequencyArgs(
            interval="HOURLY",
        ),
    ),
    schedule_id="schedule-id")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.RefreshSchedule("example", new()
    {
        DataSetId = "dataset-id",
        Schedule = new Aws.Quicksight.Inputs.RefreshScheduleScheduleArgs
        {
            RefreshType = "FULL_REFRESH",
            ScheduleFrequency = new Aws.Quicksight.Inputs.RefreshScheduleScheduleScheduleFrequencyArgs
            {
                Interval = "HOURLY",
            },
        },
        ScheduleId = "schedule-id",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewRefreshSchedule(ctx, "example", &quicksight.RefreshScheduleArgs{
			DataSetId: pulumi.String("dataset-id"),
			Schedule: &quicksight.RefreshScheduleScheduleArgs{
				RefreshType: pulumi.String("FULL_REFRESH"),
				ScheduleFrequency: &quicksight.RefreshScheduleScheduleScheduleFrequencyArgs{
					Interval: pulumi.String("HOURLY"),
				},
			},
			ScheduleId: pulumi.String("schedule-id"),
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
import com.pulumi.aws.quicksight.RefreshSchedule;
import com.pulumi.aws.quicksight.RefreshScheduleArgs;
import com.pulumi.aws.quicksight.inputs.RefreshScheduleScheduleArgs;
import com.pulumi.aws.quicksight.inputs.RefreshScheduleScheduleScheduleFrequencyArgs;
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
        var example = new RefreshSchedule("example", RefreshScheduleArgs.builder()        
            .dataSetId("dataset-id")
            .schedule(RefreshScheduleScheduleArgs.builder()
                .refreshType("FULL_REFRESH")
                .scheduleFrequency(RefreshScheduleScheduleScheduleFrequencyArgs.builder()
                    .interval("HOURLY")
                    .build())
                .build())
            .scheduleId("schedule-id")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:RefreshSchedule
    properties:
      dataSetId: dataset-id
      schedule:
        refreshType: FULL_REFRESH
        scheduleFrequency:
          interval: HOURLY
      scheduleId: schedule-id
```
{{% /example %}}
{{% example %}}
### With Weekly Refresh

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.RefreshSchedule("example", {
    dataSetId: "dataset-id",
    schedule: {
        refreshType: "INCREMENTAL_REFRESH",
        scheduleFrequency: {
            interval: "WEEKLY",
            refreshOnDay: {
                dayOfWeek: "MONDAY",
            },
            timeOfTheDay: "01:00",
            timezone: "Europe/London",
        },
    },
    scheduleId: "schedule-id",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.RefreshSchedule("example",
    data_set_id="dataset-id",
    schedule=aws.quicksight.RefreshScheduleScheduleArgs(
        refresh_type="INCREMENTAL_REFRESH",
        schedule_frequency=aws.quicksight.RefreshScheduleScheduleScheduleFrequencyArgs(
            interval="WEEKLY",
            refresh_on_day=aws.quicksight.RefreshScheduleScheduleScheduleFrequencyRefreshOnDayArgs(
                day_of_week="MONDAY",
            ),
            time_of_the_day="01:00",
            timezone="Europe/London",
        ),
    ),
    schedule_id="schedule-id")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.RefreshSchedule("example", new()
    {
        DataSetId = "dataset-id",
        Schedule = new Aws.Quicksight.Inputs.RefreshScheduleScheduleArgs
        {
            RefreshType = "INCREMENTAL_REFRESH",
            ScheduleFrequency = new Aws.Quicksight.Inputs.RefreshScheduleScheduleScheduleFrequencyArgs
            {
                Interval = "WEEKLY",
                RefreshOnDay = new Aws.Quicksight.Inputs.RefreshScheduleScheduleScheduleFrequencyRefreshOnDayArgs
                {
                    DayOfWeek = "MONDAY",
                },
                TimeOfTheDay = "01:00",
                Timezone = "Europe/London",
            },
        },
        ScheduleId = "schedule-id",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewRefreshSchedule(ctx, "example", &quicksight.RefreshScheduleArgs{
			DataSetId: pulumi.String("dataset-id"),
			Schedule: &quicksight.RefreshScheduleScheduleArgs{
				RefreshType: pulumi.String("INCREMENTAL_REFRESH"),
				ScheduleFrequency: &quicksight.RefreshScheduleScheduleScheduleFrequencyArgs{
					Interval: pulumi.String("WEEKLY"),
					RefreshOnDay: &quicksight.RefreshScheduleScheduleScheduleFrequencyRefreshOnDayArgs{
						DayOfWeek: pulumi.String("MONDAY"),
					},
					TimeOfTheDay: pulumi.String("01:00"),
					Timezone:     pulumi.String("Europe/London"),
				},
			},
			ScheduleId: pulumi.String("schedule-id"),
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
import com.pulumi.aws.quicksight.RefreshSchedule;
import com.pulumi.aws.quicksight.RefreshScheduleArgs;
import com.pulumi.aws.quicksight.inputs.RefreshScheduleScheduleArgs;
import com.pulumi.aws.quicksight.inputs.RefreshScheduleScheduleScheduleFrequencyArgs;
import com.pulumi.aws.quicksight.inputs.RefreshScheduleScheduleScheduleFrequencyRefreshOnDayArgs;
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
        var example = new RefreshSchedule("example", RefreshScheduleArgs.builder()        
            .dataSetId("dataset-id")
            .schedule(RefreshScheduleScheduleArgs.builder()
                .refreshType("INCREMENTAL_REFRESH")
                .scheduleFrequency(RefreshScheduleScheduleScheduleFrequencyArgs.builder()
                    .interval("WEEKLY")
                    .refreshOnDay(RefreshScheduleScheduleScheduleFrequencyRefreshOnDayArgs.builder()
                        .dayOfWeek("MONDAY")
                        .build())
                    .timeOfTheDay("01:00")
                    .timezone("Europe/London")
                    .build())
                .build())
            .scheduleId("schedule-id")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:RefreshSchedule
    properties:
      dataSetId: dataset-id
      schedule:
        refreshType: INCREMENTAL_REFRESH
        scheduleFrequency:
          interval: WEEKLY
          refreshOnDay:
            dayOfWeek: MONDAY
          timeOfTheDay: 01:00
          timezone: Europe/London
      scheduleId: schedule-id
```
{{% /example %}}
{{% example %}}
### With Monthly Refresh

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.RefreshSchedule("example", {
    dataSetId: "dataset-id",
    schedule: {
        refreshType: "INCREMENTAL_REFRESH",
        scheduleFrequency: {
            interval: "MONTHLY",
            refreshOnDay: {
                dayOfMonth: "1",
            },
            timeOfTheDay: "01:00",
            timezone: "Europe/London",
        },
    },
    scheduleId: "schedule-id",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.RefreshSchedule("example",
    data_set_id="dataset-id",
    schedule=aws.quicksight.RefreshScheduleScheduleArgs(
        refresh_type="INCREMENTAL_REFRESH",
        schedule_frequency=aws.quicksight.RefreshScheduleScheduleScheduleFrequencyArgs(
            interval="MONTHLY",
            refresh_on_day=aws.quicksight.RefreshScheduleScheduleScheduleFrequencyRefreshOnDayArgs(
                day_of_month="1",
            ),
            time_of_the_day="01:00",
            timezone="Europe/London",
        ),
    ),
    schedule_id="schedule-id")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.RefreshSchedule("example", new()
    {
        DataSetId = "dataset-id",
        Schedule = new Aws.Quicksight.Inputs.RefreshScheduleScheduleArgs
        {
            RefreshType = "INCREMENTAL_REFRESH",
            ScheduleFrequency = new Aws.Quicksight.Inputs.RefreshScheduleScheduleScheduleFrequencyArgs
            {
                Interval = "MONTHLY",
                RefreshOnDay = new Aws.Quicksight.Inputs.RefreshScheduleScheduleScheduleFrequencyRefreshOnDayArgs
                {
                    DayOfMonth = "1",
                },
                TimeOfTheDay = "01:00",
                Timezone = "Europe/London",
            },
        },
        ScheduleId = "schedule-id",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewRefreshSchedule(ctx, "example", &quicksight.RefreshScheduleArgs{
			DataSetId: pulumi.String("dataset-id"),
			Schedule: &quicksight.RefreshScheduleScheduleArgs{
				RefreshType: pulumi.String("INCREMENTAL_REFRESH"),
				ScheduleFrequency: &quicksight.RefreshScheduleScheduleScheduleFrequencyArgs{
					Interval: pulumi.String("MONTHLY"),
					RefreshOnDay: &quicksight.RefreshScheduleScheduleScheduleFrequencyRefreshOnDayArgs{
						DayOfMonth: pulumi.String("1"),
					},
					TimeOfTheDay: pulumi.String("01:00"),
					Timezone:     pulumi.String("Europe/London"),
				},
			},
			ScheduleId: pulumi.String("schedule-id"),
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
import com.pulumi.aws.quicksight.RefreshSchedule;
import com.pulumi.aws.quicksight.RefreshScheduleArgs;
import com.pulumi.aws.quicksight.inputs.RefreshScheduleScheduleArgs;
import com.pulumi.aws.quicksight.inputs.RefreshScheduleScheduleScheduleFrequencyArgs;
import com.pulumi.aws.quicksight.inputs.RefreshScheduleScheduleScheduleFrequencyRefreshOnDayArgs;
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
        var example = new RefreshSchedule("example", RefreshScheduleArgs.builder()        
            .dataSetId("dataset-id")
            .schedule(RefreshScheduleScheduleArgs.builder()
                .refreshType("INCREMENTAL_REFRESH")
                .scheduleFrequency(RefreshScheduleScheduleScheduleFrequencyArgs.builder()
                    .interval("MONTHLY")
                    .refreshOnDay(RefreshScheduleScheduleScheduleFrequencyRefreshOnDayArgs.builder()
                        .dayOfMonth("1")
                        .build())
                    .timeOfTheDay("01:00")
                    .timezone("Europe/London")
                    .build())
                .build())
            .scheduleId("schedule-id")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:RefreshSchedule
    properties:
      dataSetId: dataset-id
      schedule:
        refreshType: INCREMENTAL_REFRESH
        scheduleFrequency:
          interval: MONTHLY
          refreshOnDay:
            dayOfMonth: '1'
          timeOfTheDay: 01:00
          timezone: Europe/London
      scheduleId: schedule-id
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a QuickSight Refresh Schedule using the AWS account ID, data set ID and schedule ID separated by commas (`,`). For example:

```sh
 $ pulumi import aws:quicksight/refreshSchedule:RefreshSchedule example 123456789012,dataset-id,schedule-id
```
 