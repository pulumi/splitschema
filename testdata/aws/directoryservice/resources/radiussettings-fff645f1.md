Manages a directory's multi-factor authentication (MFA) using a Remote Authentication Dial In User Service (RADIUS) server.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directoryservice.RadiusSettings("example", {
    directoryId: aws_directory_service_directory.example.id,
    authenticationProtocol: "PAP",
    displayLabel: "example",
    radiusPort: 1812,
    radiusRetries: 4,
    radiusServers: ["10.0.1.5"],
    radiusTimeout: 1,
    sharedSecret: "12345678",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directoryservice.RadiusSettings("example",
    directory_id=aws_directory_service_directory["example"]["id"],
    authentication_protocol="PAP",
    display_label="example",
    radius_port=1812,
    radius_retries=4,
    radius_servers=["10.0.1.5"],
    radius_timeout=1,
    shared_secret="12345678")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DirectoryService.RadiusSettings("example", new()
    {
        DirectoryId = aws_directory_service_directory.Example.Id,
        AuthenticationProtocol = "PAP",
        DisplayLabel = "example",
        RadiusPort = 1812,
        RadiusRetries = 4,
        RadiusServers = new[]
        {
            "10.0.1.5",
        },
        RadiusTimeout = 1,
        SharedSecret = "12345678",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directoryservice.NewRadiusSettings(ctx, "example", &directoryservice.RadiusSettingsArgs{
			DirectoryId:            pulumi.Any(aws_directory_service_directory.Example.Id),
			AuthenticationProtocol: pulumi.String("PAP"),
			DisplayLabel:           pulumi.String("example"),
			RadiusPort:             pulumi.Int(1812),
			RadiusRetries:          pulumi.Int(4),
			RadiusServers: pulumi.StringArray{
				pulumi.String("10.0.1.5"),
			},
			RadiusTimeout: pulumi.Int(1),
			SharedSecret:  pulumi.String("12345678"),
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
import com.pulumi.aws.directoryservice.RadiusSettings;
import com.pulumi.aws.directoryservice.RadiusSettingsArgs;
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
        var example = new RadiusSettings("example", RadiusSettingsArgs.builder()        
            .directoryId(aws_directory_service_directory.example().id())
            .authenticationProtocol("PAP")
            .displayLabel("example")
            .radiusPort(1812)
            .radiusRetries(4)
            .radiusServers("10.0.1.5")
            .radiusTimeout(1)
            .sharedSecret("12345678")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:directoryservice:RadiusSettings
    properties:
      directoryId: ${aws_directory_service_directory.example.id}
      authenticationProtocol: PAP
      displayLabel: example
      radiusPort: 1812
      radiusRetries: 4
      radiusServers:
        - 10.0.1.5
      radiusTimeout: 1
      sharedSecret: '12345678'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import RADIUS settings using the directory ID. For example:

```sh
 $ pulumi import aws:directoryservice/radiusSettings:RadiusSettings example d-926724cf57
```
 