Data source for managing an AWS Bedrock Foundation Model.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testModels = aws.bedrockfoundation.getModels({});
const testModel = testModels.then(testModels => aws.bedrockfoundation.getModel({
    modelId: testModels.modelSummaries?.[0]?.modelId,
}));
```
```python
import pulumi
import pulumi_aws as aws

test_models = aws.bedrockfoundation.get_models()
test_model = aws.bedrockfoundation.get_model(model_id=test_models.model_summaries[0].model_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testModels = Aws.BedrockFoundation.GetModels.Invoke();

    var testModel = Aws.BedrockFoundation.GetModel.Invoke(new()
    {
        ModelId = testModels.Apply(getModelsResult => getModelsResult.ModelSummaries[0]?.ModelId),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/bedrockfoundation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testModels, err := bedrockfoundation.GetModels(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = bedrockfoundation.GetModel(ctx, &bedrockfoundation.GetModelArgs{
			ModelId: testModels.ModelSummaries[0].ModelId,
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
import com.pulumi.aws.bedrockfoundation.BedrockfoundationFunctions;
import com.pulumi.aws.bedrockfoundation.inputs.GetModelsArgs;
import com.pulumi.aws.bedrockfoundation.inputs.GetModelArgs;
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
        final var testModels = BedrockfoundationFunctions.getModels();

        final var testModel = BedrockfoundationFunctions.getModel(GetModelArgs.builder()
            .modelId(testModels.applyValue(getModelsResult -> getModelsResult.modelSummaries()[0].modelId()))
            .build());

    }
}
```
```yaml
variables:
  testModels:
    fn::invoke:
      Function: aws:bedrockfoundation:getModels
      Arguments: {}
  testModel:
    fn::invoke:
      Function: aws:bedrockfoundation:getModel
      Arguments:
        modelId: ${testModels.modelSummaries[0].modelId}
```
{{% /example %}}
{{% /examples %}}