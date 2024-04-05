Provides a SageMaker Pipeline resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.Pipeline("example", {
    pipelineName: "example",
    pipelineDisplayName: "example",
    roleArn: aws_iam_role.example.arn,
    pipelineDefinition: JSON.stringify({
        Version: "2020-12-01",
        Steps: [{
            Name: "Test",
            Type: "Fail",
            Arguments: {
                ErrorMessage: "test",
            },
        }],
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.sagemaker.Pipeline("example",
    pipeline_name="example",
    pipeline_display_name="example",
    role_arn=aws_iam_role["example"]["arn"],
    pipeline_definition=json.dumps({
        "Version": "2020-12-01",
        "Steps": [{
            "Name": "Test",
            "Type": "Fail",
            "Arguments": {
                "ErrorMessage": "test",
            },
        }],
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.Pipeline("example", new()
    {
        PipelineName = "example",
        PipelineDisplayName = "example",
        RoleArn = aws_iam_role.Example.Arn,
        PipelineDefinition = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2020-12-01",
            ["Steps"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Name"] = "Test",
                    ["Type"] = "Fail",
                    ["Arguments"] = new Dictionary<string, object?>
                    {
                        ["ErrorMessage"] = "test",
                    },
                },
            },
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2020-12-01",
			"Steps": []map[string]interface{}{
				map[string]interface{}{
					"Name": "Test",
					"Type": "Fail",
					"Arguments": map[string]interface{}{
						"ErrorMessage": "test",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = sagemaker.NewPipeline(ctx, "example", &sagemaker.PipelineArgs{
			PipelineName:        pulumi.String("example"),
			PipelineDisplayName: pulumi.String("example"),
			RoleArn:             pulumi.Any(aws_iam_role.Example.Arn),
			PipelineDefinition:  pulumi.String(json0),
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
import com.pulumi.aws.sagemaker.Pipeline;
import com.pulumi.aws.sagemaker.PipelineArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var example = new Pipeline("example", PipelineArgs.builder()        
            .pipelineName("example")
            .pipelineDisplayName("example")
            .roleArn(aws_iam_role.example().arn())
            .pipelineDefinition(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2020-12-01"),
                    jsonProperty("Steps", jsonArray(jsonObject(
                        jsonProperty("Name", "Test"),
                        jsonProperty("Type", "Fail"),
                        jsonProperty("Arguments", jsonObject(
                            jsonProperty("ErrorMessage", "test")
                        ))
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:Pipeline
    properties:
      pipelineName: example
      pipelineDisplayName: example
      roleArn: ${aws_iam_role.example.arn}
      pipelineDefinition:
        fn::toJSON:
          Version: 2020-12-01
          Steps:
            - Name: Test
              Type: Fail
              Arguments:
                ErrorMessage: test
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import pipelines using the `pipeline_name`. For example:

```sh
 $ pulumi import aws:sagemaker/pipeline:Pipeline test_pipeline pipeline
```
 