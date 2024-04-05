Provides a CodeDeploy CustomActionType

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codepipeline.CustomActionType("example", {
    category: "Build",
    inputArtifactDetails: {
        maximumCount: 1,
        minimumCount: 0,
    },
    outputArtifactDetails: {
        maximumCount: 1,
        minimumCount: 0,
    },
    providerName: "example",
    version: "1",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codepipeline.CustomActionType("example",
    category="Build",
    input_artifact_details=aws.codepipeline.CustomActionTypeInputArtifactDetailsArgs(
        maximum_count=1,
        minimum_count=0,
    ),
    output_artifact_details=aws.codepipeline.CustomActionTypeOutputArtifactDetailsArgs(
        maximum_count=1,
        minimum_count=0,
    ),
    provider_name="example",
    version="1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodePipeline.CustomActionType("example", new()
    {
        Category = "Build",
        InputArtifactDetails = new Aws.CodePipeline.Inputs.CustomActionTypeInputArtifactDetailsArgs
        {
            MaximumCount = 1,
            MinimumCount = 0,
        },
        OutputArtifactDetails = new Aws.CodePipeline.Inputs.CustomActionTypeOutputArtifactDetailsArgs
        {
            MaximumCount = 1,
            MinimumCount = 0,
        },
        ProviderName = "example",
        Version = "1",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codepipeline"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codepipeline.NewCustomActionType(ctx, "example", &codepipeline.CustomActionTypeArgs{
			Category: pulumi.String("Build"),
			InputArtifactDetails: &codepipeline.CustomActionTypeInputArtifactDetailsArgs{
				MaximumCount: pulumi.Int(1),
				MinimumCount: pulumi.Int(0),
			},
			OutputArtifactDetails: &codepipeline.CustomActionTypeOutputArtifactDetailsArgs{
				MaximumCount: pulumi.Int(1),
				MinimumCount: pulumi.Int(0),
			},
			ProviderName: pulumi.String("example"),
			Version:      pulumi.String("1"),
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
import com.pulumi.aws.codepipeline.CustomActionType;
import com.pulumi.aws.codepipeline.CustomActionTypeArgs;
import com.pulumi.aws.codepipeline.inputs.CustomActionTypeInputArtifactDetailsArgs;
import com.pulumi.aws.codepipeline.inputs.CustomActionTypeOutputArtifactDetailsArgs;
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
        var example = new CustomActionType("example", CustomActionTypeArgs.builder()        
            .category("Build")
            .inputArtifactDetails(CustomActionTypeInputArtifactDetailsArgs.builder()
                .maximumCount(1)
                .minimumCount(0)
                .build())
            .outputArtifactDetails(CustomActionTypeOutputArtifactDetailsArgs.builder()
                .maximumCount(1)
                .minimumCount(0)
                .build())
            .providerName("example")
            .version("1")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codepipeline:CustomActionType
    properties:
      category: Build
      inputArtifactDetails:
        maximumCount: 1
        minimumCount: 0
      outputArtifactDetails:
        maximumCount: 1
        minimumCount: 0
      providerName: example
      version: '1'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeDeploy CustomActionType using the `id`. For example:

```sh
 $ pulumi import aws:codepipeline/customActionType:CustomActionType example Build:pulumi:1
```
 