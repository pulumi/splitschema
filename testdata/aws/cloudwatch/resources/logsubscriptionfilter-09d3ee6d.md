Provides a CloudWatch Logs subscription filter resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testLambdafunctionLogfilter = new aws.cloudwatch.LogSubscriptionFilter("testLambdafunctionLogfilter", {
    roleArn: aws_iam_role.iam_for_lambda.arn,
    logGroup: "/aws/lambda/example_lambda_name",
    filterPattern: "logtype test",
    destinationArn: aws_kinesis_stream.test_logstream.arn,
    distribution: "Random",
});
```
```python
import pulumi
import pulumi_aws as aws

test_lambdafunction_logfilter = aws.cloudwatch.LogSubscriptionFilter("testLambdafunctionLogfilter",
    role_arn=aws_iam_role["iam_for_lambda"]["arn"],
    log_group="/aws/lambda/example_lambda_name",
    filter_pattern="logtype test",
    destination_arn=aws_kinesis_stream["test_logstream"]["arn"],
    distribution="Random")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testLambdafunctionLogfilter = new Aws.CloudWatch.LogSubscriptionFilter("testLambdafunctionLogfilter", new()
    {
        RoleArn = aws_iam_role.Iam_for_lambda.Arn,
        LogGroup = "/aws/lambda/example_lambda_name",
        FilterPattern = "logtype test",
        DestinationArn = aws_kinesis_stream.Test_logstream.Arn,
        Distribution = "Random",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewLogSubscriptionFilter(ctx, "testLambdafunctionLogfilter", &cloudwatch.LogSubscriptionFilterArgs{
			RoleArn:        pulumi.Any(aws_iam_role.Iam_for_lambda.Arn),
			LogGroup:       pulumi.Any("/aws/lambda/example_lambda_name"),
			FilterPattern:  pulumi.String("logtype test"),
			DestinationArn: pulumi.Any(aws_kinesis_stream.Test_logstream.Arn),
			Distribution:   pulumi.String("Random"),
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
import com.pulumi.aws.cloudwatch.LogSubscriptionFilter;
import com.pulumi.aws.cloudwatch.LogSubscriptionFilterArgs;
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
        var testLambdafunctionLogfilter = new LogSubscriptionFilter("testLambdafunctionLogfilter", LogSubscriptionFilterArgs.builder()        
            .roleArn(aws_iam_role.iam_for_lambda().arn())
            .logGroup("/aws/lambda/example_lambda_name")
            .filterPattern("logtype test")
            .destinationArn(aws_kinesis_stream.test_logstream().arn())
            .distribution("Random")
            .build());

    }
}
```
```yaml
resources:
  testLambdafunctionLogfilter:
    type: aws:cloudwatch:LogSubscriptionFilter
    properties:
      roleArn: ${aws_iam_role.iam_for_lambda.arn}
      logGroup: /aws/lambda/example_lambda_name
      filterPattern: logtype test
      destinationArn: ${aws_kinesis_stream.test_logstream.arn}
      distribution: Random
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch Logs subscription filter using the log group name and subscription filter name separated by `|`. For example:

```sh
 $ pulumi import aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter test_lambdafunction_logfilter /aws/lambda/example_lambda_name|test_lambdafunction_logfilter
```
 