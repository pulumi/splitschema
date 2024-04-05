Generates a Core Network policy document in JSON format for use with resources that expect core network policy documents such as `awscc_networkmanager_core_network`. It follows the API definition from the [core-network-policy documentation](https://docs.aws.amazon.com/vpc/latest/cloudwan/cloudwan-policies-json.html).

Using this data source to generate policy documents is *optional*. It is also valid to use literal JSON strings in your configuration or to use the `file` interpolation function to read a raw JSON policy document from a file.



{{% examples %}}
## Example Usage
{{% example %}}
### Basic Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.networkmanager.getCoreNetworkPolicyDocument({
    attachmentPolicies: [
        {
            action: {
                associationMethod: "constant",
                segment: "shared",
            },
            conditionLogic: "or",
            conditions: [{
                key: "segment",
                operator: "equals",
                type: "tag-value",
                value: "shared",
            }],
            ruleNumber: 100,
        },
        {
            action: {
                associationMethod: "constant",
                segment: "prod",
            },
            conditionLogic: "or",
            conditions: [{
                key: "segment",
                operator: "equals",
                type: "tag-value",
                value: "prod",
            }],
            ruleNumber: 200,
        },
    ],
    coreNetworkConfigurations: [{
        asnRanges: ["64512-64555"],
        edgeLocations: [
            {
                asn: "64512",
                location: "us-east-1",
            },
            {
                asn: "64513",
                location: "eu-central-1",
            },
        ],
        vpnEcmpSupport: false,
    }],
    segmentActions: [{
        action: "share",
        mode: "attachment-route",
        segment: "shared",
        shareWiths: ["*"],
    }],
    segments: [
        {
            description: "Segment for shared services",
            name: "shared",
            requireAttachmentAcceptance: true,
        },
        {
            description: "Segment for prod services",
            name: "prod",
            requireAttachmentAcceptance: true,
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.networkmanager.get_core_network_policy_document(attachment_policies=[
        aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyArgs(
            action=aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs(
                association_method="constant",
                segment="shared",
            ),
            condition_logic="or",
            conditions=[aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs(
                key="segment",
                operator="equals",
                type="tag-value",
                value="shared",
            )],
            rule_number=100,
        ),
        aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyArgs(
            action=aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs(
                association_method="constant",
                segment="prod",
            ),
            condition_logic="or",
            conditions=[aws.networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs(
                key="segment",
                operator="equals",
                type="tag-value",
                value="prod",
            )],
            rule_number=200,
        ),
    ],
    core_network_configurations=[aws.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs(
        asn_ranges=["64512-64555"],
        edge_locations=[
            aws.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs(
                asn="64512",
                location="us-east-1",
            ),
            aws.networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs(
                asn="64513",
                location="eu-central-1",
            ),
        ],
        vpn_ecmp_support=False,
    )],
    segment_actions=[aws.networkmanager.GetCoreNetworkPolicyDocumentSegmentActionArgs(
        action="share",
        mode="attachment-route",
        segment="shared",
        share_withs=["*"],
    )],
    segments=[
        aws.networkmanager.GetCoreNetworkPolicyDocumentSegmentArgs(
            description="Segment for shared services",
            name="shared",
            require_attachment_acceptance=True,
        ),
        aws.networkmanager.GetCoreNetworkPolicyDocumentSegmentArgs(
            description="Segment for prod services",
            name="prod",
            require_attachment_acceptance=True,
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.NetworkManager.GetCoreNetworkPolicyDocument.Invoke(new()
    {
        AttachmentPolicies = new[]
        {
            new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentAttachmentPolicyInputArgs
            {
                Action = new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentAttachmentPolicyActionInputArgs
                {
                    AssociationMethod = "constant",
                    Segment = "shared",
                },
                ConditionLogic = "or",
                Conditions = new[]
                {
                    new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentAttachmentPolicyConditionInputArgs
                    {
                        Key = "segment",
                        Operator = "equals",
                        Type = "tag-value",
                        Value = "shared",
                    },
                },
                RuleNumber = 100,
            },
            new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentAttachmentPolicyInputArgs
            {
                Action = new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentAttachmentPolicyActionInputArgs
                {
                    AssociationMethod = "constant",
                    Segment = "prod",
                },
                ConditionLogic = "or",
                Conditions = new[]
                {
                    new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentAttachmentPolicyConditionInputArgs
                    {
                        Key = "segment",
                        Operator = "equals",
                        Type = "tag-value",
                        Value = "prod",
                    },
                },
                RuleNumber = 200,
            },
        },
        CoreNetworkConfigurations = new[]
        {
            new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationInputArgs
            {
                AsnRanges = new[]
                {
                    "64512-64555",
                },
                EdgeLocations = new[]
                {
                    new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationInputArgs
                    {
                        Asn = "64512",
                        Location = "us-east-1",
                    },
                    new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationInputArgs
                    {
                        Asn = "64513",
                        Location = "eu-central-1",
                    },
                },
                VpnEcmpSupport = false,
            },
        },
        SegmentActions = new[]
        {
            new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentSegmentActionInputArgs
            {
                Action = "share",
                Mode = "attachment-route",
                Segment = "shared",
                ShareWiths = new[]
                {
                    "*",
                },
            },
        },
        Segments = new[]
        {
            new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentSegmentInputArgs
            {
                Description = "Segment for shared services",
                Name = "shared",
                RequireAttachmentAcceptance = true,
            },
            new Aws.NetworkManager.Inputs.GetCoreNetworkPolicyDocumentSegmentInputArgs
            {
                Description = "Segment for prod services",
                Name = "prod",
                RequireAttachmentAcceptance = true,
            },
        },
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
		_, err := networkmanager.GetCoreNetworkPolicyDocument(ctx, &networkmanager.GetCoreNetworkPolicyDocumentArgs{
			AttachmentPolicies: []networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicy{
				{
					Action: {
						AssociationMethod: "constant",
						Segment:           pulumi.StringRef("shared"),
					},
					ConditionLogic: pulumi.StringRef("or"),
					Conditions: []networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyCondition{
						{
							Key:      pulumi.StringRef("segment"),
							Operator: pulumi.StringRef("equals"),
							Type:     "tag-value",
							Value:    pulumi.StringRef("shared"),
						},
					},
					RuleNumber: 100,
				},
				{
					Action: {
						AssociationMethod: "constant",
						Segment:           pulumi.StringRef("prod"),
					},
					ConditionLogic: pulumi.StringRef("or"),
					Conditions: []networkmanager.GetCoreNetworkPolicyDocumentAttachmentPolicyCondition{
						{
							Key:      pulumi.StringRef("segment"),
							Operator: pulumi.StringRef("equals"),
							Type:     "tag-value",
							Value:    pulumi.StringRef("prod"),
						},
					},
					RuleNumber: 200,
				},
			},
			CoreNetworkConfigurations: []networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfiguration{
				{
					AsnRanges: []string{
						"64512-64555",
					},
					EdgeLocations: []networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocation{
						{
							Asn:      pulumi.StringRef("64512"),
							Location: "us-east-1",
						},
						{
							Asn:      pulumi.StringRef("64513"),
							Location: "eu-central-1",
						},
					},
					VpnEcmpSupport: pulumi.BoolRef(false),
				},
			},
			SegmentActions: []networkmanager.GetCoreNetworkPolicyDocumentSegmentAction{
				{
					Action:  "share",
					Mode:    pulumi.StringRef("attachment-route"),
					Segment: "shared",
					ShareWiths: []string{
						"*",
					},
				},
			},
			Segments: []networkmanager.GetCoreNetworkPolicyDocumentSegment{
				{
					Description:                 pulumi.StringRef("Segment for shared services"),
					Name:                        "shared",
					RequireAttachmentAcceptance: pulumi.BoolRef(true),
				},
				{
					Description:                 pulumi.StringRef("Segment for prod services"),
					Name:                        "prod",
					RequireAttachmentAcceptance: pulumi.BoolRef(true),
				},
			},
		}, nil)
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
import com.pulumi.aws.networkmanager.NetworkmanagerFunctions;
import com.pulumi.aws.networkmanager.inputs.GetCoreNetworkPolicyDocumentArgs;
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
        final var test = NetworkmanagerFunctions.getCoreNetworkPolicyDocument(GetCoreNetworkPolicyDocumentArgs.builder()
            .attachmentPolicies(            
                GetCoreNetworkPolicyDocumentAttachmentPolicyArgs.builder()
                    .action(GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs.builder()
                        .associationMethod("constant")
                        .segment("shared")
                        .build())
                    .conditionLogic("or")
                    .conditions(GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs.builder()
                        .key("segment")
                        .operator("equals")
                        .type("tag-value")
                        .value("shared")
                        .build())
                    .ruleNumber(100)
                    .build(),
                GetCoreNetworkPolicyDocumentAttachmentPolicyArgs.builder()
                    .action(GetCoreNetworkPolicyDocumentAttachmentPolicyActionArgs.builder()
                        .associationMethod("constant")
                        .segment("prod")
                        .build())
                    .conditionLogic("or")
                    .conditions(GetCoreNetworkPolicyDocumentAttachmentPolicyConditionArgs.builder()
                        .key("segment")
                        .operator("equals")
                        .type("tag-value")
                        .value("prod")
                        .build())
                    .ruleNumber(200)
                    .build())
            .coreNetworkConfigurations(GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs.builder()
                .asnRanges("64512-64555")
                .edgeLocations(                
                    GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs.builder()
                        .asn(64512)
                        .location("us-east-1")
                        .build(),
                    GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs.builder()
                        .asn(64513)
                        .location("eu-central-1")
                        .build())
                .vpnEcmpSupport(false)
                .build())
            .segmentActions(GetCoreNetworkPolicyDocumentSegmentActionArgs.builder()
                .action("share")
                .mode("attachment-route")
                .segment("shared")
                .shareWiths("*")
                .build())
            .segments(            
                GetCoreNetworkPolicyDocumentSegmentArgs.builder()
                    .description("Segment for shared services")
                    .name("shared")
                    .requireAttachmentAcceptance(true)
                    .build(),
                GetCoreNetworkPolicyDocumentSegmentArgs.builder()
                    .description("Segment for prod services")
                    .name("prod")
                    .requireAttachmentAcceptance(true)
                    .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:networkmanager:getCoreNetworkPolicyDocument
      Arguments:
        attachmentPolicies:
          - action:
              associationMethod: constant
              segment: shared
            conditionLogic: or
            conditions:
              - key: segment
                operator: equals
                type: tag-value
                value: shared
            ruleNumber: 100
          - action:
              associationMethod: constant
              segment: prod
            conditionLogic: or
            conditions:
              - key: segment
                operator: equals
                type: tag-value
                value: prod
            ruleNumber: 200
        coreNetworkConfigurations:
          - asnRanges:
              - 64512-64555
            edgeLocations:
              - asn: 64512
                location: us-east-1
              - asn: 64513
                location: eu-central-1
            vpnEcmpSupport: false
        segmentActions:
          - action: share
            mode: attachment-route
            segment: shared
            shareWiths:
              - '*'
        segments:
          - description: Segment for shared services
            name: shared
            requireAttachmentAcceptance: true
          - description: Segment for prod services
            name: prod
            requireAttachmentAcceptance: true
```

`data.aws_networkmanager_core_network_policy_document.test.json` will evaluate to:

```typescript
import * as pulumi from "@pulumi/pulumi";
```
```python
import pulumi
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() => 
{
});
```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
    }
}
```
```yaml
{}
```
{{% /example %}}
{{% /examples %}}