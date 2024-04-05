Resource for managing an AWS FinSpace Kx Environment.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {
    description: "Sample KMS Key",
    deletionWindowInDays: 7,
});
const exampleKxEnvironment = new aws.finspace.KxEnvironment("exampleKxEnvironment", {kmsKeyId: exampleKey.arn});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey",
    description="Sample KMS Key",
    deletion_window_in_days=7)
example_kx_environment = aws.finspace.KxEnvironment("exampleKxEnvironment", kms_key_id=example_key.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "Sample KMS Key",
        DeletionWindowInDays = 7,
    });

    var exampleKxEnvironment = new Aws.FinSpace.KxEnvironment("exampleKxEnvironment", new()
    {
        KmsKeyId = exampleKey.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/finspace"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description:          pulumi.String("Sample KMS Key"),
			DeletionWindowInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		_, err = finspace.NewKxEnvironment(ctx, "exampleKxEnvironment", &finspace.KxEnvironmentArgs{
			KmsKeyId: exampleKey.Arn,
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.finspace.KxEnvironment;
import com.pulumi.aws.finspace.KxEnvironmentArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("Sample KMS Key")
            .deletionWindowInDays(7)
            .build());

        var exampleKxEnvironment = new KxEnvironment("exampleKxEnvironment", KxEnvironmentArgs.builder()        
            .kmsKeyId(exampleKey.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: Sample KMS Key
      deletionWindowInDays: 7
  exampleKxEnvironment:
    type: aws:finspace:KxEnvironment
    properties:
      kmsKeyId: ${exampleKey.arn}
```
{{% /example %}}
{{% example %}}
### With Transit Gateway Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {
    description: "Sample KMS Key",
    deletionWindowInDays: 7,
});
const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("exampleTransitGateway", {description: "example"});
const exampleEnv = new aws.finspace.KxEnvironment("exampleEnv", {
    description: "Environment description",
    kmsKeyId: exampleKey.arn,
    transitGatewayConfiguration: {
        transitGatewayId: exampleTransitGateway.id,
        routableCidrSpace: "100.64.0.0/26",
    },
    customDnsConfigurations: [{
        customDnsServerName: "example.finspace.amazonaws.com",
        customDnsServerIp: "10.0.0.76",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey",
    description="Sample KMS Key",
    deletion_window_in_days=7)
example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway", description="example")
example_env = aws.finspace.KxEnvironment("exampleEnv",
    description="Environment description",
    kms_key_id=example_key.arn,
    transit_gateway_configuration=aws.finspace.KxEnvironmentTransitGatewayConfigurationArgs(
        transit_gateway_id=example_transit_gateway.id,
        routable_cidr_space="100.64.0.0/26",
    ),
    custom_dns_configurations=[aws.finspace.KxEnvironmentCustomDnsConfigurationArgs(
        custom_dns_server_name="example.finspace.amazonaws.com",
        custom_dns_server_ip="10.0.0.76",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "Sample KMS Key",
        DeletionWindowInDays = 7,
    });

    var exampleTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("exampleTransitGateway", new()
    {
        Description = "example",
    });

    var exampleEnv = new Aws.FinSpace.KxEnvironment("exampleEnv", new()
    {
        Description = "Environment description",
        KmsKeyId = exampleKey.Arn,
        TransitGatewayConfiguration = new Aws.FinSpace.Inputs.KxEnvironmentTransitGatewayConfigurationArgs
        {
            TransitGatewayId = exampleTransitGateway.Id,
            RoutableCidrSpace = "100.64.0.0/26",
        },
        CustomDnsConfigurations = new[]
        {
            new Aws.FinSpace.Inputs.KxEnvironmentCustomDnsConfigurationArgs
            {
                CustomDnsServerName = "example.finspace.amazonaws.com",
                CustomDnsServerIp = "10.0.0.76",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/finspace"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description:          pulumi.String("Sample KMS Key"),
			DeletionWindowInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", &ec2transitgateway.TransitGatewayArgs{
			Description: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = finspace.NewKxEnvironment(ctx, "exampleEnv", &finspace.KxEnvironmentArgs{
			Description: pulumi.String("Environment description"),
			KmsKeyId:    exampleKey.Arn,
			TransitGatewayConfiguration: &finspace.KxEnvironmentTransitGatewayConfigurationArgs{
				TransitGatewayId:  exampleTransitGateway.ID(),
				RoutableCidrSpace: pulumi.String("100.64.0.0/26"),
			},
			CustomDnsConfigurations: finspace.KxEnvironmentCustomDnsConfigurationArray{
				&finspace.KxEnvironmentCustomDnsConfigurationArgs{
					CustomDnsServerName: pulumi.String("example.finspace.amazonaws.com"),
					CustomDnsServerIp:   pulumi.String("10.0.0.76"),
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.ec2transitgateway.TransitGateway;
import com.pulumi.aws.ec2transitgateway.TransitGatewayArgs;
import com.pulumi.aws.finspace.KxEnvironment;
import com.pulumi.aws.finspace.KxEnvironmentArgs;
import com.pulumi.aws.finspace.inputs.KxEnvironmentTransitGatewayConfigurationArgs;
import com.pulumi.aws.finspace.inputs.KxEnvironmentCustomDnsConfigurationArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("Sample KMS Key")
            .deletionWindowInDays(7)
            .build());

        var exampleTransitGateway = new TransitGateway("exampleTransitGateway", TransitGatewayArgs.builder()        
            .description("example")
            .build());

        var exampleEnv = new KxEnvironment("exampleEnv", KxEnvironmentArgs.builder()        
            .description("Environment description")
            .kmsKeyId(exampleKey.arn())
            .transitGatewayConfiguration(KxEnvironmentTransitGatewayConfigurationArgs.builder()
                .transitGatewayId(exampleTransitGateway.id())
                .routableCidrSpace("100.64.0.0/26")
                .build())
            .customDnsConfigurations(KxEnvironmentCustomDnsConfigurationArgs.builder()
                .customDnsServerName("example.finspace.amazonaws.com")
                .customDnsServerIp("10.0.0.76")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: Sample KMS Key
      deletionWindowInDays: 7
  exampleTransitGateway:
    type: aws:ec2transitgateway:TransitGateway
    properties:
      description: example
  exampleEnv:
    type: aws:finspace:KxEnvironment
    properties:
      description: Environment description
      kmsKeyId: ${exampleKey.arn}
      transitGatewayConfiguration:
        transitGatewayId: ${exampleTransitGateway.id}
        routableCidrSpace: 100.64.0.0/26
      customDnsConfigurations:
        - customDnsServerName: example.finspace.amazonaws.com
          customDnsServerIp: 10.0.0.76
```
{{% /example %}}
{{% example %}}
### With Transit Gateway Attachment Network ACL Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {
    description: "Sample KMS Key",
    deletionWindowInDays: 7,
});
const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("exampleTransitGateway", {description: "example"});
const exampleEnv = new aws.finspace.KxEnvironment("exampleEnv", {
    description: "Environment description",
    kmsKeyId: exampleKey.arn,
    transitGatewayConfiguration: {
        transitGatewayId: exampleTransitGateway.id,
        routableCidrSpace: "100.64.0.0/26",
        attachmentNetworkAclConfigurations: [{
            ruleNumber: 1,
            protocol: "6",
            ruleAction: "allow",
            cidrBlock: "0.0.0.0/0",
            portRange: {
                from: 53,
                to: 53,
            },
            icmpTypeCode: {
                type: -1,
                code: -1,
            },
        }],
    },
    customDnsConfigurations: [{
        customDnsServerName: "example.finspace.amazonaws.com",
        customDnsServerIp: "10.0.0.76",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey",
    description="Sample KMS Key",
    deletion_window_in_days=7)
example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway", description="example")
example_env = aws.finspace.KxEnvironment("exampleEnv",
    description="Environment description",
    kms_key_id=example_key.arn,
    transit_gateway_configuration=aws.finspace.KxEnvironmentTransitGatewayConfigurationArgs(
        transit_gateway_id=example_transit_gateway.id,
        routable_cidr_space="100.64.0.0/26",
        attachment_network_acl_configurations=[aws.finspace.KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationArgs(
            rule_number=1,
            protocol="6",
            rule_action="allow",
            cidr_block="0.0.0.0/0",
            port_range=aws.finspace.KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationPortRangeArgs(
                from_=53,
                to=53,
            ),
            icmp_type_code=aws.finspace.KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationIcmpTypeCodeArgs(
                type=-1,
                code=-1,
            ),
        )],
    ),
    custom_dns_configurations=[aws.finspace.KxEnvironmentCustomDnsConfigurationArgs(
        custom_dns_server_name="example.finspace.amazonaws.com",
        custom_dns_server_ip="10.0.0.76",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "Sample KMS Key",
        DeletionWindowInDays = 7,
    });

    var exampleTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("exampleTransitGateway", new()
    {
        Description = "example",
    });

    var exampleEnv = new Aws.FinSpace.KxEnvironment("exampleEnv", new()
    {
        Description = "Environment description",
        KmsKeyId = exampleKey.Arn,
        TransitGatewayConfiguration = new Aws.FinSpace.Inputs.KxEnvironmentTransitGatewayConfigurationArgs
        {
            TransitGatewayId = exampleTransitGateway.Id,
            RoutableCidrSpace = "100.64.0.0/26",
            AttachmentNetworkAclConfigurations = new[]
            {
                new Aws.FinSpace.Inputs.KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationArgs
                {
                    RuleNumber = 1,
                    Protocol = "6",
                    RuleAction = "allow",
                    CidrBlock = "0.0.0.0/0",
                    PortRange = new Aws.FinSpace.Inputs.KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationPortRangeArgs
                    {
                        From = 53,
                        To = 53,
                    },
                    IcmpTypeCode = new Aws.FinSpace.Inputs.KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationIcmpTypeCodeArgs
                    {
                        Type = -1,
                        Code = -1,
                    },
                },
            },
        },
        CustomDnsConfigurations = new[]
        {
            new Aws.FinSpace.Inputs.KxEnvironmentCustomDnsConfigurationArgs
            {
                CustomDnsServerName = "example.finspace.amazonaws.com",
                CustomDnsServerIp = "10.0.0.76",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/finspace"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description:          pulumi.String("Sample KMS Key"),
			DeletionWindowInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", &ec2transitgateway.TransitGatewayArgs{
			Description: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = finspace.NewKxEnvironment(ctx, "exampleEnv", &finspace.KxEnvironmentArgs{
			Description: pulumi.String("Environment description"),
			KmsKeyId:    exampleKey.Arn,
			TransitGatewayConfiguration: &finspace.KxEnvironmentTransitGatewayConfigurationArgs{
				TransitGatewayId:  exampleTransitGateway.ID(),
				RoutableCidrSpace: pulumi.String("100.64.0.0/26"),
				AttachmentNetworkAclConfigurations: finspace.KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationArray{
					&finspace.KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationArgs{
						RuleNumber: pulumi.Int(1),
						Protocol:   pulumi.String("6"),
						RuleAction: pulumi.String("allow"),
						CidrBlock:  pulumi.String("0.0.0.0/0"),
						PortRange: &finspace.KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationPortRangeArgs{
							From: pulumi.Int(53),
							To:   pulumi.Int(53),
						},
						IcmpTypeCode: &finspace.KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationIcmpTypeCodeArgs{
							Type: -1,
							Code: -1,
						},
					},
				},
			},
			CustomDnsConfigurations: finspace.KxEnvironmentCustomDnsConfigurationArray{
				&finspace.KxEnvironmentCustomDnsConfigurationArgs{
					CustomDnsServerName: pulumi.String("example.finspace.amazonaws.com"),
					CustomDnsServerIp:   pulumi.String("10.0.0.76"),
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
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an AWS FinSpace Kx Environment using the `id`. For example:

```sh
 $ pulumi import aws:finspace/kxEnvironment:KxEnvironment example n3ceo7wqxoxcti5tujqwzs
```
 