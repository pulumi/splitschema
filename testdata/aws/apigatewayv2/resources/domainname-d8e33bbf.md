Manages an Amazon API Gateway Version 2 domain name.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).

> **Note:** This resource establishes ownership of and the TLS settings for
a particular domain name. An API stage can be associated with the domain name using the `aws.apigatewayv2.ApiMapping` resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.DomainName("example", {
    domainName: "ws-api.example.com",
    domainNameConfiguration: {
        certificateArn: aws_acm_certificate.example.arn,
        endpointType: "REGIONAL",
        securityPolicy: "TLS_1_2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.DomainName("example",
    domain_name="ws-api.example.com",
    domain_name_configuration=aws.apigatewayv2.DomainNameDomainNameConfigurationArgs(
        certificate_arn=aws_acm_certificate["example"]["arn"],
        endpoint_type="REGIONAL",
        security_policy="TLS_1_2",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.DomainName("example", new()
    {
        Domain = "ws-api.example.com",
        DomainNameConfiguration = new Aws.ApiGatewayV2.Inputs.DomainNameDomainNameConfigurationArgs
        {
            CertificateArn = aws_acm_certificate.Example.Arn,
            EndpointType = "REGIONAL",
            SecurityPolicy = "TLS_1_2",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.NewDomainName(ctx, "example", &apigatewayv2.DomainNameArgs{
			DomainName: pulumi.String("ws-api.example.com"),
			DomainNameConfiguration: &apigatewayv2.DomainNameDomainNameConfigurationArgs{
				CertificateArn: pulumi.Any(aws_acm_certificate.Example.Arn),
				EndpointType:   pulumi.String("REGIONAL"),
				SecurityPolicy: pulumi.String("TLS_1_2"),
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
import com.pulumi.aws.apigatewayv2.DomainName;
import com.pulumi.aws.apigatewayv2.DomainNameArgs;
import com.pulumi.aws.apigatewayv2.inputs.DomainNameDomainNameConfigurationArgs;
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
        var example = new DomainName("example", DomainNameArgs.builder()        
            .domainName("ws-api.example.com")
            .domainNameConfiguration(DomainNameDomainNameConfigurationArgs.builder()
                .certificateArn(aws_acm_certificate.example().arn())
                .endpointType("REGIONAL")
                .securityPolicy("TLS_1_2")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:DomainName
    properties:
      domainName: ws-api.example.com
      domainNameConfiguration:
        certificateArn: ${aws_acm_certificate.example.arn}
        endpointType: REGIONAL
        securityPolicy: TLS_1_2
```
{{% /example %}}
{{% example %}}
### Associated Route 53 Resource Record

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDomainName = new aws.apigatewayv2.DomainName("exampleDomainName", {
    domainName: "http-api.example.com",
    domainNameConfiguration: {
        certificateArn: aws_acm_certificate.example.arn,
        endpointType: "REGIONAL",
        securityPolicy: "TLS_1_2",
    },
});
const exampleRecord = new aws.route53.Record("exampleRecord", {
    name: exampleDomainName.domainName,
    type: "A",
    zoneId: aws_route53_zone.example.zone_id,
    aliases: [{
        name: exampleDomainName.domainNameConfiguration.apply(domainNameConfiguration => domainNameConfiguration.targetDomainName),
        zoneId: exampleDomainName.domainNameConfiguration.apply(domainNameConfiguration => domainNameConfiguration.hostedZoneId),
        evaluateTargetHealth: false,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_domain_name = aws.apigatewayv2.DomainName("exampleDomainName",
    domain_name="http-api.example.com",
    domain_name_configuration=aws.apigatewayv2.DomainNameDomainNameConfigurationArgs(
        certificate_arn=aws_acm_certificate["example"]["arn"],
        endpoint_type="REGIONAL",
        security_policy="TLS_1_2",
    ))
example_record = aws.route53.Record("exampleRecord",
    name=example_domain_name.domain_name,
    type="A",
    zone_id=aws_route53_zone["example"]["zone_id"],
    aliases=[aws.route53.RecordAliasArgs(
        name=example_domain_name.domain_name_configuration.target_domain_name,
        zone_id=example_domain_name.domain_name_configuration.hosted_zone_id,
        evaluate_target_health=False,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleDomainName = new Aws.ApiGatewayV2.DomainName("exampleDomainName", new()
    {
        Domain = "http-api.example.com",
        DomainNameConfiguration = new Aws.ApiGatewayV2.Inputs.DomainNameDomainNameConfigurationArgs
        {
            CertificateArn = aws_acm_certificate.Example.Arn,
            EndpointType = "REGIONAL",
            SecurityPolicy = "TLS_1_2",
        },
    });

    var exampleRecord = new Aws.Route53.Record("exampleRecord", new()
    {
        Name = exampleDomainName.Domain,
        Type = "A",
        ZoneId = aws_route53_zone.Example.Zone_id,
        Aliases = new[]
        {
            new Aws.Route53.Inputs.RecordAliasArgs
            {
                Name = exampleDomainName.DomainNameConfiguration.Apply(domainNameConfiguration => domainNameConfiguration.TargetDomainName),
                ZoneId = exampleDomainName.DomainNameConfiguration.Apply(domainNameConfiguration => domainNameConfiguration.HostedZoneId),
                EvaluateTargetHealth = false,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDomainName, err := apigatewayv2.NewDomainName(ctx, "exampleDomainName", &apigatewayv2.DomainNameArgs{
			DomainName: pulumi.String("http-api.example.com"),
			DomainNameConfiguration: &apigatewayv2.DomainNameDomainNameConfigurationArgs{
				CertificateArn: pulumi.Any(aws_acm_certificate.Example.Arn),
				EndpointType:   pulumi.String("REGIONAL"),
				SecurityPolicy: pulumi.String("TLS_1_2"),
			},
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleRecord", &route53.RecordArgs{
			Name:   exampleDomainName.DomainName,
			Type:   pulumi.String("A"),
			ZoneId: pulumi.Any(aws_route53_zone.Example.Zone_id),
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					Name: exampleDomainName.DomainNameConfiguration.ApplyT(func(domainNameConfiguration apigatewayv2.DomainNameDomainNameConfiguration) (*string, error) {
						return &domainNameConfiguration.TargetDomainName, nil
					}).(pulumi.StringPtrOutput),
					ZoneId: exampleDomainName.DomainNameConfiguration.ApplyT(func(domainNameConfiguration apigatewayv2.DomainNameDomainNameConfiguration) (*string, error) {
						return &domainNameConfiguration.HostedZoneId, nil
					}).(pulumi.StringPtrOutput),
					EvaluateTargetHealth: pulumi.Bool(false),
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
import com.pulumi.aws.apigatewayv2.DomainName;
import com.pulumi.aws.apigatewayv2.DomainNameArgs;
import com.pulumi.aws.apigatewayv2.inputs.DomainNameDomainNameConfigurationArgs;
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
            .domainName("http-api.example.com")
            .domainNameConfiguration(DomainNameDomainNameConfigurationArgs.builder()
                .certificateArn(aws_acm_certificate.example().arn())
                .endpointType("REGIONAL")
                .securityPolicy("TLS_1_2")
                .build())
            .build());

        var exampleRecord = new Record("exampleRecord", RecordArgs.builder()        
            .name(exampleDomainName.domainName())
            .type("A")
            .zoneId(aws_route53_zone.example().zone_id())
            .aliases(RecordAliasArgs.builder()
                .name(exampleDomainName.domainNameConfiguration().applyValue(domainNameConfiguration -> domainNameConfiguration.targetDomainName()))
                .zoneId(exampleDomainName.domainNameConfiguration().applyValue(domainNameConfiguration -> domainNameConfiguration.hostedZoneId()))
                .evaluateTargetHealth(false)
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleDomainName:
    type: aws:apigatewayv2:DomainName
    properties:
      domainName: http-api.example.com
      domainNameConfiguration:
        certificateArn: ${aws_acm_certificate.example.arn}
        endpointType: REGIONAL
        securityPolicy: TLS_1_2
  exampleRecord:
    type: aws:route53:Record
    properties:
      name: ${exampleDomainName.domainName}
      type: A
      zoneId: ${aws_route53_zone.example.zone_id}
      aliases:
        - name: ${exampleDomainName.domainNameConfiguration.targetDomainName}
          zoneId: ${exampleDomainName.domainNameConfiguration.hostedZoneId}
          evaluateTargetHealth: false
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_apigatewayv2_domain_name` using the domain name. For example:

```sh
 $ pulumi import aws:apigatewayv2/domainName:DomainName example ws-api.example.com
```
 