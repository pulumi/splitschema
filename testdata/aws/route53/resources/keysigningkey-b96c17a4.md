Manages a Route 53 Key Signing Key. To manage Domain Name System Security Extensions (DNSSEC) for a Hosted Zone, see the `aws.route53.HostedZoneDnsSec` resource. For more information about managing DNSSEC in Route 53, see the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html).

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
    policy: Promise.all([current, current]).then(([current, current1]) => JSON.stringify({
        Statement: [
            {
                Action: [
                    "kms:DescribeKey",
                    "kms:GetPublicKey",
                    "kms:Sign",
                ],
                Effect: "Allow",
                Principal: {
                    Service: "dnssec-route53.amazonaws.com",
                },
                Sid: "Allow Route 53 DNSSEC Service",
                Resource: "*",
                Condition: {
                    StringEquals: {
                        "aws:SourceAccount": current.accountId,
                    },
                    ArnLike: {
                        "aws:SourceArn": "arn:aws:route53:::hostedzone/*",
                    },
                },
            },
            {
                Action: "kms:CreateGrant",
                Effect: "Allow",
                Principal: {
                    Service: "dnssec-route53.amazonaws.com",
                },
                Sid: "Allow Route 53 DNSSEC Service to CreateGrant",
                Resource: "*",
                Condition: {
                    Bool: {
                        "kms:GrantIsForAWSResource": "true",
                    },
                },
            },
            {
                Action: "kms:*",
                Effect: "Allow",
                Principal: {
                    AWS: `arn:aws:iam::${current1.accountId}:root`,
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
    hostedZoneId: aws_route53_zone.test.id,
    keyManagementServiceArn: aws_kms_key.test.arn,
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
                ],
                "Effect": "Allow",
                "Principal": {
                    "Service": "dnssec-route53.amazonaws.com",
                },
                "Sid": "Allow Route 53 DNSSEC Service",
                "Resource": "*",
                "Condition": {
                    "StringEquals": {
                        "aws:SourceAccount": current.account_id,
                    },
                    "ArnLike": {
                        "aws:SourceArn": "arn:aws:route53:::hostedzone/*",
                    },
                },
            },
            {
                "Action": "kms:CreateGrant",
                "Effect": "Allow",
                "Principal": {
                    "Service": "dnssec-route53.amazonaws.com",
                },
                "Sid": "Allow Route 53 DNSSEC Service to CreateGrant",
                "Resource": "*",
                "Condition": {
                    "Bool": {
                        "kms:GrantIsForAWSResource": "true",
                    },
                },
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
    hosted_zone_id=aws_route53_zone["test"]["id"],
    key_management_service_arn=aws_kms_key["test"]["arn"])
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
        Policy = Output.Tuple(current, current).Apply(values =>
        {
            var current = values.Item1;
            var current1 = values.Item2;
            return JsonSerializer.Serialize(new Dictionary<string, object?>
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
                        },
                        ["Effect"] = "Allow",
                        ["Principal"] = new Dictionary<string, object?>
                        {
                            ["Service"] = "dnssec-route53.amazonaws.com",
                        },
                        ["Sid"] = "Allow Route 53 DNSSEC Service",
                        ["Resource"] = "*",
                        ["Condition"] = new Dictionary<string, object?>
                        {
                            ["StringEquals"] = new Dictionary<string, object?>
                            {
                                ["aws:SourceAccount"] = current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
                            },
                            ["ArnLike"] = new Dictionary<string, object?>
                            {
                                ["aws:SourceArn"] = "arn:aws:route53:::hostedzone/*",
                            },
                        },
                    },
                    new Dictionary<string, object?>
                    {
                        ["Action"] = "kms:CreateGrant",
                        ["Effect"] = "Allow",
                        ["Principal"] = new Dictionary<string, object?>
                        {
                            ["Service"] = "dnssec-route53.amazonaws.com",
                        },
                        ["Sid"] = "Allow Route 53 DNSSEC Service to CreateGrant",
                        ["Resource"] = "*",
                        ["Condition"] = new Dictionary<string, object?>
                        {
                            ["Bool"] = new Dictionary<string, object?>
                            {
                                ["kms:GrantIsForAWSResource"] = "true",
                            },
                        },
                    },
                    new Dictionary<string, object?>
                    {
                        ["Action"] = "kms:*",
                        ["Effect"] = "Allow",
                        ["Principal"] = new Dictionary<string, object?>
                        {
                            ["AWS"] = $"arn:aws:iam::{current1.AccountId}:root",
                        },
                        ["Resource"] = "*",
                        ["Sid"] = "Enable IAM User Permissions",
                    },
                },
                ["Version"] = "2012-10-17",
            });
        }),
    });

    var exampleZone = new Aws.Route53.Zone("exampleZone");

    var exampleKeySigningKey = new Aws.Route53.KeySigningKey("exampleKeySigningKey", new()
    {
        HostedZoneId = aws_route53_zone.Test.Id,
        KeyManagementServiceArn = aws_kms_key.Test.Arn,
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
					},
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "dnssec-route53.amazonaws.com",
					},
					"Sid":      "Allow Route 53 DNSSEC Service",
					"Resource": "*",
					"Condition": map[string]interface{}{
						"StringEquals": map[string]interface{}{
							"aws:SourceAccount": current.AccountId,
						},
						"ArnLike": map[string]interface{}{
							"aws:SourceArn": "arn:aws:route53:::hostedzone/*",
						},
					},
				},
				map[string]interface{}{
					"Action": "kms:CreateGrant",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "dnssec-route53.amazonaws.com",
					},
					"Sid":      "Allow Route 53 DNSSEC Service to CreateGrant",
					"Resource": "*",
					"Condition": map[string]interface{}{
						"Bool": map[string]interface{}{
							"kms:GrantIsForAWSResource": "true",
						},
					},
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
		_, err = kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			CustomerMasterKeySpec: pulumi.String("ECC_NIST_P256"),
			DeletionWindowInDays:  pulumi.Int(7),
			KeyUsage:              pulumi.String("SIGN_VERIFY"),
			Policy:                pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewZone(ctx, "exampleZone", nil)
		if err != nil {
			return err
		}
		exampleKeySigningKey, err := route53.NewKeySigningKey(ctx, "exampleKeySigningKey", &route53.KeySigningKeyArgs{
			HostedZoneId:            pulumi.Any(aws_route53_zone.Test.Id),
			KeyManagementServiceArn: pulumi.Any(aws_kms_key.Test.Arn),
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
                                "kms:Sign"
                            )),
                            jsonProperty("Effect", "Allow"),
                            jsonProperty("Principal", jsonObject(
                                jsonProperty("Service", "dnssec-route53.amazonaws.com")
                            )),
                            jsonProperty("Sid", "Allow Route 53 DNSSEC Service"),
                            jsonProperty("Resource", "*"),
                            jsonProperty("Condition", jsonObject(
                                jsonProperty("StringEquals", jsonObject(
                                    jsonProperty("aws:SourceAccount", current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                                )),
                                jsonProperty("ArnLike", jsonObject(
                                    jsonProperty("aws:SourceArn", "arn:aws:route53:::hostedzone/*")
                                ))
                            ))
                        ), 
                        jsonObject(
                            jsonProperty("Action", "kms:CreateGrant"),
                            jsonProperty("Effect", "Allow"),
                            jsonProperty("Principal", jsonObject(
                                jsonProperty("Service", "dnssec-route53.amazonaws.com")
                            )),
                            jsonProperty("Sid", "Allow Route 53 DNSSEC Service to CreateGrant"),
                            jsonProperty("Resource", "*"),
                            jsonProperty("Condition", jsonObject(
                                jsonProperty("Bool", jsonObject(
                                    jsonProperty("kms:GrantIsForAWSResource", "true")
                                ))
                            ))
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
            .hostedZoneId(aws_route53_zone.test().id())
            .keyManagementServiceArn(aws_kms_key.test().arn())
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
              Effect: Allow
              Principal:
                Service: dnssec-route53.amazonaws.com
              Sid: Allow Route 53 DNSSEC Service
              Resource: '*'
              Condition:
                StringEquals:
                  aws:SourceAccount: ${current.accountId}
                ArnLike:
                  aws:SourceArn: arn:aws:route53:::hostedzone/*
            - Action: kms:CreateGrant
              Effect: Allow
              Principal:
                Service: dnssec-route53.amazonaws.com
              Sid: Allow Route 53 DNSSEC Service to CreateGrant
              Resource: '*'
              Condition:
                Bool:
                  kms:GrantIsForAWSResource: 'true'
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
      hostedZoneId: ${aws_route53_zone.test.id}
      keyManagementServiceArn: ${aws_kms_key.test.arn}
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

Using `pulumi import`, import `aws_route53_key_signing_key` resources using the Route 53 Hosted Zone identifier and KMS Key identifier, separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:route53/keySigningKey:KeySigningKey example Z1D633PJN98FT9,example
```
 