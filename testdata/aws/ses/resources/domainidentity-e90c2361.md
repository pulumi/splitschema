Provides an SES domain identity resource

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ses.DomainIdentity("example", {domain: "example.com"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ses.DomainIdentity("example", domain="example.com")
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
		_, err := ses.NewDomainIdentity(ctx, "example", &ses.DomainIdentityArgs{
			Domain: pulumi.String("example.com"),
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

    }
}
```
```yaml
resources:
  example:
    type: aws:ses:DomainIdentity
    properties:
      domain: example.com
```
{{% /example %}}
{{% example %}}
### With Route53 Record

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ses.DomainIdentity("example", {domain: "example.com"});
const exampleAmazonsesVerificationRecord = new aws.route53.Record("exampleAmazonsesVerificationRecord", {
    zoneId: "ABCDEFGHIJ123",
    name: "_amazonses.example.com",
    type: "TXT",
    ttl: 600,
    records: [example.verificationToken],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ses.DomainIdentity("example", domain="example.com")
example_amazonses_verification_record = aws.route53.Record("exampleAmazonsesVerificationRecord",
    zone_id="ABCDEFGHIJ123",
    name="_amazonses.example.com",
    type="TXT",
    ttl=600,
    records=[example.verification_token])
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
        ZoneId = "ABCDEFGHIJ123",
        Name = "_amazonses.example.com",
        Type = "TXT",
        Ttl = 600,
        Records = new[]
        {
            example.VerificationToken,
        },
    });

});
```
```go
package main

import (
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
		_, err = route53.NewRecord(ctx, "exampleAmazonsesVerificationRecord", &route53.RecordArgs{
			ZoneId: pulumi.String("ABCDEFGHIJ123"),
			Name:   pulumi.String("_amazonses.example.com"),
			Type:   pulumi.String("TXT"),
			Ttl:    pulumi.Int(600),
			Records: pulumi.StringArray{
				example.VerificationToken,
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
        var example = new DomainIdentity("example", DomainIdentityArgs.builder()        
            .domain("example.com")
            .build());

        var exampleAmazonsesVerificationRecord = new Record("exampleAmazonsesVerificationRecord", RecordArgs.builder()        
            .zoneId("ABCDEFGHIJ123")
            .name("_amazonses.example.com")
            .type("TXT")
            .ttl("600")
            .records(example.verificationToken())
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
      zoneId: ABCDEFGHIJ123
      name: _amazonses.example.com
      type: TXT
      ttl: '600'
      records:
        - ${example.verificationToken}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SES domain identities using the domain name. For example:

```sh
 $ pulumi import aws:ses/domainIdentity:DomainIdentity example example.com
```
 