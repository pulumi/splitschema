Provides details about a specific DataPipeline Pipeline.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.datapipeline.getPipeline({
    pipelineId: "pipelineID",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datapipeline.get_pipeline(pipeline_id="pipelineID")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.DataPipeline.GetPipeline.Invoke(new()
    {
        PipelineId = "pipelineID",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/datapipeline"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datapipeline.LookupPipeline(ctx, &datapipeline.LookupPipelineArgs{
			PipelineId: "pipelineID",
		}, nil)
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
import com.pulumi.aws.datapipeline.DatapipelineFunctions;
import com.pulumi.aws.datapipeline.inputs.GetPipelineArgs;
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
        final var example = DatapipelineFunctions.getPipeline(GetPipelineArgs.builder()
            .pipelineId("pipelineID")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:datapipeline:getPipeline
      Arguments:
        pipelineId: pipelineID
```
{{% /example %}}
{{% /examples %}}