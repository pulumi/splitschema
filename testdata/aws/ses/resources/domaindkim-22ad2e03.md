Provides an SES domain DKIM generation resource.

Domain ownership needs to be confirmed first using ses_domain_identity Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDomainIdentity = new aws.ses.DomainIdentity("exampleDomainIdentity", {domain: "example.com"});
const exampleDomainDkim = new aws.ses.DomainDkim("exampleDomainDkim", {domain: exampleDomainIdentity.domain});
const exampleAmazonsesDkimRecord: aws.route53.Record[] = [];
for (const range = {value: 0}; range.value < 3; range.value++) {
    exampleAmazonsesDkimRecord.push(new aws.route53.Record(`exampleAmazonsesDkimRecord-${range.value}`, {
        zoneId: "ABCDEFGHIJ123",
        name: exampleDomainDkim.dkimTokens.apply(dkimTokens => `${dkimTokens[range.value]}._domainkey`),
        type: "CNAME",
        ttl: 600,
        records: [exampleDomainDkim.dkimTokens.apply(dkimTokens => `${dkimTokens[range.value]}.dkim.amazonses.com`)],
    }));
}
```
```python
import pulumi
import pulumi_aws as aws

example_domain_identity = aws.ses.DomainIdentity("exampleDomainIdentity", domain="example.com")
example_domain_dkim = aws.ses.DomainDkim("exampleDomainDkim", domain=example_domain_identity.domain)
example_amazonses_dkim_record = []
for range in [{"value": i} for i in range(0, 3)]:
    example_amazonses_dkim_record.append(aws.route53.Record(f"exampleAmazonsesDkimRecord-{range['value']}",
        zone_id="ABCDEFGHIJ123",
        name=example_domain_dkim.dkim_tokens.apply(lambda dkim_tokens: f"{dkim_tokens[range['value']]}._domainkey"),
        type="CNAME",
        ttl=600,
        records=[example_domain_dkim.dkim_tokens.apply(lambda dkim_tokens: f"{dkim_tokens[range['value']]}.dkim.amazonses.com")]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleDomainIdentity = new Aws.Ses.DomainIdentity("exampleDomainIdentity", new()
    {
        Domain = "example.com",
    });

    var exampleDomainDkim = new Aws.Ses.DomainDkim("exampleDomainDkim", new()
    {
        Domain = exampleDomainIdentity.Domain,
    });

    var exampleAmazonsesDkimRecord = new List<Aws.Route53.Record>();
    for (var rangeIndex = 0; rangeIndex < 3; rangeIndex++)
    {
        var range = new { Value = rangeIndex };
        exampleAmazonsesDkimRecord.Add(new Aws.Route53.Record($"exampleAmazonsesDkimRecord-{range.Value}", new()
        {
            ZoneId = "ABCDEFGHIJ123",
            Name = exampleDomainDkim.DkimTokens.Apply(dkimTokens => $"{dkimTokens[range.Value]}._domainkey"),
            Type = "CNAME",
            Ttl = 600,
            Records = new[]
            {
                exampleDomainDkim.DkimTokens.Apply(dkimTokens => $"{dkimTokens[range.Value]}.dkim.amazonses.com"),
            },
        }));
    }
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
		exampleDomainDkim, err := ses.NewDomainDkim(ctx, "exampleDomainDkim", &ses.DomainDkimArgs{
			Domain: exampleDomainIdentity.Domain,
		})
		if err != nil {
			return err
		}
		var exampleAmazonsesDkimRecord []*route53.Record
		for index := 0; index < 3; index++ {
			key0 := index
			val0 := index
			__res, err := route53.NewRecord(ctx, fmt.Sprintf("exampleAmazonsesDkimRecord-%v", key0), &route53.RecordArgs{
				ZoneId: pulumi.String("ABCDEFGHIJ123"),
				Name: exampleDomainDkim.DkimTokens.ApplyT(func(dkimTokens []string) (string, error) {
					return fmt.Sprintf("%v._domainkey", dkimTokens[val0]), nil
				}).(pulumi.StringOutput),
				Type: pulumi.String("CNAME"),
				Ttl:  pulumi.Int(600),
				Records: pulumi.StringArray{
					exampleDomainDkim.DkimTokens.ApplyT(func(dkimTokens []string) (string, error) {
						return fmt.Sprintf("%v.dkim.amazonses.com", dkimTokens[val0]), nil
					}).(pulumi.StringOutput),
				},
			})
			if err != nil {
				return err
			}
			exampleAmazonsesDkimRecord = append(exampleAmazonsesDkimRecord, __res)
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
import com.pulumi.aws.ses.DomainDkim;
import com.pulumi.aws.ses.DomainDkimArgs;
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
import com.pulumi.codegen.internal.KeyedValue;
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

        var exampleDomainDkim = new DomainDkim("exampleDomainDkim", DomainDkimArgs.builder()        
            .domain(exampleDomainIdentity.domain())
            .build());

        for (var i = 0; i < 3; i++) {
            new Record("exampleAmazonsesDkimRecord-" + i, RecordArgs.builder()            
                .zoneId("ABCDEFGHIJ123")
                .name(exampleDomainDkim.dkimTokens().applyValue(dkimTokens -> String.format("%s._domainkey", dkimTokens[range.value()])))
                .type("CNAME")
                .ttl("600")
                .records(exampleDomainDkim.dkimTokens().applyValue(dkimTokens -> String.format("%s.dkim.amazonses.com", dkimTokens[range.value()])))
                .build());

        
}
    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DKIM tokens using the `domain` attribute. For example:

```sh
 $ pulumi import aws:ses/domainDkim:DomainDkim example example.com
```
 