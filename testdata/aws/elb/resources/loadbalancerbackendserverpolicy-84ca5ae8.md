Attaches a load balancer policy to an ELB backend server.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const wu_tang = new aws.elb.LoadBalancer("wu-tang", {
    availabilityZones: ["us-east-1a"],
    listeners: [{
        instancePort: 443,
        instanceProtocol: "http",
        lbPort: 443,
        lbProtocol: "https",
        sslCertificateId: "arn:aws:iam::000000000000:server-certificate/wu-tang.net",
    }],
    tags: {
        Name: "wu-tang",
    },
});
const wu_tang_ca_pubkey_policy = new aws.elb.LoadBalancerPolicy("wu-tang-ca-pubkey-policy", {
    loadBalancerName: wu_tang.name,
    policyName: "wu-tang-ca-pubkey-policy",
    policyTypeName: "PublicKeyPolicyType",
    policyAttributes: [{
        name: "PublicKey",
        value: fs.readFileSync("wu-tang-pubkey", "utf8"),
    }],
});
const wu_tang_root_ca_backend_auth_policy = new aws.elb.LoadBalancerPolicy("wu-tang-root-ca-backend-auth-policy", {
    loadBalancerName: wu_tang.name,
    policyName: "wu-tang-root-ca-backend-auth-policy",
    policyTypeName: "BackendServerAuthenticationPolicyType",
    policyAttributes: [{
        name: "PublicKeyPolicyName",
        value: aws_load_balancer_policy["wu-tang-root-ca-pubkey-policy"].policy_name,
    }],
});
const wu_tang_backend_auth_policies_443 = new aws.elb.LoadBalancerBackendServerPolicy("wu-tang-backend-auth-policies-443", {
    loadBalancerName: wu_tang.name,
    instancePort: 443,
    policyNames: [wu_tang_root_ca_backend_auth_policy.policyName],
});
```
```python
import pulumi
import pulumi_aws as aws

wu_tang = aws.elb.LoadBalancer("wu-tang",
    availability_zones=["us-east-1a"],
    listeners=[aws.elb.LoadBalancerListenerArgs(
        instance_port=443,
        instance_protocol="http",
        lb_port=443,
        lb_protocol="https",
        ssl_certificate_id="arn:aws:iam::000000000000:server-certificate/wu-tang.net",
    )],
    tags={
        "Name": "wu-tang",
    })
wu_tang_ca_pubkey_policy = aws.elb.LoadBalancerPolicy("wu-tang-ca-pubkey-policy",
    load_balancer_name=wu_tang.name,
    policy_name="wu-tang-ca-pubkey-policy",
    policy_type_name="PublicKeyPolicyType",
    policy_attributes=[aws.elb.LoadBalancerPolicyPolicyAttributeArgs(
        name="PublicKey",
        value=(lambda path: open(path).read())("wu-tang-pubkey"),
    )])
wu_tang_root_ca_backend_auth_policy = aws.elb.LoadBalancerPolicy("wu-tang-root-ca-backend-auth-policy",
    load_balancer_name=wu_tang.name,
    policy_name="wu-tang-root-ca-backend-auth-policy",
    policy_type_name="BackendServerAuthenticationPolicyType",
    policy_attributes=[aws.elb.LoadBalancerPolicyPolicyAttributeArgs(
        name="PublicKeyPolicyName",
        value=aws_load_balancer_policy["wu-tang-root-ca-pubkey-policy"]["policy_name"],
    )])
