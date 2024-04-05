Data source for managing an AWS SFN (Step Functions) State Machine Versions.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.sfn.getStateMachineVersions({
    statemachineArn: aws_sfn_state_machine.test.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.sfn.get_state_machine_versions(statemachine_arn=aws_sfn_state_machine["test"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Sfn.GetStateMachineVersions.Invoke(new()
    {
        StatemachineArn = aws_sfn_state_machine.Test.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sfn"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sfn.GetStateMachineVersions(ctx, &sfn.GetStateMachineVersionsArgs{
			StatemachineArn: aws_sfn_state_machine.Test.Arn,
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
import com.pulumi.aws.sfn.SfnFunctions;
import com.pulumi.aws.sfn.inputs.GetStateMachineVersionsArgs;
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
        final var test = SfnFunctions.getStateMachineVersions(GetStateMachineVersionsArgs.builder()
            .statemachineArn(aws_sfn_state_machine.test().arn())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:sfn:getStateMachineVersions
      Arguments:
        statemachineArn: ${aws_sfn_state_machine.test.arn}
```
{{% /example %}}
{{% /examples %}}