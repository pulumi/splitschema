Manages a Glue Trigger resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Conditional Trigger

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Trigger("example", {
    type: "CONDITIONAL",
    actions: [{
        jobName: aws_glue_job.example1.name,
    }],
    predicate: {
        conditions: [{
            jobName: aws_glue_job.example2.name,
            state: "SUCCEEDED",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Trigger("example",
    type="CONDITIONAL",
    actions=[aws.glue.TriggerActionArgs(
        job_name=aws_glue_job["example1"]["name"],
    )],
    predicate=aws.glue.TriggerPredicateArgs(
        conditions=[aws.glue.TriggerPredicateConditionArgs(
            job_name=aws_glue_job["example2"]["name"],
            state="SUCCEEDED",
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
    var example = new Aws.Glue.Trigger("example", new()
    {
        Type = "CONDITIONAL",
        Actions = new[]
        {
            new Aws.Glue.Inputs.TriggerActionArgs
            {
                JobName = aws_glue_job.Example1.Name,
            },
        },
        Predicate = new Aws.Glue.Inputs.TriggerPredicateArgs
        {
            Conditions = new[]
            {
                new Aws.Glue.Inputs.TriggerPredicateConditionArgs
                {
                    JobName = aws_glue_job.Example2.Name,
                    State = "SUCCEEDED",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
			Type: pulumi.String("CONDITIONAL"),
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					JobName: pulumi.Any(aws_glue_job.Example1.Name),
				},
			},
			Predicate: &glue.TriggerPredicateArgs{
				Conditions: glue.TriggerPredicateConditionArray{
					&glue.TriggerPredicateConditionArgs{
						JobName: pulumi.Any(aws_glue_job.Example2.Name),
						State:   pulumi.String("SUCCEEDED"),
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
import com.pulumi.aws.glue.Trigger;
import com.pulumi.aws.glue.TriggerArgs;
import com.pulumi.aws.glue.inputs.TriggerActionArgs;
import com.pulumi.aws.glue.inputs.TriggerPredicateArgs;
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
        var example = new Trigger("example", TriggerArgs.builder()        
            .type("CONDITIONAL")
            .actions(TriggerActionArgs.builder()
                .jobName(aws_glue_job.example1().name())
                .build())
            .predicate(TriggerPredicateArgs.builder()
                .conditions(TriggerPredicateConditionArgs.builder()
                    .jobName(aws_glue_job.example2().name())
                    .state("SUCCEEDED")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Trigger
    properties:
      type: CONDITIONAL
      actions:
        - jobName: ${aws_glue_job.example1.name}
      predicate:
        conditions:
          - jobName: ${aws_glue_job.example2.name}
            state: SUCCEEDED
```
{{% /example %}}
{{% example %}}
### On-Demand Trigger

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Trigger("example", {
    type: "ON_DEMAND",
    actions: [{
        jobName: aws_glue_job.example.name,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Trigger("example",
    type="ON_DEMAND",
    actions=[aws.glue.TriggerActionArgs(
        job_name=aws_glue_job["example"]["name"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Trigger("example", new()
    {
        Type = "ON_DEMAND",
        Actions = new[]
        {
            new Aws.Glue.Inputs.TriggerActionArgs
            {
                JobName = aws_glue_job.Example.Name,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
			Type: pulumi.String("ON_DEMAND"),
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					JobName: pulumi.Any(aws_glue_job.Example.Name),
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
import com.pulumi.aws.glue.Trigger;
import com.pulumi.aws.glue.TriggerArgs;
import com.pulumi.aws.glue.inputs.TriggerActionArgs;
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
        var example = new Trigger("example", TriggerArgs.builder()        
            .type("ON_DEMAND")
            .actions(TriggerActionArgs.builder()
                .jobName(aws_glue_job.example().name())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Trigger
    properties:
      type: ON_DEMAND
      actions:
        - jobName: ${aws_glue_job.example.name}
```
{{% /example %}}
{{% example %}}
### Scheduled Trigger

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Trigger("example", {
    schedule: "cron(15 12 * * ? *)",
    type: "SCHEDULED",
    actions: [{
        jobName: aws_glue_job.example.name,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Trigger("example",
    schedule="cron(15 12 * * ? *)",
    type="SCHEDULED",
    actions=[aws.glue.TriggerActionArgs(
        job_name=aws_glue_job["example"]["name"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Trigger("example", new()
    {
        Schedule = "cron(15 12 * * ? *)",
        Type = "SCHEDULED",
        Actions = new[]
        {
            new Aws.Glue.Inputs.TriggerActionArgs
            {
                JobName = aws_glue_job.Example.Name,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
			Schedule: pulumi.String("cron(15 12 * * ? *)"),
			Type:     pulumi.String("SCHEDULED"),
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					JobName: pulumi.Any(aws_glue_job.Example.Name),
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
import com.pulumi.aws.glue.Trigger;
import com.pulumi.aws.glue.TriggerArgs;
import com.pulumi.aws.glue.inputs.TriggerActionArgs;
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
        var example = new Trigger("example", TriggerArgs.builder()        
            .schedule("cron(15 12 * * ? *)")
            .type("SCHEDULED")
            .actions(TriggerActionArgs.builder()
                .jobName(aws_glue_job.example().name())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Trigger
    properties:
      schedule: cron(15 12 * * ? *)
      type: SCHEDULED
      actions:
        - jobName: ${aws_glue_job.example.name}
```
{{% /example %}}
{{% example %}}
### Conditional Trigger with Crawler Action

**Note:** Triggers can have both a crawler action and a crawler condition, just no example provided.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Trigger("example", {
    type: "CONDITIONAL",
    actions: [{
        crawlerName: aws_glue_crawler.example1.name,
    }],
    predicate: {
        conditions: [{
            jobName: aws_glue_job.example2.name,
            state: "SUCCEEDED",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Trigger("example",
    type="CONDITIONAL",
    actions=[aws.glue.TriggerActionArgs(
        crawler_name=aws_glue_crawler["example1"]["name"],
    )],
    predicate=aws.glue.TriggerPredicateArgs(
        conditions=[aws.glue.TriggerPredicateConditionArgs(
            job_name=aws_glue_job["example2"]["name"],
            state="SUCCEEDED",
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
    var example = new Aws.Glue.Trigger("example", new()
    {
        Type = "CONDITIONAL",
        Actions = new[]
        {
            new Aws.Glue.Inputs.TriggerActionArgs
            {
                CrawlerName = aws_glue_crawler.Example1.Name,
            },
        },
        Predicate = new Aws.Glue.Inputs.TriggerPredicateArgs
        {
            Conditions = new[]
            {
                new Aws.Glue.Inputs.TriggerPredicateConditionArgs
                {
                    JobName = aws_glue_job.Example2.Name,
                    State = "SUCCEEDED",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
			Type: pulumi.String("CONDITIONAL"),
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					CrawlerName: pulumi.Any(aws_glue_crawler.Example1.Name),
				},
			},
			Predicate: &glue.TriggerPredicateArgs{
				Conditions: glue.TriggerPredicateConditionArray{
					&glue.TriggerPredicateConditionArgs{
						JobName: pulumi.Any(aws_glue_job.Example2.Name),
						State:   pulumi.String("SUCCEEDED"),
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
import com.pulumi.aws.glue.Trigger;
import com.pulumi.aws.glue.TriggerArgs;
import com.pulumi.aws.glue.inputs.TriggerActionArgs;
import com.pulumi.aws.glue.inputs.TriggerPredicateArgs;
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
        var example = new Trigger("example", TriggerArgs.builder()        
            .type("CONDITIONAL")
            .actions(TriggerActionArgs.builder()
                .crawlerName(aws_glue_crawler.example1().name())
                .build())
            .predicate(TriggerPredicateArgs.builder()
                .conditions(TriggerPredicateConditionArgs.builder()
                    .jobName(aws_glue_job.example2().name())
                    .state("SUCCEEDED")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Trigger
    properties:
      type: CONDITIONAL
      actions:
        - crawlerName: ${aws_glue_crawler.example1.name}
      predicate:
        conditions:
          - jobName: ${aws_glue_job.example2.name}
            state: SUCCEEDED
```
{{% /example %}}
{{% example %}}
### Conditional Trigger with Crawler Condition

**Note:** Triggers can have both a crawler action and a crawler condition, just no example provided.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Trigger("example", {
    type: "CONDITIONAL",
    actions: [{
        jobName: aws_glue_job.example1.name,
    }],
    predicate: {
        conditions: [{
            crawlerName: aws_glue_crawler.example2.name,
            crawlState: "SUCCEEDED",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Trigger("example",
    type="CONDITIONAL",
    actions=[aws.glue.TriggerActionArgs(
        job_name=aws_glue_job["example1"]["name"],
    )],
    predicate=aws.glue.TriggerPredicateArgs(
        conditions=[aws.glue.TriggerPredicateConditionArgs(
            crawler_name=aws_glue_crawler["example2"]["name"],
            crawl_state="SUCCEEDED",
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
    var example = new Aws.Glue.Trigger("example", new()
    {
        Type = "CONDITIONAL",
        Actions = new[]
        {
            new Aws.Glue.Inputs.TriggerActionArgs
            {
                JobName = aws_glue_job.Example1.Name,
            },
        },
        Predicate = new Aws.Glue.Inputs.TriggerPredicateArgs
        {
            Conditions = new[]
            {
                new Aws.Glue.Inputs.TriggerPredicateConditionArgs
                {
                    CrawlerName = aws_glue_crawler.Example2.Name,
                    CrawlState = "SUCCEEDED",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
			Type: pulumi.String("CONDITIONAL"),
			Actions: glue.TriggerActionArray{
				&glue.TriggerActionArgs{
					JobName: pulumi.Any(aws_glue_job.Example1.Name),
				},
			},
			Predicate: &glue.TriggerPredicateArgs{
				Conditions: glue.TriggerPredicateConditionArray{
					&glue.TriggerPredicateConditionArgs{
						CrawlerName: pulumi.Any(aws_glue_crawler.Example2.Name),
						CrawlState:  pulumi.String("SUCCEEDED"),
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
import com.pulumi.aws.glue.Trigger;
import com.pulumi.aws.glue.TriggerArgs;
import com.pulumi.aws.glue.inputs.TriggerActionArgs;
import com.pulumi.aws.glue.inputs.TriggerPredicateArgs;
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
        var example = new Trigger("example", TriggerArgs.builder()        
            .type("CONDITIONAL")
            .actions(TriggerActionArgs.builder()
                .jobName(aws_glue_job.example1().name())
                .build())
            .predicate(TriggerPredicateArgs.builder()
                .conditions(TriggerPredicateConditionArgs.builder()
                    .crawlerName(aws_glue_crawler.example2().name())
                    .crawlState("SUCCEEDED")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Trigger
    properties:
      type: CONDITIONAL
      actions:
        - jobName: ${aws_glue_job.example1.name}
      predicate:
        conditions:
          - crawlerName: ${aws_glue_crawler.example2.name}
            crawlState: SUCCEEDED
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Triggers using `name`. For example:

```sh
 $ pulumi import aws:glue/trigger:Trigger MyTrigger MyTrigger
```
 