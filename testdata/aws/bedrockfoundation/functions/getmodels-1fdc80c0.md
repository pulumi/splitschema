Data source for managing AWS Bedrock Foundation Models.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.bedrockfoundation.getModels({});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.bedrockfoundation.get_models()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.BedrockFoundation.GetModels.Invoke();

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
		_, err := bedrockfoundation.GetModels(ctx, nil, nil)
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
        final var test = BedrockfoundationFunctions.getModels();

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:bedrockfoundation:getModels
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### Filter by Inference Type

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.bedrockfoundation.getModels({
    byInferenceType: "ON_DEMAND",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.bedrockfoundation.get_models(by_inference_type="ON_DEMAND")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.BedrockFoundation.GetModels.Invoke(new()
    {
        ByInferenceType = "ON_DEMAND",
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
		_, err := bedrockfoundation.GetModels(ctx, &bedrockfoundation.GetModelsArgs{
			ByInferenceType: pulumi.StringRef("ON_DEMAND"),
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
        final var test = BedrockfoundationFunctions.getModels(GetModelsArgs.builder()
            .byInferenceType("ON_DEMAND")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:bedrockfoundation:getModels
      Arguments:
        byInferenceType: ON_DEMAND
```
{{% /example %}}
{{% /examples %}}