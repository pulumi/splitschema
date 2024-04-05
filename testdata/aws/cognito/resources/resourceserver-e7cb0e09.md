Provides a Cognito Resource Server.

{{% examples %}}
## Example Usage
{{% example %}}
### Create a basic resource server

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pool = new aws.cognito.UserPool("pool", {});
const resource = new aws.cognito.ResourceServer("resource", {
    identifier: "https://example.com",
    userPoolId: pool.id,
});
```
```python
import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")
resource = aws.cognito.ResourceServer("resource",
    identifier="https://example.com",
    user_pool_id=pool.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var pool = new Aws.Cognito.UserPool("pool");

    var resource = new Aws.Cognito.ResourceServer("resource", new()
    {
        Identifier = "https://example.com",
        UserPoolId = pool.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		pool, err := cognito.NewUserPool(ctx, "pool", nil)
		if err != nil {
			return err
		}
		_, err = cognito.NewResourceServer(ctx, "resource", &cognito.ResourceServerArgs{
			Identifier: pulumi.String("https://example.com"),
			UserPoolId: pool.ID(),
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.ResourceServer;
import com.pulumi.aws.cognito.ResourceServerArgs;
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
        var pool = new UserPool("pool");

        var resource = new ResourceServer("resource", ResourceServerArgs.builder()        
            .identifier("https://example.com")
            .userPoolId(pool.id())
            .build());

    }
}
```
```yaml
resources:
  pool:
    type: aws:cognito:UserPool
  resource:
    type: aws:cognito:ResourceServer
    properties:
      identifier: https://example.com
      userPoolId: ${pool.id}
```
{{% /example %}}
{{% example %}}
### Create a resource server with sample-scope

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pool = new aws.cognito.UserPool("pool", {});
const resource = new aws.cognito.ResourceServer("resource", {
    identifier: "https://example.com",
    scopes: [{
        scopeName: "sample-scope",
        scopeDescription: "a Sample Scope Description",
    }],
    userPoolId: pool.id,
});
```
```python
import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")
resource = aws.cognito.ResourceServer("resource",
    identifier="https://example.com",
    scopes=[aws.cognito.ResourceServerScopeArgs(
        scope_name="sample-scope",
        scope_description="a Sample Scope Description",
    )],
    user_pool_id=pool.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var pool = new Aws.Cognito.UserPool("pool");

    var resource = new Aws.Cognito.ResourceServer("resource", new()
    {
        Identifier = "https://example.com",
        Scopes = new[]
        {
            new Aws.Cognito.Inputs.ResourceServerScopeArgs
            {
                ScopeName = "sample-scope",
                ScopeDescription = "a Sample Scope Description",
            },
        },
        UserPoolId = pool.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		pool, err := cognito.NewUserPool(ctx, "pool", nil)
		if err != nil {
			return err
		}
		_, err = cognito.NewResourceServer(ctx, "resource", &cognito.ResourceServerArgs{
			Identifier: pulumi.String("https://example.com"),
			Scopes: cognito.ResourceServerScopeArray{
				&cognito.ResourceServerScopeArgs{
					ScopeName:        pulumi.String("sample-scope"),
					ScopeDescription: pulumi.String("a Sample Scope Description"),
				},
			},
			UserPoolId: pool.ID(),
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.ResourceServer;
import com.pulumi.aws.cognito.ResourceServerArgs;
import com.pulumi.aws.cognito.inputs.ResourceServerScopeArgs;
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
        var pool = new UserPool("pool");

        var resource = new ResourceServer("resource", ResourceServerArgs.builder()        
            .identifier("https://example.com")
            .scopes(ResourceServerScopeArgs.builder()
                .scopeName("sample-scope")
                .scopeDescription("a Sample Scope Description")
                .build())
            .userPoolId(pool.id())
            .build());

    }
}
```
```yaml
resources:
  pool:
    type: aws:cognito:UserPool
  resource:
    type: aws:cognito:ResourceServer
    properties:
      identifier: https://example.com
      scopes:
        - scopeName: sample-scope
          scopeDescription: a Sample Scope Description
      userPoolId: ${pool.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_cognito_resource_server` using their User Pool ID and Identifier. For example:

```sh
 $ pulumi import aws:cognito/resourceServer:ResourceServer example "us-west-2_abc123|https://example.com"
```
 