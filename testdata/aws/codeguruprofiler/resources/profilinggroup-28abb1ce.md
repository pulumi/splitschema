Resource for managing an AWS CodeGuru Profiler Profiling Group.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codeguruprofiler.ProfilingGroup("example", {
    agentOrchestrationConfig: {
        profilingEnabled: true,
    },
    computePlatform: "Default",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codeguruprofiler.ProfilingGroup("example",
    agent_orchestration_config=aws.codeguruprofiler.ProfilingGroupAgentOrchestrationConfigArgs(
        profiling_enabled=True,
    ),
    compute_platform="Default")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeGuruProfiler.ProfilingGroup("example", new()
    {
        AgentOrchestrationConfig = new Aws.CodeGuruProfiler.Inputs.ProfilingGroupAgentOrchestrationConfigArgs
        {
            ProfilingEnabled = true,
        },
        ComputePlatform = "Default",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codeguruprofiler"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codeguruprofiler.NewProfilingGroup(ctx, "example", &codeguruprofiler.ProfilingGroupArgs{
			AgentOrchestrationConfig: &codeguruprofiler.ProfilingGroupAgentOrchestrationConfigArgs{
				ProfilingEnabled: pulumi.Bool(true),
			},
			ComputePlatform: pulumi.String("Default"),
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
import com.pulumi.aws.codeguruprofiler.ProfilingGroup;
import com.pulumi.aws.codeguruprofiler.ProfilingGroupArgs;
import com.pulumi.aws.codeguruprofiler.inputs.ProfilingGroupAgentOrchestrationConfigArgs;
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
        var example = new ProfilingGroup("example", ProfilingGroupArgs.builder()        
            .agentOrchestrationConfig(ProfilingGroupAgentOrchestrationConfigArgs.builder()
                .profilingEnabled(true)
                .build())
            .computePlatform("Default")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codeguruprofiler:ProfilingGroup
    properties:
      agentOrchestrationConfig:
        profilingEnabled: true
      computePlatform: Default
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeGuru Profiler Profiling Group using the `id`. For example:

```sh
 $ pulumi import aws:codeguruprofiler/profilingGroup:ProfilingGroup example profiling_group-name-12345678
```
 