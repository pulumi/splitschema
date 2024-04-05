Manages an asynchronous invocation configuration for a Lambda Function or Alias. More information about asynchronous invocations and the configurable values can be found in the [Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Destination Configuration

> **NOTE:** Ensure the Lambda Function IAM Role has necessary permissions for the destination, such as `sqs:SendMessage` or `sns:Publish`, otherwise the API will return a generic `InvalidParameterValueException: The destination ARN arn:PARTITION:SERVICE:REGION:ACCOUNT:RESOURCE is invalid.` error.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.FunctionEventInvokeConfig("example", {
    functionName: aws_lambda_alias.example.function_name,
    destinationConfig: {
        onFailure: {
            destination: aws_sqs_queue.example.arn,
        },
        onSuccess: {
            destination: aws_sns_topic.example.arn,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lambda_.FunctionEventInvokeConfig("example",
    function_name=aws_lambda_alias["example"]["function_name"],
    destination_config=aws.lambda_.FunctionEventInvokeConfigDestinationConfigArgs(
        on_failure=aws.lambda_.FunctionEventInvokeConfigDestinationConfigOnFailureArgs(
            destination=aws_sqs_queue["example"]["arn"],
        ),
        on_success=aws.lambda_.FunctionEventInvokeConfigDestinationConfigOnSuccessArgs(
            destination=aws_sns_topic["example"]["arn"],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.FunctionEventInvokeConfig("example", new()
    {
        FunctionName = aws_lambda_alias.Example.Function_name,
        DestinationConfig = new Aws.Lambda.Inputs.FunctionEventInvokeConfigDestinationConfigArgs
        {
            OnFailure = new Aws.Lambda.Inputs.FunctionEventInvokeConfigDestinationConfigOnFailureArgs
            {
                Destination = aws_sqs_queue.Example.Arn,
            },
            OnSuccess = new Aws.Lambda.Inputs.FunctionEventInvokeConfigDestinationConfigOnSuccessArgs
            {
                Destination = aws_sns_topic.Example.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewFunctionEventInvokeConfig(ctx, "example", &lambda.FunctionEventInvokeConfigArgs{
			FunctionName: pulumi.Any(aws_lambda_alias.Example.Function_name),
			DestinationConfig: &lambda.FunctionEventInvokeConfigDestinationConfigArgs{
				OnFailure: &lambda.FunctionEventInvokeConfigDestinationConfigOnFailureArgs{
					Destination: pulumi.Any(aws_sqs_queue.Example.Arn),
				},
				OnSuccess: &lambda.FunctionEventInvokeConfigDestinationConfigOnSuccessArgs{
					Destination: pulumi.Any(aws_sns_topic.Example.Arn),
				},
			},
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
import com.pulumi.aws.lambda.FunctionEventInvokeConfig;
import com.pulumi.aws.lambda.FunctionEventInvokeConfigArgs;
import com.pulumi.aws.lambda.inputs.FunctionEventInvokeConfigDestinationConfigArgs;
import com.pulumi.aws.lambda.inputs.FunctionEventInvokeConfigDestinationConfigOnFailureArgs;
import com.pulumi.aws.lambda.inputs.FunctionEventInvokeConfigDestinationConfigOnSuccessArgs;
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
        var example = new FunctionEventInvokeConfig("example", FunctionEventInvokeConfigArgs.builder()        
            .functionName(aws_lambda_alias.example().function_name())
            .destinationConfig(FunctionEventInvokeConfigDestinationConfigArgs.builder()
                .onFailure(FunctionEventInvokeConfigDestinationConfigOnFailureArgs.builder()
                    .destination(aws_sqs_queue.example().arn())
                    .build())
                .onSuccess(FunctionEventInvokeConfigDestinationConfigOnSuccessArgs.builder()
                    .destination(aws_sns_topic.example().arn())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:FunctionEventInvokeConfig
    properties:
      functionName: ${aws_lambda_alias.example.function_name}
      destinationConfig:
        onFailure:
          destination: ${aws_sqs_queue.example.arn}
        onSuccess:
          destination: ${aws_sns_topic.example.arn}
```
{{% /example %}}
{{% example %}}
### Error Handling Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.FunctionEventInvokeConfig("example", {
    functionName: aws_lambda_alias.example.function_name,
    maximumEventAgeInSeconds: 60,
    maximumRetryAttempts: 0,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lambda_.FunctionEventInvokeConfig("example",
    function_name=aws_lambda_alias["example"]["function_name"],
    maximum_event_age_in_seconds=60,
    maximum_retry_attempts=0)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.FunctionEventInvokeConfig("example", new()
    {
        FunctionName = aws_lambda_alias.Example.Function_name,
        MaximumEventAgeInSeconds = 60,
        MaximumRetryAttempts = 0,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewFunctionEventInvokeConfig(ctx, "example", &lambda.FunctionEventInvokeConfigArgs{
			FunctionName:             pulumi.Any(aws_lambda_alias.Example.Function_name),
			MaximumEventAgeInSeconds: pulumi.Int(60),
			MaximumRetryAttempts:     pulumi.Int(0),
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
import com.pulumi.aws.lambda.FunctionEventInvokeConfig;
import com.pulumi.aws.lambda.FunctionEventInvokeConfigArgs;
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
        var example = new FunctionEventInvokeConfig("example", FunctionEventInvokeConfigArgs.builder()        
            .functionName(aws_lambda_alias.example().function_name())
            .maximumEventAgeInSeconds(60)
            .maximumRetryAttempts(0)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:FunctionEventInvokeConfig
    properties:
      functionName: ${aws_lambda_alias.example.function_name}
      maximumEventAgeInSeconds: 60
      maximumRetryAttempts: 0
```
{{% /example %}}
{{% example %}}
### Configuration for Alias Name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.FunctionEventInvokeConfig("example", {
    functionName: aws_lambda_alias.example.function_name,
    qualifier: aws_lambda_alias.example.name,
});
// ... other configuration ...
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lambda_.FunctionEventInvokeConfig("example",
    function_name=aws_lambda_alias["example"]["function_name"],
    qualifier=aws_lambda_alias["example"]["name"])
# ... other configuration ...
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.FunctionEventInvokeConfig("example", new()
    {
        FunctionName = aws_lambda_alias.Example.Function_name,
        Qualifier = aws_lambda_alias.Example.Name,
    });

    // ... other configuration ...
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewFunctionEventInvokeConfig(ctx, "example", &lambda.FunctionEventInvokeConfigArgs{
			FunctionName: pulumi.Any(aws_lambda_alias.Example.Function_name),
			Qualifier:    pulumi.Any(aws_lambda_alias.Example.Name),
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
import com.pulumi.aws.lambda.FunctionEventInvokeConfig;
import com.pulumi.aws.lambda.FunctionEventInvokeConfigArgs;
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
        var example = new FunctionEventInvokeConfig("example", FunctionEventInvokeConfigArgs.builder()        
            .functionName(aws_lambda_alias.example().function_name())
            .qualifier(aws_lambda_alias.example().name())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:FunctionEventInvokeConfig
    properties:
      functionName: ${aws_lambda_alias.example.function_name}
      qualifier: ${aws_lambda_alias.example.name}
```
{{% /example %}}
{{% example %}}
### Configuration for Function Latest Unpublished Version

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.FunctionEventInvokeConfig("example", {
    functionName: aws_lambda_function.example.function_name,
    qualifier: "$LATEST",
});
// ... other configuration ...
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lambda_.FunctionEventInvokeConfig("example",
    function_name=aws_lambda_function["example"]["function_name"],
    qualifier="$LATEST")
# ... other configuration ...
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.FunctionEventInvokeConfig("example", new()
    {
        FunctionName = aws_lambda_function.Example.Function_name,
        Qualifier = "$LATEST",
    });

    // ... other configuration ...
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewFunctionEventInvokeConfig(ctx, "example", &lambda.FunctionEventInvokeConfigArgs{
			FunctionName: pulumi.Any(aws_lambda_function.Example.Function_name),
			Qualifier:    pulumi.String("$LATEST"),
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
import com.pulumi.aws.lambda.FunctionEventInvokeConfig;
import com.pulumi.aws.lambda.FunctionEventInvokeConfigArgs;
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
        var example = new FunctionEventInvokeConfig("example", FunctionEventInvokeConfigArgs.builder()        
            .functionName(aws_lambda_function.example().function_name())
            .qualifier("$LATEST")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:FunctionEventInvokeConfig
    properties:
      functionName: ${aws_lambda_function.example.function_name}
      qualifier: $LATEST
```
{{% /example %}}
{{% example %}}
### Configuration for Function Published Version

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.FunctionEventInvokeConfig("example", {
    functionName: aws_lambda_function.example.function_name,
    qualifier: aws_lambda_function.example.version,
});
// ... other configuration ...
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lambda_.FunctionEventInvokeConfig("example",
    function_name=aws_lambda_function["example"]["function_name"],
    qualifier=aws_lambda_function["example"]["version"])
# ... other configuration ...
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.FunctionEventInvokeConfig("example", new()
    {
        FunctionName = aws_lambda_function.Example.Function_name,
        Qualifier = aws_lambda_function.Example.Version,
    });

    // ... other configuration ...
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewFunctionEventInvokeConfig(ctx, "example", &lambda.FunctionEventInvokeConfigArgs{
			FunctionName: pulumi.Any(aws_lambda_function.Example.Function_name),
			Qualifier:    pulumi.Any(aws_lambda_function.Example.Version),
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
import com.pulumi.aws.lambda.FunctionEventInvokeConfig;
import com.pulumi.aws.lambda.FunctionEventInvokeConfigArgs;
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
        var example = new FunctionEventInvokeConfig("example", FunctionEventInvokeConfigArgs.builder()        
            .functionName(aws_lambda_function.example().function_name())
            .qualifier(aws_lambda_function.example().version())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:FunctionEventInvokeConfig
    properties:
      functionName: ${aws_lambda_function.example.function_name}
      qualifier: ${aws_lambda_function.example.version}
```
{{% /example %}}
{{% /examples %}}

## Import

ARN with qualifier:

Name without qualifier (all versions and aliases):

Name with qualifier:

__Using `pulumi import` to import__ Lambda Function Event Invoke Configs using the fully qualified Function name or Amazon Resource Name (ARN). For example:

ARN without qualifier (all versions and aliases):

```sh
 $ pulumi import aws:lambda/functionEventInvokeConfig:FunctionEventInvokeConfig example arn:aws:us-east-1:123456789012:function:my_function
```
 ARN with qualifier:

```sh
 $ pulumi import aws:lambda/functionEventInvokeConfig:FunctionEventInvokeConfig example arn:aws:us-east-1:123456789012:function:my_function:production
```
 Name without qualifier (all versions and aliases):

```sh
 $ pulumi import aws:lambda/functionEventInvokeConfig:FunctionEventInvokeConfig example my_function
```
 Name with qualifier:

```sh
 $ pulumi import aws:lambda/functionEventInvokeConfig:FunctionEventInvokeConfig example my_function:production
```
 