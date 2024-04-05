Represents a successful verification of an SES domain identity.

Most commonly, this resource is used together with `aws.route53.Record` and
`aws.ses.DomainIdentity` to request an SES domain identity,
deploy the required DNS verification records, and wait for verification to complete.

> **WARNING:** This resource implements a part of the verification workflow. It does not represent a real-world entity in AWS, therefore changing or deleting this resource on its own has no immediate effect.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ses.DomainIdentity("example", {domain: "example.com"});
const exampleAmazonsesVerificationRecord = new aws.route53.Record("exampleAmazonsesVerificationRecord", {
    zoneId: aws_route53_zone.example.zone_id,
    name: pulumi.interpolate`_amazonses.${example.id}`,
    type: "TXT",
    ttl: 600,
    records: [example.verificationToken],
});
const exampleVerification = new aws.ses.DomainIdentityVerification("exampleVerification", {domain: example.id}, {
    dependsOn: [exampleAmazonsesVerificationRecord],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ses.DomainIdentity("example", domain="example.com")
example_amazonses_verification_record = aws.route53.Record("exampleAmazonsesVerificationRecord",
    zone_id=aws_route53_zone["example"]["zone_id"],
    name=example.id.apply(lambda id: f"_amazonses.{id}"),
    type="TXT",
    ttl=600,
    records=[example.verification_token])
example_verification = aws.ses.DomainIdentityVerification("exampleVerification", domain=example.id,
opts=pulumi.ResourceOptions(depends_on=[example_amazonses_verification_record]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ses.DomainIdentity("example", new()
    {
        Domain = "example.com",
    });

    var exampleAmazonsesVerificationRecord = new Aws.Route53.Record("exampleAmazonsesVerificationRecord", new()
    {
        ZoneId = aws_route53_zone.Example.Zone_id,
        Name = example.Id.Apply(id => $"_amazonses.{id}"),
        Type = "TXT",
        Ttl = 600,
        Records = new[]
        {
            example.VerificationToken,
        },
    });

    var exampleVerification = new Aws.Ses.DomainIdentityVerification("exampleVerification", new()
    {
        Domain = example.Id,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAmazonsesVerificationRecord,
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
		example, err := ses.NewDomainIdentity(ctx, "example", &ses.DomainIdentityArgs{
			Domain: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		exampleAmazonsesVerificationRecord, err := route53.NewRecord(ctx, "exampleAmazonsesVerificationRecord", &route53.RecordArgs{
			ZoneId: pulumi.Any(aws_route53_zone.Example.Zone_id),
			Name: example.ID().ApplyT(func(id string) (string, error) {
				return fmt.Sprintf("_amazonses.%v", id), nil
			}).(pulumi.StringOutput),
			Type: pulumi.String("TXT"),
			Ttl:  pulumi.Int(600),
			Records: pulumi.StringArray{
				example.VerificationToken,
			},
		})
		if err != nil {
			return err
		}
		_, err = ses.NewDomainIdentityVerification(ctx, "exampleVerification", &ses.DomainIdentityVerificationArgs{
			Domain: example.ID(),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAmazonsesVerificationRecord,
		}))
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
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
import com.pulumi.aws.ses.DomainIdentityVerification;
import com.pulumi.aws.ses.DomainIdentityVerificationArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var example = new DomainIdentity("example", DomainIdentityArgs.builder()        
            .domain("example.com")
            .build());

        var exampleAmazonsesVerificationRecord = new Record("exampleAmazonsesVerificationRecord", RecordArgs.builder()        
            .zoneId(aws_route53_zone.example().zone_id())
            .name(example.id().applyValue(id -> String.format("_amazonses.%s", id)))
            .type("TXT")
            .ttl("600")
            .records(example.verificationToken())
            .build());

        var exampleVerification = new DomainIdentityVerification("exampleVerification", DomainIdentityVerificationArgs.builder()        
            .domain(example.id())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAmazonsesVerificationRecord)
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ses:DomainIdentity
    properties:
      domain: example.com
  exampleAmazonsesVerificationRecord:
    type: aws:route53:Record
    properties:
      zoneId: ${aws_route53_zone.example.zone_id}
      name: _amazonses.${example.id}
      type: TXT
      ttl: '600'
      records:
        - ${example.verificationToken}
  exampleVerification:
    type: aws:ses:DomainIdentityVerification
    properties:
      domain: ${example.id}
    options:
      dependson:
        - ${exampleAmazonsesVerificationRecord}
```
{{% /example %}}
{{% /examples %}}