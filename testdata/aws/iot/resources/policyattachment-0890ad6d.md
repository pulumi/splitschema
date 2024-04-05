Provides an IoT policy attachment.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const pubsubPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: ["iot:*"],
        resources: ["*"],
    }],
});
const pubsubPolicy = new aws.iot.Policy("pubsubPolicy", {policy: pubsubPolicyDocument.then(pubsubPolicyDocument => pubsubPolicyDocument.json)});
const cert = new aws.iot.Certificate("cert", {
    csr: fs.readFileSync("csr.pem", "utf8"),
    active: true,
});
const att = new aws.iot.PolicyAttachment("att", {
    policy: pubsubPolicy.name,
    target: cert.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

pubsub_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["iot:*"],
    resources=["*"],
)])
pubsub_policy = aws.iot.Policy("pubsubPolicy", policy=pubsub_policy_document.json)
cert = aws.iot.Certificate("cert",
    csr=(lambda path: open(path).read())("csr.pem"),
    active=True)
att = aws.iot.PolicyAttachment("att",
    policy=pubsub_policy.name,
    target=cert.arn)
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var pubsubPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "iot:*",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var pubsubPolicy = new Aws.Iot.Policy("pubsubPolicy", new()
    {
        PolicyDocument = pubsubPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var cert = new Aws.Iot.Certificate("cert", new()
    {
        Csr = File.ReadAllText("csr.pem"),
        Active = true,
    });

    var att = new Aws.Iot.PolicyAttachment("att", new()
    {
        Policy = pubsubPolicy.Name,
        Target = cert.Arn,
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
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
		pubsubPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"iot:*",
					},
					Resources: []string{
						"*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		pubsubPolicy, err := iot.NewPolicy(ctx, "pubsubPolicy", &iot.PolicyArgs{
			Policy: *pulumi.String(pubsubPolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		cert, err := iot.NewCertificate(ctx, "cert", &iot.CertificateArgs{
			Csr:    readFileOrPanic("csr.pem"),
			Active: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = iot.NewPolicyAttachment(ctx, "att", &iot.PolicyAttachmentArgs{
			Policy: pubsubPolicy.Name,
			Target: cert.Arn,
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iot.Policy;
import com.pulumi.aws.iot.PolicyArgs;
import com.pulumi.aws.iot.Certificate;
import com.pulumi.aws.iot.CertificateArgs;
import com.pulumi.aws.iot.PolicyAttachment;
import com.pulumi.aws.iot.PolicyAttachmentArgs;
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
        final var pubsubPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("iot:*")
                .resources("*")
                .build())
            .build());

        var pubsubPolicy = new Policy("pubsubPolicy", PolicyArgs.builder()        
            .policy(pubsubPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var cert = new Certificate("cert", CertificateArgs.builder()        
            .csr(Files.readString(Paths.get("csr.pem")))
            .active(true)
            .build());

        var att = new PolicyAttachment("att", PolicyAttachmentArgs.builder()        
            .policy(pubsubPolicy.name())
            .target(cert.arn())
            .build());

    }
}
```
```yaml
resources:
  pubsubPolicy:
    type: aws:iot:Policy
    properties:
      policy: ${pubsubPolicyDocument.json}
  cert:
    type: aws:iot:Certificate
    properties:
      csr:
        fn::readFile: csr.pem
      active: true
  att:
    type: aws:iot:PolicyAttachment
    properties:
      policy: ${pubsubPolicy.name}
      target: ${cert.arn}
variables:
  pubsubPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - iot:*
            resources:
              - '*'
```
{{% /example %}}
{{% /examples %}}