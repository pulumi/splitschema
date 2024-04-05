Provides details about CodeStar Connection.

{{% examples %}}
## Example Usage
{{% example %}}
### By ARN

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.codestarconnections.getConnection({
    arn: aws_codestarconnections_connection.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codestarconnections.get_connection(arn=aws_codestarconnections_connection["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.CodeStarConnections.GetConnection.Invoke(new()
    {
        Arn = aws_codestarconnections_connection.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codestarconnections"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codestarconnections.LookupConnection(ctx, &codestarconnections.LookupConnectionArgs{
			Arn: pulumi.StringRef(aws_codestarconnections_connection.Example.Arn),
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
import com.pulumi.aws.codestarconnections.CodestarconnectionsFunctions;
import com.pulumi.aws.codestarconnections.inputs.GetConnectionArgs;
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
        final var example = CodestarconnectionsFunctions.getConnection(GetConnectionArgs.builder()
            .arn(aws_codestarconnections_connection.example().arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:codestarconnections:getConnection
      Arguments:
        arn: ${aws_codestarconnections_connection.example.arn}
```
{{% /example %}}
{{% example %}}
### By Name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.codestarconnections.getConnection({
    name: aws_codestarconnections_connection.example.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codestarconnections.get_connection(name=aws_codestarconnections_connection["example"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.CodeStarConnections.GetConnection.Invoke(new()
    {
        Name = aws_codestarconnections_connection.Example.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codestarconnections"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codestarconnections.LookupConnection(ctx, &codestarconnections.LookupConnectionArgs{
			Name: pulumi.StringRef(aws_codestarconnections_connection.Example.Name),
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
import com.pulumi.aws.codestarconnections.CodestarconnectionsFunctions;
import com.pulumi.aws.codestarconnections.inputs.GetConnectionArgs;
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
        final var example = CodestarconnectionsFunctions.getConnection(GetConnectionArgs.builder()
            .name(aws_codestarconnections_connection.example().name())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:codestarconnections:getConnection
      Arguments:
        name: ${aws_codestarconnections_connection.example.name}
```
{{% /example %}}
{{% /examples %}}