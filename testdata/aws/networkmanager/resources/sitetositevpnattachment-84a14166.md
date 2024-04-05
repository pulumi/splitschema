Resource for managing an AWS Network Manager SiteToSiteAttachment.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkmanager.SiteToSiteVpnAttachment("example", {
    coreNetworkId: awscc_networkmanager_core_network.example.id,
    vpnConnectionArn: aws_vpn_connection.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkmanager.SiteToSiteVpnAttachment("example",
    core_network_id=awscc_networkmanager_core_network["example"]["id"],
    vpn_connection_arn=aws_vpn_connection["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkManager.SiteToSiteVpnAttachment("example", new()
    {
        CoreNetworkId = awscc_networkmanager_core_network.Example.Id,
        VpnConnectionArn = aws_vpn_connection.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkmanager.NewSiteToSiteVpnAttachment(ctx, "example", &networkmanager.SiteToSiteVpnAttachmentArgs{
			CoreNetworkId:    pulumi.Any(awscc_networkmanager_core_network.Example.Id),
			VpnConnectionArn: pulumi.Any(aws_vpn_connection.Example.Arn),
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
import com.pulumi.aws.networkmanager.SiteToSiteVpnAttachment;
import com.pulumi.aws.networkmanager.SiteToSiteVpnAttachmentArgs;
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
        var example = new SiteToSiteVpnAttachment("example", SiteToSiteVpnAttachmentArgs.builder()        
            .coreNetworkId(awscc_networkmanager_core_network.example().id())
            .vpnConnectionArn(aws_vpn_connection.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkmanager:SiteToSiteVpnAttachment
    properties:
      coreNetworkId: ${awscc_networkmanager_core_network.example.id}
      vpnConnectionArn: ${aws_vpn_connection.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_networkmanager_site_to_site_vpn_attachment` using the attachment ID. For example:

```sh
 $ pulumi import aws:networkmanager/siteToSiteVpnAttachment:SiteToSiteVpnAttachment example attachment-0f8fa60d2238d1bd8
```
 