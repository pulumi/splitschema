Manages an [AWS Opensearch Inbound Connection Accepter](https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_AcceptInboundConnection.html). If connecting domains from different AWS accounts, ensure that the accepter is configured to use the AWS account where the _remote_ opensearch domain exists.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentCallerIdentity = aws.getCallerIdentity({});
const currentRegion = aws.getRegion({});
const fooOutboundConnection = new aws.opensearch.OutboundConnection("fooOutboundConnection", {
    connectionAlias: "outbound_connection",
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
const fooInboundConnectionAccepter = new aws.opensearch.InboundConnectionAccepter("fooInboundConnectionAccepter", {connectionId: fooOutboundConnection.id});
```
```python
import pulumi
import pulumi_aws as aws

current_caller_identity = aws.get_caller_identity()
current_region = aws.get_region()
foo_outbound_connection = aws.opensearch.OutboundConnection("fooOutboundConnection",
    connection_alias="outbound_connection",
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
foo_inbound_connection_accepter = aws.opensearch.InboundConnectionAccepter("fooInboundConnectionAccepter", connection_id=foo_outbound_connection.id)
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

    var fooOutboundConnection = new Aws.OpenSearch.OutboundConnection("fooOutboundConnection", new()
    {
        ConnectionAlias = "outbound_connection",
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

    var fooInboundConnectionAccepter = new Aws.OpenSearch.InboundConnectionAccepter("fooInboundConnectionAccepter", new()
    {
        ConnectionId = fooOutboundConnection.Id,
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
		fooOutboundConnection, err := opensearch.NewOutboundConnection(ctx, "fooOutboundConnection", &opensearch.OutboundConnectionArgs{
			ConnectionAlias: pulumi.String("outbound_connection"),
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
		_, err = opensearch.NewInboundConnectionAccepter(ctx, "fooInboundConnectionAccepter", &opensearch.InboundConnectionAccepterArgs{
			ConnectionId: fooOutboundConnection.ID(),
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
import com.pulumi.aws.opensearch.InboundConnectionAccepter;
import com.pulumi.aws.opensearch.InboundConnectionAccepterArgs;
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

        var fooOutboundConnection = new OutboundConnection("fooOutboundConnection", OutboundConnectionArgs.builder()        
            .connectionAlias("outbound_connection")
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

        var fooInboundConnectionAccepter = new InboundConnectionAccepter("fooInboundConnectionAccepter", InboundConnectionAccepterArgs.builder()        
            .connectionId(fooOutboundConnection.id())
            .build());

    }
}
```
```yaml
resources:
  fooOutboundConnection:
    type: aws:opensearch:OutboundConnection
    properties:
      connectionAlias: outbound_connection
      localDomainInfo:
        ownerId: ${currentCallerIdentity.accountId}
        region: ${currentRegion.name}
        domainName: ${aws_opensearch_domain.local_domain.domain_name}
      remoteDomainInfo:
        ownerId: ${currentCallerIdentity.accountId}
        region: ${currentRegion.name}
        domainName: ${aws_opensearch_domain.remote_domain.domain_name}
  fooInboundConnectionAccepter:
    type: aws:opensearch:InboundConnectionAccepter
    properties:
      connectionId: ${fooOutboundConnection.id}
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

Using `pulumi import`, import AWS Opensearch Inbound Connection Accepters using the Inbound Connection ID. For example:

```sh
 $ pulumi import aws:opensearch/inboundConnectionAccepter:InboundConnectionAccepter foo connection-id
```
 