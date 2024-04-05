Registers a custom domain name for use with AWS API Gateway. Additional information about this functionality
can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).

This resource just establishes ownership of and the TLS settings for
a particular domain name. An API can be attached to a particular path
under the registered domain name using
the `aws.apigateway.BasePathMapping` resource.

API Gateway domains can be defined as either 'edge-optimized' or 'regional'.  In an edge-optimized configuration,
API Gateway internally creates and manages a CloudFront distribution to route requests on the given hostname. In
addition to this resource it's necessary to create a DNS record corresponding to the given domain name which is an alias
(either Route53 alias or traditional CNAME) to the Cloudfront domain name exported in the `cloudfront_domain_name`
attribute.

In a regional configuration, API Gateway does not create a CloudFront distribution to route requests to the API, though
a distribution can be created if needed. In either case, it is necessary to create a DNS record corresponding to the
given domain name which is an alias (either Route53 alias or traditional CNAME) to the regional domain name exported in
the `regional_domain_name` attribute.

> **Note:** API Gateway requires the use of AWS Certificate Manager (ACM) certificates instead of Identity and Access Management (IAM) certificates in regions that support ACM. Regions that support ACM can be found in the [Regions and Endpoints Documentation](https://docs.aws.amazon.com/general/latest/gr/rande.html#acm_region). To import an existing private key and certificate into ACM or request an ACM certificate, see the `aws.acm.Certificate` resource.

> **Note:** The `aws.apigateway.DomainName` resource expects dependency on the `aws.acm.CertificateValidation` as
only verified certificates can be used. This can be made either explicitly by adding the
`depends_on = [aws_acm_certificate_validation.cert]` attribute. Or implicitly by referring certificate ARN
from the validation resource where it will be available after the resource creation:
`regional_certificate_arn = aws_acm_certificate_validation.cert.certificate_arn`.

{{% examples %}}
## Example Usage
{{% example %}}
### Edge Optimized (ACM Certificate)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDomainName = new aws.apigateway.DomainName("exampleDomainName", {
    certificateArn: aws_acm_certificate_validation.example.certificate_arn,
    domainName: "api.example.com",
});
// Example DNS record using Route53.
// Route53 is not specifically required; any DNS host can be used.
const exampleRecord = new aws.route53.Record("exampleRecord", {
    name: exampleDomainName.domainName,
    type: "A",
    zoneId: aws_route53_zone.example.id,
    aliases: [{
        evaluateTargetHealth: true,
        name: exampleDomainName.cloudfrontDomainName,
        zoneId: exampleDomainName.cloudfrontZoneId,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_domain_name = aws.apigateway.DomainName("exampleDomainName",
    certificate_arn=aws_acm_certificate_validation["example"]["certificate_arn"],
    domain_name="api.example.com")
# Example DNS record using Route53.
# Route53 is not specifically required; any DNS host can be used.
example_record = aws.route53.Record("exampleRecord",
    name=example_domain_name.domain_name,
    type="A",
    zone_id=aws_route53_zone["example"]["id"],
    aliases=[aws.route53.RecordAliasArgs(
        evaluate_target_health=True,
        name=example_domain_name.cloudfront_domain_name,
        zone_id=example_domain_name.cloudfront_zone_id,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleDomainName = new Aws.ApiGateway.DomainName("exampleDomainName", new()
    {
        CertificateArn = aws_acm_certificate_validation.Example.Certificate_arn,
        Domain = "api.example.com",
    });

    // Example DNS record using Route53.
    // Route53 is not specifically required; any DNS host can be used.
    var exampleRecord = new Aws.Route53.Record("exampleRecord", new()
    {
        Name = exampleDomainName.Domain,
        Type = "A",
        ZoneId = aws_route53_zone.Example.Id,
        Aliases = new[]
        {
            new Aws.Route53.Inputs.RecordAliasArgs
            {
                EvaluateTargetHealth = true,
                Name = exampleDomainName.CloudfrontDomainName,
                ZoneId = exampleDomainName.CloudfrontZoneId,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDomainName, err := apigateway.NewDomainName(ctx, "exampleDomainName", &apigateway.DomainNameArgs{
			CertificateArn: pulumi.Any(aws_acm_certificate_validation.Example.Certificate_arn),
			DomainName:     pulumi.String("api.example.com"),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleRecord", &route53.RecordArgs{
			Name:   exampleDomainName.DomainName,
			Type:   pulumi.String("A"),
			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					EvaluateTargetHealth: pulumi.Bool(true),
					Name:                 exampleDomainName.CloudfrontDomainName,
					ZoneId:               exampleDomainName.CloudfrontZoneId,
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
import com.pulumi.aws.apigateway.DomainName;
import com.pulumi.aws.apigateway.DomainNameArgs;
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
        var exampleDomainName = new DomainName("exampleDomainName", DomainNameArgs.builder()        
            .certificateArn(aws_acm_certificate_validation.example().certificate_arn())
            .domainName("api.example.com")
            .build());

        var exampleRecord = new Record("exampleRecord", RecordArgs.builder()        
            .name(exampleDomainName.domainName())
            .type("A")
            .zoneId(aws_route53_zone.example().id())
            .aliases(RecordAliasArgs.builder()
                .evaluateTargetHealth(true)
                .name(exampleDomainName.cloudfrontDomainName())
                .zoneId(exampleDomainName.cloudfrontZoneId())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleDomainName:
    type: aws:apigateway:DomainName
    properties:
      certificateArn: ${aws_acm_certificate_validation.example.certificate_arn}
      domainName: api.example.com
  # Example DNS record using Route53.
  # Route53 is not specifically required; any DNS host can be used.
  exampleRecord:
    type: aws:route53:Record
    properties:
      name: ${exampleDomainName.domainName}
      type: A
      zoneId: ${aws_route53_zone.example.id}
      aliases:
        - evaluateTargetHealth: true
          name: ${exampleDomainName.cloudfrontDomainName}
          zoneId: ${exampleDomainName.cloudfrontZoneId}
```
{{% /example %}}
{{% example %}}
### Edge Optimized (IAM Certificate)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const exampleDomainName = new aws.apigateway.DomainName("exampleDomainName", {
    domainName: "api.example.com",
    certificateName: "example-api",
    certificateBody: fs.readFileSync(`${path.module}/example.com/example.crt`, "utf8"),
    certificateChain: fs.readFileSync(`${path.module}/example.com/ca.crt`, "utf8"),
    certificatePrivateKey: fs.readFileSync(`${path.module}/example.com/example.key`, "utf8"),
});
// Example DNS record using Route53.
// Route53 is not specifically required; any DNS host can be used.
const exampleRecord = new aws.route53.Record("exampleRecord", {
    zoneId: aws_route53_zone.example.id,
    name: exampleDomainName.domainName,
    type: "A",
    aliases: [{
        name: exampleDomainName.cloudfrontDomainName,
        zoneId: exampleDomainName.cloudfrontZoneId,
        evaluateTargetHealth: true,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_domain_name = aws.apigateway.DomainName("exampleDomainName",
    domain_name="api.example.com",
    certificate_name="example-api",
    certificate_body=(lambda path: open(path).read())(f"{path['module']}/example.com/example.crt"),
    certificate_chain=(lambda path: open(path).read())(f"{path['module']}/example.com/ca.crt"),
    certificate_private_key=(lambda path: open(path).read())(f"{path['module']}/example.com/example.key"))
# Example DNS record using Route53.
# Route53 is not specifically required; any DNS host can be used.
example_record = aws.route53.Record("exampleRecord",
    zone_id=aws_route53_zone["example"]["id"],
    name=example_domain_name.domain_name,
    type="A",
    aliases=[aws.route53.RecordAliasArgs(
        name=example_domain_name.cloudfront_domain_name,
        zone_id=example_domain_name.cloudfront_zone_id,
        evaluate_target_health=True,
    )])
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleDomainName = new Aws.ApiGateway.DomainName("exampleDomainName", new()
    {
        Domain = "api.example.com",
        CertificateName = "example-api",
        CertificateBody = File.ReadAllText($"{path.Module}/example.com/example.crt"),
        CertificateChain = File.ReadAllText($"{path.Module}/example.com/ca.crt"),
        CertificatePrivateKey = File.ReadAllText($"{path.Module}/example.com/example.key"),
    });

    // Example DNS record using Route53.
    // Route53 is not specifically required; any DNS host can be used.
    var exampleRecord = new Aws.Route53.Record("exampleRecord", new()
    {
        ZoneId = aws_route53_zone.Example.Id,
        Name = exampleDomainName.Domain,
        Type = "A",
        Aliases = new[]
        {
            new Aws.Route53.Inputs.RecordAliasArgs
            {
                Name = exampleDomainName.CloudfrontDomainName,
                ZoneId = exampleDomainName.CloudfrontZoneId,
                EvaluateTargetHealth = true,
            },
        },
    });

});
```
```go
package main

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDomainName, err := apigateway.NewDomainName(ctx, "exampleDomainName", &apigateway.DomainNameArgs{
			DomainName:            pulumi.String("api.example.com"),
			CertificateName:       pulumi.String("example-api"),
			CertificateBody:       readFileOrPanic(fmt.Sprintf("%v/example.com/example.crt", path.Module)),
			CertificateChain:      readFileOrPanic(fmt.Sprintf("%v/example.com/ca.crt", path.Module)),
			CertificatePrivateKey: readFileOrPanic(fmt.Sprintf("%v/example.com/example.key", path.Module)),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleRecord", &route53.RecordArgs{
			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
			Name:   exampleDomainName.DomainName,
			Type:   pulumi.String("A"),
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					Name:                 exampleDomainName.CloudfrontDomainName,
					ZoneId:               exampleDomainName.CloudfrontZoneId,
					EvaluateTargetHealth: pulumi.Bool(true),
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
import com.pulumi.aws.apigateway.DomainName;
import com.pulumi.aws.apigateway.DomainNameArgs;
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
        var exampleDomainName = new DomainName("exampleDomainName", DomainNameArgs.builder()        
            .domainName("api.example.com")
            .certificateName("example-api")
            .certificateBody(Files.readString(Paths.get(String.format("%s/example.com/example.crt", path.module()))))
            .certificateChain(Files.readString(Paths.get(String.format("%s/example.com/ca.crt", path.module()))))
            .certificatePrivateKey(Files.readString(Paths.get(String.format("%s/example.com/example.key", path.module()))))
            .build());

        var exampleRecord = new Record("exampleRecord", RecordArgs.builder()        
            .zoneId(aws_route53_zone.example().id())
            .name(exampleDomainName.domainName())
            .type("A")
            .aliases(RecordAliasArgs.builder()
                .name(exampleDomainName.cloudfrontDomainName())
                .zoneId(exampleDomainName.cloudfrontZoneId())
                .evaluateTargetHealth(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleDomainName:
    type: aws:apigateway:DomainName
    properties:
      domainName: api.example.com
      certificateName: example-api
      certificateBody:
        fn::readFile: ${path.module}/example.com/example.crt
      certificateChain:
        fn::readFile: ${path.module}/example.com/ca.crt
      certificatePrivateKey:
        fn::readFile: ${path.module}/example.com/example.key
  # Example DNS record using Route53.
  # Route53 is not specifically required; any DNS host can be used.
  exampleRecord:
    type: aws:route53:Record
    properties:
      zoneId: ${aws_route53_zone.example.id} # See aws_route53_zone for how to create this
      name: ${exampleDomainName.domainName}
      type: A
      aliases:
        - name: ${exampleDomainName.cloudfrontDomainName}
          zoneId: ${exampleDomainName.cloudfrontZoneId}
          evaluateTargetHealth: true
```
{{% /example %}}
{{% example %}}
### Regional (ACM Certificate)

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.apigateway.DomainName;
import com.pulumi.aws.apigateway.DomainNameArgs;
import com.pulumi.aws.apigateway.inputs.DomainNameEndpointConfigurationArgs;
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
        var exampleDomainName = new DomainName("exampleDomainName", DomainNameArgs.builder()        
            .domainName("api.example.com")
            .regionalCertificateArn(aws_acm_certificate_validation.example().certificate_arn())
            .endpointConfiguration(DomainNameEndpointConfigurationArgs.builder()
                .types("REGIONAL")
                .build())
            .build());

        var exampleRecord = new Record("exampleRecord", RecordArgs.builder()        
            .name(exampleDomainName.domainName())
            .type("A")
            .zoneId(aws_route53_zone.example().id())
            .aliases(RecordAliasArgs.builder()
                .evaluateTargetHealth(true)
                .name(exampleDomainName.regionalDomainName())
                .zoneId(exampleDomainName.regionalZoneId())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleDomainName:
    type: aws:apigateway:DomainName
    properties:
      domainName: api.example.com
      regionalCertificateArn: ${aws_acm_certificate_validation.example.certificate_arn}
      endpointConfiguration:
        types:
          - REGIONAL
  # Example DNS record using Route53.
  # Route53 is not specifically required; any DNS host can be used.
  exampleRecord:
    type: aws:route53:Record
    properties:
      name: ${exampleDomainName.domainName}
      type: A
      zoneId: ${aws_route53_zone.example.id}
      aliases:
        - evaluateTargetHealth: true
          name: ${exampleDomainName.regionalDomainName}
          zoneId: ${exampleDomainName.regionalZoneId}
```
{{% /example %}}
{{% example %}}
### Regional (IAM Certificate)

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.apigateway.DomainName;
import com.pulumi.aws.apigateway.DomainNameArgs;
import com.pulumi.aws.apigateway.inputs.DomainNameEndpointConfigurationArgs;
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
        var exampleDomainName = new DomainName("exampleDomainName", DomainNameArgs.builder()        
            .certificateBody(Files.readString(Paths.get(String.format("%s/example.com/example.crt", path.module()))))
            .certificateChain(Files.readString(Paths.get(String.format("%s/example.com/ca.crt", path.module()))))
            .certificatePrivateKey(Files.readString(Paths.get(String.format("%s/example.com/example.key", path.module()))))
            .domainName("api.example.com")
            .regionalCertificateName("example-api")
            .endpointConfiguration(DomainNameEndpointConfigurationArgs.builder()
                .types("REGIONAL")
                .build())
            .build());

        var exampleRecord = new Record("exampleRecord", RecordArgs.builder()        
            .name(exampleDomainName.domainName())
            .type("A")
            .zoneId(aws_route53_zone.example().id())
            .aliases(RecordAliasArgs.builder()
                .evaluateTargetHealth(true)
                .name(exampleDomainName.regionalDomainName())
                .zoneId(exampleDomainName.regionalZoneId())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleDomainName:
    type: aws:apigateway:DomainName
    properties:
      certificateBody:
        fn::readFile: ${path.module}/example.com/example.crt
      certificateChain:
        fn::readFile: ${path.module}/example.com/ca.crt
      certificatePrivateKey:
        fn::readFile: ${path.module}/example.com/example.key
      domainName: api.example.com
      regionalCertificateName: example-api
      endpointConfiguration:
        types:
          - REGIONAL
  # Example DNS record using Route53.
  # Route53 is not specifically required; any DNS host can be used.
  exampleRecord:
    type: aws:route53:Record
    properties:
      name: ${exampleDomainName.domainName}
      type: A
      zoneId: ${aws_route53_zone.example.id}
      aliases:
        - evaluateTargetHealth: true
          name: ${exampleDomainName.regionalDomainName}
          zoneId: ${exampleDomainName.regionalZoneId}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import API Gateway domain names using their `name`. For example:

```sh
 $ pulumi import aws:apigateway/domainName:DomainName example dev.example.com
```
 