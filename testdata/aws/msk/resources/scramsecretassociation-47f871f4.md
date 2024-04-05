Associates SCRAM secrets stored in the Secrets Manager service with a Managed Streaming for Kafka (MSK) cluster.

> **Note:** The following assumes the MSK cluster has SASL/SCRAM authentication enabled. See below for example usage or refer to the [Username/Password Authentication](https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html) section of the MSK Developer Guide for more details.

To set up username and password authentication for a cluster, create an `aws.secretsmanager.Secret` resource and associate
a username and password with the secret with an `aws.secretsmanager.SecretVersion` resource. When creating a secret for the cluster,
the `name` must have the prefix `AmazonMSK_` and you must either use an existing custom AWS KMS key or create a new
custom AWS KMS key for your secret with the `aws.kms.Key` resource. It is important to note that a policy is required for the `aws.secretsmanager.Secret`
resource in order for Kafka to be able to read it. This policy is attached automatically when the `aws.msk.ScramSecretAssociation` is used,
however, this policy will not be in the state and as such, will present a diff on plan/apply. For that reason, you must use the `aws.secretsmanager.SecretPolicy`
resource](/docs/providers/aws/r/secretsmanager_secret_policy.html) as shown below in order to ensure that the state is in a clean state after the creation of secret and the association to the cluster.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCluster = new aws.msk.Cluster("exampleCluster", {clientAuthentication: {
    sasl: {
        scram: true,
    },
}});
const exampleKey = new aws.kms.Key("exampleKey", {description: "Example Key for MSK Cluster Scram Secret Association"});
const exampleSecret = new aws.secretsmanager.Secret("exampleSecret", {kmsKeyId: exampleKey.keyId});
const exampleSecretVersion = new aws.secretsmanager.SecretVersion("exampleSecretVersion", {
    secretId: exampleSecret.id,
    secretString: JSON.stringify({
        username: "user",
        password: "pass",
    }),
});
const exampleScramSecretAssociation = new aws.msk.ScramSecretAssociation("exampleScramSecretAssociation", {
    clusterArn: exampleCluster.arn,
    secretArnLists: [exampleSecret.arn],
}, {
    dependsOn: [exampleSecretVersion],
});
const examplePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        sid: "AWSKafkaResourcePolicy",
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["kafka.amazonaws.com"],
        }],
        actions: ["secretsmanager:getSecretValue"],
        resources: [exampleSecret.arn],
    }],
});
const exampleSecretPolicy = new aws.secretsmanager.SecretPolicy("exampleSecretPolicy", {
    secretArn: exampleSecret.arn,
    policy: examplePolicyDocument.apply(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_cluster = aws.msk.Cluster("exampleCluster", client_authentication=aws.msk.ClusterClientAuthenticationArgs(
    sasl=aws.msk.ClusterClientAuthenticationSaslArgs(
        scram=True,
    ),
))
example_key = aws.kms.Key("exampleKey", description="Example Key for MSK Cluster Scram Secret Association")
example_secret = aws.secretsmanager.Secret("exampleSecret", kms_key_id=example_key.key_id)
example_secret_version = aws.secretsmanager.SecretVersion("exampleSecretVersion",
    secret_id=example_secret.id,
    secret_string=json.dumps({
        "username": "user",
        "password": "pass",
    }))
example_scram_secret_association = aws.msk.ScramSecretAssociation("exampleScramSecretAssociation",
    cluster_arn=example_cluster.arn,
    secret_arn_lists=[example_secret.arn],
    opts=pulumi.ResourceOptions(depends_on=[example_secret_version]))
example_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="AWSKafkaResourcePolicy",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["kafka.amazonaws.com"],
    )],
    actions=["secretsmanager:getSecretValue"],
    resources=[example_secret.arn],
)])
example_secret_policy = aws.secretsmanager.SecretPolicy("exampleSecretPolicy",
    secret_arn=example_secret.arn,
    policy=example_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleCluster = new Aws.Msk.Cluster("exampleCluster", new()
    {
        ClientAuthentication = new Aws.Msk.Inputs.ClusterClientAuthenticationArgs
        {
            Sasl = new Aws.Msk.Inputs.ClusterClientAuthenticationSaslArgs
            {
                Scram = true,
            },
        },
    });

    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "Example Key for MSK Cluster Scram Secret Association",
    });

    var exampleSecret = new Aws.SecretsManager.Secret("exampleSecret", new()
    {
        KmsKeyId = exampleKey.KeyId,
    });

    var exampleSecretVersion = new Aws.SecretsManager.SecretVersion("exampleSecretVersion", new()
    {
        SecretId = exampleSecret.Id,
        SecretString = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["username"] = "user",
            ["password"] = "pass",
        }),
    });

    var exampleScramSecretAssociation = new Aws.Msk.ScramSecretAssociation("exampleScramSecretAssociation", new()
    {
        ClusterArn = exampleCluster.Arn,
        SecretArnLists = new[]
        {
            exampleSecret.Arn,
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleSecretVersion,
        },
    });

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "AWSKafkaResourcePolicy",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "kafka.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "secretsmanager:getSecretValue",
                },
                Resources = new[]
                {
                    exampleSecret.Arn,
                },
            },
        },
    });

    var exampleSecretPolicy = new Aws.SecretsManager.SecretPolicy("exampleSecretPolicy", new()
    {
        SecretArn = exampleSecret.Arn,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/msk"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleCluster, err := msk.NewCluster(ctx, "exampleCluster", &msk.ClusterArgs{
			ClientAuthentication: &msk.ClusterClientAuthenticationArgs{
				Sasl: &msk.ClusterClientAuthenticationSaslArgs{
					Scram: pulumi.Bool(true),
				},
			},
		})
		if err != nil {
			return err
		}
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description: pulumi.String("Example Key for MSK Cluster Scram Secret Association"),
		})
		if err != nil {
			return err
		}
		exampleSecret, err := secretsmanager.NewSecret(ctx, "exampleSecret", &secretsmanager.SecretArgs{
			KmsKeyId: exampleKey.KeyId,
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"username": "user",
			"password": "pass",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleSecretVersion, err := secretsmanager.NewSecretVersion(ctx, "exampleSecretVersion", &secretsmanager.SecretVersionArgs{
			SecretId:     exampleSecret.ID(),
			SecretString: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = msk.NewScramSecretAssociation(ctx, "exampleScramSecretAssociation", &msk.ScramSecretAssociationArgs{
			ClusterArn: exampleCluster.Arn,
			SecretArnLists: pulumi.StringArray{
				exampleSecret.Arn,
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleSecretVersion,
		}))
		if err != nil {
			return err
		}
		examplePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Sid:    pulumi.String("AWSKafkaResourcePolicy"),
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("Service"),
							Identifiers: pulumi.StringArray{
								pulumi.String("kafka.amazonaws.com"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("secretsmanager:getSecretValue"),
					},
					Resources: pulumi.StringArray{
						exampleSecret.Arn,
					},
				},
			},
		}, nil)
		_, err = secretsmanager.NewSecretPolicy(ctx, "exampleSecretPolicy", &secretsmanager.SecretPolicyArgs{
			SecretArn: exampleSecret.Arn,
			Policy: examplePolicyDocument.ApplyT(func(examplePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &examplePolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
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
import com.pulumi.aws.msk.Cluster;
import com.pulumi.aws.msk.ClusterArgs;
import com.pulumi.aws.msk.inputs.ClusterClientAuthenticationArgs;
import com.pulumi.aws.msk.inputs.ClusterClientAuthenticationSaslArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.secretsmanager.Secret;
import com.pulumi.aws.secretsmanager.SecretArgs;
import com.pulumi.aws.secretsmanager.SecretVersion;
import com.pulumi.aws.secretsmanager.SecretVersionArgs;
import com.pulumi.aws.msk.ScramSecretAssociation;
import com.pulumi.aws.msk.ScramSecretAssociationArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.secretsmanager.SecretPolicy;
import com.pulumi.aws.secretsmanager.SecretPolicyArgs;
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
        var exampleCluster = new Cluster("exampleCluster", ClusterArgs.builder()        
            .clientAuthentication(ClusterClientAuthenticationArgs.builder()
                .sasl(ClusterClientAuthenticationSaslArgs.builder()
                    .scram(true)
                    .build())
                .build())
            .build());

        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("Example Key for MSK Cluster Scram Secret Association")
            .build());

        var exampleSecret = new Secret("exampleSecret", SecretArgs.builder()        
            .kmsKeyId(exampleKey.keyId())
            .build());

        var exampleSecretVersion = new SecretVersion("exampleSecretVersion", SecretVersionArgs.builder()        
            .secretId(exampleSecret.id())
            .secretString(serializeJson(
                jsonObject(
                    jsonProperty("username", "user"),
                    jsonProperty("password", "pass")
                )))
            .build());

        var exampleScramSecretAssociation = new ScramSecretAssociation("exampleScramSecretAssociation", ScramSecretAssociationArgs.builder()        
            .clusterArn(exampleCluster.arn())
            .secretArnLists(exampleSecret.arn())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleSecretVersion)
                .build());

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("AWSKafkaResourcePolicy")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("kafka.amazonaws.com")
                    .build())
                .actions("secretsmanager:getSecretValue")
                .resources(exampleSecret.arn())
                .build())
            .build());

        var exampleSecretPolicy = new SecretPolicy("exampleSecretPolicy", SecretPolicyArgs.builder()        
            .secretArn(exampleSecret.arn())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(examplePolicyDocument -> examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  exampleScramSecretAssociation:
    type: aws:msk:ScramSecretAssociation
    properties:
      clusterArn: ${exampleCluster.arn}
      secretArnLists:
        - ${exampleSecret.arn}
    options:
      dependson:
        - ${exampleSecretVersion}
  exampleCluster:
    type: aws:msk:Cluster
    properties:
      clientAuthentication:
        sasl:
          scram: true
  exampleSecret:
    type: aws:secretsmanager:Secret
    properties:
      kmsKeyId: ${exampleKey.keyId}
  exampleKey:
    type: aws:kms:Key
    properties:
      description: Example Key for MSK Cluster Scram Secret Association
  exampleSecretVersion:
    type: aws:secretsmanager:SecretVersion
    properties:
      secretId: ${exampleSecret.id}
      secretString:
        fn::toJSON:
          username: user
          password: pass
  exampleSecretPolicy:
    type: aws:secretsmanager:SecretPolicy
    properties:
      secretArn: ${exampleSecret.arn}
      policy: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: AWSKafkaResourcePolicy
            effect: Allow
            principals:
              - type: Service
                identifiers:
                  - kafka.amazonaws.com
            actions:
              - secretsmanager:getSecretValue
            resources:
              - ${exampleSecret.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MSK SCRAM Secret Associations using the `id`. For example:

```sh
 $ pulumi import aws:msk/scramSecretAssociation:ScramSecretAssociation example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
```
 