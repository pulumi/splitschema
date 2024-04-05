Manages a SMB Location within AWS DataSync.

> **NOTE:** The DataSync Agents must be available before creating this resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.LocationSmb("example", {
    serverHostname: "smb.example.com",
    subdirectory: "/exported/path",
    user: "Guest",
    password: "ANotGreatPassword",
    agentArns: [aws_datasync_agent.example.arn],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datasync.LocationSmb("example",
    server_hostname="smb.example.com",
    subdirectory="/exported/path",
    user="Guest",
    password="ANotGreatPassword",
    agent_arns=[aws_datasync_agent["example"]["arn"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataSync.LocationSmb("example", new()
    {
        ServerHostname = "smb.example.com",
        Subdirectory = "/exported/path",
        User = "Guest",
        Password = "ANotGreatPassword",
        AgentArns = new[]
        {
            aws_datasync_agent.Example.Arn,
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
		_, err := datasync.NewLocationSmb(ctx, "example", &datasync.LocationSmbArgs{
			ServerHostname: pulumi.String("smb.example.com"),
			Subdirectory:   pulumi.String("/exported/path"),
			User:           pulumi.String("Guest"),
			Password:       pulumi.String("ANotGreatPassword"),
			AgentArns: pulumi.StringArray{
				aws_datasync_agent.Example.Arn,
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
import com.pulumi.aws.datasync.LocationSmb;
import com.pulumi.aws.datasync.LocationSmbArgs;
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
        var example = new LocationSmb("example", LocationSmbArgs.builder()        
            .serverHostname("smb.example.com")
            .subdirectory("/exported/path")
            .user("Guest")
            .password("ANotGreatPassword")
            .agentArns(aws_datasync_agent.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:datasync:LocationSmb
    properties:
      serverHostname: smb.example.com
      subdirectory: /exported/path
      user: Guest
      password: ANotGreatPassword
      agentArns:
        - ${aws_datasync_agent.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_location_smb` using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:datasync/locationSmb:LocationSmb example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
```
 