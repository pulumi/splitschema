Provides a CodeDeploy application to be used as a basis for deployments

{{% examples %}}
## Example Usage
{{% example %}}
### ECS Application

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codedeploy.Application("example", {computePlatform: "ECS"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codedeploy.Application("example", compute_platform="ECS")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeDeploy.Application("example", new()
    {
        ComputePlatform = "ECS",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codedeploy"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codedeploy.NewApplication(ctx, "example", &codedeploy.ApplicationArgs{
			ComputePlatform: pulumi.String("ECS"),
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
import com.pulumi.aws.codedeploy.Application;
import com.pulumi.aws.codedeploy.ApplicationArgs;
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
        var example = new Application("example", ApplicationArgs.builder()        
            .computePlatform("ECS")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codedeploy:Application
    properties:
      computePlatform: ECS
```
{{% /example %}}
{{% example %}}
### Lambda Application

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codedeploy.Application("example", {computePlatform: "Lambda"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codedeploy.Application("example", compute_platform="Lambda")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeDeploy.Application("example", new()
    {
        ComputePlatform = "Lambda",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codedeploy"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codedeploy.NewApplication(ctx, "example", &codedeploy.ApplicationArgs{
			ComputePlatform: pulumi.String("Lambda"),
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
import com.pulumi.aws.codedeploy.Application;
import com.pulumi.aws.codedeploy.ApplicationArgs;
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
        var example = new Application("example", ApplicationArgs.builder()        
            .computePlatform("Lambda")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codedeploy:Application
    properties:
      computePlatform: Lambda
```
{{% /example %}}
{{% example %}}
### Server Application

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codedeploy.Application("example", {computePlatform: "Server"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codedeploy.Application("example", compute_platform="Server")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeDeploy.Application("example", new()
    {
        ComputePlatform = "Server",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codedeploy"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codedeploy.NewApplication(ctx, "example", &codedeploy.ApplicationArgs{
			ComputePlatform: pulumi.String("Server"),
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
import com.pulumi.aws.codedeploy.Application;
import com.pulumi.aws.codedeploy.ApplicationArgs;
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
        var example = new Application("example", ApplicationArgs.builder()        
            .computePlatform("Server")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codedeploy:Application
    properties:
      computePlatform: Server
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeDeploy Applications using the `name`. For example:

```sh
 $ pulumi import aws:codedeploy/application:Application example my-application
```
 