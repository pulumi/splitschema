Resource for managing an AWS SESv2 (Simple Email V2) Contact List.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sesv2.ContactList("example", {contactListName: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.ContactList("example", contact_list_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SesV2.ContactList("example", new()
    {
        ContactListName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sesv2.NewContactList(ctx, "example", &sesv2.ContactListArgs{
			ContactListName: pulumi.String("example"),
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
import com.pulumi.aws.sesv2.ContactList;
import com.pulumi.aws.sesv2.ContactListArgs;
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
        var example = new ContactList("example", ContactListArgs.builder()        
            .contactListName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sesv2:ContactList
    properties:
      contactListName: example
```
{{% /example %}}
{{% example %}}
### Extended Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sesv2.ContactList("example", {
    contactListName: "example",
    description: "description",
    topics: [{
        defaultSubscriptionStatus: "OPT_IN",
        description: "topic description",
        displayName: "Example Topic",
        topicName: "example-topic",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.ContactList("example",
    contact_list_name="example",
    description="description",
    topics=[aws.sesv2.ContactListTopicArgs(
        default_subscription_status="OPT_IN",
        description="topic description",
        display_name="Example Topic",
        topic_name="example-topic",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SesV2.ContactList("example", new()
    {
        ContactListName = "example",
        Description = "description",
        Topics = new[]
        {
            new Aws.SesV2.Inputs.ContactListTopicArgs
            {
                DefaultSubscriptionStatus = "OPT_IN",
                Description = "topic description",
                DisplayName = "Example Topic",
                TopicName = "example-topic",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sesv2.NewContactList(ctx, "example", &sesv2.ContactListArgs{
			ContactListName: pulumi.String("example"),
			Description:     pulumi.String("description"),
			Topics: sesv2.ContactListTopicArray{
				&sesv2.ContactListTopicArgs{
					DefaultSubscriptionStatus: pulumi.String("OPT_IN"),
					Description:               pulumi.String("topic description"),
					DisplayName:               pulumi.String("Example Topic"),
					TopicName:                 pulumi.String("example-topic"),
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
import com.pulumi.aws.sesv2.ContactList;
import com.pulumi.aws.sesv2.ContactListArgs;
import com.pulumi.aws.sesv2.inputs.ContactListTopicArgs;
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
        var example = new ContactList("example", ContactListArgs.builder()        
            .contactListName("example")
            .description("description")
            .topics(ContactListTopicArgs.builder()
                .defaultSubscriptionStatus("OPT_IN")
                .description("topic description")
                .displayName("Example Topic")
                .topicName("example-topic")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sesv2:ContactList
    properties:
      contactListName: example
      description: description
      topics:
        - defaultSubscriptionStatus: OPT_IN
          description: topic description
          displayName: Example Topic
          topicName: example-topic
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SESv2 (Simple Email V2) Contact List using the `id`. For example:

```sh
 $ pulumi import aws:sesv2/contactList:ContactList example example
```
 