wu_tang_backend_auth_policies_443 = aws.elb.LoadBalancerBackendServerPolicy("wu-tang-backend-auth-policies-443",
    load_balancer_name=wu_tang.name,
    instance_port=443,
    policy_names=[wu_tang_root_ca_backend_auth_policy.policy_name])
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var wu_tang = new Aws.Elb.LoadBalancer("wu-tang", new()
    {
        AvailabilityZones = new[]
        {
            "us-east-1a",
        },
        Listeners = new[]
        {
            new Aws.Elb.Inputs.LoadBalancerListenerArgs
            {
                InstancePort = 443,
                InstanceProtocol = "http",
                LbPort = 443,
                LbProtocol = "https",
                SslCertificateId = "arn:aws:iam::000000000000:server-certificate/wu-tang.net",
            },
        },
        Tags = 
        {
            { "Name", "wu-tang" },
        },
    });

    var wu_tang_ca_pubkey_policy = new Aws.Elb.LoadBalancerPolicy("wu-tang-ca-pubkey-policy", new()
    {
        LoadBalancerName = wu_tang.Name,
        PolicyName = "wu-tang-ca-pubkey-policy",
        PolicyTypeName = "PublicKeyPolicyType",
        PolicyAttributes = new[]
        {
            new Aws.Elb.Inputs.LoadBalancerPolicyPolicyAttributeArgs
            {
                Name = "PublicKey",
                Value = File.ReadAllText("wu-tang-pubkey"),
            },
        },
    });

    var wu_tang_root_ca_backend_auth_policy = new Aws.Elb.LoadBalancerPolicy("wu-tang-root-ca-backend-auth-policy", new()
    {
        LoadBalancerName = wu_tang.Name,
        PolicyName = "wu-tang-root-ca-backend-auth-policy",
        PolicyTypeName = "BackendServerAuthenticationPolicyType",
        PolicyAttributes = new[]
        {
            new Aws.Elb.Inputs.LoadBalancerPolicyPolicyAttributeArgs
            {
                Name = "PublicKeyPolicyName",
                Value = aws_load_balancer_policy.Wu_tang_root_ca_pubkey_policy.Policy_name,
            },
        },
    });

    var wu_tang_backend_auth_policies_443 = new Aws.Elb.LoadBalancerBackendServerPolicy("wu-tang-backend-auth-policies-443", new()
    {
        LoadBalancerName = wu_tang.Name,
        InstancePort = 443,
        PolicyNames = new[]
        {
            wu_tang_root_ca_backend_auth_policy.PolicyName,
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
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
		_, err := elb.NewLoadBalancer(ctx, "wu-tang", &elb.LoadBalancerArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(443),
					InstanceProtocol: pulumi.String("http"),
					LbPort:           pulumi.Int(443),
					LbProtocol:       pulumi.String("https"),
					SslCertificateId: pulumi.String("arn:aws:iam::000000000000:server-certificate/wu-tang.net"),
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("wu-tang"),
			},
		})
		if err != nil {
			return err
		}
		_, err = elb.NewLoadBalancerPolicy(ctx, "wu-tang-ca-pubkey-policy", &elb.LoadBalancerPolicyArgs{
			LoadBalancerName: wu_tang.Name,
			PolicyName:       pulumi.String("wu-tang-ca-pubkey-policy"),
			PolicyTypeName:   pulumi.String("PublicKeyPolicyType"),
			PolicyAttributes: elb.LoadBalancerPolicyPolicyAttributeArray{
				&elb.LoadBalancerPolicyPolicyAttributeArgs{
					Name:  pulumi.String("PublicKey"),
					Value: readFileOrPanic("wu-tang-pubkey"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = elb.NewLoadBalancerPolicy(ctx, "wu-tang-root-ca-backend-auth-policy", &elb.LoadBalancerPolicyArgs{
			LoadBalancerName: wu_tang.Name,
			PolicyName:       pulumi.String("wu-tang-root-ca-backend-auth-policy"),
			PolicyTypeName:   pulumi.String("BackendServerAuthenticationPolicyType"),
			PolicyAttributes: elb.LoadBalancerPolicyPolicyAttributeArray{
				&elb.LoadBalancerPolicyPolicyAttributeArgs{
					Name:  pulumi.String("PublicKeyPolicyName"),
					Value: pulumi.Any(aws_load_balancer_policy.WuTangRootCaPubkeyPolicy.Policy_name),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = elb.NewLoadBalancerBackendServerPolicy(ctx, "wu-tang-backend-auth-policies-443", &elb.LoadBalancerBackendServerPolicyArgs{
			LoadBalancerName: wu_tang.Name,
			InstancePort:     pulumi.Int(443),
			PolicyNames: pulumi.StringArray{
				wu_tang_root_ca_backend_auth_policy.PolicyName,
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
import com.pulumi.aws.elb.LoadBalancer;
import com.pulumi.aws.elb.LoadBalancerArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerListenerArgs;
import com.pulumi.aws.elb.LoadBalancerPolicy;
import com.pulumi.aws.elb.LoadBalancerPolicyArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerPolicyPolicyAttributeArgs;
import com.pulumi.aws.elb.LoadBalancerBackendServerPolicy;
import com.pulumi.aws.elb.LoadBalancerBackendServerPolicyArgs;
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
        var wu_tang = new LoadBalancer("wu-tang", LoadBalancerArgs.builder()        
            .availabilityZones("us-east-1a")
            .listeners(LoadBalancerListenerArgs.builder()
                .instancePort(443)
                .instanceProtocol("http")
                .lbPort(443)
                .lbProtocol("https")
                .sslCertificateId("arn:aws:iam::000000000000:server-certificate/wu-tang.net")
                .build())
            .tags(Map.of("Name", "wu-tang"))
            .build());

        var wu_tang_ca_pubkey_policy = new LoadBalancerPolicy("wu-tang-ca-pubkey-policy", LoadBalancerPolicyArgs.builder()        
            .loadBalancerName(wu_tang.name())
            .policyName("wu-tang-ca-pubkey-policy")
            .policyTypeName("PublicKeyPolicyType")
            .policyAttributes(LoadBalancerPolicyPolicyAttributeArgs.builder()
                .name("PublicKey")
                .value(Files.readString(Paths.get("wu-tang-pubkey")))
                .build())
            .build());

        var wu_tang_root_ca_backend_auth_policy = new LoadBalancerPolicy("wu-tang-root-ca-backend-auth-policy", LoadBalancerPolicyArgs.builder()        
            .loadBalancerName(wu_tang.name())
            .policyName("wu-tang-root-ca-backend-auth-policy")
            .policyTypeName("BackendServerAuthenticationPolicyType")
            .policyAttributes(LoadBalancerPolicyPolicyAttributeArgs.builder()
                .name("PublicKeyPolicyName")
                .value(aws_load_balancer_policy.wu-tang-root-ca-pubkey-policy().policy_name())
                .build())
            .build());

        var wu_tang_backend_auth_policies_443 = new LoadBalancerBackendServerPolicy("wu-tang-backend-auth-policies-443", LoadBalancerBackendServerPolicyArgs.builder()        
            .loadBalancerName(wu_tang.name())
            .instancePort(443)
            .policyNames(wu_tang_root_ca_backend_auth_policy.policyName())
            .build());

    }
}
```
```yaml
resources:
  wu-tang:
    type: aws:elb:LoadBalancer
    properties:
      availabilityZones:
        - us-east-1a
      listeners:
        - instancePort: 443
          instanceProtocol: http
          lbPort: 443
          lbProtocol: https
          sslCertificateId: arn:aws:iam::000000000000:server-certificate/wu-tang.net
      tags:
        Name: wu-tang
  wu-tang-ca-pubkey-policy:
    type: aws:elb:LoadBalancerPolicy
    properties:
      loadBalancerName: ${["wu-tang"].name}
      policyName: wu-tang-ca-pubkey-policy
      policyTypeName: PublicKeyPolicyType
      # The public key of a CA certificate file can be extracted with:
      #   # $ cat wu-tang-ca.pem | openssl x509 -pubkey -noout | grep -v '\-\-\-\-' | tr -d '\n' > wu-tang-pubkey
      policyAttributes:
        - name: PublicKey
          value:
            fn::readFile: wu-tang-pubkey
  wu-tang-root-ca-backend-auth-policy:
    type: aws:elb:LoadBalancerPolicy
    properties:
      loadBalancerName: ${["wu-tang"].name}
      policyName: wu-tang-root-ca-backend-auth-policy
      policyTypeName: BackendServerAuthenticationPolicyType
      policyAttributes:
        - name: PublicKeyPolicyName
          value: ${aws_load_balancer_policy"wu-tang-root-ca-pubkey-policy"[%!s(MISSING)].policy_name}
  wu-tang-backend-auth-policies-443:
    type: aws:elb:LoadBalancerBackendServerPolicy
    properties:
      loadBalancerName: ${["wu-tang"].name}
      instancePort: 443
      policyNames:
        - ${["wu-tang-root-ca-backend-auth-policy"].policyName}
```
{{% /example %}}
{{% /examples %}}