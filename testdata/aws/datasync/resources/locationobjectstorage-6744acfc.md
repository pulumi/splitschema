Manages a Object Storage Location within AWS DataSync.

> **NOTE:** The DataSync Agents must be available before creating this resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.LocationObjectStorage("example", {
    agentArns: [aws_datasync_agent.example.arn],
    serverHostname: "example",
    bucketName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datasync.LocationObjectStorage("example",
    agent_arns=[aws_datasync_agent["example"]["arn"]],
    server_hostname="example",
    bucket_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataSync.LocationObjectStorage("example", new()
    {
        AgentArns = new[]
        {
            aws_datasync_agent.Example.Arn,
        },
        ServerHostname = "example",
        BucketName = "example",
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
		_, err := datasync.NewLocationObjectStorage(ctx, "example", &datasync.LocationObjectStorageArgs{
			AgentArns: pulumi.StringArray{
				aws_datasync_agent.Example.Arn,
			},
			ServerHostname: pulumi.String("example"),
			BucketName:     pulumi.String("example"),
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
import com.pulumi.aws.datasync.LocationObjectStorage;
import com.pulumi.aws.datasync.LocationObjectStorageArgs;
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
        var example = new LocationObjectStorage("example", LocationObjectStorageArgs.builder()        
            .agentArns(aws_datasync_agent.example().arn())
            .serverHostname("example")
            .bucketName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:datasync:LocationObjectStorage
    properties:
      agentArns:
        - ${aws_datasync_agent.example.arn}
      serverHostname: example
      bucketName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_location_object_storage` using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:datasync/locationObjectStorage:LocationObjectStorage example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
```
 