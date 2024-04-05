Provides a CodeBuild Source Credentials Resource.

> **NOTE:**
[Codebuild only allows a single credential per given server type in a given region](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_codebuild.GitHubSourceCredentials.html). Therefore, when you define `aws.codebuild.SourceCredential`, `aws.codebuild.Project` resource defined in the same module will use it.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codebuild.SourceCredential("example", {
    authType: "PERSONAL_ACCESS_TOKEN",
    serverType: "GITHUB",
    token: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codebuild.SourceCredential("example",
    auth_type="PERSONAL_ACCESS_TOKEN",
    server_type="GITHUB",
    token="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeBuild.SourceCredential("example", new()
    {
        AuthType = "PERSONAL_ACCESS_TOKEN",
        ServerType = "GITHUB",
        Token = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codebuild"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codebuild.NewSourceCredential(ctx, "example", &codebuild.SourceCredentialArgs{
			AuthType:   pulumi.String("PERSONAL_ACCESS_TOKEN"),
			ServerType: pulumi.String("GITHUB"),
			Token:      pulumi.String("example"),
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
import com.pulumi.aws.codebuild.SourceCredential;
import com.pulumi.aws.codebuild.SourceCredentialArgs;
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
        var example = new SourceCredential("example", SourceCredentialArgs.builder()        
            .authType("PERSONAL_ACCESS_TOKEN")
            .serverType("GITHUB")
            .token("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codebuild:SourceCredential
    properties:
      authType: PERSONAL_ACCESS_TOKEN
      serverType: GITHUB
      token: example
```
{{% /example %}}
{{% example %}}
### Bitbucket Server Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codebuild.SourceCredential("example", {
    authType: "BASIC_AUTH",
    serverType: "BITBUCKET",
    token: "example",
    userName: "test-user",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codebuild.SourceCredential("example",
    auth_type="BASIC_AUTH",
    server_type="BITBUCKET",
    token="example",
    user_name="test-user")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeBuild.SourceCredential("example", new()
    {
        AuthType = "BASIC_AUTH",
        ServerType = "BITBUCKET",
        Token = "example",
        UserName = "test-user",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codebuild"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codebuild.NewSourceCredential(ctx, "example", &codebuild.SourceCredentialArgs{
			AuthType:   pulumi.String("BASIC_AUTH"),
			ServerType: pulumi.String("BITBUCKET"),
			Token:      pulumi.String("example"),
			UserName:   pulumi.String("test-user"),
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
import com.pulumi.aws.codebuild.SourceCredential;
import com.pulumi.aws.codebuild.SourceCredentialArgs;
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
        var example = new SourceCredential("example", SourceCredentialArgs.builder()        
            .authType("BASIC_AUTH")
            .serverType("BITBUCKET")
            .token("example")
            .userName("test-user")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codebuild:SourceCredential
    properties:
      authType: BASIC_AUTH
      serverType: BITBUCKET
      token: example
      userName: test-user
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeBuild Source Credential using the CodeBuild Source Credential arn. For example:

```sh
 $ pulumi import aws:codebuild/sourceCredential:SourceCredential example arn:aws:codebuild:us-west-2:123456789:token:github
```
 