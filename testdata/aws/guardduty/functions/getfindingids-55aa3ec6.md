Data source for managing an AWS GuardDuty Finding Ids.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.guardduty.getFindingIds({
    detectorId: aws_guardduty_detector.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.guardduty.get_finding_ids(detector_id=aws_guardduty_detector["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.GuardDuty.GetFindingIds.Invoke(new()
    {
        DetectorId = aws_guardduty_detector.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/guardduty"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := guardduty.GetFindingIds(ctx, &guardduty.GetFindingIdsArgs{
			DetectorId: aws_guardduty_detector.Example.Id,
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
import com.pulumi.aws.guardduty.GuarddutyFunctions;
import com.pulumi.aws.guardduty.inputs.GetFindingIdsArgs;
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
        final var example = GuarddutyFunctions.getFindingIds(GetFindingIdsArgs.builder()
            .detectorId(aws_guardduty_detector.example().id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:guardduty:getFindingIds
      Arguments:
        detectorId: ${aws_guardduty_detector.example.id}
```
{{% /example %}}
{{% /examples %}}