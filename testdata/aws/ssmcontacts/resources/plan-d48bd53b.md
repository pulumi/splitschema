Resource for managing an AWS SSM Contact Plan.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssmcontacts.Plan("example", {
    contactId: "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
    stages: [{
        durationInMinutes: 1,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssmcontacts.Plan("example",
    contact_id="arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
    stages=[aws.ssmcontacts.PlanStageArgs(
        duration_in_minutes=1,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SsmContacts.Plan("example", new()
    {
        ContactId = "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
        Stages = new[]
        {
            new Aws.SsmContacts.Inputs.PlanStageArgs
            {
                DurationInMinutes = 1,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmcontacts"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssmcontacts.NewPlan(ctx, "example", &ssmcontacts.PlanArgs{
			ContactId: pulumi.String("arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias"),
			Stages: ssmcontacts.PlanStageArray{
				&ssmcontacts.PlanStageArgs{
					DurationInMinutes: pulumi.Int(1),
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
import com.pulumi.aws.ssmcontacts.Plan;
import com.pulumi.aws.ssmcontacts.PlanArgs;
import com.pulumi.aws.ssmcontacts.inputs.PlanStageArgs;
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
        var example = new Plan("example", PlanArgs.builder()        
            .contactId("arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias")
            .stages(PlanStageArgs.builder()
                .durationInMinutes(1)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssmcontacts:Plan
    properties:
      contactId: arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias
      stages:
        - durationInMinutes: 1
```
{{% /example %}}
{{% example %}}
### Usage with SSM Contact

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const contact = new aws.ssmcontacts.Contact("contact", {
    alias: "alias",
    type: "PERSONAL",
});
const plan = new aws.ssmcontacts.Plan("plan", {
    contactId: contact.arn,
    stages: [{
        durationInMinutes: 1,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

contact = aws.ssmcontacts.Contact("contact",
    alias="alias",
    type="PERSONAL")
plan = aws.ssmcontacts.Plan("plan",
    contact_id=contact.arn,
    stages=[aws.ssmcontacts.PlanStageArgs(
        duration_in_minutes=1,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var contact = new Aws.SsmContacts.Contact("contact", new()
    {
        Alias = "alias",
        Type = "PERSONAL",
    });

    var plan = new Aws.SsmContacts.Plan("plan", new()
    {
        ContactId = contact.Arn,
        Stages = new[]
        {
            new Aws.SsmContacts.Inputs.PlanStageArgs
            {
                DurationInMinutes = 1,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmcontacts"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		contact, err := ssmcontacts.NewContact(ctx, "contact", &ssmcontacts.ContactArgs{
			Alias: pulumi.String("alias"),
			Type:  pulumi.String("PERSONAL"),
		})
		if err != nil {
			return err
		}
		_, err = ssmcontacts.NewPlan(ctx, "plan", &ssmcontacts.PlanArgs{
			ContactId: contact.Arn,
			Stages: ssmcontacts.PlanStageArray{
				&ssmcontacts.PlanStageArgs{
					DurationInMinutes: pulumi.Int(1),
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
import com.pulumi.aws.ssmcontacts.Contact;
import com.pulumi.aws.ssmcontacts.ContactArgs;
import com.pulumi.aws.ssmcontacts.Plan;
import com.pulumi.aws.ssmcontacts.PlanArgs;
import com.pulumi.aws.ssmcontacts.inputs.PlanStageArgs;
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
        var contact = new Contact("contact", ContactArgs.builder()        
            .alias("alias")
            .type("PERSONAL")
            .build());

        var plan = new Plan("plan", PlanArgs.builder()        
            .contactId(contact.arn())
            .stages(PlanStageArgs.builder()
                .durationInMinutes(1)
                .build())
            .build());

    }
}
```
```yaml
resources:
  contact:
    type: aws:ssmcontacts:Contact
    properties:
      alias: alias
      type: PERSONAL
  plan:
    type: aws:ssmcontacts:Plan
    properties:
      contactId: ${contact.arn}
      stages:
        - durationInMinutes: 1
```
{{% /example %}}
{{% example %}}
### Usage With All Fields

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const escalationPlan = new aws.ssmcontacts.Contact("escalationPlan", {
    alias: "escalation-plan-alias",
    type: "ESCALATION",
});
const contactOne = new aws.ssmcontacts.Contact("contactOne", {
    alias: "alias",
    type: "PERSONAL",
});
const contactTwo = new aws.ssmcontacts.Contact("contactTwo", {
    alias: "alias",
    type: "PERSONAL",
});
const test = new aws.ssmcontacts.Plan("test", {
    contactId: escalationPlan.arn,
    stages: [{
        durationInMinutes: 0,
        targets: [
            {
                contactTargetInfo: {
                    isEssential: false,
                    contactId: contactOne.arn,
                },
            },
            {
                contactTargetInfo: {
                    isEssential: true,
                    contactId: contactTwo.arn,
                },
            },
            {
                channelTargetInfo: {
                    retryIntervalInMinutes: 2,
                    contactChannelId: aws_ssmcontacts_contact_channel.channel.arn,
                },
            },
        ],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

escalation_plan = aws.ssmcontacts.Contact("escalationPlan",
    alias="escalation-plan-alias",
    type="ESCALATION")
contact_one = aws.ssmcontacts.Contact("contactOne",
    alias="alias",
    type="PERSONAL")
contact_two = aws.ssmcontacts.Contact("contactTwo",
    alias="alias",
    type="PERSONAL")
test = aws.ssmcontacts.Plan("test",
    contact_id=escalation_plan.arn,
    stages=[aws.ssmcontacts.PlanStageArgs(
        duration_in_minutes=0,
        targets=[
            aws.ssmcontacts.PlanStageTargetArgs(
                contact_target_info=aws.ssmcontacts.PlanStageTargetContactTargetInfoArgs(
                    is_essential=False,
                    contact_id=contact_one.arn,
                ),
            ),
            aws.ssmcontacts.PlanStageTargetArgs(
                contact_target_info=aws.ssmcontacts.PlanStageTargetContactTargetInfoArgs(
                    is_essential=True,
                    contact_id=contact_two.arn,
                ),
            ),
            aws.ssmcontacts.PlanStageTargetArgs(
                channel_target_info=aws.ssmcontacts.PlanStageTargetChannelTargetInfoArgs(
                    retry_interval_in_minutes=2,
                    contact_channel_id=aws_ssmcontacts_contact_channel["channel"]["arn"],
                ),
            ),
        ],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var escalationPlan = new Aws.SsmContacts.Contact("escalationPlan", new()
    {
        Alias = "escalation-plan-alias",
        Type = "ESCALATION",
    });

    var contactOne = new Aws.SsmContacts.Contact("contactOne", new()
    {
        Alias = "alias",
        Type = "PERSONAL",
    });

    var contactTwo = new Aws.SsmContacts.Contact("contactTwo", new()
    {
        Alias = "alias",
        Type = "PERSONAL",
    });

    var test = new Aws.SsmContacts.Plan("test", new()
    {
        ContactId = escalationPlan.Arn,
        Stages = new[]
        {
            new Aws.SsmContacts.Inputs.PlanStageArgs
            {
                DurationInMinutes = 0,
                Targets = new[]
                {
                    new Aws.SsmContacts.Inputs.PlanStageTargetArgs
                    {
                        ContactTargetInfo = new Aws.SsmContacts.Inputs.PlanStageTargetContactTargetInfoArgs
                        {
                            IsEssential = false,
                            ContactId = contactOne.Arn,
                        },
                    },
                    new Aws.SsmContacts.Inputs.PlanStageTargetArgs
                    {
                        ContactTargetInfo = new Aws.SsmContacts.Inputs.PlanStageTargetContactTargetInfoArgs
                        {
                            IsEssential = true,
                            ContactId = contactTwo.Arn,
                        },
                    },
                    new Aws.SsmContacts.Inputs.PlanStageTargetArgs
                    {
                        ChannelTargetInfo = new Aws.SsmContacts.Inputs.PlanStageTargetChannelTargetInfoArgs
                        {
                            RetryIntervalInMinutes = 2,
                            ContactChannelId = aws_ssmcontacts_contact_channel.Channel.Arn,
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmcontacts"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		escalationPlan, err := ssmcontacts.NewContact(ctx, "escalationPlan", &ssmcontacts.ContactArgs{
			Alias: pulumi.String("escalation-plan-alias"),
			Type:  pulumi.String("ESCALATION"),
		})
		if err != nil {
			return err
		}
		contactOne, err := ssmcontacts.NewContact(ctx, "contactOne", &ssmcontacts.ContactArgs{
			Alias: pulumi.String("alias"),
			Type:  pulumi.String("PERSONAL"),
		})
		if err != nil {
			return err
		}
		contactTwo, err := ssmcontacts.NewContact(ctx, "contactTwo", &ssmcontacts.ContactArgs{
			Alias: pulumi.String("alias"),
			Type:  pulumi.String("PERSONAL"),
		})
		if err != nil {
			return err
		}
		_, err = ssmcontacts.NewPlan(ctx, "test", &ssmcontacts.PlanArgs{
			ContactId: escalationPlan.Arn,
			Stages: ssmcontacts.PlanStageArray{
				&ssmcontacts.PlanStageArgs{
					DurationInMinutes: pulumi.Int(0),
					Targets: ssmcontacts.PlanStageTargetArray{
						&ssmcontacts.PlanStageTargetArgs{
							ContactTargetInfo: &ssmcontacts.PlanStageTargetContactTargetInfoArgs{
								IsEssential: pulumi.Bool(false),
								ContactId:   contactOne.Arn,
							},
						},
						&ssmcontacts.PlanStageTargetArgs{
							ContactTargetInfo: &ssmcontacts.PlanStageTargetContactTargetInfoArgs{
								IsEssential: pulumi.Bool(true),
								ContactId:   contactTwo.Arn,
							},
						},
						&ssmcontacts.PlanStageTargetArgs{
							ChannelTargetInfo: &ssmcontacts.PlanStageTargetChannelTargetInfoArgs{
								RetryIntervalInMinutes: pulumi.Int(2),
								ContactChannelId:       pulumi.Any(aws_ssmcontacts_contact_channel.Channel.Arn),
							},
						},
					},
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
import com.pulumi.aws.ssmcontacts.Contact;
import com.pulumi.aws.ssmcontacts.ContactArgs;
import com.pulumi.aws.ssmcontacts.Plan;
import com.pulumi.aws.ssmcontacts.PlanArgs;
import com.pulumi.aws.ssmcontacts.inputs.PlanStageArgs;
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
        var escalationPlan = new Contact("escalationPlan", ContactArgs.builder()        
            .alias("escalation-plan-alias")
            .type("ESCALATION")
            .build());

        var contactOne = new Contact("contactOne", ContactArgs.builder()        
            .alias("alias")
            .type("PERSONAL")
            .build());

        var contactTwo = new Contact("contactTwo", ContactArgs.builder()        
            .alias("alias")
            .type("PERSONAL")
            .build());

        var test = new Plan("test", PlanArgs.builder()        
            .contactId(escalationPlan.arn())
            .stages(PlanStageArgs.builder()
                .durationInMinutes(0)
                .targets(                
                    PlanStageTargetArgs.builder()
                        .contactTargetInfo(PlanStageTargetContactTargetInfoArgs.builder()
                            .isEssential(false)
                            .contactId(contactOne.arn())
                            .build())
                        .build(),
                    PlanStageTargetArgs.builder()
                        .contactTargetInfo(PlanStageTargetContactTargetInfoArgs.builder()
                            .isEssential(true)
                            .contactId(contactTwo.arn())
                            .build())
                        .build(),
                    PlanStageTargetArgs.builder()
                        .channelTargetInfo(PlanStageTargetChannelTargetInfoArgs.builder()
                            .retryIntervalInMinutes(2)
                            .contactChannelId(aws_ssmcontacts_contact_channel.channel().arn())
                            .build())
                        .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  escalationPlan:
    type: aws:ssmcontacts:Contact
    properties:
      alias: escalation-plan-alias
      type: ESCALATION
  contactOne:
    type: aws:ssmcontacts:Contact
    properties:
      alias: alias
      type: PERSONAL
  contactTwo:
    type: aws:ssmcontacts:Contact
    properties:
      alias: alias
      type: PERSONAL
  test:
    type: aws:ssmcontacts:Plan
    properties:
      contactId: ${escalationPlan.arn}
      stages:
        - durationInMinutes: 0
          targets:
            - contactTargetInfo:
                isEssential: false
                contactId: ${contactOne.arn}
            - contactTargetInfo:
                isEssential: true
                contactId: ${contactTwo.arn}
            - channelTargetInfo:
                retryIntervalInMinutes: 2
                contactChannelId: ${aws_ssmcontacts_contact_channel.channel.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SSM Contact Plan using the Contact ARN. For example:

```sh
 $ pulumi import aws:ssmcontacts/plan:Plan example {ARNValue}
```
 