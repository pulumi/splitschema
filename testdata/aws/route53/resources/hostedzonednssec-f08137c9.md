Manages Route 53 Hosted Zone Domain Name System Security Extensions (DNSSEC). For more information about managing DNSSEC in Route 53, see the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html).

!> **WARNING:** If you disable DNSSEC signing for your hosted zone before the DNS changes have propagated, your domain could become unavailable on the internet. When you remove the DS records, you must wait until the longest TTL for the DS records that you remove has expired before you complete the step to disable DNSSEC signing. Please refer to the [Route 53 Developer Guide - Disable DNSSEC](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec-disable.html) for a detailed breakdown on the steps required to disable DNSSEC safely for a hosted zone.

> **Note:** Route53 hosted zones are global resources, and as such any `aws.kms.Key` that you use as part of a signing key needs to be located in the `us-east-1` region. In the example below, the main AWS provider declaration is for `us-east-1`, however if you are provisioning your AWS resources in a different region, you will need to specify a provider alias and use that attached to the `aws.kms.Key` resource as described in the provider alias documentation.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const exampleKey = new aws.kms.Key("exampleKey", {
    customerMasterKeySpec: "ECC_NIST_P256",
    deletionWindowInDays: 7,
    keyUsage: "SIGN_VERIFY",
    policy: current.then(current => JSON.stringify({
        Statement: [
            {
                Action: [
                    "kms:DescribeKey",
                    "kms:GetPublicKey",
                    "kms:Sign",
                    "kms:Verify",
                ],
                Effect: "Allow",
                Principal: {
                    Service: "dnssec-route53.amazonaws.com",
                },
                Resource: "*",
                Sid: "Allow Route 53 DNSSEC Service",
            },
            {
                Action: "kms:*",
                Effect: "Allow",
                Principal: {
                    AWS: `arn:aws:iam::${current.accountId}:root`,
                },
                Resource: "*",
                Sid: "Enable IAM User Permissions",
            },
        ],
        Version: "2012-10-17",
    })),
});
const exampleZone = new aws.route53.Zone("exampleZone", {});
const exampleKeySigningKey = new aws.route53.KeySigningKey("exampleKeySigningKey", {
    hostedZoneId: exampleZone.id,
    keyManagementServiceArn: exampleKey.arn,
});
const exampleHostedZoneDnsSec = new aws.route53.HostedZoneDnsSec("exampleHostedZoneDnsSec", {hostedZoneId: exampleKeySigningKey.hostedZoneId}, {
    dependsOn: [exampleKeySigningKey],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

current = aws.get_caller_identity()
example_key = aws.kms.Key("exampleKey",
    customer_master_key_spec="ECC_NIST_P256",
    deletion_window_in_days=7,
    key_usage="SIGN_VERIFY",
    policy=json.dumps({
        "Statement": [
            {
                "Action": [
                    "kms:DescribeKey",
                    "kms:GetPublicKey",
                    "kms:Sign",
                    "kms:Verify",
                ],
                "Effect": "Allow",
                "Principal": {
                    "Service": "dnssec-route53.amazonaws.com",
                },
                "Resource": "*",
                "Sid": "Allow Route 53 DNSSEC Service",
            },
            {
                "Action": "kms:*",
                "Effect": "Allow",
                "Principal": {
                    "AWS": f"arn:aws:iam::{current.account_id}:root",
                },
                "Resource": "*",
                "Sid": "Enable IAM User Permissions",
            },
        ],
        "Version": "2012-10-17",
    }))
example_zone = aws.route53.Zone("exampleZone")
example_key_signing_key = aws.route53.KeySigningKey("exampleKeySigningKey",
    hosted_zone_id=example_zone.id,
    key_management_service_arn=example_key.arn)
example_hosted_zone_dns_sec = aws.route53.HostedZoneDnsSec("exampleHostedZoneDnsSec", hosted_zone_id=example_key_signing_key.hosted_zone_id,
opts=pulumi.ResourceOptions(depends_on=[example_key_signing_key]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        CustomerMasterKeySpec = "ECC_NIST_P256",
        DeletionWindowInDays = 7,
        KeyUsage = "SIGN_VERIFY",
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[]
                    {
                        "kms:DescribeKey",
                        "kms:GetPublicKey",
                        "kms:Sign",
                        "kms:Verify",
                    },
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "dnssec-route53.amazonaws.com",
                    },
                    ["Resource"] = "*",
                    ["Sid"] = "Allow Route 53 DNSSEC Service",
                },
                new Dictionary<string, object?>
                {
                    ["Action"] = "kms:*",
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["AWS"] = $"arn:aws:iam::{current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:root",
                    },
                    ["Resource"] = "*",
                    ["Sid"] = "Enable IAM User Permissions",
                },
            },
            ["Version"] = "2012-10-17",
        }),
    });

    var exampleZone = new Aws.Route53.Zone("exampleZone");

    var exampleKeySigningKey = new Aws.Route53.KeySigningKey("exampleKeySigningKey", new()
    {
        HostedZoneId = exampleZone.Id,
        KeyManagementServiceArn = exampleKey.Arn,
    });

    var exampleHostedZoneDnsSec = new Aws.Route53.HostedZoneDnsSec("exampleHostedZoneDnsSec", new()
    {
        HostedZoneId = exampleKeySigningKey.HostedZoneId,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleKeySigningKey,
        },
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Statement": []interface{}{
				map[string]interface{}{
					"Action": []string{
						"kms:DescribeKey",
						"kms:GetPublicKey",
						"kms:Sign",
						"kms:Verify",
					},
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "dnssec-route53.amazonaws.com",
					},
					"Resource": "*",
					"Sid":      "Allow Route 53 DNSSEC Service",
				},
				map[string]interface{}{
					"Action": "kms:*",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"AWS": fmt.Sprintf("arn:aws:iam::%v:root", current.AccountId),
					},
					"Resource": "*",
					"Sid":      "Enable IAM User Permissions",
				},
			},
			"Version": "2012-10-17",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			CustomerMasterKeySpec: pulumi.String("ECC_NIST_P256"),
			DeletionWindowInDays:  pulumi.Int(7),
			KeyUsage:              pulumi.String("SIGN_VERIFY"),
			Policy:                pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		exampleZone, err := route53.NewZone(ctx, "exampleZone", nil)
		if err != nil {
			return err
		}
		exampleKeySigningKey, err := route53.NewKeySigningKey(ctx, "exampleKeySigningKey", &route53.KeySigningKeyArgs{
			HostedZoneId:            exampleZone.ID(),
			KeyManagementServiceArn: exampleKey.Arn,
		})
		if err != nil {
			return err
		}
		_, err = route53.NewHostedZoneDnsSec(ctx, "exampleHostedZoneDnsSec", &route53.HostedZoneDnsSecArgs{
			HostedZoneId: exampleKeySigningKey.HostedZoneId,
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleKeySigningKey,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.route53.Zone;
import com.pulumi.aws.route53.KeySigningKey;
import com.pulumi.aws.route53.KeySigningKeyArgs;
import com.pulumi.aws.route53.HostedZoneDnsSec;
import com.pulumi.aws.route53.HostedZoneDnsSecArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        final var current = AwsFunctions.getCallerIdentity();

        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .customerMasterKeySpec("ECC_NIST_P256")
            .deletionWindowInDays(7)
            .keyUsage("SIGN_VERIFY")
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Statement", jsonArray(
                        jsonObject(
                            jsonProperty("Action", jsonArray(
                                "kms:DescribeKey", 
                                "kms:GetPublicKey", 
                                "kms:Sign", 
                                "kms:Verify"
                            )),
                            jsonProperty("Effect", "Allow"),
                            jsonProperty("Principal", jsonObject(
                                jsonProperty("Service", "dnssec-route53.amazonaws.com")
                            )),
                            jsonProperty("Resource", "*"),
                            jsonProperty("Sid", "Allow Route 53 DNSSEC Service")
                        ), 
                        jsonObject(
                            jsonProperty("Action", "kms:*"),
                            jsonProperty("Effect", "Allow"),
                            jsonProperty("Principal", jsonObject(
                                jsonProperty("AWS", String.format("arn:aws:iam::%s:root", current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                            )),
                            jsonProperty("Resource", "*"),
                            jsonProperty("Sid", "Enable IAM User Permissions")
                        )
                    )),
                    jsonProperty("Version", "2012-10-17")
                )))
            .build());

        var exampleZone = new Zone("exampleZone");

        var exampleKeySigningKey = new KeySigningKey("exampleKeySigningKey", KeySigningKeyArgs.builder()        
            .hostedZoneId(exampleZone.id())
            .keyManagementServiceArn(exampleKey.arn())
            .build());

        var exampleHostedZoneDnsSec = new HostedZoneDnsSec("exampleHostedZoneDnsSec", HostedZoneDnsSecArgs.builder()        
            .hostedZoneId(exampleKeySigningKey.hostedZoneId())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleKeySigningKey)
                .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      customerMasterKeySpec: ECC_NIST_P256
      deletionWindowInDays: 7
      keyUsage: SIGN_VERIFY
      policy:
        fn::toJSON:
          Statement:
            - Action:
                - kms:DescribeKey
                - kms:GetPublicKey
                - kms:Sign
                - kms:Verify
              Effect: Allow
              Principal:
                Service: dnssec-route53.amazonaws.com
              Resource: '*'
              Sid: Allow Route 53 DNSSEC Service
            - Action: kms:*
              Effect: Allow
              Principal:
                AWS: arn:aws:iam::${current.accountId}:root
              Resource: '*'
              Sid: Enable IAM User Permissions
          Version: 2012-10-17
  exampleZone:
    type: aws:route53:Zone
  exampleKeySigningKey:
    type: aws:route53:KeySigningKey
    properties:
      hostedZoneId: ${exampleZone.id}
      keyManagementServiceArn: ${exampleKey.arn}
  exampleHostedZoneDnsSec:
    type: aws:route53:HostedZoneDnsSec
    properties:
      hostedZoneId: ${exampleKeySigningKey.hostedZoneId}
    options:
      dependson:
        - ${exampleKeySigningKey}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_route53_hosted_zone_dnssec` resources using the Route 53 Hosted Zone identifier. For example:

```sh
 $ pulumi import aws:route53/hostedZoneDnsSec:HostedZoneDnsSec example Z1D633PJN98FT9
```
 