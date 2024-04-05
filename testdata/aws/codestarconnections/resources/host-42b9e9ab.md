Provides a CodeStar Host.

> **NOTE:** The `aws.codestarconnections.Host` resource is created in the state `PENDING`. Authentication with the host provider must be completed in the AWS Console. For more information visit [Set up a pending host](https://docs.aws.amazon.com/dtconsole/latest/userguide/connections-host-setup.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codestarconnections.Host("example", {
    providerEndpoint: "https://example.com",
    providerType: "GitHubEnterpriseServer",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codestarconnections.Host("example",
    provider_endpoint="https://example.com",
    provider_type="GitHubEnterpriseServer")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeStarConnections.Host("example", new()
    {
        ProviderEndpoint = "https://example.com",
        ProviderType = "GitHubEnterpriseServer",
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
		_, err := codestarconnections.NewHost(ctx, "example", &codestarconnections.HostArgs{
			ProviderEndpoint: pulumi.String("https://example.com"),
			ProviderType:     pulumi.String("GitHubEnterpriseServer"),
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
import com.pulumi.aws.codestarconnections.Host;
import com.pulumi.aws.codestarconnections.HostArgs;
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
        var example = new Host("example", HostArgs.builder()        
            .providerEndpoint("https://example.com")
            .providerType("GitHubEnterpriseServer")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codestarconnections:Host
    properties:
      providerEndpoint: https://example.com
      providerType: GitHubEnterpriseServer
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeStar Host using the ARN. For example:

```sh
 $ pulumi import aws:codestarconnections/host:Host example-host arn:aws:codestar-connections:us-west-1:0123456789:host/79d4d357-a2ee-41e4-b350-2fe39ae59448
```
 