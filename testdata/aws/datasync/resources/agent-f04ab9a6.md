Manages an AWS DataSync Agent deployed on premises.

> **NOTE:** One of `activation_key` or `ip_address` must be provided for resource creation (agent activation). Neither is required for resource import. If using `ip_address`, this provider must be able to make an HTTP (port 80) GET request to the specified IP address from where it is running. The agent will turn off that HTTP server after activation.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.Agent("example", {ipAddress: "1.2.3.4"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datasync.Agent("example", ip_address="1.2.3.4")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataSync.Agent("example", new()
    {
        IpAddress = "1.2.3.4",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewAgent(ctx, "example", &datasync.AgentArgs{
			IpAddress: pulumi.String("1.2.3.4"),
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
import com.pulumi.aws.datasync.Agent;
import com.pulumi.aws.datasync.AgentArgs;
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
        var example = new Agent("example", AgentArgs.builder()        
            .ipAddress("1.2.3.4")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:datasync:Agent
    properties:
      ipAddress: 1.2.3.4
```

{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_agent` using the DataSync Agent Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:datasync/agent:Agent example arn:aws:datasync:us-east-1:123456789012:agent/agent-12345678901234567
```
 