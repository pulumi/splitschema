Provides a Step Function State Machine resource

{{% examples %}}
## Example Usage
{{% example %}}
### Basic (Standard Workflow)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ...
const sfnStateMachine = new aws.sfn.StateMachine("sfnStateMachine", {
    roleArn: aws_iam_role.iam_for_sfn.arn,
    definition: `{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "${aws_lambda_function.lambda.arn}",
      "End": true
    }
  }
}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

# ...
sfn_state_machine = aws.sfn.StateMachine("sfnStateMachine",
    role_arn=aws_iam_role["iam_for_sfn"]["arn"],
    definition=f"""{{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {{
    "HelloWorld": {{
      "Type": "Task",
      "Resource": "{aws_lambda_function["lambda"]["arn"]}",
      "End": true
    }}
  }}
}}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ...
    var sfnStateMachine = new Aws.Sfn.StateMachine("sfnStateMachine", new()
    {
        RoleArn = aws_iam_role.Iam_for_sfn.Arn,
        Definition = @$"{{
  ""Comment"": ""A Hello World example of the Amazon States Language using an AWS Lambda Function"",
  ""StartAt"": ""HelloWorld"",
  ""States"": {{
    ""HelloWorld"": {{
      ""Type"": ""Task"",
      ""Resource"": ""{aws_lambda_function.Lambda.Arn}"",
      ""End"": true
    }}
  }}
}}
",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sfn"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sfn.NewStateMachine(ctx, "sfnStateMachine", &sfn.StateMachineArgs{
			RoleArn: pulumi.Any(aws_iam_role.Iam_for_sfn.Arn),
			Definition: pulumi.String(fmt.Sprintf(`{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "%v",
      "End": true
    }
  }
}
`, aws_lambda_function.Lambda.Arn)),
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
import com.pulumi.aws.sfn.StateMachine;
import com.pulumi.aws.sfn.StateMachineArgs;
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
        var sfnStateMachine = new StateMachine("sfnStateMachine", StateMachineArgs.builder()        
            .roleArn(aws_iam_role.iam_for_sfn().arn())
            .definition("""
{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "%s",
      "End": true
    }
  }
}
", aws_lambda_function.lambda().arn()))
            .build());

    }
}
```
```yaml
resources:
  # ...
  sfnStateMachine:
    type: aws:sfn:StateMachine
    properties:
      roleArn: ${aws_iam_role.iam_for_sfn.arn}
      definition: |
        {
          "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
          "StartAt": "HelloWorld",
          "States": {
            "HelloWorld": {
              "Type": "Task",
              "Resource": "${aws_lambda_function.lambda.arn}",
              "End": true
            }
          }
        }
```
{{% /example %}}
{{% example %}}
### Basic (Express Workflow)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ...
const sfnStateMachine = new aws.sfn.StateMachine("sfnStateMachine", {
    roleArn: aws_iam_role.iam_for_sfn.arn,
    type: "EXPRESS",
    definition: `{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "${aws_lambda_function.lambda.arn}",
      "End": true
    }
  }
}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

# ...
sfn_state_machine = aws.sfn.StateMachine("sfnStateMachine",
    role_arn=aws_iam_role["iam_for_sfn"]["arn"],
    type="EXPRESS",
    definition=f"""{{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {{
    "HelloWorld": {{
      "Type": "Task",
      "Resource": "{aws_lambda_function["lambda"]["arn"]}",
      "End": true
    }}
  }}
}}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ...
    var sfnStateMachine = new Aws.Sfn.StateMachine("sfnStateMachine", new()
    {
        RoleArn = aws_iam_role.Iam_for_sfn.Arn,
        Type = "EXPRESS",
        Definition = @$"{{
  ""Comment"": ""A Hello World example of the Amazon States Language using an AWS Lambda Function"",
  ""StartAt"": ""HelloWorld"",
  ""States"": {{
    ""HelloWorld"": {{
      ""Type"": ""Task"",
      ""Resource"": ""{aws_lambda_function.Lambda.Arn}"",
      ""End"": true
    }}
  }}
}}
",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sfn"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sfn.NewStateMachine(ctx, "sfnStateMachine", &sfn.StateMachineArgs{
			RoleArn: pulumi.Any(aws_iam_role.Iam_for_sfn.Arn),
			Type:    pulumi.String("EXPRESS"),
			Definition: pulumi.String(fmt.Sprintf(`{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "%v",
      "End": true
    }
  }
}
`, aws_lambda_function.Lambda.Arn)),
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
import com.pulumi.aws.sfn.StateMachine;
import com.pulumi.aws.sfn.StateMachineArgs;
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
        var sfnStateMachine = new StateMachine("sfnStateMachine", StateMachineArgs.builder()        
            .roleArn(aws_iam_role.iam_for_sfn().arn())
            .type("EXPRESS")
            .definition("""
{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "%s",
      "End": true
    }
  }
}
", aws_lambda_function.lambda().arn()))
            .build());

    }
}
```
```yaml
resources:
  # ...
  sfnStateMachine:
    type: aws:sfn:StateMachine
    properties:
      roleArn: ${aws_iam_role.iam_for_sfn.arn}
      type: EXPRESS
      definition: |
        {
          "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
          "StartAt": "HelloWorld",
          "States": {
            "HelloWorld": {
              "Type": "Task",
              "Resource": "${aws_lambda_function.lambda.arn}",
              "End": true
            }
          }
        }
```
{{% /example %}}
{{% example %}}
### Publish (Publish SFN version)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ...
const sfnStateMachine = new aws.sfn.StateMachine("sfnStateMachine", {
    roleArn: aws_iam_role.iam_for_sfn.arn,
    publish: true,
    type: "EXPRESS",
    definition: `{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "${aws_lambda_function.lambda.arn}",
      "End": true
    }
  }
}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

# ...
sfn_state_machine = aws.sfn.StateMachine("sfnStateMachine",
    role_arn=aws_iam_role["iam_for_sfn"]["arn"],
    publish=True,
    type="EXPRESS",
    definition=f"""{{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {{
    "HelloWorld": {{
      "Type": "Task",
      "Resource": "{aws_lambda_function["lambda"]["arn"]}",
      "End": true
    }}
  }}
}}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ...
    var sfnStateMachine = new Aws.Sfn.StateMachine("sfnStateMachine", new()
    {
        RoleArn = aws_iam_role.Iam_for_sfn.Arn,
        Publish = true,
        Type = "EXPRESS",
        Definition = @$"{{
  ""Comment"": ""A Hello World example of the Amazon States Language using an AWS Lambda Function"",
  ""StartAt"": ""HelloWorld"",
  ""States"": {{
    ""HelloWorld"": {{
      ""Type"": ""Task"",
      ""Resource"": ""{aws_lambda_function.Lambda.Arn}"",
      ""End"": true
    }}
  }}
}}
",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sfn"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sfn.NewStateMachine(ctx, "sfnStateMachine", &sfn.StateMachineArgs{
			RoleArn: pulumi.Any(aws_iam_role.Iam_for_sfn.Arn),
			Publish: pulumi.Bool(true),
			Type:    pulumi.String("EXPRESS"),
			Definition: pulumi.String(fmt.Sprintf(`{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "%v",
      "End": true
    }
  }
}
`, aws_lambda_function.Lambda.Arn)),
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
import com.pulumi.aws.sfn.StateMachine;
import com.pulumi.aws.sfn.StateMachineArgs;
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
        var sfnStateMachine = new StateMachine("sfnStateMachine", StateMachineArgs.builder()        
            .roleArn(aws_iam_role.iam_for_sfn().arn())
            .publish(true)
            .type("EXPRESS")
            .definition("""
{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "%s",
      "End": true
    }
  }
}
", aws_lambda_function.lambda().arn()))
            .build());

    }
}
```
```yaml
resources:
  # ...
  sfnStateMachine:
    type: aws:sfn:StateMachine
    properties:
      roleArn: ${aws_iam_role.iam_for_sfn.arn}
      publish: true
      type: EXPRESS
      definition: |
        {
          "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
          "StartAt": "HelloWorld",
          "States": {
            "HelloWorld": {
              "Type": "Task",
              "Resource": "${aws_lambda_function.lambda.arn}",
              "End": true
            }
          }
        }
```
{{% /example %}}
{{% example %}}
### Logging

