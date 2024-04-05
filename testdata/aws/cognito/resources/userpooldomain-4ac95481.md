Provides a Cognito User Pool Domain resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Amazon Cognito domain

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cognito.UserPool("example", {});
const main = new aws.cognito.UserPoolDomain("main", {
    domain: "example-domain",
    userPoolId: example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cognito.UserPool("example")
main = aws.cognito.UserPoolDomain("main",
    domain="example-domain",
    user_pool_id=example.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Cognito.UserPool("example");

    var main = new Aws.Cognito.UserPoolDomain("main", new()
    {
        Domain = "example-domain",
        UserPoolId = example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := cognito.NewUserPool(ctx, "example", nil)
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolDomain(ctx, "main", &cognito.UserPoolDomainArgs{
			Domain:     pulumi.String("example-domain"),
			UserPoolId: example.ID(),
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolDomain;
import com.pulumi.aws.cognito.UserPoolDomainArgs;
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
        var example = new UserPool("example");

        var main = new UserPoolDomain("main", UserPoolDomainArgs.builder()        
            .domain("example-domain")
            .userPoolId(example.id())
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:cognito:UserPoolDomain
    properties:
      domain: example-domain
      userPoolId: ${example.id}
  example:
    type: aws:cognito:UserPool
```
{{% /example %}}
{{% example %}}
### Custom Cognito domain

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleUserPool = new aws.cognito.UserPool("exampleUserPool", {});
const main = new aws.cognito.UserPoolDomain("main", {
    domain: "example-domain",
    certificateArn: aws_acm_certificate.cert.arn,
    userPoolId: exampleUserPool.id,
});
const exampleZone = aws.route53.getZone({
    name: "example.com",
});
const auth_cognito_A = new aws.route53.Record("auth-cognito-A", {
    name: main.domain,
    type: "A",
    zoneId: exampleZone.then(exampleZone => exampleZone.zoneId),
    aliases: [{
        evaluateTargetHealth: false,
        name: main.cloudfrontDistribution,
        zoneId: main.cloudfrontDistributionZoneId,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_user_pool = aws.cognito.UserPool("exampleUserPool")
main = aws.cognito.UserPoolDomain("main",
    domain="example-domain",
    certificate_arn=aws_acm_certificate["cert"]["arn"],
    user_pool_id=example_user_pool.id)
example_zone = aws.route53.get_zone(name="example.com")
auth_cognito__a = aws.route53.Record("auth-cognito-A",
    name=main.domain,
    type="A",
    zone_id=example_zone.zone_id,
    aliases=[aws.route53.RecordAliasArgs(
        evaluate_target_health=False,
        name=main.cloudfront_distribution,
        zone_id=main.cloudfront_distribution_zone_id,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleUserPool = new Aws.Cognito.UserPool("exampleUserPool");

    var main = new Aws.Cognito.UserPoolDomain("main", new()
    {
        Domain = "example-domain",
        CertificateArn = aws_acm_certificate.Cert.Arn,
        UserPoolId = exampleUserPool.Id,
    });

    var exampleZone = Aws.Route53.GetZone.Invoke(new()
    {
        Name = "example.com",
    });

    var auth_cognito_A = new Aws.Route53.Record("auth-cognito-A", new()
    {
        Name = main.Domain,
        Type = "A",
        ZoneId = exampleZone.Apply(getZoneResult => getZoneResult.ZoneId),
        Aliases = new[]
        {
            new Aws.Route53.Inputs.RecordAliasArgs
            {
                EvaluateTargetHealth = false,
                Name = main.CloudfrontDistribution,
                ZoneId = main.CloudfrontDistributionZoneId,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleUserPool, err := cognito.NewUserPool(ctx, "exampleUserPool", nil)
		if err != nil {
			return err
		}
		main, err := cognito.NewUserPoolDomain(ctx, "main", &cognito.UserPoolDomainArgs{
			Domain:         pulumi.String("example-domain"),
			CertificateArn: pulumi.Any(aws_acm_certificate.Cert.Arn),
			UserPoolId:     exampleUserPool.ID(),
		})
		if err != nil {
			return err
		}
		exampleZone, err := route53.LookupZone(ctx, &route53.LookupZoneArgs{
			Name: pulumi.StringRef("example.com"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "auth-cognito-A", &route53.RecordArgs{
			Name:   main.Domain,
			Type:   pulumi.String("A"),
			ZoneId: *pulumi.String(exampleZone.ZoneId),
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					EvaluateTargetHealth: pulumi.Bool(false),
					Name:                 main.CloudfrontDistribution,
					ZoneId:               main.CloudfrontDistributionZoneId,
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolDomain;
import com.pulumi.aws.cognito.UserPoolDomainArgs;
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetZoneArgs;
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
import com.pulumi.aws.route53.inputs.RecordAliasArgs;
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
        var exampleUserPool = new UserPool("exampleUserPool");

        var main = new UserPoolDomain("main", UserPoolDomainArgs.builder()        
            .domain("example-domain")
            .certificateArn(aws_acm_certificate.cert().arn())
            .userPoolId(exampleUserPool.id())
            .build());

        final var exampleZone = Route53Functions.getZone(GetZoneArgs.builder()
            .name("example.com")
            .build());

        var auth_cognito_A = new Record("auth-cognito-A", RecordArgs.builder()        
            .name(main.domain())
            .type("A")
            .zoneId(exampleZone.applyValue(getZoneResult -> getZoneResult.zoneId()))
            .aliases(RecordAliasArgs.builder()
                .evaluateTargetHealth(false)
                .name(main.cloudfrontDistribution())
                .zoneId(main.cloudfrontDistributionZoneId())
                .build())
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:cognito:UserPoolDomain
    properties:
      domain: example-domain
      certificateArn: ${aws_acm_certificate.cert.arn}
      userPoolId: ${exampleUserPool.id}
  exampleUserPool:
    type: aws:cognito:UserPool
  auth-cognito-A:
    type: aws:route53:Record
    properties:
      name: ${main.domain}
      type: A
      zoneId: ${exampleZone.zoneId}
      aliases:
        - evaluateTargetHealth: false
          name: ${main.cloudfrontDistribution}
          zoneId: ${main.cloudfrontDistributionZoneId}
variables:
  exampleZone:
    fn::invoke:
      Function: aws:route53:getZone
      Arguments:
        name: example.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cognito User Pool Domains using the `domain`. For example:

```sh
 $ pulumi import aws:cognito/userPoolDomain:UserPoolDomain main auth.example.org
```
 