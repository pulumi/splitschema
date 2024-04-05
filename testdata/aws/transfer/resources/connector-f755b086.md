Provides a AWS Transfer AS2 Connector resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Connector("example", {
    accessRole: aws_iam_role.test.arn,
    as2Config: {
        compression: "DISABLED",
        encryptionAlgorithm: "AWS128_CBC",
        messageSubject: "For Connector",
        localProfileId: aws_transfer_profile.local.profile_id,
        mdnResponse: "NONE",
        mdnSigningAlgorithm: "NONE",
        partnerProfileId: aws_transfer_profile.partner.profile_id,
        signingAlgorithm: "NONE",
    },
    url: "http://www.test.com",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Connector("example",
    access_role=aws_iam_role["test"]["arn"],
    as2_config=aws.transfer.ConnectorAs2ConfigArgs(
        compression="DISABLED",
        encryption_algorithm="AWS128_CBC",
        message_subject="For Connector",
        local_profile_id=aws_transfer_profile["local"]["profile_id"],
        mdn_response="NONE",
        mdn_signing_algorithm="NONE",
        partner_profile_id=aws_transfer_profile["partner"]["profile_id"],
        signing_algorithm="NONE",
    ),
    url="http://www.test.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Connector("example", new()
    {
        AccessRole = aws_iam_role.Test.Arn,
        As2Config = new Aws.Transfer.Inputs.ConnectorAs2ConfigArgs
        {
            Compression = "DISABLED",
            EncryptionAlgorithm = "AWS128_CBC",
            MessageSubject = "For Connector",
            LocalProfileId = aws_transfer_profile.Local.Profile_id,
            MdnResponse = "NONE",
            MdnSigningAlgorithm = "NONE",
            PartnerProfileId = aws_transfer_profile.Partner.Profile_id,
            SigningAlgorithm = "NONE",
        },
        Url = "http://www.test.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewConnector(ctx, "example", &transfer.ConnectorArgs{
			AccessRole: pulumi.Any(aws_iam_role.Test.Arn),
			As2Config: &transfer.ConnectorAs2ConfigArgs{
				Compression:         pulumi.String("DISABLED"),
				EncryptionAlgorithm: pulumi.String("AWS128_CBC"),
				MessageSubject:      pulumi.String("For Connector"),
				LocalProfileId:      pulumi.Any(aws_transfer_profile.Local.Profile_id),
				MdnResponse:         pulumi.String("NONE"),
				MdnSigningAlgorithm: pulumi.String("NONE"),
				PartnerProfileId:    pulumi.Any(aws_transfer_profile.Partner.Profile_id),
				SigningAlgorithm:    pulumi.String("NONE"),
			},
			Url: pulumi.String("http://www.test.com"),
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
import com.pulumi.aws.transfer.Connector;
import com.pulumi.aws.transfer.ConnectorArgs;
import com.pulumi.aws.transfer.inputs.ConnectorAs2ConfigArgs;
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
        var example = new Connector("example", ConnectorArgs.builder()        
            .accessRole(aws_iam_role.test().arn())
            .as2Config(ConnectorAs2ConfigArgs.builder()
                .compression("DISABLED")
                .encryptionAlgorithm("AWS128_CBC")
                .messageSubject("For Connector")
                .localProfileId(aws_transfer_profile.local().profile_id())
                .mdnResponse("NONE")
                .mdnSigningAlgorithm("NONE")
                .partnerProfileId(aws_transfer_profile.partner().profile_id())
                .signingAlgorithm("NONE")
                .build())
            .url("http://www.test.com")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Connector
    properties:
      accessRole: ${aws_iam_role.test.arn}
      as2Config:
        compression: DISABLED
        encryptionAlgorithm: AWS128_CBC
        messageSubject: For Connector
        localProfileId: ${aws_transfer_profile.local.profile_id}
        mdnResponse: NONE
        mdnSigningAlgorithm: NONE
        partnerProfileId: ${aws_transfer_profile.partner.profile_id}
        signingAlgorithm: NONE
      url: http://www.test.com
```
{{% /example %}}
{{% example %}}
### SFTP Connector

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Connector("example", {
    accessRole: aws_iam_role.test.arn,
    sftpConfig: {
        trustedHostKeys: ["ssh-rsa AAAAB3NYourKeysHere"],
        userSecretId: aws_secretsmanager_secret.example.id,
    },
    url: "sftp://test.com",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Connector("example",
    access_role=aws_iam_role["test"]["arn"],
    sftp_config=aws.transfer.ConnectorSftpConfigArgs(
        trusted_host_keys=["ssh-rsa AAAAB3NYourKeysHere"],
        user_secret_id=aws_secretsmanager_secret["example"]["id"],
    ),
    url="sftp://test.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Connector("example", new()
    {
        AccessRole = aws_iam_role.Test.Arn,
        SftpConfig = new Aws.Transfer.Inputs.ConnectorSftpConfigArgs
        {
            TrustedHostKeys = new[]
            {
                "ssh-rsa AAAAB3NYourKeysHere",
            },
            UserSecretId = aws_secretsmanager_secret.Example.Id,
        },
        Url = "sftp://test.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewConnector(ctx, "example", &transfer.ConnectorArgs{
			AccessRole: pulumi.Any(aws_iam_role.Test.Arn),
			SftpConfig: &transfer.ConnectorSftpConfigArgs{
				TrustedHostKeys: pulumi.StringArray{
					pulumi.String("ssh-rsa AAAAB3NYourKeysHere"),
				},
				UserSecretId: pulumi.Any(aws_secretsmanager_secret.Example.Id),
			},
			Url: pulumi.String("sftp://test.com"),
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
import com.pulumi.aws.transfer.Connector;
import com.pulumi.aws.transfer.ConnectorArgs;
import com.pulumi.aws.transfer.inputs.ConnectorSftpConfigArgs;
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
        var example = new Connector("example", ConnectorArgs.builder()        
            .accessRole(aws_iam_role.test().arn())
            .sftpConfig(ConnectorSftpConfigArgs.builder()
                .trustedHostKeys("ssh-rsa AAAAB3NYourKeysHere")
                .userSecretId(aws_secretsmanager_secret.example().id())
                .build())
            .url("sftp://test.com")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Connector
    properties:
      accessRole: ${aws_iam_role.test.arn}
      sftpConfig:
        trustedHostKeys:
          - ssh-rsa AAAAB3NYourKeysHere
        userSecretId: ${aws_secretsmanager_secret.example.id}
      url: sftp://test.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transfer AS2 Connector using the `connector_id`. For example:

```sh
 $ pulumi import aws:transfer/connector:Connector example c-4221a88afd5f4362a
```
 