> *NOTE:* See the [AWS Step Functions Developer Guide](https://docs.aws.amazon.com/step-functions/latest/dg/welcome.html) for more information about enabling Step Function logging.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ...
const sfnStateMachine = new aws.sfn.StateMachine("sfnStateMachine", {
    roleArn: aws_iam_role.iam_for_sfn.arn,
    definition: `{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "${aws_lambda_function.lambda.arn}",
      "End": true
    }
  }
}
`,
    loggingConfiguration: {
        logDestination: `${aws_cloudwatch_log_group.log_group_for_sfn.arn}:*`,
        includeExecutionData: true,
        level: "ERROR",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

# ...
sfn_state_machine = aws.sfn.StateMachine("sfnStateMachine",
    role_arn=aws_iam_role["iam_for_sfn"]["arn"],
    definition=f"""{{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {{
    "HelloWorld": {{
      "Type": "Task",
      "Resource": "{aws_lambda_function["lambda"]["arn"]}",
      "End": true
    }}
  }}
}}
""",
    logging_configuration=aws.sfn.StateMachineLoggingConfigurationArgs(
        log_destination=f"{aws_cloudwatch_log_group['log_group_for_sfn']['arn']}:*",
        include_execution_data=True,
        level="ERROR",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ...
    var sfnStateMachine = new Aws.Sfn.StateMachine("sfnStateMachine", new()
    {
        RoleArn = aws_iam_role.Iam_for_sfn.Arn,
        Definition = @$"{{
  ""Comment"": ""A Hello World example of the Amazon States Language using an AWS Lambda Function"",
  ""StartAt"": ""HelloWorld"",
  ""States"": {{
    ""HelloWorld"": {{
      ""Type"": ""Task"",
      ""Resource"": ""{aws_lambda_function.Lambda.Arn}"",
      ""End"": true
    }}
  }}
}}
",
        LoggingConfiguration = new Aws.Sfn.Inputs.StateMachineLoggingConfigurationArgs
        {
            LogDestination = $"{aws_cloudwatch_log_group.Log_group_for_sfn.Arn}:*",
            IncludeExecutionData = true,
            Level = "ERROR",
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sfn"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sfn.NewStateMachine(ctx, "sfnStateMachine", &sfn.StateMachineArgs{
			RoleArn: pulumi.Any(aws_iam_role.Iam_for_sfn.Arn),
			Definition: pulumi.String(fmt.Sprintf(`{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "%v",
      "End": true
    }
  }
}
`, aws_lambda_function.Lambda.Arn)),
			LoggingConfiguration: &sfn.StateMachineLoggingConfigurationArgs{
				LogDestination:       pulumi.String(fmt.Sprintf("%v:*", aws_cloudwatch_log_group.Log_group_for_sfn.Arn)),
				IncludeExecutionData: pulumi.Bool(true),
				Level:                pulumi.String("ERROR"),
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
import com.pulumi.aws.sfn.StateMachine;
import com.pulumi.aws.sfn.StateMachineArgs;
import com.pulumi.aws.sfn.inputs.StateMachineLoggingConfigurationArgs;
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
        var sfnStateMachine = new StateMachine("sfnStateMachine", StateMachineArgs.builder()        
            .roleArn(aws_iam_role.iam_for_sfn().arn())
            .definition("""
{
  "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "%s",
      "End": true
    }
  }
}
", aws_lambda_function.lambda().arn()))
            .loggingConfiguration(StateMachineLoggingConfigurationArgs.builder()
                .logDestination(String.format("%s:*", aws_cloudwatch_log_group.log_group_for_sfn().arn()))
                .includeExecutionData(true)
                .level("ERROR")
                .build())
            .build());

    }
}
```
```yaml
resources:
  # ...
  sfnStateMachine:
    type: aws:sfn:StateMachine
    properties:
      roleArn: ${aws_iam_role.iam_for_sfn.arn}
      definition: |
        {
          "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
          "StartAt": "HelloWorld",
          "States": {
            "HelloWorld": {
              "Type": "Task",
              "Resource": "${aws_lambda_function.lambda.arn}",
              "End": true
            }
          }
        }
      loggingConfiguration:
        logDestination: ${aws_cloudwatch_log_group.log_group_for_sfn.arn}:*
        includeExecutionData: true
        level: ERROR
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import State Machines using the `arn`. For example:

```sh
 $ pulumi import aws:sfn/stateMachine:StateMachine foo arn:aws:states:eu-west-1:123456789098:stateMachine:bar
```
 