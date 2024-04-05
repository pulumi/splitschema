Provides a VPC Endpoint connection notification resource.
Connection notifications notify subscribers of VPC Endpoint events.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const topicPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["vpce.amazonaws.com"],
        }],
        actions: ["SNS:Publish"],
        resources: ["arn:aws:sns:*:*:vpce-notification-topic"],
    }],
});
const topicTopic = new aws.sns.Topic("topicTopic", {policy: topicPolicyDocument.then(topicPolicyDocument => topicPolicyDocument.json)});
const fooVpcEndpointService = new aws.ec2.VpcEndpointService("fooVpcEndpointService", {
    acceptanceRequired: false,
    networkLoadBalancerArns: [aws_lb.test.arn],
});
const fooVpcEndpointConnectionNotification = new aws.ec2.VpcEndpointConnectionNotification("fooVpcEndpointConnectionNotification", {
    vpcEndpointServiceId: fooVpcEndpointService.id,
    connectionNotificationArn: topicTopic.arn,
    connectionEvents: [
        "Accept",
        "Reject",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

topic_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["vpce.amazonaws.com"],
    )],
    actions=["SNS:Publish"],
    resources=["arn:aws:sns:*:*:vpce-notification-topic"],
)])
topic_topic = aws.sns.Topic("topicTopic", policy=topic_policy_document.json)
foo_vpc_endpoint_service = aws.ec2.VpcEndpointService("fooVpcEndpointService",
    acceptance_required=False,
    network_load_balancer_arns=[aws_lb["test"]["arn"]])
foo_vpc_endpoint_connection_notification = aws.ec2.VpcEndpointConnectionNotification("fooVpcEndpointConnectionNotification",
    vpc_endpoint_service_id=foo_vpc_endpoint_service.id,
    connection_notification_arn=topic_topic.arn,
    connection_events=[
        "Accept",
        "Reject",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var topicPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "vpce.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "SNS:Publish",
                },
                Resources = new[]
                {
                    "arn:aws:sns:*:*:vpce-notification-topic",
                },
            },
        },
    });

    var topicTopic = new Aws.Sns.Topic("topicTopic", new()
    {
        Policy = topicPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var fooVpcEndpointService = new Aws.Ec2.VpcEndpointService("fooVpcEndpointService", new()
    {
        AcceptanceRequired = false,
        NetworkLoadBalancerArns = new[]
        {
            aws_lb.Test.Arn,
        },
    });

    var fooVpcEndpointConnectionNotification = new Aws.Ec2.VpcEndpointConnectionNotification("fooVpcEndpointConnectionNotification", new()
    {
        VpcEndpointServiceId = fooVpcEndpointService.Id,
        ConnectionNotificationArn = topicTopic.Arn,
        ConnectionEvents = new[]
        {
            "Accept",
            "Reject",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		topicPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"vpce.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"SNS:Publish",
					},
					Resources: []string{
						"arn:aws:sns:*:*:vpce-notification-topic",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		topicTopic, err := sns.NewTopic(ctx, "topicTopic", &sns.TopicArgs{
			Policy: *pulumi.String(topicPolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		fooVpcEndpointService, err := ec2.NewVpcEndpointService(ctx, "fooVpcEndpointService", &ec2.VpcEndpointServiceArgs{
			AcceptanceRequired: pulumi.Bool(false),
			NetworkLoadBalancerArns: pulumi.StringArray{
				aws_lb.Test.Arn,
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcEndpointConnectionNotification(ctx, "fooVpcEndpointConnectionNotification", &ec2.VpcEndpointConnectionNotificationArgs{
			VpcEndpointServiceId:      fooVpcEndpointService.ID(),
			ConnectionNotificationArn: topicTopic.Arn,
			ConnectionEvents: pulumi.StringArray{
				pulumi.String("Accept"),
				pulumi.String("Reject"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.sns.TopicArgs;
import com.pulumi.aws.ec2.VpcEndpointService;
import com.pulumi.aws.ec2.VpcEndpointServiceArgs;
import com.pulumi.aws.ec2.VpcEndpointConnectionNotification;
import com.pulumi.aws.ec2.VpcEndpointConnectionNotificationArgs;
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
        final var topicPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("vpce.amazonaws.com")
                    .build())
                .actions("SNS:Publish")
                .resources("arn:aws:sns:*:*:vpce-notification-topic")
                .build())
            .build());

        var topicTopic = new Topic("topicTopic", TopicArgs.builder()        
            .policy(topicPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var fooVpcEndpointService = new VpcEndpointService("fooVpcEndpointService", VpcEndpointServiceArgs.builder()        
            .acceptanceRequired(false)
            .networkLoadBalancerArns(aws_lb.test().arn())
            .build());

        var fooVpcEndpointConnectionNotification = new VpcEndpointConnectionNotification("fooVpcEndpointConnectionNotification", VpcEndpointConnectionNotificationArgs.builder()        
            .vpcEndpointServiceId(fooVpcEndpointService.id())
            .connectionNotificationArn(topicTopic.arn())
            .connectionEvents(            
                "Accept",
                "Reject")
            .build());

    }
}
```
```yaml
resources:
  topicTopic:
    type: aws:sns:Topic
    properties:
      policy: ${topicPolicyDocument.json}
  fooVpcEndpointService:
    type: aws:ec2:VpcEndpointService
    properties:
      acceptanceRequired: false
      networkLoadBalancerArns:
        - ${aws_lb.test.arn}
  fooVpcEndpointConnectionNotification:
    type: aws:ec2:VpcEndpointConnectionNotification
    properties:
      vpcEndpointServiceId: ${fooVpcEndpointService.id}
      connectionNotificationArn: ${topicTopic.arn}
      connectionEvents:
        - Accept
        - Reject
variables:
  topicPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - vpce.amazonaws.com
            actions:
              - SNS:Publish
            resources:
              - arn:aws:sns:*:*:vpce-notification-topic
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Endpoint connection notifications using the VPC endpoint connection notification `id`. For example:

```sh
 $ pulumi import aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification foo vpce-nfn-09e6ed3b4efba2263
```
 