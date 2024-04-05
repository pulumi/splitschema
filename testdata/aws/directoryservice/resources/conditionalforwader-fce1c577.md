Provides a conditional forwarder for managed Microsoft AD in AWS Directory Service.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directoryservice.ConditionalForwader("example", {
    directoryId: aws_directory_service_directory.ad.id,
    remoteDomainName: "example.com",
    dnsIps: [
        "8.8.8.8",
        "8.8.4.4",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directoryservice.ConditionalForwader("example",
    directory_id=aws_directory_service_directory["ad"]["id"],
    remote_domain_name="example.com",
    dns_ips=[
        "8.8.8.8",
        "8.8.4.4",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DirectoryService.ConditionalForwader("example", new()
    {
        DirectoryId = aws_directory_service_directory.Ad.Id,
        RemoteDomainName = "example.com",
        DnsIps = new[]
        {
            "8.8.8.8",
            "8.8.4.4",
        },
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
		_, err := directoryservice.NewConditionalForwader(ctx, "example", &directoryservice.ConditionalForwaderArgs{
			DirectoryId:      pulumi.Any(aws_directory_service_directory.Ad.Id),
			RemoteDomainName: pulumi.String("example.com"),
			DnsIps: pulumi.StringArray{
				pulumi.String("8.8.8.8"),
				pulumi.String("8.8.4.4"),
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
import com.pulumi.aws.directoryservice.ConditionalForwader;
import com.pulumi.aws.directoryservice.ConditionalForwaderArgs;
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
        var example = new ConditionalForwader("example", ConditionalForwaderArgs.builder()        
            .directoryId(aws_directory_service_directory.ad().id())
            .remoteDomainName("example.com")
            .dnsIps(            
                "8.8.8.8",
                "8.8.4.4")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:directoryservice:ConditionalForwader
    properties:
      directoryId: ${aws_directory_service_directory.ad.id}
      remoteDomainName: example.com
      dnsIps:
        - 8.8.8.8
        - 8.8.4.4
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import conditional forwarders using the directory id and remote_domain_name. For example:

```sh
 $ pulumi import aws:directoryservice/conditionalForwader:ConditionalForwader example d-1234567890:example.com
```
 