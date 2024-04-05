Provides an SES domain MAIL FROM resource.

> **NOTE:** For the MAIL FROM domain to be fully usable, this resource should be paired with the aws.ses.DomainIdentity resource. To validate the MAIL FROM domain, a DNS MX record is required. To pass SPF checks, a DNS TXT record may also be required. See the [Amazon SES MAIL FROM documentation](https://docs.aws.amazon.com/ses/latest/dg/mail-from.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}
### Domain Identity MAIL FROM

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Example SES Domain Identity
const exampleDomainIdentity = new aws.ses.DomainIdentity("exampleDomainIdentity", {domain: "example.com"});
const exampleMailFrom = new aws.ses.MailFrom("exampleMailFrom", {
    domain: exampleDomainIdentity.domain,
    mailFromDomain: pulumi.interpolate`bounce.${exampleDomainIdentity.domain}`,
});
// Example Route53 MX record
const exampleSesDomainMailFromMx = new aws.route53.Record("exampleSesDomainMailFromMx", {
    zoneId: aws_route53_zone.example.id,
    name: exampleMailFrom.mailFromDomain,
    type: "MX",
    ttl: 600,
    records: ["10 feedback-smtp.us-east-1.amazonses.com"],
});
// Change to the region in which `aws_ses_domain_identity.example` is created
// Example Route53 TXT record for SPF
const exampleSesDomainMailFromTxt = new aws.route53.Record("exampleSesDomainMailFromTxt", {
    zoneId: aws_route53_zone.example.id,
    name: exampleMailFrom.mailFromDomain,
    type: "TXT",
    ttl: 600,
    records: ["v=spf1 include:amazonses.com -all"],
});
```
```python
import pulumi
import pulumi_aws as aws

# Example SES Domain Identity
example_domain_identity = aws.ses.DomainIdentity("exampleDomainIdentity", domain="example.com")
example_mail_from = aws.ses.MailFrom("exampleMailFrom",
    domain=example_domain_identity.domain,
    mail_from_domain=example_domain_identity.domain.apply(lambda domain: f"bounce.{domain}"))
# Example Route53 MX record
example_ses_domain_mail_from_mx = aws.route53.Record("exampleSesDomainMailFromMx",
    zone_id=aws_route53_zone["example"]["id"],
    name=example_mail_from.mail_from_domain,
    type="MX",
    ttl=600,
    records=["10 feedback-smtp.us-east-1.amazonses.com"])
# Change to the region in which `aws_ses_domain_identity.example` is created
# Example Route53 TXT record for SPF
example_ses_domain_mail_from_txt = aws.route53.Record("exampleSesDomainMailFromTxt",
    zone_id=aws_route53_zone["example"]["id"],
    name=example_mail_from.mail_from_domain,
    type="TXT",
    ttl=600,
    records=["v=spf1 include:amazonses.com -all"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Example SES Domain Identity
    var exampleDomainIdentity = new Aws.Ses.DomainIdentity("exampleDomainIdentity", new()
    {
        Domain = "example.com",
    });

    var exampleMailFrom = new Aws.Ses.MailFrom("exampleMailFrom", new()
    {
        Domain = exampleDomainIdentity.Domain,
        MailFromDomain = exampleDomainIdentity.Domain.Apply(domain => $"bounce.{domain}"),
    });

    // Example Route53 MX record
    var exampleSesDomainMailFromMx = new Aws.Route53.Record("exampleSesDomainMailFromMx", new()
    {
        ZoneId = aws_route53_zone.Example.Id,
        Name = exampleMailFrom.MailFromDomain,
        Type = "MX",
        Ttl = 600,
        Records = new[]
        {
            "10 feedback-smtp.us-east-1.amazonses.com",
        },
    });

    // Change to the region in which `aws_ses_domain_identity.example` is created
    // Example Route53 TXT record for SPF
    var exampleSesDomainMailFromTxt = new Aws.Route53.Record("exampleSesDomainMailFromTxt", new()
    {
        ZoneId = aws_route53_zone.Example.Id,
        Name = exampleMailFrom.MailFromDomain,
        Type = "TXT",
        Ttl = 600,
        Records = new[]
        {
            "v=spf1 include:amazonses.com -all",
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDomainIdentity, err := ses.NewDomainIdentity(ctx, "exampleDomainIdentity", &ses.DomainIdentityArgs{
			Domain: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		exampleMailFrom, err := ses.NewMailFrom(ctx, "exampleMailFrom", &ses.MailFromArgs{
			Domain: exampleDomainIdentity.Domain,
			MailFromDomain: exampleDomainIdentity.Domain.ApplyT(func(domain string) (string, error) {
				return fmt.Sprintf("bounce.%v", domain), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleSesDomainMailFromMx", &route53.RecordArgs{
			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
			Name:   exampleMailFrom.MailFromDomain,
			Type:   pulumi.String("MX"),
			Ttl:    pulumi.Int(600),
			Records: pulumi.StringArray{
				pulumi.String("10 feedback-smtp.us-east-1.amazonses.com"),
			},
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleSesDomainMailFromTxt", &route53.RecordArgs{
			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
			Name:   exampleMailFrom.MailFromDomain,
			Type:   pulumi.String("TXT"),
			Ttl:    pulumi.Int(600),
			Records: pulumi.StringArray{
				pulumi.String("v=spf1 include:amazonses.com -all"),
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
import com.pulumi.aws.ses.DomainIdentity;
import com.pulumi.aws.ses.DomainIdentityArgs;
import com.pulumi.aws.ses.MailFrom;
import com.pulumi.aws.ses.MailFromArgs;
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
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
        var exampleDomainIdentity = new DomainIdentity("exampleDomainIdentity", DomainIdentityArgs.builder()        
            .domain("example.com")
            .build());

        var exampleMailFrom = new MailFrom("exampleMailFrom", MailFromArgs.builder()        
            .domain(exampleDomainIdentity.domain())
            .mailFromDomain(exampleDomainIdentity.domain().applyValue(domain -> String.format("bounce.%s", domain)))
            .build());

        var exampleSesDomainMailFromMx = new Record("exampleSesDomainMailFromMx", RecordArgs.builder()        
            .zoneId(aws_route53_zone.example().id())
            .name(exampleMailFrom.mailFromDomain())
            .type("MX")
            .ttl("600")
            .records("10 feedback-smtp.us-east-1.amazonses.com")
            .build());

        var exampleSesDomainMailFromTxt = new Record("exampleSesDomainMailFromTxt", RecordArgs.builder()        
            .zoneId(aws_route53_zone.example().id())
            .name(exampleMailFrom.mailFromDomain())
            .type("TXT")
            .ttl("600")
            .records("v=spf1 include:amazonses.com -all")
            .build());

    }
}
```
```yaml
resources:
  exampleMailFrom:
    type: aws:ses:MailFrom
    properties:
      domain: ${exampleDomainIdentity.domain}
      mailFromDomain: bounce.${exampleDomainIdentity.domain}
  # Example SES Domain Identity
  exampleDomainIdentity:
    type: aws:ses:DomainIdentity
    properties:
      domain: example.com
  # Example Route53 MX record
  exampleSesDomainMailFromMx:
    type: aws:route53:Record
    properties:
      zoneId: ${aws_route53_zone.example.id}
      name: ${exampleMailFrom.mailFromDomain}
      type: MX
      ttl: '600'
      records:
        - 10 feedback-smtp.us-east-1.amazonses.com
  # Example Route53 TXT record for SPF
  exampleSesDomainMailFromTxt:
    type: aws:route53:Record
    properties:
      zoneId: ${aws_route53_zone.example.id}
      name: ${exampleMailFrom.mailFromDomain}
      type: TXT
      ttl: '600'
      records:
        - v=spf1 include:amazonses.com -all
```
{{% /example %}}
{{% example %}}
### Email Identity MAIL FROM

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Example SES Email Identity
const exampleEmailIdentity = new aws.ses.EmailIdentity("exampleEmailIdentity", {email: "user@example.com"});
const exampleMailFrom = new aws.ses.MailFrom("exampleMailFrom", {
    domain: exampleEmailIdentity.email,
    mailFromDomain: "mail.example.com",
});
```
```python
import pulumi
import pulumi_aws as aws

# Example SES Email Identity
example_email_identity = aws.ses.EmailIdentity("exampleEmailIdentity", email="user@example.com")
example_mail_from = aws.ses.MailFrom("exampleMailFrom",
    domain=example_email_identity.email,
    mail_from_domain="mail.example.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Example SES Email Identity
    var exampleEmailIdentity = new Aws.Ses.EmailIdentity("exampleEmailIdentity", new()
    {
        Email = "user@example.com",
    });

    var exampleMailFrom = new Aws.Ses.MailFrom("exampleMailFrom", new()
    {
        Domain = exampleEmailIdentity.Email,
        MailFromDomain = "mail.example.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleEmailIdentity, err := ses.NewEmailIdentity(ctx, "exampleEmailIdentity", &ses.EmailIdentityArgs{
			Email: pulumi.String("user@example.com"),
		})
		if err != nil {
			return err
		}
		_, err = ses.NewMailFrom(ctx, "exampleMailFrom", &ses.MailFromArgs{
			Domain:         exampleEmailIdentity.Email,
			MailFromDomain: pulumi.String("mail.example.com"),
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
import com.pulumi.aws.ses.EmailIdentity;
import com.pulumi.aws.ses.EmailIdentityArgs;
import com.pulumi.aws.ses.MailFrom;
import com.pulumi.aws.ses.MailFromArgs;
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
        var exampleEmailIdentity = new EmailIdentity("exampleEmailIdentity", EmailIdentityArgs.builder()        
            .email("user@example.com")
            .build());

        var exampleMailFrom = new MailFrom("exampleMailFrom", MailFromArgs.builder()        
            .domain(exampleEmailIdentity.email())
            .mailFromDomain("mail.example.com")
            .build());

    }
}
```
```yaml
resources:
  # Example SES Email Identity
  exampleEmailIdentity:
    type: aws:ses:EmailIdentity
    properties:
      email: user@example.com
  exampleMailFrom:
    type: aws:ses:MailFrom
    properties:
      domain: ${exampleEmailIdentity.email}
      mailFromDomain: mail.example.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MAIL FROM domain using the `domain` attribute. For example:

```sh
 $ pulumi import aws:ses/mailFrom:MailFrom example example.com
```
 