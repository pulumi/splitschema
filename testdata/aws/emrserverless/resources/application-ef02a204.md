Manages an EMR Serverless Application.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.emrserverless.Application("example", {
    releaseLabel: "emr-6.6.0",
    type: "hive",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.emrserverless.Application("example",
    release_label="emr-6.6.0",
    type="hive")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.EmrServerless.Application("example", new()
    {
        ReleaseLabel = "emr-6.6.0",
        Type = "hive",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emrserverless"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emrserverless.NewApplication(ctx, "example", &emrserverless.ApplicationArgs{
			ReleaseLabel: pulumi.String("emr-6.6.0"),
			Type:         pulumi.String("hive"),
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
import com.pulumi.aws.emrserverless.Application;
import com.pulumi.aws.emrserverless.ApplicationArgs;
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
        var example = new Application("example", ApplicationArgs.builder()        
            .releaseLabel("emr-6.6.0")
            .type("hive")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:emrserverless:Application
    properties:
      releaseLabel: emr-6.6.0
      type: hive
```
{{% /example %}}
{{% example %}}
### Initial Capacity Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.emrserverless.Application("example", {
    initialCapacities: [{
        initialCapacityConfig: {
            workerConfiguration: {
                cpu: "2 vCPU",
                memory: "10 GB",
            },
            workerCount: 1,
        },
        initialCapacityType: "HiveDriver",
    }],
    releaseLabel: "emr-6.6.0",
    type: "hive",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.emrserverless.Application("example",
    initial_capacities=[aws.emrserverless.ApplicationInitialCapacityArgs(
        initial_capacity_config=aws.emrserverless.ApplicationInitialCapacityInitialCapacityConfigArgs(
            worker_configuration=aws.emrserverless.ApplicationInitialCapacityInitialCapacityConfigWorkerConfigurationArgs(
                cpu="2 vCPU",
                memory="10 GB",
            ),
            worker_count=1,
        ),
        initial_capacity_type="HiveDriver",
    )],
    release_label="emr-6.6.0",
    type="hive")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.EmrServerless.Application("example", new()
    {
        InitialCapacities = new[]
        {
            new Aws.EmrServerless.Inputs.ApplicationInitialCapacityArgs
            {
                InitialCapacityConfig = new Aws.EmrServerless.Inputs.ApplicationInitialCapacityInitialCapacityConfigArgs
                {
                    WorkerConfiguration = new Aws.EmrServerless.Inputs.ApplicationInitialCapacityInitialCapacityConfigWorkerConfigurationArgs
                    {
                        Cpu = "2 vCPU",
                        Memory = "10 GB",
                    },
                    WorkerCount = 1,
                },
                InitialCapacityType = "HiveDriver",
            },
        },
        ReleaseLabel = "emr-6.6.0",
        Type = "hive",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emrserverless"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emrserverless.NewApplication(ctx, "example", &emrserverless.ApplicationArgs{
			InitialCapacities: emrserverless.ApplicationInitialCapacityArray{
				&emrserverless.ApplicationInitialCapacityArgs{
					InitialCapacityConfig: &emrserverless.ApplicationInitialCapacityInitialCapacityConfigArgs{
						WorkerConfiguration: &emrserverless.ApplicationInitialCapacityInitialCapacityConfigWorkerConfigurationArgs{
							Cpu:    pulumi.String("2 vCPU"),
							Memory: pulumi.String("10 GB"),
						},
						WorkerCount: pulumi.Int(1),
					},
					InitialCapacityType: pulumi.String("HiveDriver"),
				},
			},
			ReleaseLabel: pulumi.String("emr-6.6.0"),
			Type:         pulumi.String("hive"),
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
import com.pulumi.aws.emrserverless.Application;
import com.pulumi.aws.emrserverless.ApplicationArgs;
import com.pulumi.aws.emrserverless.inputs.ApplicationInitialCapacityArgs;
import com.pulumi.aws.emrserverless.inputs.ApplicationInitialCapacityInitialCapacityConfigArgs;
import com.pulumi.aws.emrserverless.inputs.ApplicationInitialCapacityInitialCapacityConfigWorkerConfigurationArgs;
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
        var example = new Application("example", ApplicationArgs.builder()        
            .initialCapacities(ApplicationInitialCapacityArgs.builder()
                .initialCapacityConfig(ApplicationInitialCapacityInitialCapacityConfigArgs.builder()
                    .workerConfiguration(ApplicationInitialCapacityInitialCapacityConfigWorkerConfigurationArgs.builder()
                        .cpu("2 vCPU")
                        .memory("10 GB")
                        .build())
                    .workerCount(1)
                    .build())
                .initialCapacityType("HiveDriver")
                .build())
            .releaseLabel("emr-6.6.0")
            .type("hive")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:emrserverless:Application
    properties:
      initialCapacities:
        - initialCapacityConfig:
            workerConfiguration:
              cpu: 2 vCPU
              memory: 10 GB
            workerCount: 1
          initialCapacityType: HiveDriver
      releaseLabel: emr-6.6.0
      type: hive
```
{{% /example %}}
{{% example %}}
### Maximum Capacity Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.emrserverless.Application("example", {
    maximumCapacity: {
        cpu: "2 vCPU",
        memory: "10 GB",
    },
    releaseLabel: "emr-6.6.0",
    type: "hive",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.emrserverless.Application("example",
    maximum_capacity=aws.emrserverless.ApplicationMaximumCapacityArgs(
        cpu="2 vCPU",
        memory="10 GB",
    ),
    release_label="emr-6.6.0",
    type="hive")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.EmrServerless.Application("example", new()
    {
        MaximumCapacity = new Aws.EmrServerless.Inputs.ApplicationMaximumCapacityArgs
        {
            Cpu = "2 vCPU",
            Memory = "10 GB",
        },
        ReleaseLabel = "emr-6.6.0",
        Type = "hive",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emrserverless"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emrserverless.NewApplication(ctx, "example", &emrserverless.ApplicationArgs{
			MaximumCapacity: &emrserverless.ApplicationMaximumCapacityArgs{
				Cpu:    pulumi.String("2 vCPU"),
				Memory: pulumi.String("10 GB"),
			},
			ReleaseLabel: pulumi.String("emr-6.6.0"),
			Type:         pulumi.String("hive"),
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
import com.pulumi.aws.emrserverless.Application;
import com.pulumi.aws.emrserverless.ApplicationArgs;
import com.pulumi.aws.emrserverless.inputs.ApplicationMaximumCapacityArgs;
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
        var example = new Application("example", ApplicationArgs.builder()        
            .maximumCapacity(ApplicationMaximumCapacityArgs.builder()
                .cpu("2 vCPU")
                .memory("10 GB")
                .build())
            .releaseLabel("emr-6.6.0")
            .type("hive")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:emrserverless:Application
    properties:
      maximumCapacity:
        cpu: 2 vCPU
        memory: 10 GB
      releaseLabel: emr-6.6.0
      type: hive
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EMR Severless applications using the `id`. For example:

```sh
 $ pulumi import aws:emrserverless/application:Application example id
```
 