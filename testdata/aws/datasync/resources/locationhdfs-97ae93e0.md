Manages an HDFS Location within AWS DataSync.

> **NOTE:** The DataSync Agents must be available before creating this resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.LocationHdfs("example", {
    agentArns: [aws_datasync_agent.example.arn],
    authenticationType: "SIMPLE",
    simpleUser: "example",
    nameNodes: [{
        hostname: aws_instance.example.private_dns,
        port: 80,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datasync.LocationHdfs("example",
    agent_arns=[aws_datasync_agent["example"]["arn"]],
    authentication_type="SIMPLE",
    simple_user="example",
    name_nodes=[aws.datasync.LocationHdfsNameNodeArgs(
        hostname=aws_instance["example"]["private_dns"],
        port=80,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataSync.LocationHdfs("example", new()
    {
        AgentArns = new[]
        {
            aws_datasync_agent.Example.Arn,
        },
        AuthenticationType = "SIMPLE",
        SimpleUser = "example",
        NameNodes = new[]
        {
            new Aws.DataSync.Inputs.LocationHdfsNameNodeArgs
            {
                Hostname = aws_instance.Example.Private_dns,
                Port = 80,
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
		_, err := datasync.NewLocationHdfs(ctx, "example", &datasync.LocationHdfsArgs{
			AgentArns: pulumi.StringArray{
				aws_datasync_agent.Example.Arn,
			},
			AuthenticationType: pulumi.String("SIMPLE"),
			SimpleUser:         pulumi.String("example"),
			NameNodes: datasync.LocationHdfsNameNodeArray{
				&datasync.LocationHdfsNameNodeArgs{
					Hostname: pulumi.Any(aws_instance.Example.Private_dns),
					Port:     pulumi.Int(80),
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
import com.pulumi.aws.datasync.LocationHdfs;
import com.pulumi.aws.datasync.LocationHdfsArgs;
import com.pulumi.aws.datasync.inputs.LocationHdfsNameNodeArgs;
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
        var example = new LocationHdfs("example", LocationHdfsArgs.builder()        
            .agentArns(aws_datasync_agent.example().arn())
            .authenticationType("SIMPLE")
            .simpleUser("example")
            .nameNodes(LocationHdfsNameNodeArgs.builder()
                .hostname(aws_instance.example().private_dns())
                .port(80)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:datasync:LocationHdfs
    properties:
      agentArns:
        - ${aws_datasync_agent.example.arn}
      authenticationType: SIMPLE
      simpleUser: example
      nameNodes:
        - hostname: ${aws_instance.example.private_dns}
          port: 80
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_location_hdfs` using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:datasync/locationHdfs:LocationHdfs example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
```
 