`aws.ec2.NetworkInsightsAnalysis` provides details about a specific Network Insights Analysis.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getNetworkInsightsAnalysis({
    networkInsightsAnalysisId: aws_ec2_network_insights_analysis.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_network_insights_analysis(network_insights_analysis_id=aws_ec2_network_insights_analysis["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetNetworkInsightsAnalysis.Invoke(new()
    {
        NetworkInsightsAnalysisId = aws_ec2_network_insights_analysis.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.LookupNetworkInsightsAnalysis(ctx, &ec2.LookupNetworkInsightsAnalysisArgs{
			NetworkInsightsAnalysisId: pulumi.StringRef(aws_ec2_network_insights_analysis.Example.Id),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetNetworkInsightsAnalysisArgs;
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
        final var example = Ec2Functions.getNetworkInsightsAnalysis(GetNetworkInsightsAnalysisArgs.builder()
            .networkInsightsAnalysisId(aws_ec2_network_insights_analysis.example().id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getNetworkInsightsAnalysis
      Arguments:
        networkInsightsAnalysisId: ${aws_ec2_network_insights_analysis.example.id}
```
{{% /example %}}
{{% /examples %}}