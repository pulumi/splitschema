Use this data source to get the ARNs and Names of IAM Users.

{{% examples %}}
## Example Usage
{{% example %}}
### All users in an account

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const users = aws.iam.getUsers({});
```
```python
import pulumi
import pulumi_aws as aws

users = aws.iam.get_users()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var users = Aws.Iam.GetUsers.Invoke();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.GetUsers(ctx, nil, nil)
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetUsersArgs;
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
        final var users = IamFunctions.getUsers();

    }
}
```
```yaml
variables:
  users:
    fn::invoke:
      Function: aws:iam:getUsers
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### Users filtered by name regex

Users whose username contains `abc`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const users = aws.iam.getUsers({
    nameRegex: ".*abc.*",
});
```
```python
import pulumi
import pulumi_aws as aws

users = aws.iam.get_users(name_regex=".*abc.*")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var users = Aws.Iam.GetUsers.Invoke(new()
    {
        NameRegex = ".*abc.*",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.GetUsers(ctx, &iam.GetUsersArgs{
			NameRegex: pulumi.StringRef(".*abc.*"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetUsersArgs;
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
        final var users = IamFunctions.getUsers(GetUsersArgs.builder()
            .nameRegex(".*abc.*")
            .build());

    }
}
```
```yaml
variables:
  users:
    fn::invoke:
      Function: aws:iam:getUsers
      Arguments:
        nameRegex: .*abc.*
```
{{% /example %}}
{{% example %}}
### Users filtered by path prefix

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const users = aws.iam.getUsers({
    pathPrefix: "/custom-path",
});
```
```python
import pulumi
import pulumi_aws as aws

users = aws.iam.get_users(path_prefix="/custom-path")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var users = Aws.Iam.GetUsers.Invoke(new()
    {
        PathPrefix = "/custom-path",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.GetUsers(ctx, &iam.GetUsersArgs{
			PathPrefix: pulumi.StringRef("/custom-path"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetUsersArgs;
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
        final var users = IamFunctions.getUsers(GetUsersArgs.builder()
            .pathPrefix("/custom-path")
            .build());

    }
}
```
```yaml
variables:
  users:
    fn::invoke:
      Function: aws:iam:getUsers
      Arguments:
        pathPrefix: /custom-path
```
{{% /example %}}
{{% /examples %}}