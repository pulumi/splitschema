Use this resource to invoke a lambda function. The lambda function is invoked with the [RequestResponse](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax) invocation type.

> **NOTE:** By default this resource _only_ invokes the function when the arguments call for a create or replace. In other words, after an initial invocation on _apply_, if the arguments do not change, a subsequent _apply_ does not invoke the function again. To dynamically invoke the function, see the `triggers` example below. To always invoke a function on each _apply_, see the `aws.lambda.Invocation` data source. To invoke the lambda function when the Pulumi resource is updated and deleted, see the CRUD Lifecycle Scope example below.

> **NOTE:** If you get a `KMSAccessDeniedException: Lambda was unable to decrypt the environment variables because KMS access was denied` error when invoking an `aws.lambda.Function` with environment variables, the IAM role associated with the function may have been deleted and recreated _after_ the function was created. You can fix the problem two ways: 1) updating the function's role to another role and then updating it back again to the recreated role, or 2) by using Pulumi to `taint` the function and `apply` your configuration again to recreate the function. (When you create a function, Lambda grants permissions on the KMS key to the function's IAM role. If the IAM role is recreated, the grant is no longer valid. Changing the function's role or recreating the function causes Lambda to update the grant.)

{{% examples %}}
## Example Usage
{{% example %}}
### Dynamic Invocation Example Using Triggers

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as crypto from "crypto";

const example = new aws.lambda.Invocation("example", {
    functionName: aws_lambda_function.lambda_function_test.function_name,
    triggers: {
        redeployment: crypto.createHash('sha1').update(JSON.stringify([aws_lambda_function.example.environment])).digest('hex'),
    },
    input: JSON.stringify({
        key1: "value1",
        key2: "value2",
    }),
});
```
```python
import pulumi
import hashlib
import json
import pulumi_aws as aws

example = aws.lambda_.Invocation("example",
    function_name=aws_lambda_function["lambda_function_test"]["function_name"],
    triggers={
        "redeployment": hashlib.sha1(json.dumps([aws_lambda_function["example"]["environment"]]).encode()).hexdigest(),
    },
    input=json.dumps({
        "key1": "value1",
        "key2": "value2",
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

	
string ComputeSHA1(string input) 
{
    var hash = SHA1.Create().ComputeHash(Encoding.UTF8.GetBytes(input));
    return BitConverter.ToString(hash).Replace("-","").ToLowerInvariant();
}

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.Invocation("example", new()
    {
        FunctionName = aws_lambda_function.Lambda_function_test.Function_name,
        Triggers = 
        {
            { "redeployment", ComputeSHA1(JsonSerializer.Serialize(new[]
            {
                aws_lambda_function.Example.Environment,
            })) },
        },
        Input = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["key1"] = "value1",
            ["key2"] = "value2",
        }),
    });

});
```
```go
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func sha1Hash(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal([]interface{}{
			aws_lambda_function.Example.Environment,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		tmpJSON1, err := json.Marshal(map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		})
		if err != nil {
			return err
		}
		json1 := string(tmpJSON1)
		_, err = lambda.NewInvocation(ctx, "example", &lambda.InvocationArgs{
			FunctionName: pulumi.Any(aws_lambda_function.Lambda_function_test.Function_name),
			Triggers: pulumi.StringMap{
				"redeployment": sha1Hash(json0),
			},
			Input: pulumi.String(json1),
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
import com.pulumi.aws.lambda.Invocation;
import com.pulumi.aws.lambda.InvocationArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var example = new Invocation("example", InvocationArgs.builder()        
            .functionName(aws_lambda_function.lambda_function_test().function_name())
            .triggers(Map.of("redeployment", computeSHA1(serializeJson(
                jsonArray(aws_lambda_function.example().environment())))))
            .input(serializeJson(
                jsonObject(
                    jsonProperty("key1", "value1"),
                    jsonProperty("key2", "value2")
                )))
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### CRUD Lifecycle Scope

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.Invocation("example", {
    functionName: aws_lambda_function.lambda_function_test.function_name,
    input: JSON.stringify({
        key1: "value1",
        key2: "value2",
    }),
    lifecycleScope: "CRUD",
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.lambda_.Invocation("example",
    function_name=aws_lambda_function["lambda_function_test"]["function_name"],
    input=json.dumps({
        "key1": "value1",
        "key2": "value2",
    }),
    lifecycle_scope="CRUD")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.Invocation("example", new()
    {
        FunctionName = aws_lambda_function.Lambda_function_test.Function_name,
        Input = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["key1"] = "value1",
            ["key2"] = "value2",
        }),
        LifecycleScope = "CRUD",
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = lambda.NewInvocation(ctx, "example", &lambda.InvocationArgs{
			FunctionName:   pulumi.Any(aws_lambda_function.Lambda_function_test.Function_name),
			Input:          pulumi.String(json0),
			LifecycleScope: pulumi.String("CRUD"),
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
import com.pulumi.aws.lambda.Invocation;
import com.pulumi.aws.lambda.InvocationArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var example = new Invocation("example", InvocationArgs.builder()        
            .functionName(aws_lambda_function.lambda_function_test().function_name())
            .input(serializeJson(
                jsonObject(
                    jsonProperty("key1", "value1"),
                    jsonProperty("key2", "value2")
                )))
            .lifecycleScope("CRUD")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:Invocation
    properties:
      functionName: ${aws_lambda_function.lambda_function_test.function_name}
      input:
        fn::toJSON:
          key1: value1
          key2: value2
      lifecycleScope: CRUD
```

> **NOTE:** `lifecycle_scope = "CRUD"` will inject a key `tf` in the input event to pass lifecycle information! This allows the lambda function to handle different lifecycle transitions uniquely.  If you need to use a key `tf` in your own input JSON, the default key name can be overridden with the `pulumi_key` argument.

The key `tf` gets added with subkeys:

* `action` - Action Pulumi performs on the resource. Values are `create`, `update`, or `delete`.
* `prev_input` - Input JSON payload from the previous invocation. This can be used to handle update and delete events.

When the resource from the example above is created, the Lambda will get following JSON payload:

```typescript
import * as pulumi from "@pulumi/pulumi";
```
```python
import pulumi
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() => 
{
});
```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
    }
}
```
```yaml
{}
```

If the input value of `key1` changes to "valueB", then the lambda will be invoked again with the following JSON payload:

```typescript
import * as pulumi from "@pulumi/pulumi";
```
```python
import pulumi
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() => 
{
});
```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
    }
}
```
```yaml
{}
```

When the invocation resource is removed, the final invocation will have the following JSON payload:

```typescript
import * as pulumi from "@pulumi/pulumi";
```
```python
import pulumi
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() => 
{
});
```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
    }
}
```
```yaml
{}
```
{{% /example %}}
{{% /examples %}}