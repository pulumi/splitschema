Manages an AWS Opensearch Outbound Connection.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentCallerIdentity = aws.getCallerIdentity({});
const currentRegion = aws.getRegion({});
const foo = new aws.opensearch.OutboundConnection("foo", {
    connectionAlias: "outbound_connection",
    connectionMode: "DIRECT",
    localDomainInfo: {
        ownerId: currentCallerIdentity.then(currentCallerIdentity => currentCallerIdentity.accountId),
        region: currentRegion.then(currentRegion => currentRegion.name),
        domainName: aws_opensearch_domain.local_domain.domain_name,
    },
    remoteDomainInfo: {
        ownerId: currentCallerIdentity.then(currentCallerIdentity => currentCallerIdentity.accountId),
        region: currentRegion.then(currentRegion => currentRegion.name),
        domainName: aws_opensearch_domain.remote_domain.domain_name,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

current_caller_identity = aws.get_caller_identity()
current_region = aws.get_region()
foo = aws.opensearch.OutboundConnection("foo",
    connection_alias="outbound_connection",
    connection_mode="DIRECT",
    local_domain_info=aws.opensearch.OutboundConnectionLocalDomainInfoArgs(
        owner_id=current_caller_identity.account_id,
        region=current_region.name,
        domain_name=aws_opensearch_domain["local_domain"]["domain_name"],
    ),
    remote_domain_info=aws.opensearch.OutboundConnectionRemoteDomainInfoArgs(
        owner_id=current_caller_identity.account_id,
        region=current_region.name,
        domain_name=aws_opensearch_domain["remote_domain"]["domain_name"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var currentRegion = Aws.GetRegion.Invoke();

    var foo = new Aws.OpenSearch.OutboundConnection("foo", new()
    {
        ConnectionAlias = "outbound_connection",
        ConnectionMode = "DIRECT",
        LocalDomainInfo = new Aws.OpenSearch.Inputs.OutboundConnectionLocalDomainInfoArgs
        {
            OwnerId = currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
            Region = currentRegion.Apply(getRegionResult => getRegionResult.Name),
            DomainName = aws_opensearch_domain.Local_domain.Domain_name,
        },
        RemoteDomainInfo = new Aws.OpenSearch.Inputs.OutboundConnectionRemoteDomainInfoArgs
        {
            OwnerId = currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
            Region = currentRegion.Apply(getRegionResult => getRegionResult.Name),
            DomainName = aws_opensearch_domain.Remote_domain.Domain_name,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = opensearch.NewOutboundConnection(ctx, "foo", &opensearch.OutboundConnectionArgs{
			ConnectionAlias: pulumi.String("outbound_connection"),
			ConnectionMode:  pulumi.String("DIRECT"),
			LocalDomainInfo: &opensearch.OutboundConnectionLocalDomainInfoArgs{
				OwnerId:    *pulumi.String(currentCallerIdentity.AccountId),
				Region:     *pulumi.String(currentRegion.Name),
				DomainName: pulumi.Any(aws_opensearch_domain.Local_domain.Domain_name),
			},
			RemoteDomainInfo: &opensearch.OutboundConnectionRemoteDomainInfoArgs{
				OwnerId:    *pulumi.String(currentCallerIdentity.AccountId),
				Region:     *pulumi.String(currentRegion.Name),
				DomainName: pulumi.Any(aws_opensearch_domain.Remote_domain.Domain_name),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.opensearch.OutboundConnection;
import com.pulumi.aws.opensearch.OutboundConnectionArgs;
import com.pulumi.aws.opensearch.inputs.OutboundConnectionLocalDomainInfoArgs;
import com.pulumi.aws.opensearch.inputs.OutboundConnectionRemoteDomainInfoArgs;
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
        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        final var currentRegion = AwsFunctions.getRegion();

        var foo = new OutboundConnection("foo", OutboundConnectionArgs.builder()        
            .connectionAlias("outbound_connection")
            .connectionMode("DIRECT")
            .localDomainInfo(OutboundConnectionLocalDomainInfoArgs.builder()
                .ownerId(currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                .region(currentRegion.applyValue(getRegionResult -> getRegionResult.name()))
                .domainName(aws_opensearch_domain.local_domain().domain_name())
                .build())
            .remoteDomainInfo(OutboundConnectionRemoteDomainInfoArgs.builder()
                .ownerId(currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                .region(currentRegion.applyValue(getRegionResult -> getRegionResult.name()))
                .domainName(aws_opensearch_domain.remote_domain().domain_name())
                .build())
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:opensearch:OutboundConnection
    properties:
      connectionAlias: outbound_connection
      connectionMode: DIRECT
      localDomainInfo:
        ownerId: ${currentCallerIdentity.accountId}
        region: ${currentRegion.name}
        domainName: ${aws_opensearch_domain.local_domain.domain_name}
      remoteDomainInfo:
        ownerId: ${currentCallerIdentity.accountId}
        region: ${currentRegion.name}
        domainName: ${aws_opensearch_domain.remote_domain.domain_name}
variables:
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS Opensearch Outbound Connections using the Outbound Connection ID. For example:

```sh
 $ pulumi import aws:opensearch/outboundConnection:OutboundConnection foo connection-id
```
 