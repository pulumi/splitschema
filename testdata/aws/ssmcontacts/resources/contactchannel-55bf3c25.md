Resource for managing an AWS SSM Contacts Contact Channel.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssmcontacts.ContactChannel("example", {
    contactId: "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
    deliveryAddress: {
        simpleAddress: "email@example.com",
    },
    type: "EMAIL",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssmcontacts.ContactChannel("example",
    contact_id="arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
    delivery_address=aws.ssmcontacts.ContactChannelDeliveryAddressArgs(
        simple_address="email@example.com",
    ),
    type="EMAIL")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SsmContacts.ContactChannel("example", new()
    {
        ContactId = "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
        DeliveryAddress = new Aws.SsmContacts.Inputs.ContactChannelDeliveryAddressArgs
        {
            SimpleAddress = "email@example.com",
        },
        Type = "EMAIL",
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
		_, err := ssmcontacts.NewContactChannel(ctx, "example", &ssmcontacts.ContactChannelArgs{
			ContactId: pulumi.String("arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias"),
			DeliveryAddress: &ssmcontacts.ContactChannelDeliveryAddressArgs{
				SimpleAddress: pulumi.String("email@example.com"),
			},
			Type: pulumi.String("EMAIL"),
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
import com.pulumi.aws.ssmcontacts.ContactChannel;
import com.pulumi.aws.ssmcontacts.ContactChannelArgs;
import com.pulumi.aws.ssmcontacts.inputs.ContactChannelDeliveryAddressArgs;
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
        var example = new ContactChannel("example", ContactChannelArgs.builder()        
            .contactId("arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias")
            .deliveryAddress(ContactChannelDeliveryAddressArgs.builder()
                .simpleAddress("email@example.com")
                .build())
            .type("EMAIL")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssmcontacts:ContactChannel
    properties:
      contactId: arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias
      deliveryAddress:
        simpleAddress: email@example.com
      type: EMAIL
```
{{% /example %}}
{{% example %}}
### Usage with SSM Contact

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleContact = new aws.ssmcontacts.Contact("exampleContact", {
    alias: "example_contact",
    type: "PERSONAL",
});
const example = new aws.ssmcontacts.ContactChannel("example", {
    contactId: exampleContact.arn,
    deliveryAddress: {
        simpleAddress: "email@example.com",
    },
    type: "EMAIL",
});
```
```python
import pulumi
import pulumi_aws as aws

example_contact = aws.ssmcontacts.Contact("exampleContact",
    alias="example_contact",
    type="PERSONAL")
example = aws.ssmcontacts.ContactChannel("example",
    contact_id=example_contact.arn,
    delivery_address=aws.ssmcontacts.ContactChannelDeliveryAddressArgs(
        simple_address="email@example.com",
    ),
    type="EMAIL")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleContact = new Aws.SsmContacts.Contact("exampleContact", new()
    {
        Alias = "example_contact",
        Type = "PERSONAL",
    });

    var example = new Aws.SsmContacts.ContactChannel("example", new()
    {
        ContactId = exampleContact.Arn,
        DeliveryAddress = new Aws.SsmContacts.Inputs.ContactChannelDeliveryAddressArgs
        {
            SimpleAddress = "email@example.com",
        },
        Type = "EMAIL",
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
		exampleContact, err := ssmcontacts.NewContact(ctx, "exampleContact", &ssmcontacts.ContactArgs{
			Alias: pulumi.String("example_contact"),
			Type:  pulumi.String("PERSONAL"),
		})
		if err != nil {
			return err
		}
		_, err = ssmcontacts.NewContactChannel(ctx, "example", &ssmcontacts.ContactChannelArgs{
			ContactId: exampleContact.Arn,
			DeliveryAddress: &ssmcontacts.ContactChannelDeliveryAddressArgs{
				SimpleAddress: pulumi.String("email@example.com"),
			},
			Type: pulumi.String("EMAIL"),
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
import com.pulumi.aws.ssmcontacts.ContactChannel;
import com.pulumi.aws.ssmcontacts.ContactChannelArgs;
import com.pulumi.aws.ssmcontacts.inputs.ContactChannelDeliveryAddressArgs;
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
        var exampleContact = new Contact("exampleContact", ContactArgs.builder()        
            .alias("example_contact")
            .type("PERSONAL")
            .build());

        var example = new ContactChannel("example", ContactChannelArgs.builder()        
            .contactId(exampleContact.arn())
            .deliveryAddress(ContactChannelDeliveryAddressArgs.builder()
                .simpleAddress("email@example.com")
                .build())
            .type("EMAIL")
            .build());

    }
}
```
```yaml
resources:
  exampleContact:
    type: aws:ssmcontacts:Contact
    properties:
      alias: example_contact
      type: PERSONAL
  example:
    type: aws:ssmcontacts:ContactChannel
    properties:
      contactId: ${exampleContact.arn}
      deliveryAddress:
        simpleAddress: email@example.com
      type: EMAIL
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SSM Contact Channel using the `ARN`. For example:

```sh
 $ pulumi import aws:ssmcontacts/contactChannel:ContactChannel example arn:aws:ssm-contacts:us-west-2:123456789012:contact-channel/example
```
 