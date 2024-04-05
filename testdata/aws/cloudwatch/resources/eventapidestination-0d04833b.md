Provides an EventBridge event API Destination resource.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.cloudwatch.EventApiDestination("test", {
    description: "An API Destination",
    invocationEndpoint: "https://api.destination.com/endpoint",
    httpMethod: "POST",
    invocationRateLimitPerSecond: 20,
    connectionArn: aws_cloudwatch_event_connection.test.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.cloudwatch.EventApiDestination("test",
    description="An API Destination",
    invocation_endpoint="https://api.destination.com/endpoint",
    http_method="POST",
    invocation_rate_limit_per_second=20,
    connection_arn=aws_cloudwatch_event_connection["test"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.CloudWatch.EventApiDestination("test", new()
    {
        Description = "An API Destination",
        InvocationEndpoint = "https://api.destination.com/endpoint",
        HttpMethod = "POST",
        InvocationRateLimitPerSecond = 20,
        ConnectionArn = aws_cloudwatch_event_connection.Test.Arn,
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
		_, err := cloudwatch.NewEventApiDestination(ctx, "test", &cloudwatch.EventApiDestinationArgs{
			Description:                  pulumi.String("An API Destination"),
			InvocationEndpoint:           pulumi.String("https://api.destination.com/endpoint"),
			HttpMethod:                   pulumi.String("POST"),
			InvocationRateLimitPerSecond: pulumi.Int(20),
			ConnectionArn:                pulumi.Any(aws_cloudwatch_event_connection.Test.Arn),
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
import com.pulumi.aws.cloudwatch.EventApiDestination;
import com.pulumi.aws.cloudwatch.EventApiDestinationArgs;
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
        var test = new EventApiDestination("test", EventApiDestinationArgs.builder()        
            .description("An API Destination")
            .invocationEndpoint("https://api.destination.com/endpoint")
            .httpMethod("POST")
            .invocationRateLimitPerSecond(20)
            .connectionArn(aws_cloudwatch_event_connection.test().arn())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:cloudwatch:EventApiDestination
    properties:
      description: An API Destination
      invocationEndpoint: https://api.destination.com/endpoint
      httpMethod: POST
      invocationRateLimitPerSecond: 20
      connectionArn: ${aws_cloudwatch_event_connection.test.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge API Destinations using the `name`. For example:

```sh
 $ pulumi import aws:cloudwatch/eventApiDestination:EventApiDestination test api-destination
```
 