Use this data source to retrieve information about a CloudFront distribution.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.cloudfront.getDistribution({
    id: "EDFDVBD632BHDS5",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.cloudfront.get_distribution(id="EDFDVBD632BHDS5")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.CloudFront.GetDistribution.Invoke(new()
    {
        Id = "EDFDVBD632BHDS5",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.LookupDistribution(ctx, &cloudfront.LookupDistributionArgs{
			Id: "EDFDVBD632BHDS5",
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
import com.pulumi.aws.cloudfront.CloudfrontFunctions;
import com.pulumi.aws.cloudfront.inputs.GetDistributionArgs;
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
        final var test = CloudfrontFunctions.getDistribution(GetDistributionArgs.builder()
            .id("EDFDVBD632BHDS5")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:cloudfront:getDistribution
      Arguments:
        id: EDFDVBD632BHDS5
```
{{% /example %}}
{{% /examples %}}