Manages a Microsoft Azure Blob Storage Location within AWS DataSync.

> **NOTE:** The DataSync Agents must be available before creating this resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.LocationAzureBlob("example", {
    agentArns: [aws_datasync_agent.example.arn],
    authenticationType: "SAS",
    containerUrl: "https://example.com/path",
    sasConfiguration: {
        token: "sp=r&st=2023-12-20T14:54:52Z&se=2023-12-20T22:54:52Z&spr=https&sv=2021-06-08&sr=c&sig=aBBKDWQvyuVcTPH9EBp%2FXTI9E%2F%2Fmq171%2BZU178wcwqU%3D",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datasync.LocationAzureBlob("example",
    agent_arns=[aws_datasync_agent["example"]["arn"]],
    authentication_type="SAS",
    container_url="https://example.com/path",
    sas_configuration=aws.datasync.LocationAzureBlobSasConfigurationArgs(
        token="sp=r&st=2023-12-20T14:54:52Z&se=2023-12-20T22:54:52Z&spr=https&sv=2021-06-08&sr=c&sig=aBBKDWQvyuVcTPH9EBp%2FXTI9E%2F%2Fmq171%2BZU178wcwqU%3D",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataSync.LocationAzureBlob("example", new()
    {
        AgentArns = new[]
        {
            aws_datasync_agent.Example.Arn,
        },
        AuthenticationType = "SAS",
        ContainerUrl = "https://example.com/path",
        SasConfiguration = new Aws.DataSync.Inputs.LocationAzureBlobSasConfigurationArgs
        {
            Token = "sp=r&st=2023-12-20T14:54:52Z&se=2023-12-20T22:54:52Z&spr=https&sv=2021-06-08&sr=c&sig=aBBKDWQvyuVcTPH9EBp%2FXTI9E%2F%2Fmq171%2BZU178wcwqU%3D",
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
		_, err := datasync.NewLocationAzureBlob(ctx, "example", &datasync.LocationAzureBlobArgs{
			AgentArns: pulumi.StringArray{
				aws_datasync_agent.Example.Arn,
			},
			AuthenticationType: pulumi.String("SAS"),
			ContainerUrl:       pulumi.String("https://example.com/path"),
			SasConfiguration: &datasync.LocationAzureBlobSasConfigurationArgs{
				Token: pulumi.String("sp=r&st=2023-12-20T14:54:52Z&se=2023-12-20T22:54:52Z&spr=https&sv=2021-06-08&sr=c&sig=aBBKDWQvyuVcTPH9EBp%2FXTI9E%2F%2Fmq171%2BZU178wcwqU%3D"),
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
import com.pulumi.aws.datasync.LocationAzureBlob;
import com.pulumi.aws.datasync.LocationAzureBlobArgs;
import com.pulumi.aws.datasync.inputs.LocationAzureBlobSasConfigurationArgs;
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
        var example = new LocationAzureBlob("example", LocationAzureBlobArgs.builder()        
            .agentArns(aws_datasync_agent.example().arn())
            .authenticationType("SAS")
            .containerUrl("https://example.com/path")
            .sasConfiguration(LocationAzureBlobSasConfigurationArgs.builder()
                .token("sp=r&st=2023-12-20T14:54:52Z&se=2023-12-20T22:54:52Z&spr=https&sv=2021-06-08&sr=c&sig=aBBKDWQvyuVcTPH9EBp%2FXTI9E%2F%2Fmq171%2BZU178wcwqU%3D")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:datasync:LocationAzureBlob
    properties:
      agentArns:
        - ${aws_datasync_agent.example.arn}
      authenticationType: SAS
      containerUrl: https://example.com/path
      sasConfiguration:
        token: sp=r&st=2023-12-20T14:54:52Z&se=2023-12-20T22:54:52Z&spr=https&sv=2021-06-08&sr=c&sig=aBBKDWQvyuVcTPH9EBp%2FXTI9E%2F%2Fmq171%2BZU178wcwqU%3D
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_location_azure_blob` using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:datasync/locationAzureBlob:LocationAzureBlob example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
```
 