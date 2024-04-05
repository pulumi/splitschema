Provides a CloudWatch Evidently Launch resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Launch("example", {
    project: aws_evidently_project.example.name,
    groups: [{
        feature: aws_evidently_feature.example.name,
        name: "Variation1",
        variation: "Variation1",
    }],
    scheduledSplitsConfig: {
        steps: [{
            groupWeights: {
                Variation1: 0,
            },
            startTime: "2024-01-07 01:43:59+00:00",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Launch("example",
    project=aws_evidently_project["example"]["name"],
    groups=[aws.evidently.LaunchGroupArgs(
        feature=aws_evidently_feature["example"]["name"],
        name="Variation1",
        variation="Variation1",
    )],
    scheduled_splits_config=aws.evidently.LaunchScheduledSplitsConfigArgs(
        steps=[aws.evidently.LaunchScheduledSplitsConfigStepArgs(
            group_weights={
                "Variation1": 0,
            },
            start_time="2024-01-07 01:43:59+00:00",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Launch("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        Groups = new[]
        {
            new Aws.Evidently.Inputs.LaunchGroupArgs
            {
                Feature = aws_evidently_feature.Example.Name,
                Name = "Variation1",
                Variation = "Variation1",
            },
        },
        ScheduledSplitsConfig = new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigArgs
        {
            Steps = new[]
            {
                new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigStepArgs
                {
                    GroupWeights = 
                    {
                        { "Variation1", 0 },
                    },
                    StartTime = "2024-01-07 01:43:59+00:00",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
			Project: pulumi.Any(aws_evidently_project.Example.Name),
			Groups: evidently.LaunchGroupArray{
				&evidently.LaunchGroupArgs{
					Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
					Name:      pulumi.String("Variation1"),
					Variation: pulumi.String("Variation1"),
				},
			},
			ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
				Steps: evidently.LaunchScheduledSplitsConfigStepArray{
					&evidently.LaunchScheduledSplitsConfigStepArgs{
						GroupWeights: pulumi.IntMap{
							"Variation1": pulumi.Int(0),
						},
						StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
					},
				},
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
import com.pulumi.aws.evidently.Launch;
import com.pulumi.aws.evidently.LaunchArgs;
import com.pulumi.aws.evidently.inputs.LaunchGroupArgs;
import com.pulumi.aws.evidently.inputs.LaunchScheduledSplitsConfigArgs;
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
        var example = new Launch("example", LaunchArgs.builder()        
            .project(aws_evidently_project.example().name())
            .groups(LaunchGroupArgs.builder()
                .feature(aws_evidently_feature.example().name())
                .name("Variation1")
                .variation("Variation1")
                .build())
            .scheduledSplitsConfig(LaunchScheduledSplitsConfigArgs.builder()
                .steps(LaunchScheduledSplitsConfigStepArgs.builder()
                    .groupWeights(Map.of("Variation1", 0))
                    .startTime("2024-01-07 01:43:59+00:00")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Launch
    properties:
      project: ${aws_evidently_project.example.name}
      groups:
        - feature: ${aws_evidently_feature.example.name}
          name: Variation1
          variation: Variation1
      scheduledSplitsConfig:
        steps:
          - groupWeights:
              Variation1: 0
            startTime: 2024-01-07 01:43:59+00:00
```
{{% /example %}}
{{% example %}}
### With description

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Launch("example", {
    project: aws_evidently_project.example.name,
    description: "example description",
    groups: [{
        feature: aws_evidently_feature.example.name,
        name: "Variation1",
        variation: "Variation1",
    }],
    scheduledSplitsConfig: {
        steps: [{
            groupWeights: {
                Variation1: 0,
            },
            startTime: "2024-01-07 01:43:59+00:00",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Launch("example",
    project=aws_evidently_project["example"]["name"],
    description="example description",
    groups=[aws.evidently.LaunchGroupArgs(
        feature=aws_evidently_feature["example"]["name"],
        name="Variation1",
        variation="Variation1",
    )],
    scheduled_splits_config=aws.evidently.LaunchScheduledSplitsConfigArgs(
        steps=[aws.evidently.LaunchScheduledSplitsConfigStepArgs(
            group_weights={
                "Variation1": 0,
            },
            start_time="2024-01-07 01:43:59+00:00",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Launch("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        Description = "example description",
        Groups = new[]
        {
            new Aws.Evidently.Inputs.LaunchGroupArgs
            {
                Feature = aws_evidently_feature.Example.Name,
                Name = "Variation1",
                Variation = "Variation1",
            },
        },
        ScheduledSplitsConfig = new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigArgs
        {
            Steps = new[]
            {
                new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigStepArgs
                {
                    GroupWeights = 
                    {
                        { "Variation1", 0 },
                    },
                    StartTime = "2024-01-07 01:43:59+00:00",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
			Project:     pulumi.Any(aws_evidently_project.Example.Name),
			Description: pulumi.String("example description"),
			Groups: evidently.LaunchGroupArray{
				&evidently.LaunchGroupArgs{
					Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
					Name:      pulumi.String("Variation1"),
					Variation: pulumi.String("Variation1"),
				},
			},
			ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
				Steps: evidently.LaunchScheduledSplitsConfigStepArray{
					&evidently.LaunchScheduledSplitsConfigStepArgs{
						GroupWeights: pulumi.IntMap{
							"Variation1": pulumi.Int(0),
						},
						StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
					},
				},
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
import com.pulumi.aws.evidently.Launch;
import com.pulumi.aws.evidently.LaunchArgs;
import com.pulumi.aws.evidently.inputs.LaunchGroupArgs;
import com.pulumi.aws.evidently.inputs.LaunchScheduledSplitsConfigArgs;
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
        var example = new Launch("example", LaunchArgs.builder()        
            .project(aws_evidently_project.example().name())
            .description("example description")
            .groups(LaunchGroupArgs.builder()
                .feature(aws_evidently_feature.example().name())
                .name("Variation1")
                .variation("Variation1")
                .build())
            .scheduledSplitsConfig(LaunchScheduledSplitsConfigArgs.builder()
                .steps(LaunchScheduledSplitsConfigStepArgs.builder()
                    .groupWeights(Map.of("Variation1", 0))
                    .startTime("2024-01-07 01:43:59+00:00")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Launch
    properties:
      project: ${aws_evidently_project.example.name}
      description: example description
      groups:
        - feature: ${aws_evidently_feature.example.name}
          name: Variation1
          variation: Variation1
      scheduledSplitsConfig:
        steps:
          - groupWeights:
              Variation1: 0
            startTime: 2024-01-07 01:43:59+00:00
```
{{% /example %}}
{{% example %}}
### With multiple groups

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Launch("example", {
    project: aws_evidently_project.example.name,
    groups: [
        {
            feature: aws_evidently_feature.example.name,
            name: "Variation1",
            variation: "Variation1",
            description: "first-group",
        },
        {
            feature: aws_evidently_feature.example.name,
            name: "Variation2",
            variation: "Variation2",
            description: "second-group",
        },
    ],
    scheduledSplitsConfig: {
        steps: [{
            groupWeights: {
                Variation1: 0,
                Variation2: 0,
            },
            startTime: "2024-01-07 01:43:59+00:00",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Launch("example",
    project=aws_evidently_project["example"]["name"],
    groups=[
        aws.evidently.LaunchGroupArgs(
            feature=aws_evidently_feature["example"]["name"],
            name="Variation1",
            variation="Variation1",
            description="first-group",
        ),
        aws.evidently.LaunchGroupArgs(
            feature=aws_evidently_feature["example"]["name"],
            name="Variation2",
            variation="Variation2",
            description="second-group",
        ),
    ],
    scheduled_splits_config=aws.evidently.LaunchScheduledSplitsConfigArgs(
        steps=[aws.evidently.LaunchScheduledSplitsConfigStepArgs(
            group_weights={
                "Variation1": 0,
                "Variation2": 0,
            },
            start_time="2024-01-07 01:43:59+00:00",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Launch("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        Groups = new[]
        {
            new Aws.Evidently.Inputs.LaunchGroupArgs
            {
                Feature = aws_evidently_feature.Example.Name,
                Name = "Variation1",
                Variation = "Variation1",
                Description = "first-group",
            },
            new Aws.Evidently.Inputs.LaunchGroupArgs
            {
                Feature = aws_evidently_feature.Example.Name,
                Name = "Variation2",
                Variation = "Variation2",
                Description = "second-group",
            },
        },
        ScheduledSplitsConfig = new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigArgs
        {
            Steps = new[]
            {
                new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigStepArgs
                {
                    GroupWeights = 
                    {
                        { "Variation1", 0 },
                        { "Variation2", 0 },
                    },
                    StartTime = "2024-01-07 01:43:59+00:00",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
			Project: pulumi.Any(aws_evidently_project.Example.Name),
			Groups: evidently.LaunchGroupArray{
				&evidently.LaunchGroupArgs{
					Feature:     pulumi.Any(aws_evidently_feature.Example.Name),
					Name:        pulumi.String("Variation1"),
					Variation:   pulumi.String("Variation1"),
					Description: pulumi.String("first-group"),
				},
				&evidently.LaunchGroupArgs{
					Feature:     pulumi.Any(aws_evidently_feature.Example.Name),
					Name:        pulumi.String("Variation2"),
					Variation:   pulumi.String("Variation2"),
					Description: pulumi.String("second-group"),
				},
			},
			ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
				Steps: evidently.LaunchScheduledSplitsConfigStepArray{
					&evidently.LaunchScheduledSplitsConfigStepArgs{
						GroupWeights: pulumi.IntMap{
							"Variation1": pulumi.Int(0),
							"Variation2": pulumi.Int(0),
						},
						StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
					},
				},
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
import com.pulumi.aws.evidently.Launch;
import com.pulumi.aws.evidently.LaunchArgs;
import com.pulumi.aws.evidently.inputs.LaunchGroupArgs;
import com.pulumi.aws.evidently.inputs.LaunchScheduledSplitsConfigArgs;
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
        var example = new Launch("example", LaunchArgs.builder()        
            .project(aws_evidently_project.example().name())
            .groups(            
                LaunchGroupArgs.builder()
                    .feature(aws_evidently_feature.example().name())
                    .name("Variation1")
                    .variation("Variation1")
                    .description("first-group")
                    .build(),
                LaunchGroupArgs.builder()
                    .feature(aws_evidently_feature.example().name())
                    .name("Variation2")
                    .variation("Variation2")
                    .description("second-group")
                    .build())
            .scheduledSplitsConfig(LaunchScheduledSplitsConfigArgs.builder()
                .steps(LaunchScheduledSplitsConfigStepArgs.builder()
                    .groupWeights(Map.ofEntries(
                        Map.entry("Variation1", 0),
                        Map.entry("Variation2", 0)
                    ))
                    .startTime("2024-01-07 01:43:59+00:00")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Launch
    properties:
      project: ${aws_evidently_project.example.name}
      groups:
        - feature: ${aws_evidently_feature.example.name}
          name: Variation1
          variation: Variation1
          description: first-group
        - feature: ${aws_evidently_feature.example.name}
          name: Variation2
          variation: Variation2
          description: second-group
      scheduledSplitsConfig:
        steps:
          - groupWeights:
              Variation1: 0
              Variation2: 0
            startTime: 2024-01-07 01:43:59+00:00
```
{{% /example %}}
{{% example %}}
### With metric_monitors

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Launch("example", {
    project: aws_evidently_project.example.name,
    groups: [{
        feature: aws_evidently_feature.example.name,
        name: "Variation1",
        variation: "Variation1",
    }],
    metricMonitors: [
        {
            metricDefinition: {
                entityIdKey: "entity_id_key1",
                eventPattern: "{\"Price\":[{\"numeric\":[\">\",11,\"<=\",22]}]}",
                name: "name1",
                unitLabel: "unit_label1",
                valueKey: "value_key1",
            },
        },
        {
            metricDefinition: {
                entityIdKey: "entity_id_key2",
                eventPattern: "{\"Price\":[{\"numeric\":[\">\",9,\"<=\",19]}]}",
                name: "name2",
                unitLabel: "unit_label2",
                valueKey: "value_key2",
            },
        },
    ],
    scheduledSplitsConfig: {
        steps: [{
            groupWeights: {
                Variation1: 0,
            },
            startTime: "2024-01-07 01:43:59+00:00",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Launch("example",
    project=aws_evidently_project["example"]["name"],
    groups=[aws.evidently.LaunchGroupArgs(
        feature=aws_evidently_feature["example"]["name"],
        name="Variation1",
        variation="Variation1",
    )],
    metric_monitors=[
        aws.evidently.LaunchMetricMonitorArgs(
            metric_definition=aws.evidently.LaunchMetricMonitorMetricDefinitionArgs(
                entity_id_key="entity_id_key1",
                event_pattern="{\"Price\":[{\"numeric\":[\">\",11,\"<=\",22]}]}",
                name="name1",
                unit_label="unit_label1",
                value_key="value_key1",
            ),
        ),
        aws.evidently.LaunchMetricMonitorArgs(
            metric_definition=aws.evidently.LaunchMetricMonitorMetricDefinitionArgs(
                entity_id_key="entity_id_key2",
                event_pattern="{\"Price\":[{\"numeric\":[\">\",9,\"<=\",19]}]}",
                name="name2",
                unit_label="unit_label2",
                value_key="value_key2",
            ),
        ),
    ],
    scheduled_splits_config=aws.evidently.LaunchScheduledSplitsConfigArgs(
        steps=[aws.evidently.LaunchScheduledSplitsConfigStepArgs(
            group_weights={
                "Variation1": 0,
            },
            start_time="2024-01-07 01:43:59+00:00",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Launch("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        Groups = new[]
        {
            new Aws.Evidently.Inputs.LaunchGroupArgs
            {
                Feature = aws_evidently_feature.Example.Name,
                Name = "Variation1",
                Variation = "Variation1",
            },
        },
        MetricMonitors = new[]
        {
            new Aws.Evidently.Inputs.LaunchMetricMonitorArgs
            {
                MetricDefinition = new Aws.Evidently.Inputs.LaunchMetricMonitorMetricDefinitionArgs
                {
                    EntityIdKey = "entity_id_key1",
                    EventPattern = "{\"Price\":[{\"numeric\":[\">\",11,\"<=\",22]}]}",
                    Name = "name1",
                    UnitLabel = "unit_label1",
                    ValueKey = "value_key1",
                },
            },
            new Aws.Evidently.Inputs.LaunchMetricMonitorArgs
            {
                MetricDefinition = new Aws.Evidently.Inputs.LaunchMetricMonitorMetricDefinitionArgs
                {
                    EntityIdKey = "entity_id_key2",
                    EventPattern = "{\"Price\":[{\"numeric\":[\">\",9,\"<=\",19]}]}",
                    Name = "name2",
                    UnitLabel = "unit_label2",
                    ValueKey = "value_key2",
                },
            },
        },
        ScheduledSplitsConfig = new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigArgs
        {
            Steps = new[]
            {
                new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigStepArgs
                {
                    GroupWeights = 
                    {
                        { "Variation1", 0 },
                    },
                    StartTime = "2024-01-07 01:43:59+00:00",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
			Project: pulumi.Any(aws_evidently_project.Example.Name),
			Groups: evidently.LaunchGroupArray{
				&evidently.LaunchGroupArgs{
					Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
					Name:      pulumi.String("Variation1"),
					Variation: pulumi.String("Variation1"),
				},
			},
			MetricMonitors: evidently.LaunchMetricMonitorArray{
				&evidently.LaunchMetricMonitorArgs{
					MetricDefinition: &evidently.LaunchMetricMonitorMetricDefinitionArgs{
						EntityIdKey:  pulumi.String("entity_id_key1"),
						EventPattern: pulumi.String("{\"Price\":[{\"numeric\":[\">\",11,\"<=\",22]}]}"),
						Name:         pulumi.String("name1"),
						UnitLabel:    pulumi.String("unit_label1"),
						ValueKey:     pulumi.String("value_key1"),
					},
				},
				&evidently.LaunchMetricMonitorArgs{
					MetricDefinition: &evidently.LaunchMetricMonitorMetricDefinitionArgs{
						EntityIdKey:  pulumi.String("entity_id_key2"),
						EventPattern: pulumi.String("{\"Price\":[{\"numeric\":[\">\",9,\"<=\",19]}]}"),
						Name:         pulumi.String("name2"),
						UnitLabel:    pulumi.String("unit_label2"),
						ValueKey:     pulumi.String("value_key2"),
					},
				},
			},
			ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
				Steps: evidently.LaunchScheduledSplitsConfigStepArray{
					&evidently.LaunchScheduledSplitsConfigStepArgs{
						GroupWeights: pulumi.IntMap{
							"Variation1": pulumi.Int(0),
						},
						StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
					},
				},
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
import com.pulumi.aws.evidently.Launch;
import com.pulumi.aws.evidently.LaunchArgs;
import com.pulumi.aws.evidently.inputs.LaunchGroupArgs;
import com.pulumi.aws.evidently.inputs.LaunchMetricMonitorArgs;
import com.pulumi.aws.evidently.inputs.LaunchMetricMonitorMetricDefinitionArgs;
import com.pulumi.aws.evidently.inputs.LaunchScheduledSplitsConfigArgs;
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
        var example = new Launch("example", LaunchArgs.builder()        
            .project(aws_evidently_project.example().name())
            .groups(LaunchGroupArgs.builder()
                .feature(aws_evidently_feature.example().name())
                .name("Variation1")
                .variation("Variation1")
                .build())
            .metricMonitors(            
                LaunchMetricMonitorArgs.builder()
                    .metricDefinition(LaunchMetricMonitorMetricDefinitionArgs.builder()
                        .entityIdKey("entity_id_key1")
                        .eventPattern("{\"Price\":[{\"numeric\":[\">\",11,\"<=\",22]}]}")
                        .name("name1")
                        .unitLabel("unit_label1")
                        .valueKey("value_key1")
                        .build())
                    .build(),
                LaunchMetricMonitorArgs.builder()
                    .metricDefinition(LaunchMetricMonitorMetricDefinitionArgs.builder()
                        .entityIdKey("entity_id_key2")
                        .eventPattern("{\"Price\":[{\"numeric\":[\">\",9,\"<=\",19]}]}")
                        .name("name2")
                        .unitLabel("unit_label2")
                        .valueKey("value_key2")
                        .build())
                    .build())
            .scheduledSplitsConfig(LaunchScheduledSplitsConfigArgs.builder()
                .steps(LaunchScheduledSplitsConfigStepArgs.builder()
                    .groupWeights(Map.of("Variation1", 0))
                    .startTime("2024-01-07 01:43:59+00:00")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Launch
    properties:
      project: ${aws_evidently_project.example.name}
      groups:
        - feature: ${aws_evidently_feature.example.name}
          name: Variation1
          variation: Variation1
      metricMonitors:
        - metricDefinition:
            entityIdKey: entity_id_key1
            eventPattern: '{"Price":[{"numeric":[">",11,"<=",22]}]}'
            name: name1
            unitLabel: unit_label1
            valueKey: value_key1
        - metricDefinition:
            entityIdKey: entity_id_key2
            eventPattern: '{"Price":[{"numeric":[">",9,"<=",19]}]}'
            name: name2
            unitLabel: unit_label2
            valueKey: value_key2
      scheduledSplitsConfig:
        steps:
          - groupWeights:
              Variation1: 0
            startTime: 2024-01-07 01:43:59+00:00
```
{{% /example %}}
{{% example %}}
### With randomization_salt

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Launch("example", {
    project: aws_evidently_project.example.name,
    randomizationSalt: "example randomization salt",
    groups: [{
        feature: aws_evidently_feature.example.name,
        name: "Variation1",
        variation: "Variation1",
    }],
    scheduledSplitsConfig: {
        steps: [{
            groupWeights: {
                Variation1: 0,
            },
            startTime: "2024-01-07 01:43:59+00:00",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Launch("example",
    project=aws_evidently_project["example"]["name"],
    randomization_salt="example randomization salt",
    groups=[aws.evidently.LaunchGroupArgs(
        feature=aws_evidently_feature["example"]["name"],
        name="Variation1",
        variation="Variation1",
    )],
    scheduled_splits_config=aws.evidently.LaunchScheduledSplitsConfigArgs(
        steps=[aws.evidently.LaunchScheduledSplitsConfigStepArgs(
            group_weights={
                "Variation1": 0,
            },
            start_time="2024-01-07 01:43:59+00:00",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Launch("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        RandomizationSalt = "example randomization salt",
        Groups = new[]
        {
            new Aws.Evidently.Inputs.LaunchGroupArgs
            {
                Feature = aws_evidently_feature.Example.Name,
                Name = "Variation1",
                Variation = "Variation1",
            },
        },
        ScheduledSplitsConfig = new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigArgs
        {
            Steps = new[]
            {
                new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigStepArgs
                {
                    GroupWeights = 
                    {
                        { "Variation1", 0 },
                    },
                    StartTime = "2024-01-07 01:43:59+00:00",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
			Project:           pulumi.Any(aws_evidently_project.Example.Name),
			RandomizationSalt: pulumi.String("example randomization salt"),
			Groups: evidently.LaunchGroupArray{
				&evidently.LaunchGroupArgs{
					Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
					Name:      pulumi.String("Variation1"),
					Variation: pulumi.String("Variation1"),
				},
			},
			ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
				Steps: evidently.LaunchScheduledSplitsConfigStepArray{
					&evidently.LaunchScheduledSplitsConfigStepArgs{
						GroupWeights: pulumi.IntMap{
							"Variation1": pulumi.Int(0),
						},
						StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
					},
				},
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
import com.pulumi.aws.evidently.Launch;
import com.pulumi.aws.evidently.LaunchArgs;
import com.pulumi.aws.evidently.inputs.LaunchGroupArgs;
import com.pulumi.aws.evidently.inputs.LaunchScheduledSplitsConfigArgs;
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
        var example = new Launch("example", LaunchArgs.builder()        
            .project(aws_evidently_project.example().name())
            .randomizationSalt("example randomization salt")
            .groups(LaunchGroupArgs.builder()
                .feature(aws_evidently_feature.example().name())
                .name("Variation1")
                .variation("Variation1")
                .build())
            .scheduledSplitsConfig(LaunchScheduledSplitsConfigArgs.builder()
                .steps(LaunchScheduledSplitsConfigStepArgs.builder()
                    .groupWeights(Map.of("Variation1", 0))
                    .startTime("2024-01-07 01:43:59+00:00")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Launch
    properties:
      project: ${aws_evidently_project.example.name}
      randomizationSalt: example randomization salt
      groups:
        - feature: ${aws_evidently_feature.example.name}
          name: Variation1
          variation: Variation1
      scheduledSplitsConfig:
        steps:
          - groupWeights:
              Variation1: 0
            startTime: 2024-01-07 01:43:59+00:00
```
{{% /example %}}
{{% example %}}
### With multiple steps

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Launch("example", {
    project: aws_evidently_project.example.name,
    groups: [
        {
            feature: aws_evidently_feature.example.name,
            name: "Variation1",
            variation: "Variation1",
        },
        {
            feature: aws_evidently_feature.example.name,
            name: "Variation2",
            variation: "Variation2",
        },
    ],
    scheduledSplitsConfig: {
        steps: [
            {
                groupWeights: {
                    Variation1: 15,
                    Variation2: 10,
                },
                startTime: "2024-01-07 01:43:59+00:00",
            },
            {
                groupWeights: {
                    Variation1: 20,
                    Variation2: 25,
                },
                startTime: "2024-01-08 01:43:59+00:00",
            },
        ],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Launch("example",
    project=aws_evidently_project["example"]["name"],
    groups=[
        aws.evidently.LaunchGroupArgs(
            feature=aws_evidently_feature["example"]["name"],
            name="Variation1",
            variation="Variation1",
        ),
        aws.evidently.LaunchGroupArgs(
            feature=aws_evidently_feature["example"]["name"],
            name="Variation2",
            variation="Variation2",
        ),
    ],
    scheduled_splits_config=aws.evidently.LaunchScheduledSplitsConfigArgs(
        steps=[
            aws.evidently.LaunchScheduledSplitsConfigStepArgs(
                group_weights={
                    "Variation1": 15,
                    "Variation2": 10,
                },
                start_time="2024-01-07 01:43:59+00:00",
            ),
            aws.evidently.LaunchScheduledSplitsConfigStepArgs(
                group_weights={
                    "Variation1": 20,
                    "Variation2": 25,
                },
                start_time="2024-01-08 01:43:59+00:00",
            ),
        ],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Launch("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        Groups = new[]
        {
            new Aws.Evidently.Inputs.LaunchGroupArgs
            {
                Feature = aws_evidently_feature.Example.Name,
                Name = "Variation1",
                Variation = "Variation1",
            },
            new Aws.Evidently.Inputs.LaunchGroupArgs
            {
                Feature = aws_evidently_feature.Example.Name,
                Name = "Variation2",
                Variation = "Variation2",
            },
        },
        ScheduledSplitsConfig = new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigArgs
        {
            Steps = new[]
            {
                new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigStepArgs
                {
                    GroupWeights = 
                    {
                        { "Variation1", 15 },
                        { "Variation2", 10 },
                    },
                    StartTime = "2024-01-07 01:43:59+00:00",
                },
                new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigStepArgs
                {
                    GroupWeights = 
                    {
                        { "Variation1", 20 },
                        { "Variation2", 25 },
                    },
                    StartTime = "2024-01-08 01:43:59+00:00",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
			Project: pulumi.Any(aws_evidently_project.Example.Name),
			Groups: evidently.LaunchGroupArray{
				&evidently.LaunchGroupArgs{
					Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
					Name:      pulumi.String("Variation1"),
					Variation: pulumi.String("Variation1"),
				},
				&evidently.LaunchGroupArgs{
					Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
					Name:      pulumi.String("Variation2"),
					Variation: pulumi.String("Variation2"),
				},
			},
			ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
				Steps: evidently.LaunchScheduledSplitsConfigStepArray{
					&evidently.LaunchScheduledSplitsConfigStepArgs{
						GroupWeights: pulumi.IntMap{
							"Variation1": pulumi.Int(15),
							"Variation2": pulumi.Int(10),
						},
						StartTime: pulumi.String("2024-01-07 01:43:59+00:00"),
					},
					&evidently.LaunchScheduledSplitsConfigStepArgs{
						GroupWeights: pulumi.IntMap{
							"Variation1": pulumi.Int(20),
							"Variation2": pulumi.Int(25),
						},
						StartTime: pulumi.String("2024-01-08 01:43:59+00:00"),
					},
				},
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
import com.pulumi.aws.evidently.Launch;
import com.pulumi.aws.evidently.LaunchArgs;
import com.pulumi.aws.evidently.inputs.LaunchGroupArgs;
import com.pulumi.aws.evidently.inputs.LaunchScheduledSplitsConfigArgs;
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
        var example = new Launch("example", LaunchArgs.builder()        
            .project(aws_evidently_project.example().name())
            .groups(            
                LaunchGroupArgs.builder()
                    .feature(aws_evidently_feature.example().name())
                    .name("Variation1")
                    .variation("Variation1")
                    .build(),
                LaunchGroupArgs.builder()
                    .feature(aws_evidently_feature.example().name())
                    .name("Variation2")
                    .variation("Variation2")
                    .build())
            .scheduledSplitsConfig(LaunchScheduledSplitsConfigArgs.builder()
                .steps(                
                    LaunchScheduledSplitsConfigStepArgs.builder()
                        .groupWeights(Map.ofEntries(
                            Map.entry("Variation1", 15),
                            Map.entry("Variation2", 10)
                        ))
                        .startTime("2024-01-07 01:43:59+00:00")
                        .build(),
                    LaunchScheduledSplitsConfigStepArgs.builder()
                        .groupWeights(Map.ofEntries(
                            Map.entry("Variation1", 20),
                            Map.entry("Variation2", 25)
                        ))
                        .startTime("2024-01-08 01:43:59+00:00")
                        .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Launch
    properties:
      project: ${aws_evidently_project.example.name}
      groups:
        - feature: ${aws_evidently_feature.example.name}
          name: Variation1
          variation: Variation1
        - feature: ${aws_evidently_feature.example.name}
          name: Variation2
          variation: Variation2
      scheduledSplitsConfig:
        steps:
          - groupWeights:
              Variation1: 15
              Variation2: 10
            startTime: 2024-01-07 01:43:59+00:00
          - groupWeights:
              Variation1: 20
              Variation2: 25
            startTime: 2024-01-08 01:43:59+00:00
```
{{% /example %}}
{{% example %}}
### With segment overrides

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Launch("example", {
    project: aws_evidently_project.example.name,
    groups: [
        {
            feature: aws_evidently_feature.example.name,
            name: "Variation1",
            variation: "Variation1",
        },
        {
            feature: aws_evidently_feature.example.name,
            name: "Variation2",
            variation: "Variation2",
        },
    ],
    scheduledSplitsConfig: {
        steps: [{
            groupWeights: {
                Variation1: 0,
                Variation2: 0,
            },
            segmentOverrides: [
                {
                    evaluationOrder: 1,
                    segment: aws_evidently_segment.example.name,
                    weights: {
                        Variation2: 10000,
                    },
                },
                {
                    evaluationOrder: 2,
                    segment: aws_evidently_segment.example.name,
                    weights: {
                        Variation1: 40000,
                        Variation2: 30000,
                    },
                },
            ],
            startTime: "2024-01-08 01:43:59+00:00",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Launch("example",
    project=aws_evidently_project["example"]["name"],
    groups=[
        aws.evidently.LaunchGroupArgs(
            feature=aws_evidently_feature["example"]["name"],
            name="Variation1",
            variation="Variation1",
        ),
        aws.evidently.LaunchGroupArgs(
            feature=aws_evidently_feature["example"]["name"],
            name="Variation2",
            variation="Variation2",
        ),
    ],
    scheduled_splits_config=aws.evidently.LaunchScheduledSplitsConfigArgs(
        steps=[aws.evidently.LaunchScheduledSplitsConfigStepArgs(
            group_weights={
                "Variation1": 0,
                "Variation2": 0,
            },
            segment_overrides=[
                aws.evidently.LaunchScheduledSplitsConfigStepSegmentOverrideArgs(
                    evaluation_order=1,
                    segment=aws_evidently_segment["example"]["name"],
                    weights={
                        "Variation2": 10000,
                    },
                ),
                aws.evidently.LaunchScheduledSplitsConfigStepSegmentOverrideArgs(
                    evaluation_order=2,
                    segment=aws_evidently_segment["example"]["name"],
                    weights={
                        "Variation1": 40000,
                        "Variation2": 30000,
                    },
                ),
            ],
            start_time="2024-01-08 01:43:59+00:00",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Launch("example", new()
    {
        Project = aws_evidently_project.Example.Name,
        Groups = new[]
        {
            new Aws.Evidently.Inputs.LaunchGroupArgs
            {
                Feature = aws_evidently_feature.Example.Name,
                Name = "Variation1",
                Variation = "Variation1",
            },
            new Aws.Evidently.Inputs.LaunchGroupArgs
            {
                Feature = aws_evidently_feature.Example.Name,
                Name = "Variation2",
                Variation = "Variation2",
            },
        },
        ScheduledSplitsConfig = new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigArgs
        {
            Steps = new[]
            {
                new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigStepArgs
                {
                    GroupWeights = 
                    {
                        { "Variation1", 0 },
                        { "Variation2", 0 },
                    },
                    SegmentOverrides = new[]
                    {
                        new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigStepSegmentOverrideArgs
                        {
                            EvaluationOrder = 1,
                            Segment = aws_evidently_segment.Example.Name,
                            Weights = 
                            {
                                { "Variation2", 10000 },
                            },
                        },
                        new Aws.Evidently.Inputs.LaunchScheduledSplitsConfigStepSegmentOverrideArgs
                        {
                            EvaluationOrder = 2,
                            Segment = aws_evidently_segment.Example.Name,
                            Weights = 
                            {
                                { "Variation1", 40000 },
                                { "Variation2", 30000 },
                            },
                        },
                    },
                    StartTime = "2024-01-08 01:43:59+00:00",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewLaunch(ctx, "example", &evidently.LaunchArgs{
			Project: pulumi.Any(aws_evidently_project.Example.Name),
			Groups: evidently.LaunchGroupArray{
				&evidently.LaunchGroupArgs{
					Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
					Name:      pulumi.String("Variation1"),
					Variation: pulumi.String("Variation1"),
				},
				&evidently.LaunchGroupArgs{
					Feature:   pulumi.Any(aws_evidently_feature.Example.Name),
					Name:      pulumi.String("Variation2"),
					Variation: pulumi.String("Variation2"),
				},
			},
			ScheduledSplitsConfig: &evidently.LaunchScheduledSplitsConfigArgs{
				Steps: evidently.LaunchScheduledSplitsConfigStepArray{
					&evidently.LaunchScheduledSplitsConfigStepArgs{
						GroupWeights: pulumi.IntMap{
							"Variation1": pulumi.Int(0),
							"Variation2": pulumi.Int(0),
						},
						SegmentOverrides: evidently.LaunchScheduledSplitsConfigStepSegmentOverrideArray{
							&evidently.LaunchScheduledSplitsConfigStepSegmentOverrideArgs{
								EvaluationOrder: pulumi.Int(1),
								Segment:         pulumi.Any(aws_evidently_segment.Example.Name),
								Weights: pulumi.IntMap{
									"Variation2": pulumi.Int(10000),
								},
							},
							&evidently.LaunchScheduledSplitsConfigStepSegmentOverrideArgs{
								EvaluationOrder: pulumi.Int(2),
								Segment:         pulumi.Any(aws_evidently_segment.Example.Name),
								Weights: pulumi.IntMap{
									"Variation1": pulumi.Int(40000),
									"Variation2": pulumi.Int(30000),
								},
							},
						},
						StartTime: pulumi.String("2024-01-08 01:43:59+00:00"),
					},
				},
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
import com.pulumi.aws.evidently.Launch;
import com.pulumi.aws.evidently.LaunchArgs;
import com.pulumi.aws.evidently.inputs.LaunchGroupArgs;
import com.pulumi.aws.evidently.inputs.LaunchScheduledSplitsConfigArgs;
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
        var example = new Launch("example", LaunchArgs.builder()        
            .project(aws_evidently_project.example().name())
            .groups(            
                LaunchGroupArgs.builder()
                    .feature(aws_evidently_feature.example().name())
                    .name("Variation1")
                    .variation("Variation1")
                    .build(),
                LaunchGroupArgs.builder()
                    .feature(aws_evidently_feature.example().name())
                    .name("Variation2")
                    .variation("Variation2")
                    .build())
            .scheduledSplitsConfig(LaunchScheduledSplitsConfigArgs.builder()
                .steps(LaunchScheduledSplitsConfigStepArgs.builder()
                    .groupWeights(Map.ofEntries(
                        Map.entry("Variation1", 0),
                        Map.entry("Variation2", 0)
                    ))
                    .segmentOverrides(                    
                        LaunchScheduledSplitsConfigStepSegmentOverrideArgs.builder()
                            .evaluationOrder(1)
                            .segment(aws_evidently_segment.example().name())
                            .weights(Map.of("Variation2", 10000))
                            .build(),
                        LaunchScheduledSplitsConfigStepSegmentOverrideArgs.builder()
                            .evaluationOrder(2)
                            .segment(aws_evidently_segment.example().name())
                            .weights(Map.ofEntries(
                                Map.entry("Variation1", 40000),
                                Map.entry("Variation2", 30000)
                            ))
                            .build())
                    .startTime("2024-01-08 01:43:59+00:00")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Launch
    properties:
      project: ${aws_evidently_project.example.name}
      groups:
        - feature: ${aws_evidently_feature.example.name}
          name: Variation1
          variation: Variation1
        - feature: ${aws_evidently_feature.example.name}
          name: Variation2
          variation: Variation2
      scheduledSplitsConfig:
        steps:
          - groupWeights:
              Variation1: 0
              Variation2: 0
            segmentOverrides:
              - evaluationOrder: 1
                segment: ${aws_evidently_segment.example.name}
                weights:
                  Variation2: 10000
              - evaluationOrder: 2
                segment: ${aws_evidently_segment.example.name}
                weights:
                  Variation1: 40000
                  Variation2: 30000
            startTime: 2024-01-08 01:43:59+00:00
```
{{% /example %}}
{{% /examples %}}

## Import

Import using the `name` of the launch and `arn` of the project separated by a `:`:

__Using `pulumi import` to import__ CloudWatch Evidently Launch using the `name` of the launch and `name` of the project or `arn` of the hosting CloudWatch Evidently Project separated by a `:`. For example:

Import using the `name` of the launch and `name` of the project separated by a `:`:

```sh
 $ pulumi import aws:evidently/launch:Launch example exampleLaunchName:exampleProjectName
```
 Import using the `name` of the launch and `arn` of the project separated by a `:`:

```sh
 $ pulumi import aws:evidently/launch:Launch example exampleLaunchName:arn:aws:evidently:us-east-1:123456789012:project/exampleProjectName
```
 