Manages an NFS Location within AWS DataSync.

> **NOTE:** The DataSync Agents must be available before creating this resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.NfsLocation("example", {
    serverHostname: "nfs.example.com",
    subdirectory: "/exported/path",
    onPremConfig: {
        agentArns: [aws_datasync_agent.example.arn],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datasync.NfsLocation("example",
    server_hostname="nfs.example.com",
    subdirectory="/exported/path",
    on_prem_config=aws.datasync.NfsLocationOnPremConfigArgs(
        agent_arns=[aws_datasync_agent["example"]["arn"]],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataSync.NfsLocation("example", new()
    {
        ServerHostname = "nfs.example.com",
        Subdirectory = "/exported/path",
        OnPremConfig = new Aws.DataSync.Inputs.NfsLocationOnPremConfigArgs
        {
            AgentArns = new[]
            {
                aws_datasync_agent.Example.Arn,
            },
        },
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
		_, err := datasync.NewNfsLocation(ctx, "example", &datasync.NfsLocationArgs{
			ServerHostname: pulumi.String("nfs.example.com"),
			Subdirectory:   pulumi.String("/exported/path"),
			OnPremConfig: &datasync.NfsLocationOnPremConfigArgs{
				AgentArns: pulumi.StringArray{
					aws_datasync_agent.Example.Arn,
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
import com.pulumi.aws.datasync.NfsLocation;
import com.pulumi.aws.datasync.NfsLocationArgs;
import com.pulumi.aws.datasync.inputs.NfsLocationOnPremConfigArgs;
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
        var example = new NfsLocation("example", NfsLocationArgs.builder()        
            .serverHostname("nfs.example.com")
            .subdirectory("/exported/path")
            .onPremConfig(NfsLocationOnPremConfigArgs.builder()
                .agentArns(aws_datasync_agent.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:datasync:NfsLocation
    properties:
      serverHostname: nfs.example.com
      subdirectory: /exported/path
      onPremConfig:
        agentArns:
          - ${aws_datasync_agent.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_location_nfs` using the DataSync Task Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:datasync/nfsLocation:NfsLocation example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
```
 